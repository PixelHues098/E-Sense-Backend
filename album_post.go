package main

import (
    "log"
    "net/http"
    
    "github.com/gin-gonic/gin"

    _ "github.com/lib/pq"
)

// PostAlbums adds an album from JSON received in the request body.
func (env Env) PostAlbum(c *gin.Context) {
    // Call BindJSON to bind the received JSON to
    // newAlbum.
    var newAlbum Album
    if err := c.BindJSON(&newAlbum); err != nil {
        log.Printf("invalid JSON body: %v", err)
        makeGinResponse(c, http.StatusNotFound, err.Error())
        return
    }

    q := `INSERT INTO artist(name,title,price) VALUES($1,$2,$3) ON CONFLICT DO NOTHING`
    result, err := env.DB.Exec(q, newAlbum.Artist, newAlbum.Title, newAlbum.Price)
    if err != nil {
        log.Printf("error occurred while inserting new record into artist table: %v", err)
        makeGinResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    // checking the number of rows affected
    n, err := result.RowsAffected()
    if err != nil {
        log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
        makeGinResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    // if no record was inserted, let us say client has failed
    if n == 0 {
        e := "could not insert the record, please try again after sometime"
        log.Println(e)
        makeGinResponse(c, http.StatusInternalServerError, e)
        return
    }

    // NOTE:
    //
    // Here I wanted to return the location for newly created Album but this
    // 'pq' library does not support, LastInsertionID functionality.
    m := "successfully created the record"
    log.Println(m)
    makeGinResponse(c, http.StatusOK, m)
}
