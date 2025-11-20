# Atividade 34: Serialização Customizada

## Objetivo
Implementar serialização customizada usando interfaces.

## Enunciado
Crie um programa que:
1. Defina interface `Serializavel` com métodos `Serializar() string` e `Deserializar(string) error`
2. Defina struct `Pessoa` que implementa `Serializavel` (formato: "nome|idade|email")
3. Defina struct `Produto` que implementa `Serializavel` (formato JSON-like simples)
4. Crie função genérica `Salvar(arquivo string, itens []Serializavel) error`
5. Crie função genérica `Carregar(arquivo string) ([]Serializavel, error)`
6. Suporte múltiplos tipos no mesmo arquivo (com tipo prefix)

## Exemplo de Saída
```
Salvando itens:
PESSOA|João|30|joao@email.com
PRODUTO|{"nome":"Notebook","preco":2500}
Carregando: 2 itens carregados
```

## Dicas
- Interface permite tratar diferentes tipos uniformemente
- Use prefix para identificar tipo ao deserializar
- `strings.Split` para parsing
- Type assertion para converter de volta ao tipo concreto



