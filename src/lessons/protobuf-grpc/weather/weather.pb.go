// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weather.proto

package weather

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WeatherInquiry struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherInquiry) Reset()         { *m = WeatherInquiry{} }
func (m *WeatherInquiry) String() string { return proto.CompactTextString(m) }
func (*WeatherInquiry) ProtoMessage()    {}
func (*WeatherInquiry) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{0}
}

func (m *WeatherInquiry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherInquiry.Unmarshal(m, b)
}
func (m *WeatherInquiry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherInquiry.Marshal(b, m, deterministic)
}
func (m *WeatherInquiry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherInquiry.Merge(m, src)
}
func (m *WeatherInquiry) XXX_Size() int {
	return xxx_messageInfo_WeatherInquiry.Size(m)
}
func (m *WeatherInquiry) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherInquiry.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherInquiry proto.InternalMessageInfo

func (m *WeatherInquiry) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type WeatherReport struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Weather              string   `protobuf:"bytes,2,opt,name=weather,proto3" json:"weather,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherReport) Reset()         { *m = WeatherReport{} }
func (m *WeatherReport) String() string { return proto.CompactTextString(m) }
func (*WeatherReport) ProtoMessage()    {}
func (*WeatherReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{1}
}

func (m *WeatherReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherReport.Unmarshal(m, b)
}
func (m *WeatherReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherReport.Marshal(b, m, deterministic)
}
func (m *WeatherReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherReport.Merge(m, src)
}
func (m *WeatherReport) XXX_Size() int {
	return xxx_messageInfo_WeatherReport.Size(m)
}
func (m *WeatherReport) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherReport.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherReport proto.InternalMessageInfo

func (m *WeatherReport) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *WeatherReport) GetWeather() string {
	if m != nil {
		return m.Weather
	}
	return ""
}

func init() {
	proto.RegisterType((*WeatherInquiry)(nil), "weather.WeatherInquiry")
	proto.RegisterType((*WeatherReport)(nil), "weather.WeatherReport")
}

func init() { proto.RegisterFile("weather.proto", fileDescriptor_231dcd72b885f4be) }

var fileDescriptor_231dcd72b885f4be = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0x4d, 0x2c,
	0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x54, 0xb8,
	0xf8, 0xc2, 0x21, 0x4c, 0xcf, 0xbc, 0xc2, 0xd2, 0xcc, 0xa2, 0x4a, 0x21, 0x21, 0x2e, 0x96, 0x94,
	0xc4, 0x92, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x96, 0x8b, 0x17,
	0xaa, 0x2a, 0x28, 0xb5, 0x20, 0xbf, 0xa8, 0x04, 0x9b, 0x22, 0x21, 0x09, 0x2e, 0x98, 0xa9, 0x12,
	0x4c, 0x60, 0x61, 0x18, 0xd7, 0xc8, 0x83, 0x8b, 0x03, 0xaa, 0xbd, 0x58, 0xc8, 0x86, 0x8b, 0x1d,
	0x62, 0x53, 0xaa, 0x90, 0xb8, 0x1e, 0xcc, 0x51, 0xa8, 0x4e, 0x90, 0x12, 0x43, 0x97, 0x80, 0xd8,
	0xaa, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x35, 0xc3,
	0xc4, 0xcf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WeathersClient is the client API for Weathers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeathersClient interface {
	Inquire(ctx context.Context, in *WeatherInquiry, opts ...grpc.CallOption) (*WeatherReport, error)
}

type weathersClient struct {
	cc *grpc.ClientConn
}

func NewWeathersClient(cc *grpc.ClientConn) WeathersClient {
	return &weathersClient{cc}
}

func (c *weathersClient) Inquire(ctx context.Context, in *WeatherInquiry, opts ...grpc.CallOption) (*WeatherReport, error) {
	out := new(WeatherReport)
	err := c.cc.Invoke(ctx, "/weather.Weathers/Inquire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeathersServer is the server API for Weathers service.
type WeathersServer interface {
	Inquire(context.Context, *WeatherInquiry) (*WeatherReport, error)
}

func RegisterWeathersServer(s *grpc.Server, srv WeathersServer) {
	s.RegisterService(&_Weathers_serviceDesc, srv)
}

func _Weathers_Inquire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherInquiry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeathersServer).Inquire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.Weathers/Inquire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeathersServer).Inquire(ctx, req.(*WeatherInquiry))
	}
	return interceptor(ctx, in, info, handler)
}

var _Weathers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weather.Weathers",
	HandlerType: (*WeathersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inquire",
			Handler:    _Weathers_Inquire_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather.proto",
}
