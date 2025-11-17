# Módulo 31: Linters - Ferramentas Avançadas de Análise de Código
## Aula 3: Exercícios e Reflexão

Olá! Agora que você entendeu os conceitos de linters, é hora de praticar! Vamos trabalhar com exercícios práticos que vão ajudar você a fixar o aprendizado e aplicar os conhecimentos em situações reais.

---

## Exercício 1: Instalação e Primeiros Passos

### Objetivo
Instalar e testar os três linters principais.

### Tarefas

1. **Instalar Revive**
   ```bash
   go install github.com/mgechev/revive@latest
   revive -version
   ```

2. **Instalar Staticcheck**
   ```bash
   go install honnef.co/go/tools/cmd/staticcheck@latest
   staticcheck -version
   ```

3. **Instalar Golangci-lint**
   ```bash
   # macOS
   brew install golangci-lint
   
   # Ou via script
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
   
   golangci-lint --version
   ```

### Verificação
Execute cada comando e confirme que todas as ferramentas foram instaladas corretamente.

---

## Exercício 2: Analisando Código com Problemas

### Objetivo
Criar código intencionalmente com problemas e ver como cada linter os detecta.

### Código com Problemas

Crie um arquivo `exercicio-02.go` com o seguinte código:

```go
package main

import (
    "fmt"
    "os"
    "time"
)

// Função exportada sem comentário
func CalculateTotal(items []int) int {
    total := 0
    for _, item := range items {
        total += item
    }
    return total
}

// Função não utilizada
func unusedFunction() {
    fmt.Println("Nunca chamada")
}

func main() {
    // Variável não utilizada
    x := 10
    
    // Erro não verificado
    os.Open("arquivo.txt")
    
    // Uso incorreto de time.Sleep
    time.Sleep(100)
    
    // Variável com nome não convencional
    var User_Name string = "João"
    
    // Comparação desnecessária
    if true == true {
        fmt.Println("Hello")
    }
    
    items := []int{1, 2, 3}
    total := CalculateTotal(items)
    fmt.Println(total)
}
```

### Tarefas

1. **Executar Revive**
   ```bash
   revive exercicio-02.go
   ```
   - Anote quantos problemas foram encontrados
   - Quais tipos de problemas o Revive detectou?

2. **Executar Staticcheck**
   ```bash
   staticcheck exercicio-02.go
   ```
   - Anote quantos problemas foram encontrados
   - Quais tipos de problemas o Staticcheck detectou?

3. **Executar Golangci-lint**
   ```bash
   golangci-lint run exercicio-02.go
   ```
   - Compare com os resultados anteriores
   - O Golangci-lint encontrou mais ou menos problemas?

### Reflexão
- Qual linter encontrou mais problemas?
- Quais problemas foram encontrados por múltiplos linters?
- Qual linter você achou mais útil para este código?

---

## Exercício 3: Corrigindo Problemas

### Objetivo
Corrigir todos os problemas encontrados pelos linters.

### Tarefas

1. **Corrigir o código do Exercício 2**
   - Adicione comentários onde necessário
   - Remova código não utilizado
   - Corrija nomenclaturas
   - Trate erros adequadamente
   - Corrija uso de APIs

2. **Verificar novamente com os linters**
   ```bash
   revive exercicio-02.go
   staticcheck exercicio-02.go
   golangci-lint run exercicio-02.go
   ```

3. **Código corrigido esperado**
   - Todos os linters devem passar sem erros (ou com mínimos avisos)

### Código Corrigido (Referência)

```go
package main

import (
    "fmt"
    "log"
    "time"
)

// CalculateTotal calcula a soma total dos itens fornecidos
func CalculateTotal(items []int) int {
    total := 0
    for _, item := range items {
        total += item
    }
    return total
}

func main() {
    // Abrir arquivo e tratar erro
    file, err := os.Open("arquivo.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    // Uso correto de time.Sleep
    time.Sleep(100 * time.Millisecond)
    
    // Variável com nome convencional
    var userName string = "João"
    
    // Comparação simplificada
    if true {
        fmt.Println("Hello")
    }
    
    items := []int{1, 2, 3}
    total := CalculateTotal(items)
    fmt.Println(total)
}
```

---

## Exercício 4: Configurando Revive

