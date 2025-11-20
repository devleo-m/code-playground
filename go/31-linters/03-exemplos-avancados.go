package main

import (
	"context"
	"fmt"
	"time"
)

// User representa um usuário do sistema
type User struct {
	ID   int
	Name string
}

// GetUserByID busca um usuário pelo ID
// Context deve ser o primeiro parâmetro - Revive detecta se não for
func GetUserByID(ctx context.Context, id int) (*User, error) {
	// Verificar se context foi cancelado
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// Simular busca
		time.Sleep(100 * time.Millisecond)
		return &User{ID: id, Name: "João"}, nil
	}
}

// ProcessUsers processa uma lista de usuários
func ProcessUsers(users []User) {
	// Range loop - Staticcheck detecta se você modificar cópia ao invés do original
	for i := range users {
		users[i].Name = "Processed: " + users[i].Name
	}
}

// CalculateAverage calcula a média de uma slice de números
func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := GetUserByID(ctx, 1)
	if err != nil {
		fmt.Printf("Erro ao buscar usuário: %v\n", err)
		return
	}

	fmt.Printf("Usuário encontrado: %+v\n", user)

	users := []User{
		{ID: 1, Name: "João"},
		{ID: 2, Name: "Maria"},
	}

	ProcessUsers(users)
	fmt.Printf("Usuários processados: %+v\n", users)

	numbers := []float64{1.5, 2.5, 3.5, 4.5}
	average := CalculateAverage(numbers)
	fmt.Printf("Média: %.2f\n", average)
}



