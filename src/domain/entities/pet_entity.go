package entities

import (
	"time"

	pets_commands "github.com/LucasToledoPereira/kyp-api/src/domain/pets/commands"
	pets_enums "github.com/LucasToledoPereira/kyp-api/src/domain/pets/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pet struct {
	ID          uuid.UUID            `gorm:"primaryKey; index; unique; type:uuid;"`
	Name        string               `gorm:"not null"`
	Description string               `gorm:"not null; default:null"`
	Status      pets_enums.PetStatus `gorm:"default:DRAFT; not null; type:pet_status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func NewPet(c pets_commands.CreatePetCommandRequest) (pet Pet) {
	return Pet{
		ID:          uuid.New(),
		Name:        c.Name,
		Description: c.Description,
	}
}
