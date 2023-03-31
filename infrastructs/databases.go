package infrastructs

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
)

// go에서는 앞에 글자가 소문자면 private, 대문자여야 public하게 사용할 수 있음.
// 앞글자가 대문자여야 다른곳에서 불러다가 사용가능
func GetDBConnection() *sql.DB {
	appConfig, _ := godotenv.Read()
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
	return db
}
