# Módulo 29: Comandos Core do Go
## Aula 1: Core Go Commands - Ferramentas Essenciais da Linha de Comando

Olá! Bem-vindo a este módulo fundamental. Até agora você aprendeu a escrever código em Go, mas uma parte crucial do desenvolvimento profissional é **dominar as ferramentas de linha de comando** que Go oferece. Esses comandos são a base do seu workflow diário como desenvolvedor Go.

Nesta aula, vamos mergulhar nos **9 comandos core** do Go que todo desenvolvedor precisa conhecer e dominar. Cada um deles resolve problemas específicos no ciclo de desenvolvimento, desde compilar e executar código até gerenciar dependências e executar testes.

---

## 1. `go run` - Compilar e Executar em Um Passo

### O Que É?

O `go run` é o comando mais simples e direto para executar código Go. Ele **compila e executa** seu programa em um único passo, sem criar arquivos executáveis permanentes.

### Sintaxe Básica

```bash
go run [arquivo.go]
go run [pacote]
go run ./...
```

### Características Principais

- ✅ **Compilação temporária**: Compila o código em memória ou em arquivos temporários
- ✅ **Execução imediata**: Executa o programa logo após compilar
- ✅ **Sem artefatos**: Não deixa arquivos `.exe` ou binários no diretório
- ✅ **Ideal para desenvolvimento**: Perfeito para testar código rapidamente

### Exemplos Práticos

**Executar um arquivo único:**

```bash
go run main.go
```

**Executar múltiplos arquivos:**

```bash
go run main.go utils.go
```

**Executar um pacote inteiro:**

```bash
go run .
```

**Executar todos os pacotes recursivamente:**

```bash
go run ./...
```

### Quando Usar?

- ✅ Durante o desenvolvimento para testar rapidamente
- ✅ Para executar scripts Go temporários
- ✅ Para demonstrações e protótipos
- ✅ Quando você não precisa do binário final

### Quando NÃO Usar?

- ❌ Para produção (use `go build` para criar binários)
- ❌ Quando precisa do executável para distribuir
- ❌ Para medir performance real (compilação pode afetar)

### Exemplo Completo

```bash
# Criar um arquivo simples
cat > hello.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Olá, mundo!")
}
EOF

# Executar diretamente
go run hello.go
# Output: Olá, mundo!
```

---

## 2. `go build` - Compilar para Binários

### O Que É?

O `go build` é o comando que **compila** seu código Go em um arquivo executável binário. É o comando essencial para criar programas que você pode distribuir e executar em qualquer máquina compatível.

### Sintaxe Básica

```bash
go build [pacote]
go build -o [nome] [pacote]
go build [flags] [pacote]
```

### Características Principais

- ✅ **Binários estáticos**: Por padrão, cria binários que não dependem de bibliotecas externas
- ✅ **Cross-compilation**: Suporta compilação para diferentes OS/arquiteturas
- ✅ **Otimização**: Inclui otimizações de compilação
- ✅ **Build constraints**: Suporta tags de build condicionais

### Flags Importantes

```bash
# Especificar nome do executável
go build -o meuapp main.go

# Compilar para diferentes sistemas operacionais
GOOS=linux go build main.go      # Linux
GOOS=windows go build main.go    # Windows
GOOS=darwin go build main.go     # macOS

# Compilar para diferentes arquiteturas
GOARCH=amd64 go build main.go     # 64-bit
GOARCH=arm64 go build main.go    # ARM 64-bit

# Combinar OS e arquitetura
GOOS=linux GOARCH=arm64 go build main.go

# Build com tags (build constraints)
go build -tags dev main.go

# Mostrar informações de build
go build -v main.go              # Verbose
go build -x main.go               # Mostra comandos executados
```

### Exemplos Práticos

**Compilar o pacote atual:**

```bash
go build
# Cria um executável com o nome do diretório
```

**Compilar com nome específico:**

```bash
go build -o minha-aplicacao main.go
```

**Cross-compilation para Windows (do Linux/macOS):**

```bash
GOOS=windows GOARCH=amd64 go build -o app.exe main.go
```

**Compilar múltiplos pacotes:**

```bash
go build ./cmd/...  # Todos os pacotes em cmd/
```

### Estrutura de Binários

Por padrão, `go build` cria binários:
- **Linux/macOS**: Arquivo executável sem extensão
- **Windows**: Arquivo `.exe`

### Quando Usar?

- ✅ Para criar executáveis para produção
- ✅ Para distribuir aplicações
- ✅ Para cross-compilation
- ✅ Quando precisa do binário final

