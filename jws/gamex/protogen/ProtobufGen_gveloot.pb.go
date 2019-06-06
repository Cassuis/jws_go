// Code generated by protoc-gen-go.
// source: ProtobufGen_gveloot.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GVELOOT struct {
	// * 关卡ID
	EGLevelID        *string             `protobuf:"bytes,1,req,def=" json:"EGLevelID,omitempty"`
	Loots            []*GVELOOT_LootRule `protobuf:"bytes,2,rep" json:"Loots,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GVELOOT) Reset()         { *m = GVELOOT{} }
func (m *GVELOOT) String() string { return proto.CompactTextString(m) }
func (*GVELOOT) ProtoMessage()    {}

func (m *GVELOOT) GetEGLevelID() string {
	if m != nil && m.EGLevelID != nil {
		return *m.EGLevelID
	}
	return ""
}

func (m *GVELOOT) GetLoots() []*GVELOOT_LootRule {
	if m != nil {
		return m.Loots
	}
	return nil
}

type GVELOOT_LootRule struct {
	// * 展示物品ID
	ShowItem *string `protobuf:"bytes,1,opt,def=" json:"ShowItem,omitempty"`
	// * 展示用数字
	ShowNum *uint32 `protobuf:"varint,4,opt,def=0" json:"ShowNum,omitempty"`
	// * 掉落模板
	LootTemplate *string `protobuf:"bytes,2,opt,def=" json:"LootTemplate,omitempty"`
	// * 模板掉落次数
	LootTimes        *uint32 `protobuf:"varint,3,opt,def=0" json:"LootTimes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GVELOOT_LootRule) Reset()         { *m = GVELOOT_LootRule{} }
func (m *GVELOOT_LootRule) String() string { return proto.CompactTextString(m) }
func (*GVELOOT_LootRule) ProtoMessage()    {}

const Default_GVELOOT_LootRule_ShowNum uint32 = 0
const Default_GVELOOT_LootRule_LootTimes uint32 = 0

func (m *GVELOOT_LootRule) GetShowItem() string {
	if m != nil && m.ShowItem != nil {
		return *m.ShowItem
	}
	return ""
}

func (m *GVELOOT_LootRule) GetShowNum() uint32 {
	if m != nil && m.ShowNum != nil {
		return *m.ShowNum
	}
	return Default_GVELOOT_LootRule_ShowNum
}

func (m *GVELOOT_LootRule) GetLootTemplate() string {
	if m != nil && m.LootTemplate != nil {
		return *m.LootTemplate
	}
	return ""
}

func (m *GVELOOT_LootRule) GetLootTimes() uint32 {
	if m != nil && m.LootTimes != nil {
		return *m.LootTimes
	}
	return Default_GVELOOT_LootRule_LootTimes
}

type GVELOOT_ARRAY struct {
	Items            []*GVELOOT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *GVELOOT_ARRAY) Reset()         { *m = GVELOOT_ARRAY{} }
func (m *GVELOOT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GVELOOT_ARRAY) ProtoMessage()    {}

func (m *GVELOOT_ARRAY) GetItems() []*GVELOOT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}