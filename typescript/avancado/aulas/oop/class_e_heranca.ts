/*
Classes e Herança em TypeScript
1. Declaração de Classes
Em TypeScript, uma classe é uma estrutura que nos permite criar objetos com propriedades e métodos. 
Ela pode ser vista como um "molde" para objetos. Aqui está a sintaxe básica para declarar uma classe:
*/
class Pessoa{
    nome: string
    idade: number
    constructor(nome: string, idade: number){
        this.nome = nome
        this.idade = idade
    }

    toString(): string {
        return `Nome Completo: ${this.nome}, Idade da Pessoa: ${this.idade}`;
    }

    cumprimentar(): void {
        console.log(`Olá, meu nome é ${this.nome} e tenho ${this.idade} anos.`);
    }
}

const pessoa = new Pessoa("Leo", 27);
pessoa.cumprimentar(); // "Olá, meu nome é Leo e tenho 27 anos."
console.log(pessoa.toString()); // "Nome Completo: Leo, Idade da Pessoa: 27"

//*******************************************************************************************************************************/

/*
Herança e Superclasses
Herança é um conceito fundamental da programação orientada a objetos que permite que uma classe 
"herde" propriedades e métodos de outra classe. A classe que herda é chamada de subclasse ou classe derivada, 
enquanto a classe da qual herda é chamada de superclasse ou classe base.
Exemplo de Herança:
*/
class Funcionario extends Pessoa{
    cargo: string

    constructor(nome: string, idade: number, cargo: string){
        super(nome, idade) // Chama o construtor da classe base (Pessoa)
        this.cargo = cargo
    }

    apresentar(): void {
        console.log(`Olá, meu nome é ${this.nome}, tenho ${this.idade} anos e sou ${this.cargo}.`);
    }
}

const funcionario = new Funcionario("Joaquim", 28, "Engenheiro");
funcionario.apresentar(); // "Olá, meu nome é Joaquim, tenho 28 anos e sou Engenheiro."

/*
Neste exemplo:
Funcionario é uma subclasse que estende Pessoa.
super(nome, idade) chama o construtor da classe Pessoa para inicializar nome e idade.
Funcionario adiciona a propriedade cargo e o método apresentar.
*/

/*
Modificadores de Acesso (public, private, protected)
Modificadores de acesso controlam a visibilidade de propriedades e métodos dentro e fora da classe.

public: O padrão. Propriedades e métodos públicos podem ser acessados de qualquer lugar.
private: Propriedades e métodos privados só podem ser acessados dentro da própria classe.
protected: Propriedades e métodos protegidos podem ser acessados dentro da classe e por subclasses.
*/

class ContaBancaria2 {
    private _saldo: number;

    constructor(saldo: number) {
        this._saldo = saldo;
    }

    toString(): string {
        return `Saldo da Conta: ${this._saldo}`;
    }

    get saldoTotal(): number {
        return this._saldo;
    }

    set saldoTotal(saldo: number) {
        if (saldo >= 0) {
            this._saldo = saldo;
        } else {
            console.log("O saldo não pode ser negativo.");
        }
    }

    depositarSaldo(saldo: number): void {
        if (saldo > 0) {
            this._saldo += saldo;
            console.log(`Depósito de R$${saldo}. Saldo atual: R$${this._saldo}.`);
        } else {
            console.log("O saldo não pode ser negativo.");
        }
    }

    sacarDinheiro(valor: number): void {
        if (valor <= this._saldo && valor > 0 && this._saldo > 0) {
            this._saldo -= valor;
            console.log(`Saque de R$${valor}. Saldo atual: R$${this._saldo}.`);
        } else {
            console.log("O valor informado é inválido ou o saldo é insuficiente.");
        }
    }
}

const conta_bancaria_nova = new ContaBancaria2(100);
console.log(conta_bancaria_nova.toString()); // "Saldo da Conta: 100"
conta_bancaria_nova.depositarSaldo(50); // "Depósito de R$50. Saldo atual: R$150."
conta_bancaria_nova.sacarDinheiro(100); // "Saque de R$100. Saldo atual: R$50."