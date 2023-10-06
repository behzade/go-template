package main

import (
	"os"

	"github.com/behzade/go-template/internal/command"
)

func main() {
	r := command.NewDefault()
	res, err := r.Run()
	statusCode := 0
	if err != nil {
		statusCode = 1
		println(err.Error(), "\n---\n")
	}
	println(res)

	os.Exit(statusCode)
}
