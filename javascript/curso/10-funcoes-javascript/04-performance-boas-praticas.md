# Aula 10 - Performance, Boas Práticas e Otimização: Funções

Agora que você entende como criar funções, é crucial aprender como criar funções **eficientes**, **manuteníveis** e que sigam **boas práticas**. Esta aula vai te ensinar a escrever código profissional.

---

## ⚡ 1. Performance de Funções

### 1.1. Recursão vs Iteração

**Problema:** Recursão pode ser mais elegante, mas geralmente é menos eficiente.

#### Exemplo: Fibonacci

```javascript
// ❌ Recursão ingênua (MUITO LENTA)
function fibonacciRecursivo(n) {
  if (n <= 1) return n;
  return fibonacciRecursivo(n - 1) + fibonacciRecursivo(n - 2);
}

// Teste de performance
console.time("Recursão");
fibonacciRecursivo(40); // Pode levar vários segundos!
console.timeEnd("Recursão");

// ✅ Iteração (RÁPIDA)
function fibonacciIterativo(n) {
  if (n <= 1) return n;
  
  let anterior = 0;
  let atual = 1;
  
  for (let i = 2; i <= n; i++) {
    let proximo = anterior + atual;
    anterior = atual;
    atual = proximo;
  }
  
  return atual;
}

console.time("Iteração");
fibonacciIterativo(40); // Milissegundos!
console.timeEnd("Iteração");
```

**Por que a recursão é lenta aqui?**
- A função calcula os mesmos valores múltiplas vezes
- `fibonacci(40)` chama `fibonacci(39)` e `fibonacci(38)`
- `fibonacci(39)` chama `fibonacci(38)` e `fibonacci(37)`
- `fibonacci(38)` é calculado **duas vezes**!
- Isso se multiplica exponencialmente

**Solução: Memoização (cache)**

```javascript
// ✅ Recursão com memoização (RÁPIDA)
function fibonacciMemo(n, cache = {}) {
  if (n in cache) {
    return cache[n]; // Retorna valor já calculado
  }
  
  if (n <= 1) {
    cache[n] = n;
    return n;
  }
  
  cache[n] = fibonacciMemo(n - 1, cache) + fibonacciMemo(n - 2, cache);
  return cache[n];
}

console.time("Memoização");
fibonacciMemo(40); // Rápido como iteração!
console.timeEnd("Memoização");
```

**Regra de Ouro:**
- Use **iteração** para loops simples
- Use **recursão** apenas quando o problema é naturalmente recursivo
- Se usar recursão, considere **memoização** para otimizar

### 1.2. Evitando Cálculos Repetidos

```javascript
// ❌ RUIM: Calcula array.length a cada iteração
function somarArray(array) {
  let total = 0;
  for (let i = 0; i < array.length; i++) { // array.length calculado toda vez!
    total += array[i];
  }
  return total;
}

// ✅ BOM: Calcula length uma vez
function somarArrayOtimizado(array) {
  let total = 0;
  let tamanho = array.length; // Calcula uma vez
  for (let i = 0; i < tamanho; i++) {
    total += array[i];
  }
  return total;
}

// ✅ MELHOR: Usa método nativo (otimizado pelo JavaScript)
function somarArrayNativo(array) {
  return array.reduce((total, num) => total + num, 0);
}
```

### 1.3. Early Return (Retorno Antecipado)

```javascript
// ❌ RUIM: Aninhamento desnecessário
function processarDados(dados) {
  if (dados) {
    if (Array.isArray(dados)) {
      if (dados.length > 0) {
        // Código principal aqui
        return dados.map(item => item * 2);
      } else {
        return [];
      }
    } else {
      return null;
    }
  } else {
    return null;
  }
}

// ✅ BOM: Early returns (mais legível e eficiente)
function processarDadosOtimizado(dados) {
  if (!dados) return null;
  if (!Array.isArray(dados)) return null;
  if (dados.length === 0) return [];
  
  // Código principal aqui (sem aninhamento)
  return dados.map(item => item * 2);
}
```

