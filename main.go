package main

import (
	"context"
	"fmt"
	cosmetic "github.com/JAY-Chan9yu/go-gin-gam-rae/apps/cosmetic"
	hello "github.com/JAY-Chan9yu/go-gin-gam-rae/apps/hello"
	cosmeticPb "github.com/JAY-Chan9yu/go-gin-gam-rae/proto/cosmetic"
	helloPb "github.com/JAY-Chan9yu/go-gin-gam-rae/proto/hello"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	// TCP 소켓 열기
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	cosmeticPb.RegisterCosmeticServiceServer(s, &cosmetic.Cosmetic{})
	helloPb.RegisterGreeterServer(s, &hello.Hello{})

	/****************************
	         Serving gRPC
	****************************/
	log.Println("Serving gRPC on 127.0.0.1:9000")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:9000",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux() // mux.Handle에 gwmux를 등록시켜서 rest 요청이 들어올 때 마다 proto에 작성해준 요청에 맞게 연결
	err = cosmeticPb.RegisterCosmeticServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register Cosmetic Proto gateway:", err)
	}

	err = helloPb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register Hello Proto gateway:", err)
	}

	/****************************
	    Serving gRPC-Gateway
	****************************/
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://127.0.0.1:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
