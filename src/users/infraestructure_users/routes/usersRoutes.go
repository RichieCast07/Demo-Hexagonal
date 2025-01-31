package routes

import (
	"demo/src/users/infraestructure_users/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, registerUserController *controllers.RegisterUserController, getAllUsersController *controllers.GetAllUsersController, editUserController *controllers.EditUserController, deleteUserController *controllers.DeleteUserController) {
	users := r.Group("/user")
	{
		users.POST("/create", registerUserController.Execute)
		users.GET("/", getAllUsersController.Execute)
		users.PUT("/update/:id", editUserController.Execute)
		users.DELETE("/delete/:id", deleteUserController.Execute)
	}
}
