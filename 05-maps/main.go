package main

import "fmt"

func main() {
	nome := "Ricardo Longa"
	idade := 31
	telefones := []string{"48999792240", "4832220978"}
	parentes := map[string]interface{}{
		"father":   "Luigi",
		"mother":   "Rosana",
		"brothers": [...]string{"Fernanda", "Gian"},
	}

	pessoa := NewPessoa(nome, idade, telefones, parentes)
	fmt.Printf("%s\n", pessoa)

	for i, telefone := range pessoa.Telefones {
		fmt.Printf("Telefone [%d]: %s\n", i, telefone)
	}

	for grauParentesco, valor := range pessoa.Parentes {
		fmt.Printf("Grau de parentesco [%s]: %s\n", grauParentesco, valor)
	}

	father, exist := pessoa.Parentes["father"]
	if exist {
		fmt.Printf("Father: %s\n", father)
	}
}
