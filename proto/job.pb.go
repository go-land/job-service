// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/job.proto

package job

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type GetJobRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobRequest) Reset()         { *m = GetJobRequest{} }
func (m *GetJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobRequest) ProtoMessage()    {}
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e2baf7640c8fed, []int{0}
}

func (m *GetJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobRequest.Unmarshal(m, b)
}
func (m *GetJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobRequest.Marshal(b, m, deterministic)
}
func (m *GetJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobRequest.Merge(m, src)
}
func (m *GetJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobRequest.Size(m)
}
func (m *GetJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobRequest proto.InternalMessageInfo

func (m *GetJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JobResponse struct {
	Job                  string   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobResponse) Reset()         { *m = JobResponse{} }
func (m *JobResponse) String() string { return proto.CompactTextString(m) }
func (*JobResponse) ProtoMessage()    {}
func (*JobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e2baf7640c8fed, []int{1}
}

func (m *JobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobResponse.Unmarshal(m, b)
}
func (m *JobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobResponse.Marshal(b, m, deterministic)
}
func (m *JobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobResponse.Merge(m, src)
}
func (m *JobResponse) XXX_Size() int {
	return xxx_messageInfo_JobResponse.Size(m)
}
func (m *JobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobResponse proto.InternalMessageInfo

func (m *JobResponse) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobRequest)(nil), "job.GetJobRequest")
	proto.RegisterType((*JobResponse)(nil), "job.JobResponse")
}

func init() { proto.RegisterFile("proto/job.proto", fileDescriptor_c5e2baf7640c8fed) }

var fileDescriptor_c5e2baf7640c8fed = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xca, 0x4f, 0xd2, 0x03, 0xb3, 0x84, 0x98, 0xb3, 0xf2, 0x93, 0x94, 0x94, 0xb9,
	0x78, 0xdd, 0x53, 0x4b, 0xbc, 0xf2, 0x93, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25,
	0x79, 0x2e, 0x6e, 0xb0, 0x8a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e, 0x90, 0x56,
	0xa8, 0x0a, 0x10, 0xd3, 0xc8, 0x8e, 0x8b, 0xcb, 0x2b, 0x3f, 0x29, 0x38, 0xb5, 0xa8, 0x2c, 0x33,
	0x39, 0x55, 0xc8, 0x80, 0x8b, 0x0d, 0x62, 0xa6, 0x90, 0x90, 0x1e, 0xc8, 0x3a, 0x14, 0x0b, 0xa4,
	0x04, 0xc0, 0x62, 0x48, 0xe6, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x5d, 0x64, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x24, 0xd3, 0xe8, 0x82, 0xa4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for JobService service

type JobServiceClient interface {
	// Queries
	GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*JobResponse, error)
}

type jobServiceClient struct {
	c           client.Client
	serviceName string
}

func NewJobServiceClient(serviceName string, c client.Client) JobServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "job"
	}
	return &jobServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*JobResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JobService.GetJob", in)
	out := new(JobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobService service

type JobServiceHandler interface {
	// Queries
	GetJob(context.Context, *GetJobRequest, *JobResponse) error
}

func RegisterJobServiceHandler(s server.Server, hdlr JobServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&JobService{hdlr}, opts...))
}

type JobService struct {
	JobServiceHandler
}

func (h *JobService) GetJob(ctx context.Context, in *GetJobRequest, out *JobResponse) error {
	return h.JobServiceHandler.GetJob(ctx, in, out)
}
