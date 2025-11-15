# Aula 17 - Simplificada: Entendendo Context

Olá! Vamos simplificar os conceitos de Context usando analogias do dia a dia para que você entenda de forma mais intuitiva.

---

## 🎯 O Que é Context? (Versão Simples)

### Analogia: O Controle Remoto Universal

Imagine que você está assistindo TV e tem um **controle remoto** que pode:
- ⏱️ **Desligar a TV após X minutos** (timeout)
- 🛑 **Desligar imediatamente quando você apertar o botão** (cancelamento)
- 📋 **Passar informações** (qual canal você está assistindo, volume, etc.)

**Context em Go é exatamente isso!**
- É um "controle remoto" que você passa para todas as operações
- Permite cancelar operações (desligar a TV)
- Permite definir timeouts (desligar após X tempo)
- Permite passar informações (valores)

**Por que é útil?**
- Se você desistir de assistir TV, pode cancelar tudo de uma vez
- Se a TV ficar ligada muito tempo, ela desliga sozinha
- Você pode passar informações (ex: "estou assistindo canal 5") para outras pessoas

---

## 🏠 Context.Background() e Context.TODO(): A Base de Tudo

### Analogia: A Fundação de uma Casa

**Context.Background()** = A **fundação** da sua casa
- É sólida, estável, nunca muda
- Tudo é construído sobre ela
- É o ponto de partida para tudo

**Context.TODO()** = Um **"a definir"** na planta da casa
- Você sabe que vai ter algo ali, mas ainda não decidiu o quê
- É temporário, só para desenvolvimento
- Em produção, você sempre usa Background

**Em Go:**
```go
// Começar com a fundação
ctx := context.Background()

// Construir coisas em cima dela
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
```

**Pense assim:**
- **Background** = "Vamos começar do zero"
- **TODO** = "Ainda não sei o que vai aqui, mas preciso de um placeholder"

---

## ⏱️ Context com Timeout: O Timer Automático

### Analogia: O Timer do Microondas

Imagine que você coloca comida no microondas e programa para **2 minutos**:

**O que acontece:**
1. Você aperta "2 minutos"
2. O microondas começa a esquentar
3. Após 2 minutos, **para automaticamente**
4. Você não precisa ficar olhando!

**Context.WithTimeout()** funciona igual:
```go
// "Programar para parar após 5 segundos"
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // Sempre "desligar" quando terminar

// Operação que pode demorar
operacaoLonga(ctx)
```

**Exemplo Real:**
- Você faz uma requisição para uma API
- Se demorar mais de 5 segundos, **cancela automaticamente**
- Não precisa esperar para sempre!

**Pense assim:**
- **Timeout** = "Faça isso, mas se demorar mais de X tempo, pare"
- É como um **timer automático** que desliga tudo sozinho

---

## 📅 Context com Deadline: O Horário Específico

### Analogia: O Despertador

**Timeout** = "Desligar após 5 minutos"
**Deadline** = "Desligar às 15:30"

**Diferença:**
- **Timeout**: "Faça por 5 minutos" (duração)
- **Deadline**: "Faça até as 15:30" (horário específico)

**Exemplo:**
```go
// Timeout: "Cancelar após 5 segundos"
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

// Deadline: "Cancelar às 15:30"
deadline := time.Date(2024, 1, 1, 15, 30, 0, 0, time.UTC)
ctx, cancel := context.WithDeadline(ctx, deadline)
```

**Quando usar cada um?**
- **Timeout**: "Cancelar após X tempo" (mais comum)
- **Deadline**: "Cancelar em um horário específico" (menos comum)

**Pense assim:**
- **Timeout** = Timer de cozinha (5 minutos)
- **Deadline** = Compromisso marcado (às 15h)

---

## 🛑 Context com Cancelamento Manual: O Botão de Emergência

### Analogia: O Botão de Parada de Emergência

Imagine uma **esteira de produção** em uma fábrica:
- Múltiplas máquinas trabalhando
- Todas têm um **botão vermelho de emergência**
- Quando você aperta, **todas param imediatamente**

**Context.WithCancel()** é esse botão:
```go
// Criar o "botão de emergência"
ctx, cancel := context.WithCancel(context.Background())

// Múltiplas "máquinas" (goroutines) trabalhando
go maquina1(ctx)
go maquina2(ctx)
go maquina3(ctx)

// Apertar o botão = cancelar tudo
cancel() // Todas as máquinas param!
```

**Exemplo Real:**
- Você inicia 10 downloads simultâneos
- Usuário clica em "Cancelar"
- Você chama `cancel()` e **todos os downloads param**

