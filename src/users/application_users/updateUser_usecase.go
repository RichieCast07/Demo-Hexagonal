package applicationusers

import (
	domain "demo/src/users/domain_users"
)

type UpdateUserUseCase struct {
	repository domain.UserInterface
}

func NewUpdateUserUseCase(repository domain.UserInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{repository: repository}
}

func (uc *UpdateUserUseCase) Execute(id int32, name string, email string, password string) error {
	user := domain.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	return uc.repository.UpdateUser(&user)
}
