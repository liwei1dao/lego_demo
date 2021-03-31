package gate

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/gate"
)

func NewModule() (module core.IModule) {
	m := new(Gate)
	return m
}

type Gate struct {
	gate.Gate
}

func (this *Gate) GetType() core.M_Modules {
	return lib.SM_GateModule
}
