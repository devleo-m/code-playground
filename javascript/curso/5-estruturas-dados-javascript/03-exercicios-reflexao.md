# Aula 5 - Exercícios Práticos e Reflexão: Estruturas de Dados

## 📝 Instruções

Complete os exercícios abaixo. Para cada exercício:
1. Escreva o código JavaScript necessário
2. Teste o código no console do navegador ou em um arquivo HTML
3. Anote os resultados que você obteve
4. Reflita sobre as perguntas de reflexão ao final

---

## 🎯 Exercício 1: Trabalhando com Arrays

### Objetivo
Praticar criação, manipulação e uso de métodos de arrays.

### Tarefa

1. **Crie um array com seus 5 filmes favoritos:**
   ```javascript
   // Seu código aqui
   const filmes = [/* seus filmes */];
   ```

2. **Adicione 2 novos filmes ao array:**
   - Um no final usando `push()`
   - Um no início usando `unshift()`

3. **Remova o primeiro filme do array usando `shift()`**

4. **Use `map()` para criar um novo array com os filmes em maiúsculas:**
   ```javascript
   // Dica: use .toUpperCase() nas strings
   ```

5. **Use `filter()` para criar um array apenas com filmes que tenham mais de 10 caracteres no nome**

6. **Use `find()` para encontrar o primeiro filme que comece com a letra "A"**

**Sua resposta:**
- Escreva o código completo
- Anote os resultados de cada operação
- Explique a diferença entre `map()`, `filter()` e `find()`

---

## 🎯 Exercício 2: Manipulando Arrays com Métodos Avançados

### Objetivo
Praticar métodos como `reduce()`, `some()`, `every()` e `sort()`.

### Tarefa

1. **Crie um array de números:**
   ```javascript
   const numeros = [10, 5, 20, 15, 30, 25];
   ```

2. **Use `reduce()` para calcular a soma de todos os números**

3. **Use `reduce()` para encontrar o maior número do array**

4. **Use `some()` para verificar se existe algum número maior que 25**

5. **Use `every()` para verificar se todos os números são maiores que 5**

6. **Ordene o array em ordem crescente usando `sort()`**
   - Lembre-se: `sort()` sem função de comparação ordena como strings!
   - Use uma função de comparação para ordenar números

7. **Ordene o array em ordem decrescente**

**Sua resposta:**
- Escreva o código completo
- Explique como funciona a função de comparação no `sort()`
- Qual a diferença entre `some()` e `every()`?

---

## 🎯 Exercício 3: Trabalhando com Map

### Objetivo
Praticar criação e manipulação de Map.

### Tarefa

1. **Crie um Map que armazene informações de produtos:**
   ```javascript
   // Chave: ID do produto (número)
   // Valor: Objeto com nome e preço
   const produtos = new Map();
   ```

2. **Adicione 5 produtos ao Map:**
   - Use IDs numéricos como chave (1, 2, 3, 4, 5)
   - Cada valor deve ser um objeto: `{ nome: "Produto X", preco: 99.90 }`

3. **Crie funções para:**
   - Buscar um produto por ID
   - Verificar se um produto existe
   - Remover um produto
   - Listar todos os produtos (usando `forEach`)

4. **Itere sobre o Map e crie um array com apenas os nomes dos produtos**

5. **Calcule o preço total de todos os produtos usando um loop**

**Sua resposta:**
- Escreva o código completo
- Qual a vantagem de usar Map ao invés de um objeto comum neste caso?
- Quando você usaria Map e quando usaria um objeto?

---

## 🎯 Exercício 4: Trabalhando com Set

### Objetivo
Praticar criação e uso de Set para valores únicos.

### Tarefa

1. **Crie um array com números duplicados:**
   ```javascript
   const numerosComDuplicatas = [1, 2, 2, 3, 3, 3, 4, 5, 5, 5, 6];
   ```

2. **Use Set para remover as duplicatas:**
   - Crie um Set a partir do array
   - Converta o Set de volta para array
   - Verifique o resultado

3. **Crie um sistema de tags para artigos de blog:**
   ```javascript
   const tags = new Set();
   ```
   - Adicione tags: 'javascript', 'programação', 'web', 'javascript', 'tutorial'
   - Verifique quantas tags únicas você tem
   - Verifique se a tag 'javascript' existe
   - Tente adicionar 'javascript' novamente e veja o que acontece

4. **Crie uma função que verifica se um email já está cadastrado:**
   ```javascript
   const emailsCadastrados = new Set(['user1@email.com', 'user2@email.com']);
   
   function verificarEmail(email) {
       // Sua lógica aqui
       // Retorne true se já existe, false se não existe
   }
   
   // Teste com emails novos e existentes
   ```

**Sua resposta:**
- Escreva o código completo
- Por que Set é melhor que Array para verificar existência?
- Em que situações você usaria Set no dia a dia?

---

## 🎯 Exercício 5: Trabalhando com JSON

### Objetivo
Praticar conversão entre objetos JavaScript e JSON.

### Tarefa

