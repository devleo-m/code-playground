# Aula 28: Realtime Communication

Olá, futuro(a) Gopher!

Bem-vindo(a) à aula 28! Na aula anterior, exploramos o framework Gin para desenvolvimento web. Agora vamos mergulhar no fascinante mundo da **comunicação em tempo real**, onde aplicações precisam enviar e receber dados instantaneamente, sem que o cliente precise fazer requisições repetidas.

Comunicação em tempo real é essencial para aplicações modernas como chats, dashboards ao vivo, notificações push, jogos multiplayer e sistemas de colaboração. Go, com sua excelente concorrência, é perfeito para construir essas aplicações!

---

## O que é Realtime Communication?

**Comunicação em tempo real** (Realtime Communication) é a capacidade de uma aplicação enviar e receber dados instantaneamente entre cliente e servidor, sem necessidade de polling (requisições repetidas) ou refresh manual.

### Características Principais

1. **Bidirecional**: Cliente e servidor podem enviar dados a qualquer momento
2. **Instantâneo**: Dados chegam quase imediatamente (baixa latência)
3. **Persistente**: Conexão mantida aberta por longos períodos
4. **Eficiente**: Não precisa fazer requisições HTTP repetidas

### Casos de Uso Comuns

- **Chats e Mensageiros**: WhatsApp, Slack, Discord
- **Dashboards ao Vivo**: Monitoramento de métricas, analytics
- **Notificações Push**: Alertas em tempo real
- **Jogos Multiplayer**: Atualizações de posição, eventos
- **Colaboração**: Google Docs, edição simultânea
- **Trading**: Cotações de ações em tempo real
- **IoT**: Telemetria de sensores

---

## Por que Go é Ideal para Realtime?

Go possui características que o tornam excelente para comunicação em tempo real:

### 1. **Concorrência Nativa**
```go
// Cada conexão pode ser uma goroutine
go handleConnection(conn)
```

### 2. **Baixo Consumo de Memória**
- Goroutines são leves (2KB inicial)
- Pode gerenciar milhares de conexões simultâneas

### 3. **Performance**
- Código compilado, execução rápida
- Garbage collector eficiente

### 4. **Simplicidade**
- Código limpo e fácil de manter
- Sem complexidade desnecessária

---

## Tecnologias de Realtime

### 1. WebSockets

**WebSocket** é um protocolo de comunicação bidirecional sobre TCP que permite comunicação full-duplex entre cliente e servidor.

#### Características

- **Conexão Persistente**: Uma vez estabelecida, permanece aberta
- **Baixa Latência**: Overhead mínimo
- **Bidirecional**: Ambos podem enviar dados
- **Protocolo**: `ws://` ou `wss://` (seguro)

#### Quando Usar WebSockets?

✅ **Use quando:**
- Precisa de comunicação bidirecional
- Baixa latência é crítica
- Muitas mensagens pequenas
- Chat, jogos, dashboards

❌ **Evite quando:**
- Apenas servidor → cliente (use SSE)
- Poucas mensagens (HTTP polling é suficiente)
- Firewalls muito restritivos

### 2. Server-Sent Events (SSE)

**SSE** permite que o servidor envie dados para o cliente automaticamente através de uma conexão HTTP persistente.

#### Características

- **Unidirecional**: Apenas servidor → cliente
- **HTTP Simples**: Usa HTTP normal, mais compatível
- **Reconexão Automática**: Cliente reconecta automaticamente
- **Formato Texto**: Fácil de debugar

#### Quando Usar SSE?

✅ **Use quando:**
- Apenas servidor precisa enviar dados
- Notificações, feeds, atualizações
- Precisa de compatibilidade máxima
- Mensagens são principalmente texto

❌ **Evite quando:**
- Precisa de comunicação bidirecional
- Precisa enviar dados binários complexos

### 3. HTTP Long Polling

**Long Polling** mantém uma requisição HTTP aberta até que o servidor tenha dados para enviar.

#### Características

- **Compatível**: Funciona em qualquer lugar
- **Simples**: Apenas HTTP
- **Overhead**: Mais requisições que WebSocket

