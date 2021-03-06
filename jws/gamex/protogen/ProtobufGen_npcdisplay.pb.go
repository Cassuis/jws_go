// Code generated by protoc-gen-go.
// source: ProtobufGen_npcdisplay.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type NPCDISPLAY struct {
	// * 对应AcDataList中的ID
	NPCID *string `protobuf:"bytes,1,req,def=" json:"NPCID,omitempty"`
	// * 战队等级
	CorpLevel *uint32 `protobuf:"varint,2,opt,def=0" json:"CorpLevel,omitempty"`
	// * 关卡ID
	StageID *string `protobuf:"bytes,3,opt,def=" json:"StageID,omitempty"`
	// * 寒暄句
	Dialogue *string `protobuf:"bytes,4,opt,def=" json:"Dialogue,omitempty"`
	// * 名字
	Name *string `protobuf:"bytes,5,opt,def=" json:"Name,omitempty"`
	// * 按钮名字
	Button *string `protobuf:"bytes,6,opt,def=" json:"Button,omitempty"`
	// * 对应功能
	Function *string `protobuf:"bytes,7,opt,def=" json:"Function,omitempty"`
	// * 对应图标
	Icon *string `protobuf:"bytes,8,opt,def=" json:"Icon,omitempty"`
	// * 对应头像
	Head             *string `protobuf:"bytes,9,opt,name=head,def=" json:"head,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NPCDISPLAY) Reset()         { *m = NPCDISPLAY{} }
func (m *NPCDISPLAY) String() string { return proto.CompactTextString(m) }
func (*NPCDISPLAY) ProtoMessage()    {}

const Default_NPCDISPLAY_CorpLevel uint32 = 0

func (m *NPCDISPLAY) GetNPCID() string {
	if m != nil && m.NPCID != nil {
		return *m.NPCID
	}
	return ""
}

func (m *NPCDISPLAY) GetCorpLevel() uint32 {
	if m != nil && m.CorpLevel != nil {
		return *m.CorpLevel
	}
	return Default_NPCDISPLAY_CorpLevel
}

func (m *NPCDISPLAY) GetStageID() string {
	if m != nil && m.StageID != nil {
		return *m.StageID
	}
	return ""
}

func (m *NPCDISPLAY) GetDialogue() string {
	if m != nil && m.Dialogue != nil {
		return *m.Dialogue
	}
	return ""
}

func (m *NPCDISPLAY) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NPCDISPLAY) GetButton() string {
	if m != nil && m.Button != nil {
		return *m.Button
	}
	return ""
}

func (m *NPCDISPLAY) GetFunction() string {
	if m != nil && m.Function != nil {
		return *m.Function
	}
	return ""
}

func (m *NPCDISPLAY) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *NPCDISPLAY) GetHead() string {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return ""
}

type NPCDISPLAY_ARRAY struct {
	Items            []*NPCDISPLAY `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NPCDISPLAY_ARRAY) Reset()         { *m = NPCDISPLAY_ARRAY{} }
func (m *NPCDISPLAY_ARRAY) String() string { return proto.CompactTextString(m) }
func (*NPCDISPLAY_ARRAY) ProtoMessage()    {}

func (m *NPCDISPLAY_ARRAY) GetItems() []*NPCDISPLAY {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
