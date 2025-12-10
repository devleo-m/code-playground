# Aula 3 - Performance, Boas Práticas e Otimização: Tipos de Dados

## 🎯 Introdução

Entender tipos de dados não é apenas sobre saber o que cada um faz - é sobre **usá-los de forma eficiente, segura e manutenível**. Nesta aula, você aprenderá como trabalhar com tipos de dados de forma profissional.

---

## ⚡ Performance: Impacto dos Tipos de Dados

### 1. Escolha do Tipo Correto

#### Number vs BigInt

**❌ Evite usar BigInt quando não for necessário:**

```javascript
// RUIM - BigInt é mais lento e ocupa mais memória
let contador = 0n;
for (let i = 0n; i < 1000000n; i++) {
    contador += 1n;
}

// BOM - Use Number para operações comuns
let contador = 0;
for (let i = 0; i < 1000000; i++) {
    contador += 1;
}
```

**Por quê?**
- `BigInt` é mais lento em operações aritméticas
- Ocupa mais memória (geralmente 2x mais)
- Não pode ser usado com `Number` diretamente (requer conversões)
- **Use apenas quando realmente precisar de números maiores que `Number.MAX_SAFE_INTEGER`**

#### String: Concatenação vs Template Literals

**❌ Evite concatenação repetida em loops:**

```javascript
// RUIM - Cria novas strings a cada iteração (muito lento!)
let resultado = "";
for (let i = 0; i < 10000; i++) {
    resultado += "item " + i + ", ";
}

// BOM - Use array e join (muito mais rápido)
let items = [];
for (let i = 0; i < 10000; i++) {
    items.push(`item ${i}`);
}
let resultado = items.join(", ");
```

**Por quê?**
- Strings são **imutáveis** em JavaScript
- Cada concatenação cria uma **nova string** na memória
- Em loops grandes, isso pode causar **garbage collection** excessivo
- Usar array + `join()` é muito mais eficiente

**Benchmark exemplo:**
```javascript
// Teste de performance
console.time("Concatenação");
let str1 = "";
for (let i = 0; i < 100000; i++) {
    str1 += i;
}
console.timeEnd("Concatenação"); // ~500ms

console.time("Array + Join");
let arr = [];
for (let i = 0; i < 100000; i++) {
    arr.push(i);
}
let str2 = arr.join("");
console.timeEnd("Array + Join"); // ~50ms (10x mais rápido!)
```

### 2. Verificação de Tipos

**❌ Evite verificar tipos repetidamente:**

```javascript
// RUIM - typeof é chamado a cada iteração
function processarItems(items) {
    for (let item of items) {
        if (typeof item === "string") {
            // processar
        }
    }
}

// BOM - Verifique o tipo uma vez, fora do loop
function processarItems(items) {
    if (!Array.isArray(items)) {
        throw new TypeError("Items deve ser um array");
    }
    
    for (let item of items) {
        // Já sabemos que é array, processe diretamente
    }
}
```

**✅ Use verificações de tipo eficientes:**

```javascript
// Para arrays - use Array.isArray() (mais rápido e confiável)
Array.isArray(valor); // ✅ BOM
valor instanceof Array; // ⚠️ Funciona, mas pode falhar em múltiplos contextos

// Para números - use Number.isInteger() ou Number.isNaN()
Number.isInteger(42); // ✅ BOM
Number.isNaN(NaN); // ✅ BOM (mais confiável que isNaN())

// Para null - use comparação direta
valor === null; // ✅ BOM
valor == null; // ⚠️ Funciona, mas verifica null E undefined
```

### 3. Conversão de Tipos

**❌ Evite conversões desnecessárias:**

```javascript
// RUIM - Conversão a cada iteração
function somarNumeros(numeros) {
    let soma = 0;
    for (let num of numeros) {
        soma += Number(num); // Conversão desnecessária se já for número
    }
    return soma;
}

// BOM - Valide uma vez, converta uma vez
function somarNumeros(numeros) {
    if (!Array.isArray(numeros)) {
        throw new TypeError("Numeros deve ser um array");
    }
    
    let soma = 0;
    for (let num of numeros) {
        if (typeof num !== "number") {
            num = Number(num);
            if (isNaN(num)) {
                throw new TypeError(`Valor inválido: ${num}`);
            }
        }
        soma += num;
    }
    return soma;
}
```

**✅ Prefira conversões explícitas:**

```javascript
// RUIM - Conversão implícita (pode causar bugs)
let resultado = "5" + 3; // "53" (não é o esperado!)

// BOM - Conversão explícita (claro e seguro)
let resultado = Number("5") + 3; // 8
// ou
let resultado = parseInt("5", 10) + 3; // 8
```

