package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"net/http"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.simple"),
		web.Handler(routers()),
		web.Version("latest"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func routers() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello /hello page.")
	})

	return router
}