### Exemplo Completo

```bash
# Estrutura do projeto
projeto/
├── main.go
└── utils.go

# Compilar
go build -o meuapp .

# Executar o binário
./meuapp
```

---

## 3. `go install` - Instalar Ferramentas e Pacotes

### O Que É?

O `go install` **compila e instala** pacotes e dependências. Para pacotes `main`, cria executáveis no diretório `$GOPATH/bin` (ou `$GOBIN` se configurado), tornando-os disponíveis globalmente no sistema.

### Sintaxe Básica

```bash
go install [pacote]@[versão]
go install [pacote]
go install ./...
```

### Características Principais

- ✅ **Instalação global**: Torna ferramentas disponíveis no PATH
- ✅ **Versionamento**: Suporta instalação de versões específicas
- ✅ **Gerenciamento de dependências**: Baixa e instala dependências automaticamente
- ✅ **Ideal para CLIs**: Perfeito para instalar ferramentas de linha de comando

### Localização dos Binários

Por padrão, `go install` instala em:
- `$GOPATH/bin` (se `GOPATH` estiver configurado)
- `$GOBIN` (se `GOBIN` estiver configurado)
- `~/go/bin` (padrão moderno do Go)

### Exemplos Práticos

**Instalar uma ferramenta específica:**

```bash
# Instalar uma versão específica
go install golang.org/x/tools/cmd/godoc@latest

# Instalar uma versão exata
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
```

**Instalar o pacote atual:**

```bash
# Se você está em um diretório com package main
go install .

# O executável será criado em $GOPATH/bin
```

**Instalar múltiplos pacotes:**

```bash
go install ./cmd/...
```

**Verificar onde será instalado:**

```bash
go env GOBIN
go env GOPATH
```

### Diferença: `go build` vs `go install`

| `go build` | `go install` |
|------------|--------------|
| Cria binário no diretório atual | Cria binário em `$GOPATH/bin` |
| Não precisa de módulo Go | Requer módulo Go configurado |
| Para uso local | Para instalação global |

### Quando Usar?

- ✅ Para instalar ferramentas CLI globalmente
- ✅ Para instalar dependências de desenvolvimento
- ✅ Para compartilhar ferramentas entre projetos
- ✅ Para versionar ferramentas específicas

### Exemplo Completo

```bash
# Instalar uma ferramenta popular
go install github.com/cosmtrek/air@latest

# Verificar instalação
which air
# Output: /home/user/go/bin/air

# Usar a ferramenta
air
```

---

## 4. `go fmt` - Formatação Automática de Código

### O Que É?

O `go fmt` é o **formatador oficial** do Go. Ele formata automaticamente seu código de acordo com as convenções oficiais da linguagem, garantindo consistência e legibilidade.

### Sintaxe Básica

```bash
go fmt [pacote]
go fmt ./...
```

### Características Principais

- ✅ **Opinião única**: Não há configurações - Go tem um estilo oficial
- ✅ **Automático**: Formata indentação, espaçamento, alinhamento
- ✅ **Padrão da comunidade**: Elimina debates sobre estilo
- ✅ **Idempotente**: Executar múltiplas vezes produz o mesmo resultado

### O Que É Formatado?

- Indentação (tabs, não espaços)
- Espaçamento em operadores
- Alinhamento de declarações
- Quebras de linha
- Organização de imports

### Exemplos Práticos

**Formatar o pacote atual:**

```bash
go fmt
```

**Formatar recursivamente:**

```bash
go fmt ./...
```

**Formatar arquivo específico:**

```bash
go fmt main.go
```

**Ver o que seria formatado (sem modificar):**

```bash
# go fmt não tem flag dry-run, mas você pode usar:
gofmt -d main.go  # Mostra diff
gofmt -l .        # Lista arquivos que precisam formatação
```

### Integração com Editores

A maioria dos editores (VS Code, GoLand, Vim) formata automaticamente ao salvar usando `gofmt` ou `goimports`.

### Quando Usar?

- ✅ **Sempre antes de commitar código**
- ✅ Como parte do pipeline CI/CD
- ✅ Após escrever código novo
- ✅ Para padronizar código legado

### Boas Práticas

```bash
# Adicionar ao .git/hooks/pre-commit
#!/bin/sh
go fmt ./...
git add -u
```

### Exemplo Completo

```go
// Antes do go fmt (código mal formatado)
package main
import "fmt"
func main(){
x:=10
y:=20
fmt.Println(x+y)
}

// Depois do go fmt
package main

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println(x + y)
}
```

