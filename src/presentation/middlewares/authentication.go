package middlewares

import (
	"log"
	"time"

	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/models"
	"github.com/LucasToledoPereira/kyp-api/src/domain/entities"
	"github.com/LucasToledoPereira/kyp-api/src/presentation/controllers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	middleware  *jwt.GinJWTMiddleware
	identityKey string
	controller  *controllers.UsersController
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func InitAuth() (a *Auth) {
	var err error
	a = &Auth{}
	a.identityKey = "id"
	a.middleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     a.identityKey,
		PayloadFunc:     a.PayloadFunc,
		IdentityHandler: a.IdentityHandler,
		Authenticator:   a.Authenticator,
		Unauthorized:    a.Unauthorized,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return a
}

func (a *Auth) SetUserController(uc *controllers.UsersController) {
	a.controller = uc
}

func (a *Auth) PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*entities.User); ok {
		return jwt.MapClaims{
			a.identityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}

func (a *Auth) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &entities.User{
		Email: claims[a.identityKey].(string),
	}
}

func (a *Auth) Authenticator(c *gin.Context) (interface{}, error) {
	return a.controller.Login(c)
}

func (a *Auth) Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, models.NewResultWrapper[any]("user.not.valid", false, nil, message))
}

// @tags Authentication
// Login User
// @Summary 	Return a JWT
// @Schemes
// @Description Return a JWT
// @Accept      json
// @Param		request	body	middlewares.Login	true	" "
// @Produce 	json
// @Success 	200
// @Router 		/auth/login [post]
func (a *Auth) LoginHandler(c *gin.Context) {
	a.middleware.LoginHandler(c)
}

// @tags Authentication
// Refresh Token
// @Summary 	Refresh JWT Token
// @Schemes
// @Description Refresh JWT Token
// @Produce 	json
// @Success 	200
// @Router 		/auth/refresh [get]
func (a *Auth) RefreshHandler(c *gin.Context) {
	a.middleware.RefreshHandler(c)
}

func (a *Auth) MiddlewareFunc() gin.HandlerFunc {
	return a.middleware.MiddlewareFunc()
}
