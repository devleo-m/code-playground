# Projeto 19: Sistema de Notificações

## 📝 Descrição
Sistema completo de notificações com múltiplos canais (email simulado, console, arquivo) e filas.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Notificações**:
   - Criar notificação (tipo, destinatário, mensagem, prioridade)
   - Enfileirar notificações
   - Processar fila

2. **Canais**:
   - Console (imprimir)
   - Arquivo (salvar em log)
   - Email (simulado - salvar em arquivo)

3. **Prioridades**: Baixa, Normal, Alta, Urgente

4. **Fila**:
   - Processar por prioridade
   - Retry em caso de falha
   - Dead letter queue

5. **Templates**: Suporte a templates de mensagem

6. **Concorrência**: Worker pool para processar

## 📚 Conceitos Utilizados
- ✅ Goroutines
- ✅ Channels
- ✅ Worker pool
- ✅ Priority queue
- ✅ Error handling
- ✅ Templates (strings)
- ✅ Concorrência avançada

## 📁 Estrutura Sugerida
```
notificacoes/
├── main.go
├── notificacao.go
├── canal.go
├── fila.go
├── worker.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Notificacao struct {
    ID          string
    Tipo        string
    Destinatario string
    Mensagem    string
    Prioridade  int
    Canal       string
    Tentativas  int
}

type Canal interface {
    Enviar(notificacao Notificacao) error
}

type FilaNotificacoes struct {
    notificacoes chan Notificacao
    workers      int
}
```

### Concorrência
- Priority queue para ordenar
- Worker pool para processar
- Retry mechanism

## ✅ Critérios de Sucesso
- [ ] Notificações são criadas
- [ ] Fila funciona
- [ ] Canais funcionam
- [ ] Prioridades são respeitadas
- [ ] Retry funciona
- [ ] Concorrência funciona

## 🚀 Extras (Desafio)
- [ ] Múltiplos canais por notificação
- [ ] Agendamento
- [ ] Webhooks
- [ ] Métricas e analytics
- [ ] Dashboard


