# Aula 3 - Exercícios e Reflexão

Para fixar o conhecimento sobre Rendering e Ciclo de Vida, realize os exercícios abaixo no projeto prático `@fundaments`.

---

## 🛠️ Exercícios Práticos

### Exercício 1: O Detetive de Renderização (Lifecycle)
Crie um componente chamado `LifecycleLogger`.
1.  Ele deve receber uma prop `message`.
2.  Use `useEffect` para logar no console:
    *   "Componente Montado" (apenas uma vez).
    *   "Mensagem mudou para: [nova mensagem]" (sempre que a prop mudar).
    *   "Componente Desmontado" (quando sair da tela).
3.  Adicione um botão no componente Pai para mostrar/esconder esse logger e testar o Unmount.

### Exercício 2: A Lista Quebrada (Keys)
1.  Crie uma lista de inputs onde o usuário pode digitar algo.
2.  Use o **Index** do array como `key`.
3.  Adicione um botão para remover o **primeiro** item da lista.
4.  Digite valores diferentes em cada input (ex: "A", "B", "C").
5.  Remova o primeiro item. Observe o que acontece com os valores dos inputs. O "B" virou "A"? O valor se manteve no input errado?
6.  **Correção:** Refatore para usar um ID único (`Date.now()` ou UUID) como key e veja o bug desaparecer.

### Exercício 3: Mouse Tracker (Render Props)
1.  Crie um componente `MouseTracker` que escuta o evento `mousemove` na janela.
2.  Ele não deve renderizar nada visual por si só.
3.  Ele deve receber uma prop `render` (função) e passar as coordenadas `{x, y}` para ela.
4.  Use-o para renderizar um "fantasma" (emoji 👻) que segue o mouse pela tela.

---

## 🤔 Perguntas de Reflexão

Responda para si mesmo ou anote:

1.  **Por que não devemos colocar chamadas de API (fetch) diretamente no corpo do componente, fora do `useEffect`?**
    *   *Dica: Pense no que acontece a cada renderização.*

2.  **Se o React é tão rápido, por que precisamos usar `key` em listas? Ele não poderia apenas comparar o conteúdo HTML?**
    *   *Dica: Pense em estado local (inputs, foco).*

3.  **Qual a diferença entre `useRef` e `useState`? Se eu quiser guardar um valor que, ao mudar, NÃO deve atualizar a tela, qual devo usar?**

4.  **No `useEffect`, o que acontece se eu esquecer de passar o array de dependências `[]`? E se eu passar o array vazio?**

