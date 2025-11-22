# Aula 10: Funções em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 9**, você aprendeu:
- ✅ Expressões e operadores aritméticos, lógicos e de comparação
- ✅ Operador ternário e operadores de atribuição
- ✅ Precedência de operadores e como controlá-la
- ✅ Como combinar operadores para criar expressões complexas

Agora vamos aprender sobre **Funções** - um dos conceitos mais importantes e poderosos do JavaScript! Funções são blocos de código reutilizáveis que executam tarefas específicas sempre que são invocadas.

---

## 🎯 O que são Funções?

**Definição:** Funções são blocos de código nomeados que executam uma tarefa específica. Elas existem para que possamos **reutilizar código** sem precisar escrever a mesma lógica várias vezes.

### Conceitos Fundamentais

1. **Declaração**: Definir uma função com um nome e um corpo
2. **Invocação/Chamada**: Executar a função quando necessário
3. **Parâmetros**: Valores que a função recebe como entrada
4. **Retorno**: Valores que a função produz como saída
5. **Escopo**: Onde a função e suas variáveis podem ser acessadas

### Por que Funções são Importantes?

Sem funções, você teria que:
- Escrever o mesmo código repetidamente
- Manter múltiplas cópias do mesmo código (difícil de atualizar)
- Ter código muito longo e difícil de entender
- Não conseguir organizar e modularizar seu código

**Com funções, você pode:**
- Escrever código uma vez e reutilizá-lo
- Organizar seu código em blocos lógicos
- Facilitar manutenção e debugging
- Criar código mais limpo e legível

---

## 📝 1. Declaração de Funções

### 1.1. Function Declaration (Declaração de Função)

A forma mais tradicional de criar uma função em JavaScript.

#### Sintaxe

```javascript
function nomeDaFuncao(parametro1, parametro2) {
  // Código a ser executado
  return valor; // Opcional
}
```

#### Exemplo Básico

```javascript
// Função que soma dois números
function somar(a, b) {
  return a + b;
}

// Chamando a função
let resultado = somar(5, 3);
console.log(resultado); // 8
```

#### Características

- **Hoisting**: Funções declaradas com `function` são "elevadas" (hoisted), podendo ser chamadas antes de serem declaradas
- **Nome obrigatório**: Deve ter um nome identificador
- **Escopo de função**: Cria seu próprio escopo

#### Exemplo com Hoisting

```javascript
// Podemos chamar antes de declarar!
let resultado = multiplicar(4, 5);
console.log(resultado); // 20

function multiplicar(x, y) {
  return x * y;
}
```

### 1.2. Function Expression (Expressão de Função)

Uma função atribuída a uma variável.

#### Sintaxe

```javascript
const nomeDaFuncao = function(parametro1, parametro2) {
  // Código a ser executado
  return valor;
};
```

#### Exemplo

```javascript
const dividir = function(a, b) {
  if (b === 0) {
    return "Erro: divisão por zero!";
  }
  return a / b;
};

console.log(dividir(10, 2)); // 5
console.log(dividir(10, 0)); // "Erro: divisão por zero!"
```

#### Características

- **Sem hoisting**: Não pode ser chamada antes da declaração
- **Pode ser anônima**: A função não precisa ter nome (mas a variável sim)
- **Pode ser reatribuída**: Como é uma variável, pode ser alterada

### 1.3. Arrow Functions (Funções de Seta) - ES6+

Uma sintaxe mais curta e moderna para criar funções.

#### Sintaxe Básica

```javascript
// Forma completa
const nomeDaFuncao = (parametro1, parametro2) => {
  return valor;
};

// Forma curta (quando há apenas uma expressão)
const nomeDaFuncao = (parametro1, parametro2) => valor;

// Sem parâmetros
const nomeDaFuncao = () => {
  return valor;
};

// Um único parâmetro (parênteses opcionais)
const nomeDaFuncao = parametro => valor;
```

#### Exemplos

