package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Peça representa as peças do jogo
type Peça string

const (
	Vazio  Peça = " "
	Branco Peça = "○"
	Preto  Peça = "●"
)

// Jogador representa um jogador do jogo
type Jogador struct {
	Nome  string
	Peça  Peça
	Pontos int
}

// Tabuleiro representa o tabuleiro 3x3 do jogo da velha
type Tabuleiro [3][3]Peça

// Jogo contém todo o estado do jogo
type Jogo struct {
	Tabuleiro     Tabuleiro
	Jogador1      Jogador
	Jogador2      Jogador
	JogadorAtual  *Jogador
	Rodada        int
	Jogadas       int
	EmAndamento   bool
}

// NovoJogo cria uma nova instância do jogo
func NovoJogo() *Jogo {
	jogo := &Jogo{
		Tabuleiro:    Tabuleiro{},
		Rodada:       1,
		Jogadas:      0,
		EmAndamento:  true,
	}
	
	// Inicializa o tabuleiro com espaços vazios
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			jogo.Tabuleiro[i][j] = Vazio
		}
	}
	
	return jogo
}

// ConfigurarJogadores configura os jogadores e suas peças
func (j *Jogo) ConfigurarJogadores() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║         🎮 CONFIGURAÇÃO DOS JOGADORES 🎮                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	
	// Nome do Jogador 1
	fmt.Print("👤 Digite o nome do Jogador 1: ")
	nome1, _ := reader.ReadString('\n')
	nome1 = strings.TrimSpace(nome1)
	if nome1 == "" {
		nome1 = "Jogador 1"
	}
	
	// Escolha de peça para Jogador 1
	var peça1 Peça
	for {
		fmt.Print("\n🎯 Escolha sua peça para " + nome1 + " (B para Branco ○, P para Preto ●): ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(strings.ToUpper(escolha))
		
		if escolha == "B" || escolha == "BRANCO" {
			peça1 = Branco
			break
		} else if escolha == "P" || escolha == "PRETO" {
			peça1 = Preto
			break
		} else {
			fmt.Println("❌ Opção inválida! Digite 'B' para Branco ou 'P' para Preto.")
		}
	}
	
	// Nome do Jogador 2
	fmt.Print("\n👤 Digite o nome do Jogador 2: ")
	nome2, _ := reader.ReadString('\n')
	nome2 = strings.TrimSpace(nome2)
	if nome2 == "" {
		nome2 = "Jogador 2"
	}
	
	// Define a peça do Jogador 2 (a oposta)
	var peça2 Peça
	if peça1 == Branco {
		peça2 = Preto
	} else {
		peça2 = Branco
	}
	
	j.Jogador1 = Jogador{Nome: nome1, Peça: peça1, Pontos: 0}
	j.Jogador2 = Jogador{Nome: nome2, Peça: peça2, Pontos: 0}
	
	// O jogador com a peça branca sempre começa
	if peça1 == Branco {
		j.JogadorAtual = &j.Jogador1
	} else {
		j.JogadorAtual = &j.Jogador2
	}
	
	fmt.Println("\n✅ Configuração concluída!")
	fmt.Printf("   %s jogará com %s\n", j.Jogador1.Nome, j.Jogador1.Peça)
	fmt.Printf("   %s jogará com %s\n", j.Jogador2.Nome, j.Jogador2.Peça)
	fmt.Printf("   %s começa (peça branca sempre começa)!\n", j.JogadorAtual.Nome)
	fmt.Println("\nPressione Enter para começar...")
	reader.ReadString('\n')
}

// ExibirTabuleiro exibe o tabuleiro formatado no terminal
func (j *Jogo) ExibirTabuleiro() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("                    TABULEIRO")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()
	
	// Cabeçalho com números das colunas
	fmt.Println("       1     2     3")
	fmt.Println("    ┌─────┬─────┬─────┐")
	
	for i := 0; i < 3; i++ {
		fmt.Printf("  %d │", i+1)
		for k := 0; k < 3; k++ {
			peça := j.Tabuleiro[i][k]
			if peça == Vazio {
				fmt.Printf("  %s  │", " ")
			} else {
				fmt.Printf("  %s  │", peça)
			}
		}
		fmt.Println()
		
		if i < 2 {
			fmt.Println("    ├─────┼─────┼─────┤")
		}
	}
	
	fmt.Println("    └─────┴─────┴─────┘")
	fmt.Println()
}

