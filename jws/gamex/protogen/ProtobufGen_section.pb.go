// Code generated by protoc-gen-go.
// source: ProtobufGen_section.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SECTION struct {
	// * 章节ID
	SectionID *uint32 `protobuf:"varint,1,req,def=0" json:"SectionID,omitempty"`
	// * 章节名
	SectionName *string `protobuf:"bytes,2,opt,def=" json:"SectionName,omitempty"`
	// * 包含故事
	Conten *string `protobuf:"bytes,3,opt,def=" json:"Conten,omitempty"`
	// * 对应拼图的编号
	PuzzlePic        *string `protobuf:"bytes,4,opt,def=" json:"PuzzlePic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SECTION) Reset()         { *m = SECTION{} }
func (m *SECTION) String() string { return proto.CompactTextString(m) }
func (*SECTION) ProtoMessage()    {}

const Default_SECTION_SectionID uint32 = 0

func (m *SECTION) GetSectionID() uint32 {
	if m != nil && m.SectionID != nil {
		return *m.SectionID
	}
	return Default_SECTION_SectionID
}

func (m *SECTION) GetSectionName() string {
	if m != nil && m.SectionName != nil {
		return *m.SectionName
	}
	return ""
}

func (m *SECTION) GetConten() string {
	if m != nil && m.Conten != nil {
		return *m.Conten
	}
	return ""
}

func (m *SECTION) GetPuzzlePic() string {
	if m != nil && m.PuzzlePic != nil {
		return *m.PuzzlePic
	}
	return ""
}

type SECTION_ARRAY struct {
	Items            []*SECTION `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SECTION_ARRAY) Reset()         { *m = SECTION_ARRAY{} }
func (m *SECTION_ARRAY) String() string { return proto.CompactTextString(m) }
func (*SECTION_ARRAY) ProtoMessage()    {}

func (m *SECTION_ARRAY) GetItems() []*SECTION {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
