// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/tx/sig_desc.proto

package tx

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/signing"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

func init() { proto.RegisterFile("cosmos/tx/sig_desc.proto", fileDescriptor_d560eb6bd3d62ef7) }

var fileDescriptor_d560eb6bd3d62ef7 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0xa9, 0xd0, 0x2f, 0xce, 0x4c, 0x8f, 0x4f, 0x49, 0x2d, 0x4e, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xc8, 0xe8, 0x95, 0x54, 0x48, 0xc9, 0xa3, 0x28, 0xca,
	0xcb, 0xcc, 0x4b, 0x87, 0xd1, 0x10, 0xb5, 0x4e, 0xf6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xa5, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f,
	0x35, 0x05, 0x42, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x82, 0x8c, 0x4d, 0x62,
	0x03, 0x9b, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x86, 0x8a, 0xa4, 0x8a, 0x8f, 0x00, 0x00,
	0x00,
}