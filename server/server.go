package server

import (
	"log"

	"github.com/GabrielVictorP/golang-restAPI.git/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("Server running on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
