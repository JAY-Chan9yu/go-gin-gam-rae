package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

// TODO: 같은 서버 안의 다른 프로세스로 떠있는 DB 연결하고 헬스체크

func healthCheckOnDB() bool {
	appConfig, err := godotenv.Read()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"]))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return err != nil
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
