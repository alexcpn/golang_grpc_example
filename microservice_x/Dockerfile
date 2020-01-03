# Build a slim GRPC Go Container
# using multistage build
# https://docs.docker.com/develop/develop-images/multistage-build/
# FROM grpc/go:1.12 AS build --> Dockerfile is in docker-go1.12 folder
FROM alexcpn/grpc-go:1.12 AS build
WORKDIR /golang_grpc_test
COPY . /golang_grpc_test
RUN cd /golang_grpc_test/microservice_x && make all

# final stage
FROM alpine
WORKDIR /app
COPY --from=build /golang_grpc_test/microservice_x/test_server/test_server /app/
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT ./test_server
