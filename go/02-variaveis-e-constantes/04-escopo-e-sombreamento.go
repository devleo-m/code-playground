package main

import "fmt"

// `nivel` tem escopo de PACOTE.
// É visível para todas as funções neste pacote.
var nivel = "Pacote"

func main() {
	// `nivel` aqui se refere à variável de escopo de pacote.
	fmt.Printf("Nível 1: %s\n", nivel)

	// `nivel` declarado aqui tem escopo de FUNÇÃO.
	// Ele "sombra" (shadows) a variável de escopo de pacote.
	nivel := "Função (main)"
	fmt.Printf("Nível 2: %s\n", nivel)

	if true {
		// Este `nivel` tem escopo de BLOCO (if).
		// Ele sombra a variável de escopo de função.
		nivel := "Bloco (if)"
		fmt.Printf("Nível 3: %s\n", nivel)

		// Não há mais nada sombreando, então esta alteração
		// afeta a variável do escopo mais próximo: o do bloco `if`.
		nivel = "Bloco (if) - Alterado"
		fmt.Printf("Nível 4: %s\n", nivel)
	}

	// De volta ao escopo da função, a variável do bloco `if` não existe mais.
	// `nivel` aqui é a variável de escopo de função novamente.
	fmt.Printf("Nível 5: %s\n", nivel)

	imprimirNivel()
}

func imprimirNivel() {
	// Esta função não tem sua própria variável `nivel`.
	// Portanto, ela "enxerga" a variável do escopo mais próximo,
	// que é o escopo de PACOTE.
	fmt.Printf("Nível 6 (outra func): %s\n", nivel)
}
