# Aula 4 - Performance e Boas Práticas: Error Handling em Go

Olá! Agora que você entende os conceitos de Error Handling em Go, é crucial aprender **quando e como** tratá-los de forma eficiente e seguir as melhores práticas da comunidade Go. Nesta aula, vamos explorar aspectos de performance, boas práticas, padrões comuns, e os erros que você deve evitar.

---

## 🚀 Performance: O Custo dos Erros em Go

### Erros São Valores: Zero Overhead

**Ponto crucial:** Em Go, erros são valores normais. Isso significa:

✅ **Sem overhead de exceções** - Não há custo de stack unwinding como em linguagens com exceções
✅ **Sem try/catch** - O código de verificação de erro é simples e rápido
✅ **Compilador otimiza** - Verificações `if err != nil` são extremamente eficientes

### Comparação de Performance

**Linguagens com Exceções (Java, Python, C#):**
- Quando uma exceção é lançada, o sistema precisa:
  1. Desenrolar a stack (unwind)
  2. Procurar o handler apropriado
  3. Executar código de cleanup
  4. Transferir controle
- **Custo:** Alto, especialmente quando exceções são comuns

**Go (Erros como Valores):**
- Quando um erro é retornado:
  1. Retorna um valor (apenas uma variável)
  2. Verificação simples `if err != nil`
  3. Continuação normal do fluxo
- **Custo:** Praticamente zero

### Benchmark: Erro vs Exceção

```go
// Go: Retornar erro
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}

// Em linguagens com exceções, o equivalente seria:
// throw new ArithmeticException("divisão por zero")
```

**Resultado:** Em Go, retornar um erro é tão rápido quanto retornar qualquer outro valor. Não há penalidade de performance.

### Quando Performance Importa

⚠️ **Atenção:** Em loops muito apertados (hot paths), verificar erros repetidamente pode ter um pequeno impacto. Mas:

1. **É mínimo** - A verificação `if err != nil` é uma operação extremamente rápida
2. **Vale a pena** - A clareza e segurança valem muito mais que micro-otimizações
3. **Não otimize prematuramente** - Só se preocupe com isso se você realmente identificar um gargalo

---

## ✅ Boas Práticas: Padrões e Convenções

### ✅ SEMPRE: Verifique Erros Imediatamente

```go
// ❌ ERRADO: Ignorar erro
resultado, _ := dividir(10, 0)
fmt.Println(resultado) // Pode usar valor inválido!

// ✅ CORRETO: Verificar imediatamente
resultado, err := dividir(10, 0)
if err != nil {
    return err // ou tratar de outra forma
}
fmt.Println(resultado)
```

**Por quê?**
- Ignorar erros é a causa #1 de bugs em produção
- O compilador Go não te força a verificar, mas você DEVE fazer isso
- Use ferramentas como `errcheck` ou `staticcheck` para detectar erros ignorados

### ✅ SEMPRE: Adicione Contexto com Error Wrapping

```go
// ❌ ERRADO: Perder contexto
func processarDados(nome string) error {
    return lerArquivo(nome) // Erro sem contexto
}

// ✅ CORRETO: Adicionar contexto
func processarDados(nome string) error {
    dados, err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("falha ao processar dados de %s: %w", nome, err)
    }
    // processar dados...
    return nil
}
```

**Por quê?**
- Facilita debugging em produção
- Mostra a cadeia completa de onde o erro ocorreu
- Ajuda a identificar problemas rapidamente

### ✅ USE: Sentinel Errors para APIs Públicas

```go
// ✅ BOM: API clara e previsível
package auth

var (
    ErrUsuarioNaoEncontrado = errors.New("auth: usuário não encontrado")
    ErrSenhaIncorreta       = errors.New("auth: senha incorreta")
)

func Autenticar(usuario, senha string) error {
    // implementação...
}
```

**Por quê?**
- Chamadores podem tratar erros específicos
- API fica auto-documentada
- Facilita testes e mocks

### ✅ USE: Erros Customizados para Informações Ricas

```go
// ✅ BOM: Quando você precisa de informações adicionais
type ErroValidacao struct {
    Campo    string
    Mensagem string
    Valor    interface{}
}

func (e ErroValidacao) Error() string {
    return fmt.Sprintf("campo '%s': %s", e.Campo, e.Mensagem)
}
```

**Quando usar:**
- Quando o chamador precisa de informações específicas do erro
- Quando você precisa de múltiplos campos de informação
- Quando diferentes tipos de erro precisam de tratamento diferente

---

## ❌ Evite: Erros Comuns e Anti-padrões

### ❌ NÃO: Ignore Erros com `_`

```go
// ❌ MUITO ERRADO
arquivo, _ := os.Open("dados.txt")
defer arquivo.Close() // Pode causar panic se arquivo for nil!
```

**Solução:**
```go
// ✅ CORRETO
arquivo, err := os.Open("dados.txt")
if err != nil {
    return fmt.Errorf("erro ao abrir arquivo: %w", err)
}
defer arquivo.Close()
```

### ❌ NÃO: Use Panic para Erros Esperados

```go
// ❌ ERRADO: Divisão por zero é um erro esperado
func dividir(a, b int) int {
    if b == 0 {
        panic("divisão por zero")
    }
    return a / b
}

// ✅ CORRETO: Retornar error
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

**Regra de ouro:** Se o erro pode acontecer em condições normais de uso, use `error`. Panic apenas para bugs no código.

### ❌ NÃO: Crie Erros Sem Prefijo de Pacote

```go
// ❌ ERRADO: Pode colidir com outros pacotes
var ErrNaoEncontrado = errors.New("não encontrado")

// ✅ CORRETO: Prefixo do pacote
var ErrNaoEncontrado = errors.New("auth: usuário não encontrado")
```

**Por quê?**
- Evita colisões de nomes entre pacotes
- Facilita identificar de onde o erro veio
- Melhora mensagens de erro em logs

### ❌ NÃO: Wrap Erros Sem Adicionar Contexto Útil

```go
// ❌ ERRADO: Contexto redundante
func processar(nome string) error {
    err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("erro: %w", err) // Não adiciona informação útil
    }
    return nil
}

