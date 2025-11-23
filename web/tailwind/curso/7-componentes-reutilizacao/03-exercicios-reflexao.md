# Aula 7 - Exercícios e Reflexão: Componentes e Reutilização com @apply

## 🎯 Objetivos dos Exercícios

Ao completar estes exercícios, você será capaz de:
- Identificar quando criar componentes vs usar utilitários diretos
- Criar componentes reutilizáveis usando @apply
- Organizar componentes de forma escalável
- Criar variantes de componentes
- Pensar criticamente sobre quando usar @apply vs CSS customizado
- Avaliar trade-offs entre flexibilidade e reutilização

---

## 📝 Exercício 1: Criando um Sistema de Botões

### Tarefa

Você precisa criar um sistema de botões para um projeto. Você notou que está repetindo as mesmas classes em vários lugares:

```html
<!-- Botão primário (repetido 15 vezes) -->
<button class="px-4 py-2 bg-blue-500 text-white rounded font-medium hover:bg-blue-600 transition-colors">
  Salvar
</button>

<!-- Botão secundário (repetido 8 vezes) -->
<button class="px-4 py-2 bg-gray-500 text-white rounded font-medium hover:bg-gray-600 transition-colors">
  Cancelar
</button>
```

### Requisitos

1. Crie um componente base `.btn` com estilos comuns
2. Crie variantes: `.btn-primary`, `.btn-secondary`, `.btn-outline`
3. Crie tamanhos: `.btn-sm`, `.btn-md`, `.btn-lg`
4. Todos os botões devem ter estados de hover, focus e disabled

### Código Base

```css
/* Seu código aqui */
```

```html
<!-- Teste seus componentes -->
<button class="btn btn-primary btn-md">Primário Médio</button>
<button class="btn btn-secondary btn-sm">Secundário Pequeno</button>
<button class="btn btn-outline btn-lg">Outline Grande</button>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```css
/* Componente base */
.btn {
  @apply px-4 py-2 rounded font-medium transition-colors duration-150;
  @apply focus:outline-none focus:ring-2 focus:ring-offset-2;
  @apply disabled:opacity-50 disabled:cursor-not-allowed;
}

/* Variantes de cor */
.btn-primary {
  @apply bg-blue-500 text-white hover:bg-blue-600 focus:ring-blue-500;
}

.btn-secondary {
  @apply bg-gray-500 text-white hover:bg-gray-600 focus:ring-gray-500;
}

.btn-outline {
  @apply border-2 border-blue-500 text-blue-500 bg-transparent;
  @apply hover:bg-blue-500 hover:text-white focus:ring-blue-500;
}

/* Tamanhos */
.btn-sm {
  @apply px-2 py-1 text-sm;
}

.btn-md {
  @apply px-4 py-2 text-base;
}

.btn-lg {
  @apply px-6 py-3 text-lg;
}
```

</details>

---

## 📝 Exercício 2: Sistema de Cards Reutilizável

### Tarefa

Crie um sistema de cards que seja flexível e reutilizável. O card deve ter:
- Header opcional
- Body (obrigatório)
- Footer opcional
- Variantes de estilo (bordered, shadowed, elevated)

### Requisitos

1. Crie componentes: `.card`, `.card-header`, `.card-body`, `.card-footer`
2. Crie variantes: `.card-bordered`, `.card-shadowed`, `.card-elevated`
3. O card deve ser responsivo (padding menor em mobile)

### Código Base

```css
/* Seu código aqui */
```

```html
<!-- Teste seus componentes -->
<div class="card card-shadowed">
  <div class="card-header">
    <h3>Título do Card</h3>
  </div>
  <div class="card-body">
    <p>Conteúdo do card aqui</p>
  </div>
  <div class="card-footer">
    <button>Ação</button>
  </div>
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```css
/* Card base */
.card {
  @apply bg-white rounded-lg overflow-hidden;
  @apply p-4 md:p-6;
}

/* Variantes de estilo */
.card-bordered {
  @apply border border-gray-200;
}

.card-shadowed {
  @apply shadow-md;
}

.card-elevated {
  @apply shadow-lg hover:shadow-xl transition-shadow;
}

/* Partes do card */
.card-header {
  @apply mb-4 pb-4 border-b border-gray-200;
}

.card-body {
  @apply mb-4;
}

.card-footer {
  @apply mt-4 pt-4 border-t border-gray-200 bg-gray-50 -mx-4 -mb-4 px-4 py-4;
}

.card-footer:first-child {
  @apply mt-0;
}
```

