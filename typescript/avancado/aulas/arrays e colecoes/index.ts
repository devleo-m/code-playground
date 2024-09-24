// 1. Métodos de Array
// Os arrays no TypeScript (e no JavaScript) vêm com uma série de métodos que facilitam o processamento de coleções de dados. Vamos explorar alguns dos métodos mais utilizados: map, filter, e reduce.

// map
// O método map cria um novo array com os resultados de uma função aplicada a cada elemento do array original.

// let novoArray = array.map(callback);
// Exemplo:
const numeros7 = [1, 2, 3, 4];
const dobrados = numeros.map(num => num * 2);
console.log(dobrados); // [2, 4, 6, 8]
// Aqui, dobrados é um novo array onde cada elemento foi multiplicado por 2.

// filter
// O método filter cria um novo array com todos os elementos que passam em um teste (função de callback).

// Sintaxe:
const numeros2 = [1, 2, 3, 4, 5];
const pares = numeros.filter(num => num % 2 === 0);
console.log(pares); // [2, 4]

// reduce
// O método reduce aplica uma função acumuladora contra um acumulador e cada elemento do array (da esquerda para a direita) para reduzi-lo a um único valor.
const numeros8 = [1, 2, 3, 4];
const soma4 = numeros.reduce((acumulador, valorAtual) => acumulador + valorAtual, 0);
console.log(soma); // 10

//forEach
// O método forEach executa uma função para cada elemento do array.

// Sintaxe:
const numeros9 = [1, 2, 3, 4];
numeros9.forEach(num => console.log(num));

//some
// O método some verifica se pelo menos um elemento do array passa em um teste (função de callback).

// Sintaxe:
const numeros5 = [1, 2, 3, 4, 5];
const temPares = numeros5.some(num => num % 2 === 0);
console.log(temPares); // true

/*
2. Sets e Maps
Além de arrays, o TypeScript também oferece Sets e Maps como estruturas de dados para armazenar coleções.

Sets
Um Set é uma coleção de valores únicos. Um valor no Set pode ocorrer apenas uma vez; ele é usado quando 
você precisa armazenar itens sem duplicatas.
*/

const meuSet = new Set<number>();

meuSet.add(1);
meuSet.add(2);
meuSet.add(2); // Ignorado, pois "2" já está no Set

console.log(meuSet); // Set { 1, 2 }

/*
Métodos Comuns:
add(value): Adiciona um valor ao Set.
has(value): Verifica se um valor está presente no Set.
delete(value): Remove um valor do Set.
clear(): Remove todos os valores do Set.
size(): Retorna o tamanho do Set.
*/

//Maps
// Um Map é uma coleção de pares chave-valor. Ao contrário de objetos, as chaves podem ser de qualquer tipo, 
// incluindo funções, objetos, ou qualquer outro tipo.
const meuMapa = new Map<string, number>();

meuMapa.set("um", 1);
meuMapa.set("dois", 2);

console.log(meuMapa.get("um")); // 1
console.log(meuMapa.has("dois")); // true

meuMapa.delete("um");

console.log(meuMapa); // Map { 'dois' => 2 }

/**
Métodos Comuns:
set(chave, valor): Adiciona ou atualiza um par chave-valor no Map.
get(chave): Retorna o valor associado à chave.
has(chave): Verifica se uma chave existe no Map.
delete(chave): Remove um par chave-valor do Map.
clear(): Remove todos os pares chave-valor do Map.
size(): Retorna o tamanho do Map.
 */





