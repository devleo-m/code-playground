# Aula 5 - Performance, Boas Práticas e Otimização: Estruturas de Dados

## 🎯 Introdução

Escolher a estrutura de dados correta não é apenas sobre funcionalidade - é sobre **performance, manutenibilidade e escalabilidade**. Nesta aula, você aprenderá como trabalhar com estruturas de dados de forma profissional e eficiente.

---

## ⚡ Performance: Escolhendo a Estrutura Correta

### 1. Array vs Set: Verificação de Existência

**❌ Evite usar Array.includes() em loops ou verificações frequentes:**

```javascript
// RUIM - O(n) - precisa verificar cada elemento
const emails = ['email1@test.com', 'email2@test.com', /* ... 1000 emails */];
function verificarEmail(email) {
    return emails.includes(email); // Lento para arrays grandes!
}
```

**✅ Use Set.has() para verificações frequentes:**

```javascript
// BOM - O(1) - busca instantânea
const emails = new Set(['email1@test.com', 'email2@test.com', /* ... 1000 emails */]);
function verificarEmail(email) {
    return emails.has(email); // Muito mais rápido!
}
```

**Por quê?**
- `Array.includes()` tem complexidade **O(n)** - precisa verificar cada elemento
- `Set.has()` tem complexidade **O(1)** - busca direta (hash table)
- Para 1000 elementos: Array pode verificar até 1000 vezes, Set verifica 1 vez

**Benchmark exemplo:**
```javascript
// Array com 10.000 elementos
const array = Array.from({length: 10000}, (_, i) => `item${i}`);
console.time("Array.includes");
array.includes('item9999'); // Verifica todos os elementos
console.timeEnd("Array.includes"); // ~0.5ms

// Set com 10.000 elementos
const set = new Set(array);
console.time("Set.has");
set.has('item9999'); // Busca direta
console.timeEnd("Set.has"); // ~0.001ms (500x mais rápido!)
```

---

### 2. Objeto vs Map: Quando Usar Cada Um

**✅ Use Objeto quando:**
- Você conhece as propriedades de antemão
- Propriedades são strings ou Symbols
- Você não precisa iterar frequentemente
- Ordem de inserção não importa (ES5) ou você usa ES6+

```javascript
// BOM - Objeto para dados estruturados fixos
const usuario = {
    nome: 'João',
    email: 'joao@email.com',
    idade: 30
};
```

**✅ Use Map quando:**
- Você precisa de chaves que não sejam strings
- Você adiciona/remove elementos dinamicamente
- Ordem de inserção importa
- Você precisa iterar frequentemente

```javascript
// BOM - Map para dados dinâmicos
const cache = new Map();
cache.set(usuarioId, dadosUsuario);
cache.set(dataHora, resultado);
```

**Performance:**
- Objetos são otimizados pelo JavaScript engine (V8, SpiderMonkey)
- Maps têm melhor performance para adicionar/remover elementos
- Para propriedades fixas: Objeto é mais rápido
- Para operações dinâmicas: Map pode ser mais rápido

---

### 3. Arrays: Métodos Eficientes vs Ineficientes

#### ❌ Evite métodos que criam novos arrays desnecessariamente

```javascript
// RUIM - Cria múltiplos arrays intermediários
const numeros = [1, 2, 3, 4, 5];
const resultado = numeros
    .filter(n => n > 2)
    .map(n => n * 2)
    .filter(n => n > 6)
    .map(n => n + 1);
```

**✅ Combine operações quando possível:**

```javascript
// BOM - Uma única iteração
const numeros = [1, 2, 3, 4, 5];
const resultado = [];
for (const n of numeros) {
    if (n > 2) {
        const dobrado = n * 2;
        if (dobrado > 6) {
            resultado.push(dobrado + 1);
        }
    }
}
```

**Ou use reduce() para operações complexas:**

```javascript
// BOM - Uma única iteração com reduce
const resultado = numeros.reduce((acc, n) => {
    if (n > 2) {
        const dobrado = n * 2;
        if (dobrado > 6) {
            acc.push(dobrado + 1);
        }
    }
    return acc;
}, []);
```

