package main

import "fmt"

/*
 * ============================================================================
 * AULA 10: FUNCTIONS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Funções são first-class citizens
 * - Múltiplos retornos (padrão: resultado, error)
 * - Retornos nomeados
 * - Variadic functions (...Tipo)
 * - Anonymous functions
 * - Closures
 * - Call by value (tudo é cópia, mas slices/maps são referências)
 */

// Função simples
func cumprimentar() {
	fmt.Println("Olá, mundo!")
}

// Função com parâmetros
func cumprimentarPessoa(nome string) {
	fmt.Printf("Olá, %s!\n", nome)
}

// Função com retorno
func somar(a int, b int) int {
	return a + b
}

// Função com múltiplos retornos (padrão idiomático)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

// Função com retornos nomeados
func calcular() (soma int, produto int) {
	soma = 10 + 20
	produto = 10 * 20
	return // Retorna soma e produto implicitamente
}

// Variadic function
func somarNumeros(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

// Função que aceita função como parâmetro
func aplicarOperacao(a, b int, operacao func(int, int) int) int {
	return operacao(a, b)
}

// Closure: função que retorna função
func criarContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// Closure com parâmetros
func criarMultiplicador(fator int) func(int) int {
	return func(numero int) int {
		return numero * fator
	}
}

func main() {
	fmt.Println("=== AULA 10: FUNCTIONS ===\n")
	
	// ===== FUNÇÕES BÁSICAS =====
	fmt.Println("1. FUNÇÕES BÁSICAS:")
	cumprimentar()
	cumprimentarPessoa("João")
	
	resultado := somar(5, 3)
	fmt.Printf("   soma(5, 3) = %d\n", resultado)
	
	// ===== MÚLTIPLOS RETORNOS =====
	fmt.Println("\n2. MÚLTIPLOS RETORNOS:")
	resultado2, err := dividir(10, 2)
	if err != nil {
		fmt.Printf("   Erro: %v\n", err)
	} else {
		fmt.Printf("   dividir(10, 2) = %.2f\n", resultado2)
	}
	
	resultado3, err2 := dividir(10, 0)
	if err2 != nil {
		fmt.Printf("   Erro: %v\n", err2)
	}
	
	// Ignorar valores
	_, _ = dividir(10, 2) // Ignora ambos
	
	// ===== RETORNOS NOMEADOS =====
	fmt.Println("\n3. RETORNOS NOMEADOS:")
	soma, produto := calcular()
	fmt.Printf("   soma: %d, produto: %d\n", soma, produto)
	
	// ===== VARIADIC FUNCTIONS =====
	fmt.Println("\n4. VARIADIC FUNCTIONS:")
	resultado4 := somarNumeros(1, 2, 3)
	fmt.Printf("   somarNumeros(1, 2, 3) = %d\n", resultado4)
	
	resultado5 := somarNumeros(10, 20, 30, 40)
	fmt.Printf("   somarNumeros(10, 20, 30, 40) = %d\n", resultado5)
	
	// Desempacotar slice
	numeros := []int{5, 10, 15}
	resultado6 := somarNumeros(numeros...)
	fmt.Printf("   somarNumeros([]int{5,10,15}...) = %d\n", resultado6)
	
	// ===== FUNÇÕES COMO VALORES =====
	fmt.Println("\n5. FUNÇÕES COMO VALORES:")
	var multiplicar func(int, int) int
	multiplicar = func(a, b int) int {
		return a * b
	}
	fmt.Printf("   multiplicar(5, 3) = %d\n", multiplicar(5, 3))
	
	// ===== FUNÇÕES COMO PARÂMETROS =====
	fmt.Println("\n6. FUNÇÕES COMO PARÂMETROS:")
	resultado7 := aplicarOperacao(10, 5, somar)
	fmt.Printf("   aplicarOperacao(10, 5, somar) = %d\n", resultado7)
	
	resultado8 := aplicarOperacao(10, 5, multiplicar)
	fmt.Printf("   aplicarOperacao(10, 5, multiplicar) = %d\n", resultado8)
	
	// ===== ANONYMOUS FUNCTIONS =====
	fmt.Println("\n7. ANONYMOUS FUNCTIONS:")
	somar2 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("   somar2(7, 3) = %d\n", somar2(7, 3))
	
	// IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("   Função anônima executada imediatamente!")
	}()
	
	// ===== CLOSURES =====
	fmt.Println("\n8. CLOSURES:")
	contador := criarContador()
	fmt.Printf("   contador(): %d\n", contador())
	fmt.Printf("   contador(): %d\n", contador())
	fmt.Printf("   contador(): %d\n", contador())
	
	// Closure com parâmetros
	dobrar := criarMultiplicador(2)
	triplicar := criarMultiplicador(3)
	
	fmt.Printf("   dobrar(5) = %d\n", dobrar(5))
	fmt.Printf("   triplicar(5) = %d\n", triplicar(5))
	
	// ===== CALL BY VALUE =====
	fmt.Println("\n9. CALL BY VALUE:")
	x := 10
	fmt.Printf("   x antes: %d\n", x)
	naoModifica(x)
	fmt.Printf("   x depois: %d (não mudou!)\n", x)
	
	modificaPonteiro(&x)
	fmt.Printf("   x depois de modificaPonteiro: %d (mudou!)\n", x)
	
	// Slices são reference types
	slice := []int{1, 2, 3}
	fmt.Printf("   slice antes: %v\n", slice)
	modificaSlice(slice)
	fmt.Printf("   slice depois: %v (elementos modificados!)\n", slice)
}

func naoModifica(valor int) {
	valor = 20 // Modifica apenas a cópia
}

func modificaPonteiro(ptr *int) {
	*ptr = 20 // Modifica o original
}

func modificaSlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Modifica o array subjacente
	}
}



