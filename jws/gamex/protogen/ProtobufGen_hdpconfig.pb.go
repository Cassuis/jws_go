// Code generated by protoc-gen-go.
// source: ProtobufGen_hdpconfig.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HDPCONFIG struct {
	// * 消耗扫荡道具个数
	SweepCost *uint32 `protobuf:"varint,1,req,def=0" json:"SweepCost,omitempty"`
	// * 扫荡解锁等级
	SweepUnlockLevel *uint32 `protobuf:"varint,2,req,def=0" json:"SweepUnlockLevel,omitempty"`
	// * 多少分以上可扫荡（含）
	SweepNeedPoint *uint32 `protobuf:"varint,3,req,def=0" json:"SweepNeedPoint,omitempty"`
	// * 购买扫荡券用到的货币
	BuySweepItemUse *string `protobuf:"bytes,4,req,def=" json:"BuySweepItemUse,omitempty"`
	// * 购买扫荡券的单价
	BuySweepItemPrice *uint32 `protobuf:"varint,5,req,def=0" json:"BuySweepItemPrice,omitempty"`
	// * 士验证用除数
	ShiDivisor       *uint32 `protobuf:"varint,6,req,def=0" json:"ShiDivisor,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HDPCONFIG) Reset()         { *m = HDPCONFIG{} }
func (m *HDPCONFIG) String() string { return proto.CompactTextString(m) }
func (*HDPCONFIG) ProtoMessage()    {}

const Default_HDPCONFIG_SweepCost uint32 = 0
const Default_HDPCONFIG_SweepUnlockLevel uint32 = 0
const Default_HDPCONFIG_SweepNeedPoint uint32 = 0
const Default_HDPCONFIG_BuySweepItemPrice uint32 = 0
const Default_HDPCONFIG_ShiDivisor uint32 = 0

func (m *HDPCONFIG) GetSweepCost() uint32 {
	if m != nil && m.SweepCost != nil {
		return *m.SweepCost
	}
	return Default_HDPCONFIG_SweepCost
}

func (m *HDPCONFIG) GetSweepUnlockLevel() uint32 {
	if m != nil && m.SweepUnlockLevel != nil {
		return *m.SweepUnlockLevel
	}
	return Default_HDPCONFIG_SweepUnlockLevel
}

func (m *HDPCONFIG) GetSweepNeedPoint() uint32 {
	if m != nil && m.SweepNeedPoint != nil {
		return *m.SweepNeedPoint
	}
	return Default_HDPCONFIG_SweepNeedPoint
}

func (m *HDPCONFIG) GetBuySweepItemUse() string {
	if m != nil && m.BuySweepItemUse != nil {
		return *m.BuySweepItemUse
	}
	return ""
}

func (m *HDPCONFIG) GetBuySweepItemPrice() uint32 {
	if m != nil && m.BuySweepItemPrice != nil {
		return *m.BuySweepItemPrice
	}
	return Default_HDPCONFIG_BuySweepItemPrice
}

func (m *HDPCONFIG) GetShiDivisor() uint32 {
	if m != nil && m.ShiDivisor != nil {
		return *m.ShiDivisor
	}
	return Default_HDPCONFIG_ShiDivisor
}

type HDPCONFIG_ARRAY struct {
	Items            []*HDPCONFIG `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HDPCONFIG_ARRAY) Reset()         { *m = HDPCONFIG_ARRAY{} }
func (m *HDPCONFIG_ARRAY) String() string { return proto.CompactTextString(m) }
func (*HDPCONFIG_ARRAY) ProtoMessage()    {}

func (m *HDPCONFIG_ARRAY) GetItems() []*HDPCONFIG {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}