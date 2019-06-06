// Code generated by protoc-gen-go.
// source: ProtobufGen_nameja.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type NAMEJA struct {
	// * 索引，第一列不能为汉字
	Index *uint32 `protobuf:"varint,1,req,def=0" json:"Index,omitempty"`
	// * 故郷
	HomeTown *string `protobuf:"bytes,2,opt,def=" json:"HomeTown,omitempty"`
	// * 姓
	FamilyName *string `protobuf:"bytes,3,opt,def=" json:"FamilyName,omitempty"`
	// * 名
	FirstName        *string `protobuf:"bytes,4,opt,def=" json:"FirstName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NAMEJA) Reset()         { *m = NAMEJA{} }
func (m *NAMEJA) String() string { return proto.CompactTextString(m) }
func (*NAMEJA) ProtoMessage()    {}

const Default_NAMEJA_Index uint32 = 0

func (m *NAMEJA) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_NAMEJA_Index
}

func (m *NAMEJA) GetHomeTown() string {
	if m != nil && m.HomeTown != nil {
		return *m.HomeTown
	}
	return ""
}

func (m *NAMEJA) GetFamilyName() string {
	if m != nil && m.FamilyName != nil {
		return *m.FamilyName
	}
	return ""
}

func (m *NAMEJA) GetFirstName() string {
	if m != nil && m.FirstName != nil {
		return *m.FirstName
	}
	return ""
}

type NAMEJA_ARRAY struct {
	Items            []*NAMEJA `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NAMEJA_ARRAY) Reset()         { *m = NAMEJA_ARRAY{} }
func (m *NAMEJA_ARRAY) String() string { return proto.CompactTextString(m) }
func (*NAMEJA_ARRAY) ProtoMessage()    {}

func (m *NAMEJA_ARRAY) GetItems() []*NAMEJA {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}