// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/ship.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ship struct {
	ID                   uint32        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	PlayerID             uint32        `protobuf:"varint,3,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	RarityID             uint32        `protobuf:"varint,5,opt,name=RarityID,proto3" json:"RarityID,omitempty"`
	Rarity               *Rarity       `protobuf:"bytes,6,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	ShipCategoryID       uint32        `protobuf:"varint,7,opt,name=ShipCategoryID,proto3" json:"ShipCategoryID,omitempty"`
	ShipCategory         *ShipCategory `protobuf:"bytes,8,opt,name=ShipCategory,proto3" json:"ShipCategory,omitempty"`
	ShipStatsID          uint32        `protobuf:"varint,9,opt,name=ShipStatsID,proto3" json:"ShipStatsID,omitempty"`
	ShipStats            *ShipStats    `protobuf:"bytes,10,opt,name=ShipStats,proto3" json:"ShipStats,omitempty"`
	Equipped             bool          `protobuf:"varint,11,opt,name=Equipped,proto3" json:"Equipped,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ship) Reset()         { *m = Ship{} }
func (m *Ship) String() string { return proto.CompactTextString(m) }
func (*Ship) ProtoMessage()    {}
func (*Ship) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{0}
}

func (m *Ship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ship.Unmarshal(m, b)
}
func (m *Ship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ship.Marshal(b, m, deterministic)
}
func (m *Ship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ship.Merge(m, src)
}
func (m *Ship) XXX_Size() int {
	return xxx_messageInfo_Ship.Size(m)
}
func (m *Ship) XXX_DiscardUnknown() {
	xxx_messageInfo_Ship.DiscardUnknown(m)
}

var xxx_messageInfo_Ship proto.InternalMessageInfo

func (m *Ship) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Ship) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ship) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Ship) GetRarityID() uint32 {
	if m != nil {
		return m.RarityID
	}
	return 0
}

func (m *Ship) GetRarity() *Rarity {
	if m != nil {
		return m.Rarity
	}
	return nil
}

func (m *Ship) GetShipCategoryID() uint32 {
	if m != nil {
		return m.ShipCategoryID
	}
	return 0
}

func (m *Ship) GetShipCategory() *ShipCategory {
	if m != nil {
		return m.ShipCategory
	}
	return nil
}

func (m *Ship) GetShipStatsID() uint32 {
	if m != nil {
		return m.ShipStatsID
	}
	return 0
}

func (m *Ship) GetShipStats() *ShipStats {
	if m != nil {
		return m.ShipStats
	}
	return nil
}

func (m *Ship) GetEquipped() bool {
	if m != nil {
		return m.Equipped
	}
	return false
}

// GetShipRepairInfo
type GetShipRepairInfoRequest struct {
	Ship                 *Ship    `protobuf:"bytes,1,opt,name=Ship,proto3" json:"Ship,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShipRepairInfoRequest) Reset()         { *m = GetShipRepairInfoRequest{} }
func (m *GetShipRepairInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetShipRepairInfoRequest) ProtoMessage()    {}
func (*GetShipRepairInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{1}
}

