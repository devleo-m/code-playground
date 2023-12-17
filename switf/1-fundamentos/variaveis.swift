/*
1. Variáveis e Constantes:
Variáveis são espaços na memória que armazenam valores mutáveis. 
Em Swift, você pode criar uma variável usando var.
*/
var idade = 40
idade = 27;

//Constantes são espaços na memória que armazenam valores imutáveis. 
//Em Swift, você pode criar uma constante usando let.
let nome = "Arolnode"
// nome = "Node" // Isso resultará em um erro, pois não alterar constantes

//Dica: Use constantes sempre que o valor não precisar ser alterado. 
//Isso torna seu código mais seguro e claro.

//Quando se trata de strings, você pode combinar texto usando operador +:
let primeiroNome = "João"
let segundoNome = "Silva"
let nomeCompleto = primeiroNome + " " + segundoNome

//***********************************************************************//
//TIPOS DE DADOS
//Swift é uma linguagem com tipagem estática, o que significa que você precisa 
//declarar o tipo de dado que uma variável ou constante irá armazenar.

//Inteiros: Armazenam números inteiros, positivos e negativos.
var idade: int = 27

//Flutuantes: Armazenam números decimais.
var dinheiro: Double = 1.99

//Booleanos: Armazenam verdadeiro ou falso.
var casado: Bool = true

//Strings: Armazenam texto.
var mundo: String = "Hello, World!!!"

//***********************************************************************//
//Inferência de Tipo:
// Swift possui inferência de tipo, o que significa que em alguns casos o tipo 
//de dado pode ser inferido pelo valor atribuído:

var contador = 0 // Swift infere que é um Int
let nome2 = "Maria" // Swift infere que é uma String

//***********************************************************************//
//Interpolação de Strings:
let nota = 10
let materia = "Fisica"
let mensagem = "Eu tirei \(nota) em \(materia)"

//Isso ajuda a criar strings mais dinâmicas, inserindo valores de variáveis 
//diretamente dentro delas.







