# Atividade 42: Processador de CSV

## Objetivo
Criar processador completo de arquivos CSV.

## Enunciado
Crie um programa que:
1. Leia arquivo CSV com header
2. Parse linhas e colunas
3. Converta para slice de maps (chave: nome da coluna)
4. Filtre linhas baseado em critérios
5. Agrupe por coluna específica
6. Calcule estatísticas (soma, média, máximo, mínimo) de coluna numérica
7. Exporte resultado filtrado para novo CSV

## Exemplo de CSV
```csv
nome,idade,salario
João,30,5000
Maria,25,6000
Pedro,35,5500
```

## Exemplo de Saída
```
CSV carregado: 3 linhas
Filtro: idade > 28
Resultado: 2 linhas
Agrupado por faixa etária:
  30-35: 2 pessoas
Estatísticas salário:
  Soma: 16500
  Média: 5500
  Máximo: 6000
  Mínimo: 5000
```

## Dicas
- Use `encoding/csv` package
- `csv.NewReader` para ler
- Parse header na primeira linha
- Converta strings para números quando necessário
- `csv.NewWriter` para escrever

