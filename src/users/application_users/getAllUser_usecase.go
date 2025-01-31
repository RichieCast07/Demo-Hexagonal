package applicationusers

import (
	domain "demo/src/users/domain_users"
)

type GetAllUsersUseCase struct {
	repository domain.UserInterface
}

func NewGetAllUsersUseCase(repository domain.UserInterface) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{repository: repository}
}

func (uc *GetAllUsersUseCase) Execute() ([]*domain.User, error) {
	return uc.repository.GetUsers()
}
