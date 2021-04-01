package gate

import (
	"github.com/liwei1dao/lego/lib/modules/gate"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type Options struct {
	gate.Options
	HeartbeatInterval int32
	MaxHeartStopNum   uint8
}

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	if err = this.Options.LoadConfig(settings); err == nil {
		if settings != nil {
			err = mapstructure.Decode(settings, this)
		}
	}
	return
}
