// Code generated by protoc-gen-go.
// source: ProtobufGen_fbbuycount.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FBBUYCOUNT struct {
	// * 购买挑战次数
	BuyNum *uint32 `protobuf:"varint,1,opt,def=0" json:"BuyNum,omitempty"`
	// * HC价格
	CostHC           *uint32 `protobuf:"varint,2,opt,def=0" json:"CostHC,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FBBUYCOUNT) Reset()         { *m = FBBUYCOUNT{} }
func (m *FBBUYCOUNT) String() string { return proto.CompactTextString(m) }
func (*FBBUYCOUNT) ProtoMessage()    {}

const Default_FBBUYCOUNT_BuyNum uint32 = 0
const Default_FBBUYCOUNT_CostHC uint32 = 0

func (m *FBBUYCOUNT) GetBuyNum() uint32 {
	if m != nil && m.BuyNum != nil {
		return *m.BuyNum
	}
	return Default_FBBUYCOUNT_BuyNum
}

func (m *FBBUYCOUNT) GetCostHC() uint32 {
	if m != nil && m.CostHC != nil {
		return *m.CostHC
	}
	return Default_FBBUYCOUNT_CostHC
}

type FBBUYCOUNT_ARRAY struct {
	Items            []*FBBUYCOUNT `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *FBBUYCOUNT_ARRAY) Reset()         { *m = FBBUYCOUNT_ARRAY{} }
func (m *FBBUYCOUNT_ARRAY) String() string { return proto.CompactTextString(m) }
func (*FBBUYCOUNT_ARRAY) ProtoMessage()    {}

func (m *FBBUYCOUNT_ARRAY) GetItems() []*FBBUYCOUNT {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}