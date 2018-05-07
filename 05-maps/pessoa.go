package main

import "fmt"

// Pessoa ...
type Pessoa struct {
	Nome      string
	Idade     int
	Telefones []string
	Parentes  map[string]interface{}
}

// NewPessoa ...
func NewPessoa(nome string, idade int, telefones []string, parentes map[string]interface{}) *Pessoa {
	return &Pessoa{nome, idade, telefones, parentes}
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Nome: %s - Idade: %d", p.Nome, p.Idade)
}
