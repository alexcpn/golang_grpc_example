
CURRENTDIR = /home/alex/coding/golang_grpc_test
all: build_proto build_go
build_proto:
	protoc -I interfaces/ interfaces/*.proto --go_out=plugins=grpc:./interfaces
build_go:
	cd test_server && go build
	cd ..
	cd test_client $$ go build
# Build inside docker	
docker-build:
	echo $(CURRENTDIR)
	docker run --rm -it -v $(CURRENTDIR)\:/go namely/protoc-all make all