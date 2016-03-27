// Code generated by protoc-gen-go.
// source: company.proto
// DO NOT EDIT!

/*
Package company is a generated protocol buffer package.

It is generated from these files:
	company.proto

It has these top-level messages:
	Company
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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Company struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Owner   string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	About   string `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
}

func (m *Company) Reset()                    { *m = Company{} }
func (m *Company) String() string            { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()               {}
func (*Company) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Company *Company `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

type CreateResponse struct {
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
	Company *Company `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

type UpdateRequest struct {
	Company *Company `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateRequest) GetCompany() *Company {
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
	proto.RegisterType((*Company)(nil), "Company")
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

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0xdd, 0x8f, 0xee, 0x76, 0x9d, 0xdd, 0x56, 0x09, 0x1e, 0x42, 0x41, 0x94, 0x80, 0x20, 0x08,
	0x39, 0x74, 0x7f, 0x81, 0xac, 0xbf, 0xa0, 0xe2, 0xd1, 0x43, 0xb6, 0x99, 0x43, 0xc1, 0x4d, 0x6a,
	0xd3, 0xa2, 0xfe, 0x7b, 0xd3, 0xa4, 0xd1, 0xba, 0x22, 0xe8, 0xa9, 0x79, 0x6f, 0xe6, 0xcd, 0x9b,
	0x37, 0x14, 0x92, 0x52, 0x1f, 0x6a, 0xa1, 0xde, 0x79, 0xdd, 0xe8, 0x56, 0x33, 0x03, 0xf1, 0xce,
	0x13, 0x24, 0x85, 0x59, 0x25, 0xe9, 0xf4, 0x6a, 0x7a, 0x73, 0x52, 0xd8, 0x17, 0x21, 0x10, 0x29,
	0x71, 0x40, 0x3a, 0x73, 0x8c, 0x7b, 0x93, 0x73, 0x58, 0xe8, 0x57, 0x85, 0x0d, 0x9d, 0x3b, 0xd2,
	0x83, 0x9e, 0x15, 0x7b, 0xdd, 0xb5, 0x34, 0xf2, 0xac, 0x03, 0x84, 0x42, 0x2c, 0xa4, 0x6c, 0xd0,
	0x18, 0xba, 0x70, 0x7c, 0x80, 0xac, 0x82, 0xf8, 0xc1, 0x7e, 0x2b, 0xad, 0x7e, 0x98, 0x66, 0xb0,
	0xea, 0x0c, 0x36, 0x23, 0xe3, 0x4f, 0xdc, 0x0f, 0x2c, 0x1b, 0x14, 0x2d, 0x4a, 0x67, 0x3f, 0x2f,
	0x02, 0xec, 0x2b, 0xf8, 0x56, 0x57, 0x76, 0xb8, 0x5b, 0xc1, 0x56, 0x06, 0xc8, 0xb6, 0x90, 0xec,
	0x5c, 0x53, 0x81, 0x2f, 0x1d, 0x9a, 0x96, 0x30, 0x3b, 0xc4, 0x07, 0x76, 0xae, 0xeb, 0x7c, 0xc5,
	0x87, 0x03, 0x14, 0xa1, 0xc0, 0xce, 0x20, 0x0d, 0x22, 0x53, 0x6b, 0x65, 0x90, 0x5d, 0x42, 0x72,
	0x8f, 0xcf, 0xf8, 0x35, 0xe6, 0x68, 0xef, 0x5e, 0x12, 0x1a, 0x06, 0xc9, 0x05, 0xac, 0x0b, 0x14,
	0xf2, 0x37, 0x41, 0x0e, 0x1b, 0x5f, 0xf6, 0xed, 0x7f, 0xda, 0xcb, 0x86, 0x79, 0xac, 0xe5, 0xff,
	0xc3, 0x04, 0x91, 0xb7, 0xca, 0x9f, 0x20, 0xbe, 0x2b, 0x4b, 0xdd, 0xa9, 0x96, 0xdc, 0xc2, 0xd2,
	0x27, 0x25, 0x29, 0xff, 0x76, 0xa7, 0xec, 0x94, 0x1f, 0x9d, 0x60, 0x42, 0xae, 0x21, 0xea, 0x57,
	0x26, 0x1b, 0x3e, 0x0a, 0x96, 0x25, 0x7c, 0x9c, 0x83, 0x4d, 0xf6, 0x4b, 0xf7, 0x67, 0x6d, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xba, 0x64, 0x3e, 0x6a, 0x02, 0x00, 0x00,
}
