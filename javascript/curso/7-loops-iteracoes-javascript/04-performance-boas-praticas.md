# Aula 7 - Performance, Boas Práticas e Otimização: Loops e Iterações

## ⚡ Performance: Escolhendo o Loop Correto

### Comparação de Performance entre Loops

Diferentes tipos de loops têm diferentes características de performance. Vamos analisar:

#### 1. `for` Tradicional vs `for...of`

```javascript
const array = new Array(1000000).fill(0).map((_, i) => i);

// Teste 1: for tradicional
console.time("for tradicional");
let soma1 = 0;
for (let i = 0; i < array.length; i++) {
  soma1 += array[i];
}
console.timeEnd("for tradicional");

// Teste 2: for tradicional com cache de length
console.time("for com cache");
let soma2 = 0;
for (let i = 0, len = array.length; i < len; i++) {
  soma2 += array[i];
}
console.timeEnd("for com cache");

// Teste 3: for...of
console.time("for...of");
let soma3 = 0;
for (const valor of array) {
  soma3 += valor;
}
console.timeEnd("for...of");
```

**Resultados típicos:**
- `for` com cache: Mais rápido (evita recalcular `length`)
- `for` tradicional: Médio (recalcula `length` a cada iteração)
- `for...of`: Mais lento em arrays grandes, mas mais legível

**Conclusão:** Para arrays muito grandes (milhões de elementos), `for` com cache é mais rápido. Para a maioria dos casos, `for...of` é uma excelente escolha por ser mais legível.

---

### 2. `for...in` vs `for...of` vs `Object.keys()`

```javascript
const objeto = {};
for (let i = 0; i < 10000; i++) {
  objeto[`propriedade${i}`] = i;
}

// Teste 1: for...in
console.time("for...in");
for (const chave in objeto) {
  if (objeto.hasOwnProperty(chave)) {
    const valor = objeto[chave];
  }
}
console.timeEnd("for...in");

// Teste 2: Object.keys() + for...of
console.time("Object.keys + for...of");
for (const chave of Object.keys(objeto)) {
  const valor = objeto[chave];
}
console.timeEnd("Object.keys + for...of");

// Teste 3: Object.entries() + for...of
console.time("Object.entries + for...of");
for (const [chave, valor] of Object.entries(objeto)) {
  // já tem chave e valor
}
console.timeEnd("Object.entries + for...of");
```

**Resultados típicos:**
- `Object.keys() + for...of`: Geralmente mais rápido e mais seguro
- `for...in`: Mais lento, especialmente com `hasOwnProperty`
- `Object.entries() + for...of`: Mais lento, mas mais conveniente

**Conclusão:** Para objetos, `Object.keys()` ou `Object.entries()` com `for...of` são geralmente melhores que `for...in`.

---

## 🎯 Boas Práticas

### 1. Use `for...of` para Arrays (Regra de Ouro)

```javascript
// ✅ BOM - Legível, moderno, seguro
const frutas = ['maçã', 'banana', 'laranja'];
for (const fruta of frutas) {
  console.log(fruta);
}

// ⚠️ ACEITÁVEL - Funciona, mas mais verboso
for (let i = 0; i < frutas.length; i++) {
  console.log(frutas[i]);
}

// ❌ EVITE - Não use for...in com arrays
for (const indice in frutas) {
  console.log(frutas[indice]);
}
```

**Por quê?**
- Mais legível e moderno
- Menos propenso a erros
- Não precisa gerenciar índices manualmente
- Funciona com qualquer iterável

---

### 2. Cache de `length` em Loops `for` Tradicionais

```javascript
const array = [1, 2, 3, 4, 5];

// ❌ INEFICIENTE - Recalcula length a cada iteração
for (let i = 0; i < array.length; i++) {
  console.log(array[i]);
}

// ✅ EFICIENTE - Cache do length
for (let i = 0, len = array.length; i < len; i++) {
  console.log(array[i]);
}

// ✅ AINDA MELHOR - Use for...of
for (const elemento of array) {
  console.log(elemento);
}
```

