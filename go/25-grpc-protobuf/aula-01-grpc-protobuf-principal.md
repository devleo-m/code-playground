# Aula 25: gRPC & Protocol Buffers

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 25! Nas aulas anteriores, exploramos desenvolvimento web com `net/http` e o framework Gin para APIs RESTful. Agora vamos mergulhar em uma tecnologia mais avançada: **gRPC e Protocol Buffers**.

gRPC é um framework RPC (Remote Procedure Call) de alto desempenho desenvolvido pelo Google. Ele é especialmente poderoso para comunicação entre microsserviços, oferecendo performance superior, type safety e suporte a streaming.

---

## Por que gRPC?

Antes de mergulharmos no código, vamos entender por que gRPC é uma excelente escolha:

### 1. **Performance Excepcional**
- Serialização binária muito mais eficiente que JSON
- Usa HTTP/2 nativamente (multiplexing, header compression)
- Menor overhead de rede

### 2. **Type Safety**
- Código gerado a partir de definições `.proto`
- Erros detectados em tempo de compilação
- Contratos claros entre serviços

### 3. **Streaming**
- Suporte a streaming unidirecional e bidirecional
- Ideal para dados em tempo real
- Comunicação eficiente para grandes volumes

### 4. **Multi-linguagem**
- Funciona com muitas linguagens (Go, Python, Java, C++, etc.)
- Mesmo contrato `.proto` gera código para todas as linguagens
- Comunicação entre serviços em diferentes linguagens

### 5. **Code Generation**
- Gera código cliente e servidor automaticamente
- Menos código boilerplate
- Menos erros manuais

### 6. **Versionamento**
- Suporte a versionamento de APIs
- Compatibilidade retroativa
- Evolução gradual de contratos

### Quando Usar gRPC?

**Use gRPC quando:**
- Comunicação entre microsserviços
- Precisa de alta performance
- Quer type safety entre serviços
- Precisa de streaming
- Serviços em diferentes linguagens
- APIs internas (não públicas)

**Use REST/HTTP quando:**
- APIs públicas
- Precisa de compatibilidade com navegadores
- Integração com sistemas legados
- Debugging simples (JSON visível)

---

## Protocol Buffers

Protocol Buffers (protobuf) é um formato de serialização de dados desenvolvido pelo Google. É mais eficiente que JSON e XML.

### Por que Protocol Buffers?

- **Eficiência**: Serialização binária compacta
- **Velocidade**: Mais rápido que JSON/XML
- **Esquema**: Define estrutura de dados claramente
- **Versionamento**: Suporta evolução de esquemas
- **Multi-linguagem**: Funciona com muitas linguagens

### Comparação com JSON

```
JSON: {"name": "João", "age": 30}  // ~25 bytes
Protobuf: [binary]                   // ~10 bytes
```

---

## Instalação

### 1. Instalar protoc (Compilador)

**macOS:**
```bash
brew install protobuf
```

**Linux:**
```bash
sudo apt-get update
sudo apt-get install protobuf-compiler
```

**Windows:**
Baixe de: https://github.com/protocolbuffers/protobuf/releases

### 2. Instalar Plugins Go

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. Verificar Instalação

```bash
protoc --version
# Deve mostrar: libprotoc 3.x.x ou superior
```

### 4. Dependências Go

```bash
go get google.golang.org/protobuf/proto
go get google.golang.org/grpc
```

---

## Definindo um Serviço .proto

Vamos criar nosso primeiro arquivo `.proto`:

### Estrutura Básica

```protobuf
syntax = "proto3";

package user;

option go_package = "./pb";

service UserService {
  rpc GetUser(GetUserRequest) returns (UserResponse);
  rpc CreateUser(CreateUserRequest) returns (UserResponse);
  rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
}
```

### Mensagens

```protobuf
message GetUserRequest {
  int32 id = 1;
}

message CreateUserRequest {
  string name = 1;
  string email = 2;
}

message ListUsersRequest {
  int32 page = 1;
  int32 page_size = 2;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

message ListUsersResponse {
  repeated UserResponse users = 1;
  int32 total = 2;
}
```

### Arquivo Completo: user.proto

