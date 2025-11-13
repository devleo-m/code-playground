# Módulo 5: Documentação e Comandos em Go

## Aula 1: Documentação e Comandos - Ferramentas Essenciais

Olá! Bem-vindo ao quinto módulo. Até agora você aprendeu a escrever código em Go, mas uma parte crucial do desenvolvimento profissional é **documentar seu código** e **saber usar as ferramentas** que Go oferece para explorar e entender código.

Nesta aula, vamos mergulhar nas ferramentas de documentação do Go e nos comandos essenciais que todo desenvolvedor Go precisa conhecer.

---

## 1. Por Que Documentar Código?

Documentação não é apenas "bom ter" - é **essencial** para:

1. **Manutenibilidade**: Código sem documentação é difícil de manter, mesmo para quem o escreveu
2. **Colaboração**: Outros desenvolvedores precisam entender seu código
3. **API Pública**: Se você cria bibliotecas, a documentação é a interface com seus usuários
4. **Auto-documentação**: Go gera documentação automaticamente a partir de comentários especiais
5. **Exploração**: Você pode explorar a biblioteca padrão do Go usando ferramentas de documentação

---

## 2. Comentários de Documentação em Go

Go usa uma convenção especial para comentários de documentação. Eles devem:

- Começar imediatamente antes da declaração (sem linha em branco)
- Começar com o nome do item sendo documentado
- Ser completos e descritivos
- Usar frases completas

### Sintaxe Básica

```go
// MinhaFuncao faz algo importante.
// Esta é a primeira linha da documentação.
func MinhaFuncao() {
    // ...
}
```

**Regra de Ouro**: Se começa com o nome do item, é documentação. Se não, é apenas um comentário.

### Exemplo: Documentando uma Função

```go
// CalculaMedia calcula a média aritmética de uma lista de números.
//
// Parâmetros:
//   - numeros: slice de números float64
//
// Retorna:
//   - float64: a média calculada
//   - error: erro se a lista estiver vazia
//
// Exemplo:
//
//	numeros := []float64{10, 20, 30}
//	media, err := CalculaMedia(numeros)
//	// media será 20.0
func CalculaMedia(numeros []float64) (float64, error) {
    if len(numeros) == 0 {
        return 0, fmt.Errorf("lista vazia")
    }
    soma := 0.0
    for _, n := range numeros {
        soma += n
    }
    return soma / float64(len(numeros)), nil
}
```

### Documentando Tipos e Structs

```go
// Usuario representa um usuário do sistema.
// Contém informações básicas de identificação.
type Usuario struct {
    // ID é o identificador único do usuário.
    ID int

    // Nome é o nome completo do usuário.
    Nome string

    // Email é o endereço de email válido.
    Email string
}

// NovoUsuario cria uma nova instância de Usuario.
// Valida os dados antes de criar.
func NovoUsuario(nome, email string) (*Usuario, error) {
    // ...
}
```

### Documentando Pacotes

A documentação do pacote deve estar no início do arquivo, antes da declaração `package`:

```go
// Package calculadora fornece funções para operações matemáticas básicas.
//
// Este pacote inclui operações como soma, subtração, multiplicação
// e divisão, além de funções mais avançadas como potenciação.
//
// Exemplo de uso:
//
//	resultado := calculadora.Soma(5, 3)
//	fmt.Println(resultado) // 8
package calculadora
```

**Nota**: Em arquivos reais, a documentação do pacote fica no mesmo arquivo, mas antes de `package`.

---

## 3. O Comando `go doc`

O `go doc` é uma ferramenta de linha de comando que exibe documentação de pacotes, tipos, funções e constantes.

### Sintaxe Básica

```bash
go doc [pacote]
go doc [pacote].[item]
go doc [pacote].[tipo].[metodo]
```

### Exemplos Práticos

**Ver documentação de um pacote:**

```bash
go doc fmt
```

**Ver documentação de uma função específica:**

```bash
go doc fmt.Println
```

**Ver documentação de um tipo:**

```bash
go doc time.Time
```

