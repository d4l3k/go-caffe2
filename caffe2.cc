#include "caffe2.h"
#include "caffe2/core/blob_serialization.h"
#include "caffe2/core/workspace.h"

caffe2::Workspace* cppW(C2Workspace w) {
  return static_cast<caffe2::Workspace*>(w);
}

caffe2::Blob* cppB(C2Blob b) { return static_cast<caffe2::Blob*>(b); }

C2Workspace C2WorkspaceInit(void) {
  caffe2::Workspace* w = new caffe2::Workspace();
  return (void*)w;
}

C2Blob C2WorkspaceCreateBlob(C2Workspace w, char* name) {
  return (void*)cppW(w)->CreateBlob(std::string(name));
}

C2Blob C2WorkspaceCreateLocalBlob(C2Workspace w, char* name) {
  return (void*)cppW(w)->CreateLocalBlob(std::string(name));
}

C2Blob C2WorkspaceGetBlob(C2Workspace w, char* name) {
  return (void*)cppW(w)->GetBlob(std::string(name));
}

bool C2WorkspaceHasBlob(C2Workspace w, char* name) {
  return cppW(w)->HasBlob(std::string(name));
}

bool C2WorkspaceRemoveBlob(C2Workspace w, char* name) {
  return cppW(w)->RemoveBlob(std::string(name));
}

void C2BlobReset(C2Blob b) { cppB(b)->Reset(); }

const char* C2BlobTypeName(C2Blob b) { return cppB(b)->TypeName(); }

void C2BlobDeserialize(C2Blob b, char* content, int len) {
  caffe2::DeserializeBlob(std::string(content, len), cppB(b));
}

char* C2BlobSerialize(C2Blob b) {
  std::string raw;
  caffe2::SerializeBlob(*cppB(b), raw);
  return strdup(raw.c_str());
}
