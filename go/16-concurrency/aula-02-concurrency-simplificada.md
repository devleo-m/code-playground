# Aula 16 - Simplificada: Entendendo Concorrência

Olá! Vamos simplificar os conceitos de concorrência usando analogias do dia a dia para que você entenda de forma mais intuitiva.

---

## 🎯 O Que é Concorrência? (Versão Simples)

### Analogia: O Garçom do Restaurante

Imagine um **restaurante** com um único garçom:

**Sem Concorrência (Sequencial):**
- O garçom atende a Mesa 1 completamente (pega pedido, leva comida, traz conta)
- Só depois vai para a Mesa 2
- Depois para a Mesa 3
- Resultado: Mesas esperam muito tempo! 😫

**Com Concorrência:**
- O garçom atende a Mesa 1 (pega pedido)
- Enquanto a cozinha prepara, vai para a Mesa 2 (pega pedido)
- Enquanto ambas cozinham, vai para a Mesa 3
- Volta para Mesa 1 quando a comida está pronta
- Resultado: Todas as mesas são atendidas mais rápido! 🎉

**Em Go:**
- **Goroutines** = o garçom atendendo múltiplas mesas
- **Channels** = a comunicação entre garçom e cozinha
- **Concorrência** = fazer várias coisas ao mesmo tempo (ou parecer que está fazendo)

---

## 🏃 Goroutines: Os "Trabalhadores Paralelos"

### Analogia: Equipe de Limpeza

Imagine que você precisa limpar uma casa grande:

**Sem Goroutines (Sequencial):**
```
1. Limpar quarto 1 (30 min)
2. Limpar quarto 2 (30 min)
3. Limpar quarto 3 (30 min)
4. Limpar cozinha (30 min)
Total: 2 horas 😴
```

**Com Goroutines (Concorrente):**
```
Trabalhador 1: Limpar quarto 1 (30 min)
Trabalhador 2: Limpar quarto 2 (30 min)
Trabalhador 3: Limpar quarto 3 (30 min)
Trabalhador 4: Limpar cozinha (30 min)
Total: 30 minutos! 🚀
```

**Em Go:**
```go
// Criar 4 "trabalhadores" (goroutines)
go limparQuarto(1)
go limparQuarto(2)
go limparQuarto(3)
go limparCozinha()
```

**Pense assim:**
- **Goroutine** = um trabalhador que faz uma tarefa
- **Leve** = criar um trabalhador é muito barato (não precisa contratar, só "falar")
- **Muitos** = você pode ter milhares de trabalhadores trabalhando ao mesmo tempo

---

## 📞 Channels: A "Linha de Comunicação"

### Analogia: Sistema de Entrega de Pizza

Imagine uma pizzaria:

**O Problema:**
- A cozinha faz as pizzas
- Os entregadores precisam saber quando uma pizza está pronta
- Como eles se comunicam?

**A Solução: Channels!**

```
Cozinha → [Channel] → Entregador 1
         → [Channel] → Entregador 2
         → [Channel] → Entregador 3
```

**Como funciona:**
1. Cozinha coloca pizza pronta no channel: `channel <- pizza`
2. Entregador pega pizza do channel: `pizza := <-channel`
3. Se não tem pizza, entregador espera (bloqueia)
4. Se tem pizza, entregador pega e sai para entregar

**Em Go:**
```go
// Cozinha envia pizza
pizzas <- "Pizza Margherita"

// Entregador recebe pizza
pizza := <-pizzas
fmt.Println("Entregando:", pizza)
```

**Pense assim:**
- **Channel** = uma caixa onde você coloca coisas e outras pessoas pegam
- **Enviar** (`ch <- valor`) = colocar algo na caixa
- **Receber** (`valor := <-ch`) = pegar algo da caixa
- **Bloqueio** = se a caixa está vazia, você espera até ter algo

---

## 🔄 Buffered vs Unbuffered: Caixa Pequena ou Grande?

### Analogia: Fila de Banco

**Unbuffered Channel = Caixa de 1 lugar (Fila Individual)**

```
Cliente → [Caixa] → Atendente
```

- Cliente chega e **espera** até o atendente estar livre
- Atendente **espera** até ter um cliente
- Só funciona se ambos estão prontos ao mesmo tempo
- **Síncrono** = tudo acontece no mesmo momento

**Buffered Channel = Caixa de Múltiplos Lugares (Fila com Espaço)**

```
Cliente 1 → [Caixa]
Cliente 2 → [Caixa] → Atendente
Cliente 3 → [Caixa]
```

- Cliente pode deixar na caixa e ir embora (se tiver espaço)
- Atendente pega quando está livre
- **Assíncrono** = não precisa esperar o outro estar pronto

