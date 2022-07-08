package common

import (
	"testing"
)


func TestNewLogFile(t *testing.T){

	logfileobj := NewLogFile("info", "./", "test", 5*1024*1024)
	logfileobj.Info("Info testing")
	//for {
	//	fileLog.Warning("Warning")
	//	//fileLog.Error("Error")
	//	//fileLog.Fatal("Fatal")
	//	consoleLog.Debug("Debug")
	//
	//}
}