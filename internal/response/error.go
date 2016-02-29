package response

import (
	"github.com/google/flatbuffers/go"
	"github.com/luizbranco/invasion/internal/protocol/client"
	"github.com/luizbranco/invasion/internal/protocol/server"
)

type Error struct {
	data []byte
}

func NewError(code uint32, description string) *Error {
	builder := flatbuffers.NewBuilder(0)

	desc := builder.CreateString(description)
	server.ErrorStart(builder)
	server.ErrorAddCode(builder, code)
	server.ErrorAddDescription(builder, desc)
	resp := server.ErrorEnd(builder)

	server.MessageStart(builder)
	server.MessageAddResponseType(builder, server.ResponseError)
	server.MessageAddResponse(builder, resp)
	msg := client.MessageEnd(builder)

	builder.Finish(msg)
	data := builder.FinishedBytes()

	return &Error{data: data}
}

func (e *Error) Data() []byte {
	return e.data
}
