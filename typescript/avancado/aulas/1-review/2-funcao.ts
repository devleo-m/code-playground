{
    // Funções
    // Tipagem de parâmetros e retorno: Você pode especificar os tipos de entrada (parâmetros)
    // e o tipo de saída (retorno) de uma função.
    function soma(a: number, b: number): number {
        return a + b;
    }

    const resultado = soma(5, 10); // resultado é 15
    // Parâmetros opcionais: Você pode indicar que um parâmetro é opcional, adicionando um ? 
    //após o nome do parâmetro. Se não for passado, o valor será undefined.
    function saudacao(nome?: string): string {
        return nome ? `Olá, ${nome}!` : "Olá!";
    }

    console.log(saudacao("Leo")); // "Olá, Leo!"
    console.log(saudacao());       // "Olá!"

    // Parâmetros default: Você pode definir um valor padrão para um parâmetro. Se o argumento 
    //não for passado, o valor padrão será utilizado.
    function multiplicar(a: number, b: number = 2): number {
        return a * b;
    }

    console.log(multiplicar(5));    // 10, usa b = 2
    console.log(multiplicar(5, 3)); // 15, usa b = 3

    // Resumo Rápido
    // Tipagem de parâmetros e retorno: Especifica tipos de entrada e saída.
    // Parâmetros opcionais (?): Permitem omitir o parâmetro.
    // Parâmetros default: Usam um valor padrão se não forem passados.
}