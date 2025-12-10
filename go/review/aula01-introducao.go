package main

import "fmt"

/*
 * ============================================================================
 * AULA 01: INTRODUÇÃO AO GO
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - Go é uma linguagem compilada, estaticamente tipada
 * - Criada no Google por Robert Griesemer, Rob Pike e Ken Thompson
 * - Focada em simplicidade, concorrência e performance
 * - Garbage collector automático
 * - Compilação rápida
 * - Binário único executável
 */

func main() {
	fmt.Println("=== AULA 01: INTRODUÇÃO AO GO ===")

	// Exemplo básico: Olá, Mundo!
	fmt.Println("Olá, Mundo!")

	// Características principais do Go
	caracteristicas := []string{
		"Simplicidade: Sintaxe limpa e fácil de aprender",
		"Concorrência: Goroutines e channels nativos",
		"Performance: Código compilado, rápido como C++",
		"Garbage Collector: Gerencia memória automaticamente",
		"Compilação Rápida: Tempos de build muito baixos",
		"Binário Único: Executável sem dependências externas",
	}

	fmt.Println("\nCaracterísticas Principais:")
	for i, carac := range caracteristicas {
		fmt.Printf("%d. %s\n", i+1, carac)
	}

	// Onde Go se destaca
	fmt.Println("\nOnde Go se destaca:")
	fmt.Println("- Serviços de backend e microsserviços")
	fmt.Println("- Ferramentas CLI (linha de comando)")
	fmt.Println("- Sistemas distribuídos e cloud computing")
	fmt.Println("- Docker, Kubernetes, Terraform são escritos em Go")

	// Estrutura básica de um programa Go
	fmt.Println("\nEstrutura básica:")
	fmt.Println("package main    // Pacote executável")
	fmt.Println("import \"fmt\"   // Importar biblioteca")
	fmt.Println("func main() {   // Ponto de entrada")
	fmt.Println("    // código")
	fmt.Println("}")
}




