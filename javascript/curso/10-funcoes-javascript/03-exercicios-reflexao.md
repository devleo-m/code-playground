# Aula 10 - Exercícios e Reflexão: Funções

Bem-vindo aos exercícios práticos! Aqui você vai colocar em prática tudo que aprendeu sobre funções. Lembre-se: a prática é essencial para dominar qualquer conceito de programação.

---

## 📝 Exercício 1: Criando Sua Primeira Função

**Objetivo:** Criar uma função que calcula a área de um retângulo.

**Instruções:**
1. Crie uma função chamada `calcularAreaRetangulo` que recebe dois parâmetros: `largura` e `altura`
2. A função deve retornar a área do retângulo (largura × altura)
3. Teste a função com diferentes valores

**Código Base:**
```javascript
// Escreva sua função aqui


// Teste sua função aqui
console.log(calcularAreaRetangulo(5, 3));  // Deve retornar 15
console.log(calcularAreaRetangulo(10, 4)); // Deve retornar 40
```

**Desafio Extra:** Crie uma versão usando arrow function também!

---

## 📝 Exercício 2: Função com Parâmetros Padrão

**Objetivo:** Criar uma função de saudação com valores padrão.

**Instruções:**
1. Crie uma função chamada `saudar` que recebe dois parâmetros:
   - `nome` (padrão: "Visitante")
   - `hora` (padrão: 12)
2. A função deve retornar uma saudação diferente baseada na hora:
   - 6-11: "Bom dia, [nome]!"
   - 12-17: "Boa tarde, [nome]!"
   - 18-23 ou 0-5: "Boa noite, [nome]!"
3. Teste a função com e sem parâmetros

**Código Base:**
```javascript
// Escreva sua função aqui


// Teste sua função aqui
console.log(saudar("Maria", 9));    // "Bom dia, Maria!"
console.log(saudar("João", 15));    // "Boa tarde, João!"
console.log(saudar("Ana", 20));     // "Boa noite, Ana!"
console.log(saudar());               // Deve usar valores padrão
```

---

## 📝 Exercício 3: Função com Rest Parameters

**Objetivo:** Criar uma função que calcula a média de vários números.

**Instruções:**
1. Crie uma função chamada `calcularMedia` que aceita um número indefinido de argumentos usando rest parameters
2. A função deve retornar a média de todos os números passados
3. Se nenhum número for passado, retorne 0
4. Teste com diferentes quantidades de números

**Código Base:**
```javascript
// Escreva sua função aqui


// Teste sua função aqui
console.log(calcularMedia(10, 20, 30));           // Deve retornar 20
console.log(calcularMedia(5, 10, 15, 20, 25));    // Deve retornar 15
console.log(calcularMedia(100));                  // Deve retornar 100
console.log(calcularMedia());                     // Deve retornar 0
```

**Desafio Extra:** Modifique a função para ignorar valores que não são números!

---

## 📝 Exercício 4: Função Recursiva

**Objetivo:** Criar uma função recursiva que calcula o fatorial de um número.

**Instruções:**
1. Crie uma função chamada `fatorial` que recebe um número `n`
2. A função deve usar recursão para calcular o fatorial
3. Lembre-se: 
   - Caso base: fatorial de 0 ou 1 é 1
   - Caso recursivo: fatorial(n) = n × fatorial(n - 1)
4. Teste com diferentes valores

**Código Base:**
```javascript
// Escreva sua função aqui


// Teste sua função aqui
console.log(fatorial(0));  // Deve retornar 1
console.log(fatorial(1));  // Deve retornar 1
console.log(fatorial(5));  // Deve retornar 120 (5! = 5 × 4 × 3 × 2 × 1)
console.log(fatorial(7));  // Deve retornar 5040
```

**Desafio Extra:** Adicione validação para números negativos!

---

## 📝 Exercício 5: Análise de Código

**Objetivo:** Analisar e corrigir o código abaixo.

**Instruções:**
1. Leia o código abaixo cuidadosamente
2. Identifique os erros (sintáticos, lógicos, ou de boas práticas)
3. Corrija os erros
4. Explique o que estava errado

