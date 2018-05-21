package user

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ricardolonga/workshop-go/domain"
)

func TestValidUser(t *testing.T) {
	user := &domain.User{ID: "1", Name: "Ricardo", Age: 31}

	userService := NewService()

	assert.True(t, userService.IsValid(user))
}