```javascript
// Forma tradicional
const somar = function(a, b) {
  return a + b;
};

// Arrow function equivalente
const somar = (a, b) => {
  return a + b;
};

// Arrow function simplificada
const somar = (a, b) => a + b;

// Sem parâmetros
const dizerOla = () => {
  console.log("Olá do Arrow Function!");
};

// Um parâmetro
const dobrar = x => x * 2;

console.log(somar(3, 4)); // 7
dizerOla(); // "Olá do Arrow Function!"
console.log(dobrar(5)); // 10
```

#### Características Importantes das Arrow Functions

1. **Sem hoisting**: Não são elevadas
2. **`this` léxico**: Não têm seu próprio `this` (veremos em aulas futuras)
3. **Não podem ser construtoras**: Não podem usar `new`
4. **Sintaxe mais curta**: Ideal para funções simples
5. **Retorno implícito**: Se houver apenas uma expressão, o `return` é automático

#### Comparação: Function vs Arrow Function

```javascript
// Function Declaration
function calcularArea(raio) {
  return Math.PI * raio * raio;
}

// Arrow Function
const calcularArea = (raio) => Math.PI * raio * raio;

// Ambas fazem a mesma coisa!
console.log(calcularArea(5)); // ~78.54
```

---

## 🔧 2. Parâmetros de Funções

### 2.1. Parâmetros Básicos

Parâmetros são variáveis que recebem valores quando a função é chamada.

```javascript
function saudar(nome, idade) {
  console.log(`Olá, ${nome}! Você tem ${idade} anos.`);
}

saudar("Maria", 25); // "Olá, Maria! Você tem 25 anos."
saudar("João", 30);  // "Olá, João! Você tem 30 anos."
```

### 2.2. Default Parameters (Parâmetros Padrão) - ES6+

Permitem definir valores padrão para parâmetros caso nenhum valor seja passado ou `undefined` seja passado.

#### Sintaxe

```javascript
function nomeDaFuncao(parametro = valorPadrao) {
  // Código
}
```

#### Exemplos

```javascript
// Função com parâmetro padrão
function saudar(nome = "Visitante") {
  console.log(`Olá, ${nome}!`);
}

saudar("Maria");    // "Olá, Maria!"
saudar();           // "Olá, Visitante!"
saudar(undefined);   // "Olá, Visitante!"

// Múltiplos parâmetros padrão
function criarUsuario(nome, idade = 18, ativo = true) {
  return {
    nome: nome,
    idade: idade,
    ativo: ativo
  };
}

console.log(criarUsuario("João"));           // { nome: "João", idade: 18, ativo: true }
console.log(criarUsuario("Maria", 25));      // { nome: "Maria", idade: 25, ativo: true }
console.log(criarUsuario("Pedro", 30, false)); // { nome: "Pedro", idade: 30, ativo: false }
```

#### Parâmetros Padrão com Expressões

```javascript
function calcularPreco(preco, desconto = preco * 0.1) {
  return preco - desconto;
}

console.log(calcularPreco(100));      // 90 (desconto padrão de 10%)
console.log(calcularPreco(100, 20));  // 80 (desconto customizado)
```

### 2.3. Rest Parameters (Parâmetros Rest) - ES6+

Permitem que uma função aceite um número indefinido de argumentos como um array.

#### Sintaxe

```javascript
function nomeDaFuncao(parametro1, parametro2, ...resto) {
  // resto é um array com os argumentos restantes
}
```

#### Exemplos

```javascript
// Função que soma todos os números passados
function somar(...numeros) {
  let total = 0;
  for (let numero of numeros) {
    total += numero;
  }
  return total;
}

console.log(somar(1, 2));           // 3
console.log(somar(1, 2, 3));        // 6
console.log(somar(1, 2, 3, 4, 5));  // 15

// Combinando parâmetros normais com rest
function criarPerfil(nome, idade, ...hobbies) {
  return {
    nome: nome,
    idade: idade,
    hobbies: hobbies
  };
}

let perfil = criarPerfil("Maria", 25, "leitura", "natação", "ciclismo");
console.log(perfil);
// {
//   nome: "Maria",
//   idade: 25,
//   hobbies: ["leitura", "natação", "ciclismo"]
// }
```

