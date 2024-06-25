/*
JavaScript fornece uma variedade de métodos poderosos para manipulação de arrays. 
Entre os mais utilizados estão map, reduce e filter, que permitem transformar, agregar 
e filtrar dados de maneira eficiente. Nesta aula, exploraremos esses métodos em detalhes, 
com exemplos práticos e atividades para reforçar o aprendizado.

Método map()
O método map cria um novo array com os resultados da aplicação de uma função a cada 
elemento do array original. É útil para transformar dados.
*/

//Sintaxe:
let novoArray = arrayOriginal.map(callback);

//Exemplo:
let numeros = [1, 2, 3, 4, 5];
let dobrados = numeros.map(numero => numero * 2);

console.log(dobrados); // [2, 4, 6, 8, 10]

/*
Método reduce()
O método reduce aplica uma função a um acumulador e a cada valor do array (da esquerda para a direita), 
reduzindo-o a um único valor. É útil para somar, concatenar ou agrupar dados.
*/

// Sintaxe:
let resultado = arrayOriginal.reduce(callback, valorInicial);

//Exemplo:
let numeros2 = [1, 2, 3, 4, 5];
let soma2 = numeros2.reduce((acumulador, numero) => acumulador + numero, 0);

console.log(soma2); // 15

// Método filter()
// O método filter cria um novo array com todos os elementos que passam no teste implementado pela função fornecida. 
// É útil para filtrar dados com base em uma condição.

//Sintaxe:
let novoArray3 = arrayOriginal.filter(callback);
//Exemplo:

let numeros3 = [1, 2, 3, 4, 5];
let pares3 = numeros3.filter(numero => numero % 2 === 0);

console.log(pares3); // [2, 4]

// Método forEach()
// O método forEach executa uma função para cada elemento do array. Diferente dos métodos anteriores, 
// forEach não retorna um novo array, apenas executa a função fornecida para cada elemento.

// Sintaxe:
arrayOriginal.forEach(callback);
// Exemplo:
let numeros4 = [1, 2, 3, 4, 5];
numeros4.forEach(numero => {
    console.log(numero * 2);
});

// Método find()
// O método find retorna o primeiro elemento do array que satisfaz a função de teste fornecida. 
// Se nenhum elemento satisfizer a função de teste, undefined é retornado.

// Sintaxe:
let elemento = arrayOriginal.find(callback);
// Exemplo:

let numeros5 = [1, 2, 3, 4, 5];
    let numeroPar = numeros5.find(numero => numero % 2 === 0);

console.log(numeroPar); // 2

// Método some()
// O método some testa se pelo menos um dos elementos no array passa no teste 
// implementado pela função fornecida. Retorna um valor booleano.

// Sintaxe:
let resultado6 = arrayOriginal.some(callback);
// Exemplo:
let numeros6 = [1, 2, 3, 4, 5];
let temPar = numeros6.some(numero => numero % 2 === 0);

console.log(temPar); // true

// Método every()
// O método every testa se todos os elementos do array passam no teste implementado 
// pela função fornecida. Retorna um valor booleano.

// Sintaxe:
let resultado7 = arrayOriginal.every(callback);
// Exemplo:

let numeros7 = [2, 4, 6, 8];
let todosPares2 = numeros.every(numero => numero % 2 === 0);

console.log(todosPares2); // true

