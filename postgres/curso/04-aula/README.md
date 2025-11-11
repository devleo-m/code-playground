# 🔒 Aula 4: Relational Model - Aprofundamento

Bem-vindo à Aula 4! Esta é a **última aula teórica** antes da avaliação final. Aqui você vai aprofundar o modelo relacional, aprendendo sobre domains customizados, constraints avançadas e tratamento correto de valores NULL.

---

## 📚 Estrutura da Aula 4

### **Tópico: Aprofundamento no Modelo Relacional**

#### 📖 Materiais Disponíveis:

1. **[aula-04-relational-model.md](./aula-04-relational-model.md)** - Aula principal completa

   - Revisão do modelo relacional de E.F. Codd
   - **Domains**: Tipos customizados reutilizáveis
   - **Attributes**: Colunas e suas propriedades
   - **Tuples**: Registros e operações
   - **Relations**: Tabelas completas (schema + dados)
   - **Constraints Completas**:
     - PRIMARY KEY (chave primária)
     - FOREIGN KEY (chave estrangeira + CASCADE)
     - UNIQUE (unicidade)
     - CHECK (validações customizadas)
     - NOT NULL (obrigatório)
     - EXCLUSION (previne conflitos)
   - **NULL Values**: Tratamento correto de valores nulos
     - IS NULL / IS NOT NULL
     - COALESCE (primeiro não-nulo)
     - NULLIF (evitar divisão por zero)
     - Lógica de três valores (TRUE, FALSE, NULL)

2. **[aula-04-simplificada.md](./aula-04-simplificada.md)** - Versão com analogias

   - Domains = Moldes reutilizáveis (carimbos)
   - Attributes = Etiquetas em caixas
   - Tuples = Fichas preenchidas
   - Relations = Fichário completo
   - Constraints = Regras da biblioteca
   - NULL = Campo em branco ("não sei")
   - COALESCE = "Use o primeiro que não for NULL"
   - NULLIF = "Se forem iguais, retorna NULL"

3. **[aula-04-exercicios.md](./aula-04-exercicios.md)** - **Exercícios finais consolidação**
   - 6 exercícios práticos extensos
   - **Modelagem completa**: Sistema de imobiliária
   - Criação de domains customizados
   - Todas as constraints aplicadas
   - Queries complexas (JOINs, agregações, NULL)
   - 3 reflexões profundas
   - **Tempo estimado: 4,5-6 horas** ⚠️ É extenso!

---

## 🎯 Como Estudar Esta Aula

### **Passo 1: Leia a Aula Principal** (60-90 minutos)

📖 Abra: `aula-04-relational-model.md`

**Foque especialmente em:**

#### Domains

- Por que criar tipos customizados?
- Como reutilizar validações?
- CREATE/ALTER/DROP DOMAIN

#### Constraints Avançadas

- Diferença entre PRIMARY KEY e UNIQUE
- FOREIGN KEY com CASCADE
- CHECK constraints complexas
- EXCLUSION (previne conflitos)

#### NULL

- NULL ≠ 0, ≠ '', ≠ false
- Por que `WHERE campo = NULL` não funciona
- COALESCE e NULLIF na prática

---

### **Passo 2: Leia a Versão Simplificada** (30-40 minutos)

🎯 Abra: `aula-04-simplificada.md`

As analogias vão fixar os conceitos:

- Domains = Moldes (eficiência)
- Constraints = Regras da biblioteca (segurança)
- NULL = "Não sei" (não é zero nem vazio)

---

### **Passo 3: Faça os Exercícios** (4,5-6 horas) ⚠️ MUITO IMPORTANTE!

✏️ Abra: `aula-04-exercicios.md`

**ATENÇÃO:** Estes são os exercícios finais de consolidação!

**Estrutura:**

1. Criar domains customizados (CPF, CEP, telefone, etc)
2. Modelar sistema completo de imobiliária
3. Aplicar TODAS as constraints
4. Povoar banco de dados
5. Queries complexas (20+ queries!)
6. Trabalhar com NULL corretamente
7. 3 reflexões profundas

**Este exercício testa TUDO das aulas 1-4!**

---

### **Passo 4: Revisão Geral**

Antes de enviar, revise:

- ✅ Aula 1: Conceitos gerais de PostgreSQL
- ✅ Aula 2: Modelo relacional teórico
- ✅ Aula 3: Implementação prática (queries, tipos)
- ✅ Aula 4: Aprofundamento (domains, constraints, NULL)

---

## ⚠️ Por Que Esta Aula é Crucial?

### É a Ponte Final

```
Aula 1-2: TEORIA (conceitos abstratos)
     ↓
Aula 3: PRÁTICA BÁSICA (queries, tipos)
     ↓
Aula 4: APROFUNDAMENTO (qualidade e integridade)
     ↓
AVALIAÇÃO FINAL
```

### Completa o Fundamento

**Sem Aula 4:**

- ❌ Código repetitivo (sem domains)
- ❌ Dados inconsistentes (sem constraints adequadas)
- ❌ Bugs com NULL

**Com Aula 4:**

- ✅ Código reutilizável (domains)
- ✅ Dados íntegros (constraints completas)
- ✅ Tratamento correto de NULL

