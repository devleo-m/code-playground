/*
A programação assíncrona permite que as operações ocorram enquanto outras tarefas continuam 
sendo executadas, sem bloquear o fluxo de execução do programa. Em JavaScript, temos várias 
formas de lidar com operações assíncronas, 
incluindo callbacks, promessas e o uso de async / await.

Callbacks
Um callback é uma função passada como argumento para outra função, que é então invocada dentro da função externa para completar algum tipo de rotina ou ação.
*/
function saudacao(nome, callback) {
    console.log("Olá, " + nome + "!");
    callback();
}

function despedida() {
    console.log("Até logo!");
}

saudacao("Mariana", despedida);
// Saída:
// Olá, Mariana!
// Até logo!

/*
Atividade 1: Utilizando Callbacks
1.1 Crie uma função chamada processarDados que receba dois parâmetros: um array de números e uma função callback.
1.2 Dentro de processarDados, use um laço para multiplicar cada número por 2 e depois chame a função callback.
1.3 Crie uma função chamada exibirResultado que receba um array e exiba seus elementos no console.
1.4 Chame processarDados com um array de números e exibirResultado como callback.
*/

/*
Promessas
Promessas são um mecanismo mais avançado e poderoso para lidar com operações assíncronas em JavaScript. 
Elas representam um valor que pode estar disponível agora, no futuro, ou nunca.
*/

// Sintaxe Básica:
let minhaPromessa = new Promise((resolver, rejeitar) => {
    // operação assíncrona
    let sucesso = true;

    if (sucesso) {
        resolver("Operação bem-sucedida!");
    } else {
        rejeitar("Operação falhou.");
    }
});

minhaPromessa
    .then(mensagem => console.log(mensagem))
    .catch(erro => console.log(erro));

// Exemplo Real:
function buscarDados() {
    return new Promise((resolver, rejeitar) => {
        setTimeout(() => {
            let dados = { nome: "João", idade: 25 };
            resolver(dados);
        }, 2000);
    });
}

buscarDados()
    .then(dados => {
        console.log(dados.nome); // João
        console.log(dados.idade); // 25
    })
    .catch(erro => console.log(erro));

// Atividade 2: Utilizando Promessas
// 2.1 Crie uma função chamada carregarDados que retorna uma promessa.
// 2.2 Dentro da promessa, use setTimeout para simular um atraso de 3 
//     segundos e depois resolva a promessa com um objeto contendo id e valor.
// 2.3 Utilize then para exibir os dados da promessa resolvida no console.
// 2.4 Utilize catch para tratar possíveis erros.

// Async / Await
// O async e await são construções que permitem escrever código assíncrono de forma mais clara e direta. 
// Funções assíncronas retornam promessas por padrão.

//Sintaxe Básica:
async function minhaFuncao() {
    try {
        let resultado = await minhaPromessa;
        console.log(resultado);
    } catch (erro) {
        console.log(erro);
    }
}

minhaFuncao();

// Exemplo Real:
function obterUsuario() {
    return new Promise((resolver, rejeitar) => {
        setTimeout(() => {
            resolver({ nome: "Ana", idade: 30 });
        }, 2000);
    });
}

async function mostrarUsuario() {
    try {
        let usuario = await obterUsuario();
        console.log(usuario.nome); // Ana
        console.log(usuario.idade); // 30
    } catch (erro) {
        console.log(erro);
    }
}

mostrarUsuario();

/*
Atividade 3: Utilizando Async / Await
3.1 Crie uma função chamada buscarInformacoes que retorna uma promessa.
3.2 Dentro da promessa, use setTimeout para simular um atraso de 2 segundos e depois resolva a promessa com um objeto contendo titulo e conteudo.
3.3 Crie uma função assíncrona chamada exibirInformacoes.
3.4 Dentro de exibirInformacoes, use await para chamar buscarInformacoes e exibir os dados no console.
3.5 Use try e catch para tratar possíveis erros.
*/

// Comparação entre Callbacks, Promessas e Async / Await

// | Característica	    | Callbacks	          | Promessas	           | Async / Await     |
// | Facilidade de Uso	| Mais complexo	      | Moderadamente complexo | Mais simples      |
// | Tratamento de Erros| Pode ser difícil    |	Mais estruturado       | Muito estruturado |
// | Leitura do Código	| Pode ser confuso    |	Moderadamente claro    | Muito claro       |
// | Encadeamento       | Usado com "callback"|	Fácil com .then()	   | Natural com await |

//OBS
//Certamente o mais utilizado eh o ASYNC E AWAIT, vou me aprofundar mais sobre