### Objetivo
Criar um arquivo de configuração personalizado para o Revive.

### Tarefas

1. **Criar arquivo `revive.toml`**
   ```toml
   ignoreGeneratedHeader = false
   severity = "warning"
   confidence = 0.8
   errorCode = 0
   warningCode = 0

   [rule.exported]
   [rule.var-naming]
   [rule.package-comments]
   [rule.error-return]
   ```

2. **Testar configuração**
   ```bash
   revive -config revive.toml exercicio-02.go
   ```

3. **Experimentar**
   - Desabilite algumas regras e veja a diferença
   - Ajuste a severidade e confiança

### Reflexão
- Como a configuração personalizada mudou os resultados?
- Quais regras você acha mais importantes para seu projeto?

---

## Exercício 5: Configurando Golangci-lint

### Objetivo
Criar um arquivo de configuração para o Golangci-lint.

### Tarefas

1. **Criar arquivo `.golangci.yml`**
   ```yaml
   linters-settings:
     revive:
       rules:
         - name: exported
           severity: warning
         - name: var-naming
           severity: warning
   
     staticcheck:
       checks: ["all"]
   
   linters:
     enable:
       - revive
       - staticcheck
       - errcheck
       - govet
     disable:
       - golint
   
   issues:
     max-issues-per-linter: 50
     max-same-issues: 3
   ```

2. **Testar configuração**
   ```bash
   golangci-lint run
   ```

3. **Experimentar**
   - Adicione mais linters
   - Ajuste limites de problemas
   - Configure exclusões

### Reflexão
- Como a configuração unificada facilita o trabalho?
- Quais linters você habilitaria para um projeto real?

---

## Exercício 6: Integração com Editor

### Objetivo
Configurar linters para rodar automaticamente no seu editor.

### Tarefas

#### VS Code

1. **Instalar extensão Go**
   - Instale a extensão oficial do Go

2. **Configurar `settings.json`**
   ```json
   {
     "go.lintTool": "golangci-lint",
     "go.lintFlags": ["--fast"],
     "editor.formatOnSave": true,
     "go.formatTool": "goimports"
   }
   ```

3. **Testar**
   - Abra um arquivo Go com problemas
   - Veja os avisos aparecerem automaticamente
   - Salve o arquivo e veja a formatação automática

#### GoLand

1. **Configurar linters**
   - Vá em: `Settings > Tools > Go > Linters`
   - Habilite `golangci-lint`
   - Configure para rodar ao salvar

### Reflexão
- Como a integração com o editor melhora sua produtividade?
- Você prefere ver avisos em tempo real ou apenas ao executar manualmente?

---

## Exercício 7: Criando um Makefile

### Objetivo
Criar um Makefile para automatizar a execução de linters.

### Tarefas

1. **Criar `Makefile`**
   ```makefile
   .PHONY: format vet lint test build

   format:
   	goimports -w .

   vet:
   	go vet ./...

   lint-revive:
   	revive ./...

   lint-staticcheck:
   	staticcheck ./...

   lint-golangci:
   	golangci-lint run

   lint: lint-revive lint-staticcheck

   test:
   	go test ./...

   build:
   	go build -o app .

   all: format vet lint test build
   ```

2. **Testar comandos**
   ```bash
   make format
   make vet
   make lint
   make all
   ```

### Reflexão
- Como o Makefile facilita o workflow?
- Quais outros comandos você adicionaria?

---

## Exercício 8: Projeto Completo

### Objetivo
Aplicar todos os conhecimentos em um projeto completo.

### Tarefas

1. **Criar estrutura de projeto**
   ```
   projeto/
   ├── cmd/
   │   └── server/
   │       └── main.go
   ├── internal/
   │   └── handlers/
   │       └── handler.go
   ├── pkg/
   │   └── utils/
   │       └── utils.go
   ├── .golangci.yml
   ├── revive.toml
   └── Makefile
   ```

2. **Criar código com problemas intencionais**
   - Adicione problemas de estilo
   - Adicione código não utilizado
   - Adicione erros não tratados

3. **Configurar linters**
   - Configure Revive
   - Configure Golangci-lint
   - Configure Makefile

4. **Executar e corrigir**
   ```bash
   make all
   ```

