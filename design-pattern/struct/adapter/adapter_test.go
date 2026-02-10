package adapter

import (
	"fmt"
	"testing"
)

func TestAlipay_AliPayMoney(t *testing.T) {
	// 1. 创建第三方支付宝实例
	aliPay := &Alipay{AppID: "202602090000001"}

	// 2. 创建适配器，将支付宝实例适配为Payment接口
	paymentAdapter := &AlipayAdapter{AliPay: aliPay}

	// 3. 系统统一调用Payment接口，无需关心底层具体实现
	err := paymentAdapter.Pay(100.50)
	if err != nil {
		fmt.Println("支付失败：", err)
	} else {
		fmt.Println("系统层面：支付流程完成")
	}

	// 测试异常场景
	err = paymentAdapter.Pay(0)
	if err != nil {
		fmt.Println("支付失败：", err)
	}
}
