# Módulo 12: Interfaces em Go - Exercícios e Reflexão

## Aula 3: Praticando Interfaces

Olá! Agora é hora de colocar a mão na massa e praticar o que você aprendeu sobre interfaces. Vamos fazer exercícios práticos e também refletir sobre os conceitos.

---

## Exercício 1: Interface de Animal

Crie uma interface `Animal` que tenha os métodos:
- `FazerSom() string` - retorna o som que o animal faz
- `Mover() string` - retorna como o animal se move

Implemente essa interface para pelo menos 3 animais diferentes (ex: Cachorro, Gato, Pássaro).

**Exemplo de uso esperado:**
```go
func main() {
    animais := []Animal{
        Cachorro{Nome: "Rex"},
        Gato{Nome: "Mimi"},
        Passaro{Nome: "Piu"},
    }
    
    for _, animal := range animais {
        fmt.Println(animal.FazerSom())
        fmt.Println(animal.Mover())
    }
}
```

---

## Exercício 2: Calculadora com Interface

Crie uma interface `Calculadora` com o método `Calcular(a, b float64) float64`.

Implemente essa interface para:
- `Soma` - soma dois números
- `Subtracao` - subtrai dois números
- `Multiplicacao` - multiplica dois números
- `Divisao` - divide dois números (cuidado com divisão por zero!)

Crie uma função que aceite uma `Calculadora` e dois números, e retorne o resultado.

---

## Exercício 3: Type Switch com Formas Geométricas

Crie uma interface `Forma` com o método `Area() float64`.

Implemente para `Retangulo`, `Circulo` e `Triangulo`.

Crie uma função que use **type switch** para:
- Se for `Retangulo`, imprimir "É um retângulo" e a área
- Se for `Circulo`, imprimir "É um círculo" e a área
- Se for `Triangulo`, imprimir "É um triângulo" e a área
- Caso contrário, imprimir "Forma desconhecida"

---

## Exercício 4: Sistema de Notificações

Crie uma interface `Notificador` com o método `Enviar(mensagem string) error`.

Implemente para:
- `EmailNotificador` - simula envio de email (apenas imprime)
- `SMSNotificador` - simula envio de SMS (apenas imprime)
- `PushNotificador` - simula notificação push (apenas imprime)

Crie uma função que aceite um slice de `Notificador` e uma mensagem, e envie a mensagem para todos os notificadores.

**Dica**: Use um loop para iterar sobre o slice de interfaces.

---

## Perguntas de Reflexão

### Reflexão 1: Por Que Interfaces São Importantes?

Pense sobre o seguinte cenário:

Você está criando um sistema de e-commerce que precisa processar pagamentos. Hoje você aceita apenas cartão de crédito, mas no futuro pode querer aceitar:
- Pix
- Boleto
- PayPal
- Criptomoedas

**Pergunta**: Como interfaces ajudam nesse cenário? O que aconteceria se você não usasse interfaces e dependesse diretamente do tipo `CartaoCredito`?

**Sua resposta deve incluir**:
- Explicação de como interfaces tornam o código flexível
- O que aconteceria se você precisasse adicionar um novo método de pagamento sem interfaces
- Por que é melhor depender de uma interface do que de um tipo concreto

---

### Reflexão 2: Type Assertions vs Type Switch

Você tem uma função que recebe `interface{}` e precisa processar diferentes tipos de dados:

```go
func Processar(dados interface{}) {
    // Como você processaria isso?
}
```

**Pergunta**: Quando você usaria **type assertion** e quando usaria **type switch**? Dê exemplos práticos de cada situação.

**Sua resposta deve incluir**:
- Uma situação onde type assertion é mais apropriado
- Uma situação onde type switch é mais apropriado
- As vantagens e desvantagens de cada abordagem
- Quando você escolheria uma sobre a outra

---

### Reflexão 3: Empty Interface vs Generics

Com o Go 1.18+, temos **generics** que podem substituir muitos usos de `interface{}`.

**Pergunta**: Quando você ainda usaria `interface{}` em vez de generics? Quando você usaria generics em vez de `interface{}`?

**Sua resposta deve incluir**:
- Uma situação onde `interface{}` ainda é apropriado
- Uma situação onde generics são melhores
- As diferenças fundamentais entre os dois
- Por que Go mantém `interface{}` mesmo com generics disponíveis

---

### Reflexão 4: Design de Interfaces

Você está criando uma biblioteca que precisa de uma interface para processar dados. Você tem duas opções:

**Opção A**: Uma interface grande com muitos métodos
```go
type Processador interface {
    Ler() []byte
    Escrever([]byte) error
    Validar() bool
    Transformar() []byte
    Salvar() error
    Limpar() error
}
```

**Opção B**: Várias interfaces pequenas
```go
type Leitor interface {
    Ler() []byte
}

type Escritor interface {
    Escrever([]byte) error
}

type Validador interface {
    Validar() bool
}

// ... etc
```

**Pergunta**: Qual abordagem você escolheria e por quê? Qual segue melhor as práticas do Go?

**Sua resposta deve incluir**:
- Qual abordagem você escolheria
- As vantagens da sua escolha
- As desvantagens da outra abordagem
- Como isso se relaciona com o princípio da responsabilidade única
- Exemplos de interfaces da biblioteca padrão do Go que seguem esse padrão

---

## Desafio Opcional: Sistema de Cache

Crie um sistema de cache genérico usando interfaces. O sistema deve:

1. Ter uma interface `Cache` com métodos:
   - `Get(chave string) (interface{}, bool)` - obtém um valor
   - `Set(chave string, valor interface{})` - armazena um valor
   - `Delete(chave string)` - remove um valor
   - `Clear()` - limpa todo o cache

2. Implementar pelo menos 2 tipos de cache:
   - `CacheMemoria` - cache em memória (use um map)
   - `CacheLimitado` - cache com limite de tamanho (remove itens antigos quando cheio)

3. Criar uma função que aceite qualquer `Cache` e teste ambas as implementações.

**Dica**: Use `interface{}` para os valores, mas considere como você poderia melhorar isso com generics (Go 1.18+).

---

## Como Entregar

Crie um arquivo `exercicios.go` com todas as suas soluções. Para as perguntas de reflexão, você pode:

1. Escrever suas respostas em comentários no código
2. Criar um arquivo separado `reflexoes.md` com suas respostas
3. Ou simplesmente me enviar suas respostas quando terminar

**Lembre-se**: 
- Não se preocupe se não conseguir fazer tudo de primeira
- O importante é tentar e pensar sobre os problemas
- As reflexões são mais importantes que os exercícios - elas mostram que você realmente entendeu os conceitos

Boa sorte! 🚀

---

**Dica**: Comece pelos exercícios mais simples e vá aumentando a complexidade. Se travar em algum exercício, revise a aula principal e tente novamente!

