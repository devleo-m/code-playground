# Aula 2: Variáveis em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 1**, você aprendeu:
- ✅ O que é JavaScript e onde ele é utilizado
- ✅ A história e evolução do JavaScript
- ✅ Como executar código JavaScript (navegador, Node.js, console)
- ✅ Diferenças entre ambientes de execução

Agora vamos aprender como **armazenar e trabalhar com informações** usando variáveis!

---

## 🎯 O que são Variáveis?

Na maioria das vezes, uma aplicação JavaScript precisa trabalhar com informações. Para armazenar e representar essas informações no código JavaScript, usamos **variáveis**.

**Definição:** Uma variável é um **container para um valor**. Pense nela como uma caixa com um rótulo (nome) que guarda algo dentro.

### Analogia Rápida

Imagine variáveis como **etiquetas em caixas de armazenamento**:
- A **etiqueta** (nome da variável) identifica o que está dentro
- O **conteúdo** (valor) pode ser alterado quando necessário
- Você pode **reutilizar** a mesma caixa para diferentes valores

---

## 🔑 Palavras-chave para Declarar Variáveis

Em JavaScript, existem três palavras-chave principais para declarar variáveis:

1. **`var`** - Declaração tradicional (ES5)
2. **`let`** - Declaração moderna com escopo de bloco (ES6+)
3. **`const`** - Declaração de constante (ES6+)

Vamos explorar cada uma em detalhes.

---

## 📦 A Palavra-chave `var`

### Definição

A declaração `var` cria uma variável com **escopo de função** ou **escopo global**, opcionalmente inicializando-a com um valor.

### Características do `var`

1. **Escopo de Função ou Global**
   - Se declarada dentro de uma função, é acessível apenas dentro dessa função
   - Se declarada fora de qualquer função, é global (acessível em todo o código)

2. **Hoisting**
   - A declaração é "elevada" (hoisted) para o topo do escopo
   - Pode ser usada antes de ser declarada (retorna `undefined`)

3. **Re-declaração Permitida**
   - Pode ser declarada múltiplas vezes no mesmo escopo

### Sintaxe

```javascript
var nomeDaVariavel;
var nomeDaVariavel = valor;
```

### Exemplos

```javascript
// Declaração sem valor inicial
var idade;
console.log(idade); // undefined

// Declaração com valor inicial
var nome = "João";
console.log(nome); // "João"

// Reatribuição
var cor = "azul";
cor = "vermelho";
console.log(cor); // "vermelho"

// Re-declaração (permitida com var)
var numero = 10;
var numero = 20; // Não gera erro
console.log(numero); // 20

// Escopo de função
function exemplo() {
    var local = "Esta variável só existe aqui";
    console.log(local);
}
exemplo(); // "Esta variável só existe aqui"
// console.log(local); // Erro: local is not defined

// Hoisting com var
console.log(x); // undefined (não gera erro)
var x = 5;
console.log(x); // 5
```

### Comportamento de Hoisting

```javascript
// O código acima é interpretado como:
var x; // Declaração movida para o topo
console.log(x); // undefined
x = 5; // Atribuição permanece no lugar
console.log(x); // 5
```

### ⚠️ Problemas com `var`

1. **Vazamento de Escopo**
   ```javascript
   for (var i = 0; i < 3; i++) {
       // i é acessível aqui
   }
   console.log(i); // 3 (i vazou do bloco for)
   ```

2. **Re-declaração Silenciosa**
   ```javascript
   var x = 10;
   var x = 20; // Não gera erro, pode causar bugs
   ```

3. **Hoisting Confuso**
   ```javascript
   console.log(y); // undefined (não gera erro, mas é confuso)
   var y = 10;
   ```

**Recomendação Moderna:** Evite usar `var` em código novo. Use `let` ou `const` em vez disso.

---

## 🔒 A Palavra-chave `let`

### Definição

