# Aula 1 - Exercícios Práticos e Reflexão

Bem-vindo aos exercícios! Esta é a etapa mais importante do aprendizado. Aqui você vai colocar em prática tudo que aprendeu e refletir sobre os conceitos.

**⚠️ IMPORTANTE:** 
- Complete TODOS os exercícios
- Responda as perguntas de reflexão com suas próprias palavras
- Não copie respostas das aulas anteriores
- Seja honesto sobre o que você entendeu ou não

---

## 📝 Exercício 1: Criando seu Primeiro JavaScript

### Objetivo
Criar uma página HTML simples com JavaScript que exibe uma mensagem personalizada.

### Tarefa
1. Crie um arquivo HTML chamado `exercicio-01.html`
2. Adicione um título `<h1>` com o texto "Meu Primeiro Exercício"
3. Adicione um parágrafo vazio com o ID `mensagem`
4. Use JavaScript inline para:
   - Selecionar o parágrafo pelo ID
   - Adicionar o texto "Olá! Este é meu primeiro código JavaScript!"
   - Exibir no console: "Exercício 1 concluído!"

### Código Base
```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Exercício 1</title>
</head>
<body>
    <!-- Seu código aqui -->
</body>
</html>
```

### Verificação
- [ ] A página exibe o título
- [ ] O parágrafo mostra a mensagem
- [ ] O console exibe a mensagem de conclusão
- [ ] O código está organizado e comentado

---

## 📝 Exercício 2: JavaScript Externo

### Objetivo
Aprender a separar JavaScript em arquivo externo e reutilizá-lo.

### Tarefa
1. Crie um arquivo HTML chamado `exercicio-02.html`
2. Crie um arquivo JavaScript chamado `exercicio-02.js`
3. No HTML:
   - Adicione um botão com o texto "Clique em mim"
   - Inclua o arquivo JavaScript externo
4. No JavaScript (`exercicio-02.js`):
   - Adicione um event listener ao botão
   - Quando clicado, o botão deve alterar sua cor de fundo para azul
   - Exiba no console: "Botão clicado com sucesso!"

### Dica
Use `document.querySelector('button')` ou `document.getElementsByTagName('button')[0]` para selecionar o botão.

### Verificação
- [ ] O JavaScript está em arquivo separado
- [ ] O botão muda de cor ao ser clicado
- [ ] O console exibe a mensagem
- [ ] O código funciona corretamente

---

## 📝 Exercício 3: Trabalhando com o Console

### Objetivo
Praticar o uso do console do navegador para testar código.

### Tarefa
Abra o console do navegador (F12) e execute os seguintes comandos, um por vez:

1. **Variáveis básicas:**
   ```javascript
   let nome = "Seu Nome";
   let idade = 25;
   console.log("Nome:", nome);
   console.log("Idade:", idade);
   ```

2. **Cálculos:**
   ```javascript
   let numero1 = 10;
   let numero2 = 5;
   let soma = numero1 + numero2;
   let multiplicacao = numero1 * numero2;
   console.log("Soma:", soma);
   console.log("Multiplicação:", multiplicacao);
   ```

3. **Manipulação da página:**
   ```javascript
   document.body.style.backgroundColor = "lightgreen";
   document.title = "Console em Ação!";
   console.log("Página modificada!");
   ```

### Perguntas
Responda em um arquivo de texto ou no papel:

a) O que aconteceu quando você executou cada bloco de código?
b) Por que o console é útil para testar código?
c) Qual a diferença entre `console.log()` e apenas digitar uma expressão no console?

### Verificação
- [ ] Executei todos os comandos no console
- [ ] Entendi o que cada comando faz
- [ ] Respondi as perguntas sobre o console

---

## 📝 Exercício 4: Análise de Código

### Objetivo
Analisar código JavaScript existente e identificar problemas e melhorias.

