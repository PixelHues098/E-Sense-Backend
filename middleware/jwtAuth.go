package middleware

import (
    "esense/helper"
    "github.com/gin-gonic/gin"
    "net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        err := helper.ValidateJWT(context)
        if err != nil {
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            context.Abort()
            return
        }
        context.Next()
    }
}
