# Aula 1 - Exercícios e Reflexão: CSS Basics

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para consolidar seu aprendizado sobre CSS Basics. Eles vão desde o básico até desafios que combinam múltiplos conceitos. Faça cada exercício com calma e pense sobre o que está fazendo.

---

## 📝 Exercício 1: Criando Sua Primeira Regra CSS

### Tarefa:
Crie um arquivo CSS externo chamado `estilos.css` e escreva regras CSS para:

1. Fazer todos os títulos `<h1>` ficarem azuis
2. Fazer todos os parágrafos terem tamanho de fonte de 18 pixels
3. Fazer todos os elementos com classe `destaque` terem fundo amarelo

### HTML de Referência:
```html
<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="estilos.css">
</head>
<body>
  <h1>Título Principal</h1>
  <p>Este é um parágrafo normal.</p>
  <p class="destaque">Este parágrafo deve ter fundo amarelo.</p>
  <h1>Outro Título</h1>
</body>
</html>
```

### O que você deve escrever no `estilos.css`:
(Escreva suas regras CSS aqui)

---

## 📝 Exercício 2: Entendendo Seletores

### Tarefa:
Analise o HTML abaixo e escreva seletores CSS para cada situação:

```html
<div id="container">
  <h1 class="titulo-principal">Título</h1>
  <p class="texto">Parágrafo 1</p>
  <p class="texto destaque">Parágrafo 2 com destaque</p>
  <div class="caixa">
    <p>Parágrafo dentro da caixa</p>
  </div>
</div>
```

**Escreva seletores para:**

1. Estilizar apenas o elemento com ID `container`
2. Estilizar todos os elementos com classe `texto`
3. Estilizar apenas parágrafos que estão dentro de elementos com classe `caixa`
4. Estilizar apenas parágrafos que são filhos diretos de `container`
5. Estilizar o parágrafo que vem logo após o `h1`

---

## 📝 Exercício 3: Propriedades de Texto

### Tarefa:
Crie regras CSS para estilizar o texto conforme as especificações:

1. **Títulos h2:**
   - Cor: verde escuro
   - Tamanho da fonte: 24 pixels
   - Alinhamento: centralizado
   - Transformação: todas as letras maiúsculas

2. **Links:**
   - Cor: azul
   - Sem decoração (sem sublinhado)
   - Estilo de fonte: itálico

3. **Parágrafos:**
   - Tamanho da fonte: 16 pixels
   - Altura da linha: 1.5
   - Alinhamento: justificado

**Escreva suas regras CSS:**

---

## 📝 Exercício 4: Combinando Conceitos

### Tarefa:
Você tem o seguinte HTML:

```html
<article class="post">
  <h1 class="titulo">Meu Artigo</h1>
  <p class="introducao">Esta é a introdução do artigo.</p>
  <p>Este é um parágrafo normal do artigo.</p>
  <p class="conclusao">Esta é a conclusão.</p>
</article>

<article class="post">
  <h1 class="titulo">Outro Artigo</h1>
  <p class="introducao">Outra introdução.</p>
</article>
```

**Crie regras CSS que:**

1. Todos os artigos tenham largura máxima de 800 pixels
2. Todos os títulos dentro de artigos sejam azuis
3. Parágrafos com classe `introducao` tenham fundo cinza claro
4. Parágrafos com classe `conclusao` tenham borda superior
5. O primeiro parágrafo após cada título tenha margem superior de 20 pixels

**Escreva suas regras CSS:**

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Métodos de Aplicação de CSS

**Pergunta:** Você está criando um site com 50 páginas. Todas as páginas precisam ter o mesmo cabeçalho azul, rodapé cinza e links vermelhos. Qual método de aplicação de CSS você escolheria e por quê?

**Pense sobre:**
- Qual método seria mais eficiente?
- Como você manteria a consistência visual?
- O que aconteceria se você precisasse mudar a cor do cabeçalho em todas as páginas?
- Qual método facilitaria o trabalho em equipe?

**Sua resposta:**

---

