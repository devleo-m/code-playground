# Aula 9 - Exercícios e Reflexão: Plugins e Extensões do Tailwind

## 🎯 Objetivos dos Exercícios

Ao completar estes exercícios, você será capaz de:
- Instalar e configurar plugins oficiais do Tailwind
- Usar o plugin Typography para estilizar conteúdo
- Aplicar o plugin Forms em formulários
- Utilizar Aspect Ratio para controlar proporções
- Implementar Line Clamp para truncar texto
- Criar plugins customizados simples
- Avaliar quando usar plugins vs CSS customizado
- Pensar criticamente sobre o impacto de plugins no bundle size

---

## 📝 Exercício 1: Criando um Blog com Typography

### Tarefa

Você precisa criar uma página de blog usando o plugin Typography. O artigo deve ter:
- Título principal estilizado
- Parágrafo de introdução (lead)
- Subtítulos (h2 e h3)
- Listas ordenadas e não ordenadas
- Links estilizados
- Citações (blockquote)
- Código inline e em blocos

### Requisitos

1. Instale o plugin `@tailwindcss/typography`
2. Configure no `tailwind.config.js`
3. Use a classe `prose` com modificadores:
   - Tamanho: `prose-lg`
   - Cor: `prose-blue`
   - Sem limite de largura: `max-w-none`
4. Crie um artigo completo com todos os elementos acima

### Código Base

```html
<!DOCTYPE html>
<html>
<head>
  <script src="https://cdn.tailwindcss.com"></script>
  <script>
    tailwind.config = {
      plugins: [
        // Configure o plugin aqui
      ],
    }
  </script>
</head>
<body class="bg-gray-100 p-8">
  <article class="prose prose-lg prose-blue mx-auto max-w-4xl">
    <!-- Seu conteúdo aqui -->
  </article>
</body>
</html>
```

### Exemplo de Conteúdo

```html
<h1>Como Aprender Tailwind CSS</h1>
<p class="lead">Este é um guia completo para dominar Tailwind CSS...</p>
<h2>Por que usar Tailwind?</h2>
<p>Tailwind oferece várias vantagens...</p>
<ul>
  <li>Produtividade</li>
  <li>Consistência</li>
</ul>
```

### Critérios de Avaliação

- ✅ Plugin instalado e configurado corretamente
- ✅ Todos os elementos tipográficos estilizados
- ✅ Modificadores aplicados corretamente
- ✅ Artigo visualmente atraente e legível

---

## 📝 Exercício 2: Formulário de Contato com Forms Plugin

### Tarefa

Crie um formulário de contato completo usando o plugin Forms. O formulário deve incluir:
- Campo de nome (text)
- Campo de email (email)
- Campo de telefone (tel)
- Textarea para mensagem
- Select para assunto
- Checkbox para aceitar termos
- Radio buttons para preferência de contato
- Botão de submit

### Requisitos

1. Instale o plugin `@tailwindcss/forms`
2. Configure no `tailwind.config.js`
3. Use a estratégia `base` (padrão)
4. Estilize o formulário com classes Tailwind adicionais
5. Adicione estados de hover e focus visíveis

### Código Base

```html
<!DOCTYPE html>
<html>
<head>
  <script src="https://cdn.tailwindcss.com"></script>
  <script>
    tailwind.config = {
      plugins: [
        // Configure o plugin aqui
      ],
    }
  </script>
</head>
<body class="bg-gray-50 p-8">
  <form class="max-w-md mx-auto bg-white p-6 rounded-lg shadow-md space-y-4">
    <!-- Seus campos aqui -->
  </form>
</body>
</html>
```

### Critérios de Avaliação

- ✅ Plugin instalado e configurado
- ✅ Todos os tipos de input estilizados
- ✅ Formulário visualmente consistente
- ✅ Estados de interação funcionando

---

## 📝 Exercício 3: Galeria de Imagens com Aspect Ratio

