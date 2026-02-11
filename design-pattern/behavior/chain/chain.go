package chain

import (
	"fmt"
)

// 1. 定义日志请求结构体：封装日志核心信息
type LogRequest struct {
	Level   string // 日志级别：Debug/Info/Warn/Error
	Message string // 日志内容
	Time    string // 日志时间
}

// 2. 定义日志处理器接口（责任链核心）
type LogHandler interface {
	// 处理日志：返回 bool 表示是否终止链
	Handle(log *LogRequest) bool
	// 设置下一个处理器
	SetNext(handler LogHandler) LogHandler
}

// 3. 具体处理器1：Debug 日志处理器（只处理 Debug 级）
type DebugHandler struct {
	next LogHandler
}

func (d *DebugHandler) Handle(log *LogRequest) bool {
	if log.Level == "Debug" {
		// 格式化输出 Debug 日志
		fmt.Printf("[%s] [DEBUG] %s\n", log.Time, log.Message)
		return true // 只处理 Debug，终止链
	}
	// 非 Debug 级，传递给下一级
	if d.next != nil {
		return d.next.Handle(log)
	}
	return true
}

func (d *DebugHandler) SetNext(handler LogHandler) LogHandler {
	d.next = handler
	return handler
}

// 4. 具体处理器2：Info 日志处理器（处理 Info 级，传递给下一级）
type InfoHandler struct {
	next LogHandler
}

func (i *InfoHandler) Handle(log *LogRequest) bool {
	if log.Level == "Info" {
		fmt.Printf("[%s] [INFO] %s\n", log.Time, log.Message)
		// Info 级处理后，继续传递给 Warn 处理器
		if i.next != nil {
			return i.next.Handle(log)
		}
		return true
	}
	// 非 Info 级，传递给下一级
	if i.next != nil {
		return i.next.Handle(log)
	}
	return true
}

func (i *InfoHandler) SetNext(handler LogHandler) LogHandler {
	i.next = handler
	return handler
}

// 5. 具体处理器3：Warn 日志处理器（处理 Warn 级，传递给下一级）
type WarnHandler struct {
	next LogHandler
}

func (w *WarnHandler) Handle(log *LogRequest) bool {
	if log.Level == "Warn" {
		fmt.Printf("[%s] [WARN] %s\n", log.Time, log.Message)
		// Warn 级处理后，继续传递给 Error 处理器
		if w.next != nil {
			return w.next.Handle(log)
		}
		return true
	}
	// 非 Warn 级，传递给下一级
	if w.next != nil {
		return w.next.Handle(log)
	}
	return true
}

func (w *WarnHandler) SetNext(handler LogHandler) LogHandler {
	w.next = handler
	return handler
}

// 6. 具体处理器4：Error 日志处理器（处理 Error 级，终止链）
type ErrorHandler struct {
	next LogHandler // 无下一级，仅占位
}

func (e *ErrorHandler) Handle(log *LogRequest) bool {
	if log.Level == "Error" {
		// Error 级日志特殊处理：标红（终端转义符）+ 记录到文件（模拟）
		fmt.Printf("\033[31m[%s] [ERROR] %s\033[0m\n", log.Time, log.Message)
		fmt.Println("【附加操作】Error 日志已写入错误日志文件")
		return true // 最高级，终止链
	}
	return true
}

func (e *ErrorHandler) SetNext(handler LogHandler) LogHandler {
	e.next = handler
	return handler
}
