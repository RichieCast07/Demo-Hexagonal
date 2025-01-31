package controllers

import (
	"net/http"

	application "demo/src/products/application_products"
	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	useCase *application.GetProductAllUseCase
}

func NewGetAllProductsController(uc *application.GetProductAllUseCase) *GetAllProductsController {
	return &GetAllProductsController{useCase: uc}
}

func (controller *GetAllProductsController) Execute(c *gin.Context) {
	result, err := controller.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener productos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": result,
	})
}
