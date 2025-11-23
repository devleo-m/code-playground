# Aula 2 - Exercícios e Reflexão: Fundamentos do Sistema de Classes Utilitárias

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para consolidar seu aprendizado sobre o sistema de classes utilitárias do Tailwind CSS. Você praticará espaçamento, cores, tipografia, bordas, sombras e opacidade. Sempre relacione as classes Tailwind com as propriedades CSS que você já conhece.

---

## 📝 Exercício 1: Traduzindo CSS para Tailwind - Espaçamento

### Tarefa:
Traduza as seguintes regras CSS para classes Tailwind equivalentes, focando em espaçamento (padding, margin, gap).

### CSS 1:
```css
.elemento {
  padding: 1.5rem;
  margin: 2rem;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 2:
```css
.container {
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  padding-left: 1rem;
  padding-right: 1rem;
  margin-bottom: 1rem;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 3:
```css
.flex-container {
  display: flex;
  gap: 1.5rem;
  padding: 2rem;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 4:
```css
.card {
  padding: 1rem 2rem;
  margin: 0 auto;
  margin-top: 1rem;
}
```

**Sua resposta (classes Tailwind):**

---

## 📝 Exercício 2: Traduzindo CSS para Tailwind - Cores e Backgrounds

### Tarefa:
Traduza as seguintes regras CSS para classes Tailwind equivalentes, focando em cores e backgrounds.

### CSS 1:
```css
.texto {
  color: rgb(59 130 246);
  background-color: rgb(243 244 246);
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 2:
```css
.botao {
  background-color: rgb(34 197 94);
  color: white;
  border: 2px solid rgb(22 163 74);
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 3:
```css
.gradiente {
  background-image: linear-gradient(to right, rgb(59 130 246), rgb(168 85 247));
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 4:
```css
.elemento {
  background-color: rgba(59, 130, 246, 0.5);
  color: rgba(31, 41, 55, 0.75);
}
```

**Sua resposta (classes Tailwind - use a sintaxe de opacidade moderna):**

---

## 📝 Exercício 3: Traduzindo CSS para Tailwind - Tipografia

### Tarefa:
Traduza as seguintes regras CSS para classes Tailwind equivalentes, focando em tipografia.

### CSS 1:
```css
.titulo {
  font-size: 1.5rem;
  font-weight: 700;
  line-height: 2rem;
  color: rgb(31 41 55);
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 2:
```css
.texto {
  font-size: 0.875rem;
  font-weight: 400;
  line-height: 1.25rem;
  letter-spacing: 0.025em;
  text-align: center;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 3:
```css
.destaque {
  text-transform: uppercase;
  text-decoration: underline;
  font-weight: 600;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 4:
```css
.truncado {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
```

**Sua resposta (classes Tailwind):**

---

## 📝 Exercício 4: Traduzindo CSS para Tailwind - Bordas e Sombras

### Tarefa:
Traduza as seguintes regras CSS para classes Tailwind equivalentes, focando em bordas e sombras.

### CSS 1:
```css
.card {
  border-width: 1px;
  border-color: rgb(229 231 235);
  border-radius: 0.5rem;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 2:
```css
.botao {
  border-width: 2px;
  border-style: solid;
  border-color: rgb(59 130 246);
  border-radius: 9999px;
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 3:
```css
.elemento {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 
              0 4px 6px -4px rgba(0, 0, 0, 0.1);
}
```

**Sua resposta (classes Tailwind):**

---

### CSS 4:
```css
.card {
  border-top-left-radius: 0.75rem;
  border-top-right-radius: 0.75rem;
  border-bottom-width: 0;
}
```

**Sua resposta (classes Tailwind):**

---

## 📝 Exercício 5: Traduzindo Tailwind para CSS

### Tarefa:
Traduza as seguintes classes Tailwind para CSS puro equivalente.

### Tailwind 1:
```html
<div class="p-8 bg-blue-500 text-white rounded-xl shadow-lg">
```

**Sua resposta (CSS):**

---

### Tailwind 2:
```html
<div class="px-6 py-4 bg-gradient-to-r from-purple-500 to-pink-500 text-white font-bold">
```

**Sua resposta (CSS):**

---

### Tailwind 3:
```html
<h1 class="text-4xl font-extrabold text-gray-900 tracking-tight mb-6">
```

**Sua resposta (CSS):**

---

### Tailwind 4:
```html
<div class="border-4 border-red-500 rounded-full p-4 bg-red-50 opacity-75">
```

**Sua resposta (CSS):**

---

## 📝 Exercício 6: Criando Componentes com Classes Utilitárias

### Tarefa 1: Card de Informação

Crie um card de informação usando apenas classes Tailwind. O card deve ter:

- Padding de 1.5rem
- Fundo branco
- Bordas arredondadas (0.5rem)
- Sombra média
- Borda sutil (1px, cinza claro)
- Título grande (2xl), negrito, cor cinza escuro
- Texto descritivo (base), cor cinza médio, line-height relaxado
- Badge/etiqueta pequena com fundo colorido e texto correspondente

**HTML de referência:**
```html
<div class="...">
  <div class="...">
    <span class="...">Novo</span>
    <h3 class="...">Título do Card</h3>
  </div>
  <p class="...">Descrição detalhada do conteúdo do card aqui.</p>
</div>
```

**Sua resposta (complete as classes):**

---

### Tarefa 2: Botão com Variações

Crie três variações de botão usando Tailwind:

1. **Botão Primário:**
   - Padding horizontal 1.5rem, vertical 0.75rem
   - Fundo azul (500)
   - Texto branco
   - Fonte semi-negrito
   - Bordas arredondadas (0.5rem)
   - Sombra pequena

2. **Botão Secundário:**
   - Mesmo padding do primário
   - Fundo transparente
   - Borda azul (500) de 2px
   - Texto azul (500)
   - Mesma fonte e bordas arredondadas

3. **Botão de Ação (Circular):**
   - Largura e altura de 3rem
   - Fundo verde (500)
   - Texto branco
   - Bordas completamente arredondadas (full)
   - Sombra média

**HTML de referência:**
```html
<!-- Botão Primário -->
<button class="...">Clique Aqui</button>

<!-- Botão Secundário -->
<button class="...">Cancelar</button>

<!-- Botão Circular -->
<button class="...">+</button>
```

**Sua resposta (complete as classes):**

---

### Tarefa 3: Badge/Etiqueta de Status

Crie três badges de status diferentes:

1. **Status Ativo:**
   - Fundo verde claro (100)
   - Texto verde escuro (800)
   - Padding pequeno (horizontal 0.75rem, vertical 0.25rem)
   - Fonte média (500)
   - Tamanho de texto pequeno (sm)
   - Bordas completamente arredondadas

2. **Status Pendente:**
   - Fundo amarelo claro (100)
   - Texto amarelo escuro (800)
   - Mesmas outras propriedades

3. **Status Inativo:**
   - Fundo cinza claro (100)
   - Texto cinza escuro (800)
   - Opacidade de 60%
   - Mesmas outras propriedades

**HTML de referência:**
```html
<span class="...">Ativo</span>
<span class="...">Pendente</span>
<span class="...">Inativo</span>
```

**Sua resposta (complete as classes):**

---

### Tarefa 4: Card com Gradiente e Overlay

Crie um card com gradiente de fundo e texto sobreposto:

- Padding de 2rem
- Gradiente linear de azul (500) para roxo (500), da esquerda para direita
- Texto branco
- Título grande (3xl), negrito
- Subtítulo médio (lg), com opacidade de 90%
- Bordas arredondadas grandes (xl)
- Sombra extra grande (2xl)

**HTML de referência:**
```html
<div class="...">
  <h2 class="...">Título Principal</h2>
  <p class="...">Subtítulo descritivo</p>
</div>
```

**Sua resposta (complete as classes):**

---

## 📝 Exercício 7: Análise de Código

### Tarefa:
Analise o seguinte código HTML com classes Tailwind e identifique:

1. Quais propriedades CSS estão sendo aplicadas
2. Se há alguma inconsistência ou problema
3. Como você melhoraria o código

### Código para Análise:

```html
<div class="p-6 bg-white rounded-lg shadow-md border border-gray-200">
  <div class="flex items-center justify-between mb-4">
    <h2 class="text-2xl font-bold text-gray-800">Título do Card</h2>
    <span class="px-3 py-1 bg-blue-100 text-blue-800 text-sm font-medium rounded-full">
      Novo
    </span>
  </div>
  <p class="text-gray-600 leading-relaxed mb-6">
    Este é um parágrafo de descrição que contém informações importantes sobre o conteúdo do card.
  </p>
  <div class="flex gap-4">
    <button class="px-6 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600">
      Ação Principal
    </button>
    <button class="px-6 py-2 border-2 border-blue-500 text-blue-500 font-semibold rounded-lg hover:bg-blue-50">
      Ação Secundária
    </button>
  </div>
</div>
```

**Sua análise:**

1. **Propriedades CSS aplicadas (liste todas):**

---

2. **Inconsistências ou problemas identificados:**

---

3. **Melhorias sugeridas:**

---

## 📝 Exercício 8: Desafio - Construindo um Perfil de Usuário

### Tarefa:
Crie um componente completo de perfil de usuário usando apenas classes Tailwind. O componente deve incluir:

**Estrutura:**
- Container principal com fundo branco, padding, bordas arredondadas e sombra
- Avatar circular (pode ser um div colorido) com borda
- Nome do usuário (grande, negrito)
- Cargo/função (médio, cinza)
- Bio/descrição (texto normal, cinza médio)
- Badges de habilidades (múltiplos badges pequenos)
- Botões de ação (seguir, mensagem)

**Requisitos específicos:**
- Use espaçamento consistente
- Use cores da paleta Tailwind
- Aplique tipografia variada
- Use bordas e sombras apropriadas
- Garanta boa hierarquia visual

**HTML de referência:**
```html
<div class="...">
  <!-- Avatar -->
  <div class="...">
    <div class="..."></div>
  </div>
  
  <!-- Informações -->
  <div class="...">
    <h2 class="...">Nome do Usuário</h2>
    <p class="...">Desenvolvedor Full Stack</p>
    <p class="...">Descrição da bio do usuário aqui...</p>
  </div>
  
  <!-- Badges -->
  <div class="...">
    <span class="...">React</span>
    <span class="...">TypeScript</span>
    <span class="...">Node.js</span>
  </div>
  
  <!-- Botões -->
  <div class="...">
    <button class="...">Seguir</button>
    <button class="...">Mensagem</button>
  </div>
</div>
```

**Sua resposta (complete todas as classes):**

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Sistema de Espaçamento

**Pergunta:** Por que o Tailwind usa uma escala de espaçamento baseada em múltiplos de 0.25rem (4px) em vez de permitir valores arbitrários como no CSS tradicional?

**Sua resposta:**

---

**Pergunta de seguimento:** Quais são as vantagens e desvantagens dessa abordagem padronizada? Em que situações você acharia útil ter valores customizados?

**Sua resposta:**

---

### Reflexão 2: Sistema de Cores

**Pergunta:** O Tailwind organiza cores em uma escala de 50 a 950. Por que essa organização é mais eficiente do que usar cores hexadecimais arbitrárias como `#3A7B9F`?

**Sua resposta:**

---

**Pergunta de seguimento:** Como a escala de cores do Tailwind ajuda na criação de designs consistentes? Pense em termos de acessibilidade e contraste.

**Sua resposta:**

---

### Reflexão 3: Tipografia e Legibilidade

**Pergunta:** O Tailwind combina `font-size` e `line-height` em uma única classe (ex: `text-xl` aplica tanto o tamanho quanto a altura da linha). Por que essa combinação faz sentido do ponto de vista de design?

**Sua resposta:**

---

**Pergunta de seguimento:** Como você garantiria que o texto seja legível em diferentes dispositivos e tamanhos de tela usando apenas as classes Tailwind? Quais classes você usaria?

**Sua resposta:**

---

### Reflexão 4: Bordas e Arredondamento

**Pergunta:** Por que o Tailwind oferece classes específicas para arredondar cantos individuais (`rounded-tl`, `rounded-tr`, etc.) além das classes gerais (`rounded-lg`)?

**Sua resposta:**

---

**Pergunta de seguimento:** Em que situações de design você usaria cantos arredondados assimétricos? Dê exemplos práticos.

**Sua resposta:**

---

### Reflexão 5: Sombras e Profundidade Visual

**Pergunta:** As sombras do Tailwind são pré-definidas (sm, md, lg, xl, 2xl). Quais são os benefícios de ter sombras padronizadas versus criar sombras customizadas com valores arbitrários?

**Sua resposta:**

---

**Pergunta de seguimento:** Como as sombras contribuem para a hierarquia visual de uma interface? Dê um exemplo de como você usaria diferentes níveis de sombra para criar profundidade.

**Sua resposta:**

---

### Reflexão 6: Opacidade e Camadas Visuais

**Pergunta:** O Tailwind oferece duas formas de aplicar opacidade: classes dedicadas (`opacity-50`) e sintaxe moderna em cores (`bg-blue-500/50`). Quando você usaria cada abordagem?

**Sua resposta:**

---

**Pergunta de seguimento:** Como a opacidade pode ser usada para criar hierarquia visual e guiar a atenção do usuário? Pense em overlays, estados de hover, e elementos desabilitados.

**Sua resposta:**

---

### Reflexão 7: Consistência vs Flexibilidade

**Pergunta:** O sistema de classes utilitárias do Tailwind prioriza consistência (valores padronizados) sobre flexibilidade (valores arbitrários). Quais são os trade-offs dessa decisão?

**Sua resposta:**

---

**Pergunta de seguimento:** Em que situações você acharia necessário usar CSS customizado mesmo tendo Tailwind disponível? Dê exemplos específicos.

**Sua resposta:**

---

### Reflexão 8: Performance e Bundle Size

**Pergunta:** Como o sistema de classes utilitárias do Tailwind impacta o tamanho do bundle CSS final? Pense em quantas classes são geradas versus quantas são realmente usadas.

**Sua resposta:**

---

**Pergunta de seguimento:** Quais estratégias o Tailwind usa para minimizar o CSS não utilizado? Como isso se compara ao CSS tradicional onde você escreve apenas o que precisa?

**Sua resposta:**

---

## ✅ Checklist de Aprendizado

Antes de avançar para a próxima aula, certifique-se de que você consegue:

- [ ] Traduzir propriedades CSS de espaçamento para classes Tailwind e vice-versa
- [ ] Usar o sistema de cores do Tailwind (escala 50-950) corretamente
- [ ] Aplicar tipografia (tamanhos, pesos, alinhamento) usando classes Tailwind
- [ ] Criar bordas e arredondamentos com diferentes variações
- [ ] Aplicar sombras apropriadas para criar profundidade visual
- [ ] Controlar opacidade de elementos e cores
- [ ] Combinar múltiplas classes para criar componentes visuais
- [ ] Entender o mapeamento mental entre classes Tailwind e CSS puro
- [ ] Identificar quando usar valores padronizados vs quando precisar de CSS customizado

---

## 🎓 Próximos Passos

Após completar estes exercícios e reflexões, você estará pronto para:
- **Aula 3:** Layout com Tailwind - Display, Position e Flexbox
- **Aula 4:** CSS Grid com Tailwind
- **Aula 5:** Responsividade com Tailwind

Continue praticando combinando essas classes fundamentais. Quanto mais você praticar, mais natural se tornará o uso do sistema de classes utilitárias!

