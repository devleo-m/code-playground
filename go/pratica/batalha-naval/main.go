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

const TAMANHO_TABULEIRO = 5

type EstadoCelula int

const (
	AGUA EstadoCelula = iota
	NAVIO
	TIRO_AGUA
	TIRO_NAVIO
)

type Tabuleiro struct {
	Grid [][]EstadoCelula
}

type Jogo struct {
	TabuleiroJogador Tabuleiro
	TabuleiroPC      Tabuleiro
	NaviosJogador    int
	NaviosPC          int
}

func NovoTabuleiro() Tabuleiro {
	grid := make([][]EstadoCelula, TAMANHO_TABULEIRO)
	for i := range grid {
		grid[i] = make([]EstadoCelula, TAMANHO_TABULEIRO)
	}
	return Tabuleiro{Grid: grid}
}

func NovoJogo() *Jogo {
	jogo := &Jogo{
		TabuleiroJogador: NovoTabuleiro(),
		TabuleiroPC:      NovoTabuleiro(),
		NaviosJogador:    3,
		NaviosPC:         3,
	}

	jogo.posicionarNaviosPC()
	return jogo
}

func (j *Jogo) posicionarNaviosPC() {
	rand.Seed(time.Now().UnixNano())
	naviosColocados := 0

	for naviosColocados < j.NaviosPC {
		linha := rand.Intn(TAMANHO_TABULEIRO)
		coluna := rand.Intn(TAMANHO_TABULEIRO)

		if j.TabuleiroPC.Grid[linha][coluna] == AGUA {
			j.TabuleiroPC.Grid[linha][coluna] = NAVIO
			naviosColocados++
		}
	}
}

func (j *Jogo) posicionarNaviosJogador() {
	reader := bufio.NewReader(os.Stdin)
	naviosColocados := 0

	fmt.Println("\nPosicione seus 3 navios:")
	for naviosColocados < j.NaviosJogador {
		fmt.Printf("\nNavio %d/%d:\n", naviosColocados+1, j.NaviosJogador)
		fmt.Print("Digite a linha (1-5): ")
		linhaStr, _ := reader.ReadString('\n')
		linhaStr = strings.TrimSpace(linhaStr)
		linha, err := strconv.Atoi(linhaStr)
		if err != nil || linha < 1 || linha > TAMANHO_TABULEIRO {
			fmt.Println("❌ Linha inválida!")
			continue
		}

		fmt.Print("Digite a coluna (1-5): ")
		colunaStr, _ := reader.ReadString('\n')
		colunaStr = strings.TrimSpace(colunaStr)
		coluna, err := strconv.Atoi(colunaStr)
		if err != nil || coluna < 1 || coluna > TAMANHO_TABULEIRO {
			fmt.Println("❌ Coluna inválida!")
			continue
		}

		linha--
		coluna--

		if j.TabuleiroJogador.Grid[linha][coluna] == NAVIO {
			fmt.Println("❌ Já existe um navio nessa posição!")
			continue
		}

		j.TabuleiroJogador.Grid[linha][coluna] = NAVIO
		naviosColocados++
		fmt.Println("✅ Navio posicionado!")
	}
}

func (j *Jogo) ExibirTabuleiros() {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("                    SEU TABULEIRO")
	fmt.Println(strings.Repeat("═", 70))
	j.exibirTabuleiro(j.TabuleiroJogador, true)

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("                  TABULEIRO DO OPONENTE")
	fmt.Println(strings.Repeat("═", 70))
	j.exibirTabuleiro(j.TabuleiroPC, false)
}

func (j *Jogo) exibirTabuleiro(tab Tabuleiro, mostrarNavios bool) {
	fmt.Print("    ")
	for j := 0; j < TAMANHO_TABULEIRO; j++ {
		fmt.Printf("  %d ", j+1)
	}
	fmt.Println()
	fmt.Println("    " + strings.Repeat("─", TAMANHO_TABULEIRO*4+1))

	for i := 0; i < TAMANHO_TABULEIRO; i++ {
		fmt.Printf("  %d │", i+1)
		for j := 0; j < TAMANHO_TABULEIRO; j++ {
			switch tab.Grid[i][j] {
			case AGUA:
				fmt.Print(" ~ │")
			case NAVIO:
				if mostrarNavios {
					fmt.Print(" 🚢│")
				} else {
					fmt.Print(" ~ │")
				}
			case TIRO_AGUA:
				fmt.Print(" ○ │")
			case TIRO_NAVIO:
				fmt.Print(" 💥│")
			}
		}
		fmt.Println()
		fmt.Println("    " + strings.Repeat("─", TAMANHO_TABULEIRO*4+1))
	}
}

