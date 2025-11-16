# Módulo 32: Security - Segurança em Aplicações Go
## Aula 3: Exercícios e Reflexão

Olá! Agora que você entendeu os conceitos de segurança e govulncheck, é hora de praticar! Vamos trabalhar com exercícios práticos que vão ajudar você a fixar o aprendizado e aplicar os conhecimentos em situações reais.

---

## Exercício 1: Instalação e Primeiros Passos

### Objetivo
Instalar e testar o govulncheck.

### Tarefas

1. **Instalar govulncheck**
   ```bash
   go install golang.org/x/vuln/cmd/govulncheck@latest
   ```

2. **Verificar instalação**
   ```bash
   govulncheck -version
   ```

3. **Criar um projeto de teste**
   ```bash
   mkdir exercicio-govulncheck
   cd exercicio-govulncheck
   go mod init exercicio-govulncheck
   ```

4. **Executar primeira verificação**
   ```bash
   govulncheck ./...
   ```

### Verificação
Execute cada comando e confirme que o govulncheck foi instalado corretamente e está funcionando.

### Reflexão
- Por que é importante verificar vulnerabilidades mesmo em projetos pequenos?
- Qual a diferença entre verificar código novo vs código legado?

---

## Exercício 2: Verificando Dependências

### Objetivo
Criar um projeto com dependências e verificar vulnerabilidades.

### Tarefas

1. **Criar arquivo main.go**
   ```go
   package main

   import (
       "fmt"
       "net/http"
       "github.com/gin-gonic/gin"
   )

   func main() {
       r := gin.Default()
       r.GET("/", func(c *gin.Context) {
           c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
       })
       r.Run(":8080")
   }
   ```

2. **Adicionar dependências**
   ```bash
   go get github.com/gin-gonic/gin
   go mod tidy
   ```

3. **Verificar vulnerabilidades**
   ```bash
   govulncheck ./...
   ```

4. **Verificar apenas dependências**
   ```bash
   govulncheck -mode=mod ./...
   ```

### Análise
- O que o govulncheck encontrou?
- Há diferença entre verificar código completo vs apenas dependências?
- Se encontrou vulnerabilidades, qual a severidade?

### Reflexão
- Por que é importante verificar dependências regularmente?
- Como você priorizaria correções se encontrasse múltiplas vulnerabilidades?

---

## Exercício 3: Trabalhando com Vulnerabilidades Encontradas

### Objetivo
Aprender a interpretar e corrigir vulnerabilidades encontradas.

### Cenário
Você executou `govulncheck ./...` e encontrou a seguinte vulnerabilidade:

```
Vulnerability #1: GO-2023-1234
  Package: golang.org/x/crypto
  Version: v0.1.0
  Severity: HIGH
  Description: Buffer overflow in crypto/rand
  
  Your code imports:
    - exemplo/cmd/server (uses crypto/rand)
  
  Recommendation: Update to v0.2.0 or later
```

### Tarefas

1. **Entender a vulnerabilidade**
   - Qual o ID da vulnerabilidade?
   - Qual pacote está afetado?
   - Qual a severidade?
   - Você realmente usa o código vulnerável?

2. **Verificar uso no código**
   - Procure no seu código onde `crypto/rand` é usado
   - Determine se você realmente precisa dessa funcionalidade

3. **Atualizar dependência**
   ```bash
   go get -u golang.org/x/crypto@latest
   go mod tidy
   ```

4. **Verificar novamente**
   ```bash
   govulncheck ./...
   ```

5. **Se não puder atualizar**
   - Documente a decisão
   - Explique por que não pode atualizar agora
   - Crie um plano de correção

### Reflexão
- Como você documentaria uma decisão de não atualizar imediatamente?
- Qual seria seu processo para priorizar correções de segurança?

---

## Exercício 4: Diferentes Modos de Operação

### Objetivo
Entender e praticar os diferentes modos do govulncheck.

### Tarefas

1. **Criar um projeto simples**
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

2. **Modo Source (padrão)**
   ```bash
   govulncheck ./...
   ```

3. **Modo Module**
   ```bash
   govulncheck -mode=mod ./...
   ```

4. **Compilar e usar Modo Binary**
   ```bash
   go build -o myapp .
   govulncheck -mode=binary ./myapp
   ```

### Análise
- Qual a diferença entre os modos?
- Quando você usaria cada modo?
- Qual modo é mais rápido? Por quê?

### Reflexão
- Por que o modo source é o padrão?
- Em que situações o modo binary seria útil?

---

## Exercício 5: Integração com CI/CD

### Objetivo
Criar uma integração básica do govulncheck com CI/CD.

### Tarefas

