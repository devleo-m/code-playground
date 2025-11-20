# Atividade 47: Error Wrapping

## Objetivo
Praticar error wrapping para adicionar contexto.

## Enunciado
Crie um programa que:
1. Crie função `lerArquivo` que simula leitura e retorna erro
2. Crie função `processarArquivo` que chama `lerArquivo` e adiciona contexto com `fmt.Errorf` e `%w`
3. Crie função `salvarDados` que chama `processarArquivo` e adiciona mais contexto
4. No `main`, chame `salvarDados` e imprima erro completo (mostra cadeia de contexto)

## Exemplo de Saída
```
Erro completo: falha ao salvar dados: erro ao processar arquivo 'dados.txt': arquivo não encontrado
```

## Dicas
- Error wrapping: `fmt.Errorf("contexto: %w", err)`
- Use `%w` para encadear erros
- Preserva erro original enquanto adiciona contexto
- Útil para rastrear onde erro ocorreu na cadeia de chamadas



