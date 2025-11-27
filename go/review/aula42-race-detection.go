package main

import (
	"fmt"
	"sync"
)

/*
 * ============================================================================
 * AULA 42: RACE DETECTION
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - go test -race
 * - Detecção de race conditions
 * - Ferramenta de debugging
 */

var contador int
var mu sync.Mutex

func incrementar() {
	mu.Lock()
	contador++
	mu.Unlock()
}

func main() {
	fmt.Println("=== AULA 42: RACE DETECTION ===\n")
	fmt.Println("Detectar races: go test -race")
	fmt.Println("go run -race: Detectar em execução")
	fmt.Println("Race conditions: Acesso concorrente não sincronizado")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementar()
		}()
	}
	wg.Wait()
	fmt.Printf("Contador final: %d\n", contador)
}


