// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: comments/comments.proto

package commentsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Comments_CreateComment_FullMethodName    = "/comments.Comments/CreateComment"
	Comments_GetComments_FullMethodName      = "/comments.Comments/GetComments"
	Comments_GetChildComments_FullMethodName = "/comments.Comments/GetChildComments"
	Comments_DeleteComment_FullMethodName    = "/comments.Comments/DeleteComment"
	Comments_UpdateComment_FullMethodName    = "/comments.Comments/UpdateComment"
)

// CommentsClient is the client API for Comments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentsClient interface {
	CreateComment(ctx context.Context, in *CreateCommentDto, opts ...grpc.CallOption) (*CommentDto, error)
	GetComments(ctx context.Context, in *GetCommentsDto, opts ...grpc.CallOption) (*CommentsResponse, error)
	GetChildComments(ctx context.Context, in *GetChildCommentsDto, opts ...grpc.CallOption) (*CommentsResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentDto, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentDto, opts ...grpc.CallOption) (*CommentDto, error)
}

type commentsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsClient(cc grpc.ClientConnInterface) CommentsClient {
	return &commentsClient{cc}
}

func (c *commentsClient) CreateComment(ctx context.Context, in *CreateCommentDto, opts ...grpc.CallOption) (*CommentDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentDto)
	err := c.cc.Invoke(ctx, Comments_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) GetComments(ctx context.Context, in *GetCommentsDto, opts ...grpc.CallOption) (*CommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, Comments_GetComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) GetChildComments(ctx context.Context, in *GetChildCommentsDto, opts ...grpc.CallOption) (*CommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, Comments_GetChildComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) DeleteComment(ctx context.Context, in *DeleteCommentDto, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, Comments_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) UpdateComment(ctx context.Context, in *UpdateCommentDto, opts ...grpc.CallOption) (*CommentDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentDto)
	err := c.cc.Invoke(ctx, Comments_UpdateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsServer is the server API for Comments service.
// All implementations must embed UnimplementedCommentsServer
// for forward compatibility.
type CommentsServer interface {
	CreateComment(context.Context, *CreateCommentDto) (*CommentDto, error)
	GetComments(context.Context, *GetCommentsDto) (*CommentsResponse, error)
	GetChildComments(context.Context, *GetChildCommentsDto) (*CommentsResponse, error)
	DeleteComment(context.Context, *DeleteCommentDto) (*DeleteCommentResponse, error)
	UpdateComment(context.Context, *UpdateCommentDto) (*CommentDto, error)
	mustEmbedUnimplementedCommentsServer()
}

// UnimplementedCommentsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentsServer struct{}

func (UnimplementedCommentsServer) CreateComment(context.Context, *CreateCommentDto) (*CommentDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentsServer) GetComments(context.Context, *GetCommentsDto) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedCommentsServer) GetChildComments(context.Context, *GetChildCommentsDto) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildComments not implemented")
}
func (UnimplementedCommentsServer) DeleteComment(context.Context, *DeleteCommentDto) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentsServer) UpdateComment(context.Context, *UpdateCommentDto) (*CommentDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentsServer) mustEmbedUnimplementedCommentsServer() {}
func (UnimplementedCommentsServer) testEmbeddedByValue()                  {}

// UnsafeCommentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentsServer will
// result in compilation errors.
type UnsafeCommentsServer interface {
	mustEmbedUnimplementedCommentsServer()
}

func RegisterCommentsServer(s grpc.ServiceRegistrar, srv CommentsServer) {
	// If the following call pancis, it indicates UnimplementedCommentsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Comments_ServiceDesc, srv)
}

func _Comments_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comments_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).CreateComment(ctx, req.(*CreateCommentDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comments_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).GetComments(ctx, req.(*GetCommentsDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_GetChildComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChildCommentsDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).GetChildComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comments_GetChildComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).GetChildComments(ctx, req.(*GetChildCommentsDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comments_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).DeleteComment(ctx, req.(*DeleteCommentDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comments_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).UpdateComment(ctx, req.(*UpdateCommentDto))
	}
	return interceptor(ctx, in, info, handler)
}

// Comments_ServiceDesc is the grpc.ServiceDesc for Comments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comments.Comments",
	HandlerType: (*CommentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _Comments_CreateComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _Comments_GetComments_Handler,
		},
		{
			MethodName: "GetChildComments",
			Handler:    _Comments_GetChildComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Comments_DeleteComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Comments_UpdateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comments/comments.proto",
}
