# Aula 2 - Simplificada: Entendendo Error Handling em Go

Olá! Na aula anterior, mergulhamos nos detalhes técnicos do tratamento de erros em Go. Agora, vamos simplificar tudo isso usando analogias do dia a dia para que você realmente **entenda** e não apenas **decorar** os conceitos.

---

## 1. Erros são Valores: A Analogia do Restaurante

Imagine que você está em um restaurante fazendo um pedido. Em algumas linguagens (com exceções), seria assim:

**Sistema com Exceções (outras linguagens):**
```
Garçom: "Um prato de lasanha, por favor!"
Cozinha: [lança uma exceção que voa pela sala]
[Todo o restaurante para, sirenes tocam, clientes fogem]
```

**Sistema do Go (erros são valores):**
```
Garçom: "Um prato de lasanha, por favor!"
Cozinha: "Desculpe, acabou a lasanha. Aqui está um prato vazio e uma nota explicando o problema."
Garçom: [lê a nota] "Entendi, vou oferecer outra opção ao cliente."
```

Em Go, o erro **não interrompe tudo**. Ele é apenas uma **informação** que você recebe junto com o resultado. Você decide o que fazer com essa informação.

### O Padrão `if err != nil`

É como verificar se você recebeu o que pediu:

```go
prato, nota := pedirLasanha()
if nota != nil {  // Se há uma nota (erro)
    fmt.Println("Problema:", nota)
    pedirOutroPrato()  // Você decide o que fazer
    return
}
// Se não há nota, tudo certo! Pode comer o prato
comer(prato)
```

---

## 2. A Interface `error`: A Analogia do Formulário

A interface `error` é como um formulário padrão que qualquer pessoa pode preencher:

```go
type error interface {
    Error() string
}
```

**Analogia:** Imagine que você trabalha em uma empresa e há um formulário de "Relatório de Problema". Qualquer pessoa pode preencher esse formulário, desde que ela escreva algo na linha "Descrição do Problema".

- Um funcionário pode escrever: "A impressora está sem papel"
- Outro pode escrever: "O computador não liga"
- Você pode criar seu próprio formulário personalizado com mais campos, mas ainda precisa ter a linha "Descrição do Problema"

Da mesma forma:
- `errors.New("algo deu errado")` cria um formulário simples
- Você pode criar seu próprio tipo de erro (formulário personalizado), mas precisa ter o método `Error() string` (a linha obrigatória)

---

## 3. `errors.New()` vs `fmt.Errorf()`: Notas Simples vs Notas Detalhadas

### `errors.New()` - Nota Simples

É como deixar um bilhete simples na mesa:

```go
errors.New("divisão por zero")
```

**Analogia:** Você deixa um bilhete: "Não funcionou"

### `fmt.Errorf()` - Nota Detalhada

É como deixar um bilhete com todos os detalhes:

```go
fmt.Errorf("não foi possível dividir %d por %d: divisão por zero", a, b)
```

**Analogia:** Você deixa um bilhete: "Não foi possível dividir 10 por 0: divisão por zero não é permitida"

**Quando usar cada um?**
- `errors.New()`: Quando a mensagem é fixa e simples
- `fmt.Errorf()`: Quando você precisa incluir informações dinâmicas (valores de variáveis)

---

## 4. Error Wrapping: A Analogia da Carta em Cadeia

Imagine que você está enviando uma carta, mas ela precisa passar por várias pessoas antes de chegar ao destinatário. Cada pessoa adiciona uma nota explicando o que aconteceu:

**Sem Error Wrapping (perde contexto):**
```
Você → João → Maria → Destinatário

João: "Não consegui entregar"
Maria: "Não consegui entregar"  (perdeu a informação de que foi João)
Destinatário: "Alguém não conseguiu entregar, mas não sei quem ou por quê"
```

