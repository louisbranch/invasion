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

func (rcv *JoinGame) GameId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func JoinGameStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func JoinGameAddGameId(builder *flatbuffers.Builder, gameId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(gameId), 0) }
func JoinGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
