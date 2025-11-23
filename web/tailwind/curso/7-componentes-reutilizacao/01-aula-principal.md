# Aula 7: Componentes e Reutilização com @apply - Conteúdo Principal

## 📖 Revisão: Componentes e Reutilização em CSS

Antes de mergulharmos em componentes com Tailwind, vamos relembrar os conceitos fundamentais que você já conhece de CSS puro:

**Componentes CSS** são classes reutilizáveis que encapsulam estilos relacionados para criar elementos visuais consistentes.

### Componentes em CSS Puro

Você já aprendeu que em CSS puro criamos componentes usando classes:

```css
/* CSS Puro - Componente de botão */
.button {
  padding: 0.5rem 1rem;
  background-color: rgb(59, 130, 246);
  color: white;
  border-radius: 0.25rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 150ms;
}

.button:hover {
  background-color: rgb(37, 99, 235);
}

.button-primary {
  background-color: rgb(59, 130, 246);
}

.button-secondary {
  background-color: rgb(107, 114, 128);
}
```

### Componentes no Tailwind

No Tailwind, você tem duas abordagens principais:
1. **Utilitários diretos** (abordagem padrão utility-first)
2. **Componentes com @apply** (quando a reutilização faz sentido)

---

## 🎯 Quando Usar Utilitários vs Componentes

### Abordagem Utility-First (Padrão)

A filosofia do Tailwind é usar utilitários diretamente no HTML:

```html
<!-- Abordagem utility-first -->
<button class="px-4 py-2 bg-blue-500 text-white rounded font-medium hover:bg-blue-600 transition-colors">
  Clique aqui
</button>
```

**Vantagens:**
- Flexibilidade total
- Sem abstrações desnecessárias
- Fácil de ver o que está acontecendo
- Menos arquivos CSS para manter

**Equivalente em CSS puro:**
```css
button {
  padding: 0.5rem 1rem;
  background-color: rgb(59, 130, 246);
  color: white;
  border-radius: 0.25rem;
  font-weight: 500;
}

button:hover {
  background-color: rgb(37, 99, 235);
}

button {
  transition-property: color, background-color, border-color;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
```

### Quando Criar Componentes com @apply

Crie componentes quando:

1. **Repetição excessiva**: Você repete o mesmo conjunto de classes muitas vezes
2. **Lógica complexa**: O componente tem muitas variantes ou estados
3. **Manutenibilidade**: Mudanças futuras precisam ser feitas em um só lugar
4. **Legibilidade**: Classes utilitárias tornam o HTML muito verboso

**Exemplo de quando criar componente:**

```html
<!-- ANTES: Repetição excessiva -->
<button class="px-4 py-2 bg-blue-500 text-white rounded font-medium hover:bg-blue-600 transition-colors">Botão 1</button>
<button class="px-4 py-2 bg-blue-500 text-white rounded font-medium hover:bg-blue-600 transition-colors">Botão 2</button>
<button class="px-4 py-2 bg-blue-500 text-white rounded font-medium hover:bg-blue-600 transition-colors">Botão 3</button>
```

```css
/* DEPOIS: Componente reutilizável */
.btn {
  @apply px-4 py-2 bg-blue-500 text-white rounded font-medium hover:bg-blue-600 transition-colors;
}
```

```html
<!-- HTML mais limpo -->
<button class="btn">Botão 1</button>
<button class="btn">Botão 2</button>
<button class="btn">Botão 3</button>
```

---

## 🔧 A Diretiva @apply

### O que é @apply?

A diretiva `@apply` permite que você **aplique classes utilitárias do Tailwind dentro de regras CSS customizadas**, criando componentes reutilizáveis.

### Sintaxe Básica

```css
.componente {
  @apply classe-utilitaria-1 classe-utilitaria-2 classe-utilitaria-3;
}
```

### Exemplo Básico

```css
/* Criando um componente de card */
.card {
  @apply p-6 bg-white rounded-lg shadow-md;
}
```

**Equivalente em CSS puro:**
```css
.card {
  padding: 1.5rem;
  background-color: white;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}
```

**Uso no HTML:**
```html
<div class="card">
  <h2>Título do Card</h2>
  <p>Conteúdo do card</p>
</div>
```

---

## 📦 Criando Componentes Comuns

### 1. Componente de Botão

```css
/* Botão base */
.btn {
  @apply px-4 py-2 rounded font-medium transition-colors duration-150;
}

/* Variantes de botão */
.btn-primary {
  @apply bg-blue-500 text-white hover:bg-blue-600;
}

.btn-secondary {
  @apply bg-gray-500 text-white hover:bg-gray-600;
}

.btn-outline {
  @apply border-2 border-blue-500 text-blue-500 hover:bg-blue-500 hover:text-white;
}
```