</details>

---

## 📝 Exercício 3: Sistema de Alertas

### Tarefa

Crie um sistema de alertas (mensagens de feedback) com diferentes tipos e ícones opcionais.

### Requisitos

1. Crie componente base `.alert`
2. Crie variantes: `.alert-info`, `.alert-success`, `.alert-warning`, `.alert-error`
3. Cada alerta deve ter um ícone opcional à esquerda
4. O alerta deve ser responsivo e ter animação de entrada suave

### Código Base

```css
/* Seu código aqui */
```

```html
<!-- Teste seus componentes -->
<div class="alert alert-info">
  <span class="alert-icon">ℹ️</span>
  <span>Esta é uma informação importante</span>
</div>

<div class="alert alert-success">
  <span>Operação realizada com sucesso!</span>
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```css
/* Alert base */
.alert {
  @apply p-4 rounded-lg border flex items-start gap-3;
  @apply animate-fade-in;
}

/* Variantes */
.alert-info {
  @apply bg-blue-50 border-blue-200 text-blue-800;
}

.alert-success {
  @apply bg-green-50 border-green-200 text-green-800;
}

.alert-warning {
  @apply bg-yellow-50 border-yellow-200 text-yellow-800;
}

.alert-error {
  @apply bg-red-50 border-red-200 text-red-800;
}

/* Ícone opcional */
.alert-icon {
  @apply text-xl flex-shrink-0;
}

/* Animação customizada (se necessário) */
@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.3s ease-out;
}
```

</details>

---

## 📝 Exercício 4: Refatoração - De Utilitários para Componentes

### Tarefa

Analise o seguinte código HTML que está sendo repetido em vários lugares do projeto:

```html
<!-- Este padrão aparece 12 vezes no projeto -->
<div class="flex items-center justify-between p-4 bg-white rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow">
  <div class="flex items-center gap-3">
    <div class="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center text-white font-bold">
      JD
    </div>
    <div>
      <h3 class="font-semibold text-gray-900">João Silva</h3>
      <p class="text-sm text-gray-500">joao@email.com</p>
    </div>
  </div>
  <button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors">
    Ver Perfil
  </button>
</div>
```

### Requisitos

1. Identifique quais partes devem virar componentes
2. Crie os componentes usando @apply
3. Refatore o HTML para usar os novos componentes
4. Mantenha a mesma aparência visual

### Análise Esperada

<details>
<summary>Clique para ver a análise</summary>

**Componentes identificados:**
1. `.user-card` - O card completo do usuário
2. `.avatar` - O círculo com iniciais
3. `.btn` - O botão (já criado no exercício 1)

**Código refatorado:**

```css
/* Card de usuário */
.user-card {
  @apply flex items-center justify-between p-4 bg-white rounded-lg shadow-md;
  @apply border border-gray-200 hover:shadow-lg transition-shadow;
}

/* Avatar */
.avatar {
  @apply w-10 h-10 rounded-full flex items-center justify-center text-white font-bold;
}

.avatar-blue {
  @apply bg-blue-500;
}
```

```html
<!-- HTML refatorado -->
<div class="user-card">
  <div class="flex items-center gap-3">
    <div class="avatar avatar-blue">JD</div>
    <div>
      <h3 class="font-semibold text-gray-900">João Silva</h3>
      <p class="text-sm text-gray-500">joao@email.com</p>
    </div>
  </div>
  <button class="btn btn-primary btn-md">Ver Perfil</button>
</div>
```

</details>

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Quando Criar Componentes?

**Situação**: Você está trabalhando em um projeto e encontra este padrão repetido 3 vezes:

```html
<div class="p-6 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl shadow-lg">
  Conteúdo especial
