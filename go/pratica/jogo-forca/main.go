package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var palavras = []string{
	"COMPUTADOR", "PROGRAMACAO", "ALGORITMO", "FUNCAO", "VARIAVEL",
	"ESTRUTURA", "LOOP", "CONDICIONAL", "STRING", "INTEIRO",
	"BOOLEANO", "ARRAY", "MATRIZ", "OBJETO", "CLASSE",
	"METODO", "PARAMETRO", "RETORNO", "RECURSAO", "ITERACAO",
}

type Jogo struct {
	PalavraSecreta string
	PalavraAtual   []rune
	LetrasUsadas   map[rune]bool
	Tentativas     int
	MaxTentativas  int
}

func NovaPalavra() string {
	rand.Seed(time.Now().UnixNano())
	return palavras[rand.Intn(len(palavras))]
}

func NovoJogo() *Jogo {
	palavra := NovaPalavra()
	palavraAtual := make([]rune, len(palavra))
	for i := range palavra {
		palavraAtual[i] = '_'
	}

	return &Jogo{
		PalavraSecreta: palavra,
		PalavraAtual:   palavraAtual,
		LetrasUsadas:   make(map[rune]bool),
		Tentativas:     0,
		MaxTentativas:  7,
	}
}

func (j *Jogo) TentarLetra(letra rune) bool {
	letra = rune(strings.ToUpper(string(letra))[0])

	if j.LetrasUsadas[letra] {
		return false
	}

	j.LetrasUsadas[letra] = true
	acertou := false

	for i, char := range j.PalavraSecreta {
		if char == letra {
			j.PalavraAtual[i] = letra
			acertou = true
		}
	}

	if !acertou {
		j.Tentativas++
	}

	return acertou
}

func (j *Jogo) Venceu() bool {
	return string(j.PalavraAtual) == j.PalavraSecreta
}

func (j *Jogo) Perdeu() bool {
	return j.Tentativas >= j.MaxTentativas
}

func (j *Jogo) ExibirBoneco() {
	partes := []string{
		"   +---+\n   |   |\n       |\n       |\n       |\n       |\n=========",
		"   +---+\n   |   |\n   O   |\n       |\n       |\n       |\n=========",
		"   +---+\n   |   |\n   O   |\n   |   |\n       |\n       |\n=========",
		"   +---+\n   |   |\n   O   |\n  /|   |\n       |\n       |\n=========",
		"   +---+\n   |   |\n   O   |\n  /|\\  |\n       |\n       |\n=========",
		"   +---+\n   |   |\n   O   |\n  /|\\  |\n  /    |\n       |\n=========",
		"   +---+\n   |   |\n   O   |\n  /|\\  |\n  / \\  |\n       |\n=========",
	}

	if j.Tentativas < len(partes) {
		fmt.Println(partes[j.Tentativas])
	} else {
		fmt.Println(partes[len(partes)-1])
	}
}

func (j *Jogo) Exibir() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	j.ExibirBoneco()
	fmt.Println()
	fmt.Println("Palavra: " + string(j.PalavraAtual))
	fmt.Printf("Tentativas restantes: %d/%d\n", j.MaxTentativas-j.Tentativas, j.MaxTentativas)

	if len(j.LetrasUsadas) > 0 {
		letras := make([]string, 0, len(j.LetrasUsadas))
		for letra := range j.LetrasUsadas {
			letras = append(letras, string(letra))
		}
		fmt.Println("Letras usadas: " + strings.Join(letras, ", "))
	}
	fmt.Println(strings.Repeat("═", 60))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║              🎯 JOGO DA FORCA (HANGMAN) 🎯              ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Adivinhe a palavra letra por letra antes que o boneco seja enforcado!")
	fmt.Println()

	for {
		jogo := NovoJogo()

		for {
			jogo.Exibir()

			if jogo.Venceu() {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Printf("║     🎉 PARABÉNS! Você acertou: %s 🎉              ║\n", jogo.PalavraSecreta)
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				fmt.Println()
				break
			}

			if jogo.Perdeu() {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Printf("║     😢 Fim de jogo! A palavra era: %s 😢          ║\n", jogo.PalavraSecreta)
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
				fmt.Println()
				break
			}

			fmt.Print("\nDigite uma letra: ")
			entrada, _ := reader.ReadString('\n')
			entrada = strings.TrimSpace(entrada)

			if len(entrada) == 0 {
				fmt.Println("❌ Digite uma letra!")
				continue
			}

			letra := rune(strings.ToUpper(entrada)[0])
			if letra < 'A' || letra > 'Z' {
				fmt.Println("❌ Digite uma letra válida (A-Z)!")
				continue
			}

			if jogo.LetrasUsadas[letra] {
				fmt.Println("❌ Você já usou essa letra!")
				continue
			}

			if jogo.TentarLetra(letra) {
				fmt.Println("✅ Letra correta!")
			} else {
				fmt.Println("❌ Letra incorreta!")
			}

			time.Sleep(500 * time.Millisecond)
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

