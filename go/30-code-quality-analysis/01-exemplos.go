package main

import (
	"fmt"
	"strings"
	"time"
)

// Exemplo 1: Código correto - go vet não encontra problemas
func exemploCorreto() {
	nome := "João"
	idade := 25
	fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}

// Exemplo 2: Problema que go vet detecta - código inalcançável
// Descomente para ver o problema:
/*
func exemploInalcançavel() {
	return
	fmt.Println("Esta linha nunca será executada") // go vet detecta isso!
}
*/

// Exemplo 3: Problema que go vet detecta - formato incorreto em printf
// Descomente para ver o problema:
/*
func exemploFormatoIncorreto() {
	nome := "Maria"
	idade := 30
	// go vet detecta: %d espera int, mas recebe string
	// go vet detecta: %s espera string, mas recebe int
	fmt.Printf("Nome: %d, Idade: %s\n", nome, idade)
}
*/

// Exemplo 4: Problema com range loops
func exemploRangeLoop() {
	numeros := []int{1, 2, 3, 4, 5}
	
	// ❌ ERRADO: Modificando a cópia, não o original
	// go vet avisa sobre isso
	for _, num := range numeros {
		_ = num * 2 // Isso não modifica o slice original!
	}
	
	// ✅ CORRETO: Modificando o original
	for i := range numeros {
		numeros[i] = numeros[i] * 2
	}
	
	fmt.Println("Números dobrados:", numeros)
}

// Exemplo 5: Uso correto de imports
func exemploImports() {
	// strings é usado aqui
	texto := strings.ToUpper("hello world")
	fmt.Println(texto)
	
	// time é usado aqui
	time.Sleep(100 * time.Millisecond)
}

// Exemplo 6: Função que demonstra boas práticas
func exemploBoasPraticas(nome string, idade int) string {
	if nome == "" {
		return "Nome não pode ser vazio"
	}
	
	if idade < 0 {
		return "Idade não pode ser negativa"
	}
	
	return fmt.Sprintf("%s tem %d anos", nome, idade)
}

func main() {
	fmt.Println("=== Exemplos de Code Quality ===")
	
	fmt.Println("\n1. Exemplo Correto:")
	exemploCorreto()
	
	fmt.Println("\n2. Exemplo Range Loop:")
	exemploRangeLoop()
	
	fmt.Println("\n3. Exemplo Imports:")
	exemploImports()
	
	fmt.Println("\n4. Exemplo Boas Práticas:")
	fmt.Println(exemploBoasPraticas("João", 25))
	fmt.Println(exemploBoasPraticas("", 30))      // Nome vazio
	fmt.Println(exemploBoasPraticas("Maria", -5)) // Idade negativa
	
	fmt.Println("\n=== Execute 'go vet' neste arquivo para ver verificações ===")
	fmt.Println("=== Execute 'goimports -w' para organizar imports ===")
}

