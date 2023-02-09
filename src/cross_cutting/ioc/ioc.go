package ioc

import (
	pets_handlers "github.com/LucasToledoPereira/kyp-api/src/domain/pets/handlers"
	users_handlers "github.com/LucasToledoPereira/kyp-api/src/domain/users/handlers"

	"github.com/LucasToledoPereira/kyp-api/src/infra/data"
	"github.com/LucasToledoPereira/kyp-api/src/infra/repositories"
	"github.com/LucasToledoPereira/kyp-api/src/presentation/controllers"
	"github.com/LucasToledoPereira/kyp-api/src/presentation/routes"
	"github.com/LucasToledoPereira/kyp-api/src/presentation/server"
)

func Bootstrap() (s *server.Server, err error) {
	data, err := data.NewData()

	if err != nil {
		return &server.Server{}, err
	}

	uc := getUsersController(data)
	routes := routes.NewRoutes()
	routes.Auth.SetUserController(uc)
	routes.BuildPetsRoutes(getPetsController(data))
	routes.BuildUsersRoutes(uc)

	return server.NewServer(routes), nil
}

func getPetsController(data *data.Data) *controllers.PetsController {
	petsRepository := repositories.NewPetsRepository(data)
	petsHandlers := pets_handlers.NewPetsHandlers(petsRepository)
	return controllers.NewPetsController(petsHandlers)
}

func getUsersController(data *data.Data) *controllers.UsersController {
	usersRepository := repositories.NewUsersRepository(data)
	usersHandlers := users_handlers.NewUsersHandlers(usersRepository)
	return controllers.NewUsersController(usersHandlers)
}