---

### 4. Arrays: Mutação vs Imutabilidade

**❌ Evite mutar arrays diretamente quando não necessário:**

```javascript
// RUIM - Modifica o array original
const numeros = [1, 2, 3, 4, 5];
numeros.sort(); // Modifica o original!
console.log(numeros); // [1, 2, 3, 4, 5] - ordenado

// Se você precisar do original depois, está perdido!
```

**✅ Crie uma cópia antes de modificar:**

```javascript
// BOM - Cria cópia antes de ordenar
const numeros = [1, 2, 3, 4, 5];
const numerosOrdenados = [...numeros].sort(); // Spread operator
// ou
const numerosOrdenados2 = numeros.slice().sort(); // slice()

console.log(numeros); // [1, 2, 3, 4, 5] - original preservado
console.log(numerosOrdenados); // [1, 2, 3, 4, 5] - ordenado
```

**Por quê?**
- Evita efeitos colaterais
- Facilita debugging
- Permite comparar antes/depois
- Alinha com princípios de programação funcional

---

### 5. JSON: Performance e Validação

**❌ Evite fazer parse de JSON sem tratamento de erros:**

```javascript
// RUIM - Pode quebrar a aplicação
const dados = JSON.parse(dadosRecebidos); // E se for inválido?
```

**✅ Sempre trate erros ao fazer parse:**

```javascript
// BOM - Tratamento de erros
function parseJSONSeguro(jsonString) {
    try {
        return JSON.parse(jsonString);
    } catch (erro) {
        console.error('Erro ao fazer parse do JSON:', erro);
        return null; // ou um valor padrão
    }
}

const dados = parseJSONSeguro(dadosRecebidos);
if (dados) {
    // Usar dados
}
```

**✅ Valide JSON antes de processar:**

```javascript
// BOM - Validação antes de processar
function validarEProcessarJSON(jsonString) {
    if (!jsonString || typeof jsonString !== 'string') {
        return null;
    }
    
    try {
        const dados = JSON.parse(jsonString);
        // Validação adicional se necessário
        if (dados && typeof dados === 'object') {
            return dados;
        }
        return null;
    } catch (erro) {
        return null;
    }
}
```

---

## 🛠️ Boas Práticas: Organização e Estrutura

### 1. Nomenclatura Clara

**❌ Evite nomes genéricos:**

```javascript
// RUIM
const arr = [1, 2, 3];
const obj = { a: 1, b: 2 };
const map = new Map();
```

**✅ Use nomes descritivos:**

```javascript
// BOM
const produtos = [1, 2, 3];
const dadosUsuario = { nome: 'João', idade: 30 };
const cacheProdutos = new Map();
```

---

### 2. Inicialização Adequada

**❌ Evite arrays/objetos vazios sem propósito:**

```javascript
// RUIM - Array vazio que nunca é usado
const dados = [];
// ... código que nunca preenche 'dados'
```

**✅ Inicialize com valores ou documente o propósito:**

```javascript
// BOM - Inicialização clara
const produtos = []; // Será preenchido com produtos da API

// ou melhor ainda, inicialize com valores padrão
const configuracoes = {
    tema: 'claro',
    idioma: 'pt-BR',
    notificacoes: true
};
```

---

### 3. Estruturas de Dados Aninhadas

**❌ Evite aninhamento excessivo:**

```javascript
// RUIM - Muito aninhado, difícil de ler
const dados = {
    usuario: {
        endereco: {
            cidade: {
                estado: {
                    pais: {
                        nome: 'Brasil'
                    }
                }
            }
        }
    }
};
```

**✅ Simplifique ou use variáveis intermediárias:**

```javascript
// BOM - Mais legível
const pais = dados.usuario.endereco.cidade.estado.pais;
const nomePais = pais.nome;

// ou melhor ainda, reestruture os dados
const endereco = {
    cidade: 'São Paulo',
    estado: 'SP',
    pais: 'Brasil'
};
```

---

### 4. Uso de Constantes para Estruturas Vazias

**❌ Evite criar novas estruturas vazias repetidamente:**

