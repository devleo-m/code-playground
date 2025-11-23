# Aula 5: Estruturas de Dados em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 4**, você aprendeu:
- ✅ Conversão de tipos (coerção de tipos)
- ✅ Conversão explícita vs implícita
- ✅ Métodos de conversão (Number(), String(), Boolean(), etc.)
- ✅ Truthy e falsy values
- ✅ Operadores de comparação e igualdade

Agora vamos aprender sobre **estruturas de dados** - como organizar, armazenar e acessar informações de forma eficiente em JavaScript!

---

## 🎯 O que são Estruturas de Dados?

**Definição:** Uma estrutura de dados é um formato para **organizar, gerenciar e armazenar dados** de forma que permita **acesso e modificação eficientes**. 

Em JavaScript, temos estruturas de dados **primitivas** (built-in) e **não-primitivas** (que precisamos implementar ou usar de forma mais avançada).

### Estruturas Primitivas (Built-in)
- Vêm por padrão com a linguagem
- Podem ser usadas imediatamente
- Exemplos: Arrays, Objetos

### Estruturas Não-Primitivas
- Não vêm por padrão
- Precisam ser implementadas ou usadas de forma mais avançada
- Exemplos: Pilhas (Stacks), Filas (Queues), Árvores, Grafos

---

## 📊 Classificação das Estruturas de Dados em JavaScript

JavaScript organiza estruturas de dados em diferentes categorias:

1. **Keyed Collections** (Coleções com Chaves)
   - Map
   - Set

2. **Indexed Collections** (Coleções Indexadas)
   - Arrays

3. **Structured Data** (Dados Estruturados)
   - JSON
   - Objetos

4. **Outras Estruturas**
   - Objetos literais
   - Arrays de objetos
   - Objetos aninhados

---

## 🔑 1. Keyed Collections (Coleções com Chaves)

### Definição

**Keyed Collections** são coleções de dados que são **ordenadas por chave, não por índice**. Elas são **associativas** por natureza, o que significa que você acessa os valores através de chaves específicas, não por posição numérica.

As principais Keyed Collections em JavaScript são:
- **Map**: Armazena pares chave-valor
- **Set**: Armazena valores únicos

### Características Importantes

- São **iteráveis** na ordem de inserção
- Permitem qualquer tipo de dado como chave (Map)
- Garantem valores únicos (Set)
- Têm métodos específicos para manipulação
- Foram introduzidas no ES6 (2015)

---

### 1.1 Map

#### O que é Map?

Um `Map` é uma coleção de pares **chave-valor** onde tanto a chave quanto o valor podem ser de qualquer tipo. Diferente de objetos, um Map mantém a ordem de inserção e permite usar qualquer tipo como chave.

#### Sintaxe Básica

```javascript
// Criando um Map vazio
const meuMap = new Map();

// Criando um Map com valores iniciais
const mapInicial = new Map([
  ['chave1', 'valor1'],
  ['chave2', 'valor2'],
  [1, 'número um'],
  [true, 'booleano']
]);
```

#### Métodos Principais do Map

```javascript
const mapa = new Map();

// set() - Adiciona ou atualiza um par chave-valor
mapa.set('nome', 'João');
mapa.set('idade', 30);
mapa.set(1, 'um');
mapa.set({ id: 1 }, 'objeto como chave');

// get() - Obtém o valor associado a uma chave
console.log(mapa.get('nome')); // "João"
console.log(mapa.get('idade')); // 30

// has() - Verifica se uma chave existe
console.log(mapa.has('nome')); // true
console.log(mapa.has('email')); // false

// delete() - Remove um par chave-valor
mapa.delete('idade');
console.log(mapa.has('idade')); // false

// clear() - Remove todos os elementos
mapa.clear();
console.log(mapa.size); // 0

// size - Propriedade que retorna o número de elementos
console.log(mapa.size); // número de elementos
```

#### Iteração sobre Map