**Equivalente em CSS puro:**
```css
.btn {
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
  font-weight: 500;
  transition-property: color, background-color, border-color;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

.btn-primary {
  background-color: rgb(59, 130, 246);
  color: white;
}

.btn-primary:hover {
  background-color: rgb(37, 99, 235);
}
```

**Uso:**
```html
<button class="btn btn-primary">Primário</button>
<button class="btn btn-secondary">Secundário</button>
<button class="btn btn-outline">Outline</button>
```

### 2. Componente de Card

```css
.card {
  @apply bg-white rounded-lg shadow-md overflow-hidden;
}

.card-header {
  @apply px-6 py-4 border-b border-gray-200;
}

.card-body {
  @apply px-6 py-4;
}

.card-footer {
  @apply px-6 py-4 border-t border-gray-200 bg-gray-50;
}
```

**Equivalente em CSS puro:**
```css
.card {
  background-color: white;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  overflow: hidden;
}

.card-header {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid rgb(229, 231, 235);
}
```

**Uso:**
```html
<div class="card">
  <div class="card-header">
    <h3>Título do Card</h3>
  </div>
  <div class="card-body">
    <p>Conteúdo do card</p>
  </div>
  <div class="card-footer">
    <button>Ação</button>
  </div>
</div>
```

### 3. Componente de Input

```css
.input {
  @apply w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent;
}

.input-error {
  @apply border-red-500 focus:ring-red-500;
}

.input-success {
  @apply border-green-500 focus:ring-green-500;
}
```

**Equivalente em CSS puro:**
```css
.input {
  width: 100%;
  padding: 0.5rem 1rem;
  border: 1px solid rgb(209, 213, 219);
  border-radius: 0.375rem;
}

.input:focus {
  outline: 2px solid transparent;
  outline-offset: 2px;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
  border-color: transparent;
}
```

**Uso:**
```html
<input type="text" class="input" placeholder="Nome">
<input type="email" class="input input-error" placeholder="Email">
<input type="text" class="input input-success" placeholder="Telefone">
```

### 4. Componente de Badge

```css
.badge {
  @apply inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium;
}

.badge-primary {
  @apply bg-blue-100 text-blue-800;
}

.badge-success {
  @apply bg-green-100 text-green-800;
}

.badge-danger {
  @apply bg-red-100 text-red-800;
}
```

**Uso:**
```html
<span class="badge badge-primary">Novo</span>
<span class="badge badge-success">Ativo</span>
<span class="badge badge-danger">Urgente</span>
```

---

## 🎨 Combinando @apply com CSS Customizado

Você pode combinar `@apply` com propriedades CSS customizadas:

```css
.btn-custom {
  @apply px-4 py-2 rounded font-medium;
  /* CSS customizado adicional */
  text-transform: uppercase;
  letter-spacing: 0.05em;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
```

**Equivalente em CSS puro:**
```css
.btn-custom {
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
```

---

## 🔄 Variantes e Estados com @apply

### Variantes Responsivas

```css
.card {
  @apply p-4 md:p-6 lg:p-8;
}
```

**Equivalente em CSS puro:**
```css
.card {
  padding: 1rem;
}

@media (min-width: 768px) {
  .card {
    padding: 1.5rem;
  }
}

@media (min-width: 1024px) {
  .card {
    padding: 2rem;
  }
}
```

### Estados com Pseudo-classes

```css
.btn {
  @apply px-4 py-2 bg-blue-500 text-white rounded;
  @apply hover:bg-blue-600;
  @apply focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2;
  @apply active:bg-blue-700;
  @apply disabled:opacity-50 disabled:cursor-not-allowed;
}
```

**Equivalente em CSS puro:**
```css
.btn {
  padding: 0.5rem 1rem;
  background-color: rgb(59, 130, 246);
  color: white;
  border-radius: 0.25rem;
}

.btn:hover {
  background-color: rgb(37, 99, 235);
}

.btn:focus {
  outline: 2px solid transparent;
  outline-offset: 2px;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

.btn:active {
  background-color: rgb(29, 78, 216);
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
```

---

## 🏗️ Estrutura de Componentes

### Organização de Arquivos

```
projeto/
├── src/
│   ├── styles/
│   │   ├── components/
│   │   │   ├── buttons.css
│   │   │   ├── cards.css
│   │   │   ├── forms.css
│   │   │   └── badges.css
│   │   └── main.css
```

### Exemplo: buttons.css

