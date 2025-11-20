# Atividade 20: Trabalhando com Datas e Horas

## Objetivo
Dominar manipulação de datas e horas usando `time` package.

## Enunciado
Crie um programa que:
1. Crie função `calcularIdade(dataNascimento time.Time) int` que calcula idade
2. Crie função `diasEntreDatas(data1, data2 time.Time) int` que calcula diferença em dias
3. Crie função `formatarData(data time.Time, formato string) string` que formata em diferentes formatos
4. Crie função `proximoDiaUtil(data time.Time) time.Time` que retorna próximo dia útil (seg-sex)
5. Crie função `agruparPorMes(datas []time.Time) map[string][]time.Time` que agrupa datas por mês
6. Implemente parsing de diferentes formatos de data

## Exemplo de Saída
```
Data de nascimento: 1990-05-15
Idade: 34 anos
Dias desde nascimento: 12345 dias

Próximo dia útil após 2024-01-13 (sábado): 2024-01-15 (segunda)

Datas agrupadas por mês:
  Janeiro 2024: [2024-01-15, 2024-01-20]
  Fevereiro 2024: [2024-02-10]
```

## Dicas
- `time.Now()` para data atual
- `time.Parse(layout, value)` para parsing
- `time.Since()` para diferença
- `time.Add()` para adicionar duração
- Layout: "2006-01-02" (data de referência do Go)



