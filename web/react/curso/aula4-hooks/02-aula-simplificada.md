# Aula 4 - Simplificada: Entendendo Hooks

Imagine que seu componente React é um **Chef de Cozinha** (uma função) que prepara um prato (a UI).

## 1. useState: A Lousa de Pedidos (Memória)

Componentes funcionais normais têm memória curta. Assim que o Chef entrega o prato, ele esquece tudo.
O `useState` é como dar uma **Lousa** para o Chef.

*   Ele pode escrever um número nela (ex: "3 Hambúrgueres").
*   Quando ele volta para fazer o próximo prato, ele olha para a lousa e vê "3".
*   Se ele quiser mudar, ele apaga e escreve "4" (`setState`).
*   **Importante:** A lousa é particular desse Chef. Outros Chefs têm suas próprias lousas.

## 2. useEffect: O Assistente Robô (Efeitos)

O Chef deve focar apenas em cozinhar (Renderizar). Mas às vezes precisamos de coisas "fora" da cozinha: buscar ingredientes no mercado (API), mudar a temperatura do ar condicionado (DOM), ou ligar o timer do forno.

O `useEffect` é um **Assistente Robô** que fica esperando o Chef terminar de servir o prato.

*   **Renderização:** O Chef serve o prato.
*   **Efeito:** O Robô entra e faz a tarefa extra (busca dados, liga timer).

### O Array de Dependências (`[]`) é a Lista de Instruções do Robô

*   **Sem lista:** O Robô faz a tarefa TODA vez que o Chef serve um prato. (Cansativo!)
*   **Lista Vazia `[]`:** O Robô faz a tarefa SÓ na primeira vez que o restaurante abre, e nunca mais. (Bom para buscar dados iniciais).
*   **Lista com itens `[ingrediente]`:** O Robô só vai ao mercado se o `ingrediente` acabou ou mudou. Se for o mesmo, ele fica quieto.

