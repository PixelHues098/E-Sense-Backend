package register

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (env Env) RegisterNewUser(c *gin.Context) {

	var NewUser User

	if err := c.BindJSON(&NewUser); err != nil {
		log.Printf("invalid JSON body: %v", err)
		makeGinResponse(c, http.StatusNotFound, err.Error())
		return
	}

	q := `INSERT INTO artist(first_name, last_name, username, email, password) VALUES($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`
	result, err := env.DB.Exec(q, NewUser.FirstName, NewUser.LastName, NewUser.Username, NewUser.Email, NewUser.Password)
	if err != nil {
		log.Printf("error occurred while inserting new record into artist table: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if n == 0 {
		e := "could not insert the record, please try again after sometime"
		log.Println(e)
		makeGinResponse(c, http.StatusInternalServerError, e)
		return
	}

	m := "successfully created the record"
	log.Println(m)
	makeGinResponse(c, http.StatusOK, m)
}
