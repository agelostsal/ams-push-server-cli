// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ams.proto

package ams

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Wrapper for subscription
type DeactivateSubscriptionResponse struct {
	// Message response
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeactivateSubscriptionResponse) Reset()         { *m = DeactivateSubscriptionResponse{} }
func (m *DeactivateSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*DeactivateSubscriptionResponse) ProtoMessage()    {}
func (*DeactivateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{0}
}

func (m *DeactivateSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeactivateSubscriptionResponse.Unmarshal(m, b)
}
func (m *DeactivateSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeactivateSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *DeactivateSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeactivateSubscriptionResponse.Merge(m, src)
}
func (m *DeactivateSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_DeactivateSubscriptionResponse.Size(m)
}
func (m *DeactivateSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeactivateSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeactivateSubscriptionResponse proto.InternalMessageInfo

func (m *DeactivateSubscriptionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Contains which subscription to deactivate
type DeactivateSubscriptionRequest struct {
	// Required. The full resource name of the subscrption.
	FullName             string   `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeactivateSubscriptionRequest) Reset()         { *m = DeactivateSubscriptionRequest{} }
func (m *DeactivateSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*DeactivateSubscriptionRequest) ProtoMessage()    {}
func (*DeactivateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{1}
}

func (m *DeactivateSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeactivateSubscriptionRequest.Unmarshal(m, b)
}
func (m *DeactivateSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeactivateSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *DeactivateSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeactivateSubscriptionRequest.Merge(m, src)
}
func (m *DeactivateSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_DeactivateSubscriptionRequest.Size(m)
}
func (m *DeactivateSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeactivateSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeactivateSubscriptionRequest proto.InternalMessageInfo

func (m *DeactivateSubscriptionRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

// Wrapper for subscription
type ActivateSubscriptionResponse struct {
	// Message response
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateSubscriptionResponse) Reset()         { *m = ActivateSubscriptionResponse{} }
func (m *ActivateSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*ActivateSubscriptionResponse) ProtoMessage()    {}
func (*ActivateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{2}
}

func (m *ActivateSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateSubscriptionResponse.Unmarshal(m, b)
}
func (m *ActivateSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *ActivateSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateSubscriptionResponse.Merge(m, src)
}
func (m *ActivateSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_ActivateSubscriptionResponse.Size(m)
}
func (m *ActivateSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateSubscriptionResponse proto.InternalMessageInfo

func (m *ActivateSubscriptionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Wrapper for subscription.
type ActivateSubscriptionRequest struct {
	// Required. A subscription.
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=Subscription,proto3" json:"Subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ActivateSubscriptionRequest) Reset()         { *m = ActivateSubscriptionRequest{} }
func (m *ActivateSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateSubscriptionRequest) ProtoMessage()    {}
func (*ActivateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{3}
}

func (m *ActivateSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateSubscriptionRequest.Unmarshal(m, b)
}
func (m *ActivateSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *ActivateSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateSubscriptionRequest.Merge(m, src)
}
func (m *ActivateSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_ActivateSubscriptionRequest.Size(m)
}
func (m *ActivateSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateSubscriptionRequest proto.InternalMessageInfo

func (m *ActivateSubscriptionRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

// Subscription holds informaton related to how the push functionality should operate.
type Subscription struct {
	// Required. The full resource name of the subscription.
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// The full topic name of the topic that the subscription is associated.
	FullTopic string `protobuf:"bytes,2,opt,name=full_topic,json=fullTopic,proto3" json:"full_topic,omitempty"`
	// Required. Information regarding the push functionality.
	PushConfig           *PushConfig `protobuf:"bytes,4,opt,name=PushConfig,proto3" json:"PushConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{4}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Subscription) GetFullTopic() string {
	if m != nil {
		return m.FullTopic
	}
	return ""
}

func (m *Subscription) GetPushConfig() *PushConfig {
	if m != nil {
		return m.PushConfig
	}
	return nil
}

// PushConfig holds information on how a push subscription functions.
type PushConfig struct {
	// Required. An https endpoint to where the messages will be pushed.
	PushEndpoint string `protobuf:"bytes,1,opt,name=push_endpoint,json=pushEndpoint,proto3" json:"push_endpoint,omitempty"`
	// Required. Retry policy.
	RetryPolicy          *RetryPolicy `protobuf:"bytes,2,opt,name=RetryPolicy,proto3" json:"RetryPolicy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PushConfig) Reset()         { *m = PushConfig{} }
func (m *PushConfig) String() string { return proto.CompactTextString(m) }
func (*PushConfig) ProtoMessage()    {}
func (*PushConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{5}
}

func (m *PushConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushConfig.Unmarshal(m, b)
}
func (m *PushConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushConfig.Marshal(b, m, deterministic)
}
func (m *PushConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushConfig.Merge(m, src)
}
func (m *PushConfig) XXX_Size() int {
	return xxx_messageInfo_PushConfig.Size(m)
}
func (m *PushConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PushConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PushConfig proto.InternalMessageInfo

func (m *PushConfig) GetPushEndpoint() string {
	if m != nil {
		return m.PushEndpoint
	}
	return ""
}

func (m *PushConfig) GetRetryPolicy() *RetryPolicy {
	if m != nil {
		return m.RetryPolicy
	}
	return nil
}

// RetryPolicy holds information regarding the retry policy.
type RetryPolicy struct {
	// Required. Type of the retry policy used (Only linear policy supported).
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Required. Retry period in milliseconds.
	Period               uint32   `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryPolicy) Reset()         { *m = RetryPolicy{} }
func (m *RetryPolicy) String() string { return proto.CompactTextString(m) }
func (*RetryPolicy) ProtoMessage()    {}
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e4db6795b5b1aa, []int{6}
}

func (m *RetryPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryPolicy.Unmarshal(m, b)
}
func (m *RetryPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryPolicy.Marshal(b, m, deterministic)
}
func (m *RetryPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryPolicy.Merge(m, src)
}
func (m *RetryPolicy) XXX_Size() int {
	return xxx_messageInfo_RetryPolicy.Size(m)
}
func (m *RetryPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RetryPolicy proto.InternalMessageInfo

func (m *RetryPolicy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RetryPolicy) GetPeriod() uint32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func init() {
	proto.RegisterType((*DeactivateSubscriptionResponse)(nil), "DeactivateSubscriptionResponse")
	proto.RegisterType((*DeactivateSubscriptionRequest)(nil), "DeactivateSubscriptionRequest")
	proto.RegisterType((*ActivateSubscriptionResponse)(nil), "ActivateSubscriptionResponse")
	proto.RegisterType((*ActivateSubscriptionRequest)(nil), "ActivateSubscriptionRequest")
	proto.RegisterType((*Subscription)(nil), "Subscription")
	proto.RegisterType((*PushConfig)(nil), "PushConfig")
	proto.RegisterType((*RetryPolicy)(nil), "RetryPolicy")
}

func init() { proto.RegisterFile("ams.proto", fileDescriptor_85e4db6795b5b1aa) }

var fileDescriptor_85e4db6795b5b1aa = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x6d, 0x7f, 0x94, 0xfe, 0xec, 0xa4, 0xbd, 0x0c, 0x52, 0x42, 0xff, 0x29, 0xeb, 0x45, 0x10,
	0x16, 0xac, 0x17, 0x15, 0x2f, 0x45, 0xbd, 0x4a, 0x49, 0xf5, 0xe4, 0xa1, 0x6c, 0xd3, 0x69, 0xbb,
	0xd0, 0x64, 0xd7, 0xec, 0xa6, 0xd2, 0x8f, 0xe7, 0x37, 0x93, 0x2c, 0x0d, 0x26, 0x50, 0x73, 0xf0,
	0xb6, 0xf3, 0xe6, 0xbd, 0x7d, 0xb3, 0x3b, 0x0f, 0x5a, 0x22, 0x32, 0x5c, 0x27, 0xca, 0x2a, 0x76,
	0x0f, 0xa3, 0x27, 0x12, 0xa1, 0x95, 0x3b, 0x61, 0x69, 0x96, 0x2e, 0x4c, 0x98, 0x48, 0x6d, 0xa5,
	0x8a, 0x03, 0x32, 0x5a, 0xc5, 0x86, 0xd0, 0x87, 0xff, 0x11, 0x19, 0x23, 0xd6, 0xe4, 0xd7, 0xcf,
	0xeb, 0x97, 0xad, 0x20, 0x2f, 0xd9, 0x03, 0x0c, 0x7f, 0xd3, 0x7e, 0xa4, 0x64, 0x2c, 0xf6, 0xa1,
	0xb5, 0x4a, 0xb7, 0xdb, 0x79, 0x2c, 0xa2, 0x5c, 0x7c, 0x92, 0x01, 0x2f, 0x22, 0x22, 0x76, 0x0b,
	0x83, 0xc9, 0xdf, 0x7c, 0xa7, 0xd0, 0x9f, 0x54, 0xb8, 0x5e, 0x43, 0xbb, 0x08, 0x3b, 0xb5, 0x37,
	0xee, 0xf0, 0x12, 0xb7, 0x44, 0x61, 0x9f, 0x65, 0x49, 0xe5, 0xe0, 0x38, 0x04, 0x70, 0x4d, 0xab,
	0xb4, 0x0c, 0xfd, 0x7f, 0xae, 0xeb, 0xe8, 0xaf, 0x19, 0x80, 0x57, 0x00, 0xd3, 0xd4, 0x6c, 0x1e,
	0x55, 0xbc, 0x92, 0x6b, 0xbf, 0xe1, 0xcc, 0x3d, 0xfe, 0x03, 0x05, 0x85, 0x36, 0x13, 0x45, 0x32,
	0x5e, 0x40, 0x47, 0xa7, 0x66, 0x33, 0xa7, 0x78, 0xa9, 0x95, 0x8c, 0xed, 0xc1, 0xba, 0x9d, 0x81,
	0xcf, 0x07, 0x0c, 0x39, 0x78, 0x01, 0xd9, 0x64, 0x3f, 0x55, 0x5b, 0x19, 0xee, 0x9d, 0xbf, 0x37,
	0x6e, 0xf3, 0x02, 0x16, 0x14, 0x09, 0xec, 0xae, 0xc4, 0x47, 0x84, 0x86, 0xdd, 0xeb, 0xfc, 0x55,
	0xee, 0x8c, 0x5d, 0x68, 0x6a, 0x4a, 0xa4, 0x5a, 0xba, 0xdb, 0x3a, 0xc1, 0xa1, 0x1a, 0x7f, 0xd5,
	0xc1, 0xcb, 0xc6, 0x9b, 0x51, 0xb2, 0x93, 0x21, 0xe1, 0x1b, 0x9c, 0x1e, 0xfb, 0x78, 0x1c, 0xf0,
	0x8a, 0x7d, 0xf4, 0x86, 0xbc, 0x6a, 0xcf, 0xac, 0x86, 0xef, 0xd0, 0x3d, 0x9e, 0x23, 0x1c, 0xf1,
	0xca, 0x80, 0xf5, 0xce, 0x78, 0x75, 0x78, 0x59, 0x6d, 0xd1, 0x74, 0x39, 0xbf, 0xf9, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xac, 0xfc, 0xb5, 0xc5, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushServiceClient interface {
	// Activates a subscription in order for the service to start handling the push functionality
	ActivateSubscription(ctx context.Context, in *ActivateSubscriptionRequest, opts ...grpc.CallOption) (*ActivateSubscriptionResponse, error)
	//  Deactivates a subscription in order for the service to stop handling the push functionality
	DeactivateSubscription(ctx context.Context, in *DeactivateSubscriptionRequest, opts ...grpc.CallOption) (*DeactivateSubscriptionResponse, error)
}

type pushServiceClient struct {
	cc *grpc.ClientConn
}

func NewPushServiceClient(cc *grpc.ClientConn) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) ActivateSubscription(ctx context.Context, in *ActivateSubscriptionRequest, opts ...grpc.CallOption) (*ActivateSubscriptionResponse, error) {
	out := new(ActivateSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/PushService/ActivateSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushServiceClient) DeactivateSubscription(ctx context.Context, in *DeactivateSubscriptionRequest, opts ...grpc.CallOption) (*DeactivateSubscriptionResponse, error) {
	out := new(DeactivateSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/PushService/DeactivateSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServiceServer is the server API for PushService service.
type PushServiceServer interface {
	// Activates a subscription in order for the service to start handling the push functionality
	ActivateSubscription(context.Context, *ActivateSubscriptionRequest) (*ActivateSubscriptionResponse, error)
	//  Deactivates a subscription in order for the service to stop handling the push functionality
	DeactivateSubscription(context.Context, *DeactivateSubscriptionRequest) (*DeactivateSubscriptionResponse, error)
}

func RegisterPushServiceServer(s *grpc.Server, srv PushServiceServer) {
	s.RegisterService(&_PushService_serviceDesc, srv)
}

func _PushService_ActivateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).ActivateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PushService/ActivateSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).ActivateSubscription(ctx, req.(*ActivateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushService_DeactivateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).DeactivateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PushService/DeactivateSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).DeactivateSubscription(ctx, req.(*DeactivateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PushService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivateSubscription",
			Handler:    _PushService_ActivateSubscription_Handler,
		},
		{
			MethodName: "DeactivateSubscription",
			Handler:    _PushService_DeactivateSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ams.proto",
}
