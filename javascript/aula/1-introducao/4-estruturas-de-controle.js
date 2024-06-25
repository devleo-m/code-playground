/*
Estruturas de Controle
Condicionais
JavaScript utiliza as instruções if, else if e else para realizar operações condicionais:
*/

let idade = 20;

if (idade < 18) {
    console.log("Menor de idade");
} else if (idade >= 18 && idade <= 60) {
    console.log("Adulto");
} else {
    console.log("Idoso");
}

/*
Laços de Repetição
Os laços de repetição (for, while e do...while) permitem executar um bloco de código várias vezes:
*/

// Usando for
for (let i = 0; i < 5; i++) {
    console.log("Repetição com for:", i);
}

// Usando while
let j = 0;
while (j < 5) {
    console.log("Repetição com while:", j);
    j++;
}

// Usando do...while
let k = 0;
do {
    console.log("Repetição com do...while:", k);
    k++;
} while (k < 5);
