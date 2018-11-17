package caffe2

import "testing"

func TestWorkspaceBlob(t *testing.T) {
	w := NewWorkspace()
	w.CreateBlob("foo")
	w.CreateLocalBlob("bar")
	if !w.HasBlob("foo") {
		t.Fatal("must have blob foo")
	}
	if !w.HasBlob("bar") {
		t.Fatal("must have blob bar")
	}
	if w.HasBlob("unknown") {
		t.Fatal("blob shouldn't exist")
	}
	if w.GetBlob("foo") == nil {
		t.Fatal("existing blob shouldn't be nil")
	}
	if w.GetBlob("unknown") != nil {
		t.Fatal("unknown blob should be nil")
	}

	if !w.RemoveBlob("foo") {
		t.Fatal("failed to remove blob")
	}
	if w.HasBlob("foo") {
		t.Fatal("blob wasn't removed")
	}
}