### Código para Análise
```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Análise de Código</title>
</head>
<body>
    <h1>Análise de Código</h1>
    <button id="botao1">Botão 1</button>
    <button id="botao2">Botão 2</button>
    
    <script>
        document.getElementById('botao1').onclick = function() {
            alert('Botão 1 clicado!');
        }
        
        document.getElementById('botao2').onclick = function() {
            alert('Botão 2 clicado!');
        }
        
        console.log('Script carregado');
    </script>
</body>
</html>
```

### Tarefa
Analise o código acima e responda:

1. **O código funciona?** Sim ou não? Por quê?

2. **Identifique problemas:**
   - Há algum problema de organização?
   - O código poderia ser melhorado?
   - Está seguindo boas práticas?

3. **Sugestões de melhoria:**
   - Como você melhoraria este código?
   - O que você mudaria e por quê?

4. **Teste o código:**
   - Crie o arquivo HTML
   - Teste no navegador
   - Verifique se funciona como esperado
   - Anote qualquer comportamento inesperado

### Verificação
- [ ] Analisei o código cuidadosamente
- [ ] Identifiquei problemas e melhorias
- [ ] Testei o código no navegador
- [ ] Documentei minhas observações

---

## 🤔 Perguntas de Reflexão

Estas perguntas são **cruciais** para seu aprendizado. Elas exigem que você pense profundamente sobre os conceitos, não apenas memorize. Responda cada uma com suas próprias palavras, sendo honesto sobre seu entendimento.

---

### 🔍 Reflexão 1: Por que JavaScript é importante?

**Pergunta:**
JavaScript é uma das linguagens mais populares do mundo. Pense sobre isso:

a) **Por que você acha que JavaScript se tornou tão popular?**
   - Considere: facilidade de uso, versatilidade, onde roda, etc.

b) **Qual o impacto de JavaScript na experiência do usuário em um site?**
   - Pense em sites com e sem JavaScript. Qual a diferença?

c) **Se JavaScript não existisse, como seria a web hoje?**
   - Reflita sobre a importância da interatividade na web moderna.

**Instruções:**
- Escreva pelo menos 3-4 frases para cada item
- Use exemplos concretos se possível
- Seja honesto sobre seu entendimento atual

---

### 🔍 Reflexão 2: Escolhendo o ambiente de execução

**Pergunta:**
JavaScript pode rodar em diferentes ambientes (navegador, Node.js, mobile, etc.).

a) **Quando você escolheria usar JavaScript no navegador vs Node.js?**
   - Dê exemplos específicos de quando cada um é mais apropriado.

b) **Quais são as limitações de segurança do JavaScript no navegador?**
   - Por que o navegador limita o que JavaScript pode fazer?
   - Isso é bom ou ruim? Por quê?

c) **Como a escolha do ambiente afeta o que você pode fazer com JavaScript?**
   - Pense em funcionalidades que só funcionam em um ambiente específico.

**Instruções:**
- Pense em cenários reais
- Considere as implicações práticas
- Escreva de forma clara e organizada

---

### 🔍 Reflexão 3: Evolução e compatibilidade

**Pergunta:**
JavaScript evoluiu muito desde sua criação, com várias versões (ES5, ES6, ES2017, etc.).

a) **Por que é importante entender as diferentes versões do JavaScript?**
   - Considere: compatibilidade com navegadores antigos, novos recursos, etc.

b) **Quais são os desafios de ter múltiplas versões do JavaScript?**
   - Pense em desenvolvedores que precisam suportar navegadores antigos.
   - Como isso afeta o desenvolvimento?

c) **Como você decidiria qual versão do JavaScript usar em um projeto?**
   - Quais fatores você consideraria?
   - Dê exemplos de quando usar ES5 vs ES6+.

**Instruções:**
- Reflita sobre implicações práticas
- Considere diferentes cenários de projeto
- Seja específico em seus exemplos

---

### 🔍 Reflexão 4: Debugging e ferramentas

**Pergunta:**
O console do navegador é uma ferramenta fundamental para desenvolvimento JavaScript.

a) **Por que o console é tão importante para aprender JavaScript?**
   - Pense em: teste rápido, ver erros, experimentação, etc.

