// Code generated by protoc-gen-go.
// source: company.proto
// DO NOT EDIT!

/*
Package company is a generated protocol buffer package.

It is generated from these files:
	company.proto

It has these top-level messages:
	CompanyItem
	Session
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
*/
package company

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion1

type CompanyItem struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Owner   string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	About   string `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	City    string `protobuf:"bytes,6,opt,name=city" json:"city,omitempty"`
}

func (m *CompanyItem) Reset()                    { *m = CompanyItem{} }
func (m *CompanyItem) String() string            { return proto.CompactTextString(m) }
func (*CompanyItem) ProtoMessage()               {}
func (*CompanyItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Session struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Created  int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Expires  int64  `protobuf:"varint,4,opt,name=expires" json:"expires,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateRequest struct {
	Company *CompanyItem `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetCompany() *CompanyItem {
	if m != nil {
		return m.Company
	}
	return nil
}

type CreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ReadResponse struct {
	Company *CompanyItem `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetCompany() *CompanyItem {
	if m != nil {
		return m.Company
	}
	return nil
}

type UpdateRequest struct {
	Company *CompanyItem `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateRequest) GetCompany() *CompanyItem {
	if m != nil {
		return m.Company
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*CompanyItem)(nil), "CompanyItem")
	proto.RegisterType((*Session)(nil), "Session")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Company service

type CompanyClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
}

type companyClient struct {
	c           client.Client
	serviceName string
}

func NewCompanyClient(serviceName string, c client.Client) CompanyClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "company"
	}
	return &companyClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *companyClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Company.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Company.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Company service

type CompanyHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
}

func RegisterCompanyHandler(s server.Server, hdlr CompanyHandler) {
	s.Handle(s.NewHandler(&Company{hdlr}))
}

type Company struct {
	CompanyHandler
}

func (h *Company) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.CompanyHandler.Create(ctx, in, out)
}

func (h *Company) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.CompanyHandler.Read(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0xed, 0x7f, 0xbf, 0xef, 0xb6, 0x33, 0x4a, 0x70, 0x11, 0x0a, 0x62, 0x09, 0x28, 0x82, 0x90,
	0x45, 0x05, 0x7d, 0x80, 0xba, 0x71, 0x3b, 0xe2, 0xd2, 0x45, 0xda, 0xdc, 0xc5, 0x80, 0x33, 0x19,
	0x27, 0x19, 0xb4, 0xcf, 0xe0, 0x4b, 0x9b, 0x9f, 0x89, 0x4e, 0x2b, 0x82, 0xb8, 0xcb, 0x39, 0xf7,
	0xe7, 0x9c, 0x7b, 0x08, 0x24, 0x5b, 0x55, 0x54, 0xa2, 0xdc, 0xf1, 0xaa, 0x56, 0x46, 0xb1, 0xf7,
	0x3e, 0xcc, 0xd6, 0x81, 0xb9, 0x37, 0x58, 0x90, 0x14, 0x06, 0xb9, 0xa4, 0xfd, 0x65, 0xff, 0xf2,
	0x7f, 0x66, 0x5f, 0x84, 0xc0, 0xa8, 0x14, 0x05, 0xd2, 0x81, 0x67, 0xfc, 0x9b, 0x9c, 0xc0, 0x58,
	0xbd, 0x96, 0x58, 0xd3, 0xa1, 0x27, 0x03, 0x70, 0xac, 0xd8, 0xa8, 0xc6, 0xd0, 0x51, 0x60, 0x3d,
	0x20, 0x14, 0xa6, 0x42, 0xca, 0x1a, 0xb5, 0xa6, 0x63, 0xcf, 0x47, 0xe8, 0x36, 0x6f, 0x73, 0xb3,
	0xa3, 0x93, 0xb0, 0xd9, 0xbd, 0x59, 0x0e, 0xd3, 0x07, 0x5b, 0xcb, 0x55, 0xf9, 0xcd, 0xc8, 0x02,
	0xfe, 0x35, 0x1a, 0xeb, 0x8e, 0x99, 0x4f, 0xec, 0x44, 0xb6, 0x35, 0x0a, 0x83, 0xd2, 0x5b, 0x1a,
	0x66, 0x11, 0xba, 0x0a, 0xbe, 0x55, 0xb9, 0x15, 0xf4, 0xb6, 0x6c, 0xa5, 0x85, 0xec, 0x16, 0x92,
	0xb5, 0x6f, 0xca, 0xf0, 0xa5, 0x41, 0x6d, 0xc8, 0x85, 0x5d, 0x12, 0x82, 0xf0, 0xaa, 0xb3, 0xd5,
	0x9c, 0x77, 0x82, 0xc9, 0x62, 0x91, 0x2d, 0x21, 0x8d, 0x83, 0xba, 0x52, 0xa5, 0xc6, 0x43, 0xab,
	0xec, 0x0c, 0x92, 0x3b, 0x7c, 0xc6, 0xaf, 0xd5, 0x87, 0x0d, 0xc7, 0x90, 0xc6, 0x86, 0xb0, 0x82,
	0x9d, 0xc2, 0x2c, 0x43, 0x21, 0x7f, 0x1a, 0xb8, 0x81, 0x79, 0x28, 0xb7, 0x8a, 0xbf, 0xf5, 0x6a,
	0x8f, 0x7c, 0xac, 0xe4, 0x1f, 0x8e, 0xb4, 0x0e, 0xe3, 0x60, 0x90, 0x5c, 0x3d, 0xc1, 0xb4, 0xed,
	0x24, 0x57, 0x30, 0x09, 0x09, 0x90, 0x94, 0xef, 0x65, 0xb8, 0x38, 0xe2, 0xfb, 0xd1, 0xb0, 0x1e,
	0x39, 0x87, 0x91, 0xb3, 0x4e, 0xe6, 0xbc, 0x73, 0xe0, 0x22, 0xe1, 0xdd, 0x7b, 0x58, 0x6f, 0x33,
	0xf1, 0xdf, 0xf1, 0xfa, 0x23, 0x00, 0x00, 0xff, 0xff, 0xed, 0x67, 0x74, 0xb0, 0x9f, 0x02, 0x00,
	0x00,
}
