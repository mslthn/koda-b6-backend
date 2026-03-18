package handler

import (
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	err := h.service.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "User registered successfully",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid email or password format",
		})
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login successful",
		Data:    models.LoginUserResponse{Token: token},
	})
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success fetch all users",
		Data:    users,
	})
}

func (h *UserHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID format",
		})
		return
	}

	user, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success fetch user",
		Data:    user,
	})
}

func (h *UserHandler) Update(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, models.Response{
            Message: "Invalid ID format",
        })
        return
    }

    var req models.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.Response{
            Message: "Invalid input data",
        })
        return
    }

    user, err := h.service.Update(id, &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.Response{
            Message: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, models.Response{
        Message: "User updated successfully",
        Data:    user,
    })
}

func (h *UserHandler) Delete(c *gin.Context) {
	email := c.Param("email")
	err := h.service.Delete(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User deleted successfully",
	})
}