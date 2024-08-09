/*
Tipagem Avançada no TypeScript
No TypeScript, a tipagem avançada permite criar tipos mais complexos e flexíveis, 
permitindo melhor expressividade e segurança no código. Vamos explorar cada um dos conceitos de Tipagem Avançada: 
Union Types, Intersection Types, Literal Types, Type Aliases, e Conditional Types.
*/

/*
1. Union Types
Union Types permitem que uma variável ou função aceite mais de um tipo de dado. 
É como dizer "essa variável pode ser desse tipo ou daquele tipo".
*/

let value: string | number;

//Aqui, value pode ser do tipo string ou number.

function printId(id: string | number) {
    console.log(`Your ID is: ${id}`);
}

printId(101);    // Funciona
printId("ABC");  // Funciona

/*
Como funciona:
O TypeScript restringe o uso de métodos/propriedades disponíveis apenas em tipos comuns aos union types.
Você pode usar narrowing para distinguir entre os tipos, usando typeof ou in para realizar uma verificação de tipo.
*/

/**
 2. Intersection Types
Intersection Types combinam múltiplos tipos em um único tipo, exigindo que uma variável satisfaça 
todos os tipos combinados.
 */

type T1 = { name: string };
type T2 = { age: number };

type Person = T1 & T2;

//Aqui, Person é um tipo que exige tanto uma propriedade name (do tipo string) 
//quanto uma propriedade age (do tipo number).

type Drivable = {
    drive: () => void;
};

type Electric = {
    charge: () => void;
};

type ElectricCar = Drivable & Electric;

const myCar: ElectricCar = {
    drive: () => console.log("Driving..."),
    charge: () => console.log("Charging...")
};

myCar.drive();
myCar.charge();

//3. Literal Types
// Literal Types permitem restringir uma variável a um valor exato, ao invés de um tipo genérico como string ou number.
let direction: "left" | "right";
// Aqui, direction só pode ser "left" ou "right".

type CardinalDirection = "North" | "South" | "East" | "West";

function move(direction: CardinalDirection) {
    console.log(`Moving ${direction}`);
}

move("North");  // Funciona
move("West");   // Funciona
// move("Up");  // Erro: "Up" não é um valor permitido

/*
Como funciona:
Literal Types são usados para criar APIs mais previsíveis e seguras, forçando o uso de valores pré-definidos.
Você pode combinar Literal Types com Union Types para permitir uma gama de valores exatos.
*/


//4. Type Aliases
// Type Aliases permitem criar um nome personalizado para um tipo. Isso é útil para tornar o código mais legível 
//e para evitar repetição de tipos complexos.
type MyString = string;
// Aqui, MyString é um alias para string.

type UserID = string | number;

function getUser(id: UserID) {
    console.log(`User ID is ${id}`);
}

getUser(101);
getUser("ABC123");










