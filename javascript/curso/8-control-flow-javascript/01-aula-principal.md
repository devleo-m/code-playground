# Aula 8: Control Flow em JavaScript - Conteúdo Principal

## 📖 Revisão da Aula Anterior

Na **Aula 7**, você aprendeu:
- ✅ Loops `for`, `while`, `do...while`
- ✅ Loops `for...of` e `for...in`
- ✅ Como usar `break` e `continue`
- ✅ Loops aninhados e suas aplicações

Agora vamos aprender sobre **Control Flow** (Fluxo de Controle) - o conceito fundamental que determina **como e em que ordem** o JavaScript executa seu código!

---

## 🎯 O que é Control Flow?

**Definição:** Control Flow (Fluxo de Controle) é a ordem em que o computador executa instruções em um programa. Por padrão, o JavaScript executa código de cima para baixo, linha por linha, mas podemos alterar esse fluxo usando estruturas especiais.

### Conceitos Fundamentais

1. **Sequential Flow** (Fluxo Sequencial): Execução padrão, linha por linha
2. **Conditional Flow** (Fluxo Condicional): Execução baseada em condições
3. **Exception Handling** (Tratamento de Exceções): Controle de erros e exceções
4. **Iterative Flow** (Fluxo Iterativo): Repetição de código (já visto em loops)

### Por que Control Flow é Importante?

Sem controle de fluxo, o código sempre executaria na mesma ordem, sem poder:
- Tomar decisões baseadas em dados
- Tratar erros adequadamente
- Repetir ações quando necessário
- Responder a diferentes situações

---

## 📊 1. Sequential Flow (Fluxo Sequencial)

### Definição

O **Sequential Flow** é o modo padrão de execução do JavaScript. O código é executado de cima para baixo, linha por linha, na ordem em que aparece.

### Como Funciona

```javascript
// Exemplo de execução sequencial
console.log("Passo 1: Iniciando programa");
let nome = "JavaScript";
console.log("Passo 2: Variável criada:", nome);
let idade = 28;
console.log("Passo 3: Segunda variável criada:", idade);
let resultado = nome + " tem " + idade + " anos";
console.log("Passo 4: Resultado final:", resultado);
```

**Saída:**
```
Passo 1: Iniciando programa
Passo 2: Variável criada: JavaScript
Passo 3: Segunda variável criada: 28
Passo 4: Resultado final: JavaScript tem 28 anos
```

### Características

- **Ordem garantida**: Cada linha executa após a anterior
- **Dependências**: Linhas posteriores podem usar valores de linhas anteriores
- **Sem pulos**: Não há saltos ou desvios (a menos que usemos outras estruturas)

### Exemplo Prático

```javascript
// Cálculo sequencial de uma compra
let precoProduto = 50.00;
console.log("Preço do produto:", precoProduto);

let desconto = 10.00;
console.log("Desconto:", desconto);

let precoFinal = precoProduto - desconto;
console.log("Preço final:", precoFinal);

let taxa = precoFinal * 0.1; // 10% de taxa
console.log("Taxa:", taxa);

let total = precoFinal + taxa;
console.log("Total a pagar:", total);
```

---

## 🔀 2. Conditional Statements (Declarações Condicionais)

As declarações condicionais permitem que o código execute diferentes ações baseadas em diferentes condições.

### 2.1. A Declaração `if`

A declaração `if` executa um bloco de código **apenas se** uma condição for verdadeira (truthy).

#### Sintaxe Básica

```javascript
if (condição) {
  // código executado se a condição for verdadeira
}
```

#### Exemplo

```javascript
let idade = 18;

if (idade >= 18) {
  console.log("Você é maior de idade!");
}
```

#### Valores Truthy e Falsy

O JavaScript avalia condições como **truthy** (verdadeiro) ou **falsy** (falso):

**Valores Falsy:**
- `false`
- `0`
- `-0`
- `0n` (BigInt zero)
- `""` (string vazia)
- `null`
- `undefined`
- `NaN`

**Valores Truthy:**
- Qualquer outro valor (incluindo objetos, arrays não vazios, strings não vazias, etc.)

```javascript
// Exemplos de condições
if (true) console.log("Executado"); // Executado
if (false) console.log("Não executado"); // Não executado
if (1) console.log("Executado"); // Executado
if (0) console.log("Não executado"); // Não executado
if ("texto") console.log("Executado"); // Executado
if ("") console.log("Não executado"); // Não executado
if ([]) console.log("Executado"); // Executado (array vazio é truthy!)
if ({}) console.log("Executado"); // Executado (objeto vazio é truthy!)
```

