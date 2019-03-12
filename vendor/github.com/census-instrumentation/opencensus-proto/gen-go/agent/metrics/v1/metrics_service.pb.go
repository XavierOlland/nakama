// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/agent/metrics/v1/metrics_service.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	v11 "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	v12 "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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

type ExportMetricsServiceRequest struct {
	// This is required only in the first message on the stream or if the
	// previous sent ExportMetricsServiceRequest message has a different Node (e.g.
	// when the same RPC is used to send Metrics from multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of metrics that belong to the last received Node.
	Metrics []*v11.Metric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// The resource for the metrics in this message that do not have an explicit
	// resource set.
	// If unset, the most recently set resource in the RPC stream applies. It is
	// valid to never be set within a stream, e.g. when no resource info is known
	// at all or when all sent metrics have an explicit resource set.
	Resource             *v12.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportMetricsServiceRequest) Reset()         { *m = ExportMetricsServiceRequest{} }
func (m *ExportMetricsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceRequest) ProtoMessage()    {}
func (*ExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e253a956287d04, []int{0}
}

func (m *ExportMetricsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceRequest.Unmarshal(m, b)
}
func (m *ExportMetricsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceRequest.Merge(m, src)
}
func (m *ExportMetricsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceRequest.Size(m)
}
func (m *ExportMetricsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceRequest proto.InternalMessageInfo

func (m *ExportMetricsServiceRequest) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ExportMetricsServiceRequest) GetMetrics() []*v11.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ExportMetricsServiceRequest) GetResource() *v12.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ExportMetricsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportMetricsServiceResponse) Reset()         { *m = ExportMetricsServiceResponse{} }
func (m *ExportMetricsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceResponse) ProtoMessage()    {}
func (*ExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e253a956287d04, []int{1}
}

func (m *ExportMetricsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceResponse.Unmarshal(m, b)
}
func (m *ExportMetricsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceResponse.Merge(m, src)
}
func (m *ExportMetricsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceResponse.Size(m)
}
func (m *ExportMetricsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportMetricsServiceRequest)(nil), "opencensus.proto.agent.metrics.v1.ExportMetricsServiceRequest")
	proto.RegisterType((*ExportMetricsServiceResponse)(nil), "opencensus.proto.agent.metrics.v1.ExportMetricsServiceResponse")
}

func init() {
	proto.RegisterFile("opencensus/proto/agent/metrics/v1/metrics_service.proto", fileDescriptor_47e253a956287d04)
}

