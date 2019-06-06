// Code generated by protoc-gen-go.
// source: ProtobufGen_guildboss.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GUILDBOSS struct {
	// * 难度等级
	DifficultyLevel *uint32 `protobuf:"varint,1,req,def=0" json:"DifficultyLevel,omitempty"`
	// * 公会等级需求
	GuildLvReqirement *uint32 `protobuf:"varint,2,req,def=0" json:"GuildLvReqirement,omitempty"`
	// * 难度名字
	LevelNameIDS *string `protobuf:"bytes,6,opt,def=" json:"LevelNameIDS,omitempty"`
	// * 每20%进度掉落（击杀那次不算）
	PartDrop *string `protobuf:"bytes,3,opt,def=" json:"PartDrop,omitempty"`
	// * 大boss大身像
	BigBossPic       *string               `protobuf:"bytes,7,opt,def=" json:"BigBossPic,omitempty"`
	BossData_Table   []*GUILDBOSS_BossData `protobuf:"bytes,4,rep" json:"BossData_Table,omitempty"`
	Loot_Table       []*GUILDBOSS_LootRule `protobuf:"bytes,5,rep" json:"Loot_Table,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *GUILDBOSS) Reset()         { *m = GUILDBOSS{} }
func (m *GUILDBOSS) String() string { return proto.CompactTextString(m) }
func (*GUILDBOSS) ProtoMessage()    {}

const Default_GUILDBOSS_DifficultyLevel uint32 = 0
const Default_GUILDBOSS_GuildLvReqirement uint32 = 0

func (m *GUILDBOSS) GetDifficultyLevel() uint32 {
	if m != nil && m.DifficultyLevel != nil {
		return *m.DifficultyLevel
	}
	return Default_GUILDBOSS_DifficultyLevel
}

func (m *GUILDBOSS) GetGuildLvReqirement() uint32 {
	if m != nil && m.GuildLvReqirement != nil {
		return *m.GuildLvReqirement
	}
	return Default_GUILDBOSS_GuildLvReqirement
}

func (m *GUILDBOSS) GetLevelNameIDS() string {
	if m != nil && m.LevelNameIDS != nil {
		return *m.LevelNameIDS
	}
	return ""
}

func (m *GUILDBOSS) GetPartDrop() string {
	if m != nil && m.PartDrop != nil {
		return *m.PartDrop
	}
	return ""
}

func (m *GUILDBOSS) GetBigBossPic() string {
	if m != nil && m.BigBossPic != nil {
		return *m.BigBossPic
	}
	return ""
}

func (m *GUILDBOSS) GetBossData_Table() []*GUILDBOSS_BossData {
	if m != nil {
		return m.BossData_Table
	}
	return nil
}

func (m *GUILDBOSS) GetLoot_Table() []*GUILDBOSS_LootRule {
	if m != nil {
		return m.Loot_Table
	}
	return nil
}

type GUILDBOSS_BossData struct {
	// * boss组ID
	BossGroupID *string `protobuf:"bytes,7,req,def=" json:"BossGroupID,omitempty"`
	// * 关卡名字
	GuildBossName *string `protobuf:"bytes,2,opt,def=" json:"GuildBossName,omitempty"`
	// * 关卡缩略名
	ShortName *string `protobuf:"bytes,12,opt,def=" json:"ShortName,omitempty"`
	// * 关卡描述
	GuildBossDesc *string `protobuf:"bytes,8,opt,def=" json:"GuildBossDesc,omitempty"`
	// * 作战需求描述
	GuildBossNeedIDS *string `protobuf:"bytes,9,opt,def=" json:"GuildBossNeedIDS,omitempty"`
	// * 挑战名额（0无限制)
	ChallengeNum *uint32 `protobuf:"varint,10,opt,def=0" json:"ChallengeNum,omitempty"`
	// * 可进入的公会职位
	PositionLevel *string `protobuf:"bytes,13,opt,def=" json:"PositionLevel,omitempty"`
	// * 限制职位描述
	PositionLevelDes *string `protobuf:"bytes,14,opt,def=" json:"PositionLevelDes,omitempty"`
	// * 挑战图片
	GuildBossPic *string `protobuf:"bytes,11,opt,def=" json:"GuildBossPic,omitempty"`
	// * 掉落ID
	DropID *string `protobuf:"bytes,3,opt,def=" json:"DropID,omitempty"`
	// * 对应的关卡
	LevelDemo *string `protobuf:"bytes,4,opt,def=" json:"LevelDemo,omitempty"`
	// * 个人奖励
	SelfLoot *string `protobuf:"bytes,5,opt,def=" json:"SelfLoot,omitempty"`
	// * 奖励的公会经验点
	SelfLootB *string `protobuf:"bytes,15,opt,def=" json:"SelfLootB,omitempty"`
	// * 公会经验数量
	SelfLootBNum *uint32 `protobuf:"varint,16,opt,def=0" json:"SelfLootBNum,omitempty"`
	// * 奖励的公会科技点
	SelfLootC *string `protobuf:"bytes,17,opt,def=" json:"SelfLootC,omitempty"`
	// * 公会科技点数量
	SelfLootCNum *uint32 `protobuf:"varint,18,opt,def=0" json:"SelfLootCNum,omitempty"`
	// * 军团币ID
	SelfLootD *string `protobuf:"bytes,19,opt,def=" json:"SelfLootD,omitempty"`
	// * 军团币数量
	SelfLootDNum *uint32 `protobuf:"varint,20,opt,def=0" json:"SelfLootDNum,omitempty"`
	// * 最大金币量
	GoldMax          *uint32 `protobuf:"varint,6,opt,def=0" json:"GoldMax,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GUILDBOSS_BossData) Reset()         { *m = GUILDBOSS_BossData{} }