**Vantagens:**
- Menos aninhamento = mais legível
- Execução mais rápida (para casos de erro)
- Mais fácil de testar

---

## 🎯 2. Boas Práticas de Nomenclatura

### 2.1. Nomes Descritivos

```javascript
// ❌ RUIM: Nomes genéricos
function calc(a, b) { }
function proc(d) { }
function fn(x, y) { }

// ✅ BOM: Nomes descritivos
function calcularAreaRetangulo(largura, altura) { }
function processarDadosUsuario(dados) { }
function verificarSeUsuarioEstaAtivo(usuario) { }
```

### 2.2. Convenções de Nomenclatura

```javascript
// Funções que retornam boolean: começam com "é", "tem", "pode"
function eMaiorDeIdade(idade) { }
function temPermissao(usuario, acao) { }
function podeAcessar(recurso, usuario) { }

// Funções que fazem ações: verbos no infinitivo
function calcularTotal(precos) { }
function validarEmail(email) { }
function criarUsuario(dados) { }
function removerItem(lista, item) { }

// Funções que retornam dados: substantivos
function obterNomeCompleto(usuario) { }
function buscarUsuarioPorId(id) { }
function listarProdutos() { }
```

### 2.3. Evitando Nomes Ambíguos

```javascript
// ❌ RUIM: Ambíguo
function processar(data) { } // O que processa?
function verificar(input) { } // Verifica o quê?

// ✅ BOM: Específico
function processarDadosDeVenda(dados) { }
function verificarSeEmailEValido(email) { }
```

---

## 🏗️ 3. Organização e Estrutura

### 3.1. Funções Pequenas e Focadas

```javascript
// ❌ RUIM: Função fazendo muitas coisas
function processarPedido(pedido) {
  // Validar pedido
  if (!pedido.itens || pedido.itens.length === 0) {
    return { erro: "Pedido vazio" };
  }
  if (!pedido.cliente) {
    return { erro: "Cliente não informado" };
  }
  
  // Calcular total
  let total = 0;
  for (let item of pedido.itens) {
    total += item.preco * item.quantidade;
  }
  
  // Aplicar desconto
  if (pedido.cliente.vip) {
    total = total * 0.9; // 10% de desconto
  }
  
  // Criar registro
  let registro = {
    id: Date.now(),
    cliente: pedido.cliente.nome,
    total: total,
    data: new Date()
  };
  
  // Salvar (simulado)
  console.log("Salvando:", registro);
  
  return registro;
}

// ✅ BOM: Funções pequenas e focadas
function validarPedido(pedido) {
  if (!pedido.itens || pedido.itens.length === 0) {
    return { valido: false, erro: "Pedido vazio" };
  }
  if (!pedido.cliente) {
    return { valido: false, erro: "Cliente não informado" };
  }
  return { valido: true };
}

function calcularTotalItens(itens) {
  return itens.reduce((total, item) => {
    return total + (item.preco * item.quantidade);
  }, 0);
}

function aplicarDescontoVip(total, cliente) {
  if (cliente.vip) {
    return total * 0.9;
  }
  return total;
}

function criarRegistroPedido(cliente, total) {
  return {
    id: Date.now(),
    cliente: cliente.nome,
    total: total,
    data: new Date()
  };
}

function processarPedidoOtimizado(pedido) {
  // Validar
  const validacao = validarPedido(pedido);
  if (!validacao.valido) {
    return { erro: validacao.erro };
  }
  
  // Calcular
  const totalItens = calcularTotalItens(pedido.itens);
  const totalComDesconto = aplicarDescontoVip(totalItens, pedido.cliente);
  
  // Criar registro
  const registro = criarRegistroPedido(pedido.cliente, totalComDesconto);
  
  // Salvar
  console.log("Salvando:", registro);
  
  return registro;
}
```

