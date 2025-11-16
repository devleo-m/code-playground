# Módulo 31: Linters - Ferramentas Avançadas de Análise de Código
## Aula 4: Performance e Boas Práticas

Olá! Agora que você domina os conceitos e práticas dos linters, vamos mergulhar em **boas práticas** e **otimizações de performance** para usar essas ferramentas de forma eficiente e profissional em seus projetos.

---

## 1. Boas Práticas Gerais

### 1.1. Comece Simples, Evolua Gradualmente

**❌ Erro Comum**: Tentar configurar tudo de uma vez

```bash
# Não faça isso no primeiro dia
golangci-lint run --enable-all
```

**✅ Boa Prática**: Comece com o básico e adicione gradualmente

```bash
# Dia 1: Básico
go vet ./...
goimports -w .

# Semana 1: Adicione Revive
revive ./...

# Semana 2: Adicione Staticcheck
staticcheck ./...

# Mês 1: Considere Golangci-lint
golangci-lint run
```

**Por quê?**
- Evita sobrecarga de informações
- Permite aprender cada ferramenta adequadamente
- Facilita a adoção pela equipe

### 1.2. Configure Adequadamente

**❌ Erro Comum**: Usar configuração padrão sem ajustes

```yaml
# Configuração genérica demais
linters:
  enable-all: true
```

**✅ Boa Prática**: Configure baseado nas necessidades do projeto

```yaml
# Configuração específica para o projeto
linters:
  enable:
    - revive      # Estilo
    - staticcheck # Bugs
    - errcheck    # Erros não tratados
    - gosec       # Segurança (se relevante)
  disable:
    - golint      # Usar revive
    - dupl        # Não relevante para este projeto
```

**Por quê?**
- Reduz ruído de avisos irrelevantes
- Foca em problemas que realmente importam
- Melhora performance (menos linters = mais rápido)

### 1.3. Integre com Editor

**❌ Erro Comum**: Executar linters apenas manualmente

```bash
# Executar apenas quando lembrar
golangci-lint run
```

**✅ Boa Prática**: Configure para rodar automaticamente

```json
// VS Code settings.json
{
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "workspace",
  "editor.formatOnSave": true
}
```

**Por quê?**
- Detecta problemas em tempo real
- Corrige antes de commitar
- Melhora produtividade

### 1.4. Use em CI/CD

**❌ Erro Comum**: Verificar qualidade apenas localmente

```bash
# Apenas local, sem garantias
golangci-lint run
git push
```

**✅ Boa Prática**: Integre no pipeline CI/CD

```yaml
# .github/workflows/lint.yml
- name: golangci-lint
  uses: golangci/golangci-lint-action@v3
```

**Por quê?**
- Garante qualidade em todos os commits
- Previne código problemático no repositório
- Padroniza qualidade na equipe

---

## 2. Performance e Otimização

### 2.1. Use Cache

**❌ Erro Comum**: Executar linters sem cache

```bash
# Sem cache, sempre lento
golangci-lint run
```

**✅ Boa Prática**: Habilite cache

```bash
# Com cache (muito mais rápido)
golangci-lint run --cache
```

**Configuração no `.golangci.yml`:**

```yaml
run:
  cache: true
  cache-duration: 1h
```

**Ganho de Performance**: 5-10x mais rápido em execuções subsequentes

### 2.2. Execute Apenas em Arquivos Modificados

**❌ Erro Comum**: Analisar todo o projeto sempre

```bash
# Analisa tudo, mesmo sem mudanças
golangci-lint run ./...
```

**✅ Boa Prática**: Analise apenas arquivos modificados

```bash
# Apenas arquivos modificados (Git)
git diff --name-only | grep '\.go$' | xargs golangci-lint run
```

**Script útil:**

```bash
#!/bin/bash
# lint-changed.sh
changed_files=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
if [ -n "$changed_files" ]; then
    echo "$changed_files" | xargs golangci-lint run
fi
```

**Ganho de Performance**: 10-100x mais rápido dependendo do tamanho do projeto

### 2.3. Use Build Tags para Excluir Arquivos

**❌ Erro Comum**: Analisar arquivos gerados ou de terceiros

```yaml
# Analisa tudo, incluindo código gerado
issues:
  # Sem exclusões
```

**✅ Boa Prática**: Exclua arquivos irrelevantes

```yaml
issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck
    - path: _gen\.go
      linters:
        - all
    - path: vendor/
      linters:
        - all
```

**Ganho de Performance**: Reduz tempo de análise significativamente

### 2.4. Limite o Número de Problemas

**❌ Erro Comum**: Mostrar todos os problemas de uma vez

```yaml
issues:
  # Sem limites, pode ser esmagador
```

**✅ Boa Prática**: Defina limites razoáveis

```yaml
issues:
  max-issues-per-linter: 50
  max-same-issues: 3
  max-issues: 100
```

**Por quê?**
- Evita sobrecarga de informações
- Foca em problemas mais importantes primeiro
- Melhora experiência do desenvolvedor

