// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/transaction.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetPlayerEconomyRequest_EconomyTypeEnum int32

const (
	GetPlayerEconomyRequest_MONEY   GetPlayerEconomyRequest_EconomyTypeEnum = 0
	GetPlayerEconomyRequest_DIAMOND GetPlayerEconomyRequest_EconomyTypeEnum = 1
	GetPlayerEconomyRequest_BANK    GetPlayerEconomyRequest_EconomyTypeEnum = 2
)

// Enum value maps for GetPlayerEconomyRequest_EconomyTypeEnum.
var (
	GetPlayerEconomyRequest_EconomyTypeEnum_name = map[int32]string{
		0: "MONEY",
		1: "DIAMOND",
		2: "BANK",
	}
	GetPlayerEconomyRequest_EconomyTypeEnum_value = map[string]int32{
		"MONEY":   0,
		"DIAMOND": 1,
		"BANK":    2,
	}
)

func (x GetPlayerEconomyRequest_EconomyTypeEnum) Enum() *GetPlayerEconomyRequest_EconomyTypeEnum {
	p := new(GetPlayerEconomyRequest_EconomyTypeEnum)
	*p = x
	return p
}

func (x GetPlayerEconomyRequest_EconomyTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPlayerEconomyRequest_EconomyTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_transaction_proto_enumTypes[0].Descriptor()
}

func (GetPlayerEconomyRequest_EconomyTypeEnum) Type() protoreflect.EnumType {
	return &file_proto_transaction_proto_enumTypes[0]
}

func (x GetPlayerEconomyRequest_EconomyTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPlayerEconomyRequest_EconomyTypeEnum.Descriptor instead.
func (GetPlayerEconomyRequest_EconomyTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{3, 0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                    uint32               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Value                 int32                `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	TransactionTypeID     uint32               `protobuf:"varint,3,opt,name=TransactionTypeID,proto3" json:"TransactionTypeID,omitempty"`
	TransactionType       *TransactionType     `protobuf:"bytes,4,opt,name=TransactionType,proto3" json:"TransactionType,omitempty"`
	TransactionCategoryID uint32               `protobuf:"varint,5,opt,name=TransactionCategoryID,proto3" json:"TransactionCategoryID,omitempty"`
	TransactionCategory   *TransactionCategory `protobuf:"bytes,6,opt,name=TransactionCategory,proto3" json:"TransactionCategory,omitempty"`
	PlayerID              uint32               `protobuf:"varint,7,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Transaction) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Transaction) GetTransactionTypeID() uint32 {
	if x != nil {
		return x.TransactionTypeID
	}
	return 0
}

func (x *Transaction) GetTransactionType() *TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return nil
}

func (x *Transaction) GetTransactionCategoryID() uint32 {
	if x != nil {
		return x.TransactionCategoryID
	}
	return 0
}

func (x *Transaction) GetTransactionCategory() *TransactionCategory {
	if x != nil {
		return x.TransactionCategory
	}
	return nil
}

func (x *Transaction) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

// CreateTransaction
type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value                 int32  `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	TransactionTypeID     uint32 `protobuf:"varint,2,opt,name=TransactionTypeID,proto3" json:"TransactionTypeID,omitempty"`
	TransactionCategoryID uint32 `protobuf:"varint,3,opt,name=TransactionCategoryID,proto3" json:"TransactionCategoryID,omitempty"`
	PlayerID              uint32 `protobuf:"varint,4,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransactionRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CreateTransactionRequest) GetTransactionTypeID() uint32 {
	if x != nil {
		return x.TransactionTypeID
	}
	return 0
}

func (x *CreateTransactionRequest) GetTransactionCategoryID() uint32 {
	if x != nil {
		return x.TransactionCategoryID
	}
	return 0
}

func (x *CreateTransactionRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type CreateTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
}

func (x *CreateTransactionResponse) Reset() {
	*x = CreateTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionResponse) ProtoMessage() {}

