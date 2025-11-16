# Módulo 30: Code Quality and Analysis
## Aula 4: Performance, Boas Práticas e Vida Profissional

Olá! Agora vamos mergulhar nas **boas práticas profissionais** relacionadas a `go vet` e `goimports`. Esta aula é sobre como usar essas ferramentas de forma eficiente, quando usá-las, quando não usá-las, e como elas se integram na vida real de um programador Go.

---

## 🚀 Performance e Eficiência

### 1. Quando Executar `go vet`?

#### ✅ Sempre Execute (Obrigatório)

- **Antes de cada commit**: Detecta problemas antes que entrem no repositório
- **Durante o desenvolvimento**: Execute periodicamente enquanto escreve código
- **No CI/CD**: Configure para rodar automaticamente em cada pull request
- **Antes de fazer merge**: Garante que código novo não introduz problemas

#### ⚠️ Cuidado com o Timing

- **Não execute em loops**: `go vet` é rápido, mas executar milhares de vezes pode ser custoso
- **Use em arquivos específicos**: Se você modificou apenas um arquivo, pode executar apenas nele
- **Cache quando possível**: Alguns sistemas de CI/CD podem cachear resultados

#### 📊 Performance Real

```bash
# go vet é MUITO rápido
time go vet ./...
# Real: 0.05s  (projeto pequeno)
# Real: 0.5s   (projeto médio)
# Real: 2-5s   (projeto grande)

# Comparado com compilação:
time go build ./...
# Real: 2-10s  (depende do projeto)
```

**Conclusão**: `go vet` é tão rápido que não há desculpa para não executá-lo!

### 2. Quando Executar `goimports`?

#### ✅ Execute Frequentemente (Recomendado)

- **Ao salvar arquivos**: Configure no editor para executar automaticamente
- **Antes de commitar**: Garante que imports estão organizados
- **Após grandes refatorações**: Quando você move código entre arquivos
- **No CI/CD**: Como verificação de formatação

#### ⚠️ Performance do `goimports`

```bash
# goimports também é rápido
time goimports -w .
# Real: 0.1s   (projeto pequeno)
# Real: 1-2s   (projeto médio)
# Real: 5-10s  (projeto grande com muitos arquivos)
```

**Dica**: Execute apenas nos arquivos modificados para ser mais rápido:

```bash
# Apenas arquivos modificados (Git)
git diff --name-only | grep '\.go$' | xargs goimports -w
```

### 3. Otimizando o Workflow

#### Workflow Lento (Não Recomendado)

```bash
# ❌ Executar em todo o projeto toda vez
goimports -w ./...  # Pode ser lento em projetos grandes
go vet ./...        # Pode ser lento em projetos grandes
```

#### Workflow Rápido (Recomendado)

```bash
# ✅ Executar apenas no que mudou
# No editor: apenas no arquivo salvo (automático)
# Antes de commit: apenas arquivos modificados
git diff --cached --name-only | grep '\.go$' | xargs goimports -w
git diff --cached --name-only | grep '\.go$' | xargs go vet
```

---

## ✅ O Que DEVE Ser Feito

### 1. Integração com Editor (Obrigatório)

**Configure `goimports` para executar ao salvar**. Isso é essencial porque:

- ✅ Você não precisa pensar nisso
- ✅ Imports sempre estão corretos
- ✅ Não perde tempo gerenciando imports manualmente
- ✅ Consistência automática

**VS Code:**
```json
{
    "go.formatTool": "goimports",
    "editor.formatOnSave": true,
    "[go]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        }
    }
}
```

**GoLand:**
- Settings → Tools → Actions on Save
- Marque "Run goimports"

### 2. Hooks de Pre-commit (Altamente Recomendado)

Crie um hook Git que executa verificações antes de permitir commit:

```bash
#!/bin/sh
# .git/hooks/pre-commit

# Formatar código
goimports -w .

# Verificar problemas
go vet ./...

# Se go vet encontrar problemas, abortar
if [ $? -ne 0 ]; then
    echo "❌ go vet encontrou problemas. Corrija antes de commitar."
    exit 1
fi

# Adicionar arquivos formatados
git add -u
```

**Benefícios:**
- ✅ Impossível commitar código com problemas
- ✅ Código sempre formatado
- ✅ Imports sempre organizados
- ✅ Qualidade garantida

### 3. CI/CD Integration (Obrigatório em Equipes)

Configure seu pipeline CI/CD para executar:

```yaml
# .github/workflows/quality.yml
name: Code Quality

on: [push, pull_request]

jobs:
  quality:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: '1.21'
      
      - name: Install goimports
        run: go install golang.org/x/tools/cmd/goimports@latest
      
      - name: Check formatting
        run: |
          if [ "$(goimports -l . | wc -l)" -gt 0 ]; then
            echo "❌ Código não está formatado"
            goimports -d .
            exit 1
          fi
      
      - name: Run go vet
        run: go vet ./...
      
      - name: Run tests
        run: go test ./...
```

