package frame

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type BaseServer struct {
	Name string //服务器名称
	Port int //服务器端口
	router *gin.Engine
}

// Init 加载必要的服务组件
func (s *BaseServer)Init(ins Server)Server{
	s.Name=ins.Name()
	s.Port=ins.Port()
	s.router=ins.Route()
	return ins
}

func (s *BaseServer)Run(){
	if err:=s.router.Run(fmt.Sprintf(":%d",s.Port));err!=nil{
		log.Fatalf("[Running] - HttpServer启动失败,%v",err)
	}

}