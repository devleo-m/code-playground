# Módulo 33: Code Generation e Build Tags em Go
## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende `go generate` e Build Tags, vamos falar sobre **como usá-los de forma eficiente e profissional**. Esta é a parte que separa programadores iniciantes de programadores experientes.

---

## 1. go generate: O Que Deve e NÃO Deve Ser Feito

### ❌ O Que NÃO Deve Ser Feito

#### 1.1. Usar go generate para Tarefas Simples que Não Precisam

```go
// ❌ RUIM: Usar go generate para algo trivial
//go:generate echo "Hello World"

package main

func main() {
    // Isso poderia ser feito manualmente sem problema
}
```

**Por quê?**: `go generate` adiciona complexidade. Se você pode fazer manualmente em 2 minutos e só precisa fazer uma vez, não use `go generate`.

**Quando NÃO usar:**
- ✅ Tarefas que você faz raramente (menos de 3 vezes)
- ✅ Código que muda constantemente e precisa de ajustes manuais
- ✅ Tarefas simples que não economizam tempo significativo

#### 1.2. Não Documentar Diretivas go:generate

```go
// ❌ RUIM: Diretiva sem explicação
//go:generate stringer -type=Status

package main
```

**Por quê?**: Outros desenvolvedores (ou você no futuro) não vão entender o que isso faz ou como usar.

**✅ BOM: Documentar sempre**

```go
//go:generate stringer -type=Status
// Gera automaticamente o método String() para o tipo Status
// Execute: go install golang.org/x/tools/cmd/stringer@latest
// Depois: go generate ./...

package main
```

#### 1.3. Não Versionar Arquivos Gerados

```go
// ❌ RUIM: Ignorar arquivos gerados no Git
# .gitignore
*.pb.go
*_string.go
*_mock.go
```

**Por quê?**: 
- Builds não são reproduzíveis
- CI/CD precisa de ferramentas instaladas
- Code review fica difícil
- Novos desenvolvedores não conseguem compilar

**✅ BOM: Versionar arquivos gerados**

```go
# .gitignore
# Apenas arquivos temporários
*.tmp.go
*.swp

# Versionar arquivos gerados importantes
# *.pb.go  <- NÃO ignorar
# *_string.go  <- NÃO ignorar
```

**Exceção**: Arquivos muito grandes ou que mudam constantemente podem ser ignorados, mas documente o motivo.

#### 1.4. Executar go generate Manualmente Sem Verificação

```bash
# ❌ RUIM: Desenvolvedor esquece de rodar go generate
git commit -m "Add new status"
# Esqueceu de rodar go generate, código quebrado!
```

**✅ BOM: Automatizar no CI/CD**

```yaml
# .github/workflows/ci.yml
- name: Generate code
  run: go generate ./...

- name: Check for uncommitted changes
  run: |
    git diff --exit-code || (echo "Generated files are out of date. Run: go generate ./..." && exit 1)
```

#### 1.5. Usar go generate para Código que Precisa de Ajustes Manuais

```go
// ❌ RUIM: Gerar código que você sempre precisa editar depois
//go:generate template-generator -output=handlers.go
// Mas você sempre precisa ajustar o código gerado manualmente
```

**Por quê?**: Se você sempre precisa editar o código gerado, a geração automática não está ajudando.

**✅ BOM: Gerar código que não precisa de ajustes**

```go
//go:generate stringer -type=Status
// O código gerado é perfeito, não precisa ajustes
```

---

### ✅ O Que Deve Ser Feito

#### 1.1. Usar go generate para Código Repetitivo e Bem Definido

```go
// ✅ BOM: Gerar métodos String() para enums
//go:generate stringer -type=Status

// ✅ BOM: Compilar protobufs
//go:generate protoc --go_out=. user.proto

// ✅ BOM: Gerar mocks para testes
//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go
```