1. **Criar GitHub Actions workflow**
   Crie um arquivo `.github/workflows/security.yml`:

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

2. **Criar script de pre-commit**
   Crie um arquivo `.git/hooks/pre-commit`:

   ```bash
   #!/bin/sh
   echo "Executando govulncheck..."
   govulncheck ./...
   
   if [ $? -ne 0 ]; then
       echo "❌ govulncheck encontrou vulnerabilidades!"
       exit 1
   fi
   
   echo "✅ Nenhuma vulnerabilidade encontrada"
   ```

3. **Criar Makefile**
   ```makefile
   .PHONY: security test build

   security:
   	govulncheck ./...

   test:
   	go test ./...

   build:
   	go build -o app .

   all: security test build
   ```

### Análise
- Como cada integração funciona?
- Qual a vantagem de cada abordagem?
- Como você melhoraria essas integrações?

### Reflexão
- Por que é importante integrar verificações de segurança no CI/CD?
- Como você convenceria sua equipe a adotar essas práticas?

---

## Exercício 6: Análise de Projeto Existente

### Objetivo
Aplicar govulncheck em um projeto real ou exemplo.

### Tarefas

1. **Escolher um projeto**
   - Pode ser um projeto seu
   - Ou um projeto open source pequeno
   - Ou criar um projeto exemplo com várias dependências

2. **Executar govulncheck**
   ```bash
   govulncheck ./...
   ```

3. **Analisar resultados**
   - Liste todas as vulnerabilidades encontradas
   - Classifique por severidade
   - Identifique quais você realmente usa

4. **Criar plano de ação**
   - Priorize correções
   - Documente decisões
   - Crie timeline de correções

5. **Executar correções**
   - Atualize dependências quando possível
   - Documente quando não puder atualizar
   - Verifique novamente

### Análise
- Quantas vulnerabilidades foram encontradas?
- Qual a distribuição por severidade?
- Quantas você realmente usa no código?

### Reflexão
- Como você comunicaria vulnerabilidades encontradas para sua equipe?
- Qual seria seu processo para corrigir vulnerabilidades em produção?

---

## Exercício 7: Comparação com Outras Ferramentas

### Objetivo
Entender quando usar govulncheck vs outras ferramentas.

### Tarefas

1. **Pesquisar outras ferramentas**
   - gosec
   - nancy
   - snyk
   - Dependabot

2. **Criar tabela comparativa**
   | Ferramenta | Foco | Quando Usar |
   |------------|------|-------------|
   | govulncheck | ? | ? |
   | gosec | ? | ? |
   | nancy | ? | ? |

3. **Testar gosec (opcional)**
   ```bash
   go install github.com/securego/gosec/v2/cmd/gosec@latest
   gosec ./...
   ```

4. **Comparar resultados**
   - O que cada ferramenta encontrou?
   - Há sobreposição?
   - Qual complementa qual?

### Reflexão
- Por que usar múltiplas ferramentas de segurança?
- Como você integraria diferentes ferramentas no workflow?

---

## Exercício 8: Documentação de Decisões de Segurança

### Objetivo
Aprender a documentar decisões de segurança adequadamente.

### Cenário
Você encontrou uma vulnerabilidade HIGH, mas não pode atualizar a dependência agora porque:
- A nova versão quebra compatibilidade
- O projeto está em fase crítica de desenvolvimento
- A atualização requer refatoração significativa

### Tarefas

1. **Criar documento de decisão**
   Crie um arquivo `SECURITY.md` ou adicione ao README:

   ```markdown
   ## Decisões de Segurança

   ### GO-2023-1234 - golang.org/x/crypto v0.1.0
   
   **Severidade:** HIGH
   **Data de Descoberta:** 2024-01-15
   **Status:** Documentado, correção planejada
   
   **Descrição:**
   Buffer overflow em crypto/rand que pode permitir leitura de memória.
   
   **Impacto:**
   Nossa aplicação usa crypto/rand para gerar tokens de sessão.
   
   **Razão para Não Atualizar Imediatamente:**
   - Nova versão (v0.2.0) requer Go 1.22+
   - Projeto atual usa Go 1.21
   - Atualização de Go requer atualização de toda infraestrutura
   
   **Plano de Correção:**
   1. Atualizar para Go 1.22 na próxima sprint (2024-02-01)
   2. Atualizar golang.org/x/crypto para v0.2.0
   3. Verificar com govulncheck
   
   **Mitigações Temporárias:**
   - Implementado rate limiting adicional
   - Monitoramento aumentado de logs
   - Revisão de código relacionado
   
   **Responsável:** Equipe de Segurança
   **Revisão:** 2024-02-01
   ```

2. **Criar issue de tracking**
   - Crie uma issue no seu sistema de tracking
   - Adicione labels apropriados
   - Defina prioridade e timeline

