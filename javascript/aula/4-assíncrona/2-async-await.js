/*
async e await são recursos introduzidos no ECMAScript 2017 (ES8) que permitem trabalhar 
com código assíncrono de maneira mais clara e concisa. Eles facilitam a leitura e escrita 
de código assíncrono, eliminando a necessidade de aninhamento excessivo de promessas 
(.then e .catch). Nesta aula, vamos explorar detalhadamente como usar async e await, 
com muitos exemplos práticos e dicas úteis.

Conceitos Básicos
Funções Assíncronas
Uma função assíncrona é declarada com a palavra-chave async. Isso indica que a 
função retornará uma promessa.
*/

async function minhaFuncao() {
    return "Olá, mundo!";
}

minhaFuncao().then(mensagem => console.log(mensagem)); // Olá, mundo!

/*
Palavra-chave await
A palavra-chave await só pode ser usada dentro de funções assíncronas. 
Ela pausa a execução da função assíncrona até que a promessa seja resolvida ou rejeitada.
*/

async function exemploAwait() {
    let promessa = new Promise((resolve, reject) => {
        setTimeout(() => resolve("Promessa resolvida!"), 2000);
    });

    let resultado = await promessa;
    console.log(resultado); // Promessa resolvida! (após 2 segundos)
}

exemploAwait();

/*
Vantagens do async / await
Código mais legível: Facilita a leitura e escrita de código assíncrono.
Fluxo de controle linear: O código é escrito de maneira linear, como se fosse síncrono.
Tratamento de erros simplificado: Erros podem ser capturados usando try / catch.
*/

// Exemplos Práticos
// Chamada Simples a uma API
// Vamos fazer uma chamada simples a uma API usando fetch e async / await.

async function buscarDadosAPI() {
    try {
        let resposta = await fetch('https://jsonplaceholder.typicode.com/posts/1');
        let dados = await resposta.json();
        console.log(dados);
    } catch (erro) {
        console.error('Erro ao buscar dados:', erro);
    }
}

buscarDadosAPI();

// Encadeamento de Promessas com async / await
// Você pode usar await múltiplas vezes dentro de uma função assíncrona para encadear várias promessas.

async function buscarPostsEComentarios() {
    try {
        let respostaPosts = await fetch('https://jsonplaceholder.typicode.com/posts/1');
        let post = await respostaPosts.json();
        
        let respostaComentarios = await fetch(`https://jsonplaceholder.typicode.com/posts/${post.id}/comments`);
        let comentarios = await respostaComentarios.json();
        
        console.log('Post:', post);
        console.log('Comentários:', comentarios);
    } catch (erro) {
        console.error('Erro:', erro);
    }
}

buscarPostsEComentarios();

// Tratamento de Erros com async / await
// O tratamento de erros em funções assíncronas é feito usando blocos try / catch.
async function exemploErro() {
    try {
        let resposta = await fetch('https://jsonplaceholder.typicode.com/invalid-url');
        let dados = await resposta.json();
        console.log(dados);
    } catch (erro) {
        console.error('Erro ao buscar dados:', erro);
    }
}

exemploErro();

//Executando Várias Promessas em Paralelo
// Às vezes, é necessário executar várias promessas em paralelo. Podemos usar Promise.all para isso.
async function buscarDadosParalelos() {
    try {
        let [resposta1, resposta2] = await Promise.all([
            fetch('https://jsonplaceholder.typicode.com/posts/1'),
            fetch('https://jsonplaceholder.typicode.com/posts/2')
        ]);

        let dados1 = await resposta1.json();
        let dados2 = await resposta2.json();

        console.log('Post 1:', dados1);
        console.log('Post 2:', dados2);
    } catch (erro) {
        console.error('Erro:', erro);
    }
}

buscarDadosParalelos();

// Dicas e Boas Práticas
// Use async / await em vez de .then e .catch para código mais limpo e legível.
// Sempre use try / catch para tratar erros em funções assíncronas.
// Prefira Promise.all para executar várias promessas em paralelo e melhorar a performance.
// Documente bem suas funções assíncronas para esclarecer o comportamento esperado e os possíveis erros.

/*
exploramos o uso de async e await em JavaScript para facilitar a escrita de código assíncrono. 
Vimos como criar funções assíncronas, usar await para esperar por promessas, tratar erros de 
maneira eficaz e executar múltiplas promessas em paralelo. Com esses conhecimentos, 
você está preparado para lidar com operações assíncronas de forma mais eficiente e clara.
*/
