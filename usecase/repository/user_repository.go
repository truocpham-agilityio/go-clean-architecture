package repository

import "go-clean-architecture/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
