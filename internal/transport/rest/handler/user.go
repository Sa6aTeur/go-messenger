package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(s string)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterEndpoints(router *gin.Engine) {
	router.GET("/user", h.Get)
}

func (h *UserHandler) Get(gc *gin.Context) {
	query := gc.Query("id")
	if query == "" {
		gc.Error(errors.New("id not found"))
	}
	fmt.Println("UserHandler: get", query)

	h.service.Create(query)
}
