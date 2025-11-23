package main

import (
	"errors"
	"fmt"
)

/*
 * ============================================================================
 * AULA 15: ERROR HANDLING
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Erros são valores (não exceções)
 * - Padrão idiomático: (resultado, error)
 * - Sempre verifique erros
 * - fmt.Errorf: Criar erros formatados
 * - errors.New: Criar erros simples
 * - errors.Is: Verificar tipo de erro
 * - errors.As: Extrair erro específico
 * - Wrapping: Preservar erro original
 */

var ErrDivisaoPorZero = errors.New("divisão por zero")
var ErrNumeroNegativo = errors.New("número negativo")

// Função que retorna erro
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil
}

// Função com erro customizado
func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("raiz quadrada de número negativo: %f", x)
	}
	// Simulação (use math.Sqrt em produção)
	return x / 2, nil
}

// Função com wrapping de erro
func processarNumero(x float64) (float64, error) {
	resultado, err := raizQuadrada(x)
	if err != nil {
		return 0, fmt.Errorf("processarNumero: %w", err)
	}
	return resultado, nil
}

func main() {
	fmt.Println("=== AULA 15: ERROR HANDLING ===\n")
	
	// ===== PADRÃO IDIOMÁTICO =====
	fmt.Println("1. PADRÃO IDIOMÁTICO (resultado, error):")
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Printf("   Erro: %v\n", err)
	} else {
		fmt.Printf("   Resultado: %.2f\n", resultado)
	}
	
	// ===== VERIFICAÇÃO DE ERRO =====
	fmt.Println("\n2. VERIFICAÇÃO DE ERRO:")
	resultado2, err2 := dividir(10, 0)
	if err2 != nil {
		fmt.Printf("   Erro: %v\n", err2)
	} else {
		fmt.Printf("   Resultado: %.2f\n", resultado2)
	}
	
	// ===== ERRORS.IS =====
	fmt.Println("\n3. ERRORS.IS (verificar tipo):")
	_, err3 := dividir(10, 0)
	if errors.Is(err3, ErrDivisaoPorZero) {
		fmt.Println("   Erro é do tipo ErrDivisaoPorZero")
	}
	
	// ===== ERRORS.AS =====
	fmt.Println("\n4. ERRORS.AS (extrair erro específico):")
	var divErr error
	if errors.As(err3, &divErr) {
		fmt.Printf("   Erro extraído: %v\n", divErr)
	}
	
	// ===== WRAPPING =====
	fmt.Println("\n5. WRAPPING (preservar erro original):")
	_, err4 := processarNumero(-5)
	if err4 != nil {
		fmt.Printf("   Erro wrapped: %v\n", err4)
		
		// Unwrap
		erroOriginal := errors.Unwrap(err4)
		if erroOriginal != nil {
			fmt.Printf("   Erro original: %v\n", erroOriginal)
		}
	}
	
	// ===== IGNORAR ERRO (NÃO RECOMENDADO) =====
	fmt.Println("\n6. IGNORAR ERRO (use com cuidado):")
	resultado3, _ := dividir(10, 2) // Ignora erro
	fmt.Printf("   Resultado (erro ignorado): %.2f\n", resultado3)
	
	// ===== MÚLTIPLOS ERROS =====
	fmt.Println("\n7. MÚLTIPLOS ERROS:")
	validarNumero(-5)
	validarNumero(0)
	validarNumero(10)
}

func validarNumero(x int) error {
	if x < 0 {
		return ErrNumeroNegativo
	}
	if x == 0 {
		return fmt.Errorf("número não pode ser zero")
	}
	fmt.Printf("   Número válido: %d\n", x)
	return nil
}


