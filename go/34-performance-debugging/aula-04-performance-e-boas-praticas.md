# Módulo 34: Performance e Debugging em Go
## Aula 4 - Performance, Boas Práticas e O Que Deve/Não Deve Ser Utilizado

Agora que você entende `pprof`, `trace` e Race Detector, vamos falar sobre **como usá-los de forma eficiente e profissional**. Esta é a parte que separa programadores iniciantes de programadores experientes.

---

## 1. pprof: O Que Deve e NÃO Deve Ser Feito

### ❌ O Que NÃO Deve Ser Feito

#### 1.1. Coletar Perfis Sem Objetivo Claro

```go
// ❌ RUIM: Coletar perfis "só por coletar"
func main() {
    _ = net/http/pprof
    // Coletando perfis sem saber o que procurar
    // Sem métricas para comparar
    // Sem hipótese do problema
}
```

**Por quê?**: Coletar perfis sem objetivo é como fazer exames médicos sem sintomas - você pode encontrar coisas, mas não sabe se são relevantes.

**Quando NÃO coletar perfis:**
- ❌ Quando o programa está funcionando bem e não há problemas
- ❌ Quando você não tem uma hipótese do que pode estar errado
- ❌ Quando você não tem baseline (medida anterior) para comparar

#### 1.2. Usar pprof em Produção Sem Proteção

```go
// ❌ RUIM: pprof aberto em produção
func main() {
    // Qualquer um pode acessar perfis!
    log.Println(http.ListenAndServe(":6060", nil))
}
```

**Por quê?**: 
- Perfis podem expor informações sensíveis sobre seu código
- Coletar perfis tem overhead que pode afetar performance
- Pode ser usado para ataques (DoS)

**✅ BOM: Proteger pprof em produção**

```go
func setupPprof(mux *http.ServeMux) {
    if !isProduction() {
        // Em desenvolvimento, sempre disponível
        _ = net/http/pprof
        return
    }

    // Em produção, proteger com autenticação
    mux.HandleFunc("/debug/pprof/", func(w http.ResponseWriter, r *http.Request) {
        if !isAuthorized(r) {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        pprof.Index(w, r)
    })
}
```

#### 1.3. Coletar Perfis por Períodos Muito Longos

```go
// ❌ RUIM: Coletar CPU profile por 1 hora
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=3600
```

**Por quê?**: 
- Perfis muito longos são difíceis de analisar
- Podem incluir períodos irrelevantes
- Arquivos ficam muito grandes

**✅ BOM: Coletar por períodos específicos e relevantes**

```bash
# Coletar durante carga específica
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Coletar durante pico de uso
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60
```

#### 1.4. Ignorar Contexto ao Interpretar Perfis

```go
// ❌ RUIM: Ver função no top e assumir que é problema
// "json.Marshal está no top, vou otimizar!"
// Mas talvez seja normal para uma API que retorna JSON
```

**Por quê?**: Uma função no top do pprof não significa necessariamente que é um problema. Pode ser que ela faça muito trabalho útil.

**✅ BOM: Interpretar com contexto**

```go
// ✅ BOM: Analisar se faz sentido
// "json.Marshal está no top, mas nossa API retorna JSON"
// "Isso é esperado. Vou verificar se há alocações desnecessárias"
// "Vou verificar se podemos usar json.Encoder em vez de Marshal"
```

#### 1.5. Não Comparar Antes e Depois

```go
// ❌ RUIM: Otimizar sem medir impacto
func optimizedFunction() {
    // Fez otimização
    // Mas não mediu se melhorou
}
```

**Por quê?**: Sem comparação, você não sabe se a otimização realmente ajudou.

**✅ BOM: Sempre comparar**

```bash
# Antes
go tool pprof -base before.prof after.prof

# Ou coletar perfis separados
go tool pprof before.prof
go tool pprof after.prof
# Comparar resultados manualmente
```

---

### ✅ O Que Deve Ser Feito

#### 1.1. Ter Hipótese Antes de Coletar Perfis

