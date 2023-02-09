package controllers

import (
	pets_handlers "github.com/LucasToledoPereira/kyp-api/src/domain/pets/handlers"
)

type Controllers struct {
	PetsHandlers *pets_handlers.PetsHandlers
}

func NewControllers(ph *pets_handlers.PetsHandlers) (controllers *Controllers) {
	return &Controllers{
		PetsHandlers: ph,
	}
}
