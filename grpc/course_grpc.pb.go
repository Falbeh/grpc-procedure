// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: grpc/course.proto

package course

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CourseAsk_AskForCourse_FullMethodName = "/course.CourseAsk/AskForCourse"
)

// CourseAskClient is the client API for CourseAsk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseAskClient interface {
	AskForCourse(ctx context.Context, in *CourseInfoRequest, opts ...grpc.CallOption) (*CourseInfoResponse, error)
}

type courseAskClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseAskClient(cc grpc.ClientConnInterface) CourseAskClient {
	return &courseAskClient{cc}
}

func (c *courseAskClient) AskForCourse(ctx context.Context, in *CourseInfoRequest, opts ...grpc.CallOption) (*CourseInfoResponse, error) {
	out := new(CourseInfoResponse)
	err := c.cc.Invoke(ctx, CourseAsk_AskForCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseAskServer is the server API for CourseAsk service.
// All implementations must embed UnimplementedCourseAskServer
// for forward compatibility
type CourseAskServer interface {
	AskForCourse(context.Context, *CourseInfoRequest) (*CourseInfoResponse, error)
	mustEmbedUnimplementedCourseAskServer()
}

// UnimplementedCourseAskServer must be embedded to have forward compatible implementations.
type UnimplementedCourseAskServer struct {
}

func (UnimplementedCourseAskServer) AskForCourse(context.Context, *CourseInfoRequest) (*CourseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskForCourse not implemented")
}
func (UnimplementedCourseAskServer) mustEmbedUnimplementedCourseAskServer() {}

// UnsafeCourseAskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseAskServer will
// result in compilation errors.
type UnsafeCourseAskServer interface {
	mustEmbedUnimplementedCourseAskServer()
}

func RegisterCourseAskServer(s grpc.ServiceRegistrar, srv CourseAskServer) {
	s.RegisterService(&CourseAsk_ServiceDesc, srv)
}

func _CourseAsk_AskForCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseAskServer).AskForCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseAsk_AskForCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseAskServer).AskForCourse(ctx, req.(*CourseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseAsk_ServiceDesc is the grpc.ServiceDesc for CourseAsk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseAsk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "course.CourseAsk",
	HandlerType: (*CourseAskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskForCourse",
			Handler:    _CourseAsk_AskForCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/course.proto",
}
