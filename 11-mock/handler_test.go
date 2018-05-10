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

type PessoaServiceMock struct {
	SaveFn func(pessoa *Pessoa) error
}

func (psm *PessoaServiceMock) Save(pessoa *Pessoa) error {
	return psm.SaveFn(pessoa)
}

func TestWithHttpServer(t *testing.T) {
	pessoaServiceMock := &PessoaServiceMock{
		SaveFn: func(pessoa *Pessoa) error {
			return nil
		},
	}

	router := gin.New()
	router.POST("/pessoas", postPessoa(pessoaServiceMock))

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
