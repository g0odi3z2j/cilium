/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/certificates/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *ClusterTrustBundle) Reset()      { *m = ClusterTrustBundle{} }
func (*ClusterTrustBundle) ProtoMessage() {}
func (*ClusterTrustBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_8915b0d419f9eda6, []int{0}
}
func (m *ClusterTrustBundle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterTrustBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterTrustBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterTrustBundle.Merge(m, src)
}
func (m *ClusterTrustBundle) XXX_Size() int {
	return m.Size()
}
func (m *ClusterTrustBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterTrustBundle.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterTrustBundle proto.InternalMessageInfo

func (m *ClusterTrustBundleList) Reset()      { *m = ClusterTrustBundleList{} }
func (*ClusterTrustBundleList) ProtoMessage() {}
func (*ClusterTrustBundleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8915b0d419f9eda6, []int{1}
}
func (m *ClusterTrustBundleList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterTrustBundleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterTrustBundleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterTrustBundleList.Merge(m, src)
}
func (m *ClusterTrustBundleList) XXX_Size() int {
	return m.Size()
}
func (m *ClusterTrustBundleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterTrustBundleList.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterTrustBundleList proto.InternalMessageInfo

func (m *ClusterTrustBundleSpec) Reset()      { *m = ClusterTrustBundleSpec{} }
func (*ClusterTrustBundleSpec) ProtoMessage() {}
func (*ClusterTrustBundleSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8915b0d419f9eda6, []int{2}
}
func (m *ClusterTrustBundleSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterTrustBundleSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterTrustBundleSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterTrustBundleSpec.Merge(m, src)
}
func (m *ClusterTrustBundleSpec) XXX_Size() int {
	return m.Size()
}
func (m *ClusterTrustBundleSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterTrustBundleSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterTrustBundleSpec proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClusterTrustBundle)(nil), "k8s.io.api.certificates.v1alpha1.ClusterTrustBundle")
	proto.RegisterType((*ClusterTrustBundleList)(nil), "k8s.io.api.certificates.v1alpha1.ClusterTrustBundleList")
	proto.RegisterType((*ClusterTrustBundleSpec)(nil), "k8s.io.api.certificates.v1alpha1.ClusterTrustBundleSpec")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/certificates/v1alpha1/generated.proto", fileDescriptor_8915b0d419f9eda6)
}

