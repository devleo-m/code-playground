# Módulo 40: Build Constraints & Tags em Go

Bem-vindo ao módulo sobre **Build Constraints & Tags** em Go!

## 📚 Estrutura

- **Aula 1**: Build Constraints (Principal)
- **Aula 2**: Versão Simplificada
- **Aula 3**: Exercícios e Reflexão
- **Aula 4**: Performance e Boas Práticas

## 🚀 Início Rápido

```go
//go:build linux
package main

func getOS() string {
    return "Linux"
}
```

```bash
go build -tags debug
```

## 📖 Conceitos

- **Build Constraints**: Controlam quais arquivos compilam
- **//go:build**: Sintaxe moderna
- **Tags**: Platform-specific, feature toggles

---

**Bons estudos! 🚀**