```protobuf
syntax = "proto3";

package user;

option go_package = "./pb";

service UserService {
  // RPC unário (request/response simples)
  rpc GetUser(GetUserRequest) returns (UserResponse);
  rpc CreateUser(CreateUserRequest) returns (UserResponse);
  rpc UpdateUser(UpdateUserRequest) returns (UserResponse);
  rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);
  rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
  
  // Server streaming (servidor envia múltiplas respostas)
  rpc StreamUsers(StreamUsersRequest) returns (stream UserResponse);
  
  // Client streaming (cliente envia múltiplas requisições)
  rpc CreateUsers(stream CreateUserRequest) returns (CreateUsersResponse);
  
  // Bidirectional streaming (ambos enviam múltiplas mensagens)
  rpc ChatUsers(stream ChatMessage) returns (stream ChatMessage);
}

// Mensagens de Request
message GetUserRequest {
  int32 id = 1;
}

message CreateUserRequest {
  string name = 1;
  string email = 2;
  int32 age = 3;
}

message UpdateUserRequest {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

message DeleteUserRequest {
  int32 id = 1;
}

message ListUsersRequest {
  int32 page = 1;
  int32 page_size = 2;
}

message StreamUsersRequest {
  int32 count = 1;
}

// Mensagens de Response
message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
  int32 age = 4;
}

message DeleteUserResponse {
  bool success = 1;
  string message = 2;
}

message ListUsersResponse {
  repeated UserResponse users = 1;
  int32 total = 2;
}

message CreateUsersResponse {
  repeated UserResponse users = 1;
  int32 count = 2;
}

message ChatMessage {
  string user = 1;
  string message = 2;
  int64 timestamp = 3;
}
```

### Tipos de Dados em Protocol Buffers

```protobuf
// Números
int32, int64      // Números inteiros
uint32, uint64    // Números inteiros sem sinal
float, double     // Números de ponto flutuante

// Booleanos
bool

// Strings
string

// Bytes
bytes

// Enums
enum Status {
  UNKNOWN = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}

// Arrays/Listas
repeated string tags = 1;

// Maps
map<string, string> metadata = 1;

// Mensagens aninhadas
message Address {
  string street = 1;
  string city = 2;
}
```

### Números de Campo

Os números de campo (1, 2, 3, etc.) são importantes:
- Não podem ser alterados após o deploy
- Devem ser únicos dentro da mensagem
- 1-15 usam 1 byte (use para campos frequentes)
- 16-2047 usam 2 bytes

---

## Gerando Código Go

Após criar o arquivo `.proto`, gere o código Go:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       user.proto
```

Isso gerará:
- `user.pb.go` - Código das mensagens
- `user_grpc.pb.go` - Código do serviço gRPC

### Estrutura de Diretórios Recomendada

```
project/
├── proto/
│   └── user.proto
├── pb/
│   ├── user.pb.go
│   └── user_grpc.pb.go
├── server/
│   └── main.go
└── client/
    └── main.go
```

---

## Implementando o Servidor gRPC

```go
package main

