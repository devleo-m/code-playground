# Módulo 33: Code Generation e Build Tags em Go
## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Gerando Métodos String() com go generate

Crie um programa que usa `go generate` para gerar automaticamente métodos `String()` para um tipo enum.

**Tarefas:**
1. Crie um arquivo `status.go` com um tipo `Status` que representa estados de uma tarefa:
   - `StatusTodo` (A fazer)
   - `StatusInProgress` (Em progresso)
   - `StatusDone` (Concluída)
   - `StatusCancelled` (Cancelada)

2. Adicione a diretiva `//go:generate` para usar o `stringer` tool

3. Instale o `stringer` e execute `go generate`

4. Crie um arquivo `main.go` que usa o tipo `Status` e imprime os valores usando o método `String()` gerado

5. Adicione um novo status `StatusBlocked` (Bloqueada) e execute `go generate` novamente para ver a atualização automática

**Dica**: 
- Instale stringer: `go install golang.org/x/tools/cmd/stringer@latest`
- Use: `//go:generate stringer -type=Status`

**Arquivos esperados:**
- `status.go` (com a diretiva go:generate)
- `status_string.go` (gerado automaticamente)
- `main.go` (que usa o Status)

---

### Exercício 2: Criando Build Tags para Diferentes Plataformas

Crie um programa que retorna caminhos diferentes dependendo do sistema operacional usando build tags.

**Tarefas:**
1. Crie três arquivos:
   - `config_linux.go` (ou .unix.go)` - retorna `/etc/myapp/config.json`
   - `config_windows.go` - retorna `C:\ProgramData\myapp\config.json`
   - `config_default.go` - retorna `./config.json` (para outros sistemas)

2. Cada arquivo deve ter a função `GetConfigPath() string` com build tags apropriadas

3. Crie um `main.go` que chama `GetConfigPath()` e imprime o resultado

4. Teste compilando para diferentes plataformas:
   ```bash
   # Para Linux
   go build -o myapp-linux
   
   # Para Windows
   GOOS=windows go build -o myapp-windows.exe
   
   # Para macOS
   GOOS=darwin go build -o myapp-darwin
   ```

**Dica**: Use `//go:build linux` para Linux, `//go:build windows` para Windows, e `//go:build !linux && !windows` para default.

**Pergunta para pensar**: Por que é melhor usar build tags nesse caso do que verificar `runtime.GOOS` em tempo de execução?

---

### Exercício 3: Feature Flags com Build Tags

Crie um sistema de feature flags usando build tags para controlar funcionalidades em desenvolvimento e produção.

**Tarefas:**
1. Crie dois arquivos:
   - `features_dev.go` com build tag `dev` contendo:
     - `EnableDebugLogging = true`
     - `EnableProfiling = true`
     - `MaxConnections = 1000`
   
   - `features_prod.go` com build tag `!dev` contendo:
     - `EnableDebugLogging = false`
     - `EnableProfiling = false`
     - `MaxConnections = 100`

2. Crie um `main.go` que:
   - Verifica se `EnableDebugLogging` está ativo e imprime uma mensagem
   - Mostra o valor de `MaxConnections`
   - Se `EnableProfiling` estiver ativo, imprime "Profiling enabled"

3. Compile duas versões:
   ```bash
   # Versão de desenvolvimento
   go build -tags=dev -o myapp-dev
   
   # Versão de produção
   go build -o myapp-prod
   ```

4. Execute ambas e compare as saídas

**Pergunta para pensar**: Quais são as vantagens de usar build tags para feature flags em vez de variáveis de ambiente ou arquivos de configuração?

---

### Exercício 4: Combinando go generate e Build Tags

Crie um exemplo mais complexo que combina `go generate` e build tags.

**Tarefas:**
1. Crie um tipo `LogLevel` com valores:
   - `LogLevelDebug`
   - `LogLevelInfo`
   - `LogLevelWarning`
   - `LogLevelError`

2. Use `go generate` para gerar o método `String()` para `LogLevel`

3. Crie dois arquivos de configuração com build tags:
   - `config_dev.go` (tag `dev`): Define `DefaultLogLevel = LogLevelDebug`
   - `config_prod.go` (tag `!dev`): Define `DefaultLogLevel = LogLevelInfo`

4. Crie um `main.go` que:
   - Usa o `DefaultLogLevel` definido pelas build tags
   - Imprime o log level usando o método `String()` gerado

5. Teste compilando com e sem a tag `dev`:
   ```bash
   go generate ./...
   go build -tags=dev -o myapp-dev
   go build -o myapp-prod
   ```

**Dica**: Este exercício demonstra como as duas ferramentas trabalham juntas - `go generate` cria código automaticamente, e build tags controlam qual código é incluído.

---

## Perguntas de Reflexão

### Reflexão 1: Quando Usar go generate vs Código Manual

Você aprendeu que `go generate` é útil para gerar código automaticamente, mas nem sempre é a melhor solução.

**Pergunta**: 
1. Dê **três situações** onde usar `go generate` é claramente vantajoso e explique por quê.
2. Dê **três situações** onde escrever código manualmente é melhor do que usar `go generate` e explique por quê.
3. Qual é o **custo** (desvantagem) de usar `go generate`? Quando esse custo pode não valer a pena?

**Sua resposta deve demonstrar**: Capacidade de avaliar quando uma ferramenta é apropriada e quando não é, considerando trade-offs práticos.

---

### Reflexão 2: Build Tags vs Runtime Checks - Decisões de Design

