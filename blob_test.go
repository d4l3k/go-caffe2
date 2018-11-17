package caffe2

import (
	"testing"
)

func TestBlob(t *testing.T) {
	w := NewWorkspace()
	b := w.CreateBlob("foo")
	b.Reset()
	t.Log("type name", b.TypeName())
}
