#ifndef CAFFE2_H_
#define CAFFE2_H_

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void* C2Workspace;
typedef void* C2Blob;
typedef void* C2Net;

// Workspace

// Result from C2WorkspaceInit must be freed by the caller.
C2Workspace C2WorkspaceInit(void);

C2Workspace C2WorkspaceChild(C2Workspace w);

C2Blob C2WorkspaceCreateBlob(C2Workspace w, char* name);
C2Blob C2WorkspaceCreateLocalBlob(C2Workspace w, char* name);
C2Blob C2WorkspaceGetBlob(C2Workspace w, char* name);
bool C2WorkspaceHasBlob(C2Workspace w, char* name);
bool C2WorkspaceRemoveBlob(C2Workspace w, char* name);

bool C2WorkspaceRunNet(C2Workspace w, char* name);
C2Net C2WorkspaceCreateNet(C2Workspace w, char* net, int len);
C2Net C2WorkspaceGetNet(C2Workspace w, char* name);
void C2WorkspaceDeleteNet(C2Workspace w, char* name);

bool C2WorkspaceRunPlan(C2Workspace w, char* plan, int len);

// Blob

void C2BlobReset(C2Blob b);
const char* C2BlobTypeName(C2Blob b);
void C2BlobDeserialize(C2Blob b, char* content, int len);

// result must be freed by caller
char* C2BlobSerialize(C2Blob b, int* len);

#ifdef __cplusplus
} /* end extern "C" */
#endif

#endif  // CAFFE2_H_
