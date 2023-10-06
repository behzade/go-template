package command

import "github.com/behzade/go-template/internal/controller"

type ServerCommand struct{}

func (ServerCommand) Run() (string, error) {
	err := controller.New().Run()
	return "", err
}

func (ServerCommand) Name() string {
	return "server"
}

func (ServerCommand) Description() string {
	return "Run the router and serve http requests"
}