// ExibirPontuação exibe a pontuação atual dos jogadores
func (j *Jogo) ExibirPontuação() {
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("                    PONTUAÇÃO")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("   %s (%s): %d pontos\n", j.Jogador1.Nome, j.Jogador1.Peça, j.Jogador1.Pontos)
	fmt.Printf("   %s (%s): %d pontos\n", j.Jogador2.Nome, j.Jogador2.Peça, j.Jogador2.Pontos)
	fmt.Println(strings.Repeat("═", 60))
}

// FazerJogada tenta fazer uma jogada na posição especificada
func (j *Jogo) FazerJogada(linha, coluna int) bool {
	// Validação de índices
	if linha < 1 || linha > 3 || coluna < 1 || coluna > 3 {
		return false
	}
	
	// Ajusta para índices do array (0-2)
	linha--
	coluna--
	
	// Verifica se a posição está vazia
	if j.Tabuleiro[linha][coluna] != Vazio {
		return false
	}
	
	// Faz a jogada
	j.Tabuleiro[linha][coluna] = j.JogadorAtual.Peça
	j.Jogadas++
	
	return true
}

// VerificarVitória verifica se o jogador atual ganhou
func (j *Jogo) VerificarVitória() bool {
	peça := j.JogadorAtual.Peça
	
	// Verifica linhas
	for i := 0; i < 3; i++ {
		if j.Tabuleiro[i][0] == peça &&
			j.Tabuleiro[i][1] == peça &&
			j.Tabuleiro[i][2] == peça {
			return true
		}
	}
	
	// Verifica colunas
	for i := 0; i < 3; i++ {
		if j.Tabuleiro[0][i] == peça &&
			j.Tabuleiro[1][i] == peça &&
			j.Tabuleiro[2][i] == peça {
			return true
		}
	}
	
	// Verifica diagonal principal
	if j.Tabuleiro[0][0] == peça &&
		j.Tabuleiro[1][1] == peça &&
		j.Tabuleiro[2][2] == peça {
		return true
	}
	
	// Verifica diagonal secundária
	if j.Tabuleiro[0][2] == peça &&
		j.Tabuleiro[1][1] == peça &&
		j.Tabuleiro[2][0] == peça {
		return true
	}
	
	return false
}

// VerificarEmpate verifica se o jogo terminou em empate
func (j *Jogo) VerificarEmpate() bool {
	return j.Jogadas == 9 && !j.VerificarVitória()
}

// AlternarJogador alterna para o próximo jogador
func (j *Jogo) AlternarJogador() {
	if j.JogadorAtual == &j.Jogador1 {
		j.JogadorAtual = &j.Jogador2
	} else {
		j.JogadorAtual = &j.Jogador1
	}
}

// ReiniciarTabuleiro reinicia o tabuleiro para uma nova partida
func (j *Jogo) ReiniciarTabuleiro() {
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			j.Tabuleiro[i][k] = Vazio
		}
	}
	j.Jogadas = 0
	j.EmAndamento = true
	
	// O jogador com peça branca sempre começa
	if j.Jogador1.Peça == Branco {
		j.JogadorAtual = &j.Jogador1
	} else {
		j.JogadorAtual = &j.Jogador2
	}
}

// LimparTela limpa a tela do terminal
func LimparTela() {
	fmt.Print("\033[2J\033[H")
}

// LerJogada lê a jogada do jogador atual
func (j *Jogo) LerJogada() (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Printf("\n🎯 Vez de %s (%s)\n", j.JogadorAtual.Nome, j.JogadorAtual.Peça)
		fmt.Print("Digite a linha (1-3): ")
		linhaStr, _ := reader.ReadString('\n')
		linhaStr = strings.TrimSpace(linhaStr)
		
		linha, err := strconv.Atoi(linhaStr)
		if err != nil {
			fmt.Println("❌ Entrada inválida! Digite um número entre 1 e 3.")
			continue
		}
		
		fmt.Print("Digite a coluna (1-3): ")
		colunaStr, _ := reader.ReadString('\n')
		colunaStr = strings.TrimSpace(colunaStr)
		
		coluna, err := strconv.Atoi(colunaStr)
		if err != nil {
			fmt.Println("❌ Entrada inválida! Digite um número entre 1 e 3.")
			continue
		}
		
		return linha, coluna, nil
	}
}

