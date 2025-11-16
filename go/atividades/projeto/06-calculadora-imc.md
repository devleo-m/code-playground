# Projeto 06: Calculadora IMC e Saúde

## 📝 Descrição
Sistema completo para calcular IMC, acompanhar histórico de peso e gerar relatórios de saúde.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Cálculos**:
   - IMC (Índice de Massa Corporal)
   - Classificação (abaixo do peso, normal, sobrepeso, etc.)
   - Peso ideal (faixa recomendada)
   - Taxa Metabólica Basal (TMB)

2. **Histórico**:
   - Registrar peso e altura
   - Armazenar histórico de medições
   - Visualizar evolução (gráfico ASCII)

3. **Relatórios**:
   - IMC atual e histórico
   - Tendência (ganhando/perdendo peso)
   - Meta de peso

4. **Validação**: Altura e peso em ranges válidos

5. **Persistência**: JSON

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Funções matemáticas
- ✅ Slices
- ✅ Error handling
- ✅ JSON
- ✅ Time package
- ✅ Formatação de saída

## 📁 Estrutura Sugerida
```
imc/
├── main.go
├── calculadora.go
├── historico.go
├── relatorio.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Medicao struct {
    Data   time.Time
    Peso   float64
    Altura float64
    IMC    float64
}

type Pessoa struct {
    Nome     string
    Altura   float64
    Medicoes []Medicao
}
```

### Fórmulas
- IMC = peso / (altura²)
- TMB Homem = 88.362 + (13.397 × peso) + (4.799 × altura) - (5.677 × idade)
- TMB Mulher = 447.593 + (9.247 × peso) + (3.098 × altura) - (4.330 × idade)

## ✅ Critérios de Sucesso
- [ ] Cálculos são precisos
- [ ] Histórico funciona
- [ ] Relatórios são informativos
- [ ] Validações funcionam
- [ ] Dados persistem
- [ ] Interface clara

## 🚀 Extras (Desafio)
- [ ] Múltiplas pessoas
- [ ] Gráficos mais elaborados
- [ ] Exportar relatório PDF
- [ ] Lembretes de medição
- [ ] Integração com metas

