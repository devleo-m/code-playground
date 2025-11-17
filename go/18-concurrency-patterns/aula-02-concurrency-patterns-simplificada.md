# Módulo 18: Concurrency Patterns em Go

## Aula 2 - Simplificada: Entendendo Padrões de Concorrência com Analogias do Dia a Dia

Olá! Na aula anterior, vimos os padrões de concorrência de forma técnica. Agora vamos entender esses mesmos conceitos usando analogias do nosso dia a dia. Isso vai ajudar a fixar o aprendizado de forma mais natural e intuitiva!

---

## 🎯 Revisão Rápida: Por que Padrões?

Antes de começarmos, lembre-se: **padrões são como receitas de bolo**. Você não precisa inventar a roda toda vez. Alguém já descobriu a melhor forma de fazer algo, e você só precisa seguir a receita (padrão) adaptando para sua situação.

Os padrões de concorrência são receitas testadas para organizar goroutines e channels de forma eficiente e segura.

---

## 1. Fan-In: O Carteiro que Coleta Cartas

### A Analogia

Imagine que você tem **três caixas de correio diferentes** (três channels):
- Caixa 1: Recebe cartas de amigos
- Caixa 2: Recebe cartas de trabalho  
- Caixa 3: Recebe cartas de família

O padrão **Fan-In** é como um **carteiro** que vai de caixa em caixa, pega todas as cartas e coloca em uma **única sacola** (um channel de saída). Assim, você só precisa olhar em um lugar para ver todas as suas cartas!

### No Código

```go
// Três pessoas escrevendo cartas (três channels de entrada)
caixaAmigos := escreverCartas("amigos", 3)
caixaTrabalho := escreverCartas("trabalho", 3)
caixaFamilia := escreverCartas("familia", 3)

// Carteiro coletando todas as cartas (Fan-In)
sacolaUnica := carteiroColetor(caixaAmigos, caixaTrabalho, caixaFamilia)

// Você lendo todas as cartas de um só lugar
for carta := range sacolaUnica {
    fmt.Println("Recebi:", carta)
}
```

### Quando Usar?

- ✅ Quando você tem **várias fontes** produzindo dados e quer **juntar tudo**
- ✅ Como um **coletor de resultados** de múltiplas tarefas paralelas
- ✅ Para **centralizar** informações que vêm de lugares diferentes

**Exemplo Real**: Você pede informações de preços em 5 sites diferentes e quer juntar todos os resultados em uma lista única.

---

## 2. Fan-Out: A Linha de Produção com Múltiplos Trabalhadores

### A Analogia

Imagine uma **fábrica de biscoitos**:
- Uma **esteira transportadora** (channel de entrada) traz massa crua
- Você tem **5 trabalhadores** (5 goroutines) ao lado da esteira
- Cada trabalhador pega um biscoito da esteira, coloca no forno, decora e embala
- Todos trabalham **ao mesmo tempo**, processando biscoitos diferentes

O padrão **Fan-Out** é exatamente isso: **distribuir o trabalho** de uma única fonte (esteira) para múltiplos trabalhadores (goroutines) que processam em paralelo.

### No Código

```go
// Esteira com biscoitos para processar
esteira := make(chan string, 10)
esteira <- "biscoito1"
esteira <- "biscoito2"
esteira <- "biscoito3"
// ... mais biscoitos

// 5 trabalhadores processando (Fan-Out)
for i := 1; i <= 5; i++ {
    go trabalhador(i, esteira, resultados)
}

// Coletar biscoitos processados
for resultado := range resultados {
    fmt.Println("Biscoito pronto:", resultado)
}
```

### Quando Usar?

- ✅ Quando você tem **muitas tarefas similares** para processar
- ✅ Para **acelerar o processamento** usando múltiplos "trabalhadores"
- ✅ Quando quer **paralelizar** trabalho que pode ser feito independentemente

**Exemplo Real**: Você tem 1000 fotos para redimensionar. Em vez de fazer uma por vez, você distribui para 10 workers que processam 100 cada, muito mais rápido!

---

## 3. Pipeline: A Linha de Montagem de Carros

### A Analogia

Imagine uma **linha de montagem de carros**:

1. **Estação 1 - Chassi**: Monta a estrutura básica do carro
2. **Estação 2 - Motor**: Instala o motor no chassi
3. **Estação 3 - Pintura**: Pinta o carro
4. **Estação 4 - Bancos**: Coloca os bancos
5. **Estação 5 - Teste**: Testa se tudo funciona

