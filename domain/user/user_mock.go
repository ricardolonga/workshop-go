package user

import "github.com/ricardolonga/workshop-go/domain"

// ServiceMock ...
type ServiceMock struct {
	IsValidFn      func(user *domain.User) bool
	IsValidInvoked bool
	IsValidCount   int

	CreateFn   func(user *domain.User) (*domain.User, error)
	RetrieveFn func(userID string) (*domain.User, error)
	DeleteFn   func(userID string) error
}

func (sm *ServiceMock) IsValid(user *domain.User) bool {
	sm.IsValidInvoked = true
	sm.IsValidCount++
	return sm.IsValidFn(user)
}

func (s *ServiceMock) Create(user *domain.User) (*domain.User, error) {
	return s.CreateFn(user)
}

func (s *ServiceMock) Retrieve(userID string) (*domain.User, error) {
	return s.RetrieveFn(userID)
}

func (s *ServiceMock) Delete(userID string) error {
	return s.DeleteFn(userID)
}
