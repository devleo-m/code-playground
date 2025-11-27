# Aula 1 - Simplificada: Entendendo Componentes e CLI Tools

## CLI Tools como oficina
Imagine que você está montando um kit LEGO: CLI Tools são as instruções de montagem que distribuem as peças certas no lugar certo. O comando `npm create vite@latest` é como abrir a caixa que já vem com blocos (pastas prontas, configuração de TypeScript e scripts). Assim, você não perde tempo procurando onde colocar `index.tsx` ou como ligar o servidor.

## Vite como motor turbo
Vite é o motor que acelera a visualização do seu componente. Em vez de esperar o Webpack empacotar tudo, ele usa módulos ES para trocar apenas as peças modificadas. Imagine um mapa-múndi digital que só redesenha o continente onde você mexeu o pincel.

## Componentes são peças de receita
Um componente funcional é como uma receita de bolo escrita em TypeScript. Ele recebe ingredientes (`props`) e devolve o bolo pronto (JSX). Se você recebe `titulo` e `subtitulo`, monta o card com eles. Se não vem `subtitulo`, simplesmente pula essa parte. Cada componente cuida de uma tarefa pequena e clara.

### Exemplo rapidinho
```tsx
const Cartao = ({ titulo }: { titulo: string }) => (
  <article>
    <h2>{titulo}</h2>
    <p>Este cartão renderiza apenas o título que recebe via props.</p>
  </article>
);
```

Esse componente não faz nada além de mostrar o que recebeu. Por isso chamamos de função pura: mesma entrada, mesmo resultado.

## Props vs estado com linguagem do dia-a-dia
Props são como ingredientes que vêm da cozinha do chefe (o componente pai). Você só lê e usa. Estado (state) é como a gaveta da sua bancada: guarda o que só aquele componente controla. Nesta aula, vamos reconhecer essa diferença sem mexer na gaveta ainda — os hooks que fazem isso serão objetos de aulas futuras.

## Por enquanto, sem renderização explicada
Ainda não vamos analisar como o React atualiza a tela. O importante agora é saber que, quando um componente recebe props, ele monta o JSX correspondente como uma foto instantânea. Renderização, memoização e hooks ficam para o segundo capítulo dessa trilha.

## Dicas de amigo
- Quando você salva um arquivo, observe o terminal do Vite: ele mostra se houve erro de sintaxe ou importação.
- Nomeie seus componentes com sentido (`BotaoPrimario` em vez de `Componente1`).
- Se quiser testar rapidamente um componente, copie o JSX para `App.tsx` e veja como ele se comporta no navegador.
