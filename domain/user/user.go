package user

import "github.com/ricardolonga/workshop-go/domain"

type userService struct {}

func NewService() *userService {
	return &userService{}
}

func(u *userService) IsValid(user *domain.User) bool {
	return user.Age > 18
}