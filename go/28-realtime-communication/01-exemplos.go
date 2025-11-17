package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/olahol/melody"
)

// ============================================
// EXEMPLO 1: WebSocket Básico com gorilla/websocket
// ============================================

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Em produção, valide a origem!
	},
}

func websocketBasicHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao fazer upgrade:", err)
		return
	}
	defer conn.Close()

	log.Println("Cliente conectado")

	// Goroutine para ler mensagens
	go func() {
		for {
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
	}()

	// Enviar mensagens periódicas
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			message := fmt.Sprintf("Servidor: %s", time.Now().Format(time.RFC3339))
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println("Erro ao escrever:", err)
				return
			}
		}
	}
}

// ============================================
// EXEMPLO 2: Hub para Múltiplas Conexões
// ============================================

type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.RWMutex
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
			h.mu.Lock()
			h.clients[conn] = true
			h.mu.Unlock()
			log.Printf("Cliente registrado. Total: %d", len(h.clients))

		case conn := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
				log.Printf("Cliente removido. Total: %d", len(h.clients))
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			for conn := range h.clients {
				select {
				default:
					conn.WriteMessage(websocket.TextMessage, message)
				}
			}
			h.mu.RUnlock()
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
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Erro: %v", err)
			}
			break
		}

		// Broadcast para todos
		h.broadcast <- message
	}
}

func hubWebSocketHandler(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Erro ao fazer upgrade:", err)
			return
		}

		hub.handleConnection(conn)
	}
}

// ============================================
// EXEMPLO 3: Melody - Chat Simples
// ============================================

func melodyChatExample() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "chat.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		log.Printf("Cliente conectado: %s", s.Request.RemoteAddr)
		s.Write([]byte("Bem-vindo ao chat!"))
	})

	m.HandleDisconnect(func(s *melody.Session) {
		log.Printf("Cliente desconectado: %s", s.Request.RemoteAddr)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Printf("Mensagem recebida: %s", string(msg))
		// Broadcast para todos (exceto o remetente)
		m.BroadcastOthers(msg, s)
	})

	log.Println("Servidor Melody rodando em :8080")
	r.Run(":8080")
}

// ============================================
// EXEMPLO 4: Melody com Salas (Rooms)
// ============================================

type ChatMessage struct {
	Room      string    `json:"room"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func melodyRoomsExample() {
	r := gin.Default()
	m := melody.New()

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

		msg := ChatMessage{
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
				msg := ChatMessage{
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

		var message ChatMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			return
		}

		message.Room, _ = room.(string)
		message.Username, _ = username.(string)
		message.Timestamp = time.Now()

		msgBytes, _ := json.Marshal(message)
		m.BroadcastToRoom(message.Room, msgBytes)
	})

	r.Run(":8081")
}

// ============================================
// EXEMPLO 5: Server-Sent Events (SSE)
// ============================================

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

func sseExample() {
	http.HandleFunc("/events", sseHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sse.html")
	})

	log.Println("Servidor SSE rodando em :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// ============================================
// EXEMPLO 6: Sistema de Notificações
// ============================================

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

func notificationSystemExample() {
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
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
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

	r.Run(":8083")
}

// ============================================
// EXEMPLO 7: Rate Limiting para WebSocket
// ============================================

type RateLimiter struct {
	clients map[*melody.Session]time.Time
	mu      sync.Mutex
	limit   time.Duration
}

func NewRateLimiter(limit time.Duration) *RateLimiter {
	return &RateLimiter{
		clients: make(map[*melody.Session]time.Time),
		limit:   limit,
	}
}

func (rl *RateLimiter) Allow(s *melody.Session) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	lastMessage, exists := rl.clients[s]
	if !exists {
		rl.clients[s] = time.Now()
		return true
	}

	if time.Since(lastMessage) < rl.limit {
		return false // Muito rápido
	}

	rl.clients[s] = time.Now()
	return true
}

func (rl *RateLimiter) Remove(s *melody.Session) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.clients, s)
}

func rateLimitedChatExample() {
	r := gin.Default()
	m := melody.New()
	limiter := NewRateLimiter(100 * time.Millisecond)

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		log.Println("Cliente conectado")
	})

	m.HandleDisconnect(func(s *melody.Session) {
		limiter.Remove(s)
		log.Println("Cliente desconectado")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if !limiter.Allow(s) {
			s.Write([]byte("Muitas mensagens! Aguarde um pouco."))
			return
		}

		m.BroadcastOthers(msg, s)
	})

	r.Run(":8084")
}

// ============================================
// MAIN - Descomente o exemplo que deseja testar
// ============================================

func main() {
	// Exemplo 1: WebSocket básico
	// http.HandleFunc("/ws", websocketBasicHandler)
	// log.Println("Servidor WebSocket básico rodando em :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// Exemplo 2: Hub com múltiplas conexões
	// hub := newHub()
	// go hub.run()
	// http.HandleFunc("/ws", hubWebSocketHandler(hub))
	// log.Println("Servidor Hub rodando em :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// Exemplo 3: Melody chat simples
	// melodyChatExample()

	// Exemplo 4: Melody com salas
	// melodyRoomsExample()

	// Exemplo 5: Server-Sent Events
	// sseExample()

	// Exemplo 6: Sistema de notificações
	// notificationSystemExample()

	// Exemplo 7: Rate limiting
	// rateLimitedChatExample()

	fmt.Println("Descomente um dos exemplos acima para testar!")
}


