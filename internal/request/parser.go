package request

import (
	"github.com/google/flatbuffers/go"
	clt "github.com/luizbranco/invasion/internal/client"
	"github.com/luizbranco/invasion/internal/protocol/client"
	"github.com/luizbranco/invasion/internal/protocol/server"
	"github.com/luizbranco/invasion/internal/response"
)

type Request struct {
	Type   byte
	Token  string
	GameID string
	Data   []byte
	Client clt.Client
}

func Parse(data []byte, c clt.Client) (res response.Response) {
	req := Request{Data: data, Client: c}

	msg := client.GetRootAsMessage(data, 0)

	defer func() {
		err := recover()
		if err != nil {
			res = response.NewError(server.CodeBadRequest, "invalid request format")
			return
		}
	}()

	unionTable := &flatbuffers.Table{}

	if !msg.Request(unionTable) {
		return response.NewError(server.CodeInvalidRequestType, "request type is not valid")
	}

	req.Type = msg.RequestType()

	switch msg.RequestType() {
	case client.RequestCreateAccount:
		res = parseCreateAccount(unionTable)
	case client.RequestCreateGame:
		//TODO
	default:
		res = parseGameRequest(msg, unionTable)
	}

	return res
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

func parseCreateAccount(unionTable *flatbuffers.Table) response.Response {
	acc := &client.CreateAccount{}
	acc.Init(unionTable.Bytes, unionTable.Pos)
	email := string(acc.Email())
	if email == "" {
		return response.NewError(server.CodeInvalidEmail, "email is invalid")
	}
	//TODO validate duplicated email
	//TODO create account and send response
	return nil
}

func parseGameRequest(msg *client.Message, unionTable *flatbuffers.Table) response.Response {
	token := string(msg.Token())
	if token == "" {
		return response.NewError(server.CodeAuthorizationDenied, "token is invalid")
	}
	//TODO validate token

	id := string(msg.GameId())
	if id == "" {
		return response.NewError(server.CodeGameNotFound, "game_id is invalid")
	}
	//TODO find game and send request
	return nil
}
