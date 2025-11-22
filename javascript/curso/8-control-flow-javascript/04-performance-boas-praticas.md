# Aula 8 - Performance, Boas Práticas e Otimização: Control Flow

## ⚡ Performance: Escolhendo a Estrutura Correta

### Comparação de Performance: `if...else` vs `switch`

Diferentes estruturas condicionais têm diferentes características de performance. Vamos analisar:

#### 1. `if...else` vs `switch` - Performance

```javascript
const valor = 5;

// Teste 1: if...else com múltiplas condições
console.time("if...else");
let resultado1;
if (valor === 1) {
  resultado1 = "um";
} else if (valor === 2) {
  resultado1 = "dois";
} else if (valor === 3) {
  resultado1 = "três";
} else if (valor === 4) {
  resultado1 = "quatro";
} else if (valor === 5) {
  resultado1 = "cinco";
} else {
  resultado1 = "outro";
}
console.timeEnd("if...else");

// Teste 2: switch
console.time("switch");
let resultado2;
switch (valor) {
  case 1:
    resultado2 = "um";
    break;
  case 2:
    resultado2 = "dois";
    break;
  case 3:
    resultado2 = "três";
    break;
  case 4:
    resultado2 = "quatro";
    break;
  case 5:
    resultado2 = "cinco";
    break;
  default:
    resultado2 = "outro";
    break;
}
console.timeEnd("switch");
```

**Resultados típicos:**
- **`switch`**: Geralmente mais rápido quando há muitos casos (5+), pois o JavaScript pode otimizar usando "jump tables"
- **`if...else`**: Mais rápido para poucos casos (2-3), pois não há overhead de estrutura
- **Diferença prática**: Em casos simples, a diferença é desprezível (< 1ms). Em loops com milhões de iterações, `switch` pode ser 10-20% mais rápido

**Conclusão:** Para poucos casos (2-4), use `if...else`. Para muitos casos (5+), use `switch`. A diferença só importa em código executado milhões de vezes.

---

### 2. Ordem das Condições em `if...else`

A ordem das condições importa para performance:

```javascript
// ❌ RUIM - Condição mais provável por último
function verificarStatus(nota) {
  if (nota < 60) return "F";      // Menos comum
  else if (nota < 70) return "D";  // Menos comum
  else if (nota < 80) return "C";  // Menos comum
  else if (nota < 90) return "B";  // Menos comum
  else return "A";                 // Mais comum, mas verificado por último!
}

// ✅ BOM - Condição mais provável primeiro
function verificarStatus(nota) {
  if (nota >= 90) return "A";     // Mais comum primeiro
  else if (nota >= 80) return "B";
  else if (nota >= 70) return "C";
  else if (nota >= 60) return "D";
  else return "F";
}
```

**Regra de Ouro:** Coloque as condições mais prováveis primeiro. Isso reduz o número médio de verificações.

---

### 3. Operador Ternário vs `if...else`

```javascript
// Teste 1: if...else
console.time("if...else");
let resultado1;
if (idade >= 18) {
  resultado1 = "maior";
} else {
  resultado1 = "menor";
}
console.timeEnd("if...else");

// Teste 2: Operador ternário
console.time("ternário");
let resultado2 = idade >= 18 ? "maior" : "menor";
console.timeEnd("ternário");
```

**Resultados:** Performance idêntica. A diferença é apenas de legibilidade.

**Conclusão:** Use ternário para decisões simples e `if...else` para lógica complexa.

---

### 4. Tratamento de Erros: Impacto na Performance

```javascript
// ❌ RUIM - try...catch em loop sem necessidade
function processarArray(array) {
  let soma = 0;
  for (let i = 0; i < array.length; i++) {
    try {
      soma += array[i];
    } catch (erro) {
      // Nunca vai executar aqui
    }
  }
  return soma;
}

// ✅ BOM - Validação antes do loop
function processarArray(array) {
  if (!Array.isArray(array)) {
    throw new TypeError("Array inválido");
  }
  
  let soma = 0;
  for (let i = 0; i < array.length; i++) {
    soma += array[i];
  }
  return soma;
}
```

