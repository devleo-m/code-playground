// Tipos Primitivos
// number
// Representa valores numéricos, sejam inteiros ou de ponto flutuante.

let age: number = 25;
let altura: number = 1.75;
// string
// Usado para representar texto.

let curso: string = "ts";
let mensagem: string = `Olá, ${curso}!`;

// boolean
// Representa valores true ou false.

let isTrue: boolean = true;
let isFalse: boolean = false;

// any
// O tipo any permite que uma variável aceite qualquer tipo de dado.

let variavelQualquer: any = 10;
variavelQualquer = "Agora é uma string";
variavelQualquer = true;

// Arrays
// Arrays são coleções de elementos do mesmo tipo.

let numeros: number[] = [1, 2, 3, 4, 5];
let frutas: Array<string> = ["maçã", "banana", "laranja"];

//any novamente
//em arrays
let anys: any[] = ["string", 10, true, false, 1.6];

// Tuplas
// Tuplas permitem expressar um array com um número fixo de elementos, com tipos definidos
// para cada posição.

let pessoa: [string, number] = ["João", 30];
// A posição 0 é uma string e a posição 1 é um número

// Enums
// Enums permitem criar um conjunto de constantes nomeadas.

enum DiaDaSemana {
    Domingo,
    Segunda,
    Terca,
    Quarta,
    Quinta,
    Sexta,
    Sabado
}

let dia: DiaDaSemana = DiaDaSemana.Segunda;
console.log(dia); // Isso imprimirá o valor numérico associado a Segunda (1)

// null e undefined
// null e undefined são tipos que representam ausência de valor.

let valorNulo: null = null;
let valorIndefinido: undefined = undefined;

// void
// É comumente utilizado para indicar que uma função não retorna nenhum valor.

function mostrarMensagem2(): void {
    console.log("Esta função não retorna nenhum valor.");
}
