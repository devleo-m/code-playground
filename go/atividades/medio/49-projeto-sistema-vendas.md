# Atividade 49: Projeto - Sistema de Vendas

## Objetivo
Integrar múltiplos conceitos em sistema completo.

## Enunciado
Crie um sistema de vendas completo com:
1. **Estruturas**: Produto, Cliente, Venda, ItemVenda
2. **Repositório**: gerencie produtos e vendas em memória (maps)
3. **Validação**: valide dados antes de criar venda
4. **Cálculos**: descontos, impostos, totais
5. **Relatórios**:
   - Vendas por período
   - Produtos mais vendidos
   - Clientes que mais compram
   - Receita total
6. **Persistência**: salve/carregue dados em JSON
7. **Testes**: testes unitários para funções principais

## Funcionalidades
- Cadastrar produtos
- Registrar vendas
- Aplicar descontos (fixo, percentual, cupom)
- Gerar relatórios
- Exportar dados
- Validações completas

## Exemplo de Saída
```
=== Sistema de Vendas ===
Produto cadastrado: Notebook (R$ 2500.00)
Venda registrada: #001 - Cliente: João - Total: R$ 2500.00
Relatório do mês:
  Total de vendas: 15
  Receita: R$ 37.500,00
  Produto mais vendido: Notebook (8 unidades)
```

## Dicas
- Organize em pacotes (opcional)
- Use interfaces para repositórios
- Error handling robusto
- Structs bem definidas
- Funções pequenas e focadas



