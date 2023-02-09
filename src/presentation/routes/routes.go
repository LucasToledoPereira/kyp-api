package routes

import (
	"fmt"
	"net/http"

	"github.com/LucasToledoPereira/kyp-api/docs"
	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/config"
	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/models"
	"github.com/LucasToledoPereira/kyp-api/src/presentation/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routes struct {
	Router  *gin.Engine
	Version *gin.RouterGroup
	Public  *gin.RouterGroup
	Private *gin.RouterGroup
	Admin   *gin.RouterGroup
	Auth    *middlewares.Auth
}

func NewRoutes() (routes *Routes) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	router := gin.Default()
	version := router.Group("v1/")
	auth := middlewares.InitAuth()

	routes = &Routes{
		Router:  router,
		Version: version,
		Public:  version,
		Private: version.Group("private/"),
		Admin:   version.Group("admin/"),
		Auth:    auth,
	}

	routes.Private.Use(auth.MiddlewareFunc())

	routes.buildAuthRoutes()
	routes.buildSwaggerRoute()
	routes.buildHealthRoute()

	return routes
}

func (r *Routes) buildSwaggerRoute() {
	docs.SwaggerInfo.BasePath = "/v1"
	r.Public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (r *Routes) buildHealthRoute() {
	r.Private.GET("/health", healthChecker)
}

func (r *Routes) buildAuthRoutes() {
	auth := r.Public.Group("/auth")
	auth.POST("/login", r.Auth.LoginHandler)
	auth.GET("/refresh", r.Auth.RefreshHandler)
}

// Health Checker godoc
// @Summary Check the api health
// @param Authorization header string false "Should be Bearer ..."
// @security Authorization
// @Schemes
// @Description Check if the api is working properly
// @Produce json
// @Success 200 {object} models.ResultWrapper[any]
// @Router /private/health [get]
func healthChecker(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.NewResultWrapper[any]("health.checker.success", true, nil, nil))
}
