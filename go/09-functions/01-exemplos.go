package main

import (
	"fmt"
	"math"
	"strings"
)

// ============================================
// 1. FUNCTION BASICS - Funções Básicas
// ============================================

// Função simples sem parâmetros e sem retorno
func cumprimentar() {
	fmt.Println("Olá, mundo!")
}

// Função com um parâmetro
func cumprimentarPessoa(nome string) {
	fmt.Printf("Olá, %s!\n", nome)
}

// Função com múltiplos parâmetros
func somar(a int, b int) int {
	return a + b
}

// Função com parâmetros do mesmo tipo (sintaxe simplificada)
func multiplicar(a, b, c int) int {
	return a * b * c
}

// Função com múltiplos tipos de retorno
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

// Função como first-class citizen - pode ser atribuída a variável
func exemploFirstClass() {
	// Atribuir função a variável
	var funcao func(int, int) int
	funcao = somar
	
	resultado := funcao(5, 3)
	fmt.Printf("Resultado: %d\n", resultado)
	
	// Passar função como argumento
	calcular(10, 20, somar)
	calcular(10, 20, multiplicar)
}

// Função que recebe outra função como parâmetro
func calcular(a, b int, operacao func(int, int) int) {
	resultado := operacao(a, b)
	fmt.Printf("Resultado da operação: %d\n", resultado)
}

// ============================================
// 2. VARIADIC FUNCTIONS - Funções Variádicas
// ============================================

// Função que aceita número variável de argumentos
func somarNumeros(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

// Função variádica com outros parâmetros (variádico deve ser o último)
func juntarStrings(separador string, textos ...string) string {
	return strings.Join(textos, separador)
}

// Passar slice para função variádica usando ...
func exemploVariadico() {
	// Chamar com múltiplos argumentos
	soma1 := somarNumeros(1, 2, 3, 4, 5)
	fmt.Printf("Soma: %d\n", soma1)
	
	// Chamar com slice usando ...
	numeros := []int{10, 20, 30}
	soma2 := somarNumeros(numeros...)
	fmt.Printf("Soma do slice: %d\n", soma2)
	
	// Chamar sem argumentos (slice vazio)
	soma3 := somarNumeros()
	fmt.Printf("Soma vazia: %d\n", soma3)
}

// ============================================
// 3. MULTIPLE RETURN VALUES - Múltiplos Retornos
// ============================================

// Padrão idiomático: retornar resultado e erro
func buscarUsuario(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("ID inválido: %d", id)
	}
	return fmt.Sprintf("Usuário-%d", id), nil
}

// Retornar múltiplos valores úteis
func calcularEstatisticas(numeros []float64) (float64, float64, float64) {
	if len(numeros) == 0 {
		return 0, 0, 0
	}
	
	var soma float64
	for _, num := range numeros {
		soma += num
	}
	media := soma / float64(len(numeros))
	
	var max float64 = numeros[0]
	var min float64 = numeros[0]
	for _, num := range numeros {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	
	return media, max, min
}

// Ignorar valores de retorno usando _
func exemploMultiplosRetornos() {
	// Receber todos os valores
	usuario, err := buscarUsuario(1)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Usuário: %s\n", usuario)
	}
	
	// Ignorar erro usando _
	usuario2, _ := buscarUsuario(2)
	fmt.Printf("Usuário: %s\n", usuario2)
	
	// Ignorar todos os valores (raro, mas possível)
	_, _, _ = calcularEstatisticas([]float64{1, 2, 3, 4, 5})
}

// ============================================
// 4. NAMED RETURN VALUES - Retornos Nomeados
// ============================================

// Retornos nomeados - inicializados com valores zero
func dividirComNome(dividendo, divisor float64) (quociente float64, resto float64, err error) {
	if divisor == 0 {
		err = fmt.Errorf("divisão por zero")
		return // Retorna os valores atuais das variáveis nomeadas
	}
	quociente = dividendo / divisor
	resto = math.Mod(dividendo, divisor)
	return // Retorna quociente, resto, err implicitamente
}

