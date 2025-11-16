# Atividade 48: Padrão Visitor

## Objetivo
Implementar padrão Visitor para operações em estruturas.

## Enunciado
Crie um programa que:
1. Defina interface `Elemento` com método `Aceitar(visitor Visitor)`
2. Defina interface `Visitor` com métodos para cada tipo de elemento
3. Defina elementos: `Arquivo`, `Pasta` (contém elementos)
4. Implemente visitors:
   - `VisitorTamanho` - calcula tamanho total
   - `VisitorListar` - lista todos arquivos
   - `VisitorBuscar` - busca por nome
5. Estrutura hierárquica: Pasta contém Arquivos e outras Pastas
6. Demonstre diferentes operações sem modificar elementos

## Exemplo de Saída
```
Estrutura:
  Pasta: documentos/
    Arquivo: doc1.txt (100 bytes)
    Pasta: imagens/
      Arquivo: foto.jpg (500 bytes)

Visitor Tamanho: 600 bytes
Visitor Listar:
  - documentos/doc1.txt
  - documentos/imagens/foto.jpg
Visitor Buscar "doc": encontrado doc1.txt
```

## Dicas
- Visitor permite adicionar operações sem modificar elementos
- Elemento aceita visitor e chama método apropriado
- Double dispatch: elemento escolhe método do visitor
- Útil para operações em estruturas complexas

