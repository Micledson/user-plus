package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"user-plus/data/repository"
	"user-plus/services"
	"user-plus/services/interfaces"
)

type UserHandler struct {
	service interfaces.IUserService
}

func NewUserHandler() *UserHandler {
	r := repository.NewUserRepository()
	s := services.NewUserService(r)
	return &UserHandler{service: s}
}

func NewUserHandlerFromService(s interfaces.IUserService) *UserHandler {
	return &UserHandler{service: s}
}

func (uh UserHandler) FindUserByEmail(ctx echo.Context) error {
	email := ctx.Param("email")
	if email == "" {
		log.Error("Email Vazio")
		return ctx.JSON(http.StatusBadRequest, "Email Vazio")
	}

	user, err := uh.service.FindUserByEmail(email)
	if err != nil {
		log.Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}
