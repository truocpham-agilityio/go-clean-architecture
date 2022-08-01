package controller_test

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/controller"
	mocks "go-clean-architecture/mocks/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_UserController_GetUsers(t *testing.T) {
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

	mockUserUsecaseInteractor := new(mocks.UserUsecase)
	mockUserUsecaseInteractor.On("Get", mock.Anything).Return(mockListUser, nil).Once()

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err = controller.NewUserController(mockUserUsecaseInteractor).GetUsers(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserUsecaseInteractor.AssertExpectations(t)
}
