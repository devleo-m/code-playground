# Módulo 30: Code Quality and Analysis

Este módulo ensina sobre ferramentas essenciais para garantir qualidade de código em Go: `go vet` e `goimports`.

## 📚 Estrutura das Aulas

1. **aula-01-code-quality-analysis-principal.md** - Aula principal com conteúdo técnico completo
2. **aula-02-code-quality-analysis-simplificada.md** - Aula simplificada com analogias
3. **aula-03-exercicios-e-reflexao.md** - Exercícios práticos e perguntas de reflexão
4. **aula-04-performance-e-boas-praticas.md** - Performance, boas práticas e vida profissional

## 📁 Arquivos de Exemplo

- **01-exemplos.go** - Exemplos corretos demonstrando boas práticas
- **02-exemplos-com-problemas.go** - Exemplos com problemas intencionais para praticar com `go vet`
- **03-exemplo-goimports.go** - Exemplo demonstrando como `goimports` funciona

## 🚀 Como Usar

### Instalar goimports

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

### Executar Verificações

```bash
# Formatar código
goimports -w .

# Verificar qualidade
go vet ./...

# Ou usar o Makefile
make quality
```

### Testar os Exemplos

```bash
# Ver problemas no arquivo de exemplo
go vet 02-exemplos-com-problemas.go

# Ver como goimports funciona
goimports -d 03-exemplo-goimports.go  # Ver diff
goimports -w 03-exemplo-goimports.go  # Aplicar mudanças
```

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você deve ser capaz de:

- ✅ Explicar o que `go vet` faz e por que é importante
- ✅ Listar tipos de problemas que `go vet` detecta
- ✅ Usar `goimports` para gerenciar imports automaticamente
- ✅ Configurar essas ferramentas no seu editor
- ✅ Integrar essas ferramentas no workflow de desenvolvimento
- ✅ Criar scripts e automações para qualidade de código

## 📖 Recursos Adicionais

- [Documentação do go vet](https://pkg.go.dev/cmd/vet)
- [Documentação do goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## 💡 Dicas

1. Configure `goimports` no seu editor para executar ao salvar
2. Execute `go vet` antes de cada commit
3. Use hooks de pre-commit para automatizar verificações
4. Integre verificações no CI/CD
5. Nunca ignore avisos do `go vet` - eles geralmente indicam bugs reais

---

**Bons estudos!** 🚀

