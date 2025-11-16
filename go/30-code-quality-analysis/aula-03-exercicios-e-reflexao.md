# Módulo 30: Code Quality and Analysis
## Aula 3: Exercícios e Reflexão

Olá! Agora é hora de colocar em prática tudo que aprendemos sobre `go vet` e `goimports`. Vamos exercitar nossos conhecimentos e refletir sobre a importância dessas ferramentas no desenvolvimento profissional.

---

## 📝 Exercícios Práticos

### Exercício 1: Detectando Problemas com `go vet`

Crie um arquivo chamado `exercicio1.go` com o seguinte código intencionalmente problemático:

```go
package main

import (
    "fmt"
    "os"  // Este import não está sendo usado
)

func main() {
    // Problema 1: Código inalcançável
    return
    fmt.Println("Esta linha nunca será executada")
    
    // Problema 2: Formato incorreto em printf
    nome := "Maria"
    idade := 25
    fmt.Printf("Nome: %d, Idade: %s\n", nome, idade)
    
    // Problema 3: Variável não utilizada
    resultado := 10 + 20
}
```

**Tarefas:**
1. Execute `go vet exercicio1.go` e anote todos os problemas encontrados
2. Corrija cada problema identificado
3. Execute `go vet` novamente para confirmar que não há mais problemas
4. Explique o que cada problema significava e por que era um bug potencial

---

### Exercício 2: Organizando Imports com `goimports`

Crie um arquivo chamado `exercicio2.go` com o seguinte código:

```go
package main

import "fmt"

func main() {
    // Você está usando strings mas não importou
    texto := strings.ToUpper("hello world")
    
    // Você está usando time mas não importou
    time.Sleep(2 * time.Second)
    
    fmt.Println(texto)
    
    // Você importou fmt mas não está usando mais (depois de remover o Println acima)
}
```

**Tarefas:**
1. Execute `goimports -d exercicio2.go` para ver o que seria modificado (sem modificar)
2. Execute `goimports -w exercicio2.go` para aplicar as mudanças
3. Verifique o arquivo e explique:
   - Quais imports foram adicionados?
   - Quais imports foram removidos (se houver)?
   - Como os imports foram organizados?

---

### Exercício 3: Problemas com Range Loops

Crie um arquivo chamado `exercicio3.go`:

```go
package main

import "fmt"

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    
    // Tentativa de dobrar os valores
    for _, num := range numeros {
        num = num * 2
    }
    
    // Imprimir os valores (esperando que estejam dobrados)
    for _, num := range numeros {
        fmt.Println(num)
    }
}
```

**Tarefas:**
1. Execute `go vet exercicio3.go` e veja o que é detectado
2. Execute o código e observe o resultado
3. Explique por que os valores não foram dobrados
4. Corrija o código para que funcione corretamente
5. Execute `go vet` novamente para confirmar que o problema foi resolvido

---

### Exercício 4: Projeto Completo - Aplicando Tudo

Crie um pequeno projeto com a seguinte estrutura:

```
meu-projeto/
├── main.go
├── utils.go
└── utils_test.go
```

**main.go:**
```go
package main

import "fmt"

func main() {
    resultado := Soma(10, 20)
    fmt.Printf("Resultado: %s\n", resultado)
    
    texto := FormatarTexto("hello")
    fmt.Println(texto)
}
```

**utils.go:**
```go
package main

import "strings"

func Soma(a, b int) int {
    return a + b
}

func FormatarTexto(s string) string {
    return strings.ToUpper(s)
}
```

**Tarefas:**
1. Execute `goimports -w .` em todo o projeto
2. Execute `go vet ./...` em todo o projeto
3. Identifique e corrija todos os problemas encontrados
4. Execute novamente para confirmar que está tudo correto
5. Crie um script ou comando que automatize esse processo

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por Que `go vet` é Importante?

**Pergunta:**
Imagine que você está trabalhando em um projeto grande com uma equipe de 10 desenvolvedores. Um desenvolvedor comete um erro comum (como usar `%d` com uma string) e commita o código sem executar `go vet`.

**Reflita sobre:**
1. Quanto tempo pode ser perdido até esse bug ser descoberto?
2. Quais são os custos (tempo, dinheiro, frustração) de descobrir bugs em produção vs. durante o desenvolvimento?
3. Como `go vet` ajuda a prevenir esses problemas?
4. Por que é importante executar `go vet` antes de cada commit, mesmo em projetos pequenos?

**Sua resposta deve ter pelo menos 3 parágrafos explicando sua reflexão.**

---

### Reflexão 2: A Importância da Automação

**Pergunta:**
`goimports` automatiza uma tarefa que você poderia fazer manualmente (adicionar/remover imports). Alguns desenvolvedores argumentam: "Por que usar uma ferramenta se eu posso fazer manualmente? Eu aprendo mais fazendo sozinho."

