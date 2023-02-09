package routes

import (
	"github.com/LucasToledoPereira/kyp-api/src/presentation/controllers"
)

func (r *Routes) BuildPetsRoutes(pc *controllers.PetsController) {
	petsRoutes := r.Public.Group("/pets")
	petsRoutes.POST("", pc.Create)
	petsRoutes.GET("", pc.GetAll)
	petsRoutes.PATCH(":id/publish", pc.Publish)
	petsRoutes.PATCH(":id/unpublish", pc.Unpublish)
}
