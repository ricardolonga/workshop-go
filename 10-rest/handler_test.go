package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWithRecorder(t *testing.T) {
	router := gin.New()
	router.POST("/pessoas", postPessoa)

	response := httptest.NewRecorder()

	endpoint := "/pessoas"
	body := []byte(`{"nome": "Ricardo Longa"}`)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	assert.NoError(t, err)

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestWithHttpServer(t *testing.T) {
	router := gin.New()
	router.POST("/pessoas", postPessoa)

	server := httptest.NewServer(router)
	defer server.Close()

	URL, _ := url.Parse(server.URL)
	endpoint := fmt.Sprintf("%s/pessoas", URL)
	body := []byte(`{"nome": "Ricardo Longa"}`)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, res.StatusCode)
}
