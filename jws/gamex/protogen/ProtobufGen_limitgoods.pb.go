// Code generated by protoc-gen-go.
// source: ProtobufGen_limitgoods.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LIMITGOODS struct {
	// * 限时商品ID
	LimitGoodsID *uint32 `protobuf:"varint,1,req,def=0" json:"LimitGoodsID,omitempty"`
	// * 是否激活，0=关闭，1=激活
	ActivityValid *uint32 `protobuf:"varint,11,opt,def=0" json:"ActivityValid,omitempty"`
	// * 开启服务器组ID
	ServerGroupID *uint32 `protobuf:"varint,12,opt,def=0" json:"ServerGroupID,omitempty"`
	// * 商品名称
	GoodsName *string `protobuf:"bytes,2,opt,def=" json:"GoodsName,omitempty"`
	// * 是否为礼包
	GoodsType *uint32 `protobuf:"varint,3,opt,def=0" json:"GoodsType,omitempty"`
	// * 礼包图标
	GoodsIcon *string `protobuf:"bytes,10,opt,def=" json:"GoodsIcon,omitempty"`
	// * 活动时间类型
	TimeType *uint32 `protobuf:"varint,16,opt,def=0" json:"TimeType,omitempty"`
	// * 上架时间
	StartTime *string `protobuf:"bytes,4,opt,def=" json:"StartTime,omitempty"`
	// * 对应VIP等级
	VIPLevel *uint32 `protobuf:"varint,14,opt,def=0" json:"VIPLevel,omitempty"`
	// * 持续时间（单位：小时）
	Duration *uint32 `protobuf:"varint,5,opt,def=0" json:"Duration,omitempty"`
	// * 购买货币ID
	CoinItemID *string `protobuf:"bytes,6,opt,def=VI_HC" json:"CoinItemID,omitempty"`
	// * 现价
	CurrentPrice *uint32 `protobuf:"varint,7,opt,def=0" json:"CurrentPrice,omitempty"`
	// * 原价
	OriginalCost *uint32 `protobuf:"varint,8,opt,def=0" json:"OriginalCost,omitempty"`
	// * 折扣
	Discount *uint32 `protobuf:"varint,13,opt,def=0" json:"Discount,omitempty"`
	// * 限购次数
	TimesLimit       *uint32            `protobuf:"varint,15,opt,def=1" json:"TimesLimit,omitempty"`
	Fixed_Loot       []*LIMITGOODS_Loot `protobuf:"bytes,9,rep" json:"Fixed_Loot,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *LIMITGOODS) Reset()         { *m = LIMITGOODS{} }
func (m *LIMITGOODS) String() string { return proto.CompactTextString(m) }
func (*LIMITGOODS) ProtoMessage()    {}

const Default_LIMITGOODS_LimitGoodsID uint32 = 0
const Default_LIMITGOODS_ActivityValid uint32 = 0
const Default_LIMITGOODS_ServerGroupID uint32 = 0
const Default_LIMITGOODS_GoodsType uint32 = 0
const Default_LIMITGOODS_TimeType uint32 = 0
const Default_LIMITGOODS_VIPLevel uint32 = 0
const Default_LIMITGOODS_Duration uint32 = 0
const Default_LIMITGOODS_CoinItemID string = "VI_HC"
const Default_LIMITGOODS_CurrentPrice uint32 = 0
const Default_LIMITGOODS_OriginalCost uint32 = 0
const Default_LIMITGOODS_Discount uint32 = 0
const Default_LIMITGOODS_TimesLimit uint32 = 1

func (m *LIMITGOODS) GetLimitGoodsID() uint32 {
	if m != nil && m.LimitGoodsID != nil {
		return *m.LimitGoodsID
	}
	return Default_LIMITGOODS_LimitGoodsID
}

func (m *LIMITGOODS) GetActivityValid() uint32 {
	if m != nil && m.ActivityValid != nil {
		return *m.ActivityValid
	}
	return Default_LIMITGOODS_ActivityValid
}

func (m *LIMITGOODS) GetServerGroupID() uint32 {
	if m != nil && m.ServerGroupID != nil {
		return *m.ServerGroupID
	}
	return Default_LIMITGOODS_ServerGroupID
}

func (m *LIMITGOODS) GetGoodsName() string {
	if m != nil && m.GoodsName != nil {
		return *m.GoodsName
	}
	return ""
}

func (m *LIMITGOODS) GetGoodsType() uint32 {
	if m != nil && m.GoodsType != nil {
		return *m.GoodsType
	}
	return Default_LIMITGOODS_GoodsType
}

func (m *LIMITGOODS) GetGoodsIcon() string {
	if m != nil && m.GoodsIcon != nil {
		return *m.GoodsIcon
	}
	return ""
}

func (m *LIMITGOODS) GetTimeType() uint32 {
	if m != nil && m.TimeType != nil {
		return *m.TimeType
	}
	return Default_LIMITGOODS_TimeType
}

func (m *LIMITGOODS) GetStartTime() string {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return ""
}

func (m *LIMITGOODS) GetVIPLevel() uint32 {
	if m != nil && m.VIPLevel != nil {
		return *m.VIPLevel
	}
	return Default_LIMITGOODS_VIPLevel
}

func (m *LIMITGOODS) GetDuration() uint32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return Default_LIMITGOODS_Duration
}

func (m *LIMITGOODS) GetCoinItemID() string {
	if m != nil && m.CoinItemID != nil {
		return *m.CoinItemID
	}
	return Default_LIMITGOODS_CoinItemID
}

func (m *LIMITGOODS) GetCurrentPrice() uint32 {
	if m != nil && m.CurrentPrice != nil {
		return *m.CurrentPrice
	}
	return Default_LIMITGOODS_CurrentPrice
}

func (m *LIMITGOODS) GetOriginalCost() uint32 {
	if m != nil && m.OriginalCost != nil {
		return *m.OriginalCost
	}
	return Default_LIMITGOODS_OriginalCost
}

func (m *LIMITGOODS) GetDiscount() uint32 {
	if m != nil && m.Discount != nil {
		return *m.Discount
	}
	return Default_LIMITGOODS_Discount
}

func (m *LIMITGOODS) GetTimesLimit() uint32 {
	if m != nil && m.TimesLimit != nil {
		return *m.TimesLimit
	}
	return Default_LIMITGOODS_TimesLimit
}

func (m *LIMITGOODS) GetFixed_Loot() []*LIMITGOODS_Loot {
	if m != nil {
		return m.Fixed_Loot
	}
	return nil
}

type LIMITGOODS_Loot struct {
	// * 物品ID
	ItemID *string `protobuf:"bytes,1,opt,def=" json:"ItemID,omitempty"`
	// * 数量
	GoodsCount       *uint32 `protobuf:"varint,2,opt,def=0" json:"GoodsCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LIMITGOODS_Loot) Reset()         { *m = LIMITGOODS_Loot{} }
func (m *LIMITGOODS_Loot) String() string { return proto.CompactTextString(m) }
func (*LIMITGOODS_Loot) ProtoMessage()    {}

const Default_LIMITGOODS_Loot_GoodsCount uint32 = 0

func (m *LIMITGOODS_Loot) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *LIMITGOODS_Loot) GetGoodsCount() uint32 {
	if m != nil && m.GoodsCount != nil {
		return *m.GoodsCount
	}
	return Default_LIMITGOODS_Loot_GoodsCount
}

type LIMITGOODS_ARRAY struct {
	Items            []*LIMITGOODS `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *LIMITGOODS_ARRAY) Reset()         { *m = LIMITGOODS_ARRAY{} }
func (m *LIMITGOODS_ARRAY) String() string { return proto.CompactTextString(m) }
func (*LIMITGOODS_ARRAY) ProtoMessage()    {}

func (m *LIMITGOODS_ARRAY) GetItems() []*LIMITGOODS {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
