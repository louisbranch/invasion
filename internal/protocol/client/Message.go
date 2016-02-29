// automatically generated, do not modify

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Message struct {
	_tab flatbuffers.Table
}

func GetRootAsMessage(buf []byte, offset flatbuffers.UOffsetT) *Message {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Message{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Message) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Message) Token() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Message) GameId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Message) RequestType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Message) Request(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func MessageStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func MessageAddToken(builder *flatbuffers.Builder, token flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(token), 0) }
func MessageAddGameId(builder *flatbuffers.Builder, gameId flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(gameId), 0) }
func MessageAddRequestType(builder *flatbuffers.Builder, requestType byte) { builder.PrependByteSlot(2, requestType, 0) }
func MessageAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(request), 0) }
func MessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
