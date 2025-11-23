# Aula 7 - Exercícios e Reflexão: Ferramentas e Metodologias CSS

## 🎯 Objetivo dos Exercícios

Estes exercícios foram criados para consolidar seu aprendizado sobre ferramentas e metodologias CSS. Eles focam em **entendimento conceitual** e **decisões práticas** sobre quando e por que usar cada ferramenta. A maioria dos exercícios não requer código, mas sim **raciocínio e reflexão**.

---

## 📝 Exercício 1: Identificando Problemas que Cada Ferramenta Resolve

### Tarefa:
Para cada situação abaixo, identifique qual ferramenta ou metodologia (Sass, PostCSS, BEM, CSS Modules, CSS-in-JS) resolveria melhor o problema. Explique **por quê**.

### Situações:

1. **Situação A:** Você está trabalhando em um projeto grande e percebe que a cor azul `#0066cc` aparece em 47 lugares diferentes do CSS. O cliente pediu para mudar para `#0055bb`. Você terá que encontrar e substituir manualmente em 47 lugares.

   **Qual ferramenta resolveria?** ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

2. **Situação B:** Você escreveu CSS moderno usando `display: grid`, mas precisa que funcione em navegadores antigos que não suportam essa propriedade sem prefixos.

   **Qual ferramenta resolveria?** ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

3. **Situação C:** Você está trabalhando em equipe e percebe que diferentes desenvolvedores estão usando nomes de classes diferentes para a mesma coisa: `.botao`, `.btn`, `.button`, `.botao-principal`. Isso está causando confusão.

   **Qual ferramenta resolveria?** ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

4. **Situação D:** Você tem dois componentes React diferentes, ambos usando a classe `.card`. Quando você estiliza um, o outro também é afetado porque ambos compartilham o mesmo nome de classe global.

   **Qual ferramenta resolveria?** ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

5. **Situação E:** Você precisa que um botão mude de cor dinamicamente baseado em uma prop do React. Se a prop `status` for "ativo", o botão fica verde. Se for "inativo", fica cinza.

   **Qual ferramenta resolveria?** ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

## 📝 Exercício 2: Entendendo BEM na Prática

### Tarefa:
Analise o HTML abaixo e crie nomes de classes seguindo a metodologia BEM. Identifique:
- Quais são os **blocos**
- Quais são os **elementos** (partes dos blocos)
- Quais são os **modificadores** (variações)

### HTML de Referência:
```html
<article class="???">
  <img class="???" src="foto.jpg" alt="Produto">
  <h2 class="???">Nome do Produto</h2>
  <p class="???">Descrição do produto</p>
  <span class="???">R$ 99,90</span>
  <button class="???">Comprar</button>
</article>

<article class="??? ???">
  <img class="???" src="foto2.jpg" alt="Produto">
  <h2 class="???">Outro Produto</h2>
  <p class="???">Este produto está em promoção</p>
  <span class="???">R$ 79,90</span>
  <button class="??? ???">Comprar Agora</button>
</article>
```

### Sua Tarefa:
1. Identifique o **bloco principal**: ________________

2. Identifique os **elementos** (partes do bloco):
   - Imagem: ________________
   - Título: ________________
   - Descrição: ________________
   - Preço: ________________
   - Botão: ________________

3. Identifique os **modificadores**:
   - Artigo em promoção: ________________
   - Botão destacado: ________________

4. **Agora escreva os nomes de classes BEM completos** no HTML acima, substituindo os `???`

---

## 📝 Exercício 3: Escolhendo a Ferramenta Certa

### Tarefa:
Para cada projeto descrito abaixo, identifique quais ferramentas/metodologias você usaria e explique sua escolha.

### Projetos:

1. **Projeto A:** Site institucional simples, 5 páginas HTML estáticas, sem JavaScript frameworks. Equipe de 2 pessoas.

   **Ferramentas que eu usaria:**
   - ________________
   - ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

2. **Projeto B:** Aplicação React grande, 50+ componentes, estilos que mudam baseado em estado e props, tema que muda dinamicamente.

   **Ferramentas que eu usaria:**
   - ________________
   - ________________
   - ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

3. **Projeto C:** Site WordPress com tema customizado, CSS grande (5000+ linhas), múltiplos desenvolvedores trabalhando, precisa funcionar em navegadores antigos.

   **Ferramentas que eu usaria:**
   - ________________
   - ________________
   - ________________
   
   **Por quê?** 
   (Escreva sua resposta aqui)

---

## 📝 Exercício 4: Analisando Código Existente

### Tarefa:
Analise o seguinte código CSS e identifique problemas que poderiam ser resolvidos com as ferramentas que você aprendeu.

### Código para Análise:
```css
/* Arquivo: estilos.css */

.botao {
  background-color: #0066cc;
  padding: 10px 20px;
  border-radius: 5px;
}

.botao-grande {
  background-color: #0066cc;
  padding: 15px 30px;
  border-radius: 5px;
}

.card {
  background-color: white;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.card-titulo {
  color: #0066cc;
  font-size: 24px;
}

.card-texto {
  color: #333;
  font-size: 16px;
}

.produto-card {
  background-color: white;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.produto-card-titulo {
  color: #0066cc;
  font-size: 24px;
}
```

### Perguntas:

1. **Problema de Repetição:**
   - Que valor aparece repetidamente? ________________
   - Qual ferramenta resolveria isso? ________________
   - Como? (Explique brevemente)

---

2. **Problema de Organização:**
   - Que padrão você nota nos nomes das classes? (está organizado ou desorganizado?)
   - Qual metodologia melhoraria isso? ________________
   - Como os nomes ficariam seguindo essa metodologia?

---

3. **Problema de Duplicação:**
   - Que estilos estão duplicados entre `.card` e `.produto-card`?
   - Como você evitaria essa duplicação? (Qual ferramenta?)

---

## 🤔 Perguntas de Reflexão

### Reflexão 1: Quando NÃO Usar Ferramentas

**Pergunta:** Em que situações você **NÃO** usaria Sass, CSS Modules ou CSS-in-JS? Por quê?

**Sua resposta:**
(Escreva sua reflexão aqui. Pense em projetos pequenos, simplicidade, performance, preferências da equipe, etc.)

---

### Reflexão 2: Trade-offs (Compensações)

**Pergunta:** Toda ferramenta tem vantagens e desvantagens. Pense em:

- **Sass:** Adiciona complexidade ao processo de build, mas facilita escrita e manutenção
- **CSS-in-JS:** Oferece estilos dinâmicos, mas adiciona overhead de JavaScript
- **CSS Modules:** Isola estilos, mas requer build tools

**Qual dessas compensações você acha mais aceitável? Por quê?**

**Sua resposta:**
(Escreva sua reflexão aqui)

---

### Reflexão 3: Evolução do CSS

**Pergunta:** CSS puro está evoluindo e adicionando recursos que antes só existiam em pré-processadores (como variáveis CSS nativas). 

**Pense:** Se CSS puro continuar evoluindo e adicionar mais recursos (mixins, funções, etc.), Sass ainda será necessário? Ou CSS puro será suficiente?

**Sua resposta:**
(Escreva sua reflexão aqui. Considere: compatibilidade, velocidade de adoção de novos recursos, necessidade de compilação, etc.)

---

### Reflexão 4: Trabalho em Equipe

**Pergunta:** Você está em uma equipe de 5 desenvolvedores trabalhando no mesmo projeto CSS grande.

**Cenário A:** Cada desenvolvedor nomeia classes como prefere (sem padrão)
**Cenário B:** Todos seguem BEM rigorosamente
**Cenário C:** Projeto usa CSS Modules, então nomes não importam tanto

**Qual cenário você prefere? Por quê? Quais são os prós e contras de cada um?**

**Sua resposta:**
(Escreva sua reflexão aqui)

---

### Reflexão 5: Performance e Complexidade

**Pergunta:** CSS-in-JS gera estilos em tempo de execução (runtime), enquanto CSS tradicional e Sass são compilados antes (build time).

**Pense:**
- Quais são as implicações de performance de cada abordagem?
- Em que situações a performance seria mais crítica?
- O ganho em desenvolvimento (estilos dinâmicos, escopo automático) compensa o custo de performance?

**Sua resposta:**
(Escreva sua reflexão aqui)

---

## 📊 Exercício Final: Criando uma Estratégia

### Tarefa:
Imagine que você está começando um novo projeto. Crie uma estratégia de CSS respondendo às perguntas abaixo.

### Informações do Projeto:
- **Tipo:** Aplicação web moderna
- **Tecnologia:** React
- **Tamanho:** Médio (20-30 componentes)
- **Equipe:** 3 desenvolvedores
- **Prazo:** 6 meses
- **Manutenção:** Projeto será mantido por anos

### Suas Decisões:

1. **Qual abordagem de CSS você escolheria?**
   - [ ] CSS tradicional + BEM
   - [ ] CSS Modules
   - [ ] CSS-in-JS (styled-components ou Emotion)
   - [ ] Sass + BEM
   - [ ] Outra: ________________

2. **Você usaria PostCSS?** [ ] Sim [ ] Não
   
   **Por quê?** 

3. **Você usaria alguma metodologia de nomenclatura?** [ ] Sim [ ] Não
   
   **Qual?** ________________

4. **Justifique suas escolhas:**
   (Explique por que você escolheu essa combinação de ferramentas para este projeto específico)

---

## ✅ Checklist de Aprendizado

Antes de avançar, verifique se você consegue:

- [ ] Explicar o que é Sass e quando usá-lo
- [ ] Explicar o que é PostCSS e sua função principal
- [ ] Aplicar nomenclatura BEM corretamente
- [ ] Entender a diferença entre CSS Modules e CSS tradicional
- [ ] Explicar quando CSS-in-JS é apropriado
- [ ] Escolher a ferramenta certa para diferentes tipos de projetos
- [ ] Identificar problemas que cada ferramenta resolve
- [ ] Entender os trade-offs (compensações) de cada abordagem

---

## 💡 Dica Final

Lembre-se: **Não existe uma ferramenta "melhor"** - existe a ferramenta certa para cada situação. O importante é entender:
- **O problema** que você está tentando resolver
- **As opções** disponíveis
- **Os trade-offs** de cada opção
- **O contexto** do seu projeto

Com esse conhecimento, você estará preparado para tomar decisões informadas sobre qual ferramenta usar em cada situação!

