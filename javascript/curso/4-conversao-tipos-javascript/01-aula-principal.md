# Aula 4: Conversão de Tipos (Type Casting) em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 3**, você aprendeu:
- ✅ Os diferentes tipos de dados em JavaScript (Number, String, Boolean, etc.)
- ✅ Diferença entre tipos primitivos e não-primitivos
- ✅ Como usar o operador `typeof` para identificar tipos
- ✅ Características específicas de cada tipo de dado

Agora vamos aprender sobre **como converter valores entre diferentes tipos de dados** - uma habilidade essencial em JavaScript!

---

## 🎯 O que é Conversão de Tipos (Type Casting)?

**Definição:** Conversão de tipos (ou type casting) é o processo de **transferir dados de um tipo de dado para outro**. Em JavaScript, isso pode acontecer de forma automática (implícita) ou manual (explícita).

### Conceitos Fundamentais

1. **Type Conversion (Conversão de Tipos)**: Processo geral de transformar um valor de um tipo para outro
2. **Type Coercion (Coerção de Tipos)**: Conversão automática e implícita realizada pelo JavaScript
3. **Type Casting (Type Casting Explícito)**: Conversão manual e explícita realizada pelo desenvolvedor

---

## 🔄 Classificação das Conversões

JavaScript possui dois tipos principais de conversão:

1. **Conversão Implícita (Type Coercion)**: JavaScript converte automaticamente
2. **Conversão Explícita (Type Casting)**: Você especifica explicitamente a conversão

---

## 🤖 1. Conversão Implícita (Type Coercion)

### Definição

A conversão implícita acontece quando o JavaScript **automaticamente converte** um valor de um tipo para outro durante operações. Isso ocorre porque JavaScript é uma linguagem **fracamente tipada** (loosely typed).

### Quando Acontece?

A coerção de tipos ocorre em várias situações:

1. **Operações aritméticas** com tipos diferentes
2. **Comparações** entre valores de tipos diferentes
3. **Concatenação de strings** com outros tipos
4. **Contextos booleanos** (if, while, operadores lógicos)

### Exemplos de Conversão Implícita

#### 1.1. Conversão em Operações Aritméticas

```javascript
// String para Number (adição vs concatenação)
console.log("10" + 5);        // "105" (concatenação - string ganha)
console.log("10" - 5);        // 5 (subtração força conversão para número)
console.log("10" * 5);        // 50 (multiplicação força conversão)
console.log("10" / 5);        // 2 (divisão força conversão)
console.log("10" % 3);        // 1 (módulo força conversão)

// Number para String
console.log(10 + "5");        // "105" (concatenação)
console.log(10 + "");         // "10" (número vira string)

// Boolean para Number
console.log(true + 1);        // 2 (true = 1)
console.log(false + 1);       // 1 (false = 0)
console.log(true + true);     // 2

// Null e Undefined
console.log(null + 5);        // 5 (null vira 0)
console.log(undefined + 5);  // NaN (undefined não converte para número)
```

#### 1.2. Conversão em Comparações

```javascript
// Comparação com == (permite coerção)
console.log("5" == 5);        // true (string "5" vira número 5)
console.log(true == 1);       // true (true vira 1)
console.log(false == 0);      // true (false vira 0)
console.log(null == undefined); // true (regra especial)
console.log("" == 0);         // true (string vazia vira 0)
console.log(" " == 0);        // true (espaço vira 0)

// Comparação com === (não permite coerção)
console.log("5" === 5);       // false (tipos diferentes)
console.log(true === 1);     // false (tipos diferentes)
console.log(false === 0);    // false (tipos diferentes)
```

#### 1.3. Conversão em Contextos Booleanos

