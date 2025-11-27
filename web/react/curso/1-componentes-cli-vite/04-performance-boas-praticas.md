# Aula 1 - Performance, Boas Práticas e Otimização

## Perfil e re-render perceptivo
Mesmo sem entrar em hooks, já podemos pensar em performance ao manter componentes pequenos. No React DevTools você verá que cada componente re-renderiza quando seu pai muda. Para evitar re-renders desnecessários:

- envolva apenas o JSX necessário em cada componente;
- prefira props imutáveis (strings, números, objetos constantes) e evite recriar funções dentro do JSX;
- mantenha ativado o **Fast Refresh** do Vite para que o browser recarregue só o que mudou.

### Exemplo sensível
Se você criar um componente `Selo` dentro do `App.tsx`, abra o DevTools e veja que ele re-renderiza toda vez que o pai muda. Esse padrão é esperado, mas dá para reduzir se o `Selo` ficar isolado, recebendo apenas `texto` e `classe`.

### Ferramentas para acompanhar
- `npm run lint` (ESLint + TypeScript) garante que os componentes sigam convenções e testam tipos básicos.
- `npm run test` já está pronto no template Vite e pode ser usado para executar testes de snapshot ou simples renderizações de componentes.
- `npm run build` mostra se a árvore de componentes consegue ser empacotada sem erros antes de adicionarmos hooks.

### Dicas rápidas
- Leia o output do ESLint e corrija avisos antes de continuar.
- Use `npm run test -- --watch` quando quiser validar um componente isoladamente.
- Pergunte se um novo bloco de UI merece ter props próprias ou pode ser um componente já existente.

## Boas práticas iniciais
- Nomeie arquivos e componentes de forma semântica (`BannerHero.tsx`, `SessaoPlanos.tsx`) para facilitar a leitura.
- Separe estilos (CSS Modules ou arquivos `.scss`) para que cada componente cuide apenas da estrutura visual que precisa.
- Organize pastas por domínio (ex: `src/components/planos/Plano.tsx`) para tornar o reuso óbvio.
- Composição é preferível a herança: crie componentes que recebem outros como children ou props tipo `header`, `footer`, `card`.

## O que evitar por ora
- Não faça chamadas de rede nem efeitos colaterais dentro do corpo do componente; isso fica para aulas futuras.
- Evite mutações diretas em objetos ou arrays recebidos via props.
- Não misture responsabilidades: um componente não deve both renderizar um layout e decidir a navegação.

> Observação: ainda não estudamos hooks como `useMemo` ou `useCallback`; o foco hoje é dominar a estrutura dos componentes e as ferramentas CLI. Esses conceitos de otimização virão depois que você sentir confiança nesta base.

