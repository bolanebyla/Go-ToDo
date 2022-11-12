package main

import (
	"github.com/bolanebyla/Go-ToDo/internal/app"
	"github.com/bolanebyla/Go-ToDo/internal/config"
)

func main() {
	conf := config.NewConfig()

	app.Run(conf)
}