```javascript
// Valores "falsy" (convertem para false)
if ("") console.log("não executa");
if (0) console.log("não executa");
if (null) console.log("não executa");
if (undefined) console.log("não executa");
if (NaN) console.log("não executa");
if (false) console.log("não executa");

// Valores "truthy" (convertem para true)
if ("texto") console.log("executa");
if (1) console.log("executa");
if ([]) console.log("executa"); // array vazio é truthy!
if ({}) console.log("executa"); // objeto vazio é truthy!
if (function(){}) console.log("executa");
```

#### 1.4. Conversão em Operadores Lógicos

```javascript
// Operador && (retorna o primeiro falsy ou o último valor)
console.log("texto" && 5);           // 5
console.log("" && 5);                // "" (primeiro falsy)
console.log(null && 5);              // null

// Operador || (retorna o primeiro truthy ou o último valor)
console.log("texto" || 5);           // "texto" (primeiro truthy)
console.log("" || 5);                // 5
console.log(null || undefined || 5); // 5

// Operador ?? (nullish coalescing - ES2020)
console.log(null ?? 5);              // 5
console.log(undefined ?? 5);         // 5
console.log(0 ?? 5);                 // 0 (não é null/undefined)
console.log("" ?? 5);                // "" (não é null/undefined)
```

### Regras de Coerção de Tipos

#### String para Number

```javascript
// Conversão automática em operações matemáticas
console.log("123" - 0);      // 123
console.log("123" * 1);      // 123
console.log("123" / 1);      // 123
console.log(+"123");         // 123 (operador unário +)

// Casos especiais
console.log("123abc" - 0);   // NaN (não é número válido)
console.log("" - 0);         // 0 (string vazia vira 0)
console.log(" " - 0);        // 0 (espaços viram 0)
console.log("12.5" - 0);     // 12.5 (decimais funcionam)
```

#### Number para String

```javascript
// Conversão automática em concatenação
console.log(123 + "");       // "123"
console.log(123 + "abc");    // "123abc"
console.log("" + 123);       // "123"

// Template literals também convertem
console.log(`${123}`);       // "123"
```

#### Boolean para Number

```javascript
console.log(true * 1);       // 1
console.log(false * 1);      // 0
console.log(true + true);    // 2
console.log(false + false);  // 0
```

#### Outros para Boolean

```javascript
// Valores Falsy (6 valores)
Boolean("");         // false
Boolean(0);          // false
Boolean(-0);         // false
Boolean(null);       // false
Boolean(undefined); // false
Boolean(NaN);        // false
Boolean(false);      // false

// Valores Truthy (tudo mais)
Boolean("texto");    // true
Boolean(1);          // true
Boolean(-1);         // true
Boolean([]);         // true (array vazio!)
Boolean({});         // true (objeto vazio!)
Boolean(function(){}); // true
```

---

## ✋ 2. Conversão Explícita (Type Casting)

### Definição

A conversão explícita acontece quando você **intencionalmente converte** um valor de um tipo para outro usando métodos específicos ou operadores. Isso dá mais controle e clareza ao código.

### Por que Usar Conversão Explícita?

1. **Clareza**: Deixa explícito o que você está fazendo
2. **Previsibilidade**: Evita comportamentos inesperados
3. **Manutenibilidade**: Código mais fácil de entender
4. **Segurança**: Reduz erros de coerção implícita

---

## 🔢 2.1. Conversão para Number

### Métodos Disponíveis

#### Number()

Converte qualquer valor para número. Retorna `NaN` se a conversão não for possível.

```javascript
// Strings numéricas
console.log(Number("123"));      // 123
console.log(Number("12.5"));     // 12.5
console.log(Number("123abc"));   // NaN
console.log(Number(""));         // 0
console.log(Number(" "));       // 0

// Boolean
console.log(Number(true));       // 1
console.log(Number(false));      // 0

// Null e Undefined
console.log(Number(null));       // 0
console.log(Number(undefined));  // NaN

// Objetos
console.log(Number([]));         // 0 (array vazio)
console.log(Number([5]));        // 5 (array com um número)
console.log(Number([1,2,3]));    // NaN (array com múltiplos elementos)
console.log(Number({}));         // NaN

// Valores especiais
console.log(Number(NaN));        // NaN
console.log(Number(Infinity));   // Infinity
```

