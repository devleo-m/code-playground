package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Torre struct {
	Nome  string
	Discos []int
}

type Jogo struct {
	TorreA *Torre
	TorreB *Torre
	TorreC *Torre
	Movimentos int
}

func NovaTorre(nome string) *Torre {
	return &Torre{
		Nome:   nome,
		Discos: []int{},
	}
}

func NovoJogo(numDiscos int) *Jogo {
	torreA := NovaTorre("A")
	for i := numDiscos; i >= 1; i-- {
		torreA.Discos = append(torreA.Discos, i)
	}

	return &Jogo{
		TorreA:     torreA,
		TorreB:     NovaTorre("B"),
		TorreC:     NovaTorre("C"),
		Movimentos: 0,
	}
}

func (j *Jogo) Exibir() {
	maxDiscos := 0
	for _, disco := range j.TorreA.Discos {
		if disco > maxDiscos {
			maxDiscos = disco
		}
	}
	for _, disco := range j.TorreB.Discos {
		if disco > maxDiscos {
			maxDiscos = disco
		}
	}
	for _, disco := range j.TorreC.Discos {
		if disco > maxDiscos {
			maxDiscos = disco
		}
	}

	altura := maxDiscos + 1

	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("                    TORRE DE HANOI")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()

	for nivel := altura - 1; nivel >= 0; nivel-- {
		for _, torre := range []*Torre{j.TorreA, j.TorreB, j.TorreC} {
			fmt.Print("  ")
			if nivel < len(torre.Discos) {
				disco := torre.Discos[nivel]
				largura := disco*2 + 1
				espacos := (maxDiscos*2 + 1 - largura) / 2
				fmt.Print(strings.Repeat(" ", espacos))
				fmt.Print(strings.Repeat("█", largura))
				fmt.Print(strings.Repeat(" ", espacos))
			} else {
				fmt.Print(strings.Repeat(" ", maxDiscos*2+1))
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}

	fmt.Print("  ")
	for range []*Torre{j.TorreA, j.TorreB, j.TorreC} {
		fmt.Print(strings.Repeat("═", maxDiscos*2+1))
		fmt.Print("  ")
	}
	fmt.Println()

	fmt.Print("  ")
	for _, torre := range []*Torre{j.TorreA, j.TorreB, j.TorreC} {
		espacos := (maxDiscos*2 + 1 - 1) / 2
		fmt.Print(strings.Repeat(" ", espacos))
		fmt.Print(torre.Nome)
		fmt.Print(strings.Repeat(" ", espacos))
		fmt.Print("  ")
	}
	fmt.Println()
	fmt.Println()

	fmt.Printf("Movimentos: %d\n", j.Movimentos)
	fmt.Println(strings.Repeat("═", 60))
}

func (j *Jogo) Mover(origem, destino string) bool {
	var torreOrigem, torreDestino *Torre

	switch origem {
	case "A", "a":
		torreOrigem = j.TorreA
	case "B", "b":
		torreOrigem = j.TorreB
	case "C", "c":
		torreOrigem = j.TorreC
	default:
		return false
	}

	switch destino {
	case "A", "a":
		torreDestino = j.TorreA
	case "B", "b":
		torreDestino = j.TorreB
	case "C", "c":
		torreDestino = j.TorreC
	default:
		return false
	}

	if len(torreOrigem.Discos) == 0 {
		return false
	}

	disco := torreOrigem.Discos[len(torreOrigem.Discos)-1]

	if len(torreDestino.Discos) > 0 {
		discoTopo := torreDestino.Discos[len(torreDestino.Discos)-1]
		if disco > discoTopo {
			return false
		}
	}

	torreOrigem.Discos = torreOrigem.Discos[:len(torreOrigem.Discos)-1]
	torreDestino.Discos = append(torreDestino.Discos, disco)
	j.Movimentos++

	return true
}

func (j *Jogo) Venceu(numDiscos int) bool {
	return len(j.TorreC.Discos) == numDiscos
}

func (j *Jogo) Resolver(numDiscos int, origem, auxiliar, destino *Torre) {
	if numDiscos == 1 {
		disco := origem.Discos[len(origem.Discos)-1]
		origem.Discos = origem.Discos[:len(origem.Discos)-1]
		destino.Discos = append(destino.Discos, disco)
		j.Movimentos++
		return
	}

	j.Resolver(numDiscos-1, origem, destino, auxiliar)
	disco := origem.Discos[len(origem.Discos)-1]
	origem.Discos = origem.Discos[:len(origem.Discos)-1]
	destino.Discos = append(destino.Discos, disco)
	j.Movimentos++
	j.Resolver(numDiscos-1, auxiliar, origem, destino)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║              🗼 TORRE DE HANOI 🗼                         ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Mova todos os discos da torre A para a torre C.")
	fmt.Println("Regras:")
	fmt.Println("  - Você só pode mover um disco por vez")
	fmt.Println("  - Você só pode mover o disco do topo")
	fmt.Println("  - Um disco maior não pode ficar sobre um menor")
	fmt.Println()

	for {
		fmt.Print("Quantos discos deseja usar? (1-7): ")
		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)
		numDiscos, err := strconv.Atoi(entrada)

		if err != nil || numDiscos < 1 || numDiscos > 7 {
			fmt.Println("❌ Número inválido! Use um número entre 1 e 7.")
			continue
		}

		jogo := NovoJogo(numDiscos)

		for {
			jogo.Exibir()

			if jogo.Venceu(numDiscos) {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Println("║         🎉 PARABÉNS! Você completou a torre! 🎉         ║")
				fmt.Printf("║         Movimentos: %d                                    ║\n", jogo.Movimentos)
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				fmt.Println()
				break
			}

			fmt.Print("Mover disco de (A/B/C) para (A/B/C) ou 'resolver' para solução automática: ")
			entrada, _ := reader.ReadString('\n')
			entrada = strings.TrimSpace(entrada)

			if strings.ToLower(entrada) == "resolver" {
				fmt.Println("\n🤖 Resolvendo automaticamente...")
				jogo = NovoJogo(numDiscos)
				jogo.Resolver(numDiscos, jogo.TorreA, jogo.TorreB, jogo.TorreC)
				jogo.Exibir()
				fmt.Println("✅ Solução completa!")
				break
			}

			partes := strings.Fields(entrada)
			if len(partes) != 2 {
				fmt.Println("❌ Formato inválido! Use: A C (para mover de A para C)")
				continue
			}

			if !jogo.Mover(partes[0], partes[1]) {
				fmt.Println("❌ Movimento inválido! Verifique as regras.")
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

