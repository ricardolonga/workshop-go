package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ricardolonga/workshop-go/domain"
	"github.com/ricardolonga/workshop-go/domain/user"
	"fmt"
)

func init() {
	fmt.Println("init...")
}

func TestMain(m *testing.M) {
	fmt.Println("antes...")
	m.Run()
	fmt.Println("...depois")
}

func TestController(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService := &user.ServiceMock{
			IsValidFn: func(user *domain.User) bool {
				assert.Equal(t, "Ricardo Longa", user.Name)
				assert.Equal(t, 31, user.Age)

				return true
			},
		}

		router := NewHandler(userService)

		response := httptest.NewRecorder()
		endpoint := "/v1/users"
		body := []byte(`{"name": "Ricardo Longa", "age": 31}`)

		req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)

		assert.Equal(t, http.StatusCreated, response.Code)

		assert.Equal(t, 1, userService.IsValidCount)
	})

	t.Run("failed", func(t *testing.T) {
		userService := &user.ServiceMock{
			IsValidFn: func(user *domain.User) bool {
				assert.Equal(t, "Ricardo Longa", user.Name)
				assert.Equal(t, 17, user.Age)

				return false
			},
		}

		router := NewHandler(userService)

		response := httptest.NewRecorder()
		endpoint := "/v1/users"
		body := []byte(`{"name": "Ricardo Longa", "age": 17}`)

		req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))

		router.ServeHTTP(response, req)

		assert.Equal(t, http.StatusBadRequest, response.Code)

		assert.Equal(t, 1, userService.IsValidCount)
	})
}
