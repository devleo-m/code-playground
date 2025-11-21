# Aula 4 - Exercícios e Reflexão: Conversão de Tipos

## 📝 Instruções

Complete os exercícios abaixo na ordem apresentada. Para cada exercício:
1. Leia cuidadosamente o enunciado
2. Escreva o código JavaScript necessário
3. Teste seu código no console do navegador
4. Verifique se o resultado está correto

**Importante:** Use conversões explícitas sempre que possível e prefira `===` em vez de `==`.

---

## 🎯 Exercício 1: Identificando Conversões Implícitas

**Objetivo:** Entender quando JavaScript faz conversões automáticas.

### Tarefa

Analise os seguintes códigos e **preveja o resultado** de cada um. Depois, execute no console para verificar:

```javascript
// Código 1
console.log("10" + 5);

// Código 2
console.log("10" - 5);

// Código 3
console.log("10" * 5);

// Código 4
console.log("10" / 5);

// Código 5
console.log("10" == 10);

// Código 6
console.log("10" === 10);

// Código 7
console.log(true + 1);

// Código 8
console.log(false + false);

// Código 9
console.log("" + 123);

// Código 10
console.log(null + 5);
```

**Sua Resposta:**
1. Anote o resultado que você espera para cada código
2. Execute no console e compare com suas previsões
3. Explique por que cada resultado aconteceu

---

## 🔢 Exercício 2: Conversão Explícita para Number

**Objetivo:** Praticar conversão de strings para números.

### Tarefa

Crie uma função chamada `converterParaNumero` que:
1. Recebe um valor (pode ser string, number, ou outro tipo)
2. Converte explicitamente para número usando `Number()`
3. Retorna o número ou `NaN` se não for possível converter
4. Testa a função com os seguintes valores:

```javascript
// Teste com estes valores:
"123"
"12.5"
"123abc"
""
" "
null
undefined
true
false
[5]
[1,2,3]
{}
```

**Código Base:**

```javascript
function converterParaNumero(valor) {
    // Seu código aqui
}

// Teste sua função
console.log(converterParaNumero("123"));      // Deve retornar: 123
console.log(converterParaNumero("12.5"));     // Deve retornar: 12.5
console.log(converterParaNumero("123abc"));    // Deve retornar: NaN
// ... continue testando com os outros valores
```

**Desafio Extra:**
Modifique a função para usar `parseInt()` quando o valor for uma string que representa um número inteiro, e `parseFloat()` quando for um número decimal.

---

## 📝 Exercício 3: Conversão Explícita para String

**Objetivo:** Praticar conversão de diferentes tipos para strings.

### Tarefa

Crie uma função chamada `converterParaString` que:
1. Recebe um valor de qualquer tipo
2. Converte explicitamente para string usando `String()`
3. Retorna a string resultante
4. Testa a função com os seguintes valores:

```javascript
// Teste com estes valores:
123
12.5
true
false
null
undefined
NaN
Infinity
[1, 2, 3]
[]
{a: 1, b: 2}
{}
```

**Código Base:**

```javascript
function converterParaString(valor) {
    // Seu código aqui
}

// Teste sua função
console.log(converterParaString(123));        // Deve retornar: "123"
console.log(converterParaString(true));       // Deve retornar: "true"
console.log(converterParaString(null));       // Deve retornar: "null"
// ... continue testando com os outros valores
```

**Desafio Extra:**
Para objetos, use `JSON.stringify()` em vez de `String()` para obter uma representação mais útil.

---

## ✅ Exercício 4: Conversão para Boolean e Valores Falsy/Truthy

**Objetivo:** Entender valores falsy e truthy.

### Tarefa

Crie uma função chamada `verificarTruthy` que:
1. Recebe um valor de qualquer tipo
2. Converte explicitamente para boolean usando `Boolean()`
3. Retorna `true` ou `false`
4. Testa a função e identifica quais valores são falsy:

