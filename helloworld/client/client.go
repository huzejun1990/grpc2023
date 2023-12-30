// @Author huzejun 2023/12/30 5:22:00
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc2023/helloworld/proto"
	"io"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	//sayHello(c)
	//sayHelloClientStream(c)
	sayHelloTwoWayStream(c)
}

func getHelloRequest() *proto.HelloRequest {
	birthday := timestamppb.New(time.Now())
	any1, _ := anypb.New(birthday)
	in := &proto.HelloRequest{
		Name:     "nick",
		Gender:   proto.Gender_MALE,
		Age:      18,
		Birthday: birthday,
		Hobys:    []string{"羽毛球", "篮球"},
		Addr: &proto.Address{
			Province: "江苏",
			City:     "南京",
		},
		Data: map[string]*anypb.Any{
			"a": any1,
		},
	}
	return in
}

func sayHello(c proto.GreeterClient) {
	ctx := context.Background()
	in := getHelloRequest()
	r, err := c.SayHello(ctx, in)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(r.Msg)
}

func sayHelloClientStream(c proto.GreeterClient) {
	ctx := context.Background()
	list := []*proto.HelloRequest{
		getHelloRequest(), getHelloRequest(), getHelloRequest(),
	}
	stream, err := c.SayHelloClientStream(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("client recv: %v\n", reply)

}

func sayHelloTwoWayStream(c proto.GreeterClient) {
	ctx := context.Background()
	list := []*proto.HelloRequest{
		getHelloRequest(), getHelloRequest(), getHelloRequest(),
	}
	stream, err := c.SayHelloTwoWayStream(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	var done = make(chan struct{}, 0)

	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Println(err)
				close(done)
				return
			}
			fmt.Printf("client recv: %v\n", reply.Msg)
		}
	}()

	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	stream.CloseSend()
	<-done
}
