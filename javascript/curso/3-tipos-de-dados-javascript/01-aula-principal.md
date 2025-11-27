# Aula 3: Tipos de Dados em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 2**, você aprendeu:
- ✅ O que são variáveis e como declará-las
- ✅ Diferenças entre `var`, `let` e `const`
- ✅ Conceitos de escopo (global, função, bloco)
- ✅ Hoisting e temporal dead zone
- ✅ Boas práticas de nomenclatura

Agora vamos aprender sobre **os diferentes tipos de dados** que podemos armazenar em variáveis!

---

## 🎯 O que são Tipos de Dados?

**Definição:** Um tipo de dado (data type) refere-se ao **tipo de informação** que uma variável JavaScript pode armazenar. Cada tipo de dado tem características específicas e comportamentos únicos.

JavaScript é uma linguagem **dinamicamente tipada**, o que significa que:
- Você não precisa declarar o tipo de uma variável explicitamente
- O tipo é determinado automaticamente pelo valor atribuído
- Uma variável pode mudar de tipo durante a execução do programa

---

## 📊 Classificação dos Tipos de Dados

JavaScript possui dois grandes grupos de tipos de dados:

1. **Tipos Primitivos** (7 tipos):
   - `Number`
   - `BigInt`
   - `String`
   - `Boolean`
   - `Null`
   - `Undefined`
   - `Symbol`

2. **Tipos Não-Primitivos** (Objetos):
   - `Object`
   - Arrays (que são objetos)
   - Funções (que são objetos)
   - Objetos built-in (Date, Math, etc.)

---

## 🔢 1. Number (Número)

### Definição

O tipo `Number` em JavaScript representa **números de ponto flutuante** (números decimais). Diferente de outras linguagens, JavaScript não possui tipos separados para inteiros e decimais - todos são tratados como números de ponto flutuante.

### Características

- JavaScript usa o padrão **IEEE 754** para representação de números
- Todos os números são armazenados como valores de 64 bits (double precision)
- Faixa segura: `-(2^53 - 1)` até `(2^53 - 1)` (aproximadamente ±9 quatrilhões)
- Números fora dessa faixa podem perder precisão

### Sintaxe e Exemplos

```javascript
// Números inteiros
let idade = 25;
let quantidade = 100;

// Números decimais
let preco = 19.99;
let temperatura = -5.5;

// Notação científica (exponencial)
let numeroGrande = 1.5e6; // 1.500.000
let numeroPequeno = 2e-3; // 0.002

// Notação hexadecimal (base 16)
let hex = 0xFF; // 255 em decimal
let hex2 = 0x1A; // 26 em decimal

// Notação binária (base 2) - ES6+
let binario = 0b11111111; // 255 em decimal
let binario2 = 0b1010; // 10 em decimal

// Notação octal (base 8) - ES6+
let octal = 0o377; // 255 em decimal

// Todos são do mesmo tipo
console.log(typeof idade); // "number"
console.log(typeof preco); // "number"
console.log(typeof hex); // "number"
console.log(typeof binario); // "number"

// Comparações
console.log(255 === 255.0); // true
console.log(255 === 0xFF); // true
console.log(255 === 0b11111111); // true
console.log(255 === 0o377); // true
```

### Valores Especiais

```javascript
// Infinity (infinito positivo)
let infinito = Infinity;
let infinito2 = 1 / 0; // Também resulta em Infinity

// -Infinity (infinito negativo)
let infinitoNegativo = -Infinity;
let infinitoNegativo2 = -1 / 0;

// NaN (Not a Number) - resultado de operações inválidas
let naoENumero = NaN;
let naoENumero2 = "texto" / 2; // NaN
let naoENumero3 = 0 / 0; // NaN

console.log(typeof Infinity); // "number"
console.log(typeof NaN); // "number" (peculiaridade do JavaScript)

// Verificando NaN
console.log(isNaN(NaN)); // true
console.log(Number.isNaN(NaN)); // true (método mais confiável)
```

