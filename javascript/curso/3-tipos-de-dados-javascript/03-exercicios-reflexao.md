# Aula 3 - Exercícios Práticos e Reflexão: Tipos de Dados

## 📝 Instruções

Complete os exercícios abaixo. Para cada exercício:
1. Escreva o código JavaScript necessário
2. Teste o código no console do navegador ou em um arquivo HTML
3. Anote os resultados que você obteve
4. Reflita sobre as perguntas de reflexão ao final

---

## 🎯 Exercício 1: Identificando Tipos de Dados

### Objetivo
Praticar o uso do operador `typeof` para identificar diferentes tipos de dados.

### Tarefa
Crie variáveis com os seguintes valores e use `typeof` para identificar o tipo de cada uma:

```javascript
// Suas variáveis aqui
let valor1 = 42;
let valor2 = "JavaScript";
let valor3 = true;
let valor4 = null;
let valor5 = undefined;
let valor6 = [1, 2, 3];
let valor7 = { nome: "João" };
let valor8 = function() {};
let valor9 = 42n;
let valor10 = Symbol("id");
```

**Sua resposta:**
- Escreva o código que verifica o tipo de cada variável
- Anote os resultados que você obteve
- Explique por que `typeof null` retorna `"object"` (pesquise se necessário)

---

## 🎯 Exercício 2: Trabalhando com Números

### Objetivo
Praticar operações com números e entender limitações de precisão.

### Tarefa
1. Crie variáveis com diferentes representações numéricas:
   - Um número inteiro
   - Um número decimal
   - Um número em notação hexadecimal
   - Um número em notação binária
   - Um número em notação exponencial

2. Realize as seguintes operações e anote os resultados:
   ```javascript
   let a = 0.1;
   let b = 0.2;
   let soma = a + b;
   
   // O que acontece quando você compara soma com 0.3?
   console.log(soma === 0.3); // Qual é o resultado?
   ```

3. Teste os limites do tipo Number:
   ```javascript
   console.log(Number.MAX_SAFE_INTEGER);
   console.log(Number.MAX_SAFE_INTEGER + 1);
   // O que acontece com a precisão?
   ```

**Sua resposta:**
- Escreva o código completo
- Explique por que `0.1 + 0.2` não é exatamente igual a `0.3`
- Quando você usaria `BigInt` ao invés de `Number`?

---

## 🎯 Exercício 3: Manipulando Strings

### Objetivo
Praticar criação e manipulação de strings usando diferentes métodos.

### Tarefa
1. Crie uma variável com seu nome completo usando template literals:
   ```javascript
   let nome = "Seu Nome";
   let sobrenome = "Seu Sobrenome";
   // Use template literal para criar nomeCompleto
   ```

2. Crie uma string multilinha usando template literals que contenha:
   - Uma saudação
   - Seu nome
   - Uma mensagem de boas-vindas

3. Use métodos de string para:
   - Converter seu nome para maiúsculas
   - Verificar o comprimento do seu nome
   - Encontrar a posição de uma letra específica
   - Extrair uma parte do seu nome usando `substring`

**Sua resposta:**
- Escreva o código completo
- Qual a diferença entre usar aspas simples, duplas e template literals?
- Quando você usaria cada uma?

---

## 🎯 Exercício 4: Lógica com Booleanos

### Objetivo
Praticar uso de valores booleanos e entender valores truthy/falsy.

### Tarefa
1. Crie variáveis e teste quais são truthy e quais são falsy:
   ```javascript
   let valores = [
       0,
       "",
       null,
       undefined,
       false,
       true,
       1,
       "texto",
       [],
       {},
       function() {}
   ];
   
   // Teste cada valor em um if
   // Anote quais são truthy e quais são falsy
   ```

2. Crie uma função que verifica se uma pessoa pode votar:
   ```javascript
   function podeVotar(idade) {
       // Sua lógica aqui
       // Retorne true se idade >= 16, false caso contrário
   }
   
   // Teste com diferentes idades
   console.log(podeVotar(15)); // false
   console.log(podeVotar(16)); // true
   console.log(poveVotar(25)); // true
   ```

3. Crie uma função que verifica se uma string está vazia:
   ```javascript
   function stringVazia(texto) {
       // Sua lógica aqui
       // Retorne true se a string estiver vazia, false caso contrário
   }
   ```

