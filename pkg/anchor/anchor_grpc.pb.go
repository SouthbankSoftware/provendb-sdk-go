// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package anchor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AnchorServiceClient is the client API for AnchorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnchorServiceClient interface {
	// GetAnchors gets all anchors
	GetAnchors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (AnchorService_GetAnchorsClient, error)
	// GetAnchor gets an anchor
	GetAnchor(ctx context.Context, in *AnchorRequest, opts ...grpc.CallOption) (*Anchor, error)
	// GetProof gets a proof
	GetProof(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*Proof, error)
	// SubmitProof submits a proof for the given hash
	SubmitProof(ctx context.Context, in *SubmitProofRequest, opts ...grpc.CallOption) (*Proof, error)
	// VerifyProof verifies the given proof. When the proof is unverifiable, an
	// exception is thrown
	VerifyProof(ctx context.Context, in *VerifyProofRequest, opts ...grpc.CallOption) (*VerifyProofReply, error)
	// GetBatch gets a batch
	GetBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*Batch, error)
	// SubscribeBatches subscribes to batch status updates
	SubscribeBatches(ctx context.Context, in *SubscribeBatchesRequest, opts ...grpc.CallOption) (AnchorService_SubscribeBatchesClient, error)
}

type anchorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnchorServiceClient(cc grpc.ClientConnInterface) AnchorServiceClient {
	return &anchorServiceClient{cc}
}

func (c *anchorServiceClient) GetAnchors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (AnchorService_GetAnchorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnchorService_ServiceDesc.Streams[0], "/anchor.AnchorService/GetAnchors", opts...)
	if err != nil {
		return nil, err
	}
	x := &anchorServiceGetAnchorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AnchorService_GetAnchorsClient interface {
	Recv() (*Anchor, error)
	grpc.ClientStream
}

type anchorServiceGetAnchorsClient struct {
	grpc.ClientStream
}