#### parseInt()

Converte string para número inteiro. Para na primeira ocorrência de caractere não numérico.

```javascript
// Conversão básica
console.log(parseInt("123"));        // 123
console.log(parseInt("12.5"));       // 12 (para no ponto)
console.log(parseInt("123abc"));     // 123 (para no 'a')
console.log(parseInt("abc123"));     // NaN (começa com letra)

// Com base numérica (radix)
console.log(parseInt("10", 10));     // 10 (decimal)
console.log(parseInt("10", 2));      // 2 (binário)
console.log(parseInt("10", 8));      // 8 (octal)
console.log(parseInt("10", 16));     // 16 (hexadecimal)
console.log(parseInt("FF", 16));     // 255

// Casos especiais
console.log(parseInt(""));           // NaN
console.log(parseInt(" "));          // NaN (espaços são ignorados, mas string vazia retorna NaN)
console.log(parseInt("   123"));     // 123 (espaços iniciais são ignorados)
console.log(parseInt(null));         // NaN
console.log(parseInt(undefined));    // NaN

// IMPORTANTE: Sempre especifique a base!
console.log(parseInt("08"));         // 8 (em alguns navegadores antigos seria 0 - octal)
console.log(parseInt("08", 10));     // 8 (sempre seguro)
```

#### parseFloat()

Converte string para número de ponto flutuante (decimal).

```javascript
// Conversão básica
console.log(parseFloat("123"));      // 123
console.log(parseFloat("12.5"));     // 12.5
console.log(parseFloat("12.5.7"));   // 12.5 (para no segundo ponto)
console.log(parseFloat("123abc"));   // 123 (para no 'a')
console.log(parseFloat("abc123"));   // NaN

// Notação científica
console.log(parseFloat("1.5e2"));    // 150
console.log(parseFloat("1.5e-2"));   // 0.015

// Casos especiais
console.log(parseFloat(""));         // NaN
console.log(parseFloat(" "));        // NaN
console.log(parseFloat(null));       // NaN
console.log(parseFloat(undefined));  // NaN
```

#### Operador Unário + (Plus)

Forma rápida de converter para número.

```javascript
console.log(+"123");         // 123
console.log(+"12.5");        // 12.5
console.log(+"123abc");      // NaN
console.log(+true);          // 1
console.log(+false);         // 0
console.log(+"");            // 0
console.log(+null);          // 0
console.log(+undefined);     // NaN
```

### Comparação dos Métodos

```javascript
let valor = "123.45";

console.log(Number(valor));      // 123.45
console.log(parseInt(valor));    // 123 (perde decimais)
console.log(parseFloat(valor));  // 123.45
console.log(+valor);             // 123.45

// Diferença importante
let valor2 = "123abc";
console.log(Number(valor2));     // NaN (conversão estrita)
console.log(parseInt(valor2));   // 123 (tenta converter o que pode)
console.log(parseFloat(valor2)); // 123
console.log(+valor2);            // NaN
```

---

## 📝 2.2. Conversão para String

### Métodos Disponíveis

#### String()

Converte qualquer valor para string.

```javascript
// Números
console.log(String(123));        // "123"
console.log(String(12.5));       // "12.5"
console.log(String(0));          // "0"
console.log(String(-0));         // "0"
console.log(String(NaN));        // "NaN"
console.log(String(Infinity));   // "Infinity"

// Boolean
console.log(String(true));       // "true"
console.log(String(false));      // "false"

// Null e Undefined
console.log(String(null));       // "null"
console.log(String(undefined));  // "undefined"

// Objetos
console.log(String([]));         // "" (array vazio)
console.log(String([1,2,3]));    // "1,2,3"
console.log(String({}));         // "[object Object]"
console.log(String({a: 1}));     // "[object Object]"
```

