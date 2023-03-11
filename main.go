package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type Config struct {
	pk            int
	googleVersion string
	iosVersion    string
}

func healthCheckOnDB() bool {
	appConfig, err := godotenv.Read()

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

func main() {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/health-check", healthChecker)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
