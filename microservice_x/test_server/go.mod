module test_server

go 1.12

require (
	google.golang.org/grpc v1.22.1
	interfaces v0.0.0
	server v0.0.0
)

replace interfaces => ../generated

replace server => ../server
