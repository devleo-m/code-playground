# 🎮 Jogo da Velha em Go

Um jogo da velha completo e funcional implementado em Go, com interface de terminal, sistema de pontuação e todas as funcionalidades necessárias para uma experiência completa de jogo.

## 🎯 Funcionalidades

- ✅ **Dois jogadores**: Sistema completo para dois jogadores humanos
- ✅ **Escolha de peças**: Cada jogador escolhe se quer jogar com peça branca (○) ou preta (●)
- ✅ **Regra do primeiro**: O jogador com peça branca sempre começa
- ✅ **Interface visual**: Tabuleiro formatado e bonito no terminal
- ✅ **Sistema de pontuação**: Acompanha os pontos de cada jogador ao longo das partidas
- ✅ **Validação completa**: Verifica jogadas inválidas, posições ocupadas e limites do tabuleiro
- ✅ **Detecção de vitória**: Verifica vitória em linhas, colunas e diagonais
- ✅ **Detecção de empate**: Identifica quando o jogo termina em empate
- ✅ **Múltiplas partidas**: Permite jogar várias partidas seguidas mantendo a pontuação
- ✅ **Código limpo**: Estrutura bem organizada, comentada e seguindo boas práticas

## 🚀 Como Executar

### Pré-requisitos

- Go 1.21 ou superior instalado

### Execução

1. Navegue até a pasta do jogo:
```bash
cd go/jogo-da-velha
```

2. Execute o jogo:
```bash
go run main.go
```

Ou compile e execute:
```bash
go build -o jogo-da-velha
./jogo-da-velha
```

## 🎲 Como Jogar

1. **Configuração Inicial**:
   - Digite o nome do Jogador 1
   - Escolha a peça (B para Branco ○ ou P para Preto ●)
   - Digite o nome do Jogador 2
   - A peça do Jogador 2 será automaticamente a oposta

2. **Durante o Jogo**:
   - O tabuleiro é exibido com coordenadas (linha 1-3, coluna 1-3)
   - Quando for sua vez, digite a linha e depois a coluna onde deseja jogar
   - O jogador com peça branca sempre começa

3. **Objetivo**:
   - Formar uma linha, coluna ou diagonal com suas 3 peças
   - Cada vitória adiciona 1 ponto ao seu placar

4. **Fim do Jogo**:
   - Após cada partida, você pode escolher jogar novamente
   - A pontuação é mantida entre as partidas
   - Ao final, o jogador com mais pontos é o grande vencedor!

## 📋 Estrutura do Código

- **Peça**: Tipo que representa as peças do jogo (Vazio, Branco, Preto)
- **Jogador**: Struct com nome, peça e pontuação
- **Tabuleiro**: Array 3x3 que representa o tabuleiro
- **Jogo**: Struct principal que gerencia todo o estado do jogo

### Principais Métodos

- `NovoJogo()`: Cria uma nova instância do jogo
- `ConfigurarJogadores()`: Configura nomes e peças dos jogadores
- `ExibirTabuleiro()`: Mostra o tabuleiro formatado
- `FazerJogada()`: Valida e executa uma jogada
- `VerificarVitória()`: Verifica se o jogador atual ganhou
- `VerificarEmpate()`: Verifica se o jogo terminou em empate
- `JogarPartida()`: Executa uma partida completa

## 🎨 Características Visuais

- Tabuleiro formatado com bordas e separadores
- Símbolos visuais: ○ (branco) e ● (preto)
- Mensagens coloridas e informativas
- Limpeza de tela para melhor experiência
- Exibição clara da pontuação

## 🔧 Melhorias Futuras (Opcional)

- [ ] Modo contra IA
- [ ] Diferentes tamanhos de tabuleiro
- [ ] Histórico de partidas
- [ ] Salvamento de pontuação em arquivo
- [ ] Interface gráfica (GUI)

## 📝 Licença

Este projeto é um exemplo educacional e pode ser usado livremente.

---

**Divirta-se jogando! 🎮**

