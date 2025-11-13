# Módulo 5: Documentação e Comandos em Go

## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Documentar uma Função

Crie um arquivo `exercicio1.go` e documente completamente a seguinte função:

```go
package main

import "fmt"

// TODO: Adicione documentação completa aqui
func CalculaIMC(peso, altura float64) float64 {
    if altura <= 0 {
        return 0
    }
    return peso / (altura * altura)
}

func main() {
    imc := CalculaIMC(70, 1.75)
    fmt.Printf("IMC: %.2f\n", imc)
}
```

**Tarefas:**

1. Adicione documentação que explique o que a função faz
2. Documente os parâmetros (peso e altura)
3. Documente o retorno (o que significa o valor retornado)
4. Adicione um exemplo de uso
5. Execute `go doc` no seu código para ver a documentação gerada

**Dica**: Use `go doc .CalculaIMC` no diretório do arquivo para ver sua documentação.

---

### Exercício 2: Explorar a Biblioteca Padrão com `go doc`

Use o comando `go doc` para explorar e responder:

1. **Pacote `strings`:**

   - Execute `go doc strings`
   - Liste 3 funções que você achou interessantes
   - Para cada função, execute `go doc strings.NomeDaFuncao` e anote o que ela faz

2. **Pacote `math`:**

   - Execute `go doc math`
   - Encontre uma função que calcula potência
   - Use `go doc` para ver a documentação completa dessa função
   - Anote a sintaxe e um exemplo de uso

3. **Pacote `time`:**
   - Execute `go doc time`
   - Encontre como obter a hora atual
   - Use `go doc` para ver a documentação
   - Anote como usar

**Crie um arquivo `exercicio2.md`** com suas descobertas, incluindo:

- Nome das funções encontradas
- O que cada uma faz (da documentação)
- Exemplo de como usar

---

### Exercício 3: Criar um Pacote Bem Documentado

Crie um pacote chamado `geometria` com as seguintes funções, todas bem documentadas:

**Arquivo: `geometria/geometria.go`**

```go
package geometria

import "math"

// TODO: Documente todas as funções abaixo seguindo as boas práticas

func AreaRetangulo(largura, altura float64) float64 {
    return largura * altura
}

func PerimetroRetangulo(largura, altura float64) float64 {
    return 2 * (largura + altura)
}

func AreaCirculo(raio float64) float64 {
    return math.Pi * raio * raio
}

func PerimetroCirculo(raio float64) float64 {
    return 2 * math.Pi * raio
}

func AreaTriangulo(base, altura float64) float64 {
    return (base * altura) / 2
}
```

**Tarefas:**

1. Adicione documentação completa para cada função
2. Inclua documentação do pacote no início
3. Adicione exemplos de uso onde apropriado
4. Documente parâmetros e retornos
5. Teste a documentação com `go doc geometria`
6. Teste funções específicas com `go doc geometria.AreaRetangulo`

**Estrutura esperada:**

```
projeto/
  geometria/
    geometria.go
  main.go (para testar)
```

---

### Exercício 4: Usar `go help` e `godoc`

**Parte 1: Explorar `go help`**

1. Execute `go help` e liste todos os comandos disponíveis
2. Escolha 3 comandos que você ainda não conhece bem
3. Para cada um, execute `go help nomeDoComando`
4. Crie um resumo explicando o que cada comando faz

**Parte 2: Configurar e Usar `godoc`**

1. Instale o `godoc` (se ainda não tiver):

   ```bash
   go install golang.org/x/tools/cmd/godoc@latest
   ```

2. Inicie o servidor:

   ```bash
   godoc -http=:6060
   ```

3. Acesse `http://localhost:6060` no navegador

4. Explore:

   - Navegue até o pacote `fmt`
   - Veja a documentação de `fmt.Println`
   - Clique em "Source" para ver o código-fonte
   - Explore outros pacotes que interessam você

5. Se você criou o pacote `geometria` do exercício anterior:
   - Navegue até ele no `godoc`
   - Veja como sua documentação aparece na interface web

**Crie um arquivo `exercicio4.md`** documentando:

- Comandos descobertos e suas funções
- Suas impressões sobre o `godoc`
- Diferenças entre `go doc` (terminal) e `godoc` (web)

---

## Perguntas de Reflexão

### Reflexão 1: Importância da Documentação

Você aprendeu que documentação é essencial para código profissional.

**Pergunta:**

1. Pense em uma situação real onde você precisou usar código que não estava documentado (pode ser código seu antigo ou de outra pessoa). Como foi a experiência?
2. Agora pense em como seria se esse código estivesse bem documentado. O que mudaria?
3. Por que você acha que muitos programadores negligenciam a documentação, mesmo sabendo que é importante?
4. Como você pode criar o hábito de documentar seu código desde o início?