```go
// ✅ BOM: Ter hipótese clara
// "Acho que o problema está na função processData"
// "Vou coletar perfil para confirmar"
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

**Quando coletar perfis:**
- ✅ Quando há problema de performance identificado
- ✅ Quando você tem hipótese do que pode estar errado
- ✅ Antes e depois de otimizações (para medir impacto)
- ✅ Durante testes de carga específicos

#### 1.2. Coletar Múltiplos Tipos de Perfis

```bash
# ✅ BOM: Coletar diferentes tipos de perfis
go tool pprof http://localhost:6060/debug/pprof/profile    # CPU
go tool pprof http://localhost:6060/debug/pprof/heap       # Memória
go tool pprof http://localhost:6060/debug/pprof/goroutine  # Goroutines
```

**Por quê?**: Diferentes problemas requerem diferentes perfis. CPU profile não mostra vazamentos de memória.

#### 1.3. Usar Visualizações Gráficas

```bash
# ✅ BOM: Usar web para visualização
(pprof) web

# Ou gerar SVG/PDF
(pprof) svg
(pprof) pdf
```

**Por quê?**: Visualizações gráficas tornam mais fácil identificar padrões e relações.

#### 1.4. Focar no Top 10

```bash
# ✅ BOM: Focar nas funções que mais consomem
(pprof) top10
(pprof) top20
```

**Por quê?**: Geralmente, as 10-20 funções no top são responsáveis pela maioria do consumo. Otimizar o resto tem pouco impacto.

#### 1.5. Integrar no Workflow de Desenvolvimento

```makefile
# ✅ BOM: Makefile com comandos úteis
.PHONY: profile-cpu profile-heap profile-goroutine

profile-cpu:
	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

profile-heap:
	go tool pprof http://localhost:6060/debug/pprof/heap

profile-goroutine:
	go tool pprof http://localhost:6060/debug/pprof/goroutine
```

---

## 2. trace: O Que Deve e NÃO Deve Ser Feito

### ❌ O Que NÃO Deve Ser Feito

#### 2.1. Coletar Traces por Períodos Muito Longos

```go
// ❌ RUIM: Coletar trace por 1 hora
trace.Start(f)
time.Sleep(1 * time.Hour) // Trace gigante!
trace.Stop()
```

**Por quê?**: 
- Traces longos geram arquivos enormes (GBs)
- Difíceis de analisar
- Interface web fica lenta

**✅ BOM: Coletar por períodos específicos**

```go
// ✅ BOM: Coletar durante período relevante
trace.Start(f)
doSpecificWork() // Período específico
trace.Stop()
```

#### 2.2. Usar trace Para Problemas Simples

```go
// ❌ RUIM: Usar trace para problema que pprof resolve
// "Meu programa está lento, vou usar trace"
// Mas pprof seria suficiente
```

**Por quê?**: trace tem mais overhead e é mais complexo. Use apenas quando pprof não é suficiente.

**Quando NÃO usar trace:**
- ❌ Quando pprof já mostra o problema claramente
- ❌ Para problemas simples de CPU ou memória
- ❌ Quando você não precisa ver timeline completa

#### 2.3. Não Entender a Interface Web

```bash
# ❌ RUIM: Abrir trace mas não saber o que procurar
go tool trace trace.out
# "O que eu estou vendo?"
```

**Por quê?**: A interface do trace é poderosa mas complexa. Sem entender, você não aproveita.

**✅ BOM: Aprender a interface**

- Use zoom (scroll) para ver detalhes
- Clique em eventos para ver informações
- Use diferentes views (goroutine analysis, etc.)
- Procure por padrões (GC frequente, goroutines bloqueadas)

---

### ✅ O Que Deve Ser Feito

#### 2.1. Usar trace Para Problemas de Concorrência

```go
// ✅ BOM: Usar trace quando há problemas de concorrência
// "Muitas goroutines, não sei por que estão bloqueadas"
trace.Start(f)
// Código com problema de concorrência
trace.Stop()
```

**Quando usar trace:**
- ✅ Problemas de concorrência difíceis de debugar
- ✅ Quando precisa ver timeline completa
- ✅ Quando pprof não mostra o problema
- ✅ Para entender comportamento de goroutines
- ✅ Para analisar GC pauses

#### 2.2. Coletar Durante Cenários Específicos

```go
// ✅ BOM: Coletar durante cenário específico
trace.Start(f)
handleHighLoad() // Cenário específico
trace.Stop()
```

**Por quê?**: Traces focados são mais fáceis de analisar e mais úteis.

#### 2.3. Usar a Interface Web Efetivamente

```bash
# ✅ BOM: Explorar diferentes views
go tool trace trace.out

