
caffe2pb: $(wildcard pytorch/caffe2/proto/*.proto)
	protoc --go_out=import_path=caffe2pb,:. -Ipytorch pytorch/caffe2/proto/{hsm,predictor_consts,caffe2,prof_dag,metanet,caffe2_legacy}.proto
	mv caffe2/proto caffe2pb
	rmdir caffe2
