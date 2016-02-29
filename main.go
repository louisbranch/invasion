package main

import (
	"log"
	"net/http"

	"github.com/luizbranco/invasion/internal/views"

	"golang.org/x/net/websocket"
)

func handler(ws *websocket.Conn) {
	for {
		var data []byte
		err := websocket.Message.Receive(ws, &data)
		if err != nil {
			break
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.Handle("/ws", websocket.Handler(handler))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			views.Render(w, "index", nil)
		} else {
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