A declaração `let` cria uma variável com **escopo de bloco**, opcionalmente inicializando-a com um valor.

### Características do `let`

1. **Escopo de Bloco**
   - Acessível apenas dentro do bloco onde foi declarada
   - Blocos são delimitados por `{}` (chaves)

2. **Hoisting com Temporal Dead Zone (TDZ)**
   - A declaração é hoisted, mas não pode ser acessada antes da declaração
   - Tentar acessar antes da declaração gera um erro

3. **Re-declaração NÃO Permitida**
   - Não pode ser declarada novamente no mesmo escopo

4. **Reatribuição Permitida**
   - O valor pode ser alterado após a declaração

### Sintaxe

```javascript
let nomeDaVariavel;
let nomeDaVariavel = valor;
```

### Exemplos

```javascript
// Declaração sem valor inicial
let idade;
console.log(idade); // undefined

// Declaração com valor inicial
let nome = "Maria";
console.log(nome); // "Maria"

// Reatribuição (permitida)
let cor = "azul";
cor = "vermelho";
console.log(cor); // "vermelho"

// Re-declaração (NÃO permitida)
let numero = 10;
// let numero = 20; // Erro: Identifier 'numero' has already been declared

// Escopo de bloco
{
    let bloco = "Esta variável só existe neste bloco";
    console.log(bloco); // "Esta variável só existe neste bloco"
}
// console.log(bloco); // Erro: bloco is not defined

// Diferentes blocos podem ter variáveis com o mesmo nome
let x = 1;
{
    let x = 2;
    console.log(x); // 2 (variável do bloco interno)
}
console.log(x); // 1 (variável do bloco externo)

// Temporal Dead Zone (TDZ)
// console.log(y); // Erro: Cannot access 'y' before initialization
let y = 10;
console.log(y); // 10

// Uso em loops
for (let i = 0; i < 3; i++) {
    console.log(i); // 0, 1, 2
}
// console.log(i); // Erro: i is not defined (não vazou do bloco)
```

### Vantagens do `let`

1. **Escopo de Bloco Previsível**
   ```javascript
   if (true) {
       let mensagem = "Dentro do if";
   }
   // mensagem não existe aqui - comportamento previsível
   ```

2. **Previne Re-declaração Acidental**
   ```javascript
   let x = 10;
   // let x = 20; // Erro imediato - previne bugs
   ```

3. **Melhor para Loops**
   ```javascript
   for (let i = 0; i < 3; i++) {
       setTimeout(() => console.log(i), 100); // 0, 1, 2 (correto)
   }
   ```

---

## 🔐 A Palavra-chave `const`

### Definição

A declaração `const` cria uma **constante** com escopo de bloco. O valor de uma constante **não pode ser alterado através de reatribuição** (usando o operador de atribuição `=`), e **não pode ser re-declarada**.

### Características do `const`

1. **Escopo de Bloco**
   - Mesmo comportamento de escopo que `let`

2. **Valor Imutável (para tipos primitivos)**
   - Não pode ser reatribuído após a declaração
   - Deve ser inicializada no momento da declaração

3. **Objetos e Arrays são Mutáveis**
   - A referência não pode mudar, mas o conteúdo pode
   - Propriedades de objetos podem ser modificadas
   - Itens de arrays podem ser adicionados/removidos/modificados

4. **Re-declaração NÃO Permitida**
   - Mesmo comportamento que `let`

5. **Temporal Dead Zone (TDZ)**
   - Mesmo comportamento que `let`

### Sintaxe

```javascript
const nomeDaConstante = valor; // OBRIGATÓRIO inicializar
```

### Exemplos

#### Tipos Primitivos (Imutáveis)

