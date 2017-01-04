/*
Copyright 2017 The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/certificates/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/certificates/v1alpha1/generated.proto

	It has these top-level messages:
		CertificateSigningRequest
		CertificateSigningRequestCondition
		CertificateSigningRequestList
		CertificateSigningRequestSpec
		CertificateSigningRequestStatus
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *CertificateSigningRequest) Reset()      { *m = CertificateSigningRequest{} }
func (*CertificateSigningRequest) ProtoMessage() {}
func (*CertificateSigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{0}
}

func (m *CertificateSigningRequestCondition) Reset()      { *m = CertificateSigningRequestCondition{} }
func (*CertificateSigningRequestCondition) ProtoMessage() {}
func (*CertificateSigningRequestCondition) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *CertificateSigningRequestList) Reset()      { *m = CertificateSigningRequestList{} }
func (*CertificateSigningRequestList) ProtoMessage() {}
func (*CertificateSigningRequestList) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{2}
}

func (m *CertificateSigningRequestSpec) Reset()      { *m = CertificateSigningRequestSpec{} }
func (*CertificateSigningRequestSpec) ProtoMessage() {}
func (*CertificateSigningRequestSpec) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{3}
}

func (m *CertificateSigningRequestStatus) Reset()      { *m = CertificateSigningRequestStatus{} }
func (*CertificateSigningRequestStatus) ProtoMessage() {}
func (*CertificateSigningRequestStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{4}
}

func init() {
	proto.RegisterType((*CertificateSigningRequest)(nil), "k8s.io.kubernetes.pkg.apis.certificates.v1alpha1.CertificateSigningRequest")
	proto.RegisterType((*CertificateSigningRequestCondition)(nil), "k8s.io.kubernetes.pkg.apis.certificates.v1alpha1.CertificateSigningRequestCondition")
	proto.RegisterType((*CertificateSigningRequestList)(nil), "k8s.io.kubernetes.pkg.apis.certificates.v1alpha1.CertificateSigningRequestList")
	proto.RegisterType((*CertificateSigningRequestSpec)(nil), "k8s.io.kubernetes.pkg.apis.certificates.v1alpha1.CertificateSigningRequestSpec")
	proto.RegisterType((*CertificateSigningRequestStatus)(nil), "k8s.io.kubernetes.pkg.apis.certificates.v1alpha1.CertificateSigningRequestStatus")
}
func (m *CertificateSigningRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CertificateSigningRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *CertificateSigningRequestCondition) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CertificateSigningRequestCondition) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Type)))
	i += copy(data[i:], m.Type)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Reason)))
	i += copy(data[i:], m.Reason)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Message)))
	i += copy(data[i:], m.Message)
	data[i] = 0x22
	i++
	i = encodeVarintGenerated(data, i, uint64(m.LastUpdateTime.Size()))
	n4, err := m.LastUpdateTime.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *CertificateSigningRequestList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CertificateSigningRequestList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n5, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CertificateSigningRequestSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CertificateSigningRequestSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		data[i] = 0xa
		i++
		i = encodeVarintGenerated(data, i, uint64(len(m.Request)))
		i += copy(data[i:], m.Request)
	}
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Username)))
	i += copy(data[i:], m.Username)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.UID)))
	i += copy(data[i:], m.UID)
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			data[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *CertificateSigningRequestStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CertificateSigningRequestStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for _, msg := range m.Conditions {
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Certificate != nil {
		data[i] = 0x12
		i++
		i = encodeVarintGenerated(data, i, uint64(len(m.Certificate)))
		i += copy(data[i:], m.Certificate)
	}
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *CertificateSigningRequest) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CertificateSigningRequestCondition) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Message)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.LastUpdateTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CertificateSigningRequestList) Size() (n int) {
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

func (m *CertificateSigningRequestSpec) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = len(m.Request)
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Username)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *CertificateSigningRequestStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.Certificate != nil {
		l = len(m.Certificate)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CertificateSigningRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CertificateSigningRequest{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_kubernetes_pkg_api_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "CertificateSigningRequestSpec", "CertificateSigningRequestSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "CertificateSigningRequestStatus", "CertificateSigningRequestStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CertificateSigningRequestCondition) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CertificateSigningRequestCondition{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`LastUpdateTime:` + strings.Replace(strings.Replace(this.LastUpdateTime.String(), "Time", "k8s_io_kubernetes_pkg_api_unversioned.Time", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CertificateSigningRequestList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CertificateSigningRequestList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_kubernetes_pkg_api_unversioned.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "CertificateSigningRequest", "CertificateSigningRequest", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CertificateSigningRequestSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CertificateSigningRequestSpec{`,
		`Request:` + valueToStringGenerated(this.Request) + `,`,
		`Username:` + fmt.Sprintf("%v", this.Username) + `,`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Groups:` + fmt.Sprintf("%v", this.Groups) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CertificateSigningRequestStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CertificateSigningRequestStatus{`,
		`Conditions:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Conditions), "CertificateSigningRequestCondition", "CertificateSigningRequestCondition", 1), `&`, ``, 1) + `,`,
		`Certificate:` + valueToStringGenerated(this.Certificate) + `,`,
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
func (m *CertificateSigningRequest) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateSigningRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateSigningRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *CertificateSigningRequestCondition) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateSigningRequestCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateSigningRequestCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = RequestConditionType(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastUpdateTime.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *CertificateSigningRequestList) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateSigningRequestList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateSigningRequestList: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CertificateSigningRequest{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *CertificateSigningRequestSpec) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateSigningRequestSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateSigningRequestSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], data[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *CertificateSigningRequestStatus) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateSigningRequestStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateSigningRequestStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, CertificateSigningRequestCondition{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = append(m.Certificate[:0], data[iNdEx:postIndex]...)
			if m.Certificate == nil {
				m.Certificate = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipGenerated(data[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xa6, 0xb4, 0x94, 0x32, 0xe5, 0x07, 0xbf, 0x4c, 0x8c, 0xa9, 0x4d, 0xa0, 0x66, 0xa3, 0x06,
	0x15, 0x66, 0x2d, 0x89, 0x09, 0x47, 0xb3, 0x98, 0x18, 0x22, 0x84, 0x38, 0x40, 0x62, 0x4c, 0x3c,
	0x6c, 0xb7, 0xc3, 0x32, 0x96, 0xee, 0x2e, 0xf3, 0x87, 0x84, 0x9b, 0x47, 0x8f, 0x7e, 0x02, 0xbf,
	0x86, 0x5f, 0x81, 0x23, 0x47, 0x4f, 0xa8, 0xf0, 0x05, 0x3c, 0x7b, 0x72, 0x66, 0x76, 0xb6, 0x5d,
	0x5b, 0x16, 0x35, 0xe1, 0x30, 0x49, 0xe7, 0x99, 0xf7, 0x7d, 0xde, 0xf7, 0x7d, 0xde, 0x67, 0x0b,
	0x9e, 0xf5, 0xd6, 0x38, 0xa2, 0xb1, 0xdb, 0x93, 0x1d, 0xc2, 0x22, 0x22, 0x08, 0x77, 0x93, 0x5e,
	0xe8, 0xfa, 0x09, 0xe5, 0x6e, 0x40, 0x98, 0xa0, 0xfb, 0x34, 0xf0, 0x35, 0x7a, 0xdc, 0xf6, 0x0f,
	0x93, 0x03, 0xbf, 0xed, 0x86, 0x24, 0x22, 0x4c, 0x41, 0x5d, 0x94, 0xb0, 0x58, 0xc4, 0xf0, 0x49,
	0xca, 0x80, 0x86, 0x0c, 0x48, 0x31, 0x20, 0xcd, 0x80, 0xf2, 0x0c, 0x28, 0x63, 0x68, 0xae, 0x84,
	0x54, 0x1c, 0xc8, 0x0e, 0x0a, 0xe2, 0xbe, 0x1b, 0xc6, 0x61, 0xec, 0x1a, 0xa2, 0x8e, 0xdc, 0x37,
	0x37, 0x73, 0x31, 0xbf, 0xd2, 0x02, 0xcd, 0xd5, 0xc2, 0x16, 0x5d, 0x46, 0x78, 0x2c, 0x59, 0x40,
	0x46, 0x9b, 0x6a, 0x3e, 0x2d, 0xce, 0x91, 0xd1, 0x31, 0x61, 0x9c, 0xc6, 0x11, 0xe9, 0x8e, 0xa5,
	0x2d, 0x17, 0xa7, 0x1d, 0x8f, 0x4d, 0xde, 0x5c, 0xb9, 0x3a, 0x9a, 0xc9, 0x48, 0xd0, 0xfe, 0x78,
	0x4f, 0xed, 0xab, 0xc3, 0xa5, 0xa0, 0x87, 0x2e, 0x8d, 0x04, 0x17, 0x6c, 0x34, 0xc5, 0xb9, 0x9c,
	0x04, 0x77, 0xd6, 0x87, 0x1a, 0xee, 0xd0, 0x30, 0xa2, 0x51, 0x88, 0xc9, 0x91, 0x24, 0x5c, 0xc0,
	0xd7, 0xa0, 0xd6, 0x27, 0xc2, 0xef, 0xfa, 0xc2, 0x6f, 0x94, 0xee, 0x96, 0x96, 0xea, 0xab, 0x4b,
	0xa8, 0x70, 0x19, 0x4a, 0x7e, 0xb4, 0xdd, 0x79, 0x47, 0x02, 0xb1, 0xa5, 0x72, 0x3c, 0x78, 0x7a,
	0xde, 0x9a, 0xb8, 0x38, 0x6f, 0x81, 0x21, 0x86, 0x07, 0x6c, 0xf0, 0x08, 0x54, 0x78, 0x42, 0x82,
	0xc6, 0xa4, 0x61, 0xdd, 0x46, 0xff, 0xba, 0x62, 0x54, 0xd8, 0xf4, 0x8e, 0xa2, 0xf5, 0x66, 0x6d,
	0xf1, 0x8a, 0xbe, 0x61, 0x53, 0x0a, 0x9e, 0x80, 0x2a, 0x17, 0xbe, 0x90, 0xbc, 0x51, 0x36, 0x45,
	0x5f, 0xdd, 0x64, 0x51, 0x43, 0xec, 0xcd, 0xd9, 0xb2, 0xd5, 0xf4, 0x8e, 0x6d, 0x41, 0xe7, 0xd3,
	0x24, 0x70, 0x0a, 0x73, 0xd7, 0xe3, 0xa8, 0x4b, 0x85, 0xb2, 0x0b, 0x5c, 0x03, 0x15, 0x71, 0x92,
	0x10, 0x23, 0xf5, 0x8c, 0x77, 0x2f, 0x9b, 0x61, 0x57, 0x61, 0x3f, 0xcf, 0x5b, 0xb7, 0x46, 0xe3,
	0x35, 0x8e, 0x4d, 0x06, 0x7c, 0x00, 0xaa, 0x8c, 0xf8, 0x3c, 0x8e, 0x8c, 0xa0, 0x33, 0xc3, 0x46,
	0xb0, 0x41, 0xb1, 0x7d, 0x85, 0x0f, 0xc1, 0x74, 0x9f, 0x70, 0xee, 0x87, 0xc4, 0x88, 0x30, 0xe3,
	0xcd, 0xdb, 0xc0, 0xe9, 0xad, 0x14, 0xc6, 0xd9, 0x3b, 0xec, 0x81, 0xb9, 0x43, 0x9f, 0x8b, 0xbd,
	0x44, 0xed, 0x8b, 0xec, 0x2a, 0xbb, 0x35, 0x2a, 0x46, 0xb6, 0xc7, 0xd7, 0x38, 0x20, 0xe7, 0x7c,
	0xa4, 0x53, 0xbc, 0xdb, 0x96, 0x7e, 0x6e, 0xf3, 0x37, 0x2a, 0x3c, 0x42, 0xed, 0xfc, 0x28, 0x81,
	0x85, 0x42, 0x81, 0x36, 0xa9, 0xb2, 0xe2, 0xdb, 0x31, 0x2b, 0xba, 0x7f, 0xd9, 0x88, 0x4e, 0x37,
	0x8e, 0xfc, 0xdf, 0x36, 0x53, 0xcb, 0x90, 0x9c, 0x1f, 0x13, 0x30, 0x45, 0x05, 0xe9, 0x73, 0xa5,
	0x5f, 0x59, 0x71, 0xbf, 0xbc, 0x41, 0x6f, 0x78, 0xff, 0xd9, 0xba, 0x53, 0x1b, 0xba, 0x02, 0x4e,
	0x0b, 0x39, 0x9f, 0xaf, 0x1b, 0x59, 0xdb, 0x16, 0xde, 0x07, 0xd3, 0x2c, 0xbd, 0x9a, 0x89, 0x67,
	0xbd, 0xba, 0x5e, 0x94, 0x8d, 0xc0, 0xd9, 0x1b, 0x5c, 0x06, 0x35, 0xc9, 0x55, 0x93, 0xbe, 0x5a,
	0x51, 0xba, 0xfd, 0xc1, 0xa0, 0x7b, 0x16, 0xc7, 0x83, 0x08, 0xb8, 0x00, 0xca, 0x92, 0x76, 0xed,
	0xf6, 0xeb, 0x36, 0xb0, 0xbc, 0xb7, 0xf1, 0x1c, 0x6b, 0x1c, 0x3a, 0xa0, 0x1a, 0xb2, 0x58, 0x26,
	0x5c, 0x6d, 0xbb, 0xac, 0x22, 0x80, 0x36, 0xd1, 0x0b, 0x83, 0x60, 0xfb, 0xe2, 0x7c, 0x2d, 0x81,
	0xd6, 0x1f, 0xbe, 0x04, 0xf8, 0xa1, 0x04, 0x40, 0x90, 0x19, 0x95, 0xab, 0xfe, 0xb5, 0xaa, 0xbb,
	0x37, 0xa8, 0xea, 0xe0, 0x2b, 0x18, 0xfe, 0xd1, 0x0c, 0x20, 0x8e, 0x73, 0xb5, 0x61, 0x1b, 0xd4,
	0x73, 0xdc, 0x46, 0xa2, 0x59, 0x6f, 0x5e, 0x25, 0xd4, 0x73, 0xe4, 0x38, 0x1f, 0xe3, 0x3d, 0x3a,
	0xfd, 0xbe, 0x38, 0x71, 0xa6, 0xce, 0x17, 0x75, 0xde, 0x5f, 0x2c, 0x96, 0x4e, 0xd5, 0x39, 0x53,
	0xe7, 0x9b, 0x3a, 0x1f, 0x2f, 0x17, 0x27, 0xde, 0xd4, 0xb2, 0x0e, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x29, 0x0b, 0xef, 0x6d, 0xe0, 0x06, 0x00, 0x00,
}
