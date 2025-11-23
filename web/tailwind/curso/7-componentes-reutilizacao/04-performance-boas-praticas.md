# Aula 7 - Performance, Boas Práticas e Otimização: Componentes e Reutilização

## 🚀 Performance de Componentes com @apply

### Como @apply Afeta o CSS Gerado

Quando você usa `@apply`, o Tailwind gera CSS equivalente. É importante entender o impacto no tamanho do bundle e na performance.

#### CSS Gerado por @apply

```css
/* Seu código */
.btn {
  @apply px-4 py-2 bg-blue-500 text-white rounded;
}
```

**CSS gerado pelo Tailwind:**
```css
.btn {
  padding-left: 1rem;
  padding-right: 1rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  background-color: rgb(59, 130, 246);
  color: white;
  border-radius: 0.25rem;
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
```

### Impacto no Bundle Size

#### ❌ Problema: Múltiplos Componentes com Classes Duplicadas

```css
/* ❌ RUIM - Classes duplicadas */
.btn-primary {
  @apply px-4 py-2 bg-blue-500 text-white rounded;
}

.card {
  @apply px-4 py-2 bg-white rounded shadow-md;
}

.alert {
  @apply px-4 py-2 bg-blue-50 rounded border;
}
```

**Problema:** Cada componente gera seu próprio CSS, mesmo que compartilhem classes. O PurgeCSS/JIT remove CSS não usado, mas ainda pode haver duplicação.

#### ✅ Solução: Componentes Base + Modificadores

```css
/* ✅ BOM - Componente base reutilizável */
.btn {
  @apply px-4 py-2 rounded;
}

.btn-primary {
  @apply bg-blue-500 text-white;
}

.card {
  @apply bg-white rounded shadow-md;
}

.alert {
  @apply bg-blue-50 rounded border;
}
```

**Vantagem:** Menos duplicação, melhor organização.

### Performance em Runtime

#### Especificidade e Cascata

Componentes com `@apply` podem ter especificidade diferente de utilitários diretos:

```css
/* Componente */
.btn {
  @apply bg-blue-500;
}

/* Utilitário direto */
<div class="bg-red-500 btn">Texto</div>
```

**Resultado:** A cor será azul (especificidade igual, mas ordem de CSS importa).

**Solução:** Use modificadores ou seja explícito:

```css
.btn {
  @apply px-4 py-2 rounded;
  /* Não defina cor aqui */
}

.btn-primary {
  @apply bg-blue-500 text-white;
}
```

---

## 📐 Boas Práticas de Organização

### Estrutura de Arquivos Recomendada

```
projeto/
├── src/
│   ├── styles/
│   │   ├── components/
│   │   │   ├── _buttons.css      <!-- Componentes de botão -->
│   │   │   ├── _cards.css        <!-- Componentes de card -->
│   │   │   ├── _forms.css        <!-- Componentes de formulário -->
│   │   │   ├── _navigation.css   <!-- Componentes de navegação -->
│   │   │   └── _index.css        <!-- Exporta todos -->
│   │   ├── utilities/
│   │   │   └── _custom.css       <!-- Utilitários customizados -->
│   │   └── main.css              <!-- Arquivo principal -->
```

### Convenção de Nomenclatura

#### ❌ Evite: Nomes Genéricos

```css
/* ❌ RUIM - Muito genérico */
.box { }
.container { }
.wrapper { }
```

**Problema:** Conflitos com classes do Tailwind ou outros componentes.

#### ✅ Prefira: Nomes Específicos e Semânticos

```css
/* ✅ BOM - Específico e semântico */
.user-card { }
.product-card { }
.dashboard-container { }
```

**Vantagem:** Evita conflitos, facilita manutenção.

#### Padrão BEM com Tailwind

```css
/* Componente */
.card { }

/* Elemento */
.card__header { }
.card__body { }
.card__footer { }

/* Modificador */
.card--elevated { }
.card--bordered { }
```

**Uso:**
```html
<div class="card card--elevated">
  <div class="card__header">Título</div>
  <div class="card__body">Conteúdo</div>
</div>
```

---

## 🎯 Quando Usar @apply vs Utilitários vs CSS Customizado

### Decisão: Utilitários Diretos

**Use quando:**
- Elemento é único ou pouco repetido
- Você está prototipando rapidamente
- Flexibilidade é mais importante que consistência
- O conjunto de classes é pequeno (< 5 classes)

**Exemplo:**
```html
<!-- ✅ BOM - Utilitários diretos -->
<div class="p-4 bg-white rounded shadow">
  Conteúdo único
</div>
```

### Decisão: @apply