```javascript
// Teste com estes valores e identifique os falsy:
""
"texto"
0
1
-1
null
undefined
NaN
false
true
[]
[1, 2, 3]
{}
{a: 1}
function() {}
```

**Código Base:**

```javascript
function verificarTruthy(valor) {
    // Seu código aqui
}

// Teste e liste quais são falsy
console.log(verificarTruthy(""));           // false (falsy)
console.log(verificarTruthy("texto"));      // true (truthy)
// ... continue testando

// Liste todos os valores falsy que você encontrou:
// 1. 
// 2. 
// 3. 
// ...
```

**Desafio Extra:**
Crie uma função que recebe um array de valores e retorna apenas os valores truthy, removendo todos os falsy.

---

## 🔄 Exercício 5: Comparação de Métodos de Conversão

**Objetivo:** Entender diferenças entre métodos de conversão.

### Tarefa

Compare os diferentes métodos de conversão para número e identifique quando usar cada um:

```javascript
let valor1 = "123";
let valor2 = "12.5";
let valor3 = "123abc";
let valor4 = "08";
let valor5 = "";

// Teste cada método com cada valor
console.log("Number():");
console.log(Number(valor1));
console.log(Number(valor2));
console.log(Number(valor3));
console.log(Number(valor4));
console.log(Number(valor5));

console.log("\nparseInt():");
console.log(parseInt(valor1));
console.log(parseInt(valor2));
console.log(parseInt(valor3));
console.log(parseInt(valor4));
console.log(parseInt(valor5));

console.log("\nparseFloat():");
console.log(parseFloat(valor1));
console.log(parseFloat(valor2));
console.log(parseFloat(valor3));
console.log(parseFloat(valor4));
console.log(parseFloat(valor5));

console.log("\nOperador +:");
console.log(+valor1);
console.log(+valor2);
console.log(+valor3);
console.log(+valor4);
console.log(+valor5);
```

**Perguntas:**
1. Qual método preserva decimais?
2. Qual método funciona melhor com strings que contêm letras?
3. Qual método você usaria para converter "08" para o número 8?
4. Qual método é mais seguro para validação?

---

## 🛡️ Exercício 6: Função de Validação e Conversão Segura

**Objetivo:** Criar uma função robusta que valida antes de converter.

### Tarefa

Crie uma função chamada `converterNumeroSeguro` que:
1. Recebe um valor
2. Valida se o valor pode ser convertido para número
3. Retorna o número se válido, ou `null` se inválido
4. Trata casos especiais (strings vazias, null, undefined)

**Requisitos:**
- Deve retornar `null` para valores inválidos (não `NaN`)
- Deve tratar strings vazias como inválidas
- Deve tratar `null` e `undefined` como inválidos
- Deve aceitar números já convertidos
- Deve usar conversão explícita

**Código Base:**

```javascript
function converterNumeroSeguro(valor) {
    // Seu código aqui
    // Dica: Use typeof, trim() para strings, e Number()
}

// Testes
console.log(converterNumeroSeguro("123"));      // 123
console.log(converterNumeroSeguro("12.5"));     // 12.5
console.log(converterNumeroSeguro(""));         // null
console.log(converterNumeroSeguro("   "));      // null
console.log(converterNumeroSeguro("abc"));      // null
console.log(converterNumeroSeguro(null));       // null
console.log(converterNumeroSeguro(undefined));   // null
console.log(converterNumeroSeguro(123));        // 123 (já é número)
```

---

## 🎯 Exercício 7: Sistema de Cálculo de Preços

**Objetivo:** Aplicar conversões em um contexto prático.

### Tarefa

Crie um sistema simples de cálculo de preços que:
1. Recebe preço e quantidade (podem vir como strings do usuário)
2. Calcula o total (preço × quantidade)
3. Aplica desconto se fornecido (em porcentagem)
4. Retorna o valor final formatado como string com 2 casas decimais

**Requisitos:**
- Use conversões explícitas
- Valide se os valores são números válidos
- Trate erros (retorne mensagem de erro se inválido)
- Formate o resultado final como "R$ XX.XX"

