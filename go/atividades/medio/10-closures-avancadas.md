# Atividade 10: Closures Avançadas

## Objetivo
Dominar closures para criar funções especializadas e manter estado.

## Enunciado
Crie um programa que:
1. Crie função `criarMultiplicador(fator int) func(int) int` que retorna função que multiplica
2. Crie função `criarAcumulador() func(int) int` que mantém soma interna
3. Crie função `criarContadorLimite(limite int) func() (int, bool)` que incrementa até limite, retorna (valor, atingiuLimite)
4. Crie função `criarCache() func(string) (string, bool)` que armazena valores (simula cache simples)
5. Demonstre uso de cada closure e como estado é mantido entre chamadas

## Exemplo de Saída
```
Multiplicador por 5:
  2 * 5 = 10
  7 * 5 = 35

Acumulador:
  +10 = 10
  +20 = 30
  +5 = 35

Contador com limite 3:
  1 (false)
  2 (false)
  3 (true - limite atingido)

Cache:
  Buscar 'chave1': não encontrado, armazenando 'valor1'
  Buscar 'chave1': encontrado 'valor1'
```

## Dicas
- Closure captura variáveis do escopo externo
- Estado persiste entre chamadas da função retornada
- Cada chamada de `criarX` cria novo estado independente
- Útil para funções especializadas e padrões de design


