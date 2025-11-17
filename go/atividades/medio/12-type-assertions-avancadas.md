# Atividade 12: Type Assertions e Type Switches Avançados

## Objetivo
Dominar type assertions e type switches para trabalhar com interfaces.

## Enunciado
Crie um programa que:
1. Defina interface `Processavel` com método `Processar()`
2. Defina structs `DadosTexto`, `DadosNumericos`, `DadosJSON` que implementam `Processavel`
3. Crie função `processarGenerico(dados interface{})` que:
   - Usa type switch para identificar tipo
   - Processa cada tipo de forma específica
   - Retorna resultado formatado
4. Crie função `extrairValor(dados interface{}) (interface{}, error)` que:
   - Usa type assertion para extrair valor concreto
   - Retorna erro se tipo não suportado
5. Crie função `converterParaString(dados interface{}) string` que converte qualquer tipo para string

## Exemplo de Saída
```
Processando dados:
  Tipo: DadosTexto -> Processado: "OLÁ, MUNDO"
  Tipo: DadosNumericos -> Processado: Soma = 150
  Tipo: DadosJSON -> Processado: {"chave": "valor"}

Extraindo valor de DadosNumericos: [10, 20, 30]
Convertendo para string: "10,20,30"
```

## Dicas
- Type switch: `switch v := valor.(type) { case Tipo1: ... }`
- Type assertion segura: `valor, ok := i.(Tipo)`
- Use type switch quando há múltiplos tipos possíveis
- Type assertion quando sabe o tipo específico


