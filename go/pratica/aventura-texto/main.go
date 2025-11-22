package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Nome        string
	Descricao   string
	Coletavel   bool
}

type Sala struct {
	Nome        string
	Descricao   string
	Itens       []*Item
	Saidas      map[string]*Sala
}

type Jogo struct {
	SalaAtual   *Sala
	Inventario  []*Item
}

func CriarMundo() *Sala {
	// Criar salas
	entrada := &Sala{
		Nome:      "Entrada da Caverna",
		Descricao: "Você está na entrada de uma caverna escura. Há uma tocha na parede.",
		Itens:     []*Item{{Nome: "tocha", Descricao: "Uma tocha que ilumina o caminho", Coletavel: true}},
		Saidas:    make(map[string]*Sala),
	}

	corredor := &Sala{
		Nome:      "Corredor",
		Descricao: "Um corredor longo e estreito. Você ouve ecos distantes.",
		Itens:     []*Item{},
		Saidas:    make(map[string]*Sala),
	}

	salaTesouro := &Sala{
		Nome:      "Sala do Tesouro",
		Descricao: "Uma sala brilhante com um baú dourado no centro!",
		Itens:     []*Item{{Nome: "tesouro", Descricao: "Um baú cheio de ouro e joias", Coletavel: true}},
		Saidas:    make(map[string]*Sala),
	}

	salaArmas := &Sala{
		Nome:      "Sala de Armas",
		Descricao: "Uma sala com armas antigas nas paredes.",
		Itens:     []*Item{{Nome: "espada", Descricao: "Uma espada afiada e brilhante", Coletavel: true}},
		Saidas:    make(map[string]*Sala),
	}

	// Conectar salas
	entrada.Saidas["norte"] = corredor
	entrada.Saidas["sul"] = nil

	corredor.Saidas["norte"] = salaTesouro
	corredor.Saidas["sul"] = entrada
	corredor.Saidas["leste"] = salaArmas

	salaTesouro.Saidas["sul"] = corredor

	salaArmas.Saidas["oeste"] = corredor

	return entrada
}

func (j *Jogo) ExibirSala() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Printf("📍 %s\n", j.SalaAtual.Nome)
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println(j.SalaAtual.Descricao)
	fmt.Println()

	if len(j.SalaAtual.Itens) > 0 {
		fmt.Println("Itens na sala:")
		for _, item := range j.SalaAtual.Itens {
			fmt.Printf("  - %s: %s\n", item.Nome, item.Descricao)
		}
		fmt.Println()
	}

	if len(j.SalaAtual.Saidas) > 0 {
		fmt.Println("Saídas disponíveis:")
		for direcao, sala := range j.SalaAtual.Saidas {
			if sala != nil {
				fmt.Printf("  - %s: %s\n", direcao, sala.Nome)
			}
		}
		fmt.Println()
	}
}

func (j *Jogo) ProcessarComando(comando string) {
	comando = strings.ToLower(strings.TrimSpace(comando))
	palavras := strings.Fields(comando)

	if len(palavras) == 0 {
		return
	}

	acao := palavras[0]

	switch acao {
	case "ir", "mover", "caminhar":
		if len(palavras) < 2 {
			fmt.Println("❌ Para onde você quer ir? (ex: ir norte)")
			return
		}
		direcao := palavras[1]
		j.Ir(direcao)

	case "pegar", "coletar":
		if len(palavras) < 2 {
			fmt.Println("❌ O que você quer pegar? (ex: pegar espada)")
			return
		}
		itemNome := strings.Join(palavras[1:], " ")
		j.PegarItem(itemNome)

	case "inventario", "itens", "mochila":
		j.ExibirInventario()

	case "olhar", "examinar", "ver":
		j.ExibirSala()

	case "ajuda", "help":
		j.ExibirAjuda()

	case "sair", "quit":
		fmt.Println("Até a próxima, aventureiro! 👋")
		os.Exit(0)

	default:
		fmt.Println("❌ Comando não reconhecido. Digite 'ajuda' para ver os comandos disponíveis.")
	}
}

func (j *Jogo) Ir(direcao string) {
	sala, existe := j.SalaAtual.Saidas[direcao]
	if !existe || sala == nil {
		fmt.Println("❌ Não há saída nessa direção!")
		return
	}

	j.SalaAtual = sala
	fmt.Printf("✅ Você foi para %s.\n", sala.Nome)
	j.ExibirSala()
}

func (j *Jogo) PegarItem(nomeItem string) {
	for i, item := range j.SalaAtual.Itens {
		if strings.ToLower(item.Nome) == strings.ToLower(nomeItem) {
			if !item.Coletavel {
				fmt.Printf("❌ Você não pode pegar %s.\n", item.Nome)
				return
			}

			j.Inventario = append(j.Inventario, item)
			j.SalaAtual.Itens = append(j.SalaAtual.Itens[:i], j.SalaAtual.Itens[i+1:]...)
			fmt.Printf("✅ Você pegou %s!\n", item.Nome)

			if item.Nome == "tesouro" {
				fmt.Println()
				fmt.Println("╔══════════════════════════════════════════════════════════╗")
				fmt.Println("║                                                          ║")
				fmt.Println("║         🎉 PARABÉNS! Você encontrou o tesouro! 🎉       ║")
				fmt.Println("║                                                          ║")
				fmt.Println("╚══════════════════════════════════════════════════════════╝")
			}
			return
		}
	}
	fmt.Printf("❌ Não há '%s' nesta sala.\n", nomeItem)
}

func (j *Jogo) ExibirInventario() {
	if len(j.Inventario) == 0 {
		fmt.Println("Seu inventário está vazio.")
		return
	}

	fmt.Println("\n📦 Inventário:")
	for _, item := range j.Inventario {
		fmt.Printf("  - %s: %s\n", item.Nome, item.Descricao)
	}
}

func (j *Jogo) ExibirAjuda() {
	fmt.Println("\n📖 Comandos disponíveis:")
	fmt.Println("  ir [direção]     - Mover para uma direção (norte, sul, leste, oeste)")
	fmt.Println("  pegar [item]     - Coletar um item da sala")
	fmt.Println("  inventario       - Ver seus itens")
	fmt.Println("  olhar            - Examinar a sala atual")
	fmt.Println("  ajuda            - Mostrar esta ajuda")
	fmt.Println("  sair             - Sair do jogo")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                          ║")
	fmt.Println("║            🗺️  AVENTURA DE TEXTO 🗺️                      ║")
	fmt.Println("║                                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Explore a caverna, colete itens e encontre o tesouro!")
	fmt.Println("Digite 'ajuda' para ver os comandos disponíveis.")
	fmt.Println()

	jogo := &Jogo{
		SalaAtual:  CriarMundo(),
		Inventario: []*Item{},
	}

	jogo.ExibirSala()

	for {
		fmt.Print("\n> ")
		comando, _ := reader.ReadString('\n')
		jogo.ProcessarComando(comando)
	}
}

