# Módulo 35: Deployment & Tooling em Go

Bem-vindo ao módulo sobre **Deployment & Tooling** em Go! Este módulo ensina como compilar executáveis, fazer cross-compilation e preparar aplicações Go para produção.

## 📚 Estrutura do Módulo

Este módulo está dividido em 4 aulas principais:

### Aula 1: Building Executables e Cross-compilation (Principal)
**Arquivo**: `aula-01-deployment-tooling-principal.md`

Conteúdo completo e detalhado sobre:
- O comando `go build` e suas flags
- Como criar executáveis standalone
- Cross-compilation para diferentes plataformas
- Otimizações de build
- Scripts e automação

**Tempo estimado**: 2-3 horas

---

### Aula 2: Versão Simplificada com Analogias
**Arquivo**: `aula-02-deployment-tooling-simplificada.md`

Explicações simplificadas com analogias do dia a dia:
- Building como "transformar receita em bolo pronto"
- Cross-compilation como "tradutor universal"
- Conceitos visuais e fáceis de entender

**Tempo estimado**: 1 hora

---

### Aula 3: Exercícios e Reflexão
**Arquivo**: `aula-03-exercicios-e-reflexao.md`

Exercícios práticos para fixar o aprendizado:
- 7 exercícios práticos progressivos
- Questões para reflexão
- Desafios avançados
- Checklist de aprendizado

**Tempo estimado**: 2-3 horas

---

### Aula 4: Performance e Boas Práticas
**Arquivo**: `aula-04-performance-e-boas-praticas.md`

Otimizações e melhores práticas:
- Otimizações de build
- Redução de tamanho de binários
- Segurança em builds
- Deploy e distribuição
- Checklist de boas práticas

**Tempo estimado**: 1-2 horas

---

## 💻 Exemplos Práticos

**Arquivo**: `01-exemplos.go`

Contém 7 exemplos práticos que demonstram:
1. Build básico
2. Build com informações de versão
3. Informações de plataforma
4. Comparação de tamanhos
5. Código específico de plataforma
6. Verificação de CGO
7. Build reproduzível

**Como usar:**
```bash
# Ver menu de exemplos
go run 01-exemplos.go

# Executar exemplo específico
go run 01-exemplos.go exemplo1
go run 01-exemplos.go exemplo2 version
```

---

## 🚀 Início Rápido

### 1. Build Básico

```bash
# Compilar programa
go build -o minha-app main.go

# Executar
./minha-app        # Linux/Mac
# ou
minha-app.exe      # Windows
```

### 2. Build com Versão

```bash
# Linux/Mac
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date) -X main.GitCommit=$(git rev-parse --short HEAD)" -o minha-app

# Windows (PowerShell)
$env:VERSION="1.0.0"
$env:BUILDTIME=(Get-Date)
$env:COMMIT=(git rev-parse --short HEAD)
go build -ldflags "-X main.Version=$env:VERSION -X main.BuildTime=$env:BUILDTIME -X main.GitCommit=$env:COMMIT" -o minha-app.exe
```

### 3. Cross-compilation

```bash
# Compilar para Linux (de Mac/Windows)
GOOS=linux GOARCH=amd64 go build -o minha-app-linux main.go

# Compilar para Windows (de Mac/Linux)
GOOS=windows GOARCH=amd64 go build -o minha-app-windows.exe main.go

# Compilar para macOS (de Linux/Windows)
GOOS=darwin GOARCH=amd64 go build -o minha-app-mac main.go
```

### 4. Build Otimizado

```bash
# Build menor e mais rápido
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o minha-app main.go
```

---

## 📖 Conceitos Principais

### Building Executables

- **`go build`**: Compila código Go em executável nativo
- **`-o`**: Especifica nome do executável
- **`-ldflags`**: Passa flags para o linker (útil para versão)
- **`-s -w`**: Remove símbolos de debug (binário menor)
- **`-trimpath`**: Remove caminhos do sistema de arquivos

### Cross-compilation

