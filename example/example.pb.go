// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stevvooe/mgrpc/example/example.proto

/*
	Package example is a generated protocol buffer package.

	It is generated from these files:
		github.com/stevvooe/mgrpc/example/example.proto

	It has these top-level messages:
		Method1Request
		Method1Response
		Method2Request
*/
package example

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import context "context"
import github_com_stevvooe_mgrpc "github.com/stevvooe/mgrpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Method1Request struct {
	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar string `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (m *Method1Request) Reset()                    { *m = Method1Request{} }
func (*Method1Request) ProtoMessage()               {}
func (*Method1Request) Descriptor() ([]byte, []int) { return fileDescriptorExample, []int{0} }

type Method1Response struct {
	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar string `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (m *Method1Response) Reset()                    { *m = Method1Response{} }
func (*Method1Response) ProtoMessage()               {}
func (*Method1Response) Descriptor() ([]byte, []int) { return fileDescriptorExample, []int{1} }

type Method2Request struct {
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
}

func (m *Method2Request) Reset()                    { *m = Method2Request{} }
func (*Method2Request) ProtoMessage()               {}
func (*Method2Request) Descriptor() ([]byte, []int) { return fileDescriptorExample, []int{2} }

func init() {
	proto.RegisterType((*Method1Request)(nil), "mgrpc.example.v1.Method1Request")
	proto.RegisterType((*Method1Response)(nil), "mgrpc.example.v1.Method1Response")
	proto.RegisterType((*Method2Request)(nil), "mgrpc.example.v1.Method2Request")
}
func (m *Method1Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Method1Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Foo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Foo)))
		i += copy(dAtA[i:], m.Foo)
	}
	if len(m.Bar) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Bar)))
		i += copy(dAtA[i:], m.Bar)
	}
	return i, nil
}

func (m *Method1Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Method1Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Foo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Foo)))
		i += copy(dAtA[i:], m.Foo)
	}
	if len(m.Bar) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Bar)))
		i += copy(dAtA[i:], m.Bar)
	}
	return i, nil
}

