# Módulo 29: Comandos Core do Go
## Aula 4: Performance e Boas Práticas - Dominando os Comandos Core

Olá! Agora que você conhece os comandos core do Go, é crucial entender **como usá-los de forma eficiente** e **quais são as melhores práticas** para cada situação. Esta aula vai te transformar de um usuário básico para um desenvolvedor que domina as ferramentas.

---

## 🚀 Performance: Quando Usar Cada Comando

### `go run` vs `go build`: Escolhendo o Comando Certo

#### ❌ **NUNCA use `go run` em produção**

**Por quê?**
- `go run` compila toda vez que executa (overhead desnecessário)
- Não cria binário otimizado
- Mais lento que executar um binário pré-compilado
- Consome mais recursos (CPU e memória)

**✅ Use `go build` para produção:**
```bash
# Desenvolvimento (rápido, para testar)
go run main.go

# Produção (otimizado, rápido)
go build -o app main.go
./app
```

#### ⚡ Performance: Build Cache

O Go mantém um **cache de build** que acelera compilações subsequentes:

```bash
# Primeira compilação (mais lenta)
go build main.go
# Tempo: ~2-3 segundos

# Segunda compilação (usa cache, muito mais rápida)
go build main.go
# Tempo: ~0.1-0.5 segundos
```

**Dica**: Não limpe o cache (`go clean -cache`) a menos que seja necessário. O cache acelera significativamente o desenvolvimento.

---

## 📦 `go build`: Otimizações e Flags Importantes

### Flags de Performance

```bash
# Build otimizado (padrão, mas explícito)
go build -ldflags="-s -w" main.go
# -s: Remove tabela de símbolos
# -w: Remove informações de debug
# Resultado: Binário ~20-30% menor

# Build com informações de versão
go build -ldflags="-X main.Version=1.0.0 -X main.BuildTime=$(date)" main.go

# Build sem otimizações (apenas para debug)
go build -gcflags="-N -l" main.go
# -N: Desabilita otimizações
# -l: Desabilita inlining
```

### Cross-Compilation: Performance e Tamanho

**Diferentes arquiteturas produzem binários de tamanhos diferentes:**

```bash
# Comparar tamanhos
GOOS=linux GOARCH=amd64 go build -o app-linux main.go
GOOS=windows GOARCH=amd64 go build -o app-windows.exe main.go
GOOS=darwin GOARCH=arm64 go build -o app-macos main.go

# Verificar tamanhos
ls -lh app-*
```

**Observação**: Binários para ARM geralmente são menores que x86_64.

---

## 🧪 `go test`: Performance e Eficiência

### Executando Testes de Forma Eficiente

#### ✅ **DO**: Executar testes em paralelo quando possível

```bash
# Testes podem rodar em paralelo (padrão)
go test -parallel 4 ./...

# Para testes que não podem ser paralelos, use t.Parallel() com cuidado
```

#### ✅ **DO**: Usar cache de testes

```bash
# Go cacheia resultados de testes que não mudaram
go test ./...
# Primeira execução: ~5 segundos
# Segunda execução (sem mudanças): ~0.1 segundos (usa cache)
```

#### ❌ **DON'T**: Executar todos os testes sempre

```bash
# ❌ RUIM: Executa TODOS os testes sempre
go test ./...

# ✅ BOM: Executa apenas testes do pacote que mudou
go test ./pacote-especifico

# ✅ MELHOR: Use ferramentas como `gotestsum` para executar apenas testes afetados
```

### Benchmarks: Medindo Performance Real

```bash
# Executar benchmarks com estatísticas
go test -bench=. -benchmem -benchtime=3s

# Comparar benchmarks entre versões
go test -bench=. > antes.txt
# ... fazer mudanças ...
go test -bench=. > depois.txt
benchcmp antes.txt depois.txt
```

**Dica**: Use `-benchmem` para ver alocações de memória, crucial para identificar vazamentos.

---

## 🧹 `go clean`: Quando e Como Limpar

### ❌ **DON'T**: Limpar cache desnecessariamente

```bash
# ❌ RUIM: Limpar cache toda vez
go clean -cache
go build  # Recompila tudo do zero, mais lento

# ✅ BOM: Limpar apenas quando necessário
# - Builds estranhos
# - Mudanças na versão do Go
# - Troubleshooting
```

### ✅ **DO**: Limpar antes de builds importantes

```bash
# Antes de build de release
go clean
go build -ldflags="-s -w" -o release/app .
```

### Gerenciando Espaço em Disco

**Cache pode crescer bastante:**

```bash
# Verificar tamanho
du -sh $(go env GOCACHE)      # Geralmente 1-5 GB
du -sh $(go env GOMODCACHE)   # Pode ser 10-50 GB+

# Limpar se necessário (mas não faça isso frequentemente!)
go clean -modcache  # Remove TODOS os módulos baixados
```

**Estratégia**: Limpe `-modcache` apenas quando realmente precisar de espaço. O cache acelera muito o desenvolvimento.

