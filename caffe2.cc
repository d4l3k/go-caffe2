#include "caffe2.h"
#include "caffe2/core/workspace.h"

C2Workspace C2WorkspaceInit(void){
  caffe2::Workspace* w = new caffe2::Workspace();
  return (void*)w;
}
