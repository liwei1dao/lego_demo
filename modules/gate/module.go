package gate

import (
	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/gate"
	"github.com/liwei1dao/lego/sys/log"
)

func NewModule() gate.IGateModule {
	m := new(Gate)
	return m
}

type Gate struct {
	gate.Gate
	service       base.IClusterService
	TcpServerComp *gate.TcpServerComp
	options       *Options
}

func (this *Gate) GetType() core.M_Modules {
	return lib.SM_GateModule
}

func (this *Gate) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *Gate) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	err = this.Gate.Init(service, module, options)
	this.service = service.(base.IClusterService)
	this.options = options.(*Options)
	return
}

func (this *Gate) Start() (err error) {
	if err = this.Gate.Start(); err != nil {
		return
	}
	this.service.RegisterGO(gate.RPC_GateAgentUnBuild, this.UnBuild)
	this.service.RegisterGO(gate.RPC_GateSendMsg, this.SendMsg)
	return
}

func (this *Gate) OnInstallComp() {
	this.Gate.OnInstallComp()
	this.AgentMgrComp = this.RegisterComp(new(AgentMgrComp)).(*AgentMgrComp)
	this.TcpServerComp = this.RegisterComp(new(gate.TcpServerComp)).(*gate.TcpServerComp)
	this.LocalRouteMgrComp.SetNewSession(NewLocalSession)
	this.RemoteRouteMgrComp.SetNewSession(NewRemoteSession)
	this.TcpServerComp.NewTcpAgent = NewAgent
}

//需重构处理  内部函数为重构代码
//代理链接
func (this *Gate) Connect(a gate.IAgent) {
	log.Debugf("有新的用户链接进来IP:[%s] Id:[%s]", a.IP(), a.Id())
	this.AgentMgrComp.Connect(a)
}

//代理关闭
func (this *Gate) DisConnect(a gate.IAgent) {
	log.Debugf("有用户链接断开IP:[%s] Id:[%s]", a.IP(), a.Id())
	this.AgentMgrComp.DisConnect(a)
}

//解绑用户
func (this *Gate) Build(aId string, uId uint32) (result interface{}, err string) {
	return this.AgentMgrComp.(*AgentMgrComp).Build(aId, uId)
}

//解绑用户
func (this *Gate) UnBuild(aId string) (result interface{}, err string) {
	return this.AgentMgrComp.(*AgentMgrComp).UnBuild(aId)
}
