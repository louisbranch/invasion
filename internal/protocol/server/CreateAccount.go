// automatically generated, do not modify

package server

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

func (rcv *CreateAccount) Token() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CreateAccountStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func CreateAccountAddToken(builder *flatbuffers.Builder, token flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(token), 0) }
func CreateAccountEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
