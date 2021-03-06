// Code generated by protoc-gen-go.
// source: ProtobufGen_geconfig.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GECONFIG struct {
	// * 兵临城下开启时间
	GEStartTime *string `protobuf:"bytes,15,req,def=" json:"GEStartTime,omitempty"`
	// * 活动等待时间（min）
	GEWaitTime *uint32 `protobuf:"varint,1,req,def=0" json:"GEWaitTime,omitempty"`
	// * 活动开启时间（min）
	GEPlayTime *uint32 `protobuf:"varint,2,req,def=0" json:"GEPlayTime,omitempty"`
	// * 挑战时间限制（min）
	FightTimeLimit *uint32 `protobuf:"varint,3,req,def=0" json:"FightTimeLimit,omitempty"`
	// * 每日次数
	DailyFightTime *uint32 `protobuf:"varint,4,req,def=0" json:"DailyFightTime,omitempty"`
	// * 公会鼓舞所需货币
	GEBuffCoin *string `protobuf:"bytes,5,req,def=" json:"GEBuffCoin,omitempty"`
	// * 公会鼓舞价格
	GEBuffCount *uint32 `protobuf:"varint,6,req,def=0" json:"GEBuffCount,omitempty"`
	// * 擂鼓手掉落加成倍率
	EncouragerLootAddition *float32 `protobuf:"fixed32,7,req,def=0" json:"EncouragerLootAddition,omitempty"`
	// * 擂鼓手结算奖励加成
	EncouragerGiftAddition *float32 `protobuf:"fixed32,8,req,def=0" json:"EncouragerGiftAddition,omitempty"`
	// * 领取结算奖励小加成百分比
	GiftLittleAddition *float32 `protobuf:"fixed32,9,req,def=0" json:"GiftLittleAddition,omitempty"`
	// * 领取结算奖励小加成货币
	GiftLittleCoin *string `protobuf:"bytes,10,req,def=" json:"GiftLittleCoin,omitempty"`
	// * 领取结算奖励小加成价格
	GiftLittleCoinCount *uint32 `protobuf:"varint,11,req,def=0" json:"GiftLittleCoinCount,omitempty"`
	// * 领取结算奖励大加成百分比
	GiftHugeAddition *float32 `protobuf:"fixed32,12,req,def=0" json:"GiftHugeAddition,omitempty"`
	// * 领取结算奖励大加成货币
	GiftHugeCoin *string `protobuf:"bytes,13,req,def=" json:"GiftHugeCoin,omitempty"`
	// * 领取结算奖励大加成价格
	GiftHugeCoinCount *uint32 `protobuf:"varint,14,req,def=0" json:"GiftHugeCoinCount,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *GECONFIG) Reset()         { *m = GECONFIG{} }
func (m *GECONFIG) String() string { return proto.CompactTextString(m) }
func (*GECONFIG) ProtoMessage()    {}

const Default_GECONFIG_GEWaitTime uint32 = 0
const Default_GECONFIG_GEPlayTime uint32 = 0
const Default_GECONFIG_FightTimeLimit uint32 = 0
const Default_GECONFIG_DailyFightTime uint32 = 0
const Default_GECONFIG_GEBuffCount uint32 = 0
const Default_GECONFIG_EncouragerLootAddition float32 = 0
const Default_GECONFIG_EncouragerGiftAddition float32 = 0
const Default_GECONFIG_GiftLittleAddition float32 = 0
const Default_GECONFIG_GiftLittleCoinCount uint32 = 0
const Default_GECONFIG_GiftHugeAddition float32 = 0
const Default_GECONFIG_GiftHugeCoinCount uint32 = 0

func (m *GECONFIG) GetGEStartTime() string {
	if m != nil && m.GEStartTime != nil {
		return *m.GEStartTime
	}
	return ""
}

func (m *GECONFIG) GetGEWaitTime() uint32 {
	if m != nil && m.GEWaitTime != nil {
		return *m.GEWaitTime
	}
	return Default_GECONFIG_GEWaitTime
}

func (m *GECONFIG) GetGEPlayTime() uint32 {
	if m != nil && m.GEPlayTime != nil {
		return *m.GEPlayTime
	}
	return Default_GECONFIG_GEPlayTime
}

func (m *GECONFIG) GetFightTimeLimit() uint32 {
	if m != nil && m.FightTimeLimit != nil {
		return *m.FightTimeLimit
	}
	return Default_GECONFIG_FightTimeLimit
}

func (m *GECONFIG) GetDailyFightTime() uint32 {
	if m != nil && m.DailyFightTime != nil {
		return *m.DailyFightTime
	}
	return Default_GECONFIG_DailyFightTime
}

func (m *GECONFIG) GetGEBuffCoin() string {
	if m != nil && m.GEBuffCoin != nil {
		return *m.GEBuffCoin
	}
	return ""
}

func (m *GECONFIG) GetGEBuffCount() uint32 {
	if m != nil && m.GEBuffCount != nil {
		return *m.GEBuffCount
	}
	return Default_GECONFIG_GEBuffCount
}

func (m *GECONFIG) GetEncouragerLootAddition() float32 {
	if m != nil && m.EncouragerLootAddition != nil {
		return *m.EncouragerLootAddition
	}
	return Default_GECONFIG_EncouragerLootAddition
}

func (m *GECONFIG) GetEncouragerGiftAddition() float32 {
	if m != nil && m.EncouragerGiftAddition != nil {
		return *m.EncouragerGiftAddition
	}
	return Default_GECONFIG_EncouragerGiftAddition
}

func (m *GECONFIG) GetGiftLittleAddition() float32 {
	if m != nil && m.GiftLittleAddition != nil {
		return *m.GiftLittleAddition
	}
	return Default_GECONFIG_GiftLittleAddition
}

func (m *GECONFIG) GetGiftLittleCoin() string {
	if m != nil && m.GiftLittleCoin != nil {
		return *m.GiftLittleCoin
	}
	return ""
}

func (m *GECONFIG) GetGiftLittleCoinCount() uint32 {
	if m != nil && m.GiftLittleCoinCount != nil {
		return *m.GiftLittleCoinCount
	}
	return Default_GECONFIG_GiftLittleCoinCount
}

func (m *GECONFIG) GetGiftHugeAddition() float32 {
	if m != nil && m.GiftHugeAddition != nil {
		return *m.GiftHugeAddition
	}
	return Default_GECONFIG_GiftHugeAddition
}

func (m *GECONFIG) GetGiftHugeCoin() string {
	if m != nil && m.GiftHugeCoin != nil {
		return *m.GiftHugeCoin
	}
	return ""
}

func (m *GECONFIG) GetGiftHugeCoinCount() uint32 {
	if m != nil && m.GiftHugeCoinCount != nil {
		return *m.GiftHugeCoinCount
	}
	return Default_GECONFIG_GiftHugeCoinCount
}

type GECONFIG_ARRAY struct {
	Items            []*GECONFIG `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GECONFIG_ARRAY) Reset()         { *m = GECONFIG_ARRAY{} }
func (m *GECONFIG_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GECONFIG_ARRAY) ProtoMessage()    {}

func (m *GECONFIG_ARRAY) GetItems() []*GECONFIG {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
