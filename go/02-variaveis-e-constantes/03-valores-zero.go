package main

import "fmt"

// Esta struct será usada para demonstrar o valor zero de um tipo composto.
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	// Declarando variáveis de diferentes tipos sem um valor inicial.
	// Go automaticamente as inicializa com seus respectivos "valores zero".

	var i int
	var f float64
	var b bool
	var s string
	var p *int // Um ponteiro para um inteiro
	var sl []int // Um slice de inteiros
	var m map[string]int // Um mapa de string para inteiro
	var pessoa Pessoa // Uma struct

	// A função Printf nos permite formatar a saída.
	// %v é um formatador "verbo" padrão que exibe o valor.
	// %T exibe o tipo da variável.
	// \n pula uma linha.
	fmt.Printf("Tipo: %T, Valor: %v\n", i, i)
	fmt.Printf("Tipo: %T, Valor: %v\n", f, f)
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)
	fmt.Printf("Tipo: %T, Valor: '%v'\n", s, s) // Aspas para mostrar a string vazia
	fmt.Printf("Tipo: %T, Valor: %v\n", p, p)
	fmt.Printf("Tipo: %T, Valor: %v, É nil? %t\n", sl, sl, sl == nil)
	fmt.Printf("Tipo: %T, Valor: %v, É nil? %t\n", m, m, m == nil)
	fmt.Printf("Tipo: %T, Valor: %+v\n", pessoa, pessoa) // %+v mostra os nomes dos campos da struct
}
