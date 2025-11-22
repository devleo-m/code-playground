package main

import "fmt"

/*
 * ============================================================================
 * AULA 02: VARIÁVEIS E CONSTANTES
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - var: Declaração explícita (pode ser usada em qualquer lugar)
 * - := : Declaração curta (apenas dentro de funções)
 * - const: Valores imutáveis
 * - iota: Gerador de sequências numéricas
 * - Valores zero: Valores padrão para tipos não inicializados
 * - Escopo: Onde uma variável é visível
 * - Sombreamento: Variável interna esconde variável externa
 */

func main() {
	fmt.Println("=== AULA 02: VARIÁVEIS E CONSTANTES ===\n")

	// ===== DECLARAÇÃO COM var =====
	fmt.Println("1. DECLARAÇÃO COM var:")
	var idade int
	var nome string = "João"
	var altura = 1.75 // Inferência de tipo

	fmt.Printf("   idade (zero value): %d\n", idade)
	fmt.Printf("   nome: %s\n", nome)
	fmt.Printf("   altura: %.2f\n", altura)

	// ===== DECLARAÇÃO CURTA (:=) =====
	fmt.Println("\n2. DECLARAÇÃO CURTA (:=):")
	cidade := "São Paulo"
	populacao := 12000000
	temperatura := 25.5

	fmt.Printf("   cidade: %s\n", cidade)
	fmt.Printf("   populacao: %d\n", populacao)
	fmt.Printf("   temperatura: %.1f\n", temperatura)

	// ===== CONSTANTES =====
	fmt.Println("\n3. CONSTANTES:")
	const Pi = 3.14159
	const PortaServidor = 8080
	const Mensagem = "Bem-vindo!"

	fmt.Printf("   Pi: %.5f\n", Pi)
	fmt.Printf("   Porta: %d\n", PortaServidor)
	fmt.Printf("   Mensagem: %s\n", Mensagem)

	// ===== IOTA (Gerador de sequências) =====
	fmt.Println("\n4. IOTA - Sequências numéricas:")
	const (
		Domingo = iota // 0
		Segunda        // 1
		Terca          // 2
		Quarta         // 3
		Quinta         // 4
		Sexta          // 5
		Sabado         // 6
	)

	fmt.Printf("   Domingo = %d\n", Domingo)
	fmt.Printf("   Segunda = %d\n", Segunda)
	fmt.Printf("   Sabado = %d\n", Sabado)

	// IOTA com expressões (flags de bits)
	fmt.Println("\n5. IOTA com expressões (flags):")
	const (
		FlagLeitura  = 1 << iota // 1 (2^0)
		FlagEscrita              // 2 (2^1)
		FlagExecucao             // 4 (2^2)
	)

	fmt.Printf("   FlagLeitura: %d\n", FlagLeitura)
	fmt.Printf("   FlagEscrita: %d\n", FlagEscrita)
	fmt.Printf("   FlagExecucao: %d\n", FlagExecucao)

	// ===== VALORES ZERO =====
	fmt.Println("\n6. VALORES ZERO:")
	var (
		num   int
		dec   float64
		texto string
		ativo bool
	)

	fmt.Printf("   int zero: %d\n", num)
	fmt.Printf("   float64 zero: %.1f\n", dec)
	fmt.Printf("   string zero: '%s'\n", texto)
	fmt.Printf("   bool zero: %t\n", ativo)

	// ===== ESCopo E SOMBREAMENTO =====
	fmt.Println("\n7. ESCOPO E SOMBREAMENTO:")
	numero := 100
	fmt.Printf("   numero (antes do if): %d\n", numero)

	if true {
		numero := 200 // Nova variável (sombra a externa)
		fmt.Printf("   numero (dentro do if): %d\n", numero)
	}

	fmt.Printf("   numero (depois do if): %d\n", numero) // Ainda 100!
}
