package main

import (
	"fmt"
	"math"
)

// ============================================
// 1. INTERFACES BÁSICAS
// ============================================

type Forma interface {
	Area() float64
	Perimetro() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

func ImprimirForma(f Forma) {
	fmt.Printf("Área: %.2f, Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// ============================================
// 2. EMPTY INTERFACE
// ============================================

func ImprimirQualquerCoisa(valor interface{}) {
	fmt.Printf("Tipo: %T, Valor: %v\n", valor, valor)
}

// ============================================
// 3. EMBEDDING INTERFACES
// ============================================

type Leitor interface {
	Ler() string
}

type Escritor interface {
	Escrever(conteudo string) error
}

type Arquivo interface {
	Leitor
	Escritor
}

type MeuArquivo struct {
	conteudo string
	fechado  bool
}

func (f *MeuArquivo) Ler() string {
	if f.fechado {
		return ""
	}
	return f.conteudo
}

func (f *MeuArquivo) Escrever(conteudo string) error {
	if f.fechado {
		return fmt.Errorf("arquivo fechado")
	}
	f.conteudo = conteudo
	return nil
}

// ============================================
// 4. TYPE ASSERTIONS
// ============================================

func ProcessarTipo(valor interface{}) {
	// Forma segura
	if str, ok := valor.(string); ok {
		fmt.Printf("É uma string: %s\n", str)
	} else if num, ok := valor.(int); ok {
		fmt.Printf("É um int: %d\n", num)
	} else {
		fmt.Printf("Tipo: %T\n", valor)
	}
}

// ============================================
// 5. TYPE SWITCH
// ============================================

func DescreverTipo(valor interface{}) {
	switch v := valor.(type) {
	case int:
		fmt.Printf("É um inteiro: %d\n", v)
	case string:
		fmt.Printf("É uma string: %s\n", v)
	case bool:
		fmt.Printf("É um booleano: %v\n", v)
	case float64:
		fmt.Printf("É um float64: %.2f\n", v)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}
}

// ============================================
// 6. EXEMPLO COMPLETO: SISTEMA DE PAGAMENTO
// ============================================

type Pagamento interface {
	Processar(valor float64) error
	Validar() bool
}

type CartaoCredito struct {
	Numero string
	Valido bool
}

func (c CartaoCredito) Processar(valor float64) error {
	if !c.Validar() {
		return fmt.Errorf("cartão inválido")
	}
	fmt.Printf("Processando R$ %.2f no cartão %s\n", valor, c.Numero)
	return nil
}

func (c CartaoCredito) Validar() bool {
	return c.Valido && len(c.Numero) > 0
}

type Pix struct {
	Chave string
}

func (p Pix) Processar(valor float64) error {
	fmt.Printf("Processando R$ %.2f via Pix para %s\n", valor, p.Chave)
	return nil
}

func (p Pix) Validar() bool {
	return len(p.Chave) > 0
}

func RealizarCompra(p Pagamento, valor float64) error {
	if !p.Validar() {
		return fmt.Errorf("método de pagamento inválido")
	}
	return p.Processar(valor)
}

// ============================================
// 7. EXEMPLO: ANIMAIS COM TYPE SWITCH
// ============================================

type Animal interface {
	FazerSom() string
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) FazerSom() string {
	return "Au au!"
}

type Gato struct {
	Nome string
}

func (g Gato) FazerSom() string {
	return "Miau!"
}

func Interagir(animal Animal) {
	fmt.Println(animal.FazerSom())

	switch a := animal.(type) {
	case Cachorro:
		fmt.Printf("%s está abanando o rabo!\n", a.Nome)
	case Gato:
		fmt.Printf("%s está ronronando!\n", a.Nome)
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	fmt.Println("=== 1. INTERFACES BÁSICAS ===")
	ret := Retangulo{Largura: 10, Altura: 5}
	circ := Circulo{Raio: 3}
	ImprimirForma(ret)
	ImprimirForma(circ)
	fmt.Println()

	fmt.Println("=== 2. EMPTY INTERFACE ===")
	ImprimirQualquerCoisa(42)
	ImprimirQualquerCoisa("Olá")
	ImprimirQualquerCoisa(3.14)
	ImprimirQualquerCoisa([]int{1, 2, 3})
	fmt.Println()

	fmt.Println("=== 3. EMBEDDING INTERFACES ===")
	arquivo := &MeuArquivo{}
	arquivo.Escrever("Conteúdo inicial")
	fmt.Println("Lendo:", arquivo.Ler())
	fmt.Println()

	fmt.Println("=== 4. TYPE ASSERTIONS ===")
	ProcessarTipo("hello")
	ProcessarTipo(42)
	ProcessarTipo(3.14)
	fmt.Println()

	fmt.Println("=== 5. TYPE SWITCH ===")
	DescreverTipo(42)
	DescreverTipo("hello")
	DescreverTipo(true)
	DescreverTipo(3.14)
	fmt.Println()

	fmt.Println("=== 6. SISTEMA DE PAGAMENTO ===")
	cartao := CartaoCredito{Numero: "1234-5678", Valido: true}
	pix := Pix{Chave: "usuario@email.com"}
	RealizarCompra(cartao, 100.50)
	RealizarCompra(pix, 250.00)
	fmt.Println()

	fmt.Println("=== 7. ANIMAIS ===")
	cachorro := Cachorro{Nome: "Rex"}
	gato := Gato{Nome: "Mimi"}
	Interagir(cachorro)
	Interagir(gato)
}

