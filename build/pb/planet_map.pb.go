// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/planet_map.proto

package pb

import (
	fmt "fmt"
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

type PlanetMap struct {
	ID             uint32     `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CellGrid       string     `protobuf:"bytes,2,opt,name=CellGrid,proto3" json:"CellGrid,omitempty"`
	StartPositionX int32      `protobuf:"varint,3,opt,name=StartPositionX,proto3" json:"StartPositionX,omitempty"`
	StartPositionY int32      `protobuf:"varint,4,opt,name=StartPositionY,proto3" json:"StartPositionY,omitempty"`
	Enemies        []*Enemy   `protobuf:"bytes,5,rep,name=Enemies,proto3" json:"Enemies,omitempty"`
	Tresures       []*Tresure `protobuf:"bytes,6,rep,name=Tresures,proto3" json:"Tresures,omitempty"`
}

func (m *PlanetMap) Reset()         { *m = PlanetMap{} }
func (m *PlanetMap) String() string { return proto.CompactTextString(m) }
func (*PlanetMap) ProtoMessage()    {}
func (*PlanetMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76a6df7491a6a50, []int{0}
}
func (m *PlanetMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlanetMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlanetMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlanetMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanetMap.Merge(m, src)
}
func (m *PlanetMap) XXX_Size() int {
	return m.Size()
}
func (m *PlanetMap) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanetMap.DiscardUnknown(m)
}

var xxx_messageInfo_PlanetMap proto.InternalMessageInfo

func (m *PlanetMap) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PlanetMap) GetCellGrid() string {
	if m != nil {
		return m.CellGrid
	}
	return ""
}

func (m *PlanetMap) GetStartPositionX() int32 {
	if m != nil {
		return m.StartPositionX
	}
	return 0
}

func (m *PlanetMap) GetStartPositionY() int32 {
	if m != nil {
		return m.StartPositionY
	}
	return 0
}

func (m *PlanetMap) GetEnemies() []*Enemy {
	if m != nil {
		return m.Enemies
	}
	return nil
}

func (m *PlanetMap) GetTresures() []*Tresure {
	if m != nil {
		return m.Tresures
	}
	return nil
}

// GetMapByID
type GetPlanetMapByIDRequest struct {
	PlanetMapID uint32 `protobuf:"varint,1,opt,name=PlanetMapID,proto3" json:"PlanetMapID,omitempty"`
}

func (m *GetPlanetMapByIDRequest) Reset()         { *m = GetPlanetMapByIDRequest{} }
func (m *GetPlanetMapByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlanetMapByIDRequest) ProtoMessage()    {}
func (*GetPlanetMapByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76a6df7491a6a50, []int{1}
}
func (m *GetPlanetMapByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlanetMapByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlanetMapByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlanetMapByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlanetMapByIDRequest.Merge(m, src)
}
func (m *GetPlanetMapByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPlanetMapByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlanetMapByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlanetMapByIDRequest proto.InternalMessageInfo

func (m *GetPlanetMapByIDRequest) GetPlanetMapID() uint32 {
	if m != nil {
		return m.PlanetMapID
	}
	return 0
}

type GetPlanetMapByIDResponse struct {
	PlanetMap *PlanetMap `protobuf:"bytes,1,opt,name=PlanetMap,proto3" json:"PlanetMap,omitempty"`
}

func (m *GetPlanetMapByIDResponse) Reset()         { *m = GetPlanetMapByIDResponse{} }
func (m *GetPlanetMapByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlanetMapByIDResponse) ProtoMessage()    {}
func (*GetPlanetMapByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76a6df7491a6a50, []int{2}
}
func (m *GetPlanetMapByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPlanetMapByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPlanetMapByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPlanetMapByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlanetMapByIDResponse.Merge(m, src)
}
func (m *GetPlanetMapByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPlanetMapByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlanetMapByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlanetMapByIDResponse proto.InternalMessageInfo

func (m *GetPlanetMapByIDResponse) GetPlanetMap() *PlanetMap {
	if m != nil {
		return m.PlanetMap
	}
	return nil
}

func init() {
	proto.RegisterType((*PlanetMap)(nil), "planet_map.PlanetMap")
	proto.RegisterType((*GetPlanetMapByIDRequest)(nil), "planet_map.GetPlanetMapByIDRequest")
	proto.RegisterType((*GetPlanetMapByIDResponse)(nil), "planet_map.GetPlanetMapByIDResponse")
}

func init() { proto.RegisterFile("proto/planet_map.proto", fileDescriptor_c76a6df7491a6a50) }

var fileDescriptor_c76a6df7491a6a50 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3d, 0x4b, 0xc3, 0x40,
	0x1c, 0xc6, 0x7b, 0x7d, 0xb3, 0xbd, 0x6a, 0xd1, 0x13, 0xf5, 0xe8, 0x70, 0x84, 0x0e, 0x25, 0x83,
	0x44, 0x68, 0xc7, 0x6e, 0x35, 0x52, 0x32, 0x88, 0xe5, 0x74, 0xb0, 0x2e, 0x92, 0xe2, 0x7f, 0x08,
	0xe4, 0xe5, 0xcc, 0x5d, 0x87, 0x7c, 0x0b, 0x3f, 0x96, 0x63, 0x47, 0x17, 0x41, 0x92, 0x2f, 0x22,
	0xe6, 0xec, 0x35, 0x58, 0xc7, 0xfc, 0x9e, 0xdf, 0x03, 0x79, 0xee, 0x8f, 0xcf, 0x45, 0x9a, 0xa8,
	0xe4, 0x4a, 0x84, 0x7e, 0x0c, 0xea, 0x39, 0xf2, 0x85, 0x53, 0x02, 0x82, 0x77, 0x64, 0x70, 0xa2,
	0x1d, 0x88, 0x21, 0xca, 0x74, 0x3c, 0x38, 0xd5, 0x48, 0xa5, 0x20, 0xd7, 0x29, 0x68, 0x38, 0xfc,
	0x44, 0xb8, 0xbb, 0x28, 0x6b, 0xb7, 0xbe, 0x20, 0x7d, 0x5c, 0xf7, 0x5c, 0x8a, 0x2c, 0x64, 0x1f,
	0xf1, 0xba, 0xe7, 0x92, 0x01, 0xee, 0x5c, 0x43, 0x18, 0xce, 0xd3, 0xe0, 0x85, 0xd6, 0x2d, 0x64,
	0x77, 0xb9, 0xf9, 0x26, 0x23, 0xdc, 0xbf, 0x57, 0x7e, 0xaa, 0x16, 0x89, 0x0c, 0x54, 0x90, 0xc4,
	0x8f, 0xb4, 0x61, 0x21, 0xbb, 0xc5, 0xff, 0xd0, 0x3d, 0x6f, 0x49, 0x9b, 0xff, 0x78, 0x4b, 0x32,
	0xc2, 0x07, 0x37, 0x31, 0x44, 0x01, 0x48, 0xda, 0xb2, 0x1a, 0x76, 0x6f, 0x7c, 0xe8, 0xe8, 0xbf,
	0xff, 0xa1, 0x19, 0xdf, 0x86, 0xe4, 0x12, 0x77, 0x1e, 0xf4, 0x04, 0x49, 0xdb, 0xa5, 0x78, 0xec,
	0x6c, 0x37, 0xfd, 0x06, 0xdc, 0x18, 0xc3, 0x29, 0xbe, 0x98, 0x83, 0x32, 0x0b, 0x67, 0x99, 0xe7,
	0x72, 0x78, 0x5d, 0x83, 0x54, 0xc4, 0xc2, 0x3d, 0xc3, 0xcd, 0xea, 0x2a, 0x1a, 0xde, 0x61, 0xba,
	0x5f, 0x96, 0x22, 0x89, 0x25, 0x90, 0x49, 0xe5, 0xdd, 0xca, 0x6e, 0x6f, 0x7c, 0xe6, 0x54, 0x4e,
	0x62, 0x42, 0xbe, 0xf3, 0x66, 0xec, 0x3d, 0x67, 0x68, 0x93, 0x33, 0xf4, 0x95, 0x33, 0xf4, 0x56,
	0xb0, 0xda, 0xa6, 0x60, 0xb5, 0x8f, 0x82, 0xd5, 0x9e, 0x9a, 0xce, 0x54, 0xac, 0x56, 0xed, 0xf2,
	0x28, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x30, 0xcd, 0xad, 0xe2, 0x01, 0x00, 0x00,
}

func (m *PlanetMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlanetMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlanetMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tresures) > 0 {
		for iNdEx := len(m.Tresures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tresures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlanetMap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Enemies) > 0 {
		for iNdEx := len(m.Enemies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Enemies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlanetMap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.StartPositionY != 0 {
		i = encodeVarintPlanetMap(dAtA, i, uint64(m.StartPositionY))
		i--
		dAtA[i] = 0x20
	}
	if m.StartPositionX != 0 {
		i = encodeVarintPlanetMap(dAtA, i, uint64(m.StartPositionX))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CellGrid) > 0 {
		i -= len(m.CellGrid)
		copy(dAtA[i:], m.CellGrid)
		i = encodeVarintPlanetMap(dAtA, i, uint64(len(m.CellGrid)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintPlanetMap(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlanetMapByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlanetMapByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlanetMapByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanetMapID != 0 {
		i = encodeVarintPlanetMap(dAtA, i, uint64(m.PlanetMapID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetPlanetMapByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPlanetMapByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPlanetMapByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanetMap != nil {
		{
			size, err := m.PlanetMap.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlanetMap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlanetMap(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlanetMap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlanetMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPlanetMap(uint64(m.ID))
	}
	l = len(m.CellGrid)
	if l > 0 {
		n += 1 + l + sovPlanetMap(uint64(l))
	}
	if m.StartPositionX != 0 {
		n += 1 + sovPlanetMap(uint64(m.StartPositionX))
	}
	if m.StartPositionY != 0 {
		n += 1 + sovPlanetMap(uint64(m.StartPositionY))
	}
	if len(m.Enemies) > 0 {
		for _, e := range m.Enemies {
			l = e.Size()
			n += 1 + l + sovPlanetMap(uint64(l))
		}
	}
	if len(m.Tresures) > 0 {
		for _, e := range m.Tresures {
			l = e.Size()
			n += 1 + l + sovPlanetMap(uint64(l))
		}
	}
	return n
}

func (m *GetPlanetMapByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanetMapID != 0 {
		n += 1 + sovPlanetMap(uint64(m.PlanetMapID))
	}
	return n
}

func (m *GetPlanetMapByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanetMap != nil {
		l = m.PlanetMap.Size()
		n += 1 + l + sovPlanetMap(uint64(l))
	}
	return n
}

func sovPlanetMap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlanetMap(x uint64) (n int) {
	return sovPlanetMap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlanetMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlanetMap
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
			return fmt.Errorf("proto: PlanetMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlanetMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellGrid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
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
				return ErrInvalidLengthPlanetMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlanetMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CellGrid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPositionX", wireType)
			}
			m.StartPositionX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartPositionX |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPositionY", wireType)
			}
			m.StartPositionY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartPositionY |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enemies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
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
				return ErrInvalidLengthPlanetMap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlanetMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Enemies = append(m.Enemies, &Enemy{})
			if err := m.Enemies[len(m.Enemies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tresures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
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
				return ErrInvalidLengthPlanetMap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlanetMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tresures = append(m.Tresures, &Tresure{})
			if err := m.Tresures[len(m.Tresures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlanetMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlanetMap
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
func (m *GetPlanetMapByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlanetMap
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
			return fmt.Errorf("proto: GetPlanetMapByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlanetMapByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanetMapID", wireType)
			}
			m.PlanetMapID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanetMapID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlanetMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlanetMap
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
func (m *GetPlanetMapByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlanetMap
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
			return fmt.Errorf("proto: GetPlanetMapByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPlanetMapByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanetMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlanetMap
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
				return ErrInvalidLengthPlanetMap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlanetMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlanetMap == nil {
				m.PlanetMap = &PlanetMap{}
			}
			if err := m.PlanetMap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlanetMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlanetMap
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
func skipPlanetMap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlanetMap
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
					return 0, ErrIntOverflowPlanetMap
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
					return 0, ErrIntOverflowPlanetMap
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
				return 0, ErrInvalidLengthPlanetMap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlanetMap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlanetMap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlanetMap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlanetMap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlanetMap = fmt.Errorf("proto: unexpected end of group")
)
