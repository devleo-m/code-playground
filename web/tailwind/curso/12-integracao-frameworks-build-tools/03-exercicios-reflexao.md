# Aula 12 - Exercícios e Reflexão: Integração com Frameworks e Build Tools

## 🎯 Exercícios Práticos

### Exercício 1: Criar Projeto React com Tailwind

**Objetivo**: Criar um projeto React do zero e configurar Tailwind CSS.

**Tarefas**:
1. Crie um novo projeto React usando Vite:
   ```bash
   npm create vite@latest meu-projeto-tailwind -- --template react
   cd meu-projeto-tailwind
   npm install
   ```

2. Instale e configure Tailwind CSS:
   ```bash
   npm install -D tailwindcss postcss autoprefixer
   npx tailwindcss init -p
   ```

3. Configure o `tailwind.config.js` para incluir seus arquivos:
   ```javascript
   content: [
     "./index.html",
     "./src/**/*.{js,ts,jsx,tsx}",
   ]
   ```

4. Adicione as diretivas Tailwind em `src/index.css`:
   ```css
   @tailwind base;
   @tailwind components;
   @tailwind utilities;
   ```

5. Importe o CSS em `src/main.jsx`:
   ```javascript
   import './index.css'
   ```

6. Crie um componente simples que use classes Tailwind:
   ```jsx
   // src/App.jsx
   function App() {
     return (
       <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
         <div className="bg-white p-8 rounded-xl shadow-2xl max-w-md">
           <h1 className="text-4xl font-bold text-gray-800 mb-4">
             Tailwind + React
           </h1>
           <p className="text-gray-600 mb-6">
             Projeto configurado com sucesso!
           </p>
           <button className="bg-blue-500 hover:bg-blue-600 text-white px-6 py-3 rounded-lg font-semibold transition-colors">
             Começar
           </button>
         </div>
       </div>
     )
   }
   
   export default App
   ```

7. Execute o projeto e verifique se está funcionando:
   ```bash
   npm run dev
   ```

**Critérios de Sucesso**:
- ✅ Projeto React criado e funcionando
- ✅ Tailwind CSS instalado e configurado
- ✅ Componente renderiza com estilos Tailwind aplicados
- ✅ Hot reload funciona (mudanças aparecem automaticamente)

---

### Exercício 2: Criar Componente Reutilizável

**Objetivo**: Criar um componente de Card reutilizável usando Tailwind.

**Tarefas**:
1. Crie um arquivo `src/components/Card.jsx`:
   ```jsx
   // src/components/Card.jsx
   export default function Card({ title, description, image, children }) {
     return (
       <div className="bg-white rounded-lg shadow-lg overflow-hidden max-w-sm">
         {image && (
           <img 
             src={image} 
             alt={title}
             className="w-full h-48 object-cover"
           />
         )}
         <div className="p-6">
           <h3 className="text-xl font-bold text-gray-800 mb-2">
             {title}
           </h3>
           {description && (
             <p className="text-gray-600 mb-4">
               {description}
             </p>
           )}
           {children}
         </div>
       </div>
     )
   }
   ```

2. Use o componente Card em `App.jsx`:
   ```jsx
   // src/App.jsx
   import Card from './components/Card'
   
   function App() {
     return (
       <div className="min-h-screen bg-gray-100 py-12 px-4">
         <div className="max-w-6xl mx-auto">
           <h1 className="text-4xl font-bold text-center mb-12 text-gray-800">
             Meus Cards
           </h1>
           
           <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
             <Card 
               title="Card 1"
               description="Descrição do primeiro card"
               image="https://via.placeholder.com/400x300"
             >
               <button className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
                 Ver Mais
               </button>
             </Card>
             
             <Card 
               title="Card 2"
               description="Descrição do segundo card"
             >
               <p className="text-sm text-gray-500">Conteúdo customizado</p>
             </Card>
             
             <Card title="Card 3">
               <div className="space-y-2">
                 <p className="text-gray-600">Item 1</p>
                 <p className="text-gray-600">Item 2</p>
                 <p className="text-gray-600">Item 3</p>
               </div>
             </Card>
           </div>
         </div>
       </div>
     )
   }
   
   export default App
   ```

**Critérios de Sucesso**:
- ✅ Componente Card criado e reutilizável
- ✅ Aceita props diferentes (title, description, image, children)
- ✅ Layout responsivo (grid adapta a diferentes tamanhos de tela)
- ✅ Estilos Tailwind aplicados corretamente

---

### Exercício 3: Configurar Next.js com Tailwind

**Objetivo**: Criar um projeto Next.js e configurar Tailwind CSS.

