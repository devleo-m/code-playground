/*
O que é um Objeto?
Um objeto é uma coleção de propriedades, onde cada propriedade é composta por uma chave 
(que é uma string ou símbolo) e um valor (que pode ser qualquer tipo de dado, 
como strings, números, arrays, outros objetos, etc.).

Como Declarar um Objeto
Aqui está um exemplo simples de como você pode declarar um objeto em TypeScript:
*/
let cliente = {
    nome: "Alice",
    idade: 30,
    cidade: "São Paulo"
};

/*
Neste exemplo:

cliente é um objeto que tem três propriedades: nome, idade, e cidade.
Cada uma dessas propriedades tem um valor associado a ela.
Acessando Propriedades de um Objeto
Você pode acessar as propriedades de um objeto usando a notação de ponto (.) ou a notação de colchetes ([]).
*/

console.log(cliente.nome); // "Alice"
console.log(cliente["idade"]); // 30
//console.log(cliente.cidade); // "São Paulo"

/*
Modificando Propriedades de um Objeto
Você também pode modificar o valor de uma propriedade existente ou adicionar novas propriedades ao objeto.
*/

cliente.idade = 31;
cliente.cidade = "Campinas";

//Uso de Interfaces para Tipagem de Objetos
// O que é uma Interface?
// Uma interface em TypeScript é uma maneira de definir um contrato para a estrutura de um objeto. 
//Ela descreve quais propriedades e métodos um objeto deve ter e seus tipos. Isso ajuda a garantir que os objetos sigam um formato específico e facilita o trabalho em equipe, já que todos sabem a estrutura exata que os objetos devem ter.

// Definindo uma Interface
// Aqui está um exemplo de como definir uma interface:

interface Cachorro {
    nome: string;
    pelagem: string;
    idade: number;
    imortal: boolean;

    som (): string;
}


//Usando a Interface
// Agora, vamos criar um objeto usando essa interface:

let cachorro2: Cachorro = {
    nome: "Simba",
    pelagem: "preta",
    idade: 3,
    imortal: false,
    som: (): string =>{
        return "Au Au";
    }
}

cachorro2.som();

/*
Interfaces com Propriedades Opcionais
Propriedades Opcionais em Interfaces
Nem todas as propriedades de um objeto precisam ser obrigatórias. 
Em TypeScript, você pode definir propriedades opcionais em uma interface usando o símbolo ? 
após o nome da propriedade.
*/

interface Leitor {
    nome: string;
    email: string;
    idade?: number;
}

let leitor: Leitor = {
    nome: "Alice",
    email: "uXqFP@example.com"
    //idade: 30
}

/*
Exercício 3: Usar uma Interface com Métodos
Defina uma interface chamada ContaBancaria que tenha as seguintes propriedades: 
titular (string), saldo (number), e um método depositar que receba um valor como parâmetro e adicione ao saldo. 
Crie um objeto que implemente essa interface e use o método depositar.
 */

interface ContaBancaria{
    titular: string;
    saldo: number;
    depositar: (valor: number) => void;
}

let conta_bancaria: ContaBancaria = {
    titular: "Beltrano de Tal",
    saldo: 1000,
    depositar: (valor: number): void =>{
        this.saldo += valor
        console.log(`O valor de ${valor} foi depositado na sua conta.`)
        console.log(`Seu novo saldo é de ${this.saldo}.`)
    }
}

