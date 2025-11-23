# Aula 3 - Exercícios e Reflexão: Position

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para consolidar seu entendimento sobre position em CSS. Eles vão desde situações básicas até problemas mais complexos que exigem raciocínio sobre posicionamento e layout.

---

## 📝 Exercício 1: Identificando o Position Correto

### Situação:

Você está criando um site e precisa decidir qual `position` usar para cada elemento abaixo. Para cada situação, identifique qual valor de position seria mais apropriado e explique por quê.

#### a) Um menu de navegação que deve permanecer visível no topo da página, mesmo quando o usuário rola para baixo.

**Qual position você usaria?**
- [ ] static
- [ ] relative
- [ ] absolute
- [ ] fixed
- [ ] sticky

**Por quê?**

---

#### b) Um badge pequeno com o texto "Novo!" que deve aparecer no canto superior direito de um card de produto.

**Qual position você usaria?**
- [ ] static
- [ ] relative
- [ ] absolute
- [ ] fixed
- [ ] sticky

**Por quê?**

---

#### c) Um botão que precisa estar ligeiramente deslocado (5 pixels para a direita e 3 pixels para baixo) para melhor alinhamento visual.

**Qual position você usaria?**
- [ ] static
- [ ] relative
- [ ] absolute
- [ ] fixed
- [ ] sticky

**Por quê?**

---

#### d) O cabeçalho de uma tabela longa que deve permanecer visível enquanto o usuário rola os dados da tabela, mas desaparecer quando a tabela sai da tela.

**Qual position você usaria?**
- [ ] static
- [ ] relative
- [ ] absolute
- [ ] fixed
- [ ] sticky

**Por quê?**

---

#### e) Um tooltip que aparece quando o usuário passa o mouse sobre um botão, posicionado logo acima do botão.

**Qual position você usaria?**
- [ ] static
- [ ] relative
- [ ] absolute
- [ ] fixed
- [ ] sticky

**Por quê?**

---

## 📝 Exercício 2: Analisando Código CSS

### Situação:

Você recebeu o seguinte código CSS de um colega. Analise o código e identifique possíveis problemas ou melhorias.

```css
/* Estilos para um card de produto */
.produto-card {
  position: relative;
  width: 300px;
  padding: 20px;
  background: white;
  border: 1px solid #ccc;
}

.produto-badge {
  position: absolute;
  top: -10px;
  right: -10px;
  background: red;
  color: white;
  padding: 5px 10px;
  z-index: 9999;
}

.produto-titulo {
  position: relative;
  top: 5px;
  left: 10px;
}

.produto-botao {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 10000;
}
```

### Perguntas:

1. **O badge do produto está usando `z-index: 9999`. Isso é necessário? Por quê?**

2. **O título do produto está usando `position: relative` com `top: 5px` e `left: 10px`. Isso é a melhor abordagem? Existe uma alternativa melhor?**

3. **O botão do produto está usando `position: fixed`. Em que situações isso pode causar problemas?**

4. **O card pai tem `position: relative`. Por que isso é importante para o badge?**

---

## 📝 Exercício 3: Criando um Layout com Position

### Situação:

Você precisa criar um card de produto com as seguintes características:
- O card tem 400px de largura e padding de 20px
- Um badge "Promoção" no canto superior direito (dentro do card, mas no canto)
- Uma imagem do produto no centro do card
- Um botão "Comprar" que deve estar sempre visível no canto inferior direito da tela (não do card, da tela toda)

### Tarefa:

Escreva o CSS necessário para criar este layout. Use position de forma apropriada para cada elemento.

**HTML de referência:**
```html
<div class="produto-card">
  <span class="badge-promocao">Promoção</span>
  <img src="produto.jpg" alt="Produto" class="produto-imagem">
  <h3 class="produto-titulo">Nome do Produto</h3>
  <p class="produto-preco">R$ 99,90</p>
</div>
<button class="botao-comprar">Comprar</button>
```

**Seu CSS:**

