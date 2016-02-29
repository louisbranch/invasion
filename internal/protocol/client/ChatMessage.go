// automatically generated, do not modify

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type ChatMessage struct {
	_tab flatbuffers.Table
}

func (rcv *ChatMessage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ChatMessage) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChatMessage) Recipients(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *ChatMessage) RecipientsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ChatMessageStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func ChatMessageAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(message), 0) }
func ChatMessageAddRecipients(builder *flatbuffers.Builder, recipients flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(recipients), 0) }
func ChatMessageStartRecipientsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func ChatMessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
