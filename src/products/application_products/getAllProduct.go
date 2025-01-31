package application

import (
	domain "demo/src/products/domain_products"
	"fmt"
)

type GetProductAllUseCase struct {
	db domain.IProduct
}

func NewGetProductAllUseCase(db domain.IProduct) *GetProductAllUseCase {
	return &GetProductAllUseCase{db: db}
}

func (uc *GetProductAllUseCase) Run() ([]*domain.Product, error) {
	fmt.Println("Obtenido exitosamente...")
	return uc.db.GetAll()
}
