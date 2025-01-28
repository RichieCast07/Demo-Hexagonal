package application

import (
	"demo/products/domain"
)

type DeleteProductUseCase struct {
	repository domain.IProduct
}

func NewDeleteProductUseCase(repository domain.IProduct) *DeleteProductUseCase {
	return &DeleteProductUseCase{repository: repository}
}

func (uc *DeleteProductUseCase) Run(productId int32) error {
	return uc.repository.Delete(productId)
}