---

## 🛡️ Boas Práticas: Segurança e Confiabilidade

### 1. Validação de Dados

**Sempre valide dados antes de usar:**

```javascript
// ❌ RUIM - Assume que o dado está correto
function calcularIdade(anoNascimento) {
    return 2024 - anoNascimento; // E se for string? E se for null?
}

// ✅ BOM - Valida antes de usar
function calcularIdade(anoNascimento) {
    // Validação de tipo
    if (typeof anoNascimento !== "number") {
        throw new TypeError("anoNascimento deve ser um número");
    }
    
    // Validação de valor
    if (isNaN(anoNascimento) || !isFinite(anoNascimento)) {
        throw new Error("anoNascimento deve ser um número válido");
    }
    
    // Validação de lógica
    if (anoNascimento < 1900 || anoNascimento > 2024) {
        throw new RangeError("anoNascimento deve estar entre 1900 e 2024");
    }
    
    return 2024 - anoNascimento;
}
```

### 2. Tratamento de null e undefined

**Sempre verifique null/undefined antes de acessar propriedades:**

```javascript
// ❌ RUIM - Pode quebrar se usuario for null
function obterNomeUsuario(usuario) {
    return usuario.nome; // TypeError se usuario for null/undefined
}

// ✅ BOM - Verifica antes de acessar
function obterNomeUsuario(usuario) {
    if (usuario == null) { // Verifica null E undefined
        return "Usuário não encontrado";
    }
    return usuario.nome || "Nome não informado";
}

// ✅ AINDA MELHOR - Optional chaining (ES2020)
function obterNomeUsuario(usuario) {
    return usuario?.nome ?? "Nome não informado";
}
```

### 3. Uso Correto de null vs undefined

**Regra geral:**
- Use `undefined` para valores que **nunca foram definidos**
- Use `null` para valores que foram **intencionalmente definidos como vazios**

```javascript
// ✅ BOM - Padrão claro
function buscarUsuario(id) {
    // Se não encontrar, retorna null (intencionalmente vazio)
    if (!usuarioExiste(id)) {
        return null;
    }
    return { id, nome: "João" };
}

// ✅ BOM - Parâmetros opcionais usam undefined
function criarUsuario(nome, email) {
    // email é undefined se não fornecido
    return {
        nome,
        email: email ?? "sem-email@exemplo.com"
    };
}
```

### 4. Sanitização de Strings

**Sempre sanitize strings de entrada do usuário:**

```javascript
// ❌ RUIM - Aceita qualquer entrada
function processarNome(nome) {
    return nome.trim(); // E se for null? E se tiver caracteres perigosos?
}

// ✅ BOM - Valida e sanitiza
function processarNome(nome) {
    // Validação de tipo
    if (typeof nome !== "string") {
        throw new TypeError("Nome deve ser uma string");
    }
    
    // Sanitização
    let nomeLimpo = nome
        .trim() // Remove espaços no início/fim
        .replace(/[<>]/g, "") // Remove caracteres perigosos para HTML
        .substring(0, 100); // Limita tamanho
    
    // Validação de conteúdo
    if (nomeLimpo.length === 0) {
        throw new Error("Nome não pode estar vazio");
    }
    
    return nomeLimpo;
}
```

---

## 📐 Padrões de Código: Clean Code

### 1. Nomenclatura Clara

**Use nomes que indiquem o tipo quando necessário:**

```javascript
// ❌ RUIM - Nome não indica tipo
let data = "2024-01-15";

// ✅ BOM - Nome indica tipo ou uso
let dataString = "2024-01-15";
let dataNascimento = "2024-01-15";
```

**Convenções comuns:**
- `is...` ou `has...` para booleanos: `isAtivo`, `hasPermissao`
- `num...` ou `...Count` para números: `numItems`, `userCount`
- `str...` para strings (quando necessário): `strNome`
- `arr...` ou `...List` para arrays: `arrUsuarios`, `userList`
- `obj...` para objetos (quando necessário): `objConfig`

### 2. Constantes para Valores Mágicos

**Evite valores "mágicos" no código:**

```javascript
// ❌ RUIM - Valores mágicos
if (idade >= 18) {
    // O que significa 18?
}

// ✅ BOM - Constantes nomeadas
const IDADE_MINIMA_VOTACAO = 16;
const IDADE_MINIMA_ADULTO = 18;

if (idade >= IDADE_MINIMA_ADULTO) {
    // Fica claro o que significa
}
```

### 3. Type Guards (Guardiões de Tipo)

**Crie funções para verificar tipos:**

