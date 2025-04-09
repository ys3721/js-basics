package main

import "fmt"

type loggable func(message string)

type Logger struct {
	loggable
}

func ConsoleLogger(message string) {
	fmt.Println("Console Log:", message)
}

func FileLogger(message string) {
	fmt.Println("File Log:", message)
}

func (l *Logger) LogMessage(message string) {
	l.loggable(message)
}

func (l loggable) LogMessageFromEmbedding(message string) {
	l(message)
}

func main0() {
	// 6. 使用 ConsoleLog 函数作为 Logger 的日志方式
	logger1 := &Logger{
		loggable: ConsoleLogger,
	}
	logger1.LogMessage("This is a message for console log!")

	// 7. 使用 FileLog 函数作为 Logger 的日志方式
	logger2 := &Logger{
		loggable: FileLogger,
	}
	logger2.LogMessage("This is a message for file log!")
	logger2.LogMessageFromEmbedding("This is a message for file log!")
}
