package repository

import (
	"go-clean-architecture/domain/model"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(users []*model.User) ([]*model.User, error) {
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
