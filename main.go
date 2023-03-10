package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthChecker(c *gin.Context) {
	/***
	헬스 체크 기획

	헬스 체크리스트(url) 받아서 그 주소들 헬스체크 결과 response 로 전달?
	***/
	naverResp, _ := http.Get("https://naver.com/")
	hwahaeResp, _ := http.Get("https://www.hwahae.co.kr/")
	var naverisOk bool
	var hwahaeisOk bool

	if naverResp.StatusCode == 200 {
		naverisOk = true
	}

	if hwahaeResp.StatusCode == 200 {
		hwahaeisOk = true
	}

	c.JSON(http.StatusOK, gin.H{
		"naver_status":  naverisOk,
		"hwahae_status": hwahaeisOk,
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