# Na interface web:
# 1. View trace - Ver timeline completa
# 2. Goroutine analysis - Ver cada goroutine
# 3. Network blocking - Ver bloqueios de rede
# 4. Synchronization blocking - Ver bloqueios de sync
```

---

## 3. Race Detector: O Que Deve e NÃO Deve Ser Feito

### ❌ O Que NÃO Deve Ser Feito

#### 3.1. Usar em Produção

```bash
# ❌ RUIM: Compilar para produção com -race
go build -race -o myapp
# Muito lento! Nunca faça isso!
```

**Por quê?**: 
- Overhead de 2-10x mais lento
- Consumo de memória muito maior
- Apenas para detectar bugs, não para executar

**✅ BOM: Usar apenas em desenvolvimento/testes**

```bash
# ✅ BOM: Apenas em desenvolvimento
go test -race ./...
go run -race main.go
```

#### 3.2. Ignorar Warnings do Race Detector

```bash
# ❌ RUIM: Ver warning e ignorar
go test -race ./...
# WARNING: DATA RACE
# "Ah, deve ser falso positivo, vou ignorar"
```

**Por quê?**: Race Detector raramente dá falsos positivos. Se avisou, há problema real.

**✅ BOM: Corrigir todos os races**

```go
// ✅ BOM: Sempre corrigir
// Encontrou race? Corrija imediatamente!
// Mesmo que pareça inofensivo
```

#### 3.3. Não Usar em Testes

```bash
# ❌ RUIM: Não usar race detector em testes
go test ./...
# Pode ter data races não detectados!
```

**Por quê?**: Testes são o lugar perfeito para detectar races. Se não usar, pode ter bugs em produção.

**✅ BOM: Sempre usar em testes**

```bash
# ✅ BOM: Sempre usar
go test -race ./...

# Ou no Makefile
test:
	go test -race ./...
```

#### 3.4. Assumir que Código Sem Races Está Correto

```go
// ❌ RUIM: "Não há races, então está correto"
// Mas pode ter outros problemas de concorrência:
// - Deadlocks
// - Livelocks
// - Race conditions lógicas (não detectadas pelo detector)
```

**Por quê?**: Race Detector detecta data races, mas não detecta todos os problemas de concorrência.

**✅ BOM: Usar outras ferramentas também**

```go
// ✅ BOM: Combinar ferramentas
// 1. Race Detector - Para data races
// 2. Testes - Para lógica
// 3. pprof/trace - Para performance
```

---

### ✅ O Que Deve Ser Feito

#### 3.1. Sempre Usar em Testes

```bash
# ✅ BOM: Sempre usar em testes
go test -race ./...

# Ou no CI/CD
- name: Run tests with race detector
  run: go test -race ./...
```

#### 3.2. Corrigir Todos os Races Encontrados

```go
// ✅ BOM: Corrigir imediatamente
// Encontrou race? Corrija antes de continuar!
// Não assuma que é inofensivo
```

#### 3.3. Usar em Desenvolvimento

```bash
# ✅ BOM: Usar durante desenvolvimento
go run -race main.go
go build -race
```

#### 3.4. Integrar no CI/CD

```yaml
# ✅ BOM: Automatizar no CI/CD
- name: Test with race detector
  run: go test -race ./...
