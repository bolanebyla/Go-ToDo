package app

import (
	"github.com/bolanebyla/Go-ToDo/internal/config"
	"github.com/bolanebyla/Go-ToDo/internal/services"
	"github.com/bolanebyla/Go-ToDo/transport"
)

func Run(conf *config.Config) {
	services.SayHello(conf.Name)
	transport.Start()
}
