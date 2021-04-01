package gate

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/modules/gate"
)

type AgentMgrComp struct {
	gate.AgentMgrComp
}

func (this *AgentMgrComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	if err = this.AgentMgrComp.Init(service, module, comp, options); err != nil {
		return
	}
	return
}

func (this *AgentMgrComp) Start() (err error) {
	if err = this.AgentMgrComp.Start(); err != nil {
		return
	}
	return
}

func (this *AgentMgrComp) Destroy() (err error) {
	err = this.AgentMgrComp.Destroy()
	return
}

func (this *AgentMgrComp) Build(aId string, uId uint32) (result interface{}, err string) {
	agent := this.Agents.Get(aId)
	if agent == nil {
		err = "No Sesssion found " + aId
		return
	}
	agent.(*Agent).Build(uId)
	result = "success"
	return
}

func (this *AgentMgrComp) UnBuild(aId string) (result interface{}, err string) {
	agent := this.Agents.Get(aId)
	if agent == nil {
		err = "No Sesssion found " + aId
		return
	}
	agent.(*Agent).UnBuild()
	result = "success"
	return
}