**Tarefas**:
1. Crie um novo projeto Next.js:
   ```bash
   npx create-next-app@latest meu-nextjs-tailwind
   cd meu-nextjs-tailwind
   ```
   
   Durante a criação, escolha:
   - TypeScript? **No** (ou Yes, se preferir)
   - ESLint? **Yes**
   - Tailwind CSS? **Yes** (se disponível, senão instale depois)
   - App Router? **Yes** (recomendado) ou **No** (Pages Router)

2. Se Tailwind não foi instalado automaticamente:
   ```bash
   npm install -D tailwindcss postcss autoprefixer
   npx tailwindcss init -p
   ```

3. Verifique se `tailwind.config.js` está configurado corretamente:
   ```javascript
   // tailwind.config.js
   module.exports = {
     content: [
       './pages/**/*.{js,ts,jsx,tsx}',
       './components/**/*.{js,ts,jsx,tsx}',
       './app/**/*.{js,ts,jsx,tsx}',
     ],
     theme: {
       extend: {},
     },
     plugins: [],
   }
   ```

4. Crie uma página inicial estilizada:
   
   **Se usar App Router** (`app/page.js`):
   ```jsx
   // app/page.js
   export default function Home() {
     return (
       <div className="min-h-screen bg-gradient-to-br from-green-400 to-blue-500">
         <div className="container mx-auto px-4 py-16">
           <div className="text-center">
             <h1 className="text-6xl font-bold text-white mb-6">
               Next.js + Tailwind
             </h1>
             <p className="text-xl text-white/90 mb-8">
               Projeto configurado com sucesso!
             </p>
             <div className="flex gap-4 justify-center">
               <button className="bg-white text-blue-500 px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition-colors">
                 Começar
               </button>
               <button className="bg-blue-600 text-white px-8 py-3 rounded-lg font-semibold hover:bg-blue-700 transition-colors">
                 Documentação
               </button>
             </div>
           </div>
         </div>
       </div>
     )
   }
   ```
   
   **Se usar Pages Router** (`pages/index.js`):
   ```jsx
   // pages/index.js
   export default function Home() {
     return (
       <div className="min-h-screen bg-gradient-to-br from-green-400 to-blue-500">
         <div className="container mx-auto px-4 py-16">
           <div className="text-center">
             <h1 className="text-6xl font-bold text-white mb-6">
               Next.js + Tailwind
             </h1>
             <p className="text-xl text-white/90 mb-8">
               Projeto configurado com sucesso!
             </p>
             <div className="flex gap-4 justify-center">
               <button className="bg-white text-blue-500 px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition-colors">
                 Começar
               </button>
               <button className="bg-blue-600 text-white px-8 py-3 rounded-lg font-semibold hover:bg-blue-700 transition-colors">
                 Documentação
               </button>
             </div>
           </div>
         </div>
       </div>
     )
   }
   ```

5. Execute o projeto:
   ```bash
   npm run dev
   ```

**Critérios de Sucesso**:
- ✅ Projeto Next.js criado
- ✅ Tailwind CSS configurado
- ✅ Página inicial renderiza com estilos
- ✅ Hot reload funciona

---

### Exercício 4: Resolver Problema de Classes Dinâmicas

**Objetivo**: Entender e resolver o problema de classes Tailwind geradas dinamicamente.

**Situação**: Você tem um componente que muda de cor baseado em uma prop, mas as classes não funcionam.

**Código Problemático**:
```jsx
// ❌ Isso NÃO funciona
function Botao({ cor }) {
  return (
    <button className={`bg-${cor}-500 text-white px-4 py-2 rounded`}>
      Botão
    </button>
  )
}

// Uso
<Botao cor="blue" />  // Não funciona!
```

**Tarefas**:
1. Identifique o problema: Por que isso não funciona?

2. Crie uma solução usando objeto de mapeamento:
   ```jsx
   // ✅ Solução correta
   function Botao({ cor = 'blue', children }) {
     const cores = {
       blue: 'bg-blue-500 hover:bg-blue-600',
       red: 'bg-red-500 hover:bg-red-600',
       green: 'bg-green-500 hover:bg-green-600',
       yellow: 'bg-yellow-500 hover:bg-yellow-600',
     }
     
     return (
       <button className={`${cores[cor]} text-white px-4 py-2 rounded transition-colors`}>
         {children}
       </button>
     )
   }
   ```