**Reflita sobre:**
1. Quais são as vantagens de automatizar tarefas repetitivas como gerenciar imports?
2. O que acontece quando você esquece de adicionar um import manualmente? E quando esquece de remover um não utilizado?
3. Em um projeto com 50 arquivos, quanto tempo você gastaria gerenciando imports manualmente vs. usando `goimports`?
4. A automação tira oportunidades de aprendizado ou libera tempo para aprender coisas mais importantes?
5. Como a automação ajuda na consistência quando trabalhamos em equipe?

**Sua resposta deve incluir:**
- Uma análise dos prós e contras da automação
- Exemplos práticos de quando a automação é valiosa
- Sua opinião sobre o equilíbrio entre automação e aprendizado manual

---

### Reflexão 3: Qualidade de Código vs. Velocidade de Desenvolvimento

**Pergunta:**
Você está em uma situação onde precisa entregar uma funcionalidade rapidamente. Seu chefe está pressionando para que você entregue o mais rápido possível. Você tem duas opções:

**Opção A:** Pular `go vet` e `goimports` para economizar tempo e entregar mais rápido
**Opção B:** Sempre executar `go vet` e `goimports`, mesmo sob pressão

**Reflita sobre:**
1. Qual opção você escolheria e por quê?
2. Quanto tempo realmente economiza ao pular essas ferramentas? (Dica: geralmente são segundos)
3. Quais são os riscos de pular essas verificações, mesmo em situações de pressão?
4. Como você pode convencer seu chefe/equipe de que vale a pena investir esses poucos segundos?
5. Existe uma situação onde seria aceitável pular essas ferramentas? Por quê?

**Sua resposta deve incluir:**
- Uma decisão clara com justificativa
- Uma análise de custo-benefício (tempo vs. qualidade)
- Uma estratégia para defender a qualidade do código mesmo sob pressão

---

## 🎯 Desafio Extra (Opcional)

### Desafio: Criar um Script de Qualidade

Crie um script (bash, PowerShell, ou Makefile) que automatize todo o processo de verificação de qualidade. O script deve:

1. ✅ Executar `goimports -w` em todo o projeto
2. ✅ Executar `go vet` em todo o projeto
3. ✅ Executar `go test` em todo o projeto
4. ✅ Mostrar um resumo colorido dos resultados
5. ✅ Retornar código de erro se algo falhar (útil para CI/CD)

**Requisitos:**
- O script deve funcionar em Linux, macOS e Windows (ou criar versões separadas)
- Deve ser fácil de executar: `./check-quality.sh` ou `make quality`
- Deve mostrar mensagens claras sobre o que está fazendo
- Deve parar na primeira falha ou continuar e mostrar tudo no final (sua escolha)

**Bônus:**
- Adicionar verificação de formatação com `gofmt -l`
- Adicionar verificação de cobertura de testes
- Criar um hook de pre-commit que usa esse script

---

## 📋 Checklist de Aprendizado

Antes de considerar que você dominou este módulo, certifique-se de que você consegue:

- [ ] Explicar o que `go vet` faz e por que é importante
- [ ] Listar pelo menos 5 tipos de problemas que `go vet` detecta
- [ ] Executar `go vet` e interpretar os resultados
- [ ] Explicar a diferença entre `gofmt` e `goimports`
- [ ] Instalar e usar `goimports` corretamente
- [ ] Entender como `goimports` organiza imports em grupos
- [ ] Configurar `goimports` no seu editor (VS Code, GoLand, etc.)
- [ ] Criar um workflow que combine `go vet` e `goimports`
- [ ] Explicar por que essas ferramentas são essenciais para qualidade de código
- [ ] Refletir sobre a importância da automação e qualidade no desenvolvimento

---

## 💡 Dicas para os Exercícios

1. **Não tenha pressa**: Leia cada problema cuidadosamente antes de começar
2. **Experimente**: Tente quebrar o código propositalmente para ver o que `go vet` detecta
3. **Documente**: Anote o que você aprendeu com cada exercício
4. **Pesquise**: Se encontrar algo que não entende, pesquise na documentação oficial
5. **Reflita profundamente**: As perguntas de reflexão são tão importantes quanto os exercícios práticos

---

## 🎓 Próximos Passos

Após completar estes exercícios e reflexões, você estará pronto para:
- Integrar essas ferramentas no seu workflow diário
- Configurar automação no seu editor
- Criar scripts de qualidade para seus projetos
- Entender a importância da qualidade de código no desenvolvimento profissional

Boa sorte com os exercícios! Lembre-se: a prática é o que transforma conhecimento em habilidade. 🚀

