# Módulo 33: Code Generation e Build Tags em Go

Este módulo ensina como usar `go generate` e Build Tags para automação de código e compilação condicional em Go.

## 📚 Estrutura do Módulo

Este módulo segue o **Ciclo de 4 Etapas** de ensino:

1. **Aula Principal** (`aula-01-code-generation-build-tags-principal.md`)
   - Conteúdo técnico completo sobre `go generate` e Build Tags
   - Exemplos práticos e casos de uso

2. **Aula Simplificada** (`aula-02-code-generation-build-tags-simplificada.md`)
   - Mesmos conceitos explicados com analogias do dia a dia
   - Facilita a compreensão e fixação

3. **Exercícios e Reflexão** (`aula-03-exercicios-e-reflexao.md`)
   - Exercícios práticos para colocar em prática
   - Perguntas de reflexão que exigem pensamento crítico

4. **Performance e Boas Práticas** (`aula-04-performance-e-boas-praticas.md`)
   - O que fazer e o que não fazer
   - Melhores práticas profissionais
   - Armadilhas comuns e como evitá-las

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você será capaz de:

- ✅ Usar `go generate` para automação de código
- ✅ Criar e usar Build Tags para compilação condicional
- ✅ Decidir quando usar cada ferramenta
- ✅ Integrar essas ferramentas no workflow profissional
- ✅ Evitar armadilhas comuns

## 📁 Arquivos

- `aula-01-code-generation-build-tags-principal.md` - Aula principal
- `aula-02-code-generation-build-tags-simplificada.md` - Aula simplificada
- `aula-03-exercicios-e-reflexao.md` - Exercícios e reflexão
- `aula-04-performance-e-boas-praticas.md` - Performance e boas práticas
- `01-exemplos.go` - Exemplos de código
- `README.md` - Este arquivo

## 🚀 Como Usar Este Módulo

1. **Leia a Aula Principal** - Entenda os conceitos técnicos
2. **Leia a Aula Simplificada** - Fixe o aprendizado com analogias
3. **Complete os Exercícios** - Pratique o que aprendeu
4. **Leia Performance e Boas Práticas** - Aprenda o que fazer e evitar
5. **Reflita sobre as Perguntas** - Desenvolva pensamento crítico

## 🛠️ Ferramentas Necessárias

Para seguir este módulo, você precisará:

- Go 1.17+ instalado
- `stringer` tool: `go install golang.org/x/tools/cmd/stringer@latest`
- (Opcional) `mockgen`: `go install github.com/golang/mock/mockgen@latest`
- (Opcional) `protoc` para exemplos de protobuf

## 📝 Conceitos Principais

### go generate
- Automação de geração de código
- Diretivas `//go:generate`
- Ferramentas comuns (stringer, mockgen, protoc)
- Integração com workflow

### Build Tags
- Compilação condicional
- Sintaxe `//go:build`
- Tags para OS, arquitetura e customizadas
- Quando usar vs runtime checks

## 🔗 Próximos Passos

Após completar este módulo:
- Pratique criando seus próprios exemplos
- Integre `go generate` e Build Tags em projetos pessoais
- Explore outras ferramentas de geração de código
- Continue para o próximo módulo do curso

## 📚 Recursos Adicionais

- [Go Generate Documentation](https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source)
- [Build Constraints Documentation](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [stringer Tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)

---

Bons estudos! 🚀


