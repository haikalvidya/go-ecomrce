package main

import (
	"fmt"
	application "go-ecomrce"
	"go-ecomrce/config"

	"go-ecomrce/docs"
)

// @title Go Ecomrce
// @version 1.0
// @description This is a simple version of e-commerce using Echo app.

// @contact.name Vidya Haikal
// @contact.email haikalvidya@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
