# Aula 6 - Exercícios e Reflexão: Estados e Interatividade

## 🎯 Objetivos dos Exercícios

Ao completar estes exercícios, você será capaz de:
- Criar elementos interativos usando estados do Tailwind
- Aplicar transições suaves em diferentes propriedades
- Usar group e peer para criar interações complexas
- Implementar feedback visual adequado para diferentes estados
- Pensar criticamente sobre acessibilidade e experiência do usuário

---

## 📝 Exercício 1: Botão Interativo Básico

### Tarefa

Crie um botão que tenha os seguintes estados:

1. **Normal**: Fundo azul (`bg-blue-500`), texto branco, padding adequado, bordas arredondadas
2. **Hover**: Fundo azul mais escuro, cresce ligeiramente (`scale-105`), sombra aumenta
3. **Active**: Fundo ainda mais escuro, encolhe ligeiramente (`scale-95`)
4. **Focus**: Anel azul ao redor (`ring-2 ring-blue-500`)
5. **Disabled**: Opacidade reduzida, cursor `not-allowed`

### Requisitos

- Use transições suaves em todas as mudanças
- O botão deve ter feedback visual claro em cada estado
- O texto do botão deve ser "Clique em mim"

### Código Base

```html
<!-- Seu código aqui -->
<button>
  Clique em mim
</button>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```html
<button class="
  bg-blue-500
  hover:bg-blue-600
  active:bg-blue-700
  focus:outline-none
  focus:ring-2
  focus:ring-blue-500
  disabled:opacity-50
  disabled:cursor-not-allowed
  text-white
  font-semibold
  px-6 py-3
  rounded-lg
  shadow-md
  hover:shadow-lg
  hover:scale-105
  active:scale-95
  transition-all
  duration-200
  ease-in-out
">
  Clique em mim
</button>
```

</details>

---

## 📝 Exercício 2: Input com Label Flutuante

### Tarefa

Crie um input com label flutuante usando `peer`:

1. **Estado inicial**: O label está dentro do input, no mesmo nível do texto
2. **Quando o input recebe focus ou tem valor**: O label sobe e fica menor
3. **Focus no input**: Borda azul e anel de foco
4. **Hover no input**: Borda muda ligeiramente de cor

### Requisitos

- Use `peer` para fazer o label reagir ao estado do input
- O label deve ter transição suave ao mover
- O input deve ter placeholder vazio (`placeholder=" "`)
- Use `peer-placeholder-shown:` para controlar quando o label está na posição inicial

### Código Base

```html
<div class="relative">
  <input 
    type="text"
    id="email"
    placeholder=" "
  />
  <label for="email">
    Email
  </label>
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```html
<div class="relative">
  <input 
    type="text"
    id="email"
    class="
      peer
      w-full
      px-4
      pt-6
      pb-2
      border-2
      border-gray-300
      rounded-lg
      focus:outline-none
      focus:border-blue-500
      focus:ring-2
      focus:ring-blue-200
      hover:border-gray-400
      transition-all
    "
    placeholder=" "
  />
  <label 
    for="email"
    class="
      absolute
      left-4
      top-4
      text-gray-500
      transition-all
      pointer-events-none
      peer-focus:text-blue-500
      peer-focus:top-2
      peer-focus:text-sm
      peer-placeholder-shown:top-4
      peer-placeholder-shown:text-base
    "
  >
    Email
  </label>
</div>
```

</details>

---

## 📝 Exercício 3: Card com Efeito Group

### Tarefa

Crie um card que usa `group` para criar uma interação onde:

1. **Normal**: Card com sombra média, elementos em estado padrão
2. **Hover no card**: 
   - Sombra aumenta
   - Card sobe ligeiramente
   - Título muda de cor para azul
   - Texto fica mais escuro
   - Botão "Ver mais" aparece (estava invisível)
   - Ícone muda de cor

### Requisitos

- Use `group` no card principal
- Use `group-hover:` nos elementos filhos
- Todos os elementos devem ter transições suaves
- O botão deve aparecer com fade-in (`opacity-0` → `opacity-100`)

### Código Base

```html
<div>
  <div>Ícone</div>
  <h3>Título do Card</h3>
  <p>Descrição do card que muda quando você passa o mouse.</p>
  <button>Ver mais</button>
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```html
<div class="
  group
  bg-white
  p-6
  rounded-lg
  shadow-md
  cursor-pointer
  hover:shadow-xl
  hover:-translate-y-2
  transition-all
  duration-300