**Vantagens:**
- Cada função tem uma responsabilidade única
- Mais fácil de testar
- Mais fácil de reutilizar
- Mais fácil de entender e manter

### 3.2. Princípio DRY (Don't Repeat Yourself)

```javascript
// ❌ RUIM: Código repetido
function calcularPrecoProduto1(preco) {
  let imposto = preco * 0.1;
  let desconto = preco * 0.05;
  return preco + imposto - desconto;
}

function calcularPrecoProduto2(preco) {
  let imposto = preco * 0.1; // Repetido!
  let desconto = preco * 0.05; // Repetido!
  return preco + imposto - desconto;
}

// ✅ BOM: Código reutilizável
function calcularImposto(preco, taxa = 0.1) {
  return preco * taxa;
}

function calcularDesconto(preco, percentual = 0.05) {
  return preco * percentual;
}

function calcularPrecoFinal(preco, taxaImposto = 0.1, percentualDesconto = 0.05) {
  const imposto = calcularImposto(preco, taxaImposto);
  const desconto = calcularDesconto(preco, percentualDesconto);
  return preco + imposto - desconto;
}
```

---

## 🛡️ 4. Validação e Tratamento de Erros

### 4.1. Validando Parâmetros

```javascript
// ❌ RUIM: Sem validação
function dividir(a, b) {
  return a / b; // E se b for 0? E se não forem números?
}

// ✅ BOM: Com validação
function dividirSeguro(a, b) {
  // Validar tipos
  if (typeof a !== 'number' || typeof b !== 'number') {
    throw new TypeError('Ambos os parâmetros devem ser números');
  }
  
  // Validar divisão por zero
  if (b === 0) {
    throw new Error('Divisão por zero não é permitida');
  }
  
  // Validar valores especiais
  if (!isFinite(a) || !isFinite(b)) {
    throw new Error('Números devem ser finitos');
  }
  
  return a / b;
}
```

### 4.2. Tratamento de Erros com Try/Catch

```javascript
function processarDadosSeguro(dados) {
  try {
    // Validação
    if (!dados || typeof dados !== 'object') {
      throw new Error('Dados inválidos');
    }
    
    // Processamento
    return dados.map(item => item * 2);
    
  } catch (erro) {
    // Log do erro (em produção, use um sistema de logging)
    console.error('Erro ao processar dados:', erro.message);
    
    // Retornar valor padrão ou re-lançar erro
    return [];
  }
}
```

### 4.3. Validação com Valores Padrão Seguros

```javascript
// ❌ RUIM: Valores padrão podem mascarar erros
function criarUsuario(nome = "Anônimo", idade = 0) {
  return { nome, idade };
}

criarUsuario(null, null); // { nome: null, idade: 0 } - Erro silencioso!

// ✅ BOM: Validação explícita
function criarUsuarioSeguro(nome, idade) {
  if (!nome || typeof nome !== 'string' || nome.trim() === '') {
    throw new Error('Nome é obrigatório e deve ser uma string não vazia');
  }
  
  if (typeof idade !== 'number' || idade < 0) {
    throw new Error('Idade deve ser um número positivo');
  }
  
  return { nome: nome.trim(), idade };
}
```

---

## 🔒 5. Segurança

### 5.1. Sanitização de Entrada

```javascript
// ❌ RUIM: Aceita qualquer entrada
function buscarUsuario(nome) {
  // Se nome contém código malicioso, pode causar problemas
  return usuarios.find(u => u.nome === nome);
}

// ✅ BOM: Sanitiza entrada
function buscarUsuarioSeguro(nome) {
  // Validar e sanitizar
  if (typeof nome !== 'string') {
    throw new TypeError('Nome deve ser uma string');
  }
  
  // Remover caracteres perigosos
  const nomeSanitizado = nome.trim().replace(/[<>]/g, '');
  
  // Validar comprimento
  if (nomeSanitizado.length === 0 || nomeSanitizado.length > 100) {
    throw new Error('Nome inválido');
  }
  
  return usuarios.find(u => u.nome === nomeSanitizado);
}
```

