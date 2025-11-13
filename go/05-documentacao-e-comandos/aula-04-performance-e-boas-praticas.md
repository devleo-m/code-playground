# Módulo 5: Documentação e Comandos em Go

## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende documentação e comandos, vamos falar sobre **como usá-los de forma eficiente e profissional**. Esta é a parte que separa desenvolvedores iniciantes de desenvolvedores experientes.

---

## 1. Performance: Documentação Não Afeta Performance

### ⚠️ Mito Comum: "Documentação Torna o Código Mais Lento"

**Verdade**: Documentação em Go **NÃO afeta a performance** do código executável.

**Por quê?**

- Comentários de documentação são **removidos durante a compilação**
- Eles existem apenas no código-fonte
- O binário compilado não contém comentários
- Zero impacto em tempo de execução

**Exemplo:**

```go
// Esta documentação extensa não afeta a performance
// Soma adiciona dois números e retorna o resultado.
// Esta função é otimizada para performance e usa
// operações de ponto flutuante nativas do processador.
func Soma(a, b float64) float64 {
    return a + b  // Código compilado é idêntico, com ou sem documentação
}
```

**Conclusão**: Documente à vontade! Não há custo de performance.

---

## 2. Performance: Uso Eficiente de `go doc`

### ✅ O Que Deve Ser Feito

#### 2.1. Usar `go doc` para Consultas Rápidas

**BOM**: Consultas diretas são rápidas e eficientes

```bash
# Consulta rápida e direta
go doc fmt.Println
```

**Por quê?**: `go doc` é otimizado para consultas rápidas. Não precisa abrir navegador ou carregar interface web.

#### 2.2. Usar `godoc` para Exploração Profunda

**BOM**: Use `godoc` quando precisar explorar e descobrir

```bash
# Iniciar servidor uma vez
godoc -http=:6060

# Explorar na interface web
# Mais eficiente para descoberta e aprendizado
```

**Por quê?**: Interface web facilita navegação e descoberta de novos pacotes.

---

## 3. Boas Práticas: O Que Deve Ser Utilizado

### Prática 1: Sempre Documente APIs Públicas

**✅ BOM:**

```go
// Package calculadora fornece operações matemáticas básicas.
package calculadora

// Soma adiciona dois números e retorna o resultado.
func Soma(a, b float64) float64 {
    return a + b
}
```

**❌ RUIM:**

```go
package calculadora

// Função que soma
func Soma(a, b float64) float64 {
    return a + b
}
```

**Por quê?**: APIs públicas são a interface do seu código. Sem documentação, outros desenvolvedores não sabem como usar.

---

### Prática 2: Use Exemplos Quando Útil

**✅ BOM:**

```go
// ParseInt converte uma string em inteiro na base especificada.
//
// Exemplo:
//
//	valor, err := ParseInt("42", 10, 64)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(valor) // 42
func ParseInt(s string, base int, bitSize int) (int64, error) {
    // ...
}
```

**⚠️ EVITE:**

```go
// ParseInt converte string para int.
// Muitos exemplos desnecessários para funções simples
func ParseInt(s string, base int, bitSize int) (int64, error) {
    // ...
}
```

**Regra**: Use exemplos para funções complexas ou quando o uso não é óbvio.

---

### Prática 3: Documente Comportamentos Especiais

**✅ BOM:**

```go
// ProcessaDados processa os dados fornecidos de forma assíncrona.
// Esta função pode bloquear por até 30 segundos.
// Retorna erro se o contexto for cancelado.
// Não é thread-safe - use mutex se necessário.
func ProcessaDados(ctx context.Context, dados []byte) error {
    // ...
}
```

**❌ RUIM:**

```go
// ProcessaDados processa dados.
func ProcessaDados(ctx context.Context, dados []byte) error {
    // ...
}
```

**Por quê?**: Comportamentos especiais (bloqueio, thread-safety, timeouts) devem ser documentados para evitar bugs.

---

### Prática 4: Mantenha Documentação Atualizada

**✅ BOM:**

```go
// Soma adiciona dois números e retorna o resultado.
// Agora suporta números negativos (atualizado em 2024).
func Soma(a, b float64) float64 {
    return a + b
}
```