">
  <div class="
    w-12
    h-12
    bg-blue-500
    rounded-full
    group-hover:bg-blue-600
    transition-colors
  "></div>
  
  <h3 class="
    mt-4
    text-xl
    font-bold
    text-gray-800
    group-hover:text-blue-600
    transition-colors
  ">
    Título do Card
  </h3>
  
  <p class="
    mt-2
    text-gray-600
    group-hover:text-gray-800
    transition-colors
  ">
    Descrição do card que muda quando você passa o mouse.
  </p>
  
  <button class="
    mt-4
    opacity-0
    group-hover:opacity-100
    bg-blue-500
    hover:bg-blue-600
    text-white
    px-4
    py-2
    rounded
    transition-all
    duration-300
  ">
    Ver mais
  </button>
</div>
```

</details>

---

## 📝 Exercício 4: Lista com Estilos Alternados

### Tarefa

Crie uma lista de itens onde:

1. **Itens ímpares**: Fundo cinza claro
2. **Itens pares**: Fundo branco
3. **Hover em qualquer item**: Fundo azul claro, texto mais escuro
4. **Primeiro item**: Sem padding no topo, bordas arredondadas no topo
5. **Último item**: Sem padding embaixo, bordas arredondadas embaixo

### Requisitos

- Use `odd:` e `even:` para cores alternadas
- Use `first:` e `last:` para estilizar extremos
- Todos os itens devem ter hover
- Use `divide-y` ou bordas para separar visualmente

### Código Base

```html
<ul>
  <li>Item 1</li>
  <li>Item 2</li>
  <li>Item 3</li>
  <li>Item 4</li>
  <li>Item 5</li>
</ul>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```html
<ul class="divide-y divide-gray-200">
  <li class="
    px-4
    py-3
    odd:bg-gray-50
    even:bg-white
    hover:bg-blue-50
    hover:text-gray-900
    first:rounded-t-lg
    first:pt-4
    last:rounded-b-lg
    last:pb-4
    transition-colors
    cursor-pointer
  ">
    Item 1
  </li>
  <li class="
    px-4
    py-3
    odd:bg-gray-50
    even:bg-white
    hover:bg-blue-50
    hover:text-gray-900
    first:rounded-t-lg
    first:pt-4
    last:rounded-b-lg
    last:pb-4
    transition-colors
    cursor-pointer
  ">
    Item 2
  </li>
  <li class="
    px-4
    py-3
    odd:bg-gray-50
    even:bg-white
    hover:bg-blue-50
    hover:text-gray-900
    first:rounded-t-lg
    first:pt-4
    last:rounded-b-lg
    last:pb-4
    transition-colors
    cursor-pointer
  ">
    Item 3
  </li>
  <li class="
    px-4
    py-3
    odd:bg-gray-50
    even:bg-white
    hover:bg-blue-50
    hover:text-gray-900
    first:rounded-t-lg
    first:pt-4
    last:rounded-b-lg
    last:pb-4
    transition-colors
    cursor-pointer
  ">
    Item 4
  </li>
  <li class="
    px-4
    py-3
    odd:bg-gray-50
    even:bg-white
    hover:bg-blue-50
    hover:text-gray-900
    first:rounded-t-lg
    first:pt-4
    last:rounded-b-lg
    last:pb-4
    transition-colors
    cursor-pointer
  ">
    Item 5
  </li>
</ul>
```

**Nota**: Uma solução mais elegante seria usar uma classe compartilhada ou componente, mas para fins didáticos, mostramos a estrutura completa.

</details>

---

## 📝 Exercício 5: Toggle Switch com Peer

### Tarefa

Crie um toggle switch (interruptor) usando `peer`:

1. **Estado desligado**: Fundo cinza, círculo à esquerda
2. **Estado ligado**: Fundo azul, círculo à direita
3. **Focus**: Anel de foco visível
4. **Hover**: Ligeira mudança visual

### Requisitos

- Use um checkbox escondido (`sr-only`) como `peer`
- Use `peer-checked:` para estilizar quando marcado
- O círculo deve se mover suavemente
- Use `after:` para criar o círculo deslizante

### Código Base

```html
<label>
  <input type="checkbox" />
  <div>Toggle</div>
  <span>Toggle me</span>
</label>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```html
<label class="relative inline-flex items-center cursor-pointer">
  <input 
    type="checkbox" 
    class="sr-only peer" 
  />
  <div class="
    w-11
    h-6
    bg-gray-200
    peer-focus:outline-none
    peer-focus:ring-4
    peer-focus:ring-blue-300
    rounded-full
    peer
    peer-checked:after:translate-x-full
    peer-checked:after:border-white
    after:content-['']
    after:absolute
    after:top-[2px]
    after:left-[2px]
    after:bg-white
    after:border-gray-300
    after:border
    after:rounded-full
    after:h-5
    after:w-5
    after:transition-all
    peer-checked:bg-blue-600
  "></div>
  <span class="ml-3 text-sm font-medium text-gray-900">
    Toggle me
  </span>
