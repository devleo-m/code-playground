# Aula 17 - Exercícios e Reflexão: Context Package

Olá! Agora é hora de colocar em prática tudo que você aprendeu sobre Context em Go. Vamos começar com exercícios práticos e depois refletir sobre os conceitos.

---

## 📝 Exercícios Práticos

### Exercício 1: Timeout em Operação Longa

Crie uma função que simula uma operação longa (como processar dados) e implemente um timeout usando context.

**Requisitos:**
- Crie uma função `processarDados(ctx context.Context, dados []int) error`
- A função deve processar cada item da lista, simulando trabalho com `time.Sleep(200 * time.Millisecond)`
- Use `context.WithTimeout` para limitar a operação a 2 segundos
- Se o timeout ocorrer, a função deve retornar o erro do context
- No `main`, teste com uma lista de 20 itens (que levaria 4 segundos sem timeout)

**Exemplo de estrutura:**
```go
func processarDados(ctx context.Context, dados []int) error {
    // Seu código aqui
    // Verificar ctx.Done() em cada iteração
    // Retornar ctx.Err() se cancelado
}

func main() {
    dados := make([]int, 20)
    for i := range dados {
        dados[i] = i + 1
    }
    
    // Criar context com timeout de 2 segundos
    // Chamar processarDados
    // Tratar erro de timeout
}
```

**Resultado esperado:**
- A função deve processar alguns itens (cerca de 10, já que cada um leva 200ms)
- Após 2 segundos, deve retornar erro de timeout
- Deve imprimir quantos itens foram processados antes do timeout

**Dica:** Use `select` com `ctx.Done()` dentro do loop para verificar cancelamento.

---

### Exercício 2: Cancelamento Manual de Múltiplas Goroutines

Crie um programa que inicia múltiplas goroutines fazendo trabalho e permite cancelá-las todas de uma vez.

**Requisitos:**
- Crie 5 goroutines que fazem trabalho contínuo (loop infinito)
- Cada goroutine deve verificar `ctx.Done()` e parar quando cancelada
- Use `context.WithCancel` para criar um contexto cancelável
- Após 3 segundos, cancele todas as goroutines
- Cada goroutine deve imprimir seu ID e quando foi cancelada

**Exemplo de estrutura:**
```go
func worker(ctx context.Context, id int) {
    // Loop infinito que verifica ctx.Done()
    // Imprimir "Worker X trabalhando..."
    // Quando cancelado, imprimir "Worker X cancelado"
}

func main() {
    // Criar context com cancel
    // Iniciar 5 workers
    // Esperar 3 segundos
    // Cancelar todos
    // Dar tempo para workers terminarem
}
```

**Resultado esperado:**
- Workers devem imprimir "trabalhando" várias vezes
- Após 3 segundos, todos devem ser cancelados
- Cada worker deve imprimir mensagem de cancelamento

**Dica:** Use `select` com `default` para fazer trabalho enquanto verifica cancelamento.

---

### Exercício 3: Context com Valores (Request-Scoped)

Implemente um sistema simples que passa informações de uma requisição através de múltiplas funções usando context.

**Requisitos:**
- Crie tipos específicos para as chaves: `userIDKey` e `requestIDKey`
- Crie função `autenticar(ctx context.Context) context.Context` que adiciona userID
- Crie função `adicionarRequestID(ctx context.Context) context.Context` que adiciona requestID
- Crie função `processarRequisicao(ctx context.Context)` que recupera e imprime ambos os valores
- No `main`, crie uma cadeia: Background → autenticar → adicionarRequestID → processarRequisicao

**Exemplo de estrutura:**
```go
type userIDKey struct{}
type requestIDKey struct{}

func autenticar(ctx context.Context) context.Context {
    // Adicionar userID "12345" ao context
}

func adicionarRequestID(ctx context.Context) context.Context {
    // Adicionar requestID "req-abc-xyz" ao context
}

func processarRequisicao(ctx context.Context) {
    // Recuperar userID e requestID
    // Imprimir: "Processando requisição req-abc-xyz para usuário 12345"
}

func main() {
    // Criar cadeia de contextos
}
```

