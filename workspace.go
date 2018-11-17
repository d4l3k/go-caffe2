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
