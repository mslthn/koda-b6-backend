package routes

import (
	"main/internal/di"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r *gin.Engine, conn *pgx.Conn) {
	ctn := di.NewContainer(conn)

	// Route
	r.POST("/register", ctn.UserHandler().Register)
	r.POST("/login", ctn.UserHandler().Login)

	users := r.Group("/users")
	{
		users.GET("/", ctn.UserHandler().GetAll)
		users.GET("/:id", ctn.UserHandler().GetById)
		users.DELETE("/:id", ctn.UserHandler().Delete)
		users.PATCH("/:id", ctn.UserHandler().Update)
	}
}