package main

import "fmt"

// Pessoa ...
type Pessoa struct {
	Nome      string
	Idade     int
	Telefones [2]string
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Nome: %s\nIdade: %d\nTelefones: %s\nCelular: %s\nResidencial: %s", p.Nome, p.Idade, p.Telefones, p.Telefones[0], p.Telefones[1])
}
