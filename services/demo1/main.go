package main

import (
	"flag"
	"fmt"
	"lego_demo/comm"
	"lego_demo/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/s_comps"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/timewheel"
)

var (
	sID = flag.String("sID", "demo1", "获取需要启动的服务id,id不同,读取的配置文件也不同") //启动服务的Id
)

func main() {
	flag.Parse()
	s := NewService(
		cluster.SetTag(comm.ServiceTag),
		cluster.SetId(*sID),
		cluster.SetType(comm.LoginService),
		cluster.SetCategory(comm.GateCategory),
		cluster.SetVersion(1),
		cluster.SetLogLvel(log.InfoLevel),
		cluster.SetDebugMode(true),
	)
	s.OnInstallComp( //装备组件
		s_comps.NewGateRouteComp(),
	)
	lego.Run(s) //运行模块

}

func NewService(ops ...cluster.Option) core.IService {
	s := new(LoginService)
	s.Configure(ops...)
	return s
}

type Demo1Service struct {
	services.HuYuServiceBase
}

func (this *LoginService) InitSys() {
	this.HuYuServiceBase.InitSys()
	if err := timewheel.OnInit(this.Service.GetSettings().Sys["timewheel"]); err != nil {
		panic(fmt.Sprintf("初始化timewheel系统失败 err:%s", err.Error()))
	}
}