func (j *Jogo) AtirarJogador(linha, coluna int) bool {
	if linha < 0 || linha >= TAMANHO_TABULEIRO || coluna < 0 || coluna >= TAMANHO_TABULEIRO {
		return false
	}

	if j.TabuleiroPC.Grid[linha][coluna] == TIRO_AGUA || j.TabuleiroPC.Grid[linha][coluna] == TIRO_NAVIO {
		return false
	}

	if j.TabuleiroPC.Grid[linha][coluna] == NAVIO {
		j.TabuleiroPC.Grid[linha][coluna] = TIRO_NAVIO
		j.NaviosPC--
		return true
	}

	j.TabuleiroPC.Grid[linha][coluna] = TIRO_AGUA
	return false
}

func (j *Jogo) AtirarPC() bool {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		linha := rand.Intn(TAMANHO_TABULEIRO)
		coluna := rand.Intn(TAMANHO_TABULEIRO)

		if j.TabuleiroJogador.Grid[linha][coluna] != TIRO_AGUA && j.TabuleiroJogador.Grid[linha][coluna] != TIRO_NAVIO {
			if j.TabuleiroJogador.Grid[linha][coluna] == NAVIO {
				j.TabuleiroJogador.Grid[linha][coluna] = TIRO_NAVIO
				j.NaviosJogador--
				return true
			}
			j.TabuleiroJogador.Grid[linha][coluna] = TIRO_AGUA
			return false
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║            ⚓ BATALHA NAVAL SIMPLIFICADA ⚓              ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Posicione seus navios e tente afundar os navios do oponente!")
	fmt.Println()

	for {
		jogo := NovoJogo()
		jogo.posicionarNaviosJogador()

		fmt.Println("\n🎮 Iniciando batalha...")
		time.Sleep(1 * time.Second)

		for {
			jogo.ExibirTabuleiros()
			fmt.Printf("\nSeus navios: %d | Navios do oponente: %d\n", jogo.NaviosJogador, jogo.NaviosPC)

			if jogo.NaviosJogador == 0 {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Println("║              😢 VOCÊ PERDEU! 😢                         ║")
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				break
			}

			if jogo.NaviosPC == 0 {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Println("║              🎉 VOCÊ VENCEU! 🎉                        ║")
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				break
			}

			fmt.Println("\n🎯 Seu turno de atirar:")
			fmt.Print("Digite a linha (1-5): ")
			linhaStr, _ := reader.ReadString('\n')
			linhaStr = strings.TrimSpace(linhaStr)
			linha, err := strconv.Atoi(linhaStr)
			if err != nil || linha < 1 || linha > TAMANHO_TABULEIRO {
				fmt.Println("❌ Linha inválida!")
				continue
			}

			fmt.Print("Digite a coluna (1-5): ")
			colunaStr, _ := reader.ReadString('\n')
			colunaStr = strings.TrimSpace(colunaStr)
			coluna, err := strconv.Atoi(colunaStr)
			if err != nil || coluna < 1 || coluna > TAMANHO_TABULEIRO {
				fmt.Println("❌ Coluna inválida!")
				continue
			}

			linha--
			coluna--

			if !jogo.AtirarJogador(linha, coluna) {
				if jogo.TabuleiroPC.Grid[linha][coluna] == TIRO_AGUA || jogo.TabuleiroPC.Grid[linha][coluna] == TIRO_NAVIO {
					fmt.Println("❌ Você já atirou nessa posição!")
				} else {
					fmt.Println("💧 Água! Você errou.")
				}
			} else {
				fmt.Println("💥 ACERTOU! Você afundou um navio!")
			}

			if jogo.NaviosPC == 0 {
				continue
			}

			fmt.Println("\n🤖 Turno do computador...")
			time.Sleep(1 * time.Second)

			if jogo.AtirarPC() {
				fmt.Println("💥 O computador acertou um dos seus navios!")
			} else {
				fmt.Println("💧 O computador errou.")
			}

			time.Sleep(1 * time.Second)
		}

		fmt.Print("\nDeseja jogar novamente? (S/N): ")
		resposta, _ := reader.ReadString('\n')
		resposta = strings.TrimSpace(strings.ToUpper(resposta))

		if resposta != "S" && resposta != "SIM" {
			fmt.Println("Obrigado por jogar! Até a próxima! 👋")
			break
		}
		fmt.Println()
	}
}