```javascript
const mapa = new Map([
  ['nome', 'Maria'],
  ['idade', 25],
  ['cidade', 'São Paulo']
]);

// Iteração com for...of
for (const [chave, valor] of mapa) {
  console.log(`${chave}: ${valor}`);
}
// nome: Maria
// idade: 25
// cidade: São Paulo

// Iteração apenas sobre chaves
for (const chave of mapa.keys()) {
  console.log(chave);
}

// Iteração apenas sobre valores
for (const valor of mapa.values()) {
  console.log(valor);
}

// Iteração com forEach
mapa.forEach((valor, chave) => {
  console.log(`${chave} = ${valor}`);
});

// Converter Map para Array
const arrayDeMap = Array.from(mapa);
console.log(arrayDeMap);
// [['nome', 'Maria'], ['idade', 25], ['cidade', 'São Paulo']]
```

#### Diferenças entre Map e Objeto

```javascript
// OBJETO
const objeto = {};
objeto[1] = 'um';
objeto['1'] = 'um string'; // Sobrescreve o anterior
console.log(objeto); // { '1': 'um string' }

// MAP
const mapa = new Map();
mapa.set(1, 'um');
mapa.set('1', 'um string'); // Mantém ambos
console.log(mapa.get(1)); // 'um'
console.log(mapa.get('1')); // 'um string'

// Objetos têm chaves apenas como strings (ou Symbols)
// Maps podem ter qualquer tipo como chave
```

#### Quando Usar Map?

- Quando você precisa de chaves que não sejam strings
- Quando precisa manter a ordem de inserção
- Quando precisa iterar frequentemente
- Quando o número de pares chave-valor muda frequentemente
- Quando precisa de melhor performance para adicionar/remover elementos

---

### 1.2 Set

#### O que é Set?

Um `Set` é uma coleção de valores **únicos** (sem duplicatas). Cada valor pode aparecer apenas uma vez no Set.

#### Sintaxe Básica

```javascript
// Criando um Set vazio
const meuSet = new Set();

// Criando um Set com valores iniciais
const setInicial = new Set([1, 2, 3, 4, 5]);
const setComStrings = new Set(['a', 'b', 'c']);
```

#### Métodos Principais do Set

```javascript
const conjunto = new Set();

// add() - Adiciona um valor ao Set
conjunto.add(1);
conjunto.add(2);
conjunto.add(3);
conjunto.add(2); // Duplicado - será ignorado
console.log(conjunto); // Set { 1, 2, 3 }

// has() - Verifica se um valor existe
console.log(conjunto.has(2)); // true
console.log(conjunto.has(5)); // false

// delete() - Remove um valor
conjunto.delete(2);
console.log(conjunto.has(2)); // false

// clear() - Remove todos os valores
conjunto.clear();
console.log(conjunto.size); // 0

// size - Propriedade que retorna o número de elementos
console.log(conjunto.size); // número de elementos únicos
```

#### Iteração sobre Set

```javascript
const conjunto = new Set([1, 2, 3, 4, 5]);

// Iteração com for...of
for (const valor of conjunto) {
  console.log(valor);
}

// Iteração com forEach
conjunto.forEach((valor) => {
  console.log(valor);
});

// Converter Set para Array
const arrayDoSet = Array.from(conjunto);
console.log(arrayDoSet); // [1, 2, 3, 4, 5]
```

#### Casos de Uso Comuns

```javascript
// Remover duplicatas de um array
const arrayComDuplicatas = [1, 2, 2, 3, 3, 3, 4, 5];
const arraySemDuplicatas = Array.from(new Set(arrayComDuplicatas));
console.log(arraySemDuplicatas); // [1, 2, 3, 4, 5]

// Verificar se um valor existe rapidamente
const emails = new Set(['user1@email.com', 'user2@email.com']);
console.log(emails.has('user1@email.com')); // true

// Armazenar valores únicos
const tags = new Set();
tags.add('javascript');
tags.add('programação');
tags.add('javascript'); // Ignorado
console.log(tags); // Set { 'javascript', 'programação' }
```

#### Quando Usar Set?

- Quando você precisa garantir valores únicos
- Quando precisa verificar existência rapidamente (melhor que array.includes())
- Quando precisa remover duplicatas
- Quando a ordem de inserção importa

---

## 📋 2. Indexed Collections (Coleções Indexadas)

### Definição

