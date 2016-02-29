package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/flatbuffers/go"
	"github.com/luizbranco/invasion/internal/protocol/client"
	"github.com/luizbranco/invasion/internal/protocol/server"
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
		msg := client.GetRootAsMessage(data, 0)

		unionTable := new(flatbuffers.Table)

		if msg.Request(unionTable) {
			switch msg.RequestType() {
			case client.RequestCreateAccount:
				acc := &client.CreateAccount{}
				acc.Init(unionTable.Bytes, unionTable.Pos)
				fmt.Printf("Joined %s\n", acc.Name())
			case client.RequestChatMessage:
				chat := &client.ChatMessage{}
				chat.Init(unionTable.Bytes, unionTable.Pos)
				builder := flatbuffers.NewBuilder(0)

				txt := builder.CreateString(string(chat.Message()))
				server.ChatMessageStart(builder)
				server.ChatMessageAddMessage(builder, txt)
				resp := client.ChatMessageEnd(builder)

				server.MessageStart(builder)
				server.MessageAddResponseType(builder, server.ResponseChatMessage)
				server.MessageAddResponse(builder, resp)
				msg := client.MessageEnd(builder)

				builder.Finish(msg)

				websocket.Message.Send(ws, builder.FinishedBytes())
			default:
				log.Fatal("Unknown type")
			}
		} else {
			log.Fatalf("No type %v", data)
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