func (x *CreateTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// GetPlayerEconomy
type GetPlayerEconomyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID    uint32                                  `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	EconomyType GetPlayerEconomyRequest_EconomyTypeEnum `protobuf:"varint,2,opt,name=EconomyType,proto3,enum=transaction.GetPlayerEconomyRequest_EconomyTypeEnum" json:"EconomyType,omitempty"`
}

func (x *GetPlayerEconomyRequest) Reset() {
	*x = GetPlayerEconomyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerEconomyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerEconomyRequest) ProtoMessage() {}

func (x *GetPlayerEconomyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerEconomyRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerEconomyRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerEconomyRequest) GetPlayerID() uint32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *GetPlayerEconomyRequest) GetEconomyType() GetPlayerEconomyRequest_EconomyTypeEnum {
	if x != nil {
		return x.EconomyType
	}
	return GetPlayerEconomyRequest_MONEY
}

type GetPlayerEconomyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetPlayerEconomyResponse) Reset() {
	*x = GetPlayerEconomyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerEconomyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerEconomyResponse) ProtoMessage() {}

func (x *GetPlayerEconomyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerEconomyResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerEconomyResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerEconomyResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_transaction_proto protoreflect.FileDescriptor

var file_proto_transaction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x4b, 0x0a, 0x0f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x5b, 0x0a,
	0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0xb0, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x56, 0x0a, 0x0b, 0x45, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x34, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0b, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x33, 0x0a, 0x0f, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x41, 0x4d, 0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x42, 0x41, 0x4e, 0x4b, 0x10, 0x02, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transaction_proto_rawDescOnce sync.Once
	file_proto_transaction_proto_rawDescData = file_proto_transaction_proto_rawDesc
)

func file_proto_transaction_proto_rawDescGZIP() []byte {
	file_proto_transaction_proto_rawDescOnce.Do(func() {
		file_proto_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transaction_proto_rawDescData)
	})
	return file_proto_transaction_proto_rawDescData
}

var file_proto_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_transaction_proto_goTypes = []interface{}{
	(GetPlayerEconomyRequest_EconomyTypeEnum)(0), // 0: transaction.GetPlayerEconomyRequest.EconomyTypeEnum
	(*Transaction)(nil),                          // 1: transaction.Transaction
	(*CreateTransactionRequest)(nil),             // 2: transaction.CreateTransactionRequest
	(*CreateTransactionResponse)(nil),            // 3: transaction.CreateTransactionResponse
	(*GetPlayerEconomyRequest)(nil),              // 4: transaction.GetPlayerEconomyRequest
	(*GetPlayerEconomyResponse)(nil),             // 5: transaction.GetPlayerEconomyResponse
	(*TransactionType)(nil),                      // 6: transaction_type.TransactionType
	(*TransactionCategory)(nil),                  // 7: transaction_category.TransactionCategory
}
var file_proto_transaction_proto_depIdxs = []int32{
	6, // 0: transaction.Transaction.TransactionType:type_name -> transaction_type.TransactionType
	7, // 1: transaction.Transaction.TransactionCategory:type_name -> transaction_category.TransactionCategory
	1, // 2: transaction.CreateTransactionResponse.Transaction:type_name -> transaction.Transaction
	0, // 3: transaction.GetPlayerEconomyRequest.EconomyType:type_name -> transaction.GetPlayerEconomyRequest.EconomyTypeEnum
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_transaction_proto_init() }
func file_proto_transaction_proto_init() {
	if File_proto_transaction_proto != nil {
		return
	}
	file_proto_transaction_type_proto_init()
	file_proto_transaction_category_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerEconomyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerEconomyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_transaction_proto_goTypes,
		DependencyIndexes: file_proto_transaction_proto_depIdxs,
		EnumInfos:         file_proto_transaction_proto_enumTypes,
		MessageInfos:      file_proto_transaction_proto_msgTypes,
	}.Build()
	File_proto_transaction_proto = out.File
	file_proto_transaction_proto_rawDesc = nil
	file_proto_transaction_proto_goTypes = nil
	file_proto_transaction_proto_depIdxs = nil
}
