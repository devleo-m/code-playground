# Atividade 50: Projeto Final - API REST Simples

## Objetivo
Criar API REST completa integrando todos os conceitos.

## Enunciado
Crie uma API REST para gerenciamento de tarefas (To-Do) com:
1. **Estruturas**: Tarefa (ID, Título, Descrição, Concluída, CriadaEm)
2. **Endpoints** (simulados, sem HTTP real):
   - `CriarTarefa(tarefa Tarefa) (Tarefa, error)`
   - `ListarTarefas() ([]Tarefa, error)`
   - `ObterTarefa(id string) (Tarefa, error)`
   - `AtualizarTarefa(id string, tarefa Tarefa) error`
   - `DeletarTarefa(id string) error`
   - `ConcluirTarefa(id string) error`
3. **Validação**: valide dados de entrada
4. **Filtros**: filtre por concluída, data, etc.
5. **Persistência**: salve em arquivo JSON
6. **Error Handling**: erros apropriados (não encontrado, inválido, etc.)
7. **Testes**: testes unitários

## Funcionalidades Extras
- Busca por título
- Ordenação (data, título)
- Estatísticas (total, concluídas, pendentes)
- Exportar/importar JSON

## Exemplo de Uso
```go
api := NovaAPI()
tarefa, _ := api.CriarTarefa(Tarefa{Titulo: "Estudar Go"})
tarefas, _ := api.ListarTarefas()
api.ConcluirTarefa(tarefa.ID)
estatisticas := api.Estatisticas()
```

## Dicas
- Separe lógica de negócio de acesso a dados
- Use interfaces para desacoplar
- Error handling consistente
- Validações em camada de serviço
- Estrutura clara e organizada

