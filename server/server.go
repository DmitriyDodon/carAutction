package server

import (
	"fmt"
	"app/config"
	"app/server/controller"

	"github.com/labstack/echo/v4"
)

type server struct {
	port int
	e    *echo.Echo
}

func NewServer(config *config.Config, controller *controller.Controller) *server {
	e := echo.New()

	server := &server{
		port: config.GetPort(),
		e:    e,
	}

	server.initRoutes(controller)

	return server
}

func (s server) Start() error {
	return s.e.Start(fmt.Sprintf(":%d", s.port))
}
