// automatically generated, do not modify

package server

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

func (rcv *CreateGame) GameId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CreateGameStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func CreateGameAddGameId(builder *flatbuffers.Builder, gameId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(gameId), 0) }
func CreateGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