Você aprendeu que build tags são resolvidos em tempo de compilação, enquanto runtime checks são resolvidos em tempo de execução.

**Pergunta**:
1. Imagine que você está criando uma aplicação que precisa funcionar com diferentes bancos de dados (PostgreSQL, MySQL, SQLite). Você usaria build tags ou runtime checks? **Justifique sua escolha** considerando:
   - Facilidade de deploy
   - Tamanho do binário
   - Flexibilidade
   - Manutenibilidade

2. Agora imagine que você precisa de um "modo debug" que pode ser ativado/desativado sem recompilar. Você usaria build tags ou runtime checks? **Justifique sua escolha**.

3. Crie um **critério de decisão** (uma "regra de ouro") que você seguiria para escolher entre build tags e runtime checks em futuros projetos. Sua regra deve ser clara e aplicável a diferentes situações.

**Sua resposta deve demonstrar**: Compreensão profunda das diferenças entre as abordagens e capacidade de tomar decisões arquiteturais informadas.

---

### Reflexão 3: Impacto de go generate e Build Tags em Projetos Reais

Pense em um projeto real que você conhece ou já trabalhou (pode ser um projeto pessoal, open source, ou até mesmo hipotético).

**Pergunta**:
1. Identifique **pelo menos 3 lugares** nesse projeto onde `go generate` poderia ser útil. Para cada um, explique:
   - O que seria gerado
   - Por que a geração automática seria melhor que código manual
   - Qual ferramenta você usaria (stringer, mockgen, protoc, etc.)

2. Identifique **pelo menos 2 lugares** onde build tags poderiam ser úteis. Para cada um, explique:
   - Qual seria a condição (OS, arquitetura, feature flag, etc.)
   - Por que build tags são melhores que alternativas (runtime checks, configuração, etc.)
   - Como isso melhoraria o projeto

3. Considere os **custos e benefícios**:
   - Quanto tempo você economizaria usando essas ferramentas?
   - Quais seriam os desafios de implementação?
   - Como isso afetaria outros desenvolvedores do projeto?
   - Vale a pena o esforço? Por quê?

**Sua resposta deve demonstrar**: Capacidade de aplicar conceitos aprendidos em contextos reais e avaliar criticamente o valor das ferramentas.

---

### Reflexão 4: Manutenibilidade e Evolução de Código Gerado

Código gerado automaticamente tem implicações importantes para manutenibilidade e evolução de projetos.

**Pergunta**:
1. **Vantagens de código gerado**: Liste e explique pelo menos 3 vantagens de manter código gerado no repositório Git (em vez de gerar apenas no CI/CD).

2. **Desvantagens de código gerado**: Liste e explique pelo menos 3 desvantagens ou riscos de usar código gerado.

3. **Estratégia de versionamento**: Imagine que você está criando uma biblioteca Go que usa `go generate`. Você versionaria os arquivos gerados? Por quê? Como isso afetaria:
   - Desenvolvedores que usam sua biblioteca
   - O processo de CI/CD
   - A facilidade de contribuição (pull requests)

4. **Cenário real**: Um desenvolvedor júnior adiciona um novo valor ao enum `Status` mas esquece de rodar `go generate`. O código compila, mas o método `String()` não funciona corretamente para o novo valor. Como você preveniria esse problema? Considere:
   - Documentação
   - Scripts/Makefiles
   - CI/CD checks
   - Outras estratégias

**Sua resposta deve demonstrar**: Compreensão dos aspectos práticos e organizacionais de usar ferramentas de geração de código em projetos reais.

---

## Como Entregar

Crie uma estrutura de diretórios organizada para seus exercícios:

```
33-code-generation-build-tags/
├── exercicio1/
│   ├── status.go
│   ├── main.go
│   └── status_string.go (gerado)
├── exercicio2/
│   ├── config_linux.go
│   ├── config_windows.go
│   ├── config_default.go
│   └── main.go
├── exercicio3/
│   ├── features_dev.go
│   ├── features_prod.go
│   └── main.go
├── exercicio4/
│   ├── loglevel.go
│   ├── config_dev.go
│   ├── config_prod.go
│   ├── main.go
│   └── loglevel_string.go (gerado)
└── reflexoes.md (ou arquivo separado para cada reflexão)
```

**Importante**: 
- Compile e execute cada programa para garantir que funciona
- Comente seu código explicando o que cada parte faz
- Para os exercícios com `go generate`, inclua instruções de como executar
- Para as reflexões, seja honesto e detalhado - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão profunda
- Se possível, teste os exercícios em diferentes sistemas operacionais ou use `GOOS` para simular

---

## Dicas Extras

### Para Exercício 1:
- Se o `stringer` não estiver instalado, use: `go install golang.org/x/tools/cmd/stringer@latest`
- Verifique se o arquivo `status_string.go` foi gerado após rodar `go generate`
- Adicione um comentário explicando que o arquivo foi gerado automaticamente

### Para Exercício 2:
- Use `GOOS` para testar compilação para diferentes plataformas sem precisar mudar de máquina
- Verifique que apenas o arquivo correto é compilado para cada plataforma
- Tente compilar sem nenhum arquivo de config e veja o erro (depois crie o default)

### Para Exercício 3:
- Compare o tamanho dos binários gerados (`ls -lh myapp-dev myapp-prod`)
- Execute ambos e compare o comportamento
- Pense em como isso se aplicaria a um projeto real

### Para Exercício 4:
- Este exercício combina tudo que você aprendeu
- Preste atenção na ordem: primeiro `go generate`, depois `go build`
- Documente o processo completo

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!

