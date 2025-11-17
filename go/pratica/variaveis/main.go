package main

import (
	index "variaveis/01-teste"
	test "variaveis/02-teste"
)

import "fmt"

const name = "Arolnode"
const age = 12
var habilitado bool

func main() {
	sobrenome := "Madeira"

	index.Teste()
	test.Test()

	fullName := fmt.Sprintf("%s %s", name, sobrenome)

	fmt.Println(fullName)

	if age >= 18 {
		fmt.Println("Eh maior de idade e ja pode dirigir")
		habilitado = true
	} else {
		fmt.Println("Voce eh demenor!!!!")
	}

	fmt.Println("Pode dirigir?")
	fmt.Println(habilitado)
}
