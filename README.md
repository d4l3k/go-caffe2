# go-caffe2

WIP caffe2 bindings for Go

You need to install caffe2 to your system first.

```sh
$ go get -u github.com/d4l3k/go-caffe2
```

## Development

Check the code out with submodules

```sh
$ git clone --recurse-submodules -j16 git@github.com:d4l3k/go-caffe2.git
```

Install go protoc: https://github.com/golang/protobuf#installation

Generate sources:

```
make
```
