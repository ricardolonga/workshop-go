package domain

type UserService interface {
	IsValid(user *User) bool
}

type User struct {
	ID string
	Name string
	Age int
}
