# Projeto 03: Todo List CLI

## 📝 Descrição
Crie um aplicativo de lista de tarefas (To-Do) completo com terminal interface, prioridades, categorias e filtros.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Gerenciamento de Tarefas**:
   - Adicionar tarefa (título, descrição, prioridade)
   - Marcar como concluída
   - Editar tarefa
   - Deletar tarefa
   - Listar tarefas (todas, pendentes, concluídas)

2. **Prioridades**: Baixa, Média, Alta

3. **Categorias**: Organize tarefas por categoria

4. **Filtros**:
   - Por status (pendente/concluída)
   - Por prioridade
   - Por categoria
   - Por data

5. **Persistência**: JSON

6. **Estatísticas**: Total, concluídas, pendentes, por categoria

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Interfaces
- ✅ Error handling
- ✅ Slices e maps
- ✅ Filtros e ordenação
- ✅ JSON
- ✅ Time package
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
todo/
├── main.go
├── todo.go
├── categoria.go
├── filtro.go
├── storage.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Prioridade int
const (
    Baixa Prioridade = iota
    Media
    Alta
)

type Tarefa struct {
    ID          string
    Titulo      string
    Descricao   string
    Prioridade  Prioridade
    Categoria   string
    Concluida   bool
    CriadaEm    time.Time
    ConcluidaEm *time.Time
}

type TodoList struct {
    tarefas []Tarefa
}
```

### Funcionalidades
- `AdicionarTarefa(tarefa Tarefa) error`
- `ConcluirTarefa(id string) error`
- `ListarTarefas(filtro Filtro) []Tarefa`
- `FiltrarPorStatus(concluida bool) []Tarefa`
- `FiltrarPorPrioridade(prioridade Prioridade) []Tarefa`
- `Estatisticas() Estatisticas`

## ✅ Critérios de Sucesso
- [ ] Todas operações CRUD funcionam
- [ ] Filtros funcionam corretamente
- [ ] Prioridades são respeitadas
- [ ] Dados persistem
- [ ] Interface clara
- [ ] Código testável

## 🚀 Extras (Desafio)
- [ ] Data de vencimento
- [ ] Lembretes
- [ ] Subtarefas
- [ ] Tags múltiplas
- [ ] Exportar relatório
- [ ] Modo interativo melhorado

