package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// ⚠️ ATENÇÃO: Este arquivo contém exemplos INTENCIONAIS de código vulnerável
// para fins educacionais. NÃO use este código em produção!

// Exemplo 1: SQL Injection
// ❌ VULNERÁVEL: Concatenação direta de string em query SQL
func exemploSQLInjection(db *sql.DB, nome string) {
	// NUNCA faça isso!
	query := "SELECT * FROM usuarios WHERE nome = '" + nome + "'"
	// Atacante pode injetar: ' OR '1'='1
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	// Processar resultados...
	_ = rows
}

// ✅ CORRETO: Usar prepared statements
func exemploSQLSeguro(db *sql.DB, nome string) {
	query := "SELECT * FROM usuarios WHERE nome = $1"
	rows, err := db.Query(query, nome)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

// Exemplo 2: Command Injection
// ❌ VULNERÁVEL: Execução de comando sem validação
func exemploCommandInjection(input string) {
	// NUNCA faça isso!
	cmd := exec.Command("echo", input)
	// Atacante pode injetar: "; rm -rf /"
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

// ✅ CORRETO: Validar e sanitizar entrada
func exemploCommandSeguro(input string) error {
	// Validar entrada
	if strings.ContainsAny(input, ";&|$`") {
		return fmt.Errorf("entrada inválida: caracteres perigosos detectados")
	}
	
	cmd := exec.Command("echo", input)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

// Exemplo 3: Exposição de Informações Sensíveis
// ❌ VULNERÁVEL: Logar informações sensíveis
func exemploExposicaoLog(senha string) {
	// NUNCA faça isso!
	log.Printf("Usuário fez login com senha: %s", senha)
}

// ✅ CORRETO: Não logar informações sensíveis
func exemploLogSeguro(usuario string) {
	log.Printf("Usuário %s fez login", usuario)
	// Senha nunca é logada
}

// Exemplo 4: Path Traversal
// ❌ VULNERÁVEL: Acesso a arquivos sem validação de path
func exemploPathTraversal(filename string) {
	// NUNCA faça isso!
	file, err := os.Open(filename)
	// Atacante pode usar: "../../etc/passwd"
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// ✅ CORRETO: Validar e sanitizar path
func exemploPathSeguro(filename string) error {
	// Validar que o arquivo está no diretório permitido
	baseDir := "/var/www/uploads"
	fullPath := baseDir + "/" + filename
	
	// Verificar que não há path traversal
	if strings.Contains(fullPath, "..") {
		return fmt.Errorf("path inválido")
	}
	
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// Exemplo 5: XSS (Cross-Site Scripting)
// ❌ VULNERÁVEL: Renderizar entrada do usuário sem sanitização
func exemploXSS(w http.ResponseWriter, r *http.Request) {
	// NUNCA faça isso!
	comentario := r.URL.Query().Get("comentario")
	// Atacante pode injetar: <script>alert('XSS')</script>
	fmt.Fprintf(w, "<p>Comentário: %s</p>", comentario)
}

// ✅ CORRETO: Sanitizar ou escapar HTML
func exemploXSSSeguro(w http.ResponseWriter, r *http.Request) {
	comentario := r.URL.Query().Get("comentario")
	// Escapar caracteres HTML
	comentario = strings.ReplaceAll(comentario, "<", "&lt;")
	comentario = strings.ReplaceAll(comentario, ">", "&gt;")
	fmt.Fprintf(w, "<p>Comentário: %s</p>", comentario)
}

// Exemplo 6: Uso de Dependências Vulneráveis
// Este exemplo demonstra a importância de verificar dependências
// com govulncheck

func main() {
	fmt.Println("⚠️ ATENÇÃO: Este arquivo contém exemplos de código VULNERÁVEL")
	fmt.Println("para fins educacionais. NÃO use em produção!")
	fmt.Println("")
	fmt.Println("Execute govulncheck para verificar vulnerabilidades:")
	fmt.Println("  govulncheck ./...")
	fmt.Println("")
	fmt.Println("Estes exemplos demonstram vulnerabilidades comuns:")
	fmt.Println("  1. SQL Injection")
	fmt.Println("  2. Command Injection")
	fmt.Println("  3. Exposição de Informações Sensíveis")
	fmt.Println("  4. Path Traversal")
	fmt.Println("  5. XSS (Cross-Site Scripting)")
	fmt.Println("")
	fmt.Println("Use os exemplos corretos como referência!")
}