// JogarPartida executa uma partida completa
func (j *Jogo) JogarPartida() {
	j.ReiniciarTabuleiro()
	
	for j.EmAndamento {
		LimparTela()
		j.ExibirPontuação()
		j.ExibirTabuleiro()
		
		// Lê a jogada
		linha, coluna, err := j.LerJogada()
		if err != nil {
			fmt.Println("❌ Erro ao ler jogada:", err)
			continue
		}
		
		// Tenta fazer a jogada
		if !j.FazerJogada(linha, coluna) {
			fmt.Println("❌ Jogada inválida! A posição já está ocupada ou está fora do tabuleiro.")
			fmt.Println("Pressione Enter para tentar novamente...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
		
		// Verifica vitória
		if j.VerificarVitória() {
			LimparTela()
			j.ExibirPontuação()
			j.ExibirTabuleiro()
			fmt.Printf("\n🎉 PARABÉNS! %s (%s) VENCEU! 🎉\n\n", j.JogadorAtual.Nome, j.JogadorAtual.Peça)
			j.JogadorAtual.Pontos++
			j.EmAndamento = false
			return
		}
		
		// Verifica empate
		if j.VerificarEmpate() {
			LimparTela()
			j.ExibirPontuação()
			j.ExibirTabuleiro()
			fmt.Println("\n🤝 EMPATE! Ninguém venceu esta rodada.\n")
			j.EmAndamento = false
			return
		}
		
		// Alterna jogador
		j.AlternarJogador()
	}
}

// PerguntarContinuar pergunta se os jogadores querem continuar
func PerguntarContinuar() bool {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\n🔄 Deseja jogar outra partida? (S/N): ")
		resposta, _ := reader.ReadString('\n')
		resposta = strings.TrimSpace(strings.ToUpper(resposta))
		
		if resposta == "S" || resposta == "SIM" || resposta == "Y" || resposta == "YES" {
			return true
		} else if resposta == "N" || resposta == "NÃO" || resposta == "NAO" || resposta == "NO" {
			return false
		} else {
			fmt.Println("❌ Resposta inválida! Digite 'S' para sim ou 'N' para não.")
		}
	}
}

// ExibirMenuPrincipal exibe o menu principal do jogo
func ExibirMenuPrincipal() {
	LimparTela()
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║              🎮 JOGO DA VELHA EM GO 🎮                   ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

// main é a função principal do programa
func main() {
	ExibirMenuPrincipal()
	
	// Cria novo jogo
	jogo := NovoJogo()
	
	// Configura jogadores
	jogo.ConfigurarJogadores()
	
	// Loop principal do jogo
	for {
		jogo.JogarPartida()
		
		// Exibe pontuação final da partida
		jogo.ExibirPontuação()
		
		// Pergunta se quer continuar
		if !PerguntarContinuar() {
			LimparTela()
			fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
			fmt.Println("║                                                          ║")
			fmt.Println("║              🏆 PONTUAÇÃO FINAL 🏆                       ║")
			fmt.Println("║                                                          ║")
			fmt.Println("╚══════════════════════════════════════════════════════════╝")
			fmt.Println()
			jogo.ExibirPontuação()
			
			// Determina o vencedor geral
			if jogo.Jogador1.Pontos > jogo.Jogador2.Pontos {
				fmt.Printf("\n🎉 %s é o GRANDE VENCEDOR com %d pontos! 🎉\n\n", 
					jogo.Jogador1.Nome, jogo.Jogador1.Pontos)
			} else if jogo.Jogador2.Pontos > jogo.Jogador1.Pontos {
				fmt.Printf("\n🎉 %s é o GRANDE VENCEDOR com %d pontos! 🎉\n\n", 
					jogo.Jogador2.Nome, jogo.Jogador2.Pontos)
			} else {
				fmt.Println("\n🤝 EMPATE GERAL! Ambos os jogadores têm a mesma pontuação! 🤝\n")
			}
			
			fmt.Println("Obrigado por jogar! Até a próxima! 👋\n")
			break
		}
	}
}

