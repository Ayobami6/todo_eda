package api

import (
	"github.com/Ayobami6/todo_read/internal/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr     string
	dbClient *mongo.Client
}

func NewAPIServer(addr string, dbClient *mongo.Client) *APIServer {
	return &APIServer{
		addr:     addr,
		dbClient: dbClient,
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