</div>
```

**Perguntas para reflexão:**

1. **Você criaria um componente para isso? Por quê?**
   - Considere: Quantas vezes será usado? Quão específico é o design? Qual a probabilidade de mudanças futuras?

2. **Quais são os trade-offs de criar um componente vs usar utilitários diretos neste caso?**
   - Pense em: Flexibilidade, manutenibilidade, legibilidade do código

3. **Em que cenário você definitivamente criaria um componente? E em que cenário definitivamente não criaria?**

---

### Reflexão 2: Organização e Escalabilidade

**Situação**: Você está trabalhando em um projeto grande com 50+ componentes diferentes.

**Perguntas para reflexão:**

1. **Como você organizaria os arquivos de componentes?**
   - Por tipo (buttons.css, cards.css)?
   - Por funcionalidade (auth.css, dashboard.css)?
   - Por página (home.css, about.css)?
   - Qual abordagem é mais escalável?

2. **Quais problemas podem surgir com muitos componentes?**
   - Pense em: Nomeação, conflitos, manutenção, onboarding de novos desenvolvedores

3. **Como você garantiria que componentes não entrem em conflito uns com os outros?**
   - Considere: Especificidade, nomenclatura, documentação

---

### Reflexão 3: Performance e Bundle Size

**Situação**: Você criou 20 componentes usando @apply. Cada componente usa várias classes utilitárias do Tailwind.

**Perguntas para reflexão:**

1. **Qual é o impacto no tamanho do CSS final quando você usa @apply?**
   - O CSS gerado é maior, menor ou igual ao usar utilitários diretos?
   - Como o PurgeCSS/JIT afeta isso?

2. **Há alguma diferença de performance em runtime entre componentes @apply e utilitários diretos?**
   - Pense em: Renderização, especificidade CSS, cache do navegador

3. **Em um projeto grande, como você monitoraria e otimizaria o tamanho do CSS?**
   - Considere: Ferramentas, métricas, estratégias de otimização

---

### Reflexão 4: Manutenibilidade e Evolução

**Situação**: Você criou um componente `.btn` há 6 meses. Agora o design system mudou e todos os botões precisam ter:
- Bordas mais arredondadas
- Sombra diferente
- Nova animação de hover

**Perguntas para reflexão:**

1. **Quais são as vantagens de ter usado @apply neste caso?**
   - Quantos arquivos você precisaria modificar?
   - Quão fácil seria fazer a mudança?

2. **E se você tivesse usado utilitários diretos em 100 lugares?**
   - Qual seria o esforço de refatoração?
   - Quais seriam os riscos?

3. **Como você documentaria componentes para facilitar manutenção futura?**
   - Pense em: Comentários, exemplos, guia de estilo

---

### Reflexão 5: Decisão Arquitetural: @apply vs CSS Customizado

**Situação**: Você precisa criar um componente de modal que tem:
- Overlay com blur
- Animação de entrada complexa (scale + fade + slide)
- Posicionamento centralizado
- Fechamento ao clicar fora

**Perguntas para reflexão:**

1. **Você usaria apenas @apply, apenas CSS customizado, ou uma combinação? Por quê?**
   - Considere: Limitações do @apply, complexidade da animação, manutenibilidade

2. **Quais são as limitações do @apply que você conhece?**
   - Pense em: Pseudo-elementos, media queries, animações complexas

3. **Como você decidiria quando "sair" do Tailwind e usar CSS puro?**
   - Qual é o seu critério pessoal?

---

### Reflexão 6: Trabalhando em Equipe

**Situação**: Você está em uma equipe de 5 desenvolvedores trabalhando no mesmo projeto. Cada um está criando componentes conforme necessário.

**Perguntas para reflexão:**

1. **Quais problemas podem surgir quando múltiplas pessoas criam componentes?**
   - Pense em: Duplicação, inconsistência, conflitos de nomenclatura

2. **Como você estabeleceria convenções para criação de componentes?**
   - Considere: Nomenclatura (BEM?), estrutura, documentação

3. **Qual seria o processo de code review para novos componentes?**
   - O que você verificaria? Quais perguntas faria?

---

## 🎯 Critérios de Avaliação

Ao completar estes exercícios e reflexões, você deve ser capaz de:

✅ **Identificar oportunidades de reutilização** - Saber quando criar componentes
✅ **Criar componentes escaláveis** - Estruturar componentes de forma organizada
✅ **Pensar criticamente** - Avaliar trade-offs e tomar decisões arquiteturais
✅ **Manter código limpo** - Organizar e documentar componentes adequadamente
✅ **Trabalhar em equipe** - Estabelecer convenções e processos

---

## 💡 Dicas para os Exercícios

1. **Comece simples**: Crie o componente base primeiro, depois adicione variantes
2. **Teste frequentemente**: Verifique se os componentes funcionam em diferentes contextos
3. **Pense em responsividade**: Sempre considere como componentes se comportam em mobile
4. **Documente decisões**: Anote por que você criou um componente e quando usá-lo
5. **Refatore gradualmente**: Não precisa criar todos os componentes de uma vez

---

## 🚀 Próximos Passos

Após completar estes exercícios, você estará pronto para aprender sobre **Customização e Configuração do Tailwind**, onde você poderá criar seus próprios utilitários e estender o sistema de design do Tailwind.