**Sua resposta deve incluir**: Experiências pessoais, análise crítica sobre a importância da documentação, e estratégias para desenvolver o hábito de documentar.

---

### Reflexão 2: `go doc` vs. Busca na Internet

Você aprendeu a usar `go doc` para explorar a biblioteca padrão do Go.

**Pergunta:**

1. Compare usar `go doc` com fazer uma busca no Google sobre funções do Go. Quais são as vantagens e desvantagens de cada abordagem?
2. Em que situações você preferiria usar `go doc`?
3. Em que situações você preferiria buscar na internet?
4. Como `go doc` pode ajudar você a se tornar um programador Go mais independente?

**Sua resposta deve demonstrar**: Capacidade de comparar ferramentas, entendimento de quando usar cada uma, e visão sobre aprendizado autônomo.

---

### Reflexão 3: Documentação como Interface

A documentação de funções públicas é a "interface" que outros desenvolvedores veem do seu código.

**Pergunta:**

1. Se você criasse uma biblioteca Go para outros desenvolvedores usarem, o que você incluiria na documentação para torná-la útil?
2. Pense em uma função complexa que você já escreveu (ou imagine uma). Como você documentaria ela de forma que alguém que nunca viu o código conseguisse usar corretamente?
3. Qual é a diferença entre "código auto-explicativo" e código que precisa de documentação? Dê exemplos.
4. Quando você acha que é "demais" documentar? Existe um ponto onde documentação excessiva pode ser ruim?

**Sua resposta deve incluir**: Exemplos concretos, análise sobre o equilíbrio na documentação, e pensamento crítico sobre boas práticas.

---

### Reflexão 4: Ferramentas e Produtividade

Você conheceu várias ferramentas: `go doc`, `go help`, `godoc`.

**Pergunta:**

1. Como essas ferramentas podem melhorar sua produtividade como desenvolvedor?
2. Pense em um problema de programação que você enfrentou recentemente. Como essas ferramentas poderiam ter ajudado você a resolver mais rápido?
3. Qual ferramenta você acha mais útil e por quê?
4. Como você incorporaria o uso dessas ferramentas no seu fluxo de trabalho diário?

**Sua resposta deve demonstrar**: Capacidade de aplicar ferramentas em situações reais, análise de produtividade, e planejamento de fluxo de trabalho.

---

### Reflexão 5: Aprendizado Ativo vs. Passivo

Usar `go doc` e `godoc` é uma forma de **aprendizado ativo** - você explora e descobre por conta própria.

**Pergunta:**

1. Compare aprender Go apenas lendo tutoriais (aprendizado passivo) com explorar usando `go doc` (aprendizado ativo). Quais são as diferenças?
2. Por que o aprendizado ativo pode ser mais eficaz?
3. Dê um exemplo de como você usaria `go doc` para aprender sobre um tópico novo em Go (ex: canais, goroutines, interfaces).
4. Como você pode usar essas ferramentas para se manter atualizado com novas funcionalidades do Go?

**Sua resposta deve incluir**: Comparação entre métodos de aprendizado, exemplos práticos de exploração, e estratégias de aprendizado contínuo.

---

## Como Entregar

Crie os seguintes arquivos na pasta `05-documentacao-e-comandos/`:

1. **`exercicio1.go`** - Função `CalculaIMC` documentada
2. **`exercicio2.md`** - Descobertas da exploração da biblioteca padrão
3. **`geometria/geometria.go`** - Pacote de geometria documentado
4. **`exercicio4.md`** - Análise de `go help` e `godoc`
5. **`reflexoes.md`** - Respostas às 5 perguntas de reflexão

**Importante**:

- Compile e teste todos os códigos para garantir que funcionam
- Execute `go doc` em seus próprios códigos para verificar a documentação
- Seja honesto nas reflexões - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão
- Use exemplos reais das suas experiências quando possível

---

## Dicas para os Exercícios

### Dica para Exercício 1:

```bash
# Depois de documentar, teste assim:
go doc .CalculaIMC

# Ou se estiver em um pacote:
go doc CalculaIMC
```

### Dica para Exercício 2:

```bash
# Explore de forma interativa
go doc strings
go doc strings.Contains
go doc strings.Split
go doc strings.Join
```

### Dica para Exercício 3:

```bash
# Estrutura de diretórios:
mkdir geometria
cd geometria
# Crie geometria.go aqui

# Depois, no diretório pai:
go doc geometria
go doc geometria.AreaRetangulo
```

### Dica para Exercício 4:

```bash
# Ver todos os comandos
go help

# Ver ajuda específica
go help build
go help test
go help mod

# Iniciar godoc (em outro terminal)
godoc -http=:6060
# Acesse http://localhost:6060
```

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!
