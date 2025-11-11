package main

import "fmt"

// Declaração de constantes no nível do pacote.
// O valor deve ser conhecido em tempo de compilação.
const Pi = 3.14159
const URLBase = "https://minhaapi.com"

// Bloco de constantes para melhor organização.
const (
	StatusOK         = 200
	StatusNotFound   = 404
	StatusInternalError = 500
)

// --- Usando `iota` para criar enumerações ---

// iota começa em 0 no início do bloco e incrementa a cada constante.
const (
	Domingo = iota // 0
	Segunda        // 1
	Terca          // 2
	Quarta         // 3
	Quinta         // 4
	Sexta          // 5
	Sabado         // 6
)

// Você pode usar expressões com iota.
// Aqui, criamos potências de 10.
const (
	B  = 1 << (10 * iota) // 1 << (10 * 0) = 1
	KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
	GB = 1 << (10 * iota) // 1 << (10 * 3) = 1073741824
)

func main() {
	fmt.Println("O valor de Pi é:", Pi)
	fmt.Println("URL Base:", URLBase)
	fmt.Println("Status OK:", StatusOK)

	// Tentar reatribuir uma constante causa um erro de compilação.
	// Pi = 3.14 // Error: cannot assign to Pi

	fmt.Println("--------------------------------")

	fmt.Println("Dias da semana:")
	fmt.Println("Domingo:", Domingo)
	fmt.Println("Segunda:", Segunda)
	fmt.Println("Sábado:", Sabado)

	fmt.Println("--------------------------------")

	fmt.Println("Unidades de armazenamento (em bytes):")
	fmt.Println("Kilobyte (KB):", KB)
	fmt.Println("Megabyte (MB):", MB)
	fmt.Println("Gigabyte (GB):", GB)
}
