package singleton

import (
	"os"
	"sync"
)

// 处理资源访问冲突

type Logger struct {
	file *os.File
	mu   sync.Mutex
}

var (
	loggerInstance *Logger
	once           sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		loggerInstance = &Logger{file: file}
	})
	return loggerInstance
}

func (l *Logger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.file.WriteString(message + "\n")
}
