# Aula 3: Exercícios e Reflexão - Methods vs Functions

Olá! Agora é hora de colocar a mão na massa e praticar o que aprendemos sobre methods e functions. Vamos fazer alguns exercícios práticos e depois refletir sobre os conceitos.

---

## Exercício 1: Criando Methods Básicos

**Objetivo**: Criar métodos para um tipo `Livro` que represente um livro em uma biblioteca.

### Tarefa

Crie um tipo `Livro` com os campos:
- `Titulo` (string)
- `Autor` (string)
- `Paginas` (int)
- `Lido` (bool)

Implemente os seguintes métodos:

1. `Info()` - Retorna uma string formatada com todas as informações do livro (value receiver)
2. `MarcarComoLido()` - Marca o livro como lido (pointer receiver)
3. `TempoLeitura()` - Calcula o tempo estimado de leitura (assumindo 2 minutos por página) (value receiver)
4. `AdicionarPaginas()` - Adiciona páginas ao livro (pointer receiver)

### Exemplo de Uso Esperado

```go
livro := Livro{
    Titulo: "Aprendendo Go",
    Autor: "João Silva",
    Paginas: 200,
    Lido: false,
}

fmt.Println(livro.Info())
// Saída esperada: "Aprendendo Go - João Silva (200 páginas) - Não lido"

livro.MarcarComoLido()
fmt.Println(livro.Info())
// Saída esperada: "Aprendendo Go - João Silva (200 páginas) - Lido"

tempo := livro.TempoLeitura()
fmt.Printf("Tempo estimado: %d minutos\n", tempo)
// Saída esperada: "Tempo estimado: 400 minutos"

livro.AdicionarPaginas(50)
fmt.Println(livro.Info())
// Saída esperada: "Aprendendo Go - João Silva (250 páginas) - Lido"
```

---

## Exercício 2: Value vs Pointer Receivers

**Objetivo**: Entender quando usar value receiver e quando usar pointer receiver.

### Tarefa

Crie um tipo `Contador` com um campo `valor` (int). Implemente:

1. `Valor()` - Retorna o valor atual (value receiver)
2. `Incrementar()` - Incrementa o contador em 1 (pointer receiver)
3. `Resetar()` - Zera o contador (pointer receiver)
4. `Duplicar()` - Retorna um novo contador com o dobro do valor, sem modificar o original (value receiver)

### Desafio Adicional

Crie um método `Copiar()` que retorna uma cópia do contador. Qual tipo de receiver você usaria? Por quê?

### Exemplo de Uso

```go
contador := Contador{valor: 5}

fmt.Println(contador.Valor())  // 5

contador.Incrementar()
fmt.Println(contador.Valor())  // 6

duplicado := contador.Duplicar()
fmt.Println(contador.Valor())   // 6 (não mudou)
fmt.Println(duplicado.Valor())  // 12

contador.Resetar()
fmt.Println(contador.Valor())  // 0
```

---

## Exercício 3: Methods em Tipos Customizados

**Objetivo**: Praticar métodos em tipos não-struct.

### Tarefa

Crie um tipo customizado `Temperatura` baseado em `float64` que representa temperatura em Celsius. Implemente:

1. `Fahrenheit()` - Converte para Fahrenheit (value receiver)
2. `Kelvin()` - Converte para Kelvin (value receiver)
3. `Formatar()` - Retorna string formatada como "25.5°C" (value receiver)

**Fórmulas**:
- Fahrenheit = (Celsius × 9/5) + 32
- Kelvin = Celsius + 273.15

### Exemplo de Uso

```go
temp := Temperatura(25.5)

fmt.Println(temp.Formatar())        // "25.5°C"
fmt.Printf("%.2f°F\n", temp.Fahrenheit())  // "77.90°F"
fmt.Printf("%.2fK\n", temp.Kelvin())       // "298.65K"
```

---

## Exercício 4: Methods vs Functions - Decisão de Design