**Com Error Wrapping (preserva contexto):**
```
Você → João → Maria → Destinatário

João: "Não consegui entregar porque estava chovendo"
Maria: "João não conseguiu entregar porque estava chovendo, e eu também não consegui porque o endereço estava errado"
Destinatário: "Maria não conseguiu porque o endereço estava errado, e antes disso João não conseguiu porque estava chovendo"
```

Em Go:

```go
// Função interna
func abrirArquivo(nome string) error {
    return fmt.Errorf("arquivo %s não encontrado", nome)
}

// Função intermediária
func lerArquivo(nome string) error {
    err := abrirArquivo(nome)
    if err != nil {
        return fmt.Errorf("erro ao ler arquivo: %w", err)  // %w preserva o erro original
    }
    return nil
}

// Função externa
func processarArquivo(nome string) error {
    err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("falha ao processar: %w", err)  // adiciona mais contexto
    }
    return nil
}
```

Quando ocorrer um erro, você terá:
```
"falha ao processar: erro ao ler arquivo: arquivo dados.txt não encontrado"
```

**Toda a cadeia de contexto está preservada!**

---

## 5. Sentinel Errors: A Analogia dos Códigos de Erro do Banco

Quando você vai ao banco e algo dá errado, o caixa não te dá uma mensagem genérica. Ele te dá um **código específico**:

- **Código 001**: "Saldo insuficiente"
- **Código 002**: "Conta não encontrada"
- **Código 003**: "Senha incorreta"

Você sabe exatamente o que fazer com cada código:
- Código 001 → Depositar mais dinheiro
- Código 002 → Verificar o número da conta
- Código 003 → Tentar novamente com a senha correta

**Sentinel Errors** funcionam da mesma forma:

```go
var (
    ErrSaldoInsuficiente = errors.New("saldo insuficiente")
    ErrContaNaoEncontrada = errors.New("conta não encontrada")
    ErrSenhaIncorreta = errors.New("senha incorreta")
)

func sacarDinheiro(conta, senha string, valor float64) error {
    // verifica conta
    if !existeConta(conta) {
        return ErrContaNaoEncontrada
    }
    // verifica senha
    if !senhaCorreta(conta, senha) {
        return ErrSenhaIncorreta
    }
    // verifica saldo
    if saldo(conta) < valor {
        return ErrSaldoInsuficiente
    }
    return nil
}

// Uso:
err := sacarDinheiro("123", "senha", 1000)
if errors.Is(err, ErrSaldoInsuficiente) {
    fmt.Println("Você precisa depositar mais dinheiro")
} else if errors.Is(err, ErrSenhaIncorreta) {
    fmt.Println("Tente novamente com a senha correta")
} else if errors.Is(err, ErrContaNaoEncontrada) {
    fmt.Println("Verifique o número da conta")
}
```

**Por que isso é útil?**
- Você pode tratar cada erro de forma diferente
- A API fica clara e previsível
- Outros programadores sabem exatamente quais erros esperar

---

## 6. `errors.Is()` e `errors.As()`: Encontrando o Erro Certo na Cadeia

### `errors.Is()`: Procurando um Erro Específico

Imagine que você está procurando uma carta específica em uma pilha de cartas. Cada carta pode ter outras cartas dentro (error wrapping). `errors.Is()` procura em **todas as camadas**:

```go
// Você tem uma cadeia de erros:
// "erro ao processar: erro ao ler: arquivo não encontrado"

err := processarArquivo("dados.txt")
if errors.Is(err, os.ErrNotExist) {
    // Encontrou! Mesmo que esteja "escondido" dentro de outros erros
    fmt.Println("O arquivo realmente não existe")
}
```

**Analogia:** É como procurar uma chave específica em uma caixa que está dentro de outra caixa, que está dentro de outra caixa. `errors.Is()` abre todas as caixas até encontrar a chave que você procura.

### `errors.As()`: Extraindo um Tipo de Erro Específico