#### Quando Usar?

✅ **Use quando:**
- Precisa de máxima compatibilidade
- Firewalls muito restritivos
- Aplicação simples, poucas conexões

---

## WebSockets com Go - Biblioteca Padrão

Vamos começar com a biblioteca padrão do Go para entender os fundamentos.

### Instalação

A biblioteca padrão `golang.org/x/net/websocket` fornece suporte básico, mas vamos usar `gorilla/websocket` que é mais completa e recomendada.

```bash
go get github.com/gorilla/websocket
```

### Exemplo Básico - Servidor WebSocket

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Em produção, valide a origem!
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade conexão HTTP para WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao fazer upgrade:", err)
		return
	}
	defer conn.Close()

	log.Println("Cliente conectado")

	// Goroutine para ler mensagens
	go readMessages(conn)

	// Enviar mensagens
	writeMessages(conn)
}

func readMessages(conn *websocket.Conn) {
	for {
		// Ler mensagem
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Erro ao ler: %v", err)
			}
			break
		}

		log.Printf("Recebido: %s (tipo: %d)", message, messageType)

		// Echo - enviar de volta
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Erro ao escrever:", err)
			break
		}
	}
}

func writeMessages(conn *websocket.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Enviar mensagem a cada 5 segundos
			message := fmt.Sprintf("Servidor: %s", time.Now().Format(time.RFC3339))
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println("Erro ao escrever:", err)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Servidor WebSocket rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Cliente HTML Simples

```html
<!DOCTYPE html>
<html>
<head>
	<title>WebSocket Test</title>
</head>
<body>
	<div id="messages"></div>
	<input type="text" id="messageInput" placeholder="Digite uma mensagem">
	<button onclick="sendMessage()">Enviar</button>

	<script>
		const ws = new WebSocket("ws://localhost:8080/ws");
		const messagesDiv = document.getElementById("messages");
		const input = document.getElementById("messageInput");

		ws.onopen = function() {
			addMessage("Conectado ao servidor!");
		};

		ws.onmessage = function(event) {
			addMessage("Servidor: " + event.data);
		};

		ws.onclose = function() {
			addMessage("Conexão fechada");
		};

		ws.onerror = function(error) {
			addMessage("Erro: " + error);
		};

		function sendMessage() {
			const message = input.value;
			if (message) {
				ws.send(message);
				addMessage("Você: " + message);
				input.value = "";
			}
		}

		function addMessage(message) {
			const p = document.createElement("p");
			p.textContent = message;
			messagesDiv.appendChild(p);
		}

		// Enviar com Enter
		input.addEventListener("keypress", function(e) {
			if (e.key === "Enter") {
				sendMessage();
			}
		});
	</script>
</body>
</html>
```

### Tipos de Mensagens WebSocket

```go
const (
	TextMessage   = 1 // Mensagem de texto (JSON, string)
	BinaryMessage = 2 // Dados binários (imagens, arquivos)
	CloseMessage  = 8 // Fechar conexão
	PingMessage   = 9 // Ping (keep-alive)
	PongMessage   = 10 // Pong (resposta ao ping)
)
```

### Gerenciando Múltiplas Conexões

```go
type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte, 256),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
			log.Println("Cliente registrado. Total:", len(h.clients))

		case conn := <-h.unregister:
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				close(conn)
				log.Println("Cliente removido. Total:", len(h.clients))
			}

		case message := <-h.broadcast:
			// Enviar para todos os clientes
			for conn := range h.clients {
				select {
				case conn.WriteMessage(websocket.TextMessage, message) <- nil:
				default:
					close(conn)
					delete(h.clients, conn)
				}
			}
		}
	}
}

func (h *Hub) handleConnection(conn *websocket.Conn) {
	defer func() {
		h.unregister <- conn
		conn.Close()
	}()

	h.register <- conn

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Broadcast para todos
		h.broadcast <- message
	}
}
```

---

## Melody - Framework WebSocket Minimalista

**Melody** é um framework WebSocket minimalista e elegante para Go, construído sobre `gorilla/websocket`. Ele fornece abstrações úteis como salas (rooms), broadcasting e gerenciamento de sessões.

### Por que Melody?

- **Simples**: API limpa e intuitiva
- **Leve**: Minimalista, sem dependências pesadas
- **Funcionalidades**: Rooms, broadcasting, ping/pong automático
- **Integração**: Funciona com qualquer framework web (Gin, Echo, net/http)

### Instalação

```bash
go get github.com/olahol/melody
```

### Exemplo Básico com Melody

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Quando cliente conecta
	m.HandleConnect(func(s *melody.Session) {
		log.Printf("Cliente conectado: %s", s.Request.RemoteAddr)
		s.Write([]byte("Bem-vindo ao chat!"))
	})

	// Quando cliente desconecta
	m.HandleDisconnect(func(s *melody.Session) {
		log.Printf("Cliente desconectado: %s", s.Request.RemoteAddr)
	})

	// Quando recebe mensagem
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Printf("Mensagem recebida: %s", string(msg))
		// Broadcast para todos (exceto o remetente)
		m.BroadcastOthers(msg, s)
	})

	log.Println("Servidor rodando em :8080")
	r.Run(":8080")
}
```

### Salas (Rooms) com Melody

Salas permitem agrupar conexões e enviar mensagens apenas para um grupo específico.

```go
func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		// Obter sala da query string
		room := s.Request.URL.Query().Get("room")
		if room == "" {
			room = "default"
		}

		// Entrar na sala
		s.Set("room", room)
		m.HandleJoinRoom(s, room)

		log.Printf("Cliente entrou na sala: %s", room)
		m.BroadcastToRoom(room, []byte(fmt.Sprintf("Novo usuário entrou na sala %s", room)))
	})

	m.HandleDisconnect(func(s *melody.Session) {
		room, _ := s.Get("room")
		if roomStr, ok := room.(string); ok {
			m.BroadcastToRoom(roomStr, []byte("Um usuário saiu da sala"))
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		room, _ := s.Get("room")
		if roomStr, ok := room.(string); ok {
			// Enviar apenas para a sala
			m.BroadcastToRoom(roomStr, msg)
		}
	})

	r.Run(":8080")
}
```

### Configurações do Melody

```go
m := melody.New()

