package main

import (
	"fmt"
	"os"
)

const (
	updateEntityArg = "update_entity"
	getAltersArg    = "get_alters"
	applyAltersArg  = "apply_alters"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		help()
	}

	arg := args[0]
	var action func() (string, error)
	switch arg {
	case updateEntityArg:
		action = updateEntity
	case getAltersArg:
		action = getAlters
	case applyAltersArg:
		action = applyAlters
	default:
		help()
	}

	msg,err := action()
	println(msg)
	
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func help() {
	fmt.Printf(
		`Available commands:
	%v  %v
	%v  %v
	%v  %v
`,
		updateEntityArg, "syncs generated entity with sql schema/query",
		getAltersArg, "generates alters by comparing db and sql schema",
		applyAltersArg, "syncs db with sql schema",
	)
	os.Exit(1)
}

func updateEntity() (string, error) {
	err := renderTemplate()
	if err != nil {
		return "", nil
	}

	msg, err := syncEntityWithSchema()
	return string(msg), err
}

func getAlters() (string, error) {
	err := renderTemplate()
	if err != nil {
		return "", nil
	}

	msg, err := diffDBWithSchema()
	return string(msg), err
}

func applyAlters() (string, error) {
	err := renderTemplate()
	if err != nil {
		return "", nil
	}

	msg, err := syncDBWithSchema()
	return string(msg), err
}
