// Code generated by protoc-gen-go.
// source: ProtobufGen_initializesave.proto
// DO NOT EDIT!

package ProtobufGen

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type INITIALIZESAVE struct {
	// * 角色ID
	Role *uint32 `protobuf:"varint,1,req,def=0" json:"Role,omitempty"`
	// * 初始经验
	InitializeXP *uint32 `protobuf:"varint,2,opt,def=0" json:"InitializeXP,omitempty"`
	// * 初始武器
	InitWeapon *string `protobuf:"bytes,3,opt,def=" json:"InitWeapon,omitempty"`
	// * 初始胸甲
	InitChest *string `protobuf:"bytes,4,opt,def=" json:"InitChest,omitempty"`
	// * 初始腰带
	InitBelt *string `protobuf:"bytes,5,opt,def=" json:"InitBelt,omitempty"`
	// * 初始护腿
	InitLeggings *string `protobuf:"bytes,6,opt,def=" json:"InitLeggings,omitempty"`
	// * 初始护腕
	InitBracers *string `protobuf:"bytes,7,opt,def=" json:"InitBracers,omitempty"`
	// * 初始项链
	InitNecklace *string `protobuf:"bytes,8,opt,def=" json:"InitNecklace,omitempty"`
	// * 初始戒指
	InitRing *string `protobuf:"bytes,9,opt,def=" json:"InitRing,omitempty"`
	// * 初始购买钻
	InitHC_Buy *uint32 `protobuf:"varint,10,opt,def=0" json:"InitHC_Buy,omitempty"`
	// * 初始赠送钻
	InitHC_Give *uint32 `protobuf:"varint,11,opt,def=0" json:"InitHC_Give,omitempty"`
	// * 初始补偿钻
	InitHC_Compensate *uint32 `protobuf:"varint,12,opt,def=0" json:"InitHC_Compensate,omitempty"`
	// * 初始SC
	InitSC *uint32 `protobuf:"varint,13,opt,def=0" json:"InitSC,omitempty"`
	// * 初始精铁
	InitFineIron     *uint32 `protobuf:"varint,14,opt,def=0" json:"InitFineIron,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *INITIALIZESAVE) Reset()         { *m = INITIALIZESAVE{} }
func (m *INITIALIZESAVE) String() string { return proto.CompactTextString(m) }
func (*INITIALIZESAVE) ProtoMessage()    {}

const Default_INITIALIZESAVE_Role uint32 = 0
const Default_INITIALIZESAVE_InitializeXP uint32 = 0
const Default_INITIALIZESAVE_InitHC_Buy uint32 = 0
const Default_INITIALIZESAVE_InitHC_Give uint32 = 0
const Default_INITIALIZESAVE_InitHC_Compensate uint32 = 0
const Default_INITIALIZESAVE_InitSC uint32 = 0
const Default_INITIALIZESAVE_InitFineIron uint32 = 0

func (m *INITIALIZESAVE) GetRole() uint32 {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return Default_INITIALIZESAVE_Role
}

func (m *INITIALIZESAVE) GetInitializeXP() uint32 {
	if m != nil && m.InitializeXP != nil {
		return *m.InitializeXP
	}
	return Default_INITIALIZESAVE_InitializeXP
}

func (m *INITIALIZESAVE) GetInitWeapon() string {
	if m != nil && m.InitWeapon != nil {
		return *m.InitWeapon
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitChest() string {
	if m != nil && m.InitChest != nil {
		return *m.InitChest
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitBelt() string {
	if m != nil && m.InitBelt != nil {
		return *m.InitBelt
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitLeggings() string {
	if m != nil && m.InitLeggings != nil {
		return *m.InitLeggings
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitBracers() string {
	if m != nil && m.InitBracers != nil {
		return *m.InitBracers
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitNecklace() string {
	if m != nil && m.InitNecklace != nil {
		return *m.InitNecklace
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitRing() string {
	if m != nil && m.InitRing != nil {
		return *m.InitRing
	}
	return ""
}

func (m *INITIALIZESAVE) GetInitHC_Buy() uint32 {
	if m != nil && m.InitHC_Buy != nil {
		return *m.InitHC_Buy
	}
	return Default_INITIALIZESAVE_InitHC_Buy
}

func (m *INITIALIZESAVE) GetInitHC_Give() uint32 {
	if m != nil && m.InitHC_Give != nil {
		return *m.InitHC_Give
	}
	return Default_INITIALIZESAVE_InitHC_Give
}

func (m *INITIALIZESAVE) GetInitHC_Compensate() uint32 {
	if m != nil && m.InitHC_Compensate != nil {
		return *m.InitHC_Compensate
	}
	return Default_INITIALIZESAVE_InitHC_Compensate
}

func (m *INITIALIZESAVE) GetInitSC() uint32 {
	if m != nil && m.InitSC != nil {
		return *m.InitSC
	}
	return Default_INITIALIZESAVE_InitSC
}

func (m *INITIALIZESAVE) GetInitFineIron() uint32 {
	if m != nil && m.InitFineIron != nil {
		return *m.InitFineIron
	}
	return Default_INITIALIZESAVE_InitFineIron
}

type INITIALIZESAVE_ARRAY struct {
	Items            []*INITIALIZESAVE `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *INITIALIZESAVE_ARRAY) Reset()         { *m = INITIALIZESAVE_ARRAY{} }
func (m *INITIALIZESAVE_ARRAY) String() string { return proto.CompactTextString(m) }
func (*INITIALIZESAVE_ARRAY) ProtoMessage()    {}

func (m *INITIALIZESAVE_ARRAY) GetItems() []*INITIALIZESAVE {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
}