### 2.5. Use Modo Fast do Golangci-lint

**❌ Erro Comum**: Executar todos os linters sempre

```bash
# Todos os linters, pode ser lento
golangci-lint run
```

**✅ Boa Prática**: Use modo fast para desenvolvimento

```bash
# Modo rápido (linters essenciais)
golangci-lint run --fast
```

**Configuração:**

```yaml
run:
  fast: true  # Apenas linters rápidos
```

**Linters rápidos incluídos:**
- errcheck
- gosec
- govet
- ineffassign
- staticcheck
- unused

**Ganho de Performance**: 2-3x mais rápido

---

## 3. Configuração por Tipo de Projeto

### 3.1. Projeto Pequeno/Pessoal

**Configuração Recomendada:**

```yaml
# .golangci.yml para projeto pequeno
linters:
  enable:
    - revive
    - staticcheck
    - errcheck
    - govet

run:
  timeout: 5m
  tests: true

issues:
  max-issues-per-linter: 20
```

**Comandos:**

```bash
# Makefile simples
lint:
	revive ./...
	staticcheck ./...
```

### 3.2. Projeto Médio/Equipe Pequena

**Configuração Recomendada:**

```yaml
# .golangci.yml para projeto médio
linters:
  enable:
    - revive
    - staticcheck
    - errcheck
    - gosec
    - govet
    - ineffassign
    - misspell
    - unused

run:
  timeout: 10m
  tests: true
  cache: true

issues:
  max-issues-per-linter: 50
  max-same-issues: 3
```

**Comandos:**

```bash
# Makefile completo
lint:
	golangci-lint run
```

### 3.3. Projeto Grande/Enterprise

**Configuração Recomendada:**

```yaml
# .golangci.yml para projeto grande
linters:
  enable:
    - revive
    - staticcheck
    - errcheck
    - gosec
    - govet
    - ineffassign
    - misspell
    - unused
    - gocritic
    - goconst
    - gocyclo
    - gofmt
    - goimports
    - golint
    - goprintffuncname
    - gosimple
    - govet
    - ineffassign
    - interfacer
    - lll
    - maligned
    - megacheck
    - nakedret
    - noctx
    - nolintlint
    - rowserrcheck
    - scopelint
    - structcheck
    - stylecheck
    - typecheck
    - unconvert
    - unparam
    - varcheck
    - whitespace

run:
  timeout: 15m
  tests: true
  cache: true
  cache-duration: 1h
  modules-download-mode: readonly

issues:
  max-issues-per-linter: 100
  max-same-issues: 5
  max-issues: 200
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck
        - gosec
    - path: _gen\.go
      linters:
        - all
```

**Comandos:**

```bash
# Makefile enterprise
lint:
	golangci-lint run --timeout=15m

lint-fast:
	golangci-lint run --fast

lint-ci:
	golangci-lint run --out-format=github-actions
```

---

## 4. Workflow Recomendado

### 4.1. Durante Desenvolvimento

**Workflow Local:**

```bash
# 1. Desenvolver código
# ... escrever código ...

# 2. Formatar automaticamente (editor faz isso)
# goimports roda ao salvar

# 3. Verificar rapidamente (modo fast)
golangci-lint run --fast

# 4. Se tudo OK, continuar desenvolvendo
# Se houver problemas, corrigir
```

**Configuração do Editor:**

```json
{
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "go.lintOnSave": "workspace",
  "editor.formatOnSave": true
}
```

### 4.2. Antes de Commitar

**Pre-commit Hook:**

```bash
#!/bin/sh
# .git/hooks/pre-commit

# Formatar
goimports -w .

# Verificar (modo completo)
golangci-lint run

if [ $? -ne 0 ]; then
    echo "❌ Linters encontraram problemas. Corrija antes de commitar."
    exit 1
fi

# Adicionar arquivos formatados
git add -u
```

### 4.3. No CI/CD

**GitHub Actions:**

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
          args: --timeout=10m
```

**Comentários em PR:**

```yaml
- name: golangci-lint
  uses: golangci/golangci-lint-action@v3
  with:
    version: latest
    github-token: ${{ secrets.GITHUB_TOKEN }}
    reporter: github-pr-review  # Comenta no PR
```

---

## 5. Tratando Falsos Positivos

### 5.1. Excluir Regras Específicas

**Problema**: Linter encontra problema que não é realmente um problema

```go
// Exemplo: Você sabe que esta função será usada no futuro
func futureFunction() {
    // ...
}
```

**Solução 1: Comentário nolint**

```go
//nolint:unused // Será usado no futuro
func futureFunction() {
    // ...
}
```

**Solução 2: Excluir no arquivo de configuração**

```yaml
issues:
  exclude-rules:
    - path: internal/future/
      linters:
        - unused
```

### 5.2. Ajustar Severidade

**Problema**: Aviso muito rigoroso para seu caso

```yaml
linters-settings:
  revive:
    rules:
      - name: exported
        severity: warning  # Ao invés de error
        disabled: false
