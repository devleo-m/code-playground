# Módulo 3: Tipos de Dados em Go
## Aula 3 - Exercícios e Reflexão

Agora é hora de colocar a mão na massa! Complete os exercícios abaixo e reflita sobre as questões propostas.

---

## Exercícios Práticos

### Exercício 1: Declaração e Uso de Tipos Básicos

Crie um programa que declare variáveis dos seguintes tipos e imprima seus valores:

- Um `int8` com valor -50
- Um `uint16` com valor 500
- Um `float64` com valor 3.14159
- Um `bool` com valor `true`
- Um `rune` com o caractere 'G'
- Uma string interpretada com a mensagem "Olá,\nMundo!"
- Uma string raw com o caminho `C:\Users\Documentos\arquivo.txt`

**Dica**: Use `fmt.Printf` com verbos de formatação (`%d` para inteiros, `%f` para floats, `%t` para bool, `%c` para rune, `%s` para strings).

---

### Exercício 2: Conversão de Tipos

Crie um programa que:

1. Declare uma variável `numeroInt` do tipo `int` com valor 42
2. Converta esse valor para `float64` e armazene em `numeroFloat`
3. Converta o `float64` de volta para `int` e armazene em `numeroInt2`
4. Imprima os três valores e seus tipos usando `fmt.Printf` com `%T` para mostrar o tipo

**Pergunta para pensar**: O que acontece se você converter `3.9` (float64) para `int`? Teste e explique o resultado.

---

### Exercício 3: Trabalhando com Runes e Strings

Crie um programa que:

1. Declare uma string com o texto: "Olá, 世界! 🚀"
2. Use um loop `for range` para iterar sobre cada caractere
3. Para cada caractere, imprima:
   - O caractere em si
   - Seu código Unicode (rune)
   - Se é um caractere ASCII (código < 128) ou não

**Dica**: Use `fmt.Printf("%c = %d (ASCII: %v)\n", char, char, char < 128)`

---

### Exercício 4: Comparando Raw e Interpreted Strings

Crie um programa que demonstre a diferença entre raw strings e interpreted strings:

1. Declare uma variável `raw` usando backticks com o conteúdo: `Linha 1\nLinha 2\tTab aqui`
2. Declare uma variável `interpreted` usando aspas duplas com o mesmo conteúdo: `"Linha 1\nLinha 2\tTab aqui"`
3. Imprima ambas as strings
4. Explique a diferença no output

---

## Perguntas de Reflexão

### Reflexão 1: Escolha de Tipos Inteiros

Imagine que você está criando um sistema para uma escola que precisa armazenar:
- A idade de cada aluno (0 a 120 anos)
- O número de matrícula (pode chegar a milhões)
- A nota de uma prova (0 a 100)
- A temperatura da sala em graus Celsius (pode ser negativa no inverno)

**Pergunta**: Para cada um desses dados, qual tipo inteiro (int8, int16, int32, int64, uint8, uint16, uint32, uint64, ou int/uint) você escolheria e **por quê**? Considere:
- O range necessário
- Economia de memória
- Facilidade de uso

**Sua resposta deve incluir**: A escolha para cada caso E a justificativa técnica por trás de cada escolha.

---

### Reflexão 2: Precisão de Ponto Flutuante e Aplicações Reais

Você aprendeu que `float32` e `float64` podem ter erros de precisão e não são adequados para cálculos financeiros.

**Pergunta**: 
1. Explique **com suas próprias palavras** por que floats têm erros de precisão. Use uma analogia se ajudar.
2. Dê **três exemplos reais** de situações onde você usaria floats (e está correto usar).
3. Dê **três exemplos reais** de situações onde você **NÃO** deveria usar floats e explique qual alternativa usaria.

**Sua resposta deve demonstrar**: Compreensão profunda do problema de precisão e capacidade de aplicar esse conhecimento em cenários práticos.

---

## Como Entregar

Crie arquivos `.go` separados para cada exercício (ex: `exercicio1.go`, `exercicio2.go`, etc.) na pasta `03-data-types/`. Para as perguntas de reflexão, você pode criar um arquivo `reflexoes.md` ou simplesmente responder diretamente aqui.

**Importante**: 
- Compile e execute cada programa para garantir que funciona
- Comente seu código explicando o que cada parte faz
- Seja honesto nas reflexões - não há resposta "errada", mas há respostas que demonstram mais ou menos compreensão

---

Após completar os exercícios e reflexões, envie suas respostas para que eu possa analisar seu desempenho e fornecer feedback construtivo!







