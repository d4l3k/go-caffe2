package caffe2

import (
	"reflect"
	"testing"

	"github.com/d4l3k/go-caffe2/caffe2pb"
	"github.com/gogo/protobuf/proto"
)

func TestNet(t *testing.T) {
	w := NewWorkspace()

	a := w.CreateBlob("a")
	dataType := caffe2pb.TensorProto_FLOAT
	if err := a.FromProto(&caffe2pb.BlobProto{
		Type: proto.String("Tensor"),
		Tensor: &caffe2pb.TensorProto{
			DataType:  &dataType,
			Dims:      []int64{3},
			FloatData: []float32{1, 2, 3},
		},
	}); err != nil {
		t.Fatal(err)
	}

	w.CreateNet(&caffe2pb.NetDef{
		Name: proto.String("test"),
		Op: []*caffe2pb.OperatorDef{
			{
				Type:   proto.String("Negative"),
				Input:  []string{"a"},
				Output: []string{"b"},
			},
		},
	})

	if err := w.RunNet("test"); err != nil {
		t.Fatal(err)
	}

	blob, err := w.GetBlob("b").Proto()
	if err != nil {
		t.Fatal(err)
	}
	want := []float32{-1, -2, -3}
	out := blob.Tensor.FloatData
	if !reflect.DeepEqual(want, out) {
		t.Fatalf("wrong output: got %+v; wanted %+v", out, want)
	}
}
