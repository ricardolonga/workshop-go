package user

import "github.com/ricardolonga/workshop-go/domain"

type UserServiceMock struct {
	IsValidFn func(user *domain.User) bool
	IsValidInvoked bool
}

func (usm *UserServiceMock) IsValid(user *domain.User) bool {
	usm.IsValidInvoked = true
	return usm.IsValidFn(user)
}
