package user

import (
	"testing"

	"encoding/json"

	"github.com/ricardolonga/workshop-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestIsAValidUser(t *testing.T) {
	user := domain.NewUser("Ricardo", 31, []string{"99979-2240"})

	userService := NewService()

	assert.True(t, userService.IsValid(user))
}

func TestIsntAValidUser(t *testing.T) {
	user := domain.NewUser("Ricardo", 17, []string{"99979-2240"})

	userService := NewService()

	assert.False(t, userService.IsValid(user))
}

func TestUserJson(t *testing.T) {
	user := domain.NewUser("Ricardo", 31, []string{"99979-2240"})

	user.Relatives = map[string]interface{}{
		"father":   "Luigi",
		"mother":   "Rosana",
		"siblings": []string{"Fernanda", "Gian"},
	}

	userBytes, err := json.Marshal(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, userBytes)

	expected := "{\"name\":\"Ricardo\", \"age\":31, \"phones\":[\"99979-2240\"]}"
	actual := string(userBytes)

	assert.JSONEq(t, expected, actual)
}
