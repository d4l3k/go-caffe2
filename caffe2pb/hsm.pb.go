// Code generated by protoc-gen-go. DO NOT EDIT.
// source: caffe2/proto/hsm.proto

/*
Package caffe2pb is a generated protocol buffer package.

It is generated from these files:
	caffe2/proto/hsm.proto
	caffe2/proto/predictor_consts.proto
	caffe2/proto/caffe2.proto
	caffe2/proto/prof_dag.proto
	caffe2/proto/metanet.proto
	caffe2/proto/caffe2_legacy.proto

It has these top-level messages:
	NodeProto
	TreeProto
	HierarchyProto
	PathProto
	PathNodeProto
	PredictorConsts
	ExternalDataProto
	TensorProto
	QTensorProto
	TensorProtos
	TensorShape
	TensorShapes
	Argument
	DeviceOption
	OperatorDef
	NetDef
	ExecutionStep
	PlanDef
	BlobProto
	DBReaderProto
	TwoNumberStatsProto
	BlobProfile
	ProfDAGProto
	ProfDAGProtos
	ModelInfo
	BlobsMap
	NetsMap
	PlansMap
	StringMap
	MetaNetDef
	CaffeDatum
*/
package caffe2pb

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

// Each node in the hierarchy contains links to either leaf nodes or more
// non-terminal nodes
type NodeProto struct {
	// Links to non-terminal children nodes
	Children []*NodeProto `protobuf:"bytes,1,rep,name=children" json:"children,omitempty"`
	// Links to terminal (leaf) nodes
	WordIds          []int32   `protobuf:"varint,2,rep,name=word_ids,json=wordIds" json:"word_ids,omitempty"`
	Offset           *int32    `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Name             *string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Scores           []float32 `protobuf:"fixed32,5,rep,name=scores" json:"scores,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NodeProto) Reset()                    { *m = NodeProto{} }
func (m *NodeProto) String() string            { return proto.CompactTextString(m) }
func (*NodeProto) ProtoMessage()               {}
func (*NodeProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NodeProto) GetChildren() []*NodeProto {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *NodeProto) GetWordIds() []int32 {
	if m != nil {
		return m.WordIds
	}
	return nil
}

func (m *NodeProto) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *NodeProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NodeProto) GetScores() []float32 {
	if m != nil {
		return m.Scores
	}
	return nil
}

