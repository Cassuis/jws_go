// Code generated by protoc-gen-go.
// source: ProtobufGen_packagegroup.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PACKAGEGROUP struct {
	// * 礼包组ID
	PackageGroupID      *string                        `protobuf:"bytes,1,req,def=" json:"PackageGroupID,omitempty"`
	StaticItem_Template []*PACKAGEGROUP_StaticItemRule `protobuf:"bytes,2,rep" json:"StaticItem_Template,omitempty"`
	RandomItem_Template []*PACKAGEGROUP_RandomItemRule `protobuf:"bytes,3,rep" json:"RandomItem_Template,omitempty"`
	XXX_unrecognized    []byte                         `json:"-"`
}

func (m *PACKAGEGROUP) Reset()         { *m = PACKAGEGROUP{} }
func (m *PACKAGEGROUP) String() string { return proto.CompactTextString(m) }
func (*PACKAGEGROUP) ProtoMessage()    {}

func (m *PACKAGEGROUP) GetPackageGroupID() string {
	if m != nil && m.PackageGroupID != nil {
		return *m.PackageGroupID
	}
	return ""
}

func (m *PACKAGEGROUP) GetStaticItem_Template() []*PACKAGEGROUP_StaticItemRule {
	if m != nil {
		return m.StaticItem_Template
	}
	return nil
}

func (m *PACKAGEGROUP) GetRandomItem_Template() []*PACKAGEGROUP_RandomItemRule {
	if m != nil {
		return m.RandomItem_Template
	}
	return nil
}

type PACKAGEGROUP_StaticItemRule struct {
	// * 固定奖励ID
	StaticItem *string `protobuf:"bytes,1,opt,def=" json:"StaticItem,omitempty"`
	// * 固定奖励数量
	StaticCount      *uint32 `protobuf:"varint,2,opt,def=1" json:"StaticCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PACKAGEGROUP_StaticItemRule) Reset()         { *m = PACKAGEGROUP_StaticItemRule{} }
func (m *PACKAGEGROUP_StaticItemRule) String() string { return proto.CompactTextString(m) }
func (*PACKAGEGROUP_StaticItemRule) ProtoMessage()    {}

const Default_PACKAGEGROUP_StaticItemRule_StaticCount uint32 = 1

func (m *PACKAGEGROUP_StaticItemRule) GetStaticItem() string {
	if m != nil && m.StaticItem != nil {
		return *m.StaticItem
	}
	return ""
}

func (m *PACKAGEGROUP_StaticItemRule) GetStaticCount() uint32 {
	if m != nil && m.StaticCount != nil {
		return *m.StaticCount
	}
	return Default_PACKAGEGROUP_StaticItemRule_StaticCount
}

type PACKAGEGROUP_RandomItemRule struct {
	// * 物品ID
	ShowItem *string `protobuf:"bytes,1,opt,def=" json:"ShowItem,omitempty"`
	// * 掉落组
	LootTemplate     *string `protobuf:"bytes,2,opt,def=" json:"LootTemplate,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PACKAGEGROUP_RandomItemRule) Reset()         { *m = PACKAGEGROUP_RandomItemRule{} }
func (m *PACKAGEGROUP_RandomItemRule) String() string { return proto.CompactTextString(m) }
func (*PACKAGEGROUP_RandomItemRule) ProtoMessage()    {}

func (m *PACKAGEGROUP_RandomItemRule) GetShowItem() string {
	if m != nil && m.ShowItem != nil {
		return *m.ShowItem
	}
	return ""
}

func (m *PACKAGEGROUP_RandomItemRule) GetLootTemplate() string {
	if m != nil && m.LootTemplate != nil {
		return *m.LootTemplate
	}
	return ""
}

type PACKAGEGROUP_ARRAY struct {
	Items            []*PACKAGEGROUP `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PACKAGEGROUP_ARRAY) Reset()         { *m = PACKAGEGROUP_ARRAY{} }
func (m *PACKAGEGROUP_ARRAY) String() string { return proto.CompactTextString(m) }
func (*PACKAGEGROUP_ARRAY) ProtoMessage()    {}

func (m *PACKAGEGROUP_ARRAY) GetItems() []*PACKAGEGROUP {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}