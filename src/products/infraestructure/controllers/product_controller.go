package controllers

import (
	"demo/products/application"
)

type ProductController struct {
	CreateProductUseCase *application.CreateProductUseCase
}

func NewProductController(createproductUseCase *application.CreateProductUseCase) *ProductController {
	return &ProductController{CreateProductUseCase: createproductUseCase}
}
