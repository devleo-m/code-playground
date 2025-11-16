package main

import (
	"fmt"
	"os"  // ❌ Este import não está sendo usado - goimports remove
	"time"
)

// Este arquivo contém problemas intencionais para demonstrar
// o que go vet detecta. Use este arquivo para praticar!

func main() {
	// Problema 1: Código inalcançável
	// go vet detecta: unreachable code
	return
	fmt.Println("Esta linha nunca será executada")
	
	// Problema 2: Formato incorreto em printf
	// go vet detecta: wrong type for format specifier
	nome := "João"
	idade := 25
	fmt.Printf("Nome: %d, Idade: %s\n", nome, idade) // Tipos trocados!
	
	// Problema 3: Variável não utilizada
	// go vet detecta: unused variable
	resultado := 10 + 20
	_ = resultado // Comentar esta linha para ver o problema
	
	// Problema 4: Range loop problemático
	// go vet detecta: range variable is never modified
	numeros := []int{1, 2, 3}
	for _, num := range numeros {
		num = num * 2 // Modificando cópia, não o original!
	}
	fmt.Println("Números (não foram modificados):", numeros)
}

// Para testar:
// 1. Execute: go vet 02-exemplos-com-problemas.go
// 2. Veja todos os problemas detectados
// 3. Corrija cada problema
// 4. Execute go vet novamente para confirmar