3. Crie uma versão mais robusta com validação:
   ```jsx
   function Botao({ cor = 'blue', children, tamanho = 'medio' }) {
     const cores = {
       blue: 'bg-blue-500 hover:bg-blue-600',
       red: 'bg-red-500 hover:bg-red-600',
       green: 'bg-green-500 hover:bg-green-600',
     }
     
     const tamanhos = {
       pequeno: 'px-3 py-1.5 text-sm',
       medio: 'px-4 py-2 text-base',
       grande: 'px-6 py-3 text-lg',
     }
     
     // Validação
     const corClasses = cores[cor] || cores.blue
     const tamanhoClasses = tamanhos[tamanho] || tamanhos.medio
     
     return (
       <button className={`${corClasses} ${tamanhoClasses} text-white rounded transition-colors`}>
         {children}
       </button>
     )
   }
   ```

4. Use o componente em diferentes variações:
   ```jsx
   <div className="space-x-4">
     <Botao cor="blue" tamanho="pequeno">Pequeno Azul</Botao>
     <Botao cor="red" tamanho="medio">Médio Vermelho</Botao>
     <Botao cor="green" tamanho="grande">Grande Verde</Botao>
   </div>
   ```

**Critérios de Sucesso**:
- ✅ Entendeu por que classes dinâmicas não funcionam
- ✅ Implementou solução com objeto de mapeamento
- ✅ Adicionou validação para valores inválidos
- ✅ Componente funciona com diferentes props

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Por que Classes Dinâmicas Não Funcionam?

**Pergunta**: Por que o código `className={`bg-${cor}-500`}` não funciona com Tailwind, mesmo que a string final seja uma classe válida como `bg-blue-500`?

**Pense sobre**:
- Como o Tailwind processa classes durante o build
- A diferença entre classes estáticas e dinâmicas
- Por que o PurgeCSS/JIT precisa "ver" as classes completas
- Qual seria o impacto na performance se Tailwind tentasse detectar todas as combinações possíveis

**Resposta Esperada** (guia de pensamento):
- Tailwind escaneia arquivos em busca de classes completas
- Classes geradas dinamicamente não são "visíveis" durante o scan
- Se Tailwind tentasse gerar todas as combinações possíveis, o CSS seria enorme
- A solução é usar mapeamento explícito ou safelist

---

### Reflexão 2: Quando Usar React vs Next.js?

**Pergunta**: Em que situações você escolheria React puro (Create React App ou Vite) vs Next.js? Quais são as vantagens e desvantagens de cada abordagem?

**Pense sobre**:
- Diferenças entre Single Page Application (SPA) e Server-Side Rendering (SSR)
- Quando SEO é importante
- Necessidade de rotas e navegação
- Performance e tempo de carregamento inicial
- Complexidade do projeto

**Resposta Esperada** (guia de pensamento):
- **React puro**: Melhor para apps internos, dashboards, aplicações que não precisam de SEO
- **Next.js**: Melhor para sites públicos, blogs, e-commerce, quando SEO é crucial
- **React**: Mais simples, mais controle, melhor para apps interativos
- **Next.js**: Mais recursos (SSR, SSG, rotas), melhor para sites estáticos/dinâmicos

---

### Reflexão 3: Impacto de Build Tools na Performance

**Pergunta**: Como a escolha do build tool (Webpack vs Vite vs Parcel) impacta a experiência de desenvolvimento e a performance final do projeto?

**Pense sobre**:
- Tempo de inicialização do servidor de desenvolvimento
- Velocidade do Hot Module Replacement (HMR)
- Tempo de build para produção
- Tamanho do bundle final
- Complexidade de configuração

**Resposta Esperada** (guia de pensamento):
- **Webpack**: Mais lento, mas maduro e estável, muitos plugins
- **Vite**: Muito mais rápido em desenvolvimento, build rápido, configuração simples
- **Parcel**: Zero config, funciona automaticamente, mas menos controle
- A escolha depende do tamanho do projeto, equipe, e necessidades específicas

---

### Reflexão 4: Organização de Componentes com Tailwind

**Pergunta**: Qual é a melhor forma de organizar componentes React que usam Tailwind? Quando você criaria um componente customizado vs quando usaria apenas classes utilitárias diretamente?

**Pense sobre**:
- Reutilização de código
- Manutenibilidade
- Legibilidade
- Performance (tamanho do bundle)
- Trabalho em equipe

**Resposta Esperada** (guia de pensamento):
- **Criar componente**: Quando há lógica reutilizável, padrões que se repetem, ou quando precisa de props dinâmicas
- **Usar classes diretamente**: Quando é uso único, layout específico, ou quando a abstração não traz benefício
- **Regra de 3**: Se você usa o mesmo padrão 3+ vezes, considere criar componente
- **Balance**: Não crie componentes demais (over-engineering), mas também não repita código

