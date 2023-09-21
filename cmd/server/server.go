package main

import (
	"fmt"
	"log"

	"github.com/behzade/fiber-template/internal/controller"
	"github.com/behzade/fiber-template/internal/service"
)

func main() {
	cfg, err := service.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cfg: %v\n", cfg)
	app := controller.New()
	app.Run()
}
