package connector

import (
	"log"
	"sync"
)

type Database struct {
	addr   string
	port   int
	dbtype string
}

var once sync.Once
var db *Database

func GetDatabase() *Database {
	once.Do(func() {
		db = new(Database)
	})
	return db
}

func init() {
	// connect to database
	log.Println("connect to database")
}

func (db *Database) Excute(sql string) {
	// execute sql script
	log.Println("excute:", sql)
}