```css
/* Escreva seu CSS aqui */




```

### Perguntas de Reflexão:

1. **Por que você escolheu cada valor de position para cada elemento?**

2. **O botão "Comprar" está fora do card no HTML, mas deve aparecer fixo na tela. Isso pode causar algum problema de organização do código? Como você resolveria isso?**

3. **Se o card tivesse muitos produtos em uma lista, todos com botões fixed, o que aconteceria? Como você resolveria isso?**

---

## 📝 Exercício 4: Problema de Z-Index

### Situação:

Você tem três elementos que se sobrepõem:
- Um card de produto (background branco)
- Um modal que aparece sobre o card (background semi-transparente)
- Um botão de fechar no modal (pequeno X no canto superior direito)

### Problema:

O botão de fechar está aparecendo **atrás** do modal, mesmo que você tenha dado `z-index: 1000` a ele. O modal tem `z-index: 100`.

### Tarefa:

Explique por que isso está acontecendo e como você resolveria o problema.

**Sua explicação:**

---

**Sua solução:**

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Performance e Position

**Pergunta**: Elementos com `position: fixed` ou `position: sticky` podem ter impacto na performance da página, especialmente durante a rolagem. Por que isso acontece? Em que situações você deve ter cuidado ao usar esses valores?

**Sua resposta:**

---

### Reflexão 2: Responsividade e Position

**Pergunta**: Elementos com `position: absolute` ou `position: fixed` podem causar problemas em dispositivos móveis ou telas pequenas. Quais são esses problemas? Como você garantiria que esses elementos funcionem bem em diferentes tamanhos de tela?

**Sua resposta:**

---

### Reflexão 3: Acessibilidade e Position

**Pergunta**: Elementos posicionados de forma absoluta ou fixa podem causar problemas de acessibilidade. Por exemplo, um menu fixo pode cobrir conteúdo importante, ou um elemento absolute pode não aparecer na ordem correta para leitores de tela. Quais cuidados você deve ter ao usar position para garantir acessibilidade?

**Sua resposta:**

---

### Reflexão 4: Quando NÃO Usar Position

**Pergunta**: Muitos desenvolvedores iniciantes usam `position: absolute` ou `position: fixed` como solução para problemas de layout que poderiam ser resolvidos de forma mais simples. Em que situações você NÃO deveria usar position? Quais alternativas existem?

**Sua resposta:**

---

### Reflexão 5: Stacking Context e Z-Index

**Pergunta**: Você tem um card com `z-index: 10` e dentro dele um botão com `z-index: 100`. Mesmo assim, o botão não aparece na frente de elementos fora do card que têm `z-index: 5`. Por que isso acontece? O que é "stacking context" e como ele afeta o z-index?

**Sua resposta:**

---

## ✅ Checklist de Aprendizado

Antes de considerar que você dominou position, verifique se você consegue:

- [ ] Explicar a diferença entre static, relative, absolute, fixed e sticky
- [ ] Saber quando usar cada valor de position
- [ ] Entender como top, right, bottom e left funcionam com cada position
- [ ] Compreender o que é z-index e quando usá-lo
- [ ] Saber o que é um "ancestral posicionado" e por que é importante
- [ ] Entender o conceito de stacking context
- [ ] Identificar problemas comuns de position e como resolvê-los
- [ ] Considerar responsividade ao usar position
- [ ] Pensar em acessibilidade ao posicionar elementos
- [ ] Saber quando NÃO usar position (e usar alternativas como flexbox ou grid)

---

## 🎯 Próximos Passos

Após completar estes exercícios e reflexões, você deve:
1. Revisar os conceitos que ainda não estão claros
2. Praticar criando layouts simples usando diferentes valores de position
3. Testar seus layouts em diferentes tamanhos de tela
4. Experimentar combinar position com outras propriedades CSS (como transform para centralizar elementos)

Lembre-se: a prática é essencial para dominar position. Experimente, teste, e veja o que acontece!

