// Code generated by protoc-gen-go.
// source: ProtobufGen_playergrowthdir.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PLAYERGROWTHDIR struct {
	// * 觉醒去处
	ArousalDir *string `protobuf:"bytes,1,req,def=" json:"ArousalDir,omitempty"`
	// * 技能去处
	SkillDir *string `protobuf:"bytes,2,req,def=" json:"SkillDir,omitempty"`
	// * 强化去处
	EnhanceDir *string `protobuf:"bytes,3,req,def=" json:"EnhanceDir,omitempty"`
	// * 升星去处
	StarDir *string `protobuf:"bytes,4,req,def=" json:"StarDir,omitempty"`
	// * 宝箱去处
	RareDir          *string `protobuf:"bytes,5,req,def=" json:"RareDir,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PLAYERGROWTHDIR) Reset()         { *m = PLAYERGROWTHDIR{} }
func (m *PLAYERGROWTHDIR) String() string { return proto.CompactTextString(m) }
func (*PLAYERGROWTHDIR) ProtoMessage()    {}

func (m *PLAYERGROWTHDIR) GetArousalDir() string {
	if m != nil && m.ArousalDir != nil {
		return *m.ArousalDir
	}
	return ""
}

func (m *PLAYERGROWTHDIR) GetSkillDir() string {
	if m != nil && m.SkillDir != nil {
		return *m.SkillDir
	}
	return ""
}

func (m *PLAYERGROWTHDIR) GetEnhanceDir() string {
	if m != nil && m.EnhanceDir != nil {
		return *m.EnhanceDir
	}
	return ""
}

func (m *PLAYERGROWTHDIR) GetStarDir() string {
	if m != nil && m.StarDir != nil {
		return *m.StarDir
	}
	return ""
}

func (m *PLAYERGROWTHDIR) GetRareDir() string {
	if m != nil && m.RareDir != nil {
		return *m.RareDir
	}
	return ""
}

type PLAYERGROWTHDIR_ARRAY struct {
	Items            []*PLAYERGROWTHDIR `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PLAYERGROWTHDIR_ARRAY) Reset()         { *m = PLAYERGROWTHDIR_ARRAY{} }
func (m *PLAYERGROWTHDIR_ARRAY) String() string { return proto.CompactTextString(m) }
func (*PLAYERGROWTHDIR_ARRAY) ProtoMessage()    {}

func (m *PLAYERGROWTHDIR_ARRAY) GetItems() []*PLAYERGROWTHDIR {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
