package routes

import (
    "github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, productController *ProductController) {
	r := r.Group("/r/product")
	{
		r.POST("/create", func(c *gin.Context) {
			c.String(http.StatusOK, "Post, Product!")
		})

		r.GET("/allproduct", func(c *gin.Context) {
			c.String(http.StatusOK, "Get, Product!")
		})
				
		r.PUT("/update/:id", func(c *gin.Context) {
			c.String(http.StatusOK, "Put, Product!")
		})
				
		r.DELETE("/delete/:id", func(c *gin.Context) {
			c.String(http.StatusOK, "Delete, Product!")
		})
	}
}
