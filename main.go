package main

import (
	"log"
	"net/http"

	"github.com/luizbranco/invasion/internal/client"
	"github.com/luizbranco/invasion/internal/message"
	"github.com/luizbranco/invasion/internal/server"
	"github.com/luizbranco/invasion/internal/views"

	"golang.org/x/net/websocket"
)

func Handler(s *server.Server, conn *websocket.Conn) {
	clt := client.NewWsClient(conn)
	for {
		var buf = make([]byte, 512) //FIXME msg limit
		msg := message.Message{Client: clt}
		_, err := conn.Read(buf)
		if err != nil {
			msg.Code = message.Disconnect
			s.Msg <- msg
			break
		}
		msg.UnmarshalText(buf)
		s.Msg <- msg
	}
}

func main() {
	hub := server.NewHub()

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/ws/"):]
		s, err := hub.GetServer(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		handler := websocket.Handler(func(conn *websocket.Conn) {
			Handler(s, conn)
		})
		handler.ServeHTTP(w, r)
	})

	http.HandleFunc("/join/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/join/"):]
		switch r.Method {
		case "GET":
			_, err := hub.GetServer(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			views.Render(w, "game", nil)
		default:
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/games/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			id := hub.NewServer()
			http.Redirect(w, r, "/join/"+id, 302)
		case "GET":
			views.Render(w, "games", hub.Status())
		default:
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
