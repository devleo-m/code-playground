# Módulo 40: Build Constraints & Tags em Go
## Aula 3 - Exercícios e Reflexão

---

## Exercícios Práticos

### Exercício 1: Criar Código Multi-plataforma

Crie um programa que detecta o sistema operacional e arquitetura usando build constraints.

**Requisitos:**
1. Função `getOS() string` com implementações para Linux, Windows, macOS
2. Função `getArch() string` com implementações para amd64, arm64
3. Arquivo padrão para sistemas não suportados
4. Função `main()` que imprime OS e Arch

**Tarefa**: Crie os arquivos necessários e teste em diferentes plataformas (ou simule com GOOS/GOARCH).

---

### Exercício 2: Feature Toggle

Crie um sistema de feature toggles usando build tags.

**Requisitos:**
1. Tag `debug`: Adiciona logging detalhado
2. Tag `metrics`: Adiciona coleta de métricas
3. Tag `tls`: Adiciona suporte a TLS
4. Teste diferentes combinações de tags

**Tarefa**: Implemente e teste com diferentes combinações de `-tags`.

---

### Exercício 3: Código Específico de Arquitetura

Crie funções otimizadas para diferentes arquiteturas.

**Requisitos:**
1. Função `processData()` otimizada para amd64
2. Função `processData()` genérica para outras arquiteturas
3. Benchmark comparando versões

**Tarefa**: Implemente e compare performance.

---

## Perguntas de Reflexão

### Reflexão 1: Organização de Código Multi-plataforma

**Perguntas:**
1. Como você organizaria código que precisa rodar em 5+ plataformas?
2. Quando faz sentido ter código específico vs código genérico?
3. Como você testaria código multi-plataforma?

**Escreva suas reflexões** (mínimo 200 palavras):

---

### Reflexão 2: Feature Toggles vs Configuração

**Perguntas:**
1. Quando usar build tags vs configuração em runtime?
2. Quais são os trade-offs de cada abordagem?
3. Como você decidiria entre as duas?

**Escreva suas reflexões** (mínimo 200 palavras):

---

## Checklist

- [ ] Sei criar build constraints
- [ ] Sei usar tags de plataforma
- [ ] Sei criar tags customizadas
- [ ] Sei organizar código multi-plataforma
- [ ] Entendo quando usar build tags

---

**Bons estudos! 🚀**

