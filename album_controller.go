package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Path string
}

func (c Controller) initRouter(route *gin.RouterGroup) {
	route.GET(c.Path, getAll)
	route.GET(c.Path+"/:id", getById)
	route.POST(c.Path, postAlbum)
	route.DELETE(c.Path+"/:id", deleteAlbum)
	route.PUT(c.Path+"/:id", updateAlbum)
}

func newAlbumController() Controller {
	albumController := Controller{
		Path: "/albums",
	}
	return albumController
}

func getAll(c *gin.Context) {
	var albs []Album
	res, err := allAlbums()
	if err != nil {
		log.Fatal(err)
	}

	albs = res

	c.IndentedJSON(http.StatusOK, albs)
}

func getById(c *gin.Context) {
	var newAlbum Album

	n, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err == nil {
		newAlbum, err = albumById(n)

		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, newAlbum)
	}
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
	}

	res, err := addAlbum(newAlbum)
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

	res, err := removeAlbum(id)
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

	var album Album

	if err := c.BindJSON(&album); err != nil {
		log.Fatal(err)
	}

	res, err := alterAlbum(album.Price, id)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}