```css
/* Botão base */
.btn {
  @apply px-4 py-2 rounded font-medium transition-colors duration-150;
  @apply focus:outline-none focus:ring-2 focus:ring-offset-2;
}

/* Variantes de cor */
.btn-primary {
  @apply bg-blue-500 text-white hover:bg-blue-600 focus:ring-blue-500;
}

.btn-secondary {
  @apply bg-gray-500 text-white hover:bg-gray-600 focus:ring-gray-500;
}

.btn-success {
  @apply bg-green-500 text-white hover:bg-green-600 focus:ring-green-500;
}

/* Variantes de tamanho */
.btn-sm {
  @apply px-2 py-1 text-sm;
}

.btn-lg {
  @apply px-6 py-3 text-lg;
}

/* Variante outline */
.btn-outline-primary {
  @apply border-2 border-blue-500 text-blue-500 hover:bg-blue-500 hover:text-white;
}
```

### Importando no main.css

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

/* Importar componentes */
@import './components/buttons.css';
@import './components/cards.css';
@import './components/forms.css';
```

---

## ⚠️ Limitações e Considerações do @apply

### 1. Não Use @apply com Classes Arbitrárias

```css
/* ❌ ERRADO */
.componente {
  @apply [alguma-classe-arbitraria];
}
```

### 2. Não Use @apply com Media Queries Diretamente

```css
/* ❌ ERRADO */
@media (min-width: 768px) {
  .componente {
    @apply p-8;
  }
}

/* ✅ CORRETO */
.componente {
  @apply p-4 md:p-8;
}
```

### 3. Não Use @apply com Pseudo-elementos Complexos

```css
/* ❌ ERRADO - pode não funcionar como esperado */
.componente::before {
  @apply content-[''] absolute;
}

/* ✅ CORRETO - use CSS customizado */
.componente::before {
  content: '';
  position: absolute;
  /* outras propriedades */
}
```

### 4. Ordem Importa

A ordem das classes no `@apply` pode afetar o resultado final devido à especificidade:

```css
/* Ordem pode importar */
.btn {
  @apply bg-blue-500 text-white;
  /* Se você adicionar mais classes depois, elas podem sobrescrever */
}
```

---

## 🎯 Padrões Comuns de Componentes

### Padrão: Componente Base + Modificadores

```css
/* Componente base */
.alert {
  @apply p-4 rounded-lg border;
}

/* Modificadores */
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
```

**Uso:**
```html
<div class="alert alert-info">Informação importante</div>
<div class="alert alert-success">Operação realizada com sucesso</div>
<div class="alert alert-warning">Atenção necessária</div>
<div class="alert alert-error">Erro encontrado</div>
```

### Padrão: Componente com Variantes de Tamanho

```css
.avatar {
  @apply rounded-full bg-gray-200 flex items-center justify-center;
}

.avatar-sm {
  @apply w-8 h-8 text-xs;
}

.avatar-md {
  @apply w-12 h-12 text-sm;
}

.avatar-lg {
  @apply w-16 h-16 text-base;
}

.avatar-xl {
  @apply w-24 h-24 text-lg;
}
```

**Uso:**
```html
<div class="avatar avatar-sm">JD</div>
<div class="avatar avatar-md">JD</div>
<div class="avatar avatar-lg">JD</div>
<div class="avatar avatar-xl">JD</div>
```

---

## 🔍 Debugging @apply

### Ver o CSS Gerado

Quando você usa `@apply`, o Tailwind gera CSS equivalente. Você pode inspecionar no DevTools:

```css
/* Seu código */
.btn {
  @apply px-4 py-2 bg-blue-500;
}
```

**CSS gerado (visível no DevTools):**
```css
.btn {
  padding-left: 1rem;
  padding-right: 1rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  background-color: rgb(59, 130, 246);
}
```

### Problemas Comuns

1. **Classes não aplicadas**: Verifique se a classe existe no Tailwind
2. **Especificidade**: Outras regras CSS podem estar sobrescrevendo
3. **Ordem de importação**: Certifique-se de que os componentes são importados após `@tailwind components`

---

## 📚 Resumo

### Conceitos Principais

1. **@apply**: Diretiva que permite aplicar classes utilitárias dentro de regras CSS
2. **Quando usar**: Quando há repetição excessiva ou necessidade de abstração
3. **Estrutura**: Organize componentes em arquivos separados
4. **Limitações**: Não use com classes arbitrárias ou media queries diretas
5. **Padrões**: Use componentes base + modificadores para variantes

### Mapeamento CSS → Tailwind

- `@apply px-4` = `padding-left: 1rem; padding-right: 1rem;`
- `@apply bg-blue-500` = `background-color: rgb(59, 130, 246);`
- `@apply hover:bg-blue-600` = `:hover { background-color: rgb(37, 99, 235); }`
- `@apply md:p-8` = `@media (min-width: 768px) { padding: 2rem; }`

### Próximos Passos

Na próxima aula, você aprenderá sobre **Customização e Configuração do Tailwind**, incluindo como personalizar cores, espaçamento, breakpoints e adicionar utilitários customizados através do arquivo `tailwind.config.js`.

