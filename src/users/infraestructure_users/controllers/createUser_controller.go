package controllers

import (
	"net/http"

	application "demo/src/users/application_users"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
	useCase *application.CreateUserUseCase
}

func NewRegisterUserController(uc *application.CreateUserUseCase) *RegisterUserController {
	return &RegisterUserController{useCase: uc}
}

func (controller *RegisterUserController) Execute(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.useCase.Execute(input.Name, input.Email, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al registrar usuario",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
	})
}