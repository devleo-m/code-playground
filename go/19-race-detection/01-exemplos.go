package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

// ============================================
// EXEMPLO 1: Race Condition (PROBLEMA)
// ============================================

var contadorSemSync int

func incrementarSemSync() {
	contadorSemSync++ // Race condition!
}

func exemplo1_RaceCondition() {
	var wg sync.WaitGroup
	
	fmt.Println("=== Exemplo 1: Race Condition ===")
	fmt.Println("Execute com: go run -race 01-exemplos.go exemplo1")
	fmt.Println()
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementarSemSync()
		}()
	}
	
	wg.Wait()
	fmt.Printf("Contador esperado: 1000\n")
	fmt.Printf("Contador obtido: %d\n", contadorSemSync)
	fmt.Println("⚠️  Execute com -race para ver o warning!")
}

// ============================================
// EXEMPLO 2: Solução com Mutex
// ============================================

var (
	contadorComMutex int
	mu               sync.Mutex
)

func incrementarComMutex() {
	mu.Lock()
	defer mu.Unlock()
	contadorComMutex++
}

func exemplo2_ComMutex() {
	var wg sync.WaitGroup
	
	fmt.Println("=== Exemplo 2: Solução com Mutex ===")
	fmt.Println()
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementarComMutex()
		}()
	}
	
	wg.Wait()
	fmt.Printf("Contador esperado: 1000\n")
	fmt.Printf("Contador obtido: %d\n", contadorComMutex)
	fmt.Println("✅ Sem race conditions!")
}

// ============================================
// EXEMPLO 3: Solução com Atomic
// ============================================

var contadorAtomic int64

func incrementarAtomic() {
	atomic.AddInt64(&contadorAtomic, 1)
}

func exemplo3_ComAtomic() {
	var wg sync.WaitGroup
	
	fmt.Println("=== Exemplo 3: Solução com Atomic ===")
	fmt.Println()
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementarAtomic()
		}()
	}
	
	wg.Wait()
	fmt.Printf("Contador esperado: 1000\n")
	fmt.Printf("Contador obtido: %d\n", atomic.LoadInt64(&contadorAtomic))
	fmt.Println("✅ Sem race conditions!")
}

// ============================================
// EXEMPLO 4: Cache Thread-Safe
// ============================================

type Cache struct {
	data map[string]int
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]int),
	}
}

func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	valor, existe := c.data[key]
	return valor, existe
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func exemplo4_CacheThreadSafe() {
	cache := NewCache()
	var wg sync.WaitGroup
	
	fmt.Println("=== Exemplo 4: Cache Thread-Safe ===")
	fmt.Println()
	
	// Múltiplas goroutines escrevendo
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				cache.Set(fmt.Sprintf("key-%d", id), j)
			}
		}(i)
	}
	
	// Múltiplas goroutines lendo
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_, _ = cache.Get(fmt.Sprintf("key-%d", id))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println("✅ Cache operations completed sem race conditions!")
}

// ============================================
// EXEMPLO 5: Race Condition Sutil (Leitura durante Escrita)
// ============================================

type Contador struct {
	valor int
	mu    sync.Mutex
}

func (c *Contador) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

// ❌ PROBLEMA: Leitura sem mutex durante escrita
func (c *Contador) Valor() int {
	// Esqueceu de usar o mutex!
	return c.valor // Race condition se houver escrita simultânea
}

// ✅ CORRETO: Leitura com mutex
func (c *Contador) ValorSeguro() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valor
}

func exemplo5_RaceConditionSutil() {
	contador := &Contador{}
	var wg sync.WaitGroup
	
	fmt.Println("=== Exemplo 5: Race Condition Sutil ===")
	fmt.Println("⚠️  Leitura sem mutex durante escrita causa race condition!")
	fmt.Println()
	
	// Goroutines escrevendo
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				contador.Incrementar()
			}
		}()
	}
	
	// Goroutine lendo (sem mutex)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			_ = contador.Valor() // Race condition!
		}
	}()
	
	wg.Wait()
	fmt.Printf("Contador final: %d\n", contador.ValorSeguro())
	fmt.Println("⚠️  Execute com -race para ver o warning no método Valor()!")
}

// ============================================
// MAIN: Execute os exemplos
// ============================================

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run 01-exemplos.go [exemplo1|exemplo2|exemplo3|exemplo4|exemplo5]")
		fmt.Println()
		fmt.Println("Exemplos disponíveis:")
		fmt.Println("  exemplo1 - Race Condition (PROBLEMA)")
		fmt.Println("  exemplo2 - Solução com Mutex")
		fmt.Println("  exemplo3 - Solução com Atomic")
		fmt.Println("  exemplo4 - Cache Thread-Safe")
		fmt.Println("  exemplo5 - Race Condition Sutil")
		fmt.Println()
		fmt.Println("💡 Dica: Execute com -race para ver os warnings:")
		fmt.Println("   go run -race 01-exemplos.go exemplo1")
		return
	}
	
	exemplo := os.Args[1]
	
	switch exemplo {
	case "exemplo1":
		exemplo1_RaceCondition()
	case "exemplo2":
		exemplo2_ComMutex()
	case "exemplo3":
		exemplo3_ComAtomic()
	case "exemplo4":
		exemplo4_CacheThreadSafe()
	case "exemplo5":
		exemplo5_RaceConditionSutil()
	default:
		fmt.Printf("Exemplo '%s' não encontrado!\n", exemplo)
		fmt.Println("Use: exemplo1, exemplo2, exemplo3, exemplo4 ou exemplo5")
	}
}

