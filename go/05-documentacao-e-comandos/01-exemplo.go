package main

import (
	"fmt"
	"math"
)

// Calculadora realiza operações matemáticas básicas.
// Este é um exemplo de documentação de pacote.
type Calculadora struct {
	historico []float64
}

// NovaCalculadora cria uma nova instância de Calculadora.
// Retorna um ponteiro para a calculadora inicializada.
func NovaCalculadora() *Calculadora {
	return &Calculadora{
		historico: make([]float64, 0),
	}
}

// Soma adiciona dois números e retorna o resultado.
//
// Parâmetros:
//   - a: primeiro número
//   - b: segundo número
//
// Retorna:
//   - float64: resultado da soma
//
// Exemplo:
//
//	resultado := calc.Soma(5, 3)
//	// resultado será 8
func (c *Calculadora) Soma(a, b float64) float64 {
	resultado := a + b
	c.historico = append(c.historico, resultado)
	return resultado
}

// Subtracao subtrai o segundo número do primeiro.
func (c *Calculadora) Subtracao(a, b float64) float64 {
	resultado := a - b
	c.historico = append(c.historico, resultado)
	return resultado
}

// Multiplicacao multiplica dois números.
func (c *Calculadora) Multiplicacao(a, b float64) float64 {
	resultado := a * b
	c.historico = append(c.historico, resultado)
	return resultado
}

// Divisao divide o primeiro número pelo segundo.
//
// Retorna um erro se b for zero.
func (c *Calculadora) Divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não permitida")
	}
	resultado := a / b
	c.historico = append(c.historico, resultado)
	return resultado, nil
}

// Potencia calcula a potência de a elevado a b.
func (c *Calculadora) Potencia(a, b float64) float64 {
	resultado := math.Pow(a, b)
	c.historico = append(c.historico, resultado)
	return resultado
}

// Historico retorna o histórico de todas as operações realizadas.
func (c *Calculadora) Historico() []float64 {
	return c.historico
}

// LimparHistorico remove todos os valores do histórico.
func (c *Calculadora) LimparHistorico() {
	c.historico = c.historico[:0]
}

// Exemplo de função não documentada (será mostrado como exemplo ruim)
func funcaoNaoDocumentada(x int) int {
	return x * 2
}

func main() {
	calc := NovaCalculadora()

	fmt.Println("=== Calculadora ===")
	fmt.Printf("5 + 3 = %.2f\n", calc.Soma(5, 3))
	fmt.Printf("10 - 4 = %.2f\n", calc.Subtracao(10, 4))
	fmt.Printf("6 * 7 = %.2f\n", calc.Multiplicacao(6, 7))

	resultado, err := calc.Divisao(15, 3)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", resultado)
	}

	fmt.Printf("2 ^ 8 = %.2f\n", calc.Potencia(2, 8))

	fmt.Println("\nHistórico de operações:")
	for i, valor := range calc.Historico() {
		fmt.Printf("  %d: %.2f\n", i+1, valor)
	}
}

