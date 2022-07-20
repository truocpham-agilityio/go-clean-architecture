package presenter

import "go-clean-architecture/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}