// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/util/intstr/generated.proto

package intstr

import (
	fmt "fmt"

	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

func (m *IntOrString) Reset()      { *m = IntOrString{} }
func (*IntOrString) ProtoMessage() {}
func (*IntOrString) Descriptor() ([]byte, []int) {
	return fileDescriptor_8984be45904ea297, []int{0}
}
func (m *IntOrString) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntOrString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IntOrString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntOrString.Merge(m, src)
}
func (m *IntOrString) XXX_Size() int {
	return m.Size()
}
func (m *IntOrString) XXX_DiscardUnknown() {
	xxx_messageInfo_IntOrString.DiscardUnknown(m)
}

var xxx_messageInfo_IntOrString proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IntOrString)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.apis.util.intstr.IntOrString")
}

func init() {
	proto.RegisterFile("github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/util/intstr/generated.proto", fileDescriptor_8984be45904ea297)
}

var fileDescriptor_8984be45904ea297 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x6d, 0x5a, 0x22, 0x08, 0x12, 0x43, 0xc4, 0x50, 0x31, 0x38, 0x11, 0x03, 0xca, 0x00,
	0xf6, 0x5a, 0x75, 0xcc, 0x56, 0x16, 0xa4, 0x14, 0x31, 0xb0, 0xa5, 0x25, 0x18, 0x2b, 0x69, 0x62,
	0xd9, 0x97, 0xa1, 0x5b, 0x1f, 0x01, 0x36, 0x46, 0x1e, 0x27, 0x63, 0xc7, 0x0e, 0xa8, 0x22, 0xe6,
	0x2d, 0x98, 0x50, 0x9c, 0x40, 0x99, 0xfe, 0xff, 0xee, 0x3e, 0x7d, 0x89, 0xdd, 0x1b, 0x2e, 0xe0,
	0xb9, 0x9a, 0xd3, 0x45, 0xb9, 0x64, 0x0b, 0x91, 0x8b, 0xea, 0x2f, 0x64, 0xc6, 0x59, 0x36, 0xd6,
	0x4c, 0xe7, 0x62, 0x69, 0x4b, 0x22, 0x85, 0x66, 0x15, 0x88, 0x9c, 0x89, 0x02, 0x34, 0x28, 0xc6,
	0xd3, 0x22, 0x55, 0x09, 0xa4, 0x8f, 0x54, 0xaa, 0x12, 0x4a, 0x6f, 0xb2, 0x77, 0xd1, 0x4e, 0xf2,
	0x1b, 0x32, 0xe3, 0x34, 0x1b, 0x6b, 0xda, 0xba, 0x6c, 0x69, 0x5d, 0xb4, 0x75, 0xd1, 0xce, 0x75,
	0x7e, 0xfd, 0xef, 0x3f, 0x78, 0xc9, 0x4b, 0x66, 0x95, 0xf3, 0xea, 0xc9, 0x4e, 0x76, 0xb0, 0xad,
	0xfb, 0xd4, 0xc5, 0x2b, 0x76, 0x4f, 0xa6, 0x05, 0xdc, 0xaa, 0x19, 0x28, 0x51, 0x70, 0x2f, 0x74,
	0x87, 0xb0, 0x92, 0xe9, 0x08, 0x07, 0x38, 0x1c, 0x44, 0x67, 0xf5, 0xce, 0x47, 0x66, 0xe7, 0x0f,
	0xef, 0x56, 0x32, 0xfd, 0xee, 0x33, 0xb6, 0x84, 0x77, 0xe9, 0x3a, 0xa2, 0x80, 0xfb, 0x24, 0x1f,
	0x1d, 0x04, 0x38, 0x3c, 0x8c, 0x4e, 0x7b, 0xd6, 0x99, 0xda, 0x6d, 0xdc, 0x5f, 0x5b, 0x4e, 0x83,
	0x6a, 0xb9, 0x41, 0x80, 0xc3, 0xe3, 0x3d, 0x37, 0xb3, 0xdb, 0xb8, 0xbf, 0x4e, 0x8e, 0xde, 0xde,
	0x7d, 0xb4, 0xfe, 0x08, 0x50, 0x74, 0x55, 0x37, 0x04, 0x6d, 0x1a, 0x82, 0xb6, 0x0d, 0x41, 0x6b,
	0x43, 0x70, 0x6d, 0x08, 0xde, 0x18, 0x82, 0xb7, 0x86, 0xe0, 0x4f, 0x43, 0xf0, 0xcb, 0x17, 0x41,
	0x0f, 0x4e, 0xf7, 0xe0, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xe1, 0xc4, 0x32, 0x79, 0x01,
	0x00, 0x00,
}

func (m *IntOrString) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntOrString) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntOrString) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.StrVal)
	copy(dAtA[i:], m.StrVal)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.StrVal)))
	i--
	dAtA[i] = 0x1a
	i = encodeVarintGenerated(dAtA, i, uint64(m.IntVal))
	i--
	dAtA[i] = 0x10
	i = encodeVarintGenerated(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IntOrString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Type))
	n += 1 + sovGenerated(uint64(m.IntVal))
	l = len(m.StrVal)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IntOrString) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IntOrString: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntOrString: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntVal", wireType)
			}
			m.IntVal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntVal |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StrVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