**Pense assim:**
- **Cancel** = "Pare tudo agora!"
- É como um **interruptor geral** que desliga tudo de uma vez

---

## 📋 Context com Valores: A Pasta de Requisição

### Analogia: A Pasta de um Processo

Imagine que você está processando um **pedido de cliente**:
- Você tem uma **pasta** com todas as informações
- A pasta passa por vários departamentos:
  - Recepção (adiciona número do pedido)
  - Financeiro (adiciona forma de pagamento)
  - Estoque (adiciona produtos)
  - Entrega (usa todas as informações)

**Context.WithValue()** é essa pasta:
```go
// Recepção adiciona número do pedido
ctx := context.WithValue(ctx, "pedidoID", "12345")

// Financeiro adiciona forma de pagamento
ctx = context.WithValue(ctx, "pagamento", "cartão")

// Entrega usa as informações
pedidoID := ctx.Value("pedidoID")
```

**Regras Importantes:**
- ✅ Use apenas para **informações da requisição** (ID do usuário, ID da requisição)
- ❌ **NÃO** use para parâmetros de função
- ❌ **NÃO** use para dependências (banco de dados, etc.)

**Pense assim:**
- **Valores no Context** = "Informações que precisam passar por vários lugares"
- É como uma **pasta** que vai de departamento em departamento

---

## 🌐 Context em Requisições HTTP: O "Controle" Automático

### Analogia: O Garçom com Timer

Imagine um **restaurante** onde:
- Cada cliente tem um **timer** na mesa
- Se o cliente demorar mais de 30 minutos, o garçom **cancela o pedido**
- Se o cliente sair antes, o garçom **cancela o pedido**

**Context em HTTP funciona assim:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Cada requisição HTTP já vem com um context!
    ctx := r.Context()

    // Se o cliente fechar a conexão, ctx é cancelado automaticamente
    // Se você adicionar timeout, cancela após X tempo
    
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    // Processar requisição
    processarRequisicao(ctx)
}
```

**O que acontece:**
1. Cliente faz requisição → Context criado automaticamente
2. Você adiciona timeout de 30 segundos
3. Se demorar mais de 30s → Cancela automaticamente
4. Se cliente fechar navegador → Cancela automaticamente

**Pense assim:**
- **Context em HTTP** = "Cada cliente tem seu próprio timer"
- Se o cliente desistir ou demorar muito, **cancela tudo**

---

## 💾 Context em Banco de Dados: A Query com Prazo

### Analogia: O Exame com Tempo Limite

Imagine que você está fazendo um **exame**:
- Você tem **1 hora** para completar
- Se passar de 1 hora, o professor **recolhe a prova**
- Você não pode continuar depois

**Context em queries funciona assim:**
```go
// "Você tem 5 segundos para fazer essa query"
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

// Fazer query
rows, err := db.QueryContext(ctx, "SELECT * FROM users")
if err != nil {
    // Se demorar mais de 5s, retorna erro
    return err
}
```

**Por que é importante?**
- Sem timeout: Query pode ficar rodando **para sempre**
- Com timeout: Query **para automaticamente** após X tempo
- Evita travar o banco de dados

**Pense assim:**
- **Context em DB** = "Essa query tem prazo"
- Se demorar muito, **cancela e retorna erro**

---

## 🔗 Context Aninhado: A Cadeia de Comandos

### Analogia: O Sistema de Hierarquia Militar

Imagine um **exército**:
- **General** (context raiz) dá ordem
- **Coronel** (context derivado) recebe ordem e adiciona mais detalhes
- **Capitão** (context derivado) recebe e adiciona mais detalhes
- **Soldado** (context derivado) executa

**Se o General cancelar, TODOS param!**

**Em Go:**
```go
// General (raiz)
ctx := context.Background()

// Coronel (adiciona timeout de 10s)
ctx, cancel1 := context.WithTimeout(ctx, 10*time.Second)

// Capitão (adiciona timeout mais restritivo de 2s)
ctx, cancel2 := context.WithTimeout(ctx, 2*time.Second)

