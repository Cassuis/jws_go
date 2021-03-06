// Code generated by protoc-gen-go.
// source: ProtobufGen_expeditionsweep.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EXPEDITIONSWEEP struct {
	// * 战力差值百分比（1为1%）
	GSDisparity *uint32 `protobuf:"varint,1,req,def=0" json:"GSDisparity,omitempty"`
	// * 高战力方损耗血量
	HPCost           *float32 `protobuf:"fixed32,2,req,def=0" json:"HPCost,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EXPEDITIONSWEEP) Reset()         { *m = EXPEDITIONSWEEP{} }
func (m *EXPEDITIONSWEEP) String() string { return proto.CompactTextString(m) }
func (*EXPEDITIONSWEEP) ProtoMessage()    {}

const Default_EXPEDITIONSWEEP_GSDisparity uint32 = 0
const Default_EXPEDITIONSWEEP_HPCost float32 = 0

func (m *EXPEDITIONSWEEP) GetGSDisparity() uint32 {
	if m != nil && m.GSDisparity != nil {
		return *m.GSDisparity
	}
	return Default_EXPEDITIONSWEEP_GSDisparity
}

func (m *EXPEDITIONSWEEP) GetHPCost() float32 {
	if m != nil && m.HPCost != nil {
		return *m.HPCost
	}
	return Default_EXPEDITIONSWEEP_HPCost
}

type EXPEDITIONSWEEP_ARRAY struct {
	Items            []*EXPEDITIONSWEEP `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *EXPEDITIONSWEEP_ARRAY) Reset()         { *m = EXPEDITIONSWEEP_ARRAY{} }
func (m *EXPEDITIONSWEEP_ARRAY) String() string { return proto.CompactTextString(m) }
func (*EXPEDITIONSWEEP_ARRAY) ProtoMessage()    {}

func (m *EXPEDITIONSWEEP_ARRAY) GetItems() []*EXPEDITIONSWEEP {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
