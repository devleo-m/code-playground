# Aula 25: Performance e Boas Práticas - gRPC & Protocol Buffers

Olá! Nesta parte da aula, vamos explorar boas práticas, otimizações de performance e padrões profissionais para desenvolvimento com gRPC.

---

## Boas Práticas Gerais

### 1. Tratamento de Erros Consistente

**❌ Ruim:**
```go
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := findUser(req.Id)
	return user, nil // Sempre retorna nil, mesmo se não encontrar
}
```

**✅ Bom:**
```go
import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID deve ser maior que zero")
	}
	
	user, err := s.store.GetUser(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Usuário com ID %d não encontrado: %v", req.Id, err)
	}
	
	return user, nil
}
```

### 2. Validação de Request

Sempre valide dados de entrada:

```go
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// Validação
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Nome é obrigatório")
	}
	if len(req.Name) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "Nome deve ter no mínimo 3 caracteres")
	}
	if !isValidEmail(req.Email) {
		return nil, status.Errorf(codes.InvalidArgument, "Email inválido: %s", req.Email)
	}
	if req.Age < 18 || req.Age > 120 {
		return nil, status.Errorf(codes.InvalidArgument, "Idade deve estar entre 18 e 120")
	}
	
	// Processar
	user, err := s.store.CreateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao criar usuário: %v", err)
	}
	
	return user, nil
}
```

### 3. Context para Timeouts e Cancelamento

Use context adequadamente:

```go
func (s *server) ProcessData(ctx context.Context, req *pb.ProcessRequest) (*pb.ProcessResponse, error) {
	// Verificar se contexto foi cancelado
	select {
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Operação cancelada")
	default:
	}
	
	// Criar context com timeout para operação
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	result, err := s.processWithContext(opCtx, req)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, status.Errorf(codes.DeadlineExceeded, "Timeout excedido")
		}
		return nil, status.Errorf(codes.Internal, "Erro ao processar: %v", err)
	}
	
	return result, nil
}
```

### 4. Logging Estruturado

Use logging estruturado:

```go
import "log/slog"

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	s.logger.Info("GetUser chamado",
		"user_id", req.Id,
		"trace_id", getTraceID(ctx),
	)
	
	user, err := s.store.GetUser(req.Id)
	if err != nil {
		s.logger.Error("Erro ao buscar usuário",
			"user_id", req.Id,
			"error", err,
		)
		return nil, status.Errorf(codes.NotFound, "Usuário não encontrado")
	}
	
	s.logger.Info("Usuário encontrado",
		"user_id", req.Id,
	)
	
	return user, nil
}
```

---

## Performance

### 1. Connection Pooling

Configure o cliente adequadamente:

```go
conn, err := grpc.Dial("localhost:50051",
	grpc.WithTransportCredentials(insecure.NewCredentials()),
	grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(4*1024*1024), // 4MB
		grpc.MaxCallSendMsgSize(4*1024*1024), // 4MB
	),
	grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	}),
)
```

### 2. Keepalive no Servidor

Configure keepalive no servidor:

```go
import "google.golang.org/grpc/keepalive"

server := grpc.NewServer(
	grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,
		MaxConnectionAge:      30 * time.Second,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               1 * time.Second,
	}),
)
```

### 3. Streaming Eficiente

Use streaming para grandes volumes de dados:

```go
func (s *server) StreamLargeData(req *pb.StreamRequest, stream pb.Service_StreamLargeDataServer) error {
	const chunkSize = 64 * 1024 // 64KB
	
	data := s.getLargeData()
	
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		
		chunk := &pb.DataChunk{
			Data: data[i:end],
		}
		
		if err := stream.Send(chunk); err != nil {
			return err
		}
		
		// Verificar se contexto foi cancelado
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}
	}
	
	return nil
}
```

### 4. Reutilização de Mensagens

Reutilize mensagens quando possível:

```go
var responsePool = sync.Pool{
	New: func() interface{} {
		return &pb.UserResponse{}
	},
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	response := responsePool.Get().(*pb.UserResponse)
	defer responsePool.Put(response)
	
	// Reset campos
	response.Reset()
	
	// Preencher dados
	response.Id = user.ID
	response.Name = user.Name
	
	// Criar cópia para retornar
	result := &pb.UserResponse{
		Id:   response.Id,
		Name: response.Name,
	}
	
	return result, nil
}
```

---

## Segurança

### 1. Autenticação com TLS

Use TLS em produção:

```go
import "google.golang.org/grpc/credentials"

// Servidor
creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
if err != nil {
	log.Fatal(err)
}

server := grpc.NewServer(grpc.Creds(creds))

// Cliente
creds, err := credentials.NewClientTLSFromFile("ca-cert.pem", "")
if err != nil {
	log.Fatal(err)
}

conn, err := grpc.Dial("server:50051", grpc.WithTransportCredentials(creds))
```