**Em Go:**
```go
// Unbuffered (caixa de 1 lugar)
ch1 := make(chan int)

// Buffered (caixa de 10 lugares)
ch2 := make(chan int, 10)
```

**Quando usar cada um?**
- **Unbuffered**: Quando precisa garantir que alguém recebeu antes de continuar (como entregar documento importante)
- **Buffered**: Quando pode deixar várias coisas na caixa e processar depois (como emails)

---

## 🎯 Select: O "Seletor de Canais"

### Analogia: Atendente de Telefone com Múltiplas Linhas

Imagine um atendente com 3 telefones:

**Sem Select:**
- Atendente só pode atender um telefone por vez
- Se o Telefone 1 tocar, ele atende
- Mas e se o Telefone 2 ou 3 tocarem? Ele não sabe!

**Com Select:**
- Atendente **escuta** todos os telefones ao mesmo tempo
- Quando **qualquer um** tocar, ele atende aquele
- Se **múltiplos** tocarem, ele escolhe um (geralmente o primeiro)

**Em Go:**
```go
select {
case chamada1 := <-telefone1:
    fmt.Println("Atendendo linha 1:", chamada1)
case chamada2 := <-telefone2:
    fmt.Println("Atendendo linha 2:", chamada2)
case chamada3 := <-telefone3:
    fmt.Println("Atendendo linha 3:", chamada3)
}
```

**Pense assim:**
- **Select** = você fica "de ouvido" em múltiplos channels
- Quando **qualquer um** tiver algo, você pega daquele
- Muito útil para **timeouts** e **cancelamento**

---

## 👷 Worker Pools: A "Equipe Organizada"

### Analogia: Fábrica com Número Fixo de Operários

Imagine uma fábrica que precisa processar 100 peças:

**Sem Worker Pool (Criar 100 Operários):**
- Contrata 100 operários
- Cada um processa 1 peça
- Problema: Muito caro! E se precisar processar 10.000 peças?

**Com Worker Pool (Equipe Fixa):**
- Contrata apenas 5 operários (número fixo)
- Cria uma **fila de trabalho**
- Operários pegam peças da fila, processam, pegam outra
- Resultado: Eficiente e controlado!

**Em Go:**
```go
// Criar 5 "operários" (workers)
for i := 1; i <= 5; i++ {
    go worker(filaDeTrabalho)
}

// Colocar 100 "peças" na fila
for i := 1; i <= 100; i++ {
    filaDeTrabalho <- peca
}
```

**Pense assim:**
- **Worker Pool** = equipe fixa de trabalhadores
- **Fila de Trabalho** = channel com tarefas
- **Controle** = você decide quantos trabalhadores ter
- **Eficiência** = não cria trabalhadores demais

---

## 🔒 Mutex: O "Cartão de Acesso Único"

### Analogia: Banheiro Público

Imagine um banheiro com apenas 1 cabine:

**Sem Mutex (Problema):**
- Pessoa 1 entra
- Pessoa 2 tenta entrar ao mesmo tempo
- **Conflito!** Duas pessoas na mesma cabine! 😱

**Com Mutex (Solução):**
- Pessoa 1 pega o "cartão de acesso" (Lock)
- Entra no banheiro
- Pessoa 2 tenta pegar o cartão, mas está ocupado
- Pessoa 2 **espera** até Pessoa 1 sair
- Pessoa 1 sai e devolve o cartão (Unlock)
- Agora Pessoa 2 pode pegar o cartão e entrar

**Em Go:**
```go
var mu sync.Mutex // O "cartão de acesso"

// Pessoa 1
mu.Lock()         // Pega o cartão
// Usa o banheiro
mu.Unlock()       // Devolve o cartão

// Pessoa 2
mu.Lock()         // Espera até o cartão estar disponível
// Usa o banheiro
mu.Unlock()       // Devolve o cartão
```

**Pense assim:**
- **Mutex** = um cartão que só uma pessoa pode ter por vez
- **Lock()** = pegar o cartão (se estiver disponível, ou esperar)
- **Unlock()** = devolver o cartão
- **Proteção** = garante que apenas uma goroutine acessa algo por vez

---

## ⏳ WaitGroup: O "Contador de Tarefas"

### Analogia: Professor Esperando Alunos Terminarem a Prova

Imagine um professor que deu uma prova:

**O Problema:**
- 30 alunos fazendo a prova
- Cada um termina em tempos diferentes
- Professor precisa esperar **todos** terminarem antes de recolher

**A Solução: WaitGroup!**

