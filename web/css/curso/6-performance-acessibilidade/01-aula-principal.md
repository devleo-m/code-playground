# Aula 6: Performance e Acessibilidade em CSS

## 🎯 Introdução

Nesta aula, você aprenderá sobre dois pilares fundamentais do desenvolvimento web moderno: **Performance** e **Acessibilidade**. Esses conceitos vão além de apenas escrever CSS que funciona - eles garantem que suas páginas sejam rápidas, eficientes e utilizáveis por todos, independentemente de limitações físicas, tecnológicas ou de conexão.

---

## ⚡ Performance em CSS

### O que é Performance em CSS?

**Performance em CSS** refere-se à eficiência com que o código CSS é processado e renderizado pelo navegador, impactando diretamente a velocidade e a responsividade de um site. Quando falamos de performance, estamos preocupados com:

- Quão rápido o CSS é carregado
- Quão rápido o navegador consegue aplicar os estilos
- Quão suave é a renderização visual
- Quanto recurso do dispositivo é consumido

### Por que Performance é Importante?

Imagine que você está visitando um site e ele demora 5 segundos para carregar completamente. Agora imagine que você está usando um celular com conexão lenta. Esses 5 segundos podem se tornar 15 ou 20 segundos. Usuários não esperam - eles simplesmente vão embora.

**Performance importa porque:**
- Usuários abandonam sites lentos
- Google e outros buscadores penalizam sites lentos
- Dispositivos móveis têm recursos limitados
- Conexões podem ser instáveis ou lentas
- Experiência do usuário é diretamente afetada

### Como o Navegador Processa CSS?

Para entender performance, você precisa saber o que acontece quando o navegador encontra CSS:

1. **Download**: O navegador baixa o arquivo CSS
2. **Parsing**: O navegador lê e interpreta o código CSS
3. **Cascata e Especificidade**: O navegador resolve conflitos entre regras
4. **Renderização**: O navegador aplica os estilos aos elementos HTML
5. **Layout (Reflow)**: O navegador calcula posições e tamanhos
6. **Pintura (Repaint)**: O navegador desenha os elementos na tela

Cada uma dessas etapas consome tempo e recursos. CSS mal otimizado pode fazer com que essas etapas sejam mais lentas do que o necessário.

### Técnicas Básicas de Otimização

#### 1. Minimizar o Tamanho do Arquivo

Quanto menor o arquivo CSS, mais rápido ele será baixado. Algumas práticas simples:

- **Remover espaços desnecessários**: Espaços, quebras de linha e indentação aumentam o tamanho do arquivo
- **Remover código não utilizado**: Se você não está usando uma regra CSS, remova-a
- **Minificar**: Ferramentas automáticas podem reduzir drasticamente o tamanho

**Exemplo conceitual:**
```
Antes: .botao { background-color: blue; padding: 10px; margin: 5px; }
Depois: .botao{background-color:blue;padding:10px;margin:5px}
```

O segundo é menor e faz exatamente a mesma coisa.

#### 2. Reduzir Complexidade de Seletores

Seletores muito complexos fazem o navegador trabalhar mais para encontrar os elementos corretos. Seletores simples são mais rápidos.

**Seletores simples são melhores:**
- `.classe` é mais rápido que `.container .wrapper .content .item`
- `#id` é muito rápido
- `elemento` é rápido

**Seletores complexos são mais lentos:**
- Múltiplos níveis de descendência
- Muitos pseudo-seletores
- Seletores de atributo complexos

#### 3. Usar CSS Crítico

**CSS Crítico** é o CSS necessário para renderizar o conteúdo que o usuário vê primeiro (above the fold). Carregar apenas o essencial primeiro faz a página parecer carregar mais rápido.

A ideia é: em vez de carregar todo o CSS de uma vez, você carrega primeiro apenas o que é necessário para a primeira tela visível, e o resto pode carregar depois.

#### 4. Evitar Propriedades que Causam Reflow

Algumas propriedades CSS fazem o navegador recalcular o layout inteiro (reflow), o que é custoso:

- `width`, `height`
- `margin`, `padding`
- `position`
- `display`
- `font-size`

Propriedades que apenas mudam a aparência visual (repaint) são mais baratas:

- `color`
- `background-color`
- `box-shadow`
- `opacity`
- `transform`

#### 5. Usar Cache do Navegador

Quando um navegador baixa um arquivo CSS, ele pode guardá-lo em cache. Na próxima visita, em vez de baixar novamente, ele usa a versão guardada. Isso torna carregamentos subsequentes muito mais rápidos.

Para isso funcionar bem, você precisa:
- Configurar cabeçalhos HTTP apropriados no servidor
- Usar nomes de arquivo com versão (ex: `styles-v2.css`) quando fizer mudanças

