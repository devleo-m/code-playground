/*
Arrays
O que são Arrays?
Arrays são listas ordenadas de elementos, que podem ser de qualquer tipo. Em TypeScript, 
você pode definir arrays que contenham elementos de um tipo específico.

Como declarar Arrays
Sintaxe com colchetes ([]):
 */

let numeros3: number[] = [1, 2, 3, 4, 5];
let palavras: string[] = ["TypeScript", "JavaScript", "Python"];

//Sintaxe com Array<tipo>:

let numeros4: Array<number> = [1, 2, 3, 4, 5];
let palavras2: Array<string> = ["TypeScript", "JavaScript", "Python"];

/*
Métodos Comuns de Arrays
push(): Adiciona um ou mais elementos ao final do array.
pop(): Remove o último elemento do array.
shift(): Remove o primeiro elemento do array.
unshift(): Adiciona um ou mais elementos no início do array.
map(): Cria um novo array com os resultados da chamada de uma função para cada elemento.
filter(): Cria um novo array com todos os elementos que passaram no teste implementado pela função fornecida.
reduce(): Aplica uma função a um acumulador e a cada elemento do array (da esquerda para a direita) para reduzi-lo a um único valor.
*/

let numbers: number[] = [1, 2, 3, 4, 5];

numbers.push(6); // [1, 2, 3, 4, 5, 6]
numbers.pop(); // [1, 2, 3, 4, 5]

let numerosPares = numbers.filter(num => num % 2 === 0); // [2, 4]
let soma = numeros.reduce((acumulador, valorAtual) => acumulador + valorAtual, 0); // 15

/*
Tuples
O que são Tuples?
Tuples são uma estrutura de dados que permite armazenar múltiplos valores de diferentes tipos em um array fixo. 
São úteis quando você deseja associar múltiplos valores de tipos diferentes.

Como declarar Tuples
*/

let pessoa4: [string, number] = ["Alice", 25];

// Acessando elementos de uma Tuple
// Você pode acessar elementos de uma tuple usando índices, assim como em arrays.
let pessoa5: [string, number] = ["Alice", 25];
let nome5: string = pessoa[0]; // "Alice"
let idade5: number = pessoa[1]; // 25

//Alterando valores em uma Tuple
pessoa5[1] = 26; // { nome: "Alice", idade: 26 }

// Enums
// O que são Enums?
// Enums são um tipo especial em TypeScript que permite definir um conjunto de constantes nomeadas, tornando o código mais legível e fácil de manter.

// Como declarar Enums
// Enums numéricos:

enum Cor2 {
    Vermelho,
    Verde,
    Azul
}

let minhaCor2: Cor = Cor.Verde; // 1

//Por padrão, enums numéricos começam com 0, mas você pode definir um valor inicial diferente.

enum Cor3 {
    Vermelho = 1,
    Verde,
    Azul
}

let minhaCor3: Cor = Cor.Verde; // 2

//Enums com strings:
enum Direcao {
    Norte = "NORTE",
    Sul = "SUL",
    Leste = "LESTE",
    Oeste = "OESTE"
}

let minhaDirecao: Direcao = Direcao.Norte; // "NORTE"

let livro: [string, number, Cor3] = ["Harry Potter", 10, Cor3.Azul]; 