**Resultado esperado:**
- Deve imprimir corretamente os valores passados através do context
- Deve demonstrar como valores podem ser passados através de múltiplas camadas

**Dica:** Use type assertion para recuperar valores: `ctx.Value(userIDKey{}).(string)`

---

### Exercício 4: Context em Operação com Múltiplas Etapas

Crie uma função que realiza múltiplas etapas de processamento, cada uma com seu próprio timeout, mas todas respeitando um timeout global.

**Requisitos:**
- Crie função `processarComEtapas(ctx context.Context) error`
- A função deve ter 3 etapas:
  1. `etapa1`: Simula 1 segundo de trabalho
  2. `etapa2`: Simula 1 segundo de trabalho
  3. `etapa3`: Simula 1 segundo de trabalho
- Cada etapa deve verificar se o context foi cancelado
- No `main`, crie um context com timeout global de 2 segundos
- Teste o comportamento quando o timeout global é menor que o tempo total das etapas

**Exemplo de estrutura:**
```go
func etapa1(ctx context.Context) error {
    // Verificar ctx.Done()
    // Simular 1 segundo de trabalho
    // Retornar erro se cancelado
}

func etapa2(ctx context.Context) error {
    // Similar a etapa1
}

func etapa3(ctx context.Context) error {
    // Similar a etapa1
}

func processarComEtapas(ctx context.Context) error {
    // Chamar etapa1, etapa2, etapa3 em sequência
    // Retornar erro se qualquer etapa falhar
}

func main() {
    // Context com timeout de 2 segundos
    // Chamar processarComEtapas
    // Tratar erro de timeout
}
```

**Resultado esperado:**
- Com timeout de 2 segundos, deve completar etapa1 e etapa2
- Deve falhar na etapa3 devido ao timeout
- Deve demonstrar como cancelamento se propaga através de múltiplas etapas

**Dica:** Verifique `ctx.Done()` antes e durante cada etapa.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por Que Context é o Primeiro Parâmetro?

Em Go, é uma convenção **sempre** passar `context.Context` como o **primeiro parâmetro** de funções.

**Perguntas para refletir:**

1. **Por que essa convenção existe?**
   - Qual é a razão prática de sempre colocar context primeiro?
   - Como isso ajuda na legibilidade e manutenção do código?

2. **O que acontece se você não seguir essa convenção?**
   - Quais problemas podem surgir?
   - Como isso afeta outros desenvolvedores que usam seu código?

3. **Pense em uma situação real:**
   - Você está criando uma função que faz uma chamada HTTP
   - A função precisa aceitar: context, URL, método HTTP, headers, body
   - Como você organizaria esses parâmetros? Por quê?

**Escreva sua resposta aqui:**
```
[Seu espaço para refletir e escrever]
```

---

### Reflexão 2: Quando Usar Valores no Context vs Parâmetros de Função?

O context pode carregar valores, mas há uma regra importante: **não use context para passar parâmetros opcionais de função**.

**Perguntas para refletir:**

1. **Qual é a diferença entre valores no context e parâmetros de função?**
   - Quando você deve usar `context.WithValue()`?
   - Quando você deve usar parâmetros normais de função?
   - Dê exemplos concretos de cada caso.

2. **Por que não devemos usar context para parâmetros opcionais?**
   - Quais problemas isso causa?
   - Como isso torna o código mais difícil de entender e manter?

3. **Cenário prático:**
   - Você está criando uma função `buscarUsuario(ctx context.Context, userID string, incluirEndereco bool)`
   - O parâmetro `incluirEndereco` é opcional
   - Alguém sugere passar `incluirEndereco` através do context: `ctx = context.WithValue(ctx, "incluirEndereco", true)`
   - Por que isso é uma má ideia? Como você faria corretamente?

4. **Pense em valores legítimos para context:**
   - Request ID: Por que faz sentido estar no context?
   - User ID (após autenticação): Por que faz sentido estar no context?
   - Timeout específico: Por que NÃO faz sentido estar no context?

**Escreva sua resposta aqui:**
```
[Seu espaço para refletir e escrever]
```

---

### Reflexão 3: Timeout vs Deadline - Quando Usar Cada Um?

