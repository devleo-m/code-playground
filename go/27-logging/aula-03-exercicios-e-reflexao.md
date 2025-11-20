# Aula 27 - Exercícios e Reflexão: Logging

Olá, futuro(a) Gopher!

Agora é hora de colocar em prática o que aprendemos sobre logging! Vamos fazer alguns exercícios práticos e depois refletir sobre os conceitos.

---

## Exercício 1: Implementando Logging Básico com `slog`

Crie um programa que simula uma aplicação de e-commerce simples. O programa deve:

1. Usar `slog` para logging estruturado
2. Simular os seguintes eventos (com logs apropriados):
   - Inicialização da aplicação (INFO)
   - Usuário fazendo login (INFO com contexto: usuário, IP)
   - Adicionando produto ao carrinho (DEBUG)
   - Processando pagamento (INFO)
   - Erro ao processar pagamento (ERROR com detalhes)
   - Aviso de estoque baixo (WARN)

**Requisitos:**
- Use formato JSON para os logs
- Adicione contexto relevante a cada log
- Use níveis de log apropriados
- Inclua informações como: timestamp, nível, mensagem e campos estruturados

**Dica**: Comece criando um logger com `slog.NewJSONHandler()` e depois adicione logs para cada evento simulado.

---

## Exercício 2: Migrando de `log` Padrão para `slog`

Você recebeu um código legado que usa o pacote `log` padrão. Sua tarefa é migrá-lo para `slog` mantendo a mesma funcionalidade, mas melhorando a estruturação.

**Código Original:**
```go
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Servidor iniciando na porta 8080")
	
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Requisição recebida: %s %s", r.Method, r.URL.Path)
		
		// Simular processamento
		time.Sleep(100 * time.Millisecond)
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"users":[]}`))
		
		duration := time.Since(start)
		log.Printf("Requisição processada em %v", duration)
	})
	
	log.Println("Servidor iniciado")
	http.ListenAndServe(":8080", nil)
}
```

**Sua Tarefa:**
1. Converta todos os `log.Println` e `log.Printf` para `slog` estruturado
2. Adicione contexto relevante (método HTTP, path, status, duração)
3. Use níveis de log apropriados
4. Configure o logger para usar formato JSON

**Resultado Esperado**: Logs estruturados em JSON com todas as informações relevantes da requisição.

---

## Exercício 3: Implementando Logger com Zerolog

Crie um sistema de logging para uma aplicação de chat usando Zerolog. O sistema deve:

1. Ter diferentes loggers para diferentes componentes:
   - `authLogger` - para autenticação
   - `messageLogger` - para mensagens
   - `connectionLogger` - para conexões

2. Simular os seguintes eventos:
   - Usuário conectando (INFO)
   - Usuário autenticando (INFO com contexto)
   - Mensagem enviada (DEBUG)
   - Mensagem recebida (DEBUG)
   - Erro de conexão (ERROR)
   - Aviso de muitas conexões simultâneas (WARN)

3. Configure o logger para:
   - Usar formato "pretty" no console para desenvolvimento
   - Adicionar um campo "component" automaticamente a cada logger
   - Filtrar logs DEBUG em produção (usar nível INFO)

**Requisitos:**
- Use `zerolog.ConsoleWriter` para desenvolvimento
- Crie sub-loggers com `.With()` para cada componente
- Use a API fluente do Zerolog

---

## Exercício 4: Middleware de Logging com Zap

Crie um middleware de logging para uma API HTTP usando Zap. O middleware deve:

1. Capturar informações de cada requisição:
   - Método HTTP
   - Path
   - IP do cliente
   - User-Agent
   - Status code da resposta
   - Tempo de processamento

2. Usar a API Structured do Zap (alta performance)

3. Criar um logger com contexto da requisição que pode ser usado nos handlers

4. Logar requisições com diferentes níveis baseado no status:
   - 2xx → INFO
   - 4xx → WARN
   - 5xx → ERROR

**Estrutura Sugerida:**
```go
func loggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Seu código aqui
		})
	}
}
```

**Dica**: Você precisará criar um `ResponseWriter` customizado para capturar o status code.

---

## Perguntas de Reflexão

### Reflexão 1: Por que Logging Estruturado?

Pense sobre a seguinte situação:

Você tem uma aplicação em produção que recebe 1 milhão de requisições por dia. Um cliente reporta que não consegue fazer login. Você precisa investigar.

**Cenário A - Logs Não Estruturados:**
```
2024-01-15 10:30:15 Usuário joao tentou fazer login do IP 192.168.1.1
2024-01-15 10:30:16 Login falhou para usuário joao
2024-01-15 10:30:20 Usuário maria tentou fazer login do IP 192.168.1.2
2024-01-15 10:30:21 Login bem-sucedido para usuário maria
```

**Cenário B - Logs Estruturados:**
```json
{"time":"2024-01-15T10:30:15Z","level":"INFO","msg":"Tentativa de login","usuario":"joao","ip":"192.168.1.1","status":"tentativa"}
{"time":"2024-01-15T10:30:16Z","level":"ERROR","msg":"Login falhou","usuario":"joao","ip":"192.168.1.1","erro":"senha incorreta","tentativa":3}
{"time":"2024-01-15T10:30:20Z","level":"INFO","msg":"Tentativa de login","usuario":"maria","ip":"192.168.1.2","status":"tentativa"}
{"time":"2024-01-15T10:30:21Z","level":"INFO","msg":"Login bem-sucedido","usuario":"maria","ip":"192.168.1.2","status":"sucesso"}
```

**Perguntas para reflexão:**

1. Como você encontraria todos os logins falhos do usuário "joao" em cada cenário? Qual é mais fácil?

2. Como você analisaria padrões (ex: "quantos logins falharam por IP?") em cada cenário?

3. Como uma ferramenta automatizada (como Elasticsearch, Splunk) processaria os logs em cada cenário?

4. Qual cenário permite melhor análise de problemas em produção? Por quê?

5. Em uma aplicação que processa milhões de eventos, qual abordagem é mais escalável? Por quê?

**Escreva suas reflexões** sobre por que logging estruturado é essencial em aplicações modernas, especialmente em produção.

---

### Reflexão 2: Escolhendo a Ferramenta Certa

Imagine que você precisa escolher uma biblioteca de logging para três projetos diferentes:

**Projeto A**: Um script simples que roda uma vez por dia para fazer backup de arquivos.

**Projeto B**: Uma API REST que recebe 10.000 requisições por minuto e precisa logar cada requisição.

**Projeto C**: Um microsserviço crítico em uma arquitetura de larga escala que processa milhões de eventos por segundo.

**Perguntas para reflexão:**

1. Para cada projeto, qual biblioteca de logging você escolheria? (`log` padrão, `slog`, Zerolog ou Zap)

2. Quais são os critérios mais importantes para cada projeto?
   - Simplicidade?
   - Performance?
   - Estruturação?
   - Zero dependências?

3. Como a escolha errada de biblioteca pode impactar cada projeto?
   - Projeto A: O que acontece se você usar Zap (overkill)?
   - Projeto B: O que acontece se você usar `log` padrão (muito lento)?
   - Projeto C: O que acontece se você usar `slog` sem otimizações?

4. Existe uma "solução única" que funciona bem para todos os casos? Por quê?

5. Como você equilibraria entre "fácil de usar" e "alta performance" na escolha?

**Escreva suas reflexões** sobre como a escolha da ferramenta de logging deve ser baseada nas necessidades específicas do projeto, não apenas em "o que é mais popular" ou "o que parece mais fácil".

---

## Dicas para os Exercícios

### Dica 1: Estrutura de Campos
Sempre adicione contexto relevante aos logs:
```go
// ❌ Ruim
logger.Info("Erro")