---

## 📚 `go mod`: Boas Práticas de Gerenciamento

### ✅ **DO**: Executar `go mod tidy` regularmente

```bash
# Antes de cada commit
go mod tidy
git add go.mod go.sum
git commit
```

**Por quê?**
- Remove dependências não usadas
- Adiciona dependências faltantes
- Atualiza `go.sum` com checksums corretos
- Mantém o projeto limpo

### ❌ **DON'T**: Editar `go.mod` manualmente (geralmente)

```bash
# ❌ RUIM: Editar go.mod diretamente
# Pode causar inconsistências

# ✅ BOM: Deixar Go gerenciar
go get github.com/algum/pacote@latest
go mod tidy
```

### Versionamento de Dependências

```bash
# ✅ BOM: Especificar versões exatas para produção
go get github.com/gin-gonic/gin@v1.9.1

# ⚠️ CUIDADO: Usar @latest em produção
go get github.com/gin-gonic/gin@latest
# Pode quebrar seu código se houver breaking changes

# ✅ MELHOR: Usar @latest apenas em desenvolvimento
# Em produção, fixe versões específicas
```

---

## ✂️ `go fmt`: Integração no Workflow

### ✅ **DO**: Automatizar formatação

**Opção 1: Git Hooks**
```bash
# .git/hooks/pre-commit
#!/bin/sh
go fmt ./...
git add -u
```

**Opção 2: Editor (VS Code)**
```json
// settings.json
{
    "editor.formatOnSave": true,
    "go.formatTool": "goimports"
}
```

**Opção 3: CI/CD**
```yaml
# .github/workflows/ci.yml
- name: Format check
  run: |
    go fmt ./...
    git diff --exit-code
```

### ❌ **DON'T**: Formatar código manualmente

```bash
# ❌ RUIM: Tentar formatar manualmente
# Você pode não seguir o padrão exato

# ✅ BOM: Sempre usar go fmt
go fmt ./...
```

---

## 🛠️ `go install`: Gerenciando Ferramentas

### ✅ **DO**: Versionar ferramentas de desenvolvimento

```bash
# ✅ BOM: Instalar versões específicas
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

# ❌ RUIM: Sempre usar @latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# Pode quebrar seu workflow se houver mudanças
```

### Organizando Ferramentas

```bash
# Criar um arquivo tools.go para versionar ferramentas
// tools.go
//go:build tools
// +build tools

package tools

import (
    _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
    _ "golang.org/x/tools/cmd/godoc"
)
```

**Vantagem**: As ferramentas ficam versionadas junto com o projeto.

---

## 🔍 `go doc`: Explorando Eficientemente

### ✅ **DO**: Usar `go doc` para aprender

```bash
# Explorar pacotes novos
go doc strings
go doc -all strings | grep -i "contains\|split\|join"

# Entender APIs rapidamente
go doc fmt.Printf
go doc -src fmt.Printf  # Ver implementação
```

### Integração com Desenvolvimento

**Workflow eficiente:**
1. Encontre a função que precisa: `go doc -all pacote | grep funcao`
2. Veja como usar: `go doc pacote.Funcao`
3. Veja exemplos: `go doc -src pacote.Funcao`
4. Implemente no seu código

---

## 📊 Workflow Otimizado: Do Desenvolvimento à Produção

### Desenvolvimento Diário

```bash
# 1. Iniciar trabalho
git pull
go mod tidy  # Atualizar dependências

# 2. Desenvolver
# ... escrever código ...
go run main.go  # Testar rapidamente

# 3. Formatar
go fmt ./...

# 4. Testar
go test -v ./pacote-atual

# 5. Verificar documentação (se necessário)
go doc pacote.Funcao
```

### Antes de Commit

```bash
# Checklist obrigatório
go fmt ./...           # ✅ Formatar
go mod tidy            # ✅ Limpar dependências
go test ./...          # ✅ Testes passando
go vet ./...           # ✅ Análise estática (se disponível)
go build ./...         # ✅ Compila sem erros
```

### Build de Produção

```bash
# 1. Limpar
go clean

# 2. Verificar dependências
go mod verify

# 3. Testar tudo
go test ./...

# 4. Build otimizado
go build -ldflags="-s -w" -o app .

# 5. Verificar binário
go version ./app
./app --version  # Se implementado
```

### CI/CD Pipeline

```bash
# Exemplo de pipeline
#!/bin/bash
set -e

# Formatação
go fmt ./...
git diff --exit-code || (echo "Código não formatado!" && exit 1)

# Dependências
go mod download
go mod verify
go mod tidy
git diff --exit-code || (echo "go.mod/go.sum desatualizados!" && exit 1)

# Testes
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# Build
go build -o app .

# Lint (se disponível)
# golangci-lint run
```

---

## ⚠️ Erros Comuns e Como Evitá-los

### Erro 1: Usar `go run` em Produção

