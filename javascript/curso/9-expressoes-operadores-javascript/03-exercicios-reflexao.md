# Aula 9 - Exercícios e Reflexão: Expressões e Operadores

## 📝 Exercícios Práticos

### Exercício 1: Calculadora Básica

Crie um programa que realiza operações aritméticas básicas. Complete o código abaixo:

```javascript
// Declare as variáveis
let numero1 = 15;
let numero2 = 4;

// Calcule e exiba os resultados
let soma = /* complete aqui */;
let subtracao = /* complete aqui */;
let multiplicacao = /* complete aqui */;
let divisao = /* complete aqui */;
let resto = /* complete aqui */;
let potencia = /* complete aqui */;

console.log("Soma:", soma);
console.log("Subtração:", subtracao);
console.log("Multiplicação:", multiplicacao);
console.log("Divisão:", divisao);
console.log("Resto:", resto);
console.log("Potência:", potencia);
```

**Resultado esperado:**
```
Soma: 19
Subtração: 11
Multiplicação: 60
Divisão: 3.75
Resto: 3
Potência: 50625
```

---

### Exercício 2: Verificador de Idade e Status

Crie um programa que verifica se uma pessoa pode votar, dirigir e beber (considerando a idade mínima de 18 anos para votar e dirigir, e 21 para beber no Brasil).

```javascript
let idade = 20; // Altere este valor para testar

// Use operadores de comparação e lógicos
let podeVotar = /* complete aqui */;
let podeDirigir = /* complete aqui */;
let podeBeber = /* complete aqui */;

console.log("Idade:", idade);
console.log("Pode votar?", podeVotar);
console.log("Pode dirigir?", podeDirigir);
console.log("Pode beber?", podeBeber);
```

**Teste com diferentes idades:**
- 16 anos
- 18 anos
- 20 anos
- 21 anos

---

### Exercício 3: Sistema de Desconto

Crie um sistema que calcula desconto baseado no valor da compra:
- Compras acima de R$ 100: 10% de desconto
- Compras acima de R$ 50: 5% de desconto
- Compras abaixo de R$ 50: sem desconto

Use o operador ternário para determinar o desconto.

```javascript
let valorCompra = 120; // Altere este valor

// Use operador ternário para calcular o desconto
let desconto = /* complete aqui */;
let valorFinal = /* complete aqui */;

console.log("Valor da compra: R$", valorCompra);
console.log("Desconto: R$", desconto);
console.log("Valor final: R$", valorFinal);
```

---

### Exercício 4: Analisador de Números

Crie um programa que analisa um número e determina:
- Se é par ou ímpar
- Se é positivo, negativo ou zero
- Se é maior que 10

```javascript
let numero = 15; // Altere este valor

// Use operador módulo (%) para verificar se é par
let ePar = /* complete aqui */;
let tipoParidade = /* complete aqui usando operador ternário */;

// Verifique se é positivo, negativo ou zero
let tipoSinal = /* complete aqui usando operador ternário */;

// Verifique se é maior que 10
let maiorQue10 = /* complete aqui */;

console.log("Número:", numero);
console.log("É par?", ePar);
console.log("Tipo:", tipoParidade);
console.log("Sinal:", tipoSinal);
console.log("Maior que 10?", maiorQue10);
```

---

### Exercício 5: Manipulador de Strings

Complete o código para manipular strings usando operadores de concatenação:

```javascript
let primeiroNome = "João";
let sobrenome = "Silva";
let idade = 30;

// Concatene para criar uma mensagem completa
let nomeCompleto = /* complete aqui */;
let apresentacao = /* complete aqui */; // "Olá, meu nome é João Silva e tenho 30 anos"

console.log(nomeCompleto);
console.log(apresentacao);
```

---

### Exercício 6: Operadores de Atribuição Compostos

Complete o código usando operadores de atribuição compostos:

```javascript
let saldo = 1000;

// Use operadores de atribuição compostos
saldo /* complete aqui */ 500;  // Adicione 500
saldo /* complete aqui */ 200;  // Subtraia 200
saldo /* complete aqui */ 2;    // Multiplique por 2
saldo /* complete aqui */ 5;    // Divida por 5

console.log("Saldo final:", saldo); // Deve ser 520
```

---

### Exercício 7: Análise de Código

Analise o seguinte código e responda:

```javascript
let a = 5;
let b = 10;
let c = "5";

let resultado1 = a == c;
let resultado2 = a === c;
let resultado3 = a + b;
let resultado4 = a + c;
let resultado5 = a - c;
let resultado6 = !resultado1;
let resultado7 = a > b || b > a;
let resultado8 = a > b && b > a;
```

