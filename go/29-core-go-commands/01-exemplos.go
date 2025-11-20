// Package exemplos demonstra exemplos práticos dos comandos core do Go.
//
// Este arquivo contém exemplos de código que podem ser usados para
// praticar os comandos go run, go build, go test, go doc, etc.
package main

import (
	"fmt"
	"math"
	"runtime"
)

// Versao armazena a versão da aplicação (pode ser definida via ldflags).
var Versao = "dev"

// main é a função principal que demonstra vários conceitos.
func main() {
	fmt.Println("=== Exemplos dos Comandos Core do Go ===\n")

	// Exemplo 1: Informações do sistema
	mostrarInfoSistema()

	// Exemplo 2: Operações matemáticas
	exemploCalculos()

	// Exemplo 3: Trabalhando com strings
	exemploStrings()

	// Exemplo 4: Versão da aplicação
	fmt.Printf("\nVersão da aplicação: %s\n", Versao)
}

// mostrarInfoSistema exibe informações sobre o sistema e runtime.
func mostrarInfoSistema() {
	fmt.Println("1. Informações do Sistema:")
	fmt.Printf("   Go Version: %s\n", runtime.Version())
	fmt.Printf("   OS: %s\n", runtime.GOOS)
	fmt.Printf("   Arch: %s\n", runtime.GOARCH)
	fmt.Printf("   Num CPU: %d\n", runtime.NumCPU())
	fmt.Println()
}

// exemploCalculos demonstra operações matemáticas básicas.
func exemploCalculos() {
	fmt.Println("2. Exemplos de Cálculos:")

	x := 10.5
	y := 20.3

	soma := x + y
	produto := x * y
	distancia := math.Sqrt(x*x + y*y)

	fmt.Printf("   Soma: %.2f + %.2f = %.2f\n", x, y, soma)
	fmt.Printf("   Produto: %.2f * %.2f = %.2f\n", x, y, produto)
	fmt.Printf("   Distância: √(%.2f² + %.2f²) = %.2f\n", x, y, distancia)
	fmt.Println()
}

// exemploStrings demonstra manipulação de strings.
func exemploStrings() {
	fmt.Println("3. Exemplos de Strings:")

	texto := "Olá, Mundo!"
	fmt.Printf("   Texto original: %s\n", texto)
	fmt.Printf("   Tamanho: %d caracteres\n", len(texto))
	fmt.Printf("   Primeiro caractere: %c\n", texto[0])
	fmt.Println()
}

// Soma adiciona dois números inteiros e retorna o resultado.
//
// Esta função é usada para demonstrar go doc e go test.
//
// Exemplo:
//
//	resultado := Soma(5, 3)
//	// resultado será 8
func Soma(a, b int) int {
	return a + b
}

// Multiplica multiplica dois números inteiros e retorna o resultado.
//
// Exemplo:
//
//	resultado := Multiplica(4, 7)
//	// resultado será 28
func Multiplica(a, b int) int {
	return a * b
}

// Divide divide a por b e retorna o resultado.
//
// Retorna um erro se b for zero.
//
// Exemplo:
//
//	resultado, err := Divide(10, 2)
//	// resultado será 5.0, err será nil
//
//	resultado, err := Divide(10, 0)
//	// resultado será 0, err não será nil
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não permitida")
	}
	return a / b, nil
}

// CalculaMedia calcula a média aritmética de uma lista de números.
//
// Parâmetros:
//   - numeros: slice de números float64
//
// Retorna:
//   - float64: a média calculada
//   - error: erro se a lista estiver vazia
//
// Exemplo:
//
//	numeros := []float64{10, 20, 30}
//	media, err := CalculaMedia(numeros)
//	// media será 20.0
func CalculaMedia(numeros []float64) (float64, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("lista vazia: não é possível calcular média")
	}

	soma := 0.0
	for _, n := range numeros {
		soma += n
	}

	return soma / float64(len(numeros)), nil
}

// Potencia calcula a potência de um número.
//
// Parâmetros:
//   - base: número base
//   - expoente: expoente (deve ser não-negativo)
//
// Retorna:
//   - float64: resultado da potência
//   - error: erro se o expoente for negativo
func Potencia(base, expoente float64) (float64, error) {
	if expoente < 0 {
		return 0, fmt.Errorf("expoente negativo não suportado")
	}
	return math.Pow(base, expoente), nil
}

/*
INSTRUÇÕES PARA USAR ESTE ARQUIVO COM OS COMANDOS CORE:

1. go run:
   $ go run 01-exemplos.go
   Executa o programa diretamente sem criar binário.

2. go build:
   $ go build -o exemplos 01-exemplos.go
   $ ./exemplos
   Compila e cria um binário executável.

3. go build com ldflags (definir versão):
   $ go build -ldflags="-X main.Versao=1.0.0" -o exemplos 01-exemplos.go
   $ ./exemplos
   Define a versão da aplicação durante a compilação.

4. go doc:
   $ go doc .
   $ go doc Soma
   $ go doc Divide
   $ go doc -src Soma
   Visualiza a documentação das funções.

5. go fmt:
   $ go fmt 01-exemplos.go
   Formata o código automaticamente.

6. go test:
   $ go test -v
   Executa os testes definidos em 01-exemplos_test.go

7. go version:
   $ go version
   $ go version ./exemplos
   Mostra a versão do Go e do binário compilado.

8. go mod (se este for um módulo):
   $ go mod init exemplos
   $ go mod tidy
   Gerencia dependências do módulo.
*/



