/*
Interfaces
As interfaces são contratos que definem a estrutura de um objeto. Elas podem conter 
propriedades, métodos e podem ser implementadas por classes ou objetos.
*/

//Declaração de Interfaces

interface Humano{
    nome: string;
    idade: number;
    saudacao(): string;
}

//Utilizando Interfaces
let saudar = (pessoa: Humano): void => {
    console.log(`Ola ${pessoa.nome} vc tem ${pessoa.idade} anos`)
}
function saudar2(ser: Humano): string{
    return `Oi ${ser.nome}! ${ser.saudacao}`
}
let usuario: Humano = {
    nome: "Fulano",
    idade: 90,
    saudacao(){
        return"Ola. teste ... texto aleatorio"
    }
}

saudar(usuario);
saudar2(usuario);

//Tipos
//Os tipos permitem criar um nome para um tipo específico, facilitando a reutilização e a legibilidade.

//Declaração de Tipos
type Coordenada = [number, number];
let ponto: Coordenada = [10345.9, 12370]

//Tipos Genéricos
// Os tipos genéricos permitem criar componentes reutilizáveis que funcionam com vários tipos.
interface Caixa<Tipo>{
    valor: Tipo;
}
let caixaDeNumero: Caixa<number> = { valor: 10 };
let caixaDeString: Caixa<string> = { valor: "Oi"};

//Extensão de Interfaces
// As interfaces podem ser estendidas por outras interfaces, adicionando novas propriedades ou métodos.
interface Animal{
    raca: string;
}
interface Mamifero extends Animal{
    mamifero: boolean;
}
let cachorro: Mamifero = {
    raca: "Shitzu",
    mamifero: true
}
