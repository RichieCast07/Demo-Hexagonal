package controllers

import (
	"net/http"
	"strconv"

	application "demo/src/products/application_products"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase *application.DeleteProductUseCase
}

func NewDeleteProductController(uc *application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{useCase: uc}
}

func (controller *DeleteProductController) Execute(c *gin.Context) {
	strID, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Falta el parámetro ID"})
		return
	}

	id, err := strconv.ParseInt(strID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro id no válido"})
		return
	}

	id32 := int32(id)
	if err := controller.useCase.Run(id32); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al ejecutar método"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Producto eliminado exitosamente",
	})
}
