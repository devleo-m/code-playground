# Módulo 29: Comandos Core do Go
## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas. Esses exercícios vão te ajudar a dominar os comandos core do Go na prática.

---

## Exercícios Práticos

### Exercício 1: Criando e Gerenciando um Projeto do Zero

Crie um projeto Go completo do zero, seguindo todos os passos:

1. **Crie um novo diretório** chamado `meu-projeto-comandos`
2. **Inicialize um módulo Go** com o nome `github.com/seu-usuario/meu-projeto-comandos`
3. **Crie um arquivo `main.go`** com o seguinte código (intencionalmente mal formatado):

```go
package main
import "fmt"
import "math"
func main(){
x:=10.5
y:=20.3
resultado:=math.Sqrt(x*x+y*y)
fmt.Printf("A distância é: %.2f\n",resultado)
}
```

4. **Execute os seguintes comandos** e documente o que cada um faz:
   - `go fmt main.go` (antes e depois - compare o código)
   - `go run main.go`
   - `go build -o calculadora main.go`
   - `go version`
   - `go doc math.Sqrt`
   - `go mod tidy`

5. **Crie um arquivo de teste** `main_test.go` com um teste para uma função que calcula a média de dois números

6. **Execute** `go test -v` e documente o resultado

**Entregável**: 
- Crie um arquivo `exercicio1.md` documentando:
  - O output de cada comando
  - As mudanças que `go fmt` fez no código
  - O resultado dos testes

---

### Exercício 2: Cross-Compilation e Instalação

Neste exercício, você vai praticar compilação para diferentes plataformas e instalação de ferramentas.

1. **Crie um programa simples** `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá do Go!")
    fmt.Printf("Versão do Go: %s\n", runtime.Version())
}
```

**Nota**: Você precisará adicionar `import "runtime"` para a função `runtime.Version()` funcionar.

2. **Compile para diferentes plataformas**:
   - Compile para Linux (mesmo que você esteja no Windows/macOS)
   - Compile para Windows (mesmo que você esteja no Linux/macOS)
   - Compile para macOS (mesmo que você esteja no Linux/Windows)
   
   Use os comandos apropriados com `GOOS` e `GOARCH`.

3. **Verifique os binários criados**:
   - Liste os arquivos criados
   - Use `go version [binário]` para verificar cada um
   - Documente as diferenças de tamanho entre os binários

4. **Instale uma ferramenta**:
   - Instale o `godoc` usando `go install`
   - Verifique onde foi instalado usando `which godoc` (Linux/macOS) ou `where godoc` (Windows)
   - Execute `godoc -http=:6060` e acesse no navegador (opcional)

**Entregável**:
- Crie um arquivo `exercicio2.md` com:
  - Os comandos usados para cross-compilation
  - O tamanho de cada binário criado
  - O caminho onde `godoc` foi instalado
  - Screenshots ou descrição do que você viu

---

### Exercício 3: Gerenciamento de Dependências e Limpeza

Este exercício foca em gerenciamento de módulos e limpeza de cache.

1. **Crie um novo projeto** `projeto-deps`:
   ```bash
   mkdir projeto-deps
   cd projeto-deps
   go mod init github.com/seu-usuario/projeto-deps
   ```

2. **Adicione uma dependência externa**:
   - Crie um `main.go` que use o pacote `github.com/fatih/color` para imprimir texto colorido
   - Execute `go run main.go` (isso vai baixar a dependência automaticamente)
   - Verifique o arquivo `go.mod` criado

3. **Explore os comandos de módulo**:
   - Execute `go mod graph` e documente o output
   - Execute `go mod why github.com/fatih/color` e documente o resultado
   - Execute `go mod download` e verifique o cache
   - Execute `go mod verify` e documente o resultado

4. **Pratique limpeza**:
   - Verifique o tamanho do cache antes: `du -sh $(go env GOCACHE)` (Linux/macOS) ou equivalente no Windows
   - Execute `go clean -cache`
   - Verifique o tamanho novamente
   - Execute `go clean -modcache` (cuidado: isso remove todos os módulos baixados!)
   - Execute `go mod download` novamente para baixar as dependências

5. **Organize as dependências**:
   - Adicione uma dependência que você não vai usar no código
   - Execute `go mod tidy`
   - Verifique se a dependência não usada foi removida do `go.mod`

**Entregável**:
- Crie um arquivo `exercicio3.md` com:
  - O conteúdo do `go.mod` antes e depois de `go mod tidy`
  - O output de `go mod graph`
  - O tamanho do cache antes e depois da limpeza
  - Suas observações sobre o processo

---

### Exercício 4: Testes, Cobertura e Documentação

Este exercício combina testes, cobertura e documentação.

1. **Crie um pacote `calculadora`** com as seguintes funções bem documentadas:

```go
// Package calculadora fornece operações matemáticas básicas.
package calculadora

// Soma adiciona dois números inteiros e retorna o resultado.
func Soma(a, b int) int {
    return a + b
}

// Multiplica multiplica dois números inteiros e retorna o resultado.
func Multiplica(a, b int) int {
    return a * b
}

// Divide divide a por b e retorna o resultado.
// Retorna um erro se b for zero.
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero não permitida")
    }
    return a / b, nil
}
```

**Nota**: Você precisará adicionar os imports necessários (`fmt`).

