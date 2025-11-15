# Aula 25: Exercícios e Reflexão - gRPC & Protocol Buffers

Olá! Agora é hora de praticar! Vamos criar alguns serviços gRPC para consolidar o conhecimento.

---

## Exercício 1: Serviço de Calculadora

Crie um serviço gRPC de calculadora.

**Requisitos:**
- Definir arquivo `.proto` com operações: Add, Subtract, Multiply, Divide
- Implementar servidor gRPC
- Criar cliente que testa todas as operações
- Tratamento de erro para divisão por zero

**Estrutura .proto:**
```protobuf
service CalculatorService {
  rpc Add(OperationRequest) returns (OperationResponse);
  rpc Subtract(OperationRequest) returns (OperationResponse);
  rpc Multiply(OperationRequest) returns (OperationResponse);
  rpc Divide(OperationRequest) returns (OperationResponse);
}

message OperationRequest {
  double a = 1;
  double b = 2;
}

message OperationResponse {
  double result = 1;
}
```

**Exemplo de uso:**
```bash
# Cliente chama Add(10, 5) → retorna 15
```

---

## Exercício 2: Serviço de Chat com Streaming

Crie um serviço de chat usando bidirectional streaming.

**Requisitos:**
- Cliente envia mensagens
- Servidor faz echo das mensagens
- Suporte a múltiplos clientes simultâneos
- Timestamp em cada mensagem

**Estrutura .proto:**
```protobuf
service ChatService {
  rpc Chat(stream ChatMessage) returns (stream ChatMessage);
}

message ChatMessage {
  string user = 1;
  string message = 2;
  int64 timestamp = 3;
}
```

**Funcionalidades:**
- Cliente envia mensagens continuamente
- Servidor responde com echo
- Suporte a comandos especiais (ex: `/quit`)

---

## Exercício 3: Serviço de Notificações

Crie um serviço de notificações com server streaming.

**Requisitos:**
- Cliente se inscreve para receber notificações
- Servidor envia notificações periodicamente
- Cliente pode cancelar a inscrição
- Diferentes tipos de notificações

**Estrutura .proto:**
```protobuf
service NotificationService {
  rpc Subscribe(SubscribeRequest) returns (stream Notification);
}

message SubscribeRequest {
  string user_id = 1;
  repeated string types = 2; // email, sms, push
}

message Notification {
  string id = 1;
  string type = 2;
  string title = 3;
  string message = 4;
  int64 timestamp = 5;
}
```

**Funcionalidades:**
- Servidor envia notificações a cada 5 segundos
- Cliente pode filtrar por tipo
- Simular diferentes tipos de notificações

---

## Exercício 4: Serviço de Upload com Client Streaming

Crie um serviço de upload de arquivos usando client streaming.

**Requisitos:**
- Cliente envia arquivo em chunks
- Servidor recebe e salva o arquivo
- Retorna hash MD5 do arquivo
- Suporte a metadados (nome, tipo, tamanho)

**Estrutura .proto:**
```protobuf
service FileService {
  rpc UploadFile(stream FileChunk) returns (UploadResponse);
}

message FileChunk {
  string filename = 1;
  string content_type = 2;
  bytes data = 3;
  bool is_last = 4;
}

message UploadResponse {
  string file_id = 1;
  string md5_hash = 2;
  int64 size = 3;
}
```

**Funcionalidades:**
- Cliente divide arquivo em chunks de 1KB
- Servidor reconstrói arquivo
- Validação de integridade

---

## Exercício 5: Serviço de Usuários Completo

Crie um serviço gRPC completo de gerenciamento de usuários.

**Requisitos:**
- CRUD completo (Create, Read, Update, Delete)
- Listagem com paginação
- Busca por diferentes critérios
- Validação de dados
- Tratamento de erros apropriado

**Estrutura .proto:**
```protobuf
service UserService {
  rpc CreateUser(CreateUserRequest) returns (UserResponse);
  rpc GetUser(GetUserRequest) returns (UserResponse);
  rpc UpdateUser(UpdateUserRequest) returns (UserResponse);
  rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);
  rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
  rpc SearchUsers(SearchUsersRequest) returns (ListUsersResponse);
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
  int32 age = 4;
  string created_at = 5;
}
```

**Funcionalidades:**
- Validação de email
- Validação de idade (18+)
- Paginação na listagem
- Busca por nome ou email

---

## Exercício 6: Interceptors e Middleware

Implemente interceptors para logging, autenticação e rate limiting.

**Requisitos:**
- Interceptor de logging (registra todas as chamadas)
- Interceptor de autenticação (verifica token)
- Interceptor de rate limiting (máx 100 req/min)
- Aplicar interceptors no servidor

**Estrutura:**
```go
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Implementar logging
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Verificar token do metadata
}

func rateLimitInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Implementar rate limiting por IP
}
```

---

## Exercício 7: Serviço com Erros Customizados

Crie um serviço que retorna erros detalhados e customizados.

