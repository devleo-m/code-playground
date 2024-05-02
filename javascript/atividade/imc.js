console.log("Ola seja bem vindo, vamos calcular o seu IMC")
console.log(relatorioIMC(72, 1.69))

function calculaIMC(peso, altura) {
    calcTotal = peso / (altura * altura)
    return calcTotal
}

function classificaIMC(calculaIMC) {
    if (calculaIMC >= 30) {
        return "Obesidade"
    } 
    if (calculaIMC >= 25 && calculaIMC <= 29.9) {
        return "Sobrepeso"

    } 
    if (calculaIMC >= 18.5 && calculaIMC <= 24.9) {
        return "Peso normal"
    }
     else {
        return "Abaixo do peso"
    }
}

function relatorioIMC(peso, altura) {
    let imc = calculaIMC(peso, altura)
    console.log(`Seu IMC eh: ${imc}`)
    console.log(`Classificacao: ${classificaIMC(imc)}`)
}
