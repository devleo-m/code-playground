# **Aula 4 - Simplificada: Entendendo o Modelo Relacional na Prática**

## 🎯 Vamos aprofundar o modelo relacional com analogias!

---

## 🎨 1. Domains = Moldes Personalizados

### A Fábrica de Formulários

Imagine que você tem uma **fábrica que produz formulários**. Você precisa de campos de email em 100 formulários diferentes.

#### ❌ Sem Domains (Modo Artesanal)

Você desenha manualmente o campo de email em cada um dos 100 formulários:

```
Formulário 1: [_____________] ← Desenhar regras: "precisa ter @"
Formulário 2: [_____________] ← Desenhar regras: "precisa ter @"
Formulário 3: [_____________] ← Desenhar regras: "precisa ter @"
... (97 vezes mais!) 😫
```

**Problema:** Se quiser mudar as regras (ex: "agora precisa ter .com"), precisa refazer 100 formulários!

#### ✅ Com Domains (Modo Molde)

Você cria um **molde de campo de email** e carimba em todos os formulários:

```
🎨 MOLDE: "Campo Email"
   Regras: precisa ter @, precisa ter ponto, etc.

Formulário 1: [usa molde] ✅
Formulário 2: [usa molde] ✅
Formulário 3: [usa molde] ✅
... (97 vezes mais, mas instantâneo!) 😊
```

**Vantagem:** Se mudar o molde, todos os formulários mudam automaticamente!

---

### Exemplos de Moldes (Domains)

#### Molde de CPF Brasileiro 🇧🇷

```
MOLDE: cpf_brasileiro
Formato: XXX.XXX.XXX-XX (onde X é número)

✅ 123.456.789-00  (segue o molde)
❌ 12345678900     (não tem pontos e traço)
❌ ABC.DEF.GHI-JK (não são números)
```

#### Molde de CEP Brasileiro 📮

```
MOLDE: cep_brasileiro
Formato: XXXXX-XXX

✅ 01234-567
❌ 01234567  (falta traço)
❌ 1234-567  (falta um número)
```

#### Molde de Nota Escolar 📝

```
MOLDE: nota_escolar
Regra: Número entre 0 e 10

✅ 7.5
✅ 10.0
✅ 0.0
❌ 11.0  (acima de 10)
❌ -2.0  (negativo)
```

---

## 🏷️ 2. Attributes = Etiquetas em uma Caixa

### A Caixa Organizada

Imagine que **cada linha do banco de dados é uma caixa** de produto.

**Attributes (atributos)** são as **etiquetas** coladas na caixa:

```
📦 CAIXA DE PRODUTO (uma linha/tupla)
├─ 🏷️ Etiqueta "Código": 12345
├─ 🏷️ Etiqueta "Nome": Mouse Gamer
├─ 🏷️ Etiqueta "Preço": R$ 150,00
└─ 🏷️ Etiqueta "Peso": 0.3kg
```

**Cada etiqueta (atributo) tem:**

- 📌 **Nome**: "Preço", "Peso", "Nome"
- 📏 **Tipo**: Número, texto, data, etc.
- ✅ **Regras**: Obrigatório? Único? Positivo?

---

### Etiquetas Obrigatórias vs Opcionais

```
📦 CAIXA DE CLIENTE
├─ 🏷️ Nome: [OBRIGATÓRIO] ✅
├─ 🏷️ Email: [OBRIGATÓRIO] ✅
└─ 🏷️ Telefone: [OPCIONAL] ⚪ (pode estar em branco)
```

**No PostgreSQL:**

```sql
nome VARCHAR(100) NOT NULL,      -- Etiqueta obrigatória
telefone VARCHAR(20)             -- Etiqueta opcional (pode ser NULL)
```

---

## 📄 3. Tuples = Fichas Preenchidas

### O Fichário

Imagine um **fichário de biblioteca** onde cada ficha representa um livro.

**Tuple (tupla)** é uma **ficha completa** com todos os campos preenchidos:

```
┌─────────────────────────────────┐
│  FICHA DE LIVRO #1              │
├─────────────────────────────────┤
│  ISBN: 978-3-16-148410-0       │
│  Título: 1984                   │
│  Autor: George Orwell           │
│  Ano: 1949                      │
│  Preço: R$ 35,90                │
└─────────────────────────────────┘
```

