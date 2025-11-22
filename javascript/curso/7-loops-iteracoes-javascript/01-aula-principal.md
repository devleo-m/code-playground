# Aula 7: Loops e Iterações em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 6**, você aprendeu:
- ✅ Operadores de comparação de igualdade (`==`, `===`, `Object.is()`)
- ✅ Diferenças entre igualdade abstrata e estrita
- ✅ Type coercion e suas implicações
- ✅ Quando usar cada tipo de comparação

Agora vamos aprender sobre **loops e iterações** - uma das ferramentas mais poderosas para repetir ações e processar dados em JavaScript!

---

## 🎯 O que são Loops e Iterações?

**Definição:** Loops são estruturas de controle que permitem executar um bloco de código **repetidamente** enquanto uma condição específica for verdadeira. Eles são fundamentais para processar coleções de dados, repetir ações e automatizar tarefas.

### Conceitos Fundamentais

1. **Iteração**: Cada execução do bloco de código dentro do loop
2. **Condição de Parada**: A condição que determina quando o loop deve parar
3. **Contador/Índice**: Variável que controla o número de iterações
4. **Loop Infinito**: Um loop que nunca para (geralmente um erro)

### Por que Usar Loops?

Imagine que você precisa:
- Imprimir números de 1 a 100
- Processar todos os itens de uma lista de compras
- Validar todos os campos de um formulário
- Calcular a soma de todos os números em um array

Sem loops, você teria que escrever o mesmo código centenas de vezes. Com loops, você escreve uma vez e o código se repete automaticamente.

---

## 🔄 1. O Loop `for`

### Definição

O loop `for` é o tipo de loop mais comum e versátil em JavaScript. Ele é ideal quando você **sabe quantas vezes** deseja repetir uma ação.

### Sintaxe

```javascript
for (inicialização; condição; incremento) {
  // código a ser executado
}
```

### Componentes do Loop `for`

1. **Inicialização**: Executada uma vez antes do loop começar. Geralmente declara uma variável contadora.
2. **Condição**: Avaliada antes de cada iteração. Se for `true`, o loop continua; se for `false`, o loop para.
3. **Incremento**: Executada após cada iteração. Geralmente incrementa ou decrementa o contador.
4. **Corpo do Loop**: O código que será executado em cada iteração.

### Exemplos de Uso

#### 1.1. Loop `for` Básico

```javascript
// Contar de 0 a 4
for (let i = 0; i < 5; i++) {
  console.log('Iteração número:', i);
}

// Saída:
// Iteração número: 0
// Iteração número: 1
// Iteração número: 2
// Iteração número: 3
// Iteração número: 4
```

**Explicação linha por linha:**
- `let i = 0`: Declara e inicializa a variável `i` com valor 0
- `i < 5`: Condição - continua enquanto `i` for menor que 5
- `i++`: Incrementa `i` em 1 após cada iteração
- O código dentro das chaves executa 5 vezes (i = 0, 1, 2, 3, 4)

#### 1.2. Contar em Ordem Decrescente

```javascript
// Contar de 10 até 1
for (let i = 10; i >= 1; i--) {
  console.log(i);
}

// Saída: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
```

#### 1.3. Incrementos Personalizados

```javascript
// Contar de 0 a 20, de 2 em 2
for (let i = 0; i <= 20; i += 2) {
  console.log(i);
}

// Saída: 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20

// Contar de 100 a 0, de 10 em 10
for (let i = 100; i >= 0; i -= 10) {
  console.log(i);
}

// Saída: 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0
```

#### 1.4. Iterar sobre Arrays

```javascript
const frutas = ['maçã', 'banana', 'laranja', 'uva'];

for (let i = 0; i < frutas.length; i++) {
  console.log(`Fruta ${i + 1}: ${frutas[i]}`);
}

// Saída:
// Fruta 1: maçã
// Fruta 2: banana
// Fruta 3: laranja
// Fruta 4: uva
```

**Importante**: Use `i < array.length` (não `i <= array.length`) porque arrays começam no índice 0.

#### 1.5. Processar Arrays com Índices

```javascript
const numeros = [10, 20, 30, 40, 50];
let soma = 0;

for (let i = 0; i < numeros.length; i++) {
  soma += numeros[i];
}

console.log('Soma total:', soma); // 150
```

#### 1.6. Múltiplas Variáveis no Loop

```javascript
// Você pode declarar múltiplas variáveis na inicialização
for (let i = 0, j = 10; i < 5; i++, j--) {
  console.log(`i: ${i}, j: ${j}`);
}

// Saída:
// i: 0, j: 10
// i: 1, j: 9
// i: 2, j: 8
// i: 3, j: 7
// i: 4, j: 6
```

