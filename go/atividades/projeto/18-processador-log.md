# Projeto 18: Processador de Logs

## 📝 Descrição
Sistema para processar, analisar e gerar relatórios de arquivos de log.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Parsing**:
   - Ler arquivo de log linha por linha
   - Parsear formato comum (Apache, Nginx, custom)
   - Extrair campos (timestamp, level, mensagem, etc.)

2. **Análise**:
   - Contar por nível (INFO, ERROR, WARN)
   - Agrupar por hora/dia
   - Top erros
   - Padrões frequentes

3. **Filtros**:
   - Por nível
   - Por período
   - Por palavra-chave
   - Por regex

4. **Relatórios**:
   - Estatísticas gerais
   - Timeline de eventos
   - Gráficos (ASCII)

5. **Processamento**:
   - Suportar arquivos grandes (streaming)
   - Processamento paralelo

## 📚 Conceitos Utilizados
- ✅ I/O (bufio para streaming)
- ✅ Regex
- ✅ Strings
- ✅ Goroutines
- ✅ Channels
- ✅ Agregações
- ✅ Time package

## 📁 Estrutura Sugerida
```
logprocessor/
├── main.go
├── parser.go
├── analisador.go
├── filtros.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type LogEntry struct {
    Timestamp time.Time
    Level     string
    Mensagem  string
    Campos    map[string]string
}

type Relatorio struct {
    Total      int
    PorNivel   map[string]int
    PorHora    map[int]int
    TopErros   []string
}
```

### Processamento
- Streaming para arquivos grandes
- Worker pool para processamento paralelo
- Agregações eficientes

## ✅ Critérios de Sucesso
- [ ] Parsing funciona
- [ ] Análises são precisas
- [ ] Filtros funcionam
- [ ] Relatórios são úteis
- [ ] Performance é boa
- [ ] Código eficiente

## 🚀 Extras (Desafio)
- [ ] Múltiplos formatos
- [ ] Real-time processing
- [ ] Alertas automáticos
- [ ] Exportar relatórios
- [ ] Dashboard interativo