// ✅ Bom
logger.Info("Erro ao processar pagamento",
    zap.String("usuario_id", "123"),
    zap.String("pedido_id", "456"),
    zap.Error(err))
```

### Dica 2: Níveis Apropriados
Use níveis de log de forma consistente:
- **DEBUG**: Informações detalhadas para desenvolvimento
- **INFO**: Eventos normais e importantes
- **WARN**: Situações que merecem atenção, mas não são erros
- **ERROR**: Erros que impedem uma operação específica
- **FATAL**: Erros críticos que param a aplicação

### Dica 3: Performance
Em produção, configure níveis apropriados:
```go
// Desenvolvimento
zerolog.SetGlobalLevel(zerolog.DebugLevel)

// Produção
zerolog.SetGlobalLevel(zerolog.InfoLevel)
```

### Dica 4: Contexto de Requisição
Use loggers com contexto para rastrear requisições:
```go
requestLogger := logger.With(
    zap.String("request_id", requestID),
    zap.String("user_id", userID),
)
```

---

## Como Entregar

Para cada exercício:

1. **Crie o código completo** que resolve o problema
2. **Execute e teste** seu código
3. **Documente** suas escolhas (por que usou determinado nível de log, por que adicionou certos campos, etc.)

Para as reflexões:

1. **Escreva suas respostas** de forma clara e completa
2. **Pense criticamente** sobre os cenários apresentados
3. **Considere** as implicações práticas de cada escolha

---

## Próximos Passos

Depois de completar os exercícios e reflexões, você estará pronto para a próxima etapa: **Aula sobre Performance e Boas Práticas de Logging**!

Lá vamos discutir:
- Quando usar cada biblioteca
- Otimizações de performance
- Padrões e anti-padrões
- Como estruturar logs em aplicações grandes
- Integração com ferramentas de monitoramento

Boa sorte com os exercícios! 🚀



