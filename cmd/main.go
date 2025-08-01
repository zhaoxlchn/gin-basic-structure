package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/zhaoxlchn/gin-basic-structure/internal/conf"
	"log"
)

func main() {
	c, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}
	app, cleanup, err := wireApp(c.Server, c.Data.Database, c.Data.Redis)
	if err != nil {
		return
	}
	defer cleanup()
	if err = app(); err != nil {
		log.Printf("server listen err: %v", err)
	}
	log.Println("Service exited...")
}

func newApp(conf *conf.Server, server *gin.Engine) func() error {
	_server := endless.NewServer(conf.Http.Addr, server)
	return func() error {
		return _server.ListenAndServe()
	}
}
