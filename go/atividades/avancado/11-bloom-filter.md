# Atividade 11: Bloom Filter

## Objetivo
Implementar Bloom Filter probabilístico para testes de pertinência eficientes.

## Enunciado
Crie um Bloom Filter completo:
1. Defina struct `BloomFilter` com bit array e hash functions
2. Implemente múltiplas hash functions (3-5 funções diferentes)
3. Métodos:
   - `Adicionar(item string)`
   - `PodeExistir(item string) bool` - retorna false se definitivamente não existe
   - `PodeNaoExistir(item string) bool` - retorna true se pode não existir
4. Calcule tamanho ótimo baseado em taxa de falso positivo desejada
5. Implemente contagem de elementos aproximada
6. Adicione métricas: taxa de falso positivo real, uso de memória
7. Compare com hash set em termos de memória e performance
8. Teste com 1.000.000 de elementos

## Exemplo de Saída
```
Bloom Filter: capacidade 1M, taxa falso positivo 1%
Tamanho bit array: 9.6MB
Hash functions: 7
Adicionando 1.000.000 elementos...
Tempo inserção: 2.3s
Taxa falso positivo real: 0.98%
Memória vs HashSet: 9.6MB vs 120MB (92% economia)
```

## Dicas
- Bit array para armazenar
- Múltiplas hash functions reduzem colisões
- Falsos positivos possíveis, falsos negativos não
- Tamanho: m = -n*ln(p) / (ln(2)^2)
- Hash functions: número ótimo = m/n * ln(2)

## Desafio Extra
Implemente Counting Bloom Filter (permite remoção) e Scalable Bloom Filter.

