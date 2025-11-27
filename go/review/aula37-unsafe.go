package main

import (
	"fmt"
	"unsafe"
)

/*
 * ============================================================================
 * AULA 37: UNSAFE PACKAGE
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - unsafe.Pointer
 * - Conversões de baixo nível
 * - Quando usar (raramente)
 */

func main() {
	fmt.Println("=== AULA 37: UNSAFE PACKAGE ===\n")
	
	var x int = 42
	ptr := unsafe.Pointer(&x)
	
	fmt.Printf("Endereço: %p\n", ptr)
	fmt.Printf("Tamanho: %d bytes\n", unsafe.Sizeof(x))
	
	fmt.Println("\n⚠️ Use apenas quando absolutamente necessário")
	fmt.Println("Quebra garantias de type safety")
}



