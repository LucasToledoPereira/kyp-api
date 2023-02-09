package repositories

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	"github.com/google/uuid"
)

type IUsersRepository interface {
	GetById(id uuid.UUID, usr *entities.User) (err error)
	GetAll(usrs *[]entities.User) (err error)
	Create(usr *entities.User) (err error)
	GetUserByEmailAndPassword(email string, password string, usr *entities.User) (err error)
}
