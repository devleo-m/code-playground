# Módulo 32: Security - Segurança em Aplicações Go
## Aula 1: govulncheck - Scanner de Vulnerabilidades Oficial do Go

Olá! Bem-vindo a este módulo essencial sobre **segurança** em Go. Até agora você aprendeu a escrever código funcional e de qualidade, mas uma parte fundamental do desenvolvimento profissional é garantir que seu código e suas dependências estejam **livres de vulnerabilidades conhecidas** que podem ser exploradas por atacantes.

Nesta aula, vamos mergulhar no **govulncheck**, o scanner de vulnerabilidades oficial do Go desenvolvido pela equipe do Go. Esta ferramenta é essencial para manter suas aplicações seguras, verificando tanto seu código quanto suas dependências contra vulnerabilidades conhecidas no banco de dados oficial do Go.

---

## 1. govulncheck - O Scanner de Vulnerabilidades Oficial

### O Que É?

O **govulncheck** é uma ferramenta oficial do Go que verifica seu código e dependências em busca de vulnerabilidades de segurança conhecidas. Ele utiliza o **Go Vulnerability Database** (banco de dados de vulnerabilidades do Go), que é mantido pela equipe do Go e pela comunidade, para identificar pacotes com vulnerabilidades conhecidas e fornecer informações sobre severidade e como corrigi-las.

### Características Principais

- ✅ **Oficial**: Desenvolvido e mantido pela equipe do Go
- ✅ **Banco de dados oficial**: Usa o Go Vulnerability Database
- ✅ **Verifica código e dependências**: Analisa tanto seu código quanto as dependências
- ✅ **Informações de severidade**: Classifica vulnerabilidades por nível de risco
- ✅ **Conselhos de correção**: Fornece orientações sobre como corrigir problemas
- ✅ **Análise estática**: Não precisa executar o código para encontrar vulnerabilidades
- ✅ **Integração fácil**: Funciona com qualquer projeto Go

### Por Que Usar govulncheck?

A segurança é uma preocupação crítica em desenvolvimento de software. Vulnerabilidades podem:

- **Expor dados sensíveis**: Informações de usuários podem ser comprometidas
- **Permitir acesso não autorizado**: Atacantes podem ganhar controle do sistema
- **Causar negação de serviço**: Aplicações podem ser derrubadas
- **Comprometer integridade**: Dados podem ser modificados ou corrompidos

O govulncheck ajuda a:

- **Detectar vulnerabilidades conhecidas**: Identifica problemas já documentados
- **Priorizar correções**: Informa a severidade de cada vulnerabilidade
- **Manter dependências seguras**: Verifica todas as dependências do projeto
- **Cumprir compliance**: Ajuda a atender requisitos de segurança

### Instalação

```bash
# Instalar govulncheck
go install golang.org/x/vuln/cmd/govulncheck@latest

# Verificar instalação
govulncheck -version
```

### Sintaxe Básica

```bash
# Verificar o pacote atual
govulncheck ./...

# Verificar um pacote específico
govulncheck ./cmd/server

# Verificar um módulo específico
govulncheck -mode=binary ./cmd/myapp

# Verificar apenas dependências (não o código)
govulncheck -mode=mod ./...

# Verificar binário compilado
govulncheck -mode=binary ./myapp

# Formato JSON (útil para CI/CD)
govulncheck -json ./...

# Mostrar apenas vulnerabilidades de alta severidade
govulncheck -severity=high ./...
```

### Modos de Operação

O govulncheck oferece três modos de operação:

#### 1. Modo Source (Padrão)
Analisa o código-fonte para encontrar vulnerabilidades que afetam o código que você realmente usa.

```bash
govulncheck ./...
```

**Vantagens:**
- Mais preciso: mostra apenas vulnerabilidades que afetam seu código
- Mais rápido: não precisa compilar
- Menos falsos positivos: ignora vulnerabilidades em código não utilizado

#### 2. Modo Binary
Analisa um binário compilado para encontrar vulnerabilidades.

```bash
govulncheck -mode=binary ./myapp
```

**Vantagens:**
- Verifica o que realmente está no binário
- Útil para verificar binários de terceiros
- Não precisa do código-fonte

#### 3. Modo Module
Analisa apenas as dependências do módulo, sem verificar o código-fonte.

