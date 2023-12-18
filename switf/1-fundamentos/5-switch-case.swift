// Estrutura Básica do switch:
// O switch é usado para verificar um valor em relação a vários casos.

let diaSemana: String = "Segunda"
switch diaSemana {
    case "Segunda":
        print("Segunda-feira. Dia de começar a trabalhar!")
    case "Terca":
        print("Terca-feira. Dia para trabalhar muito!")
    case "Quarta":
        print("Quarta-feira. Outro dia para trabalhar!")
    case "Quinta":
        print("Quinta-feira. Mais um dia de trabalho")
    case "Sexta":
        print("Sexta-feira. SEXTOUUUU!! QUASE NA FOLGA!")
    case "Sábado", "Domingo":
        print("\(diaSemana) Folga! Dia para descansar")
    default:
        print("Erro, isso nao eh dia da semana ou escreva corretamente")
}
// 2. Intervalos e Casos Compostos:
// Usando intervalos e casos compostos em um switch.

let nota: Int = 85
switch nota {
    case 90...100:
        print("A")
    case 80..<90:
        print("B")
    case 70..<80:
        print("C")
    default:
        print("Outra nota")
}
// 3. Match com Enumerações:
// Usando switch com enums, incluindo casos associados.

enum Direcao {
    case cima
    case baixo
    case esquerda
    case direita
}
let direcao: Direcao = .cima
switch direcao {
    case .cima:
        print("Para cima")
    case .baixo:
        print("Para baixo")
    case .esquerda, .direita:
        print("Esquerda ou direita")
}
// 4. Usando where:
// Adicionando condições adicionais usando where.

let numero: Int = 15
switch numero {
    case let x where x % 2 == 0:
        print("Número par")
    case let x where x % 2 != 0:
        print("Número ímpar")
    default:
        print("Outro tipo de número")
}
// 5. Switch com Valores Opcionais:
// Usando switch com valores opcionais.

let valorOpcional: Int? = 10
switch valorOpcional {
    case .some(let valor):
        print("Valor definido: \(valor)")
    case .none:
        print("Valor não definido")
}
