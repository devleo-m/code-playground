# Módulo 14: Error Handling em Go

## Aula 1: Error Handling - Tratamento de Erros em Go

Olá! Bem-vindo ao décimo quarto módulo. Até agora você aprendeu sobre funções, métodos, interfaces, generics e muitos outros conceitos fundamentais de Go. Mas você já se perguntou: **"O que acontece quando algo dá errado no meu programa? Como eu trato erros de forma adequada?"**

Em muitas linguagens de programação, erros são tratados através de **exceções** (exceptions) - mecanismos que interrompem o fluxo normal do programa e "lançam" erros que precisam ser "capturados" em algum lugar. Go, no entanto, adota uma filosofia completamente diferente: **erros são valores**.

Nesta aula, vamos mergulhar profundamente no sistema de tratamento de erros do Go: entender a filosofia por trás dele, como criar e verificar erros, como adicionar contexto através de error wrapping, como trabalhar com sentinel errors, e quando usar panic e recover.

---

## 1. A Filosofia do Error Handling em Go

### "Errors are Values" (Erros são Valores)

A filosofia central do Go é que **erros são valores normais**, não exceções especiais. Isso significa:

- ✅ **Explícito**: Você vê explicitamente quando uma função pode retornar um erro
- ✅ **Previsível**: Não há surpresas - erros são parte da assinatura da função
- ✅ **Controlado**: Você decide como tratar cada erro, não o sistema
- ✅ **Sem overhead**: Não há custo de performance como em sistemas de exceções

### Comparação com Outras Linguagens