**Use quando:**
- Padrão é repetido 3+ vezes
- Você quer garantir consistência
- Manutenção centralizada é importante
- O conjunto de classes é médio (5-10 classes)

**Exemplo:**
```css
/* ✅ BOM - @apply para padrão repetido */
.btn {
  @apply px-4 py-2 bg-blue-500 text-white rounded font-medium;
  @apply hover:bg-blue-600 transition-colors;
}
```

### Decisão: CSS Customizado

**Use quando:**
- Lógica CSS complexa (animações, pseudo-elementos)
- Propriedades não disponíveis no Tailwind
- Performance crítica (otimizações específicas)
- Integração com JavaScript (variáveis CSS dinâmicas)

**Exemplo:**
```css
/* ✅ BOM - CSS customizado para lógica complexa */
.modal-overlay {
  @apply fixed inset-0 bg-black bg-opacity-50;
  backdrop-filter: blur(4px);
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
```

---

## 🔧 Otimizações Específicas

### 1. Evite @apply Excessivo

#### ❌ Ruim: @apply em Tudo

```css
/* ❌ RUIM - @apply desnecessário */
.simple-text {
  @apply text-gray-600;
}
```

**Problema:** Cria componente desnecessário para uma única classe.

#### ✅ Bom: Use Utilitário Direto

```html
<!-- ✅ BOM - Utilitário direto -->
<p class="text-gray-600">Texto simples</p>
```

### 2. Combine @apply com CSS Customizado Quando Necessário

```css
/* ✅ BOM - Combinação inteligente */
.btn-gradient {
  @apply px-4 py-2 rounded font-medium text-white;
  /* CSS customizado para gradiente complexo */
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}
```

### 3. Use Variáveis CSS para Valores Dinâmicos

```css
/* ✅ BOM - Variáveis CSS para flexibilidade */
.card {
  @apply p-4 rounded-lg shadow-md;
  --card-bg: white;
  background-color: var(--card-bg);
}

.card-dark {
  --card-bg: #1f2937;
}
```

---

## 📊 Análise de Performance

### Ferramentas para Analisar Bundle Size

#### 1. PurgeCSS Analysis

```bash
# Ver quais classes estão sendo removidas
npx purgecss --css ./dist/output.css --content ./src/**/*.html --output ./analysis/
```

#### 2. Bundle Analyzer

```bash
# Analisar tamanho do CSS gerado
npm install -D @fullhuman/postcss-purgecss
```

#### 3. DevTools

- **Network tab**: Ver tamanho do CSS carregado
- **Coverage tab**: Ver CSS não utilizado
- **Performance tab**: Medir impacto de renderização

### Métricas Importantes

1. **Tamanho do CSS final** (após PurgeCSS/JIT)
   - Meta: < 50KB (comprimido)
   - Alerta: > 100KB

2. **Número de regras CSS**
   - Meta: < 1000 regras
   - Alerta: > 2000 regras

3. **Especificidade média**
   - Evite especificidade muito alta
   - Prefira classes simples

---

## 🛡️ Acessibilidade em Componentes

### Estados de Foco

```css
/* ✅ BOM - Foco acessível */
.btn {
  @apply px-4 py-2 bg-blue-500 text-white rounded;
  @apply focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2;
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

.btn:focus {
  outline: 2px solid transparent;
  outline-offset: 2px;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}
```

### Contraste de Cores

```css
/* ✅ BOM - Contraste adequado */
.alert-info {
  @apply bg-blue-50 text-blue-900; /* Alto contraste */
}

/* ❌ RUIM - Baixo contraste */
.alert-info {
  @apply bg-blue-100 text-blue-200; /* Contraste insuficiente */
}
```

### Estados Disabled

```css
/* ✅ BOM - Estado disabled claro */
.btn {
  @apply disabled:opacity-50 disabled:cursor-not-allowed;
  @apply disabled:hover:bg-blue-500; /* Previne hover quando disabled */
}
```

---

## 🔄 Manutenibilidade

### Documentação de Componentes

```css
/**
 * Botão primário
 * 
 * @component
 * @example
 * <button class="btn btn-primary">Clique aqui</button>
 * 
 * @variants
 * - .btn-primary - Botão primário azul
 * - .btn-secondary - Botão secundário cinza
 * - .btn-outline - Botão com borda
 */
.btn {
  @apply px-4 py-2 rounded font-medium;
}
```

### Versionamento de Componentes

```css
/* v1.0 - Versão inicial */
.btn-v1 {
  @apply px-4 py-2 bg-blue-500;
}

/* v2.0 - Nova versão (mantém v1 para compatibilidade) */
.btn {
  @apply px-4 py-2 bg-blue-600 rounded-lg;
}
```

### Testes de Componentes

