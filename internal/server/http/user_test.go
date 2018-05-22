package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ricardolonga/workshop-go/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestWithRecorder(t *testing.T) {
	userService := user.NewService()

	router := NewHandler(userService)

	response := httptest.NewRecorder()
	endpoint := "/v1/users"
	body := []byte(`{"name": "Ricardo Longa", "age": 31}`)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewReader(body))

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)
}
