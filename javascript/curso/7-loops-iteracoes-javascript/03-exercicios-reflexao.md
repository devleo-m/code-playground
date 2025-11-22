# Aula 7 - Exercícios e Reflexão: Loops e Iterações

## 📝 Exercícios Práticos

### Exercício 1: Loop `for` Básico - Contagem

Crie um loop `for` que:
- Conte de 1 até 10
- Imprima cada número no console
- Ao final, imprima "Contagem concluída!"

**Sua tarefa:** Escreva o código completo e execute no console do navegador.

**Desafio extra:** Modifique o código para contar de 10 até 1 (ordem decrescente).

---

### Exercício 2: Soma de Números em um Array

Crie uma função chamada `somarArray` que:
- Recebe um array de números como parâmetro
- Usa um loop `for` para somar todos os números
- Retorna a soma total

**Exemplo de uso:**
```javascript
console.log(somarArray([1, 2, 3, 4, 5]));        // 15
console.log(somarArray([10, 20, 30]));          // 60
console.log(somarArray([-5, 10, -3, 8]));       // 10
```

**Sua tarefa:** 
1. Escreva a função usando loop `for` tradicional
2. Reescreva usando `for...of`
3. Compare as duas versões - qual você acha mais legível?

---

### Exercício 3: Encontrar o Maior Número

Crie uma função chamada `encontrarMaior` que:
- Recebe um array de números
- Usa um loop para encontrar o maior número
- Retorna o maior número encontrado

**Exemplo de uso:**
```javascript
console.log(encontrarMaior([3, 7, 2, 9, 1]));   // 9
console.log(encontrarMaior([-5, -2, -10]));     // -2
console.log(encontrarMaior([42]));              // 42
```

**Sua tarefa:** Escreva a função e teste com diferentes arrays. O que acontece se o array estiver vazio?

---

### Exercício 4: Loop `while` - Adivinhação

Crie um jogo de adivinhação usando `while`:
- Gere um número aleatório entre 1 e 10
- Peça ao usuário para adivinhar (use `prompt`)
- Continue pedindo até o usuário acertar
- Conte quantas tentativas foram necessárias

**Exemplo de execução:**
```javascript
// Número secreto: 7
// Usuário digita: 3 → "Muito baixo! Tente novamente."
// Usuário digita: 9 → "Muito alto! Tente novamente."
// Usuário digita: 7 → "Parabéns! Você acertou em 3 tentativas!"
```

**Sua tarefa:** Escreva o código completo do jogo.

**Dica:** Use `Math.floor(Math.random() * 10) + 1` para gerar números aleatórios.

---

### Exercício 5: Loop `do...while` - Validação de Entrada

Crie uma função chamada `solicitarIdade` que:
- Usa `do...while` para pedir a idade do usuário
- Continua pedindo até que a idade seja um número válido (entre 0 e 120)
- Retorna a idade válida

**Exemplo de uso:**
```javascript
// Se usuário digitar "abc" → "Idade inválida! Digite um número entre 0 e 120."
// Se usuário digitar 150 → "Idade inválida! Digite um número entre 0 e 120."
// Se usuário digitar 25 → Retorna 25
```

**Sua tarefa:** Escreva a função completa com validação adequada.

---

### Exercício 6: Loop `for...of` - Processar Lista de Compras

Crie uma função chamada `processarListaCompras` que:
- Recebe um array de objetos com `item` e `preco`
- Usa `for...of` para processar cada item
- Calcula o total da compra
- Retorna um objeto com `total` e `quantidadeItens`

**Exemplo de uso:**
```javascript
const compras = [
  { item: 'leite', preco: 5.50 },
  { item: 'pão', preco: 3.00 },
  { item: 'ovos', preco: 8.00 }
];

console.log(processarListaCompras(compras));
// { total: 16.50, quantidadeItens: 3 }
```

**Sua tarefa:** Escreva a função usando `for...of`.

---

### Exercício 7: Loop `for...in` - Informações de Pessoa

