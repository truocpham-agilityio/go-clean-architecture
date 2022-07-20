package presenter

import "go-clean-architecture/domain/model"

type userPresenter struct{}

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(users []*model.User) []*model.User {
	for _, u := range users {
		u.Name = "Mr. " + u.Name
	}

	return users
}
