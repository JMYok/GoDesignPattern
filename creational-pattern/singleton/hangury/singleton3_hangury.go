package hangury

import (
	"os"
	"sync"
)

// 饿汉式单例
type Logger struct {
	file *os.File
	mu   sync.Mutex
}

var loggerInstance *Logger

func init() {
	file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	loggerInstance = &Logger{file: file}
}

func GetLogger() *Logger {
	return loggerInstance
}

func (l *Logger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.file.WriteString(message + "\n")
}
