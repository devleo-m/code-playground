// Este arquivo contém testes para demonstrar o uso de go test.
package main

import (
	"fmt"
	"math"
	"testing"
)

// TestSoma testa a função Soma com vários casos.
func TestSoma(t *testing.T) {
	tests := []struct {
		nome     string
		a        int
		b        int
		esperado int
	}{
		{"positivos", 2, 3, 5},
		{"zero", 0, 5, 5},
		{"negativos", -2, -3, -5},
		{"misto", -5, 10, 5},
		{"grandes", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := Soma(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("Soma(%d, %d) = %d; esperado %d", tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}

// TestMultiplica testa a função Multiplica.
func TestMultiplica(t *testing.T) {
	tests := []struct {
		nome     string
		a        int
		b        int
		esperado int
	}{
		{"básico", 3, 4, 12},
		{"zero", 5, 0, 0},
		{"um", 7, 1, 7},
		{"negativos", -3, 4, -12},
		{"ambos negativos", -2, -5, 10},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := Multiplica(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("Multiplica(%d, %d) = %d; esperado %d", tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}

// TestDivide testa a função Divide com casos de sucesso e erro.
func TestDivide(t *testing.T) {
	t.Run("divisão válida", func(t *testing.T) {
		resultado, err := Divide(10, 2)
		if err != nil {
			t.Errorf("Divide(10, 2) retornou erro inesperado: %v", err)
		}
		if resultado != 5.0 {
			t.Errorf("Divide(10, 2) = %f; esperado 5.0", resultado)
		}
	})

	t.Run("divisão por zero", func(t *testing.T) {
		resultado, err := Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) deveria retornar erro, mas não retornou")
		}
		if resultado != 0 {
			t.Errorf("Divide(10, 0) deveria retornar 0, mas retornou %f", resultado)
		}
	})

	t.Run("divisão com decimais", func(t *testing.T) {
		resultado, err := Divide(7, 2)
		if err != nil {
			t.Errorf("Divide(7, 2) retornou erro inesperado: %v", err)
		}
		esperado := 3.5
		if math.Abs(resultado-esperado) > 0.0001 {
			t.Errorf("Divide(7, 2) = %f; esperado %f", resultado, esperado)
		}
	})
}

// TestCalculaMedia testa a função CalculaMedia.
func TestCalculaMedia(t *testing.T) {
	t.Run("lista válida", func(t *testing.T) {
		numeros := []float64{10, 20, 30}
		media, err := CalculaMedia(numeros)
		if err != nil {
			t.Errorf("CalculaMedia retornou erro inesperado: %v", err)
		}
		if media != 20.0 {
			t.Errorf("CalculaMedia = %f; esperado 20.0", media)
		}
	})

	t.Run("lista vazia", func(t *testing.T) {
		numeros := []float64{}
		media, err := CalculaMedia(numeros)
		if err == nil {
			t.Error("CalculaMedia com lista vazia deveria retornar erro")
		}
		if media != 0 {
			t.Errorf("CalculaMedia com lista vazia deveria retornar 0, mas retornou %f", media)
		}
	})

	t.Run("um número", func(t *testing.T) {
		numeros := []float64{42}
		media, err := CalculaMedia(numeros)
		if err != nil {
			t.Errorf("CalculaMedia retornou erro inesperado: %v", err)
		}
		if media != 42.0 {
			t.Errorf("CalculaMedia = %f; esperado 42.0", media)
		}
	})
}

// TestPotencia testa a função Potencia.
func TestPotencia(t *testing.T) {
	tests := []struct {
		nome     string
		base     float64
		expoente float64
		esperado float64
		temErro  bool
	}{
		{"2^3", 2, 3, 8, false},
		{"5^0", 5, 0, 1, false},
		{"10^2", 10, 2, 100, false},
		{"expoente negativo", 2, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			resultado, err := Potencia(tt.base, tt.expoente)
			if tt.temErro {
				if err == nil {
					t.Error("Potencia deveria retornar erro, mas não retornou")
				}
			} else {
				if err != nil {
					t.Errorf("Potencia retornou erro inesperado: %v", err)
				}
				if math.Abs(resultado-tt.esperado) > 0.0001 {
					t.Errorf("Potencia(%.2f, %.2f) = %f; esperado %f", tt.base, tt.expoente, resultado, tt.esperado)
				}
			}
		})
	}
}

// BenchmarkSoma mede a performance da função Soma.
func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(100, 200)
	}
}

// BenchmarkMultiplica mede a performance da função Multiplica.
func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(100, 200)
	}
}

// BenchmarkCalculaMedia mede a performance da função CalculaMedia.
func BenchmarkCalculaMedia(b *testing.B) {
	numeros := []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CalculaMedia(numeros)
	}
}

// ExampleSoma demonstra como usar a função Soma.
func ExampleSoma() {
	resultado := Soma(5, 3)
	fmt.Println(resultado)
	// Output: 8
}

// ExampleDivide demonstra como usar a função Divide.
func ExampleDivide() {
	resultado, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)
	// Output: 5
}

// ExampleCalculaMedia demonstra como usar a função CalculaMedia.
func ExampleCalculaMedia() {
	numeros := []float64{10, 20, 30}
	media, err := CalculaMedia(numeros)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Média: %.1f\n", media)
	// Output: Média: 20.0
}

