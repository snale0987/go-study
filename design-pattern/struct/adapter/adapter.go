package adapter

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount float64) error
}

type Alipay struct {
	AppID string
}

func (a *Alipay) AliPayMoney(money float64) (bool, string) {
	if money <= 0 {
		return false, "支付金额必须大于0"
	}
	fmt.Printf("[支付宝] 应用ID：%s，支付金额：%.2f 元，支付成功\n", a.AppID, money)
	return true, ""
}

type AlipayAdapter struct {
	AliPay *Alipay
}

// Pay 实现Payment接口的Pay方法，内部调用Alipay的原生方法
func (a *AlipayAdapter) Pay(amount float64) error {
	// 调用支付宝原生方法，并转换返回值为目标接口的格式
	success, msg := a.AliPay.AliPayMoney(amount)
	if !success {
		return errors.New(msg)
	}
	return nil
}
