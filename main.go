package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/JAY-Chan9yu/go-gin-gam-rae/proto"
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

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
