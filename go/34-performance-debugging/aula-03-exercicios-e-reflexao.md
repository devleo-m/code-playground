# Módulo 34: Performance e Debugging em Go
## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Identificando Gargalos com pprof

Crie um programa que simula processamento de dados e use `pprof` para identificar onde está o gargalo.

**Tarefas:**
1. Crie um programa `main.go` que:
   - Processa uma lista de 1000 números
   - Tem três funções: `quickProcess()`, `mediumProcess()`, e `slowProcess()`
   - `quickProcess()`: Simula processamento rápido (loop de 100 iterações)
   - `mediumProcess()`: Simula processamento médio (loop de 10.000 iterações)
   - `slowProcess()`: Simula processamento lento (loop de 1.000.000 iterações)
   - Chama cada função 10 vezes em sequência
   - Inclui servidor pprof na porta 6060

2. Execute o programa e colete um CPU profile por 10 segundos:
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
   ```

3. No prompt do pprof, execute:
   - `top` - Para ver as funções que mais consomem CPU
   - `list slowProcess` - Para ver o código da função lenta
   - `web` - Para ver visualização gráfica

4. Identifique qual função está consumindo mais CPU e explique por quê.

5. Otimize a função mais lenta (reduza o número de iterações) e colete um novo profile para comparar.

**Dica**: Use `time.Sleep()` para simular trabalho, mas também use loops com cálculos para garantir que o CPU seja usado.

**Arquivos esperados:**
- `main.go` (com servidor pprof)
- Análise escrita dos resultados do pprof

---

### Exercício 2: Detectando Vazamento de Memória com pprof

Crie um programa que tem um vazamento de memória intencional e use `pprof` para encontrá-lo.

**Tarefas:**
1. Crie um programa `leak.go` que:
   - Tem uma variável global `data []byte`
   - Em um loop infinito, aloca 1 MB de memória e armazena em `data`
   - Nunca libera a memória (vazamento intencional)
   - Inclui servidor pprof na porta 6060
   - Imprime estatísticas de memória a cada segundo

2. Execute o programa e deixe rodar por 30 segundos.

3. Em outro terminal, colete um heap profile:
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/heap
   ```

4. No prompt do pprof, execute:
   - `top` - Para ver onde a memória está sendo alocada
   - `list` com o nome da função que aloca memória
   - `web` - Para ver visualização gráfica

5. Identifique a função que está causando o vazamento.

6. Corrija o vazamento (liberando a memória ou não armazenando em variável global) e verifique que a memória para de crescer.

**Dica**: Use `runtime.ReadMemStats()` para ver estatísticas de memória.

**Arquivos esperados:**
- `leak.go` (versão com vazamento)
- `leak_fixed.go` (versão corrigida)
- Análise escrita do problema e solução

---

### Exercício 3: Analisando Concorrência com trace

Crie um programa com múltiplas goroutines e use `trace` para analisar seu comportamento.

**Tarefas:**
1. Crie um programa `concurrent.go` que:
   - Cria 50 goroutines
   - Cada goroutine:
     - Espera um tempo aleatório (0-100ms)
     - Processa dados (loop simples)
     - Espera outro tempo aleatório
   - Usa `sync.WaitGroup` para esperar todas terminarem
   - Gera um trace durante a execução

2. Execute o programa e gere o trace:
   ```bash
   go run concurrent.go
   ```

3. Abra o trace no navegador:
   ```bash
   go tool trace trace.out
   ```

4. Na interface web, explore:
   - **View trace**: Veja a timeline de execução
   - **Goroutine analysis**: Veja informações sobre cada goroutine
   - Identifique:
     - Quantas goroutines executam simultaneamente
     - Quando goroutines estão bloqueadas
     - Se há problemas de scheduling

5. Modifique o programa para adicionar um mutex que causa contenção (muitas goroutines competindo pelo mesmo mutex) e gere um novo trace. Compare os dois traces.

**Dica**: Use `runtime/trace` para gerar o trace.

**Arquivos esperados:**
- `concurrent.go` (versão inicial)
- `concurrent_contention.go` (versão com contenção)
- Análise escrita comparando os dois traces

---

### Exercício 4: Encontrando Data Races com Race Detector

Crie programas com data races intencionais e use o Race Detector para encontrá-los.

**Tarefas:**
1. Crie um programa `race1.go` com data race em uma variável inteira:
   - Variável global `counter int`
   - 10 goroutines incrementando `counter` 1000 vezes cada
   - Sem sincronização (data race intencional)

