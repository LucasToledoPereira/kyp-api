package pets_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
)

func (gh *PetsHandlers) GetAll() (pets []entities.Pet, err error) {
	err = gh.petsRepository.GetAll(&pets)
	return pets, err
}
