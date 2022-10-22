package main

import (

	"github.com/gin-gonic/gin"
)

func main() {

    // env := new(Env)

    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    // router.POST("/albums", env.updateAlbums)
    // router.DELETE("/albums/:id", env.deleteAlbumByID)

    router.Run("localhost:8080")
}