```javascript
// ✅ BOM - Type guards reutilizáveis
function isString(valor) {
    return typeof valor === "string";
}

function isNumber(valor) {
    return typeof valor === "number" && !isNaN(valor) && isFinite(valor);
}

function isPositiveInteger(valor) {
    return Number.isInteger(valor) && valor > 0;
}

// Uso
function processarDados(dados) {
    if (!isString(dados.nome)) {
        throw new TypeError("Nome deve ser uma string");
    }
    
    if (!isPositiveInteger(dados.idade)) {
        throw new TypeError("Idade deve ser um número inteiro positivo");
    }
    
    // Processar dados com segurança
}
```

---

## 🔒 Segurança: Prevenção de Vulnerabilidades

### 1. XSS (Cross-Site Scripting)

**Nunca confie em dados do usuário:**

```javascript
// ❌ PERIGOSO - Permite XSS
function exibirMensagem(mensagem) {
    document.getElementById("mensagem").innerHTML = mensagem;
    // Se mensagem contiver <script>, será executado!
}

// ✅ SEGURO - Sanitiza antes de exibir
function exibirMensagem(mensagem) {
    // Valida tipo
    if (typeof mensagem !== "string") {
        throw new TypeError("Mensagem deve ser uma string");
    }
    
    // Sanitiza (remove tags HTML perigosas)
    let mensagemSegura = mensagem
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#x27;");
    
    // Usa textContent ao invés de innerHTML quando possível
    document.getElementById("mensagem").textContent = mensagemSegura;
}

// ✅ AINDA MELHOR - Use bibliotecas de sanitização
// Exemplo com DOMPurify (biblioteca externa)
function exibirMensagem(mensagem) {
    let mensagemSegura = DOMPurify.sanitize(mensagem);
    document.getElementById("mensagem").innerHTML = mensagemSegura;
}
```

### 2. Validação de Entrada

**Sempre valide dados de formulários e APIs:**

```javascript
// ✅ BOM - Validação completa
function validarFormularioUsuario(dados) {
    const erros = [];
    
    // Valida nome
    if (!dados.nome || typeof dados.nome !== "string") {
        erros.push("Nome é obrigatório e deve ser uma string");
    } else if (dados.nome.trim().length < 2) {
        erros.push("Nome deve ter pelo menos 2 caracteres");
    } else if (dados.nome.length > 100) {
        erros.push("Nome não pode ter mais de 100 caracteres");
    }
    
    // Valida email
    if (!dados.email || typeof dados.email !== "string") {
        erros.push("Email é obrigatório");
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(dados.email)) {
        erros.push("Email inválido");
    }
    
    // Valida idade
    if (typeof dados.idade !== "number") {
        erros.push("Idade deve ser um número");
    } else if (!Number.isInteger(dados.idade)) {
        erros.push("Idade deve ser um número inteiro");
    } else if (dados.idade < 0 || dados.idade > 150) {
        erros.push("Idade deve estar entre 0 e 150");
    }
    
    return {
        valido: erros.length === 0,
        erros
    };
}
```

---

## 🧪 Debugging: Ferramentas e Técnicas

### 1. Console Methods para Tipos

**Use os métodos corretos do console:**

```javascript
// Para objetos - use console.dir() para ver estrutura completa
let objeto = { nome: "João", idade: 25 };
console.dir(objeto); // Mostra estrutura completa

// Para arrays - console.table() é muito útil
let usuarios = [
    { nome: "João", idade: 25 },
    { nome: "Maria", idade: 30 }
];
console.table(usuarios); // Tabela formatada

// Para verificar tipos - use console.assert()
let valor = "42";
console.assert(typeof valor === "number", "Valor deve ser número");
// Se falhar, mostra erro

// Para rastrear valores - use console.trace()
function processar(valor) {
    console.trace("Processando:", valor);
    // Mostra stack trace completo
}
```

### 2. Verificação de Tipos em Runtime

**Crie helpers para debug:**

```javascript
// Helper para verificar tipo e valor
function debugTipo(nome, valor) {
    console.group(`Debug: ${nome}`);
    console.log("Valor:", valor);
    console.log("Tipo:", typeof valor);
    console.log("É null?", valor === null);
    console.log("É undefined?", valor === undefined);
    
    if (Array.isArray(valor)) {
        console.log("É array, tamanho:", valor.length);
    }
    
    if (typeof valor === "object" && valor !== null) {
        console.log("Propriedades:", Object.keys(valor));
    }
    
    console.groupEnd();
}

// Uso
let dados = { nome: "João", idade: 25 };
debugTipo("dados", dados);
```

---

## 🎯 O que Deve Ser Utilizado

