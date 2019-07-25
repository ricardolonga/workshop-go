package http

import (
	"bytes"
	"encoding/json"
	"github.com/ricardolonga/workshop-go/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

type userServiceMock struct {
	IsValidFn func(u *domain.User) bool
	CreateFn  func(u *domain.User) (*domain.User, error)
}

func (usm *userServiceMock) IsValid(u *domain.User) bool {
	return usm.IsValidFn(u)
}

func (usm *userServiceMock) Create(u *domain.User) (*domain.User, error) {
	return usm.CreateFn(u)
}

func TestController(t *testing.T) {
	t.Run("invalid user", func(t *testing.T) {
		userService := &userServiceMock{
			IsValidFn: func(u *domain.User) bool {
				assert.NotNil(t, u)
				return false
			},
		}

		router := NewHandler(userService)

		body := []byte(`{"name": "Ricardo Longa", "age": 17}`)
		req, _ := http.NewRequest(http.MethodPost, "/v1/users", bytes.NewReader(body))

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("valid user", func(t *testing.T) {
		userService := &userServiceMock{
			IsValidFn: func(u *domain.User) bool {
				return true
			},
			CreateFn: func(u *domain.User) (user *domain.User, e error) {
				return u, nil
			},
		}

		router := NewHandler(userService)

		body := []byte(`{"name": "Ricardo Longa", "age": 18}`)
		req, _ := http.NewRequest(http.MethodPost, "/v1/users", bytes.NewReader(body))

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)

		body, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		var u domain.User
		assert.NoError(t, json.Unmarshal(body, &u))
		assert.Equal(t, "Ricardo Longa", u.Name)
		assert.Equal(t, 18, u.Age)
	})
}
