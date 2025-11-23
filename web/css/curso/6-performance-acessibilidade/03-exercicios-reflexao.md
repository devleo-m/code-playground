# Aula 6 - Exercícios e Reflexão: Performance e Acessibilidade

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para você praticar e refletir sobre os conceitos de performance e acessibilidade em CSS. Lembre-se: o foco não é escrever código complexo, mas entender **por que** essas práticas são importantes e **como** aplicá-las no seu dia a dia.

---

## 📝 Exercício 1: Análise de Contraste

### Contexto
Você está criando um site e precisa escolher cores para texto e fundo. Sua tarefa é analisar se as combinações de cores têm contraste adequado.

### Tarefa
Para cada par de cores abaixo, identifique se o contraste é adequado ou não, e explique o porquê:

1. **Texto:** `#FFFFFF` (branco) | **Fundo:** `#000000` (preto)
2. **Texto:** `#CCCCCC` (cinza claro) | **Fundo:** `#FFFFFF` (branco)
3. **Texto:** `#333333` (cinza escuro) | **Fundo:** `#FFFFFF` (branco)
4. **Texto:** `#FFFF00` (amarelo) | **Fundo:** `#FFFFFF` (branco)
5. **Texto:** `#0066CC` (azul) | **Fundo:** `#FFFFFF` (branco)

### Dica
Lembre-se: contraste adequado significa que a diferença entre as cores é grande o suficiente para ser legível. Texto claro sobre fundo claro ou texto escuro sobre fundo escuro geralmente tem pouco contraste.

### Reflexão
- Por que contraste é importante para acessibilidade?
- O que acontece quando alguém com baixa visão tenta ler texto com pouco contraste?
- Como você pode verificar o contraste de cores antes de usar em produção?

---

## 📝 Exercício 2: Identificando Problemas de Performance

### Contexto
Você recebeu um arquivo CSS de um colega e precisa identificar possíveis problemas de performance.

### Tarefa
Analise os seguintes trechos de CSS e identifique o que pode ser otimizado:

**Trecho 1:**
```css
.container .wrapper .content .sidebar .menu .item .link {
  color: blue;
  text-decoration: none;
}
```

**Trecho 2:**
```css
.botao {
  width: 100px;
  height: 50px;
  background-color: blue;
  color: white;
  padding: 10px;
  margin: 5px;
  border: 1px solid black;
  border-radius: 5px;
  box-shadow: 2px 2px 5px rgba(0,0,0,0.3);
  font-size: 16px;
  font-weight: bold;
  text-align: center;
  line-height: 50px;
}
```

**Trecho 3:**
```css
/* Muitas regras não utilizadas */
.estilo-antigo-1 { color: red; }
.estilo-antigo-2 { color: blue; }
.estilo-antigo-3 { color: green; }
.estilo-antigo-4 { color: yellow; }
/* ... mais 50 regras similares que não são mais usadas */
```

### Reflexão
- Por que seletores muito complexos podem ser problemáticos para performance?
- O que acontece quando você tem muito CSS não utilizado no seu arquivo?
- Como você pode identificar CSS não utilizado no seu projeto?

---

## 📝 Exercício 3: Criando Foco Acessível

### Contexto
Você precisa criar estilos para links e botões que sejam acessíveis para navegação por teclado.

### Tarefa
Escreva CSS para um link que:
1. Tenha uma aparência normal quando não está em foco
2. Tenha um indicador de foco claro e visível quando está em foco
3. O indicador de foco deve ter contraste adequado

**Dica:** Use a pseudo-classe `:focus` e pense em como tornar o foco visível sem ser intrusivo.

### Reflexão
- Por que é importante ter foco visível em elementos interativos?
- O que acontece quando alguém tenta navegar seu site usando apenas o teclado e não há indicadores de foco?
- Qual é a diferença entre remover o outline padrão e melhorá-lo?

---

## 📝 Exercício 4: Respeitando Preferências de Movimento

### Contexto
Você criou uma animação de fade-in para elementos que aparecem na página, mas precisa garantir que pessoas sensíveis a movimento não sejam afetadas.

### Tarefa
Escreva CSS que:
1. Aplique uma transição suave de opacidade quando um elemento aparece
2. Respeite a preferência do usuário por movimento reduzido
3. Quando o usuário preferir movimento reduzido, a transição deve ser removida ou muito mais rápida

**Dica:** Use a media query `prefers-reduced-motion`.

