package repository_test

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_UserRepository_FindAll(t *testing.T) {
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("expected nill, but got an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	gormDB, err := gorm.Open("mysql", sqlDB)
	if err != nil {
		t.Errorf("expected nil, but received an error '%s'", err)
	}
	defer gormDB.Close()

	mockListUser := []*model.User{
		{
			ID:        1,
			Name:      "test1",
			Email:     "test1@gmail.com",
			Age:       "30",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "test2",
			Email:     "test2@gmail.com",
			Age:       "30",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	sqlRows := sqlmock.NewRows([]string{"id", "name", "email", "age", "created_at", "updated_at"}).
		AddRow(mockListUser[0].ID, mockListUser[0].Name, mockListUser[0].Email, mockListUser[0].Age, mockListUser[0].CreatedAt, mockListUser[0].UpdatedAt).
		AddRow(mockListUser[1].ID, mockListUser[1].Name, mockListUser[1].Email, mockListUser[1].Age, mockListUser[1].CreatedAt, mockListUser[1].UpdatedAt)
	sqlQuery := "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL"
	mock.ExpectQuery(sqlQuery).WillReturnRows(sqlRows)

	userRepository := repository.NewUserRepository(gormDB)
	users, err := userRepository.FindAll(mockListUser)

	assert.NoError(t, err)
	assert.Len(t, users, len(mockListUser))
}