**Por quê?**
- Evita recalcular `length` a cada iteração
- Em arrays grandes, isso pode fazer diferença
- Mas `for...of` já faz isso automaticamente

---

### 3. Evite Modificar Arrays Durante Iteração

```javascript
const numeros = [1, 2, 3, 4, 5];

// ❌ PROBLEMÁTICO - Pode pular elementos
for (let i = 0; i < numeros.length; i++) {
  if (numeros[i] % 2 === 0) {
    numeros.splice(i, 1); // Remove elemento
    // i não é incrementado, mas o array mudou!
  }
}

// ✅ CORRETO - Iterar de trás para frente
for (let i = numeros.length - 1; i >= 0; i--) {
  if (numeros[i] % 2 === 0) {
    numeros.splice(i, 1);
  }
}

// ✅ MELHOR - Criar novo array (imutável)
const numerosImpares = numeros.filter(num => num % 2 !== 0);
```

**Por quê?**
- Modificar durante iteração pode causar bugs sutis
- Iterar de trás para frente é mais seguro
- Criar novo array é mais funcional e seguro

---

### 4. Use `break` e `continue` com Moderação

```javascript
// ✅ BOM - break para otimizar busca
function encontrarPrimeiroPar(array) {
  for (const numero of array) {
    if (numero % 2 === 0) {
      return numero; // Mais claro que break
    }
  }
  return null;
}

// ⚠️ ACEITÁVEL - break quando necessário
let encontrado = false;
for (const item of itens) {
  if (item.id === alvo) {
    encontrado = true;
    break; // Para quando encontrar
  }
}

// ❌ EVITE - break desnecessário
for (let i = 0; i < 10; i++) {
  if (i === 5) {
    break; // Por que não usar i < 5 na condição?
  }
  console.log(i);
}
```

**Por quê?**
- `break` pode tornar o código menos legível
- Prefira retornar cedo ou ajustar a condição do loop
- Use `break` quando realmente necessário para performance

---

### 5. Escolha o Loop Correto para Cada Situação

```javascript
// ✅ for...of - Para arrays e iteráveis
const array = [1, 2, 3];
for (const elemento of array) {
  console.log(elemento);
}

// ✅ for...in - Apenas para propriedades de objetos
const objeto = { a: 1, b: 2 };
for (const chave in objeto) {
  if (objeto.hasOwnProperty(chave)) {
    console.log(chave, objeto[chave]);
  }
}

// ✅ while - Quando não sabe quantas iterações
let tentativas = 0;
while (!sucesso && tentativas < 10) {
  tentativas++;
  // tenta novamente
}

// ✅ do...while - Quando precisa executar pelo menos uma vez
let entrada;
do {
  entrada = prompt('Digite um número:');
} while (isNaN(entrada));
```

---

### 6. Nomenclatura Clara e Descritiva

```javascript
// ✅ BOM - Nomes descritivos
for (const produto of produtos) {
  calcularPreco(produto);
}

for (const usuario of usuarios) {
  validarUsuario(usuario);
}

// ❌ EVITE - Nomes genéricos
for (const item of items) {
  process(item);
}

for (const x of y) {
  doSomething(x);
}
```

**Por quê?**
- Código mais legível
- Mais fácil de entender e manter
- Menos erros

---

### 7. Evite Loops Aninhados Profundos

```javascript
// ⚠️ CUIDADO - Loops aninhados podem ser lentos
for (let i = 0; i < 1000; i++) {
  for (let j = 0; j < 1000; j++) {
    for (let k = 0; k < 1000; k++) {
      // 1 bilhão de iterações!
    }
  }
}

// ✅ MELHOR - Quebrar em funções menores
function processarLinha(linha) {
  for (const elemento of linha) {
    processarElemento(elemento);
  }
}

for (const linha of matriz) {
  processarLinha(linha);
}
```

**Por quê?**
- Loops aninhados têm complexidade O(n²) ou O(n³)
- Quebrar em funções melhora legibilidade
- Considere usar métodos de array (map, filter) quando possível

---

## 🚫 O que NÃO Fazer