#### 1.7. Loop `for` sem Corpo (Opcional)

```javascript
// Você pode omitir as chaves se houver apenas uma instrução
for (let i = 0; i < 5; i++) console.log(i);
```

---

## 🔁 2. O Loop `while`

### Definição

O loop `while` executa um bloco de código **enquanto uma condição for verdadeira**. A condição é avaliada **antes** de cada iteração. Se a condição for `false` desde o início, o loop não executa nenhuma vez.

### Sintaxe

```javascript
while (condição) {
  // código a ser executado
}
```

### Características

- ✅ Avalia a condição antes de executar
- ✅ Pode não executar nenhuma vez (se a condição for falsa desde o início)
- ✅ Ideal quando você não sabe quantas vezes precisa repetir
- ⚠️ Cuidado com loops infinitos - sempre tenha uma forma de sair

### Exemplos de Uso

#### 2.1. Loop `while` Básico

```javascript
let contador = 0;

while (contador < 5) {
  console.log('Contador:', contador);
  contador++; // IMPORTANTE: incrementar o contador
}

// Saída:
// Contador: 0
// Contador: 1
// Contador: 2
// Contador: 3
// Contador: 4
```

**⚠️ ATENÇÃO**: Se você esquecer de incrementar `contador`, o loop será infinito!

#### 2.2. Ler até Condição Específica

```javascript
let numero = 0;

while (numero !== 5) {
  numero = Math.floor(Math.random() * 10); // Número aleatório de 0 a 9
  console.log('Número gerado:', numero);
}

console.log('Encontrei o número 5!');
```

#### 2.3. Processar até Condição

```javascript
let saldo = 1000;
let saque = 100;

while (saldo >= saque) {
  saldo -= saque;
  console.log(`Saque de R$ ${saque}. Saldo restante: R$ ${saldo}`);
}

console.log('Saldo insuficiente para mais saques.');
```

#### 2.4. Loop `while` que Nunca Executa

```javascript
let x = 10;

while (x < 5) {
  console.log('Isso nunca será executado');
}

console.log('Loop terminou sem executar');
```

---

## 🔂 3. O Loop `do...while`

### Definição

O loop `do...while` é semelhante ao `while`, mas com uma diferença crucial: ele **executa o código pelo menos uma vez**, mesmo se a condição for falsa, porque a condição é avaliada **após** a execução.

### Sintaxe

```javascript
do {
  // código a ser executado
} while (condição);
```

### Características

- ✅ Executa pelo menos uma vez
- ✅ Avalia a condição após executar
- ✅ Útil para validação de entrada do usuário
- ✅ Útil quando você precisa executar algo antes de verificar a condição

### Exemplos de Uso

#### 3.1. Loop `do...while` Básico

```javascript
let contador = 0;

do {
  console.log('Contador:', contador);
  contador++;
} while (contador < 5);

// Saída:
// Contador: 0
// Contador: 1
// Contador: 2
// Contador: 3
// Contador: 4
```

#### 3.2. Executar pelo Menos Uma Vez

```javascript
let numero = 10;

do {
  console.log('Número:', numero);
  numero++;
} while (numero < 5);

// Saída: Número: 10
// Mesmo que a condição seja falsa, executa uma vez
```

#### 3.3. Validação de Entrada (Exemplo Prático)

```javascript
let senha;
let tentativas = 0;

do {
  senha = prompt('Digite sua senha (mínimo 6 caracteres):');
  tentativas++;
  
  if (senha.length < 6) {
    console.log('Senha muito curta! Tente novamente.');
  }
} while (senha.length < 6 && tentativas < 3);

if (senha.length >= 6) {
  console.log('Senha aceita!');
} else {
  console.log('Número máximo de tentativas excedido.');
}
```

---

## 🔀 4. O Loop `for...of`

### Definição

O loop `for...of` (introduzido no ES6) é uma forma moderna e limpa de iterar sobre **valores** de objetos iteráveis, como arrays, strings, Map, Set, NodeList, etc.

### Sintaxe

```javascript
for (const elemento of iteravel) {
  // código a ser executado
}
```

### Características

- ✅ Sintaxe mais limpa e legível
- ✅ Itera sobre valores (não índices)
- ✅ Funciona com qualquer objeto iterável
- ✅ Não precisa gerenciar índices manualmente
- ✅ **Recomendado para arrays** (em vez de `for...in`)