**Indexed Collections** são coleções que possuem **índices numéricos**, ou seja, coleções de dados que são **ordenadas por um valor de índice**. Em JavaScript, o principal exemplo é o **Array**.

### Características

- Acessadas por posição numérica (índice)
- Índices começam em 0 (zero-based)
- Mantêm a ordem dos elementos
- Podem conter qualquer tipo de dado
- Têm métodos poderosos para manipulação

---

## 📦 3. Arrays

### O que são Arrays?

**Arrays** são objetos que armazenam uma **coleção de itens** e podem ser atribuídos a uma variável. Eles têm métodos próprios que podem realizar operações no array.

### Características Importantes

- Arrays são **objetos especiais** em JavaScript
- Índices são numéricos e começam em 0
- Podem conter qualquer tipo de dado (números, strings, objetos, outros arrays)
- Têm propriedade `length` que indica o número de elementos
- São mutáveis (podem ser modificados após criação)

---

### 3.1 Criando Arrays

```javascript
// Método 1: Array Literal (mais comum)
const frutas = ['maçã', 'banana', 'laranja'];
const numeros = [1, 2, 3, 4, 5];
const misto = [1, 'dois', true, null, undefined];

// Método 2: Construtor Array()
const frutas2 = new Array('maçã', 'banana', 'laranja');
const numeros2 = new Array(5); // Cria array com 5 elementos vazios
const numeros3 = new Array(1, 2, 3); // Cria array com valores

// Método 3: Array.from() - ES6+
const arrayDeString = Array.from('JavaScript');
console.log(arrayDeString); // ['J', 'a', 'v', 'a', 'S', 'c', 'r', 'i', 'p', 't']

// Método 4: Array.of() - ES6+
const arrayOf = Array.of(1, 2, 3);
console.log(arrayOf); // [1, 2, 3]
```

---

### 3.2 Acessando Elementos

```javascript
const frutas = ['maçã', 'banana', 'laranja'];

// Acesso por índice
console.log(frutas[0]); // 'maçã'
console.log(frutas[1]); // 'banana'
console.log(frutas[2]); // 'laranja'
console.log(frutas[3]); // undefined (não existe)

// Propriedade length
console.log(frutas.length); // 3

// Último elemento
console.log(frutas[frutas.length - 1]); // 'laranja'

// Modificando elementos
frutas[1] = 'uva';
console.log(frutas); // ['maçã', 'uva', 'laranja']
```

---

### 3.3 Métodos de Array - Adicionar/Remover

```javascript
const frutas = ['maçã', 'banana'];

// push() - Adiciona ao final
frutas.push('laranja');
console.log(frutas); // ['maçã', 'banana', 'laranja']

// pop() - Remove do final e retorna o elemento
const ultima = frutas.pop();
console.log(ultima); // 'laranja'
console.log(frutas); // ['maçã', 'banana']

// unshift() - Adiciona ao início
frutas.unshift('uva');
console.log(frutas); // ['uva', 'maçã', 'banana']

// shift() - Remove do início e retorna o elemento
const primeira = frutas.shift();
console.log(primeira); // 'uva'
console.log(frutas); // ['maçã', 'banana']

// splice() - Adiciona/remove elementos em qualquer posição
const numeros = [1, 2, 3, 4, 5];
numeros.splice(2, 1); // Remove 1 elemento a partir do índice 2
console.log(numeros); // [1, 2, 4, 5]

numeros.splice(2, 0, 3); // Adiciona 3 no índice 2 sem remover nada
console.log(numeros); // [1, 2, 3, 4, 5]

numeros.splice(2, 1, 10); // Remove 1 elemento e adiciona 10
console.log(numeros); // [1, 2, 10, 4, 5]
```

---

### 3.4 Métodos de Array - Buscar e Verificar