### 2.2. A Declaração `if...else`

A declaração `if...else` executa um bloco se a condição for verdadeira, e outro bloco se for falsa.

#### Sintaxe

```javascript
if (condição) {
  // código executado se a condição for verdadeira
} else {
  // código executado se a condição for falsa
}
```

#### Exemplo

```javascript
let idade = 16;

if (idade >= 18) {
  console.log("Você é maior de idade!");
} else {
  console.log("Você é menor de idade!");
}
```

### 2.3. A Declaração `if...else if...else`

Permite verificar múltiplas condições em sequência.

#### Sintaxe

```javascript
if (condição1) {
  // código se condição1 for verdadeira
} else if (condição2) {
  // código se condição2 for verdadeira
} else if (condição3) {
  // código se condição3 for verdadeira
} else {
  // código se nenhuma condição for verdadeira
}
```

#### Exemplo

```javascript
let nota = 85;

if (nota >= 90) {
  console.log("Nota A - Excelente!");
} else if (nota >= 80) {
  console.log("Nota B - Muito bom!");
} else if (nota >= 70) {
  console.log("Nota C - Bom!");
} else if (nota >= 60) {
  console.log("Nota D - Aprovado!");
} else {
  console.log("Nota F - Reprovado!");
}
```

**Importante:** Apenas o primeiro bloco com condição verdadeira será executado. As condições são verificadas em ordem.

### 2.4. Operador Ternário

O operador ternário é uma forma concisa de escrever `if...else` simples.

#### Sintaxe

```javascript
condição ? valorSeVerdadeiro : valorSeFalso
```

#### Exemplo

```javascript
let idade = 20;
let status = idade >= 18 ? "Maior de idade" : "Menor de idade";
console.log(status); // "Maior de idade"

// Equivale a:
let status2;
if (idade >= 18) {
  status2 = "Maior de idade";
} else {
  status2 = "Menor de idade";
}
```

#### Ternário Aninhado

```javascript
let nota = 85;
let conceito = nota >= 90 ? "A" : 
               nota >= 80 ? "B" : 
               nota >= 70 ? "C" : 
               nota >= 60 ? "D" : "F";
console.log(conceito); // "B"
```

**⚠️ Atenção:** Ternários aninhados podem ser difíceis de ler. Use com moderação.

### 2.5. A Declaração `switch`

A declaração `switch` avalia uma expressão e executa código baseado em diferentes casos (cases).

#### Sintaxe

```javascript
switch (expressão) {
  case valor1:
    // código executado se expressão === valor1
    break;
  case valor2:
    // código executado se expressão === valor2
    break;
  case valor3:
    // código executado se expressão === valor3
    break;
  default:
    // código executado se nenhum case corresponder
    break;
}
```

#### Exemplo Básico

```javascript
let diaSemana = 3;
let nomeDia;

switch (diaSemana) {
  case 1:
    nomeDia = "Domingo";
    break;
  case 2:
    nomeDia = "Segunda-feira";
    break;
  case 3:
    nomeDia = "Terça-feira";
    break;
  case 4:
    nomeDia = "Quarta-feira";
    break;
  case 5:
    nomeDia = "Quinta-feira";
    break;
  case 6:
    nomeDia = "Sexta-feira";
    break;
  case 7:
    nomeDia = "Sábado";
    break;
  default:
    nomeDia = "Dia inválido";
    break;
}

console.log(nomeDia); // "Terça-feira"
```

#### Por que `break` é Importante?

Sem `break`, o código continua executando os cases seguintes até encontrar um `break` ou chegar ao final do `switch`. Isso é chamado de **"fall-through"**.

```javascript
let mes = 2;

switch (mes) {
  case 1:
  case 3:
  case 5:
  case 7:
  case 8:
  case 10:
  case 12:
    console.log("31 dias");
    break;
  case 4:
  case 6:
  case 9:
  case 11:
    console.log("30 dias");
    break;
  case 2:
    console.log("28 ou 29 dias");
    break;
  default:
    console.log("Mês inválido");
}
```

#### `switch` com Strings

```javascript
let cor = "vermelho";

switch (cor) {
  case "vermelho":
    console.log("Pare!");
    break;
  case "amarelo":
    console.log("Atenção!");
    break;
  case "verde":
    console.log("Siga!");
    break;
  default:
    console.log("Cor inválida");
}
```

#### Quando Usar `switch` vs `if...else`?

**Use `switch` quando:**
- Comparar uma única variável com múltiplos valores específicos
- Os valores são conhecidos e limitados
- O código fica mais legível