### Reflexão 2: Especificidade e Cascata

**Cenário:** Você tem um parágrafo que tem:
- Uma classe `texto` que define cor azul
- Um ID `destaque` que define cor verde
- Um estilo inline que define cor vermelha

**Pergunta:** Qual cor será aplicada ao parágrafo e por quê? Se você precisasse que fosse azul, como faria?

**Pense sobre:**
- Como a cascata funciona neste caso?
- Qual tem maior especificidade: classe, ID ou inline?
- Quando seria apropriado usar `!important`?
- Qual seria a melhor prática para evitar conflitos?

**Sua resposta:**

---

### Reflexão 3: Seletores e Manutenibilidade

**Cenário:** Você escreveu este CSS:

```css
div div div p {
  color: red;
}
```

**Pergunta:** Este seletor é uma boa prática? Por quê? Quais problemas ele pode causar? Como você melhoraria?

**Pense sobre:**
- Este seletor é muito específico ou muito genérico?
- O que acontece se a estrutura HTML mudar?
- Como isso afeta a performance?
- Qual seria uma alternativa melhor?

**Sua resposta:**

---

### Reflexão 4: Opacidade vs RGBA

**Pergunta:** Qual é a diferença entre usar `opacity: 0.5` e `background-color: rgba(255, 0, 0, 0.5)`? Quando você usaria cada um?

**Pense sobre:**
- O que cada um afeta?
- Como isso impacta elementos filhos?
- Qual oferece mais controle?
- Quando cada abordagem é mais apropriada?

**Sua resposta:**

---

## 🎯 Desafio Final: Criando um Estilo Completo

### Tarefa:
Crie um arquivo CSS completo que estilize uma página de blog simples. Use todos os conceitos aprendidos.

**Requisitos:**

1. **Cabeçalho (h1):**
   - Cor: azul escuro (#003366)
   - Tamanho: 32 pixels
   - Centralizado
   - Transformação: maiúsculas

2. **Parágrafos:**
   - Tamanho: 16 pixels
   - Cor: cinza escuro (#333333)
   - Altura da linha: 1.6
   - Alinhamento: justificado

3. **Links:**
   - Cor: azul (#0066cc)
   - Sem sublinhado
   - Quando passar o mouse (hover): sublinhado e cor mais escura

4. **Elementos com classe `destaque`:**
   - Fundo: amarelo claro (#fff9c4)
   - Padding: 10 pixels
   - Borda esquerda: 4 pixels sólida azul

5. **Citações (blockquote):**
   - Fonte: itálico
   - Cor: cinza (#666666)
   - Margem esquerda: 20 pixels
   - Borda esquerda: 3 pixels sólida cinza

**Escreva seu CSS completo:**

---

## 📚 Dicas para Resolver os Exercícios

1. **Leia com atenção:** Certifique-se de entender o que cada exercício pede
2. **Teste seu código:** Abra o HTML no navegador e veja se funciona
3. **Use comentários:** Adicione comentários explicando o que cada regra faz
4. **Experimente:** Não tenha medo de tentar diferentes valores
5. **Consulte a documentação:** Se esquecer uma propriedade, pesquise

---

## ✅ Checklist de Aprendizado

Após completar os exercícios, verifique se você consegue:

- [ ] Criar regras CSS com seletores básicos
- [ ] Diferenciar entre seletores de elemento, classe e ID
- [ ] Usar combinadores para selecionar elementos relacionados
- [ ] Aplicar propriedades de texto básicas
- [ ] Entender a ordem de cascata
- [ ] Escolher o método de aplicação de CSS apropriado
- [ ] Explicar quando usar cada tipo de seletor
- [ ] Criar estilos reutilizáveis com classes

---

## 🎓 Próximos Passos

Após completar estes exercícios e reflexões, você estará pronto para:
- Aprender sobre cores e backgrounds em detalhes
- Entender o modelo de caixa (box model)
- Trabalhar com layouts básicos

Lembre-se: a prática é essencial. Quanto mais você experimentar, mais confiança terá!