#### Rest vs Arguments (Legado)

```javascript
// Forma antiga (não recomendada)
function somarAntiga() {
  let total = 0;
  for (let i = 0; i < arguments.length; i++) {
    total += arguments[i];
  }
  return total;
}

// Forma moderna com rest (recomendada)
function somarModerna(...numeros) {
  return numeros.reduce((total, num) => total + num, 0);
}

// Ambas funcionam, mas rest é melhor!
console.log(somarAntiga(1, 2, 3));  // 6
console.log(somarModerna(1, 2, 3)); // 6
```

**Importante:** O parâmetro rest deve ser sempre o último!

```javascript
// ✅ Correto
function exemplo(a, b, ...resto) { }

// ❌ Errado
function exemplo(...resto, a, b) { } // SyntaxError
```

---

## 🔄 3. Return (Retorno de Valores)

### 3.1. O que é Return?

A palavra-chave `return` finaliza a execução de uma função e especifica um valor para ser retornado.

#### Sintaxe

```javascript
function nomeDaFuncao() {
  // Código
  return valor; // Retorna o valor e encerra a função
}
```

#### Exemplos

```javascript
// Função com retorno
function multiplicar(a, b) {
  return a * b;
}

let resultado = multiplicar(4, 5);
console.log(resultado); // 20

// Função sem retorno (retorna undefined)
function apenasLog(texto) {
  console.log(texto);
  // Sem return explícito
}

let valor = apenasLog("Olá");
console.log(valor); // undefined

// Múltiplos returns (com condicionais)
function verificarIdade(idade) {
  if (idade >= 18) {
    return "Maior de idade";
  } else {
    return "Menor de idade";
  }
}

console.log(verificarIdade(20)); // "Maior de idade"
console.log(verificarIdade(15)); // "Menor de idade"
```

### 3.2. Return vs Console.log

```javascript
// ❌ Errado - apenas imprime, não retorna
function somarErrado(a, b) {
  console.log(a + b);
}

let resultado1 = somarErrado(2, 3); // Imprime 5, mas resultado1 é undefined

// ✅ Correto - retorna o valor
function somarCorreto(a, b) {
  return a + b;
}

let resultado2 = somarCorreto(2, 3); // resultado2 é 5
console.log(resultado2); // 5
```

### 3.3. Retornando Múltiplos Valores

```javascript
// Retornando um objeto
function obterDados() {
  return {
    nome: "Maria",
    idade: 25,
    cidade: "São Paulo"
  };
}

// Retornando um array
function obterCoordenadas() {
  return [10, 20];
}

let [x, y] = obterCoordenadas();
console.log(x, y); // 10 20
```

---

## 🌐 4. Escopo (Scope) e Funções

### 4.1. O que é Escopo?

**Escopo** é o espaço ou ambiente onde uma variável ou função pode ser acessada. A acessibilidade depende de onde ela foi definida.

### 4.2. Tipos de Escopo em JavaScript

#### Global Scope (Escopo Global)

Variáveis declaradas fora de qualquer função são **globais** e podem ser acessadas de qualquer lugar.

```javascript
let variavelGlobal = "Eu sou global";

function funcao1() {
  console.log(variavelGlobal); // ✅ Pode acessar
}

function funcao2() {
  console.log(variavelGlobal); // ✅ Pode acessar
}

console.log(variavelGlobal); // ✅ Pode acessar
```

#### Function Scope (Escopo de Função)

Variáveis declaradas dentro de uma função são **locais** e só podem ser acessadas dentro dessa função.

```javascript
function minhaFuncao() {
  let variavelLocal = "Eu sou local";
  console.log(variavelLocal); // ✅ Funciona
}

console.log(variavelLocal); // ❌ ReferenceError: variavelLocal is not defined
```

#### Block Scope (Escopo de Bloco) - ES6+