**Impacto:** `try...catch` tem overhead. Use apenas quando realmente necessário.

**Regra:** Valide dados **antes** de processar, não durante.

---

## 🎯 Boas Práticas

### 1. Use `if...else` para Condições Complexas

```javascript
// ✅ BOM - if...else para condições complexas
if (idade >= 18 && temDocumento && !estaBloqueado) {
  permitirAcesso();
} else {
  negarAcesso();
}

// ❌ EVITE - switch não funciona bem com condições complexas
// switch não suporta operadores lógicos diretamente
```

---

### 2. Use `switch` para Múltiplas Igualdades

```javascript
// ✅ BOM - switch para múltiplas comparações de igualdade
switch (diaSemana) {
  case 1:
  case 7:
    console.log("Fim de semana");
    break;
  case 2:
  case 3:
  case 4:
  case 5:
  case 6:
    console.log("Dia útil");
    break;
  default:
    console.log("Dia inválido");
}

// ⚠️ ACEITÁVEL - if...else funciona, mas menos legível
if (diaSemana === 1 || diaSemana === 7) {
  console.log("Fim de semana");
} else if (diaSemana >= 2 && diaSemana <= 6) {
  console.log("Dia útil");
} else {
  console.log("Dia inválido");
}
```

---

### 3. Sempre Use `break` no `switch` (Exceto em Fall-Through Intencional)

```javascript
// ✅ BOM - break explícito
switch (mes) {
  case 1:
    console.log("Janeiro");
    break;
  case 2:
    console.log("Fevereiro");
    break;
  default:
    console.log("Mês inválido");
    break;
}

// ❌ RUIM - Esqueceu o break (bug!)
switch (mes) {
  case 1:
    console.log("Janeiro");
    // Sem break - vai executar o próximo case também!
  case 2:
    console.log("Fevereiro");
    break;
}
```

**Dica:** Use ESLint com regra `no-fallthrough` para detectar isso automaticamente.

---

### 4. Evite Ternários Aninhados Excessivos

```javascript
// ❌ RUIM - Ternário aninhado difícil de ler
const status = nota >= 90 ? "A" : nota >= 80 ? "B" : nota >= 70 ? "C" : nota >= 60 ? "D" : "F";

// ✅ BOM - if...else mais legível
let status;
if (nota >= 90) {
  status = "A";
} else if (nota >= 80) {
  status = "B";
} else if (nota >= 70) {
  status = "C";
} else if (nota >= 60) {
  status = "D";
} else {
  status = "F";
}

// ✅ TAMBÉM BOM - switch (se apropriado)
let status;
switch (Math.floor(nota / 10)) {
  case 10:
  case 9: status = "A"; break;
  case 8: status = "B"; break;
  case 7: status = "C"; break;
  case 6: status = "D"; break;
  default: status = "F";
}
```

**Regra:** Use ternário para 1-2 níveis. Para mais, use `if...else` ou `switch`.

---

### 5. Tratamento de Erros: Estratégias

#### Validação Antecipada (Early Return)

```javascript
// ❌ RUIM - Aninhamento excessivo
function processarDados(dados) {
  if (dados) {
    if (Array.isArray(dados)) {
      if (dados.length > 0) {
        // Processamento aqui (muito aninhado!)
      } else {
        throw new Error("Array vazio");
      }
    } else {
      throw new Error("Não é array");
    }
  } else {
    throw new Error("Dados não fornecidos");
  }
}

// ✅ BOM - Early return (retorno antecipado)
function processarDados(dados) {
  if (!dados) {
    throw new Error("Dados não fornecidos");
  }
  
  if (!Array.isArray(dados)) {
    throw new Error("Não é array");
  }
  
  if (dados.length === 0) {
    throw new Error("Array vazio");
  }
  
  // Processamento aqui (sem aninhamento!)
  return dados.map(item => item * 2);
}
```