### 4. Makefile para Automação (Recomendado)

Crie um Makefile com comandos úteis:

```makefile
.PHONY: format vet test quality all

# Formatar código
format:
	@echo "📝 Formatando código..."
	goimports -w .

# Verificar qualidade
vet:
	@echo "🔍 Verificando qualidade..."
	go vet ./...

# Executar testes
test:
	@echo "🧪 Executando testes..."
	go test -v ./...

# Tudo junto
quality: format vet test
	@echo "✅ Qualidade verificada!"

# Build
build: quality
	@echo "🏗️  Compilando..."
	go build -o app .

# Limpar
clean:
	@echo "🧹 Limpando..."
	go clean
```

**Uso:**
```bash
make quality  # Executa tudo
make format   # Apenas formatar
make vet      # Apenas verificar
```

---

## ❌ O Que NÃO DEVE Ser Feito

### 1. Ignorar Avisos do `go vet`

**❌ NUNCA faça isso:**

```bash
go vet ./...
# Output: main.go:10:2: unreachable code

# ❌ ERRADO: Ignorar e commitar mesmo assim
git commit -m "feat: nova funcionalidade"
```

**✅ CORRETO:**
```bash
go vet ./...
# Output: main.go:10:2: unreachable code

# ✅ CORRETO: Corrigir o problema
# ... corrigir código ...

# Verificar novamente
go vet ./...

# Agora sim, commitar
git commit -m "feat: nova funcionalidade"
```

**Por quê?**
- Avisos do `go vet` geralmente indicam bugs reais
- Código inalcançável pode mascarar problemas
- Erros de tipo podem causar panics em runtime
- É melhor corrigir agora do que depois em produção

### 2. Executar `goimports` Manualmente Toda Vez

**❌ NÃO faça isso:**

```bash
# ❌ ERRADO: Executar manualmente toda vez que escreve código
# Escreve código...
goimports -w main.go
# Escreve mais código...
goimports -w main.go
# Escreve mais código...
goimports -w main.go
```

**✅ CORRETO:**
```bash
# ✅ CORRETO: Configurar no editor para executar automaticamente
# Você escreve código, salva, e goimports roda automaticamente
# Não precisa pensar nisso!
```

**Por quê?**
- Você vai esquecer de executar
- É trabalho repetitivo desnecessário
- Automação é mais confiável que memória humana

### 3. Desabilitar `go vet` nos Testes

**❌ NÃO faça isso:**

```bash
# ❌ ERRADO: Desabilitar go vet nos testes
go test -vet=off ./...
```

**✅ CORRETO:**
```bash
# ✅ CORRETO: Deixar go vet rodar automaticamente
go test ./...
```

**Por quê?**
- `go vet` é executado automaticamente com `go test` por um motivo
- É rápido e não atrapalha
- Detecta problemas que testes podem não pegar
- É parte do processo de qualidade

### 4. Commitar Sem Verificar

**❌ NUNCA faça isso:**

```bash
# ❌ ERRADO: Commitar sem verificar
git add .
git commit -m "feat: algo"
git push
```

**✅ CORRETO:**
```bash
# ✅ CORRETO: Sempre verificar antes
goimports -w .
go vet ./...
go test ./...

# Se tudo passar, commitar
git add .
git commit -m "feat: algo"
git push
```

---

## 🎯 Melhores Práticas para a Vida Profissional

### 1. Workflow Diário Recomendado

**Manhã (Início do Trabalho):**
```bash
# Atualizar código
git pull

# Verificar se está tudo OK
make quality  # ou: goimports -w . && go vet ./... && go test ./...
```

**Durante o Desenvolvimento:**
- Editor formata automaticamente ao salvar (`goimports`)
- `go vet` roda automaticamente com `go test`
- Você não precisa pensar nisso!

**Antes de Commitar:**
```bash
# Verificação final
goimports -w .
go vet ./...
go test ./...

# Se tudo passar, commitar
git add .
git commit -m "feat: descrição"
```

### 2. Trabalhando em Equipe

**Padrões da Equipe:**
- ✅ Todos usam `goimports` configurado no editor
- ✅ Todos executam `go vet` antes de commitar
- ✅ CI/CD verifica automaticamente
- ✅ Pull requests são rejeitados se falharem nas verificações

**Comunicação:**
- Se alguém commitar código com problemas, seja educado mas firme
- Compartilhe conhecimento sobre essas ferramentas
- Documente o workflow da equipe

### 3. Projetos Legados

**Quando você herda código antigo:**

