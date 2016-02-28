// automatically generated, do not modify

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CreateGame struct {
	_tab flatbuffers.Table
}

func (rcv *CreateGame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func CreateGameStart(builder *flatbuffers.Builder) { builder.StartObject(0) }
func CreateGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
