# Atividade 50: Projeto Final - Sistema de Biblioteca

## Objetivo
Integrar todos os conceitos aprendidos em um projeto completo.

## Enunciado
Crie um sistema simples de biblioteca com:

### Estruturas
1. **Livro**: Titulo (string), Autor (string), ISBN (string), Disponivel (bool)
2. **Usuario**: Nome (string), Email (string), LivrosEmprestados ([]string - ISBNs)

### Funcionalidades
1. **Métodos para Livro**:
   - `Emprestar()` - marca como não disponível
   - `Devolver()` - marca como disponível
   - `String() string` - implementa fmt.Stringer

2. **Métodos para Usuario**:
   - `EmprestarLivro(livro *Livro) error` - empresta livro (verifica disponibilidade)
   - `DevolverLivro(isbn string) error` - devolve livro
   - `ListarLivros()` - imprime livros emprestados

3. **Funções auxiliares**:
   - `BuscarLivro(biblioteca map[string]*Livro, isbn string) (*Livro, error)` - busca livro
   - `ListarLivrosDisponiveis(biblioteca map[string]*Livro)` - lista livros disponíveis

### Main
1. Crie map de livros (chave: ISBN)
2. Crie usuário
3. Empreste 2 livros
4. Liste livros do usuário
5. Devolva 1 livro
6. Liste livros disponíveis

## Exemplo de Saída
```
=== Sistema de Biblioteca ===

Livro emprestado: Aprendendo Go - Autor Exemplo (ISBN: 123-456)
Livro emprestado: Go Avançado - Outro Autor (ISBN: 987-654)

Livros emprestados por João Silva:
  - Aprendendo Go (123-456)
  - Go Avançado (987-654)

Livro devolvido: Aprendendo Go

Livros disponíveis:
  - Aprendendo Go - Autor Exemplo (ISBN: 123-456)
```

## Dicas
- Use ponteiros para modificar structs
- Retorne erros quando operação não for possível
- Use map para armazenar livros (chave: ISBN)
- Implemente fmt.Stringer para impressão bonita
- Organize código em funções e métodos

## Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Ponteiros
- ✅ Slices
- ✅ Maps
- ✅ Error handling
- ✅ Interfaces (fmt.Stringer)
- ✅ Funções
- ✅ Organização de código



