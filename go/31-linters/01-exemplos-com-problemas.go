package main

import (
	"fmt"
	"os"
	"time"
)

// Função exportada sem comentário - Revive detecta
func CalculateTotal(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total
}

// Função não utilizada - Staticcheck detecta
func unusedFunction() {
	fmt.Println("Nunca chamada")
}

// Variável com nome não convencional - Revive detecta
var User_Name string = "João"

func main() {
	// Variável não utilizada - Staticcheck detecta
	x := 10

	// Erro não verificado - errcheck detecta
	os.Open("arquivo.txt")

	// Uso incorreto de time.Sleep - Staticcheck detecta
	time.Sleep(100)

	// Comparação desnecessária - Staticcheck detecta
	if true == true {
		fmt.Println("Hello")
	}

	items := []int{1, 2, 3}
	total := CalculateTotal(items)
	fmt.Println(total)
}



