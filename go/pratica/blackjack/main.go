package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Naipe string

const (
	PAUS    Naipe = "♣"
	COPAS   Naipe = "♥"
	OUROS   Naipe = "♦"
	ESPADAS Naipe = "♠"
)

type Carta struct {
	Valor int
	Naipe Naipe
	Nome  string
}

type Baralho struct {
	Cartas []Carta
}

type Jogador struct {
	Nome      string
	Cartas    []Carta
	Pontuacao int
}

func NovoBaralho() *Baralho {
	baralho := &Baralho{}
	naipes := []Naipe{PAUS, COPAS, OUROS, ESPADAS}
	nomes := []string{"Ás", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Valete", "Dama", "Rei"}

	for _, naipe := range naipes {
		for i, nome := range nomes {
			valor := i + 1
			if valor > 10 {
				valor = 10
			}
			baralho.Cartas = append(baralho.Cartas, Carta{
				Valor: valor,
				Naipe: naipe,
				Nome:  nome,
			})
		}
	}

	baralho.Embaralhar()
	return baralho
}

func (b *Baralho) Embaralhar() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(b.Cartas), func(i, j int) {
		b.Cartas[i], b.Cartas[j] = b.Cartas[j], b.Cartas[i]
	})
}

func (b *Baralho) Comprar() Carta {
	if len(b.Cartas) == 0 {
		b = NovoBaralho()
	}
	carta := b.Cartas[0]
	b.Cartas = b.Cartas[1:]
	return carta
}

func (j *Jogador) AdicionarCarta(carta Carta) {
	j.Cartas = append(j.Cartas, carta)
	j.Pontuacao = j.CalcularPontuacao()
}

func (j *Jogador) CalcularPontuacao() int {
	pontuacao := 0
	ases := 0

	for _, carta := range j.Cartas {
		if carta.Valor == 1 {
			ases++
			pontuacao += 11
		} else {
			pontuacao += carta.Valor
		}
	}

	for ases > 0 && pontuacao > 21 {
		pontuacao -= 10
		ases--
	}

	return pontuacao
}

func (j *Jogador) ExibirCartas(mostrarTodas bool) {
	if len(j.Cartas) == 0 {
		return
	}

	fmt.Printf("\n%s:\n", j.Nome)
	for i, carta := range j.Cartas {
		if !mostrarTodas && i == 0 && j.Nome == "Dealer" {
			fmt.Println("  [Carta oculta]")
		} else {
			fmt.Printf("  %s %s (%d pontos)\n", carta.Nome, carta.Naipe, carta.Valor)
		}
	}

	if mostrarTodas || j.Nome != "Dealer" {
		fmt.Printf("  Pontuação: %d\n", j.Pontuacao)
	}
}

func (j *Jogador) Resetar() {
	j.Cartas = []Carta{}
	j.Pontuacao = 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║              🃏 BLACKJACK (21) 🃏                       ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Objetivo: Ficar o mais próximo de 21 sem ultrapassar!")
	fmt.Println("O dealer deve comprar até ter pelo menos 17 pontos.")
	fmt.Println()

	jogador := &Jogador{Nome: "Jogador"}
	dealer := &Jogador{Nome: "Dealer"}

	for {
		baralho := NovoBaralho()
		jogador.Resetar()
		dealer.Resetar()

		fmt.Println(strings.Repeat("═", 60))
		fmt.Println("Nova rodada!")
		fmt.Println(strings.Repeat("═", 60))

		jogador.AdicionarCarta(baralho.Comprar())
		dealer.AdicionarCarta(baralho.Comprar())
		jogador.AdicionarCarta(baralho.Comprar())
		dealer.AdicionarCarta(baralho.Comprar())

		jogador.ExibirCartas(true)
		dealer.ExibirCartas(false)

		for jogador.Pontuacao < 21 {
			fmt.Print("\nDeseja comprar outra carta? (S/N): ")
			resposta, _ := reader.ReadString('\n')
			resposta = strings.TrimSpace(strings.ToUpper(resposta))

			if resposta == "S" || resposta == "SIM" {
				carta := baralho.Comprar()
				jogador.AdicionarCarta(carta)
				jogador.ExibirCartas(true)
				dealer.ExibirCartas(false)

				if jogador.Pontuacao > 21 {
					fmt.Println("\n💥 Você estourou! Mais de 21 pontos.")
					break
				}
			} else {
				break
			}
		}

		if jogador.Pontuacao <= 21 {
			fmt.Println("\n🤖 Turno do dealer...")
			time.Sleep(1 * time.Second)

			for dealer.Pontuacao < 17 {
				carta := baralho.Comprar()
				dealer.AdicionarCarta(carta)
				fmt.Printf("Dealer comprou: %s %s\n", carta.Nome, carta.Naipe)
				time.Sleep(1 * time.Second)
			}

			dealer.ExibirCartas(true)
		}

		fmt.Println("\n" + strings.Repeat("═", 60))
		fmt.Println("RESULTADO:")
		fmt.Println(strings.Repeat("═", 60))

		if jogador.Pontuacao > 21 {
			fmt.Println("😢 Você perdeu! Estourou 21.")
		} else if dealer.Pontuacao > 21 {
			fmt.Println("🎉 Você venceu! Dealer estourou 21.")
		} else if jogador.Pontuacao > dealer.Pontuacao {
			fmt.Println("🎉 Você venceu! Sua pontuação é maior.")
		} else if dealer.Pontuacao > jogador.Pontuacao {
			fmt.Println("😢 Você perdeu! Dealer tem pontuação maior.")
		} else {
			fmt.Println("🤝 Empate! Pontuações iguais.")
		}

		fmt.Printf("Sua pontuação: %d | Pontuação do dealer: %d\n", jogador.Pontuacao, dealer.Pontuacao)
		fmt.Println(strings.Repeat("═", 60))

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

