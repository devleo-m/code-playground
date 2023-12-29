//Funções
function saudacao(name:string): string{
    return `Olah ${name}!`;
}
console.log(saudacao("Arolnode"))

// Parâmetros Opcionais e Padrão
function bomDia(name?: string): string{
    if(name){
        return `Bom dia! ${name}.`
    }else{
        return "Bom dia!"
    }
}
console.log(bomDia("Fulano"))

//Parâmetros Padrão
function visitante(name: string): string{
    return `Seja bem vindo ${name}`
}
console.log(visitante("Leo"))

//Sobrecarga de Funções
//Permite definir múltiplas assinaturas para uma função com diferentes tipos de parâmetros 
//e retorno.

function juntar(a: number, b: number): number; // Assinatura 1
function juntar(a: string, b: string): string; // Assinatura 2

function juntar(a: any, b: any): any {
    return a + b;
}

let resultadoNumero: number = juntar(1, 2); // Saída: 3
let resultadoString: string = juntar("Olá, ", "Mundo!"); // Saída: Olá, Mundo!

//Funções de Setas (Arrow Functions)
let soma = (a: number, b:number): number =>{
    return a + b
}

let multip = (a: number, b:number): number => a+b

/*
Boas Práticas
Nomes descritivos: Escolha nomes significativos para suas funções que descrevam claramente 
o que elas fazem.
Tipagem forte: Sempre que possível, defina os tipos de parâmetros e retorno para evitar erros.
Evitar side effects: Procure manter as funções puras, ou seja, que não causem efeitos 
colaterais, o que ajuda na previsibilidade e no teste.
Reusabilidade: Promova a reutilização de código através de funções modulares.
Uso moderado de parâmetros opcionais e padrão: Eles podem aumentar a flexibilidade, mas em 
excesso podem tornar o código menos legível.
 */