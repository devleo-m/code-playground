/*
Manipulação de Promises e Assíncrono no TypeScript
O conceito de programação assíncrona é fundamental no desenvolvimento moderno, especialmente quando se 
trabalha com operações que levam tempo para serem concluídas, como chamadas de API, leitura de arquivos, 
ou operações em banco de dados. As Promises são uma das principais ferramentas usadas para lidar com 
operações assíncronas em TypeScript (e JavaScript).

1. O que são Promises?
Uma Promise é um objeto que representa a eventual conclusão (ou falha) de uma operação assíncrona e seu valor resultante. 
Uma Promise pode estar em um dos três estados:

Pendente (Pending): Estado inicial, nem concluída nem rejeitada.
Concluída (Fulfilled): A operação foi completada com sucesso e a Promise tem um valor resultante.
Rejeitada (Rejected): A operação falhou e a Promise tem um motivo de falha.
*/

/**
2. Criando Promises
Para criar uma Promise, usamos o construtor Promise, que recebe uma função como argumento. Essa função, por sua vez, 
recebe dois parâmetros: resolve e reject, que são funções usadas para completar ou rejeitar a Promise.
 */

const minhaPromise = new Promise<number>((resolve, reject) => {
    let sucesso = true;

    if (sucesso) {
        resolve(42); // Conclui a promessa
    }else{
        reject("Ops, algo deu errado!"); // Rejeita a promessa
    }
});

minhaPromise
    .then(resultado => console.log("O resultado é: ", resultado))
    .catch(erro => console.log("Houve um erro: ", erro));

//3. Encadeando Promises
// Uma das características mais poderosas das Promises é a capacidade de encadear operações assíncronas. 
//Você pode usar o método .then() para encadear múltiplas operações, onde cada .then() recebe o resultado 
//da Promise anterior.

const primeiraPromise = new Promise<number>((resolve) => {
    setTimeout(() => resolve(10), 1000);
});

primeiraPromise
    .then(result => {
        console.log(`Resultado 1: ${result}`);
        return result * 2;
    })
    .then(result => {
        console.log(`Resultado 2: ${result}`);
        return result * 2;
    })
    .then(result => {
        console.log(`Resultado 3: ${result}`);
    })
    .catch(erro => {
        console.log(`Erro: ${erro}`);
    });

/*
4. Promise.all e Promise.race
Promise.all: Permite que você execute várias Promises em paralelo e continue o processamento quando todas 
forem concluídas. Se qualquer uma das Promises falhar, Promise.all falhará.
*/

const promise1 = new Promise((resolve) => setTimeout(() => resolve("Primeira"), 1000));
const promise2 = new Promise((resolve) => setTimeout(() => resolve("Segunda"), 2000));
const promise3 = new Promise((resolve) => setTimeout(() => resolve("Terceira"), 3000));

Promise.all([promise1, promise2, promise3])
    .then(resultados => {
        console.log(resultados); // ["Primeira", "Segunda", "Terceira"]
    })
    .catch(erro => {
        console.log(`Erro: ${erro}`);
    });

//Promise.race: Similar a Promise.all, mas continua o processamento assim que qualquer uma das Promises for concluída (ou falhar).

const promise4 = new Promise((resolve) => setTimeout(() => resolve("Primeira"), 1000));
const promise5 = new Promise((resolve) => setTimeout(() => resolve("Segunda"), 2000));

Promise.race([promise4, promise5])
    .then(resultado => {
        console.log(resultado); // "Primeira"
    })
    .catch(erro => {
        console.log(`Erro: ${erro}`);
    });

/*
Uso de async e await
A palavra-chave async permite que uma função retorne uma Promise automaticamente, enquanto await pausa a execução da função 
async até que a Promise seja resolvida.
*/
const processarDados = async () => {
    try {
        const resultado1 = await new Promise<number>((resolve) => setTimeout(() => resolve(10), 1000));
        console.log(`Resultado 1: ${resultado1}`);

        const resultado2 = await new Promise<number>((resolve) => setTimeout(() => resolve(resultado1 * 2), 1000));
        console.log(`Resultado 2: ${resultado2}`);

        const resultado3 = await new Promise<number>((resolve) => setTimeout(() => resolve(resultado2 * 2), 1000));
        console.log(`Resultado 3: ${resultado3}`);
    } catch (erro) {
        console.log(`Erro: ${erro}`);
    }
}
processarDados();

// 6. Erros e Manipulação de Erros
// Quando você usa async e await, pode capturar erros com try/catch. Isso é equivalente a usar .catch() 
// com Promises, mas é geralmente mais limpo e fácil de ler.

const carregarDados = async () => {
    try {
        const dados = await fetch("https://api.exemplo.com/dados");
        const json = await dados.json();
        console.log(json);
    } catch (erro) {
        console.error("Erro ao carregar dados:", erro);
    }
}

carregarDados();

/*
7. Aplicação Prática
Imagine que você esteja desenvolvendo uma aplicação de e-commerce. Você precisa carregar dados do usuário, 
informações de produtos, e uma lista de categorias. Todas essas informações são carregadas de diferentes 
fontes, de forma assíncrona.

Você pode usar Promise.all para buscar todos os dados em paralelo e só continuar quando todas as informações 
estiverem disponíveis:
*/

const carregarDadosECommerce = async () =>{
    try {
        const [usuario, produtos, categorias] = await Promise.all([
            fetch("/api/usuario").then(res => res.json()),
            fetch("/api/produtos").then(res => res.json()),
            fetch("/api/categorias").then(res => res.json()),
        ]);

        console.log("Usuário:", usuario);
        console.log("Produtos:", produtos);
        console.log("Categorias:", categorias);
    } catch (erro) {
        console.error("Erro ao carregar dados do e-commerce:", erro);
    }
}

carregarDadosECommerce();
