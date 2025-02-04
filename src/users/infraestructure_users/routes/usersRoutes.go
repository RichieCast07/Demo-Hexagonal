package routes

import (
	"net/http"
	"strconv"

	domain "demo/src/users/domain_users"
	infraestructure_users "demo/src/users/infraestructure_users"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRepo := infraestructure_users.NewUserRepoMySQL()

	users := router.Group("/users")
	{
		users.POST("/create", func(c *gin.Context) {
			var user domain.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := userRepo.SaveUserFunction(user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el usuario"})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado con éxito"})
		})

		users.GET("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			userData, err := userRepo.GetUserFunction(int32(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuario"})
				return
			}
			c.JSON(http.StatusOK, userData)
		})

		users.PUT("/update/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
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
			c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado con éxito"})
		})

		users.DELETE("/delete/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			if err := userRepo.DeleteUserFunction(int32(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado con éxito"})
		})
	}
}