5. **Iterar até passar**
   - Corrija todos os problemas
   - Execute novamente
   - Repita até todos os linters passarem

### Reflexão
- Quanto tempo levou para corrigir todos os problemas?
- Quais problemas foram mais difíceis de corrigir?
- Como os linters ajudaram a melhorar a qualidade do código?

---

## Exercício 9: Comparando Linters

### Objetivo
Entender as diferenças entre os linters através de comparação prática.

### Tarefas

1. **Criar código de teste**
   - Crie um arquivo com vários tipos de problemas

2. **Executar cada linter separadamente**
   ```bash
   revive arquivo.go > revive-output.txt
   staticcheck arquivo.go > staticcheck-output.txt
   golangci-lint run arquivo.go > golangci-output.txt
   ```

3. **Comparar resultados**
   - Quais problemas foram encontrados por todos?
   - Quais problemas foram encontrados apenas por um linter?
   - Qual linter encontrou mais problemas?

4. **Criar tabela comparativa**
   | Problema | Revive | Staticcheck | Golangci-lint |
   |----------|--------|-------------|---------------|
   | Estilo   | ✅     | ❌          | ✅            |
   | Bugs     | ❌     | ✅          | ✅            |
   | ...      | ...    | ...         | ...           |

### Reflexão
- Qual linter você usaria para cada tipo de problema?
- Por que usar Golangci-lint ao invés de executar cada linter separadamente?

---

## Exercício 10: Integração com CI/CD

### Objetivo
Configurar linters para rodar automaticamente em CI/CD.

### Tarefas

#### GitHub Actions

1. **Criar `.github/workflows/lint.yml`**
   ```yaml
   name: Lint
   on: [push, pull_request]
   jobs:
     golangci:
       runs-on: ubuntu-latest
       steps:
         - uses: actions/checkout@v3
         - name: golangci-lint
           uses: golangci/golangci-lint-action@v3
           with:
             version: latest
   ```

2. **Testar localmente** (usando act ou similar)
   ```bash
   act -j golangci
   ```

#### GitLab CI

1. **Criar `.gitlab-ci.yml`**
   ```yaml
   lint:
     image: golangci/golangci-lint:latest
     script:
       - golangci-lint run ./...
   ```

### Reflexão
- Como a integração com CI/CD garante qualidade?
- Quais são os benefícios de rodar linters automaticamente?

---

## Questões para Reflexão

### 1. Qualidade vs Velocidade
- Você prefere corrigir todos os avisos ou apenas os críticos?
- Como equilibrar qualidade com velocidade de desenvolvimento?

### 2. Configuração
- Você prefere configuração padrão ou personalizada?
- Quais regras você desabilitaria e por quê?

### 3. Workflow
- Como você integraria linters no seu workflow diário?
- Você prefere ver avisos em tempo real ou apenas antes de commitar?

### 4. Equipe
- Como garantir que todos na equipe usem os mesmos linters?
- Como lidar com conflitos de estilo entre membros da equipe?

### 5. Projetos Legados
- Como aplicar linters em projetos legados com muitos problemas?
- Você corrigiria tudo de uma vez ou gradualmente?

---

## Desafios Extras

### Desafio 1: Criar Script de Pre-commit
Crie um script que:
- Formata código automaticamente
- Executa linters
- Bloqueia commit se houver problemas críticos

### Desafio 2: Configuração Avançada
Crie uma configuração que:
- Habilita apenas linters relevantes
- Configura exclusões para testes
- Define limites de problemas aceitáveis

### Desafio 3: Integração Completa
Configure:
- Editor (VS Code/GoLand)
- Pre-commit hooks
- CI/CD pipeline
- Makefile

---

## Conclusão

Através destes exercícios, você:

1. ✅ Instalou e testou os três linters principais
2. ✅ Entendeu como cada linter detecta problemas diferentes
3. ✅ Aprendeu a corrigir problemas encontrados
4. ✅ Configurou linters para seus projetos
5. ✅ Integrou linters com editores e CI/CD
6. ✅ Criou workflows automatizados

Continue praticando e experimentando com diferentes configurações. Os linters são ferramentas poderosas que vão melhorar significativamente a qualidade do seu código!

Na próxima aula, vamos ver boas práticas e dicas de performance para usar linters de forma eficiente.