Variáveis declaradas com `let` ou `const` dentro de blocos `{}` têm escopo de bloco.

```javascript
if (true) {
  let variavelBloco = "Eu sou de bloco";
  var variavelFuncao = "Eu sou de função";
  
  console.log(variavelBloco);  // ✅ Funciona
  console.log(variavelFuncao); // ✅ Funciona
}

console.log(variavelBloco);  // ❌ ReferenceError
console.log(variavelFuncao); // ✅ Funciona (var tem function scope)
```

### 4.3. Escopo de Funções

```javascript
// Função global
function funcaoGlobal() {
  console.log("Sou global");
}

// Função dentro de função
function funcaoExterna() {
  function funcaoInterna() {
    console.log("Sou interna");
  }
  
  funcaoInterna(); // ✅ Funciona
}

funcaoExterna();    // ✅ Funciona
funcaoInterna();    // ❌ ReferenceError: funcaoInterna is not defined
```

### 4.4. Shadowing (Sombreamento)

Quando uma variável local tem o mesmo nome de uma variável global, a local "esconde" a global.

```javascript
let nome = "Global";

function exemplo() {
  let nome = "Local";
  console.log(nome); // "Local" (usa a variável local)
}

exemplo();
console.log(nome); // "Global" (variável global não foi alterada)
```

---

## 📚 5. Function Stack (Call Stack)

### 5.1. O que é Call Stack?

A **Call Stack** (Pilha de Chamadas) é como o interpretador JavaScript rastreia sua posição em um script que chama múltiplas funções. Ela mostra qual função está sendo executada e quais funções dentro dessa função estão sendo chamadas.

### 5.2. Como Funciona

```javascript
function primeira() {
  console.log("1. Primeira função");
  segunda();
  console.log("5. Primeira função terminou");
}

function segunda() {
  console.log("2. Segunda função");
  terceira();
  console.log("4. Segunda função terminou");
}

function terceira() {
  console.log("3. Terceira função");
}

primeira();
```

**Saída:**
```
1. Primeira função
2. Segunda função
3. Terceira função
4. Segunda função terminou
5. Primeira função terminou
```

**Visualização da Call Stack:**
```
[terceira]  ← Topo (executando)
[segunda]
[primeira]
[global]    ← Base
```

### 5.3. Stack Overflow

Se uma função chama a si mesma infinitamente, ocorre um **Stack Overflow** (estouro de pilha).

```javascript
function infinito() {
  infinito(); // Chama a si mesma
}

infinito(); // ❌ RangeError: Maximum call stack size exceeded
```

---

## 🔁 6. Recursão

### 6.1. O que é Recursão?

**Recursão** é quando uma função chama a si mesma. É um dos conceitos mais poderosos e elegantes em programação.

### 6.2. Estrutura de uma Função Recursiva

Toda função recursiva precisa de:
1. **Base Case (Caso Base)**: Condição que para a recursão
2. **Recursive Case (Caso Recursivo)**: Chamada da função a si mesma

#### Exemplo: Fatorial

```javascript
function fatorial(n) {
  // Caso base
  if (n === 0 || n === 1) {
    return 1;
  }
  
  // Caso recursivo
  return n * fatorial(n - 1);
}

console.log(fatorial(5)); // 120
// 5! = 5 * 4 * 3 * 2 * 1 = 120
```

**Como funciona:**
```
fatorial(5)
  → 5 * fatorial(4)
    → 4 * fatorial(3)
      → 3 * fatorial(2)
        → 2 * fatorial(1)
          → 1 (caso base)
```

#### Exemplo: Contagem Regressiva

```javascript
function contagemRegressiva(numero) {
  // Caso base
  if (numero <= 0) {
    console.log("Fogo!");
    return;
  }
  
  // Caso recursivo
  console.log(numero);
  contagemRegressiva(numero - 1);
}

contagemRegressiva(5);
// 5
// 4
// 3
// 2
// 1
// Fogo!
```

#### Exemplo: Soma de Array

