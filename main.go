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

/*
func notMain() {
	j := join()
	l := leave()
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY

	f1, err := os.OpenFile("join.dat", flags, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f1.Write(j); err != nil {
		log.Fatal(err)
	}

	f2, err := os.OpenFile("leave.dat", flags, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f2.Write(l); err != nil {
		log.Fatal(err)
	}

	for _, f := range []string{"join", "leave"} {
		buf, err := ioutil.ReadFile(f + ".dat")
		if err != nil {
			log.Fatal(err)
		}

		msg := client.GetRootAsMessage(buf, 0)

		unionTable := new(flatbuffers.Table)

		if msg.Code(unionTable) {
			switch msg.CodeType() {
			case client.CodeJoin:
				join := &client.Join{}
				join.Init(unionTable.Bytes, unionTable.Pos)
				fmt.Printf("Joined %s\n", join.Name())
			case client.CodeLeave:
				leave := &client.Leave{}
				leave.Init(unionTable.Bytes, unionTable.Pos)
				fmt.Printf("Left %s\n", leave.Token())
			default:
				log.Fatal("Unknown")
			}
		}
	}
}

func join() []byte {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("luiz")
	email := builder.CreateString("me@luizbranco.com")
	client.JoinStart(builder)
	client.JoinAddName(builder, name)
	client.JoinAddEmail(builder, email)
	join := client.JoinEnd(builder)

	client.MessageStart(builder)
	client.MessageAddCodeType(builder, client.CodeJoin)
	client.MessageAddCode(builder, join)
	msg := client.MessageEnd(builder)

	builder.Finish(msg)

	return builder.FinishedBytes()
}

func leave() []byte {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("SeCreTT")
	client.LeaveStart(builder)
	client.LeaveAddToken(builder, name)
	leave := client.LeaveEnd(builder)

	client.MessageStart(builder)
	client.MessageAddCodeType(builder, client.CodeLeave)
	client.MessageAddCode(builder, leave)
	msg := client.MessageEnd(builder)

	builder.Finish(msg)

	return builder.FinishedBytes()
}
*/