### Exemplos de Uso

#### 4.1. Iterar sobre Arrays

```javascript
const frutas = ['maçã', 'banana', 'laranja'];

for (const fruta of frutas) {
  console.log(fruta);
}

// Saída:
// maçã
// banana
// laranja
```

#### 4.2. Iterar sobre Strings

```javascript
const palavra = 'JavaScript';

for (const letra of palavra) {
  console.log(letra);
}

// Saída: J, a, v, a, S, c, r, i, p, t (uma por linha)
```

#### 4.3. Processar Arrays com Índice (usando `entries()`)

```javascript
const cores = ['vermelho', 'verde', 'azul'];

for (const [indice, cor] of cores.entries()) {
  console.log(`Índice ${indice}: ${cor}`);
}

// Saída:
// Índice 0: vermelho
// Índice 1: verde
// Índice 2: azul
```

#### 4.4. Iterar sobre NodeList (DOM)

```javascript
// Seleciona todos os parágrafos
const paragrafos = document.querySelectorAll('p');

for (const paragrafo of paragrafos) {
  paragrafo.style.color = 'blue';
}
```

#### 4.5. Iterar sobre Map

```javascript
const mapa = new Map([
  ['nome', 'João'],
  ['idade', 30],
  ['cidade', 'São Paulo']
]);

for (const [chave, valor] of mapa) {
  console.log(`${chave}: ${valor}`);
}

// Saída:
// nome: João
// idade: 30
// cidade: São Paulo
```

#### 4.6. Iterar sobre Set

```javascript
const numeros = new Set([1, 2, 3, 4, 5]);

for (const numero of numeros) {
  console.log(numero * 2);
}

// Saída: 2, 4, 6, 8, 10
```

---

## 🔍 5. O Loop `for...in`

### Definição

O loop `for...in` itera sobre **propriedades enumeráveis** de um objeto. Ele itera sobre as **chaves** (nomes das propriedades), não sobre os valores.

### Sintaxe

```javascript
for (const chave in objeto) {
  // código a ser executado
}
```

### Características

- ✅ Itera sobre propriedades enumeráveis de objetos
- ✅ Itera sobre chaves (não valores)
- ✅ Inclui propriedades herdadas (a menos que você filtre)
- ⚠️ **NÃO use para arrays** - use `for...of` ou `for` tradicional
- ✅ Útil para objetos e suas propriedades

### Exemplos de Uso

#### 5.1. Iterar sobre Objetos

```javascript
const pessoa = {
  nome: 'Maria',
  idade: 25,
  cidade: 'Rio de Janeiro'
};

for (const propriedade in pessoa) {
  console.log(`${propriedade}: ${pessoa[propriedade]}`);
}

// Saída:
// nome: Maria
// idade: 25
// cidade: Rio de Janeiro
```

#### 5.2. Verificar se Propriedade é Própria do Objeto

```javascript
const objeto = {
  a: 1,
  b: 2
};

// Adiciona propriedade ao protótipo (não faça isso na prática)
Object.prototype.c = 3;

for (const chave in objeto) {
  if (objeto.hasOwnProperty(chave)) {
    console.log(`Propriedade própria: ${chave} = ${objeto[chave]}`);
  }
}

// Saída:
// Propriedade própria: a = 1
// Propriedade própria: b = 2
// (c não é impresso porque não é propriedade própria)
```

#### 5.3. Por que NÃO Usar `for...in` com Arrays

```javascript
const array = ['a', 'b', 'c'];

// ❌ NÃO FAÇA ISSO
for (const indice in array) {
  console.log(array[indice]);
}

// ✅ FAÇA ISSO
for (const elemento of array) {
  console.log(elemento);
}

// Ou use for tradicional
for (let i = 0; i < array.length; i++) {
  console.log(array[i]);
}
```

**Por quê?** `for...in` pode iterar sobre propriedades adicionadas ao array (não apenas elementos) e a ordem não é garantida em todos os casos.

---

## ⏸️ 6. Declarações `break` e `continue`

### Definição

- **`break`**: Sai completamente do loop, interrompendo todas as iterações futuras
- **`continue`**: Pula para a próxima iteração do loop, ignorando o código restante na iteração atual

### Sintaxe

```javascript
// break
for (let i = 0; i < 10; i++) {
  if (condicao) {
    break; // Sai do loop
  }
}

// continue
for (let i = 0; i < 10; i++) {
  if (condicao) {
    continue; // Pula para próxima iteração
  }
}
```

### Exemplos de Uso

