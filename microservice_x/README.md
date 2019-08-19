
## Demo for Protobuffer validation and documentation with GRPC


#### Using https://github.com/envoyproxy/protoc-gen-validate

Please read the documentation at protoc-gen-validate. Protobuffer is great in providing a
static typed interface. However types like strings in interface is hard to 'semantically type'.
With the protoc-gen-validate project you can sepicfy semantic validation in the proto buffer itself

Also added https://github.com/pseudomuto/protoc-gen-doc to generate better documentation
for proto files

# To build in local machine

In a (linux) machine wich has Go installed > 1.12 and  
cd to root directory and run 'make all'

# To build in container
If you just have make installed and don't have any other dependency you can build
the same output in a Docker container
Instead of make all, run 'make docker-make'

# To build A Docker Container

If you need to build a slim Docker container with Go GRPD
run 'make docker-build'

# To Test
cd to integration_test  and run 'go test' 

# To run manually
cd to test_server and run the generated ./test_server
cd to test_client and run the generated ./test_client

