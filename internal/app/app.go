package app

import (
	"github.com/bolanebyla/Go-ToDo/internal/config"
	"github.com/bolanebyla/Go-ToDo/internal/transport"
)

func Run(conf *config.Config) {
	transport.StartApiServer()
}
