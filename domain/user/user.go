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

func (s *service) Create(user *domain.User) (*domain.User, error) {
	return s.userStorage.Insert(user)
}

func (s *service) Retrieve(userID string) (*domain.User, error) {
	return s.userStorage.GetByID(userID)
}

func (s *service) Delete(userID string) error {
	return s.userStorage.Delete(userID)
}