```

### 5.3. Documentar Decisões

**Boa Prática**: Documente por que você desabilitou uma regra

```go
//nolint:unused // Esta função é chamada via reflection em runtime
func registerHandler() {
    // ...
}
```

---

## 6. Integração com Outras Ferramentas

### 6.1. Makefile Completo

```makefile
.PHONY: help format vet lint lint-fast lint-fix test build clean

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

format: ## Formata código e organiza imports
	goimports -w .

vet: ## Executa go vet
	go vet ./...

lint: ## Executa todos os linters (completo)
	golangci-lint run

lint-fast: ## Executa linters rápidos
	golangci-lint run --fast

lint-fix: ## Executa linters e tenta corrigir automaticamente
	golangci-lint run --fix

test: ## Executa testes
	go test -v ./...

build: ## Compila o projeto
	go build -o app .

clean: ## Limpa arquivos gerados
	rm -f app
	go clean

all: format vet lint-fast test ## Executa tudo (formato, vet, lint rápido, testes)
```

### 6.2. Script de Pre-commit Avançado

```bash
#!/bin/sh
# .git/hooks/pre-commit

set -e

echo "🔍 Executando linters..."

# Formatar código
echo "📝 Formatando código..."
goimports -w .

# Verificar com go vet
echo "🔎 Verificando com go vet..."
go vet ./...

# Verificar com linters (modo rápido para pre-commit)
echo "✨ Verificando com golangci-lint..."
golangci-lint run --fast

# Se tudo passar, adicionar arquivos formatados
git add -u

echo "✅ Todos os linters passaram!"
```

---

## 7. Métricas e Monitoramento

### 7.1. Acompanhar Métricas

**Script para gerar relatório:**

```bash
#!/bin/bash
# lint-report.sh

echo "=== Relatório de Linters ===" > lint-report.txt
echo "" >> lint-report.txt

echo "--- Revive ---" >> lint-report.txt
revive ./... >> lint-report.txt 2>&1
echo "" >> lint-report.txt

echo "--- Staticcheck ---" >> lint-report.txt
staticcheck ./... >> lint-report.txt 2>&1
echo "" >> lint-report.txt

echo "--- Golangci-lint ---" >> lint-report.txt
golangci-lint run >> lint-report.txt 2>&1

cat lint-report.txt
```

### 7.2. Integração com Ferramentas de Qualidade

**SonarQube:**

```yaml
# sonar-project.properties
sonar.go.golangci-lint.reportPaths=golangci-lint-report.json
```

**CodeClimate:**

```yaml
# .codeclimate.yml
engines:
  golangci-lint:
    enabled: true
```

---

## 8. Checklist de Boas Práticas

### ✅ Configuração
- [ ] Configuração adequada para o tamanho do projeto
- [ ] Cache habilitado
- [ ] Exclusões configuradas para arquivos gerados
- [ ] Limites de problemas definidos

### ✅ Integração
- [ ] Integrado com editor
- [ ] Pre-commit hooks configurados
- [ ] CI/CD pipeline configurado
- [ ] Makefile ou scripts de automação

### ✅ Workflow
- [ ] Linters rodam durante desenvolvimento
- [ ] Verificação antes de commitar
- [ ] Verificação automática no CI/CD
- [ ] Processo para tratar falsos positivos

### ✅ Equipe
- [ ] Documentação sobre configuração
- [ ] Padrões definidos e documentados
- [ ] Processo de revisão de código
- [ ] Treinamento da equipe

---

## Resumo das Boas Práticas

| Área | Boa Prática | Impacto |
|------|-------------|---------|
| **Configuração** | Configure baseado no projeto | Reduz ruído, melhora relevância |
| **Performance** | Use cache e modo fast | 5-10x mais rápido |
| **Integração** | Integre com editor e CI/CD | Detecta problemas cedo |
| **Workflow** | Verifique durante desenvolvimento | Corrige antes de commitar |
| **Equipe** | Documente e padronize | Consistência na equipe |

---

## Conclusão

Seguindo essas boas práticas, você vai:

1. ✅ **Usar linters de forma eficiente**: Performance otimizada
2. ✅ **Integrar no workflow**: Detecção automática de problemas
3. ✅ **Manter qualidade consistente**: Padrões em toda a equipe
4. ✅ **Economizar tempo**: Menos problemas em produção
5. ✅ **Melhorar produtividade**: Foco no desenvolvimento, não em correções

Linters são ferramentas poderosas, mas precisam ser usadas corretamente para maximizar seus benefícios. Configure adequadamente, integre no seu workflow e colha os frutos de código de alta qualidade!

---

## Próximos Passos

1. **Configure seu projeto**: Aplique as configurações recomendadas
2. **Integre com editor**: Configure para rodar automaticamente
3. **Configure CI/CD**: Garanta qualidade em todos os commits
4. **Experimente**: Ajuste configurações baseado na experiência
5. **Compartilhe**: Ajude a equipe a adotar as mesmas práticas

Boa sorte com seus projetos Go de alta qualidade! 🚀

