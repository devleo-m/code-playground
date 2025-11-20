# Módulo 32: Security - Segurança em Aplicações Go

Este módulo ensina sobre segurança em aplicações Go, com foco no **govulncheck**, o scanner oficial de vulnerabilidades do Go.

## 📚 Estrutura das Aulas

1. **aula-01-security-principal.md** - Aula principal com conteúdo técnico completo sobre govulncheck
2. **aula-02-security-simplificada.md** - Aula simplificada com analogias e explicações do dia a dia
3. **aula-03-exercicios-e-reflexao.md** - Exercícios práticos e perguntas de reflexão
4. **aula-04-performance-e-boas-praticas.md** - Boas práticas, otimizações e integração profissional

## 📁 Arquivos de Exemplo

- **01-exemplos.go** - Exemplos de código Go seguro demonstrando boas práticas
- **02-exemplos-com-vulnerabilidades.go** - Exemplos INTENCIONAIS de código vulnerável para fins educacionais

⚠️ **ATENÇÃO**: O arquivo `02-exemplos-com-vulnerabilidades.go` contém código vulnerável intencionalmente para fins educacionais. NÃO use esse código em produção!

## 🚀 Como Usar

### Instalar govulncheck

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
```

### Verificar Instalação

```bash
govulncheck -version
```

### Executar Verificações

```bash
# Verificar o projeto atual
govulncheck ./...

# Verificar apenas dependências
govulncheck -mode=mod ./...

# Verificar binário compilado
go build -o app .
govulncheck -mode=binary ./app

# Formato JSON (útil para CI/CD)
govulncheck -json ./... > vulnerabilities.json
```

### Testar os Exemplos

```bash
# Verificar exemplos seguros
govulncheck 01-exemplos.go

# Verificar exemplos com vulnerabilidades (educacional)
govulncheck 02-exemplos-com-vulnerabilidades.go
```

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você deve ser capaz de:

- ✅ Explicar o que é govulncheck e por que é importante
- ✅ Instalar e usar o govulncheck em projetos Go
- ✅ Entender os diferentes modos de operação (source, binary, mod)
- ✅ Interpretar resultados de vulnerabilidades
- ✅ Corrigir vulnerabilidades encontradas
- ✅ Integrar govulncheck no workflow de desenvolvimento
- ✅ Configurar govulncheck no CI/CD
- ✅ Documentar decisões de segurança
- ✅ Criar políticas de segurança para projetos
- ✅ Priorizar correções por severidade

## 📖 Conceitos Principais

### govulncheck

O **govulncheck** é o scanner oficial de vulnerabilidades do Go. Ele:

- Verifica código e dependências contra vulnerabilidades conhecidas
- Usa o Go Vulnerability Database (banco de dados oficial)
- Fornece informações sobre severidade e como corrigir
- Oferece três modos de operação: source, binary, mod

### Modos de Operação

1. **Modo Source (padrão)**: Analisa código-fonte e mostra apenas vulnerabilidades que você realmente usa
2. **Modo Binary**: Analisa binário compilado
3. **Modo Module**: Analisa apenas dependências do módulo

### Severidade de Vulnerabilidades

- **CRITICAL**: Vulnerabilidades críticas que precisam correção imediata
- **HIGH**: Vulnerabilidades de alta severidade que precisam atenção urgente
- **MEDIUM**: Vulnerabilidades médias que devem ser corrigidas
- **LOW**: Vulnerabilidades baixas que podem ser corrigidas quando possível

## 🔧 Integração com CI/CD

### GitHub Actions

```yaml
name: Security Scan
on: [push, pull_request]
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Install govulncheck
        run: go install golang.org/x/vuln/cmd/govulncheck@latest
      - name: Run govulncheck
        run: govulncheck ./...
```

### Pre-commit Hook

```bash
#!/bin/sh
# .git/hooks/pre-commit

echo "🔍 Verificando vulnerabilidades..."
govulncheck ./...

if [ $? -ne 0 ]; then
    echo "❌ Vulnerabilidades encontradas!"
    exit 1
fi

echo "✅ Nenhuma vulnerabilidade encontrada"
```

## 📊 Boas Práticas

1. **Execute regularmente**: Integre no workflow diário
2. **Use modo apropriado**: Escolha o modo baseado na necessidade
3. **Priorize por severidade**: Corrija CRITICAL e HIGH primeiro
4. **Documente decisões**: Sempre documente quando não pode corrigir
5. **Integre no CI/CD**: Automatize verificações
6. **Educar equipe**: Todos devem conhecer e usar govulncheck

## 🔗 Recursos Adicionais

- [Documentação do govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
- [Go Vulnerability Database](https://vuln.go.dev/)
- [Go Security Best Practices](https://go.dev/doc/security/best-practices)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)

## ⚠️ Avisos Importantes

1. **Código Vulnerável**: O arquivo `02-exemplos-com-vulnerabilidades.go` contém código vulnerável intencionalmente para fins educacionais. NUNCA use esse código em produção!

2. **Atualizações**: O Go Vulnerability Database é atualizado regularmente. Execute `govulncheck` regularmente para obter informações atualizadas.

3. **Não é Tudo**: O govulncheck verifica vulnerabilidades conhecidas. Ele não substitui:
   - Boas práticas de desenvolvimento seguro
   - Code reviews
   - Testes de segurança
   - Outras ferramentas de segurança (gosec, etc.)

## 🎓 Próximos Passos

Após completar este módulo:

1. Integre govulncheck no seu workflow diário
2. Configure no CI/CD dos seus projetos
3. Crie políticas de segurança para sua equipe
4. Explore outras ferramentas de segurança (gosec, nancy, etc.)
5. Continue aprendendo sobre segurança em Go

---

**Lembre-se**: Segurança não é um destino, é uma jornada contínua! 🔒



