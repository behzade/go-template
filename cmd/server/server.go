package main

import "github.com/behzade/fiber-template/internal/controller"

func main() {
	app := controller.New()
	app.Run()
}
