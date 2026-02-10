package pool

import "time"

type LogEntry struct {
	Timestamp time.Time // 日志时间
	Level     string    // 日志级别：INFO/ERROR/WARN
	Message   string    // 日志内容
}

func (l *LogEntry) Reset() {
	l.Timestamp = time.Time{} // 重置为零值
	l.Level = ""
	l.Message = ""
}