// Configurar limites
m.Config.MaxMessageSize = 512  // Tamanho máximo da mensagem (bytes)
m.Config.MessageBufferSize = 256  // Buffer de mensagens

// Ping/Pong automático
m.Config.PingPeriod = 60 * time.Second
m.Config.PongWait = 70 * time.Second
m.Config.WriteWait = 10 * time.Second

// Headers customizados
m.Config.Upgrader.CheckOrigin = func(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	// Validar origem
	return origin == "https://meusite.com"
}
```

### Exemplo Completo - Chat com Salas

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type Message struct {
	Room      string    `json:"room"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "chat.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		room := s.Request.URL.Query().Get("room")
		username := s.Request.URL.Query().Get("username")

		if room == "" {
			room = "general"
		}
		if username == "" {
			username = "Anônimo"
		}

		s.Set("room", room)
		s.Set("username", username)
		m.HandleJoinRoom(s, room)

		msg := Message{
			Room:      room,
			Username:  "Sistema",
			Text:      username + " entrou na sala",
			Timestamp: time.Now(),
		}

		msgBytes, _ := json.Marshal(msg)
		m.BroadcastToRoom(room, msgBytes)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		room, _ := s.Get("room")
		username, _ := s.Get("username")

		if roomStr, ok := room.(string); ok {
			if usernameStr, ok := username.(string); ok {
				msg := Message{
					Room:      roomStr,
					Username:  "Sistema",
					Text:      usernameStr + " saiu da sala",
					Timestamp: time.Now(),
				}
				msgBytes, _ := json.Marshal(msg)
				m.BroadcastToRoom(roomStr, msgBytes)
			}
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		room, _ := s.Get("room")
		username, _ := s.Get("username")

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			return
		}

		message.Room, _ = room.(string)
		message.Username, _ = username.(string)
		message.Timestamp = time.Now()

		msgBytes, _ := json.Marshal(message)
		m.BroadcastToRoom(message.Room, msgBytes)
	})

	r.Run(":8080")
}
```

---

## Server-Sent Events (SSE) com Go

SSE é mais simples que WebSocket quando você só precisa enviar dados do servidor para o cliente.

### Exemplo Básico SSE

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Headers para SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE não suportado", http.StatusInternalServerError)
		return
	}

	// Enviar eventos
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			log.Println("Cliente desconectado")
			return
		case <-ticker.C:
			data := fmt.Sprintf("data: %s\n\n", time.Now().Format(time.RFC3339))
			fmt.Fprint(w, data)
			flusher.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sse.html")
	})

	log.Println("Servidor SSE rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Cliente SSE (HTML)

```html
<!DOCTYPE html>
<html>
<head>
	<title>SSE Test</title>