```bash
# ❌ ERRADO
# systemd service ou dockerfile
CMD ["go", "run", "main.go"]

# ✅ CORRETO
go build -o app .
CMD ["./app"]
```

### Erro 2: Não Executar `go mod tidy`

**Sintoma**: `go.mod` e `go.sum` desatualizados, builds inconsistentes.

**Solução**: Sempre executar `go mod tidy` antes de commit.

### Erro 3: Limpar Cache Demais

**Sintoma**: Builds muito lentos.

**Solução**: Limpar cache apenas quando necessário (troubleshooting).

### Erro 4: Não Versionar Ferramentas

**Sintoma**: Diferentes desenvolvedores com versões diferentes de ferramentas.

**Solução**: Versionar ferramentas no `go.mod` ou documentar versões.

### Erro 5: Ignorar `go fmt` no CI/CD

**Sintoma**: Código inconsistente no repositório.

**Solução**: Adicionar verificação de formatação no pipeline.

---

## 🎯 Melhores Práticas por Comando

### `go run`
- ✅ Use apenas para desenvolvimento rápido
- ✅ Para testar ideias e protótipos
- ❌ Nunca em produção
- ❌ Nunca para medir performance real

### `go build`
- ✅ Sempre para produção
- ✅ Use flags de otimização (`-ldflags="-s -w"`)
- ✅ Cross-compile quando necessário
- ✅ Verifique o binário com `go version`

### `go install`
- ✅ Versionar ferramentas
- ✅ Documentar versões usadas
- ✅ Usar para CLIs globais
- ❌ Não para dependências de aplicação (use `go get`)

### `go fmt`
- ✅ Automatizar (hooks, editor, CI/CD)
- ✅ Executar antes de cada commit
- ✅ Verificar no CI/CD
- ❌ Nunca formatar manualmente

### `go mod`
- ✅ `go mod tidy` antes de cada commit
- ✅ `go mod verify` em CI/CD
- ✅ Versionar dependências em produção
- ❌ Não editar `go.mod` manualmente
- ❌ Não commitar sem `go mod tidy`

### `go test`
- ✅ Executar antes de commit
- ✅ Usar `-cover` para monitorar cobertura
- ✅ Benchmarks para performance
- ✅ Cache de testes (não limpar desnecessariamente)
- ❌ Não ignorar testes que falham

### `go clean`
- ✅ Apenas quando necessário (troubleshooting)
- ✅ Antes de builds de release importantes
- ✅ Para liberar espaço (modcache)
- ❌ Não fazer parte do workflow diário

### `go doc`
- ✅ Usar para aprender novas APIs
- ✅ Verificar documentação própria
- ✅ Explorar biblioteca padrão
- ✅ Integrar no workflow de aprendizado

### `go version`
- ✅ Verificar em CI/CD
- ✅ Troubleshooting de compatibilidade
- ✅ Documentar versão requerida
- ✅ Verificar binários compilados

---

## 📈 Métricas e Monitoramento

### Medindo Performance de Build

```bash
# Tempo de build
time go build main.go

# Comparar builds
time go build -ldflags="-s -w" main.go
time go build main.go
```

### Monitorando Cobertura de Testes

```bash
# Gerar relatório
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Meta: Manter acima de 80% de cobertura
```

### Tamanho de Binários

```bash
# Monitorar tamanho
go build -o app main.go
ls -lh app

# Comparar com otimizações
go build -ldflags="-s -w" -o app-opt main.go
ls -lh app app-opt
```

---

## 🎓 Conclusão: Tornando-se um Mestre dos Comandos

Dominar os comandos core do Go não é apenas saber o que cada um faz, mas entender:

1. **Quando usar cada comando**: Contexto é tudo
2. **Como otimizar**: Performance importa
3. **Boas práticas**: Workflow profissional
4. **Evitar erros comuns**: Aprender com os erros dos outros
5. **Automatizar**: Integrar no workflow

### Checklist de Mestre

Um desenvolvedor Go experiente:

- ✅ Sabe quando usar `go run` vs `go build`
- ✅ Executa `go fmt` e `go mod tidy` automaticamente
- ✅ Tem testes com boa cobertura
- ✅ Usa `go build` otimizado para produção
- ✅ Versiona dependências e ferramentas
- ✅ Tem CI/CD configurado corretamente
- ✅ Sabe quando limpar cache (e quando não limpar)
- ✅ Usa `go doc` para aprender eficientemente

---

## 🚀 Próximos Passos

Agora que você domina os comandos core:

1. **Pratique**: Use esses comandos diariamente
2. **Automatize**: Configure hooks e CI/CD
3. **Monitore**: Acompanhe métricas de build e testes
4. **Compartilhe**: Ensine outros desenvolvedores
5. **Evolua**: Explore comandos avançados (`go vet`, `go generate`, etc.)

Lembre-se: **Ferramentas são meios, não fins**. O objetivo é escrever código Go de qualidade, e esses comandos são suas ferramentas para alcançar isso!

