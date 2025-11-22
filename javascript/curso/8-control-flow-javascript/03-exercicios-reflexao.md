# Aula 8 - Exercícios e Reflexão: Control Flow

## 📝 Exercícios Práticos

### Exercício 1: Sistema de Notas Escolares

Crie uma função chamada `calcularConceito` que:
- Recebe uma nota (número de 0 a 100)
- Retorna o conceito baseado na nota:
  - 90-100: "A" (Excelente)
  - 80-89: "B" (Muito Bom)
  - 70-79: "C" (Bom)
  - 60-69: "D" (Aprovado)
  - 0-59: "F" (Reprovado)
- Se a nota for inválida (menor que 0 ou maior que 100), retorne "Nota inválida"

**Sua tarefa:** 
1. Escreva a função usando `if...else if...else`
2. Reescreva usando `switch` (dica: use `Math.floor(nota / 10)`)
3. Compare as duas versões - qual você acha mais legível para este caso?

**Exemplo de uso:**
```javascript
console.log(calcularConceito(95));  // "A"
console.log(calcularConceito(85));  // "B"
console.log(calcularConceito(75));  // "C"
console.log(calcularConceito(65));  // "D"
console.log(calcularConceito(45));  // "F"
console.log(calcularConceito(150)); // "Nota inválida"
console.log(calcularConceito(-10)); // "Nota inválida"
```

---

### Exercício 2: Validação de Formulário com Tratamento de Erros

Crie uma função chamada `validarFormulario` que:
- Recebe um objeto com `nome`, `email` e `idade`
- Valida cada campo:
  - `nome`: não pode estar vazio
  - `email`: deve conter "@" e ter pelo menos 5 caracteres
  - `idade`: deve ser um número entre 18 e 120
- Use `try...catch` para tratar erros
- Lance erros personalizados para cada validação
- Retorne `true` se tudo estiver válido, `false` caso contrário

**Sua tarefa:** 
1. Escreva a função completa com todas as validações
2. Use `throw new Error()` para cada erro de validação
3. Capture e exiba mensagens de erro específicas

**Exemplo de uso:**
```javascript
// Caso válido
validarFormulario({
  nome: "João Silva",
  email: "joao@email.com",
  idade: 25
}); // true

// Casos inválidos
validarFormulario({
  nome: "",
  email: "joao@email.com",
  idade: 25
}); // false - "Nome não pode estar vazio"

validarFormulario({
  nome: "João",
  email: "email-invalido",
  idade: 25
}); // false - "Email inválido"

validarFormulario({
  nome: "João",
  email: "joao@email.com",
  idade: 15
}); // false - "Idade deve ser entre 18 e 120"
```

---

### Exercício 3: Calculadora com Switch

Crie uma função chamada `calculadora` que:
- Recebe três parâmetros: `numero1`, `operacao`, `numero2`
- Usa `switch` para realizar a operação:
  - `"+"`: soma
  - `"-"`: subtração
  - `"*"`: multiplicação
  - `"/"`: divisão
  - `"%"`: resto da divisão
- Trata divisão por zero
- Retorna o resultado ou uma mensagem de erro

**Sua tarefa:**
1. Escreva a função usando `switch`
2. Adicione tratamento de erro para divisão por zero
3. Adicione um `default` para operações inválidas

**Exemplo de uso:**
```javascript
console.log(calculadora(10, "+", 5));  // 15
console.log(calculadora(10, "-", 5));  // 5
console.log(calculadora(10, "*", 5));  // 50
console.log(calculadora(10, "/", 5));  // 2
console.log(calculadora(10, "/", 0));  // "Erro: Divisão por zero!"
console.log(calculadora(10, "%", 3));  // 1
console.log(calculadora(10, "x", 5));  // "Operação inválida!"
```

---

### Exercício 4: Sistema de Desconto com Operador Ternário

Crie uma função chamada `calcularPrecoFinal` que:
- Recebe `preco` e `tipoCliente` como parâmetros
- Calcula desconto baseado no tipo de cliente:
  - `"vip"`: 20% de desconto
  - `"premium"`: 15% de desconto
  - `"regular"`: 5% de desconto
  - outros: sem desconto
- Use operador ternário para determinar o desconto
- Retorna o preço final

**Sua tarefa:**
1. Escreva a função usando operador ternário aninhado
2. Teste com diferentes tipos de cliente
3. Reflita: seria melhor usar `if...else` ou `switch` aqui? Por quê?

**Exemplo de uso:**
```javascript
console.log(calcularPrecoFinal(100, "vip"));      // 80
console.log(calcularPrecoFinal(100, "premium"));   // 85
console.log(calcularPrecoFinal(100, "regular"));   // 95
console.log(calcularPrecoFinal(100, "novo"));      // 100
```

---

### Exercício 5: Tratamento de Erros com Tipos Específicos

Crie uma função chamada `processarDados` que:
- Recebe um array de números
- Tenta calcular a média dos números
- Trata diferentes tipos de erros:
  - Se o array estiver vazio, lance `RangeError`
  - Se algum elemento não for número, lance `TypeError`
  - Se o array for `null` ou `undefined`, lance `ReferenceError`
- Use `try...catch` com verificação de tipo de erro (`instanceof`)
- Retorne a média ou uma mensagem de erro apropriada

**Sua tarefa:**
1. Escreva a função com tratamento de todos os tipos de erro
2. Use `instanceof` para identificar o tipo de erro
3. Retorne mensagens específicas para cada tipo de erro

**Exemplo de uso:**
```javascript
console.log(processarDados([10, 20, 30]));        // 20
console.log(processarDados([10, "20", 30]));     // "Erro de tipo: Elemento não é número"
console.log(processarDados([]));                  // "Erro de range: Array vazio"
console.log(processarDados(null));                // "Erro de referência: Array não fornecido"
```

