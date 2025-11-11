# **Aula 1 - Exercícios e Reflexão**

## 📝 Instruções

Complete os exercícios abaixo da melhor forma possível. Não se preocupe em estar "perfeito" - o objetivo é verificar seu entendimento e identificar áreas que precisam de atenção.

**IMPORTANTE**: Responda com suas próprias palavras. Não copie e cole das aulas!

---

## 🎯 Exercício 1: Conceitos Fundamentais (Teórico)

Responda as seguintes perguntas:

### 1.1 - O que é ACID?

Explique com suas palavras o que significa cada letra de ACID e por que isso é importante para um banco de dados.

**Sua resposta:**

```
[ESCREVA SUA RESPOSTA AQUI]
```

### 1.2 - Diferença entre Relacional e NoSQL

Imagine que você está criando dois sistemas:

- **Sistema A**: Uma loja online que vende produtos, tem clientes e pedidos
- **Sistema B**: Uma rede social onde usuários postam fotos, vídeos e textos variados

Para cada sistema, indique se você usaria PostgreSQL (relacional) ou NoSQL, e **justifique sua escolha**.

**Sua resposta:**

```
Sistema A: [PostgreSQL ou NoSQL?]
Justificativa:

Sistema B: [PostgreSQL ou NoSQL?]
Justificativa:
```

---

## 🎯 Exercício 2: Entendendo Tabelas e Relacionamentos (Conceitual)

Você está criando um sistema para uma escola. O sistema precisa guardar:

- **Alunos**: nome, email, data de nascimento
- **Turmas**: nome da turma, ano (ex: "3º Ano A", 2025)
- **Matrícula**: qual aluno está em qual turma

### 2.1 - Desenhe a estrutura das tabelas

Desenhe 3 tabelas (alunos, turmas, matriculas) mostrando quais colunas cada uma teria. Use o formato do exemplo abaixo:

**Exemplo de formato:**

```
Tabela: clientes
+----+-----------+-------------------+
| id | nome      | email             |
+----+-----------+-------------------+
```

**Sua resposta:**

```
Tabela: alunos
+----+...
[COMPLETE AQUI]

Tabela: turmas
+----+...
[COMPLETE AQUI]

Tabela: matriculas
+----+...
[COMPLETE AQUI]
```

### 2.2 - Explique o relacionamento

Como as tabelas se relacionam? Que coluna(s) conecta uma tabela à outra?

**Sua resposta:**

```
[ESCREVA SUA RESPOSTA AQUI]
```

---

## 🎯 Exercício 3: Vantagens e Limitações (Aplicação Prática)

Seu chefe quer criar um novo sistema e está em dúvida entre PostgreSQL e MongoDB (NoSQL). Analise o cenário abaixo e dê sua recomendação:

### Cenário:

**Sistema de Gestão Hospitalar**

- Precisa guardar dados de pacientes (nome, CPF, endereço, histórico médico)
- Precisa guardar consultas (data, médico, paciente, diagnóstico)
- Precisa garantir que prescrições de medicamentos estejam sempre corretas
- Dados de saúde não podem ser perdidos ou corrompidos
- Relacionamentos importantes: Paciente → Consultas → Médicos → Prescrições

### Sua análise:

**Qual banco você recomendaria? (PostgreSQL ou MongoDB)**

```
[SUA ESCOLHA]
```

**Justifique sua escolha mencionando pelo menos 2 características técnicas:**

```
1. [PRIMEIRA JUSTIFICATIVA]

2. [SEGUNDA JUSTIFICATIVA]
```

**Quais limitações do banco escolhido você precisaria considerar?**

```
[SUAS CONSIDERAÇÕES]
```

---

## 🎯 Exercício 4: Recursos do PostgreSQL (Identificação)

PostgreSQL tem vários recursos especiais. Relacione cada recurso com seu uso:

### Recursos:

A) PostGIS  
B) JSONB  
C) Full-text Search  
D) ACID

### Usos:

1. Buscar artigos que contenham as palavras "banco de dados postgresql"
2. Garantir que uma transferência bancária não perca dinheiro
3. Armazenar produtos com atributos variáveis (alguns têm cor, outros têm tamanho, outros têm ambos)
4. Encontrar todas as farmácias em um raio de 5km da minha localização

**Sua resposta:**

```
A) PostGIS → [Número do uso]
B) JSONB → [Número do uso]
C) Full-text Search → [Número do uso]
D) ACID → [Número do uso]
```

---

## 🧠 Perguntas de Reflexão (IMPORTANTE!)

Estas perguntas não têm resposta "certa" ou "errada". O objetivo é fazê-lo **pensar mais profundamente** sobre os conceitos.

### Reflexão 1: O Problema da Rigidez

Vimos que bancos relacionais como PostgreSQL têm "rigidez de schema" - é trabalhoso mudar a estrutura das tabelas depois de criadas.

**Reflita:**

- Por que você acha que essa "rigidez" existe?
- Será que essa rigidez é sempre ruim, ou pode haver benefícios nela?
- Imagine uma empresa que muda a estrutura do banco de dados toda semana sem controle. Que problemas isso poderia causar?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 4 LINHAS]
```

---

### Reflexão 2: ACID vs Velocidade

Sistemas NoSQL muitas vezes sacrificam garantias ACID em favor de velocidade e escalabilidade. O Instagram, por exemplo, pode mostrar que uma foto tem 1.000 curtidas, quando na verdade tem 1.003 - mas isso acontece muito rápido.

**Reflita:**

- Em que situações essa pequena inconsistência é aceitável?
- Em que situações essa inconsistência seria **inaceitável**?
- Como você decidiria, como desenvolvedor, se vale a pena abrir mão de garantias ACID?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 4 LINHAS]
```

---

### Reflexão 3: PostgreSQL vs Outros Bancos

PostgreSQL é gratuito e muito poderoso, mas bancos como Oracle custam milhões de reais e ainda assim são usados por grandes empresas.

**Reflita:**

- Por que você acha que empresas pagariam milhões por Oracle em vez de usar PostgreSQL gratuito?
- O que "gratuito" realmente significa em tecnologia? (dica: pense em treinamento, suporte, manutenção)
- Se você fosse CTO (Chief Technology Officer) de uma startup, qual escolheria? E de um banco multinacional?

**Sua reflexão:**

```
[ESCREVA SEUS PENSAMENTOS AQUI - MÍNIMO 4 LINHAS]
```

---

## 📤 Como Enviar Suas Respostas

1. Copie este arquivo
2. Preencha todas as seções marcadas com `[ESCREVA SUA RESPOSTA AQUI]` ou similar
3. Envie suas respostas completas para análise

**LEMBRE-SE**:

- Não existe resposta "decorada" correta
- O objetivo é entender **seu** processo de raciocínio
- Respostas curtas demais (1 linha) mostram falta de reflexão
- Seja honesto - se não souber, diga "não entendi este ponto"

---

## ⏱️ Tempo Estimado

- Exercícios 1-4: 30-40 minutos
- Reflexões: 20-30 minutos
- **Total: 50-70 minutos**

Não tenha pressa! Qualidade > Velocidade

---

## 🎯 Próximos Passos

Depois de enviar suas respostas, você receberá uma **análise detalhada do seu desempenho**, incluindo:

- Pontos que você dominou
- Conceitos que precisam ser revisados
- Lacunas no seu entendimento
- Recomendações para a próxima aula

Boa sorte! 🚀