```javascript
// Declaração OBRIGATÓRIA com valor inicial
const PI = 3.14159;
console.log(PI); // 3.14159

// Reatribuição NÃO permitida
// PI = 3.14; // Erro: Assignment to constant variable

// Re-declaração NÃO permitida
// const PI = 3.14; // Erro: Identifier 'PI' has already been declared

// Outros exemplos
const nome = "Ana";
const idade = 25;
const ativo = true;

// Tentar reatribuir gera erro
// nome = "João"; // Erro
// idade = 26; // Erro
// ativo = false; // Erro
```

#### Objetos (Mutáveis)

```javascript
// A referência não pode mudar, mas o conteúdo pode
const pessoa = {
    nome: "Carlos",
    idade: 30
};

// Modificar propriedades é permitido
pessoa.nome = "Carlos Silva";
pessoa.idade = 31;
pessoa.cidade = "São Paulo";
console.log(pessoa); // { nome: "Carlos Silva", idade: 31, cidade: "São Paulo" }

// Reatribuir o objeto inteiro NÃO é permitido
// pessoa = { nome: "Novo" }; // Erro: Assignment to constant variable

// Deletar propriedades é permitido
delete pessoa.cidade;
console.log(pessoa); // { nome: "Carlos Silva", idade: 31 }
```

#### Arrays (Mutáveis)

```javascript
// A referência não pode mudar, mas o conteúdo pode
const frutas = ["maçã", "banana"];

// Modificar itens é permitido
frutas[0] = "laranja";
console.log(frutas); // ["laranja", "banana"]

// Adicionar itens é permitido
frutas.push("uva");
console.log(frutas); // ["laranja", "banana", "uva"]

// Remover itens é permitido
frutas.pop();
console.log(frutas); // ["laranja", "banana"]

// Reatribuir o array inteiro NÃO é permitido
// frutas = ["novo"]; // Erro: Assignment to constant variable
```

### Quando Usar `const`

**Use `const` por padrão** quando:
- O valor não precisa ser reatribuído
- Você quer garantir que a referência não mude
- Trabalhando com objetos e arrays que serão modificados, mas não reatribuídos

**Use `let` quando:**
- O valor precisa ser reatribuído
- A variável será usada em loops e precisa mudar

---

## 📝 Declaração de Variáveis

### Como Declarar Variáveis

Para usar variáveis em JavaScript, primeiro precisamos **criá-las**, ou seja, **declarar uma variável**. Para declarar variáveis, usamos uma das palavras-chave: `var`, `let` ou `const`.

### Processo de Declaração

1. **Declaração**: Criar a variável
2. **Inicialização**: Atribuir um valor inicial (opcional para `var` e `let`, obrigatório para `const`)
3. **Uso**: Utilizar a variável no código

### Exemplos de Declaração

```javascript
// 1. Declaração sem inicialização (apenas var e let)
var a;
let b;
// const c; // Erro: Missing initializer in const declaration

// 2. Declaração com inicialização
var x = 10;
let y = 20;
const z = 30;

// 3. Múltiplas declarações
var nome = "João", idade = 25, cidade = "SP";
let a = 1, b = 2, c = 3;
const PI = 3.14, E = 2.71;

// 4. Declaração e uso posterior
let resultado;
resultado = 10 + 20;
console.log(resultado); // 30
```

### Boas Práticas de Declaração

```javascript
// ✅ BOM: Declaração clara e inicializada
let contador = 0;
const nomeUsuario = "João";

// ✅ BOM: Uma declaração por linha (mais legível)
let nome = "Maria";
let idade = 30;
let cidade = "Rio";

// ⚠️ EVITAR: Múltiplas declarações na mesma linha (menos legível)
let nome = "Maria", idade = 30, cidade = "Rio";

// ✅ BOM: Nomes descritivos
let quantidadeDeProdutos = 10;
let precoTotal = 99.90;

// ❌ EVITAR: Nomes genéricos
let x = 10;
let y = 20;
```

---

## ⬆️ Hoisting (Elevação)

### O que é Hoisting?

