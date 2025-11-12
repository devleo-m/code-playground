package main

import "fmt"

// 1. Bloco de Status de Conexão (iota inicia em 0)
const (
	StatusDesconectado = iota
	StatusConectando
	StatusConectado
	StatusErro
)

// 2. Bloco de Fatores (Inicie o iota para que a primeira constante seja 11)
const (
	_ = 1 << iota
	FATOR_BASE
	FATOR_DOBRO
)

func main() {
	fmt.Println("--- Status de Conexão ---")
	fmt.Printf("Desconectado: %d\n", StatusDesconectado)
	fmt.Printf("Conectando: %d\n", StatusConectando)
	fmt.Printf("Conectado: %d\n", StatusConectado)
	fmt.Printf("Erro: %d\n", StatusErro)

	// Imprima os demais status aqui...

	fmt.Println("\n--- Fatores ---")
	fmt.Printf("Fator Base: %d\n", FATOR_BASE)
	fmt.Printf("Fator Dobro: %d\n", FATOR_DOBRO)

	// Imprima o FATOR_DOBRO aqui...
}