Cada estação recebe o carro da estação anterior, faz sua parte e passa para a próxima. **Enquanto um carro está sendo pintado, outro pode estar recebendo o motor**, e assim por diante. Todas as estações trabalham **ao mesmo tempo** em carros diferentes!

O padrão **Pipeline** é exatamente isso: **encadear estágios** onde cada um faz uma transformação e passa para o próximo.

### No Código

```go
// Pipeline: números -> elevar ao quadrado -> formatar -> imprimir

// Estação 1: Gerar números
numeros := gerarNumeros(10)

// Estação 2: Elevar ao quadrado (recebe números, envia quadrados)
quadrados := elevarAoQuadrado(numeros)

// Estação 3: Formatar (recebe quadrados, envia strings)
formatados := formatar(quadrados)

// Estação 4: Imprimir (recebe strings, imprime)
for resultado := range formatados {
    fmt.Println(resultado)
}
```

### Quando Usar?

- ✅ Quando você precisa fazer **várias transformações** nos dados, uma após a outra
- ✅ Para **separar responsabilidades**: cada estágio faz uma coisa só
- ✅ Para processar dados **conforme chegam**, não precisa esperar tudo terminar

**Exemplo Real**: Processar vídeos: baixar → converter formato → adicionar legenda → comprimir → fazer upload. Cada etapa pode processar vídeos diferentes ao mesmo tempo!

---

## 4. Worker Pool: A Empresa com Número Fixo de Funcionários

### A Analogia

Imagine uma **empresa de entregas**:
- A empresa tem **exatamente 5 funcionários** (número fixo de workers)
- Quando chegam pedidos, eles vão para uma **fila** (channel de jobs)
- Os funcionários pegam pedidos da fila, fazem a entrega e voltam para pegar o próximo
- Se todos estão ocupados, os pedidos **esperam na fila**
- Se a fila está vazia, os funcionários **ficam esperando** novos pedidos

O padrão **Worker Pool** é isso: **controlar quantos "trabalhadores"** você tem trabalhando ao mesmo tempo, evitando sobrecarregar o sistema.

### No Código

```go
// Empresa com 5 funcionários
empresa := NovaEmpresa(5) // 5 workers
empresa.Iniciar()

// Pedidos chegando
for i := 1; i <= 20; i++ {
    empresa.AdicionarPedido(Pedido{ID: i, Endereco: "Rua X"})
}

// Funcionários processam os pedidos da fila
// Quando todos terminarem, encerramos
empresa.Encerrar()
```

### Quando Usar?

- ✅ Quando você quer **limitar** quantas coisas processam ao mesmo tempo
- ✅ Para **controlar recursos**: não criar goroutines demais
- ✅ Quando tem **muitas tarefas** mas quer processar de forma controlada

**Exemplo Real**: Você tem 10.000 emails para enviar, mas não quer sobrecarregar o servidor. Cria um pool de 10 workers que enviam 10 emails por vez, de forma controlada.

---

## 5. Pub-Sub: O Sistema de Notícias por Assinatura

### A Analogia

Imagine um **jornal digital**:
- O jornal **publica notícias** sobre diferentes tópicos: esportes, tecnologia, política, etc.
- Pessoas se **inscrevem** para receber notícias de tópicos que lhes interessam
- Quando o jornal publica uma notícia de "tecnologia", **todos que se inscreveram em tecnologia** recebem
- O jornal **não sabe** quem são os assinantes, apenas publica
- Os assinantes **não se conhecem**, apenas recebem o que lhes interessa

O padrão **Pub-Sub** é isso: **desacoplar** quem publica de quem recebe, usando "tópicos" como intermediário.

### No Código

```go
// Criar o sistema de notícias
jornal := NovoJornal()

// Pessoa 1 se inscreve em "tecnologia"
pessoa1 := jornal.Inscrever("tecnologia")
go func() {
    for noticia := range pessoa1 {
        fmt.Println("Pessoa 1 recebeu:", noticia)
    }
}()

// Pessoa 2 se inscreve em "tecnologia" e "esportes"
pessoa2Tech := jornal.Inscrever("tecnologia")
pessoa2Sports := jornal.Inscrever("esportes")
// ... processar notícias

// Jornal publica notícias
jornal.Publicar("tecnologia", "Nova versão do Go!")
jornal.Publicar("esportes", "Brasil vence!")
jornal.Publicar("tecnologia", "Go é incrível!")
```

### Quando Usar?

- ✅ Quando componentes **não precisam se conhecer** diretamente
- ✅ Para **sistemas de eventos**: algo acontece e vários componentes reagem
- ✅ Para **notificações**: avisar múltiplos clientes sobre algo
- ✅ Em **microserviços**: serviços se comunicam sem conhecer uns aos outros

