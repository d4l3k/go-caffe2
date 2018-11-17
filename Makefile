
caffe2pb: $(wildcard pytorch/caffe2/proto/*.proto)
	-rm -r caffe2pb
	protoc --go_out=import_path=caffe2pb,:. -Ipytorch pytorch/caffe2/proto/{hsm,predictor_consts,caffe2,prof_dag,metanet,caffe2_legacy}.proto
	mv caffe2/proto caffe2pb/
	rmdir caffe2

build: $(wildcard pytorch/*)
	mkdir -p build
	cd build && cmake \
        -DBLAS:STRING='Eigen' \
        \
        -DBUILD_BINARY:BOOL='ON' \
        -DBUILD_DOCS:BOOL='OFF' \
        -DBUILD_PYTHON:BOOL='ON' \
        -DBUILD_SHARED_LIBS:BOOL='ON' \
        \
        -DBUILD_TEST:BOOL='ON' \
        \
        -DCMAKE_CXX_COMPILER='/usr/bin/g++-7' \
        -DCMAKE_C_COMPILER='/usr/bin/gcc-7' \
        -DCMAKE_INSTALL_LIBDIR:PATH='lib' \
        -DCMAKE_INSTALL_PREFIX:PATH='/usr' \
        \
        -DOpenCV_DIR:PATH='/usr/share/OpenCV' \
        ../pytorch
	cd build && make -j16

caffe2_c.syso: caffe2.cc caffe2.h build
	g++ -c -fPIC -O2 -I./build -I./pytorch  -I./pytorch/aten/src/ -I./pytorch/third_party/protobuf/src/ -lcaffe2 -lc10 -o $@ $<
