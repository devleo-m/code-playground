# Aula 10: Funções em JavaScript

## 📚 Visão Geral

Esta aula aborda um dos conceitos mais fundamentais e poderosos do JavaScript: **Funções**. Você aprenderá como criar, usar e otimizar funções para escrever código reutilizável, organizado e eficiente.

---

## 🎯 Objetivos de Aprendizado

Ao final desta aula, você será capaz de:

- ✅ Entender o que são funções e por que são importantes
- ✅ Criar funções usando diferentes sintaxes (declaration, expression, arrow functions)
- ✅ Usar parâmetros padrão e rest parameters
- ✅ Compreender escopo (global, função, bloco)
- ✅ Entender como funciona a call stack
- ✅ Criar funções recursivas
- ✅ Usar funções nativas do JavaScript
- ✅ Escrever funções seguindo boas práticas
- ✅ Otimizar funções para melhor performance

---

## 📖 Conteúdo da Aula

### 1. Aula Principal (`01-aula-principal.md`)
Conteúdo técnico completo sobre funções, incluindo:
- Declaração de funções (function, expression, arrow functions)
- Parâmetros (básicos, padrão, rest)
- Return e valores de retorno
- Escopo e function stack
- Recursão
- Built-in functions

### 2. Aula Simplificada (`02-aula-simplificada.md`)
Versão com analogias do dia a dia:
- Funções como receitas
- Funções como máquinas
- Escopo como quartos de uma casa
- Call stack como pilha de livros
- Recursão como bonecas russas

### 3. Exercícios e Reflexão (`03-exercicios-reflexao.md`)
6 exercícios práticos + 4 perguntas de reflexão sobre:
- Criação de funções
- Parâmetros e validação
- Recursão
- Análise de código
- Eficiência e performance
- Edge cases

### 4. Performance e Boas Práticas (`04-performance-boas-praticas.md`)
Guia completo sobre:
- Otimização de funções
- Nomenclatura adequada
- Organização e estrutura
- Validação e tratamento de erros
- Segurança
- Debugging
- Testabilidade

---

## 💻 Exemplos Práticos

### `exemplo-01-funcoes-basicas.html`
Demonstra:
- Function Declaration
- Function Expression
- Arrow Functions
- Parâmetros padrão
- Rest parameters

### `exemplo-02-escopo-funcoes.html`
Demonstra:
- Escopo global vs local
- Escopo de função
- Escopo de bloco (let vs var)
- Shadowing
- Funções aninhadas

### `exemplo-03-recursao.html`
Demonstra:
- Fatorial recursivo
- Contagem regressiva
- Soma de array recursivo
- Fibonacci recursivo
- Visualização da call stack

### `exemplo-04-built-in-functions.html`
Demonstra:
- Funções globais (parseInt, parseFloat, isNaN, isFinite)
- Métodos de String
- Métodos de Array
- Objeto Math
- Objeto Date
- Funções de tempo (setTimeout, setInterval)

---

## 🗺️ Estrutura de Aprendizado

Siga esta ordem recomendada:

1. **Leia** a Aula Principal (`01-aula-principal.md`)
2. **Revisite** com a Aula Simplificada (`02-aula-simplificada.md`)
3. **Pratique** com os Exemplos HTML (abra no navegador)
4. **Resolva** os Exercícios (`03-exercicios-reflexao.md`)
5. **Aprofunde** com Performance e Boas Práticas (`04-performance-boas-praticas.md`)

---

## 🔑 Conceitos-Chave

### Funções
Blocos de código reutilizáveis que executam tarefas específicas sempre que são invocadas.

### Tipos de Declaração
- **Function Declaration**: `function nome() {}`
- **Function Expression**: `const nome = function() {}`
- **Arrow Function**: `const nome = () => {}`

### Parâmetros
- **Básicos**: Valores passados para a função
- **Padrão**: Valores padrão se nenhum for passado
- **Rest**: Aceita número indefinido de argumentos

### Escopo
- **Global**: Acessível em todo o código
- **Função**: Acessível apenas dentro da função
- **Bloco**: Acessível apenas dentro do bloco `{}`

### Recursão
Função que chama a si mesma, com caso base e caso recursivo.

### Call Stack
Pilha que rastreia quais funções estão sendo executadas.

---

## 📝 Exemplo Rápido

```javascript
// Function Declaration
function somar(a, b) {
  return a + b;
}

// Arrow Function com parâmetros padrão
const saudar = (nome = "Visitante") => {
  return `Olá, ${nome}!`;
};

// Função com rest parameters
const somarTodos = (...numeros) => {
  return numeros.reduce((total, num) => total + num, 0);
};

// Função recursiva
function fatorial(n) {
  if (n <= 1) return 1;
  return n * fatorial(n - 1);
}

// Uso
console.log(somar(5, 3));           // 8
console.log(saudar("Maria"));        // "Olá, Maria!"
console.log(somarTodos(1, 2, 3));     // 6
console.log(fatorial(5));            // 120
```

---

## 🎓 Pré-requisitos

Antes de começar esta aula, certifique-se de que você domina:
- Variáveis (let, const, var)
- Tipos de dados
- Operadores
- Estruturas condicionais
- Loops

---

## 🚀 Próximos Passos

Após dominar funções, você estará pronto para:
- **Closures**: Funções que "lembram" do escopo
- **Higher-Order Functions**: Funções que recebem ou retornam outras funções
- **Callbacks**: Funções passadas como argumentos
- **Promises e async/await**: Programação assíncrona
- **Módulos**: Organização e reutilização de código

---

## 💡 Dicas Importantes

1. **Nomenclatura**: Use nomes descritivos que expliquem o que a função faz
2. **Tamanho**: Mantenha funções pequenas e focadas em uma única tarefa
3. **Validação**: Sempre valide parâmetros de entrada
4. **Performance**: Prefira iteração para loops simples, recursão apenas quando necessário
5. **Testabilidade**: Escreva funções puras quando possível

---

## 📚 Recursos Adicionais

- [MDN: Functions](https://developer.mozilla.org/pt-BR/docs/Web/JavaScript/Guide/Functions)
- [MDN: Arrow Functions](https://developer.mozilla.org/pt-BR/docs/Web/JavaScript/Reference/Functions/Arrow_functions)
- [MDN: Scope](https://developer.mozilla.org/pt-BR/docs/Glossary/Scope)

---

## ✅ Checklist de Conclusão

Antes de avançar para a próxima aula, verifique se você:

- [ ] Consegue criar funções usando as três sintaxes principais
- [ ] Entende a diferença entre parâmetros e argumentos
- [ ] Sabe usar parâmetros padrão e rest parameters
- [ ] Compreende os diferentes tipos de escopo
- [ ] Consegue criar funções recursivas simples
- [ ] Sabe quando usar recursão vs iteração
- [ ] Conhece as principais funções nativas do JavaScript
- [ ] Consegue identificar e corrigir problemas comuns em funções
- [ ] Entende a importância de validação e tratamento de erros
- [ ] Sabe aplicar boas práticas de nomenclatura e organização

---

**Bons estudos! 🚀**

*Lembre-se: A prática é essencial. Não tenha pressa e entenda cada conceito completamente antes de prosseguir.*

