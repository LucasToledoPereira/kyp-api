package server

import (
	"fmt"

	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/config"
	"github.com/LucasToledoPereira/kyp-api/src/presentation/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(r *routes.Routes) (server *Server) {
	return &Server{
		router: r.Router,
	}
}

func (s *Server) Run() {
	fmt.Printf("collection manager listening on port '%s'", config.Server.Port)
	fmt.Println()

	panic(s.router.Run(":" + config.Server.Port))
}
