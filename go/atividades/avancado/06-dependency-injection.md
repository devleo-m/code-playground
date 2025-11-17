# Atividade 06: Dependency Injection Container

## Objetivo
Criar container de dependency injection simples mas funcional.

## Enunciado
Crie um DI container que:
1. Defina interface `Container` com métodos:
   - `Register(name string, factory func() interface{})`
   - `RegisterSingleton(name string, instance interface{})`
   - `Resolve(name string) (interface{}, error)`
   - `ResolveWithInterface(name string, iface interface{}) error`
2. Suporte injeção por interface
3. Detecte dependências circulares
4. Suporte lifecycle: transient, singleton, scoped
5. Implemente auto-wiring (detecta dependências automaticamente)
6. Adicione validação: verifica se todas dependências estão registradas
7. Crie exemplo completo: API -> Service -> Repository

## Exemplo de Uso
```go
container := NewContainer()
container.RegisterSingleton("repo", NewRepository())
container.Register("service", func() interface{} {
    return NewService(container.Resolve("repo"))
})
service := container.Resolve("service").(Service)
```

## Dicas
- Map para armazenar factories e instâncias
- Recursão para resolver dependências
- Cache de singletons
- Detecção de ciclos com set de visitados
- Reflection para auto-wiring (opcional)

## Desafio Extra
Implemente scoped lifetime (por request) e decorators/interceptors.


