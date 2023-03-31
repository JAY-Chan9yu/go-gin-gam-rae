package cosmetic

import (
	"context"
	pb "github.com/JAY-Chan9yu/go-gin-gam-rae/proto/hello"
	"log"
)

type Hello struct {
	pb.UnimplementedGreeterServer
}

func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
