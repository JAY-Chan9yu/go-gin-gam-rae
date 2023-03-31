package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"

	pb "github.com/JAY-Chan9yu/go-gin-gam-rae/proto"
	"google.golang.org/grpc"
)

type Config struct {
	pk            int
	googleVersion string
	iosVersion    string
}

func healthCheckOnDB() bool {
	appConfig, _ := godotenv.Read()

	// 이거 찾아보기
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
			appConfig["MYSQL_USER"],
			appConfig["MYSQL_PASSWORD"],
			appConfig["MYSQL_PROTOCOL"],
			appConfig["MYSQL_HOST"],
			appConfig["MYSQL_PORT"],
			appConfig["MYSQL_DBNAME"],
		),
	)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var config Config
	err = db.QueryRow("SELECT pk, googleVersion, iosVersion FROM config WHERE pk = 1").Scan(&config.pk, &config.googleVersion, &config.iosVersion)
	if err != nil {
		log.Fatal(err)
	}

	return err == nil
}

func healthCheckOnUrl(url string) bool {
	resp, _ := http.Get(url)
	return resp.StatusCode == 200
}

func healthChecker(c *gin.Context) {
	/***
	헬스 체크 기획

	헬스 체크리스트(url) 받아서 그 주소들 헬스체크 결과 response 로 전달?
	***/
	c.JSON(http.StatusOK, gin.H{
		"naver_status":  healthCheckOnUrl("https://naver.com"),
		"hwahae_status": healthCheckOnUrl("https://www.hwahae.co.kr"),
		"db_status":     healthCheckOnDB(),
	})
}

type server struct{ pb.UnimplementedGreeterServer }

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) ListCosmetics(ctx context.Context, in *pb.ListCosmeticsRequest) (*pb.ListCosmeticsResponse, error) {
	appConfig, _ := godotenv.Read()
	// 이거 찾아보기
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
			appConfig["MYSQL_USER"],
			appConfig["MYSQL_PASSWORD"],
			appConfig["MYSQL_PROTOCOL"],
			appConfig["MYSQL_HOST"],
			appConfig["MYSQL_PORT"],
			appConfig["MYSQL_DBNAME"],
		),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT COUNT(*) FROM cosmetic"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var cnt int32
	fmt.Println(rows)
	for rows.Next() {
		if err := rows.Scan(&cnt); err != nil {
			log.Fatal(err)
		}
	}

	query = "SELECT * FROM cosmetic"
	rows, err = db.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
	cosmetics := make([]*pb.Cosmetics, cnt)

	i := 0
	for rows.Next() {
		var (
			id          string
			name        string
			description string
			price       int32
		)
		if err := rows.Scan(&id, &name, &description, &price); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s\n", id, name)
		cosmetics[i] = &pb.Cosmetics{Id: id, Name: name, Description: description, Price: price}
		i += 1
	}

	return &pb.ListCosmeticsResponse{Data: cosmetics}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	//reflection.Register(s)
	//
	//log.Printf("server listening at %v", lis.Addr())
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
	log.Println("Serving gRPC on 0.0.0.0:9000")
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

	gwmux := runtime.NewServeMux()

	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://127.0.0.1:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
