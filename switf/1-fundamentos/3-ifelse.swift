// Estrutura Básica do if e else:
// O if é usado para executar um bloco de código se uma condição for verdadeira. 
// O else permite executar um bloco alternativo se a condição do if for falsa.

let idade = 18

if idade >= 18 {
    print("Você é maior de idade")
} else {
    print("Você é menor de idade")
}

// if com Várias Condições:
// Usando else if para verificar mais de uma condição.

let nota = 85

if nota >= 90 {
    print("Aprovado com A")
} else if nota >= 80 {
    print("Aprovado com B")
} else if nota >= 70 {
    print("Aprovado com C")
} else {
    print("Reprovado")
}

//if com Operadores Lógicos:
// Combinando condições com operadores lógicos && (E) e || (OU).

// Exemplo:
let salario = 3000
let horasTrabalhadas = 45

if salario > 2500 && horasTrabalhadas > 40 {
    print("Você recebeu horas extras")
} else if salario < 2500 || horasTrabalhadas <= 40 {
    print("Salário ou horas abaixo do esperado")
} else {
    print("Condições não atendidas")
}
// if com Optional Binding:
// Verificando e usando valores opcionais de forma segura.

// Exemplo:

var nome: String? = "João"

if let nome = nome {
    print("Olá, \(nome)")
} else {
    print("Nome não definido")
}
//Operador Ternário:
// Forma concisa de usar if e else.

// Exemplo:

let temperatura = 25
let mensagem = temperatura > 30 ? "Está quente" : "Está ameno"
print(mensagem)
// Estes são exemplos para explorar diferentes aspectos do if e else em Swift. 
//Praticar com eles ajudará a compreender melhor como são utilizados e como podem 
//ser aplicados em situações reais.
