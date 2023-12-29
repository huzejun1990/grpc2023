// @Author huzejun 2023/12/30 5:22:00
package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc2023/helloworld/proto"
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
	birthday := timestamppb.New(time.Now())
	ctx := context.Background()
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

	r, err := c.SayHello(ctx, in)
	if err != nil {
		//log.Println(err)
		log.Fatal(err)
		return
	}
	//fmt.Println(r.Msg)
	log.Println(r.Msg)
}