Esta ficha completa = **Uma tupla**

```
Tupla: ('978-3-16-148410-0', '1984', 'George Orwell', 1949, 35.90)
         ↑                    ↑        ↑                ↑     ↑
         ISBN                Título  Autor           Ano  Preço
```

---

### Ordem Importa!

Os valores da tupla seguem a ordem dos campos:

```
📋 Ficha de Livro:
1º campo: ISBN
2º campo: Título
3º campo: Autor
4º campo: Ano
5º campo: Preço

Tupla: (valor1, valor2, valor3, valor4, valor5)
         ↑       ↑       ↑       ↑       ↑
        ISBN   Título  Autor   Ano    Preço
```

**Se trocar a ordem, vira bagunça!**

```
❌ ERRADO: (1949, 'George Orwell', '1984', 978-3-16-148410-0, 35.90)
            ↑ Ano no lugar de ISBN? Bagunça!
```

---

## 🗂️ 4. Relations = O Fichário Completo

### Biblioteca Organizada

**Relation (relação)** é o **fichário completo** com todas as fichas organizadas:

```
🗄️ FICHÁRIO: "Livros"

   ┌─ Estrutura (Schema):
   │  - ISBN (texto)
   │  - Título (texto)
   │  - Autor (texto)
   │  - Ano (número)
   │  - Preço (decimal)
   │
   └─ Fichas (Tuplas/Dados):
      Ficha 1: 1984 de George Orwell
      Ficha 2: Dom Casmurro de Machado de Assis
      Ficha 3: O Cortiço de Aluísio Azevedo
      ...
```

**Relation = Estrutura + Dados**

---

### Propriedades do Fichário

#### 1. Não Há Fichas Duplicadas

Você não vai ter duas fichas **exatamente iguais** para o mesmo livro. Cada ficha é única (garantido por chave primária/ISBN).

#### 2. Ordem das Fichas Não Importa

As fichas podem estar em qualquer ordem no fichário. Se você quer ordem específica (alfabética, por ano), precisa pedir explicitamente (`ORDER BY`).

#### 3. Todas as Fichas Seguem o Mesmo Formato

Se o fichário é de livros, todas as fichas têm: ISBN, Título, Autor, Ano, Preço. Não pode ter uma ficha diferente!

---

## 🔒 5. Constraints = Regras da Biblioteca

### As 6 Regras Fundamentais

Imagine que você é o bibliotecário e precisa manter ordem. Você cria **regras** que as fichas precisam seguir.

---

### Regra 1: PRIMARY KEY = Número Único da Ficha

**Analogia:** Cada ficha tem um **número único** colado nela. Não pode haver duas fichas com o mesmo número.

```
Ficha #1: 1984
Ficha #2: Dom Casmurro
Ficha #3: O Cortiço

❌ NÃO PODE ter outra Ficha #1!
❌ NÃO PODE ter ficha sem número!
```

**No PostgreSQL:**

```sql
id SERIAL PRIMARY KEY  -- Número único e obrigatório
```

---

### Regra 2: FOREIGN KEY = Seta Apontando

**Analogia:** Uma ficha tem uma **seta** apontando para outra ficha.

```
📋 FICHÁRIO: Autores
   Ficha #10: George Orwell
   Ficha #20: Machado de Assis

📋 FICHÁRIO: Livros
   Ficha #1: "1984" → Autor: [Seta para Ficha #10]
   Ficha #2: "Dom Casmurro" → Autor: [Seta para Ficha #20]

❌ NÃO PODE ter seta para Ficha #99 (não existe!)
```

**Regra:** A seta sempre precisa apontar para uma ficha que existe!

---

### Regra 3: UNIQUE = Não Pode Repetir

**Analogia:** Alguns campos **não podem ter valores repetidos** entre fichas.

```
📋 Fichas de Usuários:
   Ficha #1: email@exemplo.com
   Ficha #2: email@exemplo.com  ❌ NÃO PODE! Email já usado!

📋 Fichas de Produtos:
   Ficha #1: Código ABC123
   Ficha #2: Código ABC123  ❌ NÃO PODE! Código já usado!
```

