package mocks

import (
	"go-clean-architecture/domain/model"

	"github.com/stretchr/testify/mock"
)

type UserUsecase struct {
	mock.Mock
}

func (_m *UserUsecase) Get(u []*model.User) ([]*model.User, error) {
	ret := _m.Called(u)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func([]*model.User) []*model.User); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*model.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}