#### Tipos de Erro Específicos

```javascript
// ✅ BOM - Erros específicos
function validarIdade(idade) {
  if (typeof idade !== "number") {
    throw new TypeError("Idade deve ser um número");
  }
  
  if (idade < 0) {
    throw new RangeError("Idade não pode ser negativa");
  }
  
  if (idade > 150) {
    throw new RangeError("Idade inválida (muito alta)");
  }
  
  return true;
}

// ❌ RUIM - Erro genérico
function validarIdade(idade) {
  if (typeof idade !== "number" || idade < 0 || idade > 150) {
    throw new Error("Idade inválida"); // Muito genérico!
  }
  return true;
}
```

---

### 6. Use `finally` para Limpeza de Recursos

```javascript
// ✅ BOM - finally garante limpeza
function processarArquivo(nomeArquivo) {
  let arquivo = null;
  
  try {
    arquivo = abrirArquivo(nomeArquivo);
    processar(arquivo);
  } catch (erro) {
    console.error("Erro ao processar:", erro);
    throw erro; // Re-lança o erro
  } finally {
    // Sempre fecha, mesmo com erro
    if (arquivo) {
      fecharArquivo(arquivo);
    }
  }
}

// ❌ RUIM - Pode esquecer de fechar
function processarArquivo(nomeArquivo) {
  let arquivo = abrirArquivo(nomeArquivo);
  try {
    processar(arquivo);
  } catch (erro) {
    console.error("Erro:", erro);
    // E se houver erro? Arquivo não fecha!
  }
  fecharArquivo(arquivo); // Só executa se não houver erro
}
```

---

### 7. Nomenclatura Clara em Condições

```javascript
// ❌ RUIM - Condição confusa
if (!user || !user.active || user.banned) {
  // O que isso significa?
}

// ✅ BOM - Variável com nome descritivo
const usuarioInvalido = !user || !user.active || user.banned;
if (usuarioInvalido) {
  negarAcesso();
}

// ✅ TAMBÉM BOM - Função com nome descritivo
function podeAcessar(usuario) {
  return usuario && usuario.active && !usuario.banned;
}

if (!podeAcessar(user)) {
  negarAcesso();
}
```

---

### 8. Evite Condições Múltiplas Complexas

```javascript
// ❌ RUIM - Condição muito complexa
if ((idade >= 18 && temDocumento) || (idade >= 16 && temAutorizacao && !estaBloqueado) || (idade < 16 && temResponsavel)) {
  permitirAcesso();
}

// ✅ BOM - Extrair para função
function podeAcessar(idade, temDocumento, temAutorizacao, estaBloqueado, temResponsavel) {
  if (idade >= 18 && temDocumento) return true;
  if (idade >= 16 && temAutorizacao && !estaBloqueado) return true;
  if (idade < 16 && temResponsavel) return true;
  return false;
}

if (podeAcessar(idade, temDocumento, temAutorizacao, estaBloqueado, temResponsavel)) {
  permitirAcesso();
}
```

---

## 🔒 Segurança

### 1. Validação de Entrada do Usuário

```javascript
// ✅ BOM - Validação rigorosa
function processarEntrada(entrada) {
  // Validação de tipo
  if (typeof entrada !== "string") {
    throw new TypeError("Entrada deve ser string");
  }
  
  // Validação de conteúdo
  if (entrada.trim().length === 0) {
    throw new Error("Entrada não pode estar vazia");
  }
  
  // Sanitização (remover caracteres perigosos)
  const sanitizada = entrada.replace(/[<>]/g, "");
  
  return sanitizada;
}
```

### 2. Proteção contra XSS em Condicionais