// Protobuf format to accept hierarchy for hierarchical softmax operator.
// TreeProto points to the root node.
type TreeProto struct {
	RootNode         *NodeProto `protobuf:"bytes,1,opt,name=root_node,json=rootNode" json:"root_node,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *TreeProto) Reset()                    { *m = TreeProto{} }
func (m *TreeProto) String() string            { return proto.CompactTextString(m) }
func (*TreeProto) ProtoMessage()               {}
func (*TreeProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TreeProto) GetRootNode() *NodeProto {
	if m != nil {
		return m.RootNode
	}
	return nil
}

// Internal Protobuf format which represents the path in the tree hierarchy for
// each word in the vocabulary.
type HierarchyProto struct {
	Size             *int32       `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Paths            []*PathProto `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HierarchyProto) Reset()                    { *m = HierarchyProto{} }
func (m *HierarchyProto) String() string            { return proto.CompactTextString(m) }
func (*HierarchyProto) ProtoMessage()               {}
func (*HierarchyProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HierarchyProto) GetSize() int32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *HierarchyProto) GetPaths() []*PathProto {
	if m != nil {
		return m.Paths
	}
	return nil
}

// Each PathProto belongs to a word and is an array of nodes in the
// path from the root to the leaf (which is the word itself) in the tree.
type PathProto struct {
	WordId           *int32           `protobuf:"varint,1,opt,name=word_id,json=wordId" json:"word_id,omitempty"`
	PathNodes        []*PathNodeProto `protobuf:"bytes,2,rep,name=path_nodes,json=pathNodes" json:"path_nodes,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PathProto) Reset()                    { *m = PathProto{} }
func (m *PathProto) String() string            { return proto.CompactTextString(m) }
func (*PathProto) ProtoMessage()               {}
func (*PathProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PathProto) GetWordId() int32 {
	if m != nil && m.WordId != nil {
		return *m.WordId
	}
	return 0
}

func (m *PathProto) GetPathNodes() []*PathNodeProto {
	if m != nil {
		return m.PathNodes
	}
	return nil
}

// Represents a node in the path from the root node all the way down to the
// word (leaf).
type PathNodeProto struct {
	// Parameter matrix offset for this node
	Index *int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// Number of children
	Length *int32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	// Index of the next node in the path
	Target           *int32 `protobuf:"varint,3,opt,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PathNodeProto) Reset()                    { *m = PathNodeProto{} }
func (m *PathNodeProto) String() string            { return proto.CompactTextString(m) }
func (*PathNodeProto) ProtoMessage()               {}
func (*PathNodeProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PathNodeProto) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *PathNodeProto) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PathNodeProto) GetTarget() int32 {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeProto)(nil), "caffe2.NodeProto")
	proto.RegisterType((*TreeProto)(nil), "caffe2.TreeProto")
	proto.RegisterType((*HierarchyProto)(nil), "caffe2.HierarchyProto")
	proto.RegisterType((*PathProto)(nil), "caffe2.PathProto")
	proto.RegisterType((*PathNodeProto)(nil), "caffe2.PathNodeProto")
}

func init() { proto.RegisterFile("caffe2/proto/hsm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0xa6, 0x69, 0x9b, 0x43, 0x20, 0x61, 0x41, 0x31, 0x9b, 0x95, 0x85, 0x2c, 0xa4,
	0x52, 0xc5, 0xc6, 0x17, 0x80, 0x01, 0x54, 0x59, 0xb0, 0xb0, 0x54, 0x56, 0xec, 0xd4, 0x91, 0xda,
	0xb8, 0xb2, 0x2d, 0xf1, 0xe7, 0x9b, 0xf0, 0x6d, 0x91, 0x73, 0x26, 0x2d, 0x03, 0xdb, 0xfd, 0x9e,
	0x5f, 0x5e, 0xde, 0x1d, 0xcc, 0x6b, 0xd1, 0x34, 0x6a, 0xb9, 0xd8, 0x5b, 0xe3, 0xcd, 0x42, 0xbb,
	0x5d, 0xd5, 0x4f, 0x64, 0x82, 0x7a, 0xf1, 0x9d, 0x40, 0xfe, 0x6c, 0xa4, 0x5a, 0xf5, 0xea, 0x2d,
	0xcc, 0x6a, 0xdd, 0x6e, 0xa5, 0x55, 0x1d, 0x4d, 0x58, 0x5a, 0x9e, 0x2c, 0xcf, 0x2b, 0x34, 0x56,
	0x83, 0x89, 0x0f, 0x16, 0x72, 0x0d, 0xb3, 0x77, 0x63, 0xe5, 0xba, 0x95, 0x8e, 0x8e, 0x58, 0x5a,
	0x66, 0x7c, 0x1a, 0xf8, 0x51, 0x3a, 0x32, 0x87, 0x89, 0x69, 0x1a, 0xa7, 0x3c, 0x4d, 0x59, 0x52,
	0x66, 0x3c, 0x12, 0x21, 0x30, 0xee, 0xc4, 0x4e, 0xd1, 0x31, 0x4b, 0xca, 0x9c, 0xf7, 0x73, 0xf0,
	0xba, 0xda, 0x58, 0xe5, 0x68, 0xc6, 0xd2, 0x72, 0xc4, 0x23, 0x15, 0xf7, 0x90, 0xbf, 0x58, 0x15,
	0xab, 0x55, 0x90, 0x5b, 0x63, 0xfc, 0xba, 0x33, 0x52, 0xd1, 0x84, 0x25, 0xff, 0x74, 0x0b, 0x9e,
	0x80, 0xc5, 0x13, 0x9c, 0x3d, 0xb4, 0xca, 0x0a, 0x5b, 0xeb, 0x4f, 0x4c, 0x20, 0x30, 0x76, 0xed,
	0x17, 0x7e, 0x9c, 0xf1, 0x7e, 0x26, 0x37, 0x90, 0xed, 0x85, 0xd7, 0x58, 0xff, 0x28, 0x71, 0x25,
	0xbc, 0xc6, 0x44, 0x7c, 0x2f, 0xde, 0x20, 0x1f, 0x34, 0x72, 0x05, 0xd3, 0xb8, 0x77, 0x0c, 0x9b,
	0xe0, 0xda, 0xe4, 0x0e, 0x20, 0xd8, 0xfb, 0x92, 0xbf, 0x99, 0x97, 0xc7, 0x99, 0x87, 0xa6, 0xf9,
	0x3e, 0xa2, 0x2b, 0x5e, 0xe1, 0xf4, 0xcf, 0x1b, 0xb9, 0x80, 0xac, 0xed, 0xa4, 0xfa, 0x88, 0xe9,
	0x08, 0xe1, 0x4c, 0x5b, 0xd5, 0x6d, 0xbc, 0xa6, 0x23, 0xfc, 0x29, 0x52, 0xd0, 0xbd, 0xb0, 0x9b,
	0xc3, 0xa9, 0x91, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xb3, 0x23, 0xcc, 0xfb, 0x01, 0x00,
	0x00,
}
