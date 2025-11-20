# Projeto 04: Sistema de Controle de Estoque

## 📝 Descrição
Sistema completo para gerenciar estoque de produtos com entrada, saída, relatórios e alertas de estoque baixo.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Produtos**:
   - Cadastrar produto (nome, código, preço, quantidade mínima)
   - Atualizar produto
   - Listar produtos
   - Buscar produto

2. **Movimentações**:
   - Entrada de estoque (adiciona quantidade)
   - Saída de estoque (remove quantidade)
   - Histórico de movimentações

3. **Alertas**:
   - Produtos com estoque abaixo do mínimo
   - Produtos sem estoque

4. **Relatórios**:
   - Valor total do estoque
   - Produtos mais movimentados
   - Histórico de movimentações

5. **Persistência**: JSON

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Error handling
- ✅ Slices e maps
- ✅ Agregações
- ✅ JSON
- ✅ Interfaces
- ✅ Organização de código

## 📁 Estrutura Sugerida
```
estoque/
├── main.go
├── produto.go
├── movimentacao.go
├── repositorio.go
├── service.go
├── relatorio.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type Produto struct {
    ID            string
    Nome          string
    Codigo        string
    Preco         float64
    Quantidade    int
    QuantidadeMin int
}

type Movimentacao struct {
    ID        string
    ProdutoID string
    Tipo      string // "entrada" ou "saida"
    Quantidade int
    Data      time.Time
    Motivo    string
}
```

### Funcionalidades
- `CadastrarProduto(produto Produto) error`
- `AdicionarEstoque(id string, qtd int) error`
- `RemoverEstoque(id string, qtd int) error`
- `VerificarAlertas() []Alerta`
- `CalcularValorTotal() float64`
- `GerarRelatorio() Relatorio`

## ✅ Critérios de Sucesso
- [ ] CRUD de produtos funciona
- [ ] Movimentações são registradas
- [ ] Alertas funcionam corretamente
- [ ] Relatórios são precisos
- [ ] Validações impedem estoque negativo
- [ ] Dados persistem

## 🚀 Extras (Desafio)
- [ ] Categorias de produtos
- [ ] Fornecedores
- [ ] Vendas (integração)
- [ ] Gráficos de movimentação
- [ ] Exportar para Excel/CSV
- [ ] Múltiplos depósitos



