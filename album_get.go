package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (env Env) GetAlbumByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        e := fmt.Sprintf("received invalid id path param which is not string: %v", c.Param("id"))
        log.Println(e)
        makeGinResponse(c, http.StatusNotFound, e)
        return
    }

    var name, title string
    var price string

    q := `SELECT * FROM artist where id=$1`
    row := env.DB.QueryRow(q, id)

    err = row.Scan(&id, &name, &title, &price)
    switch err {
    case sql.ErrNoRows:
        log.Printf("no rows are present for alubum with id: %d", id)
        makeGinResponse(c, http.StatusBadRequest, err.Error())
    case nil:
        log.Printf("we are able to fetch album with given id: %d", id)
        c.JSON(http.StatusOK, NewAlbum(id, title, name, price))
    default:
        e := fmt.Sprintf("error: %v occurred while reading the databse for Album record with id: %d", err, id)
        log.Println(e)
        makeGinResponse(c, http.StatusInternalServerError, err.Error())
    }
}

// GetAlbums responds with the list of all albums as JSON.
func (env Env) GetAlbums(c *gin.Context) {
    // Note:
    //
    // pagnination can be impleted in may ways, but I am following one of the way,
    // if you feel this is confusing, please read medium article that I have added below
    // For this page and perPage isseus, front end engineers can take care of it
    // by escaping and sanitization of query params.
    // Please see: https://www.enterprisedb.com/postgres-tutorials/how-use-limit-and-offset-postgresql
    // Please see: https://levelup.gitconnected.com/creating-a-data-pagination-function-in-postgresql-2a032084af54
    page := c.Query("page") // AKA limit in SQL terms
    if page == "" {
        e := "missing query param: page"
        log.Println(e)
        makeGinResponse(c, http.StatusNotFound, e)
        return
    }

    perPage := c.Query("perPage") // AKA limit in SQL terms
    if perPage == "" {
        e := "missing query param: perPage"
        log.Println(e)
        makeGinResponse(c, http.StatusNotFound, e)
        return
    }

    limit, err := strconv.Atoi(page)
    if err != nil {
        e := fmt.Sprintf("received invalid page query param which is not integer : %v", page)
        log.Println(e)
        makeGinResponse(c, http.StatusBadRequest, e)
        return
    }

    if limit > recordFetchLimit {
        // Seems some bad user or front end developer playing with query params!
        e := fmt.Sprintf("we agreed to fetch less than %d records but we received request for %d", recordFetchLimit, limit)
        log.Println(e)
        makeGinResponse(c, http.StatusBadRequest, e)
        return
    }

    offset, err := strconv.Atoi(perPage)
    if err != nil {
        e := fmt.Sprintf("received invalid offset query param which is not integer : %v", page)
        log.Println(e)
        makeGinResponse(c, http.StatusBadRequest, e)
        return
    }

    // anyway, let's check if offset is a negative value
    if offset < 0 {
        e := "offset query param cannot be negative"
        log.Println(e)
        makeGinResponse(c, http.StatusBadRequest, e)
        return
    }

    q := `SELECT id,name,title,price FROM artist LIMIT $1 OFFSET $2;`
    rows, err := env.DB.Query(q, limit, offset)
    switch err {
    case sql.ErrNoRows:
        defer rows.Close()
        e := "no rows records found in artist table to read"
        log.Println(e)
        makeGinResponse(c, http.StatusBadRequest, e)
    case nil:
        defer rows.Close()
        a := make([]Album, 0)
        var rowsReadErr bool
        for rows.Next() {
            var id int
            var name, title string
            var price string
            err = rows.Scan(&id, &name, &title, &price)
            if err != nil {
                log.Printf("error occurred while reading the database rows: %v", err)
                rowsReadErr = true
                break
            }
            a = append(a, NewAlbum(id, title, name, price))
        }

        if rowsReadErr {
            log.Println("we are not able to fetch few records")
        }

        // let's return the read rows at least
        log.Printf("we are able to fetch albums for requested limit: %d and offest: %d", limit, offset)
        c.JSON(http.StatusOK, a)
    default:
        defer rows.Close()
        // this should not happen
        e := "some internal database server error"
        log.Println(e)
        makeGinResponse(c, http.StatusInternalServerError, e)
    }
}
