package main

import (
	"fmt"
	"math"
)

// Exemplo básico de package
func main() {
	fmt.Println("Bem-vindo à organização de código em Go!")
	
	// Exemplo de uso de função de outro package
	resultado := math.Sqrt(16)
	fmt.Printf("Raiz quadrada de 16: %.2f\n", resultado)
}