```javascript
// RUIM - Cria novo array a cada chamada
function processar() {
    const resultados = [];
    // ...
    return resultados;
}
```

**✅ Reutilize quando apropriado (cuidado com mutação!):**

```javascript
// BOM - Mas cuidado: não mutar!
const ARRAY_VAZIO = Object.freeze([]); // Imutável

function processar() {
    const resultados = []; // Nova instância quando necessário
    // ...
    return resultados;
}
```

---

## 🔒 Segurança: Validação e Sanitização

### 1. Validar Dados Antes de Usar

**❌ Não confie em dados externos:**

```javascript
// RUIM - Assume que dados são válidos
function processarUsuario(dados) {
    const nome = dados.nome;
    const email = dados.email;
    // Usa sem validar
}
```

**✅ Sempre valide:**

```javascript
// BOM - Validação antes de usar
function processarUsuario(dados) {
    if (!dados || typeof dados !== 'object') {
        throw new Error('Dados inválidos');
    }
    
    const nome = dados.nome;
    if (!nome || typeof nome !== 'string' || nome.trim() === '') {
        throw new Error('Nome inválido');
    }
    
    const email = dados.email;
    if (!email || !email.includes('@')) {
        throw new Error('Email inválido');
    }
    
    // Agora pode usar com segurança
}
```

---

### 2. Sanitizar Dados de JSON

**❌ Não execute código de JSON não confiável:**

```javascript
// RUIM - Nunca faça isso!
const dados = eval('(' + jsonString + ')'); // PERIGOSO!
```

**✅ Use sempre JSON.parse():**

```javascript
// BOM - Seguro
const dados = JSON.parse(jsonString);
```

**Por quê?**
- `eval()` executa código JavaScript - pode ser explorado
- `JSON.parse()` apenas converte dados - muito mais seguro

---

## 🧹 Gerenciamento de Memória

### 1. Limpar Estruturas Não Utilizadas

**✅ Limpe Map/Set quando não precisar mais:**

```javascript
// BOM - Limpar cache antigo
const cache = new Map();

function limparCacheAntigo() {
    const agora = Date.now();
    for (const [chave, valor] of cache.entries()) {
        if (valor.timestamp < agora - 3600000) { // 1 hora
            cache.delete(chave);
        }
    }
}
```

---

### 2. Evitar Vazamentos de Memória

**❌ Evite referências circulares:**

```javascript
// RUIM - Referência circular
const obj1 = { nome: 'Objeto 1' };
const obj2 = { nome: 'Objeto 2' };
obj1.ref = obj2;
obj2.ref = obj1; // Referência circular!
// Esses objetos nunca serão coletados pelo garbage collector
```

**✅ Use WeakMap/WeakSet quando apropriado:**

```javascript
// BOM - WeakMap permite garbage collection
const cache = new WeakMap();

function cachear(objeto, dados) {
    cache.set(objeto, dados);
    // Quando 'objeto' não for mais referenciado, será coletado
}
```

---

## 📊 Debugging: Ferramentas e Técnicas

### 1. Inspecionar Estruturas de Dados

```javascript
// Use console.table() para arrays de objetos
const usuarios = [
    { nome: 'João', idade: 30 },
    { nome: 'Maria', idade: 25 }
];
console.table(usuarios); // Visualização em tabela

// Use console.dir() para objetos complexos
console.dir(usuarios, { depth: null }); // Mostra tudo

// Use JSON.stringify() para ver estrutura completa
console.log(JSON.stringify(usuarios, null, 2));
```

---

### 2. Verificar Tamanho e Conteúdo

```javascript
// Arrays
console.log(array.length);
console.log(Array.isArray(array)); // Verificar se é array

// Map/Set
console.log(map.size);
console.log(set.size);

// Objetos
console.log(Object.keys(obj).length);
console.log(Object.keys(obj)); // Ver todas as chaves
```

---

## 🎯 Padrões de Código: Clean Code

### 1. DRY (Don't Repeat Yourself)

**❌ Evite repetir código:**

