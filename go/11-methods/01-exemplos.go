package main

import (
	"fmt"
	"math"
)

// ============================================
// 1. METHODS VS FUNCTIONS
// ============================================

type Retangulo struct {
	Largura float64
	Altura  float64
}

// Method: pertence ao tipo Retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Function: independente
func AreaRetangulo(largura, altura float64) float64 {
	return largura * altura
}

// ============================================
// 2. VALUE RECEIVERS
// ============================================

type Ponto struct {
	X, Y float64
}

// Value receiver: não modifica, apenas calcula
func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Value receiver: retorna um novo ponto (imutável)
func (p Ponto) Mover(dx, dy float64) Ponto {
	return Ponto{X: p.X + dx, Y: p.Y + dy}
}

type Contador struct {
	valor int
}

// Value receiver: não modifica o original
func (c Contador) Valor() int {
	return c.valor
}

// ============================================
// 3. POINTER RECEIVERS
// ============================================

// Pointer receiver: modifica o estado
func (c *Contador) Incrementar() {
	c.valor++
}

type ContaBancaria struct {
	titular string
	saldo   float64
}

// Pointer receiver: modifica o estado
func (c *ContaBancaria) Depositar(valor float64) {
	c.saldo += valor
}

// Pointer receiver: modifica o estado
func (c *ContaBancaria) Sacar(valor float64) bool {
	if valor <= c.saldo {
		c.saldo -= valor
		return true
	}
	return false
}

// Pointer receiver: apenas lê, mas mantém consistência
func (c *ContaBancaria) Saldo() float64 {
	return c.saldo
}

// ============================================
// 4. METHODS COM PARÂMETROS E RETORNOS
// ============================================

type Calculadora struct {
	historico []float64
}

// Method com parâmetros e retorno
func (c *Calculadora) Somar(a, b float64) float64 {
	resultado := a + b
	c.historico = append(c.historico, resultado)
	return resultado
}

// Method com múltiplos retornos
func (c *Calculadora) Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	resultado := a / b
	c.historico = append(c.historico, resultado)
	return resultado, nil
}

// Method sem parâmetros, com retorno
func (c *Calculadora) Historico() []float64 {
	return c.historico
}

// ============================================
// 5. METHODS EM TIPOS NÃO-STRUCT
// ============================================

// Definindo um novo tipo baseado em int
type Dinheiro int

// Method para o tipo Dinheiro
func (d Dinheiro) Reais() string {
	return fmt.Sprintf("R$ %.2f", float64(d)/100)
}

// Method para o tipo Dinheiro
func (d Dinheiro) Centavos() int {
	return int(d)
}

// Tipo customizado baseado em slice
type ListaNumeros []int

// Method para adicionar número (precisa de pointer para modificar)
func (ln *ListaNumeros) Adicionar(num int) {
	*ln = append(*ln, num)
}

// Method para calcular média (value receiver é suficiente)
func (ln ListaNumeros) Media() float64 {
	if len(ln) == 0 {
		return 0
	}
	soma := 0
	for _, num := range ln {
		soma += num
	}
	return float64(soma) / float64(len(ln))
}

// ============================================
// 6. EXEMPLO COMPLETO: PRODUTO
// ============================================

type Produto struct {
	Nome   string
	Preco  float64
	Vendas int
}

// Value receiver: não modifica
func (p Produto) PrecoComDesconto(percentual float64) float64 {
	return p.Preco * (1 - percentual/100)
}

// Pointer receiver: modifica o estado
func (p *Produto) RegistrarVenda() {
	p.Vendas++
}

// Pointer receiver: modifica o estado
func (p *Produto) AtualizarPreco(novoPreco float64) {
	p.Preco = novoPreco
}

// ============================================
// MAIN: EXEMPLOS DE USO
// ============================================

func main() {
	fmt.Println("=== 1. METHODS VS FUNCTIONS ===")
	ret := Retangulo{Largura: 10, Altura: 5}
	fmt.Printf("Method: ret.Area() = %.2f\n", ret.Area())
	fmt.Printf("Function: AreaRetangulo(10, 5) = %.2f\n", AreaRetangulo(10, 5))
	fmt.Println()

	fmt.Println("=== 2. VALUE RECEIVERS ===")
	p1 := Ponto{X: 3, Y: 4}
	fmt.Printf("Distância da origem: %.2f\n", p1.DistanciaOrigem())
	
	p2 := p1.Mover(1, 1)
	fmt.Printf("p1: %+v (não mudou)\n", p1)
	fmt.Printf("p2: %+v (novo ponto)\n", p2)
	fmt.Println()

	fmt.Println("=== 3. POINTER RECEIVERS ===")
	contador := Contador{valor: 0}
	fmt.Printf("Antes: %d\n", contador.Valor())
	contador.Incrementar()
	fmt.Printf("Depois: %d\n", contador.Valor())
	
	conta := ContaBancaria{titular: "João", saldo: 1000}
	conta.Depositar(500)
	fmt.Printf("Saldo após depósito: %.2f\n", conta.Saldo())
	conta.Sacar(200)
	fmt.Printf("Saldo após saque: %.2f\n", conta.Saldo())
	fmt.Println()

	fmt.Println("=== 4. METHODS COM PARÂMETROS ===")
	calc := Calculadora{}
	resultado1 := calc.Somar(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", resultado1)
	
	resultado2, err := calc.Dividir(20, 4)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("20 / 4 = %.2f\n", resultado2)
	}
	fmt.Printf("Histórico: %v\n", calc.Historico())
	fmt.Println()

	fmt.Println("=== 5. METHODS EM TIPOS NÃO-STRUCT ===")
	valor := Dinheiro(1250)
	fmt.Printf("Valor: %s\n", valor.Reais())
	fmt.Printf("Centavos: %d\n", valor.Centavos())
	
	lista := ListaNumeros{}
	lista.Adicionar(10)
	lista.Adicionar(20)
	lista.Adicionar(30)
	fmt.Printf("Média: %.2f\n", lista.Media())
	fmt.Println()

	fmt.Println("=== 6. EXEMPLO COMPLETO: PRODUTO ===")
	produto := Produto{Nome: "Notebook", Preco: 2000, Vendas: 0}
	precoComDesconto := produto.PrecoComDesconto(10)
	fmt.Printf("Preço original: %.2f\n", produto.Preco)
	fmt.Printf("Preço com 10%% desconto: %.2f\n", precoComDesconto)
	
	produto.RegistrarVenda()
	fmt.Printf("Vendas: %d\n", produto.Vendas)
	
	produto.AtualizarPreco(1800)
	fmt.Printf("Novo preço: %.2f\n", produto.Preco)
}

