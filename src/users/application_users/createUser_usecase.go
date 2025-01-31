package applicationusers

import (
	domain "demo/src/users/domain_users"
)

type CreateUserUseCase struct {
	repository domain.UserInterface
}

func NewCreateUserUseCase(repository domain.UserInterface) *CreateUserUseCase {
	return &CreateUserUseCase{repository: repository}
}

func (uc *CreateUserUseCase) Execute(name string, email string, password string) error {
	user := domain.NewUser(name, email, password)
	return uc.repository.Create(user)
}