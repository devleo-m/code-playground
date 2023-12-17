//Hello word em swift
print("Hello, world!")

//1. Variáveis e Constantes:
// Variáveis: Espaços na memória que armazenam valores mutáveis.
var idade = 25
idade = 26 // Valor da variável pode ser alterado

//Constantes: Espaços na memória que armazenam valores imutáveis.
let nome = "Maria"
// nome = "João" // Isso resultará em um erro, pois não pode ser alterado

// Tipos de Dados:
// Inteiros, Flutuantes, Booleanos e Strings: Tipos de dados básicos em Swift.
let numeroInteiro: Int = 10
let numeroDecimal: Double = 3.14
let estaChovendo: Bool = true
let saudacao: String = "Olá, mundo!"

// Estruturas de Controle:
// If-else, switch-case e Loops (for, while): Controle de fluxo em Swift.
let aluno: String = "Arolnode"
var nota1: Int = 7
var nota2: Int = 10
var nota3: Int = 4
var nota4: Int = 5
var passouDeAno: Bool = false;
var recuperacao: Bool = false;

var notaTotal: Int = (nota1+nota2+nota3+nota4)/4

if notaTotal >= 7 {
    print("Parabens, \(aluno)! Passou de ano. Nota: \(notaTotal)")
    passouDeAno = true
    recuperacao = false

} else if(notaTotal >= 5) {
    print("RECUPERACAO! Caro(a) aluno(a) \(aluno)! Voce esta na corda. Nota: \(notaTotal)")
    recuperacao = true
    passouDeAno = false
} else {
    print("REPROVADO!, \(aluno)! Voce eh a vergonha da profisson. Nota: \(notaTotal)")
    passouDeAno = false
    recuperacao = false
}

var notaRecuperacao: Int = 6
var notaTotal2: Int = (notaTotal + notaRecuperacao)/2

if recuperacao {
    switch notaTotal2 {
        case 10:
            print("Passou de ano, ate ano que vem, nota: \(notaTotal2)")
            break;
        case 9:
            print("Passou de ano, ate ano que vem nota: \(notaTotal2)")
            break;
        case 8:
            print("Passou de ano, ate ano que vem nota: \(notaTotal2)")
            break;
        case 7:
            print("Passou de ano, ate ano que vem nota: \(notaTotal2)")
            break;
        case 6:
            print("Passou de ano, ate ano que vem nota: \(notaTotal2)")
            break;
        case 5:
            print("reprovado!! nota: \(notaTotal2)")
            break;
        case 4:
            print("reprovado nota: \(notaTotal2)")
            break;
        case 3:
            print("reprovado nota: \(notaTotal2)")
            break;
        case 2:
            print("reprovado nota: \(notaTotal2)")
            break;
        case 1:
            print("reprovado nota: \(notaTotal2)")
            break;
        case 0:
            print("reprovado nota: \(notaTotal2)")
            break;
        default:
            print("Nota invalida")
    }
}

//Coleções de Dados:
//Arrays, Dicionários e Sets: Estruturas para armazenar coleções de dados.
var numeros = [1, 2, 3, 4]
var dicionario = ["chave1": "valor1", "chave2": "valor2"]
var conjunto: Set<Int> = [1, 2, 3, 4]

//Funções:
// Definir Funções, Parâmetros e Retornos:
func saudacao(nome: String) -> String {
    return "Ola, \(nome)"
}

import Foundation //Importa da dependencias iguais a do Math para fazer calculos sqrt
func pitagoras(a : Float, b : Float) -> Float {
    var x: Float = (a*a)+(b*b)
    x = sqrt(x) //raiz quedrada
    return x
}

print(saudacao(nome: "Leo"))
print(pitagoras(a: 25, b: 32))

//Programação Orientada a Objetos:
class Carro{ //Criacao de uma classe carro
    var modelo:String
    var cor:String
    var ano:Int

    init(modelo:String, cor:String, ano:Int){
        self.modelo = modelo
        self.cor = cor
        self.ano = ano
    }

    func acelerar(velocidade:Int){
        print("Velocidade: \(velocidade)km/h")
    }
}

let carro1 = Carro(modelo:"Celta",cor: "Branco",ano: 2013) //Criacao do objeto
carro1.acelerar(velocidade: 60)

let carro2 = Carro(modelo: "Fusca", cor: "Amarelo", ano: 1969)
carro2.acelerar(velocidade: 360)

//Tratamento de Erros:
// Do-Try-Catch: Lidar com exceções e erros.
enum ErroPersonalizado: Error {
    case erro1
    case erro2
}
func funcaoQuePodeDarErro() throws {
    throw ErroPersonalizado.erro1
}

do{
    try funcaoQuePodeDarErro()
} catch{
    print("Ocorreu um erro: \(error)")
}





