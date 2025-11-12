package main

import (
	"fmt"
	"math"
	// "math/cmplx"
)

// Exemplos de uso dos diferentes tipos de dados em Go

func main() {
	fmt.Println("=== EXEMPLOS DE TIPOS DE DADOS EM GO ===\n")

	// 1. INTEGERS (Inteiros)
	demonstrarIntegers()

	// 2. FLOATING POINTS (Ponto Flutuante)
	demonstrarFloats()

	// 3. COMPLEX NUMBERS (Números Complexos)
	demonstrarComplexos()

	// 4. BOOLEAN (Booleano)
	demonstrarBoolean()

	// 5. RUNES
	demonstrarRunes()

	// 6. STRINGS (Raw e Interpreted)
	demonstrarStrings()

	// 7. TYPE CONVERSION (Conversão de Tipos)
	demonstrarConversao()
}

func demonstrarIntegers() {
	fmt.Println("1. INTEGERS:")
	fmt.Println("-------------")

	var a int8 = -128
	var b int16 = 32767
	var c int32 = -2147483648
	var d int64 = 9223372036854775807

	var e uint8 = 255
	var f uint16 = 65535
	var g uint = 100 // Dependente da plataforma

	var h int = 42 // Dependente da plataforma (geralmente int64 em sistemas 64-bit)

	fmt.Printf("int8:   %d\n", a)
	fmt.Printf("int16:  %d\n", b)
	fmt.Printf("int32:  %d\n", c)
	fmt.Printf("int64:  %d\n", d)
	fmt.Printf("uint8:  %d\n", e)
	fmt.Printf("uint16: %d\n", f)
	fmt.Printf("uint:   %d\n", g)
	fmt.Printf("int:    %d\n", h)
	fmt.Println()
}

func demonstrarFloats() {
	fmt.Println("2. FLOATING POINTS:")
	fmt.Println("-------------------")

	var a float32 = 3.14159
	var b float64 = 3.141592653589793
	var c = 19.99 // Inferido como float64

	// Demonstração do erro de precisão
	var d float64 = 0.1 + 0.2

	fmt.Printf("float32: %f\n", a)
	fmt.Printf("float64: %f\n", b)
	fmt.Printf("float64 (inferido): %f\n", c)
	fmt.Printf("Erro de precisão (0.1 + 0.2): %.20f\n", d)
	fmt.Printf("É exatamente 0.3? %v\n", d == 0.3)
	fmt.Println()
}

func demonstrarComplexos() {
	fmt.Println("3. COMPLEX NUMBERS:")
	fmt.Println("-------------------")

	var z1 complex64 = 3 + 4i
	var z2 complex128 = 5.5 + 7.2i
	var z3 = complex(3.0, 4.0) // Usando função complex()

	fmt.Printf("complex64:  %v\n", z1)
	fmt.Printf("complex128: %v\n", z2)
	fmt.Printf("complex():  %v\n", z3)

	// Funções úteis
	fmt.Printf("Parte real de z1: %f\n", real(z1))
	fmt.Printf("Parte imaginária de z1: %f\n", imag(z1))
	// fmt.Printf("Módulo de z1: %f\n", cmplx.Abs(z1))
	fmt.Println()
}

func demonstrarBoolean() {
	fmt.Println("4. BOOLEAN:")
	fmt.Println("-----------")

	var verdadeiro bool = true
	var falso bool = false
	var condicao bool // Valor zero é false

	idade := 18
	maiorDeIdade := idade >= 18

	fmt.Printf("true:  %t\n", verdadeiro)
	fmt.Printf("false: %t\n", falso)
	fmt.Printf("zero value: %t\n", condicao)
	fmt.Printf("idade >= 18: %t\n", maiorDeIdade)
	fmt.Println()
}

func demonstrarRunes() {
	fmt.Println("5. RUNES:")
	fmt.Println("---------")

	var letra rune = 'A'
	var chines rune = '中'
	var emoji rune = '🚀'
	var proximaLetra rune = 'A' + 1

	fmt.Printf("'A' como rune: %c (código: %d)\n", letra, letra)
	fmt.Printf("'中' como rune: %c (código: %d)\n", chines, chines)
	fmt.Printf("'🚀' como rune: %c (código: %d)\n", emoji, emoji)
	fmt.Printf("'A' + 1 = %c (código: %d)\n", proximaLetra, proximaLetra)

	// Iterando sobre uma string com caracteres internacionais
	texto := "Olá, 世界! 🎉"
	fmt.Printf("\nIterando sobre: %s\n", texto)
	for i, char := range texto {
		fmt.Printf("  Posição %d: %c (Unicode: %d, ASCII: %v)\n",
			i, char, char, char < 128)
	}
	fmt.Println()
}

func demonstrarStrings() {
	fmt.Println("6. STRINGS:")
	fmt.Println("-----------")

	// Raw String Literal (backticks)
	raw := `C:\Users\Documentos\arquivo.txt
Linha 1
Linha 2\tTab aqui
"Aspas sem escapar"`
	fmt.Println("Raw String (backticks):")
	fmt.Println(raw)
	fmt.Println()

	// Interpreted String Literal (aspas duplas)
	interpreted := "C:\\Users\\Documentos\\arquivo.txt\nLinha 1\nLinha 2\tTab aqui\n\"Aspas escapadas\""
	fmt.Println("Interpreted String (aspas duplas):")
	fmt.Println(interpreted)
	fmt.Println()

	// Comparação
	fmt.Println("Diferença:")
	fmt.Println("Raw preserva \\n e \\t literalmente")
	fmt.Println("Interpreted processa \\n (nova linha) e \\t (tab)")
	fmt.Println()
}

func demonstrarConversao() {
	fmt.Println("7. TYPE CONVERSION:")
	fmt.Println("-------------------")

	// Conversão entre tipos numéricos
	var x int = 42
	var y float64 = float64(x)
	var z int = int(y)

	fmt.Printf("int: %d\n", x)
	fmt.Printf("int -> float64: %f\n", y)
	fmt.Printf("float64 -> int: %d\n", z)

	// Conversão com perda de dados
	var grande float64 = 3.9
	var pequeno int = int(grande) // Trunca para 3
	fmt.Printf("\nfloat64 3.9 -> int: %d (perdeu a parte decimal)\n", pequeno)

	// Conversão de string para número (requer strconv)
	// Nota: Isso será mostrado em uma aula futura sobre pacotes
	// Por enquanto, apenas mencionamos que string() não faz o que você espera:
	var numero int = 65
	var caractere string = string(numero) // ⚠️ CUIDADO: Isso não cria "65"!
	fmt.Printf("\nint 65 -> string usando string(): %q (não é '65'!)\n", caractere)
	fmt.Printf("Isso criou um caractere Unicode 65, que é: %c\n", numero)

	// Demonstração de overflow
	var pequenoInt8 int8 = 127
	fmt.Printf("\nint8 antes do incremento: %d\n", pequenoInt8)
	pequenoInt8++ // Overflow!
	fmt.Printf("int8 depois do incremento: %d (overflow!)\n", pequenoInt8)

	// Comparação de floats com tolerância
	fmt.Println("\nComparação de floats:")
	var a float64 = 0.1 + 0.2
	var b float64 = 0.3
	const epsilon = 1e-9

	fmt.Printf("0.1 + 0.2 = %.20f\n", a)
	fmt.Printf("0.3 = %.20f\n", b)
	fmt.Printf("São iguais (==)? %v\n", a == b)
	fmt.Printf("São iguais (com tolerância)? %v\n", math.Abs(a-b) < epsilon)
}
