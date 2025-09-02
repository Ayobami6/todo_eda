package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}
func (uc *UserController) RegisterRoutes(router *gin.Engine) {
	// root route
	router.GET("/", uc.root)

}

func (uc *UserController) root(c *gin.Context) {
	response := map[string]string{
		"message": "Hello World",
	}
	c.JSON(200, response)
}
