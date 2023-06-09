package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/web-service-gin/src/services"
)

type AlbumController struct {
	Path string
}

func NewAlbumController() AlbumController {
	albumController := AlbumController{
		Path: "/albums",
	}
	return albumController
}

func (c AlbumController) InitRouter(route *gin.RouterGroup) {
	route.GET(c.Path, getAll)
	route.GET(c.Path+"/:id", getById)
	route.POST(c.Path, postAlbum)
	route.DELETE(c.Path+"/:id", deleteAlbum)
	route.PUT(c.Path+"/:id", updateAlbum)
}

func getAll(c *gin.Context) {
	var albs []services.Album
	res, err := services.AllAlbums()
	if err != nil {
		log.Fatal(err)
	}

	albs = res

	c.IndentedJSON(http.StatusOK, albs)
}

func getById(c *gin.Context) {
	var newAlbum services.Album

	n, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err == nil {
		newAlbum, err = services.AlbumById(n)

		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, newAlbum)
	}
}

func postAlbum(c *gin.Context) {
	var newAlbum services.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
	}

	res, err := services.AddAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}

func deleteAlbum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	res, err := services.RemoveAlbum(id)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}

func updateAlbum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var album services.Album

	if err := c.BindJSON(&album); err != nil {
		log.Fatal(err)
	}

	res, err := services.AlterAlbum(album.Price, id)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}
