package gate

import (
	"encoding/json"
	"fmt"

	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/modules/gate"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/proto"
)

type SessionData struct {
	IP           string
	SessionId    string
	GateServerId string //用户所在网关服务
	UserId       uint32
}

func NewRemoteSession(service base.IClusterService, data map[string]interface{}) (s core.IUserSession, err error) {
	session := &RemoteSession{
		service: service,
		data:    &SessionData{},
	}
	err = session.UpDateMap(data)
	if err == nil {
		s = session
	}
	return s, err
}

func NewRemoteSessionByByte(service base.IClusterService, b []byte) (s *RemoteSession, err error) {
	session := &RemoteSession{
		service: service,
		data:    &SessionData{},
	}
	err = json.Unmarshal(b, session.data)
	if err == nil {
		s = session
	}
	return s, err
}

type RemoteSession struct {
	service base.IClusterService
	data    *SessionData
}

func (this *RemoteSession) Serializable() ([]byte, error) {
	return json.Marshal(this.data)
}

func (this *RemoteSession) UpDateMap(data map[string]interface{}) (err error) {
	IP := data["IP"]
	if IP != nil {
		this.data.IP = IP.(string)
	} else {
		return fmt.Errorf("用户会话 关键数据 IP 不存在")
	}
	Sessionid := data["SessionId"]
	if Sessionid != nil {
		this.data.SessionId = Sessionid.(string)
	} else {
		return fmt.Errorf("用户会话 关键数据 SessionId 不存在")
	}
	Serverid := data["GateServerId"]
	if Serverid != nil {
		this.data.GateServerId = Serverid.(string)
	} else {
		return fmt.Errorf("用户会话 关键数据 GateServerId 不存在")
	}
	UserId := data["UserId"]
	if UserId != nil {
		this.data.UserId = UserId.(uint32)
	} else {
		this.data.UserId = 0
	}
	return nil
}
func (this *RemoteSession) GetGateServerId() string {
	return this.data.GateServerId
}
func (this *RemoteSession) GetSessionId() string {
	return this.data.SessionId
}
func (this *RemoteSession) GetIP() string {
	return this.data.IP
}
func (this *RemoteSession) GetGateId() string {
	return this.data.GateServerId
}
func (this *RemoteSession) GetUserId() uint32 {
	return this.data.UserId
}

func (this *RemoteSession) Build(userId uint32) (err error) {
	_, err = this.service.RpcInvokeById(this.data.GateServerId, gate.RPC_GateAgentBuild, false, this.data.SessionId, userId)
	if err == nil {
		this.data.UserId = userId
	}
	return
}
func (this *RemoteSession) UnBuild() (err error) {
	_, err = this.service.RpcInvokeById(this.data.GateServerId, gate.RPC_GateAgentUnBuild, false, this.data.SessionId)
	if err != nil {
		this.data.UserId = 0
	}
	return
}
func (this *RemoteSession) SendMsg(comdId uint16, msgId uint16, msg interface{}) (err error) {
	m := proto.EncodeToMesage(comdId, msgId, msg)
	if comdId != 0 {
		log.Infof("向用户【%d】发送【%d:%d】的消息:%s", this.GetUserId(), comdId, msgId, msg)
	}
	_, err = this.service.RpcInvokeById(this.data.GateServerId, gate.RPC_GateSendMsg, false, this.GetSessionId(), m)
	return err
}

func (this *RemoteSession) Close() (err error) {
	_, err = this.service.RpcInvokeById(this.data.GateServerId, gate.RPC_GateAgentClose, false, this.GetSessionId())
	return
}

func NewLocalSession(module gate.IGateModule, data map[string]interface{}) (s core.IUserSession, err error) {
	session := &LocalSession{
		module: module.(*Gate),
		data:   &SessionData{},
	}
	err = session.UpDateMap(data)
	if err == nil {
		s = session
	}
	return s, err
}

type LocalSession struct {
	module *Gate
	data   *SessionData
}

func (this *LocalSession) UpDateMap(data map[string]interface{}) (err error) {
	IP := data["IP"]
	if IP != nil {
		this.data.IP = IP.(string)
	} else {
		return fmt.Errorf("用户会话 关键数据 IP 不存在")
	}
	Sessionid := data["SessionId"]
	if Sessionid != nil {
		this.data.SessionId = Sessionid.(string)
	} else {
		return fmt.Errorf("用户会话 关键数据 SessionId 不存在")
	}
	Serverid := data["GateServerId"]
	if Serverid != nil {
		this.data.GateServerId = Serverid.(string)
	} else {
		return fmt.Errorf("用户会话 关键数据 GateServerId 不存在")
	}
	UserId := data["UserId"]
	if UserId != nil {
		this.data.UserId = UserId.(uint32)
	} else {
		this.data.UserId = 0
	}
	return nil
}
func (this *LocalSession) GetGateServerId() string {
	return this.data.GateServerId
}
func (this *LocalSession) GetSessionId() string {
	return this.data.SessionId
}
func (this *LocalSession) GetIP() string {
	return this.data.IP
}
func (this *LocalSession) GetUserId() uint32 {
	return this.data.UserId
}
func (this *LocalSession) GetGateId() string {
	return this.data.GateServerId
}
func (this *LocalSession) Build(userId uint32) (err error) {
	_, e := this.module.Build(this.GetSessionId(), userId)
	if e != "" {
		return fmt.Errorf(e)
	}
	this.data.UserId = userId
	return
}
func (this *LocalSession) UnBuild() (err error) {
	_, e := this.module.UnBuild(this.GetSessionId())
	if e != "" {
		return fmt.Errorf(e)
	}
	return err
}
func (this *LocalSession) SendMsg(comdId uint16, msgId uint16, msg interface{}) (err error) {
	m := proto.EncodeToMesage(comdId, msgId, msg)
	if comdId != 0 {
		log.Infof("向用户【%d】发送【%d:%d】的消息:%s", this.GetUserId(), comdId, msgId, msg)
	}
	if _, e := this.module.SendMsg(this.GetSessionId(), m); e != "" {
		err = fmt.Errorf(e)
	}
	return
}
func (this *LocalSession) Close() (err error) {
	if _, e := this.module.CloseAgent(this.GetSessionId()); e != "" {
		err = fmt.Errorf(e)
	}
	return
}
