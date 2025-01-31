package controllers

import (
	"net/http"
	"strconv"

	application "demo/src/users/application_users"
	"github.com/gin-gonic/gin"
)

type EditUserController struct {
	useCase *application.UpdateUserUseCase
}

func NewEditUserController(uc *application.UpdateUserUseCase) *EditUserController {
	return &EditUserController{useCase: uc}
}

func (controller *EditUserController) Execute(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.useCase.Execute(
		int32(id), input.Name, input.Email, input.Password,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo editar el usuario",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Usuario editado exitosamente",
	})
}