func (x *anchorServiceGetAnchorsClient) Recv() (*Anchor, error) {
	m := new(Anchor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *anchorServiceClient) GetAnchor(ctx context.Context, in *AnchorRequest, opts ...grpc.CallOption) (*Anchor, error) {
	out := new(Anchor)
	err := c.cc.Invoke(ctx, "/anchor.AnchorService/GetAnchor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorServiceClient) GetProof(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*Proof, error) {
	out := new(Proof)
	err := c.cc.Invoke(ctx, "/anchor.AnchorService/GetProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorServiceClient) SubmitProof(ctx context.Context, in *SubmitProofRequest, opts ...grpc.CallOption) (*Proof, error) {
	out := new(Proof)
	err := c.cc.Invoke(ctx, "/anchor.AnchorService/SubmitProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorServiceClient) VerifyProof(ctx context.Context, in *VerifyProofRequest, opts ...grpc.CallOption) (*VerifyProofReply, error) {
	out := new(VerifyProofReply)
	err := c.cc.Invoke(ctx, "/anchor.AnchorService/VerifyProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorServiceClient) GetBatch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*Batch, error) {
	out := new(Batch)
	err := c.cc.Invoke(ctx, "/anchor.AnchorService/GetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorServiceClient) SubscribeBatches(ctx context.Context, in *SubscribeBatchesRequest, opts ...grpc.CallOption) (AnchorService_SubscribeBatchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnchorService_ServiceDesc.Streams[1], "/anchor.AnchorService/SubscribeBatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &anchorServiceSubscribeBatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AnchorService_SubscribeBatchesClient interface {
	Recv() (*Batch, error)
	grpc.ClientStream
}

type anchorServiceSubscribeBatchesClient struct {
	grpc.ClientStream
}

func (x *anchorServiceSubscribeBatchesClient) Recv() (*Batch, error) {
	m := new(Batch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnchorServiceServer is the server API for AnchorService service.
// All implementations must embed UnimplementedAnchorServiceServer
// for forward compatibility
type AnchorServiceServer interface {
	// GetAnchors gets all anchors
	GetAnchors(*emptypb.Empty, AnchorService_GetAnchorsServer) error
	// GetAnchor gets an anchor
	GetAnchor(context.Context, *AnchorRequest) (*Anchor, error)
	// GetProof gets a proof
	GetProof(context.Context, *ProofRequest) (*Proof, error)
	// SubmitProof submits a proof for the given hash
	SubmitProof(context.Context, *SubmitProofRequest) (*Proof, error)
	// VerifyProof verifies the given proof. When the proof is unverifiable, an
	// exception is thrown
	VerifyProof(context.Context, *VerifyProofRequest) (*VerifyProofReply, error)
	// GetBatch gets a batch
	GetBatch(context.Context, *BatchRequest) (*Batch, error)
	// SubscribeBatches subscribes to batch status updates
	SubscribeBatches(*SubscribeBatchesRequest, AnchorService_SubscribeBatchesServer) error
	mustEmbedUnimplementedAnchorServiceServer()
}

// UnimplementedAnchorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnchorServiceServer struct {
}

func (UnimplementedAnchorServiceServer) GetAnchors(*emptypb.Empty, AnchorService_GetAnchorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAnchors not implemented")
}
func (UnimplementedAnchorServiceServer) GetAnchor(context.Context, *AnchorRequest) (*Anchor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnchor not implemented")
}
func (UnimplementedAnchorServiceServer) GetProof(context.Context, *ProofRequest) (*Proof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProof not implemented")
}
func (UnimplementedAnchorServiceServer) SubmitProof(context.Context, *SubmitProofRequest) (*Proof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProof not implemented")
}
func (UnimplementedAnchorServiceServer) VerifyProof(context.Context, *VerifyProofRequest) (*VerifyProofReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyProof not implemented")
}
func (UnimplementedAnchorServiceServer) GetBatch(context.Context, *BatchRequest) (*Batch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatch not implemented")
}
func (UnimplementedAnchorServiceServer) SubscribeBatches(*SubscribeBatchesRequest, AnchorService_SubscribeBatchesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBatches not implemented")
}
func (UnimplementedAnchorServiceServer) mustEmbedUnimplementedAnchorServiceServer() {}

// UnsafeAnchorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnchorServiceServer will
// result in compilation errors.
type UnsafeAnchorServiceServer interface {
	mustEmbedUnimplementedAnchorServiceServer()
}

func RegisterAnchorServiceServer(s grpc.ServiceRegistrar, srv AnchorServiceServer) {
	s.RegisterService(&AnchorService_ServiceDesc, srv)
}

func _AnchorService_GetAnchors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnchorServiceServer).GetAnchors(m, &anchorServiceGetAnchorsServer{stream})
}

type AnchorService_GetAnchorsServer interface {
	Send(*Anchor) error
	grpc.ServerStream
}

type anchorServiceGetAnchorsServer struct {
	grpc.ServerStream
}

func (x *anchorServiceGetAnchorsServer) Send(m *Anchor) error {
	return x.ServerStream.SendMsg(m)
}

func _AnchorService_GetAnchor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnchorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnchorServiceServer).GetAnchor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anchor.AnchorService/GetAnchor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnchorServiceServer).GetAnchor(ctx, req.(*AnchorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnchorService_GetProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnchorServiceServer).GetProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anchor.AnchorService/GetProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnchorServiceServer).GetProof(ctx, req.(*ProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnchorService_SubmitProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnchorServiceServer).SubmitProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anchor.AnchorService/SubmitProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnchorServiceServer).SubmitProof(ctx, req.(*SubmitProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnchorService_VerifyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnchorServiceServer).VerifyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anchor.AnchorService/VerifyProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnchorServiceServer).VerifyProof(ctx, req.(*VerifyProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnchorService_GetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnchorServiceServer).GetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anchor.AnchorService/GetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnchorServiceServer).GetBatch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnchorService_SubscribeBatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeBatchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnchorServiceServer).SubscribeBatches(m, &anchorServiceSubscribeBatchesServer{stream})
}

type AnchorService_SubscribeBatchesServer interface {
	Send(*Batch) error
	grpc.ServerStream
}

type anchorServiceSubscribeBatchesServer struct {
	grpc.ServerStream
}

func (x *anchorServiceSubscribeBatchesServer) Send(m *Batch) error {
	return x.ServerStream.SendMsg(m)
}

// AnchorService_ServiceDesc is the grpc.ServiceDesc for AnchorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnchorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anchor.AnchorService",
	HandlerType: (*AnchorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnchor",
			Handler:    _AnchorService_GetAnchor_Handler,
		},
		{
			MethodName: "GetProof",
			Handler:    _AnchorService_GetProof_Handler,
		},
		{
			MethodName: "SubmitProof",
			Handler:    _AnchorService_SubmitProof_Handler,
		},
		{
			MethodName: "VerifyProof",
			Handler:    _AnchorService_VerifyProof_Handler,
		},
		{
			MethodName: "GetBatch",
			Handler:    _AnchorService_GetBatch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAnchors",
			Handler:       _AnchorService_GetAnchors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeBatches",
			Handler:       _AnchorService_SubscribeBatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/SouthbankSoftware/provendb-apis/anchor/anchor.proto",
}