#### .toString()

Método disponível na maioria dos valores. **Não funciona com `null` e `undefined`**.

```javascript
// Números
console.log((123).toString());       // "123"
console.log((12.5).toString());      // "12.5"
console.log((0).toString());         // "0"
console.log((NaN).toString());       // "NaN"
console.log((Infinity).toString());  // "Infinity"

// Boolean
console.log(true.toString());        // "true"
console.log(false.toString());       // "false"

// Arrays
console.log([].toString());          // ""
console.log([1,2,3].toString());     // "1,2,3"
console.log([1,"a",true].toString()); // "1,a,true"

// Objetos
console.log({}.toString());          // "[object Object]"

// ERRO: null e undefined não têm toString()
// console.log(null.toString());     // TypeError
// console.log(undefined.toString()); // TypeError
```

#### Template Literals (Template Strings)

Usando template literals para conversão implícita.

```javascript
let numero = 123;
let booleano = true;
let nulo = null;

console.log(`${numero}`);        // "123"
console.log(`${booleano}`);      // "true"
console.log(`${nulo}`);          // "null"
console.log(`${undefined}`);     // "undefined"
```

#### Concatenação com String Vazia

Forma rápida de converter para string.

```javascript
console.log(123 + "");           // "123"
console.log(true + "");          // "true"
console.log(null + "");          // "null"
console.log(undefined + "");     // "undefined"
console.log([] + "");            // ""
console.log([1,2,3] + "");       // "1,2,3"
```

### Comparação dos Métodos

```javascript
let valor = 123;

console.log(String(valor));      // "123"
console.log(valor.toString());   // "123"
console.log(`${valor}`);         // "123"
console.log(valor + "");         // "123"

// Diferença com null/undefined
console.log(String(null));       // "null" ✅
// console.log(null.toString()); // TypeError ❌
console.log(`${null}`);          // "null" ✅
console.log(null + "");          // "null" ✅
```

---

## ✅ 2.3. Conversão para Boolean

### Métodos Disponíveis

#### Boolean()

Converte qualquer valor para boolean explicitamente.

```javascript
// Valores Falsy (convertem para false)
console.log(Boolean(""));        // false
console.log(Boolean(0));         // false
console.log(Boolean(-0));       // false
console.log(Boolean(null));      // false
console.log(Boolean(undefined)); // false
console.log(Boolean(NaN));       // false
console.log(Boolean(false));     // false

// Valores Truthy (convertem para true)
console.log(Boolean("texto"));   // true
console.log(Boolean(1));         // true
console.log(Boolean(-1));        // true
console.log(Boolean([]));        // true (array vazio é truthy!)
console.log(Boolean({}));        // true (objeto vazio é truthy!)
console.log(Boolean(function(){})); // true
console.log(Boolean("0"));       // true (string "0" é truthy!)
console.log(Boolean("false"));   // true (string "false" é truthy!)
```

#### Operador !!

Dupla negação - forma rápida de converter para boolean.

```javascript
console.log(!!"texto");          // true
console.log(!!"");               // false
console.log(!!0);                // false
console.log(!!1);                // true
console.log(!!null);             // false
console.log(!!undefined);        // false
console.log(!![]);               // true
console.log(!!{});               // true
```

### Comparação

```javascript
let valor = "texto";

console.log(Boolean(valor));     // true
console.log(!!valor);            // true

// Ambos funcionam igual, mas Boolean() é mais legível
```

---

## 🔄 2.4. Conversões Especiais

### Arrays para String

```javascript
// toString() - converte para string separada por vírgulas
console.log([1,2,3].toString());           // "1,2,3"
console.log([].toString());                // ""
console.log(["a","b","c"].toString());    // "a,b,c"

// join() - permite especificar o separador
console.log([1,2,3].join());              // "1,2,3" (vírgula padrão)
console.log([1,2,3].join(""));            // "123" (sem separador)
console.log([1,2,3].join("-"));           // "1-2-3"
console.log([1,2,3].join(" | "));         // "1 | 2 | 3"
```

