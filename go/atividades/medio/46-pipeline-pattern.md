# Atividade 46: Padrão Pipeline

## Objetivo
Implementar pipeline de processamento de dados.

## Enunciado
Crie um programa que:
1. Defina tipo `Stage func([]int) []int` (estágio do pipeline)
2. Crie estágios:
   - `FiltrarPares` - mantém apenas pares
   - `MultiplicarPor2` - dobra valores
   - `FiltrarMaioresQue10` - remove <= 10
   - `Somar1` - adiciona 1
3. Crie função `CriarPipeline(stages ...Stage) func([]int) []int`
4. Execute pipeline: dados passam por cada estágio sequencialmente
5. Teste com diferentes sequências de estágios

## Exemplo de Saída
```
Dados iniciais: [1, 2, 3, 4, 5, 6, 7, 8]
Pipeline: FiltrarPares -> MultiplicarPor2 -> FiltrarMaioresQue10
Após FiltrarPares: [2, 4, 6, 8]
Após MultiplicarPor2: [4, 8, 12, 16]
Após FiltrarMaioresQue10: [12, 16]
Resultado final: [12, 16]
```

## Dicas
- Pipeline: saída de um estágio é entrada do próximo
- Cada estágio é função que transforma dados
- Encadeie: `stage3(stage2(stage1(dados)))`
- Útil para processamento de dados em etapas



