# Atividade 31: Múltiplas Interfaces

## Objetivo
Praticar tipos que satisfazem múltiplas interfaces.

## Enunciado
Crie um programa que:
1. Defina interface `Leitor` com método `Ler() string`
2. Defina interface `Escritor` com método `Escrever(string)`
3. Defina struct `Arquivo` que implementa ambas as interfaces
4. Crie função `processar` que recebe `Leitor` e `Escritor` separadamente
5. No `main`, crie arquivo, escreva conteúdo, leia e imprima

## Exemplo de Saída
```
Escrevendo: Olá, Go!
Lendo: Olá, Go!
```

## Dicas
- Tipo pode implementar múltiplas interfaces
- Função pode receber interface como parâmetro
- Implemente métodos necessários para satisfazer interface
- Interfaces são satisfeitas implicitamente


