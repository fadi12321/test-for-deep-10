// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supply/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// SupplyRequest is the request type for the Query/TotalSupply RPC
// method.
type SupplyRequest struct {
}

func (m *SupplyRequest) Reset()         { *m = SupplyRequest{} }
func (m *SupplyRequest) String() string { return proto.CompactTextString(m) }
func (*SupplyRequest) ProtoMessage()    {}
func (*SupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cd1add7eb523c6, []int{0}
}
func (m *SupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplyRequest.Merge(m, src)
}
func (m *SupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *SupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupplyRequest proto.InternalMessageInfo

// SupplyResponse is the response type for the Query/TotalSupply RPC
// method
type SupplyResponse struct {
	// supply is the supply of the coins
	Supply github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=supply,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supply"`
}

func (m *SupplyResponse) Reset()         { *m = SupplyResponse{} }
func (m *SupplyResponse) String() string { return proto.CompactTextString(m) }
func (*SupplyResponse) ProtoMessage()    {}
func (*SupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64cd1add7eb523c6, []int{1}
}
func (m *SupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplyResponse.Merge(m, src)
}
func (m *SupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *SupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupplyResponse proto.InternalMessageInfo

func (m *SupplyResponse) GetSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Supply
	}
	return nil
}

func init() {
	proto.RegisterType((*SupplyRequest)(nil), "chainmain.supply.v1.SupplyRequest")
	proto.RegisterType((*SupplyResponse)(nil), "chainmain.supply.v1.SupplyResponse")
}

func init() { proto.RegisterFile("supply/v1/query.proto", fileDescriptor_64cd1add7eb523c6) }