### Objetos para String

```javascript
// toString() - sempre retorna "[object Object]"
console.log({}.toString());                // "[object Object]"
console.log({a: 1, b: 2}.toString());     // "[object Object]"

// JSON.stringify() - converte para JSON string
console.log(JSON.stringify({a: 1, b: 2})); // '{"a":1,"b":2}'
console.log(JSON.stringify([1,2,3]));       // '[1,2,3]'
console.log(JSON.stringify(null));         // 'null'
console.log(JSON.stringify(undefined));    // undefined (não é string!)

// Valores que não podem ser serializados
let obj = {
    a: 1,
    b: function() {},  // funções são ignoradas
    c: undefined       // undefined é ignorado
};
console.log(JSON.stringify(obj));         // '{"a":1}'
```

### Conversão de Tipos em Objetos

```javascript
// valueOf() - retorna o valor primitivo do objeto
let numObj = new Number(123);
console.log(numObj.valueOf());            // 123 (número primitivo)

let strObj = new String("texto");
console.log(strObj.valueOf());            // "texto" (string primitiva)

let boolObj = new Boolean(true);
console.log(boolObj.valueOf());           // true (boolean primitivo)

// Conversão automática em operações
console.log(numObj + 1);                  // 124 (conversão automática)
console.log(strObj + " mais");            // "texto mais" (conversão automática)
```

---

## ⚠️ 3. Armadilhas e Comportamentos Inesperados

### 3.1. Comparações com == vs ===

```javascript
// == permite coerção (pode ser perigoso)
console.log("" == 0);             // true ⚠️
console.log(" " == 0);           // true ⚠️
console.log("\n" == 0);          // true ⚠️
console.log("\t" == 0);          // true ⚠️
console.log(null == undefined);  // true (regra especial)
console.log([] == 0);            // true ⚠️ (array vazio vira 0)
console.log([0] == false);       // true ⚠️

// === não permite coerção (mais seguro)
console.log("" === 0);           // false ✅
console.log(" " === 0);          // false ✅
console.log(null === undefined); // false ✅
console.log([] === 0);           // false ✅
```

### 3.2. Arrays e Objetos Vazios

```javascript
// Arrays vazios são truthy!
if ([]) {
    console.log("Array vazio é truthy!"); // executa
}

// Objetos vazios são truthy!
if ({}) {
    console.log("Objeto vazio é truthy!"); // executa
}

// Mas em comparações...
console.log([] == false);        // true ⚠️
console.log([] == 0);            // true ⚠️
console.log([].toString());      // "" (string vazia)
console.log(Number([]));         // 0
```

### 3.3. NaN (Not a Number)

```javascript
// NaN é único - não é igual a nada, nem a si mesmo!
console.log(NaN == NaN);         // false ⚠️
console.log(NaN === NaN);        // false ⚠️
console.log(NaN != NaN);         // true ⚠️

// Como verificar NaN?
console.log(Number.isNaN(NaN));  // true ✅
console.log(isNaN(NaN));         // true (mas cuidado!)
console.log(isNaN("texto"));     // true (converte primeiro)
console.log(Number.isNaN("texto")); // false (não converte)
```

### 3.4. Null vs Undefined

```javascript
// Em comparações
console.log(null == undefined);  // true (regra especial)
console.log(null === undefined); // false (tipos diferentes)

// Em conversões numéricas
console.log(Number(null));       // 0
console.log(Number(undefined));  // NaN

// Em conversões booleanas
console.log(Boolean(null));      // false
console.log(Boolean(undefined)); // false

// Em conversões de string
console.log(String(null));       // "null"
console.log(String(undefined));  // "undefined"
```

