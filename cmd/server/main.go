package main

import (
	"chikuwait/go-grpc-sample/pb"
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	pb.UnimplementedGreeterServer
}

func main() {
	port := 50051
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	svc := &service{}
	pb.RegisterGreeterServer(server, svc)

	reflection.Register(server)
	server.Serve(listen)
}
func (*service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + in.GetName()}, nil
}
