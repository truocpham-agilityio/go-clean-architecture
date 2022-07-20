package controller

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/usecase/interactor"
	"net/http"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
}

func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{ui}
}

func (uc *userController) GetUsers(c Context) error {
	var users []*model.User

	users, err := uc.userInteractor.Get(users)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
