package connector

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type SqlWorker struct {
	db *gorm.DB
}

var once sync.Once
var sqlworker *SqlWorker
var sqlConnectionPath = "root:iam59!z$@tcp(106.52.119.98:3306)/test"

func GetSqlWorker() *SqlWorker {
	once.Do(func() {
		sqlworker = new(SqlWorker)
		sqlworker.initDB()
	})
	return sqlworker
}

func (sw *SqlWorker) initDB() {
	db, err := gorm.Open("mysql", sqlConnectionPath)
	if err != nil {
		panic(err)
	}

	log.Println("connect to database successfully.")
	sw.db = db
}
