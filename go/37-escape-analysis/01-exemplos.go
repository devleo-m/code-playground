package main

import (
	"fmt"
	"os"
	"strings"
)

// Exemplo 1: Stack allocation (não escapa)
func exemploStack() int {
	x := 42
	return x
}

// Exemplo 2: Heap allocation (escapa)
func exemploHeap() *int {
	x := 42
	return &x
}

// Exemplo 3: Interface pode causar escape
func exemploInterface(v interface{}) {
	fmt.Println(v)
}

// Exemplo 4: Tipo concreto não escapa
func exemploTipoConcreto(v int) {
	fmt.Printf("%d\n", v)
}

// Exemplo 5: Closure que captura variável
func exemploClosure() func() int {
	x := 42
	return func() int {
		return x
	}
}

// Exemplo 6: Closure sem captura
func exemploClosureSemCaptura() func() int {
	return func() int {
		return 42
	}
}

// Exemplo 7: Slice sem pré-alocação
func exemploSliceRuim(size int) []int {
	var result []int
	for i := 0; i < size; i++ {
		result = append(result, i)
	}
	return result
}

// Exemplo 8: Slice com pré-alocação
func exemploSliceBom(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, i)
	}
	return result
}

// Exemplo 9: Retornar pointer de struct pequena
func exemploPointerStruct() *Point {
	return &Point{X: 10, Y: 20}
}

// Exemplo 10: Retornar valor de struct pequena
func exemploValorStruct() Point {
	return Point{X: 10, Y: 20}
}

type Point struct {
	X, Y int
}

// Exemplo 11: Concatenação de strings (ruim)
func exemploConcatRuim(parts []string) string {
	msg := ""
	for _, part := range parts {
		msg += part
	}
	return msg
}

// Exemplo 12: strings.Builder (bom)
func exemploConcatBom(parts []string) string {
	var builder strings.Builder
	builder.Grow(len(parts) * 10)
	for _, part := range parts {
		builder.WriteString(part)
	}
	return builder.String()
}

// Exemplo 13: Goroutine que captura variável
func exemploGoroutine() {
	x := 42
	go func() {
		fmt.Println(x)
	}()
}

// Exemplo 14: Goroutine com parâmetro
func exemploGoroutineBom() {
	x := 42
	go func(val int) {
		fmt.Println(val)
	}(x)
}

// Função para demonstrar análise de escape
func demonstrarEscapeAnalysis() {
	fmt.Println("\n=== Demonstração de Escape Analysis ===")
	fmt.Println("Compile com: go build -gcflags=\"-m\" para ver decisões")
	fmt.Println("\nCenários demonstrados:")
	fmt.Println("1. Stack allocation (não escapa)")
	fmt.Println("2. Heap allocation (escapa)")
	fmt.Println("3. Interface (pode escapar)")
	fmt.Println("4. Tipo concreto (não escapa)")
	fmt.Println("5. Closure com captura (escapa)")
	fmt.Println("6. Closure sem captura (não escapa)")
	fmt.Println("7. Slice sem pré-alocação (pode escapar)")
	fmt.Println("8. Slice com pré-alocação (menos provável de escapar)")
	fmt.Println("9. Pointer de struct (escapa)")
	fmt.Println("10. Valor de struct (não escapa)")
	fmt.Println("11. Concatenação de strings (ruim)")
	fmt.Println("12. strings.Builder (bom)")
	fmt.Println("13. Goroutine com captura (escapa)")
	fmt.Println("14. Goroutine com parâmetro (não escapa)")
}

// Função para comparar versões
func compararVersoes() {
	fmt.Println("\n=== Comparação de Versões ===")

	// Comparar slices
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Println("\n1. Slices:")
	fmt.Println("   Sem pré-alocação:", len(exemploSliceRuim(5)))
	fmt.Println("   Com pré-alocação:", len(exemploSliceBom(5)))

	// Comparar structs
	fmt.Println("\n2. Structs:")
	p1 := exemploPointerStruct()
	p2 := exemploValorStruct()
	fmt.Printf("   Pointer: %+v\n", p1)
	fmt.Printf("   Valor: %+v\n", p2)

	// Comparar concatenação
	parts := []string{"Hello", " ", "World", "!"}
	fmt.Println("\n3. Concatenação:")
	fmt.Println("   Ruim:", exemploConcatRuim(parts))
	fmt.Println("   Bom:", exemploConcatBom(parts))

	// Comparar closures
	fmt.Println("\n4. Closures:")
	closure1 := exemploClosure()
	closure2 := exemploClosureSemCaptura()
	fmt.Printf("   Com captura: %d\n", closure1())
	fmt.Printf("   Sem captura: %d\n", closure2())
}

