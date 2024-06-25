/*
Introdução a Objetos
Objetos são uma das características mais poderosas e úteis do JavaScript. 
Eles permitem que você armazene dados de forma estruturada, associando valores a chaves (ou propriedades). 
Um objeto pode conter diversos tipos de dados, incluindo outros objetos e funções.

O Que São Objetos?
Em JavaScript, um objeto é uma coleção de propriedades, e uma propriedade é uma associação entre um nome 
(ou chave) e um valor. Esses valores podem ser de qualquer tipo, incluindo outros objetos e funções. 
Objetos são criados utilizando a notação de chaves ({}).
*/

let pessoa = {
    nome: "João",
    idade: 30,
    saudacao: () => "Olá, meu nome é " + this.nome + "!"
};

console.log(pessoa.nome); // João
console.log(pessoa.idade); // 30
console.log(pessoa.saudacao()); // Olá, meu nome é João!

/*
Criando e Acessando Objetos
Para criar um objeto, utilizamos a notação de chaves {}. Para acessar ou modificar as propriedades de um objeto, 
utilizamos a notação de ponto (.) ou a notação de colchetes ([]).
*/

// Criando um objeto
let carro = {
    marca: "Toyota",
    modelo: "Corolla",
    ano: 2020
};

// Acessando propriedades
console.log(carro.marca); // Toyota
console.log(carro["modelo"]); // Corolla

// Modificando propriedades
carro.ano = 2021;
carro["cor"] = "preto";

console.log(carro.ano); // 2021
console.log(carro.cor); // preto

/*
Atividade 1: Criando e Manipulando Objetos
1.1 Crie um objeto chamado livro com as propriedades titulo, autor, anoPublicacao e editora.
1.2 Acesse e exiba as propriedades titulo e autor no console.
1.3 Modifique a propriedade anoPublicacao para um novo valor e adicione uma nova propriedade chamada genero.
1.4 Exiba o objeto completo no console.
*/

let livro = {
    titulo: "Titulo do livro!",
    autor: "Fulano de tal",
    anoPublicacao: 2024,
    editora: "Tal"
}

console.log(livro.titulo)
console.log(livro.autor)

livro["genero"] = "Aventura"
livro["anoPublicacao"] = 2023

console.log(livro)

/*
Métodos em Objetos
Métodos são funções associadas a objetos. Eles permitem que os objetos realizem ações e manipulem seus próprios dados.
*/

let calculadora = {
    numero1: 10,
    numero2: 5,
    soma: () => this.numero1 + this.numero2,
    subtracao: () => this.numero1 - this.numero2
};

// Objetos Aninhados
// Objetos podem conter outros objetos como propriedades. Isso permite criar estruturas de dados complexas e hierárquicas.

let empresa = {
    nome: "Tech Solutions",
    endereco: {
        rua: "Rua Principal",
        numero: 123,
        cidade: "São Paulo",
        estado: "SP"
    },
    funcionarios: [
        { nome: "Ana", cargo: "Desenvolvedora" },
        { nome: "Bruno", cargo: "Designer" }
    ]
};

console.log(empresa.nome); // Tech Solutions
console.log(empresa.endereco.cidade); // São Paulo
console.log(empresa.funcionarios[0].nome); // Ana

// Arrays de Objetos
// É comum usar arrays para armazenar múltiplos objetos. Isso é útil quando você precisa manter uma coleção de itens 
// que compartilham a mesma estrutura.


let biblioteca = [
    { titulo: "Livro A", autor: "Autor A", ano: 2010 },
    { titulo: "Livro B", autor: "Autor B", ano: 2015 },
    { titulo: "Livro C", autor: "Autor C", ano: 2020 }
];

console.log(biblioteca[0].titulo); // Livro A

// Iterando sobre um array de objetos
biblioteca.forEach(livro => {
    console.log(livro.titulo + " por " + livro.autor);
});


