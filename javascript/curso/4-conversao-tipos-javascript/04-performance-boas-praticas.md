# Aula 4 - Performance, Boas Práticas e Otimização: Conversão de Tipos

## 🎯 Introdução

Conversão de tipos em JavaScript não é apenas sobre transformar valores - é sobre **fazer isso de forma eficiente, segura e previsível**. Nesta aula, você aprenderá como converter tipos de forma profissional, evitando armadilhas de performance e bugs comuns.

---

## ⚡ Performance: Impacto das Conversões

### 1. Escolha do Método de Conversão

#### Number() vs parseInt() vs parseFloat() vs Operador +

**Benchmark de Performance:**

```javascript
// Teste de performance - conversão de string para número
let valor = "123456789";

console.time("Number()");
for (let i = 0; i < 1000000; i++) {
    Number(valor);
}
console.timeEnd("Number()"); // ~15-20ms

console.time("parseInt()");
for (let i = 0; i < 1000000; i++) {
    parseInt(valor, 10);
}
console.timeEnd("parseInt()"); // ~25-30ms (mais lento)

console.time("parseFloat()");
for (let i = 0; i < 1000000; i++) {
    parseFloat(valor);
}
console.timeEnd("parseFloat()"); // ~30-35ms (mais lento)

console.time("Operador +");
for (let i = 0; i < 1000000; i++) {
    +valor;
}
console.timeEnd("Operador +"); // ~10-15ms (mais rápido!)
```

**Resultado:** O operador unário `+` é geralmente o mais rápido, mas `Number()` é mais legível e explícito.

**Recomendação:**
- Use `Number()` para código legível e manutenível
- Use `+valor` apenas em código onde performance é crítica e legibilidade não é comprometida
- Use `parseInt()`/`parseFloat()` quando precisar de comportamento específico (inteiros, decimais)

#### String() vs .toString() vs Template Literals vs Concatenação

**Benchmark de Performance:**

```javascript
let valor = 123456789;

console.time("String()");
for (let i = 0; i < 1000000; i++) {
    String(valor);
}
console.timeEnd("String()"); // ~20-25ms

console.time(".toString()");
for (let i = 0; i < 1000000; i++) {
    valor.toString();
}
console.timeEnd(".toString()"); // ~15-20ms (mais rápido)

console.time("Template Literal");
for (let i = 0; i < 1000000; i++) {
    `${valor}`;
}
console.timeEnd("Template Literal"); // ~25-30ms

console.time("Concatenação");
for (let i = 0; i < 1000000; i++) {
    valor + "";
}
console.timeEnd("Concatenação"); // ~10-15ms (mais rápido!)
```

**Resultado:** Concatenação com `""` é geralmente mais rápida, mas `String()` é mais seguro (funciona com null/undefined).

**Recomendação:**
- Use `String()` quando precisar de segurança (null/undefined)
- Use `.toString()` quando tiver certeza que o valor não é null/undefined
- Use template literals para legibilidade em strings complexas
- Evite concatenação repetida em loops (use array + join)

### 2. Evitando Conversões Desnecessárias

**❌ Ruim - Conversão repetida:**

```javascript
// RUIM - Converte a cada iteração
function processarNumeros(strings) {
    let soma = 0;
    for (let str of strings) {
        soma += Number(str); // Conversão a cada iteração
    }
    return soma;
}
```

**✅ Bom - Converte uma vez:**

```javascript
// BOM - Converte uma vez no início
function processarNumeros(strings) {
    // Valida e converte tudo de uma vez
    let numeros = strings.map(str => {
        let num = Number(str);
        if (isNaN(num)) {
            throw new TypeError(`Valor inválido: ${str}`);
        }
        return num;
    });
    
    // Agora trabalha apenas com números
    return numeros.reduce((soma, num) => soma + num, 0);
}
```

**Por quê?**
- Reduz chamadas de função repetidas
- Facilita otimização pelo JavaScript engine
- Código mais limpo e manutenível

