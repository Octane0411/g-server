package router

import (
	"g-server/server/ws"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	/*	r := gin.New()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())*/
	r := gin.Default()

	h := ws.NewHub()
	go h.Run()

	r.GET("/", ws.Home)
	//r.GET("/echo", ws.Echo)
	r.GET("/echo", func(c *gin.Context) {
		ws.HttpController(c, h)
	})
	return r
}
