package main

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	initRouter(route *gin.RouterGroup)
}