func (m *Method2Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Method2Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExample(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	return i, nil
}

func encodeFixed64Example(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Example(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintExample(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

type ExampleService interface {
	Method1(ctx context.Context, req *Method1Request) (*Method1Response, error)
	Method2(ctx context.Context, req *Method1Request) (*google_protobuf1.Empty, error)
}

func RegisterExampleService(srv *github_com_stevvooe_mgrpc.Server, svc ExampleService) error {
	return srv.Register("mgrpc.example.v1.Example", map[string]github_com_stevvooe_mgrpc.Handler{
		"Method1": github_com_stevvooe_mgrpc.HandlerFunc(func(ctx context.Context, req interface{}) (interface{}, error) {
			return svc.Method1(ctx, req.(*Method1Request))
		}),
		"Method2": github_com_stevvooe_mgrpc.HandlerFunc(func(ctx context.Context, req interface{}) (interface{}, error) {
			return svc.Method2(ctx, req.(*Method1Request))
		}),
	})
}

type ExampleClient struct {
	client *github_com_stevvooe_mgrpc.Client
}

func (c *ExampleClient) Method1(ctx context.Context, req *Method1Request) (*Method1Response, error) {
	resp, err := c.client.Call(ctx, "mgrpc.example.v1.Example", "Method1", req)
	if err != nil {
		return nil, err
	}
	return resp.(*Method1Response), nil
}

func (c *ExampleClient) Method2(ctx context.Context, req *Method1Request) (*google_protobuf1.Empty, error) {
	resp, err := c.client.Call(ctx, "mgrpc.example.v1.Example", "Method2", req)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf1.Empty), nil
}
func (m *Method1Request) Size() (n int) {
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Bar)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func (m *Method1Response) Size() (n int) {
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	l = len(m.Bar)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func (m *Method2Request) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	return n
}

func sovExample(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExample(x uint64) (n int) {
	return sovExample(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Method1Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Method1Request{`,
		`Foo:` + fmt.Sprintf("%v", this.Foo) + `,`,
		`Bar:` + fmt.Sprintf("%v", this.Bar) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Method1Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Method1Response{`,
		`Foo:` + fmt.Sprintf("%v", this.Foo) + `,`,
		`Bar:` + fmt.Sprintf("%v", this.Bar) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Method2Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Method2Request{`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringExample(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Method1Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Method1Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Method1Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Foo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Foo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *Method1Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Method1Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Method1Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Foo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Foo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func (m *Method2Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Method2Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Method2Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExample
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
func skipExample(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExample
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExample
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExample
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExample(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExample = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExample   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/stevvooe/mgrpc/example/example.proto", fileDescriptorExample)
}

var fileDescriptorExample = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xad, 0x41, 0x6a, 0x85, 0x07, 0xa8, 0x22, 0x54, 0x95, 0x20, 0x99, 0xd2, 0xa9, 0x2c, 0xb6,
	0x1a, 0x60, 0x62, 0xa3, 0xea, 0x08, 0x43, 0x47, 0x36, 0x27, 0xb8, 0x6e, 0xa4, 0x3a, 0x67, 0x12,
	0x27, 0x82, 0x8d, 0x2f, 0xe1, 0x7b, 0x3a, 0x32, 0x32, 0xd2, 0x7c, 0x09, 0x4a, 0x62, 0x23, 0x91,
	0x81, 0x76, 0xba, 0xf3, 0x7b, 0xf7, 0xee, 0x9d, 0x9e, 0x8c, 0x99, 0x8c, 0xcd, 0x2a, 0x0f, 0x69,
	0x04, 0x8a, 0x65, 0x46, 0x14, 0x05, 0x80, 0x60, 0x4a, 0xa6, 0x3a, 0x62, 0xe2, 0x95, 0x2b, 0xbd,
	0x16, 0xae, 0x52, 0x9d, 0x82, 0x01, 0xaf, 0x5f, 0x93, 0xd4, 0x81, 0xc5, 0xd4, 0x3f, 0x93, 0x00,
	0x72, 0x2d, 0x58, 0xcd, 0x87, 0xf9, 0x92, 0xf1, 0xe4, 0xad, 0x19, 0xf6, 0xcf, 0xdb, 0x94, 0x50,
	0xda, 0x38, 0xf2, 0x54, 0x82, 0x84, 0xba, 0x65, 0x55, 0x67, 0xd1, 0x8b, 0xb6, 0xc4, 0xc4, 0x4a,
	0x64, 0x86, 0x2b, 0xdd, 0x0c, 0x8c, 0x6f, 0xf0, 0xf1, 0x83, 0x30, 0x2b, 0x78, 0x9e, 0x2e, 0xc4,
	0x4b, 0x2e, 0x32, 0xe3, 0xf5, 0xf1, 0xe1, 0x12, 0x60, 0x88, 0x46, 0x68, 0x72, 0xb4, 0xa8, 0xda,
	0x0a, 0x09, 0x79, 0x3a, 0x3c, 0x68, 0x90, 0x90, 0xa7, 0xe3, 0x5b, 0x7c, 0xf2, 0xab, 0xca, 0x34,
	0x24, 0x99, 0xd8, 0x4b, 0x36, 0x71, 0x66, 0x81, 0x33, 0x1b, 0xe0, 0x2e, 0x8f, 0x4c, 0x0c, 0x89,
	0x15, 0xda, 0x57, 0xf0, 0x81, 0x70, 0x6f, 0xde, 0x84, 0xe2, 0x3d, 0xe2, 0x9e, 0x35, 0xf3, 0x46,
	0xb4, 0x9d, 0x17, 0xfd, 0x7b, 0xbd, 0x7f, 0xf9, 0xcf, 0x84, 0xbd, 0x74, 0xe6, 0xf6, 0x05, 0x7b,
	0xec, 0x1b, 0xd0, 0x26, 0x41, 0xea, 0x12, 0xa4, 0xf3, 0x2a, 0xf4, 0xfb, 0xd9, 0x66, 0x4b, 0x3a,
	0x5f, 0x5b, 0xd2, 0x79, 0x2f, 0x09, 0xda, 0x94, 0x04, 0x7d, 0x96, 0x04, 0x7d, 0x97, 0x04, 0x3d,
	0x5d, 0xed, 0xfc, 0x03, 0x77, 0xb6, 0x86, 0xdd, 0x7a, 0xe9, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x3f, 0x4b, 0x81, 0x37, 0x02, 0x00, 0x00,
}
