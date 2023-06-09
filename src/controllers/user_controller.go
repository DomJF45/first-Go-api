package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/web-service-gin/src/services"
)

type UserController struct {
	Path string
}

func NewUserController() UserController {
	userController := UserController{
		Path: "/users",
	}
	return userController
}

func (c UserController) InitRouter(route *gin.RouterGroup) {
	route.POST(c.Path+"/register", register)
}

func register(c *gin.Context) {
	var user services.User
	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err)
	}

	res, err := services.RegisterUser(user)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}