// Retornos nomeados melhoram legibilidade
func calcularRaizes(a, b, c float64) (x1, x2 float64, temSolucao bool) {
	discriminante := b*b - 4*a*c
	if discriminante < 0 {
		temSolucao = false
		return // x1 e x2 são zero, temSolucao é false
	}
	
	temSolucao = true
	sqrtDiscriminante := math.Sqrt(discriminante)
	x1 = (-b + sqrtDiscriminante) / (2 * a)
	x2 = (-b - sqrtDiscriminante) / (2 * a)
	return
}

func exemploRetornosNomeados() {
	quociente, resto, err := dividirComNome(10, 3)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Quociente: %.2f, Resto: %.2f\n", quociente, resto)
	}
	
	x1, x2, ok := calcularRaizes(1, -5, 6)
	if ok {
		fmt.Printf("Raízes: x1=%.2f, x2=%.2f\n", x1, x2)
	}
}

// ============================================
// 5. ANONYMOUS FUNCTIONS - Funções Anônimas
// ============================================

func exemploFuncoesAnonimas() {
	// Função anônima atribuída a variável
	somar := func(a, b int) int {
		return a + b
	}
	
	resultado := somar(5, 3)
	fmt.Printf("Soma: %d\n", resultado)
	
	// Função anônima executada imediatamente (IIFE - Immediately Invoked Function Expression)
	func() {
		fmt.Println("Função anônima executada imediatamente!")
	}()
	
	// Função anônima com parâmetros executada imediatamente
	func(nome string) {
		fmt.Printf("Olá, %s!\n", nome)
	}("João")
	
	// Função anônima passada como argumento
	numeros := []int{1, 2, 3, 4, 5}
	processarNumeros(numeros, func(n int) int {
		return n * 2
	})
}

// Função que recebe função anônima como callback
func processarNumeros(numeros []int, transformacao func(int) int) {
	for i, num := range numeros {
		numeros[i] = transformacao(num)
	}
	fmt.Printf("Números processados: %v\n", numeros)
}

// ============================================
// 6. CLOSURES - Closures
// ============================================

// Função que retorna uma closure
func criarContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// Closure que captura múltiplas variáveis
func criarMultiplicador(fator int) func(int) int {
	return func(numero int) int {
		return numero * fator
	}
}

// Closure com estado compartilhado
func criarAcumulador() func(int) int {
	soma := 0
	return func(valor int) int {
		soma += valor
		return soma
	}
}

func exemploClosures() {
	// Cada closure mantém seu próprio estado
	contador1 := criarContador()
	contador2 := criarContador()
	
	fmt.Printf("Contador 1: %d\n", contador1()) // 1
	fmt.Printf("Contador 1: %d\n", contador1()) // 2
	fmt.Printf("Contador 2: %d\n", contador2()) // 1 (independente)
	
	// Closure com parâmetro capturado
	dobrar := criarMultiplicador(2)
	triplicar := criarMultiplicador(3)
	
	fmt.Printf("Dobro de 5: %d\n", dobrar(5))      // 10
	fmt.Printf("Triplo de 5: %d\n", triplicar(5))  // 15
	
	// Closure com estado acumulado
	acumulador := criarAcumulador()
	fmt.Printf("Acumulado: %d\n", acumulador(10)) // 10
	fmt.Printf("Acumulado: %d\n", acumulador(20)) // 30
	fmt.Printf("Acumulado: %d\n", acumulador(30)) // 60
}

// Closure em loop - armadilha comum!
func exemploClosureArmadilha() {
	var funcoes []func() int
	
	// PROBLEMA: Todas as closures capturam a mesma variável i
	for i := 0; i < 3; i++ {
		funcoes = append(funcoes, func() int {
			return i // Captura a variável i, não o valor!
		})
	}
	
	// Todas retornam 3 (valor final de i)
	for _, f := range funcoes {
		fmt.Printf("Valor: %d\n", f()) // 3, 3, 3
	}
	
	// SOLUÇÃO: Passar valor como parâmetro
	var funcoesCorretas []func() int
	for i := 0; i < 3; i++ {
		i := i // Cria nova variável em cada iteração
		funcoesCorretas = append(funcoesCorretas, func() int {
			return i
		})
	}
	
	// Agora funciona corretamente
	for _, f := range funcoesCorretas {
		fmt.Printf("Valor correto: %d\n", f()) // 0, 1, 2
	}
}

