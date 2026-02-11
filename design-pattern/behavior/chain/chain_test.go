package chain

import (
	"fmt"
	"testing"
	"time"
)

func TestChain(t *testing.T) {
	// 构建责任链：Debug → Info → Warn → Error
	debugHandler := &DebugHandler{}
	infoHandler := &InfoHandler{}
	warnHandler := &WarnHandler{}
	errorHandler := &ErrorHandler{}
	debugHandler.SetNext(infoHandler).SetNext(warnHandler).SetNext(errorHandler)

	// 生成日志时间（格式化）
	now := time.Now().Format("2006-01-02 15:04:05")

	// 测试场景1：Debug 级日志（仅 DebugHandler 处理）
	fmt.Println("=== 测试场景1：Debug 级日志 ===")
	log1 := &LogRequest{
		Level:   "Debug",
		Message: "用户请求参数：{id: 1001}",
		Time:    now,
	}
	debugHandler.Handle(log1)

	// 测试场景2：Info 级日志（InfoHandler 处理后传递给 Warn/Error，无额外处理）
	fmt.Println("\n=== 测试场景2：Info 级日志 ===")
	log2 := &LogRequest{
		Level:   "Info",
		Message: "用户 1001 登录成功",
		Time:    now,
	}
	debugHandler.Handle(log2)

	// 测试场景3：Warn 级日志（WarnHandler 处理后传递给 Error，无额外处理）
	fmt.Println("\n=== 测试场景3：Warn 级日志 ===")
	log3 := &LogRequest{
		Level:   "Warn",
		Message: "用户 1001 密码错误 3 次，临时锁定 5 分钟",
		Time:    now,
	}
	debugHandler.Handle(log3)

	// 测试场景4：Error 级日志（最终由 ErrorHandler 处理）
	fmt.Println("\n=== 测试场景4：Error 级日志 ===")
	log4 := &LogRequest{
		Level:   "Error",
		Message: "数据库连接失败：timeout",
		Time:    now,
	}
	debugHandler.Handle(log4)
}
