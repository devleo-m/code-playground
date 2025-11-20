package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Importa pprof automaticamente
	"runtime"
	"sync"
	"time"
)

// Exemplo 1: Programa com pprof habilitado
// Execute: go run 01-exemplos.go
// Depois: go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
func exemplo1_pprof() {
	// Inicia servidor HTTP para pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Programa rodando. Acesse http://localhost:6060/debug/pprof/")
	fmt.Println("Para coletar CPU profile: go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10")

	// Loop principal com trabalho simulado
	for {
		processData()
		time.Sleep(100 * time.Millisecond)
	}
}

func processData() {
	quickOperation()
	slowOperation()
}

func quickOperation() {
	time.Sleep(10 * time.Millisecond)
}

func slowOperation() {
	// Simula operação lenta
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	time.Sleep(50 * time.Millisecond)
}

// Exemplo 2: Vazamento de memória intencional
// Execute: go run 01-exemplos.go exemplo2
// Depois: go tool pprof http://localhost:6060/debug/pprof/heap
func exemplo2_memory_leak() {
	// Inicia servidor pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Programa rodando. Acesse http://localhost:6060/debug/pprof/heap")

	var data []byte

	// Loop que causa vazamento de memória
	for {
		leakMemory(&data)
		time.Sleep(1 * time.Second)

		// Mostrar estatísticas de memória
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc: %d KB, TotalAlloc: %d KB, Sys: %d KB\n",
			m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024)
	}
}

func leakMemory(data *[]byte) {
	// Aloca memória mas nunca libera adequadamente
	*data = make([]byte, 1024*1024) // 1 MB
	// data deveria ser liberado, mas não é
}

// Exemplo 3: Data race intencional
// Execute: go run -race 01-exemplos.go exemplo3
// O race detector vai detectar o problema!
func exemplo3_data_race() {
	var counter int
	var wg sync.WaitGroup

	// Múltiplas goroutines escrevendo na mesma variável
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // DATA RACE!
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("Execute com: go run -race 01-exemplos.go exemplo3")
}

// Exemplo 4: Data race corrigido com mutex
// Execute: go run -race 01-exemplos.go exemplo4
// Agora não há mais data race!
func exemplo4_data_race_corrigido() {
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++ // Agora protegido!
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter) // Sempre 10000
	fmt.Println("Execute com: go run -race 01-exemplos.go exemplo4")
}

// Exemplo 5: Múltiplas goroutines para análise com trace
// Execute: go run -trace=trace.out 01-exemplos.go exemplo5
// Depois: go tool trace trace.out
func exemplo5_trace() {
	var wg sync.WaitGroup

	fmt.Println("Executando programa com trace...")
	fmt.Println("Após terminar, execute: go tool trace trace.out")

	// Cria múltiplas goroutines
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Espera tempo aleatório
			time.Sleep(time.Duration(id%10) * 10 * time.Millisecond)
			// Processa dados
			sum := 0
			for j := 0; j < 10000; j++ {
				sum += j
			}
			// Espera mais um pouco
			time.Sleep(time.Duration(id%5) * 20 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("Programa terminado!")
}

// Exemplo 6: Conteção de mutex (para análise com trace)
// Execute: go run -trace=trace.out 01-exemplos.go exemplo6
// Depois: go tool trace trace.out
func exemplo6_mutex_contention() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	fmt.Println("Executando programa com contenção de mutex...")
	fmt.Println("Após terminar, execute: go tool trace trace.out")

	// Muitas goroutines competindo por um mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mu.Lock()
			// Simula trabalho
			time.Sleep(10 * time.Millisecond)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Programa terminado!")
}

func main() {
	// Descomente o exemplo que deseja executar:

	// exemplo1_pprof()
	// exemplo2_memory_leak()
	// exemplo3_data_race()
	// exemplo4_data_race_corrigido()
	// exemplo5_trace()
	// exemplo6_mutex_contention()

	fmt.Println("Descomente um dos exemplos no main() para executar!")
	fmt.Println("\nExemplos disponíveis:")
	fmt.Println("  exemplo1_pprof()           - CPU profiling com pprof")
	fmt.Println("  exemplo2_memory_leak()        - Memory leak para análise")
	fmt.Println("  exemplo3_data_race()          - Data race (execute com -race)")
	fmt.Println("  exemplo4_data_race_corrigido() - Data race corrigido")
	fmt.Println("  exemplo5_trace()             - Múltiplas goroutines para trace")
	fmt.Println("  exemplo6_mutex_contention()  - Conteção de mutex para trace")
}



