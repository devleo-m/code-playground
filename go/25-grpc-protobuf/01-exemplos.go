package main

// Este arquivo contém exemplos práticos de gRPC e Protocol Buffers
// NOTA: Estes exemplos são referências. Para executar, você precisa:
// 1. Criar o arquivo .proto
// 2. Gerar o código com protoc
// 3. Implementar servidor e cliente

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// ============================================================================
// EXEMPLO 1: Estrutura Básica do Servidor
// ============================================================================

/*
// Estrutura do servidor (após gerar código do .proto)
type server struct {
	pb.UnimplementedUserServiceServer
	users  []*pb.UserResponse
	nextID int32
}

// Implementação de método unário
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return user, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Usuário não encontrado")
}

func exemplo1Servidor() {
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
*/

// ============================================================================
// EXEMPLO 2: Cliente Básico
// ============================================================================

/*
func exemplo2Cliente() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
	log.Printf("Usuário: %v", user)
}
*/

// ============================================================================
// EXEMPLO 3: Server Streaming
// ============================================================================

/*
func (s *server) StreamUsers(req *pb.StreamUsersRequest, stream pb.UserService_StreamUsersServer) error {
	count := int(req.Count)
	for i := 0; i < count && i < len(s.users); i++ {
		if err := stream.Send(s.users[i]); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func exemplo3ServerStreaming() {
	stream, err := client.StreamUsers(ctx, &pb.StreamUsersRequest{Count: 5})
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Erro: %v", err)
		}
		log.Printf("Usuário: %v", user)
	}
}
*/

// ============================================================================
// EXEMPLO 4: Client Streaming
// ============================================================================

/*
func (s *server) CreateUsers(stream pb.UserService_CreateUsersServer) error {
	var users []*pb.UserResponse

	for {
		req, err := stream.Recv()
		if err == io.EOF {
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
		}
		s.nextID++
		users = append(users, user)
	}
}

func exemplo4ClientStreaming() {
	stream, err := client.CreateUsers(ctx)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	users := []*pb.CreateUserRequest{
		{Name: "João", Email: "joao@example.com"},
		{Name: "Maria", Email: "maria@example.com"},
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			log.Fatalf("Erro: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
	log.Printf("Criados: %d", response.Count)
}
*/

// ============================================================================
// EXEMPLO 5: Bidirectional Streaming
// ============================================================================

/*
func (s *server) ChatUsers(stream pb.UserService_ChatUsersServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

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

func exemplo5BidirectionalStreaming() {
	stream, err := client.ChatUsers(ctx)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Erro: %v", err)
			}
			log.Printf("Recebido: %s", msg.Message)
		}
	}()

	messages := []string{"Olá", "Como vai?", "Tchau"}
	for _, text := range messages {
		msg := &pb.ChatMessage{
			User:    "Cliente",
			Message: text,
		}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Erro: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
}
*/

// ============================================================================
// EXEMPLO 6: Interceptor (Middleware)
// ============================================================================

func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(start)
	log.Printf("Método: %s, Duração: %v, Erro: %v", info.FullMethod, duration, err)
	return resp, err
}

func exemplo6Interceptor() {
	// Usar no servidor:
	// s := grpc.NewServer(grpc.UnaryInterceptor(loggingUnaryInterceptor))
	_ = loggingUnaryInterceptor
}

// ============================================================================
// EXEMPLO 7: Tratamento de Erros
// ============================================================================

func exemplo7TratamentoErros() {
	// Retornar erro no servidor
	_ = status.Errorf(codes.NotFound, "Usuário não encontrado")

	// Verificar erro no cliente
	/*
	user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				log.Println("Recurso não encontrado")
			case codes.InvalidArgument:
				log.Println("Argumento inválido")
			}
		}
	}
	*/
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	// Estes exemplos são referências de código
	// Para executar, você precisa:
	// 1. Criar arquivo .proto
	// 2. Gerar código com: protoc --go_out=. --go-grpc_out=. user.proto
	// 3. Implementar servidor e cliente

	log.Println("Estes são exemplos de referência. Veja os comentários para implementação completa.")
}