**Sua resposta:**
- Escreva o código completo
- Liste todos os valores falsy em JavaScript
- Por que é importante entender valores truthy/falsy?

---

## 🎯 Exercício 5: Diferença entre null e undefined

### Objetivo
Entender a diferença prática entre `null` e `undefined`.

### Tarefa
1. Crie situações que resultem em `undefined`:
   ```javascript
   // Situação 1: Variável declarada mas não inicializada
   
   // Situação 2: Propriedade que não existe em um objeto
   
   // Situação 3: Função sem return
   
   // Situação 4: Parâmetro não fornecido
   ```

2. Crie situações que usem `null` intencionalmente:
   ```javascript
   // Situação 1: Resetar uma variável de objeto
   
   // Situação 2: Indicar ausência intencional de valor
   
   // Situação 3: Resultado de busca que não encontrou nada
   ```

3. Compare `null` e `undefined`:
   ```javascript
   let a = null;
   let b = undefined;
   
   console.log(a == b);  // O que retorna?
   console.log(a === b); // O que retorna?
   console.log(typeof a); // O que retorna?
   console.log(typeof b); // O que retorna?
   ```

**Sua resposta:**
- Escreva o código completo para cada situação
- Explique a diferença prática entre `null` e `undefined`
- Quando você usaria cada um no desenvolvimento real?

---

## 🎯 Exercício 6: Trabalhando com Objetos

### Objetivo
Praticar criação e manipulação de objetos.

### Tarefa
1. Crie um objeto que represente uma pessoa com as seguintes propriedades:
   - `nome` (string)
   - `idade` (number)
   - `cidade` (string)
   - `ativo` (boolean)

2. Adicione um método ao objeto que retorne uma mensagem de apresentação:
   ```javascript
   let pessoa = {
       nome: "João",
       idade: 25,
       // ... outras propriedades
       
       apresentar: function() {
           // Retorne uma string como: "Olá, eu sou João, tenho 25 anos e moro em São Paulo"
       }
   };
   ```

3. Crie um objeto aninhado:
   ```javascript
   let empresa = {
       nome: "Tech Corp",
       endereco: {
           rua: "Rua das Flores",
           numero: 123,
           cidade: "São Paulo"
       },
       funcionarios: [
           { nome: "João", cargo: "Desenvolvedor" },
           { nome: "Maria", cargo: "Designer" }
       ]
   };
   
   // Acesse e exiba:
   // - O nome da empresa
   // - A cidade do endereço
   // - O nome do primeiro funcionário
   ```

**Sua resposta:**
- Escreva o código completo
- Qual a diferença entre `pessoa.nome` e `pessoa["nome"]`?
- Quando você usaria cada notação?

---

## 🎯 Exercício 7: Conversão de Tipos

### Objetivo
Praticar conversões explícitas e entender conversões implícitas.

### Tarefa
1. Realize conversões explícitas:
   ```javascript
   let numero = 42;
   let texto = "123";
   let booleano = true;
   
   // Converta número para string (3 formas diferentes)
   
   // Converta string para número (3 formas diferentes)
   
   // Converta qualquer valor para boolean (2 formas diferentes)
   ```

2. Teste conversões implícitas (cuidado!):
   ```javascript
   console.log("5" + 3);      // O que acontece?
   console.log("5" - 3);      // O que acontece?
   console.log("5" * 3);      // O que acontece?
   console.log("5" / 3);      // O que acontece?
   console.log(true + 1);     // O que acontece?
   console.log(false + 1);    // O que acontece?
   console.log("" + 42);      // O que acontece?
   ```

3. Crie uma função que valida se um valor pode ser convertido para número:
   ```javascript
   function podeSerNumero(valor) {
       // Sua lógica aqui
       // Retorne true se o valor pode ser convertido para número, false caso contrário
       // Dica: use Number() e isNaN()
   }
   
   console.log(podeSerNumero("42"));     // true
   console.log(podeSerNumero("abc"));    // false
   console.log(podeSerNumero("123abc"));  // false
   console.log(podeSerNumero(42));        // true
   ```

**Sua resposta:**
- Escreva o código completo
- Por que é importante fazer conversões explícitas?
- Quais problemas as conversões implícitas podem causar?

