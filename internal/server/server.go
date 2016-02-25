package server

import (
	"fmt"
	"time"

	"github.com/luizbranco/invasion/internal/client"
	"github.com/luizbranco/invasion/internal/message"
)

type Command int

const (
	Start Command = iota
	Stop
	Disconnect
)

type Server struct {
	clients   []client.Client
	createdAt time.Time
	Cmd       chan Command
	Msg       chan message.Message
}

func newServer() *Server {
	return &Server{
		createdAt: time.Now(),
		Cmd:       make(chan Command),
		Msg:       make(chan message.Message),
	}
}

func (s *Server) run() {
Loop:
	for {
		select {
		case cmd := <-s.Cmd:
			switch cmd {
			case Stop:
				break Loop
			case Start:
			default:
				panic(fmt.Sprintf("Unknown command %d", cmd))
			}
		case msg := <-s.Msg:

			switch msg.Code {
			case message.Join:
				s.clients = append(s.clients, msg.Client)
				msg.Client.Write([]byte(";)"))
			case message.Leave, message.Disconnect:
				for i, c := range s.clients {
					if c == msg.Client {
						s.removeClient(i)
						break
					}
				}
			case message.ChatMessage:
				s.broadcast(msg.Content)
			case message.SetName:
			default:
				//TODO reply back
			}

		}
	}
}

func (s *Server) removeClient(i int) {
	last := len(s.clients) - 1
	s.clients[i] = s.clients[last]
	s.clients = s.clients[:last]
}

func (s *Server) broadcast(msg []byte) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}