### Operações com Números

```javascript
// Operações aritméticas básicas
let soma = 10 + 5; // 15
let subtracao = 10 - 5; // 5
let multiplicacao = 10 * 5; // 50
let divisao = 10 / 5; // 2
let resto = 10 % 3; // 1 (módulo)
let potencia = 2 ** 3; // 8 (ES6+)

// Precisão de ponto flutuante
let resultado = 0.1 + 0.2;
console.log(resultado); // 0.30000000000000004 (erro de precisão)
console.log(resultado === 0.3); // false

// Solução: arredondar quando necessário
console.log(Math.round(resultado * 100) / 100); // 0.3
```

---

## 🔢 2. BigInt (Números Inteiros Grandes)

### Definição

`BigInt` é um tipo de dado introduzido no **ES2020** que permite trabalhar com **inteiros de tamanho arbitrário**. Diferente do tipo `Number`, que tem limitações de precisão, `BigInt` pode representar números muito grandes sem perder precisão.

### Características

- Criado para resolver limitações do tipo `Number`
- `Number` pode representar com precisão apenas inteiros até `±(2^53 - 1)`
- `BigInt` não tem limite superior teórico (limitado apenas pela memória)
- Não pode ser misturado diretamente com `Number` em operações

### Sintaxe

```javascript
// Criando BigInt - adicione 'n' ao final do número
let numeroGrande = 9007199254740991n; // BigInt literal
let numeroGrande2 = BigInt(9007199254740991); // Usando construtor

// A partir de string
let numeroGrande3 = BigInt("9007199254740991");

// A partir de Number (converte para BigInt)
let numeroGrande4 = BigInt(Number.MAX_SAFE_INTEGER);
```

### Exemplos

```javascript
// Limite do Number
console.log(Number.MAX_SAFE_INTEGER); // 9007199254740991
console.log(Number.MAX_SAFE_INTEGER + 1); // 9007199254740992 (perde precisão)

// Com BigInt
let big1 = 9007199254740991n;
let big2 = 1n;
console.log(big1 + big2); // 9007199254740992n (precisão mantida)

// Operações com BigInt
let a = 123456789012345678901234567890n;
let b = 987654321098765432109876543210n;
console.log(a + b); // 1111111110111111111011111111100n
console.log(a * b); // Funciona perfeitamente

// Não pode misturar BigInt com Number diretamente
let num = 10;
let big = 20n;
// console.log(num + big); // TypeError: Cannot mix BigInt and other types

// Solução: converter explicitamente
console.log(num + Number(big)); // 30
console.log(BigInt(num) + big); // 30n
```

### Quando Usar BigInt

- Criptografia e segurança
- Cálculos científicos com números muito grandes
- IDs únicos muito grandes
- Manipulação de timestamps em nanossegundos
- **Evite usar** para operações matemáticas comuns (use `Number`)

---

## 📝 3. String (Texto)

### Definição

`String` é um tipo primitivo que representa uma **sequência de caracteres** (texto). Strings em JavaScript são imutáveis - uma vez criada, não pode ser alterada diretamente.

### Características

- Strings são **imutáveis** (não podem ser alteradas após criação)
- Podem ser escritas com aspas simples, duplas ou template literals
- Suportam caracteres Unicode
- Têm propriedade `length` para obter o tamanho

### Sintaxe

```javascript
// Aspas simples
let nome1 = 'João';

// Aspas duplas
let nome2 = "Maria";

// Template literals (backticks) - ES6+
let nome3 = `Pedro`;

// Todas são equivalentes
console.log(typeof nome1); // "string"
console.log(typeof nome2); // "string"
console.log(typeof nome3); // "string"
```

### Template Literals (ES6+)

