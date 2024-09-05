package zaplog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var gwriter mywriter

type mywriter struct {
	mutex sync.Mutex
}

func (w *mywriter) Write(p []byte, flag int) (n int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if flag == 1 {
		//info
		return os.Stdout.Write(p)
	}
	if flag == 2 {
		//error
		return os.Stderr.Write(p)
	}
	return len(p), nil
}

type proxywriter int

func (pw proxywriter) Write(p []byte) (n int, err error) {
	return gwriter.Write(p, int(pw))
}

var errorLog = log.New(proxywriter(2), "", log.Llongfile|log.LstdFlags)
var infoLog = log.New(proxywriter(1), "", log.Llongfile|log.LstdFlags)

func init() {
	log.SetOutput(io.Discard)
}

func Errorf(format string, v ...any) {
	strOutput := fmt.Sprintf(format, v...)
	errorLog.Output(2, strOutput)
}

func Error(v ...any) {
	strOutput := fmt.Sprintln(v...)
	errorLog.Output(2, strOutput)
}

func Infof(format string, v ...any) {
	strOutput := fmt.Sprintf(format, v...)
	infoLog.Output(2, strOutput)
}

func Info(v ...any) {
	strOutput := fmt.Sprintln(v...)
	infoLog.Output(2, strOutput)
}

func Fatal(v ...any) {
	strOutput := fmt.Sprintln(v...)
	errorLog.Output(2, strOutput)
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	strOutput := fmt.Sprintf(format, v...)
	errorLog.Output(2, strOutput)
	os.Exit(1)
}
