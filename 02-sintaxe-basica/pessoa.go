package main

import "fmt"

// Pessoa ...
type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Nome: %s - Idade: %d", p.Nome, p.Idade)
}