// ============================================
// 7. CALL BY VALUE - Passagem por Valor
// ============================================

// Go passa TUDO por valor (cópia)
func modificarInt(numero int) {
	numero = 999 // Modifica apenas a cópia
	fmt.Printf("Dentro da função: %d\n", numero)
}

// Structs são copiados completamente
type Pessoa struct {
	Nome  string
	Idade int
}

func modificarPessoa(p Pessoa) {
	p.Nome = "Modificado"
	p.Idade = 999
	fmt.Printf("Dentro da função: %+v\n", p)
}

// Arrays são copiados completamente (pode ser custoso!)
func modificarArray(arr [5]int) {
	arr[0] = 999
	fmt.Printf("Dentro da função: %v\n", arr)
}

// Slices, maps e channels são passados por valor, mas...
// o valor é uma referência! (slice header, map pointer, etc.)
func modificarSlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Modifica o array subjacente!
	}
	s = append(s, 100) // Mas append não afeta o slice original
	fmt.Printf("Dentro da função: %v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func exemploCallByValue() {
	// Int - passado por valor
	num := 42
	modificarInt(num)
	fmt.Printf("Fora da função: %d\n", num) // 42 (não mudou)
	
	// Struct - passado por valor
	pessoa := Pessoa{Nome: "João", Idade: 30}
	modificarPessoa(pessoa)
	fmt.Printf("Fora da função: %+v\n", pessoa) // Não mudou
	
	// Array - passado por valor (cópia completa)
	arr := [5]int{1, 2, 3, 4, 5}
	modificarArray(arr)
	fmt.Printf("Fora da função: %v\n", arr) // Não mudou
	
	// Slice - passado por valor, mas valor é referência
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Antes: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
	modificarSlice(slice)
	fmt.Printf("Depois: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
	// Elementos modificados, mas length não mudou (append não afeta original)
}

// ============================================
// 8. EXEMPLOS PRÁTICOS COMBINADOS
// ============================================

// Função que usa múltiplos conceitos
func processarDados(numeros []int, filtro func(int) bool, transformacao func(int) int) ([]int, int, error) {
	if len(numeros) == 0 {
		return nil, 0, fmt.Errorf("slice vazio")
	}
	
	resultado := make([]int, 0, len(numeros))
	soma := 0
	
	for _, num := range numeros {
		if filtro(num) {
			transformado := transformacao(num)
			resultado = append(resultado, transformado)
			soma += transformado
		}
	}
	
	return resultado, soma, nil
}

func exemploCombinado() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Filtrar pares e dobrar
	filtroPar := func(n int) bool { return n%2 == 0 }
	dobrar := func(n int) int { return n * 2 }
	
	resultado, soma, err := processarDados(numeros, filtroPar, dobrar)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	
	fmt.Printf("Resultado: %v\n", resultado)
	fmt.Printf("Soma: %d\n", soma)
}

// ============================================
// MAIN - Executar Exemplos
// ============================================

func main() {
	fmt.Println("=== 1. FUNCTION BASICS ===")
	cumprimentar()
	cumprimentarPessoa("Maria")
	fmt.Printf("Soma: %d\n", somar(5, 3))
	fmt.Printf("Multiplicação: %d\n", multiplicar(2, 3, 4))
	exemploFirstClass()
	
	fmt.Println("\n=== 2. VARIADIC FUNCTIONS ===")
	exemploVariadico()
	
	fmt.Println("\n=== 3. MULTIPLE RETURN VALUES ===")
	exemploMultiplosRetornos()
	
	fmt.Println("\n=== 4. NAMED RETURN VALUES ===")
	exemploRetornosNomeados()
	
	fmt.Println("\n=== 5. ANONYMOUS FUNCTIONS ===")
	exemploFuncoesAnonimas()
	
	fmt.Println("\n=== 6. CLOSURES ===")
	exemploClosures()
	fmt.Println("\n--- Armadilha de Closure ---")
	exemploClosureArmadilha()
	
	fmt.Println("\n=== 7. CALL BY VALUE ===")
	exemploCallByValue()
	
	fmt.Println("\n=== 8. EXEMPLO COMBINADO ===")
	exemploCombinado()
}