### Reflexão
- Por que algumas pessoas preferem movimento reduzido?
- O que pode acontecer se você não respeitar essa preferência?
- Como você pode criar animações que sejam agradáveis para todos?

---

## 📝 Exercício 5: Otimizando Seletores

### Contexto
Você tem seletores CSS complexos e precisa simplificá-los para melhorar a performance.

### Tarefa
Simplifique os seguintes seletores, mantendo a mesma funcionalidade:

1. `div.container div.wrapper div.content p.texto`
2. `body > div > section > article > h2.titulo`
3. `ul.lista li.item span.texto a.link`

**Dica:** Pense em como você pode usar classes diretamente em vez de depender de hierarquia complexa.

### Reflexão
- Por que seletores simples são melhores para performance?
- Como você pode estruturar seu HTML para permitir seletores mais simples?
- Qual é o equilíbrio entre especificidade e simplicidade?

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Performance e Experiência do Usuário

**Pergunta:** Imagine que você está visitando um site em um celular com conexão lenta. O site demora 10 segundos para carregar completamente. Como você se sente? O que você faz?

**Pense sobre:**
- Qual é a relação entre performance e experiência do usuário?
- Por que sites lentos perdem usuários?
- Como performance afeta diferentes tipos de usuários (com conexões rápidas vs lentas)?

### Reflexão 2: Acessibilidade como Direito

**Pergunta:** Acessibilidade é apenas uma "boa prática" ou é um direito fundamental? Por quê?

**Pense sobre:**
- Por que acessibilidade é importante além de questões legais?
- Como você se sentiria se não conseguisse usar um site por causa de limitações de acessibilidade?
- Qual é o impacto de criar sites não acessíveis na sociedade?

### Reflexão 3: Performance vs Funcionalidade

**Pergunta:** É possível ter um site com muitas funcionalidades E ainda assim ser rápido? Como?

**Pense sobre:**
- Qual é o equilíbrio entre adicionar funcionalidades e manter performance?
- Como você pode priorizar o que é essencial vs o que é "legal de ter"?
- Quais técnicas você pode usar para carregar conteúdo não essencial depois?

### Reflexão 4: Acessibilidade Beneficia Todos

**Pergunta:** Muitas pessoas pensam que acessibilidade beneficia apenas pessoas com deficiência. Isso é verdade? Por quê?

**Pense sobre:**
- Como contraste adequado beneficia pessoas sem deficiência visual?
- Como navegação por teclado pode ser útil para pessoas que não têm deficiência?
- Quais são os benefícios "ocultos" de criar sites acessíveis?

### Reflexão 5: CSS e Responsabilidade Social

**Pergunta:** Como desenvolvedor, qual é sua responsabilidade em relação a performance e acessibilidade?

**Pense sobre:**
- Você é responsável apenas por fazer o código funcionar, ou também por garantir que funcione bem para todos?
- Como você pode incorporar performance e acessibilidade no seu processo de desenvolvimento?
- O que você pode fazer hoje para melhorar seus projetos existentes?

---

## 📚 Dicas para Resolver os Exercícios

1. **Não se preocupe com código perfeito**: O objetivo é entender os conceitos, não escrever código perfeito
2. **Pense no usuário**: Sempre considere como suas decisões afetam os usuários finais
3. **Use ferramentas**: Existem ferramentas online para verificar contraste, analisar performance, etc.
4. **Teste na prática**: Se possível, teste suas soluções em diferentes dispositivos e condições
5. **Reflita honestamente**: As perguntas de reflexão não têm resposta certa ou errada - o importante é pensar sobre elas

---

## ✅ Checklist de Aprendizado

Após completar os exercícios, você deve ser capaz de:

- [ ] Identificar problemas de contraste em combinações de cores
- [ ] Reconhecer seletores CSS que podem ser otimizados
- [ ] Criar estilos de foco acessíveis
- [ ] Respeitar preferências de movimento do usuário
- [ ] Simplificar seletores complexos
- [ ] Refletir sobre o impacto de performance e acessibilidade
- [ ] Entender a importância desses conceitos para o desenvolvimento web

---

## 🎓 Próximos Passos

Após completar estes exercícios e reflexões, você terá uma base sólida sobre performance e acessibilidade. Na próxima parte da aula, você verá boas práticas avançadas e técnicas de otimização que você pode aplicar imediatamente nos seus projetos.

Lembre-se: performance e acessibilidade não são "extras" - são fundamentais para criar sites que funcionam bem para todos!

