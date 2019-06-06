// Code generated by protoc-gen-go.
// source: ProtobufGen_destinyconfig.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DESTINYCONFIG struct {
	// * 神兽升级小暴击倍数
	DGLittleBonus *uint32 `protobuf:"varint,1,req,def=5" json:"DGLittleBonus,omitempty"`
	// * 普通/连续培养的默认次数
	TrainTimes *uint32 `protobuf:"varint,2,req,def=0" json:"TrainTimes,omitempty"`
	// * 普通/连续培养的默认次数
	TrainContinuityTimes *uint32 `protobuf:"varint,3,req,def=0" json:"TrainContinuityTimes,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *DESTINYCONFIG) Reset()         { *m = DESTINYCONFIG{} }
func (m *DESTINYCONFIG) String() string { return proto.CompactTextString(m) }
func (*DESTINYCONFIG) ProtoMessage()    {}

const Default_DESTINYCONFIG_DGLittleBonus uint32 = 5
const Default_DESTINYCONFIG_TrainTimes uint32 = 0
const Default_DESTINYCONFIG_TrainContinuityTimes uint32 = 0

func (m *DESTINYCONFIG) GetDGLittleBonus() uint32 {
	if m != nil && m.DGLittleBonus != nil {
		return *m.DGLittleBonus
	}
	return Default_DESTINYCONFIG_DGLittleBonus
}

func (m *DESTINYCONFIG) GetTrainTimes() uint32 {
	if m != nil && m.TrainTimes != nil {
		return *m.TrainTimes
	}
	return Default_DESTINYCONFIG_TrainTimes
}

func (m *DESTINYCONFIG) GetTrainContinuityTimes() uint32 {
	if m != nil && m.TrainContinuityTimes != nil {
		return *m.TrainContinuityTimes
	}
	return Default_DESTINYCONFIG_TrainContinuityTimes
}

type DESTINYCONFIG_ARRAY struct {
	Items            []*DESTINYCONFIG `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DESTINYCONFIG_ARRAY) Reset()         { *m = DESTINYCONFIG_ARRAY{} }
func (m *DESTINYCONFIG_ARRAY) String() string { return proto.CompactTextString(m) }
func (*DESTINYCONFIG_ARRAY) ProtoMessage()    {}

func (m *DESTINYCONFIG_ARRAY) GetItems() []*DESTINYCONFIG {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}