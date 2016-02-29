package server

import (
	"time"

	"github.com/luizbranco/invasion/internal/client"
	"github.com/luizbranco/invasion/internal/request"
)

type Server struct {
	clients   []client.Client
	createdAt time.Time
	In        chan request.Request
}

func newServer() *Server {
	return &Server{
		createdAt: time.Now(),
		In:        make(chan request.Request),
	}
}

func (s *Server) run() {
	for {
		_, ok := <-s.In
		if !ok {
			break
		}
	}
}

func (s *Server) addClient(c client.Client) {
	//TODO check for duplicated clients, remove old one (?)
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