b) **Quais são as limitações de testar código apenas no console?**
   - Quando o console não é suficiente?
   - Quando você precisa de arquivos HTML/JS reais?

c) **Como você usaria o console para debugar um problema em seu código?**
   - Descreva um processo passo a passo.
   - Dê um exemplo de como você investigaria um erro.

**Instruções:**
- Pense em situações práticas
- Considere diferentes tipos de problemas
- Seja específico em suas estratégias

---

### 🔍 Reflexão 5: Organização e boas práticas

**Pergunta:**
Existem diferentes formas de incluir JavaScript em uma página (inline, externo, console).

a) **Quando você usaria JavaScript inline vs arquivo externo?**
   - Dê exemplos específicos de cada caso.
   - Quais são as vantagens e desvantagens de cada abordagem?

b) **Por que a organização do código é importante?**
   - Pense em: manutenção, reutilização, trabalho em equipe, etc.
   - Qual o impacto de código mal organizado?

c) **Como você organizaria o JavaScript de um site grande?**
   - Considere: múltiplas páginas, funcionalidades diferentes, etc.
   - Descreva uma estrutura que faça sentido.

**Instruções:**
- Pense em projetos reais
- Considere escalabilidade
- Seja prático e específico

---

## 📊 Checklist de Conclusão

Antes de enviar suas respostas, verifique:

### Exercícios Práticos
- [ ] Exercício 1: Criei o arquivo HTML com JavaScript inline
- [ ] Exercício 2: Criei arquivos HTML e JS separados
- [ ] Exercício 3: Executei comandos no console
- [ ] Exercício 4: Analisei o código fornecido

### Perguntas de Reflexão
- [ ] Reflexão 1: Respondi sobre a importância do JavaScript
- [ ] Reflexão 2: Respondi sobre ambientes de execução
- [ ] Reflexão 3: Respondi sobre versões e compatibilidade
- [ ] Reflexão 4: Respondi sobre debugging e console
- [ ] Reflexão 5: Respondi sobre organização e boas práticas

### Qualidade das Respostas
- [ ] Usei minhas próprias palavras (não copiei das aulas)
- [ ] Fui honesto sobre o que entendi ou não
- [ ] Incluí exemplos concretos quando possível
- [ ] Revisei minhas respostas antes de enviar

---

## 📤 Enviando suas Respostas

Quando terminar todos os exercícios e reflexões:

1. **Organize suas respostas:**
   - Crie um arquivo de texto ou documento
   - Inclua todos os códigos que você escreveu
   - Inclua todas as respostas de reflexão

2. **Inclua os arquivos:**
   - `exercicio-01.html`
   - `exercicio-02.html`
   - `exercicio-02.js`
   - Qualquer outro arquivo que você criou

3. **Envie para análise:**
   - O tutor analisará seu código
   - Você receberá feedback construtivo
   - Identificaremos pontos fortes e áreas de melhoria

---

## ⏱️ Tempo Estimado

- **Exercícios 1-2:** 30-45 minutos
- **Exercício 3:** 15-20 minutos
- **Exercício 4:** 20-30 minutos
- **Perguntas de Reflexão:** 45-60 minutos

**Total:** 110-155 minutos (aproximadamente 2-2.5 horas)

---

## 💪 Dicas para Sucesso

1. **Não tenha pressa:** Dedique tempo de qualidade a cada exercício
2. **Teste tudo:** Sempre teste seu código no navegador
3. **Use o console:** É seu melhor amigo para debugar
4. **Seja honesto:** Nas reflexões, seja sincero sobre seu entendimento
5. **Pergunte:** Se tiver dúvidas, anote-as para discutir depois

---

## 🚀 Próximo Passo

Após completar todos os exercícios e reflexões, você estará pronto para a próxima etapa:

**Arquivo seguinte**: `04-performance-boas-praticas.md`

Mas **NÃO avance** até completar todos os exercícios e enviar suas respostas para análise!

Boa sorte! 🎓

