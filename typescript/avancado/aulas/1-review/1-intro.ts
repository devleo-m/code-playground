{
    // Tipos Básicos
    // Tipagem explícita: Especifica o tipo de uma variável ao declará-la.

    // string: Representa texto.

    let nome: string = "Leo";
    // number: Representa números (inteiros e decimais).

    let idade: number = 28;
    // boolean: Representa valores lógicos (true ou false).

    let ativo: boolean = true;
    // any: Aceita qualquer tipo de valor. Use com cuidado.

    let qualquer: any = "texto ou número";
    // void: Usado em funções que não retornam valor.

    function funcao(): void {
        console.log("Sem retorno");
    }
    // unknown: Semelhante ao any, mas mais seguro, pois requer verificação de tipo antes de uso.

    let valor: unknown = "pode ser qualquer coisa";
    // never: Indica que uma função nunca retorna, geralmente usada em funções que lançam erros.

    function erro(): never {
        throw new Error("Erro!");
    }
    // Arrays: Tipos que contêm múltiplos valores do mesmo tipo.

    // Sintaxe:

    let numeros: number[] = [1, 2, 3];
    let nomes: Array<string> = ["Leo", "Ana"];
    // Tuplas: Arrays com um número fixo de elementos de tipos diferentes.

    // Exemplo:

    let pessoa: [number, string] = [28, "Leo"];
    // Enumerações (enum): Um tipo que permite definir um conjunto de constantes nomeadas.

    // Exemplo:

    enum Cor {
        Vermelho,
        Verde,
        Azul
    }
    let minhaCor: Cor = Cor.Vermelho; // 0
    // Tipos literais: Permitem restringir uma variável a um conjunto específico de valores.

    // Exemplo:
    let estado: "ativo" | "inativo" = "ativo";

// Resumo Rápido:
// Tipagem explícita: Define tipos como string, number, boolean, etc.
// Arrays: Listas de valores do mesmo tipo.
// Tuplas: Arrays com tipos fixos e específicos.
// Enumerações: Conjunto de constantes nomeadas.
// Tipos literais: Restringem valores a um conjunto específico.
}