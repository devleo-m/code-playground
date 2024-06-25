/*
Introdução
Arrays são uma estrutura de dados fundamental em JavaScript. Eles permitem armazenar múltiplos 
valores em uma única variável, facilitando o gerenciamento e a manipulação de conjuntos de dados. 
Nesta aula, exploraremos como criar, acessar e manipular arrays, além de utilizar métodos úteis 
disponíveis para arrays em JavaScript.

Criação de Arrays
Utilizando Colchetes
A maneira mais comum de criar um array em JavaScript é utilizando colchetes ([]).
*/

var frutas = ["maçã", "banana", "laranja"];
console.log(frutas); // ["maçã", "banana", "laranja"]

// Utilizando o Construtor Array
// Outra maneira de criar um array é utilizando o construtor Array.

let numeros = new Array(1, 2, 3, 4, 5);
console.log(numeros); // [1, 2, 3, 4, 5]

// Acessando Elementos de um Array
// Os elementos de um array são acessados utilizando índices. 
// O primeiro elemento tem índice 0, o segundo tem índice 1, e assim por diante.

let cores = ["vermelho", "verde", "azul"];
console.log(cores[0]); // vermelho
console.log(cores[2]); // azul

// Manipulação de Arrays
// Adicionando e Removendo Elementos
// Podemos adicionar e remover elementos de um array utilizando métodos como 
// push(), pop(), unshift() e shift().

var frutas = ["maçã", "banana", "goiaba", "acerola", "abacate", "manga"];

// Adicionando elementos
frutas.push("laranja"); // adiciona ao final
console.log(frutas); // ["maçã", "banana", "goiaba", "acerola", "abacate", "manga", "laranja"]

frutas.unshift("uva"); // adiciona ao início
console.log(frutas); // ["uva", "maçã", "banana", "goiaba", "acerola", "abacate", "manga", "laranja"]

// Removendo elementos
frutas.pop(); // remove do final
console.log(frutas); // ["uva", "maçã", "banana", "goiaba", "acerola", "abacate", "manga"]

frutas.shift(); // remove do início
console.log(frutas); // ["maçã", "banana", "goiaba", "acerola", "abacate", "manga"]

// Métodos Úteis para Arrays
// length
// O método length retorna o número de elementos em um array.

var frutas = ["Pera", "Uva", "acerola"];
console.log(frutas.length); // 3

// indexOf()
// O método indexOf retorna o índice da primeira ocorrência de um elemento no array. 
// Se o elemento não for encontrado, retorna -1.

var frutas = ["maçã", "banana", "laranja"];
console.log(frutas.indexOf("banana")); // 1
console.log(frutas.indexOf("uva")); // -1

// slice()
// O método slice retorna uma cópia de uma parte do array, sem modificar o array original.

var frutas = ["maçã", "banana", "laranja", "uva"];
var frutasSelecionadas = frutas.slice(1, 3);
console.log(frutasSelecionadas); // ["banana", "laranja"]

// forEach()
// O método forEach executa uma função em cada elemento do array.

let numeros2 = [1, 2, 3, 4, 5];
numeros2.forEach(numero => {
    console.log(numero * 2); // exibe o dobro de cada número
});











