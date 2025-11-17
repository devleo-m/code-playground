# Atividade 29: Logger Customizado

## Objetivo
Criar sistema de logging customizado com níveis.

## Enunciado
Crie um programa que:
1. Defina níveis: DEBUG, INFO, WARN, ERROR
2. Defina struct `Logger` com nível mínimo e destino (io.Writer)
3. Implemente métodos: `Debug()`, `Info()`, `Warn()`, `Error()` que formatam e escrevem
4. Adicione timestamp e nível na mensagem
5. Permita definir formato de saída
6. Crie logger que escreve em arquivo e outro que escreve em stdout
7. Implemente rotação de logs (simplificado)

## Exemplo de Saída
```
[2024-01-15 10:30:45] [INFO] Aplicação iniciada
[2024-01-15 10:30:46] [DEBUG] Processando requisição
[2024-01-15 10:30:47] [ERROR] Erro ao conectar ao banco
```

## Dicas
- Use `time.Now().Format()` para timestamp
- `io.Writer` permite flexibilidade (arquivo, stdout, etc)
- Verifique nível antes de escrever
- Use `fmt.Fprintf` para escrever formatado


