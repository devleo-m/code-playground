# Projeto 02: Sistema de Gerenciamento de Contatos

## 📝 Descrição
Crie um sistema completo para gerenciar contatos (agenda telefônica) com CRUD completo, busca e exportação.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **CRUD Completo**:
   - Criar contato (nome, telefone, email, endereço)
   - Listar todos contatos
   - Buscar contato (por nome, telefone ou email)
   - Atualizar contato
   - Deletar contato

2. **Validação**:
   - Email válido
   - Telefone no formato correto
   - Campos obrigatórios

3. **Persistência**: Salvar em arquivo JSON

4. **Interface**: Menu interativo no terminal

5. **Estatísticas**: Total de contatos, contatos por letra inicial

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Slices e maps
- ✅ Error handling
- ✅ JSON serialização
- ✅ Validação de dados
- ✅ I/O de arquivos
- ✅ Loops e condicionais
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
contatos/
├── main.go
├── contato.go
├── repositorio.go
├── service.go
├── validacao.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Contato struct {
    ID       string
    Nome     string
    Telefone string
    Email    string
    Endereco string
    CriadoEm time.Time
}

type Repositorio struct {
    contatos map[string]Contato
    arquivo  string
}
```

### Funcionalidades
- `CriarContato(contato Contato) error`
- `BuscarPorID(id string) (Contato, error)`
- `BuscarPorNome(nome string) []Contato`
- `ListarTodos() []Contato`
- `Atualizar(id string, contato Contato) error`
- `Deletar(id string) error`
- `ValidarContato(contato Contato) []ErroValidacao`

## ✅ Critérios de Sucesso
- [ ] CRUD completo funciona
- [ ] Validações funcionam
- [ ] Dados persistem em arquivo
- [ ] Busca retorna resultados corretos
- [ ] Interface é intuitiva
- [ ] Código organizado em pacotes

## 🚀 Extras (Desafio)
- [ ] Grupos/categorias de contatos
- [ ] Exportar para CSV
- [ ] Importar de CSV
- [ ] Busca avançada (fuzzy search)
- [ ] Histórico de alterações

