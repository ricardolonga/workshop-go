package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	nome := "Ricardo Longa"
	idade := 31
	telefones := []string{"48999792240", "4832220978"}
	parentes := map[string]interface{}{
		"father":   "Luigi",
		"mother":   "Rosana",
		"brothers": [...]string{"Fernanda", "Gian"},
	}

	pessoa := NewPessoaCompleta(nome, idade, telefones, parentes)

	bytes, _ := json.Marshal(pessoa)
	fmt.Printf("%s\n", string(bytes))

	novaPessoa := NewPessoa("Fulano")

	bytes, _ = json.Marshal(novaPessoa)
	fmt.Printf("%s", string(bytes))
}
