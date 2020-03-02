package commander

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	serial "github.com/gitamenet/v2ray-core/common/serial"
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

// Config is the settings for Commander.
type Config struct {
	// Tag of the outbound handler that handles grpc connections.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Services that supported by this server. All services must implement Service interface.
	Service              []*serial.TypedMessage `protobuf:"bytes,2,rep,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_4af9cfd3f0e2019e, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Config) GetService() []*serial.TypedMessage {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.app.commander.Config")
}

func init() {
	proto.RegisterFile("github.com/gitamenet/v2ray-core/app/commander/config.proto", fileDescriptor_4af9cfd3f0e2019e)
}

var fileDescriptor_4af9cfd3f0e2019e = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x4f,
	0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x49, 0x2d, 0xd2, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0x29, 0x2d, 0x4a, 0xd5, 0x4b, 0x2c, 0x28, 0xd0,
	0x83, 0x2b, 0x93, 0x32, 0x40, 0x33, 0x04, 0x24, 0x93, 0x9f, 0xa7, 0x5f, 0x9c, 0x5a, 0x94, 0x99,
	0x98, 0xa3, 0x5f, 0x52, 0x59, 0x90, 0x9a, 0x12, 0x9f, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x0a,
	0x31, 0x4b, 0x29, 0x86, 0x8b, 0xcd, 0x19, 0x6c, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x49, 0x62, 0xba,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0xe4, 0xc0, 0xc5, 0x5e, 0x9c, 0x5a, 0x54,
	0x96, 0x99, 0x9c, 0x2a, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa6, 0x87, 0x64, 0x33, 0xc4,
	0x6c, 0x3d, 0x88, 0xd9, 0x7a, 0x21, 0x20, 0xb3, 0x7d, 0x21, 0x46, 0x07, 0xc1, 0xb4, 0x39, 0xb9,
	0x71, 0xc9, 0x24, 0xe7, 0xe7, 0xea, 0xe1, 0x72, 0x6f, 0x00, 0x63, 0x14, 0x27, 0x9c, 0xb3, 0x8a,
	0x49, 0x22, 0xcc, 0x28, 0x28, 0xb1, 0x52, 0xcf, 0x19, 0xa4, 0xce, 0xb1, 0xa0, 0x40, 0xcf, 0x19,
	0x26, 0x95, 0xc4, 0x06, 0x76, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x29, 0x02, 0xa6, 0x19,
	0x25, 0x01, 0x00, 0x00,
}