---

## 5. `go mod` - Gerenciamento de Módulos e Dependências

### O Que É?

O `go mod` é o **sistema de gerenciamento de módulos** do Go. Ele gerencia dependências, versões e módulos do seu projeto através dos arquivos `go.mod` e `go.sum`.

### Sintaxe Básica

```bash
go mod [comando] [argumentos]
```

### Comandos Principais

#### `go mod init` - Inicializar Módulo

Cria um novo módulo Go no diretório atual:

```bash
go mod init [nome-do-modulo]

# Exemplos
go mod init meu-projeto
go mod init github.com/usuario/projeto
```

**O que faz:**
- Cria arquivo `go.mod`
- Define o nome do módulo
- Estabelece a versão mínima do Go

#### `go mod tidy` - Limpar e Organizar

Remove dependências não utilizadas e adiciona as que faltam:

```bash
go mod tidy
```

**O que faz:**
- Adiciona dependências faltantes ao `go.mod`
- Remove dependências não utilizadas
- Atualiza `go.sum` com checksums

#### `go mod download` - Baixar Dependências

Baixa módulos para o cache local:

```bash
go mod download [módulo]@[versão]

# Exemplos
go mod download
go mod download github.com/gin-gonic/gin@latest
```

#### `go mod graph` - Visualizar Dependências

Mostra o grafo de dependências:

```bash
go mod graph
```

#### `go mod verify` - Verificar Integridade

Verifica se as dependências não foram modificadas:

```bash
go mod verify
```

#### `go mod why` - Explicar Dependência

Explica por que um módulo é necessário:

```bash
go mod why github.com/algum/pacote
```

### Arquivos Gerenciados

**`go.mod`**: Define o módulo e suas dependências
```go
module github.com/usuario/projeto

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
)
```

**`go.sum`**: Checksums criptográficos das dependências
```
github.com/gin-gonic/gin v1.9.1 h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=
github.com/gin-gonic/gin v1.9.1/go.mod h1:hPrL7YrpYKXt5YId3A/Tnip5kqbEAP+KLuI3SUcPTeU=
```

### Exemplos Práticos

**Inicializar novo projeto:**

```bash
mkdir meu-projeto
cd meu-projeto
go mod init github.com/meu-usuario/meu-projeto
```

**Adicionar dependência (automaticamente ao importar):**

```go
// main.go
import "github.com/gin-gonic/gin"

// Ao executar go build ou go run, a dependência é adicionada automaticamente
```

**Limpar dependências:**

```bash
go mod tidy
```

**Atualizar dependências:**

```bash
go get -u ./...           # Atualiza todas
go get github.com/pkg/errors@latest  # Atualiza específica
go mod tidy               # Limpa após atualizar
```

### Quando Usar?

- ✅ **Sempre em projetos novos**: `go mod init`
- ✅ **Antes de commitar**: `go mod tidy`
- ✅ **Para verificar dependências**: `go mod verify`
- ✅ **Para entender dependências**: `go mod graph` e `go mod why`

---

## 6. `go test` - Executar Testes

### O Que É?

O `go test` é o comando para **executar testes** em pacotes Go. Ele automaticamente encontra e executa funções que começam com `Test`, `Benchmark` ou `Example`.

### Sintaxe Básica

```bash
go test [pacote]
go test [flags] [pacote]
go test ./...
```

### Características Principais

- ✅ **Descoberta automática**: Encontra testes automaticamente
- ✅ **Múltiplos tipos**: Suporta testes, benchmarks e exemplos
- ✅ **Cobertura**: Gera relatórios de cobertura de código
- ✅ **Execução paralela**: Pode executar testes em paralelo
- ✅ **Verbose**: Múltiplos níveis de detalhamento

### Estrutura de Testes

```go
// arquivo: main_test.go
package main

import "testing"

// Teste unitário
func TestSoma(t *testing.T) {
    resultado := Soma(2, 3)
    if resultado != 5 {
        t.Errorf("Esperado 5, obteve %d", resultado)
    }
}

// Benchmark
func BenchmarkSoma(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Soma(2, 3)
    }
}

// Exemplo
func ExampleSoma() {
    resultado := Soma(2, 3)
    fmt.Println(resultado)
    // Output: 5
}
```

### Flags Importantes

