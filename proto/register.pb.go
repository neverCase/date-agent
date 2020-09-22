// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package dateagent_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RegisterRequest struct {
	Hostname             string               `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RegisterRequest) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type RegisterReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (m *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(m, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

type Task struct {
	TaskId               int32    `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{2}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *Task) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type PullTaskRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullTaskRequest) Reset()         { *m = PullTaskRequest{} }
func (m *PullTaskRequest) String() string { return proto.CompactTextString(m) }
func (*PullTaskRequest) ProtoMessage()    {}
func (*PullTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{3}
}

func (m *PullTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullTaskRequest.Unmarshal(m, b)
}
func (m *PullTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullTaskRequest.Marshal(b, m, deterministic)
}
func (m *PullTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullTaskRequest.Merge(m, src)
}
func (m *PullTaskRequest) XXX_Size() int {
	return xxx_messageInfo_PullTaskRequest.Size(m)
}
func (m *PullTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullTaskRequest proto.InternalMessageInfo

type PullTaskReply struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullTaskReply) Reset()         { *m = PullTaskReply{} }
func (m *PullTaskReply) String() string { return proto.CompactTextString(m) }
func (*PullTaskReply) ProtoMessage()    {}
func (*PullTaskReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{4}
}

func (m *PullTaskReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullTaskReply.Unmarshal(m, b)
}
func (m *PullTaskReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullTaskReply.Marshal(b, m, deterministic)
}
func (m *PullTaskReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullTaskReply.Merge(m, src)
}
func (m *PullTaskReply) XXX_Size() int {
	return xxx_messageInfo_PullTaskReply.Size(m)
}
func (m *PullTaskReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PullTaskReply.DiscardUnknown(m)
}

var xxx_messageInfo_PullTaskReply proto.InternalMessageInfo