**Ver documentação de um método:**

```bash
go doc time.Time.Format
```

### Flags Úteis do `go doc`

```bash
# Mostrar código-fonte junto com documentação
go doc -src fmt.Println

# Mostrar todos os símbolos exportados
go doc -all fmt

# Buscar por padrão
go doc -all fmt | grep Print
```

### Usando `go doc` no Seu Próprio Código

Se você está em um projeto Go, pode ver a documentação do seu próprio código:

```bash
# No diretório do seu projeto
go doc .

# Ver documentação de uma função específica
go doc .MinhaFuncao
```

---

## 4. O Comando `go help`

O `go help` é sua referência para todos os comandos disponíveis no Go.

### Ver Todos os Comandos

```bash
go help
```

Isso lista todos os comandos disponíveis:

- `go build` - compila pacotes
- `go run` - compila e executa
- `go test` - executa testes
- `go doc` - exibe documentação
- `go fmt` - formata código
- E muitos outros...

### Ver Ajuda de um Comando Específico

```bash
go help build
go help test
go help doc
go help mod
```

### Exemplo de Saída

```bash
$ go help doc

usage: go doc [-all] [-c] [-cmd] [-u] [package|[package.]symbol[.methodOrField]]

Doc prints the documentation comments associated with the item identified by its
arguments (a package, const, func, type, var, method, or struct field) followed
by a one-line summary of each of the first-level items "under" that item (package-level
declarations for a package, methods for a type, etc.).

...
```

---

## 5. Godoc: Interface Web de Documentação

O `godoc` é uma ferramenta que gera uma interface web para visualizar documentação. É especialmente útil para explorar a biblioteca padrão do Go.

### Instalando o Godoc

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

**Nota**: Em versões antigas do Go, `godoc` vinha incluído. Nas versões modernas, precisa ser instalado separadamente.

### Executando o Servidor Godoc

```bash
godoc -http=:6060
```

Isso inicia um servidor web local na porta 6060. Acesse no navegador:

```
http://localhost:6060
```

### Navegando na Interface

- **Biblioteca Padrão**: Veja todos os pacotes da biblioteca padrão do Go
- **Seu Código**: Se você estiver em um projeto Go, pode ver a documentação do seu próprio código
- **Busca**: Use a barra de busca para encontrar pacotes e funções
- **Código-Fonte**: Clique em "Source" para ver o código-fonte

### Flags Úteis do Godoc

```bash
# Especificar porta diferente
godoc -http=:8080

# Incluir código-fonte
godoc -http=:6060 -src

# Mostrar apenas pacotes locais
godoc -http=:6060 -path=.
```

---

## 6. Explorando a Biblioteca Padrão

Uma das melhores formas de aprender Go é explorar a biblioteca padrão usando `go doc` ou `godoc`.

### Exemplos de Exploração

**Descobrir funções de strings:**

```bash
go doc strings
```

**Ver funções específicas:**

```bash
go doc strings.Contains
go doc strings.Split
go doc strings.Join
```

**Explorar o pacote `time`:**

```bash
go doc time
go doc time.Now
go doc time.Duration
```

**Ver métodos de slices:**

```bash
go doc sort
go doc sort.Ints
```

### Estratégia de Aprendizado

1. **Use `go doc`** para descobrir o que um pacote faz
2. **Leia a documentação** para entender como usar
3. **Veja exemplos** no código-fonte ou na web
4. **Experimente** no seu próprio código

---

## 7. Boas Práticas de Documentação

### O Que Documentar

✅ **SEMPRE documente:**

- Funções e métodos públicos (exportados)
- Tipos e structs públicos
- Constantes públicas importantes
- Pacotes (pelo menos um arquivo do pacote)

❌ **NÃO precisa documentar:**

- Funções privadas (a menos que sejam complexas)
- Variáveis óbvias
- Código auto-explicativo

### Como Escrever Boa Documentação

**1. Seja Conciso mas Completo**

```go
// ❌ RUIM: Muito vago
// Funcao faz algo.

// ✅ BOM: Específico e útil
// CalculaMedia calcula a média aritmética de uma lista de números.
```

