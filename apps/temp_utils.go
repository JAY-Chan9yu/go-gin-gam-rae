/*
gin framework 관련 테스트 했던 코드들
*/
package apps

import (
	infra "github.com/JAY-Chan9yu/go-gin-gam-rae/infrastructs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Config struct {
	pk            int
	googleVersion string
	iosVersion    string
}

func healthCheckOnDB() bool {
	db := infra.GetDBConnection()
	defer db.Close()

	var config Config
	err := db.QueryRow("SELECT pk, googleVersion, iosVersion FROM config WHERE pk = 1").Scan(&config.pk, &config.googleVersion, &config.iosVersion)
	if err != nil {
		log.Fatal(err)
	}

	return err == nil
}

func healthCheckOnUrl(url string) bool {
	resp, _ := http.Get(url)
	return resp.StatusCode == 200
}

func HealthChecker(c *gin.Context) {
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
