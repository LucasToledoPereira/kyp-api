package pets_handlers

import (
	"github.com/LucasToledoPereira/kyp-api/src/domain/repositories"
)

type PetsHandlers struct {
	petsRepository repositories.IPetsRepository
}

func NewPetsHandlers(r repositories.IPetsRepository) (petsHandlers *PetsHandlers) {
	return &PetsHandlers{
		petsRepository: r,
	}
}