```html
<!-- Teste de diferentes estados -->
<button class="btn btn-primary">Normal</button>
<button class="btn btn-primary" disabled>Disabled</button>
<button class="btn btn-primary" aria-pressed="true">Active</button>
```

---

## ⚠️ Armadilhas Comuns

### 1. Especificidade Conflitante

```css
/* ❌ PROBLEMA - Especificidade igual pode causar conflitos */
.btn {
  @apply bg-blue-500;
}

/* Em outro arquivo */
.btn-primary {
  @apply bg-red-500;
}
```

**Solução:** Use modificadores consistentes:

```css
/* ✅ BOM - Modificadores claros */
.btn {
  @apply px-4 py-2 rounded;
  /* Sem cor aqui */
}

.btn-primary {
  @apply bg-blue-500 text-white;
}

.btn-danger {
  @apply bg-red-500 text-white;
}
```

### 2. @apply com Classes Arbitrárias

```css
/* ❌ NÃO FUNCIONA */
.component {
  @apply [alguma-classe-arbitraria];
}
```

**Solução:** Use CSS customizado para valores arbitrários:

```css
/* ✅ BOM */
.component {
  @apply p-4 rounded;
  width: calc(100% - 2rem); /* CSS customizado */
}
```

### 3. Media Queries Diretas com @apply

```css
/* ❌ NÃO RECOMENDADO */
@media (min-width: 768px) {
  .component {
    @apply p-8;
  }
}
```

**Solução:** Use prefixos responsivos do Tailwind:

```css
/* ✅ BOM */
.component {
  @apply p-4 md:p-8;
}
```

---

## 🎓 Padrões Recomendados

### Padrão 1: Componente Base + Modificadores

```css
/* Base */
.card {
  @apply bg-white rounded-lg shadow-md p-6;
}

/* Modificadores */
.card-elevated {
  @apply shadow-lg hover:shadow-xl transition-shadow;
}

.card-bordered {
  @apply border border-gray-200;
}
```

### Padrão 2: Componente com Variantes

```css
/* Componente com variantes */
.alert {
  @apply p-4 rounded-lg border;
}

.alert-info {
  @apply bg-blue-50 border-blue-200 text-blue-800;
}

.alert-success {
  @apply bg-green-50 border-green-200 text-green-800;
}
```

### Padrão 3: Componente Composto

```css
/* Componentes que trabalham juntos */
.modal {
  @apply fixed inset-0 z-50 flex items-center justify-center;
}

.modal-overlay {
  @apply absolute inset-0 bg-black bg-opacity-50;
}

.modal-content {
  @apply relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full;
}
```

---

## 📈 Métricas de Sucesso

### Indicadores de Boa Implementação

1. **Reutilização**: Componentes usados em 3+ lugares
2. **Consistência**: Mesmo componente tem mesma aparência em todo projeto
3. **Manutenibilidade**: Mudanças em um lugar refletem em todos os usos
4. **Performance**: Bundle size dentro dos limites
5. **Acessibilidade**: Componentes seguem padrões WCAG

### Checklist de Revisão

Antes de criar um componente, pergunte:

- [ ] Este padrão é repetido 3+ vezes?
- [ ] O componente terá variantes ou modificadores?
- [ ] A manutenção centralizada é importante?
- [ ] O nome do componente é específico e não conflita?
- [ ] O componente é acessível (foco, contraste, estados)?
- [ ] A documentação está clara?

---

## 🚀 Resumo de Boas Práticas

### ✅ Faça

1. **Crie componentes para padrões repetidos** (3+ vezes)
2. **Use nomes específicos e semânticos** (evite genéricos)
3. **Organize componentes em arquivos separados**
4. **Documente componentes complexos**
5. **Combine @apply com CSS customizado quando necessário**
6. **Teste componentes em diferentes contextos**
7. **Mantenha especificidade baixa**
8. **Use modificadores para variantes**

### ❌ Evite

1. **Criar componentes para uso único**
2. **Usar @apply para classes simples** (use utilitário direto)
3. **Nomes genéricos** (`.box`, `.container`)
4. **Especificidade muito alta**
5. **@apply com classes arbitrárias**
6. **Media queries diretas com @apply** (use prefixos)
7. **Duplicação desnecessária de estilos**
8. **Componentes sem estados de foco**

---

## 🎯 Próximos Passos

Agora que você domina componentes e reutilização, na próxima aula você aprenderá sobre **Customização e Configuração do Tailwind**, onde poderá:
- Personalizar cores, espaçamento e breakpoints
- Adicionar utilitários customizados
- Configurar o tema do Tailwind
- Estender o sistema de design

Isso permitirá criar componentes ainda mais poderosos e alinhados com o design system do seu projeto!

