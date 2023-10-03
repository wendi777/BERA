// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polaris/evm/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Status represents the status of a transaction.
type Status int32

const (
	// STATUS_REVERT_UNSPECIFIED indicates that the transaction reverted.
	Status_STATUS_REVERT_UNSPECIFIED Status = 0
	// STATUS_SUCCESS indicates that the transaction completed successfully.
	Status_STATUS_SUCCESS Status = 1
	// STATUS_NOT_INCLUDED indicates that the transaction was not included in the
	// `evm` block, due to an critical error.
	Status_STATUS_NOT_INCLUDED Status = 2
)

var Status_name = map[int32]string{
	0: "STATUS_REVERT_UNSPECIFIED",
	1: "STATUS_SUCCESS",
	2: "STATUS_NOT_INCLUDED",
}

var Status_value = map[string]int32{
	"STATUS_REVERT_UNSPECIFIED": 0,
	"STATUS_SUCCESS":            1,
	"STATUS_NOT_INCLUDED":       2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{0}
}

// WrappedEthereumTransaction encapsulates an Ethereum transaction as an SDK message.
type WrappedEthereumTransaction struct {
	// data is inner transaction data of the Ethereum transaction.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WrappedEthereumTransaction) Reset()         { *m = WrappedEthereumTransaction{} }
func (m *WrappedEthereumTransaction) String() string { return proto.CompactTextString(m) }
func (*WrappedEthereumTransaction) ProtoMessage()    {}
func (*WrappedEthereumTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{0}
}
func (m *WrappedEthereumTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedEthereumTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedEthereumTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedEthereumTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedEthereumTransaction.Merge(m, src)
}
func (m *WrappedEthereumTransaction) XXX_Size() int {
	return m.Size()
}
func (m *WrappedEthereumTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedEthereumTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedEthereumTransaction proto.InternalMessageInfo

func (m *WrappedEthereumTransaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// WrappedPayloadEnvelope encapsulates an Ethereum transaction as an SDK message.
type WrappedPayloadEnvelope struct {
	// data is inner transaction data of the Ethereum transaction.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WrappedPayloadEnvelope) Reset()         { *m = WrappedPayloadEnvelope{} }
func (m *WrappedPayloadEnvelope) String() string { return proto.CompactTextString(m) }
func (*WrappedPayloadEnvelope) ProtoMessage()    {}
func (*WrappedPayloadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{1}
}
func (m *WrappedPayloadEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedPayloadEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedPayloadEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedPayloadEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPayloadEnvelope.Merge(m, src)
}
func (m *WrappedPayloadEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *WrappedPayloadEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPayloadEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPayloadEnvelope proto.InternalMessageInfo

func (m *WrappedPayloadEnvelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WrappedPayloadEnvelopeResponse struct {
}

func (m *WrappedPayloadEnvelopeResponse) Reset()         { *m = WrappedPayloadEnvelopeResponse{} }
func (m *WrappedPayloadEnvelopeResponse) String() string { return proto.CompactTextString(m) }
func (*WrappedPayloadEnvelopeResponse) ProtoMessage()    {}
func (*WrappedPayloadEnvelopeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{2}
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedPayloadEnvelopeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPayloadEnvelopeResponse.Merge(m, src)
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Size() int {
	return m.Size()
}
func (m *WrappedPayloadEnvelopeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPayloadEnvelopeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPayloadEnvelopeResponse proto.InternalMessageInfo

// WrappedEthereumTransactionResult defines the Msg/EthereumTx response type.
type WrappedEthereumTransactionResult struct {
	// `status` represents a transaction's status
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=polaris.evm.v1alpha1.Status" json:"status,omitempty"`
}

func (m *WrappedEthereumTransactionResult) Reset()         { *m = WrappedEthereumTransactionResult{} }
func (m *WrappedEthereumTransactionResult) String() string { return proto.CompactTextString(m) }
func (*WrappedEthereumTransactionResult) ProtoMessage()    {}
func (*WrappedEthereumTransactionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{3}
}
func (m *WrappedEthereumTransactionResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedEthereumTransactionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedEthereumTransactionResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedEthereumTransactionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedEthereumTransactionResult.Merge(m, src)
}
func (m *WrappedEthereumTransactionResult) XXX_Size() int {
	return m.Size()
}
func (m *WrappedEthereumTransactionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedEthereumTransactionResult.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedEthereumTransactionResult proto.InternalMessageInfo

func (m *WrappedEthereumTransactionResult) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_REVERT_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("polaris.evm.v1alpha1.Status", Status_name, Status_value)
	proto.RegisterType((*WrappedEthereumTransaction)(nil), "polaris.evm.v1alpha1.WrappedEthereumTransaction")
	proto.RegisterType((*WrappedPayloadEnvelope)(nil), "polaris.evm.v1alpha1.WrappedPayloadEnvelope")
	proto.RegisterType((*WrappedPayloadEnvelopeResponse)(nil), "polaris.evm.v1alpha1.WrappedPayloadEnvelopeResponse")
	proto.RegisterType((*WrappedEthereumTransactionResult)(nil), "polaris.evm.v1alpha1.WrappedEthereumTransactionResult")
}

func init() { proto.RegisterFile("polaris/evm/v1alpha1/tx.proto", fileDescriptor_d8b33d2a2c64400f) }

var fileDescriptor_d8b33d2a2c64400f = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x93, 0x45, 0x7b, 0x18, 0xa4, 0x94, 0x51, 0x76, 0x35, 0xb8, 0xa1, 0xf4, 0x24, 0x4b,
	0x49, 0xb6, 0xeb, 0xe2, 0x5d, 0xd3, 0x59, 0x28, 0x68, 0x2d, 0x99, 0x44, 0xc5, 0x4b, 0x99, 0x4d,
	0x5e, 0x9a, 0x60, 0x92, 0x19, 0x66, 0xa6, 0xa1, 0x15, 0x0f, 0x7e, 0x01, 0xc1, 0x8f, 0xe2, 0xc7,
	0xf0, 0xd8, 0xa3, 0x47, 0x69, 0x0f, 0x7e, 0x0d, 0x31, 0x4d, 0x41, 0x4a, 0xba, 0xd0, 0xd3, 0x0c,
	0xef, 0xf3, 0x7b, 0x5f, 0xde, 0x3f, 0x0f, 0x3a, 0x17, 0x3c, 0x63, 0x32, 0x55, 0x2e, 0x94, 0xb9,
	0x5b, 0x0e, 0x58, 0x26, 0x12, 0x36, 0x70, 0xf5, 0xc2, 0x11, 0x92, 0x6b, 0x8e, 0x1f, 0xd5, 0xb2,
	0x03, 0x65, 0xee, 0xec, 0x64, 0xeb, 0x2c, 0xe2, 0x2a, 0xe7, 0xca, 0xcd, 0xd5, 0xcc, 0x2d, 0x07,
	0xff, 0x9e, 0x2d, 0xde, 0xbb, 0x44, 0xd6, 0x7b, 0xc9, 0x84, 0x80, 0x98, 0xe8, 0x04, 0x24, 0xcc,
	0xf3, 0x40, 0xb2, 0x42, 0xb1, 0x48, 0xa7, 0xbc, 0xc0, 0x18, 0xdd, 0x8b, 0x99, 0x66, 0x8f, 0xcd,
	0xae, 0xf9, 0xec, 0x81, 0x5f, 0xfd, 0x7b, 0x7d, 0x74, 0x5a, 0x67, 0x4c, 0xd8, 0x32, 0xe3, 0x2c,
	0x26, 0x45, 0x09, 0x19, 0x17, 0xd0, 0x48, 0x77, 0x91, 0xdd, 0x4c, 0xfb, 0xa0, 0x04, 0x2f, 0x14,
	0xf4, 0x3e, 0xa0, 0xee, 0xe1, 0x0e, 0x7c, 0x50, 0xf3, 0x4c, 0xe3, 0x6b, 0xd4, 0x52, 0x9a, 0xe9,
	0xb9, 0xaa, 0x6a, 0xb7, 0xaf, 0x9e, 0x3a, 0x4d, 0x53, 0x3a, 0xb4, 0x62, 0xfc, 0x9a, 0xbd, 0x08,
	0x50, 0x6b, 0x1b, 0xc1, 0xe7, 0xe8, 0x09, 0x0d, 0x5e, 0x06, 0x21, 0x9d, 0xfa, 0xe4, 0x1d, 0xf1,
	0x83, 0x69, 0x38, 0xa6, 0x13, 0xe2, 0x8d, 0x6e, 0x46, 0x64, 0xd8, 0x31, 0x30, 0x46, 0xed, 0x5a,
	0xa6, 0xa1, 0xe7, 0x11, 0x4a, 0x3b, 0x26, 0x3e, 0x43, 0x0f, 0xeb, 0xd8, 0xf8, 0x6d, 0x30, 0x1d,
	0x8d, 0xbd, 0xd7, 0xe1, 0x90, 0x0c, 0x3b, 0x27, 0x57, 0xdf, 0x4e, 0x10, 0x7a, 0xa3, 0x66, 0x14,
	0x64, 0x99, 0x46, 0x80, 0x3f, 0xa3, 0x36, 0xd1, 0xc9, 0xff, 0x4b, 0xbb, 0x6c, 0x6e, 0xee, 0xf0,
	0x90, 0xd6, 0x8b, 0x63, 0x33, 0xea, 0xb5, 0x7c, 0x41, 0xa7, 0x13, 0xc9, 0x23, 0x50, 0x6a, 0xff,
	0x14, 0xfd, 0x3b, 0x2b, 0xee, 0xd1, 0xd6, 0xf5, 0x31, 0xf4, 0xee, 0x70, 0xd6, 0xfd, 0xaf, 0x7f,
	0x7e, 0x5c, 0x98, 0xaf, 0x6e, 0x7e, 0xae, 0x6d, 0x73, 0xb5, 0xb6, 0xcd, 0xdf, 0x6b, 0xdb, 0xfc,
	0xbe, 0xb1, 0x8d, 0xd5, 0xc6, 0x36, 0x7e, 0x6d, 0x6c, 0xe3, 0x63, 0x5f, 0x7c, 0x9a, 0x39, 0xb7,
	0x20, 0x59, 0x94, 0xb0, 0xb4, 0x70, 0x62, 0x28, 0xdd, 0x9d, 0x77, 0x6b, 0x3b, 0x2e, 0x2a, 0x13,
	0xeb, 0xa5, 0x00, 0x75, 0xdb, 0xaa, 0x0c, 0xf9, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49,
	0x1a, 0xf4, 0xc3, 0xe0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// EthTransaction defines a method submitting Ethereum transactions.
	EthTransaction(ctx context.Context, in *WrappedEthereumTransaction, opts ...grpc.CallOption) (*WrappedEthereumTransactionResult, error)
	// ProcessPayloadEnvelope defines a method to process CL paylods.
	ProcessPayloadEnvelope(ctx context.Context, in *WrappedPayloadEnvelope, opts ...grpc.CallOption) (*WrappedPayloadEnvelopeResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) EthTransaction(ctx context.Context, in *WrappedEthereumTransaction, opts ...grpc.CallOption) (*WrappedEthereumTransactionResult, error) {
	out := new(WrappedEthereumTransactionResult)
	err := c.cc.Invoke(ctx, "/polaris.evm.v1alpha1.MsgService/EthTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ProcessPayloadEnvelope(ctx context.Context, in *WrappedPayloadEnvelope, opts ...grpc.CallOption) (*WrappedPayloadEnvelopeResponse, error) {
	out := new(WrappedPayloadEnvelopeResponse)
	err := c.cc.Invoke(ctx, "/polaris.evm.v1alpha1.MsgService/ProcessPayloadEnvelope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// EthTransaction defines a method submitting Ethereum transactions.
	EthTransaction(context.Context, *WrappedEthereumTransaction) (*WrappedEthereumTransactionResult, error)
	// ProcessPayloadEnvelope defines a method to process CL paylods.
	ProcessPayloadEnvelope(context.Context, *WrappedPayloadEnvelope) (*WrappedPayloadEnvelopeResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) EthTransaction(ctx context.Context, req *WrappedEthereumTransaction) (*WrappedEthereumTransactionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthTransaction not implemented")
}
func (*UnimplementedMsgServiceServer) ProcessPayloadEnvelope(ctx context.Context, req *WrappedPayloadEnvelope) (*WrappedPayloadEnvelopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayloadEnvelope not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_EthTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedEthereumTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).EthTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polaris.evm.v1alpha1.MsgService/EthTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).EthTransaction(ctx, req.(*WrappedEthereumTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ProcessPayloadEnvelope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedPayloadEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessPayloadEnvelope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polaris.evm.v1alpha1.MsgService/ProcessPayloadEnvelope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessPayloadEnvelope(ctx, req.(*WrappedPayloadEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.evm.v1alpha1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthTransaction",
			Handler:    _MsgService_EthTransaction_Handler,
		},
		{
			MethodName: "ProcessPayloadEnvelope",
			Handler:    _MsgService_ProcessPayloadEnvelope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "polaris/evm/v1alpha1/tx.proto",
}

func (m *WrappedEthereumTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedEthereumTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedEthereumTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WrappedPayloadEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedPayloadEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedPayloadEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WrappedPayloadEnvelopeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedPayloadEnvelopeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedPayloadEnvelopeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *WrappedEthereumTransactionResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedEthereumTransactionResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedEthereumTransactionResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WrappedEthereumTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *WrappedPayloadEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *WrappedPayloadEnvelopeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *WrappedEthereumTransactionResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovTx(uint64(m.Status))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WrappedEthereumTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WrappedEthereumTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedEthereumTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WrappedPayloadEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WrappedPayloadEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedPayloadEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WrappedPayloadEnvelopeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WrappedPayloadEnvelopeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedPayloadEnvelopeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WrappedEthereumTransactionResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WrappedEthereumTransactionResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedEthereumTransactionResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
