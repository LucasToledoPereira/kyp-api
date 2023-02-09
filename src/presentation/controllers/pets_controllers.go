package controllers

import (
	"net/http"
	"strings"

	pets_commands "github.com/LucasToledoPereira/kyp-api/src/domain/pets/commands"
	pets_handlers "github.com/LucasToledoPereira/kyp-api/src/domain/pets/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/models"
)

type PetsController struct {
	handlers *pets_handlers.PetsHandlers
}

func NewPetsController(ph *pets_handlers.PetsHandlers) (petsController *PetsController) {
	return &PetsController{
		handlers: ph,
	}
}

// Create a Pet godoc
// @Summary 	Create a new pet
// @Schemes
// @Description Used to create a new pet info
// @Accept      json
// @Param		request	body	pets_commands.CreatePetCommandRequest	true	" "
// @Produce 	json
// @Success 	201 {object} models.ResultWrapper[entities.Pet]
// @Router 		/pets [post]
func (pc *PetsController) Create(ctx *gin.Context) {
	var commandRequest pets_commands.CreatePetCommandRequest
	if err := ctx.ShouldBindJSON(&commandRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("", false, []string{err.Error()}, nil))
		return
	}

	pet, err := pc.handlers.Create(commandRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("error.create.pet", false, strings.Split(err.Error(), ";"), nil))
		return
	}

	ctx.JSON(http.StatusCreated, models.NewResultWrapper("success.create.pet", true, nil, pet))
}

// Get all pets
// @Summary 	Return all available pets
// @Schemes
// @Description Return all available pets
// @Produce 	json
// @Success 	200 {object} models.ResultWrapper[[]entities.Pet]
// @Router 		/pets [get]
func (pc *PetsController) GetAll(ctx *gin.Context) {
	pets, err := pc.handlers.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("error.list.pets", false, strings.Split(err.Error(), ";"), nil))
		return
	}

	ctx.JSON(http.StatusCreated, models.NewResultWrapper("success.list.pets", true, nil, pets))
}

// Publish pet
// @Summary 	Change pet status to PUBLISHED
// @Schemes
// @Description Change pet status to PUBLISHED
// @Param       id   path      string  true  "Pet ID"
// @Produce 	json
// @Success 	200 {object} models.ResultWrapper[entities.Pet]
// @Router 		/pets/{id}/publish [patch]
func (pc *PetsController) Publish(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	uid, _ := uuid.Parse(id)
	pet, err := pc.handlers.Publish(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("error.publishing.pet", false, strings.Split(err.Error(), ";"), nil))
		return
	}
	ctx.JSON(http.StatusCreated, models.NewResultWrapper("success.publishing.pet", true, nil, pet))
}

// Unpublish pet
// @Summary 	Change pet status to UNPUBLISHED
// @Schemes
// @Description Change pet status to PUBLISHED
// @Param       id   path      string  true  "Pet ID"
// @Produce 	json
// @Success 	200 {object} models.ResultWrapper[entities.Pet]
// @Router 		/pets/{id}/unpublish [patch]
func (pc *PetsController) Unpublish(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	uid, _ := uuid.Parse(id)
	pet, err := pc.handlers.Unpublish(uid)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("error.unpublishing.pet", false, strings.Split(err.Error(), ";"), nil))
		return
	}
	ctx.JSON(http.StatusCreated, models.NewResultWrapper("success.unpublishing.pet", true, nil, pet))
}