</head>
<body>
	<div id="events"></div>

	<script>
		const eventSource = new EventSource("/events");
		const eventsDiv = document.getElementById("events");

		eventSource.onmessage = function(event) {
			const p = document.createElement("p");
			p.textContent = "Evento: " + event.data;
			eventsDiv.appendChild(p);
		};

		eventSource.onerror = function(error) {
			console.error("Erro SSE:", error);
		};
	</script>
</body>
</html>
```

### SSE com Eventos Nomeados

```go
func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher := w.(http.Flusher)

	// Evento nomeado
	fmt.Fprintf(w, "event: welcome\ndata: Bem-vindo!\n\n")
	flusher.Flush()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-ticker.C:
			// Evento nomeado
			fmt.Fprintf(w, "event: update\ndata: %s\n\n", time.Now().Format(time.RFC3339))
			flusher.Flush()
		}
	}
}
```

```javascript
// Cliente
eventSource.addEventListener("welcome", function(event) {
	console.log("Bem-vindo:", event.data);
});

eventSource.addEventListener("update", function(event) {
	console.log("Atualização:", event.data);
});
```

---

## Centrifugo - Servidor de Mensagens em Tempo Real

**Centrifugo** é um servidor de mensagens em tempo real que pode ser usado com qualquer linguagem, incluindo Go. Ele fornece WebSocket, SSE e HTTP streaming.

### Características do Centrifugo

- **Escalável**: Suporta milhões de conexões
- **Redis**: Backend Redis para escalabilidade horizontal
- **Canais**: Sistema de canais/pub-sub
- **Presença**: Informações de quem está online
- **Histórico**: Histórico de mensagens
- **Múltiplos Protocolos**: WebSocket, SSE, HTTP streaming

### Quando Usar Centrifugo?

✅ **Use quando:**
- Precisa de alta escalabilidade
- Múltiplos servidores (horizontal scaling)
- Precisa de histórico de mensagens
- Sistema complexo de canais
- Precisa de informações de presença

❌ **Evite quando:**
- Aplicação simples, poucas conexões
- Não precisa de escalabilidade horizontal
- Quer controle total do código

### Instalação do Centrifugo

```bash
# Via Docker (recomendado)
docker run --ulimit nofile=65536:65536 -d \
  -p 8000:8000 \
  centrifugo/centrifugo centrifugo --config=config.json

# Ou baixar binário
# https://github.com/centrifugal/centrifugo/releases
```

### Configuração Básica (config.json)

```json
{
  "token_hmac_secret_key": "seu-secret-key-aqui",
  "admin_password": "admin",
  "admin_secret": "admin-secret",
  "api_key": "api-key",
  "allowed_origins": ["http://localhost:3000"]
}
```

### Cliente Go para Centrifugo

```bash
go get github.com/centrifugal/centrifuge-go
```

```go
package main

import (
	"log"
	"time"

	"github.com/centrifugal/centrifuge-go"
)

func main() {
	// Criar cliente
	client := centrifuge.NewJsonClient(
		"ws://localhost:8000/connection/websocket",
		centrifuge.Config{
			Token: generateToken("user123"), // Token JWT
		},
	)

	// Conectar
	if err := client.Connect(); err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer client.Close()

	// Inscrever em um canal
	sub, err := client.NewSubscription("news")
	if err != nil {
		log.Fatal("Erro ao criar subscription:", err)
	}

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		log.Printf("Mensagem recebida: %s", string(e.Data))
	})

	if err := sub.Subscribe(); err != nil {
		log.Fatal("Erro ao inscrever:", err)
	}

	// Publicar mensagem
	sub.Publish([]byte(`{"text": "Olá do Go!"}`))

	// Manter conexão
	time.Sleep(10 * time.Second)
}