---

## 📊 Conceitos-Chave Desta Aula

### Domains

- [ ] Por que criar tipos customizados
- [ ] CREATE/ALTER/DROP DOMAIN
- [ ] Quando usar domain vs CHECK inline
- [ ] Listar e consultar domains

### Constraints

- [ ] PRIMARY KEY: único + not null + índice
- [ ] FOREIGN KEY: integridade referencial + CASCADE
- [ ] UNIQUE: valores únicos (permite NULL)
- [ ] CHECK: validações customizadas
- [ ] NOT NULL: valor obrigatório
- [ ] EXCLUSION: previne conflitos (reservas, etc)

### NULL

- [ ] NULL ≠ zero, vazio, false
- [ ] IS NULL / IS NOT NULL (nunca = NULL)
- [ ] Operações com NULL retornam NULL
- [ ] Lógica de três valores
- [ ] COALESCE: primeiro não-nulo
- [ ] NULLIF: evitar divisão por zero
- [ ] NULL em agregações (COUNT, SUM, AVG)
- [ ] NULL em ORDER BY (NULLS FIRST/LAST)

---

## 🎯 Checklist de Prontidão

Antes de fazer os exercícios, você deve conseguir:

**Domains:**

- [ ] Criar domain com constraints
- [ ] Explicar vantagens de domains
- [ ] Modificar e deletar domains

**Constraints:**

- [ ] Criar PRIMARY KEY simples e composta
- [ ] Criar FOREIGN KEY com ON DELETE CASCADE
- [ ] Usar UNIQUE em uma e múltiplas colunas
- [ ] Criar CHECK constraints complexas
- [ ] Aplicar NOT NULL
- [ ] Entender quando usar EXCLUSION

**NULL:**

- [ ] Explicar diferença entre NULL, 0 e ''
- [ ] Usar IS NULL / IS NOT NULL
- [ ] Usar COALESCE para valores padrão
- [ ] Usar NULLIF para evitar erros
- [ ] Entender NULL em agregações

**Menos de 12 marcados:** Releia a aula.
**12-15 marcados:** Você está quase pronto!
**Todos marcados:** Vá para os exercícios!

---

## 💡 Dicas de Estudo

### 1. Teste TUDO no PostgreSQL 💻

Domains, constraints, NULL - teste cada conceito!

### 2. Crie Domains Realistas 📝

Pense em CPF, CEP, telefone, email - coisas do mundo real.

### 3. Quebre Constraints Propositalmente 🔨

Tente inserir dados inválidos para ver constraints funcionando!

### 4. Brinque com NULL 🎮

```sql
SELECT NULL = NULL;   -- Resultado?
SELECT NULL + 10;     -- Resultado?
SELECT COALESCE(NULL, NULL, 'oi');  -- Resultado?
```

---

## 🔄 Conectando com Aulas Anteriores

| Aula | Conceito Teórico            | Aula 4 Implementa     |
| ---- | --------------------------- | --------------------- |
| 2    | Domínio (conceito abstrato) | Domain (tipo SQL)     |
| 2    | Integridade de Entidade     | PRIMARY KEY, NOT NULL |
| 2    | Integridade Referencial     | FOREIGN KEY           |
| 2    | Integridade de Domínio      | Domain + CHECK        |
| 3    | Tipos de dados              | Domain customizado    |

---

## ⏱️ Tempo Total Estimado

| Atividade            | Tempo         |
| -------------------- | ------------- |
| Leitura principal    | 60-90 min     |
| Leitura simplificada | 30-40 min     |
| Exercícios 1-2       | 90-120 min    |
| Exercícios 3-4       | 90-120 min    |
| Exercícios 5-6       | 70-90 min     |
| Reflexões            | 40-50 min     |
| **Total**            | **6-8 horas** |

**Divida em 3-4 sessões!**

---

## 🚀 Preparação para Avaliação Final

Esta aula completa os fundamentos! Após Aula 4, você terá:

- ✅ Aula 1: Visão geral PostgreSQL
- ✅ Aula 2: Teoria do modelo relacional
- ✅ Aula 3: Prática (queries, tipos)
- ✅ Aula 4: Qualidade (domains, constraints, NULL)

**= BASE SÓLIDA COMPLETA!**

---

## 📞 Dúvidas Frequentes

### "Domain é muito importante?"

**Sim!** É um diferencial do PostgreSQL. Mostra código maduro e manutenível.

### "Preciso decorar todos os tipos de constraints?"

Não precisa decorar, mas precisa saber **quando usar cada uma**.

### "NULL é tão complicado assim?"

NULL confunde no início, mas com prática vira natural. Teste muito!

### "Quantas horas devo dedicar?"

**Mínimo 6 horas** para fazer bem feito. É a última aula antes da avaliação!

---

## 🎓 Mensagem Final

Você está na reta final dos fundamentos! A Aula 4 consolida TUDO e te prepara para criar bancos de dados robustos, íntegros e bem arquitetados.

**Invista tempo de qualidade aqui.** Cada conceito domingado agora vai economizar horas de debugging no futuro!

---

Você está pronto! Comece pela aula principal e boa sorte! 🚀💪
