package routes

import (
	"fmt"
	"go-ecomrce/server"
	"go-ecomrce/server/handlers"
	"go-ecomrce/services/token"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *server.Server) {
	productHandler := handlers.NewProductHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/api/v1/login", authHandler.Login)
	server.Echo.POST("/api/v1/register", registerHandler.Register)
	server.Echo.POST("/api/v1/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)

	server.Echo.Static("/uploads", "uploads")

	r := server.Echo.Group("/api/v1")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetAProduct)
	r.POST("/products", productHandler.CreateProduct)
	r.POST("/products/:id/image", productHandler.UploadImageProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
}
