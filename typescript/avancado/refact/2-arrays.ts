{
    //Criacao de uma array
    const frutas: string[] = ["maca", "banana", "laranja", "uva", "abacaxi"];

    /*
    Agora, vamos explorar os principais métodos de arrays e como utilizá-los.

    - push()
    Adiciona um ou mais elementos ao final do array e retorna o novo comprimento do array.
    */
   frutas.push("manga");
   console.log(frutas); // ["maca", "banana", "laranja", "uva", "abacaxi", "manga"]

    /*
    - pop()
    Remove o último elemento do array e o retorna. Se o array estiver vazio, retorna undefined.
    */
    const ultimaFruta = frutas.pop();
    frutas.pop();
    console.log(ultimaFruta); // manga
    console.log(frutas); // ["maca", "banana", "laranja", "uva"]

    /*
    - shift()
    Remove o primeiro elemento do array e o retorna. Se o array estiver vazio, retorna undefined.
    */
    const primeiraFruta = frutas.shift();
    console.log(primeiraFruta)
    frutas.shift();
    console.log(frutas); // ["laranja", "uva"]

    /*
    - unshift()
    Adiciona um ou mais elementos ao início do array e retorna o novo comprimento do array.
    */
    frutas.unshift("goiaba", "manga");
    frutas.unshift("pera");
    console.log(frutas); // ["pera", "goiaba", "manga", "laranja", "uva"]

    /*
    map()
    Cria um novo array com os resultados da chamada de uma função para cada elemento do array.
    */
   const frutasMaiusculas = frutas.map(x => x.toUpperCase())
   console.log(frutasMaiusculas)

   const frutasNumeros = frutas.map(x => x.length)
   console.log(frutasNumeros)

   const frutasCompridas = frutas.map(x => x + "!")
   console.log(frutasCompridas)

   /*
    filter()
    Cria um novo array com todos os elementos que passam em um teste definido por uma função.
   */

   const frutasComL = frutas.filter(x => x.includes("l"))
   console.log(frutasComL)

   const frutasCom5Letras = frutas.filter(x => x.length === 5)
   console.log(frutasCom5Letras)

   /*
    reduce()
    Aplica uma função acumuladora a cada elemento do array e retorna o valor final.
   */
    const num: number[] = [1, 2, 3, 4, 5]

   const somaNum = num.reduce((acumulador, valorAtual) => acumulador + valorAtual, 0)
   console.log(somaNum)

   /*forEach()
    Executa uma função fornecida uma vez para cada elemento do array. 
    */

    frutas.forEach(x => console.log(x))
    num.forEach(x => console.log(x))

    //Exemplo Completo -----

    const pessoas = [
        { nome: "Leonardo", idade: 27 },
        { nome: "Gabriel", idade: 17 },
        { nome: "Lucas", idade: 11 },
    ];
    
    // Adicionando uma nova pessoa
    pessoas.push({ nome: "Ana", idade: 22 });
    
    // Removendo a última pessoa
    pessoas.pop();
    
    // Filtrando pessoas com idade maior que 25
    const pessoasMaiorQue25 = pessoas.filter(pessoa => pessoa.idade > 25);
    
    // Mapeando para criar um array de nomes
    const nomes = pessoas.map(pessoa => pessoa.nome);
    
    // Exibindo resultados
    console.log(pessoas); // [{ nome: "Leonardo", idade: 28 }, { nome: "Gabriel", idade: 28 }]
    console.log(pessoasMaiorQue25); // [{ nome: "Leonardo", idade: 28 }, { nome: "Gabriel", idade: 28 }]
    console.log(nomes); // ["Leonardo", "Gabriel"]
    
}