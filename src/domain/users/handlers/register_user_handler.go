package users_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	user_commands "github.com/LucasToledoPereira/kyp-api/src/domain/users/commands"
)

func (uh *UsersHandlers) Register(cr user_commands.CreateUserCommandRequest) (user entities.User, err error) {
	user = entities.NewUser(cr)
	err = uh.usersRepository.Create(&user)
	return user, err
}