</label>
```

</details>

---

## 📝 Exercício 6: Menu Dropdown com Animações

### Tarefa

Crie um menu dropdown que:

1. **Estado inicial**: Menu invisível e acima da posição final
2. **Hover no botão**: Menu aparece com fade-in e desce suavemente
3. **Hover nos itens**: Cada item tem background hover
4. **Primeiro e último item**: Bordas arredondadas apropriadas

### Requisitos

- Use `group` no container do botão
- Menu deve ter `opacity-0`, `invisible`, e `-translate-y-2` inicialmente
- No `group-hover:`, menu deve aparecer e descer
- Itens devem ter hover individual

### Código Base

```html
<div>
  <button>Menu</button>
  <div>
    <a href="#">Item 1</a>
    <a href="#">Item 2</a>
    <a href="#">Item 3</a>
  </div>
</div>
```

### Solução Esperada

<details>
<summary>Clique para ver a solução</summary>

```html
<div class="relative group">
  <button class="
    px-4
    py-2
    bg-gray-100
    rounded-lg
    hover:bg-gray-200
    focus:outline-none
    focus:ring-2
    focus:ring-blue-500
    transition-colors
  ">
    Menu
  </button>
  <div class="
    absolute
    top-full
    left-0
    mt-2
    w-48
    bg-white
    rounded-lg
    shadow-lg
    opacity-0
    invisible
    group-hover:opacity-100
    group-hover:visible
    group-hover:translate-y-0
    -translate-y-2
    transition-all
    duration-200
    z-10
  ">
    <a 
      href="#" 
      class="
        block
        px-4
        py-2
        hover:bg-gray-100
        first:rounded-t-lg
        last:rounded-b-lg
        transition-colors
      "
    >
      Item 1
    </a>
    <a 
      href="#" 
      class="
        block
        px-4
        py-2
        hover:bg-gray-100
        first:rounded-t-lg
        last:rounded-b-lg
        transition-colors
      "
    >
      Item 2
    </a>
    <a 
      href="#" 
      class="
        block
        px-4
        py-2
        hover:bg-gray-100
        first:rounded-t-lg
        last:rounded-b-lg
        transition-colors
      "
    >
      Item 3
    </a>
  </div>