```javascript
// ❌ PERIGOSO - Pode permitir XSS
if (usuario.role === "admin") {
  elemento.innerHTML = usuario.nome; // Perigoso se usuario.nome contém HTML malicioso
}

// ✅ SEGURO - Usar textContent
if (usuario.role === "admin") {
  elemento.textContent = usuario.nome; // Seguro
}
```

---

## 🧪 Debugging

### 1. Console Logging Estratégico

```javascript
// ✅ BOM - Logs úteis para debugging
function processarDados(dados) {
  console.log("Iniciando processamento:", dados);
  
  try {
    if (!dados) {
      throw new Error("Dados não fornecidos");
    }
    
    const resultado = dados.map(item => item * 2);
    console.log("Processamento concluído:", resultado);
    return resultado;
    
  } catch (erro) {
    console.error("Erro no processamento:", {
      mensagem: erro.message,
      tipo: erro.name,
      dados: dados
    });
    throw erro;
  }
}
```

### 2. Usando DevTools para Analisar Performance

```javascript
// Marcar início
performance.mark("inicio-processamento");

// Seu código aqui
processarDados(arrayGrande);

// Marcar fim
performance.mark("fim-processamento");

// Medir
performance.measure("processamento", "inicio-processamento", "fim-processamento");

// Ver no DevTools > Performance
```

---

## 📊 Otimização: Quando Importa

### Quando Otimizar

- ✅ Código executado milhões de vezes (loops internos)
- ✅ Aplicações em tempo real (jogos, animações)
- ✅ Processamento de grandes volumes de dados
- ✅ Operações críticas de performance

### Quando NÃO Otimizar Prematuramente

- ❌ Código executado poucas vezes
- ❌ Protótipos e código experimental
- ❌ Quando a legibilidade é mais importante

**Regra de Ouro:** Primeiro, faça funcionar. Depois, meça. Por último, otimize.

---

## 🎓 Padrões de Código

### 1. Guard Clauses (Cláusulas de Guarda)

```javascript
// ✅ BOM - Guard clauses (early returns)
function processarUsuario(usuario) {
  if (!usuario) return null;
  if (!usuario.email) return null;
  if (!usuario.ativo) return null;
  
  // Código principal aqui (sem aninhamento)
  return enviarEmail(usuario.email);
}

// ❌ RUIM - Aninhamento excessivo
function processarUsuario(usuario) {
  if (usuario) {
    if (usuario.email) {
      if (usuario.ativo) {
        return enviarEmail(usuario.email);
      }
    }
  }
  return null;
}
```

### 2. Strategy Pattern com Switch

```javascript
// ✅ BOM - Estratégias organizadas
const estrategias = {
  "vip": (preco) => preco * 0.8,      // 20% desconto
  "premium": (preco) => preco * 0.85, // 15% desconto
  "regular": (preco) => preco * 0.95,  // 5% desconto
  "default": (preco) => preco
};

function calcularPreco(preco, tipo) {
  const estrategia = estrategias[tipo] || estrategias.default;
  return estrategia(preco);
}
```

---

## 📚 Resumo de Boas Práticas

1. ✅ Use `if...else` para condições complexas
2. ✅ Use `switch` para múltiplas igualdades (5+ casos)
3. ✅ Coloque condições mais prováveis primeiro
4. ✅ Sempre use `break` no `switch` (exceto fall-through intencional)
5. ✅ Evite ternários aninhados excessivos
6. ✅ Use early return para reduzir aninhamento
7. ✅ Use tipos de erro específicos
8. ✅ Use `finally` para limpeza de recursos
9. ✅ Valide dados antes de processar
10. ✅ Nomeie condições complexas com variáveis ou funções

---

## 🚀 Próximo Passo

Agora que você entende performance e boas práticas de Control Flow, você está pronto para praticar e receber feedback!

**Lembre-se:** Complete os exercícios da aula anterior antes de avançar para a próxima aula.