```bash
# Verbose (mostra todos os testes)
go test -v

# Executar testes específicos
go test -run TestSoma

# Executar benchmarks
go test -bench=.

# Executar com cobertura
go test -cover

# Cobertura detalhada
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Executar testes em paralelo
go test -parallel 4

# Timeout
go test -timeout 30s

# Mostrar output mesmo em testes que passam
go test -v

# Executar apenas testes que falharam anteriormente
go test -failfast
```

### Exemplos Práticos

**Executar todos os testes:**

```bash
go test
```

**Executar com detalhes:**

```bash
go test -v
```

**Executar testes recursivamente:**

```bash
go test ./...
```

**Executar teste específico:**

```bash
go test -run TestSoma
```

**Executar benchmarks:**

```bash
go test -bench=. -benchmem
```

**Gerar relatório de cobertura:**

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Executar testes de integração (com tags):**

```bash
go test -tags=integration ./...
```

### Quando Usar?

- ✅ **Sempre antes de commitar código**
- ✅ Como parte do TDD (Test-Driven Development)
- ✅ Para garantir qualidade do código
- ✅ Para medir performance com benchmarks

### Exemplo Completo

```bash
# Estrutura
projeto/
├── main.go
├── main_test.go
└── utils_test.go

# Executar todos os testes
go test -v

# Output:
# === RUN   TestSoma
# --- PASS: TestSoma (0.00s)
# === RUN   TestMultiplicacao
# --- PASS: TestMultiplicacao (0.00s)
# PASS
# ok      projeto    0.123s
```

---

## 7. `go clean` - Limpar Arquivos de Build

### O Que É?

O `go clean` **remove arquivos** gerados durante o processo de compilação e build. É útil para limpar artefatos, liberar espaço e garantir builds limpos.

### Sintaxe Básica

```bash
go clean [flags] [pacotes]
```

### Flags Principais

```bash
# Limpar arquivos de build do pacote atual
go clean

# Limpar cache de build
go clean -cache

# Limpar cache de módulos
go clean -modcache

# Limpar arquivos de teste
go clean -testcache

# Limpar tudo (cache + módulos + testes)
go clean -cache -modcache -testcache

# Mostrar o que seria removido (dry-run)
go clean -n

# Remover também arquivos de instalação
go clean -i
```

### O Que É Removido?

Por padrão, `go clean` remove:
- Arquivos objeto (`.o`)
- Arquivos de teste compilados
- Binários locais

Com flags:
- `-cache`: Cache de build (`$GOCACHE`)
- `-modcache`: Cache de módulos baixados
- `-testcache`: Cache de resultados de testes

### Exemplos Práticos

**Limpar build local:**

```bash
go clean
```

**Limpar cache de build (útil para troubleshooting):**

```bash
go clean -cache
```

**Limpar tudo (útil para liberar espaço):**

```bash
go clean -cache -modcache -testcache
```

**Ver o que seria removido:**

```bash
go clean -n -cache
```

### Quando Usar?

- ✅ **Para troubleshooting**: Quando builds estão com problemas
- ✅ **Para liberar espaço**: Cache pode ocupar muito espaço
- ✅ **Para garantir builds limpos**: Antes de builds importantes
- ✅ **Em CI/CD**: Para garantir builds reproduzíveis

### Tamanho dos Caches

Os caches podem crescer bastante:
- **Build cache**: Geralmente alguns GB
- **Module cache**: Pode ser dezenas de GB em projetos grandes
- **Test cache**: Geralmente menor, mas pode crescer

### Exemplo Completo

```bash
# Verificar tamanho do cache
du -sh $(go env GOCACHE)
du -sh $(go env GOMODCACHE)

# Limpar se necessário
go clean -cache -modcache

# Verificar espaço liberado
du -sh $(go env GOCACHE)
```

---

## 8. `go doc` - Documentação Interativa

### O Que É?

O `go doc` **exibe documentação** extraída de comentários especiais no código Go. É uma ferramenta essencial para explorar APIs, verificar documentação e entender como usar pacotes.

### Sintaxe Básica

```bash
go doc [pacote]
go doc [pacote].[item]
go doc [pacote].[tipo].[método]
```

### Características Principais

- ✅ **Documentação automática**: Extrai de comentários no código
- ✅ **Navegação hierárquica**: Do pacote até métodos específicos
- ✅ **Código-fonte**: Pode mostrar código-fonte junto
- ✅ **Busca**: Suporta busca por padrões

### Flags Úteis

```bash
# Mostrar código-fonte
go doc -src fmt.Println

# Mostrar todos os símbolos exportados
go doc -all fmt

# Mostrar símbolos não exportados também
go doc -u fmt

# Formato JSON
go doc -json fmt
```

