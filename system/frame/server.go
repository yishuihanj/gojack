package frame

import "github.com/gin-gonic/gin"

type Server interface {
	Name()string
	Port() int
	Route() *gin.Engine
	Init(Server)Server
	Run()
}