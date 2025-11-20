# Aula 3 - Exercícios e Reflexão: Error Handling em Go

Olá! Agora é hora de colocar a mão na massa e praticar tudo que aprendemos sobre tratamento de erros em Go. Vamos começar com exercícios práticos e depois refletir sobre os conceitos.

---

## Exercício 1: Criando e Verificando Erros Básicos

### Objetivo
Criar uma função que valida um email e retorna erros apropriados.

### Tarefa
Crie uma função `ValidarEmail(email string) error` que:
1. Retorna um erro se o email estiver vazio
2. Retorna um erro se o email não contiver "@"
3. Retorna um erro se o email não contiver "." após o "@"
4. Retorna `nil` se o email for válido

Use `errors.New()` ou `fmt.Errorf()` conforme apropriado.

### Exemplo de Uso
```go
func main() {
    emails := []string{
        "",
        "email-sem-arroba",
        "email@sem-ponto",
        "email@valido.com",
    }
    
    for _, email := range emails {
        err := ValidarEmail(email)
        if err != nil {
            fmt.Printf("Email '%s' inválido: %v\n", email, err)
        } else {
            fmt.Printf("Email '%s' é válido!\n", email)
        }
    }
}
```

### Dica
Lembre-se de verificar os erros na ordem: primeiro vazio, depois "@", depois ".".

---

## Exercício 2: Error Wrapping e Contexto

### Objetivo
Praticar adicionar contexto aos erros usando error wrapping.

### Tarefa
Crie um sistema de processamento de arquivos com três níveis:
1. `lerArquivo(nome string) ([]byte, error)` - Lê o arquivo do disco
2. `processarConteudo(dados []byte) (string, error)` - Processa o conteúdo
3. `salvarResultado(nome string, resultado string) error` - Salva o resultado

Cada função deve:
- Chamar a função do nível inferior
- Se houver erro, adicionar contexto usando `fmt.Errorf()` com `%w`
- Retornar o erro com contexto

### Exemplo de Implementação Parcial
```go
func lerArquivo(nome string) ([]byte, error) {
    dados, err := os.ReadFile(nome)
    if err != nil {
        return nil, fmt.Errorf("erro ao ler arquivo %s: %w", nome, err)
    }
    return dados, nil
}

// Complete as outras funções seguindo o mesmo padrão
```

### Teste
Crie um arquivo de teste que chama `salvarResultado()` com um arquivo que não existe e veja a cadeia completa de erros.

---

## Exercício 3: Sentinel Errors e Tratamento Diferenciado

### Objetivo
Criar e usar sentinel errors para tratamento diferenciado de erros.

### Tarefa
Crie um sistema de autenticação com os seguintes sentinel errors:
- `ErrUsuarioNaoEncontrado`
- `ErrSenhaIncorreta`
- `ErrContaBloqueada`
- `ErrTentativasExcedidas`

Implemente a função `Autenticar(usuario, senha string) error` que:
1. Verifica se o usuário existe (simule com um mapa)
2. Verifica se a senha está correta
3. Verifica se a conta está bloqueada
4. Retorna o sentinel error apropriado

### Estrutura Sugerida
```go
var (
    ErrUsuarioNaoEncontrado = errors.New("auth: usuário não encontrado")
    ErrSenhaIncorreta       = errors.New("auth: senha incorreta")
    ErrContaBloqueada       = errors.New("auth: conta bloqueada")
    ErrTentativasExcedidas  = errors.New("auth: tentativas excedidas")
)

type Usuario struct {
    Nome     string
    Senha    string
    Bloqueado bool
}

var usuarios = map[string]Usuario{
    "admin": {Nome: "admin", Senha: "12345", Bloqueado: false},
    "user":  {Nome: "user", Senha: "senha", Bloqueado: true},
}

func Autenticar(usuario, senha string) error {
    // Implemente aqui
}
```

### Teste
Teste cada cenário e trate cada erro de forma diferente:
```go
err := Autenticar("admin", "senha_errada")
if errors.Is(err, ErrSenhaIncorreta) {
    fmt.Println("Senha incorreta - tente novamente")
} else if errors.Is(err, ErrUsuarioNaoEncontrado) {
    fmt.Println("Usuário não existe - criar conta?")
}
// etc...
```

---

## Exercício 4: Erros Customizados com `errors.As()`

### Objetivo
Criar tipos de erro customizados e usar `errors.As()` para extrair informações.

### Tarefa
Crie um sistema de validação de formulário com um tipo de erro customizado:

```go
type ErroValidacao struct {
    Campo     string
    Mensagem  string
    ValorRecebido interface{}
}

func (e ErroValidacao) Error() string {
    return fmt.Sprintf("campo '%s': %s (valor recebido: %v)", 
        e.Campo, e.Mensagem, e.ValorRecebido)
}
```

Crie a função `ValidarFormulario(nome string, idade int, email string) error` que:
1. Valida o nome (não pode ser vazio, mínimo 3 caracteres)
2. Valida a idade (deve ser entre 18 e 120)
3. Valida o email (deve conter "@")
4. Retorna um `ErroValidacao` para cada campo inválido

### Uso com `errors.As()`
```go
err := ValidarFormulario("", 15, "email-invalido")
var errValidacao ErroValidacao
if errors.As(err, &errValidacao) {
    fmt.Printf("Problema no campo '%s': %s\n", 
        errValidacao.Campo, errValidacao.Mensagem)
    fmt.Printf("Valor recebido: %v\n", errValidacao.ValorRecebido)
}
```

### Desafio Extra
Modifique a função para retornar **múltiplos erros** (uma slice de erros) quando houver mais de um campo inválido.

