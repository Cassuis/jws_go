// Code generated by protoc-gen-go.
// source: ProtobufGen_gvghotdata.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GVGHOTDATA struct {
	// * 开服X小时内不开放（根据StartTime）
	SuspensionHour *uint32 `protobuf:"varint,1,opt,def=0" json:"SuspensionHour,omitempty"`
	// * 周几预告
	ReportWeek *uint32 `protobuf:"varint,2,opt,def=0" json:"ReportWeek,omitempty"`
	// * 预告时间
	AnnounceTime *string `protobuf:"bytes,3,opt,def=" json:"AnnounceTime,omitempty"`
	// * 周几活动重置
	RestartAndStartWeek *uint32 `protobuf:"varint,4,opt,def=0" json:"RestartAndStartWeek,omitempty"`
	// * 重置城池时间
	RestartCityTime *string `protobuf:"bytes,5,opt,def=" json:"RestartCityTime,omitempty"`
	// * 开启活动时间
	StartTime *string `protobuf:"bytes,6,opt,def=" json:"StartTime,omitempty"`
	// * 持续时间(分钟）
	GVGOpeningTime *uint32 `protobuf:"varint,7,opt,def=0" json:"GVGOpeningTime,omitempty"`
	// * 结算持续时间（分钟）
	GVGStatementTime *uint32 `protobuf:"varint,8,opt,def=0" json:"GVGStatementTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GVGHOTDATA) Reset()         { *m = GVGHOTDATA{} }
func (m *GVGHOTDATA) String() string { return proto.CompactTextString(m) }
func (*GVGHOTDATA) ProtoMessage()    {}

const Default_GVGHOTDATA_SuspensionHour uint32 = 0
const Default_GVGHOTDATA_ReportWeek uint32 = 0
const Default_GVGHOTDATA_RestartAndStartWeek uint32 = 0
const Default_GVGHOTDATA_GVGOpeningTime uint32 = 0
const Default_GVGHOTDATA_GVGStatementTime uint32 = 0

func (m *GVGHOTDATA) GetSuspensionHour() uint32 {
	if m != nil && m.SuspensionHour != nil {
		return *m.SuspensionHour
	}
	return Default_GVGHOTDATA_SuspensionHour
}

func (m *GVGHOTDATA) GetReportWeek() uint32 {
	if m != nil && m.ReportWeek != nil {
		return *m.ReportWeek
	}
	return Default_GVGHOTDATA_ReportWeek
}

func (m *GVGHOTDATA) GetAnnounceTime() string {
	if m != nil && m.AnnounceTime != nil {
		return *m.AnnounceTime
	}
	return ""
}

func (m *GVGHOTDATA) GetRestartAndStartWeek() uint32 {
	if m != nil && m.RestartAndStartWeek != nil {
		return *m.RestartAndStartWeek
	}
	return Default_GVGHOTDATA_RestartAndStartWeek
}

func (m *GVGHOTDATA) GetRestartCityTime() string {
	if m != nil && m.RestartCityTime != nil {
		return *m.RestartCityTime
	}
	return ""
}

func (m *GVGHOTDATA) GetStartTime() string {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return ""
}

func (m *GVGHOTDATA) GetGVGOpeningTime() uint32 {
	if m != nil && m.GVGOpeningTime != nil {
		return *m.GVGOpeningTime
	}
	return Default_GVGHOTDATA_GVGOpeningTime
}

func (m *GVGHOTDATA) GetGVGStatementTime() uint32 {
	if m != nil && m.GVGStatementTime != nil {
		return *m.GVGStatementTime
	}
	return Default_GVGHOTDATA_GVGStatementTime
}

type GVGHOTDATA_ARRAY struct {
	Items            []*GVGHOTDATA `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GVGHOTDATA_ARRAY) Reset()         { *m = GVGHOTDATA_ARRAY{} }
func (m *GVGHOTDATA_ARRAY) String() string { return proto.CompactTextString(m) }
func (*GVGHOTDATA_ARRAY) ProtoMessage()    {}

func (m *GVGHOTDATA_ARRAY) GetItems() []*GVGHOTDATA {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}
