# Projeto 10: Sistema de Biblioteca

## 📝 Descrição
Sistema completo para gerenciar biblioteca com livros, empréstimos, usuários e relatórios.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Livros**:
   - Cadastrar livro (título, autor, ISBN, quantidade)
   - Buscar livro
   - Listar livros disponíveis
   - Atualizar informações

2. **Usuários**:
   - Cadastrar usuário
   - Listar usuários

3. **Empréstimos**:
   - Emprestar livro (associar a usuário)
   - Devolver livro
   - Listar empréstimos ativos
   - Histórico de empréstimos
   - Verificar atrasos

4. **Relatórios**:
   - Livros mais emprestados
   - Usuários mais ativos
   - Livros disponíveis vs emprestados

5. **Validação**:
   - Não emprestar se sem estoque
   - Limite de empréstimos por usuário

6. **Persistência**: JSON

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Slices e maps
- ✅ Error handling
- ✅ Agregações
- ✅ Time package
- ✅ JSON
- ✅ Interfaces
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
biblioteca/
├── main.go
├── livro.go
├── usuario.go
├── emprestimo.go
├── repositorio.go
├── service.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Livro struct {
    ID          string
    Titulo      string
    Autor       string
    ISBN        string
    Quantidade  int
    Disponivel  int
}

type Usuario struct {
    ID    string
    Nome  string
    Email string
}

type Emprestimo struct {
    ID         string
    LivroID    string
    UsuarioID  string
    DataEmprestimo time.Time
    DataDevolucao  *time.Time
    Atrasado    bool
}
```

## ✅ Critérios de Sucesso
- [ ] CRUD completo funciona
- [ ] Empréstimos são gerenciados
- [ ] Validações funcionam
- [ ] Relatórios são precisos
- [ ] Dados persistem
- [ ] Código organizado

## 🚀 Extras (Desafio)
- [ ] Multas por atraso
- [ ] Reservas de livros
- [ ] Notificações de devolução
- [ ] Sistema de avaliações
- [ ] Busca avançada