import (
	"context"
	"log"
	"net"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	pb "./pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users  []*pb.UserResponse
	nextID int32
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return user, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Usuário não encontrado: %d", req.Id)
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &pb.UserResponse{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}
	s.nextID++
	s.users = append(s.users, user)
	return user, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	for i, user := range s.users {
		if user.Id == req.Id {
			s.users[i].Name = req.Name
			s.users[i].Email = req.Email
			return s.users[i], nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Usuário não encontrado: %d", req.Id)
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	for i, user := range s.users {
		if user.Id == req.Id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return &pb.DeleteUserResponse{
				Success: true,
				Message: "Usuário deletado com sucesso",
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Usuário não encontrado: %d", req.Id)
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	page := req.Page
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	
	start := (page - 1) * pageSize
	end := start + pageSize
	
	if start >= int32(len(s.users)) {
		return &pb.ListUsersResponse{
			Users: []*pb.UserResponse{},
			Total: int32(len(s.users)),
		}, nil
	}
	
	if end > int32(len(s.users)) {
		end = int32(len(s.users))
	}
	
	return &pb.ListUsersResponse{
		Users: s.users[start:end],
		Total: int32(len(s.users)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	
	log.Println("Servidor gRPC rodando em :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}
```

---

## Criando um Cliente gRPC

```go
package main

import (
	"context"
	"log"
	"time"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	
	pb "./pb"
)

func main() {
	// Conectar ao servidor
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()
	
	// Criar cliente
	client := pb.NewUserServiceClient(conn)
	
	// Context com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Criar usuário
	user, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "João",
		Email: "joao@example.com",
		Age:   30,
	})
	if err != nil {
		log.Fatalf("Erro ao criar usuário: %v", err)
	}
	log.Printf("Usuário criado: %v", user)
	
	// Buscar usuário
	user, err = client.GetUser(ctx, &pb.GetUserRequest{Id: user.Id})
	if err != nil {
		log.Fatalf("Erro ao buscar usuário: %v", err)
	}
	log.Printf("Usuário encontrado: %v", user)
	
	// Listar usuários
	users, err := client.ListUsers(ctx, &pb.ListUsersRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("Erro ao listar usuários: %v", err)
	}
	log.Printf("Total de usuários: %d", users.Total)
	for _, u := range users.Users {
		log.Printf("  - %s (%s)", u.Name, u.Email)
	}
}
```

---

## Streaming

gRPC suporta três tipos de streaming:

### 1. Server Streaming

O servidor envia múltiplas respostas para uma única requisição do cliente.

#### Definindo no .proto

```protobuf
service UserService {
  rpc StreamUsers(StreamUsersRequest) returns (stream UserResponse);
}
```

#### Implementando no Servidor

```go
func (s *server) StreamUsers(req *pb.StreamUsersRequest, stream pb.UserService_StreamUsersServer) error {
	count := int(req.Count)
	if count == 0 {
		count = len(s.users)
	}
	
	for i := 0; i < count && i < len(s.users); i++ {
		if err := stream.Send(s.users[i]); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond) // Simular delay
	}
	return nil
}
```

#### Usando no Cliente

```go
stream, err := client.StreamUsers(ctx, &pb.StreamUsersRequest{Count: 5})
if err != nil {
	log.Fatalf("Erro ao criar stream: %v", err)
}

for {
	user, err := stream.Recv()
	if err == io.EOF {
		break
	}
	if err != nil {
		log.Fatalf("Erro ao receber: %v", err)
	}
	log.Printf("Usuário recebido: %v", user)
}
```

### 2. Client Streaming

O cliente envia múltiplas requisições e o servidor responde uma vez.

#### Definindo no .proto

```protobuf
service UserService {
  rpc CreateUsers(stream CreateUserRequest) returns (CreateUsersResponse);
}
```

#### Implementando no Servidor

```go
func (s *server) CreateUsers(stream pb.UserService_CreateUsersServer) error {
	var users []*pb.UserResponse
	
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Cliente terminou de enviar
			return stream.SendAndClose(&pb.CreateUsersResponse{
				Users: users,
				Count: int32(len(users)),
			})
		}
		if err != nil {
			return err
		}
		
		user := &pb.UserResponse{
			Id:    s.nextID,
			Name:  req.Name,
			Email: req.Email,
			Age:   req.Age,
		}
		s.nextID++
		users = append(users, user)
	}
}
```

#### Usando no Cliente

```go
stream, err := client.CreateUsers(ctx)
if err != nil {
	log.Fatalf("Erro ao criar stream: %v", err)
}

users := []*pb.CreateUserRequest{
	{Name: "João", Email: "joao@example.com", Age: 30},
	{Name: "Maria", Email: "maria@example.com", Age: 25},
	{Name: "Pedro", Email: "pedro@example.com", Age: 35},
}

for _, user := range users {
	if err := stream.Send(user); err != nil {
		log.Fatalf("Erro ao enviar: %v", err)
	}
}

response, err := stream.CloseAndRecv()
if err != nil {
	log.Fatalf("Erro ao receber resposta: %v", err)
}
log.Printf("Usuários criados: %d", response.Count)
```

### 3. Bidirectional Streaming

Ambos enviam múltiplas mensagens simultaneamente.

#### Definindo no .proto

```protobuf
service UserService {
  rpc ChatUsers(stream ChatMessage) returns (stream ChatMessage);
}
```

#### Implementando no Servidor

```go
func (s *server) ChatUsers(stream pb.UserService_ChatUsersServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		
		// Processar mensagem e enviar resposta
		response := &pb.ChatMessage{
			User:      "Servidor",
			Message:   "Echo: " + msg.Message,
			Timestamp: time.Now().Unix(),
		}
		
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}
```

#### Usando no Cliente

```go
stream, err := client.ChatUsers(ctx)
if err != nil {
	log.Fatalf("Erro ao criar stream: %v", err)
}

// Goroutine para receber mensagens
go func() {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Erro ao receber: %v", err)
		}
		log.Printf("Recebido: %s: %s", msg.User, msg.Message)
	}
}()

// Enviar mensagens
messages := []string{"Olá", "Como vai?", "Tchau"}
for _, text := range messages {
	msg := &pb.ChatMessage{
		User:      "Cliente",
		Message:   text,
		Timestamp: time.Now().Unix(),
	}
	if err := stream.Send(msg); err != nil {
		log.Fatalf("Erro ao enviar: %v", err)
	}
	time.Sleep(1 * time.Second)
}

stream.CloseSend()
```

---

## Interceptors (Middleware)

gRPC suporta interceptors para adicionar funcionalidades como logging, autenticação, etc.

### Unary Interceptor

```go
func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	
	resp, err := handler(ctx, req)
	
	duration := time.Since(start)
	log.Printf("Método: %s, Duração: %v, Erro: %v", info.FullMethod, duration, err)
	
	return resp, err
}

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingUnaryInterceptor),
	)
	// ...
}
```

### Stream Interceptor

```go
func loggingStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	
	err := handler(srv, ss)
	
	duration := time.Since(start)
	log.Printf("Stream Método: %s, Duração: %v, Erro: %v", info.FullMethod, duration, err)
	
	return err
}

