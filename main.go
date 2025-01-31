package main

import (
	"demo/src/core"
	infraestructureproducts "demo/src/products/infraestructure_products"
	"demo/src/users/application_users"
    "demo/src/users/infraestructure_users"
	"demo/src/users/infraestructure_users/controllers"
	"demo/src/users/infraestructure_users/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    connection := core.MySQLConection()
    if connection.Err != "" {
        log.Fatalf("Error en la conexión a la base de datos: %s", connection.Err)
    }

    userRepository := infrastructure_users.NewUserRepository(connection.DB)
    getallusersUseCase := *infraestructureproducts.NewCreateProductRepoMySQL()
    getallusersUseCaseI := *applicationusers.NewGetAllUsersUseCase()
    controllers := controllers.NewRegisterUserController(userRepository,  getallusersUseCaseI)

    r := gin.Default()

    routes.RegisterUserRoutes(r,controllers)

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
