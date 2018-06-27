package user

import "github.com/ricardolonga/workshop-go/domain"

type service struct {
	userStorage domain.UserStorage
}

func NewService(userStorage domain.UserStorage) *service {
	return &service{
		userStorage: userStorage,
	}
}

func (*service) IsValid(user *domain.User) bool {
	return user.Age > 18
}