```javascript
// Interpolação de variáveis
let nome = "Ana";
let idade = 25;
let mensagem = `Olá, meu nome é ${nome} e tenho ${idade} anos.`;
console.log(mensagem); // "Olá, meu nome é Ana e tenho 25 anos."

// Expressões dentro de template literals
let a = 10;
let b = 5;
let resultado = `A soma de ${a} + ${b} é igual a ${a + b}.`;
console.log(resultado); // "A soma de 10 + 5 é igual a 15."

// Strings multilinha
let texto = `Esta é uma string
que pode ter múltiplas
linhas sem precisar de
caracteres especiais.`;

// Antes do ES6 (forma antiga)
let textoAntigo = "Esta é uma string\n" +
                  "que precisa de\n" +
                  "caracteres de escape.";
```

### Escape de Caracteres

```javascript
// Aspas dentro de strings
let texto1 = "Ele disse: \"Olá!\"";
let texto2 = 'Ele disse: \'Olá!\'';
let texto3 = `Ele disse: "Olá!"`; // Template literals permitem aspas sem escape

// Caracteres especiais
let novaLinha = "Linha 1\nLinha 2"; // \n = quebra de linha
let tab = "Coluna1\tColuna2"; // \t = tabulação
let barra = "C:\\Users\\Documentos"; // \\ = barra invertida
let unicode = "\u00A9"; // © (caractere Unicode)

console.log(novaLinha);
// Linha 1
// Linha 2
```

### Propriedades e Métodos Básicos

```javascript
let texto = "JavaScript";

// Propriedade length
console.log(texto.length); // 10

// Acesso a caracteres (indexação)
console.log(texto[0]); // "J"
console.log(texto[4]); // "S"
console.log(texto[texto.length - 1]); // "t" (último caractere)

// Concatenação
let nome = "João";
let sobrenome = "Silva";
let nomeCompleto = nome + " " + sobrenome;
console.log(nomeCompleto); // "João Silva"

// Concatenação com template literals
let nomeCompleto2 = `${nome} ${sobrenome}`;
console.log(nomeCompleto2); // "João Silva"
```

### Conversão para String

```javascript
// Conversão implícita
let numero = 42;
let texto = numero + ""; // "42"
let texto2 = String(numero); // "42" (explícita)

// toString() - não funciona com null e undefined
let num = 123;
console.log(num.toString()); // "123"

// Template literals fazem conversão automática
let idade = 25;
console.log(`Idade: ${idade}`); // "Idade: 25"
```

---

## ✅ 4. Boolean (Booleano)

### Definição

`Boolean` é um tipo primitivo que representa um **valor lógico** - apenas dois valores possíveis: `true` (verdadeiro) ou `false` (falso).

### Características

- Usado em estruturas condicionais (if/else, while, etc.)
- Essencial para controle de fluxo do programa
- Valores "truthy" e "falsy" em JavaScript

### Sintaxe

```javascript
// Valores booleanos literais
let estaChovendo = true;
let estaEnsolarado = false;

console.log(typeof estaChovendo); // "boolean"
console.log(typeof estaEnsolarado); // "boolean"
```

### Uso em Condicionais

```javascript
let idade = 18;
let podeVotar = idade >= 16;

if (podeVotar) {
    console.log("Pode votar!");
} else {
    console.log("Não pode votar ainda.");
}

// Em loops
let continuar = true;
let contador = 0;

while (continuar && contador < 5) {
    console.log(`Iteração ${contador}`);
    contador++;
    if (contador === 5) {
        continuar = false;
    }
}
```

### Conversão para Boolean

```javascript
// Conversão explícita
let valor1 = Boolean(1); // true
let valor2 = Boolean(0); // false
let valor3 = Boolean("texto"); // true
let valor4 = Boolean(""); // false

// Operador ! (negação dupla)
let valor5 = !!1; // true
let valor6 = !!0; // false
```

### Valores Truthy e Falsy

JavaScript possui valores que são considerados "falsy" (avaliam como `false`) e "truthy" (avaliam como `true`):

**Valores Falsy:**
- `false`
- `0` (zero)
- `-0` (zero negativo)
- `0n` (BigInt zero)
- `""` (string vazia)
- `null`
- `undefined`
- `NaN`

