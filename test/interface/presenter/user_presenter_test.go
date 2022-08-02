package presenter_test

import (
	"go-clean-architecture/domain/model"
	"go-clean-architecture/interface/presenter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_UserPresenter_ResponseUsers(t *testing.T) {
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

	userPresenter := presenter.NewUserPresenter()
	responseUsers := userPresenter.ResponseUsers(mockListUser)

	assert.Equal(t, "Mr. test", responseUsers[0].Name)
	assert.Len(t, responseUsers, 1)
}