**Quando usar:**
- ✅ Código que segue um padrão bem definido
- ✅ Código que você precisa atualizar frequentemente
- ✅ Código que é gerado a partir de definições (protobuf, OpenAPI, etc.)
- ✅ Código que economiza tempo significativo

#### 1.2. Integrar go generate no Workflow

**Makefile:**
```makefile
.PHONY: generate test build

generate:
	@echo "Generating code..."
	go generate ./...

test: generate
	go test ./...

build: generate
	go build ./...
```

**Script de pre-commit:**
```bash
#!/bin/sh
# .git/hooks/pre-commit

echo "Running go generate..."
go generate ./...

git add -u  # Adiciona arquivos gerados modificados
```

#### 1.3. Documentar Requisitos e Dependências

```go
//go:generate stringer -type=Status
// 
// Requisitos:
//   - Instalar: go install golang.org/x/tools/cmd/stringer@latest
//   - Executar: go generate ./...
//   - Arquivo gerado: status_string.go (não editar manualmente)
//
// Atualizar quando:
//   - Adicionar novos valores ao enum Status
//   - Modificar valores existentes
```

#### 1.4. Testar que Arquivos Gerados Estão Atualizados

```go
//go:generate stringer -type=Status

package main

import "testing"

func TestGeneratedCodeIsUpToDate(t *testing.T) {
    // Teste que verifica se todos os valores têm String()
    statuses := []Status{
        StatusPending,
        StatusInProgress,
        StatusCompleted,
    }
    
    for _, s := range statuses {
        if s.String() == "" {
            t.Errorf("Status %v não tem método String() gerado. Execute: go generate ./...", s)
        }
    }
}
```

---

## 2. Build Tags: O Que Deve e NÃO Deve Ser Feito

### ❌ O Que NÃO Deve Ser Feito

#### 2.1. Usar Build Tags para Configurações que Mudam Frequentemente

```go
// ❌ RUIM: Build tag para algo que muda constantemente
//go:build production

package main

const API_URL = "https://api.production.com"
```

**Por quê?**: Se você precisa recompilar toda vez que muda a URL, build tags não são a solução certa.

**✅ BOM: Usar variáveis de ambiente ou arquivos de config**

```go
// ✅ BOM: Configuração em runtime
var API_URL = os.Getenv("API_URL")
if API_URL == "" {
    API_URL = "https://api.default.com"
}
```

**Quando NÃO usar build tags:**
- ❌ Configurações que mudam frequentemente (URLs, chaves, etc.)
- ❌ Feature flags que precisam ser ativadas/desativadas sem recompilar
- ❌ Dados que variam por ambiente (mas não por plataforma)

#### 2.2. Criar Build Tags Excessivamente Específicos

```go
// ❌ RUIM: Build tag muito específica
//go:build linux && amd64 && ubuntu && version20.04

package main
```

**Por quê?**: Muito específico, difícil de manter, pouco reutilizável.

**✅ BOM: Build tags genéricas e úteis**

```go
// ✅ BOM: Tags genéricas e úteis
//go:build linux

//go:build dev

//go:build integration
```

#### 2.3. Esquecer de Criar Arquivo Default

```go
// ❌ PROBLEMA: Sem arquivo default
// database_sqlite.go
//go:build sqlite

// database_postgres.go  
//go:build postgres

// Se nenhuma tag for especificada, nada compila!
```

**✅ BOM: Sempre ter um default**

```go
// database_sqlite.go
//go:build sqlite || (!postgres && !mysql)

// database_postgres.go
//go:build postgres

// database_default.go
//go:build !sqlite && !postgres && !mysql
// Default para SQLite se nenhuma tag for especificada
```

#### 2.4. Usar Build Tags para Lógica de Negócio

```go
// ❌ RUIM: Build tag para lógica de negócio
//go:build premium

package main

func CalculatePrice(items []Item) float64 {
    // Lógica diferente para premium
}
```

**Por quê?**: Lógica de negócio deve ser determinada em runtime, não em compile time.

**✅ BOM: Usar interfaces e polimorfismo**

