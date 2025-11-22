# Aula 9: Expressões e Operadores em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 8**, você aprendeu:
- ✅ Control Flow (Fluxo de Controle)
- ✅ Estruturas condicionais (`if/else`, `switch`)
- ✅ Tratamento de exceções (`try/catch/finally`)
- ✅ Como controlar a execução do código

Agora vamos aprender sobre **Expressões e Operadores** - os blocos fundamentais que permitem manipular valores, fazer cálculos e tomar decisões no JavaScript!

---

## 🎯 O que são Expressões?

**Definição:** Uma expressão é uma unidade válida de código que resolve para um valor. Em outras palavras, é qualquer código que produz um resultado quando executado.

### Tipos de Expressões

Existem dois tipos principais de expressões:

1. **Expressões com Efeitos Colaterais (Side Effects)**: Expressões que fazem algo além de apenas produzir um valor
2. **Expressões Puramente Avaliativas**: Expressões que apenas calculam e retornam um valor

### Exemplos Práticos

```javascript
// Expressão com efeito colateral: atribui valor e retorna o valor atribuído
let x = 7;  // A expressão x = 7 atribui 7 a x e retorna 7

// Expressão puramente avaliativa: apenas calcula um valor
3 + 4;  // Retorna 7, mas não faz nada com esse valor (geralmente um erro do programador)

// Expressão útil: calcula e armazena o resultado
const z = 3 + 4;  // Calcula 7 e armazena em z
```

### Características das Expressões

- **Sempre produzem um valor**: Mesmo que seja `undefined` ou `null`
- **Podem ser combinadas**: Expressões complexas são formadas por expressões menores
- **Usam operadores**: Operadores conectam valores e expressões
- **Têm precedência**: Alguns operadores são avaliados antes de outros

---

## 🔧 Operadores em JavaScript

**Definição:** Operadores são símbolos especiais que realizam operações em valores (operandos) e produzem um resultado.

### Classificação dos Operadores

Os operadores podem ser classificados por:
- **Número de operandos**: Unário (1), Binário (2), Ternário (3)
- **Tipo de operação**: Aritmética, Lógica, Comparação, Atribuição, etc.
- **Precedência**: Ordem de avaliação

---

## 1️⃣ Operadores Aritméticos

Os operadores aritméticos realizam operações matemáticas básicas.

### Operadores Básicos

```javascript
// Adição (+)
let soma = 5 + 3;        // 8
let concatenacao = "Olá" + " " + "Mundo";  // "Olá Mundo"

// Subtração (-)
let subtracao = 10 - 4;  // 6

// Multiplicação (*)
let multiplicacao = 6 * 7;  // 42

// Divisão (/)
let divisao = 15 / 3;    // 5
let divisaoDecimal = 10 / 3;  // 3.3333333333333335

// Módulo/Remainder (%)
let resto = 17 % 5;      // 2 (17 dividido por 5 = 3 com resto 2)
let par = 8 % 2;         // 0 (número par)
let impar = 7 % 2;       // 1 (número ímpar)

// Exponenciação (**) - ES2016+
let potencia = 2 ** 3;   // 8 (2 elevado a 3)
let raizQuadrada = 16 ** 0.5;  // 4
```

### Operadores de Incremento e Decremento

```javascript
// Incremento (++)
let contador = 5;

// Pré-incremento: incrementa ANTES de usar o valor
let preIncremento = ++contador;  // contador vira 6, preIncremento = 6

// Pós-incremento: incrementa DEPOIS de usar o valor
contador = 5;
let posIncremento = contador++;  // posIncremento = 5, contador vira 6

// Decremento (--)
let numero = 10;
let preDecremento = --numero;    // numero vira 9, preDecremento = 9

numero = 10;
let posDecremento = numero--;    // posDecremento = 10, numero vira 9
```

### Comportamento Especial com Strings

```javascript
// O operador + pode ser usado para concatenar strings
let nome = "João";
let sobrenome = "Silva";
let nomeCompleto = nome + " " + sobrenome;  // "João Silva"

// Cuidado: números podem ser convertidos para strings
let resultado = "5" + 3;  // "53" (não 8!)
let resultado2 = "5" - 3; // 2 (subtração força conversão para número)
```

---

## 2️⃣ Operadores de Atribuição

Os operadores de atribuição atribuem valores a variáveis.

### Operador de Atribuição Simples (=)

```javascript
let x = 10;        // Atribui 10 a x
let y = x;         // Atribui o valor de x (10) a y
let z = x = 20;    // Atribui 20 a x, depois atribui o resultado (20) a z
```

### Operadores de Atribuição Compostos