Crie uma função chamada `exibirInformacoes` que:
- Recebe um objeto representando uma pessoa
- Usa `for...in` para iterar sobre as propriedades
- Retorna uma string formatada com todas as informações

**Exemplo de uso:**
```javascript
const pessoa = {
  nome: 'Maria',
  idade: 30,
  cidade: 'São Paulo',
  profissao: 'Desenvolvedora'
};

console.log(exibirInformacoes(pessoa));
// "nome: Maria\nidade: 30\ncidade: São Paulo\nprofissao: Desenvolvedora"
```

**Sua tarefa:** Escreva a função e teste com diferentes objetos.

---

### Exercício 8: Usando `break` - Encontrar Primeiro Par

Crie uma função chamada `encontrarPrimeiroPar` que:
- Recebe um array de números
- Usa um loop para encontrar o primeiro número par
- Usa `break` para parar assim que encontrar
- Retorna o número par encontrado, ou `null` se não houver

**Exemplo de uso:**
```javascript
console.log(encontrarPrimeiroPar([1, 3, 5, 8, 9]));    // 8
console.log(encontrarPrimeiroPar([1, 3, 5, 7]));      // null
console.log(encontrarPrimeiroPar([2, 4, 6]));         // 2
```

**Sua tarefa:** Escreva a função usando `break` para otimizar a busca.

---

### Exercício 9: Usando `continue` - Filtrar Números

Crie uma função chamada `filtrarPositivos` que:
- Recebe um array de números
- Usa um loop com `continue` para pular números negativos ou zero
- Retorna um novo array apenas com números positivos

**Exemplo de uso:**
```javascript
console.log(filtrarPositivos([-5, 2, -3, 0, 8, -1, 10]));  // [2, 8, 10]
console.log(filtrarPositivos([-1, -2, -3]));               // []
console.log(filtrarPositivos([1, 2, 3]));                  // [1, 2, 3]
```

**Sua tarefa:** Escreva a função usando `continue` para pular valores indesejados.

---

### Exercício 10: Loops Aninhados - Tabela de Multiplicação

Crie uma função chamada `gerarTabelaMultiplicacao` que:
- Recebe um número `n` como parâmetro
- Usa loops aninhados para gerar uma tabela de multiplicação de 1 até `n`
- Retorna um array de arrays (matriz) representando a tabela

**Exemplo de uso:**
```javascript
console.log(gerarTabelaMultiplicacao(3));
// [
//   [1, 2, 3],
//   [2, 4, 6],
//   [3, 6, 9]
// ]
```

**Sua tarefa:** Escreva a função usando loops aninhados.

**Desafio extra:** Modifique para imprimir a tabela de forma formatada no console.

---

### Exercício 11: Análise de Código - Identificar Problemas

Analise o código abaixo e identifique os problemas:

```javascript
// Código 1
const numeros = [1, 2, 3, 4, 5];
for (const numero in numeros) {
  console.log(numero);
}

// Código 2
let i = 0;
while (i < 10) {
  console.log(i);
  // Esqueceu de incrementar!
}

// Código 3
const array = [1, 2, 3];
for (let i = 0; i < array.length; i++) {
  if (array[i] === 2) {
    array.splice(i, 1);
  }
  console.log(array[i]);
}

// Código 4
const objeto = { a: 1, b: 2, c: 3 };
for (const valor of objeto) {
  console.log(valor);
}
```

**Sua tarefa:**
1. Identifique o problema em cada código
2. Explique por que é um problema
3. Corrija cada código

---

### Exercício 12: Função Completa - Processar Dados de Vendas

Crie uma função chamada `processarVendas` que:
- Recebe um array de objetos de vendas, cada um com `produto`, `quantidade`, `preco`
- Usa um loop apropriado para processar cada venda
- Calcula o total de cada venda (quantidade × preço)
- Retorna um objeto com:
  - `vendas`: array com cada venda e seu total
  - `totalGeral`: soma de todas as vendas
  - `quantidadeProdutos`: número total de produtos vendidos

