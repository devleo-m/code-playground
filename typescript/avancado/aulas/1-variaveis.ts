/*
TypeScript é uma linguagem de programação desenvolvida pela Microsoft que adiciona tipagem estática ao JavaScript. 
A tipagem estática ajuda a detectar erros em tempo de desenvolvimento e facilita a leitura e manutenção do código.

1. Variáveis em TypeScript
Declaração de Variáveis
Em TypeScript, as variáveis podem ser declaradas usando let, const, ou var. A diferença entre eles é:

let: Declara uma variável que pode ter seu valor alterado.
const: Declara uma constante, ou seja, uma variável cujo valor não pode ser alterado após a atribuição inicial.
var: Similar ao let, mas com escopo de função em vez de escopo de bloco.
*/

let variavelLet = "Pode ser alterado";
const variavelConst = "Não pode ser alterado";
var variavelVar = "Similar ao let, mas com escopo de função";

/*Tipagem de Variáveis
A tipagem pode ser explícita ou implícita. TypeScript infere o tipo quando não é explicitamente definido. */
let nome: string = "Leo" // tipagem explícita
let idade = 27 // tipagem implícita

//Tipos de Dados Básicos
// String
let mensagem2: string = "Hello, TypeScript!";

//Number (Integers e Floats)
let idade2: number = 27;
let altura2: number = 1.75;

//Boolean
let isCasado: boolean = true;
let isChovendo: boolean = false;

//Any
//O tipo any permite que uma variável contenha qualquer tipo de valor. 
//Deve ser usado com cuidado, pois desabilita a verificação de tipos.
let nome2: any = "Leo";
nome2 = 27;
let qualquerCoisa: any = "Olá";
qualquerCoisa = 42;
qualquerCoisa = true;

//Arrays
// Arrays podem ser tipados de duas formas: usando Array<tipo> ou tipo[].
let frutas2: string[] = ["Maca", "Banana", "Laranja"];
let frutas3: number[] = [1, 2, 3, 4, 5];
let frutas4: Array<string> = ["Uva", "Manga", "Goiaba"];

//Tuplas
let pessoa2: [string, number] = ["Leo", 27];

//Enums
enum Cor {Verde, Azul, Amarelo};
let minhaCor: Cor = Cor.Azul;

enum Cadastro {
    Ativo, 
    Inativo, 
    Bloqueado
};
let meuCadastro: Cadastro = Cadastro.Bloqueado;

//void
//Usado principalmente em funções que não retornam valor.

function mostrarMensagem(): void {
    console.log("Esta função não retorna nenhum valor.");
}

const mensagemOla = (): void => { console.log("Ola"); }

//Null e Undefined
let valorNulo2: null = null;
let valorIndefinido2: undefined = undefined;


