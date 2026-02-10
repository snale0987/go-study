package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLogEntry_Reset(t *testing.T) {
	logPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("📌 池中空闲对象不足，创建新的LogEntry")
			return &LogEntry{}
		},
	}

	var wg sync.WaitGroup
	// 启动10个协程，每个协程处理3条日志
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(coroutineID int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				// 从池中获取对象（复用或新建）
				entry := logPool.Get().(*LogEntry)
				// 用完后必须归还到池（defer确保执行）
				defer logPool.Put(entry)

				// 使用对象：填充日志内容
				entry.Timestamp = time.Now()
				entry.Level = "INFO"
				entry.Message = fmt.Sprintf("协程%d-日志%d：处理用户请求", coroutineID, j)

				// 模拟日志处理逻辑
				fmt.Printf("[%s] [%s] %s\n", entry.Timestamp.Format("15:04:05"), entry.Level, entry.Message)

				// 复用前重置对象（关键！避免下次使用时有残留数据）
				entry.Reset()

				// 模拟业务耗时
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("\n✅ 所有日志处理完成，对象池中的空闲对象会在GC时自动回收")
}