#### 6.1. Usando `break`

```javascript
// Encontrar o primeiro número par e parar
const numeros = [1, 3, 5, 8, 9, 10];

for (const numero of numeros) {
  if (numero % 2 === 0) {
    console.log(`Primeiro número par encontrado: ${numero}`);
    break; // Para o loop imediatamente
  }
}

// Saída: Primeiro número par encontrado: 8
// (não verifica 9 e 10)
```

#### 6.2. Usando `continue`

```javascript
// Imprimir apenas números ímpares
for (let i = 0; i < 10; i++) {
  if (i % 2 === 0) {
    continue; // Pula números pares
  }
  console.log(i);
}

// Saída: 1, 3, 5, 7, 9
```

#### 6.3. `break` em Loop `while`

```javascript
let numero = 0;

while (true) { // Loop aparentemente infinito
  numero = Math.floor(Math.random() * 100);
  console.log('Número gerado:', numero);
  
  if (numero > 90) {
    console.log('Número maior que 90 encontrado! Parando...');
    break; // Sai do loop
  }
}
```

#### 6.4. `continue` para Pular Valores Específicos

```javascript
const palavras = ['casa', 'carro', '', 'bicicleta', null, 'moto'];

for (const palavra of palavras) {
  if (!palavra) { // Se palavra for vazia, null, undefined, etc.
    continue; // Pula para próxima iteração
  }
  console.log(`Palavra: ${palavra}`);
}

// Saída:
// Palavra: casa
// Palavra: carro
// Palavra: bicicleta
// Palavra: moto
// (pula '' e null)
```

#### 6.5. `break` e `continue` em Loops Aninhados

```javascript
// break só sai do loop mais interno
for (let i = 0; i < 3; i++) {
  for (let j = 0; j < 5; j++) {
    if (j === 3) {
      break; // Sai apenas do loop interno (j)
    }
    console.log(`i: ${i}, j: ${j}`);
  }
}

// Saída:
// i: 0, j: 0
// i: 0, j: 1
// i: 0, j: 2
// i: 1, j: 0
// i: 1, j: 1
// i: 1, j: 2
// i: 2, j: 0
// i: 2, j: 1
// i: 2, j: 2
```

---

## 🏷️ 7. Labels (Rótulos) em Loops

### Definição

Labels permitem nomear loops e usar `break` ou `continue` para controlar loops externos a partir de loops internos.

### Sintaxe

```javascript
labelExterno: for (let i = 0; i < 3; i++) {
  labelInterno: for (let j = 0; j < 3; j++) {
    if (condicao) {
      break labelExterno; // Sai do loop externo
    }
  }
}
```

### Exemplos de Uso

#### 7.1. `break` com Label

```javascript
externo: for (let i = 0; i < 3; i++) {
  interno: for (let j = 0; j < 3; j++) {
    if (i === 1 && j === 1) {
      break externo; // Sai do loop externo completamente
    }
    console.log(`i: ${i}, j: ${j}`);
  }
}

// Saída:
// i: 0, j: 0
// i: 0, j: 1
// i: 0, j: 2
// i: 1, j: 0
// (para completamente quando i=1 e j=1)
```

#### 7.2. `continue` com Label

```javascript
externo: for (let i = 0; i < 3; i++) {
  interno: for (let j = 0; j < 3; j++) {
    if (j === 1) {
      continue externo; // Pula para próxima iteração do loop externo
    }
    console.log(`i: ${i}, j: ${j}`);
  }
}

// Saída:
// i: 0, j: 0
// i: 1, j: 0
// i: 2, j: 0
// (pula j=1 e j=2 em todas as iterações de i)
```

**⚠️ ATENÇÃO**: Labels são raramente necessários e podem tornar o código difícil de ler. Use com moderação e apenas quando realmente necessário.

---

## 🔄 8. Loops Aninhados

### Definição

Loops aninhados são loops dentro de outros loops. Eles são úteis para trabalhar com estruturas multidimensionais (como matrizes) ou combinações.

### Exemplos de Uso

#### 8.1. Criar Tabela de Multiplicação

```javascript
for (let i = 1; i <= 5; i++) {
  for (let j = 1; j <= 5; j++) {
    console.log(`${i} x ${j} = ${i * j}`);
  }
  console.log('---'); // Separador entre linhas
}
```

#### 8.2. Processar Matriz (Array 2D)

```javascript
const matriz = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
];

for (let i = 0; i < matriz.length; i++) {
  for (let j = 0; j < matriz[i].length; j++) {
    console.log(`matriz[${i}][${j}] = ${matriz[i][j]}`);
  }
}
```

