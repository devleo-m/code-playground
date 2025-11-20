# Módulo 43: Plugins & Dynamic Loading em Go

Bem-vindo ao módulo sobre **Plugins & Dynamic Loading** - a última aula de tópicos avançados!

## 📚 Estrutura

- **Aula 1**: Plugins & Dynamic Loading (Principal)
- **Aula 2**: Versão Simplificada
- **Aula 3**: Exercícios e Reflexão
- **Aula 4**: Performance e Boas Práticas

## 🚀 Início Rápido

```go
// plugin.go
package main

func Greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

```bash
# Compilar plugin
go build -buildmode=plugin -o plugin.so plugin.go

# Carregar e usar
p, _ := plugin.Open("plugin.so")
greetFunc, _ := p.Lookup("Greet")
greet := greetFunc.(func(string))
greet("World")
```

## ⚠️ Limitações

- Unix-only (não funciona no Windows)
- Plugin e app devem usar mesma versão do Go
- Complexidade adicional
- Pouco usado na comunidade

## 💡 Alternativas

- Interfaces e Injeção de Dependência
- RPC/HTTP services
- Scripting languages

---

**Bons estudos! 🚀**

---

**🎉 Parabéns por completar todas as aulas de tópicos avançados em Go!**



