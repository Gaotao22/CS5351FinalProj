package server

import (
	"connector"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	port string
	addr string
	sw   *connector.SqlWorker
}

var once sync.Once
var server *Server

func GetServer() *Server {
	once.Do(func() {
		server = new(Server)
		server.sw = connector.GetSqlWorker()
	})
	return server
}

func (s *Server) Listen(port string) {
	log.Println("start listen:", port)
	http.HandleFunc("/ping", pong)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	server.sw.QueryData("select * from pingpong.....")
}
