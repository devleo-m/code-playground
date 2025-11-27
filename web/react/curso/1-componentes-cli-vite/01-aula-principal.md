# Aula 1: CLI Tools, Vite e Componentes Funcionais

## 1. CLI Tools como ponto de partida
Ferramentas de linha de comando (CLI) são o atalho mais confiável para criar um projeto React padronizado. Usamos `npm`, `yarn` ou `pnpm` para instalar dependências, rodar bundlers e testar rapidamente. Em especial, o Vite nos dá a experiência mais enxuta hoje:

```bash
npx create-react-app minha-app --template typescript # alternativa clássica, mais verbosa
npm create vite@latest minha-app -- --template react-ts # Vite gera um template atualizado
cd minha-app
npm install
npm run dev
```

Explicando o que cada etapa entrega:
- `create vite@latest` gera `package.json`, `tsconfig.json`, `vite.config.ts`, `src/main.tsx`, `src/App.tsx`, `src/index.css` e já configura vitest + ESLint.
- `npm install` baixa dependências como React, React DOM e Vite.
- `npm run dev` abre o servidor que recarrega apenas o que mudou — perfeito para testar componentes pequenos.

Esses comandos eliminam Babel, Webpack e loaders manuais, deixando espaço para focar em componentes. Temos scripts (`npm run build`, `npm run lint`, `npm run test`) que garantem qualidade desde o início.

## 2. Vite e a experiência enxuta
Vite é a ferramenta recomendada para iniciarmos hoje: roda um servidor de desenvolvimento instantâneo, recompila só o que mudou e privilegia módulos ES nativos. Esse setup oferece:

- estrutura moderna: `src/`, `assets/`, `tsconfig.json`, `vite.config.ts`, testes com Vitest e regras com ESLint/Prettier.
- suporte nativo a TypeScript e JSX: não precisamos alterar `tsconfig` ou instalar presets adicionais.
- servidor com Fast Refresh: só recarrega o componente alterado, o que deixa o processo de iterar mais rápido.

Depois de gerar o projeto, abra `src/main.tsx`. Ele injeta `App` em `document.getElementById('root')`. Esse é um marco importante: agora você sabe que todo componente parte de `main.tsx` → `App.tsx` → componentes filhos.

Vite também coloca comandos extras no `package.json` (`npm run check`, `npm run preview`) que serão úteis conforme adicionamos testes e publicação.
Esses recursos deixam mais claro que componentes React são unidades pequenas e reutilizáveis, não templates gigantescos.

## 3. O que é um componente funcional?
Um componente funcional é uma função TypeScript que recebe dados (props) e devolve JSX. Não depende de classes nem de `this`, apenas de parâmetros claros:

```tsx
type BannerProps = {
  titulo: string;
  subtitulo?: string;
};

const BannerHero = ({ titulo, subtitulo }: BannerProps) => (
  <header className="banner-hero">
    <h1>{titulo}</h1>
    {subtitulo && <p>{subtitulo}</p>}
  </header>
);

export default BannerHero;
```

Note que o retorno é JSX puro e a função não muda seus argumentos: para cada entrada (`titulo`, `subtitulo`) vem a mesma saída visual. Por enquanto, **não vamos explicar como o React renderiza** (isso vem na próxima aula) e **não entraremos em hooks**. O foco é o contrato props → JSX e como separar responsabilidades.

### Exemplo completo de pasta
Depois de criar esse `BannerHero`, a pasta pode ficar assim:

```
src/
  components/
    BannerHero/
      BannerHero.tsx
      BannerHero.module.css
  App.tsx
  main.tsx
```

Isso mostra que cada componente vive em sua própria pasta com estilos associados. No futuro, você pode adicionar testes e documentos ali mesmo.

## 4. Props vs estado (visão conceitual)
Props são os argumentos que o pai envia ao filho, como `titulo`. O estado (state) existe dentro do próprio componente e representa dados que ele controla sozinho. Nesta aula nós apenas reconhecemos essa distinção: props entram de fora; estado mora dentro. Mais adiante veremos como React usa hooks para manter esse estado sem mexer diretamente na árvore de componentes.

Essa clareza evita bugs: enquanto props são sempre leitura, o state pode mudar com interações. Hoje vamos focar na leitura e garantir que cada componente receba os ingredientes certos.

## 5. Composição vs herança
React recomenda composição. Em vez de estender classes, combinamos componentes pequenos:

```tsx
const CardConteudo = () => (
  <section>
    <BannerHero titulo="Boas-vindas" subtitulo="Componentes compostos" />
    <p>Essa seção mistura blocos reutilizáveis.</p>
  </section>
);
```

Assim, `CardConteudo` reutiliza `BannerHero` (composição) sem acoplar herança. Isso aumenta reuso, facilita testes e reduz responsabilidades.

## 6. Checklist de efeitos colaterais e re-renders
- Evite lógica pesada e acesso direto ao DOM dentro do corpo do componente.
- Não modifique props recebidos; trate-os como leitura imutável.
- Separe funções auxiliares fora do JSX quando possível, assim o React não cria novas referências a cada renderização.
- Prefira componentes pequenos: cada um deve ter uma responsabilidade única.
- Nomeie props de forma descritiva (`nomePlano` em vez de `n`).
- Use variáveis locais simples para dados transitórios e deixe os dados persistentes (state) para aulas futuras.
- Documente rapidamente o que cada componente entrega (ex: `BannerHero.tsx | mostra título e subtítulo hero`).

Seguindo esse checklist desde o início mantemos o componente previsível e pronto para evoluir quando adicionarmos hooks e mecanismos de renderização.

## 7. Dicas imediatas
1. Abra o console do Vite para ver mensagens de lint enquanto digita; o ESLint já alerta erros de TS e estilos de código.
2. Crie componentes dentro de `src/components/[nome]` e prefira arquivos `.module.css` ou `.scss` para evitar vazamento de estilos.
3. Sempre rode `npm run lint` antes de enviar alterações, mesmo que ainda esteja criando componentes simples.
4. Use `Ctrl/Cmd + P` no VS Code para abrir arquivos gerados pelo CLI e entenda cada camada (main → App → components).
5. Guarde o ciclo principal: CLI gera, Vite serve, componente rende → repetimos o loop conforme criamos novos blocos.

