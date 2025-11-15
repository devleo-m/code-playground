# Aula 16 - Exercícios e Reflexão: Concorrência

Olá! Agora é hora de colocar em prática tudo que você aprendeu sobre concorrência em Go. Vamos começar com exercícios práticos e depois refletir sobre os conceitos.

---

## 📝 Exercícios Práticos

### Exercício 1: Contador Concorrente com Mutex

Crie um programa que incrementa um contador compartilhado usando múltiplas goroutines, protegendo o acesso com Mutex.

**Requisitos:**
- Crie 10 goroutines
- Cada goroutine incrementa o contador 100 vezes
- Use Mutex para proteger o contador
- O valor final deve ser exatamente 1000

**Exemplo de estrutura:**
```go
var (
    contador int
    mu       sync.Mutex
)

func incrementar() {
    // Seu código aqui
}

func main() {
    var wg sync.WaitGroup
    
    // Criar 10 goroutines
    // Cada uma incrementa 100 vezes
    // Esperar todas terminarem
    // Imprimir resultado final
}
```

**Dica:** Lembre-se de usar `defer mu.Unlock()` após `mu.Lock()`.

---

### Exercício 2: Worker Pool para Processar Números

Crie um worker pool que processa números, calculando o quadrado de cada um.

**Requisitos:**
- Crie 3 workers
- Envie números de 1 a 20 para processar
- Cada worker calcula o quadrado do número
- Use channels para distribuir trabalho e coletar resultados
- Imprima todos os resultados

**Exemplo de uso esperado:**
```go
// Estrutura esperada:
jobs := make(chan int, 20)
results := make(chan int, 20)

// Criar 3 workers
// Enviar números 1-20 para jobs
// Coletar resultados
// Imprimir: "1² = 1", "2² = 4", etc.
```

**Dica:** Use `close(jobs)` após enviar todos os números para sinalizar que não há mais trabalho.

---

### Exercício 3: Select com Timeout

Crie um programa que tenta receber um valor de um channel, mas implementa um timeout de 2 segundos.

**Requisitos:**
- Crie um channel
- Em uma goroutine, espere 3 segundos antes de enviar um valor
- No main, use `select` com `time.After` para implementar timeout de 2 segundos
- Se o timeout ocorrer antes de receber o valor, imprima "Timeout!"
- Se receber o valor a tempo, imprima o valor

**Exemplo de código base:**
```go
ch := make(chan string)

go func() {
    time.Sleep(3 * time.Second)
    ch <- "resultado"
}()

// Use select aqui com timeout
```

**Dica:** `time.After(2 * time.Second)` retorna um channel que recebe um valor após 2 segundos.

---

### Exercício 4: Cache Thread-Safe com RWMutex

Implemente um cache simples que suporta múltiplas leituras simultâneas, mas apenas uma escrita por vez.

**Requisitos:**
- Crie um tipo `Cache` com um map interno
- Use `RWMutex` para proteção
- Implemente método `Get(key string) (string, bool)`
- Implemente método `Set(key, value string)`
- Teste com múltiplos leitores e escritores simultâneos

**Exemplo de uso:**
```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
    // Seu código aqui
}

func (c *Cache) Set(key, value string) {
    // Seu código aqui
}

func main() {
    cache := NewCache()
    
    // Criar múltiplos leitores e escritores
    // Testar concorrência
}
```

**Dica:** Use `RLock()` para leitura e `Lock()` para escrita.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por Que "Don't Communicate by Sharing Memory"?

A filosofia do Go é: **"Don't communicate by sharing memory; share memory by communicating"**.

**Perguntas para refletir:**

1. **O que isso significa na prática?**
   - Explique com suas próprias palavras o que essa filosofia representa
   - Dê um exemplo prático de quando você usaria channels vs quando usaria mutex

2. **Quais são as vantagens dessa abordagem?**
   - Por que Go prefere channels sobre mutexes quando possível?
   - Quais problemas essa abordagem ajuda a evitar?

3. **Quando você deve quebrar essa regra?**
   - Em que situações faz sentido usar mutex diretamente?
   - Dê um exemplo real onde mutex é mais apropriado que channels

**Escreva suas respostas aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 2: Buffered vs Unbuffered Channels

