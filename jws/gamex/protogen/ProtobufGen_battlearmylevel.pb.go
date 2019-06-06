// Code generated by protoc-gen-go.
// source: ProtobufGen_battlearmylevel.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type BATTLEARMYLEVEL struct {
	// * 战阵将位位置
	BattleArmyLoc *uint32 `protobuf:"varint,1,req,def=0" json:"BattleArmyLoc,omitempty"`
	// * 战阵将位等级
	BattleArmyLevel *uint32 `protobuf:"varint,2,opt,def=0" json:"BattleArmyLevel,omitempty"`
	// * 当前等级属性加成
	Value *float32 `protobuf:"fixed32,3,opt,def=0" json:"Value,omitempty"`
	// * 升级将位所需的战队等级
	BattleArmyLevelLimit  *uint32                         `protobuf:"varint,4,opt,def=0" json:"BattleArmyLevelLimit,omitempty"`
	BattleArmyUp_Template []*BATTLEARMYLEVEL_BattleArmyUp `protobuf:"bytes,7,rep" json:"BattleArmyUp_Template,omitempty"`
	XXX_unrecognized      []byte                          `json:"-"`
}

func (m *BATTLEARMYLEVEL) Reset()         { *m = BATTLEARMYLEVEL{} }
func (m *BATTLEARMYLEVEL) String() string { return proto.CompactTextString(m) }
func (*BATTLEARMYLEVEL) ProtoMessage()    {}

const Default_BATTLEARMYLEVEL_BattleArmyLoc uint32 = 0
const Default_BATTLEARMYLEVEL_BattleArmyLevel uint32 = 0
const Default_BATTLEARMYLEVEL_Value float32 = 0
const Default_BATTLEARMYLEVEL_BattleArmyLevelLimit uint32 = 0

func (m *BATTLEARMYLEVEL) GetBattleArmyLoc() uint32 {
	if m != nil && m.BattleArmyLoc != nil {
		return *m.BattleArmyLoc
	}
	return Default_BATTLEARMYLEVEL_BattleArmyLoc
}

func (m *BATTLEARMYLEVEL) GetBattleArmyLevel() uint32 {
	if m != nil && m.BattleArmyLevel != nil {
		return *m.BattleArmyLevel
	}
	return Default_BATTLEARMYLEVEL_BattleArmyLevel
}

func (m *BATTLEARMYLEVEL) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_BATTLEARMYLEVEL_Value
}

func (m *BATTLEARMYLEVEL) GetBattleArmyLevelLimit() uint32 {
	if m != nil && m.BattleArmyLevelLimit != nil {
		return *m.BattleArmyLevelLimit
	}
	return Default_BATTLEARMYLEVEL_BattleArmyLevelLimit
}

func (m *BATTLEARMYLEVEL) GetBattleArmyUp_Template() []*BATTLEARMYLEVEL_BattleArmyUp {
	if m != nil {
		return m.BattleArmyUp_Template
	}
	return nil
}

type BATTLEARMYLEVEL_BattleArmyUp struct {
	// * 升级到当前等级消耗的货币
	BattleArmyCoin *string `protobuf:"bytes,1,opt,def=" json:"BattleArmyCoin,omitempty"`
	// * 消耗的货币量
	BattleArmyCost   *uint32 `protobuf:"varint,2,opt,def=0" json:"BattleArmyCost,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BATTLEARMYLEVEL_BattleArmyUp) Reset()         { *m = BATTLEARMYLEVEL_BattleArmyUp{} }
func (m *BATTLEARMYLEVEL_BattleArmyUp) String() string { return proto.CompactTextString(m) }
func (*BATTLEARMYLEVEL_BattleArmyUp) ProtoMessage()    {}

const Default_BATTLEARMYLEVEL_BattleArmyUp_BattleArmyCost uint32 = 0

func (m *BATTLEARMYLEVEL_BattleArmyUp) GetBattleArmyCoin() string {
	if m != nil && m.BattleArmyCoin != nil {
		return *m.BattleArmyCoin
	}
	return ""
}

func (m *BATTLEARMYLEVEL_BattleArmyUp) GetBattleArmyCost() uint32 {
	if m != nil && m.BattleArmyCost != nil {
		return *m.BattleArmyCost
	}
	return Default_BATTLEARMYLEVEL_BattleArmyUp_BattleArmyCost
}

type BATTLEARMYLEVEL_ARRAY struct {
	Items            []*BATTLEARMYLEVEL `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *BATTLEARMYLEVEL_ARRAY) Reset()         { *m = BATTLEARMYLEVEL_ARRAY{} }
func (m *BATTLEARMYLEVEL_ARRAY) String() string { return proto.CompactTextString(m) }
func (*BATTLEARMYLEVEL_ARRAY) ProtoMessage()    {}

func (m *BATTLEARMYLEVEL_ARRAY) GetItems() []*BATTLEARMYLEVEL {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}