func (m *GUILDBOSS_BossData) String() string { return proto.CompactTextString(m) }
func (*GUILDBOSS_BossData) ProtoMessage()    {}

const Default_GUILDBOSS_BossData_ChallengeNum uint32 = 0
const Default_GUILDBOSS_BossData_SelfLootBNum uint32 = 0
const Default_GUILDBOSS_BossData_SelfLootCNum uint32 = 0
const Default_GUILDBOSS_BossData_SelfLootDNum uint32 = 0
const Default_GUILDBOSS_BossData_GoldMax uint32 = 0

func (m *GUILDBOSS_BossData) GetBossGroupID() string {
	if m != nil && m.BossGroupID != nil {
		return *m.BossGroupID
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetGuildBossName() string {
	if m != nil && m.GuildBossName != nil {
		return *m.GuildBossName
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetShortName() string {
	if m != nil && m.ShortName != nil {
		return *m.ShortName
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetGuildBossDesc() string {
	if m != nil && m.GuildBossDesc != nil {
		return *m.GuildBossDesc
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetGuildBossNeedIDS() string {
	if m != nil && m.GuildBossNeedIDS != nil {
		return *m.GuildBossNeedIDS
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetChallengeNum() uint32 {
	if m != nil && m.ChallengeNum != nil {
		return *m.ChallengeNum
	}
	return Default_GUILDBOSS_BossData_ChallengeNum
}

func (m *GUILDBOSS_BossData) GetPositionLevel() string {
	if m != nil && m.PositionLevel != nil {
		return *m.PositionLevel
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetPositionLevelDes() string {
	if m != nil && m.PositionLevelDes != nil {
		return *m.PositionLevelDes
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetGuildBossPic() string {
	if m != nil && m.GuildBossPic != nil {
		return *m.GuildBossPic
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetDropID() string {
	if m != nil && m.DropID != nil {
		return *m.DropID
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetLevelDemo() string {
	if m != nil && m.LevelDemo != nil {
		return *m.LevelDemo
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetSelfLoot() string {
	if m != nil && m.SelfLoot != nil {
		return *m.SelfLoot
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetSelfLootB() string {
	if m != nil && m.SelfLootB != nil {
		return *m.SelfLootB
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetSelfLootBNum() uint32 {
	if m != nil && m.SelfLootBNum != nil {
		return *m.SelfLootBNum
	}
	return Default_GUILDBOSS_BossData_SelfLootBNum
}

func (m *GUILDBOSS_BossData) GetSelfLootC() string {
	if m != nil && m.SelfLootC != nil {
		return *m.SelfLootC
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetSelfLootCNum() uint32 {
	if m != nil && m.SelfLootCNum != nil {
		return *m.SelfLootCNum
	}
	return Default_GUILDBOSS_BossData_SelfLootCNum
}

func (m *GUILDBOSS_BossData) GetSelfLootD() string {
	if m != nil && m.SelfLootD != nil {
		return *m.SelfLootD
	}
	return ""
}

func (m *GUILDBOSS_BossData) GetSelfLootDNum() uint32 {
	if m != nil && m.SelfLootDNum != nil {
		return *m.SelfLootDNum
	}
	return Default_GUILDBOSS_BossData_SelfLootDNum
}

func (m *GUILDBOSS_BossData) GetGoldMax() uint32 {
	if m != nil && m.GoldMax != nil {
		return *m.GoldMax
	}
	return Default_GUILDBOSS_BossData_GoldMax
}

type GUILDBOSS_LootRule struct {
	// * 道具ID
	ItemID *string `protobuf:"bytes,1,opt,def=" json:"ItemID,omitempty"`
	// * 道具数量
	ItemNum          *uint32 `protobuf:"varint,2,opt,def=0" json:"ItemNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GUILDBOSS_LootRule) Reset()         { *m = GUILDBOSS_LootRule{} }
func (m *GUILDBOSS_LootRule) String() string { return proto.CompactTextString(m) }
func (*GUILDBOSS_LootRule) ProtoMessage()    {}

const Default_GUILDBOSS_LootRule_ItemNum uint32 = 0

func (m *GUILDBOSS_LootRule) GetItemID() string {
	if m != nil && m.ItemID != nil {
		return *m.ItemID
	}
	return ""
}

func (m *GUILDBOSS_LootRule) GetItemNum() uint32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return Default_GUILDBOSS_LootRule_ItemNum
}

type GUILDBOSS_ARRAY struct {
	Items            []*GUILDBOSS `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GUILDBOSS_ARRAY) Reset()         { *m = GUILDBOSS_ARRAY{} }
func (m *GUILDBOSS_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GUILDBOSS_ARRAY) ProtoMessage()    {}

func (m *GUILDBOSS_ARRAY) GetItems() []*GUILDBOSS {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}