```javascript
const frutas = ['maçã', 'banana', 'laranja', 'banana'];

// indexOf() - Retorna o índice da primeira ocorrência
console.log(frutas.indexOf('banana')); // 1
console.log(frutas.indexOf('uva')); // -1 (não encontrado)

// lastIndexOf() - Retorna o índice da última ocorrência
console.log(frutas.lastIndexOf('banana')); // 3

// includes() - Verifica se contém o elemento (ES6+)
console.log(frutas.includes('maçã')); // true
console.log(frutas.includes('uva')); // false

// find() - Encontra o primeiro elemento que satisfaz a condição
const numeros = [1, 5, 10, 15, 20];
const maiorQue10 = numeros.find(num => num > 10);
console.log(maiorQue10); // 15

// findIndex() - Retorna o índice do primeiro elemento que satisfaz
const indice = numeros.findIndex(num => num > 10);
console.log(indice); // 3
```

---

### 3.5 Métodos de Array - Transformação

```javascript
const numeros = [1, 2, 3, 4, 5];

// map() - Cria novo array transformando cada elemento
const dobrados = numeros.map(num => num * 2);
console.log(dobrados); // [2, 4, 6, 8, 10]

// filter() - Cria novo array com elementos que passam no teste
const pares = numeros.filter(num => num % 2 === 0);
console.log(pares); // [2, 4]

// reduce() - Reduz array a um único valor
const soma = numeros.reduce((acc, num) => acc + num, 0);
console.log(soma); // 15

// reduceRight() - Reduz da direita para esquerda
const subtracao = numeros.reduceRight((acc, num) => acc - num);
console.log(subtracao); // -13 (5 - 4 - 3 - 2 - 1)

// flat() - Achata arrays aninhados (ES2019)
const aninhado = [1, [2, 3], [4, [5, 6]]];
const achatado = aninhado.flat(2); // profundidade 2
console.log(achatado); // [1, 2, 3, 4, 5, 6]

// flatMap() - Combina map() e flat() (ES2019)
const palavras = ['olá mundo', 'javascript é legal'];
const letras = palavras.flatMap(palavra => palavra.split(' '));
console.log(letras); // ['olá', 'mundo', 'javascript', 'é', 'legal']
```

---

### 3.6 Métodos de Array - Ordenação e Reversão

```javascript
const frutas = ['banana', 'maçã', 'laranja', 'uva'];

// sort() - Ordena o array (modifica o original)
frutas.sort();
console.log(frutas); // ['banana', 'laranja', 'maçã', 'uva'] (alfabética)

const numeros = [10, 5, 40, 25, 1000];
numeros.sort(); // Ordenação como strings!
console.log(numeros); // [10, 1000, 25, 40, 5] (errado!)

// Ordenação numérica correta
numeros.sort((a, b) => a - b); // Crescente
console.log(numeros); // [5, 10, 25, 40, 1000]

numeros.sort((a, b) => b - a); // Decrescente
console.log(numeros); // [1000, 40, 25, 10, 5]

// reverse() - Inverte a ordem do array
const arr = [1, 2, 3, 4, 5];
arr.reverse();
console.log(arr); // [5, 4, 3, 2, 1]
```

---

### 3.7 Métodos de Array - Iteração

```javascript
const frutas = ['maçã', 'banana', 'laranja'];

// forEach() - Executa função para cada elemento
frutas.forEach((fruta, indice) => {
  console.log(`${indice}: ${fruta}`);
});
// 0: maçã
// 1: banana
// 2: laranja

// for...of - Loop moderno
for (const fruta of frutas) {
  console.log(fruta);
}

// for tradicional
for (let i = 0; i < frutas.length; i++) {
  console.log(frutas[i]);
}
```

---

### 3.8 Métodos de Array - Verificação

```javascript
const numeros = [1, 2, 3, 4, 5];

// every() - Verifica se todos os elementos passam no teste
const todosPares = numeros.every(num => num % 2 === 0);
console.log(todosPares); // false

// some() - Verifica se pelo menos um elemento passa no teste
const algumPar = numeros.some(num => num % 2 === 0);
console.log(algumPar); // true
```

---

### 3.9 Métodos de Array - Criação de Novos Arrays

