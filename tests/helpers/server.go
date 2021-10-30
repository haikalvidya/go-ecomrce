package helpers

import (
	"go-ecomrce/config"
	"go-ecomrce/server"
	"github.com/labstack/echo/v4"
)

func NewServer() *server.Server {
	s := &server.Server{
		Echo: echo.New(),
		DB:   Init(),
		Config: config.NewConfig(),
	}

	return s
}
