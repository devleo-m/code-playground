/*
Vamos mergulhar na sintaxe básica do TypeScript, cobrindo variáveis, tipos de dados, 
operadores, estruturas de controle e loops, além de algumas boas práticas ao longo do caminho.

Variáveis e Tipos de Dados
Em TypeScript, podemos declarar variáveis usando as palavras-chave let ou const.

let: Usado para declarar variáveis que podem ser alteradas.
const: Usado para declarar constantes imutáveis.
*/

let idade: number = 90;
let nome: string = "Leonardo";
let casado: boolean = true

/*
Os tipos de dados incluem number, string, boolean, array, object, null, undefined, 
tuple, enum, entre outros.
*/

/*
Operadores
Os operadores em TypeScript são semelhantes aos de outras linguagens e incluem aritméticos 
(+, -, *, /)
operadores de comparação (==, !=, ===, !==, <, >, <=, >=)
operadores lógicos (&&, ||, !) e 
operadores de atribuição (=, +=, -=, *=, /=).
*/

//Estruturas de Controle
//if/else

if (idade >= 18) { //idade 90 condicao true
    console.log("Maior de idade");
} else{
    console.log("Menor de idade");
}

//switch

let diaDaSemana: string = "Segunda";
switch (diaDaSemana) {
    case "Segunda":
        console.log(`${diaDaSemana}-feira, dia de trabalho!`)
        break;
    case "Terca":
        console.log(`${diaDaSemana}-feira, dia de trabalho!`)
        break;
    case "Quarta":
        console.log(`${diaDaSemana}-feira, dia de trabalho!`)
        break;
    case "Quinta":
        console.log(`${diaDaSemana}-feira, dia de trabalho!`)
        break;
    case "Sexta":
        console.log(`${diaDaSemana}-feira, Finalmente. SEXTOU!!!`)
        break;
    case "Sabado":
        console.log(`${diaDaSemana}! FOLGA`)
        break;
    case "Domingo":
        console.log(`${diaDaSemana} DIA DE PRAIA!`)
        break;

    default:
        console.log("Dia da semana invalido. Escreva corretamente!")
        break;
}

//Loops
//for

for (let i: number = 0; i < 5; i++) {
    console.log(i);
}

//while
let contador: number = 0;
while (contador < 5) {
    console.log(contador);
    contador++;
}
/**
Boas Práticas:::

Nomenclatura significativa: Use nomes de variáveis claros e descritivos.

Comentários: Escreva comentários explicativos para tornar o código mais compreensível.

Evite variáveis globais: Prefira escopo local sempre que possível para evitar colisões 
e efeitos colaterais indesejados.

Modularização: Divida seu código em funções e módulos para facilitar a manutenção.

Evite repetição de código: Use funções ou loops para evitar repetições desnecessárias.

Código legível: Mantenha um estilo consistente, usando indentação e formatação clara.

O TypeScript é poderoso por sua tipagem estática, que ajuda a detectar erros no código 
durante a fase de desenvolvimento.

tmj :P
 */