**❌ RUIM:**

```go
// Soma adiciona dois números positivos.
// (Mas a função agora aceita negativos também - documentação desatualizada!)
func Soma(a, b float64) float64 {
    return a + b
}
```

**Regra**: Quando mudar código, atualize a documentação também.

---

### Prática 5: Use `go doc` Regularmente

**✅ BOM:**

```bash
# Hábito: Consultar documentação antes de usar função nova
go doc strings.Contains
# Lê a documentação
# Entende como usar
# Escreve código correto
```

**❌ RUIM:**

```bash
# Adivinha como usar
# Tenta várias vezes
# Procura no Google
# Perde tempo
```

**Por quê?**: `go doc` é mais rápido que buscar na internet e mostra documentação oficial.

---

## 4. O Que NÃO Deve Ser Feito

### ❌ Erro 1: Documentação Obvia ou Redundante

**RUIM:**

```go
// Soma soma dois números.
// a é o primeiro número.
// b é o segundo número.
// Retorna a soma de a e b.
func Soma(a, b float64) float64 {
    return a + b
}
```

**BOM:**

```go
// Soma adiciona dois números e retorna o resultado.
func Soma(a, b float64) float64 {
    return a + b
}
```

**Por quê?**: Documentação óbvia não adiciona valor e polui o código.

---

### ❌ Erro 2: Documentação Incorreta ou Enganosa

**RUIM:**

```go
// Multiplica dois números e retorna o resultado.
func Soma(a, b float64) float64 {  // Função soma, mas documentação diz multiplica!
    return a + b
}
```

**BOM:**

```go
// Soma adiciona dois números e retorna o resultado.
func Soma(a, b float64) float64 {
    return a + b
}
```

**Por quê?**: Documentação incorreta é pior que sem documentação - engana desenvolvedores.

---

### ❌ Erro 3: Não Documentar APIs Públicas

**RUIM:**

```go
package calculadora

// Função pública sem documentação
func Soma(a, b float64) float64 {
    return a + b
}
```

**BOM:**

```go
package calculadora

// Soma adiciona dois números e retorna o resultado.
func Soma(a, b float64) float64 {
    return a + b
}
```

**Por quê?**: Funções públicas são a interface do seu código. Devem ser documentadas.

---

### ❌ Erro 4: Documentação Excessiva em Código Privado

**RUIM:**

```go
// somaInterna soma dois números internamente.
// Esta função é usada apenas dentro do pacote.
// Não deve ser chamada externamente.
// Implementa algoritmo otimizado de soma.
// Usa operações de ponto flutuante.
func somaInterna(a, b float64) float64 {
    return a + b
}
```

**BOM:**

```go
// somaInterna é uma função auxiliar para cálculos internos.
func somaInterna(a, b float64) float64 {
    return a + b
}
```

**Por quê?**: Documentação excessiva em código privado pode ser desnecessária. Foque em APIs públicas.

---

### ❌ Erro 5: Ignorar `go help`

**RUIM:**

```bash
# Nunca consulta go help
# Adivinha sintaxe de comandos
# Comete erros
go build --wrong-flag
```

**BOM:**

```bash
# Consulta ajuda quando necessário
go help build
# Lê a documentação
# Usa corretamente
go build -o meuapp ./cmd
```

**Por quê?**: `go help` é a fonte oficial de verdade sobre comandos Go.

---

## 5. Padrões de Documentação por Tipo

### Documentando Funções

**Estrutura recomendada:**

```go
// NomeDaFuncao faz uma ação específica.
//
// Parâmetros:
//   - param1: descrição do primeiro parâmetro
//   - param2: descrição do segundo parâmetro
//
// Retorna:
//   - tipo1: descrição do primeiro retorno
//   - error: descrição do erro (se aplicável)
//
// Exemplo:
//
//	resultado, err := NomeDaFuncao(valor1, valor2)
//	if err != nil {
//	    log.Fatal(err)
//	}
func NomeDaFuncao(param1 Tipo1, param2 Tipo2) (TipoRetorno, error) {
    // ...
}
```

---

### Documentando Tipos