### 2. Autenticação com Tokens

Implemente autenticação:

```go
func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Metadata não encontrada")
	}
	
	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Token não fornecido")
	}
	
	token := tokens[0]
	if !isValidToken(token) {
		return nil, status.Errorf(codes.Unauthenticated, "Token inválido")
	}
	
	// Adicionar user ID ao context
	ctx = context.WithValue(ctx, "user_id", extractUserID(token))
	
	return handler(ctx, req)
}
```

### 3. Rate Limiting

Implemente rate limiting:

```go
type rateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	maxReqs  int
	window   time.Duration
}

func (rl *rateLimiter) allow(identifier string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	reqs := rl.requests[identifier]
	validReqs := []time.Time{}
	for _, req := range reqs {
		if req.After(cutoff) {
			validReqs = append(validReqs, req)
		}
	}
	
	if len(validReqs) >= rl.maxReqs {
		return false
	}
	
	validReqs = append(validReqs, now)
	rl.requests[identifier] = validReqs
	return true
}

func rateLimitInterceptor(limiter *rateLimiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		identifier := getClientIdentifier(ctx)
		if !limiter.allow(identifier) {
			return nil, status.Errorf(codes.ResourceExhausted, "Rate limit excedido")
		}
		return handler(ctx, req)
	}
}
```

---

## Organização de Código

### 1. Estrutura de Projeto

Organize o projeto adequadamente:

```
project/
├── proto/
│   └── user.proto
├── pb/
│   ├── user.pb.go
│   └── user_grpc.pb.go
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── internal/
│   ├── server/
│   │   └── user.go
│   ├── service/
│   │   └── user.go
│   ├── store/
│   │   └── user.go
│   └── interceptor/
│       ├── auth.go
│       └── logging.go
└── go.mod
```

### 2. Separação de Responsabilidades

Separe servidor, serviço e store:

```go
// internal/service/user.go
type UserService interface {
	GetUser(id int32) (*pb.UserResponse, error)
	CreateUser(req *pb.CreateUserRequest) (*pb.UserResponse, error)
}

// internal/server/user.go
type server struct {
	pb.UnimplementedUserServiceServer
	service UserService
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return s.service.GetUser(req.Id)
}
```

---

## Versionamento

### 1. Versionamento de APIs

Use versionamento em seus serviços:

```protobuf
// v1/user.proto
package user.v1;
service UserService {
  rpc GetUser(GetUserRequest) returns (UserResponse);
}

// v2/user.proto
package user.v2;
service UserService {
  rpc GetUser(GetUserRequest) returns (UserResponseV2);
}
```

### 2. Compatibilidade Retroativa

Mantenha compatibilidade:

```protobuf
message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
  // Campos novos sempre no final
  string phone = 4; // Novo campo
}
```

---

## Testes

### 1. Testes de Servidor

Teste implementações do servidor:

```go
func TestGetUser(t *testing.T) {
	ctx := context.Background()
	server := &server{store: mockStore}
	
	req := &pb.GetUserRequest{Id: 1}
	resp, err := server.GetUser(ctx, req)
	
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}
	
	if resp.Id != 1 {
		t.Errorf("Esperado ID 1, obteve %d", resp.Id)
	}
}
```

### 2. Testes de Cliente

Teste cliente:

```go
func TestClientGetUser(t *testing.T) {
	// Criar servidor de teste
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, mockServer)
	
	go func() {
		s.Serve(lis)
	}()
	defer s.Stop()
	
	// Conectar cliente
	conn, _ := grpc.Dial("", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	
	client := pb.NewUserServiceClient(conn)
	user, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})
	
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
	
	if user.Id != 1 {
		t.Errorf("Esperado ID 1, obteve %d", user.Id)
	}
}
```

---

## Monitoramento

### 1. Métricas com Prometheus

Integre métricas:

```go
import "github.com/prometheus/client_golang/prometheus"

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total de requisições gRPC",
		},
		[]string{"method", "status"},
	)
	
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "grpc_request_duration_seconds",
			Help: "Duração das requisições gRPC",
		},
		[]string{"method"},
	)
)

func metricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(start)
	
	status := "success"
	if err != nil {
		status = "error"
	}
	
	requestsTotal.WithLabelValues(info.FullMethod, status).Inc()
	requestDuration.WithLabelValues(info.FullMethod).Observe(duration.Seconds())
	
	return resp, err
}
```

---

## Conclusão

Seguindo essas boas práticas, você criará serviços gRPC robustos, seguros e performáticos. Lembre-se:

- Valide sempre entrada
- Trate erros com códigos apropriados
- Use context adequadamente
- Configure segurança (TLS, auth)
- Organize o código
- Teste adequadamente
- Monitore em produção
- Versionamento de APIs
- Otimize performance

Pratique essas técnicas e você estará pronto para criar microsserviços profissionais com gRPC!

