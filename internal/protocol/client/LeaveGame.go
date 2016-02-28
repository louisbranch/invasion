// automatically generated, do not modify

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type LeaveGame struct {
	_tab flatbuffers.Table
}

func (rcv *LeaveGame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LeaveGame) GameId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LeaveGameStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func LeaveGameAddGameId(builder *flatbuffers.Builder, gameId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(gameId), 0) }
func LeaveGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