```go
// ✅ BOM: Lógica de negócio em runtime
type PricingStrategy interface {
    CalculatePrice(items []Item) float64
}

type PremiumPricing struct{}
type StandardPricing struct{}

func NewPricingStrategy(isPremium bool) PricingStrategy {
    if isPremium {
        return &PremiumPricing{}
    }
    return &StandardPricing{}
}
```

#### 2.5. Não Testar Todas as Combinações de Build Tags

```bash
# ❌ RUIM: Só testar uma combinação
go test ./...
```

**✅ BOM: Testar todas as combinações relevantes**

```bash
# ✅ BOM: Testar todas as combinações
go test ./...                    # Sem tags
go test -tags=dev ./...          # Com tag dev
go test -tags=prod ./...         # Com tag prod
go test -tags=integration ./...  # Com tag integration
```

---

### ✅ O Que Deve Ser Feito

#### 2.1. Usar Build Tags para Código Específico de Plataforma

```go
// ✅ BOM: Código específico de OS
//go:build linux

package main

func getConfigPath() string {
    return "/etc/myapp/config.json"
}
```

**Quando usar:**
- ✅ Código que só funciona em uma plataforma
- ✅ Implementações diferentes para cada plataforma
- ✅ APIs específicas de sistema operacional

#### 2.2. Usar Build Tags para Feature Flags em Compile Time

```go
// ✅ BOM: Feature flag em compile time
//go:build dev

package main

const EnableDebugLogging = true
const EnableProfiling = true
```

**Quando usar:**
- ✅ Features que não devem estar no binário final (debug, profiling)
- ✅ Features experimentais que podem ser removidas
- ✅ Código de desenvolvimento que não deve ir para produção

#### 2.3. Usar Build Tags para Otimização

```go
// ✅ BOM: Build tag para otimização
//go:build !no_optimization

package main

// Código otimizado que pode ser desabilitado para debug
func optimizedFunction() {
    // Implementação otimizada
}
```

**Quando usar:**
- ✅ Código que melhora performance mas dificulta debug
- ✅ Implementações alternativas (uma rápida, outra legível)
- ✅ Código que pode ser removido para reduzir tamanho do binário

#### 2.4. Documentar Build Tags

```go
//go:build dev
// Este arquivo contém configurações e código de desenvolvimento.
// Compile com: go build -tags=dev
// 
// Inclui:
//   - Logging detalhado
//   - Profiling habilitado
//   - Informações de debug
//
// NÃO use em produção!

package main
```

#### 2.5. Usar Nomes Descritivos e Padronizados

```go
// ✅ BOM: Nomes claros e padronizados
//go:build dev
//go:build prod
//go:build test
//go:build integration
//go:build debug

// ❌ EVITE: Nomes confusos
//go:build x
//go:build flag1
//go:build mode2
```

---

## 3. Performance: Impacto de go generate e Build Tags

### 3.1. go generate: Zero Overhead em Runtime

**Importante**: `go generate` roda **antes** da compilação. Ele não afeta a performance do código em execução.

```go
//go:generate stringer -type=Status

// O código gerado é idêntico ao que você escreveria manualmente
// Performance: EXATAMENTE A MESMA
```

**Vantagens de performance:**
- ✅ Código gerado pode ser otimizado pela ferramenta
- ✅ Menos chance de erros que afetam performance
- ✅ Consistência garante código eficiente

**Desvantagens:**
- ❌ Nenhuma em runtime (go generate não roda em produção)

### 3.2. Build Tags: Redução de Tamanho e Overhead

**Build tags reduzem o tamanho do binário:**

```go
// Sem build tags: Binário inclui código de todas as plataformas
// Tamanho: ~10 MB

// Com build tags: Binário inclui apenas código da plataforma atual
// Tamanho: ~5 MB (50% menor!)
```

**Exemplo prático:**