**Hoisting** (elevação) refere-se ao processo pelo qual o interpretador JavaScript **aparentemente move** as declarações de funções, variáveis ou classes para o topo de seu escopo, **antes da execução do código**.

### Como Funciona

O JavaScript processa o código em duas fases:
1. **Fase de Compilação**: Declarações são "elevadas" (hoisted)
2. **Fase de Execução**: Código é executado linha por linha

### Hoisting com `var`

```javascript
// Código escrito:
console.log(x); // undefined (não gera erro)
var x = 5;
console.log(x); // 5

// Como o JavaScript interpreta:
var x; // Declaração movida para o topo
console.log(x); // undefined
x = 5; // Atribuição permanece no lugar
console.log(x); // 5
```

### Hoisting com `let` e `const` (Temporal Dead Zone)

```javascript
// Código escrito:
console.log(y); // Erro: Cannot access 'y' before initialization
let y = 10;

// Como o JavaScript interpreta:
// let y; // Declaração é hoisted, mas não inicializada
console.log(y); // Erro: TDZ (Temporal Dead Zone)
y = 10; // Inicialização

// O mesmo acontece com const
console.log(z); // Erro: Cannot access 'z' before initialization
const z = 20;
```

### Temporal Dead Zone (TDZ)

A **Temporal Dead Zone** é o período entre o início do escopo e a declaração da variável onde a variável não pode ser acessada.

```javascript
// TDZ começa aqui (início do escopo)
console.log(a); // Erro: TDZ
let a = 10; // TDZ termina aqui
console.log(a); // 10 (OK)
```

### Exemplos Práticos de Hoisting

```javascript
// Exemplo 1: var
function exemplo1() {
    console.log(nome); // undefined
    var nome = "João";
    console.log(nome); // "João"
}
exemplo1();

// Exemplo 2: let (TDZ)
function exemplo2() {
    // console.log(idade); // Erro: Cannot access 'idade' before initialization
    let idade = 25;
    console.log(idade); // 25
}
exemplo2();

// Exemplo 3: const (TDZ)
function exemplo3() {
    // console.log(PI); // Erro: Cannot access 'PI' before initialization
    const PI = 3.14159;
    console.log(PI); // 3.14159
}
exemplo3();
```

### ⚠️ Por que Hoisting Pode Ser Problemático?

```javascript
// Código confuso devido ao hoisting
function exemplo() {
    console.log(x); // undefined (não o valor esperado)
    var x = 10;
}

// Melhor prática: sempre declarar no topo
function exemploMelhor() {
    var x = 10; // Declaração no topo
    console.log(x); // 10 (comportamento esperado)
}
```

**Dica:** Sempre declare variáveis no topo do escopo para evitar confusão com hoisting.

---

## 📛 Regras de Nomenclatura

### Por que Nomes Importam?

Um nome de variável deve **identificar com precisão** sua variável. Quando você cria bons nomes de variáveis, seu código JavaScript se torna:
- ✅ Mais fácil de entender
- ✅ Mais fácil de trabalhar
- ✅ Mais fácil de manter
- ✅ Menos propenso a erros

### Regras Obrigatórias

JavaScript tem regras específicas para nomes de variáveis:

1. **Caracteres Permitidos**
   - Letras (a-z, A-Z)
   - Dígitos (0-9)
   - Underscore (_)
   - Símbolo de dólar ($)

2. **Primeiro Caractere**
   - **NÃO pode** começar com um dígito
   - **Pode** começar com letra, underscore ou dólar

3. **Case Sensitive**
   - `nome` e `Nome` são variáveis diferentes
   - `idade` e `Idade` são variáveis diferentes

4. **Palavras Reservadas**
   - Não pode usar palavras reservadas do JavaScript como nome de variável

### Exemplos de Nomes Válidos e Inválidos

