// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agoric/swingset/storage.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Storage is the data storage.
type Storage struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value" yaml:"value"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e196a9f45e310a8, []int{0}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return m.Size()
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Keys are the storage node subkeys.
type Keys struct {
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys" yaml:"keys"`
}

func (m *Keys) Reset()         { *m = Keys{} }
func (m *Keys) String() string { return proto.CompactTextString(m) }
func (*Keys) ProtoMessage()    {}
func (*Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e196a9f45e310a8, []int{1}
}
func (m *Keys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keys.Merge(m, src)
}
func (m *Keys) XXX_Size() int {
	return m.Size()
}
func (m *Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_Keys proto.InternalMessageInfo

func (m *Keys) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

// Egress is the format for a swingset egress.
type Egress struct {
	Nickname   string                                        `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname" yaml:"nickname"`
	Peer       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=peer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"peer" yaml:"peer"`
	PowerFlags []string                                      `protobuf:"bytes,3,rep,name=power_flags,json=powerFlags,proto3" json:"powerFlags" yaml:"powerFlags"`
}

func (m *Egress) Reset()         { *m = Egress{} }
func (m *Egress) String() string { return proto.CompactTextString(m) }
func (*Egress) ProtoMessage()    {}
func (*Egress) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e196a9f45e310a8, []int{2}
}
func (m *Egress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Egress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Egress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Egress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Egress.Merge(m, src)
}
func (m *Egress) XXX_Size() int {
	return m.Size()
}
func (m *Egress) XXX_DiscardUnknown() {
	xxx_messageInfo_Egress.DiscardUnknown(m)
}

var xxx_messageInfo_Egress proto.InternalMessageInfo

func (m *Egress) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Egress) GetPeer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *Egress) GetPowerFlags() []string {
	if m != nil {
		return m.PowerFlags
	}
	return nil
}

func init() {
	proto.RegisterType((*Storage)(nil), "agoric.swingset.Storage")
	proto.RegisterType((*Keys)(nil), "agoric.swingset.Keys")
	proto.RegisterType((*Egress)(nil), "agoric.swingset.Egress")
}

func init() { proto.RegisterFile("agoric/swingset/storage.proto", fileDescriptor_4e196a9f45e310a8) }

var fileDescriptor_4e196a9f45e310a8 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x41, 0xeb, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0x57, 0xa7, 0xcb, 0x06, 0xc3, 0x22, 0x58, 0x05, 0x9b, 0x91, 0xd3, 0x40, 0xd6,
	0x22, 0x9e, 0xdc, 0x2e, 0xae, 0xa0, 0x08, 0x5e, 0xa4, 0xe2, 0x45, 0x04, 0xe9, 0xba, 0x18, 0xcb,
	0xda, 0xa6, 0x34, 0x9d, 0xb3, 0xdf, 0xc2, 0x8f, 0xe0, 0xc7, 0xf1, 0xb8, 0xa3, 0xa7, 0x20, 0xdd,
	0x45, 0x7a, 0xec, 0x51, 0xf8, 0xc3, 0x9f, 0x24, 0x6b, 0xb7, 0x53, 0xfb, 0xfc, 0xde, 0xbc, 0x79,
	0xde, 0x87, 0x37, 0xf0, 0x69, 0x48, 0x59, 0x11, 0x47, 0x1e, 0x3f, 0xc4, 0x19, 0xe5, 0xa4, 0xf4,
	0x78, 0xc9, 0x8a, 0x90, 0x12, 0x37, 0x2f, 0x58, 0xc9, 0xac, 0xa9, 0x2e, 0xbb, 0x5d, 0xf9, 0xc9,
	0x43, 0xca, 0x28, 0x53, 0x35, 0x4f, 0xfe, 0xe9, 0x63, 0xf8, 0x15, 0xbc, 0xf7, 0x41, 0xf7, 0x59,
	0x1e, 0xbc, 0xfb, 0x3d, 0x4c, 0xf6, 0xc4, 0x06, 0x33, 0x30, 0x1f, 0xf9, 0x8f, 0x1b, 0x81, 0x34,
	0x68, 0x05, 0x9a, 0x54, 0x61, 0x9a, 0x2c, 0xb1, 0x92, 0x38, 0xd0, 0x78, 0x69, 0xfe, 0xfb, 0x85,
	0x0c, 0xfc, 0x12, 0x9a, 0xef, 0x48, 0xc5, 0xad, 0x67, 0xd0, 0xdc, 0x91, 0x8a, 0xdb, 0x60, 0x36,
	0x98, 0x8f, 0xfc, 0x47, 0x8d, 0x40, 0x4a, 0xb7, 0x02, 0x8d, 0x75, 0xb3, 0x54, 0x38, 0x50, 0xf0,
	0xdc, 0x7a, 0x03, 0xe0, 0xf0, 0x35, 0x2d, 0x08, 0xe7, 0xd6, 0x0a, 0xde, 0xcf, 0xe2, 0x68, 0x97,
	0x85, 0x69, 0xe7, 0x8f, 0x1a, 0x81, 0x7a, 0xd6, 0x0a, 0x34, 0xd5, 0xb7, 0x74, 0x04, 0x07, 0x7d,
	0xd1, 0xfa, 0x0c, 0xcd, 0x9c, 0x90, 0xc2, 0xbe, 0x33, 0x03, 0xf3, 0x89, 0xff, 0x56, 0x5a, 0x4b,
	0x7d, 0xb1, 0x96, 0x0a, 0xff, 0x17, 0x68, 0x41, 0xe3, 0xf2, 0xdb, 0x7e, 0xe3, 0x46, 0x2c, 0xf5,
	0x22, 0xc6, 0x53, 0xc6, 0xcf, 0x9f, 0x05, 0xdf, 0xee, 0xbc, 0xb2, 0xca, 0x09, 0x77, 0xd7, 0x51,
	0xb4, 0xde, 0x6e, 0xe5, 0x50, 0x81, 0xba, 0xc5, 0x0a, 0xe0, 0x38, 0x67, 0x07, 0x52, 0x7c, 0xf9,
	0x9a, 0x84, 0x94, 0xdb, 0x03, 0x95, 0xef, 0x79, 0x2d, 0x10, 0x7c, 0x2f, 0xf1, 0x1b, 0x49, 0x1b,
	0x81, 0x60, 0xde, 0xab, 0x56, 0xa0, 0x07, 0x67, 0xe3, 0x9e, 0xe1, 0xe0, 0xea, 0x80, 0xce, 0xef,
	0x7f, 0xfc, 0x5d, 0x3b, 0xe0, 0x58, 0x3b, 0xe0, 0x6f, 0xed, 0x80, 0x9f, 0x27, 0xc7, 0x38, 0x9e,
	0x1c, 0xe3, 0xcf, 0xc9, 0x31, 0x3e, 0xad, 0xae, 0x06, 0x5d, 0xeb, 0x3d, 0xeb, 0x7d, 0xaa, 0x41,
	0x29, 0x4b, 0xc2, 0x8c, 0x76, 0x09, 0x7e, 0x5c, 0x9e, 0x80, 0x4a, 0xb0, 0x19, 0xaa, 0xd5, 0xbe,
	0xb8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x60, 0x8b, 0xfa, 0x22, 0x02, 0x00, 0x00,
}

func (m *Storage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Storage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Storage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Keys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintStorage(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Egress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Egress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Egress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PowerFlags) > 0 {
		for iNdEx := len(m.PowerFlags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PowerFlags[iNdEx])
			copy(dAtA[i:], m.PowerFlags[iNdEx])
			i = encodeVarintStorage(dAtA, i, uint64(len(m.PowerFlags[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Peer) > 0 {
		i -= len(m.Peer)
		copy(dAtA[i:], m.Peer)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Peer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Storage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *Keys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			l = len(s)
			n += 1 + l + sovStorage(uint64(l))
		}
	}
	return n
}

func (m *Egress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Peer)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if len(m.PowerFlags) > 0 {
		for _, s := range m.PowerFlags {
			l = len(s)
			n += 1 + l + sovStorage(uint64(l))
		}
	}
	return n
}

func sovStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Storage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Storage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Storage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
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
func (m *Keys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Keys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
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
func (m *Egress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Egress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Egress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peer = append(m.Peer[:0], dAtA[iNdEx:postIndex]...)
			if m.Peer == nil {
				m.Peer = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PowerFlags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PowerFlags = append(m.PowerFlags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
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
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
				return 0, ErrInvalidLengthStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStorage = fmt.Errorf("proto: unexpected end of group")
)