### 3. Comparações: == vs ===

**Performance:**

```javascript
// Teste de performance
let valor = "123";

console.time("== (com coerção)");
for (let i = 0; i < 1000000; i++) {
    if (valor == 123) {}
}
console.timeEnd("== (com coerção)"); // ~15-20ms

console.time("=== (sem coerção)");
for (let i = 0; i < 1000000; i++) {
    if (valor === 123) {}
}
console.timeEnd("=== (sem coerção)"); // ~10-15ms (mais rápido!)
```

**Resultado:** `===` é mais rápido porque não precisa fazer coerção de tipos.

**Além da performance, `===` é mais seguro:**
- Evita bugs sutis de coerção
- Código mais previsível
- Melhor para debugging

**Recomendação:** **Sempre use `===` e `!==`** a menos que tenha uma razão muito específica para usar `==`.

---

## 🛡️ Segurança: Validação e Sanitização

### 1. Validação Antes de Converter

**❌ Ruim - Sem validação:**

```javascript
// PERIGOSO - Aceita qualquer entrada
function calcularTotal(preco, quantidade) {
    return Number(preco) * Number(quantidade);
}

// Pode retornar NaN sem avisar!
calcularTotal("abc", "2"); // NaN
```

**✅ Bom - Com validação:**

```javascript
// SEGURO - Valida antes de converter
function calcularTotal(preco, quantidade) {
    // Validação de entrada
    if (preco == null || quantidade == null) {
        throw new Error("Preço e quantidade são obrigatórios");
    }
    
    // Conversão e validação
    let precoNum = Number(preco);
    let qtdNum = Number(quantidade);
    
    if (isNaN(precoNum) || isNaN(qtdNum)) {
        throw new TypeError("Preço e quantidade devem ser números válidos");
    }
    
    if (precoNum < 0 || qtdNum < 0) {
        throw new RangeError("Preço e quantidade devem ser positivos");
    }
    
    return precoNum * qtdNum;
}
```

### 2. Sanitização de Entrada do Usuário

**❌ Ruim - Aceita entrada sem sanitizar:**

```javascript
// PERIGOSO - Não sanitiza entrada
function processarIdade(idadeInput) {
    let idade = Number(idadeInput);
    return idade;
}

// Problemas:
processarIdade("  25  ");  // 25 (espaços são ignorados, mas pode ser confuso)
processarIdade("25abc");   // NaN (mas não avisa)
processarIdade("");        // 0 (string vazia vira 0!)
```

**✅ Bom - Sanitiza antes de converter:**

```javascript
// SEGURO - Sanitiza entrada
function processarIdade(idadeInput) {
    // 1. Verifica se é string e remove espaços
    if (typeof idadeInput === 'string') {
        idadeInput = idadeInput.trim();
        
        // Rejeita strings vazias
        if (idadeInput === '') {
            throw new Error("Idade não pode ser vazia");
        }
    }
    
    // 2. Converte
    let idade = Number(idadeInput);
    
    // 3. Valida resultado
    if (isNaN(idade)) {
        throw new TypeError(`Idade inválida: ${idadeInput}`);
    }
    
    if (idade < 0 || idade > 150) {
        throw new RangeError("Idade deve estar entre 0 e 150");
    }
    
    return idade;
}
```

### 3. Proteção Contra XSS em Conversões de String

**⚠️ Cuidado ao converter dados do usuário para string:**

```javascript
// PERIGOSO - Se usado em innerHTML pode causar XSS
let userInput = "<script>alert('XSS')</script>";
let html = String(userInput); // Contém código malicioso!

// SEGURO - Sanitiza antes de usar em HTML
function sanitizarHTML(str) {
    return String(str)
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#x27;");
}

// Ou use textContent em vez de innerHTML
element.textContent = String(userInput); // Seguro!
```

---

## 📋 Boas Práticas

### 1. Sempre Use Conversão Explícita

