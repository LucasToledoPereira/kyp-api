package pets_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	"github.com/google/uuid"
)

func (gh *PetsHandlers) Publish(id uuid.UUID) (pet entities.Pet, err error) {
	err = gh.petsRepository.GetById(id, &pet)
	if err != nil {
		return entities.Pet{}, err
	}
	err = gh.petsRepository.Publish(id, &pet)
	return pet, err
}
