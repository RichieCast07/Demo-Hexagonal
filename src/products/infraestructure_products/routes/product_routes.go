package routes

import (
	"net/http"
	"strconv"

	domain "demo/src/products/domain_products"
	infraestructure_products "demo/src/products/infraestructure_products"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine) {

	productRepo := infraestructure_products.NewCreateProductRepoMySQL()

	products := router.Group("/products")
	{
		products.POST("/", func(c *gin.Context) {
			var product domain.Product
			if err := c.ShouldBindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := productRepo.Save(&product); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar el producto"})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "Producto creado con éxito"})
		})

		products.GET("/", func(c *gin.Context) {
			productsList, err := productRepo.GetAll()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener productos"})
				return
			}
			c.JSON(http.StatusOK, productsList)
		})

		products.GET("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			product, err := productRepo.GetByID(int32(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el producto"})
				return
			}
			c.JSON(http.StatusOK, product)
		})

		products.PUT("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			var product domain.Product
			if err := c.ShouldBindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := productRepo.Edit(product.Name, product.Price, product.Amount, int32(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar producto"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado con éxito"})
		})

		products.DELETE("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
				return
			}

			if err := productRepo.Delete(int32(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar producto"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado con éxito"})
		})
	}
}
