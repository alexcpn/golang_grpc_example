# Make a Docker container for Go 1.12 with GRPC first

Note This is the base container -only needed if THIS Container (grpc/go:1.12) does not exist
in internet/ from Google (at the time of writing this was not there)

Example docker build --build-arg  http_proxy=$(HTTP_PROXY) \
	  --build-arg https_proxy=$(HTTP_PROXY)  -t alexcpn/grpc-go:1.12 - < Dockerfile.go_grpc.1.12 
(change the tag in Dockerfile too)