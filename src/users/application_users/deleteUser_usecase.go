package applicationusers

import (
	domain "demo/src/users/domain_users"
)

type DeleteUserUseCase struct {
	repository domain.UserInterface
}

func NewDeleteUserUseCase(repository domain.UserInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{repository: repository}
}

func (uc *DeleteUserUseCase) Execute(id int32) error {
	return uc.repository.DeleteUser(id)
}