### Tarefa

Crie uma galeria de imagens responsiva usando o plugin Aspect Ratio. A galeria deve ter:
- Imagens em diferentes proporções (1:1, 16:9, 4:3)
- Layout em grid responsivo
- Imagens que mantêm proporção em qualquer tamanho de tela
- Hover effects

### Requisitos

1. Instale o plugin `@tailwindcss/aspect-ratio`
2. Configure no `tailwind.config.js`
3. Crie um grid com 3 colunas no desktop, 2 no tablet, 1 no mobile
4. Use diferentes proporções para criar visual interessante
5. Adicione efeitos de hover (scale, opacity)

### Código Base

```html
<!DOCTYPE html>
<html>
<head>
  <script src="https://cdn.tailwindcss.com"></script>
  <script>
    tailwind.config = {
      plugins: [
        // Configure o plugin aqui
      ],
    }
  </script>
</head>
<body class="bg-gray-100 p-8">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <!-- Suas imagens aqui -->
  </div>
</body>
</html>
```

### Exemplo de Estrutura

```html
<div class="aspect-w-1 aspect-h-1 overflow-hidden rounded-lg">
  <img src="imagem.jpg" alt="Imagem" class="object-cover w-full h-full hover:scale-110 transition-transform" />
</div>
```

### Critérios de Avaliação

- ✅ Plugin instalado e configurado
- ✅ Proporções mantidas em diferentes tamanhos
- ✅ Layout responsivo funcionando
- ✅ Efeitos de hover implementados

---

## 📝 Exercício 4: Cards de Produto com Line Clamp

### Tarefa

Crie cards de produto para um e-commerce usando o plugin Line Clamp. Cada card deve ter:
- Imagem do produto
- Título (truncado em 1 linha)
- Descrição (truncada em 3 linhas)
- Preço
- Botão de compra

### Requisitos

1. Instale o plugin `@tailwindcss/line-clamp`
2. Configure no `tailwind.config.js`
3. Crie pelo menos 4 cards
4. Use `line-clamp-1` para títulos
5. Use `line-clamp-3` para descrições
6. Layout em grid responsivo

### Código Base

```html
<!DOCTYPE html>
<html>
<head>
  <script src="https://cdn.tailwindcss.com"></script>
  <script>
    tailwind.config = {
      plugins: [
        // Configure o plugin aqui
      ],
    }
  </script>
</head>
<body class="bg-gray-100 p-8">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <!-- Seus cards aqui -->
  </div>
</body>
</html>
```

### Exemplo de Card

```html
<div class="bg-white rounded-lg shadow-md overflow-hidden">
  <img src="produto.jpg" alt="Produto" class="w-full h-48 object-cover" />
  <div class="p-4">
    <h3 class="line-clamp-1 font-bold text-lg mb-2">Nome do Produto Muito Longo Que Será Truncado</h3>
    <p class="line-clamp-3 text-gray-600 mb-4">Descrição muito longa do produto que será truncada após três linhas com ellipsis no final da terceira linha...</p>
    <div class="flex items-center justify-between">
      <span class="text-2xl font-bold text-blue-600">R$ 99,90</span>
      <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">Comprar</button>
    </div>
  </div>
</div>
```

### Critérios de Avaliação

- ✅ Plugin instalado e configurado
- ✅ Texto truncado corretamente
- ✅ Cards visualmente consistentes
- ✅ Layout responsivo

---

## 📝 Exercício 5: Criando um Plugin Customizado

### Tarefa

Crie seu próprio plugin que adiciona classes utilitárias para **text-shadow** (sombra de texto). O plugin deve incluir:
- `.text-shadow-sm` - sombra pequena
- `.text-shadow` - sombra padrão
- `.text-shadow-md` - sombra média
- `.text-shadow-lg` - sombra grande
- `.text-shadow-none` - sem sombra
- Variantes hover e focus

### Requisitos

