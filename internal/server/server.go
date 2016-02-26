package server

import (
	"time"

	"github.com/luizbranco/invasion/internal/client"
)

type Message struct {
	Data   []byte
	Client client.Client
}

type Server struct {
	clients   []client.Client
	createdAt time.Time
	Msg       chan Message
}

func newServer() *Server {
	return &Server{
		createdAt: time.Now(),
		Msg:       make(chan Message),
	}
}

func (s *Server) run() {
	for {
		_, ok := <-s.Msg
		if !ok {
			break
		}
	}
}

func (s *Server) addClient(c client.Client) {
}

func (s *Server) removeClient(c client.Client) {
	for i, client := range s.clients {
		if c == client {
			last := len(s.clients) - 1
			s.clients[i] = s.clients[last]
			s.clients = s.clients[:last]
		}
	}
}

func (s *Server) broadcast(msg []byte) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}