### Exemplos Práticos

**Ver documentação de um pacote:**

```bash
go doc fmt
```

**Ver documentação de uma função:**

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

**Ver código-fonte com documentação:**

```bash
go doc -src fmt.Println
```

**Buscar em todos os símbolos:**

```bash
go doc -all fmt | grep Print
```

**Ver documentação do seu próprio código:**

```bash
# No diretório do projeto
go doc .
go doc .MinhaFuncao
```

### Quando Usar?

- ✅ **Para explorar APIs**: Descobrir o que um pacote oferece
- ✅ **Para verificar documentação**: Confirmar como usar funções
- ✅ **Para aprender**: Estudar a biblioteca padrão
- ✅ **Para verificar sua própria documentação**: Testar se está correta

### Exemplo Completo

```bash
# Explorar pacote strings
go doc strings

# Ver função específica
go doc strings.Contains

# Ver código-fonte
go doc -src strings.Contains

# Output mostra:
# func Contains(s, substr string) bool
# Contains reports whether substr is within s.
```

---

## 9. `go version` - Informações da Versão

### O Que É?

O `go version` **exibe informações** sobre a versão do Go instalada, incluindo versão, sistema operacional alvo e arquitetura.

### Sintaxe Básica

```bash
go version
go version [arquivo]
```

### Características Principais

- ✅ **Versão do Go**: Mostra a versão exata instalada
- ✅ **Informações do sistema**: OS e arquitetura
- ✅ **Análise de binários**: Pode analisar binários compilados
- ✅ **Verificação de compatibilidade**: Útil para troubleshooting

### Exemplos Práticos

**Ver versão instalada:**

```bash
go version
# Output: go version go1.21.5 darwin/arm64
```

**Analisar um binário:**

```bash
go version ./meuapp
# Output: ./meuapp: go1.21.5
```

**Ver versão de múltiplos binários:**

```bash
go version bin1 bin2 bin3
```

### Informações Exibidas

O output típico inclui:
- **Versão do Go**: `go1.21.5`
- **Sistema operacional**: `darwin`, `linux`, `windows`
- **Arquitetura**: `amd64`, `arm64`, `386`

### Quando Usar?

- ✅ **Verificar instalação**: Confirmar que Go está instalado
- ✅ **Troubleshooting**: Verificar compatibilidade de versões
- ✅ **CI/CD**: Verificar versão em pipelines
- ✅ **Análise de binários**: Ver com qual versão foi compilado

### Exemplo Completo

```bash
# Verificar versão
go version
# Output: go version go1.21.5 darwin/arm64

# Verificar variáveis de ambiente relacionadas
go env GOVERSION
go env GOOS
go env GOARCH

# Analisar binário
go build -o app main.go
go version ./app
# Output: ./app: go1.21.5
```

---

## Resumo dos Comandos Core

| Comando | Propósito Principal | Quando Usar |
|---------|-------------------|-------------|
| `go run` | Compilar e executar | Desenvolvimento rápido, testes |
| `go build` | Compilar binários | Produção, distribuição, cross-compilation |
| `go install` | Instalar ferramentas | CLIs globais, ferramentas de desenvolvimento |
| `go fmt` | Formatar código | Antes de commitar, padronização |
| `go mod` | Gerenciar dependências | Inicializar projetos, gerenciar módulos |
| `go test` | Executar testes | TDD, qualidade, benchmarks |
| `go clean` | Limpar artefatos | Troubleshooting, liberar espaço |
| `go doc` | Ver documentação | Explorar APIs, aprender |
| `go version` | Ver versão | Verificação, troubleshooting |

---

## Workflow Típico de Desenvolvimento

Um workflow típico usando esses comandos:

```bash
# 1. Inicializar projeto
go mod init meu-projeto

# 2. Desenvolver código
# ... escrever código ...

# 3. Formatar
go fmt ./...

# 4. Testar durante desenvolvimento
go run main.go

# 5. Executar testes
go test -v

# 6. Verificar documentação
go doc .

# 7. Limpar antes de build final
go clean

# 8. Build para produção
go build -o meuapp .

# 9. Verificar versão
go version
```

---

## Conclusão

Dominar esses 9 comandos core do Go é fundamental para:

1. **Produtividade**: Trabalhar eficientemente no dia a dia
2. **Qualidade**: Manter código formatado e testado
3. **Colaboração**: Seguir padrões da comunidade
4. **Profissionalismo**: Usar ferramentas corretamente

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!