// Soldado executa (tempo limite de 2s, não 10s!)
executarTarefa(ctx)
```

**O que acontece:**
- Se `cancel1()` for chamado → Tudo para (General cancelou)
- Se `cancel2()` for chamado → Só essa parte para (Capitão cancelou)
- Se passar de 2 segundos → Para automaticamente (timeout do Capitão)

**Pense assim:**
- **Context aninhado** = "Hierarquia de controle"
- Contextos filhos **herdam** cancelamento dos pais
- Contextos filhos podem ter **regras mais restritivas**

---

## ✅ Verificando Status: O "Está Ligado?"

### Analogia: Verificar se a TV Está Ligada

**Antes de fazer algo, você verifica:**
- A TV está ligada? (`ctx.Err() == nil`)
- Foi desligada? (`ctx.Err() != nil`)
- Por que foi desligada? (`ctx.Err()` retorna o motivo)

**Em Go:**
```go
// Verificar se está cancelado
if ctx.Err() != nil {
    // Foi cancelado!
    return ctx.Err()
}

// Ou usando select
select {
case <-ctx.Done():
    // Foi cancelado!
    return ctx.Err()
default:
    // Ainda está ativo, pode continuar
}
```

**Pense assim:**
- **Verificar status** = "A operação ainda pode continuar?"
- É como **verificar se a TV está ligada** antes de mudar de canal

---

## 🎯 Padrões de Uso: As "Regras de Ouro"

### Regra 1: Context Sempre Primeiro

**Como passar um documento importante:**
- Você sempre passa **primeiro** (não no meio, não no final)
- Context é igual: sempre primeiro parâmetro

```go
// ✅ CORRETO
func processar(ctx context.Context, dados []string) error

// ❌ ERRADO
func processar(dados []string, ctx context.Context) error
```

### Regra 2: Sempre Cancelar

**Como desligar um aparelho:**
- Você sempre desliga quando termina
- Context é igual: sempre use `defer cancel()`

```go
// ✅ CORRETO
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel() // Sempre cancelar!

// ❌ ERRADO
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
// Esqueceu de cancelar!
```

### Regra 3: Verificar em Loops

**Como verificar se ainda pode continuar:**
- Em loops longos, sempre verifique se foi cancelado
- É como verificar se ainda tem tempo antes de cada tarefa

```go
// ✅ CORRETO
for i := 0; i < 1000000; i++ {
    select {
    case <-ctx.Done():
        return ctx.Err() // Parar se foi cancelado
    default:
    }
    processar(i)
}
```

---

## 🚫 Erros Comuns: O Que NÃO Fazer

### Erro 1: Não Cancelar = Vazamento de Recursos

**Analogia:** Deixar a TV ligada 24 horas por dia
- Gasta energia desnecessariamente
- Context não cancelado = recursos não liberados

**Solução:** Sempre use `defer cancel()`

### Erro 2: Usar Context para Parâmetros

**Analogia:** Usar o controle remoto para guardar o número do canal
- Controle remoto é para **controlar**, não para **guardar coisas**
- Context é para **cancelamento**, não para **parâmetros**

**Solução:** Use parâmetros normais de função

### Erro 3: Não Verificar Cancelamento

**Analogia:** Continuar assistindo TV mesmo depois que desligou
- Não faz sentido!
- Se context foi cancelado, pare de trabalhar

**Solução:** Sempre verifique `ctx.Done()` em loops

---

## 📝 Resumo com Analogias

| Conceito | Analogia | Em Go |
|----------|----------|-------|
| **Context** | Controle remoto universal | `context.Context` |
| **Background** | Fundação da casa | `context.Background()` |
| **Timeout** | Timer do microondas | `context.WithTimeout()` |
| **Deadline** | Despertador | `context.WithDeadline()` |
| **Cancel** | Botão de emergência | `context.WithCancel()` |
| **Valores** | Pasta de processo | `context.WithValue()` |
| **Verificar** | "Está ligado?" | `ctx.Done()`, `ctx.Err()` |

---

## 🎓 Pensando de Forma Simples

**Context é como um "controle remoto" que você passa para todas as operações:**

1. **Timeout** = "Desligar após X tempo"
2. **Cancel** = "Desligar agora"
3. **Valores** = "Passar informações"
4. **Verificar** = "Ainda está ligado?"

**Use quando:**
- ✅ Operações que podem demorar (HTTP, DB, I/O)
- ✅ Precisa cancelar operações
- ✅ Precisa passar informações de requisição
- ✅ Precisa definir limites de tempo

**Não use para:**
- ❌ Parâmetros de função
- ❌ Dependências (banco, serviços)
- ❌ Configurações globais

---

E assim simplificamos o Context! Agora você entende que:
- Context é como um **controle remoto** que controla operações
- Timeout é como um **timer automático**
- Cancel é como um **botão de emergência**
- Valores são como uma **pasta** que passa informações

Na próxima aula, vamos praticar com exercícios para fixar ainda mais esses conceitos!

Sinta-se à vontade para reler este material. Se tiver qualquer dúvida, pode perguntar!