func main() {
	s := grpc.NewServer(
		grpc.StreamInterceptor(loggingStreamInterceptor),
	)
	// ...
}
```

### Múltiplos Interceptors

```go
func chainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for i := len(interceptors) - 1; i >= 0; i-- {
			handler = func(current grpc.UnaryHandler, interceptor grpc.UnaryServerInterceptor) grpc.UnaryHandler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					return interceptor(ctx, req, info, current)
				}
			}(handler, interceptors[i])
		}
		return handler(ctx, req)
	}
}

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(chainUnaryInterceptors(
			loggingUnaryInterceptor,
			authUnaryInterceptor,
		)),
	)
	// ...
}
```

---

## Autenticação

### Autenticação com Tokens

```go
func authUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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
	
	return handler(ctx, req)
}

func isValidToken(token string) bool {
	// Implementar validação do token
	return token == "valid-token"
}
```

### Usando no Cliente

```go
md := metadata.New(map[string]string{
	"authorization": "valid-token",
})
ctx := metadata.NewOutgoingContext(context.Background(), md)

user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
```

---

## Erros e Status Codes

gRPC usa códigos de status específicos:

```go
import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Retornar erro com código
return nil, status.Errorf(codes.NotFound, "Usuário não encontrado")

// Verificar código de erro no cliente
st, ok := status.FromError(err)
if ok {
	switch st.Code() {
	case codes.NotFound:
		log.Println("Recurso não encontrado")
	case codes.Unauthenticated:
		log.Println("Não autenticado")
	case codes.InvalidArgument:
		log.Println("Argumento inválido")
	}
}
```

### Códigos de Status Comuns

- `OK` - Sucesso
- `InvalidArgument` - Argumento inválido
- `NotFound` - Recurso não encontrado
- `AlreadyExists` - Recurso já existe
- `Unauthenticated` - Não autenticado
- `PermissionDenied` - Permissão negada
- `Internal` - Erro interno
- `Unavailable` - Serviço indisponível

---

## Boas Práticas

### 1. **Versionamento de APIs**
Use versionamento em seus serviços:

```protobuf
service UserServiceV1 {
  rpc GetUser(GetUserRequest) returns (UserResponse);
}

service UserServiceV2 {
  rpc GetUser(GetUserRequestV2) returns (UserResponseV2);
}
```

### 2. **Tratamento de Erros**
Sempre retorne erros apropriados:

```go
if user == nil {
	return nil, status.Errorf(codes.NotFound, "Usuário não encontrado")
}
```

### 3. **Timeouts**
Use context com timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### 4. **Validação**
Valide dados de entrada:

```go
if req.Id <= 0 {
	return nil, status.Errorf(codes.InvalidArgument, "ID inválido")
}
```

### 5. **Logging**
Use logging estruturado:

```go
log.Printf("Processando requisição: método=%s, id=%d", info.FullMethod, req.Id)
```

### 6. **Testes**
Teste seus serviços:

```go
func TestGetUser(t *testing.T) {
	s := &server{}
	req := &pb.GetUserRequest{Id: 1}
	resp, err := s.GetUser(context.Background(), req)
	// Verificar resposta
}
```

---

## Conclusão

Nesta aula, exploramos gRPC e Protocol Buffers:

1. **Por que gRPC**: Performance, type safety, streaming
2. **Protocol Buffers**: Formato de serialização eficiente
3. **Definindo Serviços**: Como criar arquivos `.proto`
4. **Implementando Servidor**: Criando servidores gRPC
5. **Criando Cliente**: Fazendo requisições gRPC
6. **Streaming**: Server, client e bidirectional streaming
7. **Interceptors**: Middleware para gRPC
8. **Autenticação**: Como autenticar requisições

gRPC é uma excelente escolha para comunicação entre microsserviços, oferecendo performance superior e type safety. Ele é especialmente útil em arquiteturas de microsserviços onde serviços precisam se comunicar de forma eficiente.

Com isso, completamos nossa jornada pelo desenvolvimento web em Go! Você agora conhece:
- `net/http` para controle total
- Gin para produtividade
- gRPC para microsserviços

Cada ferramenta tem seu lugar. Escolha baseado nas necessidades do seu projeto!

Até a próxima aula!

