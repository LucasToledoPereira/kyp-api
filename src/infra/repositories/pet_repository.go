package repositories

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	pets_enums "github.com/LucasToledoPereira/kyp-api/src/domain/pets/enums"
	"github.com/LucasToledoPereira/kyp-api/src/infra/data"
	"github.com/google/uuid"
)

type PetsRepository struct {
	data *data.Data
}

func NewPetsRepository(d *data.Data) (petsRepository *PetsRepository) {
	return &PetsRepository{
		data: d,
	}
}

func (gr PetsRepository) GetById(bi uuid.UUID, pet *entities.Pet) (err error) {
	err = gr.data.DB.Where(entities.Pet{ID: bi}).First(&pet).Error
	return err
}

func (gr PetsRepository) Create(pet *entities.Pet) (err error) {
	err = gr.data.DB.Create(&pet).Error
	return err
}

func (gr PetsRepository) GetAll(pets *[]entities.Pet) (err error) {
	err = gr.data.DB.Find(&pets).Error
	return err
}

func (gr PetsRepository) Publish(id uuid.UUID, pet *entities.Pet) (err error) {
	err = gr.data.DB.Model(&pet).Update("status", pets_enums.PUBLISHED).Error
	return err
}

func (gr PetsRepository) Unpublish(id uuid.UUID, pet *entities.Pet) (err error) {
	err = gr.data.DB.Model(&pet).Update("status", pets_enums.UNPUBLISHED).Error
	return err
}