var fileDescriptor_64cd1add7eb523c6 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbd, 0x4e, 0xe3, 0x40,
	0x14, 0x85, 0xed, 0x44, 0x9b, 0xc2, 0xd9, 0x1f, 0xc9, 0xbb, 0x2b, 0xed, 0x7a, 0xb3, 0x0e, 0x72,
	0x9a, 0x34, 0x9e, 0xc1, 0x41, 0xbc, 0x40, 0x68, 0x91, 0x10, 0x81, 0x8a, 0x6e, 0xec, 0x8c, 0x9c,
	0x11, 0xce, 0x5c, 0xc7, 0x33, 0x8e, 0x30, 0xa2, 0xa2, 0x40, 0x94, 0x48, 0xbc, 0x05, 0x4f, 0x92,
	0x32, 0x12, 0x0d, 0x15, 0xa0, 0x84, 0x07, 0x41, 0x9e, 0x71, 0xf8, 0x91, 0x22, 0x2a, 0xaa, 0xb9,
	0xba, 0x73, 0xef, 0x39, 0xdf, 0x9c, 0xb1, 0x7e, 0x8b, 0x3c, 0x4d, 0x93, 0x02, 0x4f, 0x03, 0x3c,
	0xc9, 0x69, 0x56, 0xa0, 0x34, 0x03, 0x09, 0xf6, 0xcf, 0x68, 0x44, 0x18, 0x1f, 0x13, 0xc6, 0x91,
	0x1e, 0x40, 0xd3, 0xc0, 0x71, 0x23, 0x10, 0x63, 0x10, 0x38, 0x24, 0x82, 0xe2, 0x69, 0x10, 0x52,
	0x49, 0x02, 0x1c, 0x01, 0xe3, 0x7a, 0xc9, 0xf9, 0x15, 0x43, 0x0c, 0xaa, 0xc4, 0x65, 0x55, 0x75,
	0x5b, 0x31, 0x40, 0x9c, 0x50, 0x4c, 0x52, 0x86, 0x09, 0xe7, 0x20, 0x89, 0x64, 0xc0, 0x85, 0xbe,
	0xf5, 0x7e, 0x58, 0xdf, 0x0e, 0x94, 0xc1, 0x80, 0x4e, 0x72, 0x2a, 0xa4, 0x97, 0x5b, 0xdf, 0x57,
	0x0d, 0x91, 0x02, 0x17, 0xd4, 0x8e, 0xac, 0x86, 0x66, 0xf8, 0x63, 0x6e, 0xd4, 0xbb, 0xcd, 0xde,
	0x5f, 0xa4, 0x39, 0x50, 0xc9, 0x81, 0x2a, 0x0e, 0xb4, 0x03, 0x8c, 0xf7, 0x37, 0x67, 0xf7, 0x6d,
	0xe3, 0xe6, 0xa1, 0xdd, 0x8d, 0x99, 0x1c, 0xe5, 0x21, 0x8a, 0x60, 0x8c, 0x2b, 0x68, 0x7d, 0xf8,
	0x62, 0x78, 0x8c, 0x65, 0x91, 0x52, 0xa1, 0x16, 0xc4, 0xa0, 0x92, 0xee, 0x5d, 0xd4, 0xac, 0x2f,
	0xfb, 0x65, 0x00, 0xf6, 0xa9, 0xd5, 0x3c, 0x04, 0x49, 0x12, 0x4d, 0x61, 0x7b, 0x68, 0x4d, 0x14,
	0xe8, 0x1d, 0xb3, 0xd3, 0xf9, 0x70, 0x46, 0x3f, 0xc3, 0xf3, 0xce, 0x6f, 0x9f, 0xae, 0x6b, 0x2d,
	0xdb, 0xc1, 0x2f, 0xc3, 0xf8, 0x35, 0x7c, 0x59, 0x5a, 0xda, 0x67, 0xd6, 0xd7, 0x5d, 0x36, 0xc9,
	0xd9, 0xf0, 0xb3, 0xcd, 0x3b, 0xca, 0xfc, 0xbf, 0xfd, 0x6f, 0xad, 0x79, 0xa2, 0x3c, 0x9d, 0xfa,
	0x65, 0xcd, 0xec, 0xef, 0xcd, 0x16, 0xae, 0x39, 0x5f, 0xb8, 0xe6, 0xe3, 0xc2, 0x35, 0xaf, 0x96,
	0xae, 0x31, 0x5f, 0xba, 0xc6, 0xdd, 0xd2, 0x35, 0x8e, 0xb6, 0xdf, 0x86, 0x9a, 0x15, 0xa9, 0x04,
	0x1f, 0xb2, 0xd8, 0x57, 0x82, 0x5a, 0xd6, 0x57, 0xba, 0x27, 0x2b, 0x65, 0x95, 0x73, 0xd8, 0x50,
	0x1f, 0xbd, 0xf5, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x38, 0xcd, 0x08, 0x12, 0x6a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type QueryClient interface {
	// TotalSupply queries the total supply of all coins.
	TotalSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
	// LiquidSupply queries the liquid supply of all coins.
	LiquidSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

// Deprecated: Do not use.
func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TotalSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/chainmain.supply.v1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidSupply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/chainmain.supply.v1.Query/LiquidSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
//
// Deprecated: Do not use.
type QueryServer interface {
	// TotalSupply queries the total supply of all coins.
	TotalSupply(context.Context, *SupplyRequest) (*SupplyResponse, error)
	// LiquidSupply queries the liquid supply of all coins.
	LiquidSupply(context.Context, *SupplyRequest) (*SupplyResponse, error)
}

// Deprecated: Do not use.
// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TotalSupply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (*UnimplementedQueryServer) LiquidSupply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidSupply not implemented")
}

// Deprecated: Do not use.
func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.supply.v1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainmain.supply.v1.Query/LiquidSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidSupply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainmain.supply.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "LiquidSupply",
			Handler:    _Query_LiquidSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supply/v1/query.proto",
}

func (m *SupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Supply) > 0 {
		for iNdEx := len(m.Supply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Supply) > 0 {
		for _, e := range m.Supply {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SupplyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *SupplyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supply = append(m.Supply, types.Coin{})
			if err := m.Supply[len(m.Supply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
