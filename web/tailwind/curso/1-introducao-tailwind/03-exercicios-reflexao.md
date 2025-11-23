# Aula 1 - Exercícios e Reflexão: Introdução ao Tailwind CSS

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para consolidar seu aprendizado sobre Tailwind CSS e a filosofia utility-first. Eles vão desde o básico até desafios que combinam múltiplos conceitos. Faça cada exercício com calma e sempre relacione as classes Tailwind com as propriedades CSS que você já conhece.

---

## 📝 Exercício 1: Traduzindo CSS para Tailwind

### Tarefa:
Traduza as seguintes regras CSS para classes Tailwind equivalentes. Use o Play CDN do Tailwind para testar.

### CSS 1:
```css
.elemento {
  padding: 1rem;
  background-color: rgb(59 130 246);
  color: white;
  border-radius: 0.5rem;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 2:
```css
.container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 1.5rem;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 3:
```css
.titulo {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  color: rgb(31 41 55);
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 4:
```css
.botao {
  padding: 0.5rem 1rem;
  background-color: rgb(34 197 94);
  color: white;
  border-radius: 0.375rem;
  font-weight: 500;
}

.botao:hover {
  background-color: rgb(22 163 74);
}
```

**Sua resposta (classes Tailwind):**

---

## 📝 Exercício 2: Traduzindo Tailwind para CSS

### Tarefa:
Traduza as seguintes classes Tailwind para CSS puro equivalente.

### Tailwind 1:
```html
<div class="p-6 bg-gray-100 rounded-lg shadow-md">
```

**Sua resposta (CSS):**

---

### Tailwind 2:
```html
<div class="flex flex-col items-center gap-4 p-8">
```

**Sua resposta (CSS):**

---

### Tailwind 3:
```html
<h1 class="text-3xl font-bold text-blue-600 mb-4">
```

**Sua resposta (CSS):**

---

### Tailwind 4:
```html
<button class="px-6 py-3 bg-red-500 text-white rounded-full hover:bg-red-600 transition-colors">
```

**Sua resposta (CSS):**

---

## 📝 Exercício 3: Criando Componentes com Tailwind

### Tarefa 1: Card de Produto

Crie um card de produto usando apenas classes Tailwind. O card deve ter:

- Largura máxima de 300px
- Fundo branco
- Padding de 1.5rem
- Bordas arredondadas
- Sombra suave
- Título com fonte grande e negrito
- Descrição com texto cinza
- Preço destacado em azul
- Botão de ação

**HTML de referência:**
```html
<div class="...">
  <img src="produto.jpg" alt="Produto" class="...">
  <div class="...">
    <h3 class="...">Nome do Produto</h3>
    <p class="...">Descrição do produto aqui.</p>
    <div class="...">
      <span class="...">R$ 99,90</span>
      <button class="...">Comprar</button>
    </div>
  </div>
</div>
```

**Sua resposta (complete as classes):**

---

### Tarefa 2: Header de Navegação

Crie um header de navegação horizontal usando Tailwind. Deve ter:

- Fundo escuro (cinza escuro ou preto)
- Texto branco
- Display flex
- Itens centralizados verticalmente
- Espaçamento entre logo e menu
- Links com hover effect
- Padding adequado

**HTML de referência:**
```html
<header class="...">
  <div class="...">Logo</div>
  <nav class="...">
    <a href="#" class="...">Home</a>
    <a href="#" class="...">Sobre</a>
    <a href="#" class="...">Contato</a>
  </nav>
</header>
```

**Sua resposta (complete as classes):**

---

### Tarefa 3: Layout de Grid Simples

Crie um layout de grid com 3 colunas usando Tailwind. Cada item deve ter:

- Fundo cinza claro
- Padding
- Bordas arredondadas
- Espaçamento entre itens

**HTML de referência:**
```html
<div class="...">
  <div class="...">Item 1</div>
  <div class="...">Item 2</div>
  <div class="...">Item 3</div>