1. **Crie um objeto JavaScript representando um usuário:**
   ```javascript
   const usuario = {
       nome: "João Silva",
       idade: 30,
       email: "joao@email.com",
       ativo: true,
       hobbies: ["leitura", "programação", "música"],
       endereco: {
           rua: "Rua das Flores",
           numero: 123,
           cidade: "São Paulo"
       }
   };
   ```

2. **Converta o objeto para JSON:**
   - Use `JSON.stringify()` sem formatação
   - Use `JSON.stringify()` com formatação (indentação de 2 espaços)
   - Compare os resultados

3. **Crie uma string JSON manualmente:**
   ```javascript
   const jsonString = '{"nome":"Maria","idade":25,"ativo":true}';
   ```

4. **Converta a string JSON de volta para objeto:**
   - Use `JSON.parse()`
   - Acesse as propriedades do objeto resultante

5. **Simule salvamento e recuperação de dados:**
   ```javascript
   // Simular localStorage
   const dadosParaSalvar = { tema: "escuro", idioma: "pt-BR" };
   
   // 1. Converter para JSON e "salvar"
   const jsonSalvo = JSON.stringify(dadosParaSalvar);
   console.log("Dados salvos:", jsonSalvo);
   
   // 2. "Recuperar" e converter de volta
   const dadosRecuperados = JSON.parse(jsonSalvo);
   console.log("Dados recuperados:", dadosRecuperados);
   ```

6. **Trate erros ao fazer parse de JSON inválido:**
   ```javascript
   const jsonInvalido = '{nome: "João"}'; // JSON inválido (chave sem aspas)
   
   // Use try/catch para tratar o erro
   ```

**Sua resposta:**
- Escreva o código completo
- Por que JSON não suporta funções?
- Quando você usaria JSON.stringify() e JSON.parse() em uma aplicação real?

---

## 🎯 Exercício 6: Estruturas de Dados Complexas

### Objetivo
Praticar combinação de diferentes estruturas de dados.

### Tarefa

Crie um **sistema de gerenciamento de biblioteca** usando diferentes estruturas:

1. **Array de livros (lista ordenada):**
   ```javascript
   const livros = [
       { id: 1, titulo: "Livro A", autor: "Autor 1" },
       { id: 2, titulo: "Livro B", autor: "Autor 2" },
       { id: 3, titulo: "Livro C", autor: "Autor 1" }
   ];
   ```

2. **Map para busca rápida por ID:**
   ```javascript
   const livrosPorId = new Map();
   // Preencha o Map usando os dados do array
   ```

3. **Set para autores únicos:**
   ```javascript
   const autores = new Set();
   // Adicione todos os autores únicos
   ```

4. **Crie funções para:**
   - Buscar livro por ID (usando Map)
   - Listar todos os autores únicos (usando Set)
   - Filtrar livros por autor (usando Array.filter)
   - Adicionar novo livro (atualizar Array, Map e Set)

5. **Converta os dados da biblioteca para JSON:**
   - Crie um objeto com todas as informações
   - Converta para JSON formatado

**Sua resposta:**
- Escreva o código completo
- Explique por que usar Map para busca por ID é mais eficiente que Array.find()
- Como você organizaria os dados se tivesse milhares de livros?

---

## 🎯 Exercício 7: Análise de Código

### Objetivo
Analisar código existente e identificar problemas ou melhorias.

### Tarefa

Analise o seguinte código e responda:

```javascript
// Código para gerenciar uma lista de tarefas
const tarefas = [];

function adicionarTarefa(tarefa) {
    tarefas.push(tarefa);
}

function removerTarefa(tarefa) {
    const indice = tarefas.indexOf(tarefa);
    if (indice !== -1) {
        tarefas.splice(indice, 1);
    }
}

function marcarComoConcluida(tarefa) {
    const indice = tarefas.indexOf(tarefa);
    if (indice !== -1) {
        tarefas[indice] = tarefa + " [CONCLUÍDA]";
    }
}

function listarTarefas() {
    tarefas.forEach((tarefa, indice) => {
        console.log(`${indice + 1}. ${tarefa}`);
    });
}

// Teste
adicionarTarefa("Comprar leite");
adicionarTarefa("Estudar JavaScript");
adicionarTarefa("Comprar leite"); // Duplicata!
marcarComoConcluida("Estudar JavaScript");
listarTarefas();
```

**Perguntas:**
1. O que acontece se você tentar adicionar a mesma tarefa duas vezes?
2. Como você melhoraria este código para evitar tarefas duplicadas?
3. Qual estrutura de dados você usaria para rastrear tarefas concluídas separadamente?
4. Como você modificaria o código para usar objetos ao invés de strings simples?

**Sua resposta:**
- Analise o código e identifique problemas
- Proponha melhorias
- Reescreva o código com suas melhorias

---

## 🎯 Exercício 8: Desafio - Sistema de Carrinho de Compras

### Objetivo
Criar um sistema completo usando múltiplas estruturas de dados.

### Tarefa

