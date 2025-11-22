package main

import "fmt"

/*
 * ============================================================================
 * AULA 07: STRUCTS
 * ============================================================================
 * 
 * RESUMO DOS CONCEITOS:
 * - Agrupa dados relacionados sob um tipo
 * - Modela entidades do mundo real
 * - Campos nomeados ou por ordem
 * - Métodos podem ser adicionados
 * - Comparáveis se todos campos forem comparáveis
 */

// Declaração de struct
type Pessoa struct {
	Nome   string
	Idade  int
	Email  string
	Altura float64
}

// Método com value receiver (não modifica)
func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, eu sou %s e tenho %d anos", p.Nome, p.Idade)
}

// Método com pointer receiver (pode modificar)
func (p *Pessoa) FazerAniversario() {
	p.Idade++
}

func main() {
	fmt.Println("=== AULA 07: STRUCTS ===\n")
	
	// ===== DECLARAÇÃO COM VALORES ZERO =====
	fmt.Println("1. DECLARAÇÃO COM VALORES ZERO:")
	var pessoa1 Pessoa
	fmt.Printf("   pessoa1: %+v\n", pessoa1)
	
	// ===== LITERAL COM CAMPOS NOMEADOS (RECOMENDADO) =====
	fmt.Println("\n2. LITERAL COM CAMPOS NOMEADOS:")
	pessoa2 := Pessoa{
		Nome:   "João Silva",
		Idade:  30,
		Email:  "joao@email.com",
		Altura: 1.75,
	}
	fmt.Printf("   pessoa2: %+v\n", pessoa2)
	
	// ===== LITERAL POR ORDEM =====
	fmt.Println("\n3. LITERAL POR ORDEM:")
	pessoa3 := Pessoa{"Maria", 25, "maria@email.com", 1.65}
	fmt.Printf("   pessoa3: %+v\n", pessoa3)
	
	// ===== INICIALIZAÇÃO PARCIAL =====
	fmt.Println("\n4. INICIALIZAÇÃO PARCIAL:")
	pessoa4 := Pessoa{
		Nome:  "Pedro",
		Idade: 28,
		// Email e Altura recebem valores zero
	}
	fmt.Printf("   pessoa4: %+v\n", pessoa4)
	
	// ===== ACESSO A CAMPOS =====
	fmt.Println("\n5. ACESSO A CAMPOS:")
	fmt.Printf("   pessoa2.Nome: %s\n", pessoa2.Nome)
	fmt.Printf("   pessoa2.Idade: %d\n", pessoa2.Idade)
	
	// Modificar campos
	pessoa2.Idade = 31
	pessoa2.Email = "joao.novo@email.com"
	fmt.Printf("   Após modificar: %+v\n", pessoa2)
	
	// ===== COM PONTEIROS =====
	fmt.Println("\n6. COM PONTEIROS:")
	pessoa5 := &Pessoa{
		Nome:  "Ana",
		Idade: 22,
	}
	
	// Go permite acesso direto (sem (*p).Nome)
	fmt.Printf("   pessoa5.Nome: %s\n", pessoa5.Nome)
	pessoa5.Idade = 23
	fmt.Printf("   pessoa5.Idade: %d\n", pessoa5.Idade)
	
	// ===== MÉTODOS =====
	fmt.Println("\n7. MÉTODOS:")
	pessoa6 := Pessoa{Nome: "Carlos", Idade: 35}
	
	// Value receiver (não modifica)
	fmt.Printf("   %s\n", pessoa6.Apresentar())
	fmt.Printf("   Idade antes: %d\n", pessoa6.Idade)
	
	// Pointer receiver (modifica)
	pessoa6.FazerAniversario()
	fmt.Printf("   Idade depois: %d\n", pessoa6.Idade)
	
	// ===== COMPARAÇÃO =====
	fmt.Println("\n8. COMPARAÇÃO:")
	p1 := Pessoa{Nome: "João", Idade: 30}
	p2 := Pessoa{Nome: "João", Idade: 30}
	p3 := Pessoa{Nome: "Maria", Idade: 25}
	
	fmt.Printf("   p1 == p2: %t\n", p1 == p2)
	fmt.Printf("   p1 == p3: %t\n", p1 == p3)
	
	// ===== STRUCTS ANINHADAS =====
	fmt.Println("\n9. STRUCTS ANINHADAS:")
	type Endereco struct {
		Rua    string
		Cidade string
		CEP    string
	}
	
	type Usuario struct {
		Nome     string
		Idade    int
		Endereco Endereco
	}
	
	usuario := Usuario{
		Nome:  "Lucas",
		Idade: 28,
		Endereco: Endereco{
			Rua:    "Rua das Flores, 123",
			Cidade: "São Paulo",
			CEP:    "01234-567",
		},
	}
	fmt.Printf("   usuario: %+v\n", usuario)
	fmt.Printf("   usuario.Endereco.Cidade: %s\n", usuario.Endereco.Cidade)
	
	// ===== SLICES DE STRUCTS =====
	fmt.Println("\n10. SLICES DE STRUCTS:")
	pessoas := []Pessoa{
		{Nome: "João", Idade: 30},
		{Nome: "Maria", Idade: 25},
		{Nome: "Pedro", Idade: 35},
	}
	
	for i, p := range pessoas {
		fmt.Printf("   [%d] %s tem %d anos\n", i, p.Nome, p.Idade)
	}
}

