package domain

type Service interface {
	IsValid(u *User) bool
	Create(u *User) (*User, error)
}

type User struct {
	ID        string                 `json:"id,omitempty" bson:"_id"`
	Name      string                 `json:"name,omitempty" bson:"name"`
	Age       int                    `json:"age,omitempty" bson:"age"`
	Phones    []string               `json:"phones,omitempty" bson:"phones,omitempty"`
	Relatives map[string]interface{} `json:"-" bson:"-"`
}