---

### Exercício 6: Sistema de Autenticação Completo

Crie uma função chamada `autenticarUsuario` que:
- Recebe `usuario` e `senha`
- Valida:
  1. Se `usuario` está vazio → lance erro
  2. Se `senha` está vazia → lance erro
  3. Se `senha` tem menos de 6 caracteres → lance erro
  4. Se `usuario` não existe no sistema → lance erro
  5. Se `senha` está incorreta → lance erro
- Use `try...catch...finally`
- No `finally`, registre uma tentativa de login (mesmo que tenha falhado)
- Retorne `true` se autenticação for bem-sucedida

**Sua tarefa:**
1. Crie um objeto simples de "banco de dados" de usuários
2. Implemente todas as validações
3. Use `finally` para registrar tentativas de login

**Exemplo de uso:**
```javascript
// Banco de dados simulado
const usuarios = {
  "joao": "senha123",
  "maria": "senha456"
};

autenticarUsuario("joao", "senha123");  // true
autenticarUsuario("joao", "senhaErrada"); // false - "Senha incorreta"
autenticarUsuario("", "senha123");       // false - "Usuário não pode estar vazio"
autenticarUsuario("joao", "123");       // false - "Senha deve ter pelo menos 6 caracteres"
```

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Performance e Legibilidade

Analise o seguinte código:

```javascript
// Versão 1: if...else
function verificarStatus(nota) {
  if (nota >= 90) return "A";
  else if (nota >= 80) return "B";
  else if (nota >= 70) return "C";
  else if (nota >= 60) return "D";
  else return "F";
}

// Versão 2: switch
function verificarStatus(nota) {
  switch (Math.floor(nota / 10)) {
    case 10:
    case 9: return "A";
    case 8: return "B";
    case 7: return "C";
    case 6: return "D";
    default: return "F";
  }
}

// Versão 3: Operador ternário
function verificarStatus(nota) {
  return nota >= 90 ? "A" :
         nota >= 80 ? "B" :
         nota >= 70 ? "C" :
         nota >= 60 ? "D" : "F";
}
```

**Perguntas:**
1. Qual versão você considera mais legível? Por quê?
2. Qual versão tem melhor performance? (Dica: `switch` pode ser mais rápido em alguns casos)
3. Em uma aplicação com milhões de verificações por segundo, qual impacto cada abordagem teria?
4. Como você escolheria qual versão usar em um projeto real? Quais fatores consideraria?

---

### Reflexão 2: Tratamento de Erros e Robustez

Analise este código de uma função que processa pagamentos:

```javascript
function processarPagamento(valor, metodo) {
  if (valor <= 0) {
    console.log("Valor inválido");
    return false;
  }
  
  if (metodo === "cartao") {
    // Processa pagamento com cartão
    return true;
  } else if (metodo === "pix") {
    // Processa pagamento com PIX
    return true;
  } else {
    console.log("Método inválido");
    return false;
  }
}
```

**Perguntas:**
1. Quais problemas você identifica neste código em relação ao tratamento de erros?
2. O que aconteceria se `valor` fosse `null` ou `undefined`? O código está preparado para isso?
3. Como você melhoraria este código usando `try...catch` e tipos de erro apropriados?
4. Em uma aplicação de e-commerce real, quais seriam as consequências se um erro não tratado ocorresse durante um pagamento?
5. Como você garantiria que, mesmo com erro, o sistema registrasse a tentativa de pagamento?

---

### Reflexão 3: Edge Cases e Validação

Considere esta função que valida idade:

```javascript
function validarIdade(idade) {
  if (idade >= 18) {
    return true;
  } else {
    return false;
  }
}
```

**Perguntas:**
1. Quais edge cases (casos extremos) esta função não trata?
2. O que acontece se `idade` for `null`, `undefined`, uma string, ou um número negativo?
3. Como você reescreveria esta função para ser mais robusta?
4. Em um sistema de cadastro de usuários, quais seriam os riscos de segurança de não validar adequadamente a idade?
5. Como você implementaria validação que também considere performance? (Ex: validar tipos antes de fazer comparações)

---

### Reflexão 4: Switch vs If...Else em Aplicações Escaláveis

Imagine que você está construindo um sistema de processamento de pedidos com 20 tipos diferentes de produtos, cada um com regras de processamento específicas.

**Perguntas:**
1. Você usaria `switch` ou `if...else`? Por quê?
2. Como você organizaria o código para facilitar a adição de novos tipos de produtos no futuro?
3. Qual abordagem seria mais fácil de manter e testar?
4. Em termos de performance, qual seria mais eficiente para 20+ casos?
5. Como você estruturaria o código para que outros desenvolvedores pudessem adicionar novos tipos sem modificar código existente? (Dica: pense em padrões de design)

---

## 📋 Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Escrever estruturas condicionais (`if`, `if...else`, `if...else if...else`)
- [ ] Usar operador ternário para decisões simples
- [ ] Implementar `switch` com `break` e `default`
- [ ] Escolher entre `if...else` e `switch` adequadamente
- [ ] Usar `try...catch` para tratar erros
- [ ] Implementar `finally` para código que sempre deve executar
- [ ] Lançar erros com `throw`
- [ ] Identificar e tratar diferentes tipos de erros (`ReferenceError`, `TypeError`, `RangeError`)
- [ ] Criar erros personalizados
- [ ] Validar dados antes de processar
- [ ] Pensar em edge cases ao escrever código

---

## 🚀 Próximo Passo

Após completar os exercícios e refletir sobre as perguntas, você estará pronto para aprender sobre **Performance, Boas Práticas e Otimização**!

**Arquivo seguinte**: `04-performance-boas-praticas.md`

