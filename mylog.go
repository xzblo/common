package common

import mylogger "github.com/xzblo/common/log/mylog"

func NewLogFile(pattern, filepath, prefix string, logfilesize int64)*mylogger.FileLogger{
	//return  mylogger.NewFileLogger("debug", "./", "test", 5*1024*1024)
	return  mylogger.NewFileLogger(pattern, filepath, prefix, logfilesize)
}
