/*
Generics
O que são Generics?
Generics são uma maneira de criar componentes que podem funcionar com diferentes tipos de dados, 
ao mesmo tempo em que fornecem forte tipagem. Em outras palavras, generics permitem que você escreva 
código que é mais flexível e reutilizável, mas sem perder a segurança dos tipos.

Por exemplo, ao invés de criar uma função que só funciona com números ou uma função que só funciona c
om strings, você pode criar uma função genérica que funcione com qualquer tipo de dado.

Por que usar Generics?
Reutilização de Código: Você pode criar funções ou classes que funcionam com múltiplos tipos de dados.
Segurança de Tipos: Ao contrário de usar o tipo any, os generics mantêm a segurança de tipos, permitindo 
que o TypeScript infira e valide os tipos usados.
Flexibilidade: Você pode definir restrições para os tipos, garantindo que apenas certos tipos sejam permitidos.
Sintaxe Básica de Generics
A sintaxe básica de generics usa colchetes angulares (<>) para definir um tipo genérico.
*/

function identityExemple<T>(arg: T): T {
    return arg;
}

/*
Nesse exemplo:

T é um parâmetro de tipo que atua como um substituto para o tipo real.
arg: T define que a função identity aceita um argumento de tipo T.
: T especifica que o tipo de retorno da função é T.
Como Usar Generics em Funções
Vamos começar com funções genéricas.

Função Genérica Simples:
*/

function identity<T>(arg: T): T {
    return arg;
}

let output1 = identity<string>("Hello, TypeScript!");  // T é string
let output2 = identity<number>(100);                   // T é number

/*
Aqui, identity é uma função genérica que pode trabalhar com qualquer tipo de dado. Ao chamar a função, 
você pode especificar o tipo (por exemplo, <string> ou <number>), mas o TypeScript frequentemente pode 
inferir o tipo automaticamente.
*/

//Função Genérica com Arrays:
function logArray<T>(arg: T[]): void {
    console.log(arg.length);  // Podemos usar a propriedade 'length' pois sabemos que é um array.
    console.log(arg);
}

logArray<number>([1, 2, 3, 4]); // T é inferido como number
logArray<string>(["a", "b", "c"]); // T é inferido como string
//Nesse exemplo, T[] representa um array de elementos do tipo T.

/*
Como Usar Generics em Classes
Generics também podem ser usados em classes, permitindo que você crie classes que funcionam com 
diferentes tipos de dados.

Classe Genérica Simples:
*/

class Box<T> {
    private _content: T;

    constructor(content: T) {
        this._content = content;
    }

    public getContent(): T {
        return this._content;
    }
}

let stringBox = new Box<string>("Hello");
console.log(stringBox.getContent()); // Saída: Hello

let numberBox = new Box<number>(100);
console.log(numberBox.getContent()); // Saída: 100
//Aqui, a classe Box é genérica e pode ser usada para armazenar qualquer tipo de dado.

interface Identifiable {
    id: number;
}

class IdentifiableBox<T extends Identifiable> {
    private _content: T;

    constructor(content: T) {
        this._content = content;
    }

    public getContent(): T {
        return this._content;
    }

    public getId(): number {
        return this._content.id;
    }
}

let identifiableItem = { id: 42, name: "Item" };
let box = new IdentifiableBox(identifiableItem);
console.log(box.getId()); // Saída: 42

/**
Aqui, IdentifiableBox só pode ser instanciado com tipos que implementam a interface Identifiable.

Recapitulando:
 * Generics permitem criar código reutilizável e flexível que ainda mantém a segurança de tipos.
 * Funções genéricas são funções que podem trabalhar com qualquer tipo, e os tipos podem ser 
especificados ou inferidos pelo TypeScript.
 * Classes genéricas funcionam de forma similar às funções genéricas, mas no contexto de classes, 
permitindo que você crie objetos que funcionam com diferentes tipos de dados.
 * Restrições de tipo permitem limitar os tipos que podem ser usados com generics, garantindo que 
certas propriedades ou métodos estejam disponíveis.

Generics são uma poderosa ferramenta no TypeScript, permitindo que você escreva código mais 
abstrato e flexível sem sacrificar a segurança de tipos que torna o TypeScript tão valioso.
 */