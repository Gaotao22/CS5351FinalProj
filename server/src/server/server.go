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
	db   *connector.Database
}

var once sync.Once
var server *Server

func GetServer() *Server {
	once.Do(func() {
		server = new(Server)
		server.db = connector.GetDatabase()
	})
	return server
}

func (s *Server) Listen(port string) {
	log.Println("start listen:", s.port)
	http.HandleFunc("/ping", pong)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	server.db.Excute("select * from pingpong.....")
}
