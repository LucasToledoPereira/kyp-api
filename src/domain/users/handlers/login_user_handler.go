package users_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/utils"
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	user_commands "github.com/LucasToledoPereira/kyp-api/src/domain/users/commands"
)

func (uh *UsersHandlers) Login(lr user_commands.LoginUserCommandRequest) (user entities.User, err error) {
	err = uh.usersRepository.GetUserByEmailAndPassword(lr.Email, utils.EncryptPassword(lr.Password), &user)
	return user, err
}