```

---

## 4. Performance: Impacto e Otimizações

### 4.1. Overhead das Ferramentas

| Ferramenta | Overhead | Uso em Produção |
|------------|----------|-----------------|
| **pprof** | Baixo (1-5%) | ✅ Pode usar (com proteção) |
| **trace** | Médio (5-10%) | ⚠️ Use com cuidado |
| **Race Detector** | Alto (200-1000%) | ❌ Nunca use |

### 4.2. Quando Coletar Perfis em Produção

**✅ Pode coletar:**
- Durante períodos de baixa carga
- Em servidores de staging/testes
- Com autenticação e proteção
- Por períodos curtos e específicos

**❌ Evite coletar:**
- Durante picos de carga
- Sem proteção/autenticação
- Por períodos longos
- Em servidores críticos sem necessidade

### 4.3. Estratégia de Coleta em Produção

```go
// ✅ BOM: Coletar sob demanda em produção
func handleProfileRequest(w http.ResponseWriter, r *http.Request) {
    if !isAuthorized(r) {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    duration, _ := time.ParseDuration(r.URL.Query().Get("seconds"))
    if duration == 0 || duration > 60*time.Second {
        duration = 30 * time.Second // Limite de 60s
    }

    // Coletar CPU profile
    f, _ := os.Create(fmt.Sprintf("cpu-%d.prof", time.Now().Unix()))
    pprof.StartCPUProfile(f)
    time.Sleep(duration)
    pprof.StopCPUProfile()
    f.Close()

    // Enviar arquivo
    http.ServeFile(w, r, f.Name())
}
```

---

## 5. Boas Práticas para a Vida do Programador

### 5.1. Workflow de Debugging

**Passo a passo recomendado:**

1. **Identificar o problema**
   - O que está errado? (lento, vazamento, crash, etc.)
   - Quando acontece?
   - Como reproduzir?

2. **Formular hipótese**
   - O que pode estar causando?
   - Onde pode estar o problema?

3. **Escolher ferramenta**
   - Race Detector: Se há suspeita de data race
   - pprof: Se há problema de performance
   - trace: Se há problema de concorrência complexo

4. **Coletar dados**
   - Coletar perfil/trace durante cenário específico
   - Garantir que dados são representativos

5. **Analisar resultados**
   - Interpretar com contexto
   - Focar no top 10
   - Usar visualizações

6. **Implementar correção**
   - Corrigir problema identificado
   - Medir impacto da correção

7. **Verificar correção**
   - Coletar novo perfil
   - Comparar antes e depois
   - Verificar que problema foi resolvido

### 5.2. Integração com CI/CD

```yaml
# .github/workflows/ci.yml
name: CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Run tests with race detector
        run: go test -race ./...
      
      - name: Build
        run: go build ./...

  # Opcional: Coletar perfis em testes de carga
  load-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Run load test
        run: |
          go run loadtest.go &
          sleep 10
          go tool pprof -proto http://localhost:6060/debug/pprof/profile?seconds=30 > profile.pb
```

### 5.3. Documentação

**Sempre documente:**
- Como coletar perfis
- Como interpretar resultados
- Onde encontrar perfis históricos
- Quais otimizações foram feitas e por quê

```go
// ✅ BOM: Documentar no código
// Para coletar CPU profile:
//   1. Execute o programa
//   2. go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
//   3. No prompt: top, list FunctionName, web
//
// Perfis históricos: /var/profiles/
// Última otimização: 2024-01-15 - Otimizada função processData
```

### 5.4. Monitoramento Contínuo

```go
// ✅ BOM: Coletar métricas continuamente
func collectMetrics() {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()

    for range ticker.C {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        
        // Enviar para sistema de monitoramento
        metrics.Gauge("memory.alloc", float64(m.Alloc))
        metrics.Gauge("memory.num_gc", float64(m.NumGC))
        metrics.Gauge("goroutines", float64(runtime.NumGoroutine()))
    }
}
```

---

## 6. Armadilhas Comuns e Como Evitá-las

### 6.1. Otimizar Sem Medir

**Problema**: Otimizar código sem medir se realmente melhora.

**Solução**: Sempre medir antes e depois.

```bash
# Antes
go tool pprof before.prof
(pprof) top

# Otimizar código

# Depois
go tool pprof after.prof
(pprof) top

# Comparar resultados
```

### 6.2. Focar em Micro-otimizações

**Problema**: Otimizar funções que consomem 1% do tempo.

**Solução**: Focar no top 10 que consome 80% do tempo.

```bash
# ✅ BOM: Focar no que importa
(pprof) top10
# Otimizar essas 10 funções primeiro
```

### 6.3. Ignorar Contexto

**Problema**: Ver função no top e assumir que é problema.

**Solução**: Interpretar com contexto do que o programa faz.

```go
// "json.Marshal está no top"
// Mas nossa API retorna JSON - isso é esperado!
// Focar em otimizar alocações, não remover json.Marshal
```

### 6.4. Não Usar Race Detector

**Problema**: Não usar race detector e ter bugs em produção.

**Solução**: Sempre usar em testes.

```bash
# ✅ BOM: Sempre usar
go test -race ./...
```

---

## 7. Checklist de Revisão

Antes de considerar seu código "otimizado", pergunte-se:

**Para pprof:**
- [ ] Coletou perfis com objetivo claro?
- [ ] Interpretou resultados com contexto?
- [ ] Focou no top 10?
- [ ] Comparou antes e depois?
- [ ] Protegeu pprof em produção?

**Para trace:**
- [ ] Usou trace apenas quando necessário?
- [ ] Coletou durante período específico?
- [ ] Entendeu a interface web?
- [ ] Identificou problemas reais?

**Para Race Detector:**
- [ ] Usou em todos os testes?
- [ ] Corrigiu todos os races encontrados?
- [ ] NÃO usou em produção?
- [ ] Integrou no CI/CD?

**Geral:**
- [ ] Tem workflow de debugging definido?
- [ ] Documentou como usar as ferramentas?
- [ ] Tem monitoramento contínuo?
- [ ] Mediu impacto das otimizações?

---

## 8. Resumo: Regras de Ouro

### Para pprof:
1. **Tenha objetivo claro** - Não colete perfis sem saber o que procurar
2. **Proteja em produção** - Sempre use autenticação
3. **Foque no top 10** - Otimize o que mais consome
4. **Compare antes e depois** - Meça impacto das otimizações
5. **Use visualizações** - Gráficos ajudam a entender

### Para trace:
1. **Use apenas quando necessário** - pprof primeiro, trace depois
2. **Colete por períodos específicos** - Não muito longos
3. **Aprenda a interface** - Explore as diferentes views
4. **Procure por padrões** - GC frequente, goroutines bloqueadas

### Para Race Detector:
1. **Sempre use em testes** - `go test -race ./...`
2. **Nunca use em produção** - Muito lento
3. **Corrija todos os races** - Não ignore warnings
4. **Integre no CI/CD** - Automatize detecção

### Performance:
1. **Medir antes de otimizar** - Sempre tenha baseline
2. **Focar no que importa** - Top 10 consome 80%
3. **Interpretar com contexto** - Nem tudo no top é problema
4. **Usar ferramentas certas** - Cada problema tem sua ferramenta

---

## Conclusão

Dominar `pprof`, `trace` e Race Detector vai além de saber a sintaxe. É sobre:

- **Escolher a ferramenta certa** para cada situação
- **Coletar dados relevantes** com objetivo claro
- **Interpretar resultados** com contexto e conhecimento
- **Otimizar sistematicamente** medindo impacto
- **Integrar no workflow** de desenvolvimento

Lembre-se: **Ferramentas são meios, não fins**. Use pprof, trace e Race Detector quando realmente ajudam, não apenas porque são "legais".

**Priorize:**
1. **Correção** sobre otimização prematura
2. **Medição** sobre suposição
3. **Contexto** sobre números isolados
4. **Sistemático** sobre aleatório

---

Na próxima etapa, você completará os exercícios e eu analisarei seu desempenho. Boa sorte!


