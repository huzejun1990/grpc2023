// @Author huzejun 2023/12/30 5:01:00
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc2023/helloworld/proto"
	"io"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "")
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("server recv : %v\n", in)
	return &proto.HelloReply{
		Msg: "hello client",
	}, nil
}

func (server) SayHelloClientStream(stream proto.Greeter_SayHelloClientStreamServer) error {
	var i = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloReply{Msg: fmt.Sprintf("total recv count : %d ", i)})
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("server recv : %v\n", req)
		i++
	}

	//return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}

func (server) SayHelloServerStream(in *proto.HelloRequest, stream proto.Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}

func (server) SayHelloTwoWayStream(stream proto.Greeter_SayHelloTwoWayStreamServer) error {
	var i = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("server recv : %v\n", req)
		stream.Send(&proto.HelloReply{Msg: fmt.Sprintf("%dth request od\n", i)})
		i++
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()

	proto.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %s\n", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
