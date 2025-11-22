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

const (
	TAMANHO_TABULEIRO = 8
	NUM_BOMBAS        = 10
)

type Celula struct {
	Revelada bool
	Bomba    bool
	Vizinhas int
}

type Campo struct {
	Grid     [][]Celula
	Reveladas int
	Bombas   int
}

func NovoCampo() *Campo {
	campo := &Campo{
		Grid:     make([][]Celula, TAMANHO_TABULEIRO),
		Reveladas: 0,
		Bombas:   NUM_BOMBAS,
	}

	for i := range campo.Grid {
		campo.Grid[i] = make([]Celula, TAMANHO_TABULEIRO)
	}

	rand.Seed(time.Now().UnixNano())
	bombasColocadas := 0

	for bombasColocadas < NUM_BOMBAS {
		linha := rand.Intn(TAMANHO_TABULEIRO)
		coluna := rand.Intn(TAMANHO_TABULEIRO)

		if !campo.Grid[linha][coluna].Bomba {
			campo.Grid[linha][coluna].Bomba = true
			bombasColocadas++
		}
	}

	campo.contarBombasVizinhas()
	return campo
}

func (c *Campo) contarBombasVizinhas() {
	for i := 0; i < TAMANHO_TABULEIRO; i++ {
		for j := 0; j < TAMANHO_TABULEIRO; j++ {
			if !c.Grid[i][j].Bomba {
				count := 0
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						ni, nj := i+di, j+dj
						if ni >= 0 && ni < TAMANHO_TABULEIRO && nj >= 0 && nj < TAMANHO_TABULEIRO {
							if c.Grid[ni][nj].Bomba {
								count++
							}
						}
					}
				}
				c.Grid[i][j].Vizinhas = count
			}
		}
	}
}

func (c *Campo) Revelar(linha, coluna int) bool {
	if linha < 0 || linha >= TAMANHO_TABULEIRO || coluna < 0 || coluna >= TAMANHO_TABULEIRO {
		return false
	}

	if c.Grid[linha][coluna].Revelada {
		return false
	}

	c.Grid[linha][coluna].Revelada = true
	c.Reveladas++

	if c.Grid[linha][coluna].Bomba {
		return true
	}

	if c.Grid[linha][coluna].Vizinhas == 0 {
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				if di == 0 && dj == 0 {
					continue
				}
				ni, nj := linha+di, coluna+dj
				if ni >= 0 && ni < TAMANHO_TABULEIRO && nj >= 0 && nj < TAMANHO_TABULEIRO {
					if !c.Grid[ni][nj].Revelada && !c.Grid[ni][nj].Bomba {
						c.Revelar(ni, nj)
					}
				}
			}
		}
	}

	return false
}

func (c *Campo) Exibir() {
	fmt.Println("\n    ", strings.Repeat("─", TAMANHO_TABULEIRO*4+1))
	fmt.Print("    │")
	for j := 0; j < TAMANHO_TABULEIRO; j++ {
		fmt.Printf(" %d │", j+1)
	}
	fmt.Println()
	fmt.Println("    ", strings.Repeat("─", TAMANHO_TABULEIRO*4+1))

	for i := 0; i < TAMANHO_TABULEIRO; i++ {
		fmt.Printf("  %d │", i+1)
		for j := 0; j < TAMANHO_TABULEIRO; j++ {
			celula := c.Grid[i][j]
			if !celula.Revelada {
				fmt.Print(" ? │")
			} else if celula.Bomba {
				fmt.Print(" 💣│")
			} else if celula.Vizinhas == 0 {
				fmt.Print("   │")
			} else {
				fmt.Printf(" %d │", celula.Vizinhas)
			}
		}
		fmt.Println()
		fmt.Println("    ", strings.Repeat("─", TAMANHO_TABULEIRO*4+1))
	}
	fmt.Println()
}

func (c *Campo) Venceu() bool {
	return c.Reveladas == (TAMANHO_TABULEIRO*TAMANHO_TABULEIRO - c.Bombas)
}

func (c *Campo) ExibirCompleto() {
	fmt.Println("\n    ", strings.Repeat("─", TAMANHO_TABULEIRO*4+1))
	fmt.Print("    │")
	for j := 0; j < TAMANHO_TABULEIRO; j++ {
		fmt.Printf(" %d │", j+1)
	}
	fmt.Println()
	fmt.Println("    ", strings.Repeat("─", TAMANHO_TABULEIRO*4+1))

	for i := 0; i < TAMANHO_TABULEIRO; i++ {
		fmt.Printf("  %d │", i+1)
		for j := 0; j < TAMANHO_TABULEIRO; j++ {
			celula := c.Grid[i][j]
			if celula.Bomba {
				fmt.Print(" 💣│")
			} else if celula.Vizinhas == 0 {
				fmt.Print("   │")
			} else {
				fmt.Printf(" %d │", celula.Vizinhas)
			}
		}
		fmt.Println()
		fmt.Println("    ", strings.Repeat("─", TAMANHO_TABULEIRO*4+1))
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║         💣 CAÇA AO TESOURO (MINESWEEPER) 💣             ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("Revele todas as células sem bombas!\n")
	fmt.Printf("Tabuleiro: %dx%d | Bombas: %d\n", TAMANHO_TABULEIRO, TAMANHO_TABULEIRO, NUM_BOMBAS)
	fmt.Println()

	for {
		campo := NovoCampo()

		for {
			campo.Exibir()

			fmt.Print("Digite a linha (1-8): ")
			linhaStr, _ := reader.ReadString('\n')
			linhaStr = strings.TrimSpace(linhaStr)
			linha, err := strconv.Atoi(linhaStr)
			if err != nil || linha < 1 || linha > TAMANHO_TABULEIRO {
				fmt.Println("❌ Linha inválida!")
				continue
			}

			fmt.Print("Digite a coluna (1-8): ")
			colunaStr, _ := reader.ReadString('\n')
			colunaStr = strings.TrimSpace(colunaStr)
			coluna, err := strconv.Atoi(colunaStr)
			if err != nil || coluna < 1 || coluna > TAMANHO_TABULEIRO {
				fmt.Println("❌ Coluna inválida!")
				continue
			}

			linha--
			coluna--

			if campo.Revelar(linha, coluna) {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Println("║              💥 BOOM! Você pisou em uma bomba! 💥       ║")
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				fmt.Println()
				campo.ExibirCompleto()
				break
			}

			if campo.Venceu() {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Println("║         🎉 PARABÉNS! Você encontrou o tesouro! 🎉       ║")
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				fmt.Println()
				campo.ExibirCompleto()
				break
			}
		}

		fmt.Print("Deseja jogar novamente? (S/N): ")
		resposta, _ := reader.ReadString('\n')
		resposta = strings.TrimSpace(strings.ToUpper(resposta))

		if resposta != "S" && resposta != "SIM" {
			fmt.Println("Obrigado por jogar! Até a próxima! 👋")
			break
		}
		fmt.Println()
	}
}

