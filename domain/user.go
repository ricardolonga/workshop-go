package domain

type UserService interface {
	IsValid(user *User) bool
}

type User struct {
	ID        string                 `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Age       int                    `json:"age,omitempty"`
	Phones    []string               `json:"phones,omitempty"`
	Relatives map[string]interface{} `json:"-"`
}

func NewUser(name string, age int, phones []string) *User {
	return &User{Name: name, Age: age, Phones: phones}
}
