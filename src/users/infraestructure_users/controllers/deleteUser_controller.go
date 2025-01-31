package controllers

import (
	"net/http"
	"strconv"

	application "demo/src/users/application_users"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	useCase *application.DeleteUserUseCase
}

func NewDeleteUserController(uc *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{useCase: uc}
}

func (controller *DeleteUserController) Execute(c *gin.Context) {
	idStr, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener parámetros",
		})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parámetro no válido",
		})
		return
	}

	if err := controller.useCase.Execute(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo eliminar al usuario",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Usuario eliminado exitosamente",
	})
}
