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

	if err := w.RemoveBlob("foo"); err != nil {
		t.Fatal("failed to remove blob", err)
	}
	if w.HasBlob("foo") {
		t.Fatal("blob wasn't removed")
	}
}

func TestWorkspaceChild(t *testing.T) {
	parent := NewWorkspace()
	parent.CreateBlob("parent")

	child := parent.NewChild()
	child.CreateBlob("child")

	parent.CreateBlob("parent2")
	if !child.HasBlob("parent") || !child.HasBlob("parent2") {
		t.Fatal("missing blobs from parent net")
	}

	if parent.HasBlob("child") {
		t.Fatal("child blobs leaking to parent")
	}
}
