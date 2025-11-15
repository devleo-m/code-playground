package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ============================================
// Exemplo 1: Context Básico - Background e TODO
// ============================================
func exemploContextBasico() {
	fmt.Println("\n=== Exemplo 1: Context Básico ===")

	// context.Background() - contexto raiz, nunca cancelado
	ctx := context.Background()
	fmt.Printf("Context Background: %v\n", ctx)

	// context.TODO() - quando não tem certeza qual contexto usar
	ctxTODO := context.TODO()
	fmt.Printf("Context TODO: %v\n", ctxTODO)
}

// ============================================
// Exemplo 2: Context com Timeout
// ============================================
func exemploContextTimeout() {
	fmt.Println("\n=== Exemplo 2: Context com Timeout ===")

	// Criar contexto com timeout de 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Sempre cancelar para liberar recursos

	// Simular operação longa
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second) // Simula operação de 3 segundos
		ch <- "Operação concluída"
	}()

	select {
	case resultado := <-ch:
		fmt.Printf("Resultado: %s\n", resultado)
	case <-ctx.Done():
		fmt.Printf("Timeout! Operação cancelada: %v\n", ctx.Err())
	}
}

// ============================================
// Exemplo 3: Context com Deadline
// ============================================
func exemploContextDeadline() {
	fmt.Println("\n=== Exemplo 3: Context com Deadline ===")

	// Criar contexto com deadline (tempo absoluto)
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Verificar se já passou do deadline
	if ctx.Err() != nil {
		fmt.Printf("Context já cancelado: %v\n", ctx.Err())
		return
	}

	// Simular trabalho
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Trabalho %d concluído\n", i+1)
		case <-ctx.Done():
			fmt.Printf("Deadline atingido! Cancelando: %v\n", ctx.Err())
			return
		}
	}
}

// ============================================
// Exemplo 4: Context com Cancelamento Manual
// ============================================
func exemploContextCancel() {
	fmt.Println("\n=== Exemplo 4: Context com Cancelamento Manual ===")

	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine que faz trabalho
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Goroutine cancelada: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("Trabalhando... %d\n", i+1)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Cancelar após 1 segundo
	time.Sleep(1 * time.Second)
	fmt.Println("Cancelando contexto...")
	cancel()

	time.Sleep(500 * time.Millisecond)
}

// ============================================
// Exemplo 5: Context com Valores (Request-scoped)
// ============================================
func exemploContextValores() {
	fmt.Println("\n=== Exemplo 5: Context com Valores ===")

	// Adicionar valores ao contexto
	ctx := context.WithValue(context.Background(), "userID", "12345")
	ctx = context.WithValue(ctx, "requestID", "req-abc-xyz")

	// Função que usa valores do contexto
	processarRequisicao := func(ctx context.Context) {
		userID := ctx.Value("userID")
		requestID := ctx.Value("requestID")

		fmt.Printf("Processando requisição - UserID: %v, RequestID: %v\n", userID, requestID)
	}

	processarRequisicao(ctx)
}

// ============================================
// Exemplo 6: Context em HTTP Request
// ============================================
func exemploContextHTTP() {
	fmt.Println("\n=== Exemplo 6: Context em HTTP Request ===")

	// Handler HTTP que usa context
	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // Context vem automaticamente do request

		// Verificar se request foi cancelado
		select {
		case <-ctx.Done():
			fmt.Printf("Request cancelado: %v\n", ctx.Err())
			return
		default:
			// Processar request normalmente
			fmt.Println("Processando HTTP request...")
		}
	}

	// Simular (não vamos iniciar servidor real)
	fmt.Println("Handler criado com suporte a context")
	_ = handler
}

// ============================================
// Exemplo 7: Context com Timeout em Operação de Banco
// ============================================
func exemploContextBancoDados() {
	fmt.Println("\n=== Exemplo 7: Context com Timeout em Operação de Banco ===")

	// Simular operação de banco de dados
	consultarBanco := func(ctx context.Context, query string) (string, error) {
		// Verificar se contexto foi cancelado
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		// Simular consulta demorada
		time.Sleep(3 * time.Second)

		// Verificar novamente antes de retornar
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			return "Resultado da query: " + query, nil
		}
	}

	// Criar contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resultado, err := consultarBanco(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("Erro na consulta: %v\n", err)
	} else {
		fmt.Printf("Sucesso: %s\n", resultado)
	}
}

// ============================================
// Exemplo 8: Context Aninhado (Context Chain)
// ============================================
func exemploContextAninhado() {
	fmt.Println("\n=== Exemplo 8: Context Aninhado ===")

	// Context raiz
	ctx := context.Background()

	// Adicionar timeout
	ctx, cancel1 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel1()

	// Adicionar valores
	ctx = context.WithValue(ctx, "userID", "999")

	// Adicionar outro timeout (mais restritivo)
	ctx, cancel2 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel2()

	fmt.Printf("Context aninhado criado com timeout de 2s e userID: %v\n", ctx.Value("userID"))

	// Simular operação
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operação concluída")
	case <-ctx.Done():
		fmt.Printf("Context cancelado: %v\n", ctx.Err())
	}
}

// ============================================
// Exemplo 9: Verificando Status do Context
// ============================================
func exemploVerificarContext() {
	fmt.Println("\n=== Exemplo 9: Verificando Status do Context ===")

	ctx, cancel := context.WithCancel(context.Background())

	// Verificar se está cancelado
	fmt.Printf("Context cancelado? %v\n", ctx.Err() == nil)

	// Cancelar
	cancel()

	// Verificar novamente
	fmt.Printf("Context cancelado? %v\n", ctx.Err() != nil)
	fmt.Printf("Erro do context: %v\n", ctx.Err())

	// Verificar deadline
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()

	deadline, ok := ctx2.Deadline()
	if ok {
		fmt.Printf("Deadline: %v\n", deadline)
	} else {
		fmt.Println("Sem deadline definido")
	}
}

// ============================================
// Exemplo 10: Context em Worker Pool
// ============================================
func exemploContextWorkerPool() {
	fmt.Println("\n=== Exemplo 10: Context em Worker Pool ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Worker
	worker := func(id int) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d cancelado\n", id)
				return
			case job := <-jobs:
				fmt.Printf("Worker %d processando job %d\n", id, job)
				time.Sleep(500 * time.Millisecond)
				results <- job * 2
			}
		}
	}

	// Iniciar workers
	for i := 1; i <= 2; i++ {
		go worker(i)
	}

	// Enviar jobs
	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
		close(jobs)
	}()

	// Coletar resultados
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Timeout! Processamento cancelado\n")
			return
		case result := <-results:
			fmt.Printf("Resultado: %d\n", result)
		}
	}
}

// ============================================
// Função Main
// ============================================
func main() {
	fmt.Println("=== Exemplos de Context Package em Go ===")

	exemploContextBasico()
	exemploContextTimeout()
	exemploContextDeadline()
	exemploContextCancel()
	exemploContextValores()
	exemploContextHTTP()
	exemploContextBancoDados()
	exemploContextAninhado()
	exemploVerificarContext()
	exemploContextWorkerPool()

	fmt.Println("\n=== Fim dos Exemplos ===")
}

