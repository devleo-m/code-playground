# Módulo 43: Plugins & Dynamic Loading em Go
## Aula 3 - Exercícios e Reflexão

---

## Exercícios Práticos

### Exercício 1: Criar Plugin Simples

Crie um plugin que exporta uma função `Calculate` que recebe dois números e retorna a soma.

**Tarefa**: 
1. Crie o código do plugin
2. Compile como plugin
3. Crie aplicação que carrega e usa o plugin

---

### Exercício 2: Sistema de Plugins com Interface

Crie um sistema onde plugins implementam uma interface `Transformer`:

```go
type Transformer interface {
    Transform(input string) string
    Name() string
}
```

**Tarefa**:
1. Crie interface compartilhada
2. Crie pelo menos 2 plugins (ex: Uppercase, Reverse)
3. Crie aplicação que carrega e usa plugins dinamicamente

---

### Exercício 3: Gerenciador de Plugins

Crie um gerenciador que:
1. Carrega todos os plugins de um diretório
2. Lista plugins disponíveis
3. Permite escolher qual plugin usar

**Tarefa**: Implemente o gerenciador completo.

---

## Perguntas de Reflexão

### Reflexão 1: Plugins vs Alternativas

Por que plugins do Go não são amplamente usados? Quais alternativas são melhores e quando?

**Escreva suas reflexões** (mínimo 250 palavras):

---

### Reflexão 2: Quando Plugins Fazem Sentido

Em que situações reais plugins do Go seriam a melhor solução? Dê exemplos concretos.

**Escreva suas reflexões** (mínimo 200 palavras):

---

## Checklist

- [ ] Entendi o que são plugins
- [ ] Sei criar plugins
- [ ] Sei carregar plugins
- [ ] Entendo limitações
- [ ] Sei quando usar plugins
- [ ] Conheço alternativas

---

**Bons estudos! 🚀**

---

**🎉 Parabéns por completar todas as aulas de tópicos avançados em Go!**

