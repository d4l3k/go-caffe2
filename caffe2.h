#ifndef CAFFE2_H_
#define CAFFE2_H_

#include <stddef.h>
#include <stdint.h>
#ifdef __cplusplus
extern "C" {
#endif

typedef void* C2Workspace;
// Result from C2WorkspaceInit must be freed.
C2Workspace C2WorkspaceInit(void);

#ifdef __cplusplus
} /* end extern "C" */
#endif

#endif  // CAFFE2_H_
