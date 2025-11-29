# Estrutura do Projeto - React Fundamentos

Este projeto foi desenvolvido para ensinar React de forma prática e organizada, com componentes e páginas bem separadas para facilitar o aprendizado e manutenção.

## 📁 Estrutura de Pastas

```
src/
├── components/          # Componentes reutilizáveis
│   ├── interface/       # Componentes de interface (Menu, Footer)
│   │   ├── Menu.jsx     # Menu de navegação
│   │   └── Footer.jsx   # Rodapé
│   ├── Layout.jsx       # Layout principal (Menu + Conteúdo + Footer)
│   └── CodeExample.jsx  # Componente para exibir exemplos de código
│
├── pages/               # Páginas da aplicação
│   ├── Home.jsx         # Página inicial
│   ├── Aulas.jsx        # Lista de aulas
│   └── Aula.jsx         # Página individual de cada aula
│
├── App.jsx              # Componente principal com rotas
└── main.jsx             # Ponto de entrada da aplicação
```

## 🎯 Como Adicionar uma Nova Aula

### 1. Criar o componente da aula (se necessário)

Se a aula tiver componentes práticos específicos, crie em `src/components/aulas/`:

```jsx
// src/components/aulas/Aula1Exemplo.jsx
function Aula1Exemplo() {
  return <div>Meu exemplo prático</div>
}

export default Aula1Exemplo
```

### 2. Atualizar a página Aula.jsx

Adicione o conteúdo da aula na página `Aula.jsx`:

```jsx
import CodeExample from '../components/CodeExample'
import Aula1Exemplo from '../components/aulas/Aula1Exemplo'

function Aula() {
  const { aulaId } = useParams()
  
  // Conteúdo específico para cada aula
  if (aulaId === '1') {
    const teoria = `
      <h3>Título da Teoria</h3>
      <p>Explicação teórica aqui...</p>
    `
    
    const codigoExemplo = `
      function Exemplo() {
        return <h1>Olá!</h1>
      }
    `
    
    return (
      <div className="aula-page">
        <div className="aula-header">
          <h1>Aula 1: Introdução ao React</h1>
        </div>
        
        <div className="aula-content">
          <div className="aula-section">
            <h2>📖 Teoria</h2>
            <div 
              className="aula-theory"
              dangerouslySetInnerHTML={{ __html: teoria }}
            />
          </div>
          
          <div className="aula-section">
            <h2>💻 Prática</h2>
            <CodeExample
              title="Meu Primeiro Componente"
              description="Este é um exemplo básico de componente React"
              code={codigoExemplo}
              ExampleComponent={Aula1Exemplo}
            />
          </div>
        </div>
      </div>
    )
  }
  
  // ... outras aulas
}
```

### 3. Adicionar na lista de aulas

Atualize `Aulas.jsx` para incluir a nova aula:

```jsx
const aulas = [
  {
    id: '1',
    titulo: 'Aula 1: Introdução ao React',
    descricao: 'Aprenda os conceitos básicos do React'
  },
  // ... outras aulas
]
```

## 🎨 Componentes Disponíveis

### CodeExample

Componente para exibir exemplos práticos com código e demonstração:

```jsx
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
```

## 📝 Boas Práticas

1. **Separação de Responsabilidades**: Cada componente tem uma responsabilidade única
2. **Organização**: Componentes de interface em `interface/`, páginas em `pages/`
3. **Reutilização**: Use o componente `CodeExample` para todos os exemplos práticos
4. **CSS Modular**: Cada componente tem seu próprio arquivo CSS
5. **JavaScript Puro**: Sem TypeScript, sem gambiarras, código limpo e simples

## 🎓 Estrutura de uma Aula

Cada aula deve ter:
- **Teoria**: Explicação conceitual do tema
- **Prática**: Exemplos de código funcionando
- **Visualização**: Componentes React em ação

Isso permite que o estudante:
1. Entenda o conceito (teoria)
2. Veja como implementar (código)
3. Veja funcionando (componente em ação)

