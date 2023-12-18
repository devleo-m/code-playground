//Loop while
var number: Int = 0; //Criando a variavel number com valor 0
while number < 5 { // se 0 for menor que 5 execute
    print("While: \(number)")
    number += 1 // soma number ate chegar a 5
}

// do{
//     print("k")
//     number += 1
// } while(number < 10)

//Nao existe do-while em swift, mas existe outro nome para isso
//repeat-while
var num1: Int = 0

repeat {
    print("Repeat-While: \(num1)")
    num1 += 1
} while num1 < 6

//Em Swift, o operador de incremento ++ foi removido na versão Swift 3. 
//Agora, para incrementar uma variável, você pode usar o operador +=
//por exemplo num1 += 1 ou num1 = num1 + 1

//for

for var i: Int in 0..<5 {
    print("For: \(i)")
    i += 1
}

//Loop for-in:
// Itera sobre uma coleção de itens (arrays, dicionários, intervalos etc.).

let numeros: [Any] = [1,2,3,5.5,34.2,"Leo"]

for i in numeros {
    print("For-in: \(i)")
}