**Estrutura recomendada:**

```go
// TipoNome representa um conceito específico.
// Use este tipo quando precisar fazer X.
//
// Exemplo:
//
//	valor := TipoNome{
//	    Campo1: "valor",
//	    Campo2: 42,
//	}
type TipoNome struct {
    // Campo1 faz algo específico.
    Campo1 string

    // Campo2 faz outra coisa.
    Campo2 int
}
```

---

### Documentando Métodos

**Estrutura recomendada:**

```go
// MetodoNome faz uma ação no receptor.
//
// Exemplo:
//
//	obj := NovoObjeto()
//	obj.MetodoNome(parametro)
func (o *Objeto) MetodoNome(param Tipo) Retorno {
    // ...
}
```

---

### Documentando Pacotes

**Estrutura recomendada:**

```go
// Package nomepacote fornece funcionalidade X.
//
// Este pacote é útil para fazer Y.
// Use quando precisar de Z.
//
// Exemplo básico:
//
//	resultado := nomepacote.FuncaoPrincipal()
package nomepacote
```

---

## 6. Ferramentas de Verificação

### `go vet` - Verificação Básica

```bash
go vet ./...
```

**O que verifica:**

- Erros comuns
- Problemas de estilo
- Alguns problemas de documentação

**Limitação**: Não verifica se funções públicas estão documentadas.

---

### `golangci-lint` - Análise Avançada

```bash
# Instalar
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Executar
golangci-lint run
```

**O que verifica:**

- Funções públicas sem documentação
- Problemas de estilo
- Muitos outros problemas

**Configuração mínima (`.golangci.yml`):**

```yaml
linters:
  enable:
    - godot # Verifica pontos finais em documentação
    - godox # Verifica comentários TODO/FIXME
    - gofmt # Verifica formatação
    - goimports # Verifica imports
    - misspell # Verifica erros de ortografia
```

---

## 7. Integração com IDEs

### VS Code

**Extensão Go oficial** mostra documentação automaticamente:

- Hover sobre função mostra documentação
- Autocomplete inclui documentação
- `go doc` integrado

**Configuração:**

```json
{
  "go.docsTool": "godoc",
  "go.lintTool": "golangci-lint"
}
```

---

### GoLand

**IDE da JetBrains** tem suporte completo:

- Documentação inline
- Navegação para definição
- Geração automática de documentação

---

## 8. Workflow Recomendado

### Ao Escrever Código Novo

1. **Escreva a função**
2. **Documente imediatamente** (não deixe para depois)
3. **Teste a documentação** com `go doc`
4. **Verifique com `go vet`**

### Ao Modificar Código Existente

1. **Modifique o código**
2. **Atualize a documentação** se necessário
3. **Verifique se ainda está correta** com `go doc`
4. **Execute `go vet`**

### Ao Explorar Bibliotecas

1. **Use `go doc`** para consultas rápidas
2. **Use `godoc`** para exploração profunda
3. **Leia exemplos** na documentação
4. **Experimente** no seu código

---

## 9. Checklist de Revisão

Antes de considerar seu código "pronto", pergunte-se:

- [ ] Todas as funções públicas estão documentadas?
- [ ] A documentação do pacote existe e é clara?
- [ ] Exemplos foram incluídos onde necessário?
- [ ] Parâmetros e retornos estão documentados?
- [ ] Comportamentos especiais estão documentados?
- [ ] A documentação está atualizada com o código?
- [ ] `go doc` mostra a documentação corretamente?
- [ ] `go vet` não reporta problemas?
- [ ] A documentação segue o padrão Go?

---

## 10. Casos Especiais

### Documentando Interfaces

```go
// Processador processa dados de entrada.
type Processador interface {
    // Processa processa os dados fornecidos.
    // Retorna erro se o processamento falhar.
    Processa(dados []byte) error
}
```

---

### Documentando Constantes

```go
const (
    // StatusAtivo indica que o item está ativo.
    StatusAtivo = 1

    // StatusInativo indica que o item está inativo.
    StatusInativo = 2
)
```

---

### Documentando Variáveis

```go
// MaxTentativas é o número máximo de tentativas permitidas.
var MaxTentativas = 3
```