#### 8.3. Encontrar Duplicatas

```javascript
const array1 = [1, 2, 3, 4];
const array2 = [3, 4, 5, 6];

for (let i = 0; i < array1.length; i++) {
  for (let j = 0; j < array2.length; j++) {
    if (array1[i] === array2[j]) {
      console.log(`Duplicata encontrada: ${array1[i]}`);
    }
  }
}

// Saída:
// Duplicata encontrada: 3
// Duplicata encontrada: 4
```

---

## 📊 9. Comparação entre Tipos de Loops

| Tipo | Quando Usar | Vantagens | Desvantagens |
|------|-------------|-----------|--------------|
| `for` | Número conhecido de iterações | Controle total, flexível | Mais verboso |
| `while` | Condição desconhecida | Simples, direto | Risco de loop infinito |
| `do...while` | Executar pelo menos uma vez | Garante execução | Menos comum |
| `for...of` | ✅ **Arrays, strings, iteráveis** | Sintaxe limpa, moderno | Não tem acesso direto ao índice |
| `for...in` | Propriedades de objetos | Acesso a chaves | ⚠️ Não use para arrays |

---

## ⚠️ 10. Armadilhas Comuns e Como Evitá-las

### 10.1. Loop Infinito

```javascript
// ❌ ERRADO - Loop infinito
let i = 0;
while (i < 5) {
  console.log(i);
  // Esqueceu de incrementar i!
}

// ✅ CORRETO
let i = 0;
while (i < 5) {
  console.log(i);
  i++; // Sempre incremente!
}
```

### 10.2. Modificar Array Durante Iteração

```javascript
const numeros = [1, 2, 3, 4, 5];

// ⚠️ CUIDADO - Pode causar problemas
for (let i = 0; i < numeros.length; i++) {
  if (numeros[i] === 3) {
    numeros.splice(i, 1); // Remove elemento
    // i não é incrementado, mas o array mudou!
  }
}

// ✅ MELHOR - Iterar de trás para frente
for (let i = numeros.length - 1; i >= 0; i--) {
  if (numeros[i] === 3) {
    numeros.splice(i, 1);
  }
}
```

### 10.3. Usar `for...in` com Arrays

```javascript
const array = ['a', 'b', 'c'];

// ❌ NÃO FAÇA
for (const indice in array) {
  console.log(array[indice]);
}

// ✅ FAÇA
for (const elemento of array) {
  console.log(elemento);
}
```

### 10.4. Recálculo de `length` em Cada Iteração

```javascript
const array = [1, 2, 3, 4, 5];

// ⚠️ INEFICIENTE - Recalcula length a cada iteração
for (let i = 0; i < array.length; i++) {
  console.log(array[i]);
}

// ✅ MAIS EFICIENTE - Cache do length
for (let i = 0, len = array.length; i < len; i++) {
  console.log(array[i]);
}

// ✅ AINDA MELHOR - Use for...of
for (const elemento of array) {
  console.log(elemento);
}
```

---

## 🎓 Resumo dos Conceitos

### Loops Fundamentais

1. **`for`**: Loop mais comum, ideal quando você sabe quantas vezes repetir
2. **`while`**: Executa enquanto a condição for verdadeira (avalia antes)
3. **`do...while`**: Executa pelo menos uma vez (avalia depois)

### Loops Modernos (ES6+)

4. **`for...of`**: ✅ **Use para arrays e iteráveis** - sintaxe limpa
5. **`for...in`**: Use apenas para propriedades de objetos

### Controle de Fluxo

6. **`break`**: Sai completamente do loop
7. **`continue`**: Pula para a próxima iteração
8. **Labels**: Controlam loops externos a partir de loops internos

### Regras de Ouro

- ✅ Use `for...of` para arrays
- ✅ Use `for...in` apenas para objetos
- ✅ Sempre tenha uma condição de saída clara
- ✅ Evite loops infinitos
- ✅ Cache `length` em loops `for` tradicionais
- ⚠️ Cuidado ao modificar arrays durante iteração

---

## 🚀 Próximos Passos

Agora que você entendeu os conceitos fundamentais de loops, você está pronto para:
- Aprender métodos de array (map, filter, reduce) que são alternativas funcionais aos loops
- Aplicar loops na manipulação do DOM
- Trabalhar com eventos e loops assíncronos
- Otimizar performance de loops

Na próxima etapa, você verá uma **versão simplificada** deste conteúdo com analogias e exemplos do dia a dia!

