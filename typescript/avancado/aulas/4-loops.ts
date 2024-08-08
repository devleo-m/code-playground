// 1. for Loop
// O que é?
// O for loop é o loop mais comum e flexível em TypeScript. Ele é ideal para quando você sabe de 
// antemão quantas vezes deseja repetir um bloco de código.

// Sintaxe

for (let i = 0; i < 5; i++) {
    console.log(`O valor de i é: ${i}`);
}

// 2. while Loop
// O que é?
// O while loop é útil quando você deseja repetir um bloco de código enquanto uma condição for verdadeira. 
// Diferente do for, você não precisa saber de antemão o número de iterações.

// Sintaxe

let i = 0;
while (i < 5) {
    console.log(`O valor de i é: ${i}`);
    i++;
}

// 3. do...while Loop
// O que é?
// O do...while loop é semelhante ao while, mas com uma diferença fundamental: o bloco de código é executado pelo menos uma vez, independentemente da condição. A condição é verificada após a execução do bloco.

// Sintaxe

//let i = 0;
do {
    console.log(`O valor de i é: ${i}`);
    i++;
} while (i < 5);

// 4. for...of Loop
// O que é?
// O for...of loop é utilizado para iterar sobre elementos de coleções iteráveis, como arrays, strings, ou objetos que implementam o protocolo iterável. Este loop é útil quando você deseja acessar diretamente os valores dos elementos.

// Sintaxe

let number: number[] = [1, 2, 3, 4, 5];
for (const num of number) {
    console.log(`Número: ${num}`);
}

// 5. for...in Loop
// O que é?
// O for...in loop é usado para iterar sobre as propriedades enumeráveis de um objeto ou elementos de um array. Ele retorna o índice (ou chave) e não o valor.

// Sintaxe
let pessoa9 = { nome: "Alice", idade: 25, cidade: "São Paulo" };
for (const chave in pessoa9) {
    console.log(`${chave}: ${pessoa9[chave as keyof typeof pessoa9]}`);
}

// Diferença entre for...of e for...in
// for...of: Itera sobre os valores dos elementos.
// for...in: Itera sobre as chaves (ou índices) dos elementos.
// Exemplos Práticos
// Exemplo 1: Usando for Loop para Iterar sobre um Array

let frutas9: string[] = ["Maçã", "Banana", "Manga", "Laranja"];
for (let i = 0; i < frutas.length; i++) {
    console.log(frutas[i]);
}

//Exemplo 2: Usando while Loop para Contagem Regressiva
let contador: number = 5;
while (contador > 0) {
    console.log(contador);
    contador--;
}

//Exemplo 3: Usando do...while para Repetir até Condição ser Falsa
let numero: number = 0;
do {
    console.log(`Número: ${numero}`);
    numero++;
} while (numero < 3);

//Exemplo 4: Usando for...of para Iterar sobre uma String
let palavra: string = "TypeScript";
for (const letra of palavra) {
    console.log(letra);
}

//Exemplo 5: Usando for...in para Iterar sobre as Propriedades de um Objeto
let carro = { marca: "Toyota", modelo: "Corolla", ano: 2020 };
for (const chave in carro) {
    console.log(`${chave}: ${carro[chave as keyof typeof carro]}`);
}

