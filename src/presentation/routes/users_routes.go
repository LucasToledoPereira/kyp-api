package routes

import (
	"github.com/LucasToledoPereira/kyp-api/src/presentation/controllers"
)

func (r *Routes) BuildUsersRoutes(uc *controllers.UsersController) {
	r.Public.POST("register", uc.Register)
	//usersRoutes := r.Public.Group("/users")
}