```go
// features_dev.go
//go:build dev

package main

var debugInfo = map[string]interface{}{
    "version": "1.0.0",
    "build_time": "...",
    // Muitos outros dados de debug
}

// features_prod.go
//go:build !dev

package main

// Nada aqui - binário de produção não tem código de debug!
```

**Impacto:**
- ✅ Binários menores (menos memória)
- ✅ Menos código para carregar (startup mais rápido)
- ✅ Menos superfície de ataque (segurança)

### 3.3. Quando Build Tags Afetam Performance

**Cenário 1: Código Específico de Plataforma**

```go
// ❌ Sem build tags: Verificação em runtime
func getConfigPath() string {
    if runtime.GOOS == "windows" {
        return "C:\\..."
    } else if runtime.GOOS == "linux" {
        return "/etc/..."
    }
    // Overhead: Verificação a cada chamada
}

// ✅ Com build tags: Sem verificação
//go:build windows
func getConfigPath() string {
    return "C:\\..."  // Sem overhead!
}
```

**Cenário 2: Feature Flags**

```go
// ❌ Runtime check: Overhead mínimo mas constante
if os.Getenv("DEBUG") == "true" {
    log.Debug("...")
}

// ✅ Build tag: Zero overhead
//go:build dev
const EnableDebug = true
if EnableDebug {  // Compilador pode otimizar isso!
    log.Debug("...")
}
```

---

## 4. Boas Práticas para a Vida do Programador

### 4.1. Estratégia de Versionamento

**Arquivos gerados devem ser versionados quando:**
- ✅ São necessários para compilar o projeto
- ✅ Mudam raramente
- ✅ São pequenos ou médios
- ✅ Facilitam onboarding de novos desenvolvedores

**Arquivos gerados podem ser ignorados quando:**
- ✅ São muito grandes (ex: assets embedados)
- ✅ Mudam constantemente
- ✅ Podem ser gerados rapidamente
- ✅ Requerem ferramentas pesadas

**Regra de ouro**: Se um novo desenvolvedor consegue clonar e compilar sem rodar `go generate`, versionar. Caso contrário, documente bem como gerar.

### 4.2. Integração com CI/CD

**Sempre verificar que arquivos gerados estão atualizados:**

```yaml
# .github/workflows/ci.yml
- name: Generate code
  run: go generate ./...

- name: Check for uncommitted changes
  run: |
    git diff --exit-code || (echo "❌ Generated files are out of date!" && exit 1)
```

**Benefícios:**
- ✅ Detecta quando desenvolvedor esqueceu de rodar `go generate`
- ✅ Garante que código gerado está sempre atualizado
- ✅ Previne bugs sutis

### 4.3. Documentação

**Sempre documente:**
1. **O que cada diretiva faz**
2. **Como instalar dependências**
3. **Como executar**
4. **Quando atualizar**

```go
//go:generate stringer -type=Status
//
// Gera automaticamente o método String() para o tipo Status.
//
// Instalação:
//   go install golang.org/x/tools/cmd/stringer@latest
//
// Uso:
//   go generate ./...
//
// Arquivo gerado:
//   status_string.go (não editar manualmente)
//
// Atualizar quando:
//   - Adicionar novos valores ao enum Status
//   - Modificar valores existentes do enum
```

### 4.4. Testes com Build Tags

**Teste todas as combinações relevantes:**

```bash
#!/bin/bash
# test-all-tags.sh

echo "Testing without tags..."
go test ./...

echo "Testing with dev tag..."
go test -tags=dev ./...

echo "Testing with prod tag..."
go test -tags=prod ./...

echo "Testing with integration tag..."
go test -tags=integration ./...
```

**Inclua no Makefile:**

```makefile
.PHONY: test test-all

test:
	go test ./...

test-all: test
	go test -tags=dev ./...
	go test -tags=prod ./...
	go test -tags=integration ./...
```

---

## 5. Armadilhas Comuns e Como Evitá-las

### 5.1. Esquecer de Rodar go generate

**Problema**: Código gerado desatualizado causa bugs.

