package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Exemplos básicos de error handling
func exemploBasico() {
	result, err := dividir(10, 2)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("Resultado: %d\n", result)
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}

// Exemplo com fmt.Errorf
func exemploFmtErrorf() {
	err := validarIdade(-5)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}
}

func validarIdade(idade int) error {
	if idade < 0 {
		return fmt.Errorf("idade inválida: %d (deve ser positiva)", idade)
	}
	if idade > 150 {
		return fmt.Errorf("idade inválida: %d (muito alta)", idade)
	}
	return nil
}

// Exemplo de error wrapping
func exemploErrorWrapping() {
	err := processarArquivo("arquivo.txt")
	if err != nil {
		fmt.Printf("Erro ao processar: %v\n", err)
		fmt.Printf("Erro original: %v\n", errors.Unwrap(err))
	}
}

func processarArquivo(nome string) error {
	arquivo, err := abrirArquivo(nome)
	if err != nil {
		return fmt.Errorf("falha ao processar arquivo %s: %w", nome, err)
	}
	defer arquivo.Close()
	// processamento...
	return nil
}

func abrirArquivo(nome string) (*os.File, error) {
	return nil, errors.New("arquivo não encontrado")
}

// Exemplo de sentinel errors
var (
	ErrUsuarioNaoEncontrado = errors.New("usuário não encontrado")
	ErrSenhaInvalida        = errors.New("senha inválida")
)

func exemploSentinelErrors() {
	err := autenticar("usuario", "senha")
	if errors.Is(err, ErrUsuarioNaoEncontrado) {
		fmt.Println("Usuário não existe no sistema")
	} else if errors.Is(err, ErrSenhaInvalida) {
		fmt.Println("Senha incorreta")
	} else if err != nil {
		fmt.Printf("Erro inesperado: %v\n", err)
	}
}

func autenticar(usuario, senha string) error {
	// simulação
	if usuario != "admin" {
		return ErrUsuarioNaoEncontrado
	}
	if senha != "12345" {
		return ErrSenhaInvalida
	}
	return nil
}

// Exemplo de errors.As
type ErroValidacao struct {
	Campo   string
	Mensagem string
}

func (e ErroValidacao) Error() string {
	return fmt.Sprintf("erro no campo %s: %s", e.Campo, e.Mensagem)
}

func exemploErrorsAs() {
	err := validarFormulario("", 0)
	var errValidacao ErroValidacao
	if errors.As(err, &errValidacao) {
		fmt.Printf("Campo com erro: %s\n", errValidacao.Campo)
		fmt.Printf("Mensagem: %s\n", errValidacao.Mensagem)
	}
}

func validarFormulario(nome string, idade int) error {
	if nome == "" {
		return ErroValidacao{Campo: "nome", Mensagem: "não pode ser vazio"}
	}
	if idade < 18 {
		return ErroValidacao{Campo: "idade", Mensagem: "deve ser maior que 18"}
	}
	return nil
}

// Exemplo de panic e recover
func exemploPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recuperado: %v\n", r)
		}
	}()
	
	operacaoPerigosa()
	fmt.Println("Esta linha não será executada se houver panic")
}

func operacaoPerigosa() {
	panic("algo deu muito errado!")
}

// Exemplo prático: leitura de arquivo
func exemploLeituraArquivo() {
	conteudo, err := lerArquivo("dados.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Arquivo não existe")
		} else if errors.Is(err, io.EOF) {
			fmt.Println("Arquivo vazio")
		} else {
			fmt.Printf("Erro ao ler arquivo: %v\n", err)
		}
		return
	}
	fmt.Printf("Conteúdo: %s\n", conteudo)
}

func lerArquivo(nome string) (string, error) {
	arquivo, err := os.Open(nome)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer arquivo.Close()
	
	conteudo, err := io.ReadAll(arquivo)
	if err != nil {
		return "", fmt.Errorf("erro ao ler conteúdo: %w", err)
	}
	
	return string(conteudo), nil
}

func main() {
	fmt.Println("=== Exemplo Básico ===")
	exemploBasico()
	
	fmt.Println("\n=== Exemplo fmt.Errorf ===")
	exemploFmtErrorf()
	
	fmt.Println("\n=== Exemplo Error Wrapping ===")
	exemploErrorWrapping()
	
	fmt.Println("\n=== Exemplo Sentinel Errors ===")
	exemploSentinelErrors()
	
	fmt.Println("\n=== Exemplo errors.As ===")
	exemploErrorsAs()
	
	fmt.Println("\n=== Exemplo Panic/Recover ===")
	exemploPanicRecover()
	fmt.Println("Programa continuou após recover")
	
	fmt.Println("\n=== Exemplo Leitura de Arquivo ===")
	exemploLeituraArquivo()
}








