package application

import "demo/products/domain"

type CreateProductUseCase struct {
	repository domain.IProduct
}

func NewCreateProductUseCase(repository domain.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{repository: repository}
}

func (uc *CreateProductUseCase) Run(name string, price float32) error {
	product := domain.NewProduct(name, price)
	return uc.repository.Save(product)
}
