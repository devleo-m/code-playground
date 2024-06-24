/*
Aula 1: Fundamentos de JavaScript
Introdução ao JavaScript
JavaScript é uma das linguagens de programação mais populares no desenvolvimento web. 
Com ele, podemos criar desde interações simples em páginas web até aplicações complexas. 
Nesta aula, vamos explorar os fundamentos do JavaScript, sua sintaxe básica, estrutura e alguns 
exemplos práticos para você começar.

1.1 Sintaxe e Estrutura Básica
1.1.1 Declaração de Variáveis
No JavaScript, podemos declarar variáveis de várias maneiras, utilizando as palavras-chave 
var, let e const. 
Vamos entender cada uma delas:

var: Declara uma variável que pode ser redeclarada e cujo escopo é a função na qual está 
inserida ou o escopo global, se não estiver dentro de uma função.
let: Declara uma variável com escopo de bloco (entre {}), e não pode ser redeclarada no mesmo escopo.
const: Declara uma constante, ou seja, uma variável cujo valor não pode ser alterado após a 
sua atribuição inicial. Também tem escopo de bloco.
*/

// Usando var
var nome = "Leo";
console.log(nome); // Leo

// Usando let
let idade = 99;
console.log(idade); // 99

// Usando const
const pais = "Brasil";
console.log(pais); // Brasil

/*
Atividade 1:
1.1 Crie uma variável chamada cidade usando var e atribua o nome de sua cidade a ela.
1.2 Crie uma variável chamada ano usando let e atribua o ano atual a ela.
1.3 Crie uma constante chamada sobrenome e atribua o seu sobrenome a ela.
1.4 Exiba todas as variáveis no console.
*/

var cidade = "mcp"
let ano = 2024
const sobrenome = "storn"

console.log(`cidade: ${cidade}, ano: ${ano}, sobrenome: ${sobrenome}`)