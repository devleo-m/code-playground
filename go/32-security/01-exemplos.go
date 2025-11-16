package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Exemplo 1: Uso básico de crypto/rand
// Este exemplo demonstra uso correto de crypto/rand
func exemploCryptoRand() {
	buffer := make([]byte, 32)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Random bytes: %x\n", buffer)
}

// Exemplo 2: Servidor HTTP básico
// Este exemplo demonstra um servidor HTTP simples
func exemploServidorHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, World!",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Exemplo 3: Handler com validação de entrada
// Este exemplo demonstra validação básica de entrada
func exemploHandlerValidado(w http.ResponseWriter, r *http.Request) {
	// Validar método HTTP
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Validar parâmetros de query
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Parâmetro 'q' é obrigatório", http.StatusBadRequest)
		return
	}

	// Processar requisição
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"query": query,
		"status": "ok",
	})
}

// Exemplo 4: Geração de token seguro
// Este exemplo demonstra geração de token usando crypto/rand
func gerarTokenSeguro() (string, error) {
	buffer := make([]byte, 32)
	if _, err := rand.Read(buffer); err != nil {
		return "", fmt.Errorf("erro ao gerar token: %w", err)
	}
	return fmt.Sprintf("%x", buffer), nil
}

// Exemplo 5: Handler com tratamento de erro adequado
// Este exemplo demonstra tratamento de erro adequado
func exemploHandlerComErro(w http.ResponseWriter, r *http.Request) {
	token, err := gerarTokenSeguro()
	if err != nil {
		// Log do erro (sem expor detalhes sensíveis)
		log.Printf("Erro ao gerar token: %v", err)
		
		// Resposta genérica ao cliente (não expor detalhes internos)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"status": "ok",
	})
}

func main() {
	// Exemplo de uso
	fmt.Println("Exemplos de código Go seguro")
	
	// Gerar token seguro
	token, err := gerarTokenSeguro()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token gerado: %s\n", token)
	
	// Nota: Para testar o servidor HTTP, descomente a linha abaixo
	// exemploServidorHTTP()
}