### Reflow e Repaint

Esses são dois conceitos importantes para entender performance:

**Reflow (Layout):** Quando o navegador precisa recalcular as posições e tamanhos dos elementos. É custoso porque pode afetar muitos elementos.

**Repaint (Pintura):** Quando o navegador apenas precisa redesenhar elementos sem recalcular layout. É mais barato que reflow.

**Regra geral:** Evite causar reflows desnecessários. Se você precisa mudar muitas propriedades que causam reflow, tente agrupá-las ou usar `transform` e `opacity` quando possível, pois eles não causam reflow.

### Ferramentas para Analisar Performance

Navegadores modernos têm ferramentas (DevTools) que ajudam a analisar performance:

- **Performance Tab**: Mostra quanto tempo cada operação leva
- **Network Tab**: Mostra quanto tempo leva para baixar arquivos
- **Coverage Tab**: Mostra quanto CSS está sendo usado vs não usado

Essas ferramentas ajudam você a identificar problemas de performance no seu CSS.

---

## ♿ Acessibilidade em CSS

### O que é Acessibilidade?

**Acessibilidade** (também chamada de a11y) é a prática de criar conteúdo web que pode ser usado por todas as pessoas, independentemente de suas habilidades ou limitações. Isso inclui pessoas com:

- Deficiências visuais (cegueira, baixa visão, daltonismo)
- Deficiências auditivas
- Deficiências motoras
- Limitações cognitivas
- Dispositivos ou conexões limitadas

### Por que Acessibilidade é Importante?

Acessibilidade não é apenas uma questão legal ou ética - é uma questão de inclusão. A web deve ser para todos. Além disso:

- **Mais usuários**: Você alcança mais pessoas
- **Melhor SEO**: Sites acessíveis tendem a ter melhor posicionamento
- **Melhor experiência**: Práticas de acessibilidade melhoram a experiência para todos
- **Obrigação legal**: Em muitos países, é obrigatório por lei

### Como CSS Afeta Acessibilidade?

CSS não pode tornar conteúdo acessível sozinho - você precisa de HTML semântico também. Mas CSS pode:

- **Melhorar** a acessibilidade: Tornando conteúdo mais legível e navegável
- **Piorar** a acessibilidade: Escondendo conteúdo importante ou tornando difícil de usar

### Contraste de Cores

**Contraste** é a diferença entre a cor do texto e a cor do fundo. Texto com pouco contraste é difícil ou impossível de ler.

#### Por que Contraste Importa?

Pessoas com baixa visão ou daltonismo precisam de mais contraste para conseguir ler. Mesmo pessoas com visão normal podem ter dificuldade em ler texto com pouco contraste, especialmente em telas pequenas ou em ambientes com muita luz.

#### Níveis de Contraste (WCAG)

O WCAG (Web Content Accessibility Guidelines) define níveis mínimos de contraste:

- **AA (Mínimo)**: Texto normal precisa de contraste de pelo menos 4.5:1. Texto grande precisa de pelo menos 3:1.
- **AAA (Recomendado)**: Texto normal precisa de contraste de pelo menos 7:1. Texto grande precisa de pelo menos 4.5:1.

**Como verificar:** Existem ferramentas online que calculam o contraste entre duas cores. Use-as sempre que estiver escolhendo cores para texto.

#### Exemplos de Bom e Mau Contraste