```bash
govulncheck -mode=mod ./...
```

**Vantagens:**
- Mais rápido: não analisa código-fonte
- Útil para verificar apenas dependências
- Bom para auditorias rápidas

---

## 2. Entendendo Vulnerabilidades

### O Que São Vulnerabilidades?

Vulnerabilidades são falhas de segurança em software que podem ser exploradas por atacantes para causar danos. Em Go, vulnerabilidades podem estar em:

1. **Biblioteca padrão**: Bugs na biblioteca padrão do Go
2. **Dependências de terceiros**: Pacotes externos com vulnerabilidades
3. **Seu próprio código**: Problemas de segurança no código que você escreveu

### Tipos Comuns de Vulnerabilidades

#### 1. Vulnerabilidades de Entrada (Input Validation)
Código que não valida adequadamente entradas do usuário:

```go
// ❌ Vulnerável: não valida entrada
func processUserInput(input string) {
    // Processa sem validação
    exec.Command(input) // Pode executar comandos maliciosos
}
```

#### 2. Vulnerabilidades de Injeção
Permitem que atacantes injetem código ou comandos:

```go
// ❌ Vulnerável: SQL injection
func queryUser(name string) {
    query := "SELECT * FROM users WHERE name = '" + name + "'"
    // Atacante pode injetar SQL malicioso
}
```

#### 3. Vulnerabilidades de Exposição de Dados
Expõem informações sensíveis:

```go
// ❌ Vulnerável: expõe informações sensíveis em logs
func logError(err error) {
    log.Printf("Erro: %v", err) // Pode conter dados sensíveis
}
```

#### 4. Vulnerabilidades de Autenticação/Autorização
Problemas com controle de acesso:

```go
// ❌ Vulnerável: não verifica permissões
func deleteUser(userID int) {
    // Deleta sem verificar se usuário tem permissão
    db.Delete("users", userID)
}
```

### Severidade de Vulnerabilidades

O govulncheck classifica vulnerabilidades por severidade:

- **CRITICAL**: Vulnerabilidades críticas que podem causar danos graves
- **HIGH**: Vulnerabilidades de alta severidade que precisam atenção imediata
- **MEDIUM**: Vulnerabilidades médias que devem ser corrigidas
- **LOW**: Vulnerabilidades baixas que podem ser corrigidas quando possível

---

## 3. Usando govulncheck na Prática

### Exemplo 1: Verificando um Projeto Simples

Vamos criar um exemplo simples para demonstrar o govulncheck:

```go
// main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    http.ListenAndServe(":8080", nil)
}
```

Execute o govulncheck:

```bash
govulncheck ./...
```

**Saída esperada:**
```
govulncheck: scanning for known vulnerabilities...

No vulnerabilities found.
```

### Exemplo 2: Projeto com Dependências

Vamos criar um projeto que usa dependências externas:

```go
// go.mod
module exemplo

go 1.21

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/lib/pq v1.10.0
)
```

```bash
# Verificar vulnerabilidades
govulncheck ./...
```

O govulncheck verificará todas as dependências e mostrará vulnerabilidades conhecidas.

### Exemplo 3: Interpretando Resultados

Quando o govulncheck encontra vulnerabilidades, a saída pode ser assim:

```
govulncheck: scanning for known vulnerabilities...

Vulnerability #1: GO-2023-1234
  Package: golang.org/x/crypto
  Version: v0.1.0
  Severity: HIGH
  Description: Buffer overflow in crypto/rand
  
  Your code imports:
    - exemplo/cmd/server (uses crypto/rand)
  
  Recommendation: Update to v0.2.0 or later
```

### Exemplo 4: Formato JSON

Para integração com CI/CD, use formato JSON:

```bash
govulncheck -json ./... > vulnerabilities.json
```

```json
{
  "vulnerabilities": [
    {
      "id": "GO-2023-1234",
      "package": "golang.org/x/crypto",
      "version": "v0.1.0",
      "severity": "HIGH",
      "description": "Buffer overflow in crypto/rand",
      "recommendation": "Update to v0.2.0 or later"
    }
  ]
}
```

---

## 4. Trabalhando com Vulnerabilidades Encontradas

### Passo 1: Entender a Vulnerabilidade

