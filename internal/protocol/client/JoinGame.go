// automatically generated, do not modify

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type JoinGame struct {
	_tab flatbuffers.Table
}

func (rcv *JoinGame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func JoinGameStart(builder *flatbuffers.Builder) { builder.StartObject(0) }
func JoinGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
