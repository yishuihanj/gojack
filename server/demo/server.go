package demo

import (
	"github.com/gin-gonic/gin"
	"gojack/system/frame"
	"gojack/system/global"
)

type server struct {
	frame.BaseServer
}


func NewServer()*server{
	return  &server{}
}

func (s *server)Name()string  {
	return global.DemoServer
}

func (s *server)Port()int {
	return global.DemoServerPort
}
func(s *server) Route()*gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	return Router()
}
