package caffe2

// #include <stdio.h>
// #include <stdlib.h>
// #include "caffe2.h"
import "C"
import (
	"log"
	"unsafe"

	"github.com/d4l3k/go-caffe2/caffe2pb"
	"github.com/gogo/protobuf/proto"
)

type Blob struct {
	b C.C2Blob

	// we hold on to a pointer to the workspace to make sure it isn't freed before
	// the blob is since we don't own the blob.
	w *Workspace
}

func (b *Blob) Reset() {
	C.C2BlobReset(b.b)
}

func (b *Blob) TypeName() string {
	return C.GoString(C.C2BlobTypeName(b.b))
}

func (b *Blob) FromProto(p *caffe2pb.BlobProto) error {
	buf, err := proto.Marshal(p)
	if err != nil {
		return err
	}
	b.Deserialize(buf)
	return nil
}

func (b *Blob) Proto() (*caffe2pb.BlobProto, error) {
	buf := b.Serialize()
	log.Printf("length: %v, %q", len(buf), buf)
	var p caffe2pb.BlobProto
	if err := proto.Unmarshal(buf, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (b *Blob) Deserialize(content []byte) {
	C.C2BlobDeserialize(
		b.b, (*C.char)(unsafe.Pointer(&content[0])), C.int(len(content)),
	)
}

// TODO: make this not do 4 memory copies to get from C to Go
func (b *Blob) Serialize() []byte {
	var len C.int
	raw := C.C2BlobSerialize(b.b, &len)
	defer C.free(unsafe.Pointer(raw))

	return []byte(C.GoStringN(raw, len))
}
