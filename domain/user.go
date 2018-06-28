package domain

import (
	"fmt"
)

type UserService interface {
	IsValid(user *User) bool
	Create(user *User) (*User, error)
	Retrieve(userID string) (*User, error)
	Delete(userID string) error
}

type UserStorage interface {
	Insert(user *User) (*User, error)
	GetByID(userID string) (*User, error)
	Delete(userID string) error
}

type User struct {
	ID        string                 `json:"id,omitempty" bson:"_id"`
	Name      string                 `json:"name,omitempty" bson:"name"`
	Age       int                    `json:"age,omitempty" bson:"age"`
	Phones    []string               `json:"phones,omitempty" bson:"phones,omitempty"`
	Relatives map[string]interface{} `json:"-" bson:"-"`
}

func NewUser(name string, age int, phones []string) *User {
	return &User{Name: name, Age: age, Phones: phones}
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s - Age: %d", u.Name, u.Age)
}