**Mas NULL é permitido:**

```
📋 Fichas de Usuários:
   Ficha #1: email@exemplo.com
   Ficha #2: [sem email]  ✅ PODE! (NULL)
   Ficha #3: [sem email]  ✅ PODE! (outro NULL)
```

---

### Regra 4: CHECK = Porteiro Vigilante

**Analogia:** Há um **porteiro** que verifica se os valores fazem sentido antes de aceitar a ficha.

```
👮 PORTEIRO: "Regra de Idade"

Você: "Quero cadastrar pessoa com 25 anos"
Porteiro: "25 anos? Entre 18 e 100? ✅ PODE ENTRAR!"

Você: "Quero cadastrar pessoa com 10 anos"
Porteiro: "10 anos? Menor de 18? ❌ NÃO PODE!"

Você: "Quero cadastrar pessoa com 200 anos"
Porteiro: "200 anos? Acima de 100? ❌ NÃO PODE!"
```

**Exemplos de porteiros:**

- Preço deve ser positivo
- Nota deve estar entre 0 e 10
- Data de fim deve ser depois da data de início
- Email deve conter "@"

---

### Regra 5: NOT NULL = Campo Obrigatório

**Analogia:** Certos campos **não podem ficar em branco**.

```
📋 Ficha de Cliente:
   Nome: [OBRIGATÓRIO] ❌ Não pode estar vazio!
   Email: [OBRIGATÓRIO] ❌ Não pode estar vazio!
   Telefone: [OPCIONAL] ✅ Pode estar vazio
```

**Tentando cadastrar:**

```
✅ Nome: "João", Email: "joao@email.com", Telefone: [vazio]
   → ACEITO! (telefone é opcional)

❌ Nome: [vazio], Email: "maria@email.com"
   → REJEITADO! (nome é obrigatório)
```

---

### Regra 6: EXCLUSION = Não Pode Conflitar

**Analogia:** Reserva de salas - não pode ter duas reservas que se sobrepõem.

```
🏢 SALA DE REUNIÃO 1

Reserva #1: Segunda, 14:00 - 16:00 ✅ OK
Reserva #2: Segunda, 15:00 - 17:00 ❌ CONFLITO!
            └─ Sobrepõe com Reserva #1!

Reserva #3: Segunda, 16:00 - 18:00 ✅ OK (não sobrepõe)
```

**Regra:** Mesma sala não pode ter horários sobrepostos!

---

## ❓ 6. NULL = Campo em Branco

### A Etiqueta Misteriosa

**NULL** não é:

- ❌ Zero (0)
- ❌ Texto vazio ("")
- ❌ Falso (false)
- ❌ Espaço em branco (" ")

**NULL é:** "**Não sei / Não tenho essa informação**"

---

### Analogia do Formulário

Imagine um formulário em papel:

```
┌────────────────────────────────┐
│  FORMULÁRIO DE CLIENTE         │
├────────────────────────────────┤
│  Nome: João Silva          ✍️   │
│  Email: joao@email.com     ✍️   │
│  Telefone: ___________     ❓   │ ← Campo vazio (NULL)
└────────────────────────────────┘
```

**O campo "Telefone" está em branco.** Isso significa:

- 🤔 Cliente não tem telefone? (possível)
- 🤔 Cliente não quis informar? (possível)
- 🤔 Esqueceram de perguntar? (possível)

**NULL = "Não sabemos"**

---

### NULL vs Zero vs Vazio

```
📦 Caixas de Produtos

Caixa A: Peso = 5 kg     → Tem peso!
Caixa B: Peso = 0 kg     → Tem peso! (É uma caixa vazia conhecida)
Caixa C: Peso = NULL     → NÃO SABEMOS o peso!

🍎 Avaliações de Restaurante

Avaliação 1: 5 estrelas  → Ótimo!
Avaliação 2: 0 estrelas  → Péssimo!
Avaliação 3: NULL        → Ainda não avaliou
```

**Diferença crucial:**

- **0** = Sabemos que é zero
- **NULL** = Não sabemos quanto é

---