**Valores Truthy:**
- Todos os outros valores são truthy
- `true`
- Números diferentes de zero
- Strings não vazias
- Objetos (mesmo vazios)
- Arrays (mesmo vazios)
- Funções

```javascript
// Exemplos
if (0) console.log("não executa");
if ("") console.log("não executa");
if (null) console.log("não executa");
if (undefined) console.log("não executa");

if (1) console.log("executa");
if ("texto") console.log("executa");
if ([]) console.log("executa");
if ({}) console.log("executa");
```

---

## ❓ 5. Undefined (Indefinido)

### Definição

`undefined` é um tipo primitivo que representa a **ausência de valor** - uma variável que foi declarada mas não foi inicializada ou atribuída.

### Características

- Valor padrão de variáveis não inicializadas
- Retornado por funções que não têm `return`
- Propriedade de objeto que não existe retorna `undefined`
- Não deve ser usado intencionalmente (use `null` para isso)

### Sintaxe e Exemplos

```javascript
// Variável declarada mas não inicializada
let variavel;
console.log(variavel); // undefined
console.log(typeof variavel); // "undefined"

// Função sem return
function semRetorno() {
    // não retorna nada
}
console.log(semRetorno()); // undefined

// Acesso a propriedade inexistente
let objeto = { nome: "João" };
console.log(objeto.idade); // undefined

// Parâmetro não fornecido
function saudacao(nome) {
    console.log(`Olá, ${nome}!`);
}
saudacao(); // "Olá, undefined!"

// Comparações
let valor;
console.log(valor === undefined); // true
console.log(valor == undefined); // true (mas evite ==)
```

### Diferença entre `undefined` e `null`

```javascript
let a; // undefined (não foi atribuído valor)
let b = null; // null (foi explicitamente definido como vazio)

console.log(a); // undefined
console.log(b); // null
console.log(typeof a); // "undefined"
console.log(typeof b); // "object" (peculiaridade do JavaScript)

// Comparações
console.log(a === undefined); // true
console.log(b === null); // true
console.log(a == b); // true (mas são diferentes!)
console.log(a === b); // false
```

---

## 🚫 6. Null (Nulo)

### Definição

`null` é um tipo primitivo que representa a **ausência intencional de um valor de objeto**. É usado para indicar que uma variável foi explicitamente definida como vazia ou sem referência.

### Características

- Representa ausência **deliberada** de valor
- É um valor **falsy**
- `typeof null` retorna `"object"` (bug histórico do JavaScript)
- Deve ser usado quando você quer indicar explicitamente "sem valor"

### Sintaxe e Exemplos

```javascript
// Atribuição explícita de null
let usuario = null; // Indica que não há usuário no momento

// Uso comum: resetar uma variável
let dados = { nome: "João", idade: 25 };
console.log(dados); // { nome: "João", idade: 25 }

dados = null; // Resetar para indicar ausência de dados
console.log(dados); // null

// Verificação
if (usuario === null) {
    console.log("Nenhum usuário logado");
}

// typeof null (peculiaridade)
console.log(typeof null); // "object" (bug histórico, mas mantido para compatibilidade)

// Verificação correta de null
let valor = null;
console.log(valor === null); // true (forma correta)
console.log(valor == null); // true (mas evite ==)
```

### Quando Usar `null` vs `undefined`

```javascript
// Use null quando:
// - Você quer explicitamente indicar "sem valor"
// - Resetar uma variável de objeto
// - Indicar que algo foi intencionalmente removido

let elemento = document.getElementById("inexistente");
console.log(elemento); // null (não encontrado)

// Use undefined quando:
// - Variável não foi inicializada
// - Propriedade não existe
// - Parâmetro não foi fornecido

let variavel;
console.log(variavel); // undefined (não inicializada)
```

---

## 🔍 7. Symbol (Símbolo)

### Definição

`Symbol` é um tipo primitivo introduzido no **ES6** que representa um **identificador único e imutável**. Cada símbolo é único, mesmo que tenha a mesma descrição.

