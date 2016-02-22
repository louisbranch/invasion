package server

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Hub struct {
	ids     int
	servers map[string]*Server
	mutex   sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		servers: make(map[string]*Server),
	}
}

func (hub *Hub) NewServer() string {
	hub.mutex.Lock()
	hub.ids++
	id := strconv.Itoa(hub.ids)
	server := newServer()
	hub.servers[id] = server
	hub.fork(id, server)
	hub.mutex.Unlock()
	return id
}

func (hub *Hub) fork(id string, server *Server) {
	go func() {
		server.run()
		hub.remove(id)
	}()
}

func (hub *Hub) remove(id string) {
	hub.mutex.Lock()
	delete(hub.servers, id)
	hub.mutex.Unlock()
}

func (hub *Hub) Status() []string {
	var status []string
	hub.mutex.RLock()
	for id, g := range hub.servers {
		s := fmt.Sprintf("Server %s has %d players and is running for %s",
			id, len(g.clients), time.Since(g.createdAt),
		)
		status = append(status, s)
	}
	hub.mutex.RUnlock()
	return status
}

func (hub *Hub) GetServer(i string) (*Server, error) {
	hub.mutex.RLock()
	server, ok := hub.servers[i]
	hub.mutex.RUnlock()
	if !ok {
		return nil, errors.New("server not found")
	}
	return server, nil
}
