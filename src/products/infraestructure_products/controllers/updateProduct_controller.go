package controllers

import (
	"net/http"
	"strconv"

	application "demo/src/products/application_products"
	domain "demo/src/products/domain_products"

	"github.com/gin-gonic/gin"
)

type EditProductController struct {
	useCase *application.UpdateProductUseCase
}

func NewEditProductController(uc *application.UpdateProductUseCase) *EditProductController {
	return &EditProductController{useCase: uc}
}

func (controller *EditProductController) Execute(c *gin.Context) {
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

	product := &domain.Product{
		ID:     int32(id),
		Name:   input.Name,
		Price:  input.Price,
		Amount: input.Amount,
	}

	if err := controller.useCase.Run(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo editar el producto",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Producto editado exitosamente",
	})
}
