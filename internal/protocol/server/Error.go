// automatically generated, do not modify

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Error struct {
	_tab flatbuffers.Table
}

func (rcv *Error) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Error) Code() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Error) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ErrorStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func ErrorAddCode(builder *flatbuffers.Builder, code uint32) { builder.PrependUint32Slot(0, code, 0) }
func ErrorAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0) }
func ErrorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
