package main

import (

	"net/http"
	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}