package main

import "fmt"

// A declaração de variáveis com `var` pode ser feita no nível do pacote.
// O valor inicial é obrigatório se o tipo não for especificado.
var versao = "1.0"

func main() {
	// --- Usando `var` ---

	// 1. Declaração explícita: declara o tipo, depois atribui o valor.
	// Útil quando o valor inicial não está disponível ainda.
	var idade int
	idade = 30
	fmt.Println("Idade (var explícito):", idade)

	// 2. Declaração com inicialização: declara e atribui na mesma linha.
	var altura float64 = 1.75
	fmt.Println("Altura (var com tipo):", altura)

	// 3. Declaração com inferência de tipo: Go adivinha o tipo pelo valor.
	// Esta forma é mais comum quando se usa `var`.
	var peso = 70.5
	fmt.Println("Peso (var com inferência):", peso)
	fmt.Println("Versão do programa:", versao)

	fmt.Println("--------------------------------")

	// --- Usando `:=` (Declaração Curta) ---
	// Só pode ser usado dentro de funções.
	// Declara e inicializa ao mesmo tempo, com inferência de tipo.
	// É a forma mais comum e preferida dentro de funções.
	nome := "Carlos"
	email := "carlos@email.com"
	estaAtivo := true

	fmt.Println("Nome:", nome)
	fmt.Println("Email:", email)
	fmt.Println("Está Ativo?:", estaAtivo)

	// Para mudar o valor de uma variável que JÁ EXISTE, use apenas o sinal de igual (=)
	nome = "Carlos Silva"
	fmt.Println("Nome atualizado:", nome)

	// O código abaixo causaria um erro de compilação, pois 'nome' já foi declarado neste escopo.
	// nome := "Carlos de novo" // Erro: no new variables on left side of :=
}
