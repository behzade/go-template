package command

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type CommandRegistry struct {
	commands map[string]Command
}

func New() *CommandRegistry {
	return &CommandRegistry{
		commands: map[string]Command{},
	}
}

func NewDefault() *CommandRegistry {
	r := CommandRegistry{
		commands: map[string]Command{},
	}

	r.MustRegister(ServerCommand{})
	r.MustRegister(SyncEntityCommand{})
	r.MustRegister(GetAlterCommand{})
	r.MustRegister(ApplyAltersCommand{})

	return &r
}

func (r *CommandRegistry) MustRegister(command Command) {
	err := r.Register(command)
	if err != nil {
		panic(err)
	}
}

func (r *CommandRegistry) Register(command Command) error {
	err := r.validateCommandName(command.Name())
	if err != nil {
		return err
	}

	r.commands[command.Name()] = command
	return nil
}

const namePattern = "^[a-z0-9_]+$"

func (r *CommandRegistry) validateCommandName(name string) error {
	if _, ok := r.commands[name]; ok {
		return fmt.Errorf("duplicate command name")
	}

	if !regexp.MustCompile(namePattern).MatchString(name) {
		return fmt.Errorf("command name must match %v: %v given", namePattern, name)
	}

	return nil
}

func (r *CommandRegistry) Help() string {
	var b strings.Builder
	b.WriteString("Available commands:\n")
	for i := range r.commands {
		b.WriteString(r.commands[i].Name())
		b.WriteString("    ")
		b.WriteString(r.commands[i].Description())
		b.WriteString("\n---\n")
	}

	return b.String()
}

var (
	ErrNoArg      = errors.New("a single argument is required")
	ErrInvalidArg = errors.New("invalid argument given")
)

func (r *CommandRegistry) Run() (string, error) {
	args := os.Args[1:]

	if len(args) < 1 {
		return r.Help(), ErrNoArg
	}

	command, ok := r.commands[args[0]]
	if !ok {
		return r.Help(), ErrInvalidArg
	}

	return command.Run()
}