Quando o govulncheck encontra uma vulnerabilidade, leia cuidadosamente:

1. **ID da vulnerabilidade**: Identificador único (ex: GO-2023-1234)
2. **Pacote afetado**: Qual pacote tem a vulnerabilidade
3. **Versão afetada**: Qual versão do pacote está vulnerável
4. **Severidade**: Quão crítica é a vulnerabilidade
5. **Descrição**: O que a vulnerabilidade permite
6. **Recomendação**: Como corrigir

### Passo 2: Verificar Se Você Está Usando o Código Vulnerável

O govulncheck mostra se você realmente usa o código vulnerável:

```go
// Se você importa mas não usa:
import "golang.org/x/crypto" // ❌ Vulnerável, mas não usado

// Se você usa:
import "golang.org/x/crypto/rand" // ✅ Vulnerável E usado
rand.Read(buffer) // Código vulnerável está sendo usado
```

### Passo 3: Atualizar Dependências

A maneira mais comum de corrigir é atualizar a dependência:

```bash
# Atualizar para versão segura
go get -u golang.org/x/crypto@latest

# Ou versão específica
go get golang.org/x/crypto@v0.2.0

# Verificar novamente
govulncheck ./...
```

### Passo 4: Alternativas Se Não Puder Atualizar

Se não puder atualizar imediatamente:

1. **Remover código vulnerável**: Se possível, remova o uso do código vulnerável
2. **Workaround temporário**: Implemente uma solução temporária
3. **Monitorar**: Acompanhe atualizações do pacote
4. **Documentar**: Documente o risco e o plano de correção

---

## 5. Integração com Workflow de Desenvolvimento

### Integração com go.mod

O govulncheck funciona automaticamente com `go.mod`:

```bash
# Verificar dependências do módulo
govulncheck -mode=mod ./...
```

### Integração com CI/CD

#### GitHub Actions

```yaml
name: Security Scan
on: [push, pull_request]
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Install govulncheck
        run: go install golang.org/x/vuln/cmd/govulncheck@latest
      - name: Run govulncheck
        run: govulncheck ./...
```

#### GitLab CI

```yaml
security-scan:
  image: golang:1.21
  before_script:
    - go install golang.org/x/vuln/cmd/govulncheck@latest
  script:
    - govulncheck ./...
```

### Script de Pre-commit

```bash
#!/bin/sh
# .git/hooks/pre-commit

echo "Executando govulncheck..."
govulncheck ./...

if [ $? -ne 0 ]; then
    echo "❌ govulncheck encontrou vulnerabilidades!"
    exit 1
fi

echo "✅ Nenhuma vulnerabilidade encontrada"
```

### Makefile

```makefile
.PHONY: security test build

# Verificar segurança
security:
	govulncheck ./...

# Testar
test:
	go test ./...

# Build
build:
	go build -o app .

# Tudo junto
all: security test build
```

---

## 6. Boas Práticas com govulncheck

### 1. Execute Regularmente

**❌ Erro Comum**: Verificar apenas antes de releases

```bash
# Não faça isso
# Verificar apenas uma vez por mês
```

**✅ Boa Prática**: Verificar em cada commit ou pull request

```bash
# Integre no workflow diário
govulncheck ./...
```

### 2. Verifique Dependências Novas

**❌ Erro Comum**: Adicionar dependências sem verificar

```bash
go get github.com/algum/pacote
# Esqueceu de verificar vulnerabilidades
```

**✅ Boa Prática**: Verificar após adicionar dependências

```bash
go get github.com/algum/pacote
govulncheck ./...  # Verificar imediatamente
```

### 3. Use no CI/CD

**❌ Erro Comum**: Verificar apenas localmente

```bash
# Apenas no seu computador
govulncheck ./...
```

**✅ Boa Prática**: Integrar no pipeline de CI/CD

```yaml
# GitHub Actions, GitLab CI, etc.
- name: Security Scan
  run: govulncheck ./...
```

### 4. Priorize por Severidade

**❌ Erro Comum**: Tratar todas as vulnerabilidades igualmente

```bash
# Corrigir tudo de uma vez (pode ser impossível)
```

**✅ Boa Prática**: Priorizar por severidade

```bash
# Corrigir CRITICAL e HIGH primeiro
# Depois MEDIUM e LOW
```