**Código com Problemas:**
```javascript
function calcularPreco(preco, desconto) {
  let precoFinal = preco - desconto
  return precoFinal
  console.log("Preço calculado")
}

function somar(...numeros) {
  let total = 0
  for (let i = 0; i <= numeros.length; i++) {
    total = total + numeros[i]
  }
  return total
}

function verificarIdade(idade) {
  if (idade >= 18) {
    return "Maior de idade"
  } else if (idade < 18) {
    return "Menor de idade"
  }
}

let resultado = calcularPreco(100, 20)
console.log(resultado)
```

**Sua Análise:**
```
Erros encontrados:
1. 
2. 
3. 

Código corrigido:
[Escreva o código corrigido aqui]
```

---

## 📝 Exercício 6: Criando um Sistema Completo

**Objetivo:** Criar um sistema de gerenciamento de notas usando múltiplas funções.

**Instruções:**
1. Crie uma função `calcularMedia` que recebe um array de notas e retorna a média
2. Crie uma função `verificarAprovacao` que recebe uma média e retorna:
   - "Aprovado" se média >= 7
   - "Recuperação" se média >= 5 e < 7
   - "Reprovado" se média < 5
3. Crie uma função `processarAluno` que recebe o nome do aluno e um array de notas, e retorna um objeto com:
   - nome
   - notas
   - media
   - status
4. Teste o sistema com diferentes alunos

**Código Base:**
```javascript
// Escreva suas funções aqui




// Teste seu sistema aqui
let aluno1 = processarAluno("Maria", [8, 7, 9, 6]);
console.log(aluno1);
// Deve retornar algo como:
// {
//   nome: "Maria",
//   notas: [8, 7, 9, 6],
//   media: 7.5,
//   status: "Aprovado"
// }

let aluno2 = processarAluno("João", [4, 5, 3, 6]);
console.log(aluno2);
```

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Eficiência e Performance

Analise estas duas funções que fazem a mesma coisa:

```javascript
// Versão 1: Loop tradicional
function somarArray1(array) {
  let total = 0;
  for (let i = 0; i < array.length; i++) {
    total += array[i];
  }
  return total;
}

// Versão 2: Recursão
function somarArray2(array, indice = 0) {
  if (indice >= array.length) {
    return 0;
  }
  return array[indice] + somarArray2(array, indice + 1);
}

// Versão 3: Método nativo
function somarArray3(array) {
  return array.reduce((total, num) => total + num, 0);
}
```

**Perguntas:**
1. Qual versão você acha mais eficiente? Por quê?
2. O que aconteceria se você tentasse somar um array com 10.000 elementos usando a versão recursiva?
3. Qual versão é mais legível e fácil de entender?
4. Em uma aplicação web com muitos usuários, qual versão você escolheria? Por quê?
5. Existe algum caso onde a recursão seria melhor que a iteração?

**Sua Resposta:**
```
[Escreva suas reflexões aqui]
```

---

### Reflexão 2: Escopo e Vazamentos de Memória

Analise este código:

```javascript
let contador = 0;

function criarContador() {
  let valor = 0;
  return function() {
    valor++;
    contador++;
    return valor;
  };
}

let contador1 = criarContador();
let contador2 = criarContador();

console.log(contador1()); // 1
console.log(contador1()); // 2
console.log(contador2()); // 1
console.log(contador);    // 3
```

**Perguntas:**
1. Por que `contador1()` e `contador2()` retornam valores diferentes?
2. O que acontece com a variável `valor` dentro de `criarContador`? Ela é destruída quando a função termina?
3. A variável `contador` global pode causar problemas? Quais?
4. Como você poderia melhorar este código para evitar problemas?
5. Em uma aplicação grande, qual seria o impacto de ter muitas variáveis globais?

**Sua Resposta:**
```
[Escreva suas reflexões aqui]
```

---

### Reflexão 3: Edge Cases e Validação

Analise esta função:

```javascript
function dividir(a, b) {
  return a / b;
}
```

**Perguntas:**
1. Quais são os possíveis problemas com esta função?
2. O que acontece se `b` for 0?
3. O que acontece se `a` ou `b` não forem números?
4. O que acontece se `a` ou `b` forem `null` ou `undefined`?
5. Como você melhoraria esta função para lidar com todos esses casos?
6. Em uma aplicação real, qual seria o impacto de não validar os parâmetros?

**Sua Resposta:**
```
[Escreva suas reflexões aqui]
```