// Menu interativo
func showMenu() {
	fmt.Println("\n=== Exemplos de Escape Analysis ===")
	fmt.Println("1. Stack vs Heap allocation")
	fmt.Println("2. Interface vs Tipo concreto")
	fmt.Println("3. Closures (com vs sem captura)")
	fmt.Println("4. Slices (sem vs com pré-alocação)")
	fmt.Println("5. Structs (pointer vs valor)")
	fmt.Println("6. Concatenação de strings")
	fmt.Println("7. Goroutines (captura vs parâmetro)")
	fmt.Println("8. Informações sobre escape analysis")
	fmt.Println("9. Comparar todas as versões")
	fmt.Println("0. Sair")
	fmt.Print("\nEscolha uma opção: ")
}

func main() {
	if len(os.Args) > 1 {
		// Modo não-interativo: executar exemplo específico
		switch os.Args[1] {
		case "1":
			fmt.Println("Stack:", exemploStack())
			fmt.Println("Heap:", *exemploHeap())
		case "2":
			exemploInterface(42)
			exemploTipoConcreto(42)
		case "3":
			closure1 := exemploClosure()
			closure2 := exemploClosureSemCaptura()
			fmt.Println("Com captura:", closure1())
			fmt.Println("Sem captura:", closure2())
		case "4":
			fmt.Println("Sem pré-alocação:", exemploSliceRuim(5))
			fmt.Println("Com pré-alocação:", exemploSliceBom(5))
		case "5":
			p1 := exemploPointerStruct()
			p2 := exemploValorStruct()
			fmt.Printf("Pointer: %+v\n", p1)
			fmt.Printf("Valor: %+v\n", p2)
		case "6":
			parts := []string{"Hello", " ", "World"}
			fmt.Println("Ruim:", exemploConcatRuim(parts))
			fmt.Println("Bom:", exemploConcatBom(parts))
		case "7":
			fmt.Println("Executando goroutines...")
			exemploGoroutine()
			exemploGoroutineBom()
		case "8":
			demonstrarEscapeAnalysis()
		case "9":
			compararVersoes()
		default:
			fmt.Println("Opção inválida")
		}
		return
	}

	// Modo interativo
	for {
		showMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nStack:", exemploStack())
			fmt.Println("Heap:", *exemploHeap())
		case 2:
			fmt.Println("\nInterface:")
			exemploInterface(42)
			fmt.Println("Tipo concreto:")
			exemploTipoConcreto(42)
		case 3:
			fmt.Println("\nClosures:")
			closure1 := exemploClosure()
			closure2 := exemploClosureSemCaptura()
			fmt.Println("Com captura:", closure1())
			fmt.Println("Sem captura:", closure2())
		case 4:
			fmt.Println("\nSlices:")
			fmt.Println("Sem pré-alocação:", exemploSliceRuim(5))
			fmt.Println("Com pré-alocação:", exemploSliceBom(5))
		case 5:
			fmt.Println("\nStructs:")
			p1 := exemploPointerStruct()
			p2 := exemploValorStruct()
			fmt.Printf("Pointer: %+v\n", p1)
			fmt.Printf("Valor: %+v\n", p2)
		case 6:
			fmt.Println("\nConcatenação:")
			parts := []string{"Hello", " ", "World"}
			fmt.Println("Ruim:", exemploConcatRuim(parts))
			fmt.Println("Bom:", exemploConcatBom(parts))
		case 7:
			fmt.Println("\nExecutando goroutines...")
			exemploGoroutine()
			exemploGoroutineBom()
		case 8:
			demonstrarEscapeAnalysis()
		case 9:
			compararVersoes()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}