### 1. ❌ NÃO Use `for...in` com Arrays

```javascript
// ❌ ERRADO
const array = [1, 2, 3];
for (const indice in array) {
  console.log(array[indice]);
}

// ✅ CORRETO
for (const elemento of array) {
  console.log(elemento);
}
```

**Por quê?**
- `for...in` itera sobre propriedades enumeráveis, não apenas elementos
- Pode incluir propriedades herdadas ou adicionadas
- Ordem não é garantida em todos os casos
- Performance pior que `for...of`

---

### 2. ❌ NÃO Crie Loops Infinitos Acidentalmente

```javascript
// ❌ ERRADO - Loop infinito
let i = 0;
while (i < 10) {
  console.log(i);
  // Esqueceu de incrementar!
}

// ✅ CORRETO
let i = 0;
while (i < 10) {
  console.log(i);
  i++; // Sempre incremente!
}
```

**Por quê?**
- Loops infinitos travam o navegador
- Difíceis de debugar
- Sempre tenha uma condição de saída clara

---

### 3. ❌ NÃO Use `var` em Loops `for`

```javascript
// ❌ PROBLEMÁTICO - var tem escopo de função
for (var i = 0; i < 3; i++) {
  setTimeout(() => {
    console.log(i); // Sempre imprime 3!
  }, 100);
}

// ✅ CORRETO - let tem escopo de bloco
for (let i = 0; i < 3; i++) {
  setTimeout(() => {
    console.log(i); // Imprime 0, 1, 2
  }, 100);
}
```

**Por quê?**
- `var` tem escopo de função, não de bloco
- Pode causar bugs com closures e callbacks assíncronos
- Sempre use `let` ou `const` em loops

---

### 4. ❌ NÃO Recalcule Valores Dentro do Loop

```javascript
const array = [1, 2, 3, 4, 5];

// ❌ INEFICIENTE - Recalcula a cada iteração
for (let i = 0; i < array.length; i++) {
  const resultado = array[i] * Math.sqrt(array.length);
  console.log(resultado);
}

// ✅ EFICIENTE - Calcula uma vez antes
const sqrtLength = Math.sqrt(array.length);
for (let i = 0; i < array.length; i++) {
  const resultado = array[i] * sqrtLength;
  console.log(resultado);
}
```

**Por quê?**
- Evita cálculos desnecessários
- Melhora performance significativamente
- Código mais claro

---

## 🔧 Otimização de Loops

### 1. Cache de Propriedades e Cálculos

```javascript
// ❌ INEFICIENTE
for (let i = 0; i < objetos.length; i++) {
  const resultado = objetos[i].propriedade.nested.calcular();
  processar(resultado);
}

// ✅ EFICIENTE
for (let i = 0; i < objetos.length; i++) {
  const obj = objetos[i];
  const nested = obj.propriedade.nested;
  const resultado = nested.calcular();
  processar(resultado);
}
```

---

### 2. Use `break` para Otimizar Buscas

```javascript
// ✅ BOM - Para quando encontrar
function encontrarItem(array, alvo) {
  for (const item of array) {
    if (item.id === alvo) {
      return item; // Para imediatamente
    }
  }
  return null;
}
```

---

### 3. Evite Operações Pesadas Dentro de Loops

```javascript
// ❌ INEFICIENTE - DOM query dentro do loop
for (let i = 0; i < 1000; i++) {
  const elemento = document.querySelector('.item');
  elemento.textContent = i;
}

// ✅ EFICIENTE - DOM query fora do loop
const elemento = document.querySelector('.item');
for (let i = 0; i < 1000; i++) {
  elemento.textContent = i;
}
```

---

## 🧹 Clean Code e Padrões

### 1. Prefira Métodos de Array Quando Apropriado

```javascript
// ⚠️ FUNCIONAL, mas loops são mais flexíveis
const numeros = [1, 2, 3, 4, 5];

// Loop tradicional
let soma = 0;
for (const numero of numeros) {
  soma += numero;
}

// Método de array (mais funcional)
const soma = numeros.reduce((acc, num) => acc + num, 0);

// Escolha baseado no contexto:
// - Loops: Mais controle, lógica complexa
// - Métodos: Mais funcional, código mais declarativo
```

