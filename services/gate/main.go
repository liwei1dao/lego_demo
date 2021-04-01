package main

import (
	"flag"
	"lego_demo/modules/gate"
	"lego_demo/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

var (
	sID = flag.String("sID", "gate_1", "获取需要启动的服务id,id不同,读取的配置文件也不同") //启动服务的Id
)

func main() {
	flag.Parse()
	s := NewService(
		cluster.SetId(*sID),
	)
	s.OnInstallComp( //装备组件
	)
	lego.Run(s, //运行模块
		gate.NewModule(),
	)

}

func NewService(ops ...cluster.Option) core.IService {
	s := new(Demo1Service)
	s.Configure(ops...)
	return s
}

type Demo1Service struct {
	services.ServiceBase
}

func (this *Demo1Service) InitSys() {
	this.ServiceBase.InitSys()
}
