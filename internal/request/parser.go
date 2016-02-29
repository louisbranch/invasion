package request

import (
	"github.com/google/flatbuffers/go"
	clt "github.com/luizbranco/invasion/internal/client"
	"github.com/luizbranco/invasion/internal/protocol/client"
)

type Request struct {
	Type   byte
	Token  string
	GameID string
	Data   []byte
	Client clt.Client
}

func New(data []byte, c clt.Client) (Request, error) {
	req := Request{Data: data, Client: c}

	msg := client.GetRootAsMessage(data, 0)

	defer func() {
		err := recover()
		if err != nil {
			//TODO invalid format
		}
	}()

	unionTable := new(flatbuffers.Table)

	if msg.Request(unionTable) {
		req.Type = msg.RequestType()

		switch msg.RequestType() {
		case client.RequestCreateAccount:
			acc := &client.CreateAccount{}
			acc.Init(unionTable.Bytes, unionTable.Pos)
		case client.RequestCreateGame:

		default:
			//TODO extract token and game id
		}
	} else {
		//TODO
	}

	return req, nil
}

func ChatMessage() {
	/*
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
	*/
}
