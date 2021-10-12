// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

package hlmirbft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// BlockMetadata stores data used by the MirBFT OSNs when
// coordinating with each other, to be serialized into
// block meta dta field and used after failres and restarts.
type BlockMetadata struct {
	// Maintains a mapping between the cluster's OSNs
	// and their Mir-BFT IDs.
	ConsenterIds []uint64 `protobuf:"varint,1,rep,packed,name=consenter_ids,json=consenterIds,proto3" json:"consenter_ids,omitempty"`
	// Carries the Mir-BFT ID value that will be assigned
	// to the next OSN that will join this cluster.
	NextConsenterId uint64 `protobuf:"varint,2,opt,name=next_consenter_id,json=nextConsenterId,proto3" json:"next_consenter_id,omitempty"`
	// Index of MirBft entry for current block.
	MirbftIndex uint64 `protobuf:"varint,3,opt,name=mirbft_index,json=mirbftIndex,proto3" json:"mirbft_index,omitempty"`
	// Last Mir-BFT checkpoint sequence number reached by consensus
	LastCheckpointSeqNo  uint64   `protobuf:"varint,4,opt,name=last_checkpoint_seq_no,json=lastCheckpointSeqNo,proto3" json:"last_checkpoint_seq_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{0}
}

func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMetadata.Unmarshal(m, b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
}
func (m *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(m, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return xxx_messageInfo_BlockMetadata.Size(m)
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetConsenterIds() []uint64 {
	if m != nil {
		return m.ConsenterIds
	}
	return nil
}

func (m *BlockMetadata) GetNextConsenterId() uint64 {
	if m != nil {
		return m.NextConsenterId
	}
	return 0
}

func (m *BlockMetadata) GetMirbftIndex() uint64 {
	if m != nil {
		return m.MirbftIndex
	}
	return 0
}

func (m *BlockMetadata) GetLastCheckpointSeqNo() uint64 {
	if m != nil {
		return m.LastCheckpointSeqNo
	}
	return 0
}

// ClusterMetadata encapsulates metadata that is exchanged among cluster nodes
type ClusterMetadata struct {
	// Indicates active nodes in cluster that are reacheable by Raft leader
	ActiveNodes          []uint64 `protobuf:"varint,1,rep,packed,name=active_nodes,json=activeNodes,proto3" json:"active_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterMetadata) Reset()         { *m = ClusterMetadata{} }
func (m *ClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*ClusterMetadata) ProtoMessage()    {}
func (*ClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{1}
}

func (m *ClusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterMetadata.Unmarshal(m, b)
}
func (m *ClusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterMetadata.Marshal(b, m, deterministic)
}
func (m *ClusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMetadata.Merge(m, src)
}
func (m *ClusterMetadata) XXX_Size() int {
	return xxx_messageInfo_ClusterMetadata.Size(m)
}
func (m *ClusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMetadata proto.InternalMessageInfo

func (m *ClusterMetadata) GetActiveNodes() []uint64 {
	if m != nil {
		return m.ActiveNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockMetadata)(nil), "hlmirbft.BlockMetadata")
	proto.RegisterType((*ClusterMetadata)(nil), "hlmirbft.ClusterMetadata")
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor_56d9f74966f40d04) }

var fileDescriptor_56d9f74966f40d04 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x86, 0x83, 0x10, 0x63, 0x0a, 0x48, 0x9c, 0x89, 0xd9, 0x11, 0xf0, 0x42, 0x34, 0x74, 0x89,
	0xe8, 0x1f, 0x80, 0x13, 0x07, 0x39, 0xe0, 0xcd, 0x4b, 0xd3, 0xb5, 0xdf, 0xb6, 0x86, 0xae, 0xdf,
	0x68, 0x8b, 0xc1, 0x7f, 0xe6, 0xcf, 0x33, 0x5b, 0xd9, 0xf4, 0xfa, 0xbc, 0xcf, 0x93, 0x36, 0x1f,
	0xb9, 0x2d, 0xc1, 0x73, 0xc9, 0x3d, 0xa7, 0x95, 0x45, 0x8f, 0xd1, 0x4d, 0xa1, 0x4b, 0x65, 0xd3,
	0xcc, 0xcf, 0x7f, 0x7a, 0x64, 0xbc, 0xd6, 0x28, 0x0e, 0xef, 0x17, 0x23, 0x7a, 0x24, 0x63, 0x81,
	0xc6, 0x81, 0xf1, 0x60, 0x99, 0x92, 0x2e, 0xee, 0x4d, 0xfb, 0x8b, 0xc1, 0x7e, 0xd4, 0xc1, 0xad,
	0x74, 0xd1, 0x13, 0xb9, 0x33, 0x70, 0xf6, 0xec, 0xbf, 0x19, 0x5f, 0x4d, 0x7b, 0x8b, 0xc1, 0x7e,
	0x52, 0x0f, 0x9b, 0x3f, 0x39, 0x9a, 0x91, 0x51, 0x78, 0x8c, 0x29, 0x23, 0xe1, 0x1c, 0xf7, 0x1b,
	0x6d, 0x18, 0xd8, 0xb6, 0x46, 0xd1, 0x8a, 0x3c, 0x68, 0xee, 0x3c, 0x13, 0x05, 0x88, 0x43, 0x85,
	0xca, 0x78, 0xe6, 0xe0, 0xc8, 0x0c, 0xc6, 0x83, 0x46, 0xbe, 0xaf, 0xd7, 0x4d, 0x37, 0x7e, 0xc0,
	0x71, 0x87, 0xf3, 0x57, 0x32, 0xd9, 0xe8, 0x93, 0xf3, 0x60, 0xbb, 0xbf, 0xcf, 0xc8, 0x88, 0x0b,
	0xaf, 0xbe, 0x80, 0x19, 0x94, 0xd0, 0x7e, 0x7d, 0x18, 0xd8, 0xae, 0x46, 0x6b, 0x49, 0x9e, 0xd1,
	0xe6, 0x34, 0xd3, 0xdf, 0x2f, 0x95, 0xe6, 0x86, 0x66, 0x3c, 0xb5, 0x4a, 0x84, 0xbb, 0x38, 0x8a,
	0x56, 0x82, 0x05, 0x4b, 0xdb, 0xfb, 0x7c, 0xbe, 0xe5, 0xca, 0x17, 0xa7, 0x94, 0x0a, 0x2c, 0x93,
	0xb6, 0x49, 0x42, 0xb3, 0x0c, 0xcd, 0x32, 0xc7, 0xe4, 0x92, 0x25, 0x6d, 0x96, 0x5e, 0x37, 0xdb,
	0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x30, 0xe8, 0xc0, 0x16, 0x79, 0x01, 0x00, 0x00,
}
