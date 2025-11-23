# Aula 6 - Exercícios e Reflexão: Comparações de Igualdade

## 📝 Exercícios Práticos

### Exercício 1: Identificando o Resultado das Comparações

Analise cada comparação abaixo e determine qual será o resultado (`true` ou `false`). Depois, teste no console do navegador para verificar suas respostas.

```javascript
// 1.1. Comparações com ==
console.log("5" == 5);
console.log(true == 1);
console.log(false == 0);
console.log(null == undefined);
console.log("" == 0);
console.log([] == "");

// 1.2. Comparações com ===
console.log("5" === 5);
console.log(true === 1);
console.log(false === 0);
console.log(null === undefined);
console.log("" === 0);
console.log([] === "");

// 1.3. Comparações com Object.is()
console.log(Object.is("5", 5));
console.log(Object.is(NaN, NaN));
console.log(Object.is(-0, +0));
console.log(Object.is(-0, 0));
```

**Sua tarefa:** 
1. Anote suas respostas antes de testar
2. Execute cada linha no console
3. Compare seus resultados com o que você pensou
4. Explique por que cada resultado aconteceu

---

### Exercício 2: Função de Validação de Idade

Crie uma função chamada `validarIdade` que:
- Recebe um parâmetro `idade`
- Retorna `true` se a idade for exatamente 18 (número)
- Retorna `false` caso contrário
- Deve usar `===` para a comparação

**Exemplo de uso:**
```javascript
console.log(validarIdade(18));      // true
console.log(validarIdade("18"));    // false (é string, não número)
console.log(validarIdade(17));      // false
console.log(validarIdade(19));      // false
```

**Sua tarefa:** Escreva a função e teste com diferentes valores.

---

### Exercício 3: Verificação de NaN

Crie uma função chamada `ehNaN` que:
- Recebe um parâmetro `valor`
- Retorna `true` se o valor for `NaN`
- Retorna `false` caso contrário
- Deve usar `Object.is()` para a verificação

**Dica:** Lembre-se que `NaN === NaN` retorna `false`, então você precisa usar `Object.is()`.

**Exemplo de uso:**
```javascript
console.log(ehNaN(NaN));                    // true
console.log(ehNaN(Number("abc")));         // true
console.log(ehNaN(5));                     // false
console.log(ehNaN("texto"));               // false
```

**Sua tarefa:** Escreva a função e teste com diferentes valores.

---

### Exercício 4: Sistema de Comparação de Senhas

Crie uma função chamada `verificarSenha` que:
- Recebe dois parâmetros: `senhaDigitada` e `senhaCorreta`
- Retorna um objeto com:
  - `valida`: `true` se as senhas forem iguais (usando `===`)
  - `tipoCorreto`: `true` se os tipos forem iguais
  - `mensagem`: uma mensagem explicando o resultado

**Exemplo de uso:**
```javascript
console.log(verificarSenha(1234, 1234));
// { valida: true, tipoCorreto: true, mensagem: "Senha correta!" }

console.log(verificarSenha("1234", 1234));
// { valida: false, tipoCorreto: false, mensagem: "Senha incorreta: tipos diferentes" }

console.log(verificarSenha(1234, 5678));
// { valida: false, tipoCorreto: true, mensagem: "Senha incorreta: valores diferentes" }
```

**Sua tarefa:** Escreva a função completa com todas as verificações.

---

### Exercício 5: Comparação de Objetos

Analise o código abaixo e responda:

```javascript
const pessoa1 = { nome: "João", idade: 25 };
const pessoa2 = { nome: "João", idade: 25 };
const pessoa3 = pessoa1;

// Comparações
console.log(pessoa1 === pessoa2);  // ?
console.log(pessoa1 === pessoa3);  // ?
console.log(pessoa2 === pessoa3);  // ?

// Modificações
pessoa1.idade = 30;
console.log(pessoa3.idade);       // ?
console.log(pessoa2.idade);       // ?
```

**Sua tarefa:**
1. Determine o resultado de cada `console.log`
2. Explique por que `pessoa1 === pessoa2` retorna `false` mesmo tendo os mesmos valores
3. Explique o que acontece quando você modifica `pessoa1.idade` e por que `pessoa3.idade` também muda

---

### Exercício 6: Função de Comparação Segura

Crie uma função chamada `compararValores` que:
- Recebe três parâmetros: `valor1`, `valor2`, e `modo` (que pode ser `"relaxado"`, `"rigido"`, ou `"preciso"`)
- Retorna o resultado da comparação usando:
  - `==` se modo for `"relaxado"`
  - `===` se modo for `"rigido"`
  - `Object.is()` se modo for `"preciso"`

**Exemplo de uso:**
```javascript
console.log(compararValores("5", 5, "relaxado"));  // true
console.log(compararValores("5", 5, "rigido"));    // false
console.log(compararValores("5", 5, "preciso"));  // false

console.log(compararValores(NaN, NaN, "rigido"));  // false
console.log(compararValores(NaN, NaN, "preciso")); // true
```

**Sua tarefa:** Escreva a função usando uma estrutura condicional (if/else ou switch).

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por que === é mais seguro que ==?

**Pergunta:** Em um sistema de e-commerce, você precisa verificar se o preço de um produto é R$ 0,00 (produto grátis). 

