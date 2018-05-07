package main

import "fmt"

func main() {
	pessoa := &Pessoa{Nome: "Ricardo Longa", Idade: 31}
	fmt.Printf("%s", pessoa)
}
