package bridge

import "fmt"

// SendImplementor 发送方式接口（实现层）：抽象发送行为
type SendImplementor interface {
	Send(message string) // 核心发送方法
}

// 短信发送（具体实现1）
type SmsSender struct{}

func (s *SmsSender) Send(message string) {
	fmt.Printf("【短信发送】%s\n", message)
}

// 邮件发送（具体实现2）
type EmailSender struct{}

func (e *EmailSender) Send(message string) {
	fmt.Printf("【邮件发送】%s\n", message)
}

type MessageAbstraction struct {
	sender SendImplementor // 桥接核心：组合实现层接口
}

// 设置发送方式（动态切换实现）
func (m *MessageAbstraction) SetSender(sender SendImplementor) {
	m.sender = sender
}

// 发送消息（抽象行为）
func (m *MessageAbstraction) SendMessage(message string) {
	if m.sender == nil {
		fmt.Println("请先设置发送方式")
		return
	}
	m.sender.Send(message)
}

// 普通消息（抽象层的扩展1）
type NormalMessage struct {
	MessageAbstraction // 嵌入抽象类（Go 模拟继承）
}

// 紧急消息（抽象层的扩展2）
type UrgentMessage struct {
	MessageAbstraction // 嵌入抽象类
}

// 重写紧急消息的发送逻辑（增强行为）
func (u *UrgentMessage) SendMessage(message string) {
	// 紧急消息前置处理：添加【紧急】标签
	urgentMsg := fmt.Sprintf("[紧急] %s", message)
	u.sender.Send(urgentMsg)
}
