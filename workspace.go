package caffe2

// #include <stdio.h>
// #include <stdlib.h>
// #include "caffe2.h"
import "C"
import (
	"runtime"
	"unsafe"
)

type Workspace struct {
	w C.C2Workspace
}

func finalizeWorkspace(w *Workspace) {
	C.free(unsafe.Pointer(w.w))
}

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

func (w *Workspace) RemoveBlob(blob string) bool {
	blobC := C.CString(blob)
	defer C.free(unsafe.Pointer(blobC))

	return bool(C.C2WorkspaceRemoveBlob(w.w, blobC))
}
