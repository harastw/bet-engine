package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type PlaceBetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId uint64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount int64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Odds   float64 `protobuf:"fixed64,4,opt,name=odds,proto3" json:"odds,omitempty"`
}

func (x *PlaceBetRequest) ProtoReflect() protoreflect.Message { return nil }
func (x *PlaceBetRequest) Reset()                             {}
func (x *PlaceBetRequest) String() string                     { return "" }
func (*PlaceBetRequest) ProtoMessage()                        {}

type PlaceBetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PlaceBetResponse) ProtoReflect() protoreflect.Message { return nil }
func (x *PlaceBetResponse) Reset()                             {}
func (x *PlaceBetResponse) String() string                     { return "" }
func (*PlaceBetResponse) ProtoMessage()                        {}

type BetServiceClient interface {
	PlaceBet(ctx context.Context, in *PlaceBetRequest, opts ...grpc.CallOption) (*PlaceBetResponse, error)
}

type BetServiceServer interface {
	PlaceBet(context.Context, *PlaceBetRequest) (*PlaceBetResponse, error)
	mustEmbedUnimplementedBetServiceServer()
}

type UnimplementedBetServiceServer struct{}

func (UnimplementedBetServiceServer) PlaceBet(context.Context, *PlaceBetRequest) (*PlaceBetResponse, error) {
	return nil, nil
}
func (UnimplementedBetServiceServer) mustEmbedUnimplementedBetServiceServer() {}

func RegisterBetServiceServer(s grpc.ServiceRegistrar, srv BetServiceServer) {
	s.RegisterService(&BetService_ServiceDesc, srv)
}

var BetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BetService",
	HandlerType: (*BetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBet",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PlaceBetRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(BetServiceServer).PlaceBet(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/pb.BetService/PlaceBet",
				}
				handler := func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(BetServiceServer).PlaceBet(ctx, req.(*PlaceBetRequest))
				}
				return interceptor(ctx, in, info, handler)
			},
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bet.proto",
}
