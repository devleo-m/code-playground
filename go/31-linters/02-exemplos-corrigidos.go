package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// CalculateTotal calcula a soma total dos itens fornecidos
func CalculateTotal(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total
}

func main() {
	// Abrir arquivo e tratar erro adequadamente
	file, err := os.Open("arquivo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Uso correto de time.Sleep
	time.Sleep(100 * time.Millisecond)

	// Comparação simplificada
	if true {
		fmt.Println("Hello")
	}

	// Variável com nome convencional
	var userName string = "João"
	fmt.Println(userName)

	items := []int{1, 2, 3}
	total := CalculateTotal(items)
	fmt.Println(total)
}