```javascript
// ✅ VÁLIDOS
let nome = "João";
let nomeCompleto = "João Silva";
let _privado = "valor";
let $elemento = document.getElementById("id");
let nome123 = "teste";
let nome_completo = "João Silva";

// ❌ INVÁLIDOS
// let 123nome = "teste"; // Erro: Não pode começar com dígito
// let nome completo = "João"; // Erro: Não pode ter espaços
// let nome-completo = "João"; // Erro: Hífen não é permitido
// let var = "teste"; // Erro: var é palavra reservada
// let function = "teste"; // Erro: function é palavra reservada
// let if = "teste"; // Erro: if é palavra reservada
```

### Palavras Reservadas do JavaScript

Algumas palavras não podem ser usadas como nomes de variáveis:

```javascript
// Palavras reservadas principais:
// break, case, catch, class, const, continue, debugger, default,
// delete, do, else, export, extends, finally, for, function, if,
// import, in, instanceof, new, return, super, switch, this, throw,
// try, typeof, var, void, while, with, yield

// Exemplos de erros:
// let class = "teste"; // Erro
// let function = "teste"; // Erro
// let return = "teste"; // Erro
```

### Convenções de Nomenclatura

#### 1. camelCase (Recomendado para JavaScript)

```javascript
// Primeira palavra em minúscula, palavras seguintes com primeira letra maiúscula
let nomeCompleto = "João Silva";
let idadeDoUsuario = 25;
let quantidadeDeProdutos = 10;
let precoTotal = 99.90;
```

#### 2. snake_case (Menos comum em JavaScript)

```javascript
// Palavras separadas por underscore
let nome_completo = "João Silva";
let idade_do_usuario = 25;
```

#### 3. UPPER_SNAKE_CASE (Para constantes)

```javascript
// Geralmente usado para constantes
const PI = 3.14159;
const MAX_TENTATIVAS = 3;
const URL_BASE = "https://api.exemplo.com";
```

#### 4. PascalCase (Para classes, não variáveis)

```javascript
// Usado para classes, não para variáveis
class Usuario {
    // ...
}

// Não use para variáveis:
// let NomeCompleto = "João"; // Não é convenção padrão
```

### Boas Práticas de Nomenclatura

```javascript
// ✅ BOM: Nomes descritivos e claros
let quantidadeDeProdutos = 10;
let precoTotal = 99.90;
let nomeDoUsuario = "João";
let estaAtivo = true;

// ❌ EVITAR: Nomes genéricos ou abreviados
let qtd = 10; // O que é qtd?
let pt = 99.90; // O que é pt?
let n = "João"; // O que é n?
let flag = true; // O que a flag representa?

// ✅ BOM: Nomes que indicam o tipo (quando útil)
let listaDeFrutas = ["maçã", "banana"];
let objetoUsuario = { nome: "João" };
let numeroDeTentativas = 3;

// ✅ BOM: Nomes booleanos começam com "é", "esta", "tem", "pode"
let estaAtivo = true;
let temPermissao = false;
let podeEditar = true;
let eValido = false;

// ✅ BOM: Nomes de arrays no plural
let frutas = ["maçã", "banana"];
let usuarios = [];
let produtos = [];

// ✅ BOM: Nomes de objetos no singular
let usuario = { nome: "João" };
let produto = { nome: "Notebook" };
```

---

## 🌍 Escopo de Variáveis

### O que é Escopo?

**Escopo de variável** determina a **acessibilidade** de variáveis em diferentes partes do seu código. O escopo define onde uma variável pode ser acessada e modificada.

### Tipos de Escopo em JavaScript

JavaScript tem três tipos principais de escopo:

1. **Global Scope** (Escopo Global)
2. **Function Scope** (Escopo de Função)
3. **Block Scope** (Escopo de Bloco)

Além disso, existe:
4. **Module Scope** (Escopo de Módulo)

---

## 🌐 Global Scope (Escopo Global)

### Definição