Considere este código:
```javascript
const preco = "";  // String vazia vinda de um formulário

// Opção 1: Usando ==
if (preco == 0) {
    console.log("Produto grátis!");
}

// Opção 2: Usando ===
if (preco === 0) {
    console.log("Produto grátis!");
}
```

**Reflita sobre:**
1. Qual das duas opções é mais segura? Por quê?
2. O que aconteceria se um usuário acidentalmente enviasse uma string vazia no formulário?
3. Qual seria o impacto se o sistema aceitasse string vazia como preço zero?
4. Como você poderia melhorar essa verificação para ser ainda mais robusta?

---

### Reflexão 2: Performance e Eficiência

**Pergunta:** Considere este código que verifica se um array contém um valor específico:

```javascript
// Versão 1: Usando ==
function contemValor(arr, valor) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] == valor) {
            return true;
        }
    }
    return false;
}

// Versão 2: Usando ===
function contemValor(arr, valor) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === valor) {
            return true;
        }
    }
    return false;
}
```

**Reflita sobre:**
1. Qual versão seria mais rápida? Por quê?
2. Em uma aplicação com milhões de comparações por segundo, qual seria o impacto da diferença?
3. Além da performance, quais outros fatores você deve considerar ao escolher entre == e ===?
4. Como você poderia medir a diferença de performance entre as duas versões?

---

### Reflexão 3: Edge Cases e Comportamentos Inesperados

**Pergunta:** Analise este código de um sistema de autenticação:

```javascript
function autenticar(usuario, senha) {
    const usuarioCorreto = "admin";
    const senhaCorreta = 12345;
    
    // Usando ==
    if (usuario == usuarioCorreto && senha == senhaCorreta) {
        return "Acesso permitido";
    }
    return "Acesso negado";
}

// Testes
console.log(autenticar("admin", 12345));        // ?
console.log(autenticar("admin", "12345"));      // ?
console.log(autenticar("admin", 012345));       // ? (número em octal)
console.log(autenticar("admin", true));         // ? (true == 1, mas 1 != 12345)
```

**Reflita sobre:**
1. Quais são os possíveis problemas de segurança neste código?
2. Como um atacante poderia explorar o uso de `==` para burlar a autenticação?
3. Quais edge cases (casos extremos) você consegue identificar?
4. Como você reescreveria esta função para ser mais segura?
5. Que tipo de validação adicional você adicionaria?

---

### Reflexão 4: Quando Object.is() é Necessário?

**Pergunta:** Em um sistema de cálculo científico, você precisa verificar se o resultado de uma operação é `NaN`:

```javascript
function calcularRaizQuadrada(numero) {
    const resultado = Math.sqrt(numero);
    
    // Opção 1: Usando ===
    if (resultado === NaN) {
        return "Erro: não é possível calcular";
    }
    
    // Opção 2: Usando Object.is()
    if (Object.is(resultado, NaN)) {
        return "Erro: não é possível calcular";
    }
    
    // Opção 3: Usando isNaN()
    if (isNaN(resultado)) {
        return "Erro: não é possível calcular";
    }
    
    return resultado;
}
```

**Reflita sobre:**
1. Qual das três opções funcionaria corretamente? Por quê?
2. Qual seria a melhor opção? Justifique sua escolha.
3. Em que situações `Object.is()` seria essencial e não poderia ser substituído?
4. Quando você deveria usar `Object.is()` em vez de `===` no dia a dia?
5. Qual seria o impacto se você usasse a opção errada em um sistema crítico?

---

### Reflexão 5: Manutenibilidade e Legibilidade

**Pergunta:** Você está revisando código de um colega e encontra:

```javascript
// Código misturado
function processarDados(dados) {
    if (dados == null) {
        return [];
    }
    
    if (dados.length === 0) {
        return [];
    }
    
    if (dados.tipo == "array") {
        return dados.valores;
    }
    
    if (typeof dados === "object") {
        return Object.values(dados);
    }
    
    return [dados];
}
```

**Reflita sobre:**
1. Quais problemas você identifica neste código?
2. Como a mistura de `==` e `===` afeta a legibilidade?
3. Como isso poderia confundir outros desenvolvedores que precisam manter este código?
4. Qual seria o impacto em um projeto grande com múltiplos desenvolvedores?
5. Como você padronizaria este código seguindo boas práticas?
6. Que regras você estabeleceria em um guia de estilo para sua equipe?

---

## 📋 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Explicar a diferença entre `==`, `===` e `Object.is()`
- [ ] Identificar quando cada operador retorna `true` ou `false`
- [ ] Escrever código usando `===` corretamente
- [ ] Usar `Object.is()` para verificar `NaN`
- [ ] Entender por que objetos são comparados por referência
- [ ] Identificar problemas de segurança relacionados ao uso de `==`
- [ ] Explicar edge cases e comportamentos inesperados
- [ ] Escolher o operador correto para diferentes situações

---

## 🎯 Próximos Passos

Após completar os exercícios e refletir sobre as perguntas:

1. **Teste seus códigos** no console do navegador
2. **Experimente variações** dos exercícios
3. **Pense em situações reais** onde você usaria cada operador
4. **Discuta suas respostas** com outras pessoas (se possível)
5. **Revise os conceitos** que você teve dificuldade

**Lembre-se:** A prática constante é essencial para dominar esses conceitos! 🚀

---

**Quando terminar os exercícios, envie suas respostas para análise e feedback!**


