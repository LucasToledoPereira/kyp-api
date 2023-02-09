package repositories

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	"github.com/google/uuid"
)

type IPetsRepository interface {
	GetById(id uuid.UUID, pet *entities.Pet) (err error)
	GetAll(pets *[]entities.Pet) (err error)
	Create(p *entities.Pet) (err error)
	Publish(id uuid.UUID, p *entities.Pet) (err error)
	Unpublish(id uuid.UUID, p *entities.Pet) (err error)
}
