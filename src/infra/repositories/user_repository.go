package repositories

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	"github.com/LucasToledoPereira/kyp-api/src/infra/data"
	"github.com/google/uuid"
)

type UsersRepository struct {
	data *data.Data
}

func NewUsersRepository(d *data.Data) (usersRepository *UsersRepository) {
	return &UsersRepository{
		data: d,
	}
}

func (ur UsersRepository) GetById(bi uuid.UUID, usr *entities.User) (err error) {
	err = ur.data.DB.Where(entities.User{ID: bi}).First(&usr).Error
	return err
}

func (ur UsersRepository) Create(usr *entities.User) (err error) {
	err = ur.data.DB.Create(&usr).Error
	return err
}

func (ur UsersRepository) GetAll(usrs *[]entities.User) (err error) {
	err = ur.data.DB.Find(&usrs).Error
	return err
}

func (ur UsersRepository) GetUserByEmailAndPassword(email string, password string, usr *entities.User) (err error) {
	err = ur.data.DB.Where(entities.User{
		Email:    email,
		Password: password,
	}).First(usr).Error
	return err
}
