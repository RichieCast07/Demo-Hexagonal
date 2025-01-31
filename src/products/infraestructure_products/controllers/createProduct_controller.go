package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	application "demo/src/products/application_products"
)

type CreateProductController struct{
	useCase *application.CreateProductUseCase
}

func NewCreateProductController(uc *application.CreateProductUseCase) *CreateProductController{
	return &CreateProductController{useCase: uc}
}

func (controller *CreateProductController) Execute(c *gin.Context){
	var input struct {
		Name   string  `json:"name"`
		Price  float32 `json:"price"`
		Amount float32 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.useCase.Run(input.Name, input.Price, input.Amount); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al registrar producto",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Producto registrado",
	})
}
