package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbums)
    router.DELETE("/albums/:id", deleteAlbumByID)
	

	router.Run("localhost:8080")
}


