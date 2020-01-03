
CURRENTDIR =  $(shell pwd)
PARENTDIR = $(shell dirname $(CURRENTDIR))

HTTP_PROXY = http://10.158.100.6:8080 #change this as per your proxy

all: clean  _get_deps1 _get_deps2  _build_proto_with_validate _build_go _build_proto_documentation

clean:
	-rm -rf generated/*.pb
# get the protoc validator
# https://github.com/envoyproxy/protoc-gen-validate
_get_deps1:
# Note this is using GO MODULES
# after first build all below will be in Dokcer cache and build will be fast
	go get github.com/golang/protobuf/protoc-gen-go@v1.2.0
	go get  github.com/envoyproxy/protoc-gen-validate@v0.1.0
	go get  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.3.0
_get_deps2:
 # You may need to keep updating the golangci - there are bugs and it is in dev    
	-go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.22.2
#build_proto:
#	protoc -I ./interfaces/ interfaces/test_server/*.proto --go_out=plugins=grpc:./gen_interfaces
_build_proto_with_validate:
	protoc   -I ../interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0/ \
	 --go_out=plugins=grpc:./generated   --validate_out=lang=go:./generated ../interfaces/test_server/*.proto
_build_proto_documentation:
	protoc   -I ../interfaces/ -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0/ \
	--plugin=protoc-gen-doc=${GOPATH}/bin/protoc-gen-doc \
     --doc_out=./doc --doc_opt=html,index.html  ../interfaces/test_server/*.proto
	 #--doc_out=./doc --doc_opt=markdown,index.md

_build_go:
	# let us do the test, get the coverage
	cd integration_test && go test -coverprofile=cover.out -coverpkg=server \
	&& go tool cover -html=cover.out -o cover.html
	# lets run the linter  -ignore on fail
	-golangci-lint run > linter.out
	cd test_server && go build
	cd test_client && go build

# Build a Docker conatiner for the MS - USE this for all
make_in_docker:
	docker build --build-arg  http_proxy=$(HTTP_PROXY) \
	 --build-arg https_proxy=$(HTTP_PROXY)  -t mysample_go_grpc:1.0 -f  Dockerfile $(PARENTDIR) 


# Build inside docker with volume mapping to generate the documentation or test locally
# !! Note !! that when you run in docker the files that is generated in volume map will have
# docker user permissions  EX below
#-rw------- 1 root root  1261 Jan  3 11:33 go.mod
#-rw------- 1 root root 49400 Jan  3 11:27 go.sum

make_in_docker_with_vol_map:
	echo "ENV variables check" $(CURRENTDIR) $(HTTP_PROXY)  $(PARENTDIR)

	docker run  --env http_proxy=$(HTTP_PROXY) --env https_proxy=$(HTTP_PROXY) \
	--rm -it -v $(PARENTDIR):/go/golang_grpc_test grpc/go:1.12 /bin/sh -c \
	 'cd golang_grpc_test/microservice_x && make all'

