# Atividade 47: Padrão Command

## Objetivo
Implementar padrão Command para desacoplar requisições.

## Enunciado
Crie um programa que:
1. Defina interface `Command` com métodos `Executar() error` e `Desfazer() error`
2. Implemente comandos:
   - `ComandoAdicionar` - adiciona item a lista
   - `ComandoRemover` - remove item
   - `ComandoAtualizar` - atualiza item
3. Defina struct `Invoker` que armazena histórico de comandos
4. Implemente undo/redo: desfazer e refazer comandos
5. Suporte macro (executar múltiplos comandos)
6. Simule editor de texto simples com comandos

## Exemplo de Saída
```
Executando: Adicionar "Hello"
Lista: [Hello]
Executando: Adicionar "World"
Lista: [Hello World]
Undo: removendo "World"
Lista: [Hello]
Redo: adicionando "World"
Lista: [Hello World]
```

## Dicas
- Command encapsula ação e seus parâmetros
- Histórico permite undo/redo
- Cada comando sabe como se desfazer
- Útil para operações reversíveis