Tanto `WithTimeout` quanto `WithDeadline` criam contextos que são cancelados após um período de tempo, mas há diferenças importantes.

**Perguntas para refletir:**

1. **Qual é a diferença prática entre timeout e deadline?**
   - Em que situações você usaria `WithTimeout`?
   - Em que situações você usaria `WithDeadline`?
   - Dê exemplos reais de cada caso.

2. **Cenário 1: Requisição HTTP**
   - Você está criando um cliente HTTP que faz requisições para uma API externa
   - Você quer que cada requisição tenha no máximo 5 segundos
   - Qual você usaria: `WithTimeout` ou `WithDeadline`? Por quê?

3. **Cenário 2: Processamento em Lote**
   - Você tem um job que processa dados em lotes
   - O job deve terminar antes das 23:59:59 de hoje (deadline absoluto)
   - Cada lote pode levar até 30 segundos (timeout por lote)
   - Como você combinaria timeout e deadline? Por quê?

4. **Pense em um sistema real:**
   - Um sistema de backup que roda diariamente
   - O backup deve começar à meia-noite
   - Cada arquivo tem timeout de 10 segundos
   - O backup completo tem deadline de 6:00 AM
   - Como você estruturaria os contextos? Explique sua escolha.

**Escreva sua resposta aqui:**
```
[Seu espaço para refletir e escrever]
```

---

### Reflexão 4: Cancelamento e Limpeza de Recursos

Quando um context é cancelado, as operações devem parar e recursos devem ser liberados.

**Perguntas para refletir:**

1. **Por que é importante sempre chamar `defer cancel()`?**
   - O que acontece se você não cancelar um context?
   - Quais recursos podem vazar?
   - Dê um exemplo de um vazamento de recurso que pode ocorrer.

2. **Cenário: Operação de Banco de Dados**
   - Você cria um context com timeout de 5 segundos
   - Você inicia uma query que pode demorar
   - A query é cancelada após 5 segundos
   - O que mais você precisa fazer além de cancelar o context?
   - Como garantir que a conexão com o banco seja fechada corretamente?

3. **Pense em uma função que abre um arquivo:**
   ```go
   func processarArquivo(ctx context.Context, nomeArquivo string) error {
       file, err := os.Open(nomeArquivo)
       // ... processar arquivo
   }
   ```
   - O que está faltando nessa função?
   - Como você garantiria que o arquivo seja fechado mesmo se o context for cancelado?
   - Como você estruturaria isso corretamente?

4. **Hierarquia de cancelamento:**
   - Você tem um context pai com timeout de 10 segundos
   - Você deriva um context filho com timeout de 2 segundos
   - Se o context filho for cancelado, o pai também é cancelado?
   - Se o context pai for cancelado, o filho também é cancelado?
   - Por que essa hierarquia é importante para limpeza de recursos?

**Escreva sua resposta aqui:**
```
[Seu espaço para refletir e escrever]
```

---

## 📋 Checklist de Aprendizado

Antes de prosseguir, verifique se você consegue:

- [ ] Criar contextos com timeout, deadline e cancelamento manual
- [ ] Passar valores através do context usando tipos específicos
- [ ] Verificar se um context foi cancelado usando `ctx.Done()` e `ctx.Err()`
- [ ] Implementar cancelamento em loops e operações longas
- [ ] Usar context em funções seguindo a convenção (primeiro parâmetro)
- [ ] Entender quando usar valores no context vs parâmetros de função
- [ ] Entender a diferença entre timeout e deadline
- [ ] Sempre usar `defer cancel()` após criar contextos com cancelamento
- [ ] Entender como cancelamento se propaga em contextos aninhados

---

## 🎯 Próximos Passos

Após completar os exercícios e reflexões:

1. **Revise suas respostas** às perguntas de reflexão
2. **Teste seus códigos** e certifique-se de que funcionam corretamente
3. **Experimente variações** dos exercícios para fixar o aprendizado
4. **Pense em situações reais** onde você usaria context no seu trabalho

Na próxima aula, vamos ver performance, boas práticas e o que você deve ou não fazer com context!

Boa sorte com os exercícios! 🚀

