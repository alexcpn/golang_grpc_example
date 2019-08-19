
## Demo for Protobuffer validation and documentation with GRPC

Also you can use this as a sample template to generate Go Microservices. The interfaces folder
is kept out so that multiple microservice interfaces can be added in different directories here
and between microservice only interfaces need to be shared

See also https://medium.com/techlogs/micro-services-git-d840305e6220

#### Using https://github.com/envoyproxy/protoc-gen-validate

Please read the documentation at protoc-gen-validate. Protobuffer is great in providing a
static typed interface. However types like strings in interface is hard to 'semantically type'.
With the protoc-gen-validate project you can sepicfy semantic validation in the proto buffer itself

Also added https://github.com/pseudomuto/protoc-gen-doc to generate better documentation
for proto files

# To build in local machine

In a (linux) machine wich has Go installed > 1.12 and  
cd to microservice_x directory and run 'make all'

# To build Binary/ run tests in Go container

If you just have 'make' installed and don't have any other dependency even Go you can build
the same output in a Docker container
First cd to microservice_x and  run 'make docker-make'

Note - This works in Windows 10 onwards too (Docker support). You can install make via Cholclatey
(https://stackoverflow.com/a/54086635/429476) and run . Only the shell scripts for CURRENTDIR and PARENTDIR
needs to be adopted. I recommend doing this in Linux though.

# To build A Docker Container

If you need to build a slim Docker container with Go GRPC
cd to microservice_x directory and run 'make docker-build'

# To Test

cd to microservice_x directory
cd to integration_test  and run 'go test' 

# To run manually

cd to microservice_x directory
cd to test_server and run the generated ./test_server
cd to test_client and run the generated ./test_client

