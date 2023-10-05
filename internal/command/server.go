package command

import "github.com/behzade/go-template/internal/controller"

type ServeCommand struct{}

func (ServeCommand) Run() (string, error) {
	err := controller.New().Run()
	return "", err
}

func (ServeCommand) Name() string {
	return "serve"
}

func (ServeCommand) Description() string {
	return "Run the router and serve http requests"
}
