package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router *gin.Engine
}

func (c *Controller) initRouter() {
	c.Router.GET("/albums", getAll)
	c.Router.GET("/albums/:id", getById)
	c.Router.POST("/albums", postAlbum)
	c.Router.DELETE("/albums/:id", deleteAlbum)
	c.Router.PUT("/albums/:id", updateAlbum)
}

func (c *Controller) run() {
	c.Router.Run("localhost:8080")
}

func newController() Controller {
	albumController := Controller{
		Router: gin.Default(),
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
