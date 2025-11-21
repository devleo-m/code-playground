# Aula 2 - Exercícios e Reflexão: Variáveis em JavaScript

Bem-vindo aos exercícios práticos! Aqui você vai colocar em prática tudo que aprendeu sobre variáveis. Lembre-se: **a prática é essencial para fixar o conhecimento**.

---

## 📝 Exercício 1: Declaração de Variáveis

### Objetivo
Praticar a declaração de variáveis usando `let`, `const` e entender quando usar cada uma.

### Tarefa
Crie variáveis para armazenar as seguintes informações de um usuário:

1. Nome completo (não muda)
2. Idade (pode mudar)
3. Email (não muda)
4. Quantidade de login (aumenta a cada login)
5. Status de ativo (pode mudar entre true/false)
6. Lista de hobbies (pode adicionar/remover itens)

### Código Base
```javascript
// Escreva seu código aqui:

```

### Resposta Esperada
```javascript
// Sua resposta deve ser algo como:
const nomeCompleto = "João Silva";
let idade = 25;
const email = "joao@email.com";
let quantidadeDeLogin = 0;
let estaAtivo = true;
const hobbies = ["ler", "correr", "programar"];
```

### Verificação
- [ ] Usei `const` para valores que não mudam (nome, email)
- [ ] Usei `let` para valores que podem mudar (idade, quantidadeDeLogin, estaAtivo)
- [ ] Usei `const` para o array de hobbies (pode modificar conteúdo, mas não reatribuir)
- [ ] Usei nomes descritivos em camelCase

---

## 📝 Exercício 2: Escopo de Variáveis

### Objetivo
Entender como o escopo funciona com `var`, `let` e `const`.

### Tarefa
Analise o código abaixo e responda:

1. Quais variáveis serão impressas corretamente?
2. Quais vão gerar erro?
3. Por quê?

### Código
```javascript
var globalVar = "Sou global (var)";
let globalLet = "Sou global (let)";
const globalConst = "Sou global (const)";

function exemplo() {
    var funcaoVar = "Sou da função (var)";
    let funcaoLet = "Sou da função (let)";
    const funcaoConst = "Sou da função (const)";
    
    if (true) {
        var blocoVar = "Sou do bloco (var)";
        let blocoLet = "Sou do bloco (let)";
        const blocoConst = "Sou do bloco (const)";
        
        console.log(globalVar); // ?
        console.log(globalLet); // ?
        console.log(globalConst); // ?
        console.log(funcaoVar); // ?
        console.log(funcaoLet); // ?
        console.log(funcaoConst); // ?
        console.log(blocoVar); // ?
        console.log(blocoLet); // ?
        console.log(blocoConst); // ?
    }
    
    console.log(blocoVar); // ?
    console.log(blocoLet); // ?
    console.log(blocoConst); // ?
}

exemplo();

console.log(globalVar); // ?
console.log(globalLet); // ?
console.log(globalConst); // ?
console.log(funcaoVar); // ?
console.log(funcaoLet); // ?
console.log(funcaoConst); // ?
```

### Sua Resposta
Escreva aqui quais linhas vão funcionar e quais vão gerar erro, e explique o porquê:

```
Linha X: ✅ Funciona / ❌ Erro - Explicação...
Linha Y: ✅ Funciona / ❌ Erro - Explicação...
...
```

---

## 📝 Exercício 3: Hoisting e Temporal Dead Zone

### Objetivo
Entender o comportamento de hoisting com `var`, `let` e `const`.

### Tarefa
Analise os seguintes códigos e explique o que acontece em cada um:

### Código A
```javascript
console.log(x);
var x = 10;
console.log(x);
```

### Código B
```javascript
console.log(y);
let y = 20;
console.log(y);
```

### Código C
```javascript
console.log(z);
const z = 30;
console.log(z);
```

### Sua Resposta
Explique o comportamento de cada código:

```
Código A: 
- Linha 1: ?
- Linha 3: ?
- Explicação: ...

Código B:
- Linha 1: ?
- Linha 3: ?
- Explicação: ...

Código C:
- Linha 1: ?
- Linha 3: ?
- Explicação: ...
```

---

## 📝 Exercício 4: Nomenclatura e Boas Práticas

### Objetivo
Praticar a criação de nomes de variáveis seguindo boas práticas.

### Tarefa
Reescreva as seguintes variáveis com nomes melhores seguindo as boas práticas:

### Código Original (Ruim)
```javascript
var n = "João";
var a = 25;
var e = "joao@email.com";
var q = 10;
var f = true;
var l = ["ler", "correr"];
var o = { n: "Maria", i: 30 };
```

### Sua Resposta
Reescreva com nomes descritivos:

```javascript
// Sua resposta aqui:

```

### Critérios de Avaliação
- [ ] Nomes descritivos e claros
- [ ] Uso correto de camelCase
- [ ] Uso adequado de `let` e `const`
- [ ] Nomes booleanos começam com "é", "esta", "tem", "pode"
- [ ] Arrays no plural, objetos no singular

