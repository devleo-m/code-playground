/*
Introdução
Funções são blocos de código que podem ser reutilizados em diferentes partes do programa. 
Elas permitem modularizar o código, tornando-o mais organizado e fácil de manter. 
Nesta aula, vamos explorar a declaração de funções, funções anônimas e expressões de função, 
funções de seta (arrow functions), escopo de função e fechamentos (closures) e funções de ordem superior.

Declaração de Funções
A declaração de funções é a maneira mais comum de definir uma função em JavaScript. 
Ela consiste na palavra-chave function, seguida pelo nome da função, uma lista de parâmetros entre parênteses 
e o corpo da função entre chaves.
*/

function saudacao(nome) {
    return "Olá, " + nome + "!";
}

console.log(saudacao("João")); // Olá, João!

/*
Atividade 1:
1.1 Crie uma função chamada soma que receba dois parâmetros e retorne a soma deles.
1.2 Crie uma função chamada multiplicacao que receba dois parâmetros e retorne o produto deles.
1.3 Chame as funções e exiba os resultados no console.
*/

function soma(x,y) {
    return x + y
}
function multiplicacao(x,y) {
    return x * y
}
console.log(soma(1,1)) // 2
console.log(multiplicacao(2,2)) // 4

// Funções Anônimas e Expressões de Função
// Funções anônimas são funções sem nome. Elas são frequentemente usadas em expressões de função, 
//onde a função é atribuída a uma variável.

const ola_mundo = function() {
    return "Olá, mundo!";
};

console.log(ola_mundo()); // Olá, mundo!

/*
Atividade 2:
Crie uma expressão de função que calcule a diferença entre dois números e atribua-a a uma variável chamada subtracao.
Crie uma expressão de função que divida dois números e atribua-a a uma variável chamada divisao.
Chame as funções e exiba os resultados no console.
*/

const subtracao = function(x,y) {
    return x - y
}
const divisao = function(x,y) {
    return x / y
}
console.log(subtracao(5,3)) // 2
console.log(divisao(8,2)) // 4

/*
Funções de Seta (Arrow Functions)
Introduzidas no ES6, as funções de seta (arrow functions) fornecem uma sintaxe mais curta para escrever funções. 
Elas são particularmente úteis para funções de callback e funções de uma única linha.
*/

const sobrenome = (nome) => {
    return "Olá, " + nome + " Madeira";
};

console.log(sobrenome("Leo"));

// Função de uma única linha
const sobrenomeCurto = nome => "Olá, " + nome + " Madeira";
console.log(sobrenomeCurto("Leonardo"));

/*
Atividade 3:
Reescreva a função soma da Atividade 1 como uma arrow function.
Reescreva a função multiplicacao da Atividade 1 como uma arrow function.
Chame as funções e exiba os resultados no console.
*/

const soma2 = (x,y) => x+y
const multiplicacao2 = (x,y) => x*y
console.log(soma(1,1)) // 2
console.log(multiplicacao(2,2)) // 4

