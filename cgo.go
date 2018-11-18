package caffe2

// #cgo LDFLAGS: -lcaffe2 -lglog -lprotobuf
// #include "caffe2.h"
import "C"

func init() {
	C.C2Init()
}