var fileDescriptor_47e253a956287d04 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xf9, 0x2b, 0x55, 0xa6, 0xe0, 0x62, 0xdc, 0x94, 0x2a, 0x52, 0xab, 0x48, 0x45,
	0x32, 0x63, 0xea, 0x42, 0x10, 0x54, 0x28, 0xb8, 0x11, 0x94, 0x12, 0x77, 0x6e, 0xa4, 0x4d, 0x2f,
	0x71, 0x16, 0x99, 0x1b, 0x67, 0x26, 0xc1, 0x57, 0x70, 0xe5, 0x3b, 0xf8, 0x5c, 0x3e, 0x8c, 0x24,
	0x93, 0xb4, 0x94, 0x18, 0x0b, 0xee, 0x2e, 0x99, 0xf3, 0x9d, 0x9c, 0x33, 0x73, 0xe9, 0x05, 0x26,
	0xa0, 0x42, 0x50, 0x26, 0x35, 0x22, 0xd1, 0x68, 0x51, 0x4c, 0x23, 0x50, 0x56, 0xc4, 0x60, 0xb5,
	0x0c, 0x8d, 0xc8, 0xfc, 0x6a, 0x7c, 0x36, 0xa0, 0x33, 0x19, 0x02, 0x2f, 0x64, 0xec, 0x60, 0x09,
	0xba, 0x2f, 0xbc, 0x00, 0x79, 0xa9, 0xe6, 0x99, 0xdf, 0xf3, 0x1a, 0xbc, 0x43, 0x8c, 0x63, 0x54,
	0xb9, 0xb5, 0x9b, 0x1c, 0xdf, 0x3b, 0xa9, 0xc9, 0xeb, 0x21, 0x4a, 0xe9, 0x69, 0x4d, 0xaa, 0xc1,
	0x60, 0xaa, 0x43, 0xc8, 0xb5, 0xd5, 0xec, 0xc4, 0x83, 0x2f, 0x42, 0x77, 0x6f, 0xdf, 0x12, 0xd4,
	0xf6, 0xde, 0x99, 0x3c, 0xba, 0x22, 0x01, 0xbc, 0xa6, 0x60, 0x2c, 0xbb, 0xa4, 0x1b, 0x0a, 0xe7,
	0xd0, 0x25, 0x7d, 0x32, 0xec, 0x8c, 0x8e, 0x79, 0x43, 0xb1, 0x32, 0x6b, 0xe6, 0xf3, 0x07, 0x9c,
	0x43, 0x50, 0x30, 0xec, 0x8a, 0x6e, 0x96, 0xc9, 0xba, 0xff, 0xfb, 0xad, 0x61, 0x67, 0x74, 0x58,
	0xc7, 0x97, 0x37, 0xc2, 0x5d, 0x80, 0xa0, 0x62, 0xd8, 0x98, 0x6e, 0x55, 0x61, 0xbb, 0xad, 0xa6,
	0xdf, 0x2f, 0xea, 0x64, 0x3e, 0x0f, 0xca, 0x39, 0x58, 0x70, 0x83, 0x7d, 0xba, 0xf7, 0x73, 0x3b,
	0x93, 0xa0, 0x32, 0x30, 0xfa, 0x24, 0x74, 0x7b, 0xf5, 0x88, 0x7d, 0x10, 0xda, 0x76, 0x0c, 0xbb,
	0xe6, 0x6b, 0xdf, 0x91, 0xff, 0x72, 0x79, 0xbd, 0x9b, 0x3f, 0xf3, 0x2e, 0xde, 0xe0, 0xdf, 0x90,
	0x9c, 0x91, 0xf1, 0x3b, 0xa1, 0x47, 0x12, 0xd7, 0x7b, 0x8d, 0x77, 0x56, 0x6d, 0x26, 0xb9, 0x6a,
	0x42, 0x9e, 0xee, 0x22, 0x69, 0x5f, 0xd2, 0x59, 0xfe, 0x48, 0xc2, 0x19, 0x78, 0x52, 0x19, 0xab,
	0xd3, 0x18, 0x94, 0x9d, 0x5a, 0x89, 0x4a, 0x2c, 0xbd, 0x3d, 0xb7, 0x32, 0x11, 0x28, 0x2f, 0xaa,
	0xef, 0xfb, 0xac, 0x5d, 0x1c, 0x9f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x16, 0x61, 0x3b, 0xc3,
	0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ExportClient, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) Export(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/opencensus.proto.agent.metrics.v1.MetricsService/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceExportClient{stream}
	return x, nil
}

type MetricsService_ExportClient interface {
	Send(*ExportMetricsServiceRequest) error
	Recv() (*ExportMetricsServiceResponse, error)
	grpc.ClientStream
}

type metricsServiceExportClient struct {
	grpc.ClientStream
}

func (x *metricsServiceExportClient) Send(m *ExportMetricsServiceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceExportClient) Recv() (*ExportMetricsServiceResponse, error) {
	m := new(ExportMetricsServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(MetricsService_ExportServer) error
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).Export(&metricsServiceExportServer{stream})
}

type MetricsService_ExportServer interface {
	Send(*ExportMetricsServiceResponse) error
	Recv() (*ExportMetricsServiceRequest, error)
	grpc.ServerStream
}

type metricsServiceExportServer struct {
	grpc.ServerStream
}

func (x *metricsServiceExportServer) Send(m *ExportMetricsServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceExportServer) Recv() (*ExportMetricsServiceRequest, error) {
	m := new(ExportMetricsServiceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opencensus.proto.agent.metrics.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Export",
			Handler:       _MetricsService_Export_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "opencensus/proto/agent/metrics/v1/metrics_service.proto",
}