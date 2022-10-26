package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to DB first
	env := new(Env)
	var err error
	env.DB, err = ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5000"}
	router := gin.Default()
	router.POST("/albums", env.PostAlbum)
	router.POST("/register-user", env.RegisterNewUser)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(cors.New(config))
	router.Run()
}



