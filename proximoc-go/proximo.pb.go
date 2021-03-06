// Code generated by protoc-gen-go.
// source: proximo.proto
// DO NOT EDIT!

/*
Package proximoc is a generated protocol buffer package.

It is generated from these files:
	proximo.proto

It has these top-level messages:
	Message
	ConsumerRequest
	StartConsumeRequest
	Confirmation
	PublisherRequest
	StartPublishRequest
*/
package proximoc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Message struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ConsumerRequest struct {
	// expected if this is a start request
	StartRequest *StartConsumeRequest `protobuf:"bytes,2,opt,name=startRequest" json:"startRequest,omitempty"`
	// expected if this is a confirmation
	Confirmation *Confirmation `protobuf:"bytes,3,opt,name=confirmation" json:"confirmation,omitempty"`
}

func (m *ConsumerRequest) Reset()                    { *m = ConsumerRequest{} }
func (m *ConsumerRequest) String() string            { return proto.CompactTextString(m) }
func (*ConsumerRequest) ProtoMessage()               {}
func (*ConsumerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConsumerRequest) GetStartRequest() *StartConsumeRequest {
	if m != nil {
		return m.StartRequest
	}
	return nil
}

func (m *ConsumerRequest) GetConfirmation() *Confirmation {
	if m != nil {
		return m.Confirmation
	}
	return nil
}

type StartConsumeRequest struct {
	Topic    string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Consumer string `protobuf:"bytes,2,opt,name=consumer" json:"consumer,omitempty"`
}

func (m *StartConsumeRequest) Reset()                    { *m = StartConsumeRequest{} }
func (m *StartConsumeRequest) String() string            { return proto.CompactTextString(m) }
func (*StartConsumeRequest) ProtoMessage()               {}
func (*StartConsumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StartConsumeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *StartConsumeRequest) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

type Confirmation struct {
	MsgID string `protobuf:"bytes,1,opt,name=msgID" json:"msgID,omitempty"`
}

