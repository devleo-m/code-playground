package main

import (
	"context"
	"fmt"
	"time"
)

/*
 * ============================================================================
 * AULA 17: CONTEXT
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Cancelamento de operações
 * - Timeouts
 * - Valores de requisição
 * - Propagação de escopo
 * - ctx.Done() para verificar cancelamento
 */

func main() {
	fmt.Println("=== AULA 17: CONTEXT ===\n")
	
	// ===== CONTEXT BACKGROUND =====
	fmt.Println("1. CONTEXT BACKGROUND:")
	ctx := context.Background()
	fmt.Printf("   ctx: %v\n", ctx)
	
	// ===== CONTEXT WITH CANCEL =====
	fmt.Println("\n2. CONTEXT WITH CANCEL:")
	ctx2, cancel := context.WithCancel(ctx)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		cancel() // Cancela após 200ms
	}()
	
	select {
	case <-ctx2.Done():
		fmt.Println("   Context cancelado")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("   Timeout")
	}
	
	// ===== CONTEXT WITH TIMEOUT =====
	fmt.Println("\n3. CONTEXT WITH TIMEOUT:")
	ctx3, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()
	
	select {
	case <-ctx3.Done():
		fmt.Println("   Context expirou (timeout)")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("   Operação completada")
	}
	
	// ===== CONTEXT WITH VALUE =====
	fmt.Println("\n4. CONTEXT WITH VALUE:")
	ctx4 := context.WithValue(context.Background(), "userID", 123)
	ctx5 := context.WithValue(ctx4, "requestID", "abc-123")
	
	userID := ctx5.Value("userID")
	requestID := ctx5.Value("requestID")
	
	fmt.Printf("   userID: %v\n", userID)
	fmt.Printf("   requestID: %v\n", requestID)
	
	// ===== VERIFICAÇÃO DE CANCELAMENTO =====
	fmt.Println("\n5. VERIFICAÇÃO DE CANCELAMENTO:")
	ctx6, cancel3 := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel3()
	}()
	
	processarComContexto(ctx6)
}

func processarComContexto(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("   Processamento cancelado: %v\n", ctx.Err())
			return
		default:
			fmt.Printf("   Processando... %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