```javascript
let valor = 10;

// Adição e atribuição (+=)
valor += 5;        // Equivale a: valor = valor + 5 (valor = 15)

// Subtração e atribuição (-=)
valor -= 3;        // Equivale a: valor = valor - 3 (valor = 12)

// Multiplicação e atribuição (*=)
valor *= 2;        // Equivale a: valor = valor * 2 (valor = 24)

// Divisão e atribuição (/=)
valor /= 4;        // Equivale a: valor = valor / 4 (valor = 6)

// Módulo e atribuição (%=)
valor %= 4;        // Equivale a: valor = valor % 4 (valor = 2)

// Exponenciação e atribuição (**=)
valor **= 3;       // Equivale a: valor = valor ** 3 (valor = 8)
```

### Atribuição com Strings

```javascript
let mensagem = "Olá";
mensagem += " Mundo";  // "Olá Mundo" (concatenação)
```

---

## 3️⃣ Operadores de Comparação

Os operadores de comparação comparam dois valores e retornam um booleano (`true` ou `false`).

### Operadores de Comparação de Valor

```javascript
// Igualdade (==) - Compara valores com conversão de tipo
5 == 5;        // true
5 == "5";      // true (converte string para número)
true == 1;     // true (converte boolean para número)
false == 0;    // true

// Desigualdade (!=)
5 != 3;        // true
5 != "5";      // false (mesmo valor após conversão)
```

### Operadores de Comparação Estrita

```javascript
// Igualdade estrita (===) - Compara valores E tipos
5 === 5;       // true
5 === "5";     // false (tipos diferentes: number vs string)
true === 1;    // false (tipos diferentes: boolean vs number)

// Desigualdade estrita (!==)
5 !== "5";     // true (valores ou tipos diferentes)
5 !== 5;       // false
```

**⚠️ IMPORTANTE:** Sempre prefira `===` e `!==` sobre `==` e `!=` para evitar bugs sutis!

### Operadores de Relação

```javascript
// Maior que (>)
10 > 5;        // true
5 > 10;        // false
"b" > "a";     // true (comparação lexicográfica)

// Menor que (<)
5 < 10;        // true
10 < 5;        // false

// Maior ou igual (>=)
10 >= 10;      // true
10 >= 5;       // true
5 >= 10;       // false

// Menor ou igual (<=)
5 <= 5;        // true
5 <= 10;       // true
10 <= 5;       // false
```

### Comparações Especiais

```javascript
// Comparações com NaN
NaN == NaN;    // false (NaN nunca é igual a nada, nem a si mesmo)
NaN === NaN;   // false
isNaN(NaN);    // true (use isNaN() para verificar NaN)

// Comparações com null e undefined
null == undefined;   // true
null === undefined;  // false (tipos diferentes)
null == 0;          // false
undefined == 0;     // false
```

---

## 4️⃣ Operadores Lógicos

Os operadores lógicos são usados para combinar ou inverter valores booleanos.

### Operador AND (&&)

Retorna `true` se **ambos** os operandos forem verdadeiros.

```javascript
true && true;      // true
true && false;     // false
false && true;     // false
false && false;    // false

// Exemplos práticos
let idade = 25;
let temCarteira = true;
let podeDirigir = idade >= 18 && temCarteira;  // true

// Comportamento de curto-circuito
let resultado = false && console.log("Não executa");  // Não imprime nada
let resultado2 = true && console.log("Executa");      // Imprime "Executa"
```

**Comportamento de Curto-Circuito:**
- Se o primeiro operando for `false`, o segundo **não é avaliado**
- Retorna o primeiro valor falsy ou o último valor se todos forem truthy

```javascript
let valor = 0 && 10;        // 0 (primeiro falsy)
let valor2 = 5 && 10;       // 10 (último truthy)
let valor3 = null && 10;    // null (primeiro falsy)
```

### Operador OR (||)

Retorna `true` se **pelo menos um** dos operandos for verdadeiro.

```javascript
true || true;      // true
true || false;     // true
false || true;     // true
false || false;    // false

// Exemplos práticos
let nome = "";
let nomePadrao = nome || "Anônimo";  // "Anônimo" (se nome for falsy)

// Comportamento de curto-circuito
let resultado = true || console.log("Não executa");  // Não imprime nada
```

**Comportamento de Curto-Circuito:**
- Se o primeiro operando for `true`, o segundo **não é avaliado**
- Retorna o primeiro valor truthy ou o último valor se todos forem falsy

```javascript
let valor = 0 || 10;        // 10 (primeiro truthy encontrado)
let valor2 = 5 || 10;       // 5 (primeiro truthy)
let valor3 = null || 0 || "padrão";  // "padrão" (primeiro truthy)
```

