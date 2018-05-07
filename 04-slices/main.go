package main

import "fmt"

func main() {
	pessoa := &Pessoa{Nome: "Ricardo Longa", Idade: 31, Telefones: []string{"48999792240", "4832220978"}}
	fmt.Printf("%s\n", pessoa)

	for i, telefone := range pessoa.Telefones {
		fmt.Printf("Telefone [%d]: %s\n", i, telefone)
	}
}
