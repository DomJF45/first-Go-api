package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/web-service-gin/src/interfaces"
	"github.com/web-service-gin/src/services"
	"github.com/web-service-gin/src/util"
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
	route.POST(c.Path+"/login", login)
	route.GET(c.Path+"/me", getUser)
}

func register(c *gin.Context) {
	var user interfaces.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, err := services.RegisterUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func login(c *gin.Context) {
	var user *interfaces.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// c.SetCookie("gin_cookie", uuid.NewString(), 3600, "/", "localhost", false, true)

	util.SessionManager.Put(c.Request.Context(), "message", "Hello from a session!")

	c.IndentedJSON(http.StatusOK, res)
}

func getUser(c *gin.Context) {
	var user *interfaces.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	msg := util.SessionManager.GetString(c.Request.Context(), "message")
	c.IndentedJSON(http.StatusOK, msg)
}
