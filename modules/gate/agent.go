package gate

import (
	"lego_demo/pb"
	"time"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/proto"
	"github.com/liwei1dao/lego/sys/timewheel"
	"github.com/liwei1dao/lego/lib/modules/gate"
)

func NewAgent(gate gate.IGateModule, coon gate.IConn) (gate.IAgent, error) {
	a := &Agent{}
	err := a.OnInit(gate, coon, a)
	return a, err
}

type Agent struct {
	gate.AgentBase
	gate               *Gate
	UserId             uint32
	hearttimer         *timewheel.Task
	heartbeatfrequency uint8
}

func (this *Agent) GetSessionData() map[string]interface{} {
	return map[string]interface{}{
		"SessionId":    this.Id(),
		"IP":           this.Conn.RemoteAddr().String(),
		"GateServerId": this.gate.service.GetId(),
		"UserId":       this.UserId,
	}
}

func (this *Agent) OnInit(module gate.IGateModule, coon gate.IConn, agent gate.IAgent) (err error) {
	this.AgentBase.OnInit(module, coon, agent)
	this.gate = module.(*Gate)
	this.UserId = 0
	this.heartbeatfrequency = 0
	this.StartHeartListen()
	return
}

func (this *Agent) Destory() {
	timewheel.Remove(this.hearttimer)
	this.AgentBase.Destory()
}

//启动心跳监听
func (this *Agent) StartHeartListen() {
	this.hearttimer = timewheel.AddCron(time.Second*time.Duration(this.gate.options.HeartbeatInterval), func(task *timewheel.Task, i ...interface{}) {
		if this.heartbeatfrequency >= this.gate.options.MaxHeartStopNum { //超过三次 死
			log.Infof("远程连接心跳异常 断开用户【%d】连接【 %s】", this.UserId, this.Id())
			this.OnClose()
		} else {
			if !this.Isclose {
				this.heartbeatfrequency++
			}
		}
	})
}

func (this *Agent) Build(uId uint32) {
	this.UserId = uId
}

func (this *Agent) UnBuild() {
	this.UserId = 0
}

//发送代理消息
func (this *Agent) WriteMsg(msg proto.IMessage) error {
	return this.AgentBase.WriteMsg(msg)
}

//接收到消息 需要重构
func (this *Agent) OnRecover(msg proto.IMessage) {
	this.AgentBase.OnRecover(msg)
	switch msg.GetComId() {
	case SystemComId:
		switch msg.GetMsgId() {
		case HeartbeatReq: //心跳请求
			this.heartbeatfrequency = 0
			msg := proto.EncodeToMesage(SystemComId, HeartbeatResp, &pb.HeartbeatResp{})
			this.WriteMsg(msg)
			break
		}
		break
	default:
		this.Module.OnRoute(this, msg)
	}
}
