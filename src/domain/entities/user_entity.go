package entities

import (
	"time"

	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/utils"
	user_commands "github.com/LucasToledoPereira/kyp-api/src/domain/users/commands"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Name      string    `gorm:"not null; default:null"`
	Email     string    `gorm:"not null; default:null; unique"`
	Password  string    `gorm:"not null; default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewUser(c user_commands.CreateUserCommandRequest) (user User) {
	return User{
		ID:       uuid.New(),
		Name:     c.Name,
		Email:    c.Email,
		Password: utils.EncryptPassword(c.Password),
	}
}

func (usr *User) CheckPassword(s string) bool {
	if usr.Password != utils.EncryptPassword(s) {
		return false
	}

	return true
}