---

### 2. Extraia Lógica Complexa para Funções

```javascript
// ❌ DIFÍCIL DE LER - Tudo em um loop
for (const usuario of usuarios) {
  if (usuario.idade >= 18 && usuario.ativo && !usuario.banido) {
    if (usuario.assinatura === 'premium' || usuario.assinatura === 'vip') {
      enviarEmail(usuario);
      registrarAcao(usuario, 'email_enviado');
      atualizarEstatisticas(usuario);
    }
  }
}

// ✅ MELHOR - Lógica extraída
function podeReceberEmail(usuario) {
  return usuario.idade >= 18 
    && usuario.ativo 
    && !usuario.banido
    && (usuario.assinatura === 'premium' || usuario.assinatura === 'vip');
}

function processarUsuario(usuario) {
  enviarEmail(usuario);
  registrarAcao(usuario, 'email_enviado');
  atualizarEstatisticas(usuario);
}

for (const usuario of usuarios) {
  if (podeReceberEmail(usuario)) {
    processarUsuario(usuario);
  }
}
```

---

### 3. Use Early Returns Quando Possível

```javascript
// ✅ BOM - Early return
function processarArray(array) {
  if (!array || array.length === 0) {
    return [];
  }
  
  const resultado = [];
  for (const item of array) {
    if (!item.valido) {
      continue; // Pula itens inválidos
    }
    resultado.push(processarItem(item));
  }
  
  return resultado;
}
```

---

## 🐛 Debugging de Loops

### 1. Use `console.log` Estrategicamente

```javascript
// ✅ BOM - Debug claro
for (let i = 0; i < array.length; i++) {
  console.log(`Iteração ${i}:`, array[i]);
  // processar...
}

// ✅ MELHOR - Use debugger
for (let i = 0; i < array.length; i++) {
  debugger; // Pausa aqui no DevTools
  // processar...
}
```

---

### 2. Adicione Validações

```javascript
// ✅ BOM - Validação antes do loop
function processarArray(array) {
  if (!Array.isArray(array)) {
    throw new Error('Array esperado');
  }
  
  if (array.length === 0) {
    return [];
  }
  
  for (const item of array) {
    // processar...
  }
}
```

---

## 📊 Resumo: Quando Usar Cada Loop

| Situação | Loop Recomendado | Por quê? |
|---------|------------------|----------|
| Array/Iterável | `for...of` | Legível, moderno, seguro |
| Objeto (propriedades) | `Object.keys()` + `for...of` | Mais seguro que `for...in` |
| Número conhecido de iterações | `for` | Controle total |
| Condição desconhecida | `while` | Flexível |
| Executar pelo menos uma vez | `do...while` | Garante execução |
| Performance crítica (arrays grandes) | `for` com cache | Mais rápido |

---

## 🎓 Regras de Ouro

1. ✅ **Use `for...of` para arrays** - É a escolha padrão moderna
2. ✅ **Use `Object.keys()` ou `Object.entries()` para objetos** - Mais seguro que `for...in`
3. ✅ **Cache `length` em loops `for` tradicionais** - Melhora performance
4. ✅ **Evite modificar arrays durante iteração** - Use filter ou itere de trás para frente
5. ✅ **Sempre use `let` ou `const` em loops** - Nunca `var`
6. ✅ **Extraia lógica complexa para funções** - Melhora legibilidade
7. ✅ **Use `break` e `continue` com moderação** - Prefira early returns
8. ✅ **Valide antes de iterar** - Evite erros em runtime

---

## 🚀 Próximos Passos

Agora que você entendeu performance e boas práticas de loops, você está pronto para:
- Aplicar loops em manipulação do DOM
- Trabalhar com métodos de array (map, filter, reduce)
- Entender programação assíncrona e loops
- Otimizar código em projetos reais

**Lembre-se**: Performance é importante, mas legibilidade também é. Encontre o equilíbrio certo para cada situação!