---

### Reflexão 4: Organização e Manutenibilidade

Analise estes dois códigos que fazem a mesma coisa:

**Código A:**
```javascript
function processarDados(dados) {
  let resultado = [];
  for (let i = 0; i < dados.length; i++) {
    if (dados[i].idade >= 18) {
      if (dados[i].ativo === true) {
        let nome = dados[i].nome.toUpperCase();
        let idade = dados[i].idade;
        resultado.push({nome: nome, idade: idade});
      }
    }
  }
  return resultado;
}
```

**Código B:**
```javascript
function eMaiorDeIdade(pessoa) {
  return pessoa.idade >= 18;
}

function estaAtivo(pessoa) {
  return pessoa.ativo === true;
}

function formatarNome(pessoa) {
  return pessoa.nome.toUpperCase();
}

function processarDados(dados) {
  return dados
    .filter(eMaiorDeIdade)
    .filter(estaAtivo)
    .map(pessoa => ({
      nome: formatarNome(pessoa),
      idade: pessoa.idade
    }));
}
```

**Perguntas:**
1. Qual código é mais fácil de entender? Por quê?
2. Qual código é mais fácil de testar? Por quê?
3. Qual código é mais fácil de modificar? Por quê?
4. Se você precisasse adicionar uma nova validação (ex: verificar se tem email), qual código seria mais fácil de modificar?
5. Em um projeto grande com vários desenvolvedores, qual abordagem você preferiria? Por quê?
6. Qual código segue melhor o princípio DRY (Don't Repeat Yourself)?

**Sua Resposta:**
```
[Escreva suas reflexões aqui]
```

---

## 🎯 Desafio Final: Calculadora Completa

**Objetivo:** Criar uma calculadora completa usando funções.

**Instruções:**
1. Crie funções para cada operação: `somar`, `subtrair`, `multiplicar`, `dividir`, `potencia`, `raizQuadrada`
2. Crie uma função `calculadora` que recebe:
   - operação (string: "somar", "subtrair", etc.)
   - números (usando rest parameters)
3. A função deve validar os parâmetros e retornar mensagens de erro apropriadas
4. Crie uma função `calcularExpressao` que recebe uma string como "10 + 5 * 2" e retorna o resultado
5. Teste todas as funcionalidades

**Código Base:**
```javascript
// Escreva suas funções aqui





// Teste sua calculadora aqui
console.log(calculadora("somar", 10, 5, 3));        // 18
console.log(calculadora("multiplicar", 4, 5));      // 20
console.log(calculadora("dividir", 10, 0));         // "Erro: divisão por zero"
console.log(calculadora("potencia", 2, 3));         // 8
```

---

## ✅ Checklist de Aprendizado

Antes de prosseguir, verifique se você consegue:

- [ ] Criar funções usando function declaration, function expression e arrow functions
- [ ] Usar parâmetros padrão em funções
- [ ] Usar rest parameters para aceitar múltiplos argumentos
- [ ] Entender a diferença entre return e console.log
- [ ] Explicar o conceito de escopo (global, função, bloco)
- [ ] Entender como funciona a call stack
- [ ] Criar funções recursivas com casos base e recursivos
- [ ] Usar funções nativas do JavaScript (Math, String, Array, etc.)
- [ ] Identificar e corrigir erros comuns em funções
- [ ] Organizar código usando funções para melhorar legibilidade

---

## 📚 Próximos Passos

Depois de completar os exercícios e reflexões, você estará pronto para aprender sobre:
- Performance e otimização de funções
- Boas práticas e padrões de código
- Closures (conceito avançado)
- Higher-order functions
- Callbacks e programação assíncrona

**Lembre-se:** Não tenha pressa! Entenda cada exercício completamente antes de prosseguir. A programação é uma habilidade que se desenvolve com prática constante.

---

## 💬 Dica do Professor

Funções são a base de código organizado e reutilizável. Quanto mais você praticar criando funções, mais natural se tornará pensar em termos de "blocos de código reutilizáveis". 

Sempre pergunte a si mesmo:
- "Esta lógica será usada mais de uma vez?" → Crie uma função
- "Este código está muito longo?" → Quebre em funções menores
- "Esta função faz mais de uma coisa?" → Divida em funções menores

Boa prática! 🚀