### 5.2. Evitando Eval e Funções Perigosas

```javascript
// ❌ MUITO PERIGOSO: Nunca use eval com entrada do usuário
function calcularExpressao(expressao) {
  return eval(expressao); // PERIGOSO! Permite execução de código arbitrário
}

// ✅ SEGURO: Parser próprio ou biblioteca validada
function calcularExpressaoSegura(expressao) {
  // Usar uma biblioteca de parsing matemático validada
  // ou criar seu próprio parser seguro
  // Nunca use eval com dados do usuário!
}
```

---

## 📊 6. Debugging e Ferramentas

### 6.1. Console Methods Úteis

```javascript
function processarDadosComplexos(dados) {
  // console.log - Informação geral
  console.log('Iniciando processamento:', dados);
  
  // console.table - Para arrays/objetos
  console.table(dados);
  
  // console.time/timeEnd - Medir performance
  console.time('processamento');
  const resultado = dados.map(item => item * 2);
  console.timeEnd('processamento');
  
  // console.group - Agrupar logs
  console.group('Detalhes do processamento');
  console.log('Total de itens:', dados.length);
  console.log('Resultado:', resultado);
  console.groupEnd();
  
  // console.assert - Verificações
  console.assert(resultado.length === dados.length, 
    'Resultado deve ter o mesmo tamanho dos dados');
  
  return resultado;
}
```

### 6.2. Breakpoints e DevTools

```javascript
function funcaoParaDebugar(parametro) {
  // Adicione 'debugger;' para pausar a execução
  debugger;
  
  // Ou use breakpoints no DevTools do navegador
  // 1. Abra DevTools (F12)
  // 2. Vá para a aba Sources
  // 3. Encontre seu arquivo
  // 4. Clique na linha para adicionar breakpoint
  
  let resultado = parametro * 2;
  return resultado;
}
```

---

## 🧪 7. Testabilidade

### 7.1. Funções Puras (Pure Functions)

Funções puras são mais fáceis de testar porque sempre retornam o mesmo resultado para as mesmas entradas.

```javascript
// ✅ BOM: Função pura (fácil de testar)
function somar(a, b) {
  return a + b;
}

// Sempre retorna 5 para somar(2, 3)
// Não depende de estado externo
// Não causa efeitos colaterais

// ❌ RUIM: Função impura (difícil de testar)
let contador = 0;

function somarComContador(a, b) {
  contador++; // Efeito colateral
  return a + b; // Depende de estado externo
}

// Resultado pode variar dependendo de quantas vezes foi chamada
```

### 7.2. Separando Lógica de Efeitos Colaterais

```javascript
// ❌ RUIM: Lógica misturada com efeitos
function processarEImprimir(dados) {
  const resultado = dados.map(item => item * 2);
  console.log(resultado); // Efeito colateral
  document.getElementById('output').textContent = resultado; // Efeito colateral
  return resultado;
}

// ✅ BOM: Separado
function processar(dados) {
  // Lógica pura (fácil de testar)
  return dados.map(item => item * 2);
}

function imprimir(resultado) {
  // Efeitos colaterais separados
  console.log(resultado);
  document.getElementById('output').textContent = resultado;
}

// Uso
const resultado = processar(dados);
imprimir(resultado);
```

---

## 🎨 8. Padrões de Código

### 8.1. Factory Functions

```javascript
// ✅ BOM: Factory function para criar objetos
function criarUsuario(nome, email, idade) {
  return {
    nome: nome,
    email: email,
    idade: idade,
    ativo: true,
    obterNomeCompleto() {
      return this.nome;
    },
    desativar() {
      this.ativo = false;
    }
  };
}

const usuario1 = criarUsuario("Maria", "maria@email.com", 25);
const usuario2 = criarUsuario("João", "joao@email.com", 30);
```

