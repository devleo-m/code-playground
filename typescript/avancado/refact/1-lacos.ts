{
/*
1. Laço for
O laço for é utilizado quando sabemos quantas vezes precisamos repetir um bloco de código. 
Ele é composto por três partes:

Inicialização: define uma variável de controle.
Condição: a condição que deve ser verdadeira para o laço continuar.
Incremento/Decremento: altera a variável de controle após cada iteração.
*/
for(let i = 0; i < 5; i++) {
    console.log(`for - Número: ${i}`);
}

/*
2. Laço while
O laço while continua executando enquanto a condição for verdadeira. Ele é útil quando não 
sabemos quantas iterações serão necessárias, mas sabemos a condição que deve ser atendida para continuar.
*/
let i = 0;
while (i < 5) {
    console.log(`While - Número: ${i}`);
    i++;
}

/*
Laço do...while
O laço do...while é semelhante ao while, mas garante que o bloco de código será executado 
pelo menos uma vez, pois a condição é verificada após a execução do bloco.
*/
let x = 0;
do {
    console.log(`Do/while - Número: ${x}`);
    x++;
} while (x < 5);

/*
4. Laço for...in
O laço for...in é utilizado para iterar sobre as propriedades enumeráveis de um objeto. 
É útil quando queremos acessar as chaves de um objeto.
*/
const pessoas = [
    {
        nome: "Leonardo",
        idade: 28,
        cidade: "Florianópolis"
    },
    {
        nome: "Gabriel",
        idade: 28,
        cidade: "Florianópolis"
    },
    {
        nome: "Lucas",
        idade: 28,
        cidade: "Florianópolis"
    }
];

for (const key in pessoas) {
    console.log(`Pessoa: ${key} | ${pessoas[key].nome} | ${pessoas[key].idade} | ${pessoas[key].cidade}`);
}

/*
5. Laço for...of
O laço for...of é utilizado para iterar sobre objetos iteráveis, como arrays. É uma forma simplificada e mais 
legível de iterar sobre os elementos de um array ou outros objetos iteráveis.
*/
const frutas = ["maçã", "banana", "laranja"];

for (const fruta of frutas) {
    console.log(`Fruta: ${fruta}`);
}

for (const pessoa of pessoas) {
    console.log(`Pessoa: ${pessoa.nome}`)
}

}