### Características

- Cada `Symbol` é **único** (não há dois símbolos iguais)
- Usado principalmente como chaves de propriedades de objetos
- Não pode ser convertido para string diretamente
- Útil para criar propriedades "privadas" em objetos

### Sintaxe

```javascript
// Criando símbolos
let simbolo1 = Symbol();
let simbolo2 = Symbol("descricao"); // Descrição opcional (apenas para debug)

// Cada símbolo é único
let s1 = Symbol("teste");
let s2 = Symbol("teste");
console.log(s1 === s2); // false (são diferentes!)

// Symbol.for() - símbolos globais (mesma descrição = mesmo símbolo)
let global1 = Symbol.for("chave");
let global2 = Symbol.for("chave");
console.log(global1 === global2); // true
```

### Uso Prático

```javascript
// Criando propriedades "privadas" em objetos
let id = Symbol("id");
let usuario = {
    nome: "João",
    [id]: 12345 // Propriedade com símbolo como chave
};

console.log(usuario.nome); // "João"
console.log(usuario[id]); // 12345
console.log(usuario["id"]); // undefined (não acessa a propriedade com símbolo)

// Iteração não mostra propriedades com símbolo
for (let chave in usuario) {
    console.log(chave); // apenas "nome"
}

// Acessar símbolos explicitamente
let simbolos = Object.getOwnPropertySymbols(usuario);
console.log(simbolos); // [Symbol(id)]
```

---

## 🔧 8. typeof Operator (Operador de Tipo)

### Definição

O operador `typeof` retorna uma **string** indicando o tipo do operando. É a principal ferramenta para verificar tipos de dados em JavaScript.

### Sintaxe

```javascript
typeof operando
typeof (operando) // Parênteses são opcionais
```

### Exemplos com Todos os Tipos

```javascript
// Tipos primitivos
console.log(typeof 42); // "number"
console.log(typeof 42n); // "bigint"
console.log(typeof "texto"); // "string"
console.log(typeof true); // "boolean"
console.log(typeof undefined); // "undefined"
console.log(typeof null); // "object" (bug histórico)
console.log(typeof Symbol("id")); // "symbol"

// Tipos não-primitivos
console.log(typeof {}); // "object"
console.log(typeof []); // "object" (arrays são objetos)
console.log(typeof function() {}); // "function"
console.log(typeof new Date()); // "object"

// Verificações úteis
let valor = null;
console.log(valor === null); // true (forma correta de verificar null)

let arr = [];
console.log(Array.isArray(arr)); // true (forma correta de verificar array)
```

### Verificações de Tipo Mais Confiáveis

```javascript
// Para null
let valor = null;
console.log(valor === null); // true

// Para arrays
let arr = [1, 2, 3];
console.log(Array.isArray(arr)); // true

// Para NaN
let num = NaN;
console.log(Number.isNaN(num)); // true (mais confiável que isNaN())

// Para números inteiros seguros
console.log(Number.isSafeInteger(42)); // true
console.log(Number.isSafeInteger(9007199254740992)); // false
```

---

## 🏗️ 9. Object (Objeto)

### Definição

`Object` é um tipo **não-primitivo** que permite armazenar coleções de dados em formato **chave-valor**. É uma estrutura de dados fundamental em JavaScript.

### Características

- Estrutura de dados chave-valor
- Chaves são strings ou símbolos
- Valores podem ser de qualquer tipo
- Objetos são **mutáveis** (podem ser alterados)
- Arrays e funções são tipos especiais de objetos

### Sintaxe

```javascript
// Objeto literal (forma mais comum)
let pessoa = {
    nome: "João",
    idade: 25,
    cidade: "São Paulo"
};

// Acesso a propriedades
console.log(pessoa.nome); // "João" (notação de ponto)
console.log(pessoa["idade"]); // 25 (notação de colchetes)

// Adicionar propriedades
pessoa.email = "joao@email.com";
pessoa["telefone"] = "123456789";

// Modificar propriedades
pessoa.idade = 26;

// Remover propriedades
delete pessoa.telefone;

console.log(pessoa);
// { nome: "João", idade: 26, cidade: "São Paulo", email: "joao@email.com" }
```

