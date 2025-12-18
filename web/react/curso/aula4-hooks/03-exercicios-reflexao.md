# Aula 4 - Exercícios e Reflexão

## Exercício 1: O Cronômetro
Crie um componente `Cronometro` que:
1.  Tenha um estado `segundos` (inicia em 0).
2.  Use `useEffect` para iniciar um `setInterval` que aumenta os segundos a cada 1000ms.
3.  **Desafio:** Garanta que o timer pare (clearInterval) se o componente for removido da tela (unmount).
4.  Adicione botões "Pausar" e "Zerar".

## Exercício 2: Fetch com Feedback
Crie um componente `UsuarioGitHub` que recebe uma prop `username`.
1.  Use `useEffect` para buscar dados da API `https://api.github.com/users/{username}` sempre que a prop `username` mudar.
2.  Enquanto espera, mostre "Carregando...".
3.  Se der erro (ou usuário não existir), mostre "Erro".
4.  Se der certo, mostre o avatar e o nome.

---

## Reflexão (Responda para si mesmo)

1.  **Por que não podemos colocar Hooks dentro de `if` ou `for`?** (Dica: O React confia na ordem das chamadas).
2.  No Exercício 1, o que acontece se você esquecer de passar o array de dependências `[]` no `useEffect` do `setInterval`? E se você esquecer o `clearInterval`?
3.  Qual a diferença entre `useState(0)` e `const [count, setCount] = useState(() => 0)`? (Lazy initialization).

