package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	// Remove o banco de dados existente se houver
	os.Remove("biblioteca.db")

	// Abre ou cria o banco de dados
	db, err := sql.Open("sqlite", "biblioteca.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("📚 Criando banco de dados da biblioteca...")

	// Cria as tabelas
	createTables(db)

	// Insere dados de exemplo
	insertData(db)

	fmt.Println("✅ Banco de dados criado com sucesso!")
	fmt.Println("📖 Banco: biblioteca.db")
	fmt.Println("\nTabelas criadas:")
	fmt.Println("  - autores")
	fmt.Println("  - livros")
	fmt.Println("  - usuarios")
	fmt.Println("  - emprestimos")
	fmt.Println("  - categorias")
	fmt.Println("\n💡 Você pode começar a praticar SQL agora!")
}

func createTables(db *sql.DB) {
	// Tabela de categorias
	categoriasSQL := `
	CREATE TABLE IF NOT EXISTS categorias (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL UNIQUE,
		descricao TEXT
	);`

	// Tabela de autores
	autoresSQL := `
	CREATE TABLE IF NOT EXISTS autores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		nacionalidade TEXT,
		data_nascimento DATE
	);`

	// Tabela de livros
	livrosSQL := `
	CREATE TABLE IF NOT EXISTS livros (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT NOT NULL,
		isbn TEXT UNIQUE,
		ano_publicacao INTEGER,
		editora TEXT,
		autor_id INTEGER,
		categoria_id INTEGER,
		quantidade_disponivel INTEGER DEFAULT 0,
		FOREIGN KEY (autor_id) REFERENCES autores(id),
		FOREIGN KEY (categoria_id) REFERENCES categorias(id)
	);`

	// Tabela de usuários
	usuariosSQL := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		telefone TEXT,
		data_cadastro DATE DEFAULT CURRENT_DATE
	);`

	// Tabela de empréstimos
	emprestimosSQL := `
	CREATE TABLE IF NOT EXISTS emprestimos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		livro_id INTEGER NOT NULL,
		usuario_id INTEGER NOT NULL,
		data_emprestimo DATE DEFAULT CURRENT_DATE,
		data_devolucao_prevista DATE,
		data_devolucao_real DATE,
		status TEXT DEFAULT 'ativo',
		FOREIGN KEY (livro_id) REFERENCES livros(id),
		FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
	);`

	queries := []struct {
		name string
		sql  string
	}{
		{"categorias", categoriasSQL},
		{"autores", autoresSQL},
		{"livros", livrosSQL},
		{"usuarios", usuariosSQL},
		{"emprestimos", emprestimosSQL},
	}

	for _, q := range queries {
		_, err := db.Exec(q.sql)
		if err != nil {
			log.Fatalf("Erro ao criar tabela %s: %v", q.name, err)
		}
		fmt.Printf("  ✓ Tabela '%s' criada\n", q.name)
	}
}

func insertData(db *sql.DB) {
	fmt.Println("\n📝 Inserindo dados de exemplo...")

	// Inserir categorias
	categorias := []struct {
		nome      string
		descricao string
	}{
		{"Ficção Científica", "Livros de ficção científica e fantasia"},
		{"Romance", "Romances e histórias de amor"},
		{"Técnico", "Livros técnicos de programação e tecnologia"},
		{"História", "Livros sobre história e biografias"},
		{"Filosofia", "Livros de filosofia e pensamento crítico"},
		{"Mistério", "Livros de mistério e suspense"},
	}

	for _, cat := range categorias {
		_, err := db.Exec("INSERT INTO categorias (nome, descricao) VALUES (?, ?)", cat.nome, cat.descricao)
		if err != nil {
			log.Printf("Erro ao inserir categoria %s: %v", cat.nome, err)
		}
	}

	// Inserir autores
	autores := []struct {
		nome          string
		nacionalidade string
		dataNasc      string
	}{
		{"Isaac Asimov", "Russo-Americano", "1920-01-02"},
		{"George Orwell", "Britânico", "1903-06-25"},
		{"J.K. Rowling", "Britânica", "1965-07-31"},
		{"Robert C. Martin", "Americano", "1952-12-05"},
		{"Yuval Noah Harari", "Israelense", "1976-02-24"},
		{"Agatha Christie", "Britânica", "1890-09-15"},
		{"Machado de Assis", "Brasileiro", "1839-06-21"},
		{"Clarice Lispector", "Brasileira", "1920-12-10"},
		{"Jorge Amado", "Brasileiro", "1912-08-10"},
		{"Carlos Drummond de Andrade", "Brasileiro", "1902-10-31"},
	}

	for _, aut := range autores {
		_, err := db.Exec("INSERT INTO autores (nome, nacionalidade, data_nascimento) VALUES (?, ?, ?)",
			aut.nome, aut.nacionalidade, aut.dataNasc)
		if err != nil {
			log.Printf("Erro ao inserir autor %s: %v", aut.nome, err)
		}
	}

	// Inserir livros
	livros := []struct {
		titulo               string
		isbn                 string
		anoPublicacao        int
		editora              string
		autorID              int
		categoriaID          int
		quantidadeDisponivel int
	}{
		{"Fundação", "978-8535914849", 1951, "Aleph", 1, 1, 5},
		{"Eu, Robô", "978-8535914848", 1950, "Aleph", 1, 1, 3},
		{"1984", "978-8535914847", 1949, "Companhia das Letras", 2, 2, 8},
		{"A Revolução dos Bichos", "978-8535914846", 1945, "Companhia das Letras", 2, 2, 6},
		{"Harry Potter e a Pedra Filosofal", "978-8535914845", 1997, "Rocco", 3, 1, 10},
		{"Código Limpo", "978-8535914844", 2008, "Alta Books", 4, 3, 4},
		{"Sapiens", "978-8535914843", 2011, "L&PM", 5, 4, 7},
		{"Assassinato no Expresso do Oriente", "978-8535914842", 1934, "Globo", 6, 6, 5},
		{"Dom Casmurro", "978-8535914841", 1899, "Globo", 7, 2, 9},
		{"A Hora da Estrela", "978-8535914840", 1977, "Rocco", 8, 2, 6},
		{"Capitães da Areia", "978-8535914839", 1937, "Companhia das Letras", 9, 2, 8},
		{"Claro Enigma", "978-8535914838", 1951, "Companhia das Letras", 10, 5, 4},
		{"O Guia do Mochileiro das Galáxias", "978-8535914837", 1979, "Arqueiro", 1, 1, 7},
		{"O Programador Pragmático", "978-8535914836", 1999, "Bookman", 4, 3, 3},
		{"Homo Deus", "978-8535914835", 2015, "Companhia das Letras", 5, 4, 5},
	}

	for _, liv := range livros {
		_, err := db.Exec(`INSERT INTO livros 
			(titulo, isbn, ano_publicacao, editora, autor_id, categoria_id, quantidade_disponivel) 
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
			liv.titulo, liv.isbn, liv.anoPublicacao, liv.editora, liv.autorID, liv.categoriaID, liv.quantidadeDisponivel)
		if err != nil {
			log.Printf("Erro ao inserir livro %s: %v", liv.titulo, err)
		}
	}

	// Inserir usuários
	usuarios := []struct {
		nome         string
		email        string
		telefone     string
		dataCadastro string
	}{
		{"João Silva", "joao.silva@email.com", "(11) 98765-4321", "2024-01-15"},
		{"Maria Santos", "maria.santos@email.com", "(11) 97654-3210", "2024-01-20"},
		{"Pedro Oliveira", "pedro.oliveira@email.com", "(21) 96543-2109", "2024-02-01"},
		{"Ana Costa", "ana.costa@email.com", "(21) 95432-1098", "2024-02-05"},
		{"Carlos Pereira", "carlos.pereira@email.com", "(31) 94321-0987", "2024-02-10"},
		{"Juliana Ferreira", "juliana.ferreira@email.com", "(31) 93210-9876", "2024-02-15"},
		{"Roberto Alves", "roberto.alves@email.com", "(41) 92109-8765", "2024-02-20"},
		{"Fernanda Lima", "fernanda.lima@email.com", "(41) 91098-7654", "2024-03-01"},
	}

	for _, usu := range usuarios {
		_, err := db.Exec("INSERT INTO usuarios (nome, email, telefone, data_cadastro) VALUES (?, ?, ?, ?)",
			usu.nome, usu.email, usu.telefone, usu.dataCadastro)
		if err != nil {
			log.Printf("Erro ao inserir usuário %s: %v", usu.nome, err)
		}
	}

	// Inserir empréstimos
	emprestimos := []struct {
		livroID               int
		usuarioID             int
		dataEmprestimo        string
		dataDevolucaoPrevista string
		dataDevolucaoReal     string
		status                string
	}{
		{1, 1, "2024-03-01", "2024-03-15", "", "ativo"},
		{3, 2, "2024-03-02", "2024-03-16", "2024-03-10", "devolvido"},
		{5, 3, "2024-03-03", "2024-03-17", "", "ativo"},
		{6, 4, "2024-03-04", "2024-03-18", "", "ativo"},
		{7, 1, "2024-02-20", "2024-03-05", "2024-03-05", "devolvido"},
		{8, 5, "2024-03-05", "2024-03-19", "", "ativo"},
		{9, 6, "2024-02-25", "2024-03-10", "2024-03-08", "devolvido"},
		{10, 7, "2024-03-06", "2024-03-20", "", "ativo"},
		{11, 8, "2024-03-07", "2024-03-21", "", "ativo"},
		{12, 2, "2024-02-15", "2024-03-01", "2024-02-28", "devolvido"},
	}

	for _, emp := range emprestimos {
		_, err := db.Exec(`INSERT INTO emprestimos 
			(livro_id, usuario_id, data_emprestimo, data_devolucao_prevista, data_devolucao_real, status) 
			VALUES (?, ?, ?, ?, ?, ?)`,
			emp.livroID, emp.usuarioID, emp.dataEmprestimo, emp.dataDevolucaoPrevista, emp.dataDevolucaoReal, emp.status)
		if err != nil {
			log.Printf("Erro ao inserir empréstimo: %v", err)
		}
	}

	fmt.Println("  ✓ Dados inseridos com sucesso!")
}
