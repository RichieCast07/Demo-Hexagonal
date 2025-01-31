package controllers

import (
	"net/http"

	application "demo/src/users/application_users"
	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	useCase *application.GetAllUsersUseCase
}

func NewGetAllUsersController(uc *application.GetAllUsersUseCase) *GetAllUsersController {
	return &GetAllUsersController{useCase: uc}
}

func (controller *GetAllUsersController) Execute(c *gin.Context) {
	results, err := controller.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo obtener la lista de usuarios",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": results,
	})
}