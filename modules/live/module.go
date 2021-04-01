package live

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/live"
)

func NewModule() (module core.IModule) {
	m := new(Live)
	return m
}

type Live struct {
	live.Live
}

func (this *Live) GetType() core.M_Modules {
	return lib.SM_LiveModule
}

func (this *Live) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	err = this.Live.Init(service, module, options)
	return
}
