# Aula 4 - Performance e Boas Práticas

Hooks são poderosos, mas têm regras estritas ("Rules of Hooks"). Quebrá-las causa bugs bizarros.

## 1. As Regras de Ouro

1.  **Apenas no Top Level:** Nunca chame Hooks dentro de loops, condições (`if`) ou funções aninhadas. O React depende da ordem exata de chamada dos Hooks para saber qual estado pertence a qual variável.
    *   *Errado:* `if (condicao) { useState() }`
    *   *Certo:* `const [val, set] = useState(); if (condicao) ...`
2.  **Apenas em Funções React:** Não chame Hooks em funções JS comuns. Apenas em Componentes ou Custom Hooks (`useSomething`).

## 2. A Mentira das Dependências
Um erro comum de performance (e lógica) é **mentir** para o React no array de dependências do `useEffect`.

```javascript
// ERRO COMUM
useEffect(() => {
  console.log(usuario.nome);
}, []); // Você usou 'usuario' mas disse que não depende de nada!
```

O ESLint (`react-hooks/exhaustive-deps`) vai te avisar. **Obedeça o ESLint.** Se ele pedir para adicionar uma dependência e isso causar loops infinitos, o problema está na sua lógica, não na regra.

## 3. Limpeza de Efeitos (Memory Leaks)
Sempre limpe assinaturas. Em SPAs (Single Page Applications), se você navega de uma página para outra, os componentes velhos morrem, mas os `setInterval` ou `addEventListener` que eles criaram continuam vivos comendo memória se você não der `clearInterval` ou `removeEventListener` no retorno do `useEffect`.

## 4. useState Lazy Initialization
Se o estado inicial for pesado para calcular (ex: ler localStorage complexo), use a forma de função para que o cálculo só ocorra na primeira renderização:

```javascript
// Executa o parse TODA renderização (lento)
const [state, setState] = useState(JSON.parse(localStorage.getItem('bigData')));

// Executa SÓ na primeira (rápido)
const [state, setState] = useState(() => JSON.parse(localStorage.getItem('bigData')));
```