**Exemplo Real**: Sistema de e-commerce. Quando um pedido é feito, vários componentes precisam saber: estoque (diminuir), financeiro (cobrar), logística (preparar envio), marketing (enviar email). Com Pub-Sub, o sistema de pedidos apenas "publica" o evento, e cada componente "se inscreve" no que precisa.

---

## 🎨 Comparando os Padrões: Uma Tabela Visual

| Padrão | Analogia | Quando Usar |
|--------|----------|-------------|
| **Fan-In** | Carteiro coletando cartas de várias caixas | Juntar resultados de várias fontes |
| **Fan-Out** | Múltiplos trabalhadores em uma linha de produção | Distribuir trabalho para acelerar |
| **Pipeline** | Linha de montagem com várias estações | Transformações sequenciais de dados |
| **Worker Pool** | Empresa com número fixo de funcionários | Controlar recursos e throughput |
| **Pub-Sub** | Jornal com assinantes por tópico | Comunicação desacoplada via eventos |

---

## 🔄 Combinando Padrões: Um Exemplo Completo

Vamos imaginar uma **fábrica de biscoitos gourmet** que usa vários padrões:

1. **Pipeline**: Massa → Cortar → Assar → Decorar → Embalar
2. **Fan-Out**: Na etapa de "Decorar", temos 5 decoradores trabalhando em paralelo
3. **Fan-In**: Todos os biscoitos decorados vão para uma única esteira de embalagem
4. **Worker Pool**: Temos exatamente 3 embaladores (não mais, não menos)
5. **Pub-Sub**: Quando biscoitos ficam prontos, publicamos evento "biscoito-pronto". Sistema de estoque, sistema de vendas e sistema de logística se inscrevem para reagir.

```go
// 1. Pipeline: etapas de produção
massa := prepararMassa()
cortados := cortar(massa)
assados := assar(cortados)

// 2. Fan-Out: 5 decoradores trabalhando
decorado1 := decorar("Decorador 1", assados)
decorado2 := decorar("Decorador 2", assados)
decorado3 := decorar("Decorador 3", assados)
decorado4 := decorar("Decorador 4", assados)
decorado5 := decorar("Decorador 5", assados)

// 3. Fan-In: juntar todos os biscoitos decorados
todosDecorados := juntar(decorado1, decorado2, decorado3, decorado4, decorado5)

// 4. Worker Pool: 3 embaladores
pool := NovoPool(3)
embalados := pool.Processar(todosDecorados)

// 5. Pub-Sub: notificar sistemas
for biscoito := range embalados {
    sistemaEventos.Publicar("biscoito-pronto", biscoito)
}
```

---

## 💡 Dicas para Escolher o Padrão Certo

### Use Fan-In quando:
- Você tem **várias goroutines produzindo** e quer **juntar tudo**
- Precisa **agregar resultados** de múltiplas fontes

### Use Fan-Out quando:
- Você tem **muitas tarefas similares** para processar
- Quer **acelerar** usando paralelismo
- Tarefas são **independentes** umas das outras

### Use Pipeline quando:
- Você precisa fazer **várias transformações** sequenciais
- Quer **separar responsabilidades** (cada estágio faz uma coisa)
- Dados podem ser processados **conforme chegam**

### Use Worker Pool quando:
- Você quer **limitar recursos** (não criar goroutines demais)
- Precisa **controlar throughput** (quantos processam por vez)
- Tem **muitas tarefas** mas quer processar de forma controlada

### Use Pub-Sub quando:
- Componentes **não precisam se conhecer**
- Você quer **desacoplar** quem produz de quem consome
- Precisa de **comunicação assíncrona** via eventos

---

## 🎯 Resumo em Uma Frase

- **Fan-In**: "Juntar várias fontes em uma só"
- **Fan-Out**: "Distribuir trabalho para vários trabalhadores"
- **Pipeline**: "Linha de montagem: cada estágio faz sua parte e passa adiante"
- **Worker Pool**: "Empresa com número fixo de funcionários processando fila"
- **Pub-Sub**: "Jornal: publica notícias, assinantes recebem o que lhes interessa"

---

E assim terminamos nossa aula simplificada! Agora você entende os padrões de concorrência não apenas tecnicamente, mas também de forma intuitiva através das analogias. 

Na próxima aula, vamos praticar com exercícios para fixar ainda mais esses conceitos!

Sinta-se à vontade para reler este material. Se tiver qualquer dúvida, pode perguntar!


