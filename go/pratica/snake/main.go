package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	LARGURA = 20
	ALTURA  = 10
)

type Posicao struct {
	X, Y int
}

type Jogo struct {
	Snake      []Posicao
	Comida     Posicao
	Direcao    Posicao
	Pontuacao  int
	GameOver   bool
}

func LimparTela() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func NovoJogo() *Jogo {
	jogo := &Jogo{
		Snake:     []Posicao{{X: LARGURA / 2, Y: ALTURA / 2}},
		Direcao:   Posicao{X: 1, Y: 0},
		Pontuacao: 0,
		GameOver:  false,
	}
	jogo.GerarComida()
	return jogo
}

func (j *Jogo) GerarComida() {
	rand.Seed(time.Now().UnixNano())
	for {
		j.Comida = Posicao{
			X: rand.Intn(LARGURA),
			Y: rand.Intn(ALTURA),
		}
		colisao := false
		for _, segmento := range j.Snake {
			if segmento.X == j.Comida.X && segmento.Y == j.Comida.Y {
				colisao = true
				break
			}
		}
		if !colisao {
			break
		}
	}
}

func (j *Jogo) Atualizar() {
	if j.GameOver {
		return
	}

	cabeca := j.Snake[0]
	novaCabeca := Posicao{
		X: cabeca.X + j.Direcao.X,
		Y: cabeca.Y + j.Direcao.Y,
	}

	if novaCabeca.X < 0 || novaCabeca.X >= LARGURA || novaCabeca.Y < 0 || novaCabeca.Y >= ALTURA {
		j.GameOver = true
		return
	}

	for _, segmento := range j.Snake {
		if novaCabeca.X == segmento.X && novaCabeca.Y == segmento.Y {
			j.GameOver = true
			return
		}
	}

	j.Snake = append([]Posicao{novaCabeca}, j.Snake...)

	if novaCabeca.X == j.Comida.X && novaCabeca.Y == j.Comida.Y {
		j.Pontuacao++
		j.GerarComida()
	} else {
		j.Snake = j.Snake[:len(j.Snake)-1]
	}
}

func (j *Jogo) Exibir() {
	LimparTela()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║              🐍 SNAKE (JOGO DA COBRINHA) 🐍             ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Printf("\nPontuação: %d\n", j.Pontuacao)
	fmt.Println()

	fmt.Print("  ")
	for x := 0; x < LARGURA+2; x++ {
		fmt.Print("═")
	}
	fmt.Println()

	for y := 0; y < ALTURA; y++ {
		fmt.Print("  ║")
		for x := 0; x < LARGURA; x++ {
			celula := " "
			if x == j.Comida.X && y == j.Comida.Y {
				celula = "@"
			} else {
				for i, segmento := range j.Snake {
					if segmento.X == x && segmento.Y == y {
						if i == 0 {
							celula = "#"
						} else {
							celula = "o"
						}
						break
					}
				}
			}
			fmt.Print(celula)
		}
		fmt.Println("║")
	}

	fmt.Print("  ")
	for x := 0; x < LARGURA+2; x++ {
		fmt.Print("═")
	}
	fmt.Println()

	if j.GameOver {
		fmt.Println("\n💥 GAME OVER!")
		fmt.Printf("Pontuação final: %d\n", j.Pontuacao)
	}
}

func (j *Jogo) MudarDirecao(direcao string) {
	direcao = strings.ToLower(direcao)
	switch direcao {
	case "w", "cima", "up":
		if j.Direcao.Y == 0 {
			j.Direcao = Posicao{X: 0, Y: -1}
		}
	case "s", "baixo", "down":
		if j.Direcao.Y == 0 {
			j.Direcao = Posicao{X: 0, Y: 1}
		}
	case "a", "esquerda", "left":
		if j.Direcao.X == 0 {
			j.Direcao = Posicao{X: -1, Y: 0}
		}
	case "d", "direita", "right":
		if j.Direcao.X == 0 {
			j.Direcao = Posicao{X: 1, Y: 0}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║              🐍 SNAKE (JOGO DA COBRINHA) 🐍             ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Use W/A/S/D ou setas para controlar a cobra.")
	fmt.Println("Colete @ para crescer e ganhar pontos!")
	fmt.Println("Pressione Enter para começar...")
	reader.ReadString('\n')

	for {
		jogo := NovoJogo()

		go func() {
			for !jogo.GameOver {
				entrada, _ := reader.ReadString('\n')
				entrada = strings.TrimSpace(entrada)
				jogo.MudarDirecao(entrada)
			}
		}()

		for !jogo.GameOver {
			jogo.Exibir()
			jogo.Atualizar()
			time.Sleep(200 * time.Millisecond)
		}

		jogo.Exibir()

		fmt.Print("\nDeseja jogar novamente? (S/N): ")
		resposta, _ := reader.ReadString('\n')
		resposta = strings.TrimSpace(strings.ToUpper(resposta))

		if resposta != "S" && resposta != "SIM" {
			fmt.Println("Obrigado por jogar! Até a próxima! 👋")
			break
		}
	}
}

