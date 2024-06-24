/*
vamos explorar em detalhes dois conceitos fundamentais em JavaScript: variáveis e tipos de dados. 
Esses conceitos são a base para qualquer programação em JavaScript e entender como funcionam é 
essencial para escrever código eficiente e eficaz.

2.1 Variáveis
As variáveis são usadas para armazenar dados que podem ser manipulados e utilizados ao longo do 
código. Em JavaScript, temos três formas principais de declarar variáveis: var, let e const.

2.1.1 var
A palavra-chave var foi a maneira original de declarar variáveis em JavaScript. Variáveis 
declaradas com var têm escopo de função ou escopo global, se não estiverem dentro de uma função. 
Elas podem ser redeclaradas e têm uma característica chamada de "hoisting" (i.e., são movidas para 
o topo do contexto de execução durante a compilação).
*/

var nome = "Ana";
console.log(nome); // Ana

var nome = "Carlos";
console.log(nome); // Carlos

function exemploVar() {
    var idade = 30;
    if (true) {
        var idade = 25; // mesma variável dentro do escopo da função
        console.log(idade); // 25
    }
    console.log(idade); // 25
}

exemploVar();

/*
let
Introduzida no ES6 (ECMAScript 2015), a palavra-chave let permite declarar variáveis 
com escopo de bloco. Isso significa que a variável só está acessível dentro do bloco 
onde foi declarada. Variáveis let não podem ser redeclaradas no mesmo escopo.
*/

let nome2 = "Pedro";
console.log(nome2); // Pedro

nome2 = "Mariana";
console.log(nome2); // Mariana

function exemploLet() {
    let idade = 30;
    if (true) {
        let idade = 25; // nova variável dentro do bloco if
        console.log(idade); // 25
    }
    console.log(idade); // 30
}

exemploLet();

/*
const
Também introduzida no ES6, a palavra-chave const é usada para declarar constantes, 
ou seja, variáveis cujo valor não pode ser reatribuído após a inicialização. 
Como let, const também tem escopo de bloco.
*/

const nome3 = "Lucas";
console.log(nome3); // Lucas

// nome = "Gabriel"; // Erro: Assignment to constant variable.

const pessoa = { nome3: "Lucas", idade: 28 };
pessoa.idade = 29; // permitido alterar propriedades de objetos

console.log(pessoa.idade); // 29

/*
Atividade 2: Declaração de Variáveis
2.1 Crie uma variável chamada cidade usando var e atribua o nome de sua cidade a ela.
2.2 Crie uma variável chamada anoAtual usando let e atribua o ano atual a ela.
2.3 Crie uma constante chamada pais e atribua o nome de seu país a ela.
2.4 Exiba todas as variáveis no console.
*/
var cidade = "brasilia"
console.log(cidade)
let anoAtual = 2024
console.log(anoAtual)
const pais = "espanha"
console.log(pais)

/*
Tipos de Dados
JavaScript possui vários tipos de dados primitivos, além de tipos de objetos. 
Conhecer os tipos de dados é fundamental para manipular corretamente as 
informações no seu código.

Tipos Primitivos
* String: Sequência de caracteres. 
    Pode ser definida entre aspas simples ('), duplas ("), ou crases (`) para template literals.
* Number: Números, incluindo inteiros e de ponto flutuante.
* Boolean: Valores booleanos, verdadeiro (true) ou falso (false).
* Undefined: Quando uma variável é declarada mas não inicializada, ela tem o valor undefined.
* Null: Representa a ausência intencional de qualquer valor.
* Symbol: Introduzido no ES6, representa um valor único e imutável.
* BigInt: Permite representar números inteiros maiores do que o limite do tipo Number.
*/

let texto = "Olá, mundo!"; // String
let numero = 42; // Number
let decimal = 3.14; // Number
let isEstudante = true; // Boolean
let indefinido; // Undefined
let nulo = null; // Null
let simbolo = Symbol('descricao'); // Symbol
let grandeNumero = BigInt(9007199254740991); // BigInt

console.log(texto); // "Olá, mundo!"
console.log(numero); // 42
console.log(decimal); // 3.14
console.log(isEstudante); // true
console.log(indefinido); // undefined
console.log(nulo); // null
console.log(simbolo); // Symbol(descricao)
console.log(grandeNumero); // 9007199254740991n

/*

Tipos de Objetos
Object: Coleção de propriedades, onde cada propriedade é uma associação entre chave e valor.
Array: Lista ordenada de valores, onde cada valor pode ser acessado pelo seu índice.
Function: Funções são objetos em JavaScript e podem ser atribuídas a variáveis, passadas como 
argumentos e retornadas de outras funções.
*/

// Object
let pessoa2 = {
    nome: "Maria",
    idade: 30
};

console.log(pessoa2.nome); // Maria
console.log(pessoa2["idade"]); // 30

// Array
let lista = [1, 2, 3, 4, 5];
console.log(lista[0]); // 1
console.log(lista.length); // 5

// Function
function saudacao(nome) {
    return "Olá, " + nome + "!";
}

console.log(saudacao("João")); // Olá, João!

/*
Conversão de Tipos
Às vezes, é necessário converter entre diferentes tipos de dados. 
JavaScript fornece várias maneiras de fazer isso.
*/

let numero2 = "42";
let numeroConvertido = Number(numero2); // converte string para número
console.log(numeroConvertido); // 42

let booleano2 = true;
let booleanoConvertido = String(booleano2); // converte booleano para string
console.log(booleanoConvertido); // "true"

let texto2 = "123.45";
let textoConvertido = parseFloat(texto2); // converte string para número de ponto flutuante
console.log(textoConvertido); // 123.45

/*
Atividade 3: Conversão de Tipos
Converta uma string para um número usando Number().
Converta um número para uma string usando String().
Converta uma string para um número de ponto flutuante usando parseFloat().
Exiba os resultados no console.
*/

let name2 = "Leo"
let string_para_number = Number(name2)
console.log(string_para_number)
let number = 20
let number_para_string = String(number)
console.log(number_para_string)
let info = "text..."
let string_para_float = parseFloat(info)
console.log(string_para_float)

/**
 * Nesta aula, exploramos em profundidade variáveis e tipos de dados em JavaScript. 
 * Entendemos como declarar variáveis usando var, let e const, e exploramos os diferentes 
 * tipos de dados primitivos e objetos disponíveis na linguagem. Também vimos como realizar 
 * conversões entre diferentes tipos de dados. Com esses conhecimentos, você está pronto para 
 * manipular e utilizar dados de forma eficaz em seus programas JavaScript.
 */