- **`GOOS`**: Sistema operacional de destino (linux, darwin, windows)
- **`GOARCH`**: Arquitetura de destino (amd64, arm64, 386)
- **`CGO_ENABLED=0`**: Desabilita CGO (facilita cross-compilation)

### Plataformas Comuns

| GOOS    | GOARCH | Descrição                    |
|---------|--------|------------------------------|
| linux   | amd64  | Linux 64-bit (servidores)    |
| linux   | arm64  | Linux ARM64 (Raspberry Pi 4) |
| darwin  | amd64  | macOS Intel                  |
| darwin  | arm64  | macOS Apple Silicon (M1/M2)  |
| windows | amd64  | Windows 64-bit               |

---

## 🎯 Objetivos de Aprendizado

Ao final deste módulo, você será capaz de:

- ✅ Compilar programas Go em executáveis standalone
- ✅ Criar executáveis com informações de versão
- ✅ Fazer cross-compilation para diferentes plataformas
- ✅ Otimizar tamanho e performance de binários
- ✅ Criar scripts de build automatizados
- ✅ Entender boas práticas de deployment
- ✅ Preparar aplicações para produção

---

## 📝 Checklist de Progresso

Marque conforme avança:

- [ ] Li a aula principal (aula-01)
- [ ] Li a aula simplificada (aula-02)
- [ ] Completei os exercícios (aula-03)
- [ ] Li sobre boas práticas (aula-04)
- [ ] Executei os exemplos práticos
- [ ] Compilei meu primeiro executável
- [ ] Fiz cross-compilation para outra plataforma
- [ ] Criei um script de build multiplataforma
- [ ] Entendi quando usar cada técnica

---

## 🔗 Recursos Adicionais

### Documentação Oficial

- [Go Build Command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [Cross Compilation](https://golang.org/doc/install/source#environment)
- [Build Constraints](https://golang.org/pkg/go/build/#hdr-Build_Constraints)

### Ferramentas Úteis

- **goreleaser**: Ferramenta para releases automatizados
- **Make**: Automação de builds
- **Docker**: Testar binários cross-compilados

### Comandos Úteis

```bash
# Ver todas as plataformas suportadas
go tool dist list

# Ver informações do ambiente Go
go env

# Ver ajuda do comando build
go help build

# Limpar cache de build
go clean -cache

# Ver tamanho do binário
ls -lh minha-app        # Linux/Mac
dir minha-app.exe       # Windows
```

---

## 🐛 Troubleshooting

### Problema: Binário cross-compilado não funciona

**Solução**: Teste em ambiente similar (Docker, QEMU) ou compile na plataforma de destino.

### Problema: Binário muito grande

**Solução**: Use `-ldflags "-s -w"` e `CGO_ENABLED=0`.

### Problema: Erro ao fazer cross-compilation com CGO

**Solução**: Desabilite CGO (`CGO_ENABLED=0`) ou use toolchain específico.

### Problema: Build muito lento

**Solução**: Verifique se está usando cache do Go. Use `go build -x` para ver o que está sendo recompilado.

---

## 📚 Próximos Módulos

Depois de dominar deployment e tooling, você pode avançar para:

- **Módulo 36**: Containerização e Docker
- **Módulo 37**: CI/CD com GitHub Actions
- **Módulo 38**: Monitoramento e Observabilidade
- **Módulo 39**: Deploy em Cloud (AWS, GCP, Azure)

---

## 💡 Dicas Finais

1. **Sempre teste binários cross-compilados** antes de distribuir
2. **Inclua informações de versão** em todos os binários de produção
3. **Use scripts/Makefiles** para automatizar builds
4. **Mantenha builds reproduzíveis** quando possível
5. **Nunca inclua secrets** no binário

---

## 📞 Suporte

Se tiver dúvidas ou problemas:

1. Revise a aula simplificada (aula-02)
2. Consulte os exemplos práticos (01-exemplos.go)
3. Verifique a seção de troubleshooting
4. Consulte a documentação oficial do Go

---

**Bons estudos e happy building! 🚀**


