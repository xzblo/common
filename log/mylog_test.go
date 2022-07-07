package mylog

import (
	 mylogger "mylog/log/mylog"
	"testing"
)


func TestNewFileLogger(t *testing.T){
	fileLog := mylogger.NewFileLogger("debug", "./", "test", 2*1024*1024) // 向文件打印
	consoleLog := mylogger.NewConsoleLogger("FATAL")

	for {
		//fileLog.Debug("Debug%v", "试试")
		//fileLog.Info("Info")
		fileLog.Warning("Warning")
		//fileLog.Error("Error")
		//fileLog.Fatal("Fatal")
		consoleLog.Debug("Debug")

	}
}

