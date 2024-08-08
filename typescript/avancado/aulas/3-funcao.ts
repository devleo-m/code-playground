/*
Sintaxe de Funções
O que são Funções?
Funções são blocos de código que podem ser executados quando chamados. 
Eles podem aceitar parâmetros, realizar operações e retornar um valor.

Como declarar uma função
A sintaxe básica para declarar uma função em TypeScript é semelhante ao JavaScript, 
mas você pode (e deve) adicionar tipos de parâmetros e tipos de retorno.
*/

function multiplicacao(x:number, y:number): number{
    return x * y;
}

console.log(multiplicacao(5, 5)); // 25

/*
Neste exemplo:

a e b são os parâmetros da função, ambos do tipo number.
A função retorna um valor do tipo number.
O tipo de retorno é explicitamente declarado após os parâmetros.
*/

/*
Parâmetros Opcionais e Padrão
Parâmetros Opcionais
Em TypeScript, você pode declarar parâmetros opcionais, que podem ou não ser passados quando a função é chamada. 
Parâmetros opcionais são indicados com um ponto de interrogação (?) após o nome do parâmetro.
*/

function adcao(x:number, y?:number): number{
    if (y){
        return x + y;
    } else {
        return x + 10;
    }
}

console.log(adcao(5)); // 15
console.log(adcao(5, 5)); // 10

/*
Parâmetros Padrão
Em TypeScript, você pode declarar parâmetros padrão, que são valores que serão usados quando nenhum parâmetro é passado.
*/

function subtracao(x:number, y:number = 100): number{
    return x - y;
}   

console.log(subtracao(5)); // -95
console.log(subtracao(5, 5)); // 0

/**
Tipos de Retorno
Especificando Tipos de Retorno
Você pode (e deve) especificar o tipo de retorno de uma função para garantir que a função retorne sempre o 
tipo esperado. Se a função não retornar um valor, use o tipo void.
*/

function subtrair(a: number, b: number): number {
    return a - b;
}

function exibirMensagem(mensagem: string): void {
    console.log(mensagem);
}

// Funções que Retornam Objetos
// Você pode definir funções que retornam objetos complexos.

function criarPessoa(nome: string, sobrenome: string) {
    return {
        nome: nome,
        sobrenome: sobrenome
    }
}

const pessoa6 = criarPessoa("Leonardo", "Madeira");
console.log(pessoa6);

// Funções Anônimas e Arrow Functions
// Funções Anônimas
// Funções anônimas são funções que não têm nome e são geralmente usadas como argumentos para outras 
// funções ou atribuídas a variáveis.

let soma7 = function (a: number, b: number): number {
    return a + b;
};

console.log(soma7(2, 3)); // 5

// Arrow Functions
// Arrow functions são uma sintaxe mais curta para escrever funções anônimas. 
// Elas são particularmente úteis para funções simples e para manter o this do contexto correto.
let soma8 = (a: number, b: number): number => {
    return a + b;
}

// Se a função for simples e tiver apenas uma linha, você pode omitir as chaves e o `return`:
let multiplicar = (a: number, b: number): number => a * b;

console.log(soma8(2, 3)); // 5
console.log(multiplicar(2, 3)); // 6

/*
Exercício 1: Criar uma função de saudação com parâmetro opcional
Crie uma função saudacao que aceita um nome e uma saudação opcional. 
Se a saudação for fornecida, retorne uma saudação personalizada; caso contrário, retorne "Ola, tudo bem? Sr.[nome]".
*/

let saudacao = (nome:string, saudacao?:string): string => {
    return saudacao ? `${saudacao} ${nome}`: `Ola, tudo bem? Sr. ${nome}`;
}

// Exercício 3: Criar uma arrow function para multiplicação
// Escreva uma arrow function multiplicar que aceita dois números e retorna o produto deles.

let mult = (x:number, y:number): number => x * y;
console.log(mult(3, 6)); // 18