// ✅ CORRETO: Adiciona contexto útil
func processar(nome string) error {
    err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("falha ao processar arquivo %s: %w", nome, err)
    }
    return nil
}
```

### ❌ NÃO: Retorne Erros Genéricos Demais

```go
// ❌ ERRADO: Muito genérico
func buscarUsuario(id int) (*Usuario, error) {
    if id < 0 {
        return nil, errors.New("erro") // Muito vago!
    }
    // ...
}

// ✅ CORRETO: Específico e útil
func buscarUsuario(id int) (*Usuario, error) {
    if id < 0 {
        return nil, fmt.Errorf("ID inválido: %d (deve ser positivo)", id)
    }
    // ...
}
```

---

## 🎯 Padrões Avançados e Melhores Práticas

### Padrão 1: Error Wrapping Estruturado

Em aplicações grandes, estabeleça uma convenção de como adicionar contexto:

```go
// Convenção: [ação] [recurso] [detalhes]: [erro original]
func processarArquivo(nome string) error {
    dados, err := lerArquivo(nome)
    if err != nil {
        return fmt.Errorf("processar arquivo %s: %w", nome, err)
    }
    
    resultado, err := parsearDados(dados)
    if err != nil {
        return fmt.Errorf("processar arquivo %s: parsear dados: %w", nome, err)
    }
    
    err = salvarResultado(resultado)
    if err != nil {
        return fmt.Errorf("processar arquivo %s: salvar resultado: %w", nome, err)
    }
    
    return nil
}
```

**Benefícios:**
- Mensagens consistentes
- Fácil de ler em logs
- Contexto claro em cada nível

### Padrão 2: Error Types para Categorização

Crie tipos de erro para categorizar diferentes classes de problemas:

```go
type ErroTemporario interface {
    error
    Temporario() bool
}

type ErroRede struct {
    Mensagem string
}

func (e ErroRede) Error() string {
    return e.Mensagem
}

func (e ErroRede) Temporario() bool {
    return true // Erros de rede são geralmente temporários
}

// Uso:
func fazerRequisicao() error {
    // ...
    return ErroRede{Mensagem: "timeout na conexão"}
}

// Tratamento:
err := fazerRequisicao()
var errTemp ErroTemporario
if errors.As(err, &errTemp) && errTemp.Temporario() {
    // Tentar novamente
    return tentarNovamente()
}
```

### Padrão 3: Logging Estruturado de Erros

Use logging estruturado para facilitar análise:

```go
import "log"