1. Crie o arquivo `tailwindcss-text-shadow.js`
2. Use a API do plugin do Tailwind
3. Adicione as classes com `addUtilities`
4. Inclua variantes hover e focus
5. Configure no `tailwind.config.js`
6. Teste usando as classes em um HTML

### Código Base do Plugin

```javascript
// tailwindcss-text-shadow.js
const plugin = require('tailwindcss/plugin')

module.exports = plugin(function({ addUtilities, theme }) {
  // Seu código aqui
})
```

### Estrutura Esperada

```javascript
const textShadows = {
  'sm': '1px 1px 2px rgba(0, 0, 0, 0.1)',
  'DEFAULT': '2px 2px 4px rgba(0, 0, 0, 0.1)',
  'md': '4px 4px 8px rgba(0, 0, 0, 0.12)',
  'lg': '8px 8px 16px rgba(0, 0, 0, 0.15)',
  'none': 'none',
}
```

### Configuração

```javascript
// tailwind.config.js
module.exports = {
  content: ['./src/**/*.{html,js}'],
  plugins: [
    require('./tailwindcss-text-shadow'),
  ],
}
```

### Teste

```html
<h1 class="text-shadow-lg hover:text-shadow-xl">Título com Sombra</h1>
<p class="text-shadow-md focus:text-shadow-lg">Texto com sombra média</p>
```

### Critérios de Avaliação

- ✅ Plugin criado corretamente
- ✅ Todas as classes funcionando
- ✅ Variantes hover e focus implementadas
- ✅ Plugin configurado e testado

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Bundle Size e Performance

**Pergunta:** Você instalou 4 plugins oficiais (Typography, Forms, Aspect Ratio, Line Clamp) em um projeto. Como isso afeta o tamanho do CSS final? Quais estratégias você usaria para minimizar o impacto?

**Pense sobre:**
- Quantas classes cada plugin gera?
- Todas as classes são necessárias no seu projeto?
- Como o PurgeCSS/JIT ajuda?
- Quando vale a pena instalar um plugin completo vs criar CSS customizado?

**Resposta esperada:** Você deve considerar que cada plugin adiciona CSS ao bundle. Typography, por exemplo, gera muitas classes (prose, prose-sm, prose-lg, etc.). Se você só usa `prose`, muitas classes ficam não utilizadas. O PurgeCSS remove classes não usadas, mas ainda assim, plugins grandes podem aumentar o bundle. Estratégias: usar apenas o necessário, considerar CSS customizado para casos simples, monitorar o tamanho do bundle.

---

### Reflexão 2: Manutenibilidade e Escalabilidade

**Pergunta:** Você está trabalhando em um projeto grande com uma equipe de 5 desenvolvedores. Você decidiu criar um plugin customizado para adicionar classes específicas da marca. Quais são os prós e contras dessa decisão? Como você garantiria que todos na equipe entendam e usem o plugin corretamente?

**Pense sobre:**
- Documentação do plugin
- Onboarding de novos desenvolvedores
- Manutenção a longo prazo
- Alternativas (CSS customizado, componentes)
- Versionamento e atualizações

**Resposta esperada:** Prós: reutilização, consistência, centralização. Contras: curva de aprendizado, necessidade de documentação, manutenção adicional. Estratégias: documentar bem, criar exemplos, considerar se realmente precisa ser um plugin ou pode ser CSS/componentes, versionar adequadamente.

---

### Reflexão 3: Quando NÃO Usar Plugins

**Pergunta:** Em que situações você NÃO deveria usar um plugin, mesmo que ele resolva seu problema? Dê exemplos práticos.

**Pense sobre:**
- Complexidade vs benefício
- Tamanho do projeto
- Necessidade única vs reutilização
- Performance
- Dependências

**Resposta esperada:** Não usar quando: é uma necessidade única (CSS customizado é suficiente), projeto muito pequeno (overhead não vale a pena), plugin adiciona muito CSS para pouco uso, você precisa de controle total sobre a implementação, ou quando CSS nativo moderno já resolve (ex: aspect-ratio nativo em navegadores modernos).

