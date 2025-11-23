package main

import (
	"fmt"
	"math"
)

/*
 * ============================================================================
 * AULA 13: INTERFACES
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Contratos que especificam métodos
 * - Satisfação implícita (não precisa declarar "implements")
 * - Permite polimorfismo
 * - Empty interface (interface{}): Aceita qualquer tipo
 * - Type assertion: Verificar e converter tipo
 * - Type switch: Verificar tipo de interface{}
 */

// Interface que define contrato
type Forma interface {
	Area() float64
	Perimetro() float64
}

// Retangulo implementa Forma (implicitamente)
type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

// Circulo implementa Forma (implicitamente)
type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

// Função que aceita qualquer Forma
func ImprimirArea(f Forma) {
	fmt.Printf("   Área: %.2f, Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// Empty interface
func ImprimirQualquerCoisa(valor interface{}) {
	fmt.Printf("   Tipo: %T, Valor: %v\n", valor, valor)
}

func main() {
	fmt.Println("=== AULA 13: INTERFACES ===\n")
	
	// ===== INTERFACES BÁSICAS =====
	fmt.Println("1. INTERFACES BÁSICAS:")
	ret := Retangulo{Largura: 10, Altura: 5}
	circ := Circulo{Raio: 3}
	
	// Ambos podem ser usados como Forma
	ImprimirArea(ret)
	ImprimirArea(circ)
	
	// ===== POLIMORFISMO =====
	fmt.Println("\n2. POLIMORFISMO:")
	formas := []Forma{
		Retangulo{Largura: 5, Altura: 3},
		Circulo{Raio: 2},
		Retangulo{Largura: 4, Altura: 4},
	}
	
	for i, forma := range formas {
		fmt.Printf("   Forma %d: ", i+1)
		ImprimirArea(forma)
	}
	
	// ===== EMPTY INTERFACE =====
	fmt.Println("\n3. EMPTY INTERFACE (interface{}):")
	ImprimirQualquerCoisa(42)
	ImprimirQualquerCoisa("Olá")
	ImprimirQualquerCoisa(3.14)
	ImprimirQualquerCoisa([]int{1, 2, 3})
	
	// ===== TYPE ASSERTION =====
	fmt.Println("\n4. TYPE ASSERTION:")
	var valor interface{} = 42
	
	// Type assertion
	if v, ok := valor.(int); ok {
		fmt.Printf("   É um int: %d\n", v)
	}
	
	if v, ok := valor.(string); ok {
		fmt.Printf("   É uma string: %s\n", v)
	} else {
		fmt.Println("   Não é uma string")
	}
	
	// ===== TYPE SWITCH =====
	fmt.Println("\n5. TYPE SWITCH:")
	processarTipo(42)
	processarTipo("Olá")
	processarTipo(3.14)
	processarTipo(true)
	processarTipo([]int{1, 2, 3})
}

func processarTipo(valor interface{}) {
	switch v := valor.(type) {
	case int:
		fmt.Printf("   É um inteiro: %d\n", v)
	case string:
		fmt.Printf("   É uma string: %s\n", v)
	case float64:
		fmt.Printf("   É um float64: %.2f\n", v)
	case bool:
		fmt.Printf("   É um booleano: %t\n", v)
	default:
		fmt.Printf("   Tipo desconhecido: %T\n", v)
	}
}