3. **Comunicar para equipe**
   - Documente em reunião de equipe
   - Adicione ao changelog se relevante
   - Informe stakeholders se necessário

### Reflexão
- Por que é importante documentar decisões de segurança?
- Como você garantiria que vulnerabilidades documentadas sejam corrigidas?
- Qual seria seu processo de revisão de decisões de segurança?

---

## Exercício 9: Criando Política de Segurança

### Objetivo
Criar uma política de segurança para seu projeto ou equipe.

### Tarefas

1. **Definir frequência de verificação**
   - Com que frequência executar govulncheck?
   - Quando verificar no CI/CD?
   - Quando verificar manualmente?

2. **Definir processo de correção**
   - Quanto tempo para corrigir CRITICAL?
   - Quanto tempo para corrigir HIGH?
   - Como documentar quando não pode corrigir?

3. **Definir responsabilidades**
   - Quem é responsável por executar verificações?
   - Quem é responsável por corrigir vulnerabilidades?
   - Quem aprova decisões de não corrigir?

4. **Criar documento de política**
   ```markdown
   # Política de Segurança

   ## Verificações
   - govulncheck executado em cada commit via CI/CD
   - Verificação manual semanal
   - Verificação completa antes de releases

   ## Processo de Correção
   - CRITICAL: Corrigir em 24 horas
   - HIGH: Corrigir em 1 semana
   - MEDIUM: Corrigir em 1 mês
   - LOW: Corrigir quando possível

   ## Documentação
   - Todas as vulnerabilidades devem ser documentadas
   - Decisões de não corrigir devem ser aprovadas
   - Revisão trimestral de vulnerabilidades documentadas
   ```

### Reflexão
- Como você adaptaria essa política para diferentes tipos de projetos?
- Como você garantiria que a política seja seguida?
- Como você mediria a efetividade da política?

---

## Exercício 10: Cenário Real - Projeto com Múltiplas Vulnerabilidades

### Objetivo
Praticar gerenciamento de múltiplas vulnerabilidades em um cenário realista.

### Cenário
Você é responsável por um projeto que tem:
- 1 vulnerabilidade CRITICAL
- 3 vulnerabilidades HIGH
- 5 vulnerabilidades MEDIUM
- 2 vulnerabilidades LOW

O projeto está em produção e tem 10.000 usuários ativos.

### Tarefas

1. **Priorizar correções**
   - Crie uma lista priorizada
   - Justifique a ordem
   - Estime tempo para cada correção

2. **Criar plano de ação**
   - Defina timeline
   - Identifique recursos necessários
   - Identifique riscos

3. **Comunicar para stakeholders**
   - Crie um resumo executivo
   - Explique impacto
   - Apresente plano de correção

4. **Implementar correções**
   - Comece pelas mais críticas
   - Documente progresso
   - Verifique após cada correção

### Reflexão
- Como você balancearia correções de segurança com novas funcionalidades?
- Como você comunicaria vulnerabilidades para usuários se necessário?
- Como você garantiria que vulnerabilidades não sejam reintroduzidas?

---

## Perguntas de Reflexão Final

1. **Compreensão Conceitual**
   - Qual a diferença entre vulnerabilidade e bug?
   - Por que govulncheck é importante mesmo que você escreva código seguro?
   - Como o Go Vulnerability Database é mantido?

2. **Aplicação Prática**
   - Como você integraria govulncheck no workflow da sua equipe?
   - Qual seria seu processo para lidar com vulnerabilidades em código legado?
   - Como você priorizaria correções quando há múltiplas vulnerabilidades?

3. **Pensamento Crítico**
   - É sempre possível corrigir vulnerabilidades imediatamente? Por quê?
   - Como você balancearia segurança com velocidade de desenvolvimento?
   - Qual o papel da documentação em decisões de segurança?

4. **Aprendizado Contínuo**
   - Como você se manteria atualizado sobre novas vulnerabilidades?
   - Como você compartilharia conhecimento de segurança com sua equipe?
   - Como você melhoraria continuamente o processo de segurança?

---

## Conclusão

Parabéns por completar os exercícios! Você agora tem experiência prática com:

- ✅ Instalação e uso básico do govulncheck
- ✅ Interpretação de resultados
- ✅ Correção de vulnerabilidades
- ✅ Integração com CI/CD
- ✅ Documentação de decisões
- ✅ Criação de políticas de segurança

Continue praticando e aplicando esses conceitos em seus projetos reais. Segurança é uma jornada contínua, não um destino!

Na próxima aula, vamos mergulhar em boas práticas e otimizações para usar o govulncheck de forma eficiente e profissional!

