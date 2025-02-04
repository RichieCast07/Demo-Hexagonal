package main

import (
	productRoutes "demo/src/products/infraestructure_products/routes"
	usersRoutes "demo/src/users/infraestructure_users/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	productRoutes.RegisterProductRoutes(r)
	usersRoutes.RegisterUserRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