var fileDescriptor_8915b0d419f9eda6 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0x6a, 0x0b, 0xed, 0x44, 0x41, 0x56, 0x90, 0x90, 0xc3, 0x34, 0xe4, 0xd4, 0x8b,
	0x33, 0x26, 0x54, 0xe9, 0x79, 0x05, 0xa1, 0xe0, 0x0f, 0xd8, 0x7a, 0xb1, 0x78, 0x70, 0x32, 0x79,
	0xdd, 0x8c, 0xc9, 0xee, 0x0e, 0x33, 0xb3, 0x01, 0x6f, 0x82, 0xff, 0x80, 0x7f, 0x56, 0x8e, 0xd5,
	0x53, 0x4f, 0xc5, 0xac, 0xff, 0x88, 0xcc, 0x64, 0x93, 0x5d, 0x5c, 0x25, 0xd2, 0xdb, 0xbe, 0x1f,
	0x9f, 0xef, 0x7b, 0xdf, 0xb7, 0x0c, 0x3e, 0x9f, 0x9d, 0x19, 0x2a, 0x73, 0x36, 0x2b, 0xc6, 0xa0,
	0x33, 0xb0, 0x60, 0xd8, 0x02, 0xb2, 0x49, 0xae, 0x59, 0x55, 0xe0, 0x4a, 0x32, 0x01, 0xda, 0xca,
	0x2b, 0x29, 0xb8, 0x2f, 0x0f, 0xf9, 0x5c, 0x4d, 0xf9, 0x90, 0x25, 0x90, 0x81, 0xe6, 0x16, 0x26,
	0x54, 0xe9, 0xdc, 0xe6, 0x61, 0x7f, 0x4d, 0x50, 0xae, 0x24, 0x6d, 0x12, 0x74, 0x43, 0xf4, 0x9e,
	0x24, 0xd2, 0x4e, 0x8b, 0x31, 0x15, 0x79, 0xca, 0x92, 0x3c, 0xc9, 0x99, 0x07, 0xc7, 0xc5, 0x95,
	0x8f, 0x7c, 0xe0, 0xbf, 0xd6, 0x82, 0xbd, 0xd3, 0x7a, 0x85, 0x94, 0x8b, 0xa9, 0xcc, 0x40, 0x7f,
	0x66, 0x6a, 0x96, 0xb8, 0x84, 0x61, 0x29, 0x58, 0xce, 0x16, 0xad, 0x35, 0x7a, 0xec, 0x5f, 0x94,
	0x2e, 0x32, 0x2b, 0x53, 0x68, 0x01, 0xcf, 0x77, 0x01, 0x46, 0x4c, 0x21, 0xe5, 0x7f, 0x72, 0x83,
	0x1f, 0x08, 0x87, 0x2f, 0xe6, 0x85, 0xb1, 0xa0, 0xdf, 0xe9, 0xc2, 0xd8, 0xa8, 0xc8, 0x26, 0x73,
	0x08, 0x3f, 0xe2, 0x43, 0xb7, 0xda, 0x84, 0x5b, 0xde, 0x45, 0x7d, 0x74, 0xd2, 0x19, 0x3d, 0xa5,
	0xf5, 0x65, 0xb6, 0x13, 0xa8, 0x9a, 0x25, 0x2e, 0x61, 0xa8, 0xeb, 0xa6, 0x8b, 0x21, 0x7d, 0x3b,
	0xfe, 0x04, 0xc2, 0xbe, 0x06, 0xcb, 0xa3, 0x70, 0x79, 0x7b, 0x1c, 0x94, 0xb7, 0xc7, 0xb8, 0xce,
	0xc5, 0x5b, 0xd5, 0xf0, 0x12, 0xef, 0x1b, 0x05, 0xa2, 0xbb, 0xe7, 0xd5, 0xcf, 0xe8, 0xae, 0xbb,
	0xd3, 0xf6, 0x96, 0x17, 0x0a, 0x44, 0x74, 0xbf, 0x9a, 0xb2, 0xef, 0xa2, 0xd8, 0x6b, 0x0e, 0xbe,
	0x23, 0xfc, 0xb8, 0xdd, 0xfe, 0x4a, 0x1a, 0x1b, 0x7e, 0x68, 0x19, 0xa3, 0xff, 0x67, 0xcc, 0xd1,
	0xde, 0xd6, 0xc3, 0x6a, 0xe0, 0xe1, 0x26, 0xd3, 0x30, 0xf5, 0x1e, 0x1f, 0x48, 0x0b, 0xa9, 0xe9,
	0xee, 0xf5, 0xef, 0x9d, 0x74, 0x46, 0xa7, 0x77, 0x71, 0x15, 0x3d, 0xa8, 0x06, 0x1c, 0x9c, 0x3b,
	0xa9, 0x78, 0xad, 0x38, 0xf8, 0xfa, 0x57, 0x4f, 0xce, 0x74, 0x38, 0xc2, 0xd8, 0xc8, 0x24, 0x03,
	0xfd, 0x86, 0xa7, 0xe0, 0x5d, 0x1d, 0xd5, 0xc7, 0xbf, 0xd8, 0x56, 0xe2, 0x46, 0x57, 0xf8, 0x0c,
	0x77, 0x6c, 0x2d, 0xe3, 0xff, 0xc2, 0x51, 0xf4, 0xa8, 0x82, 0x3a, 0x8d, 0x09, 0x71, 0xb3, 0x2f,
	0x7a, 0xb9, 0x5c, 0x91, 0xe0, 0x7a, 0x45, 0x82, 0x9b, 0x15, 0x09, 0xbe, 0x94, 0x04, 0x2d, 0x4b,
	0x82, 0xae, 0x4b, 0x82, 0x6e, 0x4a, 0x82, 0x7e, 0x96, 0x04, 0x7d, 0xfb, 0x45, 0x82, 0xcb, 0xfe,
	0xae, 0x67, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x05, 0xe9, 0xaa, 0x07, 0xb2, 0x03, 0x00, 0x00,
}

func (m *ClusterTrustBundle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterTrustBundle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterTrustBundle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClusterTrustBundleList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterTrustBundleList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterTrustBundleList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClusterTrustBundleSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterTrustBundleSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterTrustBundleSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.TrustBundle)
	copy(dAtA[i:], m.TrustBundle)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.TrustBundle)))
	i--
	dAtA[i] = 0x12
	i -= len(m.SignerName)
	copy(dAtA[i:], m.SignerName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.SignerName)))
	i--
	dAtA[i] = 0xa
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
func (m *ClusterTrustBundle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ClusterTrustBundleList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ClusterTrustBundleSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SignerName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.TrustBundle)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterTrustBundle) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterTrustBundle{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ClusterTrustBundleSpec", "ClusterTrustBundleSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterTrustBundleList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]ClusterTrustBundle{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "ClusterTrustBundle", "ClusterTrustBundle", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&ClusterTrustBundleList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterTrustBundleSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterTrustBundleSpec{`,
		`SignerName:` + fmt.Sprintf("%v", this.SignerName) + `,`,
		`TrustBundle:` + fmt.Sprintf("%v", this.TrustBundle) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClusterTrustBundle) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClusterTrustBundle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterTrustBundle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *ClusterTrustBundleList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClusterTrustBundleList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterTrustBundleList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, ClusterTrustBundle{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *ClusterTrustBundleSpec) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClusterTrustBundleSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterTrustBundleSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerName", wireType)
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
			m.SignerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustBundle", wireType)
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
			m.TrustBundle = string(dAtA[iNdEx:postIndex])
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
