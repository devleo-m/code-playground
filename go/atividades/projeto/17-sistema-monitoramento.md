# Projeto 17: Sistema de Monitoramento

## 📝 Descrição
Sistema para monitorar recursos do sistema (CPU, memória, disco) e gerar alertas.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Métricas**:
   - Coletar uso de CPU
   - Coletar uso de memória
   - Coletar uso de disco
   - Coletar processos ativos

2. **Monitoramento**:
   - Coletar métricas periodicamente
   - Armazenar histórico
   - Calcular médias e picos

3. **Alertas**:
   - Definir thresholds
   - Alertar quando exceder
   - Histórico de alertas

4. **Relatórios**:
   - Uso atual
   - Tendências
   - Gráficos (ASCII)

5. **Persistência**: Salvar métricas em arquivo

## 📚 Conceitos Utilizados
- ✅ Runtime package (métricas)
- ✅ Goroutines (coleta periódica)
- ✅ Channels
- ✅ Time package
- ✅ Slices
- ✅ JSON
- ✅ Concorrência

## 📁 Estrutura Sugerida
```
monitor/
├── main.go
├── metricas.go
├── coletor.go
├── alertas.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Metricas struct {
    Timestamp time.Time
    CPU       float64
    Memoria   float64
    Disco     float64
    Processos int
}

type Alerta struct {
    Tipo      string
    Mensagem  string
    Valor     float64
    Threshold float64
    Timestamp time.Time
}
```

### Coleta
- Goroutine para coletar periodicamente
- Ticker para intervalo
- Channel para comunicação

## ✅ Critérios de Sucesso
- [ ] Métricas são coletadas
- [ ] Monitoramento funciona
- [ ] Alertas são gerados
- [ ] Relatórios são precisos
- [ ] Dados persistem
- [ ] Concorrência funciona

## 🚀 Extras (Desafio)
- [ ] Múltiplos hosts
- [ ] Dashboard web
- [ ] Notificações (email, etc)
- [ ] Análise preditiva
- [ ] Exportar métricas