### 3.5. Strings que Parecem Números

```javascript
// Cuidado com strings que contêm números
console.log("123" == 123);       // true
console.log("123" === 123);      // false
console.log("0" == false);       // true ⚠️
console.log("0" === false);      // false ✅
console.log("" == 0);            // true ⚠️
console.log("" === 0);           // false ✅
```

---

## 🎯 4. Boas Práticas

### 4.1. Sempre Use Conversão Explícita

```javascript
// ❌ Ruim - conversão implícita
let idade = "25";
let novaIdade = idade + 1; // "251" (erro!)

// ✅ Bom - conversão explícita
let idade = "25";
let novaIdade = Number(idade) + 1; // 26
// ou
let novaIdade = parseInt(idade, 10) + 1; // 26
```

### 4.2. Use === em vez de ==

```javascript
// ❌ Ruim - pode ter comportamentos inesperados
if (valor == 0) { }

// ✅ Bom - comparação estrita
if (valor === 0) { }
```

### 4.3. Valide Antes de Converter

```javascript
// ❌ Ruim - pode retornar NaN
let numero = Number(entradaUsuario);

// ✅ Bom - valida primeiro
function converterParaNumero(valor) {
    if (typeof valor === 'number') return valor;
    if (typeof valor === 'string' && valor.trim() === '') return NaN;
    let numero = Number(valor);
    return isNaN(numero) ? NaN : numero;
}
```

### 4.4. Use parseInt com Base Numérica

```javascript
// ❌ Ruim - pode ter comportamento inconsistente
let numero = parseInt("08"); // pode ser 0 ou 8 dependendo do navegador

// ✅ Bom - sempre especifique a base
let numero = parseInt("08", 10); // sempre 8
```

### 4.5. Prefira Number() para Conversões Numéricas

```javascript
// ❌ Ruim - parseInt pode perder decimais
let preco = parseInt("12.99"); // 12 (perdeu os centavos!)

// ✅ Bom - use Number() ou parseFloat()
let preco = Number("12.99"); // 12.99
// ou
let preco = parseFloat("12.99"); // 12.99
```

---

## 📚 5. Resumo das Conversões

### Para Number

| Método | Uso | Quando Usar |
|--------|-----|-------------|
| `Number()` | Conversão geral | Quando quer conversão estrita |
| `parseInt()` | String → Inteiro | Quando precisa de número inteiro |
| `parseFloat()` | String → Decimal | Quando precisa preservar decimais |
| `+valor` | Operador unário | Forma rápida e concisa |

### Para String

| Método | Uso | Quando Usar |
|--------|-----|-------------|
| `String()` | Conversão geral | Funciona com null/undefined |
| `.toString()` | Método do objeto | Não funciona com null/undefined |
| `` `${valor}` `` | Template literal | Moderno e legível |
| `valor + ""` | Concatenação | Forma rápida |

### Para Boolean

| Método | Uso | Quando Usar |
|--------|-----|-------------|
| `Boolean()` | Conversão explícita | Mais legível |
| `!!valor` | Dupla negação | Forma rápida |

---

## 🎓 Conclusão

Nesta aula, você aprendeu:

- ✅ O que é conversão de tipos (type casting) e coerção de tipos (type coercion)
- ✅ Como JavaScript converte tipos automaticamente (conversão implícita)
- ✅ Como converter tipos manualmente (conversão explícita)
- ✅ Métodos para converter para Number, String e Boolean
- ✅ Armadilhas e comportamentos inesperados
- ✅ Boas práticas para evitar erros

**Próximo passo:** Na próxima aula, você verá como aplicar essas conversões em situações práticas e aprenderá sobre operadores em JavaScript!

---

**Lembre-se:** Em JavaScript, a conversão de tipos é poderosa, mas pode ser traiçoeira. Sempre prefira conversões explícitas e use `===` para comparações estritas!

