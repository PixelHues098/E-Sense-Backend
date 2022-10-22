package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// UpdateAlbum updates the Album with the given details if record found.
func (env Env) UpdateAlbum(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// toBeUpdatedAlbum.
	var toBeUpdatedAlbum Album
	if err := c.BindJSON(&toBeUpdatedAlbum); err != nil {
		e := fmt.Sprintf("invalid JSON body: %v", err)
		log.Println(e)
		makeGinResponse(c, http.StatusBadRequest, e)
		return
	}

	q := `UPDATE artist 
    SET name=$1,title=$2, price=$3
    WHERE id=$4;`
	result, err := env.DB.Exec(q, toBeUpdatedAlbum.Artist, toBeUpdatedAlbum.Title, toBeUpdatedAlbum.Price, toBeUpdatedAlbum.ID)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating artist record with id: %d", err, toBeUpdatedAlbum.ID)
		log.Println(e)
		makeGinResponse(c, http.StatusInternalServerError, e)
		return
	}

	// checking the number of rows affected
	n, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred while checking the returned result from database after updation: %v", err)
		log.Println(e)
		makeGinResponse(c, http.StatusInternalServerError, e)
	}

	// if no record was updated, let us say client has failed
	if n == 0 {
		e := "could not update the record, please try again after sometime"
		log.Println(e)
		makeGinResponse(c, http.StatusInternalServerError, e)
		return
	}

	m := "successfully updated the record"
	log.Println(m)
	makeGinResponse(c, http.StatusOK, m)
}
