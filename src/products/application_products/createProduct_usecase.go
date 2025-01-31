package application

import (
	domain "demo/src/products/domain_products"
)

type CreateProductUseCase struct {
	db domain.IProduct
}

func NewCreateProductUseCase(db domain.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{db: db}
}

func (uc *CreateProductUseCase) Run(name string, price float32, amount float32) error {
	product := domain.NewProduct(name, price, amount)
	return uc.db.Save(product)
}
