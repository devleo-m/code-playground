package main

import (
	"flag"
	"fmt"
)

/*
 * ============================================================================
 * AULA 43: CLI BUILDING
 * ============================================================================
 *
 * RESUMO DOS CONCEITOS:
 * - flag package
 * - Cobra (biblioteca popular)
 * - Argumentos
 * - Output formatting
 */

func main() {
	fmt.Println("=== AULA 43: CLI BUILDING ===\n")

	// Exemplo com flag package
	nome := flag.String("nome", "Mundo", "Nome para cumprimentar")
	idade := flag.Int("idade", 0, "Idade")
	flag.Parse()

	fmt.Printf("Olá, %s!\n", *nome)
	if *idade > 0 {
		fmt.Printf("Você tem %d anos\n", *idade)
	}

	fmt.Println("\nBibliotecas populares:")
	fmt.Println("- flag: Padrão do Go")
	fmt.Println("- Cobra: Framework completo")
	fmt.Println("- urfave/cli: Alternativa popular")
}




