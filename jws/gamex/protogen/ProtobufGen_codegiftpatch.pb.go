// Code generated by protoc-gen-go.
// source: ProtobufGen_codegiftpatch.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CODEGIFTPATCH struct {
	// * 批号
	PatchID *uint32 `protobuf:"varint,1,req,def=0" json:"PatchID,omitempty"`
	// * 开始时间
	StartTime *string `protobuf:"bytes,2,opt,def=" json:"StartTime,omitempty"`
	// * 结束时间
	EndTime *string `protobuf:"bytes,3,opt,def=" json:"EndTime,omitempty"`
	// * 礼包标题
	GiftName         *string `protobuf:"bytes,4,opt,def=" json:"GiftName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CODEGIFTPATCH) Reset()         { *m = CODEGIFTPATCH{} }
func (m *CODEGIFTPATCH) String() string { return proto.CompactTextString(m) }
func (*CODEGIFTPATCH) ProtoMessage()    {}

const Default_CODEGIFTPATCH_PatchID uint32 = 0

func (m *CODEGIFTPATCH) GetPatchID() uint32 {
	if m != nil && m.PatchID != nil {
		return *m.PatchID
	}
	return Default_CODEGIFTPATCH_PatchID
}

func (m *CODEGIFTPATCH) GetStartTime() string {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return ""
}

func (m *CODEGIFTPATCH) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

func (m *CODEGIFTPATCH) GetGiftName() string {
	if m != nil && m.GiftName != nil {
		return *m.GiftName
	}
	return ""
}

type CODEGIFTPATCH_ARRAY struct {
	Items            []*CODEGIFTPATCH `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CODEGIFTPATCH_ARRAY) Reset()         { *m = CODEGIFTPATCH_ARRAY{} }
func (m *CODEGIFTPATCH_ARRAY) String() string { return proto.CompactTextString(m) }
func (*CODEGIFTPATCH_ARRAY) ProtoMessage()    {}

func (m *CODEGIFTPATCH_ARRAY) GetItems() []*CODEGIFTPATCH {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