func (m *PullTaskReply) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type CompleteTaskRequest struct {
	TaskId               int32    `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteTaskRequest) Reset()         { *m = CompleteTaskRequest{} }
func (m *CompleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteTaskRequest) ProtoMessage()    {}
func (*CompleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{5}
}

func (m *CompleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteTaskRequest.Unmarshal(m, b)
}
func (m *CompleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *CompleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteTaskRequest.Merge(m, src)
}
func (m *CompleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteTaskRequest.Size(m)
}
func (m *CompleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteTaskRequest proto.InternalMessageInfo

func (m *CompleteTaskRequest) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type CompleteTaskReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteTaskReply) Reset()         { *m = CompleteTaskReply{} }
func (m *CompleteTaskReply) String() string { return proto.CompactTextString(m) }
func (*CompleteTaskReply) ProtoMessage()    {}
func (*CompleteTaskReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{6}
}

func (m *CompleteTaskReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteTaskReply.Unmarshal(m, b)
}
func (m *CompleteTaskReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteTaskReply.Marshal(b, m, deterministic)
}
func (m *CompleteTaskReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteTaskReply.Merge(m, src)
}
func (m *CompleteTaskReply) XXX_Size() int {
	return xxx_messageInfo_CompleteTaskReply.Size(m)
}
func (m *CompleteTaskReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteTaskReply.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteTaskReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "dateagent.v1.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "dateagent.v1.RegisterReply")
	proto.RegisterType((*Task)(nil), "dateagent.v1.Task")
	proto.RegisterType((*PullTaskRequest)(nil), "dateagent.v1.PullTaskRequest")
	proto.RegisterType((*PullTaskReply)(nil), "dateagent.v1.PullTaskReply")
	proto.RegisterType((*CompleteTaskRequest)(nil), "dateagent.v1.CompleteTaskRequest")
	proto.RegisterType((*CompleteTaskReply)(nil), "dateagent.v1.CompleteTaskReply")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x59, 0x83, 0xfc, 0x19, 0x40, 0x42, 0x39, 0x48, 0xd6, 0x18, 0xb0, 0x07, 0xc3, 0xa9,
	0x44, 0x3c, 0x18, 0x8f, 0x46, 0x0f, 0x7a, 0x33, 0x0d, 0x57, 0x63, 0x8a, 0x3b, 0xae, 0x84, 0x96,
	0x56, 0x5a, 0x4c, 0xf8, 0xe0, 0xde, 0x4d, 0x0b, 0x2b, 0xec, 0x46, 0x3c, 0xce, 0xce, 0xdb, 0xdf,
	0x7b, 0xf3, 0x0a, 0x27, 0x4b, 0x4c, 0x67, 0xd6, 0xe1, 0x92, 0x99, 0xa5, 0x76, 0x9a, 0x34, 0x13,
	0xe1, 0x50, 0xa4, 0xb8, 0x70, 0xec, 0xeb, 0x2a, 0xee, 0xa7, 0x5a, 0xa7, 0x12, 0x47, 0x61, 0x37,
	0x5d, 0xbd, 0x8f, 0xdc, 0x4c, 0xa1, 0x75, 0x42, 0x99, 0x8d, 0x9c, 0xbe, 0x40, 0x9b, 0x6f, 0x01,
	0x1c, 0x3f, 0x57, 0x68, 0x1d, 0x89, 0xa1, 0xf6, 0xa1, 0xad, 0x5b, 0x08, 0x85, 0xbd, 0x68, 0x10,
	0x0d, 0xeb, 0xfc, 0x77, 0x26, 0x0c, 0xca, 0x9e, 0xd0, 0x3b, 0x1a, 0x44, 0xc3, 0xc6, 0x38, 0x66,
	0x1b, 0x3c, 0xcb, 0xf0, 0x6c, 0x92, 0xe1, 0x79, 0xd0, 0xd1, 0x36, 0xb4, 0x76, 0x78, 0x23, 0xd7,
	0xf4, 0x16, 0xca, 0x13, 0x61, 0xe7, 0xe4, 0x14, 0xaa, 0x4e, 0xd8, 0xf9, 0xeb, 0x2c, 0x09, 0x1e,
	0xc7, 0xbc, 0xe2, 0xc7, 0xa7, 0x84, 0xf4, 0xa0, 0xfa, 0xa6, 0x95, 0x12, 0x8b, 0x24, 0x98, 0xd4,
	0x79, 0x36, 0xd2, 0x0e, 0xb4, 0x9f, 0x57, 0x52, 0xfa, 0xdf, 0xb7, 0x51, 0xe9, 0x0d, 0xb4, 0x76,
	0x9f, 0x8c, 0x5c, 0x93, 0x4b, 0x28, 0x7b, 0x4e, 0x60, 0x36, 0xc6, 0x84, 0xed, 0x97, 0xc1, 0x82,
	0x2c, 0xec, 0x29, 0x83, 0xee, 0xbd, 0x56, 0x46, 0xa2, 0xc3, 0x3d, 0xde, 0xc1, 0x54, 0xb4, 0x0b,
	0x9d, 0xbc, 0xde, 0xc8, 0xf5, 0xf8, 0x3b, 0x82, 0xfa, 0x83, 0x70, 0x78, 0xe7, 0x0d, 0xc8, 0x23,
	0xd4, 0xb2, 0x53, 0xc9, 0x79, 0xde, 0xb8, 0xd0, 0x70, 0x7c, 0x76, 0x68, 0xed, 0x1b, 0x2a, 0x79,
	0x52, 0x76, 0x55, 0x91, 0x54, 0x28, 0xa0, 0x48, 0xca, 0x95, 0x41, 0x4b, 0x64, 0x02, 0xcd, 0xfd,
	0xd8, 0xe4, 0x22, 0x2f, 0xff, 0xa3, 0x82, 0xb8, 0xff, 0x9f, 0x24, 0x50, 0xa7, 0x95, 0xf0, 0xdc,
	0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x81, 0x03, 0xb6, 0x7b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DateAgentClient is the client API for DateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DateAgentClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	PullTask(ctx context.Context, in *PullTaskRequest, opts ...grpc.CallOption) (*PullTaskReply, error)
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskReply, error)
}

type dateAgentClient struct {
	cc *grpc.ClientConn
}

func NewDateAgentClient(cc *grpc.ClientConn) DateAgentClient {
	return &dateAgentClient{cc}
}

func (c *dateAgentClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/dateagent.v1.DateAgent/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateAgentClient) PullTask(ctx context.Context, in *PullTaskRequest, opts ...grpc.CallOption) (*PullTaskReply, error) {
	out := new(PullTaskReply)
	err := c.cc.Invoke(ctx, "/dateagent.v1.DateAgent/PullTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateAgentClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskReply, error) {
	out := new(CompleteTaskReply)
	err := c.cc.Invoke(ctx, "/dateagent.v1.DateAgent/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DateAgentServer is the server API for DateAgent service.
type DateAgentServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	PullTask(context.Context, *PullTaskRequest) (*PullTaskReply, error)
	CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskReply, error)
}

// UnimplementedDateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedDateAgentServer struct {
}

func (*UnimplementedDateAgentServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedDateAgentServer) PullTask(ctx context.Context, req *PullTaskRequest) (*PullTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullTask not implemented")
}
func (*UnimplementedDateAgentServer) CompleteTask(ctx context.Context, req *CompleteTaskRequest) (*CompleteTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}

func RegisterDateAgentServer(s *grpc.Server, srv DateAgentServer) {
	s.RegisterService(&_DateAgent_serviceDesc, srv)
}

func _DateAgent_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateAgentServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dateagent.v1.DateAgent/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateAgentServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateAgent_PullTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateAgentServer).PullTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dateagent.v1.DateAgent/PullTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateAgentServer).PullTask(ctx, req.(*PullTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateAgent_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateAgentServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dateagent.v1.DateAgent/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateAgentServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dateagent.v1.DateAgent",
	HandlerType: (*DateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DateAgent_Register_Handler,
		},
		{
			MethodName: "PullTask",
			Handler:    _DateAgent_PullTask_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _DateAgent_CompleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}