### Propriedades e Métodos

```javascript
let carro = {
    marca: "Toyota",
    modelo: "Corolla",
    ano: 2020,
    
    // Método (função dentro do objeto)
    ligar: function() {
        return "Carro ligado!";
    },
    
    // Método com sintaxe ES6+
    desligar() {
        return "Carro desligado!";
    },
    
    // Método com arrow function (cuidado com 'this')
    acelerar: () => {
        return "Acelerando!";
    }
};

console.log(carro.marca); // "Toyota"
console.log(carro.ligar()); // "Carro ligado!"
console.log(carro.desligar()); // "Carro desligado!"
```

### Objetos Aninhados

```javascript
let empresa = {
    nome: "Tech Corp",
    endereco: {
        rua: "Rua das Flores",
        numero: 123,
        cidade: "São Paulo",
        cep: "01234-567"
    },
    funcionarios: [
        { nome: "João", cargo: "Desenvolvedor" },
        { nome: "Maria", cargo: "Designer" }
    ]
};

console.log(empresa.endereco.cidade); // "São Paulo"
console.log(empresa.funcionarios[0].nome); // "João"
```

### Verificação de Tipo

```javascript
let obj = {};
let arr = [];
let func = function() {};

console.log(typeof obj); // "object"
console.log(typeof arr); // "object" (arrays são objetos)
console.log(typeof func); // "function"

// Verificações mais específicas
console.log(Array.isArray(arr)); // true
console.log(obj instanceof Object); // true
console.log(arr instanceof Array); // true
```

---

## 🛠️ 10. Built-in Objects (Objetos Integrados)

### Definição

Objetos built-in (ou objetos globais) são objetos **incorporados na especificação do JavaScript** e disponíveis globalmente. Eles fornecem funcionalidades essenciais da linguagem.

### Principais Objetos Built-in

#### Number

```javascript
// Constantes
console.log(Number.MAX_VALUE); // Maior número possível
console.log(Number.MIN_VALUE); // Menor número positivo possível
console.log(Number.MAX_SAFE_INTEGER); // Maior inteiro seguro
console.log(Number.MIN_SAFE_INTEGER); // Menor inteiro seguro

// Métodos
console.log(Number.isInteger(42)); // true
console.log(Number.isNaN(NaN)); // true
console.log(Number.parseFloat("3.14")); // 3.14
console.log(Number.parseInt("42")); // 42
```

#### Math

```javascript
// Constantes
console.log(Math.PI); // 3.141592653589793
console.log(Math.E); // 2.718281828459045

// Métodos
console.log(Math.abs(-5)); // 5 (valor absoluto)
console.log(Math.round(3.7)); // 4 (arredondar)
console.log(Math.floor(3.7)); // 3 (arredondar para baixo)
console.log(Math.ceil(3.2)); // 4 (arredondar para cima)
console.log(Math.max(1, 2, 3)); // 3 (máximo)
console.log(Math.min(1, 2, 3)); // 1 (mínimo)
console.log(Math.random()); // Número aleatório entre 0 e 1
console.log(Math.pow(2, 3)); // 8 (potência)
console.log(Math.sqrt(16)); // 4 (raiz quadrada)
```

#### String

```javascript
// Métodos de string
let texto = "JavaScript";

console.log(texto.length); // 10
console.log(texto.toUpperCase()); // "JAVASCRIPT"
console.log(texto.toLowerCase()); // "javascript"
console.log(texto.charAt(0)); // "J"
console.log(texto.indexOf("S")); // 4
console.log(texto.substring(0, 4)); // "Java"
console.log(texto.includes("Script")); // true
console.log(texto.replace("Java", "Type")); // "TypeScript"
```

