package di

import (
	"github.com/jackc/pgx/v5"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
)

type Container struct {
	conn *pgx.Conn

	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handler.UserHandler

}

func NewContainer(conn *pgx.Conn) *Container {
	container := Container{
		conn: conn,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.conn)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)


}

func (c *Container) UserHandler() *handler.UserHandler{
	return c.userHandler
}