**Bom contraste:**
- Texto preto (#000000) sobre fundo branco (#FFFFFF) = 21:1 (excelente)
- Texto azul escuro (#003366) sobre fundo branco (#FFFFFF) = 12.6:1 (excelente)

**Mau contraste:**
- Texto cinza claro (#CCCCCC) sobre fundo branco (#FFFFFF) = 1.6:1 (muito ruim)
- Texto amarelo (#FFFF00) sobre fundo branco (#FFFFFF) = 1.07:1 (impossível de ler)

### Tamanho de Fonte e Legibilidade

Texto muito pequeno é difícil de ler para todos, especialmente em dispositivos móveis ou para pessoas com dificuldades visuais.

#### Boas Práticas

- **Tamanho mínimo**: Use pelo menos 16px para texto do corpo (body text)
- **Unidades relativas**: Prefira `rem` ou `em` em vez de `px` fixos, pois respeitam as preferências do usuário
- **Line-height**: Use pelo menos 1.5 para texto do corpo, facilitando a leitura
- **Largura de linha**: Limite a largura de linhas de texto (idealmente entre 50-75 caracteres)

### Estados de Foco

**Foco** é o indicador visual de qual elemento está ativo quando alguém navega usando teclado (Tab). É crucial para pessoas que não usam mouse.

#### Por que Foco é Importante?

Sem indicadores de foco visíveis, usuários que navegam com teclado não sabem onde estão na página. Isso torna o site completamente inutilizável para eles.

#### Boas Práticas de Foco

- **Sempre torne o foco visível**: Não remova o outline padrão sem substituir por algo melhor
- **Foco claro e visível**: O indicador de foco deve ter contraste suficiente
- **Foco consistente**: Todos os elementos interativos devem ter foco visível

**Exemplo básico:**
```css
/* Não faça isso - remove o foco */
a:focus {
  outline: none;
}

/* Faça isso - melhora o foco */
a:focus {
  outline: 3px solid blue;
  outline-offset: 2px;
}
```

### Redução de Movimento

Algumas pessoas são sensíveis a movimento e animações podem causar desconforto, tontura ou até mesmo náusea.

#### Prefers-Reduced-Motion

CSS oferece uma media query especial chamada `prefers-reduced-motion` que detecta se o usuário prefere menos movimento:

```css
/* Animação normal */
.elemento {
  transition: transform 0.3s;
}

/* Respeitar preferência do usuário */
@media (prefers-reduced-motion: reduce) {
  .elemento {
    transition: none;
  }
}
```

**Sempre respeite essa preferência** quando criar animações ou transições.

### Leitores de Tela e CSS

**Leitores de tela** são tecnologias assistivas que leem o conteúdo da página em voz alta para pessoas cegas ou com baixa visão.

#### Como CSS Afeta Leitores de Tela?

CSS pode esconder conteúdo visualmente, mas isso não significa que o leitor de tela não vai ler. Algumas práticas importantes:

- **`display: none`** e **`visibility: hidden`**: Escondem do leitor de tela também
- **`opacity: 0`** ou posicionamento fora da tela: Esconde visualmente, mas o leitor de tela ainda lê

**Regra importante:** Se você esconde algo visualmente mas quer que leitores de tela leiam, use técnicas específicas. Se você esconde algo e não quer que leitores de tela leiam, use `display: none`.

### Navegação por Teclado

CSS pode melhorar a experiência de navegação por teclado:

- **Ordem lógica**: A ordem visual dos elementos deve corresponder à ordem de navegação por teclado
- **Foco visível**: Como já mencionado, foco deve ser sempre visível
- **Áreas clicáveis grandes**: Botões e links devem ter tamanho suficiente para serem fáceis de clicar/toque

### Responsividade e Acessibilidade

Layouts responsivos são parte da acessibilidade porque:

- Pessoas podem usar diferentes tamanhos de tela
- Zoom do navegador pode mudar o tamanho efetivo
- Dispositivos assistivos podem ter telas pequenas

Um layout que funciona bem em diferentes tamanhos é mais acessível.

---

## 🔗 Performance e Acessibilidade Trabalhando Juntos

Performance e acessibilidade não são opostos - na verdade, elas se complementam:

- **Sites rápidos são mais acessíveis**: Pessoas com conexões lentas ou dispositivos limitados se beneficiam de performance
- **CSS otimizado é mais fácil de processar**: Tecnologias assistivas também se beneficiam
- **Boas práticas se sobrepõem**: Código limpo e organizado é bom para ambos

### Exemplo Prático

Imagine um botão:

**Versão ruim (lenta e inacessível):**
- CSS complexo e verboso
- Contraste baixo
- Sem foco visível
- Tamanho pequeno

**Versão boa (rápida e acessível):**
- CSS simples e otimizado
- Contraste adequado
- Foco claro e visível
- Tamanho adequado para toque

A segunda versão é melhor em todos os aspectos.

---

## 📝 Resumo dos Conceitos Principais

### Performance
- **Objetivo**: Tornar CSS rápido de carregar e processar
- **Técnicas**: Minimizar arquivos, simplificar seletores, evitar reflows desnecessários
- **Ferramentas**: DevTools do navegador para análise

### Acessibilidade
- **Objetivo**: Tornar conteúdo utilizável por todos
- **Aspectos**: Contraste, tamanho de fonte, foco, movimento, leitores de tela
- **Padrões**: Seguir WCAG (níveis AA ou AAA)

### Boas Práticas Gerais
- Sempre pense em performance ao escrever CSS
- Sempre pense em acessibilidade ao criar estilos
- Teste em diferentes dispositivos e condições
- Use ferramentas para verificar contraste e performance
- Mantenha código simples e limpo

---

## 🎓 Próximos Passos

Agora que você entende os conceitos fundamentais de performance e acessibilidade, na próxima parte da aula você verá versões simplificadas com analogias do dia a dia, exercícios práticos e boas práticas avançadas. Esses conceitos devem ser parte do seu processo de desenvolvimento desde o início - não são algo para adicionar depois.


