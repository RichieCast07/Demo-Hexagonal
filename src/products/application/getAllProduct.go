package application

import "demo/products/domain"

type GetProductAllUseCase struct {
	repository domain.IProduct
}

func NewGetProductAllUseCase(repository domain.IProduct) *GetProductAllUseCase {
	return &GetProductAllUseCase{repository: repository}
}

func (uc *GetProductAllUseCase) Run() ([]*domain.Product, error) {
	return uc.repository.GetAll()
}
