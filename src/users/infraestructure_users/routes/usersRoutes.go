package routes

import (
	"net/http"
	"strconv"
	"time"

	application_users "demo/src/users/application_users"
	domain "demo/src/users/domain_users"
	infraestructure_users "demo/src/users/infraestructure_users"
	controllers "demo/src/users/infraestructure_users/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRepo := infraestructure_users.NewUserRepoMySQL()
	getAllUsersUseCase := application_users.NewGetAllUsersUseCase(userRepo)
	getAllUsersController := controllers.NewGetAllUsersController(getAllUsersUseCase)

	users := router.Group("/users")
	{
		users.POST("/", func(c *gin.Context) {
			var user domain.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := userRepo.SaveUserFunction(user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el usuario"})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado con exito"})
		})

		users.GET("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
				return
			}
			userData, err := userRepo.GetUserFunction(int32(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuario"})
				return
			}
			c.JSON(http.StatusOK, userData)
		})

		users.PUT("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
				return
			}
			var user domain.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			user.ID = int32(id)
			if err := userRepo.EditUserFunction(user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado con exito"})
		})

		users.DELETE("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
				return
			}
			if err := userRepo.DeleteUser(int32(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado con exito"})
		})

		users.GET("/sync", func(c *gin.Context) {
			getAllUsersController.Execute(c)
		})

		users.GET("/sync/recent", func(c *gin.Context) {
			desdeParam := c.Query("desde")
			desde, err := strconv.Atoi(desdeParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro 'desde' inválido"})
				return
			}
			allUsers, err := userRepo.GetAllUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al sincronizar usuarios"})
				return
			}
			var recentUsers []domain.User
			for _, u := range allUsers {
				if int(u.ID) > desde {
					recentUsers = append(recentUsers, *u)
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"usuarios_recientes": recentUsers,
				"timestamp":          time.Now().Unix(),
			})
		})
		
		users.GET("/longsync", func(c *gin.Context) {
			getAllUsersController.LongPollingExecute(c)
		})
	}
}