**Requisitos:**
- Erros com códigos apropriados (NotFound, InvalidArgument, etc.)
- Mensagens de erro detalhadas
- Cliente trata diferentes tipos de erro
- Logging de erros no servidor

**Exemplo:**
```go
// Servidor
if user == nil {
	return nil, status.Errorf(codes.NotFound, "Usuário com ID %d não encontrado", req.Id)
}

// Cliente
st, ok := status.FromError(err)
if ok {
	switch st.Code() {
	case codes.NotFound:
		// Tratar usuário não encontrado
	case codes.InvalidArgument:
		// Tratar argumento inválido
	}
}
```

---

## Exercício 8: Serviço de Monitoramento

Crie um serviço de monitoramento com métricas em tempo real.

**Requisitos:**
- Server streaming de métricas
- Métricas: CPU, memória, requisições por segundo
- Cliente recebe atualizações a cada segundo
- Suporte a múltiplos clientes

**Estrutura .proto:**
```protobuf
service MonitoringService {
  rpc StreamMetrics(MetricsRequest) returns (stream Metric);
}

message MetricsRequest {
  repeated string metric_types = 1; // cpu, memory, requests
}

message Metric {
  string type = 1;
  double value = 2;
  int64 timestamp = 3;
  map<string, string> labels = 4;
}
```

**Funcionalidades:**
- Simular métricas do sistema
- Filtrar por tipo de métrica
- Atualizações em tempo real

---

## Exercício 9: Serviço de Cache Distribuído

Crie um serviço de cache distribuído usando gRPC.

**Requisitos:**
- `Set(key, value)` - Armazenar valor
- `Get(key)` - Recuperar valor
- `Delete(key)` - Remover valor
- `ListKeys()` - Listar todas as chaves
- TTL (Time To Live) para valores
- Persistência em memória

**Estrutura .proto:**
```protobuf
service CacheService {
  rpc Set(SetRequest) returns (SetResponse);
  rpc Get(GetRequest) returns (GetResponse);
  rpc Delete(DeleteRequest) returns (DeleteResponse);
  rpc ListKeys(ListKeysRequest) returns (ListKeysResponse);
}

message SetRequest {
  string key = 1;
  string value = 2;
  int32 ttl_seconds = 3;
}
```

---

## Exercício 10: Sistema Completo de Microsserviços

Crie um sistema completo com múltiplos serviços gRPC.

**Requisitos:**
- **Serviço de Usuários**: CRUD de usuários
- **Serviço de Produtos**: CRUD de produtos
- **Serviço de Pedidos**: Criar pedidos, listar pedidos
- **Serviço de Notificações**: Enviar notificações
- Comunicação entre serviços
- Interceptors compartilhados
- Tratamento de erros consistente

**Arquitetura:**
```
Cliente
  ↓
API Gateway (opcional)
  ↓
┌─────────────┬─────────────┬─────────────┐
│   Users     │  Products   │   Orders    │
│   Service   │   Service   │   Service   │
└─────────────┴─────────────┴─────────────┘
  ↓
Notification Service
```

**Funcionalidades:**
- Criar pedido que consulta usuário e produtos
- Enviar notificação quando pedido é criado
- Validação entre serviços
- Tratamento de erros em cascata

---

## Reflexão

Após completar os exercícios, reflita sobre:

1. **gRPC vs REST**
   - Quando usar gRPC vs REST?
   - Quais são as principais vantagens de cada um?

2. **Streaming**
   - Quais casos de uso se beneficiam de streaming?
   - Como streaming melhora a performance?

3. **Protocol Buffers**
   - Por que Protocol Buffers é mais eficiente que JSON?
   - Como versionar APIs com Protocol Buffers?

4. **Microsserviços**
   - Como gRPC facilita comunicação entre microsserviços?
   - Quais são os desafios de usar gRPC em microsserviços?

5. **Performance**
   - Como gRPC melhora a performance em relação a REST?
   - Quais otimizações você pode fazer?

---

## Desafio Final

Crie um sistema completo de e-commerce com microsserviços gRPC:

**Serviços:**
1. **User Service**: Autenticação, perfis, gerenciamento de usuários
2. **Product Service**: Catálogo de produtos, busca, categorias
3. **Cart Service**: Carrinho de compras, cálculos
4. **Order Service**: Processamento de pedidos, histórico
5. **Payment Service**: Processamento de pagamentos
6. **Notification Service**: Notificações em tempo real
7. **Inventory Service**: Controle de estoque

**Funcionalidades:**
- Comunicação entre todos os serviços
- Streaming para notificações em tempo real
- Validação e tratamento de erros robusto
- Interceptors para logging, auth, rate limiting
- Versionamento de APIs
- Documentação completa

**Desafios extras:**
- Implementar circuit breaker
- Implementar retry logic
- Implementar distributed tracing
- Implementar service discovery

---

Boa sorte com os exercícios! Pratique bastante e explore todas as funcionalidades do gRPC!

