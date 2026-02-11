package bridge

import "testing"

func Test_Bridge(t *testing.T) {
	// 1. 创建实现层实例（发送方式）
	smsSender := &SmsSender{}
	emailSender := &EmailSender{}

	// 2. 创建抽象层实例（消息类型）
	normalMsg := &NormalMessage{}
	urgentMsg := &UrgentMessage{}

	// 3. 桥接组合：动态绑定发送方式
	normalMsg.SetSender(smsSender)
	normalMsg.SendMessage("您的订单已发货") // 输出：【短信发送】您的订单已发货

	normalMsg.SetSender(emailSender)
	normalMsg.SendMessage("您的订单已发货") // 输出：【邮件发送】您的订单已发货

	urgentMsg.SetSender(smsSender)
	urgentMsg.SendMessage("系统异常，请及时处理") // 输出：【短信发送】[紧急] 系统异常，请及时处理
}