**Solução**: Automatizar e verificar no CI/CD.

### 5.2. Build Tags em Arquivos de Teste

**Problema**: Testes podem não rodar se você esquecer a tag.

```go
//go:build integration

package main

import "testing"

func TestIntegration(t *testing.T) {
    // Este teste só roda com -tags=integration
    // Se você esquecer, o teste não roda!
}
```

**Solução**: Documentar claramente e usar scripts/Makefile.

### 5.3. Múltiplas Implementações Sem Default

**Problema**: Nada compila se nenhuma tag for especificada.

**Solução**: Sempre ter um default ou documentar claramente qual tag é obrigatória.

### 5.4. Build Tags Confusos

**Problema**: Tags com nomes ruins causam confusão.

```go
// ❌ RUIM
//go:build x
//go:build flag1

// ✅ BOM
//go:build dev
//go:build prod
```

**Solução**: Usar nomes descritivos e padronizados.

---

## 6. Checklist de Revisão de Código

Antes de considerar seu código "pronto", pergunte-se:

**Para go generate:**
- [ ] Diretivas estão documentadas?
- [ ] Arquivos gerados estão versionados (ou há razão clara para não versionar)?
- [ ] CI/CD verifica que arquivos gerados estão atualizados?
- [ ] Requisitos e dependências estão documentados?
- [ ] go generate está integrado no workflow (Makefile, scripts)?

**Para Build Tags:**
- [ ] Build tags têm nomes descritivos?
- [ ] Há arquivo default quando necessário?
- [ ] Build tags são testadas em todas as combinações relevantes?
- [ ] Build tags são usadas para casos apropriados (não para configuração runtime)?
- [ ] Build tags estão documentadas?

**Geral:**
- [ ] Código funciona sem go generate (se arquivos gerados estão versionados)?
- [ ] Código funciona com diferentes combinações de build tags?
- [ ] Novos desenvolvedores conseguem compilar facilmente?
- [ ] Documentação está clara e completa?

---

## 7. Resumo: Regras de Ouro

### Para go generate:
1. **Use para código repetitivo e bem definido** - Não para tarefas simples ou únicas
2. **Sempre documente** - Explique o que, como e quando
3. **Versionar arquivos gerados** - A menos que haja razão clara para não fazer
4. **Automatizar verificação** - CI/CD deve verificar que arquivos estão atualizados
5. **Integrar no workflow** - Makefile, scripts, pre-commit hooks

### Para Build Tags:
1. **Use para código específico de plataforma** - Não para configurações runtime
2. **Use para feature flags em compile time** - Não para configuração que muda frequentemente
3. **Sempre tenha default** - Ou documente qual tag é obrigatória
4. **Teste todas as combinações** - Não assuma que funciona
5. **Use nomes descritivos** - `dev`, `prod`, `test` são melhores que `x`, `y`, `z`

### Performance:
1. **go generate**: Zero overhead em runtime (roda antes da compilação)
2. **Build Tags**: Reduzem tamanho do binário e overhead de código não usado
3. **Prefira build tags a runtime checks** quando possível (zero overhead)
4. **Meça antes de otimizar** - Build tags são úteis, mas não são solução mágica

---

## Conclusão

Dominar `go generate` e Build Tags vai além de saber a sintaxe. É sobre:

- **Escolher a ferramenta certa** para cada situação
- **Equilibrar automação e simplicidade**
- **Manter código gerado atualizado** e testado
- **Otimizar builds** sem sacrificar manutenibilidade
- **Facilitar vida de outros desenvolvedores** com documentação clara

Lembre-se: **Ferramentas são meios, não fins**. Use `go generate` e Build Tags quando realmente ajudam, não apenas porque são "legais".

**Priorize:**
1. **Clareza** sobre complexidade
2. **Manutenibilidade** sobre automação excessiva
3. **Documentação** sobre "código auto-explicativo"
4. **Testes** sobre "deve funcionar"

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!



