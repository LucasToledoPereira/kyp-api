package users_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/repositories"
)

type UsersHandlers struct {
	usersRepository repositories.IUsersRepository
}

func NewUsersHandlers(r repositories.IUsersRepository) (usersHandlers *UsersHandlers) {
	return &UsersHandlers{
		usersRepository: r,
	}
}