</div>
```

</details>

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Performance e Transições

**Pergunta**: Você criou um botão com `transition-all duration-300`. Analise:

1. **Quais são as vantagens de usar `transition-all`?**
2. **Quais são as desvantagens de usar `transition-all`?**
3. **Em que situações seria melhor usar `transition-colors` ou `transition-transform` separadamente?**
4. **Como isso impacta a performance da página, especialmente em dispositivos móveis?**

**Pense sobre:**
- Reflow e repaint no navegador
- Uso de GPU para animações
- Propriedades que causam reflow vs propriedades que não causam
- Quando animar múltiplas propriedades simultaneamente é necessário vs desnecessário

---

### Reflexão 2: Acessibilidade e Estados de Focus

**Pergunta**: Você criou um input com `focus:ring-2 focus:ring-blue-500`. Analise:

1. **Por que é importante ter um estado de focus visível?**
2. **Qual a diferença entre `focus:` e `focus-visible:`? Quando usar cada um?**
3. **Como usuários que navegam apenas com teclado se beneficiariam de estados de focus bem implementados?**
4. **O que acontece se você remover completamente os estilos de focus? Qual o impacto na acessibilidade?**

**Pense sobre:**
- Navegação por teclado (Tab, Shift+Tab)
- Leitores de tela e acessibilidade
- Diretrizes WCAG sobre foco visível
- Diferença entre interação por mouse vs teclado

---

### Reflexão 3: Group vs Peer - Quando Usar Cada Um?

**Pergunta**: Você usou tanto `group` quanto `peer` nos exercícios. Analise:

1. **Qual a diferença fundamental entre `group` e `peer`?**
2. **Em que situações `group` é mais apropriado?**
3. **Em que situações `peer` é mais apropriado?**
4. **Você conseguiria recriar o efeito do `peer` usando apenas `group`? E vice-versa?**
5. **Qual abordagem é mais semântica e manutenível?**

**Pense sobre:**
- Relação pai-filho (group) vs relação irmão-irmão (peer)
- Estrutura HTML e semântica
- Manutenibilidade do código
- Performance de renderização

---

### Reflexão 4: Estados Combinados e Especificidade

**Pergunta**: Você criou elementos com múltiplos estados (`hover:`, `focus:`, `active:`). Analise:

1. **O que acontece quando um elemento está em múltiplos estados simultaneamente? (ex: hover + focus)**
2. **Como o Tailwind resolve conflitos entre estados?**
3. **Qual a ordem de especificidade dos estados no Tailwind?**
4. **Como você garantiria que um estado específico sempre tenha prioridade?**

**Pense sobre:**
- Cascata CSS e especificidade
- Ordem das classes no HTML
- Estados que podem coexistir vs estados mutuamente exclusivos
- Como testar todos os cenários possíveis

---

### Reflexão 5: Mobile e Touch Devices

**Pergunta**: Você criou interações baseadas principalmente em hover. Analise:

1. **Como hover funciona em dispositivos touch (tablets, smartphones)?**
2. **Quais estados são mais apropriados para dispositivos móveis?**
3. **Como você adaptaria suas interações para funcionar bem tanto em desktop quanto mobile?**
4. **O que acontece quando um usuário toca em um elemento com `hover:` em um dispositivo touch?**

**Pense sobre:**
- Diferenças entre mouse e touch
- Estados que funcionam universalmente vs estados específicos de dispositivo
- Responsividade de interações, não apenas de layout
- Testes em diferentes dispositivos

---

### Reflexão 6: Feedback Visual e UX

**Pergunta**: Você criou vários elementos com feedback visual. Analise:

1. **Qual a importância do feedback visual imediato em interfaces?**
2. **Como você determina se um feedback visual é suficiente, insuficiente ou excessivo?**
3. **Quais são os riscos de ter muito feedback visual? E de ter pouco?**
4. **Como você garantiria que usuários com deficiência visual também recebam feedback adequado?**

**Pense sobre:**
- Princípios de design de interação
- Feedback háptico, visual e auditivo
- Acessibilidade além do visual
- Consistência de feedback em toda a aplicação

---

### Reflexão 7: Tailwind vs CSS Puro para Estados

**Pergunta**: Compare criar estados com Tailwind vs CSS puro. Analise:

1. **Quais são as vantagens de usar Tailwind para estados?**
2. **Quais são as desvantagens?**
3. **Em que situações você preferiria escrever CSS puro para estados complexos?**
4. **Como você decidiria entre usar `@apply` para criar um componente vs usar classes utilitárias diretamente?**

**Pense sobre:**
- Produtividade vs controle
- Manutenibilidade
- Reutilização
- Complexidade de estados muito específicos
- Performance de bundle

---

## 📊 Checklist de Aprendizado

Marque os itens que você conseguiu completar:

### Conceitos Fundamentais
- [ ] Entendo o que são estados e pseudo-classes
- [ ] Sei usar `hover:` para criar interações no mouse
- [ ] Sei usar `focus:` e suas variantes (`focus-visible:`, `focus-within:`)
- [ ] Sei usar `active:` para feedback no clique
- [ ] Sei usar `disabled:` e `required:` para estados de formulário
- [ ] Sei usar `first:`, `last:`, `odd:`, `even:` para estilizar por posição

### Conceitos Avançados
- [ ] Entendo e uso `group` para interações em elementos pais
- [ ] Entendo e uso `peer` para interações entre irmãos
- [ ] Sei criar transições suaves com `transition-*`
- [ ] Sei controlar duração com `duration-*`
- [ ] Sei usar timing functions (`ease-*`)
- [ ] Sei aplicar transforms (`scale`, `rotate`, `translate`, `skew`)

### Boas Práticas
- [ ] Sempre forneço estados de focus visíveis
- [ ] Uso transições apropriadas (não `transition-all` quando desnecessário)
- [ ] Considero acessibilidade em minhas interações
- [ ] Testo em diferentes dispositivos (mouse e touch)
- [ ] Penso sobre performance ao criar animações

---

## 🎯 Próximos Passos

Após completar estes exercícios e reflexões:

1. **Revise suas respostas** às perguntas de reflexão
2. **Teste seus componentes** em diferentes navegadores e dispositivos
3. **Peça feedback** de outros desenvolvedores sobre acessibilidade
4. **Experimente combinações** de estados que não foram cobertas nos exercícios
5. **Analise sites reais** e identifique como eles implementam estados e interatividade

---

**Lembre-se**: Estados e interatividade não são apenas sobre fazer coisas "bonitas". São sobre criar uma experiência de usuário clara, acessível e agradável. Sempre pense no "porquê" além do "como".

