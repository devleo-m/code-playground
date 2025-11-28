# 📚 Guia de Estudos React

Um guia interativo para aprender os fundamentos do React de forma prática e organizada.

## 🎯 Como Funciona

1. Escolha um tópico na lista de estudos
2. Clique no card do tópico que deseja aprender
3. O conteúdo aparecerá abaixo para você estudar

## 📁 Estrutura do Projeto

```
src/
├── components/          # Componentes reutilizáveis
│   ├── Header.jsx      # Cabeçalho da aplicação
│   ├── TopicList.jsx   # Lista de tópicos para escolher
│   └── ContentViewer.jsx # Visualizador de conteúdo
├── data/               # Dados da aplicação
│   └── topics.js      # Array com todos os tópicos de estudo
├── App.jsx            # Componente principal
└── main.jsx           # Ponto de entrada
```

## 🎨 Organização

- **components/**: Todos os componentes React organizados por funcionalidade
- **data/**: Dados estáticos da aplicação (tópicos, conteúdo, etc)
- Cada componente tem seu próprio arquivo CSS para manter organizado

## 🚀 Como Adicionar Novos Tópicos

Edite o arquivo `src/data/topics.js` e adicione um novo objeto ao array:

```javascript
{
  id: 7,
  title: "Seu Tópico",
  description: "Descrição do tópico",
  content: `
    <h2>Seu Conteúdo</h2>
    <p>Conteúdo em HTML aqui...</p>
  `
}
```

## 💡 Conceitos Usados

- **useState**: Para gerenciar qual tópico está selecionado
- **Props**: Para passar dados entre componentes
- **Componentes Funcionais**: Forma moderna de criar componentes
- **Eventos**: onClick para selecionar tópicos
