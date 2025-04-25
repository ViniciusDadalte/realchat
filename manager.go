package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

func checkOrigin(r *http.Request) bool {

	origin := r.Header.Get("Origin")

	switch origin {
	case "https://localhost:8080":
		return true
	default:
		return false
	}
}

type Manager struct {
    clients  ClientList
    handlers map[string]EventHandler
    otps     RetentionMap
    logins   map[string]string 
    sync.RWMutex
}

func NewManager(ctx context.Context) *Manager {
    m := &Manager{
        clients:  make(ClientList),
        handlers: make(map[string]EventHandler),
        otps:     NewRetentionMap(ctx, 5*time.Second),
        logins: map[string]string{ 
            "percy":    "123",
            "vinicius": "171",
            "gabi":     "6969",
        },
    }
    m.setupEventHandlers()
    return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessageHandler
	m.handlers[EventChangeRoom] = ChatRoomHandler
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) loginHandler(w http.ResponseWriter, r *http.Request) {
    type userLoginRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    var req userLoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if password, ok := m.logins[req.Username]; ok && password == req.Password {
        type response struct {
            OTP string `json:"otp"`
        }

        otp := m.otps.NewOTP()

        resp := response{
            OTP: otp.Key,
        }

        data, err := json.Marshal(resp)
        if err != nil {
            log.Println(err)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write(data)
        return
    }

    w.WriteHeader(http.StatusUnauthorized)
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {

	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !m.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("New connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}
