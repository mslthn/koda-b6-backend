package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func cors() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Header("Access-Control-Allow-Origin", "http://localhost:8888")
        ctx.Header("Access-control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        ctx.Header("Access-Cortrol-Allow-Headers", "Content-Type")
        
        if ctx.Request.Method == "OPTIONS"{
            ctx.Data(http.StatusOK, "", []byte(""))
        } else {
            ctx.Next()
        }
    }
}

func main() {
    godotenv.Load()
    r := gin.Default()
    r.Use(cors())

    

    r.Run("localhost:8888")
}