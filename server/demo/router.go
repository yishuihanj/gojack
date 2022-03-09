package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router()*gin.Engine{
	route:=gin.New()
	initRoute(route)
	return route
}

func initRoute(route *gin.Engine){
	demo:=route.Group("/demo")
	{
		demo.GET("/greet", func(c *gin.Context) {
			c.JSON(http.StatusOK,"hello world")
		})
	}
}