**Exemplo de uso:**
```javascript
const vendas = [
  { produto: 'Notebook', quantidade: 2, preco: 3000 },
  { produto: 'Mouse', quantidade: 5, preco: 50 },
  { produto: 'Teclado', quantidade: 3, preco: 150 }
];

console.log(processarVendas(vendas));
// {
//   vendas: [
//     { produto: 'Notebook', quantidade: 2, preco: 3000, total: 6000 },
//     { produto: 'Mouse', quantidade: 5, preco: 50, total: 250 },
//     { produto: 'Teclado', quantidade: 3, preco: 150, total: 450 }
//   ],
//   totalGeral: 6700,
//   quantidadeProdutos: 10
// }
```

**Sua tarefa:** Escreva a função completa usando o loop mais apropriado.

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Performance e Eficiência

Analise os dois códigos abaixo:

```javascript
// Código A
const array = [1, 2, 3, 4, 5];
for (let i = 0; i < array.length; i++) {
  console.log(array[i]);
}

// Código B
const array = [1, 2, 3, 4, 5];
for (let i = 0, len = array.length; i < len; i++) {
  console.log(array[i]);
}
```

**Perguntas:**
1. Qual código é mais eficiente? Por quê?
2. Qual é a diferença de performance entre eles?
3. Em que situações essa diferença seria significativa?
4. Existe uma forma ainda mais eficiente de iterar sobre este array?

---

### Reflexão 2: Escolha do Loop Correto

Considere os seguintes cenários:

**Cenário A:** Você precisa processar todos os elementos de um array de 1000 itens.

**Cenário B:** Você precisa processar propriedades de um objeto com informações do usuário.

**Cenário C:** Você precisa executar uma ação até que o usuário digite "sair".

**Cenário D:** Você precisa validar uma entrada do usuário, garantindo que execute pelo menos uma vez.

**Perguntas:**
1. Qual loop você escolheria para cada cenário? Por quê?
2. Quais são as implicações de escolher o loop errado?
3. Como a escolha do loop afeta a legibilidade do código?
4. Existe uma situação onde múltiplos loops seriam igualmente apropriados?

---

### Reflexão 3: Modificação de Arrays Durante Iteração

Analise este código:

```javascript
const numeros = [1, 2, 3, 4, 5];

for (let i = 0; i < numeros.length; i++) {
  if (numeros[i] % 2 === 0) {
    numeros.splice(i, 1);
  }
}

console.log(numeros); // Qual será o resultado?
```

**Perguntas:**
1. Qual será o resultado final do array `numeros`?
2. Por que esse resultado acontece?
3. Quais são os problemas de modificar um array durante a iteração?
4. Como você resolveria esse problema de forma segura?
5. Existem alternativas melhores para remover elementos de um array?

---

### Reflexão 4: Loops Infinitos e Segurança

Considere este código:

```javascript
let contador = 0;
while (contador < 10) {
  console.log(contador);
  // Esqueceu de incrementar contador
}
```

**Perguntas:**
1. O que acontece quando este código é executado?
2. Como você detectaria esse problema em um código maior?
3. Quais são as melhores práticas para evitar loops infinitos?
4. Como você protegeria seu código contra loops infinitos acidentais?
5. Em um ambiente de produção, quais seriam as consequências de um loop infinito?

---

### Reflexão 5: Performance em Loops Aninhados

Analise este código:

```javascript
const matriz1 = [[1, 2, 3], [4, 5, 6], [7, 8, 9]];
const matriz2 = [[9, 8, 7], [6, 5, 4], [3, 2, 1]];

for (let i = 0; i < matriz1.length; i++) {
  for (let j = 0; j < matriz1[i].length; j++) {
    console.log(matriz1[i][j] + matriz2[i][j]);
  }
}
```

**Perguntas:**
1. Quantas iterações este código executa no total?
2. Se cada matriz tivesse 1000 linhas e 1000 colunas, quantas iterações seriam?
3. Como a complexidade de tempo cresce com o tamanho das matrizes?
4. Quais são as implicações de performance para loops aninhados profundos?
5. Existem alternativas mais eficientes para processar matrizes grandes?

