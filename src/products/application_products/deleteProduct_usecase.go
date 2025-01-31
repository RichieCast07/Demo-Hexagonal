package application

import (
	domain "demo/src/products/domain_products"
	"fmt"
)

type DeleteProductUseCase struct {
	db domain.IProduct
}

func NewDeleteProductUseCase(db domain.IProduct) *DeleteProductUseCase {
	return &DeleteProductUseCase{db: db}
}

func (uc *DeleteProductUseCase) Run(productId int32) error {
	fmt.Println("Eliminado exitasamente...")
	return uc.db.Delete(productId)
}