```javascript
function somarArray(array, indice = 0) {
  // Caso base
  if (indice >= array.length) {
    return 0;
  }
  
  // Caso recursivo
  return array[indice] + somarArray(array, indice + 1);
}

console.log(somarArray([1, 2, 3, 4, 5])); // 15
```

### 6.3. Recursão vs Iteração

```javascript
// Versão iterativa (com loop)
function fatorialIterativo(n) {
  let resultado = 1;
  for (let i = 1; i <= n; i++) {
    resultado *= i;
  }
  return resultado;
}

// Versão recursiva
function fatorialRecursivo(n) {
  if (n === 0 || n === 1) return 1;
  return n * fatorialRecursivo(n - 1);
}

// Ambas produzem o mesmo resultado
console.log(fatorialIterativo(5));  // 120
console.log(fatorialRecursivo(5));  // 120
```

**Quando usar cada uma?**
- **Iteração**: Geralmente mais eficiente em termos de memória
- **Recursão**: Geralmente mais elegante e fácil de entender para problemas que são naturalmente recursivos

---

## 🛠️ 7. Built-in Functions (Funções Nativas)

JavaScript oferece uma variedade de funções nativas que simplificam tarefas comuns, disponíveis globalmente ou dentro de objetos específicos.

### 7.1. Funções Globais

```javascript
// parseInt() - Converte string para número inteiro
let numero1 = parseInt("42");        // 42
let numero2 = parseInt("42.7");     // 42 (trunca)
let numero3 = parseInt("abc");      // NaN

// parseFloat() - Converte string para número decimal
let decimal1 = parseFloat("42.7");  // 42.7
let decimal2 = parseFloat("42");    // 42

// isNaN() - Verifica se é NaN
console.log(isNaN(42));      // false
console.log(isNaN("abc"));   // true

// isFinite() - Verifica se é um número finito
console.log(isFinite(42));           // true
console.log(isFinite(Infinity));     // false
console.log(isFinite("42"));         // true (converte)
```

### 7.2. Funções de Objetos Nativos

#### String Methods

```javascript
let texto = "JavaScript é incrível";

console.log(texto.length);                    // 23
console.log(texto.toUpperCase());             // "JAVASCRIPT É INCRÍVEL"
console.log(texto.toLowerCase());             // "javascript é incrível"
console.log(texto.indexOf("é"));              // 11
console.log(texto.substring(0, 10));          // "JavaScript"
console.log(texto.replace("incrível", "ótimo")); // "JavaScript é ótimo"
```

#### Array Methods

```javascript
let numeros = [1, 2, 3, 4, 5];

console.log(numeros.length);              // 5
console.log(numeros.push(6));             // 6 (retorna novo length)
console.log(numeros);                     // [1, 2, 3, 4, 5, 6]
console.log(numeros.pop());               // 6 (remove e retorna)
console.log(numeros);                     // [1, 2, 3, 4, 5]
console.log(numeros.includes(3));         // true
console.log(numeros.indexOf(3));          // 2
```

#### Math Object

```javascript
console.log(Math.PI);                    // 3.141592653589793
console.log(Math.max(1, 5, 3, 9));       // 9
console.log(Math.min(1, 5, 3, 9));       // 1
console.log(Math.round(4.7));            // 5
console.log(Math.floor(4.7));            // 4
console.log(Math.ceil(4.2));             // 5
console.log(Math.random());               // Número aleatório entre 0 e 1
console.log(Math.sqrt(16));               // 4
console.log(Math.pow(2, 3));              // 8
```

#### Date Object

```javascript
let agora = new Date();
console.log(agora);                       // Data e hora atual

let dataEspecifica = new Date(2024, 0, 1); // 1 de janeiro de 2024
console.log(dataEspecifica.getFullYear());  // 2024
console.log(dataEspecifica.getMonth());     // 0 (janeiro é 0)
console.log(dataEspecifica.getDate());     // 1
```

### 7.3. Funções de Tempo