2. **Crie testes completos** em `calculadora_test.go`:
   - Teste para `Soma` com casos positivos e negativos
   - Teste para `Multiplica`
   - Teste para `Divide` com casos de sucesso e erro (divisão por zero)

3. **Execute os testes**:
   - `go test -v` (mostre o output)
   - `go test -cover` (mostre a cobertura)
   - `go test -coverprofile=coverage.out` e depois `go tool cover -html=coverage.out` (abra o HTML gerado)

4. **Crie benchmarks**:
   - Adicione benchmarks para `Soma` e `Multiplica`
   - Execute `go test -bench=. -benchmem`
   - Documente os resultados

5. **Explore a documentação**:
   - Execute `go doc calculadora`
   - Execute `go doc calculadora.Soma`
   - Execute `go doc calculadora.Divide`
   - Execute `go doc -src calculadora.Soma`

**Entregável**:
- Crie um arquivo `exercicio4.md` com:
  - O código completo dos testes
  - O output de `go test -v`
  - A porcentagem de cobertura
  - Os resultados dos benchmarks
  - Screenshots ou descrição da documentação gerada

---

## Perguntas de Reflexão

### Reflexão 1: Escolhendo o Comando Correto

Imagine os seguintes cenários. Para cada um, **escolha qual comando Go você usaria** e **explique por quê**:

**Cenário A**: Você acabou de escrever 50 linhas de código e quer testar rapidamente se funciona.

**Cenário B**: Você precisa criar um executável para distribuir para 100 usuários que usam Windows, mas você está desenvolvendo no macOS.

**Cenário C**: Você está trabalhando em um projeto grande com 20 desenvolvedores. Antes de fazer commit, você quer garantir que o código está no padrão.

**Cenário D**: Você encontrou um bug estranho que só acontece às vezes. Você suspeita que pode ser cache corrompido.

**Cenário E**: Você quer instalar uma ferramenta de linting (`golangci-lint`) para usar em todos os seus projetos.

**Cenário F**: Você está aprendendo Go e quer entender como usar a função `strings.Split`.

**Sua resposta deve incluir**:
- O comando escolhido para cada cenário
- A justificativa técnica por trás de cada escolha
- Alternativas que você considerou e por que não as escolheu

---

### Reflexão 2: Workflow e Boas Práticas

Pense sobre o **workflow de desenvolvimento** de um programador Go profissional.

**Pergunta 1**: Crie um **workflow ideal** (sequência de comandos) que um desenvolvedor deveria seguir desde o momento que começa a escrever código até fazer commit. Inclua:

- Quando usar cada comando
- A ordem ideal de execução
- Por que essa ordem é importante
- O que acontece se você pular algum passo

**Pergunta 2**: Analise a seguinte situação:

> "Um desenvolvedor júnior sempre usa `go run` para testar o código durante o desenvolvimento. Quando precisa fazer deploy, ele simplesmente executa `go run main.go` no servidor de produção."

**Identifique os problemas** dessa abordagem e explique:
- Por que isso é problemático?
- Quais comandos deveriam ser usados em cada etapa?
- Quais são as consequências de usar `go run` em produção?
- Como você explicaria a diferença para esse desenvolvedor júnior?

**Pergunta 3**: Considere a importância de `go mod tidy`:

- Por que é importante executar `go mod tidy` regularmente?
- O que pode acontecer se você não executar esse comando?
- Como isso afeta colaboração em equipe?
- Qual é a melhor prática: executar manualmente ou automatizar?

**Sua resposta deve demonstrar**:
- Compreensão profunda do ciclo de desenvolvimento
- Capacidade de identificar problemas em workflows
- Conhecimento sobre quando usar cada ferramenta
- Pensamento crítico sobre boas práticas

---

## Como Entregar

Crie os seguintes arquivos na pasta `29-core-go-commands/`:

1. `exercicio1.md` - Documentação do Exercício 1
2. `exercicio2.md` - Documentação do Exercício 2
3. `exercicio3.md` - Documentação do Exercício 3
4. `exercicio4.md` - Documentação do Exercício 4
5. `reflexoes.md` - Respostas às perguntas de reflexão

**Estrutura de pastas sugerida**:
```
29-core-go-commands/
├── exercicios/
│   ├── exercicio1/
│   │   ├── main.go
│   │   └── main_test.go
│   ├── exercicio2/
│   │   └── hello.go
│   ├── exercicio3/
│   │   └── main.go
│   └── exercicio4/
│       ├── calculadora.go
│       └── calculadora_test.go
├── exercicio1.md
├── exercicio2.md
├── exercicio3.md
├── exercicio4.md
└── reflexoes.md
```

**Importante**: 
- Execute cada comando e documente os resultados reais
- Seja honesto nas reflexões - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão
- Inclua screenshots quando relevante (especialmente para cobertura de testes e godoc)
- Comente seu código explicando o que cada parte faz

---

## Dicas para Sucesso

1. **Leia a documentação**: Use `go help [comando]` quando tiver dúvidas
2. **Experimente**: Não tenha medo de testar diferentes flags e opções
3. **Observe os outputs**: Preste atenção nas mensagens e erros
4. **Compare**: Veja as diferenças antes e depois de executar comandos
5. **Documente**: Anote tudo que você descobrir durante os exercícios

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!