```javascript
// concat() - Combina arrays
const arr1 = [1, 2, 3];
const arr2 = [4, 5, 6];
const combinado = arr1.concat(arr2);
console.log(combinado); // [1, 2, 3, 4, 5, 6]

// Spread operator (ES6+) - Alternativa moderna
const combinado2 = [...arr1, ...arr2];
console.log(combinado2); // [1, 2, 3, 4, 5, 6]

// slice() - Extrai parte do array (não modifica original)
const numeros = [1, 2, 3, 4, 5];
const parte = numeros.slice(1, 4); // do índice 1 até 3 (4 não incluso)
console.log(parte); // [2, 3, 4]
console.log(numeros); // [1, 2, 3, 4, 5] (inalterado)

// join() - Junta elementos em string
const palavras = ['Olá', 'mundo', 'JavaScript'];
const frase = palavras.join(' ');
console.log(frase); // "Olá mundo JavaScript"

// toString() - Converte para string
const arr = [1, 2, 3];
console.log(arr.toString()); // "1,2,3"
```

---

### 3.10 Arrays Multidimensionais

```javascript
// Array de arrays (matriz)
const matriz = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
];

console.log(matriz[0][0]); // 1
console.log(matriz[1][2]); // 6

// Iterando sobre matriz
for (let i = 0; i < matriz.length; i++) {
  for (let j = 0; j < matriz[i].length; j++) {
    console.log(matriz[i][j]);
  }
}
```

---

## 📄 4. Structured Data (Dados Estruturados)

### Definição

**Structured Data** (dados estruturados) são dados organizados em um formato padronizado que permite que máquinas (como motores de busca) entendam o conteúdo. Em JavaScript, isso geralmente se refere a:

- **JSON** (JavaScript Object Notation)
- **Objetos** estruturados
- Dados organizados de forma hierárquica

### Uso por Motores de Busca

Dados estruturados são usados por motores de busca (como Google) para:
- Entender o conteúdo da página
- Coletar informações sobre a web
- Melhorar resultados de busca
- Exibir rich snippets (resultados enriquecidos)

São codificados usando marcação na página (como JSON-LD, Microdata, RDFa).

---

## 📋 5. JSON (JavaScript Object Notation)

### O que é JSON?

**JSON** (JavaScript Object Notation) é um formato de texto padrão para representar dados estruturados baseado na sintaxe de objetos JavaScript. É comumente usado para:

- Transmitir dados em aplicações web
- Enviar dados do servidor para o cliente
- Enviar dados do cliente para o servidor
- Armazenar dados de configuração
- APIs REST

### Características do JSON

- Formato de texto legível por humanos
- Baseado em pares chave-valor
- Suporta strings, números, booleanos, null, arrays e objetos
- **NÃO suporta** funções, undefined, Symbols, ou comentários
- Chaves sempre entre aspas duplas

---

### 5.1 Sintaxe JSON

```javascript
// Exemplo de JSON válido
const jsonExemplo = {
  "nome": "João",
  "idade": 30,
  "ativo": true,
  "hobbies": ["leitura", "programação"],
  "endereco": {
    "rua": "Rua das Flores",
    "numero": 123,
    "cidade": "São Paulo"
  },
  "telefone": null
};
```

### Regras do JSON

1. Chaves devem estar entre aspas duplas
2. Strings devem estar entre aspas duplas
3. Não pode ter vírgula no final
4. Não pode ter comentários
5. Não pode ter funções
6. Não pode ter undefined

```javascript
// ❌ JSON INVÁLIDO
const invalido1 = {
  nome: "João", // chave sem aspas
  idade: 30,
};

const invalido2 = {
  "nome": "João",
  "funcao": function() {} // funções não são permitidas
};

// ✅ JSON VÁLIDO
const valido = {
  "nome": "João",
  "idade": 30
};
```

---

### 5.2 JSON.stringify() - Converter para JSON

```javascript
const objeto = {
  nome: "Maria",
  idade: 25,
  ativo: true,
  hobbies: ["leitura", "programação"],
  endereco: {
    rua: "Rua das Flores",
    numero: 123
  }
};

// Converter objeto JavaScript para string JSON
const jsonString = JSON.stringify(objeto);
console.log(jsonString);
// {"nome":"Maria","idade":25,"ativo":true,"hobbies":["leitura","programação"],"endereco":{"rua":"Rua das Flores","numero":123}}

// Com formatação (indentação)
const jsonFormatado = JSON.stringify(objeto, null, 2);
console.log(jsonFormatado);
// {
//   "nome": "Maria",
//   "idade": 25,
//   "ativo": true,
//   ...
// }

// Filtrar propriedades
const jsonFiltrado = JSON.stringify(objeto, ['nome', 'idade'], 2);
console.log(jsonFiltrado);
// Apenas nome e idade serão incluídos
```

