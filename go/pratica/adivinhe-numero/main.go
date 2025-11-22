package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║         🎯 ADIVINHE O NÚMERO 🎯                         ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("O computador escolheu um número secreto entre 1 e 100.")
	fmt.Println("Tente adivinhar! Você receberá dicas como 'Muito alto' ou 'Muito baixo'.")
	fmt.Println()

	for {
		numeroSecreto := rand.Intn(100) + 1
		tentativas := 0
		maxTentativas := 10

		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Printf("Número secreto gerado! Você tem %d tentativas.\n", maxTentativas)
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()

		for tentativas < maxTentativas {
			tentativas++
			tentativasRestantes := maxTentativas - tentativas

			fmt.Printf("Tentativa %d/%d (restam %d)\n", tentativas, maxTentativas, tentativasRestantes)
			fmt.Print("Digite seu palpite (1-100): ")

			entrada, _ := reader.ReadString('\n')
			entrada = strings.TrimSpace(entrada)

			palpite, err := strconv.Atoi(entrada)
			if err != nil {
				fmt.Println("❌ Entrada inválida! Digite um número entre 1 e 100.")
				tentativas--
				fmt.Println()
				continue
			}

			if palpite < 1 || palpite > 100 {
				fmt.Println("❌ O número deve estar entre 1 e 100!")
				tentativas--
				fmt.Println()
				continue
			}

			if palpite == numeroSecreto {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Printf("║     🎉 PARABÉNS! Você acertou em %d tentativa(s)! 🎉      ║\n", tentativas)
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				fmt.Println()
				break
			} else if palpite < numeroSecreto {
				fmt.Printf("📉 Muito baixo! Tente um número maior.\n")
			} else {
				fmt.Printf("📈 Muito alto! Tente um número menor.\n")
			}

			if tentativasRestantes > 0 {
				fmt.Printf("💡 Dica: O número está entre %d e %d.\n", 
					max(1, numeroSecreto-20), min(100, numeroSecreto+20))
			}

			fmt.Println()
		}

		if tentativas >= maxTentativas {
			fmt.Println()
			fmt.Println("╔══════════════════════════════════════════════════════════╗")
			fmt.Println("║                                                          ║")
			fmt.Printf("║     😢 Fim de jogo! O número secreto era %d! 😢           ║\n", numeroSecreto)
			fmt.Println("║                                                          ║")
			fmt.Println("╚══════════════════════════════════════════════════════════╝")
			fmt.Println()
		}

		fmt.Print("Deseja jogar novamente? (S/N): ")
		resposta, _ := reader.ReadString('\n')
		resposta = strings.TrimSpace(strings.ToUpper(resposta))

		if resposta != "S" && resposta != "SIM" && resposta != "Y" && resposta != "YES" {
			fmt.Println()
			fmt.Println("Obrigado por jogar! Até a próxima! 👋")
			break
		}

		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