**Código Base:**

```javascript
function calcularPreco(preco, quantidade, desconto = 0) {
    // Seu código aqui
    // 1. Converter preco e quantidade para números
    // 2. Validar se são números válidos
    // 3. Calcular total = preco * quantidade
    // 4. Aplicar desconto se fornecido
    // 5. Formatar resultado como "R$ XX.XX"
    // 6. Retornar erro se valores inválidos
}

// Testes
console.log(calcularPreco("10.50", "2", 10));     // "R$ 18.90"
console.log(calcularPreco("25", 3, 0));            // "R$ 75.00"
console.log(calcularPreco("abc", "2"));            // "Erro: valores inválidos"
console.log(calcularPreco("10", "", 5));           // "Erro: valores inválidos"
```

**Desafio Extra:**
Adicione validação para garantir que preço e quantidade sejam positivos.

---

## 🧩 Exercício 8: Analisando Código com Conversões

**Objetivo:** Identificar problemas em código existente.

### Tarefa

Analise o seguinte código e identifique **todos os problemas** relacionados a conversão de tipos:

```javascript
function processarDados(entrada) {
    let resultado = "";
    
    // Processamento 1
    if (entrada == 0) {
        resultado = "zero";
    }
    
    // Processamento 2
    let numero = entrada + 0;
    
    // Processamento 3
    if (entrada) {
        resultado += numero;
    }
    
    // Processamento 4
    let total = resultado + 10;
    
    return total;
}

// Teste com diferentes valores
console.log(processarDados("5"));
console.log(processarDados(0));
console.log(processarDados(""));
console.log(processarDados(null));
```

**Sua Análise:**
1. Liste todos os problemas que você encontrou
2. Explique por que cada problema acontece
3. Reescreva a função corrigindo todos os problemas
4. Use conversões explícitas e comparações estritas

---

## 🤔 Perguntas de Reflexão

Responda as seguintes perguntas com base no que você aprendeu:

### 1. Por que a conversão implícita pode ser perigosa?

**Sua Resposta:**

---

### 2. Qual é a diferença prática entre `Number("123abc")` e `parseInt("123abc")`? Quando você usaria cada um?

**Sua Resposta:**

---

### 3. Por que arrays e objetos vazios são considerados "truthy" mesmo estando vazios? Isso faz sentido?

**Sua Resposta:**

---

### 4. Em uma aplicação web real, os dados do usuário geralmente chegam como strings (de formulários HTML). Como você garantiria que esses dados sejam convertidos corretamente antes de usar em cálculos?

**Sua Resposta:**

---

### 5. Qual seria o impacto de usar `==` em vez de `===` em uma aplicação com milhares de usuários? Pense em performance e bugs potenciais.

**Sua Resposta:**

---

### 6. O método `parseInt("08")` pode retornar resultados diferentes em navegadores antigos. Por que isso acontece e como você evitaria esse problema?

**Sua Resposta:**

---

### 7. Quando você recebe dados de uma API, eles geralmente vêm como JSON (strings). Como você garantiria que os tipos numéricos sejam preservados corretamente?

**Sua Resposta:**

---

### 8. Em termos de performance, há diferença significativa entre `Number()`, `parseInt()`, `parseFloat()` e o operador `+`? Quando cada um seria mais apropriado?

**Sua Resposta:**

---

## 📋 Checklist de Conclusão

Antes de enviar suas respostas, verifique:

- [ ] Completei todos os 8 exercícios
- [ ] Testei cada código no console do navegador
- [ ] Respondi todas as 8 perguntas de reflexão
- [ ] Usei conversões explícitas nos meus códigos
- [ ] Usei `===` em vez de `==` nas comparações
- [ ] Revisei meu código procurando por possíveis erros

---

## 🚀 Próximos Passos

Após completar os exercícios e reflexões:

1. Revise suas respostas
2. Teste todos os códigos no console
3. Compare seus resultados com os esperados
4. Envie suas respostas para análise e feedback

**Boa sorte! 🎯**