#### Date

```javascript
// Criar data atual
let agora = new Date();
console.log(agora); // Data e hora atual

// Criar data específica
let data = new Date(2024, 0, 15); // 15 de janeiro de 2024
let data2 = new Date("2024-01-15");

// Métodos
console.log(agora.getFullYear()); // Ano atual
console.log(agora.getMonth()); // Mês (0-11)
console.log(agora.getDate()); // Dia do mês
console.log(agora.getDay()); // Dia da semana (0-6)
console.log(agora.getHours()); // Hora
console.log(agora.getMinutes()); // Minutos
```

#### Error

```javascript
// Criar erro
let erro = new Error("Algo deu errado!");
console.log(erro.message); // "Algo deu errado!"
console.log(erro.name); // "Error"

// Tipos de erro
let erroTipo = new TypeError("Tipo inválido");
let erroRef = new ReferenceError("Referência inválida");
let erroSintaxe = new SyntaxError("Erro de sintaxe");
```

#### Function

```javascript
// Funções são objetos
function minhaFuncao() {
    return "Olá!";
}

console.log(typeof minhaFuncao); // "function"
console.log(minhaFuncao.name); // "minhaFuncao"
console.log(minhaFuncao.length); // 0 (número de parâmetros)
```

#### Boolean

```javascript
// Construtor Boolean (evite usar como construtor)
let bool1 = Boolean(true); // true
let bool2 = new Boolean(true); // Objeto Boolean (não primitivo)

console.log(typeof bool1); // "boolean"
console.log(typeof bool2); // "object"
```

---

## 🔄 Conversão de Tipos (Type Coercion)

JavaScript realiza conversões automáticas de tipos em certas situações. É importante entender isso para evitar bugs.

### Conversão Implícita

```javascript
// String + Number = String
console.log("5" + 3); // "53" (concatenação)
console.log("5" - 3); // 2 (subtração força conversão para número)

// Boolean em operações numéricas
console.log(true + 1); // 2 (true vira 1)
console.log(false + 1); // 1 (false vira 0)

// Comparações
console.log("5" == 5); // true (conversão implícita)
console.log("5" === 5); // false (sem conversão, tipos diferentes)
```

### Conversão Explícita

```javascript
// Para String
let num = 42;
console.log(String(num)); // "42"
console.log(num.toString()); // "42"
console.log(`${num}`); // "42"

// Para Number
let texto = "42";
console.log(Number(texto)); // 42
console.log(parseInt(texto)); // 42
console.log(parseFloat("3.14")); // 3.14
console.log(+texto); // 42 (operador unário +)

// Para Boolean
let valor = 1;
console.log(Boolean(valor)); // true
console.log(!!valor); // true (negação dupla)
```

---

## 📋 Resumo dos Tipos de Dados

| Tipo | Descrição | Exemplo | typeof |
|------|-----------|---------|--------|
| **Number** | Números (inteiros e decimais) | `42`, `3.14` | `"number"` |
| **BigInt** | Números inteiros muito grandes | `42n` | `"bigint"` |
| **String** | Sequência de caracteres | `"texto"` | `"string"` |
| **Boolean** | Valores lógicos | `true`, `false` | `"boolean"` |
| **Undefined** | Variável não inicializada | `undefined` | `"undefined"` |
| **Null** | Ausência intencional de valor | `null` | `"object"` |
| **Symbol** | Identificador único | `Symbol("id")` | `"symbol"` |
| **Object** | Estrutura chave-valor | `{}`, `[]`, `function(){}` | `"object"` |

---

## 🎯 Próximos Passos

Agora que você entende os tipos de dados em JavaScript, você está pronto para:
- ✅ Trabalhar com operadores (próxima aula)
- ✅ Entender como os tipos se comportam em operações
- ✅ Evitar erros comuns de conversão de tipos
- ✅ Usar os tipos corretamente em suas aplicações

**Continue para a Aula Simplificada para ver esses conceitos explicados de forma mais acessível!**



