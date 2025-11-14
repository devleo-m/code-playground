package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ============================================
// EXEMPLO 1: Struct Básica
// ============================================

type Pessoa struct {
	Nome   string
	Idade  int
	Email  string
	Altura float64
}

// Método com receptor por valor
func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, eu sou %s e tenho %d anos", p.Nome, p.Idade)
}

// Método com receptor por ponteiro (pode modificar)
func (p *Pessoa) FazerAniversario() {
	p.Idade++
}

// ============================================
// EXEMPLO 2: Struct com Tags JSON
// ============================================

type Usuario struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Senha     string    `json:"-"`              // Nunca serializar
	CriadoEm  time.Time `json:"criado_em"`
	Ativo     bool      `json:"ativo"`
	Telefone  string    `json:"telefone,omitempty"` // Omitir se vazio
}

// ============================================
// EXEMPLO 3: Structs Aninhadas
// ============================================

type Endereco struct {
	Rua    string `json:"rua"`
	Cidade string `json:"cidade"`
	CEP    string `json:"cep,omitempty"`
}

type PessoaCompleta struct {
	Nome     string   `json:"nome"`
	Idade    int      `json:"idade"`
	Endereco Endereco `json:"endereco"`
}

// ============================================
// EXEMPLO 4: Embedding (Composição)
// ============================================

// Timestamps fornece campos de data/hora
type Timestamps struct {
	CriadoEm    time.Time
	AtualizadoEm time.Time
}

func (t *Timestamps) Atualizar() {
	t.AtualizadoEm = time.Now()
}

// Auditoria fornece campos de auditoria
type Auditoria struct {
	CriadoPor    string
	AtualizadoPor string
}

// Produto incorpora Timestamps e Auditoria
type Produto struct {
	ID       int
	Nome     string
	Preco    float64
	Timestamps  // Embedding - campos acessíveis diretamente
	Auditoria   // Embedding
}

// ============================================
// EXEMPLO 5: Sistema de Biblioteca
// ============================================

type Livro struct {
	Titulo        string
	Autor         string
	ISBN          string
	AnoPublicacao int
	Disponivel    bool
}

func (l *Livro) Emprestar() {
	l.Disponivel = false
}

func (l *Livro) Devolver() {
	l.Disponivel = true
}

type Biblioteca struct {
	Nome   string
	Livros []Livro
}

func (b *Biblioteca) AdicionarLivro(livro Livro) {
	b.Livros = append(b.Livros, livro)
}

func (b *Biblioteca) BuscarPorTitulo(titulo string) *Livro {
	for i := range b.Livros {
		if b.Livros[i].Titulo == titulo {
			return &b.Livros[i]
		}
	}
	return nil
}

// ============================================
// FUNÇÃO MAIN - EXEMPLOS DE USO
// ============================================

func main() {
	fmt.Println("=== EXEMPLO 1: Struct Básica ===\n")

	pessoa := Pessoa{
		Nome:   "João Silva",
		Idade:  30,
		Email:  "joao@email.com",
		Altura: 1.75,
	}

	fmt.Println(pessoa.Apresentar())
	pessoa.FazerAniversario()
	fmt.Printf("Depois do aniversário: %d anos\n\n", pessoa.Idade)

	fmt.Println("=== EXEMPLO 2: Tags JSON ===\n")

	usuario := Usuario{
		ID:       1,
		Nome:     "Maria",
		Email:    "maria@email.com",
		Senha:    "senha123", // Não será serializado
		CriadoEm: time.Now(),
		Ativo:    true,
		// Telefone omitido (vazio) - não aparecerá no JSON
	}

	jsonBytes, err := json.MarshalIndent(usuario, "", "  ")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("JSON serializado:")
	fmt.Println(string(jsonBytes))
	fmt.Println()

	fmt.Println("=== EXEMPLO 3: Structs Aninhadas ===\n")

	pessoaCompleta := PessoaCompleta{
		Nome:  "Pedro",
		Idade: 25,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Cidade: "São Paulo",
			CEP:    "01234-567",
		},
	}

	jsonBytes2, _ := json.MarshalIndent(pessoaCompleta, "", "  ")
	fmt.Println("Pessoa com endereço:")
	fmt.Println(string(jsonBytes2))
	fmt.Println()

	fmt.Println("=== EXEMPLO 4: Embedding ===\n")

	produto := Produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 2500.00,
		Timestamps: Timestamps{
			CriadoEm:    time.Now(),
			AtualizadoEm: time.Now(),
		},
		Auditoria: Auditoria{
			CriadoPor: "admin",
		},
	}

	// Acesso direto a campos incorporados!
	fmt.Printf("Produto: %s\n", produto.Nome)
	fmt.Printf("Criado em: %s\n", produto.CriadoEm.Format("2006-01-02"))
	fmt.Printf("Criado por: %s\n", produto.CriadoPor)

	// Método de Timestamps acessível diretamente
	produto.Atualizar()
	fmt.Printf("Atualizado em: %s\n\n", produto.AtualizadoEm.Format("2006-01-02"))

	fmt.Println("=== EXEMPLO 5: Sistema de Biblioteca ===\n")

	biblioteca := Biblioteca{
		Nome: "Biblioteca Central",
	}

	livro1 := Livro{
		Titulo:        "Aprendendo Go",
		Autor:         "Autor Exemplo",
		ISBN:          "123-456-789",
		AnoPublicacao: 2024,
		Disponivel:    true,
	}

	livro2 := Livro{
		Titulo:        "Go Avançado",
		Autor:         "Outro Autor",
		ISBN:          "987-654-321",
		AnoPublicacao: 2023,
		Disponivel:    true,
	}

	biblioteca.AdicionarLivro(livro1)
	biblioteca.AdicionarLivro(livro2)

	fmt.Printf("Biblioteca: %s\n", biblioteca.Nome)
	fmt.Printf("Total de livros: %d\n", len(biblioteca.Livros))

	// Buscar e emprestar
	livro := biblioteca.BuscarPorTitulo("Aprendendo Go")
	if livro != nil {
		fmt.Printf("Livro encontrado: %s\n", livro.Titulo)
		livro.Emprestar()
		fmt.Printf("Disponível após empréstimo: %v\n", livro.Disponivel)
	}
}

