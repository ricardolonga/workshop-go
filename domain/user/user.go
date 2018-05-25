package user

import "github.com/ricardolonga/workshop-go/domain"

type service struct{}

func NewService() *service {
	return &service{}
}

func (*service) IsValid(user *domain.User) bool {
	return user.Age > 18
}