Variáveis declaradas **globalmente** (fora de qualquer função ou bloco) têm **escopo global**. Elas podem ser acessadas e modificadas de **qualquer lugar** no seu código.

### Características

- Acessível em todo o código
- Pode ser modificada de qualquer lugar
- Permanece disponível durante toda a execução do programa
- Pode causar conflitos de nomes

### Exemplos

```javascript
// Variáveis globais
var globalVar = "Sou global (var)";
let globalLet = "Sou global (let)";
const globalConst = "Sou global (const)";

function funcao1() {
    console.log(globalVar); // "Sou global (var)"
    console.log(globalLet); // "Sou global (let)"
    console.log(globalConst); // "Sou global (const)"
}

function funcao2() {
    globalVar = "Modificado em funcao2";
    globalLet = "Modificado em funcao2";
    // globalConst = "Erro"; // Não pode reatribuir const
    console.log(globalVar); // "Modificado em funcao2"
}

funcao1();
funcao2();
console.log(globalVar); // "Modificado em funcao2"
```

### Variáveis Globais Implícitas

```javascript
// ⚠️ ATENÇÃO: Se você atribuir um valor a uma variável não declarada,
// ela se torna automaticamente GLOBAL (mesmo dentro de funções)

function exemplo() {
    // Variável não declarada - torna-se global automaticamente
    variavelGlobal = "Sou global automaticamente";
}

exemplo();
console.log(variavelGlobal); // "Sou global automaticamente" (acessível globalmente)

// ⚠️ Isso é um PROBLEMA! Sempre declare variáveis.
// Use 'use strict' para evitar isso:
'use strict';
function exemplo2() {
    // variavelGlobal2 = "Erro"; // Erro em strict mode
    let variavelLocal = "Sou local";
}
```

### ⚠️ Problemas com Variáveis Globais

```javascript
// Problema 1: Conflito de nomes
var contador = 0;

function incrementar() {
    contador++; // Qual contador? Pode haver conflito
}

// Problema 2: Modificação acidental
var nome = "João";

function processar() {
    nome = "Maria"; // Modificou a variável global acidentalmente
}

// Problema 3: Poluição do escopo global
// Muitas variáveis globais tornam o código difícil de manter
```

**Dica:** Evite variáveis globais sempre que possível. Use `let` e `const` com escopo apropriado.

---

## 🏠 Function Scope (Escopo de Função)

### Definição

Quando uma variável é declarada **dentro de uma função**, ela é acessível **apenas dentro dessa função** e não pode ser usada fora dela.

### Características

- Acessível apenas dentro da função onde foi declarada
- Não acessível fora da função
- Cada chamada da função cria um novo escopo
- Variáveis `var` têm escopo de função

### Exemplos

```javascript
function exemplo() {
    var funcaoVar = "Sou local à função (var)";
    let funcaoLet = "Sou local à função (let)";
    const funcaoConst = "Sou local à função (const)";
    
    console.log(funcaoVar); // "Sou local à função (var)"
    console.log(funcaoLet); // "Sou local à função (let)"
    console.log(funcaoConst); // "Sou local à função (const)"
}

exemplo();
// console.log(funcaoVar); // Erro: funcaoVar is not defined
// console.log(funcaoLet); // Erro: funcaoLet is not defined
// console.log(funcaoConst); // Erro: funcaoConst is not defined
```

### Escopo de Função com `var`

```javascript
function exemploVar() {
    if (true) {
        var x = 10; // var tem escopo de função, não de bloco
    }
    console.log(x); // 10 (acessível porque var tem escopo de função)
}

exemploVar();
// console.log(x); // Erro: x is not defined (fora da função)
```

### Escopo de Função com `let` e `const`

```javascript
function exemploLet() {
    if (true) {
        let y = 20; // let tem escopo de bloco
        const z = 30; // const tem escopo de bloco
    }
    // console.log(y); // Erro: y is not defined (fora do bloco)
    // console.log(z); // Erro: z is not defined (fora do bloco)
}

exemploLet();
```

