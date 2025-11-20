# Atividade 27: Padrão Builder

## Objetivo
Implementar padrão Builder para construção complexa de objetos.

## Enunciado
Crie um programa que:
1. Defina struct `Query` complexa com múltiplos campos opcionais
2. Defina struct `QueryBuilder` com métodos encadeados
3. Implemente métodos: `Select()`, `From()`, `Where()`, `Join()`, `OrderBy()`, `Limit()`, `Build()`
4. Permita construir query passo a passo
5. Valide query antes de construir (ex: FROM obrigatório)
6. Crie builder para struct `Email` com campos: De, Para, Assunto, Corpo, Anexos

## Exemplo de Saída
```
Query construída:
SELECT nome, email 
FROM usuarios 
WHERE idade > 18 
ORDER BY nome 
LIMIT 10

Email construído:
De: remetente@email.com
Para: destinatario@email.com
Assunto: Teste
Corpo: Mensagem...
```

## Dicas
- Builder armazena estado durante construção
- Métodos retornam builder para encadeamento
- Valide no método `Build()`
- Útil para objetos com muitos campos opcionais



