# Aula 25: gRPC & Protocol Buffers (Versão Simplificada)

Olá! Esta é uma versão resumida da aula 25 sobre gRPC e Protocol Buffers.

---

## Por que gRPC?

- ✅ **Performance** (serialização binária)
- ✅ **Type safety** (código gerado)
- ✅ **Streaming** (unidirecional e bidirecional)
- ✅ **Multi-linguagem** (mesmo contrato .proto)
- ✅ **HTTP/2** nativo

---

## 1. Instalação

```bash
# Instalar protoc
brew install protobuf  # macOS
sudo apt-get install protobuf-compiler  # Linux

# Plugins Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## 2. Arquivo .proto

```protobuf
syntax = "proto3";

package user;
option go_package = "./pb";

service UserService {
  rpc GetUser(GetUserRequest) returns (UserResponse);
  rpc CreateUser(CreateUserRequest) returns (UserResponse);
}

message GetUserRequest {
  int32 id = 1;
}

message CreateUserRequest {
  string name = 1;
  string email = 2;
}

message UserResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

---

## 3. Gerar Código

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       user.proto
```

---

## 4. Servidor

```go
type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:    req.Id,
		Name:  "João",
		Email: "joao@example.com",
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	s.Serve(lis)
}
```

---

## 5. Cliente

```go
conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
defer conn.Close()

client := pb.NewUserServiceClient(conn)
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

user, _ := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
```

---

## 6. Streaming

### Server Streaming

```protobuf
rpc StreamUsers(StreamUsersRequest) returns (stream UserResponse);
```

```go
func (s *server) StreamUsers(req *pb.StreamUsersRequest, stream pb.UserService_StreamUsersServer) error {
	for _, user := range s.users {
		stream.Send(user)
	}
	return nil
}
```

### Client Streaming

```protobuf
rpc CreateUsers(stream CreateUserRequest) returns (CreateUsersResponse);
```

```go
func (s *server) CreateUsers(stream pb.UserService_CreateUsersServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CreateUsersResponse{...})
		}
		// Processar req
	}
}
```

---

## 7. Erros

```go
import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Servidor
return nil, status.Errorf(codes.NotFound, "Usuário não encontrado")

// Cliente
st, ok := status.FromError(err)
if ok && st.Code() == codes.NotFound {
	// Tratar erro
}
```

---

## 8. Interceptor

```go
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("Método: %s, Duração: %v", info.FullMethod, time.Since(start))
	return resp, err
}

s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
```

---

## Resumo Rápido

| Tarefa | Código |
|--------|--------|
| Definir serviço | `service UserService { rpc GetUser(...) returns (...); }` |
| Gerar código | `protoc --go_out=. --go-grpc_out=. user.proto` |
| Implementar | `func (s *server) GetUser(...) (*pb.UserResponse, error)` |
| Criar servidor | `grpc.NewServer()` |
| Criar cliente | `pb.NewUserServiceClient(conn)` |
| Streaming | `stream.Send(...)` / `stream.Recv()` |
| Erro | `status.Errorf(codes.NotFound, "mensagem")` |

---

Pronto! Agora você tem o básico de gRPC. Pratique criando um serviço completo!