### Variáveis Locais vs Globais

```javascript
var global = "Sou global";

function exemplo() {
    var local = "Sou local";
    console.log(global); // "Sou global" (acessível)
    console.log(local); // "Sou local" (acessível)
    
    // Variável local com mesmo nome "esconde" a global
    var global = "Sou local, escondendo a global";
    console.log(global); // "Sou local, escondendo a global"
}

exemplo();
console.log(global); // "Sou global" (não foi modificada)
// console.log(local); // Erro: local is not defined
```

---

## 📦 Block Scope (Escopo de Bloco)

### Definição

Este escopo **restringe a variável** que é declarada dentro de um **bloco específico** (delimitado por `{}`), impedindo o acesso de fora do bloco. As palavras-chave `let` e `const` facilitam variáveis com escopo de bloco.

### Características

- Acessível apenas dentro do bloco onde foi declarada
- Blocos são delimitados por chaves `{}`
- `let` e `const` têm escopo de bloco
- `var` **NÃO** tem escopo de bloco

### Exemplos Básicos

```javascript
{
    let blocoLet = "Sou do bloco (let)";
    const blocoConst = "Sou do bloco (const)";
    var blocoVar = "Sou do bloco (var)";
    
    console.log(blocoLet); // "Sou do bloco (let)"
    console.log(blocoConst); // "Sou do bloco (const)"
    console.log(blocoVar); // "Sou do bloco (var)"
}

// console.log(blocoLet); // Erro: blocoLet is not defined
// console.log(blocoConst); // Erro: blocoConst is not defined
console.log(blocoVar); // "Sou do bloco (var)" (var vazou do bloco)
```

### Block Scope em Estruturas Condicionais

```javascript
if (true) {
    let condicionalLet = "Dentro do if (let)";
    const condicionalConst = "Dentro do if (const)";
    var condicionalVar = "Dentro do if (var)";
    
    console.log(condicionalLet); // OK
    console.log(condicionalConst); // OK
    console.log(condicionalVar); // OK
}

// console.log(condicionalLet); // Erro: não acessível fora do bloco
// console.log(condicionalConst); // Erro: não acessível fora do bloco
console.log(condicionalVar); // "Dentro do if (var)" (var vazou)
```

### Block Scope em Loops

```javascript
// Loop com let (cada iteração tem seu próprio escopo)
for (let i = 0; i < 3; i++) {
    console.log(i); // 0, 1, 2
}
// console.log(i); // Erro: i is not defined

// Loop com var (vaza do bloco)
for (var j = 0; j < 3; j++) {
    console.log(j); // 0, 1, 2
}
console.log(j); // 3 (var vazou do bloco)

// Problema clássico com var em loops e callbacks
for (var k = 0; k < 3; k++) {
    setTimeout(() => {
        console.log(k); // 3, 3, 3 (todos imprimem 3)
    }, 100);
}

// Solução com let
for (let l = 0; l < 3; l++) {
    setTimeout(() => {
        console.log(l); // 0, 1, 2 (correto)
    }, 100);
}
```

### Blocos Aninhados

```javascript
let externo = "Escopo externo";

{
    let medio = "Escopo médio";
    console.log(externo); // "Escopo externo" (acessível)
    
    {
        let interno = "Escopo interno";
        console.log(externo); // "Escopo externo" (acessível)
        console.log(medio); // "Escopo médio" (acessível)
        console.log(interno); // "Escopo interno" (acessível)
    }
    
    // console.log(interno); // Erro: interno não acessível aqui
    console.log(medio); // "Escopo médio" (acessível)
}

// console.log(medio); // Erro: medio não acessível aqui
console.log(externo); // "Escopo externo" (acessível)
```

### Variáveis com Mesmo Nome em Diferentes Escopos

