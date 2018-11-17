package caffe2

// #include "caffe2.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/d4l3k/go-caffe2/caffe2pb"
)

type Workspace struct {
	w C.C2Workspace
}

func finalizeWorkspace(w *Workspace) {
	C.free(unsafe.Pointer(w.w))
}

// NewWorkspace creates a new workspace.
// The underlying C++ workspace is freed via a finalizer and all child Blobs and
// Nets have a pointer to the workspace so there won't ever be any unsafe memory
// accesses.
func NewWorkspace() *Workspace {
	w := &Workspace{
		w: C.C2WorkspaceInit(),
	}
	runtime.SetFinalizer(w, finalizeWorkspace)
	return w
}

func (w *Workspace) CreateBlob(blob string) *Blob {
	blobC := C.CString(blob)
	defer C.free(unsafe.Pointer(blobC))

	return &Blob{
		b: C.C2WorkspaceCreateBlob(w.w, blobC),
		w: w,
	}
}

func (w *Workspace) CreateLocalBlob(blob string) *Blob {
	blobC := C.CString(blob)
	defer C.free(unsafe.Pointer(blobC))

	return &Blob{
		b: C.C2WorkspaceCreateLocalBlob(w.w, blobC),
		w: w,
	}
}

func (w *Workspace) GetBlob(blob string) *Blob {
	blobC := C.CString(blob)
	defer C.free(unsafe.Pointer(blobC))

	b := C.C2WorkspaceGetBlob(w.w, blobC)
	if b == nil {
		return nil
	}

	return &Blob{
		b: b,
		w: w,
	}
}

func (w *Workspace) HasBlob(blob string) bool {
	blobC := C.CString(blob)
	defer C.free(unsafe.Pointer(blobC))

	return bool(C.C2WorkspaceHasBlob(w.w, blobC))
}

func (w *Workspace) RemoveBlob(blob string) error {
	blobC := C.CString(blob)
	defer C.free(unsafe.Pointer(blobC))

	if !bool(C.C2WorkspaceRemoveBlob(w.w, blobC)) {
		return errors.Errorf("failed to remove blob %q", blob)
	}
	return nil
}

func (w *Workspace) RunNet(name string) error {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	if !bool(C.C2WorkspaceRunNet(w.w, nameC)) {
		return errors.Errorf("failed to run net %q", name)
	}
	return nil
}

// createNet creates a net from a serialized NetDef proto.
func (w *Workspace) createNet(netDef []byte) *Net {
	return &Net{
		n: C.C2WorkspaceCreateNet(
			w.w, (*C.char)(unsafe.Pointer(&netDef[0])), C.int(len(netDef)),
		),
		w: w,
	}
}

func (w *Workspace) CreateNet(net *caffe2pb.NetDef) (*Net, error) {
	buf, err := proto.Marshal(net)
	if err != nil {
		return nil, err
	}
	return w.createNet(buf), nil
}

func (w *Workspace) GetNet(name string) *Net {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return &Net{
		n: C.C2WorkspaceGetNet(w.w, nameC),
		w: w,
	}
}

func (w *Workspace) DeleteNet(name string) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	C.C2WorkspaceDeleteNet(w.w, nameC)
}

func (w *Workspace) RunPlan(plan *caffe2pb.PlanDef) error {
	buf, err := proto.Marshal(plan)
	if err != nil {
		return err
	}
	if !w.runPlan(buf) {
		return errors.Errorf("failed to run plan %q", plan.GetName())
	}
	return nil
}

// runPlan runs a serialized PlanDef proto.
func (w *Workspace) runPlan(planDef []byte) bool {
	return bool(C.C2WorkspaceRunPlan(
		w.w, (*C.char)(unsafe.Pointer(&planDef[0])), C.int(len(planDef)),
	))
}
