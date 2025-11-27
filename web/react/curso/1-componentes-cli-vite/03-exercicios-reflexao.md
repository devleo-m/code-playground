# Aula 1 - Exercícios e Reflexão

## Exercícios práticos
1. **CLI e scaffolding**: use `npm create vite@latest meu-componente -- --template react-ts` para gerar um projeto mínimo e liste quais arquivos e pastas o CLI criou. Explique em poucas frases como cada um desses arquivos ajuda a manter os componentes organizados.
2. **Componente funcional simples**: dentro do `src`, crie um arquivo `Plano.tsx` que exporte um componente funcional. Ele deve receber `nome`, `preco` e `descricao` via props e renderizar esses dados em um card estilizado. Foque apenas em JSX e props; não tente adicionar estado ainda.
3. **Composição deliberada**: com o `Plano` criado, construa outro componente `SessaoPlanos` que reutilize `Plano` três vezes com dados diferentes. Mostre como a composição permite montar UI complexa sem herança.

### Dicas para registrar o que você fez
- Copie e cole o `Plano.tsx` e o `SessaoPlanos.tsx` dentro de `src/components` e organize com um CSS simples (pode ser inline ou um `.module.css`).
- Tire uma captura rápida da tela do Vite rodando `npm run dev` antes de enviar; isso ajuda a mostrar que o projeto sobe sem erros.
- Se quiser, adicione comentários curtos no código (`// props imutáveis`) para lembrar da intenção das props.

## Perguntas de reflexão
- Como este componente respeita a regra de responsabilidade única?
- Quando ainda vale a pena passar dados via props e quando vai ser necessário um estado local (mesmo que não esteja sendo definido nesta aula)?
- Quais trade-offs ocorrem se você tentar criar um componente gigante em vez de dividir em blocos menores?
- Este padrão escala bem em aplicações maiores? Que sinais ajudam a identificar quando quebrar em novos componentes?

