// Code generated by protoc-gen-go.
// source: ProtobufGen_gvgguildgift.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GVGGUILDGIFT struct {
	// * ID
	CityID     *uint32                  `protobuf:"varint,1,req,def=0" json:"CityID,omitempty"`
	Loot_Table []*GVGGUILDGIFT_LootRule `protobuf:"bytes,2,rep" json:"Loot_Table,omitempty"`
	// * 奖励1
	GuildBagID *string `protobuf:"bytes,3,opt,def=" json:"GuildBagID,omitempty"`
	// * 数量1
	GuildBagNum *uint32 `protobuf:"varint,4,opt,def=0" json:"GuildBagNum,omitempty"`
	// * 第二名的数量
	SecondGuildBagNum *uint32 `protobuf:"varint,5,opt,def=0" json:"SecondGuildBagNum,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *GVGGUILDGIFT) Reset()         { *m = GVGGUILDGIFT{} }
func (m *GVGGUILDGIFT) String() string { return proto.CompactTextString(m) }
func (*GVGGUILDGIFT) ProtoMessage()    {}

const Default_GVGGUILDGIFT_CityID uint32 = 0
const Default_GVGGUILDGIFT_GuildBagNum uint32 = 0
const Default_GVGGUILDGIFT_SecondGuildBagNum uint32 = 0

func (m *GVGGUILDGIFT) GetCityID() uint32 {
	if m != nil && m.CityID != nil {
		return *m.CityID
	}
	return Default_GVGGUILDGIFT_CityID
}

func (m *GVGGUILDGIFT) GetLoot_Table() []*GVGGUILDGIFT_LootRule {
	if m != nil {
		return m.Loot_Table
	}
	return nil
}

func (m *GVGGUILDGIFT) GetGuildBagID() string {
	if m != nil && m.GuildBagID != nil {
		return *m.GuildBagID
	}
	return ""
}

func (m *GVGGUILDGIFT) GetGuildBagNum() uint32 {
	if m != nil && m.GuildBagNum != nil {
		return *m.GuildBagNum
	}
	return Default_GVGGUILDGIFT_GuildBagNum
}

func (m *GVGGUILDGIFT) GetSecondGuildBagNum() uint32 {
	if m != nil && m.SecondGuildBagNum != nil {
		return *m.SecondGuildBagNum
	}
	return Default_GVGGUILDGIFT_SecondGuildBagNum
}

type GVGGUILDGIFT_LootRule struct {
	// * 奖励1
	GuildItemID *string `protobuf:"bytes,1,opt,def=" json:"GuildItemID,omitempty"`
	// * 数量1
	GuildItemNum     *uint32 `protobuf:"varint,2,opt,def=0" json:"GuildItemNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GVGGUILDGIFT_LootRule) Reset()         { *m = GVGGUILDGIFT_LootRule{} }
func (m *GVGGUILDGIFT_LootRule) String() string { return proto.CompactTextString(m) }
func (*GVGGUILDGIFT_LootRule) ProtoMessage()    {}

const Default_GVGGUILDGIFT_LootRule_GuildItemNum uint32 = 0

func (m *GVGGUILDGIFT_LootRule) GetGuildItemID() string {
	if m != nil && m.GuildItemID != nil {
		return *m.GuildItemID
	}
	return ""
}

func (m *GVGGUILDGIFT_LootRule) GetGuildItemNum() uint32 {
	if m != nil && m.GuildItemNum != nil {
		return *m.GuildItemNum
	}
	return Default_GVGGUILDGIFT_LootRule_GuildItemNum
}

type GVGGUILDGIFT_ARRAY struct {
	Items            []*GVGGUILDGIFT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GVGGUILDGIFT_ARRAY) Reset()         { *m = GVGGUILDGIFT_ARRAY{} }
func (m *GVGGUILDGIFT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GVGGUILDGIFT_ARRAY) ProtoMessage()    {}

func (m *GVGGUILDGIFT_ARRAY) GetItems() []*GVGGUILDGIFT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