2. Execute com race detector:
   ```bash
   go run -race race1.go
   ```

3. Analise a saída do race detector:
   - Onde está o data race?
   - Quais goroutines estão envolvidas?
   - Qual é o problema?

4. Corrija o programa usando `sync.Mutex` e verifique que não há mais data races.

5. Crie um segundo programa `race2.go` com data race em um map:
   - Map global `m map[string]int`
   - Múltiplas goroutines escrevendo no map simultaneamente
   - Sem sincronização

6. Execute com race detector, analise e corrija.

7. Crie um terceiro programa `race3.go` com data race em um slice:
   - Slice global `data []int`
   - Uma goroutine adicionando elementos
   - Outra goroutine lendo elementos
   - Sem sincronização

8. Execute com race detector, analise e corrija.

**Dica**: Lembre-se que maps e slices não são thread-safe em Go.

**Arquivos esperados:**
- `race1.go` (com data race e versão corrigida)
- `race2.go` (com data race e versão corrigida)
- `race3.go` (com data race e versão corrigida)
- Análise escrita de cada data race encontrado

---

## Perguntas de Reflexão

### Reflexão 1: Quando Usar Cada Ferramenta

Você aprendeu três ferramentas diferentes: `pprof`, `trace` e Race Detector. Cada uma tem seu propósito específico.

**Pergunta**: 
1. Crie um **guia de decisão** (fluxograma ou árvore de decisão) que ajude a decidir qual ferramenta usar em diferentes situações. Considere cenários como:
   - Programa está lento
   - Programa está usando muita memória
   - Há suspeita de data race
   - Há problemas de concorrência
   - Há muitas goroutines
   - GC está causando pausas

2. Para cada ferramenta, dê **pelo menos 2 exemplos concretos** de situações onde ela é a melhor escolha e explique por quê.

3. Dê **pelo menos 1 exemplo** de situação onde usar uma ferramenta seria um erro (usar a ferramenta errada) e explique qual ferramenta deveria ser usada e por quê.

**Sua resposta deve demonstrar**: Capacidade de escolher a ferramenta certa para cada situação, considerando as características e limitações de cada uma.

---

### Reflexão 2: Trade-offs de Performance e Debugging

Usar ferramentas de profiling e debugging tem custos e benefícios. O Race Detector, por exemplo, torna o programa 2-10x mais lento.

**Pergunta**:
1. **Overhead das ferramentas**: 
   - Qual é o overhead de cada ferramenta (pprof, trace, Race Detector)?
   - Em que situações esse overhead é aceitável?
   - Em que situações esse overhead é inaceitável?
   - Como você equilibraria a necessidade de debugging com a performance?

2. **Uso em Produção**:
   - Quais ferramentas podem ser usadas em produção? Quais não podem? Por quê?
   - Se você precisasse usar uma ferramenta em produção, quais precauções tomaria?
   - Como você coletaria dados de performance em produção sem afetar os usuários?

3. **Estratégia de Debugging**:
   - Crie uma estratégia para debugging de problemas de performance em produção. Considere:
     - Como detectar problemas (monitoramento, alertas)
     - Como coletar dados (pprof, trace) sem afetar performance
     - Como analisar dados coletados
     - Como implementar correções
     - Como verificar que as correções funcionaram

**Sua resposta deve demonstrar**: Compreensão dos trade-offs práticos e capacidade de criar estratégias realistas para debugging em ambientes de produção.

---

### Reflexão 3: Interpretando Resultados de pprof

Ler e interpretar resultados de `pprof` requer prática e entendimento. Os resultados podem ser confusos para iniciantes.

**Pergunta**:
1. **Interpretando CPU Profile**:
   - O que significa quando uma função aparece no topo do `top` do pprof?
   - O que significa quando uma função consome 80% do tempo de CPU?
   - Quando isso é um problema? Quando não é?
   - Como você distinguiria entre "função que consome muito CPU porque faz muito trabalho útil" vs "função que consome muito CPU porque é ineficiente"?

2. **Interpretando Heap Profile**:
   - O que significa quando uma função aloca muita memória no heap profile?
   - Quando alocações grandes são aceitáveis? Quando não são?
   - Como você distinguiria entre "alocação necessária" vs "vazamento de memória"?
   - O que você procuraria em um heap profile para identificar vazamentos?