func (m *Confirmation) Reset()                    { *m = Confirmation{} }
func (m *Confirmation) String() string            { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()               {}
func (*Confirmation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Confirmation) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

type PublisherRequest struct {
	// expected if this is a start request
	StartRequest *StartPublishRequest `protobuf:"bytes,2,opt,name=startRequest" json:"startRequest,omitempty"`
	// expected if this is a message
	Msg *Message `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *PublisherRequest) Reset()                    { *m = PublisherRequest{} }
func (m *PublisherRequest) String() string            { return proto.CompactTextString(m) }
func (*PublisherRequest) ProtoMessage()               {}
func (*PublisherRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublisherRequest) GetStartRequest() *StartPublishRequest {
	if m != nil {
		return m.StartRequest
	}
	return nil
}

func (m *PublisherRequest) GetMsg() *Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

type StartPublishRequest struct {
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *StartPublishRequest) Reset()                    { *m = StartPublishRequest{} }
func (m *StartPublishRequest) String() string            { return proto.CompactTextString(m) }
func (*StartPublishRequest) ProtoMessage()               {}
func (*StartPublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StartPublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "proximo.Message")
	proto.RegisterType((*ConsumerRequest)(nil), "proximo.ConsumerRequest")
	proto.RegisterType((*StartConsumeRequest)(nil), "proximo.StartConsumeRequest")
	proto.RegisterType((*Confirmation)(nil), "proximo.Confirmation")
	proto.RegisterType((*PublisherRequest)(nil), "proximo.PublisherRequest")
	proto.RegisterType((*StartPublishRequest)(nil), "proximo.StartPublishRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageSource service

type MessageSourceClient interface {
	Consume(ctx context.Context, opts ...grpc.CallOption) (MessageSource_ConsumeClient, error)
}

type messageSourceClient struct {
	cc *grpc.ClientConn
}

func NewMessageSourceClient(cc *grpc.ClientConn) MessageSourceClient {
	return &messageSourceClient{cc}
}

func (c *messageSourceClient) Consume(ctx context.Context, opts ...grpc.CallOption) (MessageSource_ConsumeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageSource_serviceDesc.Streams[0], c.cc, "/proximo.MessageSource/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageSourceConsumeClient{stream}
	return x, nil
}

type MessageSource_ConsumeClient interface {
	Send(*ConsumerRequest) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageSourceConsumeClient struct {
	grpc.ClientStream
}

func (x *messageSourceConsumeClient) Send(m *ConsumerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageSourceConsumeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MessageSource service

type MessageSourceServer interface {
	Consume(MessageSource_ConsumeServer) error
}

func RegisterMessageSourceServer(s *grpc.Server, srv MessageSourceServer) {
	s.RegisterService(&_MessageSource_serviceDesc, srv)
}

func _MessageSource_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageSourceServer).Consume(&messageSourceConsumeServer{stream})
}

type MessageSource_ConsumeServer interface {
	Send(*Message) error
	Recv() (*ConsumerRequest, error)
	grpc.ServerStream
}

type messageSourceConsumeServer struct {
	grpc.ServerStream
}

func (x *messageSourceConsumeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageSourceConsumeServer) Recv() (*ConsumerRequest, error) {
	m := new(ConsumerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageSource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proximo.MessageSource",
	HandlerType: (*MessageSourceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _MessageSource_Consume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proximo.proto",
}

// Client API for MessageSink service

type MessageSinkClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (MessageSink_PublishClient, error)
}

type messageSinkClient struct {
	cc *grpc.ClientConn
}

func NewMessageSinkClient(cc *grpc.ClientConn) MessageSinkClient {
	return &messageSinkClient{cc}
}

func (c *messageSinkClient) Publish(ctx context.Context, opts ...grpc.CallOption) (MessageSink_PublishClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageSink_serviceDesc.Streams[0], c.cc, "/proximo.MessageSink/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageSinkPublishClient{stream}
	return x, nil
}

type MessageSink_PublishClient interface {
	Send(*PublisherRequest) error
	Recv() (*Confirmation, error)
	grpc.ClientStream
}

type messageSinkPublishClient struct {
	grpc.ClientStream
}

func (x *messageSinkPublishClient) Send(m *PublisherRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageSinkPublishClient) Recv() (*Confirmation, error) {
	m := new(Confirmation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MessageSink service

type MessageSinkServer interface {
	Publish(MessageSink_PublishServer) error
}

func RegisterMessageSinkServer(s *grpc.Server, srv MessageSinkServer) {
	s.RegisterService(&_MessageSink_serviceDesc, srv)
}

func _MessageSink_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageSinkServer).Publish(&messageSinkPublishServer{stream})
}

type MessageSink_PublishServer interface {
	Send(*Confirmation) error
	Recv() (*PublisherRequest, error)
	grpc.ServerStream
}

type messageSinkPublishServer struct {
	grpc.ServerStream
}

func (x *messageSinkPublishServer) Send(m *Confirmation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageSinkPublishServer) Recv() (*PublisherRequest, error) {
	m := new(PublisherRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageSink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proximo.MessageSink",
	HandlerType: (*MessageSinkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _MessageSink_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proximo.proto",
}

func init() { proto.RegisterFile("proximo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x56, 0x8d, 0x9d, 0xa6, 0x5a, 0x46, 0x85, 0x58, 0x3c, 0x94, 0xc5, 0x43, 0x41,
	0x2c, 0x52, 0x4f, 0xe2, 0x45, 0xa9, 0x20, 0x82, 0x42, 0xd9, 0x3e, 0x41, 0x9a, 0xae, 0x75, 0xd1,
	0x64, 0xeb, 0xee, 0x06, 0xfa, 0x14, 0x3e, 0xb3, 0x74, 0x3b, 0x0d, 0x89, 0xe6, 0xe2, 0x2d, 0x3f,
	0x33, 0xff, 0x37, 0xfb, 0xcf, 0x04, 0x3a, 0x4b, 0xa3, 0x57, 0x2a, 0xd5, 0xc3, 0xa5, 0xd1, 0x4e,
	0x63, 0x40, 0x92, 0x5f, 0x41, 0xf0, 0x2a, 0xad, 0x8d, 0x17, 0x12, 0x11, 0x76, 0xe7, 0xb1, 0x8b,
	0x23, 0xd6, 0x67, 0x83, 0x50, 0xf8, 0x6f, 0x3c, 0x84, 0x86, 0x9a, 0x47, 0x8d, 0x3e, 0x1b, 0xb4,
	0x44, 0x43, 0xcd, 0xf9, 0x37, 0x83, 0xa3, 0xb1, 0xce, 0x6c, 0x9e, 0x4a, 0x23, 0xe4, 0x57, 0x2e,
	0xad, 0xc3, 0x7b, 0x08, 0xad, 0x8b, 0x8d, 0x23, 0xed, 0xbb, 0xdb, 0xa3, 0xf3, 0xe1, 0x76, 0xe2,
	0x74, 0x5d, 0x24, 0x13, 0xf5, 0x88, 0x8a, 0x03, 0x6f, 0x21, 0x4c, 0x74, 0xf6, 0xa6, 0x4c, 0x1a,
	0x3b, 0xa5, 0xb3, 0xa8, 0xe9, 0x09, 0xa7, 0x05, 0x61, 0x5c, 0x2a, 0x8a, 0x4a, 0x2b, 0x7f, 0x82,
	0xe3, 0x1a, 0x3e, 0x9e, 0xc0, 0x9e, 0xd3, 0x4b, 0x95, 0xf8, 0x30, 0x2d, 0xb1, 0x11, 0xd8, 0x83,
	0x83, 0x84, 0x1e, 0x4f, 0x99, 0x0a, 0xcd, 0x2f, 0x20, 0x2c, 0x8f, 0x59, 0x13, 0x52, 0xbb, 0x78,
	0x7e, 0xdc, 0x12, 0xbc, 0xe0, 0x2b, 0xe8, 0x4e, 0xf2, 0xd9, 0xa7, 0xb2, 0xef, 0xff, 0xcc, 0x4f,
	0xae, 0xfa, 0xfc, 0x1c, 0x9a, 0xa9, 0x5d, 0x50, 0xec, 0x6e, 0x61, 0xa4, 0xc3, 0x88, 0x75, 0x91,
	0x5f, 0x52, 0xd0, 0x2a, 0xa8, 0x3e, 0xe8, 0xe8, 0x05, 0x3a, 0x64, 0x9e, 0xea, 0xdc, 0x24, 0x12,
	0xef, 0x20, 0xa0, 0x0d, 0x61, 0x54, 0x5e, 0x6b, 0xf9, 0x90, 0xbd, 0x3f, 0x93, 0xf9, 0xce, 0x80,
	0x5d, 0xb3, 0xd1, 0x04, 0xda, 0x5b, 0x9a, 0xca, 0x3e, 0xf0, 0x01, 0x02, 0x7a, 0x04, 0x9e, 0x15,
	0x8e, 0xdf, 0x5b, 0xe9, 0xd5, 0x5f, 0x6f, 0x43, 0x9c, 0xed, 0xfb, 0xbf, 0xf0, 0xe6, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0x93, 0x80, 0xcf, 0x96, 0x02, 0x00, 0x00,
}
