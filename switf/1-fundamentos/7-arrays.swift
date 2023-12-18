//ARRAYS
//os arrays em Swift são usados para armazenar múltiplos valores 
//do mesmo tipo em uma única coleção ordenada. Eles oferecem uma 
//maneira flexível de trabalhar com conjuntos de dados.

//Criando Arrays:

//array varios:
var vazio = [Int]()

//Array com Valores Iniciais:
var numeros = [1, 2, 3, 4, 5]

//Acessando e Modificando Elementos:
// Acesso por Índice:
print(numeros[0]) //1
print(numeros[3]) //4

//Modificação de Elementos
numeros[4] = 7

//Propriedades e Métodos dos Arrays:
let tamanhoDoArray = numeros.count // Número de elementos no array
let estaVazio = numeros.isEmpty // Verifica se o array está vazio

//Adicionando Elementos:
numeros.append(6) // Adiciona um novo elemento ao final do array
numeros.insert(7, at: 0) // Insere o valor 7 no índice 0

//Removendo Elementos:
let ultimo = numeros.removeLast() // Remove e retorna o último elemento
let primeiro = numeros.removeFirst() // Remove e retorna o primeiro elemento
numeros.remove(at: 2) // Remove o elemento no índice 2

// Iterando sobre um Array:
// Usando for-in
for i in numeros {
    print(i)
}

//Arrays com Tipos Diversos:
// Array de Tipos Diversos:

var elementos: [Any] = [1, "dois", 3.0, true]

//Arrays são Tipados:
//Arrays Imutáveis:
let letras: [Character] = ["a", "b", "c"]

//Arrays Mutáveis
var numeros: [Int] = [1, 2, 3, 4]

//OBS
//Os arrays são uma estrutura de dados poderosa em Swift, 
//oferecendo várias maneiras de manipular e trabalhar com 
//conjuntos de elementos do mesmo tipo. Praticar com arrays é 
//fundamental para entender sua flexibilidade e utilidade em 
//diversos cenários de programação.