### Comparações com NULL São Estranhas!

#### Você NÃO pode perguntar "é igual a NULL?"

```
❌ ERRADO:
SELECT * FROM clientes WHERE telefone = NULL;
└─ Não retorna nada! (mesmo que haja telefones NULL)

✅ CERTO:
SELECT * FROM clientes WHERE telefone IS NULL;
└─ Retorna clientes sem telefone
```

#### Por que isso?

Imagine perguntar:

```
Pergunta: "Seu telefone é igual a 'não sei'?"
Resposta: "Não sei!" 🤷

NULL = NULL → NULL (não sabemos se são iguais!)
```

---

### Operações com NULL = NULL

```
10 + NULL = ?
"Não sei quanto é, então não sei o resultado" = NULL

10 > NULL = ?
"Não sei o valor, então não sei se 10 é maior" = NULL

NULL / 5 = ?
"Não sei o número, então não sei o resultado" = NULL
```

**Tudo envolvendo NULL vira NULL!**

---

### COALESCE = "Use o Primeiro que Não For NULL"

**Analogia:** Lista de contatos de emergência.

```
Contato de emergência de João:
1. Telefone celular: NULL (não tem)
2. Telefone fixo: NULL (não tem)
3. Email: joao@email.com ✅
4. "Sem contato"

COALESCE(celular, fixo, email, 'Sem contato')
→ Retorna: joao@email.com (primeiro não-NULL)
```

**Outro exemplo:**

```
Contato de emergência de Maria:
1. Telefone celular: 11-9999-9999 ✅
2. Telefone fixo: 11-3333-3333
3. Email: maria@email.com

COALESCE(celular, fixo, email)
→ Retorna: 11-9999-9999 (já achou no primeiro!)
```

---

### NULLIF = "Se Forem Iguais, Retorna NULL"

**Analogia:** Evitar divisão por zero.

```
Você quer calcular: Total / Quantidade

Se Quantidade = 0:
Total / 0 = ERRO! 💥 (divisão por zero)

Solução: Transformar 0 em NULL antes de dividir

Total / NULLIF(Quantidade, 0)

Se Quantidade = 5:
  NULLIF(5, 0) = 5
  Total / 5 = resultado normal ✅

Se Quantidade = 0:
  NULLIF(0, 0) = NULL
  Total / NULL = NULL (não dá erro, retorna NULL) ✅
```

---

## 🎓 Resumo Ultra-Simplificado

| Conceito PostgreSQL | Analogia do Mundo Real                       |
| ------------------- | -------------------------------------------- |
| **Domain**          | Molde reutilizável (carimbo com regras)      |
| **Attribute**       | Etiqueta colada na caixa                     |
| **Tuple**           | Ficha preenchida completa                    |
| **Relation**        | Fichário completo (estrutura + fichas)       |
| **PRIMARY KEY**     | Número único da ficha (obrigatório)          |
| **FOREIGN KEY**     | Seta apontando para outra ficha              |
| **UNIQUE**          | Campo que não pode repetir entre fichas      |
| **CHECK**           | Porteiro que valida se valor faz sentido     |
| **NOT NULL**        | Campo obrigatório (não pode ficar em branco) |
| **EXCLUSION**       | Previne conflitos (ex: horários sobrepostos) |
| **NULL**            | "Não sei / Não tenho essa informação"        |
| **COALESCE**        | "Use o primeiro que não for NULL"            |
| **NULLIF**          | "Se forem iguais, retorna NULL"              |

---

## 💡 A Grande Lição

### Domains = Eficiência

Ao invés de repetir regras 100 vezes, crie um **molde** e use em todos os lugares!

### Constraints = Segurança

São **guardas** que protegem seus dados de ficarem bagunçados ou errados.

### NULL ≠ Zero/Vazio

NULL é **"não sei"**, não é **"sei que é zero"** ou **"sei que está vazio"**.

---

## 🎯 Próximo Passo

Agora você vai fazer exercícios onde vai:

- Criar domains customizados (moldes)
- Aplicar todas as constraints (regras)
- Trabalhar com NULL de forma correta
- Usar COALESCE e NULLIF na prática

Prepare-se para consolidar tudo! 💪
