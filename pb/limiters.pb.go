// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/ratelimitd/limiters.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LimiterList struct {
	List                 []*Limiter `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LimiterList) Reset()         { *m = LimiterList{} }
func (m *LimiterList) String() string { return proto.CompactTextString(m) }
func (*LimiterList) ProtoMessage()    {}
func (*LimiterList) Descriptor() ([]byte, []int) {
	return fileDescriptor_08617d6a6bfd79d5, []int{0}
}
func (m *LimiterList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimiterList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimiterList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimiterList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimiterList.Merge(m, src)
}
func (m *LimiterList) XXX_Size() int {
	return m.Size()
}
func (m *LimiterList) XXX_DiscardUnknown() {
	xxx_messageInfo_LimiterList.DiscardUnknown(m)
}

var xxx_messageInfo_LimiterList proto.InternalMessageInfo

func (m *LimiterList) GetList() []*Limiter {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*LimiterList)(nil), "ratelimitd.LimiterList")
}

func init() { proto.RegisterFile("pb/ratelimitd/limiters.proto", fileDescriptor_08617d6a6bfd79d5) }

var fileDescriptor_08617d6a6bfd79d5 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x48, 0xd2, 0x2f,
	0x4a, 0x2c, 0x49, 0xcd, 0xc9, 0xcc, 0xcd, 0x2c, 0x49, 0xd1, 0x07, 0x53, 0xa9, 0x45, 0xc5, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x29, 0x29, 0x69, 0xac, 0x2a, 0x21, 0x0a, 0x95,
	0xcc, 0xb8, 0xb8, 0x7d, 0x20, 0x02, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0xea, 0x5c, 0x2c, 0x39, 0x99,
	0xc5, 0x25, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x08, 0x7d, 0x7a, 0x50, 0x65,
	0x41, 0x60, 0x05, 0x4e, 0x86, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0xf2, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x49, 0xa5, 0x49, 0xa5, 0x49, 0x89, 0xc8, 0xb6, 0x16, 0x24, 0x05, 0x30, 0x24,
	0xb1, 0x81, 0xed, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x57, 0x06, 0xd8, 0x10, 0xbc, 0x00,
	0x00, 0x00,
}

func (m *LimiterList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimiterList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimiterList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLimiters(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimiters(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimiters(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimiterList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovLimiters(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLimiters(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimiters(x uint64) (n int) {
	return sovLimiters(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimiterList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimiters
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
			return fmt.Errorf("proto: LimiterList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimiterList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimiters
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
				return ErrInvalidLengthLimiters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimiters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &Limiter{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimiters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimiters
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLimiters(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimiters
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
					return 0, ErrIntOverflowLimiters
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
					return 0, ErrIntOverflowLimiters
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
				return 0, ErrInvalidLengthLimiters
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimiters
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimiters
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimiters        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimiters          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimiters = fmt.Errorf("proto: unexpected end of group")
)
