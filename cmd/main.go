package main

import (
	"context"
	"fmt"
	"main/internal/routes"
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func corsMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        origin := ctx.Request.Header.Get("Origin")
        frontendURL := os.Getenv("FRONTEND_URL")

        if origin == frontendURL || frontendURL == "" {
            ctx.Header("Access-Control-Allow-Origin", origin)
        }

        ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        ctx.Header("Access-Control-Allow-Credentials", "true")

        if ctx.Request.Method == "OPTIONS" {
            ctx.AbortWithStatus(204)
            return
        }
        ctx.Next()
    }
}

func main() {
	godotenv.Load()

	r := gin.Default()
	r.Use(corsMiddleware())

	dbUrl := os.Getenv("DATABASE_URL")
	connConfig, err := pgx.ParseConfig(dbUrl)
	if err != nil {
		fmt.Println("Failed to parse config")
		return
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		fmt.Println("Connection Failed")
		return
	}
	defer conn.Close(context.Background())

	routes.SetupRoutes(r, conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	r.Run(fmt.Sprintf(":%s", port))
}