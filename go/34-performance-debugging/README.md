# Módulo 34: Performance e Debugging em Go

Bem-vindo ao módulo sobre **Performance e Debugging** em Go! Este módulo ensina as ferramentas essenciais para analisar, otimizar e debugar programas Go.

---

## 📚 Estrutura do Módulo

Este módulo segue o **Ciclo de 4 Etapas**:

```
┌─────────────────────────────────────────────────────────┐
│  1. AULA PRINCIPAL                                      │
│     └─ Conteúdo técnico completo sobre pprof, trace    │
│        e Race Detector                                  │
├─────────────────────────────────────────────────────────┤
│  2. AULA SIMPLIFICADA                                   │
│     └─ Mesmos conceitos com analogias do cotidiano      │
├─────────────────────────────────────────────────────────┤
│  3. EXERCÍCIOS E REFLEXÃO                               │
│     └─ Práticas + perguntas que exigem pensamento      │
├─────────────────────────────────────────────────────────┤
│  4. PERFORMANCE E BOAS PRÁTICAS                         │
│     └─ O que fazer e não fazer, otimizações           │
└─────────────────────────────────────────────────────────┘
```

---

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você será capaz de:

- ✅ Usar `pprof` para analisar performance (CPU, memória, goroutines)
- ✅ Usar `trace` para analisar execução e concorrência
- ✅ Usar Race Detector para encontrar data races
- ✅ Interpretar resultados de profiling
- ✅ Otimizar código baseado em dados reais
- ✅ Integrar ferramentas de debugging no workflow

---

## 📖 Conteúdo das Aulas

### Aula 1: pprof, trace e Race Detector - Ferramentas Essenciais
**Arquivo**: `aula-01-performance-debugging-principal.md`

Conteúdo técnico completo sobre:
- **pprof**: Profiling de CPU, memória, goroutines
- **trace**: Análise de execução e concorrência
- **Race Detector**: Detecção de data races

### Aula 2: Entendendo Performance e Debugging
**Arquivo**: `aula-02-performance-debugging-simplificada.md`

Mesmos conceitos explicados com analogias:
- pprof como médico/detetive
- trace como câmera de segurança
- Race Detector como guarda de segurança

### Aula 3: Exercícios e Reflexão
**Arquivo**: `aula-03-exercicios-e-reflexao.md`

Exercícios práticos:
- Identificar gargalos com pprof
- Detectar vazamentos de memória
- Analisar concorrência com trace
- Encontrar data races

Perguntas de reflexão sobre:
- Quando usar cada ferramenta
- Trade-offs de performance
- Interpretação de resultados
- Estratégias de correção

### Aula 4: Performance e Boas Práticas
**Arquivo**: `aula-04-performance-e-boas-praticas.md`

O que fazer e não fazer:
- Boas práticas com pprof
- Boas práticas com trace
- Boas práticas com Race Detector
- Estratégias de otimização
- Integração com CI/CD

---

## 🛠️ Ferramentas Necessárias

Para seguir este módulo, você precisa:

- Go 1.16+ instalado
- Navegador web (para interface do pprof e trace)
- Terminal/CLI
- Editor de código

**Ferramentas opcionais (mas recomendadas):**
- `graphviz` (para visualizações do pprof): `brew install graphviz` (macOS) ou `apt-get install graphviz` (Linux)

---

## 🚀 Como Usar Este Módulo

### 1. Leia a Aula Principal
Comece lendo `aula-01-performance-debugging-principal.md` para entender os conceitos técnicos.

### 2. Leia a Aula Simplificada
Continue com `aula-02-performance-debugging-simplificada.md` para fixar os conceitos com analogias.

### 3. Faça os Exercícios
Complete os exercícios em `aula-03-exercicios-e-reflexao.md`:
- Execute os programas
- Colete perfis e traces
- Analise os resultados
- Responda as perguntas de reflexão

### 4. Estude Boas Práticas
Leia `aula-04-performance-e-boas-praticas.md` para aprender:
- O que fazer e não fazer
- Como integrar no workflow
- Estratégias de otimização

---

## 📝 Exemplos Práticos

Este módulo inclui exemplos práticos que você pode executar:

### Exemplo 1: CPU Profile
```bash
# Executar programa com pprof
go run examples/cpu-profile/main.go

# Em outro terminal, coletar profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
```

### Exemplo 2: Memory Profile
```bash
# Executar programa com vazamento
go run examples/memory-leak/main.go

# Coletar heap profile
go tool pprof http://localhost:6060/debug/pprof/heap
```

### Exemplo 3: Trace
```bash
# Executar programa com trace
go run examples/trace/main.go

# Abrir trace no navegador
go tool trace trace.out
```

### Exemplo 4: Race Detector
```bash
# Executar com race detector
go run -race examples/race/main.go

# Ou testar
go test -race ./...
```

---

## 🔍 Comandos Úteis

### pprof
```bash
# Coletar CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Coletar heap profile
go tool pprof http://localhost:6060/debug/pprof/heap

# Coletar goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine

# Comandos no pprof interativo
(pprof) top          # Top funções
(pprof) top10        # Top 10 funções
(pprof) list Func    # Ver código da função
(pprof) web          # Abrir visualização
(pprof) svg          # Gerar SVG
```

### trace
```bash
# Gerar trace
go run -trace=trace.out main.go

# Abrir trace
go tool trace trace.out
```

### Race Detector
```bash
# Executar com race detector
go run -race main.go

# Testar com race detector
go test -race ./...

# Build com race detector (apenas para testes!)
go build -race
```

---

## 📊 Estrutura de Arquivos

```
34-performance-debugging/
├── README.md                                    # Este arquivo
├── aula-01-performance-debugging-principal.md   # Aula técnica completa
├── aula-02-performance-debugging-simplificada.md # Aula com analogias
├── aula-03-exercicios-e-reflexao.md              # Exercícios práticos
├── aula-04-performance-e-boas-praticas.md       # Boas práticas
└── 01-exemplos.go                                # Exemplos de código
```

---

## 🎓 Pré-requisitos

Antes de começar este módulo, você deve ter conhecimento de:

- ✅ Conceitos básicos de Go (variáveis, funções, structs)
- ✅ Concorrência em Go (goroutines, channels)
- ✅ Sincronização (sync.Mutex, sync.WaitGroup)
- ✅ HTTP básico (para servidor pprof)

**Módulos recomendados antes deste:**
- Módulo 16: Concorrência
- Módulo 17: Context
- Módulo 18: Padrões de Concorrência
- Módulo 19: Race Detection (se disponível)

---

## 💡 Dicas Importantes

1. **Sempre meça antes de otimizar**: Não otimize baseado em suposições
2. **Use Race Detector sempre em testes**: `go test -race ./...`
3. **Nunca use Race Detector em produção**: Muito lento (2-10x)
4. **Proteja pprof em produção**: Use autenticação
5. **Foque no top 10**: As 10 funções que mais consomem são as mais importantes

---

## 🔗 Recursos Adicionais

- [Go pprof Documentation](https://pkg.go.dev/net/http/pprof)
- [Go trace Documentation](https://pkg.go.dev/runtime/trace)
- [Go Race Detector](https://go.dev/blog/race-detector)
- [Profiling Go Programs](https://go.dev/blog/pprof)

---

## ❓ Dúvidas Frequentes

**P: Posso usar pprof em produção?**
R: Sim, mas com cuidado. Proteja com autenticação e colete apenas quando necessário.

**P: Posso usar Race Detector em produção?**
R: Não! É muito lento (2-10x). Use apenas em desenvolvimento e testes.

**P: Qual ferramenta devo usar primeiro?**
R: Depende do problema:
- Data race? → Race Detector
- Performance? → pprof
- Concorrência complexa? → trace

**P: Como interpreto resultados do pprof?**
R: Foque no top 10. Funções no top consomem mais recursos. Mas interprete com contexto - nem tudo no top é problema.

---

## 🎯 Próximos Passos

Após completar este módulo:

1. Integre as ferramentas no seu workflow de desenvolvimento
2. Adicione `go test -race` ao seu CI/CD
3. Pratique coletando e analisando perfis
4. Aplique em projetos reais

---

Boa sorte com o aprendizado! 🚀