**2. Comece com o Nome do Item**

```go
// ✅ BOM
// Soma adiciona dois números.

// ❌ RUIM
// Esta função adiciona dois números.
```

**3. Use Exemplos Quando Útil**

```go
// ParseInt converte uma string em um inteiro na base especificada.
//
// Exemplo:
//
//	valor, err := ParseInt("42", 10, 64)
//	// valor será 42
```

**4. Documente Parâmetros e Retornos**

```go
// Divide divide a por b e retorna o resultado.
//
// Parâmetros:
//   - a: o dividendo
//   - b: o divisor (não pode ser zero)
//
// Retorna:
//   - float64: resultado da divisão
//   - error: erro se b for zero
```

**5. Documente Comportamentos Especiais**

```go
// ProcessaDados processa os dados fornecidos.
// Esta função pode bloquear por até 30 segundos.
// Retorna erro se o contexto for cancelado.
```

---

## 8. Comandos Relacionados à Documentação

### `go fmt` - Formatação de Código

Embora não seja diretamente sobre documentação, `go fmt` garante que seu código (incluindo comentários) siga o padrão Go:

```bash
go fmt ./...
```

### `go vet` - Análise Estática

`go vet` pode detectar problemas comuns, incluindo documentação faltando:

```bash
go vet ./...
```

### `golint` (Ferramenta Externa)

O `golint` (agora `golangci-lint` ou `staticcheck`) pode verificar se funções públicas estão documentadas:

```bash
# Instalar
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Executar
golangci-lint run
```

---

## 9. Exemplo Completo: Pacote Bem Documentado

```go
// Package calculadora fornece operações matemáticas básicas.
//
// Este pacote implementa uma calculadora simples com suporte
// a operações aritméticas fundamentais. Todas as operações
// são thread-safe e podem ser usadas em programas concorrentes.
//
// Exemplo básico:
//
//	calc := calculadora.Nova()
//	resultado := calc.Soma(5, 3)
//	fmt.Println(resultado) // 8
package calculadora

import "fmt"

// Calculadora representa uma calculadora com histórico de operações.
type Calculadora struct {
	historico []float64
}

// Nova cria uma nova instância de Calculadora.
// Retorna um ponteiro para a calculadora inicializada.
func Nova() *Calculadora {
	return &Calculadora{
		historico: make([]float64, 0),
	}
}

// Soma adiciona dois números e retorna o resultado.
//
// Parâmetros:
//   - a: primeiro número
//   - b: segundo número
//
// Retorna:
//   - float64: resultado da soma
//
// Exemplo:
//
//	resultado := calc.Soma(5, 3)
//	// resultado será 8.0
func (c *Calculadora) Soma(a, b float64) float64 {
	resultado := a + b
	c.historico = append(c.historico, resultado)
	return resultado
}

// Divisao divide a por b.
//
// Retorna um erro se b for zero.
func (c *Calculadora) Divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não permitida")
	}
	resultado := a / b
	c.historico = append(c.historico, resultado)
	return resultado, nil
}
```

---

## 10. Resumo dos Comandos Essenciais

| Comando   | Descrição                    | Exemplo              |
| --------- | ---------------------------- | -------------------- |
| `go doc`  | Exibe documentação           | `go doc fmt.Println` |
| `go help` | Ajuda sobre comandos         | `go help doc`        |
| `godoc`   | Servidor web de documentação | `godoc -http=:6060`  |
| `go fmt`  | Formata código               | `go fmt ./...`       |
| `go vet`  | Análise estática             | `go vet ./...`       |

---

## Conclusão

Documentação e conhecimento das ferramentas são fundamentais para:

1. **Escrever código profissional**: Código bem documentado é mais fácil de manter
2. **Explorar Go**: Use `go doc` e `godoc` para aprender a biblioteca padrão
3. **Colaborar**: Documentação ajuda equipes a trabalhar juntas
4. **Criar bibliotecas**: Se você compartilha código, documentação é essencial

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!
