package main

import (
	"fmt"
	"strconv"
)

/*
 * ============================================================================
 * AULA 03: TIPOS DE DADOS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Integers: int8, int16, int32, int64, int, uint
 * - Floating Points: float32, float64
 * - Complex Numbers: complex64, complex128
 * - Boolean: bool (true/false)
 * - Runes: int32 (caracteres Unicode)
 * - Strings: Raw (`texto`) e Interpreted ("texto")
 * - Conversão de tipos: Tipo(valor) - EXPLÍCITA
 */

func main() {
	fmt.Println("=== AULA 03: TIPOS DE DADOS ===\n")
	
	// ===== INTEGERS =====
	fmt.Println("1. INTEGERS:")
	var a int8 = -128
	var b int16 = -32768
	var c int = 42
	var d uint = 100
	
	fmt.Printf("   int8: %d\n", a)
	fmt.Printf("   int16: %d\n", b)
	fmt.Printf("   int: %d\n", c)
	fmt.Printf("   uint: %d\n", d)
	
	// ===== FLOATING POINTS =====
	fmt.Println("\n2. FLOATING POINTS:")
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	
	fmt.Printf("   float32: %.2f\n", f32)
	fmt.Printf("   float64: %.15f\n", f64)
	
	// ⚠️ CUIDADO: Erros de precisão
	resultado := 0.1 + 0.2
	fmt.Printf("   0.1 + 0.2 = %.20f (não exatamente 0.3!)\n", resultado)
	
	// ===== COMPLEX NUMBERS =====
	fmt.Println("\n3. COMPLEX NUMBERS:")
	var z1 complex64 = 3 + 4i
	var z2 = complex(5.5, 7.2)
	
	fmt.Printf("   complex64: %v\n", z1)
	fmt.Printf("   complex128: %v\n", z2)
	fmt.Printf("   Parte real: %.1f\n", real(z2))
	fmt.Printf("   Parte imaginária: %.1f\n", imag(z2))
	
	// ===== BOOLEAN =====
	fmt.Println("\n4. BOOLEAN:")
	var estaChovendo bool = true
	var temSol bool = false
	var condicao bool // Valor zero: false
	
	fmt.Printf("   estaChovendo: %t\n", estaChovendo)
	fmt.Printf("   temSol: %t\n", temSol)
	fmt.Printf("   condicao (zero): %t\n", condicao)
	
	// Operações lógicas
	idade := 18
	maiorDeIdade := idade >= 18
	fmt.Printf("   maiorDeIdade (18 >= 18): %t\n", maiorDeIdade)
	
	// ===== RUNES =====
	fmt.Println("\n5. RUNES (Unicode):")
	var letra rune = 'A'
	var chines rune = '中'
	var emoji rune = '🚀'
	
	fmt.Printf("   letra 'A': %c (Unicode: %d)\n", letra, letra)
	fmt.Printf("   chinês '中': %c (Unicode: %d)\n", chines, chines)
	fmt.Printf("   emoji '🚀': %c (Unicode: %d)\n", emoji, emoji)
	
	// Processar string como runes
	texto := "Olá, 世界! 🎉"
	fmt.Printf("   Processando string: %s\n", texto)
	for i, char := range texto {
		fmt.Printf("      Posição %d: %c (Unicode: %d)\n", i, char, char)
		if i >= 10 { // Limitar output
			fmt.Printf("      ...\n")
			break
		}
	}
	
	// ===== STRINGS =====
	fmt.Println("\n6. STRINGS:")
	// Raw string (backticks) - literal
	rawString := `C:\Users\Documentos\arquivo.txt
Linha 1
Linha 2`
	fmt.Printf("   Raw string:\n%s\n", rawString)
	
	// Interpreted string (aspas) - processa escapes
	interpretedString := "Olá,\nMundo!\nTab:\tTabulação"
	fmt.Printf("   Interpreted string:\n%s\n", interpretedString)
	
	// ===== CONVERSÃO DE TIPOS =====
	fmt.Println("\n7. CONVERSÃO DE TIPOS (EXPLÍCITA):")
	var x int = 42
	var y int64 = int64(x)
	var z float64 = float64(x)
	
	fmt.Printf("   int para int64: %d → %d\n", x, y)
	fmt.Printf("   int para float64: %d → %.1f\n", x, z)
	
	// Conversão float para int (trunca)
	var f float64 = 3.99
	var i int = int(f)
	fmt.Printf("   float64 para int: %.2f → %d (trunca!)\n", f, i)
	
	// Conversão número para string (use strconv)
	numero := 42
	textoNumero := strconv.Itoa(numero)
	fmt.Printf("   int para string: %d → '%s'\n", numero, textoNumero)
	
	// String para número
	texto := "123"
	numero2, err := strconv.Atoi(texto)
	if err == nil {
		fmt.Printf("   string para int: '%s' → %d\n", texto, numero2)
	}
	
	// ⚠️ CUIDADO: string(numero) converte para caractere Unicode!
	unicodeChar := string(65)
	fmt.Printf("   string(65) = '%s' (caractere Unicode 65, não '65'!)\n", unicodeChar)
}