---

## 🎯 Exercício 8: Desafio - Sistema de Cadastro

### Objetivo
Aplicar todos os conceitos aprendidos em um projeto prático.

### Tarefa
Crie um sistema simples de cadastro de usuários que:

1. Armazene informações de usuários em objetos:
   ```javascript
   // Cada usuário deve ter:
   // - id (number ou BigInt)
   // - nome (string)
   // - email (string)
   // - idade (number)
   // - ativo (boolean)
   // - dataCadastro (Date ou string)
   ```

2. Crie funções para:
   - Adicionar um novo usuário
   - Buscar um usuário por ID (retornar `null` se não encontrar)
   - Listar todos os usuários ativos
   - Desativar um usuário (mudar `ativo` para `false`)

3. Valide os dados antes de adicionar:
   - Nome não pode estar vazio
   - Email deve conter "@"
   - Idade deve ser um número positivo
   - ID deve ser único

**Sua resposta:**
- Escreva o código completo do sistema
- Use diferentes tipos de dados apropriadamente
- Inclua tratamento para casos onde dados podem ser `null` ou `undefined`

---

## 🤔 Perguntas de Reflexão

Responda as seguintes perguntas com base no que você aprendeu e nos exercícios que realizou:

### 1. Por que JavaScript tem `null` e `undefined`? Não seria mais simples ter apenas um?

**Sua resposta:**

---

### 2. Quando você usaria `BigInt` ao invés de `Number`? Dê exemplos práticos de situações reais onde `BigInt` seria necessário.

**Sua resposta:**

---

### 3. Por que `typeof null` retorna `"object"` mesmo sendo um tipo primitivo? Pesquise sobre o histórico do JavaScript e explique.

**Sua resposta:**

---

### 4. Qual é o impacto de usar conversões implícitas de tipos em uma aplicação grande? Pense em:
   - Manutenibilidade do código
   - Possibilidade de bugs
   - Performance
   - Legibilidade

**Sua resposta:**

---

### 5. Em uma aplicação web com milhares de usuários, quais problemas de performance podem ocorrer se você:
   - Não validar tipos de dados antes de processar?
   - Usar conversões de tipo desnecessárias em loops?
   - Não verificar se valores são `null` ou `undefined` antes de usar?

**Sua resposta:**

---

### 6. Como você garantiria que um valor recebido de um formulário HTML (que sempre vem como string) seja tratado corretamente antes de ser usado em cálculos?

**Sua resposta:**

---

### 7. Qual seria a melhor forma de criar um sistema que precisa armazenar IDs muito grandes (maiores que `Number.MAX_SAFE_INTEGER`)? Considere:
   - Uso de `BigInt`
   - Uso de strings
   - Impacto na performance
   - Compatibilidade com APIs e bancos de dados

**Sua resposta:**

---

### 8. Pense em um cenário onde você precisa processar dados de um arquivo JSON. Como você garantiria que:
   - Todos os tipos de dados estão corretos?
   - Valores não esperados (como `null` ou `undefined`) são tratados adequadamente?
   - O código não quebra se algum campo estiver faltando?

**Sua resposta:**

---

## ✅ Checklist de Aprendizado

Antes de prosseguir, certifique-se de que você:

- [ ] Consegue identificar todos os tipos primitivos em JavaScript
- [ ] Entende a diferença entre `null` e `undefined`
- [ ] Sabe quando usar `BigInt` vs `Number`
- [ ] Consegue criar e manipular strings usando template literals
- [ ] Entende valores truthy e falsy
- [ ] Sabe como verificar tipos usando `typeof`
- [ ] Consegue criar e manipular objetos
- [ ] Entende conversões explícitas e implícitas de tipos
- [ ] Sabe quando usar cada tipo de dado apropriadamente
- [ ] Consegue validar tipos de dados antes de usar

---

## 📤 Envie Suas Respostas

Após completar todos os exercícios e responder as perguntas de reflexão, envie suas respostas para análise. Inclua:
1. Todo o código que você escreveu
2. Os resultados que você obteve
3. Suas respostas às perguntas de reflexão
4. Qualquer dúvida ou dificuldade que você encontrou

**Boa sorte! 🚀**