**❌ Ruim - Conversão implícita:**

```javascript
// CONFUSO - Conversão implícita
function somar(a, b) {
    return a + b; // Pode concatenar strings!
}

somar("10", 5); // "105" (erro!)
```

**✅ Bom - Conversão explícita:**

```javascript
// CLARO - Conversão explícita
function somar(a, b) {
    let numA = Number(a);
    let numB = Number(b);
    
    if (isNaN(numA) || isNaN(numB)) {
        throw new TypeError("Ambos os valores devem ser números");
    }
    
    return numA + numB;
}

somar("10", 5); // 15 (correto!)
```

### 2. Use Constantes para Valores Mágicos

**❌ Ruim - Valores mágicos:**

```javascript
// CONFUSO - O que significa 0, "", null?
if (valor == 0 || valor == "" || valor == null) {
    // ...
}
```

**✅ Bom - Constantes nomeadas:**

```javascript
// CLARO - Constantes nomeadas
const VALOR_VAZIO = "";
const VALOR_ZERO = 0;
const VALOR_NULO = null;

if (valor === VALOR_ZERO || valor === VALOR_VAZIO || valor === VALOR_NULO) {
    // ...
}

// Ou melhor ainda - função helper
function isEmpty(valor) {
    return valor === "" || valor === 0 || valor === null || valor === undefined;
}
```

### 3. Funções Helper para Conversões Comuns

**✅ Crie funções reutilizáveis:**

```javascript
// Funções helper para conversões comuns
const Conversor = {
    // Converte para número com validação
    paraNumero(valor, valorPadrao = null) {
        if (valor == null) return valorPadrao;
        
        let numero = Number(valor);
        return isNaN(numero) ? valorPadrao : numero;
    },
    
    // Converte para string com tratamento de null/undefined
    paraString(valor, valorPadrao = "") {
        if (valor == null) return valorPadrao;
        return String(valor);
    },
    
    // Converte para boolean explicitamente
    paraBoolean(valor) {
        return Boolean(valor);
    },
    
    // Valida se é número válido
    ehNumeroValido(valor) {
        return typeof valor === 'number' && !isNaN(valor) && isFinite(valor);
    }
};

// Uso
let idade = Conversor.paraNumero(entradaUsuario, 0);
let nome = Conversor.paraString(entradaUsuario, "Anônimo");
```

### 4. Tratamento de Erros Consistente

**❌ Ruim - Retorna valores inesperados:**

```javascript
// CONFUSO - Retorna NaN sem avisar
function converterIdade(valor) {
    return Number(valor); // Pode retornar NaN!
}
```

**✅ Bom - Tratamento de erros explícito:**

```javascript
// CLARO - Trata erros explicitamente
function converterIdade(valor) {
    if (valor == null) {
        throw new Error("Valor não pode ser null ou undefined");
    }
    
    let idade = Number(valor);
    
    if (isNaN(idade)) {
        throw new TypeError(`Não foi possível converter "${valor}" para número`);
    }
    
    return idade;
}

// Ou retorna null em vez de lançar erro (depende do caso)
function converterIdadeSeguro(valor) {
    try {
        return converterIdade(valor);
    } catch (error) {
        console.error("Erro ao converter idade:", error);
        return null; // Valor padrão seguro
    }
}
```

---

## 🚫 O Que NÃO Fazer

### 1. Não Use == para Comparações

**❌ Ruim:**

```javascript
// PERIGOSO - Comportamento inesperado
if (valor == 0) { }        // Pode ser true para "", "0", null, false
if (valor == null) { }    // true para null E undefined
if (valor == "") { }      // true para 0, false, null
```

**✅ Bom:**

```javascript
// SEGURO - Comportamento previsível
if (valor === 0) { }      // Apenas 0
if (valor === null) { }   // Apenas null
if (valor === "") { }     // Apenas string vazia
```

### 2. Não Confie em Conversões Implícitas em Operações Críticas

