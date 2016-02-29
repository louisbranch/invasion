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

func LeaveGameStart(builder *flatbuffers.Builder) { builder.StartObject(0) }
func LeaveGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