**Perguntas:**
1. Qual é o valor de `resultado1`? Por quê?
2. Qual é o valor de `resultado2`? Por quê?
3. Qual é o valor de `resultado3`? Por quê?
4. Qual é o valor de `resultado4`? Por quê?
5. Qual é o valor de `resultado5`? Por quê?
6. Qual é o valor de `resultado6`? Por quê?
7. Qual é o valor de `resultado7`? Por quê?
8. Qual é o valor de `resultado8`? Por quê?

---

### Exercício 8: Precedência de Operadores

Analise as seguintes expressões e determine o resultado. Depois, verifique no console:

```javascript
// Expressão 1
let resultado1 = 2 + 3 * 4;

// Expressão 2
let resultado2 = (2 + 3) * 4;

// Expressão 3
let resultado3 = 10 > 5 && 3 < 7;

// Expressão 4
let resultado4 = 10 > 5 || 3 > 7;

// Expressão 5
let resultado5 = !true || false;

// Expressão 6
let resultado6 = 5 === 5 && "texto" === "texto";

// Expressão 7
let resultado7 = 10 > 5 ? "maior" : "menor";

// Expressão 8
let resultado8 = (5, 10, 15);
```

**Tarefa:** 
1. Anote o que você acha que cada resultado será
2. Execute o código e verifique
3. Explique por que cada resultado é o que é

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Operadores de Comparação

**Cenário:** Você está desenvolvendo um sistema de autenticação onde precisa verificar se a senha digitada pelo usuário corresponde à senha armazenada.

```javascript
let senhaDigitada = "12345";
let senhaArmazenada = "12345";

// Opção 1
if (senhaDigitada == senhaArmazenada) {
    console.log("Acesso permitido");
}

// Opção 2
if (senhaDigitada === senhaArmazenada) {
    console.log("Acesso permitido");
}
```

**Perguntas:**
1. Qual operador você usaria (`==` ou `===`) e por quê?
2. Em que situação o uso de `==` poderia causar um bug de segurança?
3. Por que é importante entender a diferença entre comparação estrita e não estrita em sistemas reais?
4. Como você garantiria que a comparação seja sempre segura, mesmo se os tipos forem diferentes?

---

### Reflexão 2: Operadores Lógicos e Curto-Circuito

**Cenário:** Você está verificando se um usuário pode acessar um recurso:

```javascript
function verificarAcesso(usuario) {
    // Opção 1
    if (usuario && usuario.permissao && usuario.permissao.admin) {
        return true;
    }
    
    // Opção 2
    if (usuario?.permissao?.admin) {
        return true;
    }
    
    return false;
}
```

**Perguntas:**
1. Como o comportamento de curto-circuito do operador `&&` ajuda neste código?
2. O que aconteceria se `usuario` fosse `null` ou `undefined` na Opção 1?
3. Qual é a diferença entre usar `&&` encadeado e o optional chaining (`?.`)?
4. Em uma aplicação com milhões de usuários, qual abordagem seria mais eficiente? Por quê?
5. Como você garantiria que este código não quebre se a estrutura do objeto `usuario` mudar no futuro?

---

### Reflexão 3: Conversão de Tipos e Bugs Sutis

**Cenário:** Você está calculando o total de uma compra:

```javascript
let preco1 = "10.50";
let preco2 = "20.30";
let quantidade = 2;

// Código 1
let total1 = preco1 + preco2 * quantidade;

// Código 2
let total2 = (parseFloat(preco1) + parseFloat(preco2)) * quantidade;

// Código 3
let total3 = Number(preco1) + Number(preco2) * quantidade;
```

**Perguntas:**
1. Qual é o resultado de cada cálculo? Por quê?
2. Qual código está correto para calcular o total?
3. Como a conversão automática de tipos do JavaScript pode causar bugs difíceis de encontrar?
4. Em um sistema de e-commerce real, qual seria o impacto de um bug como este?
5. Como você garantiria que todos os valores numéricos sejam tratados corretamente, mesmo vindo de formulários HTML (que sempre retornam strings)?

---

### Reflexão 4: Operador Ternário vs If/Else

**Cenário:** Você precisa determinar o status de um pedido:

```javascript
// Abordagem 1: Operador Ternário
let status = pedido.valor > 100 ? "premium" : "standard";

// Abordagem 2: If/Else
let status;
if (pedido.valor > 100) {
    status = "premium";
} else {
    status = "standard";
}

// Abordagem 3: Ternário Aninhado
let status = pedido.valor > 100 ? "premium" : 
             pedido.valor > 50 ? "gold" : "standard";
```