```javascript
let nome = "Global";

{
    let nome = "Bloco 1";
    console.log(nome); // "Bloco 1" (esconde a variável global)
    
    {
        let nome = "Bloco 2";
        console.log(nome); // "Bloco 2" (esconde a variável do bloco 1)
    }
    
    console.log(nome); // "Bloco 1" (ainda acessível)
}

console.log(nome); // "Global" (variável global não foi modificada)
```

---

## 🔄 Comparação de Escopos

### Tabela Comparativa

| Característica | `var` | `let` | `const` |
|----------------|-------|-------|---------|
| **Escopo Global** | ✅ Sim | ✅ Sim | ✅ Sim |
| **Escopo de Função** | ✅ Sim | ✅ Sim | ✅ Sim |
| **Escopo de Bloco** | ❌ Não | ✅ Sim | ✅ Sim |
| **Hoisting** | ✅ Sim (undefined) | ✅ Sim (TDZ) | ✅ Sim (TDZ) |
| **Re-declaração** | ✅ Permitida | ❌ Não permitida | ❌ Não permitida |
| **Reatribuição** | ✅ Permitida | ✅ Permitida | ❌ Não permitida |
| **Inicialização Obrigatória** | ❌ Não | ❌ Não | ✅ Sim |

### Exemplo Comparativo Completo

```javascript
// === GLOBAL SCOPE ===
var globalVar = "Global (var)";
let globalLet = "Global (let)";
const globalConst = "Global (const)";

function exemplo() {
    // === FUNCTION SCOPE ===
    var funcaoVar = "Função (var)";
    let funcaoLet = "Função (let)";
    const funcaoConst = "Função (const)";
    
    // Acessa variáveis globais
    console.log(globalVar); // "Global (var)"
    console.log(globalLet); // "Global (let)"
    console.log(globalConst); // "Global (const)"
    
    if (true) {
        // === BLOCK SCOPE ===
        var blocoVar = "Bloco (var)";
        let blocoLet = "Bloco (let)";
        const blocoConst = "Bloco (const)";
        
        // Acessa variáveis da função
        console.log(funcaoVar); // "Função (var)"
        console.log(funcaoLet); // "Função (let)"
        console.log(funcaoConst); // "Função (const)"
        
        // Acessa variáveis globais
        console.log(globalVar); // "Global (var)"
    }
    
    // var vazou do bloco
    console.log(blocoVar); // "Bloco (var)" (acessível)
    // console.log(blocoLet); // Erro: não acessível
    // console.log(blocoConst); // Erro: não acessível
}

exemplo();
// Variáveis da função não são acessíveis aqui
// console.log(funcaoVar); // Erro
```

---

## 📚 Resumo

Nesta aula você aprendeu:

- ✅ **Variáveis** são containers para valores
- ✅ Três palavras-chave: `var`, `let`, `const`
- ✅ **`var`**: Escopo de função/global, hoisting, permite re-declaração
- ✅ **`let`**: Escopo de bloco, TDZ, não permite re-declaração, permite reatribuição
- ✅ **`const`**: Escopo de bloco, TDZ, não permite re-declaração, não permite reatribuição (mas objetos/arrays são mutáveis)
- ✅ **Hoisting**: Declarações são movidas para o topo do escopo
- ✅ **Regras de nomenclatura**: camelCase, não pode começar com dígito, case sensitive
- ✅ **Escopos**: Global, Function, Block
- ✅ **Block Scope**: Restringe variáveis ao bloco (let/const)
- ✅ **Function Scope**: Restringe variáveis à função (var tem function scope)
- ✅ **Global Scope**: Acessível em todo o código

---

## 🚀 Próximo Passo

Agora que você entendeu as variáveis tecnicamente, está pronto para a **Aula Simplificada**, onde vamos revisar esses conceitos com analogias e exemplos do dia a dia.

**Arquivo seguinte**: `02-aula-simplificada.md`





