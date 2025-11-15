package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Exemplo 1: Goroutine Básica
// ============================================
func exemploGoroutineBasica() {
	fmt.Println("\n=== Exemplo 1: Goroutine Básica ===")

	go func() {
		fmt.Println("Executando em uma goroutine")
	}()

	fmt.Println("Executando na função main")
	time.Sleep(100 * time.Millisecond)
}

// ============================================
// Exemplo 2: Múltiplas Goroutines
// ============================================
func exemploMultiplasGoroutines() {
	fmt.Println("\n=== Exemplo 2: Múltiplas Goroutines ===")

	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d executando\n", id)
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
}

// ============================================
// Exemplo 3: Channel Básico
// ============================================
func exemploChannelBasico() {
	fmt.Println("\n=== Exemplo 3: Channel Básico ===")

	ch := make(chan string)

	go func() {
		ch <- "Olá do channel!"
	}()

	mensagem := <-ch
	fmt.Println(mensagem)
}

// ============================================
// Exemplo 4: Channel com Múltiplos Valores
// ============================================
func exemploChannelMultiplosValores() {
	fmt.Println("\n=== Exemplo 4: Channel com Múltiplos Valores ===")

	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for valor := range ch {
		fmt.Printf("Recebido: %d\n", valor)
	}
}

// ============================================
// Exemplo 5: Unbuffered Channel (Síncrono)
// ============================================
func exemploUnbufferedChannel() {
	fmt.Println("\n=== Exemplo 5: Unbuffered Channel ===")

	ch := make(chan int) // Unbuffered

	go func() {
		fmt.Println("Enviando 42...")
		ch <- 42
		fmt.Println("42 foi enviado!")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Recebendo...")
	valor := <-ch
	fmt.Printf("Recebido: %d\n", valor)
}

// ============================================
// Exemplo 6: Buffered Channel (Assíncrono)
// ============================================
func exemploBufferedChannel() {
	fmt.Println("\n=== Exemplo 6: Buffered Channel ===")

	ch := make(chan int, 3) // Buffered com capacidade 3

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("3 valores enviados sem bloquear")

	fmt.Printf("Recebido: %d\n", <-ch)
	fmt.Printf("Recebido: %d\n", <-ch)
	fmt.Printf("Recebido: %d\n", <-ch)
}

// ============================================
// Exemplo 7: Select Statement
// ============================================
func exemploSelect() {
	fmt.Println("\n=== Exemplo 7: Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "do channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "do channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Printf("Recebido %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Recebido %s\n", msg2)
	}
}

// ============================================
// Exemplo 8: Select com Timeout
// ============================================
func exemploSelectTimeout() {
	fmt.Println("\n=== Exemplo 8: Select com Timeout ===")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "resultado"
	}()

	select {
	case resultado := <-ch:
		fmt.Printf("Recebido: %s\n", resultado)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Operação demorou muito")
	}
}

// ============================================
// Exemplo 9: Worker Pool Básico
// ============================================
func exemploWorkerPool() {
	fmt.Println("\n=== Exemplo 9: Worker Pool ===")

	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Criar workers
	for w := 1; w <= numWorkers; w++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("Worker %d processando job %d\n", id, job)
				time.Sleep(200 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Enviar jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Coletar resultados
	for r := 1; r <= numJobs; r++ {
		fmt.Printf("Resultado: %d\n", <-results)
	}
}

// ============================================
// Exemplo 10: Mutex para Proteção
// ============================================
func exemploMutex() {
	fmt.Println("\n=== Exemplo 10: Mutex ===")

	var (
		contador int
		mu       sync.Mutex
	)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			contador++
		}()
	}

	wg.Wait()
	fmt.Printf("Contador final: %d\n", contador)
}

// ============================================
// Exemplo 11: RWMutex (Múltiplos Leitores)
// ============================================
func exemploRWMutex() {
	fmt.Println("\n=== Exemplo 11: RWMutex ===")

	type Cache struct {
		mu   sync.RWMutex
		data map[string]string
	}

	cache := &Cache{
		data: make(map[string]string),
	}

	var wg sync.WaitGroup

	// Escritor
	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.mu.Lock()
		defer cache.mu.Unlock()
		cache.data["chave"] = "valor"
		fmt.Println("Valor escrito")
	}()

	// Múltiplos leitores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cache.mu.RLock()
			defer cache.mu.RUnlock()
			valor := cache.data["chave"]
			fmt.Printf("Leitor %d leu: %s\n", id, valor)
		}(i)
	}

	wg.Wait()
}

// ============================================
// Exemplo 12: WaitGroup
// ============================================
func exemploWaitGroup() {
	fmt.Println("\n=== Exemplo 12: WaitGroup ===")

	var wg sync.WaitGroup

	worker := func(id int) {
		defer wg.Done()
		fmt.Printf("Worker %d começando\n", id)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Worker %d terminando\n", id)
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	fmt.Println("Todos os workers terminaram")
}

// ============================================
// Exemplo 13: Contador Thread-Safe
// ============================================
func exemploContadorThreadSafe() {
	fmt.Println("\n=== Exemplo 13: Contador Thread-Safe ===")

	type ContadorSeguro struct {
		mu    sync.Mutex
		valor int
	}

	contador := &ContadorSeguro{}
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.mu.Lock()
			defer contador.mu.Unlock()
			contador.valor++
		}()
	}

	wg.Wait()
	fmt.Printf("Valor final: %d\n", contador.valor)
}

// ============================================
// Exemplo 14: Pipeline
// ============================================
func exemploPipeline() {
	fmt.Println("\n=== Exemplo 14: Pipeline ===")

	gerar := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	quadrado := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Pipeline: gerar -> quadrado -> imprimir
	for n := range quadrado(gerar(1, 2, 3, 4, 5)) {
		fmt.Printf("Quadrado: %d\n", n)
	}
}

// ============================================
// Exemplo 15: Select com Default (Non-Blocking)
// ============================================
func exemploSelectDefault() {
	fmt.Println("\n=== Exemplo 15: Select com Default ===")

	ch := make(chan int)

	select {
	case valor := <-ch:
		fmt.Printf("Recebido: %d\n", valor)
	default:
		fmt.Println("Nenhum valor disponível, não bloqueando")
	}

	fmt.Println("Continuando...")
}

// ============================================
// Main: Executar todos os exemplos
// ============================================
func main() {
	exemploGoroutineBasica()
	exemploMultiplasGoroutines()
	exemploChannelBasico()
	exemploChannelMultiplosValores()
	exemploUnbufferedChannel()
	exemploBufferedChannel()
	exemploSelect()
	exemploSelectTimeout()
	exemploWorkerPool()
	exemploMutex()
	exemploRWMutex()
	exemploWaitGroup()
	exemploContadorThreadSafe()
	exemploPipeline()
	exemploSelectDefault()

	fmt.Println("\n=== Todos os exemplos foram executados ===")
}
