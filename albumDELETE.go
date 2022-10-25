package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that corresponding message.
func RemoveIndex(s []album, index int) []album {
	return append(s[:index], s[index+1:]...)
}
func deleteAlbumByID(c *gin.Context) {

	id := c.Param("id")

	for i, album := range albums {

		if album.ID == id {
			albums = RemoveIndex(albums, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
