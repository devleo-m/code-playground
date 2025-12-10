package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * ============================================================================
 * AULA 16: CONCURRENCY
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Goroutines: Threads leves (go funcao())
 * - Channels: Comunicação entre goroutines
 * - Unbuffered channels: Síncrono
 * - Buffered channels: Assíncrono até capacidade
 * - select: Multiplexa channels
 * - sync.Mutex: Exclusão mútua
 * - sync.WaitGroup: Esperar goroutines
 */

func main() {
	fmt.Println("=== AULA 16: CONCURRENCY ===\n")
	
	// ===== GOROUTINES =====
	fmt.Println("1. GOROUTINES:")
	go dizerOla()
	time.Sleep(100 * time.Millisecond)
	
	// ===== CHANNELS UNBUFFERED =====
	fmt.Println("\n2. CHANNELS UNBUFFERED (síncrono):")
	ch := make(chan string)
	
	go func() {
		ch <- "Olá do channel!"
	}()
	
	mensagem := <-ch
	fmt.Printf("   Recebido: %s\n", mensagem)
	
	// ===== CHANNELS BUFFERED =====
	fmt.Println("\n3. CHANNELS BUFFERED:")
	ch2 := make(chan int, 3)
	
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	
	fmt.Printf("   Enviados 3 valores (capacity: 3)\n")
	fmt.Printf("   Recebido: %d\n", <-ch2)
	fmt.Printf("   Recebido: %d\n", <-ch2)
	fmt.Printf("   Recebido: %d\n", <-ch2)
	
	// ===== SELECT =====
	fmt.Println("\n4. SELECT:")
	ch3 := make(chan string)
	ch4 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch3 <- "Channel 1"
	}()
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch4 <- "Channel 2"
	}()
	
	select {
	case msg := <-ch3:
		fmt.Printf("   Recebido de ch3: %s\n", msg)
	case msg := <-ch4:
		fmt.Printf("   Recebido de ch4: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("   Timeout")
	}
	
	// ===== SYNC.MUTEX =====
	fmt.Println("\n5. SYNC.MUTEX:")
	var mu sync.Mutex
	contador := 0
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			contador++
			mu.Unlock()
		}()
	}
	
	wg.Wait()
	fmt.Printf("   Contador final: %d\n", contador)
	
	// ===== SYNC.WAITGROUP =====
	fmt.Println("\n6. SYNC.WAITGROUP:")
	var wg2 sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			fmt.Printf("   Goroutine %d executando\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	wg2.Wait()
	fmt.Println("   Todas as goroutines terminaram")
	
	// ===== WORKER POOL =====
	fmt.Println("\n7. WORKER POOL:")
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	// Workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Results
	for r := 1; r <= 5; r++ {
		fmt.Printf("   Resultado: %d\n", <-results)
	}
}

func dizerOla() {
	fmt.Println("   Olá do goroutine!")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("   Worker %d processando job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
	}
}