```bash
# 1. Formatar tudo de uma vez
goimports -w ./...

# 2. Verificar problemas
go vet ./...

# 3. Corrigir problemas críticos primeiro
# 4. Criar issues para problemas menores
# 5. Configurar CI/CD para prevenir novos problemas
```

**Estratégia:**
- Não tente corrigir tudo de uma vez
- Priorize problemas que podem causar bugs
- Configure ferramentas para prevenir novos problemas
- Corrija gradualmente durante refatorações

### 4. Performance em Projetos Grandes

**Projetos com milhares de arquivos:**

```bash
# ❌ Lento: Executar em tudo
goimports -w ./...  # Pode levar minutos

# ✅ Rápido: Apenas arquivos modificados
git diff --name-only | grep '\.go$' | xargs goimports -w
git diff --name-only | grep '\.go$' | xargs go vet
```

**Ou use ferramentas especializadas:**
```bash
# Usar golangci-lint que é mais eficiente em projetos grandes
golangci-lint run
```

---

## 🔧 Resolvendo Problemas Comuns

### Problema 1: `go vet` Encontra Muitos Problemas

**Situação:** Você executou `go vet` e encontrou 50 problemas.

**Solução:**
1. Não entre em pânico
2. Priorize problemas críticos (bugs reais)
3. Crie uma lista de tarefas
4. Corrija gradualmente
5. Configure CI/CD para prevenir novos problemas

### Problema 2: `goimports` Adiciona Imports Errados

**Situação:** `goimports` adiciona imports que você não quer.

**Solução:**
```go
// Use comentários especiais para controlar imports
import (
    _ "github.com/pacote/indesejado"  // Força import mas não usa
)
```

**Ou remova manualmente e use:**
```bash
goimports -format-only main.go  # Apenas formata, não adiciona/remove
```

### Problema 3: Conflitos de Formatação

**Situação:** Você e seu colega têm formatação diferente.

**Solução:**
- Ambos devem usar `goimports` (não `gofmt`)
- Configure no editor para executar automaticamente
- Use o mesmo formato (padrão do Go)
- Não há discussão: Go tem um formato oficial

---

## 📊 Métricas de Qualidade

### Como Medir Melhoria?

**Antes de usar as ferramentas:**
- Quantos bugs são encontrados em produção?
- Quanto tempo é gasto corrigindo imports?
- Quantos PRs são rejeitados por formatação?

**Depois de usar as ferramentas:**
- Bugs em produção: ↓
- Tempo com imports: ↓ (automatizado)
- PRs rejeitados: ↓ (verificação automática)

### KPIs Sugeridos

1. **Taxa de sucesso do `go vet`**: % de commits que passam
2. **Tempo de revisão de PR**: Deve diminuir
3. **Bugs em produção**: Deve diminuir
4. **Consistência de código**: Deve aumentar

---

## 🎓 Conclusão: A Vida Profissional

### Resumo das Boas Práticas

1. **Automatize tudo possível**
   - `goimports` no editor (ao salvar)
   - `go vet` no CI/CD
   - Hooks de pre-commit

2. **Nunca ignore avisos**
   - `go vet` geralmente está certo
   - Corrija problemas antes de commitar
   - Qualidade > Velocidade

3. **Integre no workflow**
   - Faça parte do processo natural
   - Não seja uma etapa extra
   - Seja automático e invisível

4. **Trabalhe em equipe**
   - Padrões compartilhados
   - CI/CD garante consistência
   - Comunicação sobre qualidade

### O Caminho para a Excelência

```
Nível 1: Usuário Iniciante
├── Executa go vet manualmente
├── Executa goimports manualmente
└── Esquece às vezes

Nível 2: Usuário Intermediário
├── Configura goimports no editor
├── Executa go vet antes de commitar
└── Lembra da maioria das vezes

Nível 3: Usuário Avançado
├── Tudo automatizado no editor
├── Hooks de pre-commit configurados
├── CI/CD verifica automaticamente
└── Qualidade é parte natural do processo

Nível 4: Líder Técnico
├── Garante que toda equipe usa as ferramentas
├── Documenta padrões e workflows
├── Monitora métricas de qualidade
└── Melhora continuamente o processo
```

**Qual nível você quer alcançar?** 🚀

---

## 📚 Recursos Adicionais

- **Documentação oficial do `go vet`: https://pkg.go.dev/cmd/vet
- **Documentação do `goimports`: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- **Go Code Review Comments: https://github.com/golang/go/wiki/CodeReviewComments
- **Effective Go: https://go.dev/doc/effective_go

---

**Lembre-se**: Qualidade de código não é opcional, é essencial. Essas ferramentas não são "nice to have", são **necessárias** para desenvolvimento profissional em Go. Use-as sempre, automatize-as, e torne-as parte natural do seu workflow. Seu futuro eu (e sua equipe) agradecerão! 🎯