</div>
```

**Sua resposta (complete as classes):**

---

## 📝 Exercício 4: Comparando Abordagens

### Tarefa:
Analise o seguinte componente criado de duas formas diferentes e responda as perguntas.

**Abordagem CSS Tradicional:**
```html
<div class="card">
  <h2 class="card-title">Título</h2>
  <p class="card-text">Texto</p>
</div>
```

```css
.card {
  padding: 1.5rem;
  background: white;
  border-radius: 0.5rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.card-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.card-text {
  color: #666;
  line-height: 1.6;
}
```

**Abordagem Tailwind:**
```html
<div class="p-6 bg-white rounded-lg shadow-sm">
  <h2 class="text-2xl font-semibold mb-2">Título</h2>
  <p class="text-gray-500 leading-relaxed">Texto</p>
</div>
```

**Perguntas:**

1. Qual abordagem tem mais linhas de código total?
2. Qual abordagem permite ver todos os estilos diretamente no HTML?
3. Se você precisasse mudar o padding de `1.5rem` para `2rem`, onde faria a mudança em cada abordagem?
4. Qual abordagem força mais consistência de design?
5. Qual abordagem você acha mais fácil de entender para um desenvolvedor que nunca viu o código antes?

**Suas respostas:**

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Filosofia Utility-First

**Pergunta:** Você está trabalhando em um projeto onde precisa criar 20 cards diferentes, cada um com pequenas variações (diferentes cores, tamanhos, espaçamentos). Como a abordagem utility-first do Tailwind ajudaria nessa situação? Quais seriam as vantagens em relação ao CSS tradicional?

**Pense sobre:**
- Como você criaria 20 cards com CSS tradicional?
- Como você criaria 20 cards com Tailwind?
- Qual abordagem seria mais rápida?
- Qual abordagem seria mais fácil de manter?
- Como você garantiria consistência em cada abordagem?

**Sua resposta:**

---

### Reflexão 2: Mapeamento Mental CSS → Tailwind

**Cenário:** Você está aprendendo Tailwind e vê esta classe: `bg-blue-500 hover:bg-blue-600 transition-colors`

**Pergunta:** 
1. Que propriedades CSS você identifica nessa classe?
2. Como você explicaria para alguém que nunca viu Tailwind o que essa classe faz?
3. Por que é importante entender CSS para usar Tailwind efetivamente?
4. Como você treinaria seu "mapeamento mental" para se tornar mais rápido com Tailwind?

**Pense sobre:**
- A relação entre conhecimento de CSS e facilidade com Tailwind
- Como o mapeamento mental se torna automático com prática
- A importância de entender o "porquê" além do "como"

**Sua resposta:**

---

### Reflexão 3: Quando Usar Tailwind vs CSS Puro

**Cenário 1:** Você precisa criar uma animação complexa de partículas que se movem em padrões específicos baseados em cálculos matemáticos.

**Cenário 2:** Você precisa criar 50 botões diferentes, cada um com cores e tamanhos ligeiramente diferentes, mas seguindo um design system consistente.

**Cenário 3:** Você precisa criar um layout de dashboard com cards, gráficos, tabelas e formulários, todos seguindo o mesmo design system.

**Pergunta:** Para cada cenário, você usaria Tailwind, CSS puro, ou uma combinação? Por quê?

**Pense sobre:**
- Quando Tailwind é mais apropriado?
- Quando CSS puro é mais apropriado?
- Como decidir entre as duas abordagens?
- A abordagem híbrida faz sentido em algum caso?

**Sua resposta:**

---

### Reflexão 4: Produtividade e Manutenibilidade

**Pergunta:** Considere um projeto grande com 100 componentes diferentes. Compare a manutenibilidade entre CSS tradicional e Tailwind em termos de:

1. **Tempo para criar novos componentes**
2. **Facilidade de encontrar e modificar estilos**
3. **Consistência visual**
4. **Onboarding de novos desenvolvedores**
5. **Refatoração e mudanças de design**

Qual abordagem você acha melhor para cada aspecto? Por quê?

**Pense sobre:**
- O que acontece quando você precisa mudar a cor primária do site?
- O que acontece quando um novo desenvolvedor entra no projeto?
- Como você garante que todos os botões tenham o mesmo estilo?
- O que é mais fácil de debugar quando algo não está funcionando?

**Sua resposta:**

---

### Reflexão 5: Play CDN vs Build Process

**Pergunta:** Você está começando um novo projeto. Como você decidiria entre usar Play CDN ou Build Process? Quais fatores influenciariam sua decisão?

**Pense sobre:**
- Em que fase do projeto você está?
- Qual é o tamanho do projeto?
- Você precisa de customização?
- Você precisa de otimização?
- Qual é o prazo do projeto?
- Você trabalha sozinho ou em equipe?

**Sua resposta:**

---

## 🎯 Desafio Final: Criando uma Página Completa

### Tarefa:
Crie uma página de landing page simples usando apenas Tailwind CSS (Play CDN). A página deve ter:

1. **Header:**
   - Fundo escuro
   - Logo à esquerda
   - Menu de navegação à direita
   - Links com hover effect

2. **Hero Section:**
   - Título grande e chamativo
   - Subtítulo
   - Botão de call-to-action
   - Fundo com gradiente ou cor sólida

3. **Seção de Features (3 colunas):**
   - Título da seção
   - 3 cards lado a lado
   - Cada card com ícone (ou emoji), título e descrição
   - Layout responsivo (empilha em mobile)

4. **Footer:**
   - Fundo escuro
   - Texto centralizado
   - Links ou informações de contato

**Requisitos técnicos:**
- Use apenas classes Tailwind
- Layout responsivo (pense em mobile-first)
- Cores consistentes
- Espaçamento adequado
- Hover effects onde apropriado

**HTML de estrutura (complete com classes Tailwind):**
```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Landing Page - Tailwind</title>
  <script src="https://cdn.tailwindcss.com"></script>
</head>
<body>
  <!-- Header -->
  <header class="...">
    <!-- Complete aqui -->
  </header>

  <!-- Hero Section -->
  <section class="...">
    <!-- Complete aqui -->
  </section>

  <!-- Features Section -->
  <section class="...">
    <!-- Complete aqui -->
  </section>

  <!-- Footer -->
  <footer class="...">
    <!-- Complete aqui -->
  </footer>
</body>
</html>
```

**Sua resposta (complete o HTML com classes Tailwind):**

---

## 📚 Dicas para Resolver os Exercícios

1. **Use o Play CDN:** Acesse https://play.tailwindcss.com para testar rapidamente
2. **Consulte a documentação:** https://tailwindcss.com/docs quando precisar
3. **Pense em CSS primeiro:** Antes de escrever Tailwind, pense no CSS equivalente
4. **Use DevTools:** Inspecione elementos para ver o CSS gerado
5. **Experimente:** Não tenha medo de tentar diferentes classes e ver o resultado

---

## ✅ Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Traduzir propriedades CSS comuns para classes Tailwind
- [ ] Traduzir classes Tailwind para CSS puro
- [ ] Criar componentes básicos usando apenas Tailwind
- [ ] Entender a diferença entre abordagem tradicional e utility-first
- [ ] Decidir quando usar Tailwind vs CSS puro
- [ ] Explicar a filosofia utility-first
- [ ] Mapear mentalmente classes Tailwind para propriedades CSS
- [ ] Criar layouts responsivos básicos com Tailwind

---

## 🎓 Próximos Passos

Após completar estes exercícios e reflexões, você estará pronto para:
- Aprender o sistema de espaçamento detalhado do Tailwind
- Trabalhar com cores e backgrounds em profundidade
- Dominar tipografia com Tailwind
- Trabalhar com bordas, arredondamento e sombras

Lembre-se: a prática é essencial. Quanto mais você experimentar, mais confiança terá com Tailwind. E sempre relacione com o CSS que você já conhece - essa é a chave para aprender Tailwind rapidamente!

