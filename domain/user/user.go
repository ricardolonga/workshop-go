package user

import (
	"github.com/ricardolonga/workshop-go/domain"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (*Service) IsValid(user *domain.User) bool {
	return false
}

func (s *Service) Create(user *domain.User) (*domain.User, error) {
	user.ID = uuid.NewV4().String()
	return user, nil
}
