package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateAlbum updates the Album with the given details if record found.
func updateAlbums(c *gin.Context) {

	id := c.Param("id")

	var updateData album

	if error := c.BindJSON(&updateData); error != nil {
		return
	}
	for i, album := range albums {

		if album.ID == id {
			albums[i] = updateData
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "album not found"})
}
