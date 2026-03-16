package main

import (
	"context"
	"fmt"
	"main/internal/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8888")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
			ctx.Abort()
		} else {
			ctx.Next()
		}
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
	r.Run(fmt.Sprintf("localhost:%s", port))
}