### 8.2. Higher-Order Functions

```javascript
// ✅ BOM: Função que retorna função
function criarMultiplicador(multiplicador) {
  return function(numero) {
    return numero * multiplicador;
  };
}

const dobrar = criarMultiplicador(2);
const triplicar = criarMultiplicador(3);

console.log(dobrar(5));    // 10
console.log(triplicar(5)); // 15
```

### 8.3. Funções como Parâmetros (Callbacks)

```javascript
// ✅ BOM: Função que aceita outra função
function processarArray(array, funcaoProcessamento) {
  const resultado = [];
  for (let item of array) {
    resultado.push(funcaoProcessamento(item));
  }
  return resultado;
}

// Uso flexível
const numeros = [1, 2, 3, 4, 5];
const dobrados = processarArray(numeros, x => x * 2);
const quadrados = processarArray(numeros, x => x * x);
```

---

## 📈 9. Gerenciamento de Memória

### 9.1. Evitando Vazamentos de Memória

```javascript
// ❌ RUIM: Referências que não são limpas
let callbacks = [];

function adicionarCallback(callback) {
  callbacks.push(callback);
  // Se nunca remover, a memória vaza!
}

// ✅ BOM: Sistema de remoção
function criarGerenciadorCallbacks() {
  const callbacks = [];
  
  return {
    adicionar(callback) {
      callbacks.push(callback);
      // Retorna função para remover
      return () => {
        const index = callbacks.indexOf(callback);
        if (index > -1) {
          callbacks.splice(index, 1);
        }
      };
    },
    executarTodos() {
      callbacks.forEach(cb => cb());
    }
  };
}

const gerenciador = criarGerenciadorCallbacks();
const remover = gerenciador.adicionar(() => console.log("Callback"));
// Quando não precisar mais:
remover(); // Remove da memória
```

### 9.2. Evitando Closures Desnecessários

```javascript
// ❌ RUIM: Closure mantém referência grande
function criarProcessador(dadosGrandes) {
  return function(item) {
    // Closure mantém referência a dadosGrandes na memória
    return dadosGrandes.processar(item);
  };
}

// ✅ BOM: Passar apenas o necessário
function criarProcessadorOtimizado(processador) {
  return function(item) {
    // Apenas a função processar é mantida
    return processador(item);
  };
}
```

---

## 🎓 Resumo: Checklist de Boas Práticas

Ao criar funções, sempre verifique:

- [ ] **Nomenclatura**: Nome descritivo que explica o que a função faz
- [ ] **Tamanho**: Função pequena e focada em uma única responsabilidade
- [ ] **Parâmetros**: Validados e com tipos corretos
- [ ] **Retorno**: Sempre retorna um valor consistente
- [ ] **Validação**: Edge cases tratados (null, undefined, arrays vazios, etc.)
- [ ] **Erros**: Erros tratados adequadamente
- [ ] **Performance**: Sem cálculos desnecessários ou repetidos
- [ ] **Legibilidade**: Código fácil de entender sem comentários excessivos
- [ ] **Testabilidade**: Função pura quando possível, efeitos colaterais separados
- [ ] **Segurança**: Entrada sanitizada, sem eval ou código perigoso
- [ ] **DRY**: Sem código duplicado
- [ ] **Documentação**: Comentários apenas quando necessário (código auto-explicativo é melhor)

---

## 🚀 Próximos Passos

Agora que você sabe criar funções eficientes e seguir boas práticas, você está pronto para:
- Aprender sobre Closures (funções que "lembram" do escopo)
- Entender Higher-Order Functions em profundidade
- Explorar programação assíncrona com callbacks, Promises e async/await
- Aprender sobre módulos e organização de código

Lembre-se: **código bom não é apenas código que funciona, é código que outros desenvolvedores (incluindo você no futuro) conseguem entender e manter facilmente!**

