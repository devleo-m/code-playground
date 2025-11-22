package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	NUM_DADOS      = 5
	MAX_RELANCAMENTOS = 2
)

type Dados struct {
	Valores []int
}

type Combinacao struct {
	Nome  string
	Pontos int
}

func NovoDados() *Dados {
	return &Dados{Valores: make([]int, NUM_DADOS)}
}

func (d *Dados) Rolar() {
	rand.Seed(time.Now().UnixNano())
	for i := range d.Valores {
		d.Valores[i] = rand.Intn(6) + 1
	}
}

func (d *Dados) Relancar(indices []int) {
	rand.Seed(time.Now().UnixNano())
	for _, idx := range indices {
		if idx >= 0 && idx < NUM_DADOS {
			d.Valores[idx] = rand.Intn(6) + 1
		}
	}
}

func (d *Dados) Exibir() {
	fmt.Println("\nDados:")
	for i, valor := range d.Valores {
		fmt.Printf("  [%d]: %d\n", i+1, valor)
	}
}

func (d *Dados) Analisar() Combinacao {
	contagem := make(map[int]int)
	for _, valor := range d.Valores {
		contagem[valor]++
	}

	valores := make([]int, 0, len(contagem))
	for v := range contagem {
		valores = append(valores, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(valores)))

	maxContagem := 0
	for _, count := range contagem {
		if count > maxContagem {
			maxContagem = count
		}
	}

	soma := 0
	for _, valor := range d.Valores {
		soma += valor
	}

	switch {
	case maxContagem == 5:
		return Combinacao{Nome: "Yahtzee (5 iguais)", Pontos: 50}
	case maxContagem == 4:
		return Combinacao{Nome: "Quadra (4 iguais)", Pontos: 30}
	case maxContagem == 3:
		temPar := false
		for _, count := range contagem {
			if count == 2 {
				temPar = true
				break
			}
		}
		if temPar {
			return Combinacao{Nome: "Full House (trinca + par)", Pontos: 25}
		}
		return Combinacao{Nome: "Trinca (3 iguais)", Pontos: soma}
	case maxContagem == 2:
		pares := 0
		for _, count := range contagem {
			if count == 2 {
				pares++
			}
		}
		if pares == 2 {
			return Combinacao{Nome: "Dois Pares", Pontos: soma}
		}
		return Combinacao{Nome: "Par (2 iguais)", Pontos: soma}
	default:
		sort.Ints(d.Valores)
		sequencia := true
		for i := 1; i < len(d.Valores); i++ {
			if d.Valores[i] != d.Valores[i-1]+1 {
				sequencia = false
				break
			}
		}
		if sequencia {
			return Combinacao{Nome: "Sequência", Pontos: 40}
		}
		return Combinacao{Nome: "Nada", Pontos: soma}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║         🎲 PÔQUER-DADOS (YAHTZEE) 🎲                    ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Role os dados e tente obter as melhores combinações!")
	fmt.Printf("Você tem %d relançamentos por rodada.\n", MAX_RELANCAMENTOS)
	fmt.Println()

	pontuacaoTotal := 0

	for {
		dados := NovoDados()
		dados.Rolar()
		dados.Exibir()

		relancamentos := 0

		for relancamentos < MAX_RELANCAMENTOS {
			fmt.Printf("\nRelançamentos restantes: %d/%d\n", MAX_RELANCAMENTOS-relancamentos, MAX_RELANCAMENTOS)
			fmt.Print("Deseja relançar alguns dados? (S/N): ")
			resposta, _ := reader.ReadString('\n')
			resposta = strings.TrimSpace(strings.ToUpper(resposta))

			if resposta != "S" && resposta != "SIM" {
				break
			}

			fmt.Print("Digite os números dos dados para relançar (ex: 1 3 5): ")
			entrada, _ := reader.ReadString('\n')
			entrada = strings.TrimSpace(entrada)

			if entrada == "" {
				fmt.Println("❌ Nenhum dado selecionado!")
				continue
			}

			partes := strings.Fields(entrada)
			indices := make([]int, 0)

			for _, parte := range partes {
				idx, err := strconv.Atoi(parte)
				if err != nil || idx < 1 || idx > NUM_DADOS {
					fmt.Printf("❌ '%s' não é um número de dado válido (1-%d)!\n", parte, NUM_DADOS)
					continue
				}
				indices = append(indices, idx-1)
			}

			if len(indices) > 0 {
				dados.Relancar(indices)
				dados.Exibir()
				relancamentos++
			}
		}

		combinacao := dados.Analisar()
		fmt.Println()
		fmt.Println(strings.Repeat("═", 60))
		fmt.Printf("Resultado: %s\n", combinacao.Nome)
		fmt.Printf("Pontos: %d\n", combinacao.Pontos)
		fmt.Println(strings.Repeat("═", 60))

		pontuacaoTotal += combinacao.Pontos
		fmt.Printf("\nPontuação total: %d\n", pontuacaoTotal)

		fmt.Print("\nDeseja jogar outra rodada? (S/N): ")
		resposta, _ := reader.ReadString('\n')
		resposta = strings.TrimSpace(strings.ToUpper(resposta))

		if resposta != "S" && resposta != "SIM" {
			fmt.Println()
			fmt.Println("╔══════════════════════════════════════════════════════════╗")
			fmt.Println("║                                                          ║")
			fmt.Printf("║         🏆 PONTUAÇÃO FINAL: %d 🏆                  ║\n", pontuacaoTotal)
			fmt.Println("║                                                          ║")
			fmt.Println("╚══════════════════════════════════════════════════════════╝")
			fmt.Println()
			fmt.Println("Obrigado por jogar! Até a próxima! 👋")
			break
		}
		fmt.Println()
	}
}

