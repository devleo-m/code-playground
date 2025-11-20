# Projeto 08: Gerenciador Financeiro Pessoal

## 📝 Descrição
Sistema completo para gerenciar finanças pessoais com receitas, despesas, categorias e relatórios.

## 🎯 Requisitos

### Funcionalidades Obrigatórias
1. **Transações**:
   - Registrar receita (valor, descrição, categoria, data)
   - Registrar despesa
   - Editar transação
   - Deletar transação
   - Listar transações

2. **Categorias**:
   - Criar categorias personalizadas
   - Filtrar por categoria

3. **Relatórios**:
   - Saldo atual
   - Receitas vs Despesas (período)
   - Gastos por categoria
   - Transações por mês
   - Gráficos simples (ASCII)

4. **Filtros**:
   - Por período (mês, ano)
   - Por tipo (receita/despesa)
   - Por categoria

5. **Persistência**: JSON

## 📚 Conceitos Utilizados
- ✅ Structs e métodos
- ✅ Slices e maps
- ✅ Agregações
- ✅ Filtros
- ✅ Time package
- ✅ JSON
- ✅ Error handling
- ✅ Formatação monetária

## 📁 Estrutura Sugerida
```
financas/
├── main.go
├── transacao.go
├── categoria.go
├── relatorio.go
├── repositorio.go
└── README.md
```

## 💡 Implementação Sugerida

### Structs
```go
type TipoTransacao string
const (
    Receita TipoTransacao = "receita"
    Despesa TipoTransacao = "despesa"
)

type Transacao struct {
    ID        string
    Tipo      TipoTransacao
    Valor     float64
    Categoria string
    Descricao string
    Data      time.Time
}

type Relatorio struct {
    Saldo        float64
    TotalReceitas float64
    TotalDespesas float64
    PorCategoria map[string]float64
}
```

## ✅ Critérios de Sucesso
- [ ] CRUD de transações funciona
- [ ] Cálculos são precisos
- [ ] Relatórios são corretos
- [ ] Filtros funcionam
- [ ] Dados persistem
- [ ] Interface clara

## 🚀 Extras (Desafio)
- [ ] Metas de gastos
- [ ] Orçamento mensal
- [ ] Exportar para CSV/Excel
- [ ] Gráficos mais elaborados
- [ ] Múltiplas contas
- [ ] Previsões baseadas em histórico



