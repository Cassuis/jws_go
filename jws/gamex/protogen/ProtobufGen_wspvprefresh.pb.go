// Code generated by protoc-gen-go.
// source: ProtobufGen_wspvprefresh.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WSPVPREFRESH struct {
	// * 刷新次数
	WSPVPRefreshTime *uint32 `protobuf:"varint,3,req,def=0" json:"WSPVPRefreshTime,omitempty"`
	// * SC消耗
	WsPvPPrice       *uint32 `protobuf:"varint,2,req,def=0" json:"WsPvPPrice,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WSPVPREFRESH) Reset()         { *m = WSPVPREFRESH{} }
func (m *WSPVPREFRESH) String() string { return proto.CompactTextString(m) }
func (*WSPVPREFRESH) ProtoMessage()    {}

const Default_WSPVPREFRESH_WSPVPRefreshTime uint32 = 0
const Default_WSPVPREFRESH_WsPvPPrice uint32 = 0

func (m *WSPVPREFRESH) GetWSPVPRefreshTime() uint32 {
	if m != nil && m.WSPVPRefreshTime != nil {
		return *m.WSPVPRefreshTime
	}
	return Default_WSPVPREFRESH_WSPVPRefreshTime
}

func (m *WSPVPREFRESH) GetWsPvPPrice() uint32 {
	if m != nil && m.WsPvPPrice != nil {
		return *m.WsPvPPrice
	}
	return Default_WSPVPREFRESH_WsPvPPrice
}

type WSPVPREFRESH_ARRAY struct {
	Items            []*WSPVPREFRESH `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *WSPVPREFRESH_ARRAY) Reset()         { *m = WSPVPREFRESH_ARRAY{} }
func (m *WSPVPREFRESH_ARRAY) String() string { return proto.CompactTextString(m) }
func (*WSPVPREFRESH_ARRAY) ProtoMessage()    {}

func (m *WSPVPREFRESH_ARRAY) GetItems() []*WSPVPREFRESH {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
