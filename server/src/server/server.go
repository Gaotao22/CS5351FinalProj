package server

import (
	"connector"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

type Server struct {
	sw *connector.SqlWorker
}

var once sync.Once
var server *Server
var router = gin.Default()

func GetServer() *Server {
	once.Do(func() {
		server = new(Server)
		server.sw = connector.GetSqlWorker()
	})
	return server
}

func (s *Server) BindAndListen(port string) {
	// method and url define here
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
	// run and listen
	router.Run(":" + port)
}

func getting(c *gin.Context) {
	log.Println("getting...")
}

func posting(c *gin.Context) {
	log.Println("posting...")
}

func putting(c *gin.Context) {
	log.Println("putting...")
}

func deleting(c *gin.Context) {
	log.Println("deleting...")
}

func patching(c *gin.Context) {
	log.Println("patching...")
}

func head(c *gin.Context) {
	log.Println("head...")
}

func options(c *gin.Context) {
	log.Println("options...")
}
