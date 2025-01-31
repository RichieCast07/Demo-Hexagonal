package routes

import (
	//"demo/src/products/application_products"
	//"demo/src/products/domain_products"
	//"demo/src/products/infraestructure_products"
	"demo/src/products/infraestructure_products/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
//	ps := infraestructure_products.NewCreateProductRepoMySQL()

//	cpUseCase := application_products.NewCreateProductUseCase(ps)
	cpController := controllers.NewCreateProductController(cpUseCase)

//	gapUseCase := application_products.NewGetProductAllUseCase(ps)
	gapController := controllers.NewGetAllProductsController(gapUseCase)

	//epUseCase := application_products.NewUpdateProductUseCase(ps)
	epController := controllers.NewEditProductController(epUseCase)

	//dpUseCase := application_products.NewDeleteProductUseCase(ps)
	dpController := controllers.NewDeleteProductController(dpUseCase)

	products := r.Group("/products")
	{
		products.POST("/create", cpController.Execute)
		products.GET("/", gapController.Execute)
		products.PUT("/update/:id", epController.Execute)
		products.DELETE("/delete/:id", dpController.Execute)
	}
}