# Módulo 29: Comandos Core do Go

Este módulo ensina os **9 comandos core** do Go que todo desenvolvedor precisa dominar para trabalhar eficientemente com a linguagem.

## 📚 Estrutura do Módulo

Este módulo segue a metodologia de ensino em 4 etapas:

1. **Aula Principal** (`aula-01-core-go-commands-principal.md`)
   - Conteúdo técnico completo sobre os 9 comandos
   - Exemplos práticos e flags importantes
   - Quando usar cada comando

2. **Aula Simplificada** (`aula-02-core-go-commands-simplificada.md`)
   - Mesmos conceitos com analogias do dia a dia
   - Exemplos práticos e didáticos
   - Facilita a compreensão e fixação

3. **Exercícios e Reflexão** (`aula-03-exercicios-e-reflexao.md`)
   - 4 exercícios práticos completos
   - 2 perguntas de reflexão profundas
   - Aplicação prática dos conhecimentos

4. **Performance e Boas Práticas** (`aula-04-performance-e-boas-praticas.md`)
   - Como usar cada comando de forma eficiente
   - Workflows otimizados
   - Erros comuns e como evitá-los

## 🛠️ Comandos Abordados

1. **`go run`** - Compilar e executar em um passo
2. **`go build`** - Compilar para binários
3. **`go install`** - Instalar ferramentas e pacotes
4. **`go fmt`** - Formatação automática de código
5. **`go mod`** - Gerenciamento de módulos e dependências
6. **`go test`** - Executar testes
7. **`go clean`** - Limpar arquivos de build
8. **`go doc`** - Documentação interativa
9. **`go version`** - Informações da versão

## 📁 Arquivos do Módulo

```
29-core-go-commands/
├── README.md                                    # Este arquivo
├── aula-01-core-go-commands-principal.md        # Aula principal técnica
├── aula-02-core-go-commands-simplificada.md    # Aula com analogias
├── aula-03-exercicios-e-reflexao.md            # Exercícios práticos
├── aula-04-performance-e-boas-praticas.md      # Performance e boas práticas
├── 01-exemplos.go                              # Exemplos de código
└── 01-exemplos_test.go                         # Testes de exemplo
```

## 🚀 Como Usar Este Módulo

### Ordem Recomendada de Estudo

1. **Leia a Aula Principal** para entender os conceitos técnicos
2. **Leia a Aula Simplificada** para fixar com analogias
3. **Execute os Exemplos** usando os comandos aprendidos:
   ```bash
   go run 01-exemplos.go
   go build -o exemplos 01-exemplos.go
   go test -v
   go doc .
   ```
4. **Complete os Exercícios** em `aula-03-exercicios-e-reflexao.md`
5. **Estude Performance e Boas Práticas** para otimizar seu workflow

### Comandos Rápidos para Praticar

```bash
# Executar exemplos
go run 01-exemplos.go

# Compilar
go build -o exemplos 01-exemplos.go

# Executar testes
go test -v
go test -cover
go test -bench=.

# Ver documentação
go doc .
go doc Soma
go doc -src Divide

# Formatar código
go fmt ./...

# Ver versão
go version
```

## 📖 Recursos Adicionais

- [Documentação Oficial do Go](https://go.dev/doc/)
- [Go Command Documentation](https://pkg.go.dev/cmd/go)
- [Effective Go](https://go.dev/doc/effective_go)

## ✅ Objetivos de Aprendizado

Ao final deste módulo, você será capaz de:

- ✅ Entender quando usar cada comando core do Go
- ✅ Compilar e executar programas Go eficientemente
- ✅ Gerenciar dependências com `go mod`
- ✅ Executar e criar testes
- ✅ Formatar código seguindo padrões
- ✅ Explorar documentação usando `go doc`
- ✅ Otimizar seu workflow de desenvolvimento
- ✅ Evitar erros comuns

## 🎯 Próximos Passos

Após dominar estes comandos core, você estará pronto para:
- Trabalhar em projetos Go profissionais
- Contribuir para projetos open source
- Criar suas próprias ferramentas CLI
- Gerenciar projetos Go complexos

---

**Bons estudos!** 🚀



