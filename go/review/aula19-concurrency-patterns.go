package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * ============================================================================
 * AULA 19: CONCURRENCY PATTERNS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Worker pools
 * - Fan-out/Fan-in
 * - Pipeline
 * - Rate limiting
 */

func main() {
	fmt.Println("=== AULA 19: CONCURRENCY PATTERNS ===\n")
	
	// Worker Pool
	fmt.Println("1. WORKER POOL:")
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Enviar jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Coletar results
	for r := 1; r <= 5; r++ {
		fmt.Printf("   Resultado: %d\n", <-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("   Worker %d processando job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
	}
}