**❌ Ruim:**

```javascript
// PERIGOSO - Pode dar resultado errado
let total = preco + quantidade; // Se preco for string, concatena!
```

**✅ Bom:**

```javascript
// SEGURO - Conversão explícita
let total = Number(preco) + Number(quantidade);
```

### 3. Não Use parseInt() sem Especificar a Base

**❌ Ruim:**

```javascript
// PERIGOSO - Comportamento inconsistente
parseInt("08");    // Pode ser 0 ou 8 dependendo do navegador
parseInt("0x10");  // 16 (hexadecimal automático)
```

**✅ Bom:**

```javascript
// SEGURO - Sempre especifique a base
parseInt("08", 10);   // Sempre 8
parseInt("0x10", 16); // 16 (explícito)
```

### 4. Não Converta Objetos para String com .toString()

**❌ Ruim:**

```javascript
// INÚTIL - Sempre retorna "[object Object]"
let obj = {a: 1, b: 2};
console.log(obj.toString()); // "[object Object]" (não útil)
```

**✅ Bom:**

```javascript
// ÚTIL - Use JSON.stringify()
let obj = {a: 1, b: 2};
console.log(JSON.stringify(obj)); // '{"a":1,"b":2}' (útil!)
```

### 5. Não Ignore NaN em Validações

**❌ Ruim:**

```javascript
// PERIGOSO - NaN passa na validação
let numero = Number("abc");
if (numero) { // NaN é falsy, mas...
    // Não entra aqui, mas e se usar em cálculo?
}
let resultado = numero * 10; // NaN (propaga!)
```

**✅ Bom:**

```javascript
// SEGURO - Valida explicitamente
let numero = Number("abc");
if (isNaN(numero)) {
    throw new TypeError("Valor inválido");
}
// Agora sabemos que numero é válido
```

---

## 🔍 Debugging: Identificando Problemas de Conversão

### 1. Use typeof para Debugging

```javascript
function debugTipo(valor) {
    console.log("Valor:", valor);
    console.log("Tipo:", typeof valor);
    console.log("É NaN?", isNaN(valor));
    console.log("É null?", valor === null);
    console.log("É undefined?", valor === undefined);
    console.log("Valor após Number():", Number(valor));
    console.log("Valor após String():", String(valor));
    console.log("Valor após Boolean():", Boolean(valor));
}

// Use para debugar problemas de conversão
debugTipo(entradaProblema);
```

### 2. Validação em Desenvolvimento

```javascript
// Função de validação para desenvolvimento
function validarConversao(valor, tipoEsperado) {
    if (process.env.NODE_ENV === 'development') {
        let tipoAtual = typeof valor;
        
        if (tipoAtual !== tipoEsperado) {
            console.warn(
                `⚠️ Conversão de tipo: esperado ${tipoEsperado}, recebido ${tipoAtual}`,
                valor
            );
        }
    }
    
    // Conversão
    switch(tipoEsperado) {
        case 'number':
            return Number(valor);
        case 'string':
            return String(valor);
        case 'boolean':
            return Boolean(valor);
        default:
            return valor;
    }
}
```

### 3. Logging de Conversões

```javascript
// Wrapper para logging de conversões
function converterComLog(valor, metodo, resultado) {
    if (process.env.NODE_ENV === 'development') {
        console.log(`[Conversão] ${metodo}(${JSON.stringify(valor)}) = ${JSON.stringify(resultado)}`);
    }
    return resultado;
}

// Uso
let numero = converterComLog("123", "Number", Number("123"));
```

---

## 📊 Performance: Métricas e Otimização

### 1. Medindo Performance de Conversões

```javascript
// Função helper para medir performance
function medirPerformance(nome, funcao, iteracoes = 1000000) {
    console.time(nome);
    for (let i = 0; i < iteracoes; i++) {
        funcao();
    }
    console.timeEnd(nome);
}

// Teste diferentes métodos
let valor = "123456";

medirPerformance("Number()", () => Number(valor));
medirPerformance("parseInt()", () => parseInt(valor, 10));
medirPerformance("Operador +", () => +valor);
```

