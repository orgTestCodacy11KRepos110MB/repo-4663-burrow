// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpcdump

import (
	context "context"

	dump "github.com/hyperledger/burrow/dump"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DumpClient is the client API for Dump service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DumpClient interface {
	GetDump(ctx context.Context, in *GetDumpParam, opts ...grpc.CallOption) (Dump_GetDumpClient, error)
}

type dumpClient struct {
	cc grpc.ClientConnInterface
}

func NewDumpClient(cc grpc.ClientConnInterface) DumpClient {
	return &dumpClient{cc}
}

func (c *dumpClient) GetDump(ctx context.Context, in *GetDumpParam, opts ...grpc.CallOption) (Dump_GetDumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dump_serviceDesc.Streams[0], "/rpcdump.Dump/GetDump", opts...)
	if err != nil {
		return nil, err
	}
	x := &dumpGetDumpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dump_GetDumpClient interface {
	Recv() (*dump.Dump, error)
	grpc.ClientStream
}

type dumpGetDumpClient struct {
	grpc.ClientStream
}

func (x *dumpGetDumpClient) Recv() (*dump.Dump, error) {
	m := new(dump.Dump)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DumpServer is the server API for Dump service.
// All implementations must embed UnimplementedDumpServer
// for forward compatibility
type DumpServer interface {
	GetDump(*GetDumpParam, Dump_GetDumpServer) error
	mustEmbedUnimplementedDumpServer()
}

// UnimplementedDumpServer must be embedded to have forward compatible implementations.
type UnimplementedDumpServer struct {
}

func (UnimplementedDumpServer) GetDump(*GetDumpParam, Dump_GetDumpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDump not implemented")
}
func (UnimplementedDumpServer) mustEmbedUnimplementedDumpServer() {}

// UnsafeDumpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DumpServer will
// result in compilation errors.
type UnsafeDumpServer interface {
	mustEmbedUnimplementedDumpServer()
}

func RegisterDumpServer(s grpc.ServiceRegistrar, srv DumpServer) {
	s.RegisterService(&_Dump_serviceDesc, srv)
}

func _Dump_GetDump_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDumpParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DumpServer).GetDump(m, &dumpGetDumpServer{stream})
}

type Dump_GetDumpServer interface {
	Send(*dump.Dump) error
	grpc.ServerStream
}

type dumpGetDumpServer struct {
	grpc.ServerStream
}

func (x *dumpGetDumpServer) Send(m *dump.Dump) error {
	return x.ServerStream.SendMsg(m)
}

var _Dump_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcdump.Dump",
	HandlerType: (*DumpServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDump",
			Handler:       _Dump_GetDump_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpcdump.proto",
}