---

### Reflexão 6: Uso de `break` e `continue`

Considere estes dois códigos que fazem a mesma coisa:

```javascript
// Código A - Usando break
const numeros = [1, 2, 3, 4, 5];
let encontrado = false;
for (const numero of numeros) {
  if (numero === 3) {
    encontrado = true;
    break;
  }
}

// Código B - Sem break
const numeros = [1, 2, 3, 4, 5];
let encontrado = false;
for (const numero of numeros) {
  if (!encontrado && numero === 3) {
    encontrado = true;
  }
}
```

**Perguntas:**
1. Qual código é mais eficiente? Por quê?
2. Qual código é mais legível?
3. Em que situações usar `break` é apropriado?
4. Em que situações usar `break` pode ser considerado uma "má prática"?
5. Como `break` e `continue` afetam a manutenibilidade do código?

---

### Reflexão 7: Edge Cases e Validação

Considere esta função:

```javascript
function somarArray(array) {
  let soma = 0;
  for (let i = 0; i < array.length; i++) {
    soma += array[i];
  }
  return soma;
}
```

**Perguntas:**
1. O que acontece se `array` for `null` ou `undefined`?
2. O que acontece se `array` for vazio `[]`?
3. O que acontece se `array` contiver valores não numéricos?
4. O que acontece se `array` contiver `NaN`?
5. Como você modificaria a função para lidar com esses edge cases?
6. Qual é a importância de considerar edge cases em loops?

---

### Reflexão 8: Legibilidade vs Performance

Compare estas três formas de iterar sobre um array:

```javascript
// Forma 1 - for tradicional
for (let i = 0; i < array.length; i++) {
  console.log(array[i]);
}

// Forma 2 - for...of
for (const elemento of array) {
  console.log(elemento);
}

// Forma 3 - forEach (método de array)
array.forEach(elemento => {
  console.log(elemento);
});
```

**Perguntas:**
1. Qual forma é mais legível? Por quê?
2. Qual forma tem melhor performance? (pesquise se necessário)
3. Em que situações cada forma seria mais apropriada?
4. Como você equilibra legibilidade e performance na escolha de loops?
5. A diferença de performance é sempre significativa o suficiente para justificar código menos legível?

---

## 📋 Checklist de Entrega

Antes de enviar suas respostas, verifique:

- [ ] Completei todos os 12 exercícios práticos
- [ ] Testei cada função no console do navegador
- [ ] Respondi todas as 8 perguntas de reflexão
- [ ] Expliquei meu raciocínio nas respostas
- [ ] Identifiquei e corrigi problemas nos exercícios de análise de código
- [ ] Considerei edge cases nas minhas soluções
- [ ] Revisei meu código para garantir que está correto

---

## 🎯 Objetivos dos Exercícios

Estes exercícios foram projetados para:

1. **Praticar** todos os tipos de loops aprendidos
2. **Aplicar** loops em situações práticas do dia a dia
3. **Refletir** sobre performance, eficiência e boas práticas
4. **Identificar** problemas comuns e como evitá-los
5. **Desenvolver** pensamento crítico sobre escolhas de código

---

## 💡 Dicas para Resolução

1. **Teste no Console**: Sempre teste seu código no console do navegador
2. **Comece Simples**: Resolva a versão básica primeiro, depois adicione complexidade
3. **Pense em Edge Cases**: O que acontece com arrays vazios? Com valores nulos?
4. **Compare Soluções**: Tente resolver o mesmo problema com diferentes tipos de loops
5. **Leia Erros**: Se algo der errado, leia a mensagem de erro cuidadosamente

---

## 🚀 Próximos Passos

Após completar estes exercícios e reflexões, você estará pronto para:
- Aprender sobre performance e otimização de loops
- Aplicar loops em manipulação do DOM
- Trabalhar com métodos de array (map, filter, reduce)
- Entender programação assíncrona

**Lembre-se**: A prática constante é essencial. Não tenha pressa - entenda cada conceito antes de avançar!

Boa sorte! 🎓

