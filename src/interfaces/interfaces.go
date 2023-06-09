package interfaces

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	InitRouter(route *gin.RouterGroup)
}

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registeredAt"`
	LastLogin    time.Time `json:"lastLogin"`
}