---

## 📝 Exercício 5: Const com Objetos e Arrays

### Objetivo
Entender que `const` não impede modificação de objetos e arrays.

### Tarefa
Complete o código abaixo demonstrando que:

1. Você pode modificar propriedades de um objeto declarado com `const`
2. Você pode modificar itens de um array declarado com `const`
3. Você NÃO pode reatribuir o objeto/array inteiro

### Código Base
```javascript
// Crie um objeto pessoa com const
const pessoa = {
    nome: "João",
    idade: 25
};

// 1. Modifique a propriedade nome
// Seu código aqui:

// 2. Adicione uma nova propriedade cidade
// Seu código aqui:

// 3. Tente reatribuir o objeto inteiro (comente o código que gera erro)
// Seu código aqui:

// Crie um array frutas com const
const frutas = ["maçã", "banana"];

// 4. Modifique o primeiro item
// Seu código aqui:

// 5. Adicione um novo item
// Seu código aqui:

// 6. Tente reatribuir o array inteiro (comente o código que gera erro)
// Seu código aqui:

console.log(pessoa);
console.log(frutas);
```

### Resposta Esperada
```javascript
// Sua resposta deve demonstrar:
const pessoa = {
    nome: "João",
    idade: 25
};

pessoa.nome = "João Silva"; // ✅ Modifica propriedade
pessoa.cidade = "São Paulo"; // ✅ Adiciona propriedade
// pessoa = { nome: "Novo" }; // ❌ Erro: não pode reatribuir

const frutas = ["maçã", "banana"];
frutas[0] = "laranja"; // ✅ Modifica item
frutas.push("uva"); // ✅ Adiciona item
// frutas = ["novo"]; // ❌ Erro: não pode reatribuir
```

---

## 📝 Exercício 6: Problema Clássico com var em Loops

### Objetivo
Entender o problema do vazamento de escopo com `var` em loops.

### Tarefa
Analise os dois códigos abaixo e explique a diferença:

### Código A (com var)
```javascript
for (var i = 0; i < 3; i++) {
    setTimeout(() => {
        console.log(i);
    }, 100);
}
```

### Código B (com let)
```javascript
for (let j = 0; j < 3; j++) {
    setTimeout(() => {
        console.log(j);
    }, 100);
}
```

### Perguntas
1. O que será impresso no Código A? Por quê?
2. O que será impresso no Código B? Por quê?
3. Qual é a solução correta e por quê?

### Sua Resposta
```
Código A:
- Saída: ?
- Explicação: ...

Código B:
- Saída: ?
- Explicação: ...

Solução correta: ?
Razão: ...
```

---

## 📝 Exercício 7: Criar um Sistema de Contador

### Objetivo
Criar um sistema simples de contador usando variáveis adequadas.

### Tarefa
Crie um sistema de contador que:

1. Tem um valor inicial (constante)
2. Tem um contador atual (pode aumentar)
3. Tem um limite máximo (constante)
4. Pode incrementar o contador
5. Pode verificar se atingiu o limite

### Código Base
```javascript
// Crie seu sistema de contador aqui:

// Exemplo de uso:
// incrementar(); // contador atual: 1
// incrementar(); // contador atual: 2
// incrementar(); // contador atual: 3
// verificarLimite(); // "Limite atingido!"
```

### Resposta Esperada
```javascript
// Exemplo de solução:
const VALOR_INICIAL = 0;
const LIMITE_MAXIMO = 10;
let contadorAtual = VALOR_INICIAL;

function incrementar() {
    if (contadorAtual < LIMITE_MAXIMO) {
        contadorAtual++;
        console.log(`Contador atual: ${contadorAtual}`);
    } else {
        console.log("Limite atingido!");
    }
}

function verificarLimite() {
    if (contadorAtual >= LIMITE_MAXIMO) {
        console.log("Limite atingido!");
    } else {
        console.log(`Ainda faltam ${LIMITE_MAXIMO - contadorAtual} incrementos`);
    }
}
```

### Critérios de Avaliação
- [ ] Usei `const` para valores que não mudam (VALOR_INICIAL, LIMITE_MAXIMO)
- [ ] Usei `let` para o contador que muda
- [ ] Nomes descritivos e claros
- [ ] Código funcional e testado

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por que `var` pode ser problemático?

**Pergunta:** 
Analise o código abaixo e explique por que `var` pode causar problemas:

```javascript
function exemplo() {
    for (var i = 0; i < 3; i++) {
        // alguma lógica
    }
    console.log(i); // Por que isso funciona?
    
    if (true) {
        var x = 10;
    }
    console.log(x); // Por que isso funciona?
}
```

**Sua Resposta:**
```
Explique aqui por que var pode ser problemático neste contexto:

1. ...
2. ...
3. ...
```

**Perguntas para pensar:**
- Qual seria o impacto se este código estivesse em uma aplicação com muitos usuários?
- Como isso poderia causar bugs difíceis de encontrar?
- Por que `let` e `const` resolvem esses problemas?

