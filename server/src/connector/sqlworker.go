package connector

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type SqlWorker struct {
	path string
	db   *sql.DB
}

var once sync.Once
var sqlworker *SqlWorker

func GetSqlWorker() *SqlWorker {
	once.Do(func() {
		sqlworker = new(SqlWorker)
		sqlworker.initDB()
	})
	return sqlworker
}

func (sw *SqlWorker) initDB() {
	sw.path = "root:iam59!z$@tcp(106.52.119.98:3306)/test"
	db, err := sql.Open("mysql", sw.path)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("connect to database successfully.")
	sw.db = db
}

func (sw *SqlWorker) QueryData(s string) {

}

func (sw *SqlWorker) InsertData(s string) {

}

func (sw *SqlWorker) DeleteData(s string) {

}

func (sw *SqlWorker) UpdateData(s string) {

}
