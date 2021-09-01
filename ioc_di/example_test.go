package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var channelSet = []IChannel{
		&Email{},
		&BusinessWechat{},
	}

	msg := "错误告警"

	for _, c := range channelSet {
		NewAlarm2(c).SendMsg(msg)
	}
}

/**1**/
type Alarm1 struct {}

func (*Alarm1) collectErrorLog() {
	fmt.Println("收集错误日志")
}

func (*Alarm1) SendMsg(channel, content string) {
	if channel == "email" {
		fmt.Println(fmt.Sprintf("通过email发送告警信息:%s", content))
	} else if channel == "business wechat" {
		fmt.Println(fmt.Sprintf("通过企业微信发送告警信息:%s", content))
	}
}

/**1**/
type IChannel interface {
	SendMsg(content string)
}

type Email struct {}
func (*Email) SendMsg(content string) {
	fmt.Println(fmt.Sprintf("email发送信息:%s", content))
}

type BusinessWechat struct {}
func (*BusinessWechat) SendMsg(content string) {
	fmt.Println(fmt.Sprintf("企业微信发送信息:%s", content))
}

type Alarm2 struct {
	channel IChannel
}

/**2**/
func NewAlarm2(channel IChannel) *Alarm2 {
	return &Alarm2{channel: channel}
}

func (*Alarm2) collectErrorLog() {
	fmt.Println("收集错误日志")
}

func (a *Alarm2) SendMsg(content string) {
	a.channel.SendMsg(content)
}

/**3**/