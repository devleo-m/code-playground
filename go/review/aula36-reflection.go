package main

import (
	"fmt"
	"reflect"
)

/*
 * ============================================================================
 * AULA 36: REFLECTION
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - reflect package
 * - TypeOf, ValueOf
 * - Type assertions
 * - Use cases
 */

func main() {
	fmt.Println("=== AULA 36: REFLECTION ===\n")
	
	var x int = 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	
	fmt.Printf("Tipo: %s\n", t)
	fmt.Printf("Valor: %v\n", v.Int())
	
	fmt.Println("\nReflection permite inspecionar tipos em runtime")
	fmt.Println("Use com cuidado (performance)")
}





