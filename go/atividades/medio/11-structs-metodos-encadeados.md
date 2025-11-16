# Atividade 11: Métodos Encadeados (Method Chaining)

## Objetivo
Implementar padrão de method chaining para APIs fluentes.

## Enunciado
Crie um programa que:
1. Defina struct `QueryBuilder` para construir queries SQL
2. Implemente métodos que retornam `*QueryBuilder` (ponteiro):
   - `Select(campos ...string) *QueryBuilder`
   - `From(tabela string) *QueryBuilder`
   - `Where(condicao string) *QueryBuilder`
   - `OrderBy(campo string) *QueryBuilder`
   - `Build() string` (retorna query final)
3. Crie struct `StringBuilder` com métodos encadeados:
   - `Append(texto string) *StringBuilder`
   - `AppendLine(texto string) *StringBuilder`
   - `Clear() *StringBuilder`
   - `String() string`
4. Demonstre uso encadeado: `builder.Select("nome").From("usuarios").Where("idade > 18").Build()`

## Exemplo de Saída
```
Query construída:
SELECT nome, email FROM usuarios WHERE idade > 18 ORDER BY nome

StringBuilder:
Linha 1
Linha 2
Linha 3
```

## Dicas
- Métodos retornam ponteiro para permitir encadeamento
- Armazene estado na struct
- Útil para APIs fluentes e builders
- Retorne `*Tipo` para encadear: `return b`

