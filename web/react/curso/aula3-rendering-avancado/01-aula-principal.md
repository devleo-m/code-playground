# Aula 3: Rendering e Conceitos Avançados do React

## Introdução

Nesta aula, mergulharemos no coração do React. Vamos entender não apenas *como* usar a biblioteca, mas *como ela funciona* por baixo dos panos. Isso é o que separa um programador que "sabe React" de um Engenheiro de Software React capaz de otimizar aplicações complexas.

---

## 1. O Processo de Renderização (The Rendering Process)

### 1.1 Declarativo vs Imperativo

O conceito mais fundamental do React é sua natureza **Declarativa**.

*   **Imperativo (Como fazer):** Você dá instruções passo a passo.
    *   Exemplo (jQuery/JS Vanilla): "Selecione a div X, crie um elemento P, adicione o texto 'Olá', anexe P em X, mude a cor de fundo de X para azul."
*   **Declarativo (O que eu quero):** Você descreve o estado final da interface.
    *   Exemplo (React): "Eu quero uma div azul com um parágrafo dizendo 'Olá' dentro."

O React cuida de *como* chegar nesse estado.

### 1.2 O Ciclo de Renderização

Quando dizemos que o React "renderiza", isso acontece em três fases:

1.  **Render Phase (Fase de Renderização):**
    *   O React chama seus componentes (funções).
    *   Seus componentes retornam JSX (que vira objetos JavaScript - `React.createElement`).
    *   O React constrói uma nova árvore de **Virtual DOM**.
    *   **Reconciliation (Reconciliação):** O React compara (diffing) essa nova árvore com a anterior para descobrir o que mudou.

2.  **Commit Phase (Fase de Commit):**
    *   O React pega as mudanças calculadas e as aplica ao **DOM Real** (o navegador).
    *   É aqui que o usuário vê as mudanças.
    *   Criação, atualização e remoção de nós do DOM acontecem aqui.

3.  **Passive Effects Phase (Fase de Efeitos):**
    *   Após o DOM ser atualizado, o React roda os `useEffects`.

---

## 2. Component Life Cycle (Ciclo de Vida)

Todo componente React tem um ciclo de vida. Com a chegada dos Hooks, pensamos menos em "métodos de ciclo de vida" e mais em "sincronização com o estado", mas os momentos ainda existem:

### 2.1 Mounting (Montagem - Nascimento)
Acontece quando o componente é criado e inserido no DOM pela primeira vez.
*   Inicialização do State.
*   Primeira Renderização.
*   Execução do `useEffect` (com array de dependências vazio `[]`).

### 2.2 Updating (Atualização - Vida)
Acontece quando:
*   As **Props** mudam.
*   O **State** muda.
*   O componente **Pai** re-renderiza.

Nesta fase:
*   O componente re-executa.
*   O Virtual DOM é comparado.
*   `useEffect` roda novamente (se as dependências mudaram).

### 2.3 Unmounting (Desmontagem - Morte)
Acontece quando o componente é removido da tela (ex: renderização condicional `false` ou navegação de rota).
*   Funções de limpeza (Cleanup) do `useEffect` rodam.
*   O componente sai da memória.

---

## 3. Listas e Keys (Chaves)

Ao renderizar listas, o React precisa de uma forma de rastrear cada item individualmente para manter a performance e o estado correto.

### O Problema do Índice (Index)
```jsx
// 🚫 EVITE ISSO SE A LISTA MUDA DE ORDEM
{items.map((item, index) => <li key={index}>{item}</li>)}
```
Se você deletar o item 0, o item 1 vira o 0. O React vai achar que o item 0 *mudou de conteúdo*, em vez de perceber que o item 0 *sumiu* e o 1 subiu. Isso causa bugs bizarros em inputs e perda de performance.

### A Solução: IDs Únicos e Estáveis
```jsx
// ✅ USE IDs DO BANCO DE DADOS OU GERE IDs ÚNICOS
{items.map((item) => <li key={item.id}>{item}</li>)}
```
A `key` deve ser:
1.  **Única** (entre os irmãos).
2.  **Estável** (não deve mudar a cada render - evite `Math.random()` na key).

---

## 4. Render Props

Um padrão avançado para compartilhar lógica entre componentes. Em vez de passar um dado simples, você passa uma **função que retorna JSX**.

```jsx
<MouseTracker render={({ x, y }) => (
  <h1>O mouse está em {x}, {y}</h1>
)} />
```
O componente `MouseTracker` não sabe *o que* vai renderizar, ele apenas gerencia a lógica (posição do mouse) e "delega" a renderização para quem o chamou.

---

## 5. Refs (Referências)

`useRef` é uma "porta de escape" do fluxo declarativo do React.

### Casos de Uso:
1.  **Acessar o DOM Real:** Focar um input, medir o tamanho de uma div, integrar com libs de terceiros (D3, mapas).
2.  **Persistência de Valor:** Guardar um valor que deve sobreviver a re-renderizações mas **não deve causar** uma re-renderização quando mudar (ex: armazenar um ID de `setInterval`).

---

## 6. Events (Eventos)

O React usa **SyntheticEvents**.
*   É um wrapper em volta dos eventos nativos do browser.
*   Garante que `e.preventDefault()` ou `e.stopPropagation()` funcionem igual no Chrome, Firefox, Safari, etc.
*   Event Delegation: O React na verdade anexa um único ouvinte de evento na raiz do app para otimizar a memória, em vez de um em cada botão.

---

## Resumo

Entender Rendering e Ciclo de Vida é crucial para não cair na armadilha de "Componentes que renderizam demais" ou "Efeitos que rodam em loop infinito". Na próxima aula, focaremos em como medir e otimizar isso.

