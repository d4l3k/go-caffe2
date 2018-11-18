# go-caffe2 [![GoDoc](https://godoc.org/github.com/d4l3k/go-caffe2?status.svg)](https://godoc.org/github.com/d4l3k/go-caffe2)

This provides caffe2 bindings for Go.

You need to install caffe2 to your system first.

```sh
$ go get -u github.com/d4l3k/go-caffe2
```

See TestNet for an example of creating a new net with blobs and running it:
https://github.com/d4l3k/go-caffe2/blob/master/net_test.go

## Development

Check the code out with submodules

```sh
$ git clone --recurse-submodules -j16 git@github.com:d4l3k/go-caffe2.git
```

Install go protoc: https://github.com/golang/protobuf#installation

Generate protobuf sources and test:

```
make
```

## License

go-caffe2 is licensed under the MIT license.
