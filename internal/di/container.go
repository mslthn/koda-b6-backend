package di

import (
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	userHandler *handler.UserHandler
}

func NewContainer(conn *pgx.Conn) *Container {
	userRepo := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &Container{
		userHandler: userHandler,
	}
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}