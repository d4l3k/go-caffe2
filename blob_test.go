package caffe2

import (
	"reflect"
	"testing"

	"github.com/d4l3k/go-caffe2/caffe2pb"
	"github.com/gogo/protobuf/proto"
)

func TestBlob(t *testing.T) {
	w := NewWorkspace()
	b := w.CreateBlob("foo")
	b.Reset()
	t.Log("type name", b.TypeName())
}

func TestBlobTensorSerialization(t *testing.T) {
	w := NewWorkspace()
	b := w.CreateBlob("foo")
	dataType := caffe2pb.TensorProto_FLOAT
	in := &caffe2pb.BlobProto{
		Type: proto.String("Tensor"),
		Tensor: &caffe2pb.TensorProto{
			DataType:  &dataType,
			Dims:      []int64{3},
			FloatData: []float32{1, 2, 3},
		},
	}
	if err := b.FromProto(in); err != nil {
		t.Fatal(err)
	}

	out := b.Serialize()
	b.Deserialize(out)
}

func TestBlobTensorSerializationProto(t *testing.T) {
	w := NewWorkspace()
	b := w.CreateBlob("foo")
	dataType := caffe2pb.TensorProto_FLOAT
	in := &caffe2pb.BlobProto{
		Type: proto.String("Tensor"),
		Tensor: &caffe2pb.TensorProto{
			DataType:  &dataType,
			Dims:      []int64{3},
			FloatData: []float32{1, 2, 3},
		},
	}
	if err := b.FromProto(in); err != nil {
		t.Fatal(err)
	}

	out, err := b.Proto()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in.Tensor.FloatData, out.Tensor.FloatData) {
		t.Fatalf("Float data: %+v != %+v", in, out)
	}
}