```
Professor: "Vou esperar 30 alunos terminarem"
Aluno 1 termina → Contador: 29
Aluno 2 termina → Contador: 28
...
Aluno 30 termina → Contador: 0
Professor: "Todos terminaram! Posso recolher as provas"
```

**Em Go:**
```go
var wg sync.WaitGroup

// Professor diz: "Espero 30 alunos"
wg.Add(30)

// Cada aluno, quando termina:
go func() {
    defer wg.Done() // "Terminei!"
    fazerProva()
}()

// Professor espera todos
wg.Wait() // Bloqueia até contador chegar a 0
fmt.Println("Todos terminaram!")
```

**Pense assim:**
- **WaitGroup** = um contador de tarefas
- **Add(n)** = "espero n tarefas terminarem"
- **Done()** = "terminei uma tarefa" (decrementa contador)
- **Wait()** = "espero até contador chegar a zero"

---

## 🔄 RWMutex: O "Cartão de Leitura vs Escrita"

### Analogia: Biblioteca

Imagine uma biblioteca:

**Mutex Normal (Apenas 1 pessoa por vez):**
- Se alguém está lendo, ninguém mais pode entrar
- Se alguém está escrevendo, ninguém mais pode entrar
- **Ineficiente** para leitura!

**RWMutex (Múltiplos Leitores, 1 Escritor):**
- **Múltiplas pessoas podem ler ao mesmo tempo** ✅
- Mas apenas **1 pessoa pode escrever** por vez ✅
- Quando alguém escreve, todos os leitores esperam

**Em Go:**
```go
var mu sync.RWMutex

// Leitores (podem ler juntos)
mu.RLock()   // Lock para leitura
ler()
mu.RUnlock()

// Escritor (exclusivo)
mu.Lock()    // Lock para escrita (bloqueia todos)
escrever()
mu.Unlock()
```

**Pense assim:**
- **RWMutex** = cartão especial de biblioteca
- **RLock()** = "vou ler" (várias pessoas podem ter)
- **Lock()** = "vou escrever" (apenas 1 pessoa pode ter, bloqueia leitores)

---

## 🎭 Padrões Comuns: Fan-Out e Fan-In

### Analogia: Distribuição de Jornais

**Fan-Out (Distribuir):**
```
Editora → [Channel] → Entregador 1
        → [Channel] → Entregador 2
        → [Channel] → Entregador 3
```
- Uma fonte (editora) distribui para múltiplos workers (entregadores)
- Cada entregador pega jornais do mesmo channel

**Fan-In (Combinar):**
```
Entregador 1 → [Channel]
Entregador 2 → [Channel] → Central de Recebimento
Entregador 3 → [Channel]
```
- Múltiplas fontes (entregadores) enviam para um único destino (central)
- Central recebe de todos e processa

**Em Go:**
```go
// Fan-Out: Múltiplos workers lendo do mesmo channel
in := producer()
c1 := worker(in)
c2 := worker(in)
c3 := worker(in)

// Fan-In: Combinar resultados
for resultado := range merge(c1, c2, c3) {
    processar(resultado)
}
```

---

## 🎯 Resumo com Analogias

| Conceito | Analogia | Em Go |
|----------|----------|-------|
| **Goroutine** | Trabalhador | `go funcao()` |
| **Channel** | Caixa de comunicação | `ch := make(chan int)` |
| **Unbuffered** | Caixa de 1 lugar | `make(chan int)` |
| **Buffered** | Caixa de múltiplos lugares | `make(chan int, 10)` |
| **Select** | Atendente de múltiplos telefones | `select { case ... }` |
| **Worker Pool** | Equipe fixa de operários | Múltiplos workers + channel |
| **Mutex** | Cartão de acesso único | `mu.Lock()` / `mu.Unlock()` |
| **RWMutex** | Biblioteca (múltiplos leitores) | `mu.RLock()` / `mu.Lock()` |
| **WaitGroup** | Contador de tarefas | `wg.Add()` / `wg.Done()` / `wg.Wait()` |

---

## 💡 Dicas Finais

1. **Goroutines são leves**: Crie quantas precisar, mas não abuse!
2. **Channels para comunicação**: Use channels para goroutines se comunicarem
3. **Mutex para proteção**: Use mutex para proteger dados compartilhados
4. **WaitGroup para esperar**: Use quando precisar esperar múltiplas goroutines
5. **Worker Pools para controle**: Use quando quiser limitar concorrência

**Lembre-se:**
- Concorrência = fazer várias coisas ao mesmo tempo (ou parecer que está)
- Go torna isso fácil com goroutines e channels
- A prática leva à perfeição!

Na próxima aula, vamos colocar tudo isso em prática com exercícios! 🚀

