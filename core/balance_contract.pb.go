// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/contract/balance_contract.proto

package core

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

type FreezeBalanceContract struct {
	OwnerAddress         []byte       `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	FrozenBalance        int64        `protobuf:"varint,2,opt,name=frozen_balance,json=frozenBalance,proto3" json:"frozen_balance,omitempty"`
	FrozenDuration       int64        `protobuf:"varint,3,opt,name=frozen_duration,json=frozenDuration,proto3" json:"frozen_duration,omitempty"`
	Resource             ResourceCode `protobuf:"varint,10,opt,name=resource,proto3,enum=protocol.ResourceCode" json:"resource,omitempty"`
	ReceiverAddress      []byte       `protobuf:"bytes,15,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FreezeBalanceContract) Reset()         { *m = FreezeBalanceContract{} }
func (m *FreezeBalanceContract) String() string { return proto.CompactTextString(m) }
func (*FreezeBalanceContract) ProtoMessage()    {}
func (*FreezeBalanceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{0}
}

func (m *FreezeBalanceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreezeBalanceContract.Unmarshal(m, b)
}
func (m *FreezeBalanceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreezeBalanceContract.Marshal(b, m, deterministic)
}
func (m *FreezeBalanceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreezeBalanceContract.Merge(m, src)
}
func (m *FreezeBalanceContract) XXX_Size() int {
	return xxx_messageInfo_FreezeBalanceContract.Size(m)
}
func (m *FreezeBalanceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_FreezeBalanceContract.DiscardUnknown(m)
}

var xxx_messageInfo_FreezeBalanceContract proto.InternalMessageInfo

func (m *FreezeBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *FreezeBalanceContract) GetFrozenBalance() int64 {
	if m != nil {
		return m.FrozenBalance
	}
	return 0
}

func (m *FreezeBalanceContract) GetFrozenDuration() int64 {
	if m != nil {
		return m.FrozenDuration
	}
	return 0
}

func (m *FreezeBalanceContract) GetResource() ResourceCode {
	if m != nil {
		return m.Resource
	}
	return ResourceCode_BANDWIDTH
}

func (m *FreezeBalanceContract) GetReceiverAddress() []byte {
	if m != nil {
		return m.ReceiverAddress
	}
	return nil
}

type UnfreezeBalanceContract struct {
	OwnerAddress         []byte       `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Resource             ResourceCode `protobuf:"varint,10,opt,name=resource,proto3,enum=protocol.ResourceCode" json:"resource,omitempty"`
	ReceiverAddress      []byte       `protobuf:"bytes,15,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UnfreezeBalanceContract) Reset()         { *m = UnfreezeBalanceContract{} }
func (m *UnfreezeBalanceContract) String() string { return proto.CompactTextString(m) }
func (*UnfreezeBalanceContract) ProtoMessage()    {}
func (*UnfreezeBalanceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{1}
}

func (m *UnfreezeBalanceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeBalanceContract.Unmarshal(m, b)
}
func (m *UnfreezeBalanceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeBalanceContract.Marshal(b, m, deterministic)
}
func (m *UnfreezeBalanceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeBalanceContract.Merge(m, src)
}
func (m *UnfreezeBalanceContract) XXX_Size() int {
	return xxx_messageInfo_UnfreezeBalanceContract.Size(m)
}
func (m *UnfreezeBalanceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeBalanceContract.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeBalanceContract proto.InternalMessageInfo

func (m *UnfreezeBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *UnfreezeBalanceContract) GetResource() ResourceCode {
	if m != nil {
		return m.Resource
	}
	return ResourceCode_BANDWIDTH
}

func (m *UnfreezeBalanceContract) GetReceiverAddress() []byte {
	if m != nil {
		return m.ReceiverAddress
	}
	return nil
}

type WithdrawBalanceContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawBalanceContract) Reset()         { *m = WithdrawBalanceContract{} }
func (m *WithdrawBalanceContract) String() string { return proto.CompactTextString(m) }
func (*WithdrawBalanceContract) ProtoMessage()    {}
func (*WithdrawBalanceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{2}
}

func (m *WithdrawBalanceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawBalanceContract.Unmarshal(m, b)
}
func (m *WithdrawBalanceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawBalanceContract.Marshal(b, m, deterministic)
}
func (m *WithdrawBalanceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawBalanceContract.Merge(m, src)
}
func (m *WithdrawBalanceContract) XXX_Size() int {
	return xxx_messageInfo_WithdrawBalanceContract.Size(m)
}
func (m *WithdrawBalanceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawBalanceContract.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawBalanceContract proto.InternalMessageInfo

func (m *WithdrawBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type TransferContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress            []byte   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferContract) Reset()         { *m = TransferContract{} }
func (m *TransferContract) String() string { return proto.CompactTextString(m) }
func (*TransferContract) ProtoMessage()    {}
func (*TransferContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{3}
}

func (m *TransferContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferContract.Unmarshal(m, b)
}
func (m *TransferContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferContract.Marshal(b, m, deterministic)
}
func (m *TransferContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferContract.Merge(m, src)
}
func (m *TransferContract) XXX_Size() int {
	return xxx_messageInfo_TransferContract.Size(m)
}
func (m *TransferContract) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferContract.DiscardUnknown(m)
}

var xxx_messageInfo_TransferContract proto.InternalMessageInfo

func (m *TransferContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*FreezeBalanceContract)(nil), "protocol.FreezeBalanceContract")
	proto.RegisterType((*UnfreezeBalanceContract)(nil), "protocol.UnfreezeBalanceContract")
	proto.RegisterType((*WithdrawBalanceContract)(nil), "protocol.WithdrawBalanceContract")
	proto.RegisterType((*TransferContract)(nil), "protocol.TransferContract")
}

