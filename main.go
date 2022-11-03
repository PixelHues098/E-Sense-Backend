package main

import (
    "esense/controller"
    "esense/database"
    "esense/model"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    loadEnv()
    loadDatabase()
    serveApplication()
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
}

func serveApplication() {
    router := gin.Default()

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register-user", controller.Register)
    publicRoutes.POST("/login-user", controller.Login)
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:5000"}
    router.Use(cors.New(config))
    router.Run()
}