3. **Cenário Real**:
   - Imagine que você coletou um CPU profile e viu que `json.Marshal()` consome 30% do tempo. Isso é um problema? Por quê?
   - E se `runtime.gcBgMarkWorker` consumisse 20% do tempo? Isso seria um problema?
   - Como você priorizaria otimizações baseado em resultados de pprof?

**Sua resposta deve demonstrar**: Capacidade de interpretar resultados de profiling de forma crítica, distinguindo entre problemas reais e comportamento normal, e priorizando otimizações de forma inteligente.

---

### Reflexão 4: Data Races e Correção de Código Concorrente

Data races são bugs sutis e perigosos. Eles podem causar comportamento imprevisível que é difícil de reproduzir e debugar.

**Pergunta**:
1. **Por que Data Races são Perigosos**:
   - Explique por que data races são considerados bugs graves, mesmo quando o programa "funciona" na maioria das vezes.
   - Dê um exemplo concreto de como um data race pode causar comportamento incorreto que é difícil de detectar.
   - Por que data races são mais perigosos em produção do que em desenvolvimento?

2. **Estratégias de Prevenção**:
   - Além de usar Race Detector, quais outras estratégias você usaria para prevenir data races?
   - Como você estruturaria código concorrente para minimizar o risco de data races?
   - Qual é o papel de testes em detectar data races?

3. **Corrigindo Data Races**:
   - Quando você encontra um data race, quais são as opções para corrigi-lo?
   - Quando usar `sync.Mutex` vs `sync.RWMutex` vs channels?
   - Dê um exemplo de data race e mostre pelo menos 2 formas diferentes de corrigi-lo, explicando as vantagens e desvantagens de cada abordagem.

4. **Cenário Real**:
   - Imagine que você encontrou um data race em código de produção que está rodando há meses sem problemas aparentes. Como você lidaria com isso?
   - Quais seriam seus passos para:
     - Entender o impacto do data race
     - Corrigir o problema
     - Verificar que a correção não introduziu novos problemas
     - Fazer deploy da correção

**Sua resposta deve demonstrar**: Compreensão profunda dos riscos de data races, capacidade de prevenir e corrigir problemas de concorrência, e habilidade de tomar decisões arquiteturais informadas sobre sincronização.

---

## Como Entregar

Crie uma estrutura de diretórios organizada para seus exercícios:

```
34-performance-debugging/
├── exercicio1/
│   ├── main.go
│   └── analise.md (análise dos resultados)
├── exercicio2/
│   ├── leak.go
│   ├── leak_fixed.go
│   └── analise.md
├── exercicio3/
│   ├── concurrent.go
│   ├── concurrent_contention.go
│   └── analise.md
├── exercicio4/
│   ├── race1.go
│   ├── race2.go
│   ├── race3.go
│   └── analise.md
└── reflexoes.md (ou arquivo separado para cada reflexão)
```

**Importante**: 
- Execute cada programa e colete os perfis/traces conforme solicitado
- Documente seus achados em arquivos de análise
- Para os exercícios com pprof, inclua screenshots ou descrições dos resultados
- Para os exercícios com trace, descreva o que você observou na interface web
- Para os exercícios com race detector, copie as mensagens de warning e explique o que significam
- Para as reflexões, seja honesto e detalhado - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão profunda
- Se possível, experimente diferentes cenários e compare os resultados

---

## Dicas Extras

### Para Exercício 1:
- Use loops com cálculos (ex: `sum += i`) em vez de apenas `time.Sleep()` para garantir uso de CPU
- Deixe o programa rodar por tempo suficiente para coletar dados significativos
- Compare os resultados antes e depois da otimização

### Para Exercício 2:
- Monitore a memória por tempo suficiente para ver o crescimento
- Use `runtime.GC()` antes de coletar heap profile para ver alocações atuais
- Compare o comportamento antes e depois da correção

### Para Exercício 3:
- Use a interface web do trace - é muito mais visual e útil
- Experimente zoom in/out na timeline
- Clique em diferentes goroutines para ver seus detalhes
- Procure por padrões (ex: muitas goroutines bloqueadas ao mesmo tempo)

### Para Exercício 4:
- Execute múltiplas vezes - data races podem não aparecer sempre
- Leia cuidadosamente as mensagens do race detector - elas mostram exatamente onde está o problema
- Teste suas correções executando novamente com race detector
- Considere usar `go test -race` para testar suas correções

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!