---

### 5.3 JSON.parse() - Converter de JSON

```javascript
// String JSON
const jsonString = '{"nome":"João","idade":30,"ativo":true}';

// Converter string JSON para objeto JavaScript
const objeto = JSON.parse(jsonString);
console.log(objeto);
// { nome: 'João', idade: 30, ativo: true }
console.log(objeto.nome); // "João"
console.log(objeto.idade); // 30

// JSON com array
const jsonArray = '[1, 2, 3, 4, 5]';
const array = JSON.parse(jsonArray);
console.log(array); // [1, 2, 3, 4, 5]

// Tratamento de erros
try {
  const jsonInvalido = '{nome: "João"}'; // JSON inválido
  const objeto = JSON.parse(jsonInvalido);
} catch (erro) {
  console.error('Erro ao fazer parse do JSON:', erro.message);
}
```

---

### 5.4 Casos de Uso do JSON

```javascript
// 1. Armazenar dados no localStorage
const dados = { nome: "João", idade: 30 };
localStorage.setItem('usuario', JSON.stringify(dados));
const dadosRecuperados = JSON.parse(localStorage.getItem('usuario'));

// 2. Enviar dados para servidor (fetch API)
const enviarDados = async () => {
  const dados = { nome: "Maria", email: "maria@email.com" };
  
  const resposta = await fetch('https://api.exemplo.com/usuarios', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(dados)
  });
  
  const resultado = await resposta.json();
  console.log(resultado);
};

// 3. Receber dados do servidor
const receberDados = async () => {
  const resposta = await fetch('https://api.exemplo.com/usuarios');
  const dados = await resposta.json(); // Converte automaticamente
  console.log(dados);
};

// 4. Configurações de aplicação
const config = {
  "apiUrl": "https://api.exemplo.com",
  "timeout": 5000,
  "retry": 3
};
const configString = JSON.stringify(config, null, 2);
// Salvar em arquivo config.json
```

---

## 🔄 6. Comparação entre Estruturas de Dados

### Quando Usar Cada Uma?

| Estrutura | Quando Usar | Características |
|-----------|-------------|-----------------|
| **Array** | Lista ordenada de elementos | Índices numéricos, ordenado, permite duplicatas |
| **Objeto** | Dados com propriedades nomeadas | Chaves string/Symbol, não ordenado (ES5), rápido acesso |
| **Map** | Chaves de qualquer tipo, ordem importa | Qualquer tipo de chave, mantém ordem, melhor para adicionar/remover |
| **Set** | Valores únicos | Sem duplicatas, verificação rápida de existência |
| **JSON** | Transmissão/armazenamento de dados | Formato de texto, interoperável, não executa código |

---

## 📝 Resumo da Aula

Nesta aula, você aprendeu:

✅ **Estruturas de Dados**: Formatos para organizar e armazenar dados eficientemente

✅ **Keyed Collections**:
- **Map**: Pares chave-valor com qualquer tipo de chave
- **Set**: Coleção de valores únicos

✅ **Indexed Collections**:
- **Arrays**: Coleções ordenadas por índice numérico

✅ **Structured Data**:
- Dados organizados em formato padronizado
- Usado por motores de busca

✅ **JSON**:
- Formato de texto para dados estruturados
- `JSON.stringify()` para converter objeto → JSON
- `JSON.parse()` para converter JSON → objeto

✅ **Métodos de Arrays**: push, pop, map, filter, reduce, find, etc.

---

## 🎯 Próximos Passos

Agora que você entende as estruturas de dados básicas, na próxima etapa você verá:
- Versão simplificada com analogias do dia a dia
- Exemplos práticos e visuais
- Comparações com situações reais

Continue para a **Aula Simplificada** para consolidar seu aprendizado! 🚀