Imagine que você tem uma pilha de formulários de diferentes tipos (formulário de reclamação, formulário de sugestão, etc.). `errors.As()` procura um **tipo específico** de formulário e extrai todas as informações:

```go
type ErroValidacao struct {
    Campo    string
    Mensagem string
}

func (e ErroValidacao) Error() string {
    return fmt.Sprintf("erro no campo %s: %s", e.Campo, e.Mensagem)
}

// Uso:
err := validarFormulario("", 0)
var errValidacao ErroValidacao
if errors.As(err, &errValidacao) {
    // Encontrou um ErroValidacao! Agora posso acessar os campos
    fmt.Printf("O campo '%s' tem problema: %s\n", 
        errValidacao.Campo, errValidacao.Mensagem)
}
```

**Analogia:** É como procurar um formulário de reclamação específico na pilha. Quando você encontra, pode ler todos os campos detalhados (Campo, Mensagem) que não estariam disponíveis em um erro simples.

**Diferença:**
- `errors.Is()`: "Existe este erro específico aqui?" (sim/não)
- `errors.As()`: "Existe um erro deste tipo? Se sim, me dê todas as informações dele"

---

## 7. Panic e Recover: A Analogia do Incêndio e do Extintor

### Panic: Quando Algo Está Realmente Queimando

**Panic** é como um **incêndio** - algo que não deveria acontecer e que precisa de ação imediata:

```go
panic("algo deu muito errado!")
```

**Analogia do dia a dia:**
- ❌ **NÃO é panic**: "Acabou o leite" (erro esperado, você pode comprar mais)
- ✅ **É panic**: "A casa está pegando fogo!" (situação irrecuperável, precisa parar tudo)

**Em programação:**
- ❌ **NÃO use panic**: Divisão por zero (erro esperado, retorne um error)
- ✅ **Use panic**: Acesso a índice inválido de array quando você tinha certeza que o índice era válido (bug no código)

### Recover: O Extintor de Incêndio

`recover()` é como um **extintor de incêndio** - ele apaga o fogo (panic) antes que queime tudo:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Incêndio apagado! Causa: %v\n", r)
            // Limpeza, logging, etc.
        }
    }()
    
    fmt.Println("Tudo normal...")
    causarPanic()  // "Fogo!"
    fmt.Println("Esta linha será executada se o extintor funcionar")
}
```

**Analogia:** Imagine que você tem um extintor automático em casa. Se começar um incêndio, o extintor apaga automaticamente, e você pode continuar vivendo na casa (mas precisa investigar o que causou o incêndio).

**Quando usar recover:**
- Em servidores web: Evitar que um panic derrube o servidor inteiro
- Em goroutines: Evitar que um panic em uma goroutine derrube todo o programa
- Em bibliotecas: Garantir cleanup mesmo em caso de panic

**Regra de ouro:** Se você pode tratar com `error`, **use error**. Panic é apenas para emergências reais (bugs no código).

---

## 8. Stack Traces: A Analogia do Rastreamento de Encomenda

Quando você envia uma encomenda e ela se perde, a empresa de correios te mostra **todo o caminho** que a encomenda percorreu:

```
Encomenda #12345:
- 10:00 - Coletada em São Paulo
- 12:00 - Chegou no centro de distribuição de Campinas
- 14:00 - Saiu para entrega em Ribeirão Preto
- 16:00 - [PERDIDA AQUI]
```

Um **stack trace** em Go funciona da mesma forma. Quando um panic ocorre, você vê **todo o caminho** que o código percorreu até chegar no erro:

```go
func main() {
    nivel1()  // linha 3
}

func nivel1() {
    nivel2()  // linha 7
}

func nivel2() {
    nivel3()  // linha 11
}

func nivel3() {
    panic("erro aqui!")  // linha 15
}
```

**Stack trace:**
```
panic: erro aqui!

