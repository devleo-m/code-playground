func calculadora(num1: Double, operador: String, num2: Double) -> Double {
    switch operador {
    case "+":
        return num1 + num2
    case "-":
        return num1 - num2
    case "*":
        return num1 * num2
    case "/":
        if num2 != 0 {
            return num1 / num2
        } else {
            return 0 // Retornando um valor padrão para divisão por zero
        }
    default:
        return 0 // Retornando um valor padrão para casos de operação inválida
    }
}

print("Digite o primeiro número: ")
if let numero1 = Double(readLine() ?? "") {
    print("Digite o operador (+, -, *, /)")
    if let operacao = readLine(), ["+", "-", "*", "/"].contains(operacao) {
        print("Digite o segundo número: ")
        if let numero2 = Double(readLine() ?? "") {
            let resultado = calculadora(num1: numero1, operador: operacao, num2: numero2)
            print("Total: \(resultado)")
        } else {
            print("Segundo número inválido")
        }
    } else {
        print("Operação inválida")
    }
} else {
    print("Primeiro número inválido")
}
