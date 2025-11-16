# Projeto 09: Sistema de Votação

## 📝 Descrição
Sistema completo para criar enquetes, votar e visualizar resultados com estatísticas.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Enquetes**:
   - Criar enquete (pergunta, opções, data de encerramento)
   - Listar enquetes (ativas, encerradas)
   - Visualizar enquete

2. **Votação**:
   - Votar em uma opção
   - Verificar se usuário já votou
   - Contar votos

3. **Resultados**:
   - Exibir resultados (contagem, percentual)
   - Gráfico de barras (ASCII)
   - Estatísticas (total de votos, opção vencedora)

4. **Validação**:
   - Não permitir votar duas vezes
   - Verificar se enquete está ativa
   - Validar opção escolhida

5. **Persistência**: JSON

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Slices e maps
- ✅ Error handling
- ✅ Agregações
- ✅ Time package
- ✅ JSON
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
votacao/
├── main.go
├── enquete.go
├── voto.go
├── resultado.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Enquete struct {
    ID            string
    Pergunta      string
    Opcoes        []string
    Votos         map[string]int
    Votantes      map[string]bool
    DataCriacao   time.Time
    DataEncerramento time.Time
    Ativa         bool
}

type Resultado struct {
    EnqueteID     string
    TotalVotos    int
    Opcoes        []OpcaoResultado
    Vencedora     string
}

type OpcaoResultado struct {
    Opcao     string
    Votos     int
    Percentual float64
}
```

## ✅ Critérios de Sucesso
- [ ] CRUD de enquetes funciona
- [ ] Votação funciona corretamente
- [ ] Resultados são precisos
- [ ] Validações impedem fraudes
- [ ] Dados persistem
- [ ] Interface clara

## 🚀 Extras (Desafio)
- [ ] Votação múltipla (escolher várias opções)
- [ ] Ranking de enquetes mais votadas
- [ ] Exportar resultados
- [ ] Sistema de comentários
- [ ] Enquetes privadas (com senha)