**Linguagens com Exceções (Java, Python, C#):**
```java
// Java - exceções não são visíveis na assinatura
public int dividir(int a, int b) {
    if (b == 0) {
        throw new ArithmeticException("Divisão por zero");
    }
    return a / b;
}
// O chamador pode não saber que precisa tratar o erro!
```

**Go - Erros Explícitos:**
```go
// Go - erro é parte da assinatura
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
// O chamador SABE que precisa verificar o erro
```

### O Padrão `if err != nil`

Em Go, o padrão mais comum é verificar erros imediatamente após uma chamada de função:

```go
resultado, err := dividir(10, 2)
if err != nil {
    // Tratar o erro
    return err // ou log.Fatal(err), ou fazer algo específico
}
// Continuar com o resultado
```

Este padrão pode parecer verboso no início, mas oferece:
- **Clareza**: Fica explícito que você está tratando erros
- **Controle**: Você decide o que fazer com cada erro
- **Sem surpresas**: Não há erros "escondidos" que podem quebrar seu programa

---

## 2. A Interface `error`

### Definição

A interface `error` é uma das interfaces mais simples e importantes do Go:

```go
type error interface {
    Error() string
}
```

Qualquer tipo que implemente o método `Error() string` automaticamente implementa a interface `error`. Isso significa que você pode criar seus próprios tipos de erro!

### Tipos de Erro Comuns

#### 1. Erros Simples com `errors.New()`

```go
import "errors"

func validarEmail(email string) error {
    if email == "" {
        return errors.New("email não pode ser vazio")
    }
    if !strings.Contains(email, "@") {
        return errors.New("email deve conter @")
    }
    return nil
}
```

#### 2. Erros Formatados com `fmt.Errorf()`

```go
func validarIdade(idade int) error {
    if idade < 0 {
        return fmt.Errorf("idade inválida: %d (deve ser positiva)", idade)
    }
    if idade > 150 {
        return fmt.Errorf("idade inválida: %d (muito alta, máximo 150)", idade)
    }
    return nil
}
```

#### 3. Erros Customizados

Você pode criar seus próprios tipos de erro:

```go
type ErroValidacao struct {
    Campo    string
    Mensagem string
    Valor    interface{}
}

func (e ErroValidacao) Error() string {
    return fmt.Sprintf("erro no campo '%s': %s (valor: %v)", 
        e.Campo, e.Mensagem, e.Valor)
}

func validarFormulario(nome string, idade int) error {
    if nome == "" {
        return ErroValidacao{
            Campo:    "nome",
            Mensagem: "não pode ser vazio",
            Valor:    nome,
        }
    }
    if idade < 18 {
        return ErroValidacao{
            Campo:    "idade",
            Mensagem: "deve ser maior ou igual a 18",
            Valor:    idade,
        }
    }
    return nil
}
```

---

## 3. Criando Erros

### `errors.New()`

A forma mais simples de criar um erro:

```go
import "errors"

err := errors.New("algo deu errado")
```

**Quando usar:**
- Erros simples e estáticos
- Mensagens de erro que não precisam de formatação
- Erros pré-definidos (sentinel errors)

**Exemplo:**
```go
var ErrDivisaoPorZero = errors.New("divisão por zero não permitida")

func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivisaoPorZero
    }
    return a / b, nil
}
```

### `fmt.Errorf()`

Cria erros com formatação, similar ao `fmt.Printf()`:

```go
nome := "João"
idade := -5
err := fmt.Errorf("usuário %s tem idade inválida: %d", nome, idade)
```

**Verbos comuns:**
- `%v` - valor padrão
- `%s` - string
- `%d` - inteiro
- `%f` - float
- `%w` - **error wrapping** (veremos depois)

**Exemplo prático:**
```go
func processarArquivo(caminho string) error {
    arquivo, err := os.Open(caminho)
    if err != nil {
        return fmt.Errorf("falha ao abrir arquivo %s: %v", caminho, err)
    }
    defer arquivo.Close()
    
    // processamento...
    return nil
}
```

---

## 4. Error Wrapping (Encadeamento de Erros)

### O Problema: Perder Contexto

Imagine esta cadeia de chamadas:

```go
func main() {
    err := processarDados("arquivo.txt")
    if err != nil {
        log.Printf("Erro: %v", err)
        // Output: Erro: arquivo não encontrado
        // Mas onde? Em qual função? Com quais parâmetros?
    }
}

func processarDados(nome string) error {
    return lerArquivo(nome) // retorna erro sem contexto
}

func lerArquivo(nome string) error {
    return os.Open(nome) // retorna: "arquivo não encontrado"
}
```

O erro original se perde! Não sabemos em qual camada o erro ocorreu.

### A Solução: Error Wrapping

Go 1.13 introduziu o verbo `%w` no `fmt.Errorf()` para **encadear erros**, preservando o erro original enquanto adiciona contexto:

```go
func processarDados(nome string) error {
    err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("falha ao processar dados do arquivo %s: %w", nome, err)
    }
    return nil
}

func lerArquivo(nome string) error {
    arquivo, err := os.Open(nome)
    if err != nil {
        return fmt.Errorf("erro ao abrir arquivo %s: %w", nome, err)
    }
    defer arquivo.Close()
    return nil
}
```

Agora, quando ocorrer um erro:
```
Erro: falha ao processar dados do arquivo arquivo.txt: erro ao abrir arquivo arquivo.txt: open arquivo.txt: no such file or directory
```

Temos o **contexto completo** da cadeia de erros!

### Desempacotando Erros: `errors.Unwrap()`

Para acessar o erro original em uma cadeia:

```go
err := processarDados("arquivo.txt")
if err != nil {
    fmt.Printf("Erro atual: %v\n", err)
    fmt.Printf("Erro original: %v\n", errors.Unwrap(err))
}
```

### Verificando Erros Específicos: `errors.Is()`

`errors.Is()` verifica se um erro específico está na cadeia de erros:

```go
var ErrArquivoNaoEncontrado = errors.New("arquivo não encontrado")

func lerArquivo(nome string) error {
    arquivo, err := os.Open(nome)
    if err != nil {
        if os.IsNotExist(err) {
            return ErrArquivoNaoEncontrado
        }
        return err
    }
    defer arquivo.Close()
    return nil
}

// Em outro lugar:
err := processarDados("arquivo.txt")
if errors.Is(err, ErrArquivoNaoEncontrado) {
    fmt.Println("Arquivo não existe - criar novo?")
}
```

**Importante:** `errors.Is()` funciona mesmo com error wrapping! Ela percorre toda a cadeia de erros.

### Verificando Tipos de Erro: `errors.As()`

`errors.As()` verifica se um erro é de um tipo específico e extrai o erro:

```go
type ErroValidacao struct {
    Campo    string
    Mensagem string
}

func (e ErroValidacao) Error() string {
    return fmt.Sprintf("erro no campo %s: %s", e.Campo, e.Mensagem)
}

func validarFormulario(nome string, idade int) error {
    if nome == "" {
        return ErroValidacao{Campo: "nome", Mensagem: "não pode ser vazio"}
    }
    if idade < 18 {
        return ErroValidacao{Campo: "idade", Mensagem: "deve ser maior que 18"}
    }
    return nil
}

// Uso:
err := validarFormulario("", 0)
var errValidacao ErroValidacao
if errors.As(err, &errValidacao) {
    fmt.Printf("Campo com erro: %s\n", errValidacao.Campo)
    fmt.Printf("Mensagem: %s\n", errValidacao.Mensagem)
}
```

**Diferença entre `errors.Is()` e `errors.As()`:**
- `errors.Is()`: Verifica se um **valor de erro específico** está na cadeia
- `errors.As()`: Verifica se um **tipo de erro específico** está na cadeia e extrai o erro

---

## 5. Sentinel Errors (Erros Sentinelas)

### O Que São?

**Sentinel errors** são erros pré-definidos que representam condições específicas e conhecidas. Eles são tipicamente definidos como variáveis de pacote:

```go
var (
    ErrUsuarioNaoEncontrado = errors.New("usuário não encontrado")
    ErrSenhaInvalida        = errors.New("senha inválida")
    ErrSessaoExpirada       = errors.New("sessão expirada")
)
```

### Por Que Usar?

1. **APIs Previsíveis**: Chamadores podem verificar erros específicos
2. **Tratamento Diferenciado**: Diferentes erros podem ter diferentes tratamentos
3. **Documentação**: Erros sentinelas servem como documentação da API

### Exemplo Prático

```go
package auth

var (
    ErrUsuarioNaoEncontrado = errors.New("auth: usuário não encontrado")
    ErrSenhaInvalida        = errors.New("auth: senha inválida")
    ErrContaBloqueada       = errors.New("auth: conta bloqueada")
)

func Autenticar(usuario, senha string) error {
    user, err := buscarUsuario(usuario)
    if err != nil {
        return ErrUsuarioNaoEncontrado
    }
    
    if user.Bloqueado {
        return ErrContaBloqueada
    }
    
    if !verificarSenha(user, senha) {
        return ErrSenhaInvalida
    }
    
    return nil
}

// Uso:
err := auth.Autenticar("joao", "senha123")
if errors.Is(err, auth.ErrUsuarioNaoEncontrado) {
    fmt.Println("Usuário não existe - criar conta?")
} else if errors.Is(err, auth.ErrSenhaInvalida) {
    fmt.Println("Senha incorreta - tentar novamente?")
} else if errors.Is(err, auth.ErrContaBloqueada) {
    fmt.Println("Conta bloqueada - contatar suporte")
} else if err != nil {
    fmt.Printf("Erro inesperado: %v\n", err)
}
```

### Sentinel Errors na Biblioteca Padrão

A biblioteca padrão do Go usa muitos sentinel errors:

```go
import (
    "io"
    "os"
)

// io.EOF - fim do arquivo
// os.ErrNotExist - arquivo não existe
// os.ErrPermission - permissão negada

arquivo, err := os.Open("dados.txt")
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("Arquivo não existe")
} else if errors.Is(err, os.ErrPermission) {
    fmt.Println("Sem permissão para ler o arquivo")
}
```

### Boas Práticas para Sentinel Errors

1. **Prefixo do Pacote**: Use prefixo para evitar colisões
   ```go
   var ErrUsuarioNaoEncontrado = errors.New("auth: usuário não encontrado")
   ```

2. **Documentação**: Documente quando cada erro é retornado
   ```go
   // ErrUsuarioNaoEncontrado é retornado quando o usuário não existe no sistema
   var ErrUsuarioNaoEncontrado = errors.New("auth: usuário não encontrado")
   ```

3. **Não Abuse**: Use apenas para erros que realmente precisam de tratamento específico

---

## 6. Panic e Recover

### Quando Usar Panic?

**Panic** deve ser usado **muito raramente** e apenas para situações verdadeiramente **irrecuperáveis**:

- ✅ Erros de programação (bugs) que indicam um problema no código
- ✅ Situações onde continuar a execução não faz sentido
- ❌ **NÃO** use para erros esperados (use `error`!)
- ❌ **NÃO** use como mecanismo de controle de fluxo

### Exemplo de Uso Correto de Panic

```go
func dividir(a, b int) int {
    if b == 0 {
        panic("divisão por zero: isso não deveria acontecer!")
        // Se chegou aqui, há um bug no código
    }
    return a / b
}
```

**Mas na prática, é melhor usar error:**
```go
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

### O Que Acontece com Panic?

Quando `panic()` é chamado:
1. A execução do programa **para imediatamente**
2. Todas as funções **defer** são executadas
3. O programa **imprime uma stack trace** e termina

```go
func main() {
    fmt.Println("Início")
    causarPanic()
    fmt.Println("Esta linha nunca será executada")
}

func causarPanic() {
    panic("algo deu muito errado!")
}
```

**Output:**
```
Início
panic: algo deu muito errado!

goroutine 1 [running]:
main.causarPanic(...)
    /path/to/file.go:10
main.main()
    /path/to/file.go:6
exit code 2
```

### Recover: Capturando Panics

`recover()` só funciona dentro de uma função **defer** e captura panics:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic recuperado: %v\n", r)
            // Aqui você pode fazer cleanup, logging, etc.
        }
    }()
    
    fmt.Println("Início")
    causarPanic()
    fmt.Println("Esta linha será executada se o panic for recuperado")
}

func causarPanic() {
    panic("algo deu muito errado!")
}
```

**Output:**
```
Início
Panic recuperado: algo deu muito errado!
Esta linha será executada se o panic for recuperado
```

### Quando Usar Recover?

**Casos válidos:**
1. **Servidores HTTP**: Evitar que um panic derrube o servidor inteiro
   ```go
   func handler(w http.ResponseWriter, r *http.Request) {
       defer func() {
           if r := recover(); r != nil {
               http.Error(w, "Erro interno", http.StatusInternalServerError)
               log.Printf("Panic recuperado: %v", r)
           }
       }()
       // código do handler
   }
   ```

2. **Goroutines**: Evitar que um panic em uma goroutine derrube o programa
   ```go
   go func() {
       defer func() {
           if r := recover(); r != nil {
               log.Printf("Panic na goroutine: %v", r)
           }
       }()
       // código da goroutine
   }()
   ```

3. **Cleanup em Bibliotecas**: Garantir que recursos sejam liberados mesmo em caso de panic

**Regra de ouro:** Se você pode tratar o erro com `error`, **use error**. Panic é apenas para situações verdadeiramente irrecuperáveis.

---

## 7. Stack Traces e Debugging

### Stack Traces Automáticos

Quando um panic ocorre, Go automaticamente imprime uma **stack trace** mostrando:
- A cadeia completa de chamadas de função
- O arquivo e linha onde cada função foi chamada
- Os valores dos argumentos (em alguns casos)

```go
func main() {
    nivel1()
}

func nivel1() {
    nivel2()
}

func nivel2() {
    nivel3()
}

func nivel3() {
    panic("erro no nível 3")
}
```

**Stack trace:**
```
panic: erro no nível 3

goroutine 1 [running]:
main.nivel3()
    /path/to/file.go:15
main.nivel2()
    /path/to/file.go:11
main.nivel1()
    /path/to/file.go:7
main.main()
    /path/to/file.go:3
```

### Ferramentas de Debugging

#### 1. Delve (Debugger)

Delve é o debugger oficial do Go:

```bash
# Instalar
go install github.com/go-delve/delve/cmd/dlv@latest

# Usar
dlv debug main.go
```

**Comandos úteis:**
- `break` - definir breakpoint
- `continue` - continuar execução
- `next` - próxima linha
- `print <var>` - imprimir variável
- `stack` - mostrar stack trace

#### 2. pprof (Profiling)

Para análise de performance e identificação de problemas:

```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // seu código
}
```

Acesse `http://localhost:6060/debug/pprof/` para ver perfis.

#### 3. Race Detector

Detecta condições de corrida (race conditions):

```bash
go run -race main.go
go test -race ./...
```

### Logging para Debugging

Use logging estruturado para facilitar debugging:

```go
import "log"

func processarDados(nome string) error {
    log.Printf("Iniciando processamento de %s", nome)
    
    dados, err := lerArquivo(nome)
    if err != nil {
        log.Printf("Erro ao ler arquivo %s: %v", nome, err)
        return fmt.Errorf("falha ao processar: %w", err)
    }
    
    log.Printf("Processados %d bytes de %s", len(dados), nome)
    return nil
}
```

---

## 8. Padrões e Boas Práticas

### 1. Sempre Verifique Erros

```go
// ❌ ERRADO
resultado, _ := dividir(10, 0) // ignorando erro!

// ✅ CORRETO
resultado, err := dividir(10, 0)
if err != nil {
    return err
}
```

### 2. Adicione Contexto com Error Wrapping

```go
// ❌ ERRADO - perde contexto
func processar(nome string) error {
    return lerArquivo(nome) // erro sem contexto
}

// ✅ CORRETO - adiciona contexto
func processar(nome string) error {
    err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("falha ao processar %s: %w", nome, err)
    }
    return nil
}
```

### 3. Use Sentinel Errors para APIs Públicas

```go
// ✅ BOM - API clara
var ErrUsuarioNaoEncontrado = errors.New("usuário não encontrado")

func BuscarUsuario(id int) (*Usuario, error) {
    // ...
    return nil, ErrUsuarioNaoEncontrado
}
```

### 4. Documente Quando Erros São Retornados

```go
// BuscarUsuario busca um usuário pelo ID.
// Retorna ErrUsuarioNaoEncontrado se o usuário não existir.
func BuscarUsuario(id int) (*Usuario, error) {
    // ...
}
```

### 5. Não Use Panic para Erros Esperados

```go
// ❌ ERRADO
func dividir(a, b int) int {
    if b == 0 {
        panic("divisão por zero")
    }
    return a / b
}

// ✅ CORRETO
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

---

## Resumo da Aula

Nesta aula, você aprendeu:

1. **Filosofia do Go**: Erros são valores, não exceções
2. **Interface error**: Qualquer tipo com `Error() string` é um erro
3. **Criar erros**: `errors.New()` e `fmt.Errorf()`
4. **Error wrapping**: Adicionar contexto com `%w` e `fmt.Errorf()`
5. **Desempacotar erros**: `errors.Unwrap()`, `errors.Is()`, `errors.As()`
6. **Sentinel errors**: Erros pré-definidos para APIs previsíveis
7. **Panic e recover**: Para situações verdadeiramente irrecuperáveis
8. **Debugging**: Stack traces, Delve, pprof, race detector

**Próximos passos:**
- Na próxima aula, vamos simplificar esses conceitos com analogias e exemplos do cotidiano
- Depois, exercícios práticos para fixar o aprendizado
- Por fim, boas práticas e otimizações de performance

Até a próxima!