**Objetivo**: Decidir quando usar method e quando usar function.

### Tarefa

Você precisa implementar um sistema de cálculo de área para diferentes formas geométricas. Analise cada caso e decida se deve ser um **method** ou uma **function**. Justifique sua escolha.

1. Calcular área de um retângulo
2. Comparar se dois retângulos têm a mesma área
3. Calcular área de um círculo
4. Verificar se um ponto está dentro de um retângulo
5. Calcular área total de uma lista de retângulos

**Implemente pelo menos 3 delas** (escolha as que você achar mais interessantes) e explique por que escolheu method ou function para cada uma.

---

## Perguntas de Reflexão

Agora vamos pensar mais profundamente sobre os conceitos. Responda as seguintes perguntas com suas próprias palavras:

### Reflexão 1: Por Que Methods?

**Pergunta**: Por que Go permite definir methods fora da declaração do tipo, ao invés de dentro como em outras linguagens? Qual você acha que é a vantagem dessa abordagem?

**Dica para pensar**: Considere como isso afeta a organização do código, a capacidade de adicionar métodos a tipos de bibliotecas externas, e a flexibilidade do design.

---

### Reflexão 2: Value vs Pointer - Performance e Semântica

**Pergunta**: Além da questão de modificar ou não o valor, quando mais você deveria usar pointer receiver mesmo que o método não modifique nada? Pense em termos de performance e semântica do código.

**Dica para pensar**: 
- O que acontece quando você passa uma struct muito grande por valor?
- O que significa para outros desenvolvedores quando veem um método com pointer receiver?
- Qual é o custo de copiar uma struct pequena vs uma struct grande?

---

### Reflexão 3: Consistência e Manutenibilidade

**Pergunta**: Por que é importante manter consistência entre os receivers de um tipo? Ou seja, por que se um método usa pointer receiver, geralmente todos os métodos do tipo deveriam usar pointer receiver?

**Dica para pensar**:
- Como isso afeta a legibilidade do código?
- O que acontece quando você mistura value e pointer receivers no mesmo tipo?
- Como isso impacta a manutenção do código a longo prazo?

---

### Reflexão 4: Methods vs Functions - Filosofia de Design

**Pergunta**: Em Go, methods e functions podem fazer coisas muito similares. Quando você está projetando uma API, quais critérios você usaria para decidir entre criar um method ou uma function? Pense além do "pertence ao tipo ou não".

**Dica para pensar**:
- Como a escolha afeta a usabilidade da API?
- O que é mais idiomático em Go?
- Como isso se relaciona com interfaces (que veremos depois)?
- Qual torna o código mais legível e intuitivo?

---

## Desafio Opcional: Sistema de Estoque

Crie um sistema simples de estoque com os seguintes requisitos:

1. Tipo `Produto` com: Nome, Preço, Quantidade em estoque
2. Métodos para:
   - Adicionar estoque (pointer receiver)
   - Remover estoque (pointer receiver) - retornar erro se não houver estoque suficiente
   - Calcular valor total do estoque (value receiver)
   - Aplicar desconto percentual (pointer receiver)
3. Tipo `Estoque` (slice de Produto) com métodos para:
   - Adicionar produto
   - Buscar produto por nome
   - Calcular valor total de todo o estoque
   - Listar produtos com estoque baixo (menos de 10 unidades)

Implemente tudo usando methods e explique suas escolhas de receivers.

---

## Como Enviar Suas Respostas

Quando terminar os exercícios e reflexões, envie:

1. **Código dos exercícios**: Os arquivos `.go` com suas implementações
2. **Respostas das reflexões**: Suas respostas escritas para as 4 perguntas de reflexão
3. **Dúvidas**: Qualquer dúvida que surgiu durante os exercícios

Lembre-se: O objetivo não é apenas fazer funcionar, mas **entender o porquê** de cada decisão. Seja honesto sobre o que você entendeu e o que ainda tem dúvidas.

Boa sorte e bons estudos! 🚀

