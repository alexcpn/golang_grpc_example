module test_server

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	google.golang.org/grpc v1.22.1
	interfaces v0.0.0
	server v0.0.0
)

replace interfaces => ../interfaces

replace server => ../server