**Use `if...else` quando:**
- As condições são complexas (operadores lógicos, comparações múltiplas)
- Você precisa verificar ranges (faixas de valores)
- As condições não são simples igualdades

---

## ⚠️ 3. Exception Handling (Tratamento de Exceções)

Exception Handling permite que você trate erros de forma controlada, evitando que o programa pare completamente quando algo dá errado.

### 3.1. O que são Exceções?

Em JavaScript, **exceções** são erros que ocorrem durante a execução do código. Quando uma exceção não é tratada, o programa para e exibe uma mensagem de erro.

### 3.2. A Declaração `try...catch`

A declaração `try...catch` permite testar um bloco de código e capturar erros que possam ocorrer.

#### Sintaxe

```javascript
try {
  // código que pode gerar erro
} catch (erro) {
  // código executado se ocorrer um erro
}
```

#### Exemplo Básico

```javascript
try {
  let resultado = 10 / 0;
  console.log("Resultado:", resultado); // Infinity (não é erro)
  
  let texto = null;
  let tamanho = texto.length; // Erro! null não tem propriedade length
} catch (erro) {
  console.log("Ocorreu um erro:", erro.message);
  console.log("Tipo do erro:", erro.name);
}
```

#### Exemplo Prático: Divisão por Zero

```javascript
function dividir(a, b) {
  try {
    if (b === 0) {
      throw new Error("Divisão por zero não é permitida!");
    }
    return a / b;
  } catch (erro) {
    console.log("Erro capturado:", erro.message);
    return null;
  }
}

console.log(dividir(10, 2)); // 5
console.log(dividir(10, 0)); // null (erro tratado)
```

### 3.3. A Declaração `try...catch...finally`

O bloco `finally` é executado **sempre**, independentemente de ocorrer erro ou não.

#### Sintaxe

```javascript
try {
  // código que pode gerar erro
} catch (erro) {
  // tratamento de erro
} finally {
  // código sempre executado
}
```

#### Exemplo

```javascript
function processarArquivo() {
  let arquivo = null;
  
  try {
    arquivo = abrirArquivo("dados.txt");
    processar(arquivo);
  } catch (erro) {
    console.log("Erro ao processar:", erro.message);
  } finally {
    // Sempre fecha o arquivo, mesmo se houver erro
    if (arquivo) {
      fecharArquivo(arquivo);
    }
    console.log("Processamento finalizado");
  }
}
```

### 3.4. A Declaração `throw`

A declaração `throw` permite lançar exceções personalizadas.

#### Sintaxe

```javascript
throw expressão;
```

#### Exemplo

```javascript
function verificarIdade(idade) {
  if (idade < 0) {
    throw new Error("Idade não pode ser negativa!");
  }
  if (idade > 150) {
    throw new Error("Idade inválida! Muito alta.");
  }
  return idade;
}

try {
  verificarIdade(-5);
} catch (erro) {
  console.log("Erro:", erro.message); // "Idade não pode ser negativa!"
}
```

### 3.5. Tipos de Erros em JavaScript

JavaScript possui diferentes tipos de erros, cada um representando um tipo específico de problema.

#### Error (Erro Genérico)

```javascript
try {
  throw new Error("Erro genérico");
} catch (erro) {
  console.log(erro.name); // "Error"
  console.log(erro.message); // "Erro genérico"
}
```

#### ReferenceError

Ocorre quando você tenta acessar uma variável que não existe.

```javascript
try {
  console.log(variavelInexistente);
} catch (erro) {
  if (erro instanceof ReferenceError) {
    console.log("Erro de referência:", erro.message);
    // "variavelInexistente is not defined"
  }
}
```

#### TypeError

Ocorre quando você tenta usar um valor de forma inadequada.

```javascript
try {
  let numero = 10;
  numero.toUpperCase(); // Número não tem método toUpperCase
} catch (erro) {
  if (erro instanceof TypeError) {
    console.log("Erro de tipo:", erro.message);
    // "numero.toUpperCase is not a function"
  }
}
```

#### RangeError

Ocorre quando um valor está fora do range permitido.

```javascript
try {
  let array = new Array(-1); // Tamanho negativo não é permitido
} catch (erro) {
  if (erro instanceof RangeError) {
    console.log("Erro de range:", erro.message);
    // "Invalid array length"
  }
}
```

#### SyntaxError

Ocorre quando há um erro de sintaxe no código.

