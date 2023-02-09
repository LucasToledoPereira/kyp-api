package controllers

import (
	"net/http"
	"strings"

	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/models"
	user_commands "github.com/LucasToledoPereira/kyp-api/src/domain/users/commands"
	users_handlers "github.com/LucasToledoPereira/kyp-api/src/domain/users/handlers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	handlers *users_handlers.UsersHandlers
}

func NewUsersController(uh *users_handlers.UsersHandlers) (usersController *UsersController) {
	return &UsersController{
		handlers: uh,
	}
}

// @tag Users
// Register a new user
// @Summary 	Register a new user
// @Schemes
// @Description Register a new user
// @Accept      json
// @Param		request	body	user_commands.CreateUserCommandRequest	true	" "
// @Produce 	json
// @Success 	201 {object} models.ResultWrapper[entities.User]
// @Router 		/register [post]
func (uc *UsersController) Register(ctx *gin.Context) {
	var commandRequest user_commands.CreateUserCommandRequest
	if err := ctx.ShouldBindJSON(&commandRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("", false, []string{err.Error()}, nil))
		return
	}

	pet, err := uc.handlers.Register(commandRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("error.create.user", false, strings.Split(err.Error(), ";"), nil))
		return
	}

	ctx.JSON(http.StatusCreated, models.NewResultWrapper("success.create.user", true, nil, pet))
}

func (uc *UsersController) Login(ctx *gin.Context) (interface{}, error) {
	var command user_commands.LoginUserCommandRequest

	if err := ctx.ShouldBind(&command); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	usr, err := uc.handlers.Login(command)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
