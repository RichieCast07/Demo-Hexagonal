package application

import "demo/products/domain"

type UpdateProductUseCase struct {
	repository domain.IProduct
}

func NewUpdateProductUseCase(repository domain.IProduct) *UpdateProductUseCase {
	return &UpdateProductUseCase{repository: repository}
}

func (uc *UpdateProductUseCase) Run(product *domain.Product) error {
	return uc.repository.Update(product)
}