func (m *GetShipRepairInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShipRepairInfoRequest.Unmarshal(m, b)
}
func (m *GetShipRepairInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShipRepairInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetShipRepairInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipRepairInfoRequest.Merge(m, src)
}
func (m *GetShipRepairInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetShipRepairInfoRequest.Size(m)
}
func (m *GetShipRepairInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipRepairInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipRepairInfoRequest proto.InternalMessageInfo

func (m *GetShipRepairInfoRequest) GetShip() *Ship {
	if m != nil {
		return m.Ship
	}
	return nil
}

type GetShipRepairInfoResponse struct {
	NeedRepairs          bool     `protobuf:"varint,1,opt,name=NeedRepairs,proto3" json:"NeedRepairs,omitempty"`
	RepairTime           int32    `protobuf:"varint,2,opt,name=RepairTime,proto3" json:"RepairTime,omitempty"`
	QuantityResources    int32    `protobuf:"varint,3,opt,name=QuantityResources,proto3" json:"QuantityResources,omitempty"`
	TypeResources        string   `protobuf:"bytes,4,opt,name=TypeResources,proto3" json:"TypeResources,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShipRepairInfoResponse) Reset()         { *m = GetShipRepairInfoResponse{} }
func (m *GetShipRepairInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetShipRepairInfoResponse) ProtoMessage()    {}
func (*GetShipRepairInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{2}
}

func (m *GetShipRepairInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShipRepairInfoResponse.Unmarshal(m, b)
}
func (m *GetShipRepairInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShipRepairInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetShipRepairInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipRepairInfoResponse.Merge(m, src)
}
func (m *GetShipRepairInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetShipRepairInfoResponse.Size(m)
}
func (m *GetShipRepairInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipRepairInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipRepairInfoResponse proto.InternalMessageInfo

func (m *GetShipRepairInfoResponse) GetNeedRepairs() bool {
	if m != nil {
		return m.NeedRepairs
	}
	return false
}

func (m *GetShipRepairInfoResponse) GetRepairTime() int32 {
	if m != nil {
		return m.RepairTime
	}
	return 0
}

func (m *GetShipRepairInfoResponse) GetQuantityResources() int32 {
	if m != nil {
		return m.QuantityResources
	}
	return 0
}

func (m *GetShipRepairInfoResponse) GetTypeResources() string {
	if m != nil {
		return m.TypeResources
	}
	return ""
}

// StartShipRepair
type StartShipRepairRequest struct {
	Ship                 *Ship    `protobuf:"bytes,1,opt,name=Ship,proto3" json:"Ship,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartShipRepairRequest) Reset()         { *m = StartShipRepairRequest{} }
func (m *StartShipRepairRequest) String() string { return proto.CompactTextString(m) }
func (*StartShipRepairRequest) ProtoMessage()    {}
func (*StartShipRepairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{3}
}

func (m *StartShipRepairRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartShipRepairRequest.Unmarshal(m, b)
}
func (m *StartShipRepairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartShipRepairRequest.Marshal(b, m, deterministic)
}
func (m *StartShipRepairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartShipRepairRequest.Merge(m, src)
}
func (m *StartShipRepairRequest) XXX_Size() int {
	return xxx_messageInfo_StartShipRepairRequest.Size(m)
}
func (m *StartShipRepairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartShipRepairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartShipRepairRequest proto.InternalMessageInfo

func (m *StartShipRepairRequest) GetShip() *Ship {
	if m != nil {
		return m.Ship
	}
	return nil
}

type StartShipRepair struct {
	ResourceID           uint32   `protobuf:"varint,1,opt,name=ResourceID,proto3" json:"ResourceID,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartShipRepair) Reset()         { *m = StartShipRepair{} }
func (m *StartShipRepair) String() string { return proto.CompactTextString(m) }
func (*StartShipRepair) ProtoMessage()    {}
func (*StartShipRepair) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{4}
}

func (m *StartShipRepair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartShipRepair.Unmarshal(m, b)
}
func (m *StartShipRepair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartShipRepair.Marshal(b, m, deterministic)
}
func (m *StartShipRepair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartShipRepair.Merge(m, src)
}
func (m *StartShipRepair) XXX_Size() int {
	return xxx_messageInfo_StartShipRepair.Size(m)
}
func (m *StartShipRepair) XXX_DiscardUnknown() {
	xxx_messageInfo_StartShipRepair.DiscardUnknown(m)
}

var xxx_messageInfo_StartShipRepair proto.InternalMessageInfo

func (m *StartShipRepair) GetResourceID() uint32 {
	if m != nil {
		return m.ResourceID
	}
	return 0
}

func (m *StartShipRepair) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type StartShipRepairResponse struct {
	StartShipRepair      []*StartShipRepair `protobuf:"bytes,1,rep,name=StartShipRepair,proto3" json:"StartShipRepair,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StartShipRepairResponse) Reset()         { *m = StartShipRepairResponse{} }
func (m *StartShipRepairResponse) String() string { return proto.CompactTextString(m) }
func (*StartShipRepairResponse) ProtoMessage()    {}
func (*StartShipRepairResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{5}
}

func (m *StartShipRepairResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartShipRepairResponse.Unmarshal(m, b)
}
func (m *StartShipRepairResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartShipRepairResponse.Marshal(b, m, deterministic)
}
func (m *StartShipRepairResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartShipRepairResponse.Merge(m, src)
}
func (m *StartShipRepairResponse) XXX_Size() int {
	return xxx_messageInfo_StartShipRepairResponse.Size(m)
}
func (m *StartShipRepairResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartShipRepairResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartShipRepairResponse proto.InternalMessageInfo

func (m *StartShipRepairResponse) GetStartShipRepair() []*StartShipRepair {
	if m != nil {
		return m.StartShipRepair
	}
	return nil
}

// StartShipRepair
type EndShipRepairRequest struct {
	Ship                 *Ship    `protobuf:"bytes,1,opt,name=Ship,proto3" json:"Ship,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndShipRepairRequest) Reset()         { *m = EndShipRepairRequest{} }
func (m *EndShipRepairRequest) String() string { return proto.CompactTextString(m) }
func (*EndShipRepairRequest) ProtoMessage()    {}
func (*EndShipRepairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{6}
}

func (m *EndShipRepairRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndShipRepairRequest.Unmarshal(m, b)
}
func (m *EndShipRepairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndShipRepairRequest.Marshal(b, m, deterministic)
}
func (m *EndShipRepairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndShipRepairRequest.Merge(m, src)
}
func (m *EndShipRepairRequest) XXX_Size() int {
	return xxx_messageInfo_EndShipRepairRequest.Size(m)
}
func (m *EndShipRepairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndShipRepairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndShipRepairRequest proto.InternalMessageInfo

func (m *EndShipRepairRequest) GetShip() *Ship {
	if m != nil {
		return m.Ship
	}
	return nil
}

type EndShipRepairResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndShipRepairResponse) Reset()         { *m = EndShipRepairResponse{} }
func (m *EndShipRepairResponse) String() string { return proto.CompactTextString(m) }
func (*EndShipRepairResponse) ProtoMessage()    {}
func (*EndShipRepairResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{7}
}

func (m *EndShipRepairResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndShipRepairResponse.Unmarshal(m, b)
}
func (m *EndShipRepairResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndShipRepairResponse.Marshal(b, m, deterministic)
}
func (m *EndShipRepairResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndShipRepairResponse.Merge(m, src)
}
func (m *EndShipRepairResponse) XXX_Size() int {
	return xxx_messageInfo_EndShipRepairResponse.Size(m)
}
func (m *EndShipRepairResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndShipRepairResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndShipRepairResponse proto.InternalMessageInfo

// GetShipExplorationInfo
type GetShipExplorationInfoRequest struct {
	Ship                 *Ship    `protobuf:"bytes,1,opt,name=Ship,proto3" json:"Ship,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShipExplorationInfoRequest) Reset()         { *m = GetShipExplorationInfoRequest{} }
func (m *GetShipExplorationInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetShipExplorationInfoRequest) ProtoMessage()    {}
func (*GetShipExplorationInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{8}
}

func (m *GetShipExplorationInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShipExplorationInfoRequest.Unmarshal(m, b)
}
func (m *GetShipExplorationInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShipExplorationInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetShipExplorationInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipExplorationInfoRequest.Merge(m, src)
}
func (m *GetShipExplorationInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetShipExplorationInfoRequest.Size(m)
}
func (m *GetShipExplorationInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipExplorationInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipExplorationInfoRequest proto.InternalMessageInfo

func (m *GetShipExplorationInfoRequest) GetShip() *Ship {
	if m != nil {
		return m.Ship
	}
	return nil
}

type GetShipExplorationInfo struct {
	Planet               *Planet  `protobuf:"bytes,1,opt,name=Planet,proto3" json:"Planet,omitempty"`
	Distance             float64  `protobuf:"fixed64,2,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Time                 int32    `protobuf:"varint,3,opt,name=Time,proto3" json:"Time,omitempty"`
	Fuel                 float64  `protobuf:"fixed64,4,opt,name=Fuel,proto3" json:"Fuel,omitempty"`
	Integrity            uint32   `protobuf:"varint,5,opt,name=Integrity,proto3" json:"Integrity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShipExplorationInfo) Reset()         { *m = GetShipExplorationInfo{} }
func (m *GetShipExplorationInfo) String() string { return proto.CompactTextString(m) }
func (*GetShipExplorationInfo) ProtoMessage()    {}
func (*GetShipExplorationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{9}
}

func (m *GetShipExplorationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShipExplorationInfo.Unmarshal(m, b)
}
func (m *GetShipExplorationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShipExplorationInfo.Marshal(b, m, deterministic)
}
func (m *GetShipExplorationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipExplorationInfo.Merge(m, src)
}
func (m *GetShipExplorationInfo) XXX_Size() int {
	return xxx_messageInfo_GetShipExplorationInfo.Size(m)
}
func (m *GetShipExplorationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipExplorationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipExplorationInfo proto.InternalMessageInfo

func (m *GetShipExplorationInfo) GetPlanet() *Planet {
	if m != nil {
		return m.Planet
	}
	return nil
}

func (m *GetShipExplorationInfo) GetDistance() float64 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *GetShipExplorationInfo) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *GetShipExplorationInfo) GetFuel() float64 {
	if m != nil {
		return m.Fuel
	}
	return 0
}

func (m *GetShipExplorationInfo) GetIntegrity() uint32 {
	if m != nil {
		return m.Integrity
	}
	return 0
}

type GetShipExplorationInfoResponse struct {
	GetShipExplorationInfo []*GetShipExplorationInfo `protobuf:"bytes,1,rep,name=GetShipExplorationInfo,proto3" json:"GetShipExplorationInfo,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                  `json:"-"`
	XXX_unrecognized       []byte                    `json:"-"`
	XXX_sizecache          int32                     `json:"-"`
}

func (m *GetShipExplorationInfoResponse) Reset()         { *m = GetShipExplorationInfoResponse{} }
func (m *GetShipExplorationInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetShipExplorationInfoResponse) ProtoMessage()    {}
func (*GetShipExplorationInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{10}
}

func (m *GetShipExplorationInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShipExplorationInfoResponse.Unmarshal(m, b)
}
func (m *GetShipExplorationInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShipExplorationInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetShipExplorationInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipExplorationInfoResponse.Merge(m, src)
}
func (m *GetShipExplorationInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetShipExplorationInfoResponse.Size(m)
}
func (m *GetShipExplorationInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipExplorationInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipExplorationInfoResponse proto.InternalMessageInfo

func (m *GetShipExplorationInfoResponse) GetGetShipExplorationInfo() []*GetShipExplorationInfo {
	if m != nil {
		return m.GetShipExplorationInfo
	}
	return nil
}

// EndShipExploration
type EndShipExplorationRequest struct {
	Integrity            uint32   `protobuf:"varint,1,opt,name=Integrity,proto3" json:"Integrity,omitempty"`
	Tank                 float64  `protobuf:"fixed64,2,opt,name=Tank,proto3" json:"Tank,omitempty"`
	PositionX            float64  `protobuf:"fixed64,3,opt,name=PositionX,proto3" json:"PositionX,omitempty"`
	PositionY            float64  `protobuf:"fixed64,4,opt,name=PositionY,proto3" json:"PositionY,omitempty"`
	PositionZ            float64  `protobuf:"fixed64,5,opt,name=PositionZ,proto3" json:"PositionZ,omitempty"`
	ShipID               uint32   `protobuf:"varint,6,opt,name=ShipID,proto3" json:"ShipID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndShipExplorationRequest) Reset()         { *m = EndShipExplorationRequest{} }
func (m *EndShipExplorationRequest) String() string { return proto.CompactTextString(m) }
func (*EndShipExplorationRequest) ProtoMessage()    {}
func (*EndShipExplorationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{11}
}

func (m *EndShipExplorationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndShipExplorationRequest.Unmarshal(m, b)
}
func (m *EndShipExplorationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndShipExplorationRequest.Marshal(b, m, deterministic)
}
func (m *EndShipExplorationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndShipExplorationRequest.Merge(m, src)
}
func (m *EndShipExplorationRequest) XXX_Size() int {
	return xxx_messageInfo_EndShipExplorationRequest.Size(m)
}
func (m *EndShipExplorationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndShipExplorationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndShipExplorationRequest proto.InternalMessageInfo

func (m *EndShipExplorationRequest) GetIntegrity() uint32 {
	if m != nil {
		return m.Integrity
	}
	return 0
}

func (m *EndShipExplorationRequest) GetTank() float64 {
	if m != nil {
		return m.Tank
	}
	return 0
}

func (m *EndShipExplorationRequest) GetPositionX() float64 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *EndShipExplorationRequest) GetPositionY() float64 {
	if m != nil {
		return m.PositionY
	}
	return 0
}

func (m *EndShipExplorationRequest) GetPositionZ() float64 {
	if m != nil {
		return m.PositionZ
	}
	return 0
}

func (m *EndShipExplorationRequest) GetShipID() uint32 {
	if m != nil {
		return m.ShipID
	}
	return 0
}

type EndShipExplorationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndShipExplorationResponse) Reset()         { *m = EndShipExplorationResponse{} }
func (m *EndShipExplorationResponse) String() string { return proto.CompactTextString(m) }
func (*EndShipExplorationResponse) ProtoMessage()    {}
func (*EndShipExplorationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{12}
}

func (m *EndShipExplorationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndShipExplorationResponse.Unmarshal(m, b)
}
func (m *EndShipExplorationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndShipExplorationResponse.Marshal(b, m, deterministic)
}
func (m *EndShipExplorationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndShipExplorationResponse.Merge(m, src)
}
func (m *EndShipExplorationResponse) XXX_Size() int {
	return xxx_messageInfo_EndShipExplorationResponse.Size(m)
}
func (m *EndShipExplorationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndShipExplorationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndShipExplorationResponse proto.InternalMessageInfo

// GetPlayerShips
type GetPlayerShipsRequest struct {
	PlayerID             uint32   `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerShipsRequest) Reset()         { *m = GetPlayerShipsRequest{} }
func (m *GetPlayerShipsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlayerShipsRequest) ProtoMessage()    {}
func (*GetPlayerShipsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{13}
}

func (m *GetPlayerShipsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerShipsRequest.Unmarshal(m, b)
}
func (m *GetPlayerShipsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerShipsRequest.Marshal(b, m, deterministic)
}
func (m *GetPlayerShipsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerShipsRequest.Merge(m, src)
}
func (m *GetPlayerShipsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPlayerShipsRequest.Size(m)
}
func (m *GetPlayerShipsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerShipsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerShipsRequest proto.InternalMessageInfo

func (m *GetPlayerShipsRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

type GetPlayerShipsResponse struct {
	Ships                []*Ship  `protobuf:"bytes,1,rep,name=Ships,proto3" json:"Ships,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerShipsResponse) Reset()         { *m = GetPlayerShipsResponse{} }
func (m *GetPlayerShipsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlayerShipsResponse) ProtoMessage()    {}
func (*GetPlayerShipsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{14}
}

func (m *GetPlayerShipsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerShipsResponse.Unmarshal(m, b)
}
func (m *GetPlayerShipsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerShipsResponse.Marshal(b, m, deterministic)
}
func (m *GetPlayerShipsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerShipsResponse.Merge(m, src)
}
func (m *GetPlayerShipsResponse) XXX_Size() int {
	return xxx_messageInfo_GetPlayerShipsResponse.Size(m)
}
func (m *GetPlayerShipsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerShipsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerShipsResponse proto.InternalMessageInfo

func (m *GetPlayerShipsResponse) GetShips() []*Ship {
	if m != nil {
		return m.Ships
	}
	return nil
}

// GetPlayerShipEquipped
type GetPlayerShipEquippedRequest struct {
	PlayerID             uint32   `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerShipEquippedRequest) Reset()         { *m = GetPlayerShipEquippedRequest{} }
func (m *GetPlayerShipEquippedRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlayerShipEquippedRequest) ProtoMessage()    {}
func (*GetPlayerShipEquippedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{15}
}

func (m *GetPlayerShipEquippedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerShipEquippedRequest.Unmarshal(m, b)
}
func (m *GetPlayerShipEquippedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerShipEquippedRequest.Marshal(b, m, deterministic)
}
func (m *GetPlayerShipEquippedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerShipEquippedRequest.Merge(m, src)
}
func (m *GetPlayerShipEquippedRequest) XXX_Size() int {
	return xxx_messageInfo_GetPlayerShipEquippedRequest.Size(m)
}
func (m *GetPlayerShipEquippedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerShipEquippedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerShipEquippedRequest proto.InternalMessageInfo

func (m *GetPlayerShipEquippedRequest) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

type GetPlayerShipEquippedResponse struct {
	Ship                 *Ship    `protobuf:"bytes,1,opt,name=Ship,proto3" json:"Ship,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerShipEquippedResponse) Reset()         { *m = GetPlayerShipEquippedResponse{} }
func (m *GetPlayerShipEquippedResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlayerShipEquippedResponse) ProtoMessage()    {}
func (*GetPlayerShipEquippedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a527018f6722ef, []int{16}
}

func (m *GetPlayerShipEquippedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerShipEquippedResponse.Unmarshal(m, b)
}
func (m *GetPlayerShipEquippedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerShipEquippedResponse.Marshal(b, m, deterministic)
}
func (m *GetPlayerShipEquippedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerShipEquippedResponse.Merge(m, src)
}
func (m *GetPlayerShipEquippedResponse) XXX_Size() int {
	return xxx_messageInfo_GetPlayerShipEquippedResponse.Size(m)
}
func (m *GetPlayerShipEquippedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerShipEquippedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerShipEquippedResponse proto.InternalMessageInfo

func (m *GetPlayerShipEquippedResponse) GetShip() *Ship {
	if m != nil {
		return m.Ship
	}
	return nil
}

func init() {
	proto.RegisterType((*Ship)(nil), "ship.Ship")
	proto.RegisterType((*GetShipRepairInfoRequest)(nil), "ship.GetShipRepairInfoRequest")
	proto.RegisterType((*GetShipRepairInfoResponse)(nil), "ship.GetShipRepairInfoResponse")
	proto.RegisterType((*StartShipRepairRequest)(nil), "ship.StartShipRepairRequest")
	proto.RegisterType((*StartShipRepair)(nil), "ship.StartShipRepair")
	proto.RegisterType((*StartShipRepairResponse)(nil), "ship.StartShipRepairResponse")
	proto.RegisterType((*EndShipRepairRequest)(nil), "ship.EndShipRepairRequest")
	proto.RegisterType((*EndShipRepairResponse)(nil), "ship.EndShipRepairResponse")
	proto.RegisterType((*GetShipExplorationInfoRequest)(nil), "ship.GetShipExplorationInfoRequest")
	proto.RegisterType((*GetShipExplorationInfo)(nil), "ship.GetShipExplorationInfo")
	proto.RegisterType((*GetShipExplorationInfoResponse)(nil), "ship.GetShipExplorationInfoResponse")
	proto.RegisterType((*EndShipExplorationRequest)(nil), "ship.EndShipExplorationRequest")
	proto.RegisterType((*EndShipExplorationResponse)(nil), "ship.EndShipExplorationResponse")
	proto.RegisterType((*GetPlayerShipsRequest)(nil), "ship.GetPlayerShipsRequest")
	proto.RegisterType((*GetPlayerShipsResponse)(nil), "ship.GetPlayerShipsResponse")
	proto.RegisterType((*GetPlayerShipEquippedRequest)(nil), "ship.GetPlayerShipEquippedRequest")
	proto.RegisterType((*GetPlayerShipEquippedResponse)(nil), "ship.GetPlayerShipEquippedResponse")
}

func init() { proto.RegisterFile("rpc/ship.proto", fileDescriptor_f4a527018f6722ef) }

var fileDescriptor_f4a527018f6722ef = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x37, 0xa7, 0x3f, 0x99, 0xd0, 0x00, 0xab, 0x1e, 0xdc, 0x10, 0xaa, 0xc8, 0x42, 0x55,
	0x2e, 0x50, 0x90, 0x5a, 0x09, 0xa1, 0xde, 0x54, 0x02, 0x97, 0xca, 0x17, 0x54, 0x66, 0xdb, 0x0b,
	0xc8, 0x0d, 0x32, 0xe9, 0xd2, 0x5a, 0x04, 0xdb, 0xf5, 0x6e, 0x10, 0x7d, 0x1d, 0xde, 0x80, 0x27,
	0xe0, 0x6d, 0x78, 0x0e, 0xb4, 0xb3, 0x6b, 0x67, 0xed, 0x24, 0xa2, 0xbd, 0x9b, 0xf9, 0xe6, 0xb0,
	0x33, 0xdf, 0x8c, 0xc7, 0xd0, 0xcb, 0xd2, 0xe9, 0x0b, 0x71, 0x1d, 0xa5, 0xe3, 0x34, 0x4b, 0x64,
	0x42, 0x1b, 0x4a, 0xee, 0x3f, 0x52, 0x68, 0x16, 0x66, 0x91, 0xbc, 0xd5, 0x78, 0x7f, 0x27, 0xf7,
	0xfb, 0x34, 0x0d, 0x25, 0xbf, 0x4a, 0xb2, 0xdc, 0xb0, 0x59, 0x18, 0x84, 0x0c, 0xa5, 0x30, 0x28,
	0x26, 0x48, 0x67, 0x61, 0xcc, 0xa5, 0x46, 0xdc, 0x3f, 0x35, 0x68, 0x9c, 0x5f, 0x47, 0x29, 0xed,
	0x41, 0xcd, 0xf7, 0x1c, 0x32, 0x24, 0xa3, 0x0d, 0x56, 0xf3, 0x3d, 0x4a, 0xa1, 0x71, 0x16, 0x7e,
	0xe3, 0x4e, 0x6d, 0x48, 0x46, 0x1d, 0x86, 0x32, 0xed, 0x43, 0x3b, 0x98, 0x85, 0xb7, 0x3c, 0xf3,
	0x3d, 0xa7, 0x8e, 0x9e, 0x85, 0xae, 0x6c, 0x0c, 0x2b, 0xf3, 0x3d, 0xa7, 0xa9, 0x6d, 0xb9, 0x4e,
	0xf7, 0xa1, 0xa5, 0x65, 0xa7, 0x35, 0x24, 0xa3, 0xee, 0x41, 0x6f, 0x6c, 0x9a, 0xd0, 0x28, 0x33,
	0x56, 0xba, 0x0f, 0x3d, 0x55, 0xcb, 0x1b, 0xd3, 0x8a, 0xef, 0x39, 0xff, 0x63, 0xa6, 0x0a, 0x4a,
	0x8f, 0xe1, 0x81, 0x8d, 0x38, 0x6d, 0xcc, 0xfa, 0x64, 0x5c, 0x26, 0xc2, 0x76, 0x61, 0xa5, 0x00,
	0x3a, 0x84, 0xae, 0xd2, 0xcf, 0x15, 0x35, 0xbe, 0xe7, 0x74, 0xf0, 0x15, 0x1b, 0xa2, 0x87, 0xd0,
	0x29, 0x54, 0x07, 0x30, 0xff, 0xd6, 0xd8, 0xe2, 0xb3, 0x30, 0xb2, 0x85, 0x9f, 0xe2, 0xe0, 0xe4,
	0x66, 0x1e, 0xa5, 0x29, 0xbf, 0x74, 0xba, 0x43, 0x32, 0x6a, 0xb3, 0x42, 0x77, 0x8f, 0xc0, 0x39,
	0xe5, 0x52, 0xf9, 0x32, 0x9e, 0x86, 0x51, 0xe6, 0xc7, 0x5f, 0x12, 0xc6, 0x6f, 0xe6, 0x5c, 0x48,
	0xba, 0xa7, 0x67, 0x80, 0xec, 0x77, 0x0f, 0x00, 0xdf, 0xc1, 0x17, 0x18, 0xe2, 0xee, 0x2f, 0x02,
	0xbb, 0x2b, 0x82, 0x45, 0x9a, 0xc4, 0x82, 0xab, 0x66, 0xce, 0x38, 0xbf, 0xd4, 0x16, 0x81, 0x49,
	0xda, 0xcc, 0x86, 0xe8, 0x1e, 0x80, 0x16, 0x2f, 0x22, 0x33, 0xd1, 0x26, 0xb3, 0x10, 0xfa, 0x1c,
	0x1e, 0xbf, 0x9f, 0x87, 0xb1, 0x54, 0xb3, 0xe0, 0x22, 0x99, 0x67, 0x53, 0x2e, 0x70, 0xc0, 0x4d,
	0xb6, 0x6c, 0xa0, 0xcf, 0x60, 0xe3, 0xe2, 0x36, 0xe5, 0x0b, 0xcf, 0x06, 0xae, 0x48, 0x19, 0x74,
	0x5f, 0xc1, 0xf6, 0xb9, 0x0c, 0x33, 0xab, 0xe8, 0xbb, 0x76, 0xfb, 0x0e, 0x1e, 0x56, 0x22, 0x75,
	0x03, 0x3a, 0x73, 0xb1, 0xa4, 0x16, 0xa2, 0x88, 0xcf, 0xeb, 0x34, 0xed, 0x15, 0xba, 0x3b, 0x81,
	0x9d, 0xa5, 0x42, 0x0c, 0x73, 0xc7, 0x4b, 0x2f, 0x39, 0x64, 0x58, 0x2f, 0x46, 0x3d, 0xae, 0xc6,
	0x55, 0xbd, 0xdd, 0x97, 0xb0, 0x79, 0x12, 0x5f, 0xde, 0xbf, 0xc5, 0x1d, 0xd8, 0xaa, 0xc4, 0xe9,
	0x8a, 0xdc, 0x63, 0x78, 0x6a, 0x06, 0x7d, 0xf2, 0x23, 0x9d, 0x25, 0x59, 0x28, 0xa3, 0x24, 0xbe,
	0xcf, 0xaa, 0xfc, 0x24, 0xb0, 0xbd, 0x3a, 0x83, 0xfa, 0x0a, 0x03, 0xfc, 0xf4, 0x4d, 0x70, 0x6f,
	0x6c, 0x2e, 0x81, 0x46, 0x99, 0xb1, 0x2a, 0x32, 0xbd, 0x48, 0xc8, 0x30, 0x9e, 0xea, 0x5d, 0x21,
	0xac, 0xd0, 0xd5, 0x55, 0xc0, 0x1d, 0xd2, 0xcb, 0x81, 0xb2, 0xc2, 0xde, 0xce, 0xf9, 0x0c, 0xd7,
	0x80, 0x30, 0x94, 0xe9, 0x00, 0x3a, 0x7e, 0x2c, 0xf9, 0x15, 0x7e, 0xf4, 0xfa, 0x1c, 0x2c, 0x00,
	0xf7, 0x3b, 0xec, 0xad, 0xeb, 0xd2, 0x4c, 0xe6, 0x62, 0x5d, 0x17, 0x66, 0x40, 0x03, 0xdd, 0xf8,
	0x9a, 0x2c, 0x6b, 0x62, 0xdd, 0xdf, 0x04, 0x76, 0x0d, 0xef, 0x96, 0x29, 0xa7, 0xb6, 0x54, 0x33,
	0xa9, 0xd4, 0x8c, 0x9d, 0x87, 0xf1, 0x57, 0xc3, 0x08, 0xca, 0x2a, 0x22, 0x48, 0x44, 0xa4, 0x92,
	0x7c, 0x40, 0x4a, 0x08, 0x5b, 0x00, 0xb6, 0xf5, 0xa3, 0x21, 0x67, 0x01, 0xd8, 0xd6, 0x09, 0x32,
	0x64, 0x59, 0x27, 0x74, 0x1b, 0x5a, 0xaa, 0x4a, 0xdf, 0xc3, 0x8b, 0xb9, 0xc1, 0x8c, 0xe6, 0x0e,
	0xa0, 0xbf, 0xaa, 0x01, 0xb3, 0x3d, 0x87, 0xb0, 0x75, 0xca, 0xa5, 0x3e, 0xc9, 0xca, 0x47, 0xe4,
	0xad, 0xd9, 0x87, 0x9b, 0x94, 0x0f, 0xb7, 0x7b, 0x84, 0x54, 0x97, 0x82, 0x8a, 0xc3, 0xd2, 0x44,
	0xc0, 0x70, 0x6e, 0x2f, 0x9b, 0x36, 0xb8, 0x47, 0x30, 0x28, 0xc5, 0xe6, 0xd7, 0xee, 0x2e, 0xef,
	0xea, 0x55, 0x5f, 0x15, 0x6b, 0x9e, 0xff, 0xc7, 0xaa, 0xbf, 0x6e, 0x4e, 0xea, 0x59, 0x3a, 0x0d,
	0xfe, 0x0b, 0x48, 0x50, 0x0b, 0xea, 0x9f, 0x5b, 0xf8, 0x43, 0x3b, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x32, 0xe3, 0x90, 0x43, 0x3b, 0x07, 0x00, 0x00,
}