### 2. Cache de Conversões

**✅ Quando possível, cache conversões:**

```javascript
// BOM - Cache conversões quando o valor não muda
class ConversorCacheado {
    constructor() {
        this.cache = new Map();
    }
    
    paraNumero(valor) {
        // Se já convertemos este valor, retorna do cache
        if (this.cache.has(valor)) {
            return this.cache.get(valor);
        }
        
        let numero = Number(valor);
        this.cache.set(valor, numero);
        return numero;
    }
    
    limparCache() {
        this.cache.clear();
    }
}
```

### 3. Lazy Conversion (Conversão Preguiçosa)

**✅ Converta apenas quando necessário:**

```javascript
// BOM - Converte apenas quando realmente precisa
class ValorLazy {
    constructor(valor) {
        this._valorOriginal = valor;
        this._valorConvertido = null;
        this._tipoConvertido = null;
    }
    
    comoNumero() {
        if (this._tipoConvertido !== 'number') {
            this._valorConvertido = Number(this._valorOriginal);
            this._tipoConvertido = 'number';
        }
        return this._valorConvertido;
    }
    
    comoString() {
        if (this._tipoConvertido !== 'string') {
            this._valorConvertido = String(this._valorOriginal);
            this._tipoConvertido = 'string';
        }
        return this._valorConvertido;
    }
}
```

---

## 🎯 Resumo: Checklist de Boas Práticas

### Conversões

- [ ] ✅ Sempre use conversões explícitas em código crítico
- [ ] ✅ Valide entrada antes de converter
- [ ] ✅ Trate erros de conversão (NaN, null, undefined)
- [ ] ✅ Use `Number()` para conversões numéricas gerais
- [ ] ✅ Use `parseInt()` com base numérica especificada
- [ ] ✅ Use `String()` para segurança com null/undefined
- [ ] ✅ Use `Boolean()` ou `!!` para conversões booleanas

### Comparações

- [ ] ✅ Sempre use `===` e `!==` em vez de `==` e `!=`
- [ ] ✅ Use `Number.isNaN()` em vez de `isNaN()`
- [ ] ✅ Use `Array.isArray()` para verificar arrays

### Performance

- [ ] ✅ Evite conversões repetidas em loops
- [ ] ✅ Cache conversões quando possível
- [ ] ✅ Use métodos mais rápidos quando performance é crítica
- [ ] ✅ Meça performance antes de otimizar

### Segurança

- [ ] ✅ Valide e sanitize entrada do usuário
- [ ] ✅ Trate valores null/undefined explicitamente
- [ ] ✅ Use sanitização para prevenir XSS
- [ ] ✅ Lance erros descritivos em vez de retornar NaN silenciosamente

### Manutenibilidade

- [ ] ✅ Crie funções helper para conversões comuns
- [ ] ✅ Use constantes nomeadas em vez de valores mágicos
- [ ] ✅ Documente comportamento de conversões complexas
- [ ] ✅ Use TypeScript ou JSDoc para tipos (quando possível)

---

## 🚀 Conclusão

Conversão de tipos em JavaScript é uma habilidade fundamental que, quando feita corretamente, resulta em:

- ✅ **Código mais seguro** - Menos bugs e vulnerabilidades
- ✅ **Código mais rápido** - Menos conversões desnecessárias
- ✅ **Código mais legível** - Conversões explícitas são mais claras
- ✅ **Código mais manutenível** - Fácil de entender e modificar

**Lembre-se:** Em JavaScript, a conversão de tipos é poderosa, mas pode ser perigosa. Sempre prefira ser explícito, validar entrada e usar comparações estritas!

---

**Próximo passo:** Aplique essas práticas em seus projetos e sempre questione: "Esta conversão é segura? É eficiente? É clara?"

