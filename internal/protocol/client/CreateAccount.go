// automatically generated, do not modify

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CreateAccount struct {
	_tab flatbuffers.Table
}

func (rcv *CreateAccount) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CreateAccount) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CreateAccount) Email() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CreateAccountStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func CreateAccountAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0) }
func CreateAccountAddEmail(builder *flatbuffers.Builder, email flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(email), 0) }
func CreateAccountEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
