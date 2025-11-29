# Estrutura do Projeto - React Fundamentos

Este projeto foi desenvolvido para ensinar React de forma prática e organizada, seguindo as melhores práticas de organização de código React.

## 📁 Estrutura de Pastas

```
src/
├── components/          # Componentes reutilizáveis
│   ├── ui/             # Componentes de interface reutilizáveis
│   │   ├── CodeExample/
│   │   │   ├── CodeExample.jsx
│   │   │   ├── CodeExample.css
│   │   │   └── index.js
│   │   └── index.js    # Exports centralizados
│   │
│   └── layout/         # Componentes de layout
│       ├── Layout/
│       │   ├── Layout.jsx
│       │   ├── Layout.css
│       │   └── index.js
│       ├── Menu/
│       │   ├── Menu.jsx
│       │   ├── Menu.css
│       │   └── index.js
│       ├── Footer/
│       │   ├── Footer.jsx
│       │   ├── Footer.css
│       │   └── index.js
│       └── index.js    # Exports centralizados
│
├── pages/               # Páginas da aplicação
│   ├── Home/
│   │   ├── Home.jsx
│   │   └── Home.css
│   ├── Aulas/
│   │   ├── Aulas.jsx
│   │   └── Aulas.css
│   └── Aula/
│       ├── Aula.jsx
│       └── Aula.css
│
├── hooks/              # Custom hooks (quando necessário)
├── utils/              # Funções utilitárias
├── constants/          # Dados estáticos e constantes
│   └── aulas.js       # Lista de aulas
│
├── App.jsx             # Componente principal com rotas
└── main.jsx            # Ponto de entrada da aplicação
```

## 🎯 Padrões de Organização

### Colocation (Arquivos Relacionados Juntos)
Cada componente tem sua própria pasta com todos os arquivos relacionados:
- Componente JSX
- Estilos CSS
- Arquivo `index.js` para export limpo

### Exports Centralizados
Cada pasta de componentes tem um `index.js` que exporta todos os componentes:
```javascript
// components/ui/index.js
export { default as CodeExample } from './CodeExample'

// components/layout/index.js
export { default as Layout } from './Layout'
export { default as Menu } from './Menu'
export { default as Footer } from './Footer'
```

### Imports Limpos
Use os exports centralizados para imports mais limpos:
```javascript
// ✅ Bom
import { Layout } from './components/layout'
import { CodeExample } from './components/ui'

// ❌ Evitar
import Layout from './components/layout/Layout/Layout'
```

## 🎯 Como Adicionar uma Nova Aula

### 1. Criar o componente da aula (se necessário)

Se a aula tiver componentes práticos específicos, crie em `src/components/aulas/`:

```jsx
// src/components/aulas/Aula1Exemplo/Aula1Exemplo.jsx
function Aula1Exemplo() {
  return <div>Meu exemplo prático</div>
}

export default Aula1Exemplo
```

### 2. Adicionar dados da aula em `constants/aulas.js`

```javascript
// src/constants/aulas.js
export const AULAS = [
  {
    id: '1',
    titulo: 'Aula 1: Introdução ao React',
    descricao: 'Aprenda os conceitos básicos do React',
    teoria: `
      <h3>O que é React?</h3>
      <p>React é uma biblioteca JavaScript para construir interfaces de usuário...</p>
    `,
    exemplos: [
      {
        title: 'Meu Primeiro Componente',
        description: 'Este é um exemplo básico de componente React',
        code: `
          function Exemplo() {
            return <h1>Olá!</h1>
          }
        `,
        ExampleComponent: Aula1Exemplo // Componente opcional
      }
    ]
  }
]
```

### 3. Atualizar a página Aula.jsx

A página `Aula.jsx` buscará automaticamente os dados de `constants/aulas.js`:

```jsx
import { useParams } from 'react-router-dom'
import { CodeExample } from '../components/ui'
import { AULAS } from '../constants/aulas'
import './Aula.css'

function Aula() {
  const { aulaId } = useParams()
  const aula = AULAS.find(a => a.id === aulaId)
  
  if (!aula) {
    return <div>Aula não encontrada</div>
  }
  
  return (
    <div className="aula-page">
      <div className="aula-header">
        <h1>{aula.titulo}</h1>
      </div>
      
      <div className="aula-content">
        <div className="aula-section">
          <h2>📖 Teoria</h2>
          <div 
            className="aula-theory"
            dangerouslySetInnerHTML={{ __html: aula.teoria }}
          />
        </div>
        
        <div className="aula-section">
          <h2>💻 Prática</h2>
          {aula.exemplos?.map((exemplo, index) => (
            <CodeExample
              key={index}
              title={exemplo.title}
              description={exemplo.description}
              code={exemplo.code}
              ExampleComponent={exemplo.ExampleComponent}
            />
          ))}
        </div>
      </div>
    </div>
  )
}
```

### 4. A lista de aulas será atualizada automaticamente

A página `Aulas.jsx` já importa de `constants/aulas.js`, então as aulas aparecerão automaticamente na lista.

## 🎨 Componentes Disponíveis

### CodeExample

Componente para exibir exemplos práticos com código e demonstração:

```jsx
import { CodeExample } from '../components/ui'

<CodeExample
  title="Título do Exemplo"
  description="Descrição do exemplo"
  code={`código aqui`}
  ExampleComponent={MeuComponente}
/>
```

**Props:**
- `title` (string): Título do exemplo
- `description` (string, opcional): Descrição do exemplo
- `code` (string, opcional): Código a ser exibido
- `ExampleComponent` (component, opcional): Componente React para demonstrar

## 🚀 Executando o Projeto

```bash
npm run dev    # Desenvolvimento
npm run build  # Build para produção
npm run preview # Preview do build
npm run lint   # Verificar erros de lint
```

## 📝 Boas Práticas

1. **Colocation**: Cada componente tem sua própria pasta com todos os arquivos relacionados
2. **Separação de Responsabilidades**: Componentes UI em `ui/`, layout em `layout/`
3. **Exports Centralizados**: Use `index.js` para exports limpos
4. **Dados Estáticos**: Mantenha dados estáticos em `constants/`
5. **CSS Modular**: Cada componente tem seu próprio arquivo CSS
6. **Hooks Customizados**: Coloque em `hooks/` quando necessário
7. **Funções Utilitárias**: Coloque em `utils/` quando necessário
8. **JavaScript Puro**: Sem TypeScript, código limpo e simples

## 🎓 Estrutura de uma Aula

Cada aula deve ter:
- **Teoria**: Explicação conceitual do tema (HTML formatado)
- **Prática**: Exemplos de código funcionando
- **Visualização**: Componentes React em ação

Isso permite que o estudante:
1. Entenda o conceito (teoria)
2. Veja como implementar (código)
3. Veja funcionando (componente em ação)

## 🔄 Vantagens desta Estrutura

- **Escalável**: Fácil adicionar novos componentes e features
- **Organizada**: Fácil encontrar código relacionado
- **Manutenível**: Mudanças isoladas em componentes específicos
- **Reutilizável**: Componentes bem separados e exportados
- **Profissional**: Segue padrões da indústria
