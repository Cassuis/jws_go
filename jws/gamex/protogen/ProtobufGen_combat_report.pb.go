// Code generated by protoc-gen-go.
// source: ProtobufGen_combat_report.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type COMBAT_REPORT struct {
	// * 索引
	Index *int32 `protobuf:"varint,1,req,name=index,def=0" json:"index,omitempty"`
	// * 触发类型
	TriggerType *int32 `protobuf:"varint,2,req,def=0" json:"TriggerType,omitempty"`
	// * 对应参数
	TriggerValue *string `protobuf:"bytes,3,opt,def=" json:"TriggerValue,omitempty"`
	// * 对应参数1
	TriggerValue1 *float32 `protobuf:"fixed32,4,opt,def=0" json:"TriggerValue1,omitempty"`
	// * 报文文本IDS
	IDS *string `protobuf:"bytes,5,opt,def=" json:"IDS,omitempty"`
	// * 延迟秒
	Delay            *int32 `protobuf:"varint,6,req,def=0" json:"Delay,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *COMBAT_REPORT) Reset()         { *m = COMBAT_REPORT{} }
func (m *COMBAT_REPORT) String() string { return proto.CompactTextString(m) }
func (*COMBAT_REPORT) ProtoMessage()    {}

const Default_COMBAT_REPORT_Index int32 = 0
const Default_COMBAT_REPORT_TriggerType int32 = 0
const Default_COMBAT_REPORT_TriggerValue1 float32 = 0
const Default_COMBAT_REPORT_Delay int32 = 0

func (m *COMBAT_REPORT) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_COMBAT_REPORT_Index
}

func (m *COMBAT_REPORT) GetTriggerType() int32 {
	if m != nil && m.TriggerType != nil {
		return *m.TriggerType
	}
	return Default_COMBAT_REPORT_TriggerType
}

func (m *COMBAT_REPORT) GetTriggerValue() string {
	if m != nil && m.TriggerValue != nil {
		return *m.TriggerValue
	}
	return ""
}

func (m *COMBAT_REPORT) GetTriggerValue1() float32 {
	if m != nil && m.TriggerValue1 != nil {
		return *m.TriggerValue1
	}
	return Default_COMBAT_REPORT_TriggerValue1
}

func (m *COMBAT_REPORT) GetIDS() string {
	if m != nil && m.IDS != nil {
		return *m.IDS
	}
	return ""
}

func (m *COMBAT_REPORT) GetDelay() int32 {
	if m != nil && m.Delay != nil {
		return *m.Delay
	}
	return Default_COMBAT_REPORT_Delay
}

type COMBAT_REPORT_ARRAY struct {
	Items            []*COMBAT_REPORT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *COMBAT_REPORT_ARRAY) Reset()         { *m = COMBAT_REPORT_ARRAY{} }
func (m *COMBAT_REPORT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*COMBAT_REPORT_ARRAY) ProtoMessage()    {}

func (m *COMBAT_REPORT_ARRAY) GetItems() []*COMBAT_REPORT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}