### ✅ Faça:

1. **Sempre valide tipos antes de usar**
   ```javascript
   if (typeof valor !== "string") {
       throw new TypeError("Esperado string");
   }
   ```

2. **Use conversões explícitas**
   ```javascript
   let numero = Number(texto); // Explícito
   ```

3. **Verifique null/undefined antes de acessar propriedades**
   ```javascript
   if (objeto?.propriedade) { // Optional chaining
       // usar
   }
   ```

4. **Use const para valores que não mudam**
   ```javascript
   const IDADE_MINIMA = 18;
   ```

5. **Sanitize dados do usuário**
   ```javascript
   let entradaSegura = entrada.trim().replace(/[<>]/g, "");
   ```

6. **Use type guards reutilizáveis**
   ```javascript
   function isString(valor) {
       return typeof valor === "string";
   }
   ```

---

## 🚫 O que NÃO Deve Ser Utilizado

### ❌ Evite:

1. **Conversões implícitas sem entender o comportamento**
   ```javascript
   // Evite - pode causar bugs
   let resultado = "5" + 3; // "53" não é o esperado!
   ```

2. **Usar == ao invés de ===**
   ```javascript
   // Evite - permite conversão implícita
   if (valor == null) { }
   
   // Use - comparação estrita
   if (valor === null || valor === undefined) { }
   ```

3. **Assumir que valores sempre existem**
   ```javascript
   // Evite - pode quebrar
   let nome = usuario.nome;
   
   // Use - verifica antes
   let nome = usuario?.nome ?? "Sem nome";
   ```

4. **Usar BigInt quando Number é suficiente**
   ```javascript
   // Evite - desnecessário e mais lento
   let contador = 0n;
   
   // Use - Number é suficiente
   let contador = 0;
   ```

5. **Concatenar strings em loops grandes**
   ```javascript
   // Evite - muito lento
   let resultado = "";
   for (let i = 0; i < 10000; i++) {
       resultado += i;
   }
   
   // Use - array + join
   let items = [];
   for (let i = 0; i < 10000; i++) {
       items.push(i);
   }
   let resultado = items.join("");
   ```

6. **Confiar em dados do usuário sem validação**
   ```javascript
   // Evite - perigoso
   document.innerHTML = dadosUsuario;
   
   // Use - sanitiza primeiro
   document.textContent = sanitizar(dadosUsuario);
   ```

---

## 📊 Gerenciamento de Memória

### 1. Evite Vazamentos com Strings

**Strings grandes podem causar problemas de memória:**

```javascript
// ❌ RUIM - Mantém referências desnecessárias
let dadosGrandes = "string muito grande...";
function processar() {
    // Mesmo depois de processar, dadosGrandes ainda está na memória
}

// ✅ BOM - Limpe referências quando não precisar mais
let dadosGrandes = "string muito grande...";
function processar() {
    // processar
    dadosGrandes = null; // Libera memória
}
```

### 2. Reutilize Objetos Quando Possível

**Evite criar objetos desnecessariamente:**

```javascript
// ❌ RUIM - Cria novo objeto a cada chamada
function criarConfig() {
    return {
        timeout: 5000,
        retries: 3
    };
}

// ✅ BOM - Reutiliza objeto constante
const CONFIG_PADRAO = {
    timeout: 5000,
    retries: 3
};

function criarConfig() {
    return { ...CONFIG_PADRAO }; // Spread apenas se precisar modificar
}
```

---

## 🎓 Resumo: Melhores Práticas para Tipos de Dados

### Checklist de Boas Práticas

- [ ] **Valide tipos** antes de usar valores
- [ ] **Use conversões explícitas** ao invés de implícitas
- [ ] **Verifique null/undefined** antes de acessar propriedades
- [ ] **Sanitize dados do usuário** para prevenir XSS
- [ ] **Use o tipo correto** para cada situação (Number vs BigInt)
- [ ] **Evite concatenação** de strings em loops grandes
- [ ] **Use type guards** para código mais limpo
- [ ] **Nomeie variáveis** de forma clara
- [ ] **Use const** para valores imutáveis
- [ ] **Prefira ===** ao invés de ==
- [ ] **Trate erros** adequadamente
- [ ] **Documente** tipos esperados em funções

---

## 🚀 Próximos Passos

Agora que você entende como trabalhar com tipos de dados de forma profissional, você está pronto para:
- ✅ Aplicar essas práticas em projetos reais
- ✅ Evitar bugs comuns relacionados a tipos
- ✅ Escrever código mais seguro e performático
- ✅ Depurar problemas de tipo mais facilmente

**Continue para a Análise de Desempenho após completar os exercícios!**