goroutine 1 [running]:
main.nivel3()           ← Onde o erro aconteceu (linha 15)
    /path/to/file.go:15
main.nivel2()           ← Quem chamou nivel3 (linha 11)
    /path/to/file.go:11
main.nivel1()           ← Quem chamou nivel2 (linha 7)
    /path/to/file.go:7
main.main()             ← Quem chamou nivel1 (linha 3)
    /path/to/file.go:3
```

**Analogia:** É como um rastreamento completo mostrando exatamente onde a encomenda (erro) foi perdida e como ela chegou até lá.

---

## 9. O Padrão Completo: A Analogia do Processo de Contratação

Vamos juntar tudo em uma analogia completa. Imagine que você está contratando um funcionário:

### Passo 1: Verificar se há Erro (Receber a Resposta)

```go
candidato, err := buscarCandidato("João")
if err != nil {
    // Há um problema! Preciso lidar com isso
    return err
}
// Tudo certo, posso continuar
```

**Analogia:** Você pediu informações sobre um candidato. Se receber uma resposta dizendo "candidato não encontrado", você precisa lidar com isso antes de continuar.

### Passo 2: Adicionar Contexto (Error Wrapping)

```go
func contratarFuncionario(nome string) error {
    candidato, err := buscarCandidato(nome)
    if err != nil {
        return fmt.Errorf("falha ao contratar %s: %w", nome, err)
        // Adiciona contexto: "não foi possível contratar porque não encontrei o candidato"
    }
    // processo de contratação...
    return nil
}
```

**Analogia:** Se algo der errado, você não apenas diz "candidato não encontrado". Você diz "não foi possível contratar João porque o candidato não foi encontrado no banco de dados".

### Passo 3: Tratar Erros Específicos (Sentinel Errors)

```go
var (
    ErrCandidatoNaoEncontrado = errors.New("candidato não encontrado")
    ErrCandidatoJaContratado  = errors.New("candidato já está contratado")
)

func contratarFuncionario(nome string) error {
    candidato, err := buscarCandidato(nome)
    if errors.Is(err, ErrCandidatoNaoEncontrado) {
        // Tratamento específico: criar novo candidato?
        return criarNovoCandidato(nome)
    } else if errors.Is(err, ErrCandidatoJaContratado) {
        // Tratamento específico: apenas atualizar?
        return atualizarCandidato(nome)
    } else if err != nil {
        // Outro tipo de erro
        return err
    }
    // contratação normal...
    return nil
}
```

**Analogia:** Diferentes problemas têm diferentes soluções:
- Candidato não encontrado → Criar novo registro
- Candidato já contratado → Apenas atualizar informações
- Outro erro → Retornar o erro para quem chamou

---

## Resumo com Analogias

1. **Erros são valores**: Como receber uma nota junto com seu pedido no restaurante
2. **Interface error**: Como um formulário padrão que qualquer um pode preencher
3. **errors.New()**: Nota simples
4. **fmt.Errorf()**: Nota detalhada com informações
5. **Error wrapping**: Carta em cadeia onde cada pessoa adiciona contexto
6. **Sentinel errors**: Códigos de erro específicos do banco
7. **errors.Is()**: Procurar uma chave específica em caixas dentro de caixas
8. **errors.As()**: Procurar um tipo específico de formulário e extrair informações
9. **Panic**: Incêndio - situação irrecuperável
10. **Recover**: Extintor - apaga o incêndio antes que queime tudo
11. **Stack trace**: Rastreamento de encomenda mostrando todo o caminho

---

## Próximos Passos

Agora que você entendeu os conceitos de forma simplificada, na próxima aula vamos colocar a mão na massa com exercícios práticos! Você vai criar funções que tratam erros, usar error wrapping, criar sentinel errors, e muito mais.

Até lá, tente pensar em situações do seu dia a dia que se parecem com cada um desses conceitos. Isso vai ajudar a fixar o aprendizado!



