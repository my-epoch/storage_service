FROM golang:latest

RUN mkdir "build"
ADD . /build
WORKDIR /build
RUN cd /build

RUN go build cmd/grpc_server/grpc_server.go
EXPOSE 50051
ENTRYPOINT "./grpc_server"