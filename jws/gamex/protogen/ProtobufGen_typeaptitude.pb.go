// Code generated by protoc-gen-go.
// source: ProtobufGen_typeaptitude.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TYPEAPTITUDE struct {
	// * 灵宠的ID
	PetID *uint32 `protobuf:"varint,5,req,def=0" json:"PetID,omitempty"`
	// * 次数
	ExistTimes *uint32 `protobuf:"varint,1,req,def=0" json:"ExistTimes,omitempty"`
	// * 攻权重
	ATKWeight *uint32 `protobuf:"varint,2,opt,def=0" json:"ATKWeight,omitempty"`
	// * 防权重
	DEFWeight *uint32 `protobuf:"varint,3,opt,def=0" json:"DEFWeight,omitempty"`
	// * 血权重
	HPWeight         *uint32 `protobuf:"varint,4,opt,def=0" json:"HPWeight,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TYPEAPTITUDE) Reset()         { *m = TYPEAPTITUDE{} }
func (m *TYPEAPTITUDE) String() string { return proto.CompactTextString(m) }
func (*TYPEAPTITUDE) ProtoMessage()    {}

const Default_TYPEAPTITUDE_PetID uint32 = 0
const Default_TYPEAPTITUDE_ExistTimes uint32 = 0
const Default_TYPEAPTITUDE_ATKWeight uint32 = 0
const Default_TYPEAPTITUDE_DEFWeight uint32 = 0
const Default_TYPEAPTITUDE_HPWeight uint32 = 0

func (m *TYPEAPTITUDE) GetPetID() uint32 {
	if m != nil && m.PetID != nil {
		return *m.PetID
	}
	return Default_TYPEAPTITUDE_PetID
}

func (m *TYPEAPTITUDE) GetExistTimes() uint32 {
	if m != nil && m.ExistTimes != nil {
		return *m.ExistTimes
	}
	return Default_TYPEAPTITUDE_ExistTimes
}

func (m *TYPEAPTITUDE) GetATKWeight() uint32 {
	if m != nil && m.ATKWeight != nil {
		return *m.ATKWeight
	}
	return Default_TYPEAPTITUDE_ATKWeight
}

func (m *TYPEAPTITUDE) GetDEFWeight() uint32 {
	if m != nil && m.DEFWeight != nil {
		return *m.DEFWeight
	}
	return Default_TYPEAPTITUDE_DEFWeight
}

func (m *TYPEAPTITUDE) GetHPWeight() uint32 {
	if m != nil && m.HPWeight != nil {
		return *m.HPWeight
	}
	return Default_TYPEAPTITUDE_HPWeight
}

type TYPEAPTITUDE_ARRAY struct {
	Items            []*TYPEAPTITUDE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *TYPEAPTITUDE_ARRAY) Reset()         { *m = TYPEAPTITUDE_ARRAY{} }
func (m *TYPEAPTITUDE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*TYPEAPTITUDE_ARRAY) ProtoMessage()    {}

func (m *TYPEAPTITUDE_ARRAY) GetItems() []*TYPEAPTITUDE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
