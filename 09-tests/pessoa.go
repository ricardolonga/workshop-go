package main

import "fmt"

// Pessoa ...
type Pessoa struct {
	Nome      string                 `json:"nome"`
	Idade     int                    `json:"idade,omitempty"`
	Telefones []string               `json:"telefones,omitempty"`
	Parentes  map[string]interface{} `json:"-"`
}

// NewPessoa ...
func NewPessoa(nome string) *Pessoa {
	return &Pessoa{Nome: nome}
}

// NewPessoa ...
func NewPessoaCompleta(nome string, idade int, telefones []string, parentes map[string]interface{}) *Pessoa {
	return &Pessoa{nome, idade, telefones, parentes}
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Nome: %s", p.Nome)
}