### Operador NOT (!)

Inverte o valor booleano.

```javascript
!true;           // false
!false;          // true
!0;              // true (0 é falsy)
!1;              // false (1 é truthy)
!"";             // true (string vazia é falsy)
!"texto";        // false (string não vazia é truthy)
```

### Operador Nullish Coalescing (??) - ES2020

Retorna o operando direito apenas se o esquerdo for `null` ou `undefined`.

```javascript
let valor1 = null ?? "padrão";        // "padrão"
let valor2 = undefined ?? "padrão";   // "padrão"
let valor3 = 0 ?? "padrão";           // 0 (não é null nem undefined)
let valor4 = "" ?? "padrão";          // "" (não é null nem undefined)
let valor5 = false ?? "padrão";       // false

// Diferença entre || e ??
let valor6 = 0 || "padrão";           // "padrão" (0 é falsy)
let valor7 = 0 ?? "padrão";           // 0 (0 não é null nem undefined)
```

**Quando usar `??` vs `||`:**
- Use `||` quando quiser usar um valor padrão para qualquer valor falsy
- Use `??` quando quiser usar um valor padrão apenas para `null` ou `undefined`

---

## 5️⃣ Operador Ternário (Condicional)

O operador ternário é o único operador que usa **três operandos**. É uma forma concisa de escrever `if/else`.

### Sintaxe

```javascript
condicao ? valor_se_verdadeiro : valor_se_falso
```

### Exemplos

```javascript
// Exemplo básico
let idade = 20;
let status = idade >= 18 ? "Adulto" : "Menor";
// status = "Adulto"

// Exemplo com múltiplas condições
let nota = 85;
let conceito = nota >= 90 ? "A" : 
               nota >= 80 ? "B" : 
               nota >= 70 ? "C" : "D";
// conceito = "B"

// Exemplo prático
let preco = 100;
let desconto = preco > 50 ? preco * 0.1 : 0;
let precoFinal = preco - desconto;
```

### Quando Usar o Operador Ternário

**✅ Use quando:**
- A lógica é simples e direta
- Você precisa de uma expressão (não uma declaração)
- O código fica mais legível

**❌ Evite quando:**
- A lógica é complexa
- Você precisa de múltiplas declarações
- A legibilidade é prejudicada

---

## 6️⃣ Operadores Unários

Operadores unários trabalham com **um único operando**.

### Operador Unário Plus (+)

Converte o operando para número.

```javascript
+"5";           // 5 (converte string para número)
+true;          // 1
+false;         // 0
+null;          // 0
+undefined;     // NaN
+"abc";         // NaN
```

### Operador Unário Negation (-)

Converte o operando para número e inverte o sinal.

```javascript
-"5";           // -5
-true;          // -1
-false;         // -0
```

### Operador de Incremento/Decremento (já visto)

```javascript
let x = 5;
++x;            // Pré-incremento
x++;            // Pós-incremento
--x;            // Pré-decremento
x--;            // Pós-decremento
```

### Operador typeof

Retorna uma string indicando o tipo do operando.

```javascript
typeof 42;              // "number"
typeof "texto";         // "string"
typeof true;            // "boolean"
typeof undefined;       // "undefined"
typeof null;            // "object" (⚠️ bug conhecido do JavaScript)
typeof {};              // "object"
typeof [];              // "object"
typeof function(){};    // "function"
```

### Operador delete

Remove uma propriedade de um objeto.

```javascript
let objeto = { nome: "João", idade: 30 };
delete objeto.idade;    // true
console.log(objeto);    // { nome: "João" }

// Não pode deletar variáveis
let x = 10;
delete x;               // false (em modo estrito, gera erro)
```

### Operador void

Avalia uma expressão e retorna `undefined`.

```javascript
void 0;                // undefined
void (5 + 3);          // undefined
```

---

## 7️⃣ Operador de Vírgula (,)

O operador de vírgula avalia múltiplas expressões da esquerda para a direita e retorna o valor da última expressão.

### Sintaxe

```javascript
expressao1, expressao2, expressao3
```

### Exemplos

```javascript
// Retorna o último valor
let resultado = (5, 10, 15);  // resultado = 15

// Comum em loops for
for (let i = 0, j = 10; i < 5; i++, j--) {
    console.log(i, j);
}

// Múltiplas atribuições
let a, b, c;
a = 1, b = 2, c = 3;  // Todas são executadas, c = 3 é retornado
```

### Quando Usar

- Em loops `for` para inicializar ou atualizar múltiplas variáveis
- Quando você precisa executar múltiplas expressões mas só precisa do último valor

