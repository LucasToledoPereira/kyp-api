package pets_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	pets_commands "github.com/LucasToledoPereira/kyp-api/src/domain/pets/commands"
)

func (gh *PetsHandlers) Create(cr pets_commands.CreatePetCommandRequest) (pet entities.Pet, err error) {
	pet = entities.NewPet(cr)
	err = gh.petsRepository.Create(&pet)
	return pet, err
}