---

## Exercício 5: Panic e Recover em Situação Real

### Objetivo
Implementar recover em um handler HTTP para evitar que panics derrubem o servidor.

### Tarefa
Crie um servidor HTTP simples com um handler que pode causar panic:

```go
func handlerPerigoso(w http.ResponseWriter, r *http.Request) {
    // Simule uma operação que pode causar panic
    valores := r.URL.Query()["numero"]
    if len(valores) == 0 {
        panic("nenhum número fornecido")
    }
    
    numero, err := strconv.Atoi(valores[0])
    if err != nil {
        panic("número inválido")
    }
    
    resultado := 100 / numero  // Pode causar panic se numero == 0
    fmt.Fprintf(w, "Resultado: %d", resultado)
}
```

**Sua tarefa:** Modifique o handler para usar `recover()` e retornar uma resposta HTTP apropriada (status 500) em caso de panic, sem derrubar o servidor.

### Estrutura Sugerida
```go
func handlerSeguro(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            // Implemente: retornar HTTP 500 e logar o erro
        }
    }()
    
    // código do handler...
}
```

### Teste
Teste com diferentes URLs:
- `http://localhost:8080/?numero=10` (deve funcionar)
- `http://localhost:8080/?numero=0` (deve retornar 500 sem derrubar)
- `http://localhost:8080/` (deve retornar 500 sem derrubar)

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Erros São Valores?

**Pergunta:** Por que Go escolheu fazer erros serem valores explícitos em vez de usar exceções como outras linguagens? Quais são as vantagens e desvantagens dessa abordagem?

**Pense sobre:**
- Como isso afeta a legibilidade do código?
- Como isso afeta a performance?
- Como isso afeta o controle que você tem sobre o tratamento de erros?
- Em que situações você preferiria exceções? Em que situações prefere erros explícitos?

**Dica para reflexão:** Compare com uma linguagem que você conhece que usa exceções (Java, Python, C#). Qual abordagem você acha mais clara? Por quê?

---

### Reflexão 2: Quando Usar Error Wrapping?

**Pergunta:** Em uma aplicação grande com muitas camadas (API → Service → Repository → Database), quando você deve adicionar contexto aos erros e quando deve apenas repassar o erro original?

**Pense sobre:**
- Se você adicionar contexto em TODAS as camadas, o erro final pode ficar muito longo. Isso é um problema?
- Se você NÃO adicionar contexto em nenhuma camada, como você vai debugar quando algo der errado?
- Qual é o equilíbrio ideal?
- Em que ponto da cadeia você deve parar de adicionar contexto?

**Cenário para pensar:**
```
Database retorna: "connection timeout"
Repository adiciona: "erro ao buscar usuário: connection timeout"
Service adiciona: "erro ao autenticar: erro ao buscar usuário: connection timeout"
API adiciona: "erro na requisição: erro ao autenticar: erro ao buscar usuário: connection timeout"
```

Isso é útil ou excessivo? Por quê?

---

### Reflexão 3: Sentinel Errors vs Erros Customizados

**Pergunta:** Quando você deve usar sentinel errors (erros pré-definidos) e quando deve usar tipos de erro customizados? Qual abordagem é melhor para APIs públicas?

**Pense sobre:**
- Sentinel errors são mais simples, mas limitados. Quando isso é suficiente?
- Erros customizados são mais flexíveis, mas mais complexos. Quando vale a pena?
- Se você está criando uma biblioteca que outros vão usar, qual abordagem facilita mais o uso?
- Como você documenta cada tipo de erro para que outros desenvolvedores saibam como tratá-los?

**Cenário:** Você está criando uma biblioteca de autenticação. Você usaria sentinel errors ou erros customizados? Por quê?

---

### Reflexão 4: Panic vs Error - A Linha Tênue

**Pergunta:** A regra geral é "use error para erros esperados, panic para bugs". Mas como você determina se algo é um "erro esperado" ou um "bug"?

**Pense sobre:**
- Divisão por zero: é um erro esperado ou um bug? Depende do contexto?
- Acesso a índice inválido de array: sempre é um bug, ou pode ser um erro esperado?
- Falha ao conectar no banco de dados: erro esperado ou bug?
- O que acontece se você usar panic demais? E se usar error demais?

**Cenário 1:** Você tem uma função que recebe um ID e busca um usuário. Se o usuário não existir, é um erro esperado (retornar error) ou um bug (panic)?

**Cenário 2:** Você tem uma função que recebe um índice e acessa um elemento de um slice. Você garantiu que o índice sempre será válido através de validação prévia. Se o índice for inválido, é um bug (panic)?

Como você decide?

---

## Dicas para os Exercícios

1. **Comece simples**: Faça o Exercício 1 primeiro para se acostumar com o padrão básico
2. **Teste cada função**: Crie funções `main()` de teste para cada exercício
3. **Leia as mensagens de erro**: Quando algo der errado, leia a mensagem completa - ela te diz exatamente o que está errado
4. **Use o compilador a seu favor**: O compilador do Go vai te avisar se você não estiver tratando erros corretamente
5. **Experimente**: Tente quebrar seu código propositalmente para ver como os erros se comportam

---

## Entrega

Para cada exercício:
1. Escreva o código completo
2. Teste com diferentes cenários (casos válidos e inválidos)
3. Para as reflexões, escreva suas respostas (não precisa ser código, pode ser texto mesmo)

**Lembre-se:** O objetivo não é apenas fazer funcionar, mas entender **por que** você está fazendo dessa forma. Se você não souber responder "por quê?", volte e releia as aulas anteriores!

Boa sorte e bons estudos! 🚀