---

## 11. Performance: Impacto de Ferramentas

### `go doc` - Muito Rápido

- Consulta local (sem rede)
- Análise de código-fonte local
- Resposta em milissegundos
- Zero impacto em build

**Uso recomendado**: Consultas frequentes durante desenvolvimento.

---

### `godoc` - Servidor Local

- Servidor web local
- Consome alguns MB de RAM
- Inicialização em segundos
- Útil para exploração

**Uso recomendado**: Sessões de exploração e aprendizado.

---

### `go help` - Instantâneo

- Informação embutida
- Resposta imediata
- Zero overhead

**Uso recomendado**: Sempre que tiver dúvida sobre comandos.

---

## 12. Resumo: Regras de Ouro

1. **Documente APIs públicas**: Sempre, sem exceção
2. **Seja claro e conciso**: Documentação não é romance
3. **Use exemplos quando útil**: Mas não exagere
4. **Mantenha atualizada**: Documentação desatualizada é pior que sem documentação
5. **Use `go doc` regularmente**: É sua ferramenta de aprendizado
6. **Consulte `go help`**: Fonte oficial de verdade
7. **Explore com `godoc`**: Aprenda a biblioteca padrão
8. **Verifique com ferramentas**: `go vet`, `golangci-lint`
9. **Documente comportamentos especiais**: Bloqueios, thread-safety, etc.
10. **Não documente o óbvio**: Mas documente o não-óbvio

---

## 13. Comparação: Com vs. Sem Boas Práticas

### Cenário: Biblioteca Pública

**SEM boas práticas:**

- Funções sem documentação
- Desenvolvedores não sabem como usar
- Muitas issues no GitHub
- Biblioteca pouco usada

**COM boas práticas:**

- Tudo documentado
- Desenvolvedores usam facilmente
- Poucas dúvidas
- Biblioteca amplamente adotada

---

### Cenário: Código Interno

**SEM boas práticas:**

- Time precisa adivinhar como usar
- Muitas perguntas em reuniões
- Refatoração difícil
- Bugs por uso incorreto

**COM boas práticas:**

- Time entende rapidamente
- Poucas perguntas
- Refatoração mais segura
- Menos bugs

---

## 14. Ferramentas Adicionais

### `go list` - Listar Pacotes

```bash
# Listar todos os pacotes
go list ./...

# Ver informações de um pacote
go list -json ./meupacote
```

---

### `go mod` - Gerenciamento de Módulos

```bash
# Ver ajuda
go help mod

# Comandos úteis
go mod tidy      # Limpar dependências
go mod download  # Baixar dependências
go mod graph     # Ver grafo de dependências
```

---

### `go env` - Variáveis de Ambiente

```bash
# Ver todas as variáveis
go env

# Ver variável específica
go env GOPATH
go env GOROOT
```

---

## 15. Integração com CI/CD

### Verificação Automática

**GitHub Actions exemplo:**

```yaml
name: Lint
on: [push, pull_request]
jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
      - run: go vet ./...
      - run: golangci-lint run
```

**Benefícios:**

- Documentação verificada automaticamente
- Problemas detectados antes do merge
- Qualidade mantida

---

## 16. Recursos Adicionais

### Documentação Oficial

- [Effective Go](https://go.dev/doc/effective_go) - Guia de boas práticas
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) - Comentários de revisão
- [Go Blog](https://go.dev/blog) - Artigos sobre Go

### Ferramentas

- `godoc` - Interface web
- `golangci-lint` - Linter avançado
- `staticcheck` - Análise estática

---

## Conclusão

Documentação e conhecimento de ferramentas são fundamentais para:

1. **Código profissional**: Bem documentado é mais fácil de manter
2. **Produtividade**: Ferramentas certas economizam tempo
3. **Colaboração**: Time trabalha melhor com código documentado
4. **Aprendizado**: `go doc` e `godoc` são ferramentas de aprendizado
5. **Qualidade**: Documentação é parte da qualidade do código

Lembre-se: **Código é lido muito mais vezes do que escrito**. Documentação ajuda quem lê (incluindo você no futuro) a entender rapidamente.

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!
