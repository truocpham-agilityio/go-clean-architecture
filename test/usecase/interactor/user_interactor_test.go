package interactor_test

import (
	"errors"
	"go-clean-architecture/domain/mocks"
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/presenter"
	"go-clean-architecture/usecase/interactor"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserInteractor_Get(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUser := model.User{
		ID:        1,
		Name:      "test",
		Email:     "test@gmail.com",
		Age:       "30",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockListUser := make([]*model.User, 0)
	mockListUser = append(mockListUser, &mockUser)

	t.Run("success", func(t *testing.T) {
		// mock FindAll call in UserRepository
		mockUserRepository.On("FindAll", mock.Anything).Return(mockListUser, nil).Once()

		// mock usecase interactor instance
		usecaseInteractor := interactor.NewUserInteractor(mockUserRepository, presenter.NewUserPresenter())

		// try to get users from usecase interactor
		users, err := usecaseInteractor.Get(mockListUser)

		assert.NoError(t, err)
		assert.Len(t, users, len(mockListUser))
	})

	t.Run("error", func(t *testing.T) {
		// mock FindAll call in UserRepository
		mockUserRepository.On("FindAll", mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()

		// mock usecase interactor instance
		usecaseInteractor := interactor.NewUserInteractor(mockUserRepository, presenter.NewUserPresenter())

		// try to get users from usecase interactor
		users, err := usecaseInteractor.Get(mockListUser)

		assert.Error(t, err)
		assert.Len(t, users, 0)
		mockUserRepository.AssertExpectations(t)
	})
}
