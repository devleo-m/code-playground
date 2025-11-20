# Projeto 07: Jogo da Forca

## 📝 Descrição
Implemente o clássico jogo da Forca com palavras aleatórias, níveis de dificuldade e pontuação.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Jogo**:
   - Selecionar palavra aleatória
   - Mostrar letras descobertas
   - Permitir chutar letra
   - Contar tentativas restantes
   - Detectar vitória/derrota

2. **Dificuldades**:
   - Fácil: 8 tentativas
   - Médio: 6 tentativas
   - Difícil: 4 tentativas

3. **Categorias**: Palavras por tema (animais, países, etc.)

4. **Pontuação**: Sistema de pontos baseado em dificuldade e tentativas

5. **Histórico**: Salvar partidas e estatísticas

6. **Interface**: ASCII art da forca

## 📚 Conceitos Utilizados
- ✅ Structs
- ✅ Slices e strings
- ✅ Loops e condicionais
- ✅ Random
- ✅ Error handling
- ✅ JSON
- ✅ I/O

## 📁 Estrutura Sugerida
```
forca/
├── main.go
├── jogo.go
├── palavras.go
├── pontuacao.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Jogo struct {
    Palavra        string
    LetrasDescobertas []bool
    LetrasTentadas []rune
    Tentativas     int
    MaxTentativas  int
    Pontuacao      int
}

type Estatisticas struct {
    PartidasJogadas int
    Vitorias        int
    Derrotas        int
    PontuacaoTotal  int
}
```

### Funcionalidades
- `NovaPartida(dificuldade string) Jogo`
- `TentarLetra(jogo *Jogo, letra rune) bool`
- `VerificarVitoria(jogo Jogo) bool`
- `DesenharForca(tentativas int)`
- `CalcularPontuacao(jogo Jogo) int`

## ✅ Critérios de Sucesso
- [ ] Jogo funciona corretamente
- [ ] Dificuldades funcionam
- [ ] Interface é clara
- [ ] Pontuação é calculada
- [ ] Estatísticas são salvas
- [ ] Código organizado

## 🚀 Extras (Desafio)
- [ ] Modo multiplayer
- [ ] Dicas
- [ ] Ranking de jogadores
- [ ] Modo contra o tempo
- [ ] Palavras customizadas



