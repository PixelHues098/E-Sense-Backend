package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    _ "github.com/lib/pq"
)

// DeleteAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that corresponding message.
func (env Env) DeleteAlbumByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
        log.Println(e)
        makeGinResponse(c, http.StatusNotFound, e)
        return
    }

    q := `DELETE FROM artist WHERE id = $1;`
    result, err := env.DB.Exec(q, id)
    if err != nil {
        e := fmt.Sprintf("error occurred while deleting artist record with id: %d and error is: %v", id, err)
        log.Println(e)
        makeGinResponse(c, http.StatusInternalServerError, e)
        return
    }

    // checking the number of rows affected
    n, err := result.RowsAffected()
    if err != nil {
        e := fmt.Sprintf("error occurred while checking the returned result from database after deletion: %v", err)
        log.Println(e)
        makeGinResponse(c, http.StatusInternalServerError, e)
        return
    }

    // if no record was deleted, let us inform that there might be no
    // records to delete for this given album ID.
    if n == 0 {
        e := "could not delete the record, there might be no records for the given ID"
        log.Println(e)
        makeGinResponse(c, http.StatusBadRequest, e)
        return
    }

    m := "successfully deleted the record"
    log.Println(m)
    makeGinResponse(c, http.StatusOK, m)
}

