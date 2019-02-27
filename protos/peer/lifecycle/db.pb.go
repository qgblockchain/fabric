// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/lifecycle/db.proto

package lifecycle // import "github.com/hyperledger/fabric/protos/peer/lifecycle"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StateMetadata describes the keys in a namespace.  It is necessary because
// in collections, range scans are not possible during transactions which
// write.  Therefore we must track the keys in our namespace ourselves.
type StateMetadata struct {
	Datatype             string   `protobuf:"bytes,1,opt,name=datatype,proto3" json:"datatype,omitempty"`
	Fields               []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateMetadata) Reset()         { *m = StateMetadata{} }
func (m *StateMetadata) String() string { return proto.CompactTextString(m) }
func (*StateMetadata) ProtoMessage()    {}
func (*StateMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_8951d509ec26cc08, []int{0}
}
func (m *StateMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateMetadata.Unmarshal(m, b)
}
func (m *StateMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateMetadata.Marshal(b, m, deterministic)
}
func (dst *StateMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateMetadata.Merge(dst, src)
}
func (m *StateMetadata) XXX_Size() int {
	return xxx_messageInfo_StateMetadata.Size(m)
}
func (m *StateMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_StateMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_StateMetadata proto.InternalMessageInfo

func (m *StateMetadata) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *StateMetadata) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

// StateData encodes a particular field of a datatype
type StateData struct {
	// Types that are valid to be assigned to Type:
	//	*StateData_Int64
	//	*StateData_Bytes
	Type                 isStateData_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StateData) Reset()         { *m = StateData{} }
func (m *StateData) String() string { return proto.CompactTextString(m) }
func (*StateData) ProtoMessage()    {}
func (*StateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_8951d509ec26cc08, []int{1}
}
func (m *StateData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateData.Unmarshal(m, b)
}
func (m *StateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateData.Marshal(b, m, deterministic)
}
func (dst *StateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateData.Merge(dst, src)
}
func (m *StateData) XXX_Size() int {
	return xxx_messageInfo_StateData.Size(m)
}
func (m *StateData) XXX_DiscardUnknown() {
	xxx_messageInfo_StateData.DiscardUnknown(m)
}

var xxx_messageInfo_StateData proto.InternalMessageInfo

type isStateData_Type interface {
	isStateData_Type()
}

type StateData_Int64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=Int64,proto3,oneof"`
}

type StateData_Bytes struct {
	Bytes []byte `protobuf:"bytes,2,opt,name=Bytes,proto3,oneof"`
}

func (*StateData_Int64) isStateData_Type() {}

func (*StateData_Bytes) isStateData_Type() {}

func (m *StateData) GetType() isStateData_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *StateData) GetInt64() int64 {
	if x, ok := m.GetType().(*StateData_Int64); ok {
		return x.Int64
	}
	return 0
}

func (m *StateData) GetBytes() []byte {
	if x, ok := m.GetType().(*StateData_Bytes); ok {
		return x.Bytes
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StateData) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StateData_OneofMarshaler, _StateData_OneofUnmarshaler, _StateData_OneofSizer, []interface{}{
		(*StateData_Int64)(nil),
		(*StateData_Bytes)(nil),
	}
}

func _StateData_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StateData)
	// Type
	switch x := m.Type.(type) {
	case *StateData_Int64:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64))
	case *StateData_Bytes:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Bytes)
	case nil:
	default:
		return fmt.Errorf("StateData.Type has unexpected type %T", x)
	}
	return nil
}

func _StateData_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StateData)
	switch tag {
	case 1: // Type.Int64
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &StateData_Int64{int64(x)}
		return true, err
	case 2: // Type.Bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Type = &StateData_Bytes{x}
		return true, err
	default:
		return false, nil
	}
}

func _StateData_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StateData)
	// Type
	switch x := m.Type.(type) {
	case *StateData_Int64:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64))
	case *StateData_Bytes:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Bytes)))
		n += len(x.Bytes)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*StateMetadata)(nil), "lifecycle.StateMetadata")
	proto.RegisterType((*StateData)(nil), "lifecycle.StateData")
}

func init() { proto.RegisterFile("peer/lifecycle/db.proto", fileDescriptor_db_8951d509ec26cc08) }

var fileDescriptor_db_8951d509ec26cc08 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x40, 0x09, 0x85, 0x08, 0x9f, 0x60, 0xc9, 0x50, 0x2a, 0xa6, 0xaa, 0x53, 0x06, 0x64, 0x0f,
	0x45, 0xfc, 0x80, 0x94, 0x01, 0x06, 0x96, 0xc0, 0xc4, 0xe6, 0x8f, 0x4b, 0x6a, 0xc9, 0x60, 0xcb,
	0x3d, 0x06, 0xff, 0x7b, 0x64, 0x1b, 0x45, 0x30, 0x59, 0xef, 0x49, 0xef, 0x74, 0x3e, 0xb8, 0x0d,
	0x88, 0x51, 0x38, 0x3b, 0xa1, 0x4e, 0xda, 0xa1, 0x30, 0x8a, 0x87, 0xe8, 0xc9, 0x77, 0x6c, 0x71,
	0xbb, 0x03, 0xdc, 0xbc, 0x91, 0x24, 0x7c, 0x45, 0x92, 0x46, 0x92, 0xec, 0xee, 0xe0, 0x2a, 0xbf,
	0x94, 0x02, 0x6e, 0x9a, 0x6d, 0xd3, 0xb3, 0x71, 0xe1, 0x6e, 0x0d, 0xed, 0x64, 0xd1, 0x99, 0xd3,
	0xe6, 0x7c, 0xbb, 0xea, 0xd9, 0xf8, 0x4b, 0xbb, 0x03, 0xb0, 0x32, 0xe4, 0x29, 0x0f, 0x58, 0xc3,
	0xe5, 0xcb, 0x17, 0x3d, 0x3e, 0x94, 0x7a, 0xf5, 0x7c, 0x36, 0x56, 0xcc, 0x7e, 0x48, 0x84, 0xb9,
	0x6d, 0xfa, 0xeb, 0xec, 0x0b, 0x0e, 0x2d, 0x5c, 0xbc, 0xa7, 0x80, 0x83, 0x86, 0x7b, 0x1f, 0x67,
	0x7e, 0x4c, 0x01, 0xa3, 0x43, 0x33, 0x63, 0xe4, 0x93, 0x54, 0xd1, 0xea, 0xba, 0xf4, 0x89, 0xe7,
	0xdf, 0xf0, 0x65, 0xf3, 0x8f, 0xfd, 0x6c, 0xe9, 0xf8, 0xad, 0xb8, 0xf6, 0x9f, 0xe2, 0x4f, 0x24,
	0x6a, 0x24, 0x6a, 0x24, 0xfe, 0x9f, 0x40, 0xb5, 0x45, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0xb7, 0x06, 0x5b, 0x1b, 0x01, 0x00, 0x00,
}