---

## 8️⃣ Operadores de String

Além dos operadores de comparação, o JavaScript tem operadores específicos para strings.

### Operador de Concatenação (+)

Concatena duas ou mais strings.

```javascript
let nome = "João";
let sobrenome = "Silva";
let nomeCompleto = nome + " " + sobrenome;  // "João Silva"

// Concatenação múltipla
let mensagem = "Olá" + " " + "Mundo" + "!";  // "Olá Mundo!"
```

### Operador de Atribuição com Concatenação (+=)

```javascript
let texto = "Olá";
texto += " ";        // "Olá "
texto += "Mundo";   // "Olá Mundo"
```

### Template Literals (ES6+) - Alternativa Moderna

```javascript
let nome = "João";
let idade = 30;
let mensagem = `Olá, meu nome é ${nome} e tenho ${idade} anos.`;
// "Olá, meu nome é João e tenho 30 anos."
```

---

## 📊 Precedência de Operadores

A precedência determina a ordem em que os operadores são avaliados. Operadores com maior precedência são avaliados primeiro.

### Tabela de Precedência (Principais)

1. **Agrupamento**: `()`
2. **Acesso/Membro**: `.`, `[]`, `()`
3. **Unários**: `!`, `typeof`, `+`, `-`, `++`, `--`
4. **Multiplicativos**: `*`, `/`, `%`
5. **Aditivos**: `+`, `-`
6. **Relacionais**: `<`, `>`, `<=`, `>=`
7. **Igualdade**: `==`, `!=`, `===`, `!==`
8. **Lógicos AND**: `&&`
9. **Lógicos OR**: `||`, `??`
10. **Ternário**: `? :`
11. **Atribuição**: `=`, `+=`, `-=`, etc.
12. **Vírgula**: `,`

### Exemplos de Precedência

```javascript
// Multiplicação antes de adição
let resultado1 = 2 + 3 * 4;        // 14 (não 20!)
// Equivale a: 2 + (3 * 4)

// Parênteses mudam a ordem
let resultado2 = (2 + 3) * 4;      // 20

// Lógica AND antes de OR
let resultado3 = true || false && false;  // true
// Equivale a: true || (false && false)

// Ternário tem baixa precedência
let resultado4 = 5 > 3 ? "sim" : "não";  // "sim"
```

### Boas Práticas

- **Use parênteses** quando a precedência não for óbvia
- **Não confie apenas na memória** - parênteses tornam o código mais legível
- **Considere a legibilidade** sobre a brevidade

---

## 🔄 Conversão de Tipos em Expressões

JavaScript realiza conversões automáticas de tipo em expressões (type coercion).

### Conversões Comuns

```javascript
// String + Number = String
"5" + 3;        // "53" (número convertido para string)

// Number - String = Number (se possível)
"5" - 3;        // 2 (string convertida para número)
"abc" - 3;      // NaN

// Boolean em operações numéricas
true + 1;       // 2 (true = 1)
false + 1;      // 1 (false = 0)

// Comparações com conversão
"5" == 5;       // true (conversão de tipo)
"5" === 5;      // false (sem conversão)
```

### Valores Falsy e Truthy

**Valores Falsy** (convertidos para `false`):
- `false`
- `0`
- `-0`
- `0n` (BigInt zero)
- `""` (string vazia)
- `null`
- `undefined`
- `NaN`

**Valores Truthy** (convertidos para `true`):
- Todos os outros valores, incluindo:
  - `"0"` (string)
  - `"false"` (string)
  - `[]` (array vazio)
  - `{}` (objeto vazio)
  - `function(){}` (função)

---

## 🎯 Resumo dos Operadores

### Operadores Aritméticos
- `+`, `-`, `*`, `/`, `%`, `**`
- `++`, `--`

### Operadores de Atribuição
- `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `**=`

### Operadores de Comparação
- `==`, `!=`, `===`, `!==`
- `>`, `<`, `>=`, `<=`

### Operadores Lógicos
- `&&`, `||`, `!`, `??`

### Operadores Especiais
- `? :` (ternário)
- `,` (vírgula)
- `typeof`, `delete`, `void`
- `+` (unário), `-` (unário)

---

## 📝 Próximos Passos

Agora que você entende expressões e operadores, você está pronto para:
- Combinar operadores em expressões complexas
- Entender como o JavaScript avalia expressões
- Escrever código mais eficiente e legível
- Aplicar esses conceitos em funções e estruturas de controle

Na próxima aula, você aprenderá sobre **Funções** - uma das partes mais importantes do JavaScript!