**Perguntas:**
1. Quando você escolheria usar o operador ternário em vez de if/else?
2. Em que situação o ternário aninhado se torna difícil de ler e manter?
3. Como você equilibraria a brevidade do código com a legibilidade?
4. Em um código que será mantido por uma equipe, qual abordagem você recomendaria e por quê?
5. Como você documentaria a lógica de negócio por trás dessas decisões?

---

### Reflexão 5: Performance e Operadores

**Cenário:** Você está processando uma lista grande de usuários:

```javascript
// Código 1
for (let i = 0; i < usuarios.length; i++) {
    if (usuarios[i] && usuarios[i].ativo && usuarios[i].permissao && usuarios[i].permissao.admin) {
        // processar admin
    }
}

// Código 2
for (let i = 0; i < usuarios.length; i++) {
    if (usuarios[i]?.ativo && usuarios[i]?.permissao?.admin) {
        // processar admin
    }
}

// Código 3
const admins = usuarios.filter(u => u?.ativo && u?.permissao?.admin);
admins.forEach(admin => {
    // processar admin
});
```

**Perguntas:**
1. Como o curto-circuito dos operadores lógicos afeta a performance neste caso?
2. Qual código seria mais eficiente para uma lista com 1 milhão de usuários? Por quê?
3. Como você mediria a performance de cada abordagem?
4. Em que situação a otimização prematura pode ser prejudicial?
5. Como você equilibraria legibilidade, manutenibilidade e performance?

---

### Reflexão 6: Edge Cases e Valores Especiais

**Cenário:** Você está validando entrada de dados:

```javascript
function calcularMedia(notas) {
    let soma = 0;
    for (let i = 0; i < notas.length; i++) {
        soma += notas[i];
    }
    return soma / notas.length;
}

// Testes
calcularMedia([10, 8, 9]);
calcularMedia([]);
calcularMedia([null, undefined, 0]);
calcularMedia(["10", "8", "9"]);
```

**Perguntas:**
1. O que acontece em cada um desses casos? Por quê?
2. Como você trataria cada edge case (caso extremo)?
3. Por que é importante considerar valores como `null`, `undefined`, `NaN`, e arrays vazios?
4. Como você garantiria que a função sempre retorne um valor válido?
5. Em um sistema de produção, quais seriam as consequências de não tratar esses casos?

---

## 📋 Checklist de Aprendizado

Antes de prosseguir, certifique-se de que você:

- [ ] Consegue usar todos os operadores aritméticos corretamente
- [ ] Entende a diferença entre `==` e `===`
- [ ] Sabe quando usar `&&`, `||`, e `!`
- [ ] Compreende como funciona o operador ternário
- [ ] Entende a precedência de operadores básica
- [ ] Sabe usar operadores de atribuição compostos
- [ ] Consegue identificar valores falsy e truthy
- [ ] Entende o comportamento de curto-circuito
- [ ] Sabe quando usar `??` em vez de `||`
- [ ] Consegue escrever expressões complexas de forma legível

---

## 🎯 Desafio Final

Crie um **sistema de cálculo de frete** que:

1. Recebe o peso do produto (em kg)
2. Recebe a distância (em km)
3. Calcula o frete baseado nas regras:
   - Peso até 1kg: R$ 5,00 base
   - Peso de 1kg a 5kg: R$ 5,00 + R$ 2,00 por kg adicional
   - Peso acima de 5kg: R$ 15,00 + R$ 1,50 por kg adicional
   - Distância até 50km: sem taxa adicional
   - Distância de 50km a 200km: + R$ 0,50 por km adicional
   - Distância acima de 200km: + R$ 0,30 por km adicional
4. Aplica desconto de 10% se o valor total for acima de R$ 50,00
5. Exibe o valor final formatado

**Requisitos:**
- Use operadores de comparação para as condições
- Use operador ternário onde apropriado
- Use operadores aritméticos para os cálculos
- Valide se os valores de entrada são válidos (números positivos)
- Trate casos extremos (peso 0, distância 0, etc.)

**Exemplo de saída:**
```
Peso: 3kg
Distância: 100km
Frete base: R$ 9,00
Taxa de distância: R$ 25,00
Subtotal: R$ 34,00
Desconto: R$ 0,00
Total: R$ 34,00
```

---

## 📝 Respostas e Discussão

Após completar os exercícios e refletir sobre as perguntas, você estará pronto para a próxima etapa: **Performance, Boas Práticas e Otimização**!

Lembre-se: a prática constante é essencial para dominar expressões e operadores. Experimente criar seus próprios exemplos e testar diferentes cenários!



