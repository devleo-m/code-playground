# Aula 4: Hooks - O Coração do React Moderno

## 1. A Era dos Hooks (React 16.8+)

Antes de 2019, componentes funcionais eram "stateless" (sem estado). Se você precisasse de estado ou ciclo de vida (como `componentDidMount`), era obrigado a reescrever tudo como uma **Class Component**. Isso gerava códigos verbosos e difíceis de reutilizar.

Os **Hooks** mudaram tudo. Eles permitem que você use recursos do React (como estado e efeitos) dentro de funções simples.

> **Definição Oficial:** "Hooks allow you to reuse stateful logic without changing your component hierarchy."

---

## 2. useState: Gerenciando Memória

O `useState` é o hook fundamental. Ele permite que um componente "lembre" de informações entre renderizações.

### Anatomia
```javascript
const [state, setState] = useState(initialValue);
```

1.  **state**: O valor atual (na renderização atual).
2.  **setState**: A função que atualiza o valor e **agenda um re-render**.
3.  **initialValue**: O valor apenas na *primeira* renderização.

### O "Pulo do Gato": Updates Funcionais
Muitos iniciantes cometem este erro:

```javascript
// Errado se chamado múltiplas vezes rápido
setCount(count + 1);
setCount(count + 1);
setCount(count + 1);
// Resultado: count aumentou apenas 1, pois 'count' era o mesmo valor no escopo.
```

A forma correta quando a atualização depende do valor anterior:

```javascript
// Correto: Functional Update
setCount(prevCount => prevCount + 1);
setCount(prevCount => prevCount + 1);
setCount(prevCount => prevCount + 1);
// Resultado: count aumentou 3.
```

---

## 3. useEffect: Sincronização e Efeitos Colaterais

O `useEffect` é frequentemente mal compreendido como apenas "ciclo de vida" (mount/unmount). Na verdade, ele serve para **sincronizar** seu componente com sistemas externos (DOM, API, Timers, LocalStorage).

> **Definição:** "useEffect is a special hook that lets you run side effects in React. It only runs when the component (or some of its props) changes."

### Estrutura
```javascript
useEffect(() => {
  // 1. Código do efeito (Executa após o render e commit)
  console.log('Sincronizando...');

  // 2. Cleanup Function (Opcional)
  return () => {
    console.log('Limpando antes do próximo efeito ou unmount');
  };
}, [dependencias]); // 3. Array de Dependências
```

### O Array de Dependências (`[]`)
Ele diz ao React: "Só execute este efeito de novo se **estas variáveis** mudarem".

| Dependências | Comportamento | Analogia |
| :--- | :--- | :--- |
| **Sem array** | Executa a cada render. | "Obsessivo" (Verifica tudo sempre) |
| **`[]` (Vazio)** | Executa apenas no Mount (1x). | "Meteoro" (Cai uma vez e pronto) |
| **`[id, data]`** | Executa no Mount e quando `id` OU `data` mudarem. | "Vigilante" (Só age se o alvo mover) |

### Cleanup Function (A Faxina)
Sempre que seu efeito cria algo que persiste (assinatura de evento, timer, conexão socket), você **deve** limpá-lo. Se não, você cria "Memory Leaks" (vazamentos de memória).

```javascript
useEffect(() => {
  const handleResize = () => console.log(window.innerWidth);
  window.addEventListener('resize', handleResize);

  // Cleanup obrigatório!
  return () => {
    window.removeEventListener('resize', handleResize);
  };
}, []);
```