**Perguntas para refletir:**

1. **Qual é a diferença fundamental?**
   - Explique a diferença entre buffered e unbuffered channels
   - Quando cada um bloqueia e por quê?

2. **Quando usar cada um?**
   - Dê um exemplo de situação onde unbuffered é melhor
   - Dê um exemplo de situação onde buffered é melhor
   - O que acontece se você escolher errado?

3. **Qual é o risco de usar buffered channels?**
   - Por que buffered channels podem "mascarar" problemas?
   - Como você pode detectar se está usando o tamanho correto do buffer?

**Escreva suas respostas aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 3: Worker Pools e Controle de Recursos

**Perguntas para refletir:**

1. **Por que limitar o número de goroutines?**
   - Se goroutines são "leves", por que não criar milhares delas?
   - Quais são os limites práticos de goroutines?

2. **Como você decide o número de workers?**
   - Se você tem 4 núcleos de CPU, quantos workers você deve criar?
   - A resposta é diferente para operações I/O vs CPU-bound?
   - Como você descobriria o número ideal?

3. **Worker Pools vs Criar Goroutine para Cada Tarefa:**
   - Quando faz sentido criar uma goroutine para cada tarefa?
   - Quando faz sentido usar worker pool?
   - Dê exemplos práticos de cada caso

**Escreva suas respostas aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 4: Race Conditions e Sincronização

**Perguntas para refletir:**

1. **O que é uma race condition?**
   - Explique com suas próprias palavras
   - Por que elas são perigosas?
   - Dê um exemplo de como uma race condition pode causar bugs difíceis de encontrar

2. **Como Go ajuda a prevenir race conditions?**
   - Quais ferramentas Go oferece?
   - Por que channels são mais seguros que mutexes em muitos casos?

3. **Quando você precisa de sincronização?**
   - Sempre que múltiplas goroutines acessam os mesmos dados?
   - Há casos onde não precisa de sincronização?
   - Como você identifica se precisa proteger algo?

**Escreva suas respostas aqui:**
```
[Seu espaço para escrever]
```

---

### Reflexão 5: Aplicação Prática

**Cenário:** Você precisa criar um sistema que:
- Faz requisições HTTP para 100 URLs diferentes
- Processa as respostas (extrai dados)
- Salva os resultados em um arquivo

**Perguntas para refletir:**

1. **Como você estruturaria isso usando concorrência?**
   - Quantas goroutines você criaria? Por quê?
   - Usaria worker pool? Por quê?
   - Como gerenciaria erros?

2. **Quais primitivos de sincronização você usaria?**
   - Channels? Mutex? WaitGroup? Todos?
   - Explique sua escolha

3. **Quais problemas potenciais você precisa considerar?**
   - O que pode dar errado?
   - Como você garantiria que todos os resultados sejam salvos?
   - Como lidaria com timeouts?

**Escreva suas respostas aqui:**
```
[Seu espaço para escrever]
```

---

## 📋 Checklist de Aprendizado

Antes de prosseguir, certifique-se de que você consegue:

- [ ] Criar e usar goroutines
- [ ] Criar e usar channels (buffered e unbuffered)
- [ ] Usar `select` para multiplexar channels
- [ ] Implementar worker pools
- [ ] Usar Mutex para proteger dados compartilhados
- [ ] Usar RWMutex quando apropriado
- [ ] Usar WaitGroup para coordenar goroutines
- [ ] Entender quando usar channels vs mutexes
- [ ] Identificar e evitar race conditions
- [ ] Implementar timeouts com select

---

## 🎯 Próximos Passos

Após completar os exercícios e reflexões:

1. **Revise suas respostas** - Certifique-se de que entendeu os conceitos
2. **Teste seus códigos** - Execute os exercícios e verifique se funcionam
3. **Experimente variações** - Tente modificar os exercícios para entender melhor
4. **Pesquise mais** - Se algo não ficou claro, pesquise e pratique mais

**Lembre-se:** Concorrência é um tópico complexo. Não se preocupe se não entender tudo de primeira. A prática constante é a chave!

Na próxima aula, vamos aprender sobre performance, boas práticas e quando usar cada ferramenta de concorrência! 🚀