### 5. Documente Decisões

**❌ Erro Comum**: Ignorar vulnerabilidades sem documentar

```go
// Vulnerabilidade conhecida, mas ignorada
// (sem documentação do porquê)
```

**✅ Boa Prática**: Documentar decisões de segurança

```go
// NOTA DE SEGURANÇA: Esta dependência tem uma vulnerabilidade
// conhecida (GO-2023-1234), mas não podemos atualizar devido a
// incompatibilidades. Monitorando atualizações.
// Data: 2024-01-15
// Responsável: Equipe de Segurança
```

---

## 7. Comparação com Outras Ferramentas

### govulncheck vs Outras Ferramentas

| Ferramenta | Foco | Banco de Dados | Oficial |
|------------|------|----------------|---------|
| **govulncheck** | Vulnerabilidades Go | Go Vulnerability DB | ✅ Sim |
| **gosec** | Análise de segurança | Próprio | ❌ Não |
| **nancy** | Vulnerabilidades de dependências | OWASP | ❌ Não |
| **snyk** | Vulnerabilidades gerais | Próprio | ❌ Não |

### Por Que Escolher govulncheck?

- ✅ **Oficial**: Mantido pela equipe do Go
- ✅ **Específico para Go**: Focado em vulnerabilidades Go
- ✅ **Banco de dados oficial**: Usa fonte confiável
- ✅ **Integração nativa**: Funciona perfeitamente com Go modules
- ✅ **Precisão**: Mostra apenas vulnerabilidades que você realmente usa

### Usando Junto com Outras Ferramentas

Você pode usar govulncheck junto com outras ferramentas:

```bash
# Verificar vulnerabilidades conhecidas
govulncheck ./...

# Verificar problemas de segurança no código
gosec ./...

# Verificar dependências
nancy sleuth
```

---

## 8. Exemplos Práticos Completos

### Exemplo 1: Projeto Web com Dependências

```go
// go.mod
module meuapp

go 1.21

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/lib/pq v1.10.0
    golang.org/x/crypto v0.1.0
)
```

```bash
# Verificar vulnerabilidades
govulncheck ./...

# Se encontrar vulnerabilidades, atualizar:
go get -u golang.org/x/crypto@latest
go mod tidy
govulncheck ./...  # Verificar novamente
```

### Exemplo 2: Verificando Binário Compilado

```bash
# Compilar aplicação
go build -o myapp ./cmd/server

# Verificar binário
govulncheck -mode=binary ./myapp
```

### Exemplo 3: Verificando Apenas Dependências

```bash
# Verificar apenas dependências (não código-fonte)
govulncheck -mode=mod ./...
```

### Exemplo 4: Integração com Scripts

```bash
#!/bin/bash
# check-security.sh

echo "🔍 Verificando vulnerabilidades..."

if govulncheck ./...; then
    echo "✅ Nenhuma vulnerabilidade encontrada"
    exit 0
else
    echo "❌ Vulnerabilidades encontradas!"
    exit 1
fi
```

---

## Resumo dos Conceitos

| Conceito | Descrição |
|----------|-----------|
| **govulncheck** | Scanner oficial de vulnerabilidades do Go |
| **Go Vulnerability DB** | Banco de dados oficial de vulnerabilidades |
| **Modo Source** | Analisa código-fonte (padrão) |
| **Modo Binary** | Analisa binário compilado |
| **Modo Module** | Analisa apenas dependências |
| **Severidade** | Classificação de risco (CRITICAL, HIGH, MEDIUM, LOW) |

---

## Conclusão

Dominar o govulncheck é essencial para:

1. **Segurança Profissional**: Manter aplicações livres de vulnerabilidades conhecidas
2. **Compliance**: Atender requisitos de segurança e auditoria
3. **Confiança**: Garantir que dependências são seguras
4. **Produtividade**: Detectar problemas antes de produção
5. **Padrão da Indústria**: Usar ferramenta oficial recomendada

O govulncheck é uma ferramenta poderosa e essencial para qualquer desenvolvedor Go sério sobre segurança. Integre-o no seu workflow diário e mantenha suas aplicações seguras!

Na próxima parte desta aula, vamos simplificar esses conceitos com analogias do dia a dia para fixar melhor o aprendizado!

