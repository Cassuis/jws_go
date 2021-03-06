// Code generated by protoc-gen-go.
// source: ProtobufGen_euiptrickdetail.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EUIPTRICKDETAIL struct {
	// * 技能ID
	TrickID *string `protobuf:"bytes,1,req,def=" json:"TrickID,omitempty"`
	// * 属性ID
	Property *string `protobuf:"bytes,2,opt,def=" json:"Property,omitempty"`
	// * 数值属性
	Value *uint32 `protobuf:"varint,3,opt,def=0" json:"Value,omitempty"`
	// * 特技名称
	TrickName *string `protobuf:"bytes,4,opt,def=" json:"TrickName,omitempty"`
	// * 在T0上的颜色
	Color0 *uint32 `protobuf:"varint,5,opt,def=0" json:"Color0,omitempty"`
	// * 在T1上的颜色
	Color1 *uint32 `protobuf:"varint,6,opt,def=0" json:"Color1,omitempty"`
	// * 在T2上的颜色
	Color2 *uint32 `protobuf:"varint,7,opt,def=0" json:"Color2,omitempty"`
	// * 在T3上的颜色
	Color3 *uint32 `protobuf:"varint,8,opt,def=0" json:"Color3,omitempty"`
	// * 在T4上的颜色
	Color4 *uint32 `protobuf:"varint,9,opt,def=0" json:"Color4,omitempty"`
	// * 在T5上的颜色
	Color5 *uint32 `protobuf:"varint,10,opt,def=0" json:"Color5,omitempty"`
	// * 在T6上的颜色
	Color6 *uint32 `protobuf:"varint,11,opt,def=0" json:"Color6,omitempty"`
	// * 在T7上的颜色
	Color7 *uint32 `protobuf:"varint,12,opt,def=0" json:"Color7,omitempty"`
	// * 在T8上的颜色
	Color8 *uint32 `protobuf:"varint,13,opt,def=0" json:"Color8,omitempty"`
	// * 在T9上的颜色
	Color9 *uint32 `protobuf:"varint,14,opt,def=0" json:"Color9,omitempty"`
	// * 在T10上的颜色
	Color10 *uint32 `protobuf:"varint,15,opt,def=0" json:"Color10,omitempty"`
	// * 特技类型
	TrickType *uint32 `protobuf:"varint,16,opt,def=0" json:"TrickType,omitempty"`
	// * 特技战力
	TrickGS *uint32 `protobuf:"varint,17,opt,def=0" json:"TrickGS,omitempty"`
	// * 属性人为评分
	ValueScore       *float32 `protobuf:"fixed32,18,opt,def=0" json:"ValueScore,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EUIPTRICKDETAIL) Reset()         { *m = EUIPTRICKDETAIL{} }
func (m *EUIPTRICKDETAIL) String() string { return proto.CompactTextString(m) }
func (*EUIPTRICKDETAIL) ProtoMessage()    {}

const Default_EUIPTRICKDETAIL_Value uint32 = 0
const Default_EUIPTRICKDETAIL_Color0 uint32 = 0
const Default_EUIPTRICKDETAIL_Color1 uint32 = 0
const Default_EUIPTRICKDETAIL_Color2 uint32 = 0
const Default_EUIPTRICKDETAIL_Color3 uint32 = 0
const Default_EUIPTRICKDETAIL_Color4 uint32 = 0
const Default_EUIPTRICKDETAIL_Color5 uint32 = 0
const Default_EUIPTRICKDETAIL_Color6 uint32 = 0
const Default_EUIPTRICKDETAIL_Color7 uint32 = 0
const Default_EUIPTRICKDETAIL_Color8 uint32 = 0
const Default_EUIPTRICKDETAIL_Color9 uint32 = 0
const Default_EUIPTRICKDETAIL_Color10 uint32 = 0
const Default_EUIPTRICKDETAIL_TrickType uint32 = 0
const Default_EUIPTRICKDETAIL_TrickGS uint32 = 0
const Default_EUIPTRICKDETAIL_ValueScore float32 = 0

func (m *EUIPTRICKDETAIL) GetTrickID() string {
	if m != nil && m.TrickID != nil {
		return *m.TrickID
	}
	return ""
}

func (m *EUIPTRICKDETAIL) GetProperty() string {
	if m != nil && m.Property != nil {
		return *m.Property
	}
	return ""
}

func (m *EUIPTRICKDETAIL) GetValue() uint32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_EUIPTRICKDETAIL_Value
}

func (m *EUIPTRICKDETAIL) GetTrickName() string {
	if m != nil && m.TrickName != nil {
		return *m.TrickName
	}
	return ""
}

func (m *EUIPTRICKDETAIL) GetColor0() uint32 {
	if m != nil && m.Color0 != nil {
		return *m.Color0
	}
	return Default_EUIPTRICKDETAIL_Color0
}

func (m *EUIPTRICKDETAIL) GetColor1() uint32 {
	if m != nil && m.Color1 != nil {
		return *m.Color1
	}
	return Default_EUIPTRICKDETAIL_Color1
}

func (m *EUIPTRICKDETAIL) GetColor2() uint32 {
	if m != nil && m.Color2 != nil {
		return *m.Color2
	}
	return Default_EUIPTRICKDETAIL_Color2
}

func (m *EUIPTRICKDETAIL) GetColor3() uint32 {
	if m != nil && m.Color3 != nil {
		return *m.Color3
	}
	return Default_EUIPTRICKDETAIL_Color3
}

func (m *EUIPTRICKDETAIL) GetColor4() uint32 {
	if m != nil && m.Color4 != nil {
		return *m.Color4
	}
	return Default_EUIPTRICKDETAIL_Color4
}

func (m *EUIPTRICKDETAIL) GetColor5() uint32 {
	if m != nil && m.Color5 != nil {
		return *m.Color5
	}
	return Default_EUIPTRICKDETAIL_Color5
}

func (m *EUIPTRICKDETAIL) GetColor6() uint32 {
	if m != nil && m.Color6 != nil {
		return *m.Color6
	}
	return Default_EUIPTRICKDETAIL_Color6
}

func (m *EUIPTRICKDETAIL) GetColor7() uint32 {
	if m != nil && m.Color7 != nil {
		return *m.Color7
	}
	return Default_EUIPTRICKDETAIL_Color7
}

func (m *EUIPTRICKDETAIL) GetColor8() uint32 {
	if m != nil && m.Color8 != nil {
		return *m.Color8
	}
	return Default_EUIPTRICKDETAIL_Color8
}

func (m *EUIPTRICKDETAIL) GetColor9() uint32 {
	if m != nil && m.Color9 != nil {
		return *m.Color9
	}
	return Default_EUIPTRICKDETAIL_Color9
}

func (m *EUIPTRICKDETAIL) GetColor10() uint32 {
	if m != nil && m.Color10 != nil {
		return *m.Color10
	}
	return Default_EUIPTRICKDETAIL_Color10
}

func (m *EUIPTRICKDETAIL) GetTrickType() uint32 {
	if m != nil && m.TrickType != nil {
		return *m.TrickType
	}
	return Default_EUIPTRICKDETAIL_TrickType
}

func (m *EUIPTRICKDETAIL) GetTrickGS() uint32 {
	if m != nil && m.TrickGS != nil {
		return *m.TrickGS
	}
	return Default_EUIPTRICKDETAIL_TrickGS
}

func (m *EUIPTRICKDETAIL) GetValueScore() float32 {
	if m != nil && m.ValueScore != nil {
		return *m.ValueScore
	}
	return Default_EUIPTRICKDETAIL_ValueScore
}

type EUIPTRICKDETAIL_ARRAY struct {
	Items            []*EUIPTRICKDETAIL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *EUIPTRICKDETAIL_ARRAY) Reset()         { *m = EUIPTRICKDETAIL_ARRAY{} }
func (m *EUIPTRICKDETAIL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*EUIPTRICKDETAIL_ARRAY) ProtoMessage()    {}

func (m *EUIPTRICKDETAIL_ARRAY) GetItems() []*EUIPTRICKDETAIL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
