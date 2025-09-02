package api

import (
	"github.com/change_me/go_starter_rest/internal/controller"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}
func (s *APIServer) Start() error {
	// set up gin
	router := gin.Default()
	// set up routes
	// set up user controller
	userController := controller.NewUserController()
	userController.RegisterRoutes(router)

	return router.Run(s.addr)

}