// Gerar token JWT (simplificado)
func generateToken(userID string) string {
	// Em produção, use uma biblioteca JWT adequada
	// Exemplo com github.com/golang-jwt/jwt
	return "token-jwt-aqui"
}
```

### Integração Centrifugo com Go Backend

```go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CentrifugoConfig struct {
	SecretKey string
	APIURL    string
	APIKey    string
}

func generateConnectionToken(userID string, secret string) string {
	// Token simples (em produção, use JWT)
	payload := map[string]interface{}{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	payloadJSON, _ := json.Marshal(payload)
	payloadB64 := base64.URLEncoding.EncodeToString(payloadJSON)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payloadB64))
	signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))

	return payloadB64 + "." + signature
}

func main() {
	r := gin.Default()
	config := CentrifugoConfig{
		SecretKey: "seu-secret-key",
		APIURL:    "http://localhost:8000/api",
		APIKey:    "api-key",
	}

	// Endpoint para obter token de conexão
	r.GET("/centrifuge/token", func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.JSON(400, gin.H{"error": "user_id requerido"})
			return
		}

		token := generateConnectionToken(userID, config.SecretKey)
		c.JSON(200, gin.H{
			"token": token,
			"url":   "ws://localhost:8000/connection/websocket",
		})
	})

	// Publicar mensagem via API do Centrifugo
	r.POST("/centrifuge/publish", func(c *gin.Context) {
		var req struct {
			Channel string      `json:"channel"`
			Data    interface{} `json:"data"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Chamar API do Centrifugo
		// (implementar chamada HTTP para Centrifugo)
		c.JSON(200, gin.H{"status": "published"})
	})

	r.Run(":8080")
}
```

---

## Comparação: WebSocket vs SSE vs Centrifugo

| Característica | WebSocket | SSE | Centrifugo |
|---------------|-----------|-----|------------|
| **Bidirecional** | ✅ Sim | ❌ Não | ✅ Sim |
| **Compatibilidade** | Boa | Excelente | Boa |
| **Complexidade** | Média | Baixa | Alta |
| **Escalabilidade** | Média | Média | Excelente |
| **Histórico** | Manual | Manual | ✅ Automático |
| **Presença** | Manual | Manual | ✅ Automático |
| **Uso** | Chat, jogos | Notificações, feeds | Sistemas complexos |

---

## Boas Práticas

### 1. **Validação de Origem (CORS)**

```go
upgrader.CheckOrigin = func(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	allowedOrigins := []string{
		"https://meusite.com",
		"https://app.meusite.com",
	}
	
	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}
```

### 2. **Limites de Mensagem**

```go
m.Config.MaxMessageSize = 1024 * 1024 // 1MB máximo
```

### 3. **Timeouts e Keep-Alive**

```go
m.Config.PingPeriod = 30 * time.Second
m.Config.PongWait = 40 * time.Second
m.Config.WriteWait = 10 * time.Second
```

### 4. **Tratamento de Erros**

```go
m.HandleError(func(s *melody.Session, err error) {
	log.Printf("Erro na sessão: %v", err)
	// Fechar conexão se necessário
	s.Close()
})
```

### 5. **Rate Limiting**

```go
type RateLimiter struct {
	clients map[*melody.Session]time.Time
	mu      sync.Mutex
}

func (rl *RateLimiter) Allow(s *melody.Session) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	lastMessage, exists := rl.clients[s]
	if !exists {
		rl.clients[s] = time.Now()
		return true
	}

	if time.Since(lastMessage) < 100*time.Millisecond {
		return false // Muito rápido
	}

	rl.clients[s] = time.Now()
	return true
}
```

### 6. **Autenticação**

```go
m.HandleConnect(func(s *melody.Session) {
	token := s.Request.URL.Query().Get("token")
	
	// Validar token
	userID, err := validateToken(token)
	if err != nil {
		s.Close()
		return
	}

	s.Set("user_id", userID)
})
```

### 7. **Logging e Monitoramento**

```go
m.HandleConnect(func(s *melody.Session) {
	log.Printf("Conexão: %s, Total: %d", 
		s.Request.RemoteAddr, 
		len(m.Sessions))
})

// Métricas
func getMetrics(m *melody.Melody) map[string]interface{} {
	return map[string]interface{}{
		"total_connections": len(m.Sessions),
		"rooms":             len(m.Rooms),
	}
}
```

---

## Exemplo Completo - Sistema de Notificações

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"` // info, warning, error, success
	Timestamp time.Time `json:"timestamp"`
	Read      bool      `json:"read"`
}

type NotificationService struct {
	mu            sync.RWMutex
	notifications map[string][]Notification
	melody        *melody.Melody
}

func NewNotificationService(m *melody.Melody) *NotificationService {
	return &NotificationService{
		notifications: make(map[string][]Notification),
		melody:        m,
	}
}

func (ns *NotificationService) SendNotification(userID string, notif Notification) {
	ns.mu.Lock()
	ns.notifications[userID] = append(ns.notifications[userID], notif)
	ns.mu.Unlock()

	// Enviar via WebSocket se usuário estiver conectado
	notifBytes, _ := json.Marshal(notif)
	ns.melody.BroadcastFilter(notifBytes, func(s *melody.Session) bool {
		uid, _ := s.Get("user_id")
		return uid == userID
	})
}

func (ns *NotificationService) GetNotifications(userID string) []Notification {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.notifications[userID]
}

func main() {
	r := gin.Default()
	m := melody.New()
	ns := NewNotificationService(m)

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Autenticação e registro
	m.HandleConnect(func(s *melody.Session) {
		userID := s.Request.URL.Query().Get("user_id")
		if userID == "" {
			s.Close()
			return
		}

		s.Set("user_id", userID)

		// Enviar notificações pendentes
		notifications := ns.GetNotifications(userID)
		for _, notif := range notifications {
			if !notif.Read {
				notifBytes, _ := json.Marshal(notif)
				s.Write(notifBytes)
			}
		}
	})

	// API para criar notificação
	r.POST("/notifications", func(c *gin.Context) {
		var req struct {
			UserID  string `json:"user_id" binding:"required"`
			Title   string `json:"title" binding:"required"`
			Message string `json:"message" binding:"required"`
			Type    string `json:"type"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Type == "" {
			req.Type = "info"
		}

		notif := Notification{
			ID:        generateID(),
			UserID:    req.UserID,
			Title:     req.Title,
			Message:   req.Message,
			Type:      req.Type,
			Timestamp: time.Now(),
			Read:      false,
		}

		ns.SendNotification(req.UserID, notif)
		c.JSON(201, notif)
	})

	// API para listar notificações
	r.GET("/notifications/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		notifications := ns.GetNotifications(userID)
		c.JSON(200, notifications)
	})

	r.Run(":8080")
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
```

---

## Conclusão

Nesta aula, exploramos comunicação em tempo real em Go:

1. **Conceitos Fundamentais**: O que é realtime communication e quando usar
2. **WebSockets**: Protocolo bidirecional com `gorilla/websocket`
3. **Melody**: Framework minimalista para WebSockets
4. **Server-Sent Events**: Para comunicação unidirecional servidor→cliente
5. **Centrifugo**: Servidor escalável de mensagens em tempo real
6. **Boas Práticas**: Validação, rate limiting, autenticação, monitoramento

Go é uma excelente escolha para aplicações em tempo real devido à sua concorrência nativa, performance e simplicidade. Com as ferramentas que vimos, você pode construir desde chats simples até sistemas complexos de mensagens escaláveis.

Na próxima aula, vamos explorar mais padrões avançados e otimizações para sistemas em tempo real!

Até lá, experimente criar seu próprio chat ou sistema de notificações. Pratique com WebSockets, explore Melody e teste diferentes cenários!