---

### Reflexão 4: Plugins vs CSS Customizado vs @apply

**Pergunta:** Você precisa de uma funcionalidade que não existe no Tailwind. Como você decide entre:
1. Criar um plugin
2. Escrever CSS customizado
3. Usar @apply para criar componentes

Dê exemplos de quando cada abordagem é mais apropriada.

**Pense sobre:**
- Reutilização
- Complexidade
- Manutenibilidade
- Performance
- Contexto do projeto

**Resposta esperada:** 
- **Plugin**: Quando é uma funcionalidade reutilizável em múltiplos projetos, adiciona muitas classes relacionadas, ou é parte do sistema de design
- **CSS customizado**: Quando é uma necessidade única, muito específica, ou quando CSS nativo já resolve
- **@apply**: Quando é um componente específico do projeto, não precisa ser reutilizável em outros projetos, ou quando você quer manter tudo no Tailwind

---

### Reflexão 5: Ecossistema e Dependências

**Pergunta:** Você encontrou um plugin da comunidade que resolve exatamente seu problema. Quais fatores você consideraria antes de adicioná-lo ao projeto? O que poderia dar errado?

**Pense sobre:**
- Manutenção do plugin (última atualização, issues abertas)
- Compatibilidade com versões do Tailwind
- Tamanho e impacto no bundle
- Licença
- Alternativas oficiais ou nativas

**Resposta esperada:** Considerar: última atualização, número de downloads/stars, issues abertas, compatibilidade, licença, se há alternativa oficial, impacto no bundle, se o autor mantém ativamente. Problemas potenciais: plugin abandonado, incompatibilidade com atualizações do Tailwind, bugs não corrigidos, dependências conflitantes.

---

## 📊 Checklist de Aprendizado

Marque o que você conseguiu fazer:

- [ ] Instalei e configurei o plugin Typography
- [ ] Usei modificadores do Typography (tamanho, cor)
- [ ] Instalei e configurei o plugin Forms
- [ ] Criei formulários estilizados com o plugin Forms
- [ ] Instalei e configurei o plugin Aspect Ratio
- [ ] Usei Aspect Ratio para manter proporções
- [ ] Instalei e configurei o plugin Line Clamp
- [ ] Usei Line Clamp para truncar texto
- [ ] Criei um plugin customizado do zero
- [ ] Entendi quando usar plugins vs CSS customizado
- [ ] Refleti sobre impacto de plugins no bundle size
- [ ] Considerei manutenibilidade ao escolher plugins

---

## 🎯 Desafio Final: Projeto Completo

Crie uma página completa que combine TODOS os plugins aprendidos:

1. **Seção de Blog** (Typography)
   - Artigo com título, subtítulos, listas, citações

2. **Formulário de Contato** (Forms)
   - Todos os tipos de input estilizados

3. **Galeria de Imagens** (Aspect Ratio)
   - Diferentes proporções mantidas

4. **Cards de Produto** (Line Clamp)
   - Títulos e descrições truncados

5. **Elementos Customizados** (Seu plugin)
   - Use seu plugin de text-shadow em títulos

### Requisitos do Desafio

- Layout responsivo
- Design consistente
- Todos os plugins funcionando
- Código organizado e comentado
- Performance considerada (não instale plugins desnecessários)

---

## 💡 Dicas para os Exercícios

1. **Comece simples**: Instale um plugin por vez e teste antes de adicionar mais
2. **Leia a documentação**: Cada plugin tem suas particularidades
3. **Teste responsividade**: Plugins devem funcionar em todos os tamanhos de tela
4. **Monitore o bundle**: Use DevTools para ver o tamanho do CSS gerado
5. **Documente seus plugins**: Se criar plugins customizados, documente bem

---

**Bons exercícios! 🚀**