```javascript
// RUIM - Código repetido
const usuarios1 = usuarios.filter(u => u.idade > 18);
const usuarios2 = usuarios.filter(u => u.idade > 18);
const usuarios3 = usuarios.filter(u => u.idade > 18);
```

**✅ Extraia para função:**

```javascript
// BOM - Função reutilizável
function filtrarMaioresDeIdade(usuarios, idadeMinima = 18) {
    return usuarios.filter(u => u.idade >= idadeMinima);
}

const usuariosMaiores = filtrarMaioresDeIdade(usuarios);
```

---

### 2. Separação de Responsabilidades

**❌ Evite funções que fazem muitas coisas:**

```javascript
// RUIM - Faz muitas coisas
function processarDados(dados) {
    const json = JSON.parse(dados);
    const filtrado = json.filter(/* ... */);
    const mapeado = filtrado.map(/* ... */);
    const ordenado = mapeado.sort(/* ... */);
    localStorage.setItem('dados', JSON.stringify(ordenado));
    return ordenado;
}
```

**✅ Separe em funções menores:**

```javascript
// BOM - Funções com responsabilidade única
function parsearJSON(dados) {
    return JSON.parse(dados);
}

function filtrarDados(dados) {
    return dados.filter(/* ... */);
}

function mapearDados(dados) {
    return dados.map(/* ... */);
}

function ordenarDados(dados) {
    return dados.sort(/* ... */);
}

function salvarDados(dados) {
    localStorage.setItem('dados', JSON.stringify(dados));
}

function processarDados(dados) {
    const parseado = parsearJSON(dados);
    const filtrado = filtrarDados(parseado);
    const mapeado = mapearDados(filtrado);
    const ordenado = ordenarDados(mapeado);
    salvarDados(ordenado);
    return ordenado;
}
```

---

## 🚀 Otimizações Avançadas

### 1. Lazy Evaluation

**✅ Use generators para arrays grandes:**

```javascript
// BOM - Gera valores sob demanda
function* gerarNumeros(limite) {
    for (let i = 0; i < limite; i++) {
        yield i;
    }
}

// Não cria array completo na memória
for (const numero of gerarNumeros(1000000)) {
    // Processa um por vez
}
```

---

### 2. Memoização com Map

**✅ Cache resultados de funções custosas:**

```javascript
// BOM - Memoização
const cache = new Map();

function calcularCustoso(n) {
    if (cache.has(n)) {
        return cache.get(n);
    }
    
    // Cálculo custoso
    const resultado = /* cálculo complexo */;
    cache.set(n, resultado);
    return resultado;
}
```

---

## 📝 Checklist de Boas Práticas

Ao trabalhar com estruturas de dados, sempre:

- [ ] Escolha a estrutura correta para o caso de uso
- [ ] Use Set para verificações de existência frequentes
- [ ] Use Map para chaves dinâmicas ou não-string
- [ ] Valide dados antes de usar
- [ ] Trate erros ao fazer parse de JSON
- [ ] Evite mutar estruturas quando não necessário
- [ ] Use nomes descritivos
- [ ] Limpe estruturas não utilizadas
- [ ] Evite aninhamento excessivo
- [ ] Documente estruturas complexas

---

## 🎯 Resumo

**Performance:**
- Set.has() é muito mais rápido que Array.includes()
- Map é melhor para operações dinâmicas
- Evite criar arrays intermediários desnecessariamente
- Crie cópias antes de mutar

**Boas Práticas:**
- Nomenclatura clara e descritiva
- Validação de dados
- Separação de responsabilidades
- Tratamento de erros

**Segurança:**
- Sempre use JSON.parse(), nunca eval()
- Valide dados externos
- Sanitize entrada do usuário

**Memória:**
- Limpe estruturas não utilizadas
- Use WeakMap/WeakSet quando apropriado
- Evite referências circulares

---

## 🚀 Próximos Passos

Agora que você entende performance e boas práticas:
- Aplique esses conceitos nos exercícios
- Pratique escolhendo a estrutura correta
- Sempre pense em performance ao escrever código
- Continue para a análise de desempenho após completar os exercícios

Boa prática! 💪



