# Projeto 11: Chat em Terminal

## 📝 Descrição
Sistema de chat simples que permite múltiplos usuários conversarem via terminal usando arquivos compartilhados.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Usuários**:
   - Criar usuário (nome, ID único)
   - Listar usuários online

2. **Mensagens**:
   - Enviar mensagem (para usuário específico ou broadcast)
   - Receber mensagens (polling)
   - Timestamp em cada mensagem

3. **Canais**:
   - Criar canal
   - Entrar em canal
   - Enviar mensagem para canal

4. **Persistência**: Mensagens em arquivo JSON

5. **Interface**: Terminal interativo

## 📚 Conceitos Utilizados
- ✅ Structs
- ✅ Goroutines (para polling)
- ✅ Channels (comunicação)
- ✅ Slices e maps
- ✅ JSON
- ✅ Time package
- ✅ I/O
- ✅ Concorrência básica

## 📁 Estrutura Sugerida
```
chat/
├── main.go
├── usuario.go
├── mensagem.go
├── canal.go
├── servidor.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Mensagem struct {
    ID        string
    De        string
    Para      string // "" para broadcast
    Canal     string
    Conteudo  string
    Timestamp time.Time
}

type Usuario struct {
    ID   string
    Nome string
    Online bool
}
```

### Concorrência
- Goroutine para receber mensagens (polling)
- Channel para comunicação entre goroutines
- Mutex para acesso thread-safe aos dados

## ✅ Critérios de Sucesso
- [ ] Múltiplos usuários podem conversar
- [ ] Mensagens são entregues
- [ ] Canais funcionam
- [ ] Dados persistem
- [ ] Concorrência funciona corretamente
- [ ] Interface é responsiva

## 🚀 Extras (Desafio)
- [ ] Mensagens privadas
- [ ] Histórico de conversas
- [ ] Emojis e formatação
- [ ] Comandos (/help, /users, /quit)
- [ ] Notificações

