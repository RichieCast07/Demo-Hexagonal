package application

import (
	domain "demo/src/products/domain_products"
	"fmt"
)

type UpdateProductUseCase struct {
	db domain.IProduct
}

func NewUpdateProductUseCase(db domain.IProduct) *UpdateProductUseCase {
	return &UpdateProductUseCase{db: db}
}

func (uc *UpdateProductUseCase) Run(product *domain.Product) error {
	fmt.Println("Actualizado exitosamente...")
	return uc.db.Update(product)
}