```javascript
// setTimeout() - Executa após um delay
setTimeout(() => {
  console.log("Isso aparece após 2 segundos");
}, 2000);

// setInterval() - Executa repetidamente
let contador = 0;
let intervalo = setInterval(() => {
  contador++;
  console.log(`Contador: ${contador}`);
  if (contador >= 5) {
    clearInterval(intervalo); // Para o intervalo
  }
}, 1000);
```

---

## 💡 8. Exemplos Práticos Completos

### Exemplo 1: Calculadora Simples

```javascript
function calculadora(operacao, a, b) {
  switch(operacao) {
    case 'somar':
      return a + b;
    case 'subtrair':
      return a - b;
    case 'multiplicar':
      return a * b;
    case 'dividir':
      if (b === 0) {
        return "Erro: divisão por zero!";
      }
      return a / b;
    default:
      return "Operação inválida";
  }
}

console.log(calculadora('somar', 10, 5));        // 15
console.log(calculadora('subtrair', 10, 5));     // 5
console.log(calculadora('multiplicar', 10, 5));  // 50
console.log(calculadora('dividir', 10, 5));      // 2
console.log(calculadora('dividir', 10, 0));      // "Erro: divisão por zero!"
```

### Exemplo 2: Validação de Formulário

```javascript
function validarEmail(email) {
  if (!email || email.length === 0) {
    return "Email não pode estar vazio";
  }
  
  if (!email.includes("@")) {
    return "Email deve conter @";
  }
  
  if (!email.includes(".")) {
    return "Email deve conter um ponto";
  }
  
  return "Email válido";
}

console.log(validarEmail("usuario@email.com"));  // "Email válido"
console.log(validarEmail("usuarioemail.com"));   // "Email deve conter @"
console.log(validarEmail(""));                   // "Email não pode estar vazio"
```

### Exemplo 3: Sistema de Notas

```javascript
function calcularMedia(...notas) {
  if (notas.length === 0) {
    return "Nenhuma nota fornecida";
  }
  
  let soma = 0;
  for (let nota of notas) {
    soma += nota;
  }
  
  return soma / notas.length;
}

function verificarAprovacao(media) {
  if (media >= 7) {
    return "Aprovado";
  } else if (media >= 5) {
    return "Recuperação";
  } else {
    return "Reprovado";
  }
}

let media = calcularMedia(8, 7, 9, 6);
let status = verificarAprovacao(media);

console.log(`Média: ${media}`);        // Média: 7.5
console.log(`Status: ${status}`);      // Status: Aprovado
```

### Exemplo 4: Função Recursiva - Fibonacci

```javascript
function fibonacci(n) {
  // Casos base
  if (n === 0) return 0;
  if (n === 1) return 1;
  
  // Caso recursivo
  return fibonacci(n - 1) + fibonacci(n - 2);
}

console.log(fibonacci(0));  // 0
console.log(fibonacci(1));  // 1
console.log(fibonacci(2));  // 1
console.log(fibonacci(3));  // 2
console.log(fibonacci(4));  // 3
console.log(fibonacci(5));  // 5
console.log(fibonacci(6));  // 8
```

---

## 🎓 Resumo da Aula

Nesta aula, você aprendeu:

✅ **O que são funções**: Blocos de código reutilizáveis que executam tarefas específicas

✅ **Tipos de declaração**:
- Function Declaration (com hoisting)
- Function Expression (sem hoisting)
- Arrow Functions (sintaxe moderna ES6+)

✅ **Parâmetros**:
- Parâmetros básicos
- Default Parameters (valores padrão)
- Rest Parameters (número indefinido de argumentos)

✅ **Return**: Como retornar valores de funções

✅ **Escopo**: Global, Function e Block scope

✅ **Call Stack**: Como o JavaScript rastreia chamadas de funções

✅ **Recursão**: Funções que chamam a si mesmas, com casos base e recursivos

✅ **Built-in Functions**: Funções nativas do JavaScript para tarefas comuns

---

## 🚀 Próximos Passos

Na próxima etapa, você verá uma versão simplificada desta aula com analogias do dia a dia para fixar ainda mais os conceitos!