---

### Reflexão 2: Quando usar `const` vs `let`?

**Pergunta:**
Analise os seguintes cenários e decida se deve usar `const` ou `let`:

1. Uma variável que armazena o nome do usuário (não muda depois de definida)
2. Uma variável de contador em um loop
3. Um array que receberá novos itens ao longo do tempo
4. Um objeto de configuração que será modificado
5. Uma variável que armazena o resultado de um cálculo que pode mudar

**Sua Resposta:**
```
1. Nome do usuário: const / let - Por quê: ...
2. Contador em loop: const / let - Por quê: ...
3. Array que recebe itens: const / let - Por quê: ...
4. Objeto de configuração: const / let - Por quê: ...
5. Resultado de cálculo: const / let - Por quê: ...
```

**Perguntas para pensar:**
- Qual é a regra geral para decidir entre `const` e `let`?
- Por que usar `const` por padrão é uma boa prática?
- Quando é realmente necessário usar `let`?

---

### Reflexão 3: Impacto de Variáveis Globais

**Pergunta:**
Analise o código abaixo e pense sobre os problemas:

```javascript
// arquivo1.js
var contador = 0;

function incrementar() {
    contador++;
}

// arquivo2.js
var contador = 0; // Mesmo nome!

function decrementar() {
    contador--;
}

// O que acontece quando ambos os arquivos são carregados?
```

**Sua Resposta:**
```
Explique os problemas que podem ocorrer:

1. ...
2. ...
3. ...
```

**Perguntas para pensar:**
- O que acontece quando dois arquivos usam o mesmo nome de variável global?
- Como isso pode causar bugs difíceis de encontrar?
- Qual seria a melhor forma de organizar este código?
- Como módulos (ES6 modules) resolvem esse problema?

---

### Reflexão 4: Performance e Memória

**Pergunta:**
Considere os seguintes cenários:

1. Criar 1000 variáveis com `var` vs `let` vs `const`
2. Variáveis globais vs variáveis locais
3. Re-declarar variáveis com `var` múltiplas vezes

**Sua Resposta:**
```
Analise o impacto em cada cenário:

1. 1000 variáveis:
   - var: ...
   - let: ...
   - const: ...

2. Globais vs Locais:
   - Impacto na memória: ...
   - Impacto na performance: ...

3. Re-declaração:
   - Impacto: ...
```

**Perguntas para pensar:**
- Há diferença de performance entre `var`, `let` e `const`?
- Variáveis globais consomem mais memória que locais?
- Como o garbage collector (coletor de lixo) afeta variáveis de diferentes escopos?
- Qual é o impacto de ter muitas variáveis globais em uma aplicação grande?

---

### Reflexão 5: Edge Cases e Possíveis Erros

**Pergunta:**
Identifique possíveis problemas nestes códigos:

### Código A
```javascript
const usuario = {
    nome: "João",
    idade: 25
};

usuario = {
    nome: "Maria",
    idade: 30
};
```

### Código B
```javascript
let x = 10;
let x = 20;
```

### Código C
```javascript
console.log(y);
let y = 10;
```

### Código D
```javascript
function exemplo() {
    if (true) {
        var x = 10;
    }
    console.log(x);
    if (false) {
        var y = 20;
    }
    console.log(y);
}
```

**Sua Resposta:**
```
Código A:
- Problema: ...
- Solução: ...

Código B:
- Problema: ...
- Solução: ...

Código C:
- Problema: ...
- Solução: ...

Código D:
- Problema: ...
- Solução: ...
```

**Perguntas para pensar:**
- Quais são os erros mais comuns ao trabalhar com variáveis?
- Como você pode prevenir esses erros?
- Qual ferramenta ou técnica ajuda a identificar esses problemas antes da execução?

---

## ✅ Checklist de Aprendizado

Antes de prosseguir, verifique se você consegue:

- [ ] Declarar variáveis usando `let`, `const` e `var`
- [ ] Explicar a diferença entre `var`, `let` e `const`
- [ ] Entender o conceito de escopo (global, função, bloco)
- [ ] Explicar o que é hoisting e Temporal Dead Zone
- [ ] Criar nomes de variáveis seguindo boas práticas
- [ ] Entender que `const` não impede modificação de objetos/arrays
- [ ] Identificar problemas comuns com variáveis
- [ ] Escolher entre `const` e `let` adequadamente
- [ ] Explicar por que evitar `var` em código novo
- [ ] Resolver exercícios práticos sobre variáveis

---

## 🚀 Próximo Passo

Após completar todos os exercícios e reflexões, você estará pronto para a **Aula sobre Performance, Boas Práticas e Otimização**!

**Arquivo seguinte**: `04-performance-boas-praticas.md`

**Dica:** Não pule os exercícios! Eles são essenciais para fixar o conhecimento. Se tiver dúvidas, revise as aulas anteriores.

---

## 📝 Notas Pessoais

Use este espaço para anotar suas dúvidas, descobertas ou observações:

```
[Seu espaço para anotações]
```

