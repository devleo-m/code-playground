package main

// Este arquivo demonstra como goimports funciona
// Execute: goimports -d 03-exemplo-goimports.go (para ver diff)
// Execute: goimports -w 03-exemplo-goimports.go (para aplicar)

// ANTES do goimports:
// - Falta importar "fmt" (usado na linha 15)
// - Falta importar "strings" (usado na linha 16)
// - Falta importar "time" (usado na linha 17)
// - Tem import "os" que não está sendo usado

import "os"

func main() {
	// Usando fmt mas não importado
	fmt.Println("Hello, World!")
	
	// Usando strings mas não importado
	resultado := strings.ToUpper("hello")
	fmt.Println(resultado)
	
	// Usando time mas não importado
	time.Sleep(1 * time.Second)
}

// DEPOIS do goimports -w:
// 
// package main
//
// import (
//     "fmt"
//     "strings"
//     "time"
// )
//
// func main() {
//     fmt.Println("Hello, World!")
//     resultado := strings.ToUpper("hello")
//     fmt.Println(resultado)
//     time.Sleep(1 * time.Second)
// }
//
// Observe que:
// - "fmt" foi adicionado (estava sendo usado)
// - "strings" foi adicionado (estava sendo usado)
// - "time" foi adicionado (estava sendo usado)
// - "os" foi removido (não estava sendo usado)
// - Imports foram organizados em ordem alfabética

