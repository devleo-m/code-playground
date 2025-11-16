package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// Exemplo 1: Stack Allocation
func exemploStack() int {
	x := 42 // Alocado no stack
	return x
}

// Exemplo 2: Heap Allocation (escape)
func exemploHeap() *int {
	x := 42 // Escapa para o heap
	return &x
}

// Exemplo 3: Comparação de alocações - sem pré-alocação
func processItemsRuim(items []string) []string {
	var result []string // Capacidade zero
	for _, item := range items {
		result = append(result, "Processed: "+item) // Múltiplas realocações
	}
	return result
}

// Exemplo 4: Comparação de alocações - com pré-alocação
func processItemsBom(items []string) []string {
	result := make([]string, 0, len(items)) // Pré-aloca capacidade
	for _, item := range items {
		result = append(result, "Processed: "+item) // Sem realocações
	}
	return result
}

// Exemplo 5: sync.Pool - Pool de buffers
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func processWithPool(data string) string {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	buf.Reset() // Resetar antes de usar
	buf.WriteString("Processed: ")
	buf.WriteString(data)

	return buf.String()
}

// Exemplo 6: Reutilização de slice em loop
func processBatchRuim(items []int) {
	for i := 0; i < len(items); i++ {
		buffer := make([]byte, 1024) // Nova alocação a cada iteração
		_ = buffer
	}
}

func processBatchBom(items []int) {
	buffer := make([]byte, 0, 1024) // Uma única alocação
	for i := 0; i < len(items); i++ {
		buffer = buffer[:0] // Resetar sem realocar
		_ = buffer
	}
}

// Exemplo 7: Monitoramento de GC
func printGCStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("\n=== %s ===\n", label)
	fmt.Printf("Alloc: %d KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys: %d KB\n", m.Sys/1024)
	fmt.Printf("NumGC: %d\n", m.NumGC)
	fmt.Printf("PauseTotalNs: %d ns\n", m.PauseTotalNs)
	if m.NumGC > 0 {
		fmt.Printf("Avg Pause: %d ns\n", m.PauseTotalNs/uint64(m.NumGC))
	}
}

// Exemplo 8: Comparação de performance
func benchmarkComparison() {
	items := make([]string, 1000)
	for i := range items {
		items[i] = fmt.Sprintf("item-%d", i)
	}

	// Versão ruim
	start := time.Now()
	_ = processItemsRuim(items)
	timeRuim := time.Since(start)

	// Versão boa
	start = time.Now()
	_ = processItemsBom(items)
	timeBom := time.Since(start)

	fmt.Printf("\n=== Comparação de Performance ===\n")
	fmt.Printf("Sem pré-alocação: %v\n", timeRuim)
	fmt.Printf("Com pré-alocação: %v\n", timeBom)
	fmt.Printf("Melhoria: %.2fx mais rápido\n", float64(timeRuim)/float64(timeBom))
}

// Exemplo 9: Demonstração de escape analysis
func demonstrateEscape() {
	fmt.Println("\n=== Demonstração de Escape Analysis ===")
	fmt.Println("Compile com: go build -gcflags=\"-m\" para ver decisões de escape")
	fmt.Println("\nVariáveis que escapam:")
	fmt.Println("- Retornadas como pointer")
	fmt.Println("- Armazenadas em estruturas globais")
	fmt.Println("- Compartilhadas entre goroutines")
	fmt.Println("- Muito grandes para o stack")
}

// Exemplo 10: Uso correto de sync.Pool
func demonstratePool() {
	fmt.Println("\n=== Demonstração de sync.Pool ===")

	requests := []string{"req1", "req2", "req3", "req4", "req5"}

	fmt.Println("Processando com pool:")
	for _, req := range requests {
		result := processWithPool(req)
		fmt.Printf("  %s\n", result)
	}

	fmt.Println("\nPool reutiliza buffers, reduzindo alocações!")
}

// Menu interativo
func showMenu() {
	fmt.Println("\n=== Exemplos de Memory Management ===")
	fmt.Println("1. Stack vs Heap allocation")
	fmt.Println("2. Comparação: sem vs com pré-alocação")
	fmt.Println("3. sync.Pool demonstration")
	fmt.Println("4. Reutilização de slices")
	fmt.Println("5. Monitoramento de GC")
	fmt.Println("6. Benchmark comparison")
	fmt.Println("7. Escape analysis info")
	fmt.Println("8. Executar todos os exemplos")
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
			items := []string{"a", "b", "c"}
			fmt.Println("Ruim:", processItemsRuim(items))
			fmt.Println("Bom:", processItemsBom(items))
		case "3":
			demonstratePool()
		case "4":
			items := []int{1, 2, 3, 4, 5}
			fmt.Println("Processando batch...")
			processBatchBom(items)
			fmt.Println("Concluído!")
		case "5":
			printGCStats("Inicial")
			// Criar algumas alocações
			for i := 0; i < 1000; i++ {
				_ = make([]byte, 1024)
			}
			runtime.GC()
			printGCStats("Após alocações e GC")
		case "6":
			benchmarkComparison()
		case "7":
			demonstrateEscape()
		case "8":
			runAllExamples()
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
			items := []string{"a", "b", "c"}
			fmt.Println("\nSem pré-alocação:", processItemsRuim(items))
			fmt.Println("Com pré-alocação:", processItemsBom(items))
		case 3:
			demonstratePool()
		case 4:
			items := []int{1, 2, 3, 4, 5}
			fmt.Println("\nProcessando batch com reutilização...")
			processBatchBom(items)
			fmt.Println("Concluído!")
		case 5:
			printGCStats("Inicial")
			for i := 0; i < 1000; i++ {
				_ = make([]byte, 1024)
			}
			runtime.GC()
			printGCStats("Após alocações e GC")
		case 6:
			benchmarkComparison()
		case 7:
			demonstrateEscape()
		case 8:
			runAllExamples()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}

func runAllExamples() {
	fmt.Println("\n=== Executando Todos os Exemplos ===\n")

	// 1. Stack vs Heap
	fmt.Println("1. Stack vs Heap:")
	fmt.Println("   Stack:", exemploStack())
	heapPtr := exemploHeap()
	fmt.Println("   Heap:", *heapPtr)

	// 2. Pré-alocação
	fmt.Println("\n2. Pré-alocação:")
	items := []string{"item1", "item2", "item3"}
	fmt.Println("   Sem pré-alocação:", len(processItemsRuim(items)))
	fmt.Println("   Com pré-alocação:", len(processItemsBom(items)))

	// 3. Pool
	fmt.Println("\n3. sync.Pool:")
	demonstratePool()

	// 4. Reutilização
	fmt.Println("\n4. Reutilização de slices:")
	processBatchBom([]int{1, 2, 3, 4, 5})
	fmt.Println("   Concluído!")

	// 5. GC Stats
	fmt.Println("\n5. GC Stats:")
	printGCStats("Inicial")
	for i := 0; i < 100; i++ {
		_ = make([]byte, 1024)
	}
	runtime.GC()
	printGCStats("Após alocações")

	// 6. Benchmark
	fmt.Println("\n6. Benchmark:")
	benchmarkComparison()

	// 7. Escape
	fmt.Println("\n7. Escape Analysis:")
	demonstrateEscape()

	fmt.Println("\n=== Todos os Exemplos Concluídos ===")
}