func init() {
	proto.RegisterFile("core/contract/balance_contract.proto", fileDescriptor_54daa91fef75922f)
}

var fileDescriptor_54daa91fef75922f = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xdf, 0x6a, 0xea, 0x40,
	0x10, 0x87, 0x89, 0x82, 0x78, 0x06, 0xff, 0x11, 0x38, 0x1a, 0x84, 0x82, 0xd8, 0x16, 0xed, 0x85,
	0x89, 0xb5, 0xf7, 0x85, 0x6a, 0xe9, 0x03, 0x84, 0x96, 0x42, 0x6f, 0x64, 0xdd, 0x4c, 0x62, 0xc0,
	0xec, 0xc8, 0x64, 0x53, 0x5b, 0xdf, 0xa5, 0x4f, 0xd8, 0x97, 0x28, 0x26, 0x1b, 0x69, 0x2f, 0x2b,
	0xf4, 0x2a, 0xe4, 0x37, 0xdf, 0x0c, 0xdf, 0xb0, 0x03, 0x17, 0x92, 0x18, 0x3d, 0x49, 0x4a, 0xb3,
	0x90, 0xda, 0x5b, 0x89, 0x8d, 0x50, 0x12, 0x97, 0x65, 0xe0, 0x6e, 0x99, 0x34, 0xd9, 0xf5, 0xfc,
	0x23, 0x69, 0xd3, 0xef, 0xff, 0xe4, 0x25, 0x25, 0x09, 0xa9, 0x82, 0x1a, 0x7e, 0x5a, 0xf0, 0xff,
	0x81, 0x11, 0xf7, 0x38, 0x2f, 0xc6, 0x2c, 0x0c, 0x66, 0x9f, 0x43, 0x93, 0x76, 0x0a, 0x79, 0x29,
	0x82, 0x80, 0x31, 0x4d, 0x1d, 0x6b, 0x60, 0x8d, 0x1b, 0x7e, 0x23, 0x0f, 0xef, 0x8a, 0xcc, 0xbe,
	0x84, 0x56, 0xc8, 0xb4, 0x47, 0xb5, 0x34, 0x16, 0x4e, 0x65, 0x60, 0x8d, 0xab, 0x7e, 0xb3, 0x48,
	0xcd, 0x4c, 0x7b, 0x04, 0x6d, 0x83, 0x05, 0x19, 0x0b, 0x1d, 0x93, 0x72, 0xaa, 0x39, 0x67, 0xba,
	0xef, 0x4d, 0x6a, 0xcf, 0xa0, 0xce, 0x98, 0x52, 0xc6, 0x12, 0x1d, 0x18, 0x58, 0xe3, 0xd6, 0xac,
	0xeb, 0x96, 0x7b, 0xb8, 0xbe, 0xa9, 0x2c, 0x28, 0x40, 0xff, 0xc8, 0xd9, 0x57, 0xd0, 0x61, 0x94,
	0x18, 0xbf, 0x7e, 0x73, 0x6d, 0xe7, 0xae, 0xed, 0x32, 0x37, 0xba, 0xc3, 0x0f, 0x0b, 0x7a, 0x4f,
	0x2a, 0x3c, 0x7d, 0xdf, 0x3f, 0xf6, 0xbb, 0x85, 0xde, 0x73, 0xac, 0xd7, 0x01, 0x8b, 0xdd, 0x29,
	0x7a, 0x43, 0x05, 0x9d, 0x47, 0x16, 0x2a, 0x0d, 0x91, 0x7f, 0xb7, 0xd7, 0x19, 0x80, 0xa6, 0x23,
	0x51, 0xc9, 0x89, 0x7f, 0x9a, 0xca, 0x72, 0x17, 0x6a, 0x22, 0xa1, 0x4c, 0x69, 0xf3, 0x6c, 0xe6,
	0x6f, 0xbe, 0x00, 0x87, 0x38, 0x72, 0x35, 0x97, 0xf7, 0x94, 0xba, 0xe5, 0x99, 0xbd, 0x8c, 0xa2,
	0x58, 0xaf, 0xb3, 0x95, 0x2b, 0x29, 0xf1, 0xa6, 0x6f, 0xd7, 0xd3, 0xd0, 0x3b, 0x60, 0x93, 0x88,
	0xb7, 0x72, 0x12, 0x09, 0x8d, 0x3b, 0xf1, 0xee, 0x1d, 0x0e, 0x73, 0x55, 0xcb, 0x3b, 0x6f, 0xbe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0x97, 0x74, 0xcf, 0xd7, 0x02, 0x00, 0x00,
}