Crie um **sistema de carrinho de compras** com as seguintes funcionalidades:

**Requisitos:**
1. **Array** para manter a ordem dos itens adicionados
2. **Map** para armazenar detalhes dos produtos (ID → {nome, preco})
3. **Set** para IDs de produtos em promoção
4. **JSON** para salvar/recuperar o carrinho

**Funcionalidades:**
- Adicionar produto ao carrinho
- Remover produto do carrinho
- Calcular total do carrinho
- Aplicar desconto de 10% para produtos em promoção
- Listar todos os itens do carrinho
- Salvar carrinho em JSON
- Recuperar carrinho de JSON

**Estrutura sugerida:**
```javascript
// Map de produtos disponíveis
const produtos = new Map();
produtos.set(1, { nome: "Notebook", preco: 2500 });
produtos.set(2, { nome: "Mouse", preco: 50 });
// ... mais produtos

// Set de produtos em promoção
const produtosPromocao = new Set([1, 3]);

// Array do carrinho (armazena IDs)
const carrinho = [];

// Funções aqui...
```

**Sua resposta:**
- Escreva o código completo
- Teste todas as funcionalidades
- Explique suas escolhas de estruturas de dados

---

## 🤔 Perguntas de Reflexão

Responda as seguintes perguntas de forma honesta e detalhada:

### 1. Performance e Eficiência

**a)** Compare o tempo de execução de `array.find()` vs `map.get()` para buscar um elemento:
   - Qual é mais eficiente? Por quê?
   - Em que situação você usaria cada um?

**b)** Se você tivesse uma lista de 1 milhão de emails e precisasse verificar se um email existe:
   - Você usaria Array com `includes()` ou Set com `has()`?
   - Por quê? Qual seria a diferença de performance?

### 2. Escolha de Estrutura de Dados

**a)** Você precisa armazenar informações de alunos em uma escola:
   - Cada aluno tem: nome, idade, matrícula, notas
   - Você precisa buscar aluno por matrícula frequentemente
   - Você precisa listar todos os alunos em ordem alfabética
   - Qual(is) estrutura(s) de dados você usaria? Justifique.

**b)** Você está criando um sistema de tags para posts de blog:
   - Cada post pode ter múltiplas tags
   - Você precisa listar todas as tags únicas do blog
   - Você precisa verificar rapidamente se uma tag existe
   - Qual estrutura você usaria? Por quê?

### 3. JSON e Dados Estruturados

**a)** Por que JSON não suporta funções, undefined ou comentários?
   - Qual o impacto disso no desenvolvimento?
   - Como você contornaria a necessidade de enviar uma função via JSON?

**b)** Você precisa enviar dados de um formulário para um servidor:
   - Por que converter para JSON antes de enviar?
   - O que aconteceria se você tentasse enviar um objeto JavaScript diretamente?

### 4. Edge Cases e Tratamento de Erros

**a)** O que acontece se você tentar fazer `JSON.parse()` em uma string inválida?
   - Como você trataria esse erro em uma aplicação real?
   - Qual a importância de validar JSON antes de fazer parse?

**b)** Se você tentar acessar `array[100]` em um array com apenas 5 elementos:
   - O que retorna?
   - Como você verificaria se um índice existe antes de acessá-lo?

### 5. Impacto em Aplicações Reais

**a)** Imagine uma aplicação de e-commerce com milhares de produtos:
   - Como você organizaria os dados para busca eficiente?
   - Qual estrutura usaria para o carrinho de compras? Por quê?
   - Como você lidaria com produtos duplicados no carrinho?

**b)** Em uma aplicação de chat:
   - Como você armazenaria as mensagens? (Array, Map, ou outro?)
   - Como evitaria mensagens duplicadas?
   - Como você salvaria o histórico de mensagens para recuperar depois?

### 6. Boas Práticas

**a)** É uma boa prática modificar arrays diretamente com métodos como `sort()` e `reverse()`?
   - Por quê?
   - Como você criaria uma cópia antes de modificar?

**b)** Quando você criaria uma cópia de uma estrutura de dados ao invés de modificar a original?
   - Dê exemplos práticos
   - Qual a importância disso?

---

## ✅ Checklist de Conclusão

Antes de prosseguir, certifique-se de que:

- [ ] Completei todos os exercícios práticos
- [ ] Testei cada código no console do navegador
- [ ] Entendi a diferença entre Array, Map, Set e Objeto
- [ ] Sei quando usar cada estrutura de dados
- [ ] Entendi como funciona JSON.stringify() e JSON.parse()
- [ ] Respondi todas as perguntas de reflexão
- [ ] Analisei o código do Exercício 7 e propus melhorias
- [ ] Completei o desafio do Exercício 8

---

## 🚀 Próximo Passo

Após completar todos os exercícios e responder as perguntas de reflexão, você estará pronto para:
- Receber feedback sobre seu desempenho
- Aprender sobre performance e boas práticas
- Avançar para os próximos tópicos

**Envie suas respostas para análise!** 📝