---

### Reflexão 5: PostCSS e o Processo de Build

**Pergunta**: Por que o Tailwind precisa do PostCSS? O que aconteceria se tentássemos usar Tailwind sem PostCSS? Qual é o papel de cada ferramenta no processo?

**Pense sobre**:
- O que PostCSS faz
- Como Tailwind gera CSS
- O papel do Autoprefixer
- O fluxo completo: código → CSS final

**Resposta Esperada** (guia de pensamento):
- **PostCSS**: Processador que transforma CSS usando plugins
- **Tailwind plugin**: Gera classes utilitárias a partir das diretivas @tailwind
- **Autoprefixer**: Adiciona vendor prefixes automaticamente
- **Sem PostCSS**: Tailwind não conseguiria processar as diretivas @tailwind
- **Fluxo**: CSS com @tailwind → PostCSS processa → Tailwind gera classes → Autoprefixer adiciona prefixes → CSS final

---

## 📝 Exercício de Análise de Código

### Análise: Componente com Problemas

Analise o seguinte componente e identifique problemas:

```jsx
// Componente com problemas
function ProdutoCard({ produto }) {
  const cor = produto.categoria === 'eletronicos' ? 'blue' : 'green'
  
  return (
    <div className={`bg-${cor}-100 border-${cor}-500 border-2 p-4 rounded`}>
      <h3 className="text-xl font-bold">{produto.nome}</h3>
      <p className="text-gray-600">{produto.descricao}</p>
      <div className="flex justify-between items-center mt-4">
        <span className={`text-${cor}-700 font-bold`}>
          R$ {produto.preco}
        </span>
        <button className={`bg-${cor}-500 hover:bg-${cor}-600 text-white px-4 py-2 rounded`}>
          Comprar
        </button>
      </div>
    </div>
  )
}
```

**Tarefas**:
1. Identifique todos os problemas neste código
2. Explique por que cada problema ocorre
3. Reescreva o componente de forma correta
4. Adicione melhorias (validação, acessibilidade, etc.)

**Problemas Identificados**:
- ❌ Classes dinâmicas não funcionam (`bg-${cor}-100`, etc.)
- ❌ Falta validação de props
- ❌ Falta tratamento de valores undefined/null
- ❌ Pode melhorar acessibilidade (alt text, aria labels)

**Solução Correta**:
```jsx
function ProdutoCard({ produto }) {
  // Validação
  if (!produto) return null
  
  // Mapeamento de cores
  const estilosPorCategoria = {
    eletronicos: {
      bg: 'bg-blue-100',
      border: 'border-blue-500',
      text: 'text-blue-700',
      button: 'bg-blue-500 hover:bg-blue-600',
    },
    default: {
      bg: 'bg-green-100',
      border: 'border-green-500',
      text: 'text-green-700',
      button: 'bg-green-500 hover:bg-green-600',
    },
  }
  
  const estilos = estilosPorCategoria[produto.categoria] || estilosPorCategoria.default
  
  return (
    <div className={`${estilos.bg} ${estilos.border} border-2 p-4 rounded-lg shadow-md`}>
      <h3 className="text-xl font-bold text-gray-800 mb-2">
        {produto.nome || 'Produto sem nome'}
      </h3>
      <p className="text-gray-600 mb-4">
        {produto.descricao || 'Sem descrição'}
      </p>
      <div className="flex justify-between items-center mt-4">
        <span className={`${estilos.text} font-bold text-lg`}>
          R$ {produto.preco?.toFixed(2) || '0.00'}
        </span>
        <button 
          className={`${estilos.button} text-white px-4 py-2 rounded transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2`}
          aria-label={`Comprar ${produto.nome}`}
        >
          Comprar
        </button>
      </div>
    </div>
  )
}
```

---

## ✅ Checklist de Aprendizado

Marque o que você conseguiu fazer:

- [ ] Criei um projeto React do zero e configurei Tailwind
- [ ] Criei componentes reutilizáveis com Tailwind
- [ ] Configurei Next.js com Tailwind
- [ ] Resolvi problemas de classes dinâmicas
- [ ] Entendi o papel do PostCSS
- [ ] Compreendi diferenças entre build tools
- [ ] Sei quando usar React vs Next.js
- [ ] Consigo debugar problemas de integração
- [ ] Organizo projetos de forma escalável
- [ ] Entendo o processo de build completo

---

**Bons exercícios! Pratique bastante e reflita sobre as decisões de arquitetura! 🚀**

