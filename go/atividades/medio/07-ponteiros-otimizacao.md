# Atividade 07: Otimização com Ponteiros

## Objetivo
Entender quando usar ponteiros para otimizar performance e evitar cópias.

## Enunciado
Crie um programa que:
1. Defina struct `DadosGrandes` com array de 1000 inteiros
2. Crie função `processarPorValor(d DadosGrandes)` que recebe por valor
3. Crie função `processarPorPonteiro(d *DadosGrandes)` que recebe por ponteiro
4. Meça tempo de execução de ambas as funções
5. Crie função `modificarSlice(s []int)` e `modificarSlicePonteiro(s *[]int)` e compare comportamento
6. Demonstre quando cópias são feitas e quando são evitadas

## Exemplo de Saída
```
Processando struct grande (1000 elementos):
  Por valor: 2.5ms (cópia completa)
  Por ponteiro: 0.1ms (sem cópia)

Modificando slice:
  Slice original modificado: [10, 20, 30] -> [100, 20, 30]
  (Slice é referência, não precisa de ponteiro para modificar elementos)
```

## Dicas
- Structs grandes: use ponteiro para evitar cópia
- Slices: já são referências, ponteiro só se quiser modificar o slice em si
- Use `time` package para medir performance
- Considere trade-off: ponteiros vs cópias


