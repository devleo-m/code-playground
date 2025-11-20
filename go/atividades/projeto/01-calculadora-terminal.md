# Projeto 01: Calculadora de Terminal

## 📝 Descrição
Crie uma calculadora interativa que roda no terminal, permitindo ao usuário realizar operações matemáticas básicas e avançadas.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Operações Básicas**: Soma, subtração, multiplicação, divisão
2. **Interface Interativa**: Menu que permite escolher operação
3. **Histórico**: Armazene últimas 10 operações realizadas
4. **Validação**: Trate divisão por zero e entradas inválidas
5. **Formatação**: Exiba resultados com 2 casas decimais
6. **Persistência**: Salve histórico em arquivo JSON

### Operações Extras
- Potenciação (x^y)
- Raiz quadrada
- Porcentagem
- Operações com múltiplos números

## 📚 Conceitos Utilizados
- ✅ Variáveis e tipos
- ✅ Funções e métodos
- ✅ Error handling
- ✅ Structs
- ✅ Slices e maps
- ✅ I/O (fmt, os)
- ✅ JSON (encoding/json)
- ✅ Loops e condicionais

## 📁 Estrutura Sugerida
```
calculadora/
├── main.go
├── calculadora.go
├── historico.go
├── storage.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs Principais
```go
type Operacao struct {
    Tipo      string
    Numeros   []float64
    Resultado float64
    Data      time.Time
}

type Calculadora struct {
    Historico []Operacao
}
```

### Funções Principais
- `Somar(a, b float64) float64`
- `Subtrair(a, b float64) float64`
- `Multiplicar(a, b float64) float64`
- `Dividir(a, b float64) (float64, error)`
- `AdicionarHistorico(op Operacao)`
- `ExibirHistorico()`
- `SalvarHistorico() error`
- `CarregarHistorico() error`

## ✅ Critérios de Sucesso
- [ ] Todas operações básicas funcionam
- [ ] Menu interativo funciona
- [ ] Histórico armazena e exibe corretamente
- [ ] Erros são tratados adequadamente
- [ ] Histórico persiste entre execuções
- [ ] Código organizado e legível

## 🚀 Extras (Desafio)
- [ ] Modo expressão (ex: "2 + 3 * 4")
- [ ] Variáveis (ex: "x = 10")
- [ ] Gráficos simples (ASCII art)
- [ ] Modo científico (sen, cos, log)