```javascript
try {
  eval("let x = ;"); // Sintaxe inválida
} catch (erro) {
  if (erro instanceof SyntaxError) {
    console.log("Erro de sintaxe:", erro.message);
  }
}
```

#### Tratando Múltiplos Tipos de Erro

```javascript
function processarDados(dados) {
  try {
    if (!dados) {
      throw new ReferenceError("Dados não fornecidos");
    }
    
    if (typeof dados !== "object") {
      throw new TypeError("Dados devem ser um objeto");
    }
    
    if (dados.length < 0) {
      throw new RangeError("Tamanho inválido");
    }
    
    // Processamento normal
    return dados.map(item => item * 2);
    
  } catch (erro) {
    if (erro instanceof ReferenceError) {
      console.log("Erro de referência:", erro.message);
      return [];
    } else if (erro instanceof TypeError) {
      console.log("Erro de tipo:", erro.message);
      return null;
    } else if (erro instanceof RangeError) {
      console.log("Erro de range:", erro.message);
      return [];
    } else {
      console.log("Erro desconhecido:", erro.message);
      return null;
    }
  }
}
```

### 3.6. Criando Erros Personalizados

Você pode criar suas próprias classes de erro:

```javascript
class ErroValidacao extends Error {
  constructor(mensagem, campo) {
    super(mensagem);
    this.name = "ErroValidacao";
    this.campo = campo;
  }
}

function validarEmail(email) {
  if (!email) {
    throw new ErroValidacao("Email é obrigatório", "email");
  }
  if (!email.includes("@")) {
    throw new ErroValidacao("Email inválido", "email");
  }
  return true;
}

try {
  validarEmail("");
} catch (erro) {
  if (erro instanceof ErroValidacao) {
    console.log(`Erro no campo ${erro.campo}: ${erro.message}`);
  }
}
```

---

## 🔄 4. Combinando Control Flow

Na prática, você combinará diferentes estruturas de control flow:

### Exemplo: Validação de Formulário

```javascript
function validarFormulario(dados) {
  try {
    // Verificações condicionais
    if (!dados.nome) {
      throw new Error("Nome é obrigatório");
    }
    
    if (!dados.email) {
      throw new Error("Email é obrigatório");
    }
    
    if (!dados.email.includes("@")) {
      throw new Error("Email inválido");
    }
    
    if (dados.idade < 18) {
      throw new Error("Idade mínima é 18 anos");
    }
    
    // Processamento
    console.log("Formulário válido!");
    return true;
    
  } catch (erro) {
    console.log("Erro de validação:", erro.message);
    return false;
  } finally {
    console.log("Validação finalizada");
  }
}
```

### Exemplo: Sistema de Notas

```javascript
function calcularConceito(nota) {
  try {
    // Validação
    if (nota < 0 || nota > 100) {
      throw new RangeError("Nota deve estar entre 0 e 100");
    }
    
    if (typeof nota !== "number") {
      throw new TypeError("Nota deve ser um número");
    }
    
    // Lógica condicional
    let conceito;
    if (nota >= 90) {
      conceito = "A";
    } else if (nota >= 80) {
      conceito = "B";
    } else if (nota >= 70) {
      conceito = "C";
    } else if (nota >= 60) {
      conceito = "D";
    } else {
      conceito = "F";
    }
    
    return conceito;
    
  } catch (erro) {
    if (erro instanceof RangeError) {
      console.log("Erro de range:", erro.message);
    } else if (erro instanceof TypeError) {
      console.log("Erro de tipo:", erro.message);
    } else {
      console.log("Erro desconhecido:", erro.message);
    }
    return null;
  }
}
```

---

## 📚 Resumo

Nesta aula você aprendeu:

- ✅ **Sequential Flow**: Execução padrão linha por linha
- ✅ **Conditional Statements**: `if`, `if...else`, `if...else if...else`
- ✅ **Operador Ternário**: Forma concisa de escrever condicionais simples
- ✅ **Switch Statement**: Para múltiplas comparações de igualdade
- ✅ **Exception Handling**: `try...catch...finally` para tratar erros
- ✅ **Throw Statement**: Para lançar exceções personalizadas
- ✅ **Tipos de Erros**: `Error`, `ReferenceError`, `TypeError`, `RangeError`, `SyntaxError`
- ✅ **Quando usar cada estrutura**: Escolha baseada na situação

---

## 🚀 Próximo Passo

Agora que você entende Control Flow, está pronto para a **Aula Simplificada**, onde vamos revisar esses conceitos com analogias e exemplos do dia a dia.

**Arquivo seguinte**: `02-aula-simplificada.md`

