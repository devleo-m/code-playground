# Aula 3 - Performance, Boas Práticas e Otimização

Rendering é o custo mais caro de uma aplicação Frontend. Controlar quem renderiza e quando renderiza é a chave para apps de 60 FPS.

---

## 1. Otimização de Re-renderização

### O Problema
Por padrão, quando um componente Pai renderiza, **todos** os seus Filhos renderizam também, recursivamente. Mesmo que as props dos filhos não tenham mudado.

### A Solução: `React.memo`
`React.memo` é um Higher-Order Component (HOC) que memoiza o componente. Ele faz uma comparação superficial (shallow compare) das props. Se as props não mudaram, o React **pula** a renderização desse componente.

```jsx
const MeuComponenteCaro = React.memo(function MeuComponente({ dados }) {
  // Só renderiza se 'dados' mudar
  return <div>{dados}</div>;
});
```

### A Pegadinha das Funções e Objetos
Em JS, `{}` !== `{}` e `function(){}` !== `function(){}`.
Se você passa uma função ou objeto criado dentro do Pai para um componente memoizado, o `React.memo` vai falhar, porque a referência muda a cada render.

**Solução:** `useCallback` (para funções) e `useMemo` (para objetos/valores).

## 2. Limpeza de Efeitos (Cleanup)

Vazamento de memória (Memory Leak) é real no Frontend.
Se você cria um `setInterval` ou um `addEventListener` no `useEffect` e não o limpa no retorno, cada vez que o componente desmontar e montar, você cria um novo ouvinte.
Em 10 minutos de uso, seu app fica lento.

**Boas Práticas:**
*   Sempre retorne uma função de limpeza se o efeito cria uma subscrição/timer.
*   Use ferramentas como ESLint (`exhaustive-deps`) para garantir que você não esqueceu dependências.

## 3. Virtualização de Listas

Se você tem uma lista com 1.000 ou 10.000 itens, o React vai sofrer para criar 10.000 nós no DOM, mesmo com keys perfeitas. O navegador não aguenta tanto DOM.

**Solução:** **Windowing (ou Virtualization).**
Renderize apenas o que está visível na tela (ex: 10 itens). Conforme o usuário rola, recicle os componentes.
*   Libs recomendadas: `react-window` ou `react-virtualized`.

## 4. Code Splitting e Lazy Loading

Não carregue o código da "Página de Configurações" se o usuário está na "Home".
Use `React.lazy` e `Suspense` para dividir o bundle.

```jsx
const ConfigPage = React.lazy(() => import('./ConfigPage'));

function App() {
  return (
    <Suspense fallback={<div>Carregando...</div>}>
      <ConfigPage />
    </Suspense>
  );
}
```

## 5. Profiler (React DevTools)

Não adivinhe onde está lento. Use a aba **Profiler** no React DevTools.
*   Ative a opção "Highlight updates when components render".
*   Grave uma sessão de uso.
*   Veja quais componentes renderizaram sem necessidade (barras cinzas/amarelas) e porquê ("Did not change props" vs "Changed props").

---

## Checklist de Qualidade

1.  [ ] Estou usando keys únicas e estáveis em todas as listas?
2.  [ ] Limpei todos os meus timers e event listeners no `useEffect`?
3.  [ ] Estou criando componentes DENTRO de outros componentes? (NÃO FAÇA ISSO! Isso destroi o estado e remonta o componente a cada render).
4.  [ ] Estou passando objetos literais como props para componentes puros? (Use useMemo se necessário).

