package user

import "github.com/ricardolonga/workshop-go/domain"

// ServiceMock ...
type ServiceMock struct {
	IsValidFn      func(user *domain.User) bool
	IsValidInvoked bool
	IsValidCount   int
}

func (sm *ServiceMock) IsValid(user *domain.User) bool {
	sm.IsValidInvoked = true
	sm.IsValidCount++
	return sm.IsValidFn(user)
}