func processarDados(nome string) error {
    log.Printf("Iniciando processamento de %s", nome)
    
    dados, err := lerArquivo(nome)
    if err != nil {
        log.Printf("ERRO: falha ao ler arquivo %s: %v", nome, err)
        return fmt.Errorf("processar %s: %w", nome, err)
    }
    
    log.Printf("Processados %d bytes de %s", len(dados), nome)
    return nil
}
```

**Para produção, considere bibliotecas como:**
- `logrus` - Logging estruturado
- `zap` - Logging de alta performance
- `zerolog` - Logging zero-allocation

### Padrão 4: Retry com Backoff para Erros Temporários

```go
func tentarComRetry(operacao func() error, maxTentativas int) error {
    for i := 0; i < maxTentativas; i++ {
        err := operacao()
        if err == nil {
            return nil
        }
        
        var errTemp ErroTemporario
        if !errors.As(err, &errTemp) || !errTemp.Temporario() {
            return err // Erro não é temporário, não tente novamente
        }
        
        // Backoff exponencial
        tempoEspera := time.Duration(1<<uint(i)) * time.Second
        time.Sleep(tempoEspera)
    }
    return fmt.Errorf("falhou após %d tentativas", maxTentativas)
}
```

---

## 🔧 Ferramentas e Verificações

### 1. errcheck - Detectar Erros Ignorados

```bash
go install github.com/kisielk/errcheck@latest
errcheck ./...
```

**O que faz:** Detecta quando você ignora erros com `_`

### 2. staticcheck - Análise Estática

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

**O que faz:** Detecta muitos problemas comuns, incluindo erros não verificados

### 3. go vet - Verificador Oficial

```bash
go vet ./...
```

**O que faz:** Verifica problemas comuns, incluindo alguns relacionados a erros

### 4. Testes de Erro

Sempre teste os caminhos de erro:

```go
func TestDividir(t *testing.T) {
    tests := []struct {
        nome     string
        a        int
        b        int
        esperado int
        temErro  bool
    }{
        {"divisão normal", 10, 2, 5, false},
        {"divisão por zero", 10, 0, 0, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.nome, func(t *testing.T) {
            resultado, err := dividir(tt.a, tt.b)
            if tt.temErro {
                if err == nil {
                    t.Error("esperava erro, mas não houve")
                }
            } else {
                if err != nil {
                    t.Errorf("não esperava erro, mas obteve: %v", err)
                }
                if resultado != tt.esperado {
                    t.Errorf("esperava %d, obteve %d", tt.esperado, resultado)
                }
            }
        })
    }
}
```

---

## 📊 Performance: Quando Otimizar

### Caso 1: Hot Paths (Caminhos Críticos)

Se você tem uma função chamada milhões de vezes por segundo:

```go
// Se esta função é chamada milhões de vezes:
func processarItem(item Item) error {
    if err := validar(item); err != nil {
        return err // Verificação de erro é OK, é rápida
    }
    // ...
}
```

**Conclusão:** Verificações de erro são tão rápidas que não vale a pena otimizar. Foque em outras partes do código.

### Caso 2: Error Allocation

Criar novos erros aloca memória. Em loops muito apertados:

```go
// ❌ Se chamado milhões de vezes, aloca muitos erros
func processar(items []Item) error {
    for _, item := range items {
        if err := validar(item); err != nil {
            return fmt.Errorf("item inválido: %w", err) // Aloca memória
        }
    }
    return nil
}

// ✅ Para hot paths, considere reutilizar erros ou usar sentinel errors
var ErrItemInvalido = errors.New("item inválido")

func processar(items []Item) error {
    for _, item := range items {
        if !validar(item) {
            return ErrItemInvalido // Não aloca
        }
    }
    return nil
}
```

**Mas lembre-se:** Só otimize se você realmente medir um problema de performance!

---

## 🎓 Resumo: Checklist de Boas Práticas

### ✅ Sempre Faça:

- [ ] Verifique erros imediatamente após chamadas de função
- [ ] Adicione contexto com error wrapping (`%w` em `fmt.Errorf`)
- [ ] Use sentinel errors para APIs públicas
- [ ] Prefira prefixos de pacote em sentinel errors
- [ ] Documente quando cada erro é retornado
- [ ] Teste os caminhos de erro
- [ ] Use `errors.Is()` para verificar sentinel errors
- [ ] Use `errors.As()` para extrair tipos de erro customizados

### ❌ Nunca Faça:

- [ ] Ignorar erros com `_`
- [ ] Usar panic para erros esperados
- [ ] Criar erros sem contexto útil
- [ ] Retornar erros genéricos demais
- [ ] Esquecer de verificar erros em defer
- [ ] Misturar panic e error handling

### 🎯 Considere:

- [ ] Logging estruturado para produção
- [ ] Error types para categorização
- [ ] Retry com backoff para erros temporários
- [ ] Ferramentas de análise estática (errcheck, staticcheck)
- [ ] Convenções de error wrapping na equipe

---

## 🚀 Próximos Passos

Agora que você domina Error Handling em Go:

1. **Pratique:** Implemente os exercícios da aula anterior
2. **Leia código real:** Veja como bibliotecas populares tratam erros
3. **Estabeleça convenções:** Se trabalhar em equipe, defina padrões de error handling
4. **Use ferramentas:** Configure errcheck e staticcheck no seu CI/CD

**Lembre-se:** Error handling em Go é sobre clareza e controle. Não é sobre performance (que já é excelente), mas sobre escrever código robusto e fácil de debugar.

Bons estudos e código sem bugs! 🐛❌



